package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint64; F2 string }
type GenS3 struct { F0 bool }

func main() {
    v1 := []uint32{}
    v2 := []int{32767, 0, -451, 0}
    v3 := make(map[int]bool)
    var v4 uint8 = 231
    var v5 int32 = 42
    var v6 uint8 = (v4 | v4)
    var v7 float32 = ((float32(1e10) * float32(1e10)) / 1e10)
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v5, v7)
        v2 = append(v2, (-(fi8)))
        var v9 uint8 = (v6 % ((191) | 1))
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: (^(v6))}
    if (!(false && (((!false) || (!(false && false))) || false))) {
        fmt.Println(v6, v7)
    } else {
        var v11 uint64 = 1
        var v12 int16 = -694
        var v13 int64 = -1
        _ = v11
        _ = v12
        _ = v13
    }
    switch ("a" + ("foo bar" + "\n\t")) {
    case "a":
        switch "foo bar" {
        case "":
            fmt.Println(v7, v4, v5)
        case "a":
            var v14 uint8 = (v4 % ((v4) | 1))
            _ = v14
        case "hello":
            v15 := GenS1{F0: 128}
            _ = v15
        default:
            var v16 int8 = -128
            _ = v16
        }
    case "":
        var v17 int8 = 100
        _ = v17
    }
    v4 -= v6
    var v18 uint16 = 0
    for fi19 := 0; fi19 < 10; fi19++ {
        v18 = v18
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(v7, fi20)
            v21 := GenS2{F0: 372, F1: 1, F2: ""}
            _ = fi20
            _ = v21
        }
        if len(v1) > 0 {
            v1[0] = 236
        }
        _ = fi19
    }
    var v22 uint16 = (v18 % (((v18 | (v18 & (v18 & 80)))) | 1))
    for fi23 := 0; fi23 < 3; fi23++ {
        v24 := GenS2{F0: 1, F1: 1, F2: (("\n\t" + ("hello" + ("a" + ""))) + "a")}
        if true {
            var v25 float64 = float64(v7)
            v5 = 2147483647
            var v26 uint8 = 43
            _ = v25
            _ = v26
        }
        v7 = 3.14159
        _ = fi23
        _ = v24
    }
    switch 42 {
    case 42:
        if (true && ((^(v4)) != v6)) {
            fmt.Println(v7, v5)
        }
    default:
        v4 += (v6 * v4)
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        if (1 <= int64(v6)) {
            v28 := GenS3{F0: false}
            _ = v28
        } else {
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 uint16 = 2
                v4 += 1
                var v31 float64 = 1.5
                _ = fi29
                _ = v30
                _ = v31
            }
            for fi32 := 0; fi32 < 18; fi32++ {
                v22 += (v18 % (((v18 + (v18 ^ (v22 - (^(v18)))))) | 1))
                _ = fi32
            }
            if true {
                v7 = float32(1.5)
                var v33 int32 = (v5 & (v5 - (-(v5))))
                fmt.Println(v33, v7, fi27)
                _ = v33
            } else {
                var v34 uint = 255
                v5 -= (v5 % ((-557) | 1))
                var v35 string = "\n\t"
                _ = v34
                _ = v35
            }
        }
        _ = fi27
    }
    v7 = (v7 / 1.0)
    v36 := GenS2{F0: 2, F1: 1, F2: ""}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
