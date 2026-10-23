package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 bool }
func genFn0(p0 uint16) uint64 {
    if (!(((1.5 - 0.1) < ((1e10 * 0.0) + (1.5 / 1.0))) && (255 > 255))) {
        for fi4 := 0; fi4 < 6; fi4++ {
            v5 := GenS1{F0: 42, F1: 1, F2: 524}
            var v6 int = (-(fi4))
            p0 = p0
            _ = fi4
            _ = v5
            _ = v6
        }
    } else {
        p0 = (p0 & (p0 | (^(p0))))
        var v7 int = 32768
        var v8 float64 = 0.1
        _ = v7
        _ = v8
    }
    switch "a" {
    case "":
        p0 = p0
    }
    _ = p0
    return 2
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1e10
    var v3 float32 = v2
    v9 := GenS1{F0: -9223372036854775808, F1: 2, F2: 0}
    fmt.Println(v3)
    fmt.Println(v3)
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 uint64 = 0
            var v13 int64 = -3
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v3 += v2
        if (false || true) {
            v2 -= 1.5
        }
        v1[42] = fi14
        _ = fi14
    }
    if (!(!false)) {
        for fi15 := 0; fi15 < 10; fi15++ {
            fmt.Println(v2, v2, v3)
            for fi16 := 0; fi16 < 12; fi16++ {
                _ = len(v1)
                var v17 float32 = v2
                switch "\n\t" {
                case "":
                    fmt.Println(fi16, v2, v17)
                case "foo bar":
                    var v18 uint32 = 4294967295
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    } else {
        if (1 > 93) {
            if false {
                v19 := GenS1{F0: 9223372036854775807, F1: 42, F2: 597}
                var v20 float64 = 1.5
                v21 := GenS1{F0: -9223372036854775808, F1: 65535, F2: 0}
                _ = v19
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 int64 = 42
                    _ = fi22
                    _ = v23
                }
                v24 := GenS1{F0: 1, F1: 32767, F2: genFn0(255)}
                _ = v24
            }
            v25 := GenS2{F0: 218, F1: 9223372036854775807, F2: 2147483647}
            fmt.Println(v3)
            _ = v25
        } else {
            v26 := []int{32768, 0}
            v27 := GenS3{F0: 144}
            var v28 uint32 = 673
            _ = v26
            _ = v27
            _ = v28
        }
    }
    fmt.Println(v3, v2)
    v3 -= 1e10
    if true {
        v29 := GenS4{F0: 0, F1: (128 != 42)}
        _ = v29
    } else {
        for fi30 := 0; fi30 < 10; fi30++ {
            for fi31 := 0; fi31 < 14; fi31++ {
                var v32 bool = (0.0 < 0.1)
                v33 := GenS2{F0: 118, F1: (fi31 ^ -2147483648), F2: 9223372036854775807}
                var v34 int = (fi30 * fi30)
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            var v35 float64 = float64(v2)
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 int32 = 332
                var v38 int16 = -1
                v38 -= (v38 | v38)
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = fi30
            _ = v35
        }
        for fi39 := 0; fi39 < 6; fi39++ {
            fmt.Println(v2, v2)
            if (-609 < -576) {
                v2 = (float32((float64(v3) / ((0.0 / (0.1 + 1.5)) - 1.5))) - ((v3 + (v2 / float32(-2.25))) / (float32(-2.25) - 1e10)))
            }
            for fi40 := 0; fi40 < 8; fi40++ {
                v41 := []int{127, (fi40 / (((fi39 & -688)) | 1))}
                _ = fi40
                _ = v41
            }
            _ = fi39
        }
    }
    v3 -= v2
    v2 = float32(1.5)
    if true {
        v42 := GenS2{F0: 2, F1: 32768, F2: 4294967295}
        var v43 bool = true
        _ = v42
        _ = v43
    }
    v44 := GenS3{F0: 42}
    v45 := v1[0]
    v2 = v3
    fmt.Println(v2, v2, v3)
    v3 += 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v9
}
