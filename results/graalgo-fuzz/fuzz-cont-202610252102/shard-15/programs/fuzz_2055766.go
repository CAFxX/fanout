package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int64 }
func genFn0(p0 int) uint {
    var v6 uint32 = 0
    switch p0 {
    case 960:
        var v7 uint = 18446744073709551615
        _ = v7
    case 974:
        for fi8 := 0; fi8 < 12; fi8++ {
            if true {
                v9 := GenS1{F0: -32768, F1: 42, F2: 0}
                _ = v9
            }
            var v10 uint8 = 100
            _ = fi8
            _ = v10
        }
    case 9223372036854775807:
        v11 := []int{4294967295, 127}
        _ = v11
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1() uint {
    v12 := GenS1{F0: 32767, F1: 42, F2: -1}
    var v13 int16 = 470
    v13 = v13
    var v14 uint64 = 2
    _ = v12
    _ = v13
    _ = v14
    return 4
}

func main() {
    v1 := []string{"hello"}
    var v2 uint16 = 42
    v3 := make(map[string]uint32)
    v4 := make(map[int]int)
    var v5 uint16 = (^(v2))
    if true {
        var v15 int = 9223372036854775807
        v2 -= v2
        _ = v15
    }
    v2 = (v2 & (^(v5)))
    v5 = v2
    if (1 < 42) {
        v5 += (v5 << ((807) & 63))
        var v16 float64 = (float64(float32(0.1)) + 1e10)
        _ = v16
    } else {
        _ = len(v1)
    }
    v17 := GenS1{F0: 0, F1: v2, F2: 297}
    var v18 uint64 = 523
    if (!(928 >= -9223372036854775808)) {
        var v19 uint8 = uint8(v2)
        v19 -= v19
        if true {
            var v20 uint8 = (v19 * v19)
            _ = v20
        } else {
            if (((true && true) || true) && ((718 != uint(v19)) || ((((0.1 + (1.5 - 3.14159)) * 1e10) * 3.14159) < 0.0))) {
                for fi21 := 0; fi21 < 18; fi21++ {
                    switch (fi21 | (fi21 << ((2) & 63))) {
                    case 32767:
                        if ((v18 >> ((980) & 63)) > v18) {
                            v22 := GenS1{F0: 1, F1: v5, F2: 1}
                            v19 -= (v19 & v19)
                            switch "hello" {
                            case "foo bar":
                                v18 = v18
                            case "hello":
                                var v23 uint8 = v19
                                _ = v23
                            case "\n\t":
                                var v24 float32 = 1e10
                                _ = v24
                            }
                            _ = v22
                        }
                    case 2147483647:
                        var v25 bool = ("foo bar" <= "hello")
                        _ = v25
                    default:
                        v2 -= (v2 + (v2 | 425))
                    }
                    v19 = v19
                    if true {
                        v18 += (v18 << ((555) & 63))
                        var v26 bool = false
                        _ = v26
                    } else {
                        v27 := v3[("foo bar" + (("foo bar" + "") + "\n\t"))]
                        _ = v27
                    }
                    _ = fi21
                }
            }
            v28 := GenS1{F0: 32767, F1: 65535, F2: -1}
            _ = v28
        }
        _ = v19
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 uint16 = (v2 / ((42) | 1))
        fmt.Println(v5, v2)
        v1 = append(v1, "hello")
        _ = fi29
        _ = v30
    }
    var v31 int16 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
