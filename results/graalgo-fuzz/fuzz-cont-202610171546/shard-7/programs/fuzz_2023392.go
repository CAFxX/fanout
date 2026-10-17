package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int, p1 string, p2 uint64) int32 {
    p1 = "foo bar"
    var v5 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 uint64, p1 uint32) int16 {
    var v6 bool = true
    fmt.Println(p0)
    v7 := GenS1{F0: 743, F1: -55, F2: -1}
    p1 = p1
    if true {
        switch ("" + ("a" + ("" + "a"))) {
        case "foo bar":
            v8 := []int{-9223372036854775808, -1, 1}
            _ = v8
        case "":
            var v9 float32 = -2.25
            _ = v9
        case "\n\t":
            v10 := GenS1{F0: 1, F1: 1, F2: 235}
            _ = v10
        default:
            v11 := GenS1{F0: -374, F1: 1, F2: 42}
            _ = v11
        }
        var v12 uint16 = 65535
        var v13 uint32 = 2
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 644
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 42
    v3 := make(map[string]string)
    v4 := []int{127, 127}
    v2 = 793
    v14 := GenS2{F0: 42}
    if true {
        v2 -= 758
    }
    v15 := v1["a"]
    if (172 <= 0) {
        fmt.Println(v2)
        var v16 int16 = 546
        _ = v16
    } else {
        v4 = append(v4, 127)
        fmt.Println(v2)
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 int64 = 42
        _ = fi17
        _ = v18
    }
    switch -9223372036854775808 {
    case 1:
        if (false && true) {
            v19 := GenS1{F0: 1, F1: -128, F2: 603}
            delete(v3, "\n\t")
            v2 = (v2 >> ((0) & 63))
            _ = v19
        } else {
            var v20 uint64 = uint64(v15)
            var v21 uint32 = 69
            fmt.Println(v2, v20)
            _ = v20
            _ = v21
        }
    case 32768:
        switch -850 {
        case 773:
            v15 = (v15 >> (((v2 | v2)) & 63))
        case 669:
            var v22 float64 = -2.25
            _ = v22
        case 4294967295:
            v15 -= v15
        default:
            v2 = v2
        }
    case -2147483648:
        var v23 string = ""
        _ = v23
    default:
        fmt.Println(v2, v2)
    }
    fmt.Println(v15, v15)
    for fi24 := 0; fi24 < 4; fi24++ {
        var v25 string = "foo bar"
        var v26 int = (fi24 ^ fi24)
        _ = fi24
        _ = v25
        _ = v26
    }
    if true {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            var v27 int32 = 2147483647
            fmt.Println(v2, v2)
            _ = v27
        }
        fmt.Println(v15, v15)
        var v28 uint32 = (v15 + (^(v15)))
        _ = v28
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v15 = 4294967295
    v29 := v3[("\n\t" + "hello")]
    var v30 uint32 = v15
    _ = len(v4)
    var v31 uint32 = v30
    var v32 int = 42
    if false {
        fmt.Println(v2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
