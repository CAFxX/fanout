package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int16; F1 string; F2 string }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 1
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 690, F1: 127}
        _ = fi4
        _ = v5
    }
    switch 435 {
    case -9223372036854775808:
        switch 4294967295 {
        case 1:
            for fi6 := 0; fi6 < 2; fi6++ {
                for fi7 := 0; fi7 < 17; fi7++ {
                    switch (fi7 + (fi7 >> ((407) & 63))) {
                    case 481:
                        fmt.Println(v3, v3, fi7)
                    case -819:
                        v3 = (v3 ^ (v3 - 0))
                    case 128:
                        v3 = (v3 - v3)
                        v1[("hello" + "a")] = fi7
                    default:
                        v3 = (v3 << ((2) & 63))
                    }
                    var v8 float32 = (3.14159 - (1.5 - (float32(float64((0.0 + ((1e10 * 0.1) * 1e10)))) + 3.14159)))
                    if true {
                        var v9 uint32 = 42
                        v10 := []int{}
                        v11 := GenS1{F0: 2147483647, F1: -128}
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                    _ = fi7
                    _ = v8
                }
                v12 := []int{128, (fi6 * fi6), 42, (fi6 * -9223372036854775808)}
                _ = fi6
                _ = v12
            }
        default:
            fmt.Println(v3)
        }
    case 127:
        fmt.Println(v3)
    default:
        if false {
            v3 = (^(v3))
            var v13 int = 42
            _ = v13
        } else {
            v2[42] = false
            fmt.Println(v3)
        }
    }
    v3 = (v3 & v3)
    if false {
        var v14 float32 = 1.5
        v14 += v14
        v15 := v2[42]
        _ = v14
        _ = v15
    } else {
        for fi16 := 0; fi16 < 6; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                v3 -= v3
                var v18 uint8 = v3
                _ = fi17
                _ = v18
            }
            fmt.Println(v3, fi16)
            v19 := GenS1{F0: 0, F1: 100}
            _ = fi16
            _ = v19
        }
        v20 := v2[65535]
        _ = v20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 uint8 = v3
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 128, F1: 1}
    v24 := GenS2{F0: 167, F1: "", F2: "hello"}
    for fi25 := 0; fi25 < 9; fi25++ {
        if false {
            var v26 uint64 = 4294967295
            _ = v26
        }
        v27 := GenS2{F0: -622, F1: "hello", F2: ("" + "\n\t")}
        var v28 uint64 = uint64(v3)
        _ = fi25
        _ = v27
        _ = v28
    }
    var v29 int16 = -117
    fmt.Println(v29, v3)
    v29 += v29
    v30 := []int{}
    v3 += (v3 | v3)
    var v31 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
