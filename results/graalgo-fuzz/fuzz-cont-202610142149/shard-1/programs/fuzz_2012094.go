package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
type GenS3 struct { F0 uint8; F1 int32 }

func main() {
    v1 := []uint32{2, 971, 2147483648, 0}
    v2 := make(map[uint32]string)
    var v3 float64 = -2.25
    v4 := make(map[string]bool)
    var v5 int = -1
    v6 := GenS1{F0: 650, F1: 766, F2: 977}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint8 = 30
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (!(!true)) {
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 bool = true
                switch fi7 {
                case -1:
                    var v11 uint16 = 32767
                    _ = v11
                case 101:
                    var v12 float64 = -2.25
                    _ = v12
                case 32768:
                    var v13 uint = 65535
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
        } else {
            v8 -= 149
            var v14 uint64 = 18446744073709551615
            for fi15 := 0; fi15 < 12; fi15++ {
                v14 = (v14 - 1)
                var v16 uint8 = (v8 % ((v8) | 1))
                v17 := GenS2{F0: 63, F1: 65535, F2: (false || true)}
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                switch ("\n\t" + "hello") {
                case "hello":
                    v8 += (^(v8))
                case "a":
                    var v18 string = (("" + (("a" + "") + "foo bar")) + "a")
                    _ = v18
                case "foo bar":
                    var v19 int16 = 1000
                    _ = v19
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "":
        var v20 uint32 = 42
        _ = v20
    case "\n\t":
        var v21 uint8 = 127
        _ = v21
    default:
        v5 -= v5
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 bool = (0 != 0)
        var v24 uint16 = 396
        if len(v1) > 0 {
            v1[0] = 2
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    _ = len(v1)
    v25 := GenS3{F0: 89, F1: -996}
    for fi26 := 0; fi26 < 17; fi26++ {
        for fi27 := 0; fi27 < 16; fi27++ {
            v3 = (v3 * 0.0)
            v4["\n\t"] = true
            var v28 bool = (!(v3 == v3))
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    var v29 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
