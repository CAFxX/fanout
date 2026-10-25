package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint32, p1 uint8, p2 bool) bool {
    v6 := []int{9223372036854775807, 42}
    p2 = (((1e10 + 0.0) / 3.14159) <= (3.14159 + -2.25))
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (2 >= 0)
}
func genFn1(p0 uint64, p1 bool, p2 bool) int32 {
    v7 := []int{1, 72}
    var v8 int32 = -133
    p2 = (false && true)
    switch ("hello" + ("hello" + "foo bar")) {
    case "foo bar":
        var v9 int16 = -32768
        _ = v9
    default:
        v10 := GenS1{F0: -1}
        _ = v10
    }
    var v11 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = (-2.25 + 0.0)
    var v3 uint16 = 769
    v4 := []int{-2147483648, 552, 182}
    v5 := make(map[uint32]int)
    if len(v4) > 0 {
        v4[0] = 32768
    }
    v2 = v2
    var v12 int8 = 100
    var v13 int32 = 1
    var v14 int8 = 100
    switch -109 {
    case 1:
        if ((127 != 128) && false) {
            var v15 uint = 708
            delete(v1, "")
            v16 := GenS2{F0: v15}
            _ = v15
            _ = v16
        } else {
            v12 += (v14 | (v14 / (((v14 << ((65535) & 63))) | 1)))
        }
    case -1:
        for fi17 := 0; fi17 < 4; fi17++ {
            fmt.Println(v3, v2, v12)
            v18 := GenS3{F0: 889, F1: 255, F2: 128}
            v12 = (-(v14))
            _ = fi17
            _ = v18
        }
    case 127:
        v19 := GenS2{F0: 4294967295}
        _ = v19
    default:
        switch "a" {
        case "a":
            switch ((("foo bar" + "\n\t") + "foo bar") + "\n\t") {
            case "foo bar":
                fmt.Println(v12, v12)
            case "a":
                if false {
                    v14 = v14
                    v3 -= v3
                }
            default:
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(v12, v2)
                    v21 := GenS4{F0: (true || (42 < -890))}
                    var v22 string = ("" + "hello")
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        case "":
            var v23 uint = 704
            _ = v23
        case "hello":
            v5[2147483648] = 32767
        }
    }
    v3 -= v3
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 uint16 = (v3 | 1)
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: v12}
    for fi27 := 0; fi27 < 15; fi27++ {
        _ = len(v4)
        fmt.Println(v3, v14)
        switch fi27 {
        case 32767:
            if len(v4) > 0 {
                v4[0] = 128
            }
        }
        _ = fi27
    }
    var v28 uint8 = 0
    v29 := v5[0]
    if true {
        v12 = 118
        var v30 string = "hello"
        _ = len(v4)
        _ = v30
    } else {
        v31 := GenS4{F0: true}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v31
    }
    v2 = (v2 / v2)
    var v32 uint8 = (v28 * (v28 + v28))
    var v33 float32 = 0.0
    var v34 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
