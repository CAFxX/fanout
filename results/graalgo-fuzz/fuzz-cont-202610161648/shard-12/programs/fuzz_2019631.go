package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 uint32; F2 uint32 }

func main() {
    var v1 int8 = 127
    var v2 float64 = ((0.1 + 3.14159) - 1e10)
    var v3 uint = 18446744073709551615
    var v4 uint64 = 4294967295
    v5 := make(map[string]uint32)
    v6 := make(map[string]string)
    var v7 string = "\n\t"
    v7 = v7
    v8 := GenS1{F0: 127, F1: 32768, F2: 103}
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int = 4294967295
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := GenS1{F0: 58, F1: 0, F2: (fi11 / ((fi11) | 1))}
        switch fi11 {
        case 309:
            v2 -= v2
        case 128:
            if (v7 != v7) {
                v13 := v5[v7]
                fmt.Println(v4, v13)
                _ = v13
            } else {
                if true {
                    var v14 int8 = v1
                    v14 -= v14
                    delete(v6, "foo bar")
                    _ = v14
                }
                var v15 uint8 = 2
                v16 := GenS2{F0: 1, F1: int(v15)}
                _ = v15
                _ = v16
            }
        default:
            _ = len(v6)
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := GenS3{F0: v3}
            switch v7 {
            case "foo bar":
                v19 := GenS4{F0: -1}
                _ = v19
            case "\n\t":
                if (!(!true)) {
                    var v20 uint64 = 740
                    _ = v20
                } else {
                    v2 += v2
                }
                v7 = v7
            case "a":
                fmt.Println(v7)
            }
            v21 := GenS5{F0: 1, F1: 287, F2: 2}
            _ = fi17
            _ = v18
            _ = v21
        }
        _ = fi11
        _ = v12
    }
    v2 += (v2 / v2)
    var v22 float64 = (0.0 - v2)
    fmt.Println(v22)
    if (false || true) {
        v6[v7] = (v7 + v7)
        v23 := GenS5{F0: 42, F1: 2, F2: 129}
        _ = v23
    }
    switch 1 {
    case 9223372036854775807:
        switch ((v7 + "hello") + "a") {
        case "a":
            v1 -= v1
        }
    }
    if false {
        v24 := GenS5{F0: 0, F1: 2147483648, F2: 42}
        v4 -= v4
        _ = v24
    }
    switch "a" {
    case "hello":
        fmt.Println(v22, v1, v3)
    case "a":
        if false {
            switch v7 {
            case "":
                v2 += (v2 / (1e10 + 1e10))
            case "a":
                var v25 uint = v3
                _ = v25
            default:
                var v26 uint16 = 369
                _ = v26
            }
        }
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        v4 -= uint64(v3)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
