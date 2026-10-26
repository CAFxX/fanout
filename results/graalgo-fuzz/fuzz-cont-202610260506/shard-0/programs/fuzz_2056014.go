package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 bool; F1 int32; F2 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int; F2 uint }
type GenS7 struct { F0 int8 }
func genFn0(p0 int8, p1 uint8, p2 int16) bool {
    v6 := GenS1{F0: (p0 >> ((1) & 63))}
    v7 := []int{0, 2147483647, 365}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 uint32, p1 uint, p2 int64) uint {
    switch "" {
    case "\n\t":
        v8 := GenS1{F0: -1}
        _ = v8
    case "":
        v9 := GenS1{F0: 0}
        _ = v9
    case "a":
        switch "hello" {
        case "foo bar":
            switch (("foo bar" + ("foo bar" + "foo bar")) + "foo bar") {
            case "a":
                for fi10 := 0; fi10 < 18; fi10++ {
                    fmt.Println(p2, p0)
                    _ = fi10
                }
            case "":
                v11 := []int{}
                _ = v11
            case "hello":
                v12 := GenS2{F0: 42, F1: 875, F2: 2147483647}
                _ = v12
            default:
                v13 := GenS3{F0: "a", F1: 2}
                _ = v13
            }
        }
    }
    var v14 uint32 = 2147483648
    var v15 int64 = (p2 & 9223372036854775807)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return (p1 & 785)
}

func main() {
    var v1 uint8 = 250
    var v2 int = 9223372036854775807
    var v3 uint32 = 42
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    v16 := GenS3{F0: ("hello" + ""), F1: 546}
    if true {
        v2 -= v2
        v17 := GenS4{F0: (((v3 | v3) < (v3 + (^(v3)))) || (false && (!(v2 <= v2)))), F1: -1, F2: -168}
        v18 := GenS5{F0: 42}
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint64 = 0
                v2 = 127
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        var v22 string = ""
        if ((!true) || true) {
            if ((1e10 + -2.25) <= 3.14159) {
                var v23 int = (v2 / ((v2) | 1))
                _ = v23
            }
        }
        _ = v22
    }
    v4[212] = (2 == 4294967295)
    v24 := []int{(^(v2)), 128, -161, -622}
    v25 := GenS2{F0: -1, F1: 4294967295, F2: v3}
    fmt.Println(v3)
    v24 = append(v24, -723)
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 uint8 = (v1 & v1)
        v28 := GenS2{F0: -34, F1: 4294967295, F2: (v3 / ((v3) | 1))}
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v1, v3, v3)
    if (-2147483648 <= 1) {
        _ = len(v5)
        var v29 uint64 = 4294967295
        _ = v29
    } else {
        v30 := GenS6{F0: 2, F1: (v2 - (v2 - v2)), F2: genFn1(892, 827, 42)}
        if true {
            if false {
                v31 := v4[(v2 ^ v2)]
                var v32 uint32 = 4294967295
                _ = v31
                _ = v32
            }
        }
        _ = v30
    }
    delete(v4, (v2 ^ 805))
    var v33 float32 = (float32(((1.5 / (-2.25 + 0.0)) + -2.25)) * (3.14159 + 3.14159))
    v34 := GenS7{F0: -128}
    var v35 int32 = 0
    if false {
        v2 = 2147483647
    }
    switch 32768 {
    case 2147483647:
        var v36 float32 = ((v33 / v33) / (((v33 + v33) / v33) - v33))
        _ = v36
    case 128:
        delete(v4, v2)
    default:
        v24 = append(v24, (v2 | 4294967295))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
