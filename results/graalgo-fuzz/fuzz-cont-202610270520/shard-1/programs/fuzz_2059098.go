package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int32 = 42
    var v2 uint8 = 0
    v3 := []string{"\n\t", "foo bar", "foo bar"}
    var v4 uint16 = 42
    v5 := []uint32{2, 65535, 634}
    var v6 float64 = 1e10
    v7 := GenS1{F0: 127}
    if true {
        if (2 < 65535) {
            var v8 int8 = -89
            _ = v8
        } else {
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 uint8 = 2
                _ = fi9
                _ = v10
            }
        }
        var v11 int16 = -32768
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            if (42 < 2147483648) {
                v4 -= v4
                var v14 int8 = -124
                var v15 int8 = v14
                _ = v14
                _ = v15
            }
            _ = fi13
        }
        _ = fi12
    }
    fmt.Println(v2, v1)
    if (!true) {
        switch "a" {
        case "a":
            _ = len(v5)
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := GenS2{F0: 0, F1: 42}
            _ = fi16
            _ = v17
        }
        var v18 int16 = 1000
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        if (true || true) {
            _ = len(v3)
            var v20 int8 = -8
            v21 := GenS1{F0: (^(v20))}
            _ = v20
            _ = v21
        }
        var v22 uint16 = v4
        v6 += v6
        _ = fi19
        _ = v22
    }
    fmt.Println(v4, v1, v1)
    var v23 int8 = 42
    if (!true) {
        v1 = v1
    } else {
        var v24 uint16 = 857
        var v25 bool = true
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v23 += v23
        v5 = append(v5, 707)
        switch fi26 {
        case -9223372036854775808:
            if (((2 <= 0) || false) || (!false)) {
                var v27 uint8 = (v2 & 79)
                for fi28 := 0; fi28 < 19; fi28++ {
                    _ = len(v3)
                    v23 = (v23 ^ (-(v23)))
                    v29 := GenS2{F0: (v23 | (v23 / (((^(v23))) | 1))), F1: 255}
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
        case 42:
            var v30 uint = 18446744073709551615
            _ = v30
        default:
            switch "\n\t" {
            case "\n\t":
                _ = len(v5)
            case "":
                if len(v3) > 0 {
                    v3[0] = ""
                }
            }
        }
        _ = fi26
    }
    v31 := GenS3{F0: 42}
    v5 = append(v5, 929)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
