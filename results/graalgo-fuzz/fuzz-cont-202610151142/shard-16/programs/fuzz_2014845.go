package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int; F1 uint64; F2 uint64 }
func genFn0(p0 int, p1 uint32) uint8 {
    v4 := GenS1{F0: 1, F1: 57}
    v5 := GenS2{F0: int16(p0)}
    var v6 uint8 = 98
    v7 := GenS1{F0: 76, F1: 42}
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 42, F1: 292}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (v6 & (v6 ^ (v6 ^ v6)))
}

func main() {
    var v1 bool = ((!(!true)) && ((-280 < 42) || ((true && false) && false)))
    var v2 int = 32768
    v3 := make(map[string]int)
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            if v1 {
                v12 := []int{(fi10 - v2), (-(v2))}
                if v1 {
                    switch 814 {
                    case -450:
                        var v13 uint8 = 17
                        _ = v13
                    case -1:
                        switch "hello" {
                        case "a":
                            if ((genFn0((^(fi11)), 65535) != 2) || true) {
                                v14 := GenS3{F0: 42}
                                _ = v14
                            }
                        case "\n\t":
                            v15 := GenS2{F0: 32767}
                            _ = v15
                        case "foo bar":
                            var v16 uint = 4294967295
                            _ = v16
                        }
                        _ = len(v3)
                        var v17 bool = (!((float32(0.1) - (0.0 + 0.0)) < 3.14159))
                        _ = v17
                    }
                    v2 = 4294967295
                }
                switch v2 {
                case 9223372036854775807:
                    var v18 uint64 = 140
                    _ = v18
                }
                _ = v12
            }
            var v19 uint16 = 2
            v20 := GenS3{F0: 42}
            _ = fi11
            _ = v19
            _ = v20
        }
        v2 += fi10
        _ = fi10
    }
    v21 := GenS4{F0: 1, F1: -1}
    if v1 {
        if (42 != 4294967295) {
            if false {
                switch "foo bar" {
                case "":
                    fmt.Println(v2, v2)
                default:
                    var v22 uint = 18446744073709551615
                    _ = v22
                }
                var v23 uint32 = 2
                _ = v23
            } else {
                v1 = v1
                v2 = v2
                v24 := GenS1{F0: 42, F1: 255}
                _ = v24
            }
            var v25 float64 = 0.0
            _ = v25
        }
        var v26 bool = false
        _ = v26
    } else {
        switch (v2 % ((v2) | 1)) {
        case -997:
            var v27 int16 = -1
            _ = v27
        case 4294967295:
            fmt.Println(v2, v1)
        case 406:
            v28 := v3[("foo bar" + "a")]
            _ = v28
        }
    }
    var v29 uint16 = 625
    v30 := GenS5{F0: v2, F1: 2, F2: 1}
    var v31 int = v2
    var v32 int64 = -9223372036854775808
    v33 := []int{661, (v2 * v2)}
    v29 += (^(v29))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
