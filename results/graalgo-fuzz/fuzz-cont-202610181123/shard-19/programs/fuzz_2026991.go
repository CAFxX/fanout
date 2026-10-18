package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint }
func genFn0(p0 int16) uint64 {
    var v4 float32 = 0.1
    var v5 float64 = 0.0
    v5 -= ((float64(v4) * v5) / (((v5 * ((1e10 - 0.0) * 1.5)) / v5) - (v5 + (v5 / (3.14159 + 1.5)))))
    _ = p0
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 32768
    v3 := make(map[string]uint32)
    v6 := GenS1{F0: 128}
    v7 := GenS1{F0: 32}
    if false {
        var v8 int64 = 456
        var v9 int32 = 2147483647
        var v10 int32 = v9
        _ = v10
        _ = v8
        _ = v9
    }
    fmt.Println(v2)
    if false {
        fmt.Println(v2)
        v11 := v3["\n\t"]
        var v12 int64 = 9223372036854775807
        _ = v11
        _ = v12
    }
    v2 = v2
    v2 -= v2
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint16 = 42
        if (!false) {
            for fi15 := 0; fi15 < 10; fi15++ {
                switch "foo bar" {
                case "a":
                    v2 = -914
                case "\n\t":
                    var v16 int32 = 2147483647
                    var v17 int32 = (v16 & (v16 ^ -438))
                    if false {
                        var v18 int = fi13
                        switch (v18 + (fi15 << ((1) & 63))) {
                        case 542:
                            v19 := GenS2{F0: 68, F1: 2, F2: -43}
                            _ = v19
                        case -976:
                            v20 := GenS1{F0: 1}
                            _ = v20
                        default:
                            v17 -= (v17 + v17)
                        }
                        var v21 uint64 = 564
                        _ = v18
                        _ = v21
                    }
                    _ = v16
                    _ = v17
                case "":
                    v22 := GenS3{F0: ("\n\t" <= "\n\t"), F1: 151, F2: 2}
                    _ = v22
                }
                v23 := GenS2{F0: 0, F1: 152, F2: -128}
                _ = fi15
                _ = v23
            }
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 uint64 = genFn0(889)
                _ = fi24
                _ = v25
            }
            var v26 string = ""
            _ = v26
        }
        _ = fi13
        _ = v14
    }
    v27 := GenS4{F0: 183}
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := GenS4{F0: 623}
        v30 := []int{(-(v2))}
        if (0 > 164) {
            delete(v1, (fi28 - -9223372036854775808))
            switch 32768 {
            case -513:
                v31 := GenS5{F0: 395}
                _ = v31
            case -85:
                v2 = (^(v2))
            case 32767:
                switch -212 {
                case 4294967295:
                    fmt.Println(v2, v2)
                default:
                    var v32 uint64 = 4294967295
                    _ = v32
                }
            }
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    if true {
        var v33 string = ""
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v6
    _ = v7
}
