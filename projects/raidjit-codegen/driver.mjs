// GHA shard driver for raidjit codegen review campaign.
// Shard contract: $SHARD_IDX / $SHARD_COUNT select configs; outputs to $OUT_DIR.
// Layout: <tenant>/driver.mjs, <tenant>/wasm/*.js, <tenant>/configs.jsonl
// Run: node driver.mjs
import { readFileSync, writeFileSync, mkdirSync } from 'fs';
import { createRequire } from 'module';
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
let dmod = null;
async function getDisasm() {
  if (!dmod) {
    const m = req(join(WASM, 'disasm.js'));
    dmod = await (m.default || m)();
  }
  return dmod;
}

function compileKernel(mod, doc) {
  const json = JSON.stringify(doc);
  const p = mod.stringToNewUTF8(json);
  let rp;
  try { rp = mod._rjw_compile(p); }
  catch (e) { mod._free(p); return { ok: false, err: 'trap: ' + e.message }; }
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

function disasm(arch, bytes) {
  const p = dmod._malloc(bytes.length || 1);
  if (bytes.length) dmod.HEAPU8.set(bytes, p);
  const ap = dmod.stringToNewUTF8(arch);
  let rp;
  try { rp = dmod._dj_disasm(ap, p, bytes.length, 0n); }
  catch (e) { dmod._free(ap); dmod._free(p); return null; }
  const s = dmod.UTF8ToString(rp);
  dmod._free(ap); dmod._free(p);
  return s;
}

function fmtAsm(cfg, res, disJson) {
  const L = [];
  L.push(`; config_t${cfg.tier}_${cfg.idx} target=${cfg.target} twin=${cfg.meta.pinned} topo=${cfg.meta.topo} op=${cfg.meta.op} U=${cfg.meta.U}`);
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

async function main() {
  const lines = readFileSync(join(HERE, 'configs.jsonl'), 'utf8').trim().split('\n');
  const all = lines.map(l => JSON.parse(l));
  const mine = all.filter((_, i) => i % SHARD_COUNT === SHARD_IDX);
  console.log(`shard ${SHARD_IDX}/${SHARD_COUNT}: ${mine.length}/${all.length} configs`);
  await getDisasm();
  let ok = 0, rejected = 0, errors = 0;
  for (const cfg of mine) {
    const name = `config_t${cfg.tier}_${cfg.idx}`;
    try {
      const mod = await getMod(cfg.isa_module);
      const res = compileKernel(mod, cfg.doc);
      const meta = {
        tier: cfg.tier, idx: cfg.idx,
        isa_module: cfg.isa_module, target: cfg.target,
        doc: cfg.doc, meta: cfg.meta,
        result: res.ok
          ? { ok: true, code_size: res.code_size, twins: res.twins, uops_per_byte: res.uops_per_byte,
              selection: res.selection, const_size: res.cnst.length }
          : { ok: false, err: res.err },
      };
      if (res.ok) {
        ok++;
        const d = disasm(ARCH_OF[cfg.isa_module], res.code);
        writeFileSync(join(OUT_DIR, name + '.asm'), fmtAsm(cfg, res, d));
      } else {
        rejected++;
      }
      writeFileSync(join(OUT_DIR, name + '.json'), JSON.stringify(meta));
    } catch (e) {
      errors++;
      writeFileSync(join(OUT_DIR, name + '.json'),
        JSON.stringify({ tier: cfg.tier, idx: cfg.idx, isa_module: cfg.isa_module,
          target: cfg.target, doc: cfg.doc, meta: cfg.meta,
          result: { ok: false, err: 'driver exception: ' + e.message } }));
    }
  }
  const status = { shard: SHARD_IDX, total: mine.length, ok, rejected, errors };
  writeFileSync(join(OUT_DIR, '_shard_status.json'), JSON.stringify(status));
  console.log(JSON.stringify(status));
}
main().catch(e => { console.error('FATAL', e); process.exit(1); });
