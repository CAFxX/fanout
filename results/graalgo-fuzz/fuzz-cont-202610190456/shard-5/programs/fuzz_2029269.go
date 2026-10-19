package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 uint16, p1 int16) uint16 {
    var v7 int16 = p1
    p0 -= (p0 - p0)
    _ = p0
    _ = p1
    _ = v7
    return p0
}

func main() {
    var v1 float32 = float32((1e10 - 1e10))
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 0
    v5 := make(map[string]int)
    var v6 uint = 0
    if (true && (-32768 >= int16(v4))) {
        _ = len(v5)
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            if true {
                v6 = v6
                v4 = (v4 >> (((v6 ^ (v6 | v6))) & 63))
                var v9 uint16 = 42
                _ = v9
            } else {
                var v10 bool = (651 >= -353)
                var v11 uint64 = 4294967295
                v4 += (v4 % (((v4 + v4)) | 1))
                _ = v10
                _ = v11
            }
            var v12 string = (("hello" + "hello") + ("foo bar" + (("foo bar" + "\n\t") + "hello")))
            v13 := GenS1{F0: 42, F1: (v4 & 157)}
            _ = fi8
            _ = v12
            _ = v13
        }
        v1 = float32((((float64((v1 * 1.5)) / 3.14159) * (float64(0.1) * ((((float64(((float32(3.14159) * 0.1) * v1)) + 0.1) + 1e10) - (3.14159 - 1.5)) / 0.1))) + ((0.1 / 1.5) - 0.1)))
    }
    var v14 uint32 = 1
    var v15 int = 429
    switch (v15 & (v15 & v15)) {
    case -2147483648:
        if (!true) {
            v16 := GenS1{F0: 35, F1: 42}
            v15 += (-(v15))
            v4 -= v4
            _ = v16
        }
    case 0:
        v17 := GenS1{F0: uint8(v14), F1: v4}
        _ = v17
    }
    v18 := GenS1{F0: v4, F1: (v4 % (((^(v4))) | 1))}
    v19 := GenS2{F0: genFn0(genFn0(1, -330), -32768)}
    for fi20 := 0; fi20 < 19; fi20++ {
        switch 42 {
        case 2147483647:
            var v21 uint64 = 42
            _ = v21
        case -2147483648:
            if true {
                for fi22 := 0; fi22 < 3; fi22++ {
                    v23 := GenS3{F0: -32768, F1: 1}
                    _ = fi22
                    _ = v23
                }
                v4 = v4
                fmt.Println(v6, fi20, v1)
            }
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            v14 += (v14 ^ v14)
            v25 := []int{(v15 * fi20), v15, (v15 >> ((uint(v14)) & 63)), 59}
            switch (fi20 - (v15 * fi24)) {
            case 42:
                if true {
                    v15 = (v15 / (((v15 & (fi20 << ((273) & 63)))) | 1))
                }
            case 198:
                if false {
                    v14 = v14
                    v4 -= v4
                } else {
                    v26 := GenS4{F0: (v4 ^ 0), F1: 18446744073709551615, F2: 897}
                    if (0.0 == 0.0) {
                        switch -2147483648 {
                        case 32768:
                            var v27 int32 = 0
                            _ = v27
                        }
                    }
                    _ = v26
                }
            case -9223372036854775808:
                v25 = append(v25, (fi24 ^ v15))
            default:
                v14 += v14
            }
            if true {
                switch "\n\t" {
                case "\n\t":
                    v28 := GenS5{F0: 0, F1: 4294967295, F2: "a"}
                    _ = v28
                default:
                    var v29 uint32 = v14
                    _ = v29
                }
            }
            _ = fi24
            _ = v25
        }
        _ = fi20
    }
    v30 := GenS1{F0: (v4 << ((65535) & 63)), F1: (^(v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
