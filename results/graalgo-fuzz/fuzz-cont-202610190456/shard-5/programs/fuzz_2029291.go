package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int }
func genFn0(p0 string) uint {
    v4 := GenS1{F0: -1, F1: 42, F2: 0}
    v5 := []int{42, -362}
    switch "" {
    case "a":
        p0 = p0
    case "hello":
        for fi6 := 0; fi6 < 2; fi6++ {
            p0 = p0
            _ = fi6
        }
    case "\n\t":
        var v7 float64 = 1e10
        _ = v7
    default:
        _ = len(v5)
    }
    v8 := GenS2{F0: 0, F1: 2, F2: ((true || true) || false)}
    p0 = ((p0 + p0) + ("hello" + p0))
    _ = p0
    _ = v4
    _ = v5
    _ = v8
    return 4294967295
}

func main() {
    var v1 int = 127
    var v2 int16 = -32768
    var v3 int32 = -1
    var v9 float32 = -2.25
    v10 := GenS1{F0: v2, F1: 12, F2: -1}
    var v11 uint16 = 1
    v12 := []int{(v1 * (v1 | 42)), (v1 >> ((983) & 63)), (-(v1))}
    v11 = 497
    switch v1 {
    case -469:
        v3 += v3
    default:
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v3)
        }
    }
    switch (v1 & v1) {
    case -39:
        var v13 float64 = 0.1
        _ = v13
    case 631:
        switch (v1 * (v1 ^ (^(v1)))) {
        case 587:
            var v14 string = "hello"
            _ = v14
        case -1:
            v9 = v9
        case 17:
            v3 += (v3 - v3)
        }
    default:
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    v9 = (v9 * (v9 - v9))
    v12 = append(v12, v1)
    for fi15 := 0; fi15 < 6; fi15++ {
        fmt.Println(v1, v9, v2)
        _ = fi15
    }
    if false {
        _ = len(v12)
        switch "foo bar" {
        case "":
            var v16 uint = 491
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 11; fi17++ {
            fmt.Println(fi17, v3, v2)
            v12 = append(v12, -2147483648)
            fmt.Println(v1, fi17)
            _ = fi17
        }
        v2 += (-(v2))
    }
    v18 := GenS2{F0: 374, F1: 255, F2: (false && false)}
    fmt.Println(v1)
    v19 := GenS3{F0: genFn0("a")}
    if false {
        var v20 int = v1
        fmt.Println(v3)
        if false {
            switch "\n\t" {
            case "foo bar":
                for fi21 := 0; fi21 < 12; fi21++ {
                    v22 := GenS3{F0: 2}
                    _ = fi21
                    _ = v22
                }
            }
            v20 = (^(v1))
        } else {
            for fi23 := 0; fi23 < 10; fi23++ {
                if true {
                    v24 := GenS4{F0: 127, F1: v11}
                    _ = v24
                }
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi23
            }
        }
        _ = v20
    } else {
        v1 = v1
        v9 = v9
    }
    v25 := GenS5{F0: (-(v1))}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v9
}
