package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint8; F1 int8 }

func main() {
    v1 := []uint32{2147483648}
    var v2 bool = (!(54 < 36))
    var v3 float32 = 0.1
    var v4 float64 = 0.1
    v5 := []uint32{}
    fmt.Println(v4)
    v6 := GenS1{F0: 1, F1: 4294967295}
    if (("" < ("a" + "hello")) || true) {
        v4 = float64(0.0)
    } else {
        v7 := GenS1{F0: 2, F1: 2147483647}
        switch "" {
        case "":
            if v2 {
                if len(v5) > 0 {
                    v5[0] = 822
                }
                var v8 float64 = ((v4 * (0.1 / v4)) * 1e10)
                if len(v1) > 0 {
                    v1[0] = 2
                }
                _ = v8
            }
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = 116
            }
        default:
            switch "\n\t" {
            case "foo bar":
                if v2 {
                    v9 := GenS1{F0: 4294967295, F1: 4294967295}
                    _ = v9
                }
            case "":
                if (((v2 || (121 != 255)) && v2) && (v2 && (230 < 649))) {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        v11 := GenS2{F0: 79, F1: 59}
                        var v12 int8 = 100
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    v4 -= v4
                } else {
                    var v13 bool = v2
                    v3 = 3.14159
                    _ = v13
                }
            case "hello":
                fmt.Println(v3, v4)
            }
        }
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 uint64 = 18446744073709551615
    var v15 float32 = 3.14159
    v16 := GenS1{F0: 65535, F1: 128}
    var v17 int16 = -32768
    v18 := GenS2{F0: 0, F1: int8(v14)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
