package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0(p0 uint, p1 uint8) int32 {
    var v4 int8 = 42
    var v5 int = 32768
    if true {
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 int = v5
            _ = fi6
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 86
}
func genFn1(p0 uint32, p1 int) string {
    var v8 int = p1
    v9 := []int{v8}
    if true {
        p0 += p0
        p0 = 1
        if true {
            var v10 uint32 = p0
            var v11 uint8 = 0
            _ = v10
            _ = v11
        }
    }
    p1 -= 1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "a"
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -58
    var v3 string = "foo bar"
    v12 := GenS1{F0: -82, F1: 2, F2: 128}
    fmt.Println(v3, v2)
    switch ("" + v3) {
    case "a":
        v13 := []int{-871}
        _ = v13
    case "":
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 int64 = 9223372036854775807
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := GenS2{F0: -889}
                _ = fi16
                _ = v17
            }
            v18 := GenS2{F0: -32768}
            _ = fi14
            _ = v15
            _ = v18
        }
    default:
        v19 := v1[32768]
        _ = v19
    }
    v3 = "foo bar"
    v20 := GenS3{F0: v3, F1: 431}
    fmt.Println(v3, v2)
    v2 = v2
    if true {
        switch v3 {
        case "a":
            v21 := []int{}
            _ = v21
        case "hello":
            if false {
                v22 := GenS4{F0: false}
                for fi23 := 0; fi23 < 18; fi23++ {
                    v24 := GenS5{F0: 1, F1: v2}
                    _ = fi23
                    _ = v24
                }
                fmt.Println(v2)
                _ = v22
            }
        case "":
            v2 += 75
        default:
            fmt.Println(v3)
        }
    }
    fmt.Println(v2, v2)
    var v25 int32 = 42
    if (!(296 < -548)) {
        fmt.Println(v2)
    } else {
        switch 1 {
        case -1:
            var v26 int = -119
            _ = v26
        case 4294967295:
            fmt.Println(v3, v3, v3)
        default:
            for fi27 := 0; fi27 < 4; fi27++ {
                v25 -= genFn0(65535, 50)
                v3 = (((v3 + v3) + ((v3 + v3) + v3)) + (v3 + ("\n\t" + v3)))
                _ = fi27
            }
        }
        var v28 int8 = 42
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v25
    _ = v3
}
