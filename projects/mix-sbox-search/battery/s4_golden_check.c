/* Minimal golden-vector check for the S4 canary (fast path).
 * Compiled with -DGOLDEN_KEY=... -DGOLDEN_OUT=... alongside the bundle's
 * impl.c. This is the same assertion s4_smallcrush.c performs before
 * launching SmallCrush; the canary runs it standalone so the pass canary
 * completes in seconds instead of 45 minutes. */
#include <stdio.h>
#include <stdint.h>
#include <inttypes.h>
uint64_t mix_hash(uint64_t v, uint64_t k);
#ifndef GOLDEN_KEY
#define GOLDEN_KEY 0x123456789abcdef0ULL
#endif
#ifndef GOLDEN_VAL
#define GOLDEN_VAL 0x0ULL
#endif
#ifndef GOLDEN_OUT
#define GOLDEN_OUT 0x0ULL
#endif
int main(void) {
    uint64_t o = mix_hash(GOLDEN_VAL, GOLDEN_KEY);
    if (o != GOLDEN_OUT) {
        printf("golden MISMATCH got=%016" PRIx64 " want=%016" PRIx64 "\n",
               o, (uint64_t)GOLDEN_OUT);
        return 1;
    }
    printf("golden OK\n");
    return 0;
}
