package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS2 struct { F0 bool; F1 bool }
func genFn0(p0 uint16, p1 uint16, p2 int32) int64 {
    switch "hello" {
    case "a":
        fmt.Println(p2, p2)
    case "":
        switch "foo bar" {
        case "foo bar":
            var v5 uint16 = p0
            _ = v5
        }
    case "\n\t":
        for fi6 := 0; fi6 < 4; fi6++ {
            p1 += 977
            fmt.Println(fi6, p1, p2)
            _ = fi6
        }
    default:
        p0 = p0
    }
    p0 += p1
    var v7 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int8 = 100
    var v3 uint = 4294967295
    v4 := make(map[int]bool)
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        var v11 int8 = (-(v2))
        v3 -= (v3 / (((v3 | 42)) | 1))
        _ = fi8
        _ = v11
    }
    delete(v4, 127)
    delete(v4, -211)
    v12 := GenS1{F0: 564, F1: 127, F2: -1}
    if (0 >= 255) {
        fmt.Println(v2, v2, v2)
        fmt.Println(v2, v3, v3)
        var v13 uint32 = 0
        _ = v13
    } else {
        v3 += 2
        switch "hello" {
        case "hello":
            var v14 uint16 = 115
            _ = v14
        case "foo bar":
            for fi15 := 0; fi15 < 17; fi15++ {
                v3 += (v3 & (v3 % ((923) | 1)))
                if true {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v1 = v1
                        _ = fi16
                    }
                    var v17 float64 = 1e10
                    _ = v17
                }
                var v18 uint32 = 4294967295
                _ = fi15
                _ = v18
            }
        case "":
            v19 := GenS1{F0: 319, F1: 128, F2: 42}
            _ = v19
        }
        fmt.Println(v3, v2)
    }
    v20 := []int{}
    if ((1 >= (^(v1))) || (2 <= 2147483648)) {
        v21 := GenS1{F0: 255, F1: 152, F2: -2147483648}
        _ = v21
    } else {
        if (0.1 < 1e10) {
            v3 += v3
        }
        if false {
            if len(v20) > 0 {
                v20[0] = 2147483647
            }
            for fi22 := 0; fi22 < 7; fi22++ {
                switch -826 {
                case -877:
                    var v23 float32 = 0.0
                    _ = v23
                case 633:
                    fmt.Println(fi22, v1, v2)
                default:
                    v24 := GenS1{F0: 2, F1: 0, F2: 872}
                    _ = v24
                }
                _ = len(v4)
                v25 := GenS2{F0: false, F1: true}
                _ = fi22
                _ = v25
            }
        }
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v1 += v1
    var v26 int64 = 9223372036854775807
    switch ((("a" + ("hello" + "foo bar")) + (("hello" + ("a" + "a")) + "hello")) + ("foo bar" + "")) {
    case "a":
        v2 = (v2 ^ (^(v2)))
    case "\n\t":
        v3 += v3
    }
    var v27 uint64 = (^(v1))
    var v28 int16 = -77
    var v29 bool = false
    v4[128] = (!(!v29))
    for fi30 := 0; fi30 < 10; fi30++ {
        v3 = (v3 + 380)
        var v31 bool = (!false)
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        fmt.Println(fi32, v28)
        v33 := GenS2{F0: (!((42 < 425) && v29)), F1: true}
        _ = fi32
        _ = v33
    }
    v3 += (v3 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
