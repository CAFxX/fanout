package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8 }
func genFn0() uint16 {
    var v7 bool = false
    v8 := []int{}
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 int, p1 bool, p2 int32) uint16 {
    if (p1 || p1) {
        var v9 bool = ((p2 * (^(p2))) == p2)
        var v10 float64 = (1.5 + 1e10)
        _ = v10
        _ = v9
    }
    if (!p1) {
        var v11 int = p0
        switch (v11 - -2147483648) {
        case 465:
            v12 := []int{(p0 * v11), 2147483647}
            _ = v12
        default:
            v13 := []int{v11, (p0 + v11), 3, (p0 / ((1) | 1))}
            _ = v13
        }
        var v14 uint = 2
        _ = v11
        _ = v14
    } else {
        if p1 {
            if p1 {
                fmt.Println(p0, p2, p1)
                v15 := []int{}
                _ = v15
            }
            v16 := GenS1{F0: 1, F1: 1, F2: "hello"}
            if (0.1 > 1e10) {
                v17 := GenS2{F0: "", F1: 2147483648}
                v18 := GenS3{F0: 42}
                for fi19 := 0; fi19 < 15; fi19++ {
                    v20 := GenS2{F0: "\n\t", F1: 65535}
                    _ = fi19
                    _ = v20
                }
                _ = v17
                _ = v18
            } else {
                var v21 uint8 = 0
                for fi22 := 0; fi22 < 16; fi22++ {
                    switch ("" + "hello") {
                    case "\n\t":
                        fmt.Println(fi22, p0, p1)
                    case "a":
                        v21 = 234
                    case "foo bar":
                        fmt.Println(p1, v21)
                    }
                    _ = fi22
                }
                _ = v21
            }
            _ = v16
        }
    }
    var v23 uint = 18446744073709551615
    p1 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 0
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int16 = -692
    var v3 int16 = v2
    var v4 uint16 = 0
    v5 := make(map[int]bool)
    var v6 float64 = ((-2.25 - 1e10) + (((1e10 * -2.25) + ((1.5 * float64(0.1)) / 0.1)) * 3.14159))
    if false {
        v4 -= genFn1(int(v1), false, 413)
        v24 := GenS3{F0: 42}
        _ = v24
    }
    var v25 uint64 = 4294967295
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 uint16 = genFn0()
        v28 := []int{fi26, fi26, (fi26 & fi26)}
        v5[fi26] = true
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS4{F0: false}
    v30 := GenS1{F0: 42, F1: v1, F2: "hello"}
    fmt.Println(v4, v6)
    if false {
        v31 := GenS5{F0: -915}
        delete(v5, 32768)
        _ = v31
    } else {
        for fi32 := 0; fi32 < 4; fi32++ {
            switch fi32 {
            case 4294967295:
                switch "\n\t" {
                case "hello":
                    var v33 uint8 = 203
                    _ = v33
                default:
                    var v34 uint = 153
                    _ = v34
                }
            case -1:
                v1 = uint32(v4)
            case 42:
                var v35 int32 = 902
                _ = v35
            default:
                v36 := GenS6{F0: (v1 / ((65535) | 1))}
                _ = v36
            }
            v2 = v3
            _ = fi32
        }
        v37 := []int{-658}
        var v38 bool = (uint32(v25) < (v1 * v1))
        _ = v37
        _ = v38
    }
    v39 := GenS7{F0: 219}
    v40 := []int{1, -2147483648, 59}
    var v41 uint32 = v1
    for fi42 := 0; fi42 < 18; fi42++ {
        v5[(^(fi42))] = true
        _ = fi42
    }
    delete(v5, -2147483648)
    switch "\n\t" {
    case "\n\t":
        if (200 >= 42) {
            v43 := GenS1{F0: 42, F1: 501, F2: "\n\t"}
            v2 = 1000
            v5[212] = ((false && false) && ((v1 - (v1 | v1)) <= 4294967295))
            _ = v43
        }
    case "foo bar":
        _ = len(v40)
    case "":
        v44 := GenS2{F0: ("hello" + ("\n\t" + ("foo bar" + "hello"))), F1: v41}
        _ = v44
    default:
        v5[4294967295] = true
    }
    v4 = (^(v4))
    v2 = v2
    switch 9223372036854775807 {
    case 42:
        var v45 int8 = 0
        _ = v45
    case 9223372036854775807:
        v46 := GenS7{F0: 123}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
