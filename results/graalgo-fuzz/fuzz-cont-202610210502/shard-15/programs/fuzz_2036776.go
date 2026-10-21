package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 string }

func main() {
    var v1 int = 2147483647
    v2 := []string{("foo bar" + "hello"), ("\n\t" + "hello"), "a"}
    v3 := make(map[string]string)
    v4 := []int{}
    var v5 uint8 = 42
    var v6 int32 = 0
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: fi7, F1: 18446744073709551615, F2: (v5 | 219)}
        _ = fi7
        _ = v8
    }
    _ = len(v3)
    var v9 float64 = 1.5
    if true {
        if (((v6 | (v6 << ((4294967295) & 63))) <= (v6 - v6)) && true) {
            _ = len(v3)
            switch (v1 * (v1 >> ((2) & 63))) {
            case -508:
                v10 := GenS1{F0: v1, F1: 692, F2: v5}
                _ = v10
            case -2147483648:
                v6 -= v6
            }
        }
        var v11 uint = 168
        _ = v11
    } else {
        if (507 == 1) {
            var v12 uint = 65535
            if (1.5 <= v9) {
                var v13 uint8 = (v5 | 27)
                v1 = 127
                var v14 bool = true
                _ = v13
                _ = v14
            }
            _ = v12
        }
        var v15 int8 = 1
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := v3["foo bar"]
        if (!true) {
            v6 = v6
            v1 = (v1 % (((^(fi16))) | 1))
            var v18 uint64 = 1
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    var v19 float32 = -2.25
    v20 := GenS2{F0: 760, F1: 128}
    v9 = (0.0 * (3.14159 * -2.25))
    for fi21 := 0; fi21 < 8; fi21++ {
        v19 = (((v19 - 0.0) - v19) * (1e10 + (float32(v9) / 1.5)))
        switch (^(v1)) {
        case -1:
            v9 += v9
        default:
            if false {
                v5 = v5
                for fi22 := 0; fi22 < 2; fi22++ {
                    v23 := GenS3{F0: 518, F1: 0}
                    var v24 int16 = int16(v6)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            } else {
                switch 2147483647 {
                case 128:
                    if false {
                        var v25 float32 = ((float32(v9) + v19) - (float32(v9) + (float32(-2.25) * (v19 / (((1.5 * v19) / 1.0) - 0.1)))))
                        _ = v25
                    }
                default:
                    fmt.Println(v5, v9, v19)
                }
            }
        }
        _ = fi21
    }
    v26 := v3["\n\t"]
    v27 := GenS4{F0: (v26 + v26)}
    v28 := GenS1{F0: (v1 >> ((255) & 63)), F1: 1, F2: (v5 + v5)}
    v9 -= v9
    if false {
        v1 = -377
        _ = len(v3)
    }
    v29 := GenS2{F0: 1, F1: (^(v5))}
    v6 -= (v6 & (-(v6)))
    var v30 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
