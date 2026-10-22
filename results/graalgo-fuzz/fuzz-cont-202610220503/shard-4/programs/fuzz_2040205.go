package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 uint64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("" + ("a" + "\n\t"))}
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 uint = 255
    fmt.Println(v5)
    _ = len(v2)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: (^(v5))}
        var v8 float32 = ((3.14159 / 1.0) * 3.14159)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 string = "a"
    var v10 int32 = 1
    fmt.Println(v10, v9, v9)
    if true {
        v11 := GenS1{F0: 255}
        var v12 int16 = 418
        _ = v11
        _ = v12
    }
    _ = len(v3)
    var v13 uint32 = 2147483647
    v9 = ((v9 + "a") + v9)
    v13 -= v13
    for fi14 := 0; fi14 < 13; fi14++ {
        switch v9 {
        case "a":
            v10 += v10
        case "":
            if false {
                var v15 float32 = (3.14159 + -2.25)
                v13 += v13
                var v16 int16 = -714
                _ = v15
                _ = v16
            }
        case "foo bar":
            var v17 int = (fi14 % (((fi14 ^ fi14)) | 1))
            _ = v17
        default:
            v5 -= (v5 << ((v5) & 63))
        }
        v9 = "a"
        _ = fi14
    }
    if true {
        var v18 uint32 = (v13 ^ (v13 + 38))
        _ = v18
    }
    if (42 >= 1) {
        v19 := GenS1{F0: (v5 + v5)}
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 17; fi21++ {
                var v22 int32 = (v10 | v10)
                _ = len(v4)
                v23 := GenS2{F0: 0, F1: v5, F2: uint64(v13)}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        v13 = v13
        _ = v19
    }
    v10 -= (v10 + (v10 ^ v10))
    var v24 uint = v5
    if (1 >= v10) {
        var v25 int8 = -1
        _ = v25
    } else {
        var v26 uint16 = 477
        for fi27 := 0; fi27 < 1; fi27++ {
            v5 += (v24 / (((^(v5))) | 1))
            v10 = (v10 << ((v24) & 63))
            var v28 uint8 = uint8(v13)
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
