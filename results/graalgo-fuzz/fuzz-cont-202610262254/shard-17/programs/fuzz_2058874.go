package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint; F1 uint; F2 string }
func genFn0(p0 uint8) uint16 {
    for fi7 := 0; fi7 < 3; fi7++ {
        p0 -= (^(p0))
        v8 := GenS1{F0: (p0 ^ p0)}
        switch "hello" {
        case "a":
            for fi9 := 0; fi9 < 10; fi9++ {
                if false {
                    p0 = p0
                    p0 += p0
                    var v10 float32 = 1e10
                    _ = v10
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        var v11 bool = false
                        _ = v11
                    case "":
                        fmt.Println(fi7)
                    case "hello":
                        v12 := GenS1{F0: p0}
                        _ = v12
                    }
                }
                if false {
                    var v13 uint16 = 81
                    v14 := GenS2{F0: fi7}
                    _ = v13
                    _ = v14
                }
                p0 = (p0 ^ (p0 | p0))
                _ = fi9
            }
        case "hello":
            var v15 int = (fi7 & -2147483648)
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        switch "\n\t" {
        case "hello":
            switch "foo bar" {
            case "":
                p0 += (p0 + p0)
            case "hello":
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 string = ("a" + "")
                    var v19 int16 = -1
                    v20 := GenS3{F0: 42, F1: 1}
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
            default:
                var v21 string = ((("hello" + ("\n\t" + ("a" + ("hello" + "\n\t")))) + "a") + "hello")
                _ = v21
            }
        case "a":
            fmt.Println(fi16)
        case "foo bar":
            var v22 int = fi16
            _ = v22
        }
        v23 := []int{(fi16 - 2147483647), (fi16 / ((fi16) | 1))}
        _ = fi16
        _ = v23
    }
    _ = p0
    return 304
}

func main() {
    var v1 uint64 = 345
    v2 := []int{32768, -9223372036854775808, 127}
    var v3 string = ""
    var v4 float32 = (3.14159 * 0.0)
    var v5 int8 = -109
    v6 := make(map[int]bool)
    v24 := GenS4{F0: 1, F1: 255}
    for fi25 := 0; fi25 < 5; fi25++ {
        for fi26 := 0; fi26 < 16; fi26++ {
            v4 += v4
            _ = fi26
        }
        var v27 int16 = 42
        v27 = -191
        _ = fi25
        _ = v27
    }
    v28 := GenS5{F0: v5, F1: 127, F2: 818}
    if false {
        for fi29 := 0; fi29 < 13; fi29++ {
            v5 += v5
            _ = fi29
        }
    }
    var v30 int8 = (v5 & v5)
    v2 = append(v2, 1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        v32 := GenS6{F0: 906, F1: 338, F2: v3}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v4)
        _ = fi31
        _ = v32
    }
    var v33 uint64 = 1
    v6[127] = (false || true)
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
