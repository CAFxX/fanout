/* r23 1-round SPN (S1 KILL canary): same construction as
 * r23_spn_mix4r_pba19b01_nw but a single round. Expected to be flagged by
 * the S1 DIFF family (insufficient diffusion).
 * Bit-exact vs battery/canaries kill_s1 model.py (proven, see PROOF.md). */
#include <stdint.h>

static const uint8_t SBOXT[16] = {0,1,2,4,3,8,15,12,9,5,11,6,7,14,13,10};

static uint64_t rotl64(uint64_t x, unsigned r) {
    r &= 63u;
    return r ? (x << r) | (x >> (64u - r)) : x;
}
static uint64_t rk64(uint64_t k, int i) {
    static const uint64_t RC[4] = {0x9E3779B97F4A7C15ULL, 0x3C6EF372FE94F82AULL, 0xDAA66D2C7DDF743FULL, 0x78DDE6E5FD29F054ULL};
    unsigned r = (13u * (unsigned)(i + 1)) & 63u;
    return rotl64(k, r) ^ RC[i];
}
static uint64_t sbox64(uint64_t x) {
    uint64_t r = 0;
    for (int i = 0; i < 16; i++)
        r |= (uint64_t)SBOXT[(x >> (4*i)) & 0xFu] << (4*i);
    return r;
}
/* P(i) = (19*i+1) mod 64, generated */
static uint64_t pbox64(uint64_t x) {
    uint64_t y = 0;
    for (int i = 0; i < 64; i++) {
        int p = (19 * i + 1) & 63;
        y |= ((x >> i) & 1u) << p;
    }
    return y;
}
static uint64_t lin_spec(uint64_t x) {
    return x ^ rotl64(x, 3) ^ rotl64(x, 11);
}
uint64_t mix_hash(uint64_t val, uint64_t key) {
    uint64_t t = sbox64(val ^ rk64(key, 0));
    return lin_spec(pbox64(t));
}
