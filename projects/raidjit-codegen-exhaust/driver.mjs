// GHA shard driver, raidjit-codegen-exhaust tenant (compact outputs).
// Compile/disasm logic is bit-exact from raidjit-codegen/driver.mjs;
// only the output stage differs: per shard it writes results.jsonl
// (one compact record per config, incl. code_sha256 + asm_sha256) and
// asm_bundle.txt.gz (all disassemblies, header-delimited), instead of
// one .asm/.json file pair per config. Shard contract unchanged:
// $SHARD_IDX/$SHARD_COUNT select configs; outputs under $OUT_DIR.
//
// BUG-COMBO-TRAP-002 fix (2026-10-11): a library abort (assert) kills the
// shared WASM instance; the old driver kept serving the dead instance, so
// every later config in the shard threw 'memory access out of bounds'
// (992 of E1's 1,024 'traps' were this artifact). Now: any throw from a
// compile/disasm call drops the instance, and the SAME config is retried
// once on a fresh instance. Only a throw on the fresh instance is recorded
// as a trap verdict; a successful retry yields the config's true verdict.
import { readFileSync, writeFileSync, mkdirSync, createWriteStream } from 'fs';
import { createRequire } from 'module';
import { createHash } from 'crypto';
import { createGzip } from 'zlib';
import { dirname, join } from 'path';
import { fileURLToPath } from 'url';

const HERE = dirname(fileURLToPath(import.meta.url));
const OUT_DIR = process.env.OUT_DIR || join(HERE, 'out');
const SHARD_IDX = parseInt(process.env.SHARD_IDX || '0', 10);
const SHARD_COUNT = parseInt(process.env.SHARD_COUNT || '1', 10);
mkdirSync(OUT_DIR, { recursive: true });

const req = createRequire(HERE + '/');
const WASM = join(HERE, 'wasm');
const ARCH_OF = { x86: 'x86_64', aarch64: 'aarch64', riscv: 'riscv64' };
const MOD_FILE = { x86: 'raidjit_x86.js', aarch64: 'raidjit_aarch64.js', riscv: 'raidjit_riscv.js' };

const mods = {};
async function getMod(isa_module) {
  if (!mods[isa_module]) {
    const m = req(join(WASM, MOD_FILE[isa_module]));
    mods[isa_module] = await (m.default || m)();
  }
  return mods[isa_module];
}
function dropMod(isa_module) { delete mods[isa_module]; }
let dmod = null;
async function getDisasm() {
  if (!dmod) {
    const m = req(join(WASM, 'disasm.js'));
    dmod = await (m.default || m)();
  }
  return dmod;
}

// One compile attempt on the given instance. Throws propagate to the
// caller (instance may be dead afterwards); cleanup is best-effort.
// A JSON parse failure is a returned verdict, not a throw (as before).
function compileOnce(mod, doc) {
  const json = JSON.stringify(doc);
  const p = mod.stringToNewUTF8(json);
  let rp;
  try { rp = mod._rjw_compile(p); }
  catch (e) { try { mod._free(p); } catch { /* dead instance */ } throw e; }
  let res;
  try { res = JSON.parse(mod.UTF8ToString(rp)); }
  catch (e) { mod._rjw_free_string(rp); mod._free(p); return { ok: false, err: 'bad JSON: ' + e.message }; }
  mod._rjw_free_string(rp); mod._free(p);
  if (!res || !res.ok) return res;
  const lenp = mod._malloc(8);
  const cp = mod._rjw_code_bytes(lenp);
  const clen = mod.getValue(lenp, 'i32');
  const code = Buffer.from(mod.HEAPU8.slice(cp, cp + clen));
  const dp = mod._rjw_const_bytes(lenp);
  const dlen = mod.getValue(lenp, 'i32');
  const cnst = Buffer.from(mod.HEAPU8.slice(dp, dp + dlen));
  mod._free(lenp);
  res.code = code; res.cnst = cnst;
  return res;
}

// Compile with instance-poisoning protection (BUG-COMBO-TRAP-002): on any
// throw, drop the instance and retry the same config once on a fresh one.
async function compileKernel(isa_module, doc) {
  const mod = await getMod(isa_module);
  try { return compileOnce(mod, doc); }
  catch (e1) {
    dropMod(isa_module);
    const fresh = await getMod(isa_module);
    try { return compileOnce(fresh, doc); }
    catch (e2) { dropMod(isa_module); return { ok: false, err: 'trap: ' + e2.message }; }
  }
}

function disasmOnce(arch, bytes) {
  const p = dmod._malloc(bytes.length || 1);
  if (bytes.length) dmod.HEAPU8.set(bytes, p);
  const ap = dmod.stringToNewUTF8(arch);
  let rp;
  try { rp = dmod._dj_disasm(ap, p, bytes.length, 0n); }
  catch (e) { try { dmod._free(ap); dmod._free(p); } catch { /* dead */ } throw e; }
  const s = dmod.UTF8ToString(rp);
  dmod._free(ap); dmod._free(p);
  return s;
}

async function disasm(arch, bytes) {
  await getDisasm();
  try { return disasmOnce(arch, bytes); }
  catch (e1) {
    dmod = null;
    await getDisasm();
    try { return disasmOnce(arch, bytes); }
    catch (e2) { dmod = null; return null; }
  }
}

function fmtAsm(cfg, res, disJson) {
  const L = [];
  L.push(`; config_t${cfg.tier}_${cfg.idx} target=${cfg.target} meta=${JSON.stringify(cfg.meta)}`);
  L.push(`; code_size=${res.code_size} twins=${JSON.stringify(res.twins)} uops_per_byte=${res.uops_per_byte}`);
  let insns;
  try { insns = JSON.parse(disJson); } catch { return L.join('\n') + '\n; DISASM_PARSE_ERROR\n'; }
  for (const i of insns) {
    const addr = String(i.addr).padStart(6, ' ');
    const bytes = (i.bytes || '').padEnd(24, ' ');
    L.push(`${addr}: ${bytes} ${i.mnemonic} ${i.op_str}`);
  }
  return L.join('\n') + '\n';
}

const sha = (buf) => createHash('sha256').update(buf).digest('hex');

async function main() {
  const lines = readFileSync(join(HERE, 'configs.jsonl'), 'utf8').trim().split('\n');
  const all = lines.map(l => JSON.parse(l));
  const mine = all.filter((_, i) => i % SHARD_COUNT === SHARD_IDX);
  console.log(`shard ${SHARD_IDX}/${SHARD_COUNT}: ${mine.length}/${all.length} configs`);
  await getDisasm();
  const gz = createGzip();
  const bundle = gz.pipe(createWriteStream(join(OUT_DIR, 'asm_bundle.txt.gz')));
  const recs = [];
  let ok = 0, rejected = 0, errors = 0;
  for (const cfg of mine) {
    const name = `config_t${cfg.tier}_${cfg.idx}`;
    try {
      const res = await compileKernel(cfg.isa_module, cfg.doc);
      if (res.ok) {
        ok++;
        const d = await disasm(ARCH_OF[cfg.isa_module], res.code);
        const asm = fmtAsm(cfg, res, d);
        gz.write(`===== ${name} =====\n${asm}`);
        recs.push(JSON.stringify({ name, tier: cfg.tier, idx: cfg.idx, isa_module: cfg.isa_module,
          target: cfg.target, doc: cfg.doc, meta: cfg.meta, ok: true,
          code_size: res.code_size, code_sha256: sha(res.code), asm_sha256: sha(Buffer.from(asm)),
          twins: res.twins, uops_per_byte: res.uops_per_byte, selection: res.selection,
          const_size: res.cnst.length }));
      } else {
        rejected++;
        recs.push(JSON.stringify({ name, tier: cfg.tier, idx: cfg.idx, isa_module: cfg.isa_module,
          target: cfg.target, doc: cfg.doc, meta: cfg.meta, ok: false, err: res.err }));
      }
    } catch (e) {
      errors++;
      recs.push(JSON.stringify({ name, tier: cfg.tier, idx: cfg.idx, isa_module: cfg.isa_module,
        target: cfg.target, doc: cfg.doc, meta: cfg.meta, ok: false, err: 'driver exception: ' + e.message }));
    }
  }
  gz.end();
  await new Promise(r => bundle.on('finish', r));
  writeFileSync(join(OUT_DIR, 'results.jsonl'), recs.join('\n') + '\n');
  const status = { shard: SHARD_IDX, total: mine.length, ok, rejected, errors };
  writeFileSync(join(OUT_DIR, '_shard_status.json'), JSON.stringify(status));
  console.log(JSON.stringify(status));
}
main().catch(e => { console.error('FATAL', e); process.exit(1); });
