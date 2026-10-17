package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 bool; F2 uint8 }

func main() {
    var v1 uint = 2
    var v2 string = ""
    v3 := []int{}
    v4 := make(map[string]bool)
    var v5 int32 = -1
    var v6 int = 42
    if false {
        var v7 float64 = 3.14159
        switch (v6 * (v6 ^ 127)) {
        case 32767:
            v8 := GenS1{F0: false, F1: -465, F2: 18446744073709551615}
            _ = v8
        }
        if false {
            var v9 int = (v6 ^ (v6 >> ((0) & 63)))
            fmt.Println(v9)
            _ = v9
        }
        _ = v7
    }
    v5 = 42
    v10 := GenS2{F0: (v6 * int(v5))}
    v11 := GenS1{F0: (!false), F1: -321, F2: 4294967295}
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 float32 = 1e10
            if (!false) {
                var v15 string = ""
                if false {
                    if (((255 < 255) || (("hello" + "") == (v15 + v2))) && true) {
                        fmt.Println(v5, v15)
                        var v16 float32 = v14
                        _ = len(v4)
                        v17 := GenS3{F0: v2, F1: false, F2: 0}
                        v16 -= v16
                        _ = v16
                        _ = v17
                    }
                    var v18 uint64 = 485
                    _ = v18
                } else {
                    _ = len(v3)
                }
                _ = v15
            } else {
                v4[v2] = false
                if (!(false && false)) {
                    var v19 uint16 = 0
                    v1 = 969
                    v6 = fi12
                    _ = v19
                }
                v6 = (v6 / ((fi12) | 1))
            }
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    switch ("" + v2) {
    case "a":
        v2 = "hello"
    case "foo bar":
        fmt.Println(v6, v6, v2)
    default:
        var v20 uint8 = 81
        _ = v20
    }
    switch v6 {
    case -785:
        var v21 uint16 = 1
        _ = v21
    case 32768:
        _ = len(v3)
    default:
        var v22 uint8 = 42
        _ = v22
    }
    v3 = append(v3, 42)
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 uint16 = 0
        fmt.Println(v5, v24, fi23)
        _ = fi23
        _ = v24
    }
    switch v6 {
    case -1:
        switch v6 {
        case 1:
            var v25 int16 = 1000
            _ = v25
        case -2147483648:
            var v26 uint = v1
            _ = v26
        }
    case 512:
        var v27 bool = true
        _ = v27
    case 4294967295:
        fmt.Println(v6, v5, v6)
    default:
        v28 := v4["foo bar"]
        _ = v28
    }
    if ((true && (1 == 0)) || (true || false)) {
        if (!false) {
            v29 := GenS3{F0: v2, F1: false, F2: 2}
            var v30 uint8 = 63
            _ = v29
            _ = v30
        }
    } else {
        switch "a" {
        case "\n\t":
            v3 = append(v3, (^(v6)))
        case "":
            var v31 int8 = int8(v1)
            _ = v31
        default:
            v6 -= -748
        }
    }
    var v32 int = 2147483647
    var v33 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
