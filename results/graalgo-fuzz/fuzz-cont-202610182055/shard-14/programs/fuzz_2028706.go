package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 bool; F2 int64 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 bool = (-14 > 42)
    v2 := []int{-697, 0}
    var v3 int64 = -1
    v4 := []uint32{65535, 1, 192}
    v5 := []string{}
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 19; fi7++ {
        if (454 <= 65535) {
            var v8 int64 = v3
            v9 := GenS1{F0: -128}
            _ = v8
            _ = v9
        } else {
            if len(v2) > 0 {
                v2[0] = (fi7 / (((^(fi7))) | 1))
            }
            var v10 int = (fi7 + (fi7 * fi7))
            _ = v10
        }
        _ = fi7
    }
    v3 -= (v3 % ((v3) | 1))
    for fi11 := 0; fi11 < 18; fi11++ {
        if (false && (!((v1 || v1) || true))) {
            v12 := GenS1{F0: 42}
            if (v1 && (!v1)) {
                _ = len(v2)
            } else {
                v13 := GenS1{F0: -128}
                var v14 uint16 = 32767
                switch "\n\t" {
                case "foo bar":
                    var v15 uint8 = 42
                    _ = v15
                case "":
                    fmt.Println(v14, v14)
                default:
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                }
                _ = v13
                _ = v14
            }
            v6 = 0.1
            _ = v12
        }
        _ = fi11
    }
    v3 += 9223372036854775807
    if false {
        if false {
            var v16 int8 = 1
            v17 := GenS2{F0: 1000, F1: v1, F2: v3}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v16
            _ = v17
        }
    }
    if false {
        v18 := GenS3{F0: 1}
        v19 := GenS2{F0: 867, F1: v1, F2: (-(v3))}
        _ = v18
        _ = v19
    }
    v3 += (v3 ^ -398)
    if (((v3 | (v3 << ((65535) & 63))) <= 292) || false) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi20 := 0; fi20 < 17; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                switch "" {
                case "\n\t":
                    v3 -= v3
                default:
                    fmt.Println(v6)
                }
                _ = len(v4)
                var v22 int32 = 1
                v6 += (v6 * v6)
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS1{F0: 0}
            if v1 {
                v6 += v6
                fmt.Println(fi23, v1)
                if (((!false) && (1 >= -1)) || ((v3 * v3) <= (v3 | v3))) {
                    v26 := GenS1{F0: -91}
                    fmt.Println(fi23)
                    _ = v26
                } else {
                    var v27 uint8 = 81
                    _ = v27
                }
                v2 = append(v2, -240)
            }
            var v28 bool = (true && ((790 >= 4294967295) || (42 == -20)))
            if len(v2) > 0 {
                v2[0] = fi24
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi24
            _ = v25
            _ = v28
        }
        v29 := GenS3{F0: 255}
        v1 = (1 <= 42)
        _ = fi23
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
