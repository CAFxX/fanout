package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS2 struct { F0 string; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 int8 }
func genFn0(p0 int8) int64 {
    v4 := GenS1{F0: 389, F1: 127, F2: 0}
    fmt.Println(p0)
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 float64 = -2.25
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = v4
    return -307
}
func genFn1(p0 int8) uint {
    var v7 bool = (-128 != p0)
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint16 = 713
        v10 := GenS2{F0: "foo bar", F1: int8(v9), F2: 42}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: "foo bar", F1: (p0 ^ p0), F2: 42}
    if true {
        if (504 > 2) {
            var v12 int8 = p0
            fmt.Println(v7)
            _ = v12
        }
        if v7 {
            var v13 uint32 = 4294967295
            v14 := GenS2{F0: ("" + "\n\t"), F1: p0, F2: 1}
            if (!v7) {
                v15 := GenS3{F0: 107, F1: 182}
                _ = v15
            } else {
                v13 += (v13 - v13)
                var v16 uint = 2
                for fi17 := 0; fi17 < 14; fi17++ {
                    v18 := GenS4{F0: false, F1: 63}
                    v7 = (v7 || (-1 == 1))
                    fmt.Println(v16, p0)
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            _ = v13
            _ = v14
        } else {
            if v7 {
                var v19 int = 32767
                _ = v19
            }
        }
    } else {
        fmt.Println(p0, v7)
        var v20 float64 = 1.5
        _ = v20
    }
    _ = p0
    _ = v11
    _ = v7
    return 386
}

func main() {
    var v1 int8 = 0
    var v2 int8 = (v1 / (((v1 - (v1 << ((768) & 63)))) | 1))
    var v3 int16 = -984
    if ("foo bar" >= "foo bar") {
        var v21 uint = 723
        v22 := GenS5{F0: 2}
        var v23 uint64 = 906
        _ = v21
        _ = v22
        _ = v23
    } else {
        var v24 uint8 = 2
        var v25 float64 = 3.14159
        _ = v24
        _ = v25
    }
    v1 = (v2 << ((255) & 63))
    v26 := GenS6{F0: 4}
    fmt.Println(v2, v2, v1)
    v1 = (v1 | v1)
    if false {
        var v27 int16 = (v3 | v3)
        v27 += v27
        _ = v27
    }
    switch "hello" {
    case "a":
        var v28 uint8 = 195
        _ = v28
    case "hello":
        for fi29 := 0; fi29 < 9; fi29++ {
            v30 := []int{(-(fi29))}
            _ = fi29
            _ = v30
        }
    case "":
        v31 := GenS7{F0: -1, F1: (v1 ^ 114)}
        _ = v31
    }
    var v32 int8 = 0
    v1 = -121
    var v33 uint8 = 255
    v34 := GenS6{F0: genFn1((^(v1)))}
    v35 := GenS5{F0: 42}
    v32 -= (v1 | 86)
    fmt.Println(v33, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
