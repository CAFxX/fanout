package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
func genFn0(p0 string, p1 string, p2 int64) int64 {
    v4 := []int{32767, 32767, 1, 4294967295}
    var v5 int64 = -1
    fmt.Println(p2)
    var v6 bool = true
    v7 := GenS1{F0: -86}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (p2 + (p2 + -119))
}

func main() {
    var v1 float64 = 1.5
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    v2 = "hello"
    v1 = (((((v1 + float64(1e10)) / (1e10 * v1)) - float64((3.14159 * float32(1.5)))) + v1) + v1)
    v8 := GenS1{F0: -82}
    var v9 uint8 = 42
    v10 := GenS1{F0: -108}
    _ = len(v3)
    _ = len(v3)
    v11 := []int{42, -642}
    v2 = v2
    if true {
        switch 9223372036854775807 {
        case 42:
            var v12 int8 = -106
            _ = v12
        case 106:
            v1 = 0.1
        case 32768:
            fmt.Println(v1)
        }
        var v13 int8 = 100
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = v13
    }
    var v14 int = 652
    for fi15 := 0; fi15 < 13; fi15++ {
        if ((!false) || false) {
            var v16 uint = 255
            var v17 int = (fi15 * fi15)
            var v18 int16 = 42
            _ = v16
            _ = v17
            _ = v18
        } else {
            v19 := GenS2{F0: v2}
            _ = v19
        }
        fmt.Println(v14, v2, fi15)
        _ = fi15
    }
    if (v9 == v9) {
        var v20 int = 127
        v20 -= v14
        _ = v20
    } else {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        var v21 uint64 = 18446744073709551615
        v21 -= 348
        _ = v21
    }
    var v22 int64 = -758
    var v23 uint32 = 546
    switch 4294967295 {
    case 42:
        v14 -= -144
    }
    switch v2 {
    case "":
        var v24 int32 = 0
        _ = v24
    default:
        if ((2 <= 854) || false) {
            for fi25 := 0; fi25 < 15; fi25++ {
                var v26 int16 = -427
                v9 -= v9
                _ = fi25
                _ = v26
            }
        }
    }
    switch v2 {
    case "hello":
        var v27 string = v2
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v8
    _ = v9
}
