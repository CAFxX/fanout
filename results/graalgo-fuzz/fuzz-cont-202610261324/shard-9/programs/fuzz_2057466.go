package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int16) int8 {
    p1 = -1
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 = (^(v6))
        v6 += 0
        _ = fi7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            if false {
                var v10 uint32 = (p0 & (p0 | (p0 + 42)))
                var v11 uint64 = 18446744073709551615
                _ = v10
                _ = v11
            }
            fmt.Println(p0)
            var v12 int8 = 42
            _ = fi9
            _ = v12
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS1{F0: -70, F1: -1}
            p0 += 2
            var v15 uint64 = 1
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v6
    return 100
}

func main() {
    var v1 float64 = 3.14159
    v2 := []uint32{585, 1, 0, 0}
    var v3 int8 = -128
    v4 := []string{"\n\t", "foo bar"}
    var v5 uint8 = 255
    fmt.Println(v3, v3)
    for fi16 := 0; fi16 < 15; fi16++ {
        if (8 > (v3 / (((v3 >> ((625) & 63))) | 1))) {
            switch "foo bar" {
            case "\n\t":
                if len(v2) > 0 {
                    v2[0] = 1
                }
            case "":
                if true {
                    var v17 uint16 = 55
                    _ = v17
                }
            case "a":
                v3 = -55
            default:
                v3 -= v3
            }
        }
        var v18 int16 = 42
        v3 = v3
        _ = fi16
        _ = v18
    }
    v3 += -2
    v1 -= 1.5
    switch ("hello" + ((("" + "a") + "a") + "")) {
    case "\n\t":
        v19 := GenS2{F0: (((0.1 - 1e10) * v1) < 1.5), F1: 41, F2: 0}
        _ = v19
    case "hello":
        _ = len(v2)
    }
    fmt.Println(v1)
    if (!false) {
        fmt.Println(v5, v3)
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v3, fi20, v5)
            v1 = v1
            var v21 int16 = 42
            _ = fi20
            _ = v21
        }
    }
    fmt.Println(v5)
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v22 string = "a"
    v23 := GenS3{F0: v5}
    v24 := GenS3{F0: (v5 - 2)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
