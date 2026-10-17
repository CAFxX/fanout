package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 bool; F2 uint }
func genFn0() string {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint16 = 45
        if (!((180 != 1) && true)) {
            switch "\n\t" {
            case "foo bar":
                var v9 uint32 = 2
                _ = v9
            case "hello":
                v8 = 189
            default:
                if true {
                    v10 := []int{(fi7 / (((^(fi7))) | 1)), fi7, 0}
                    _ = len(v10)
                    _ = v10
                }
            }
            fmt.Println(fi7, fi7)
            v11 := []int{}
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v12 := []int{-235, 0}
    var v13 uint = 23
    _ = v12
    _ = v13
    return "foo bar"
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 int32 = -326
    var v5 float32 = 3.14159
    var v6 int16 = -1
    v4 = v4
    var v14 bool = (!(42 > 12))
    for fi15 := 0; fi15 < 14; fi15++ {
        v16 := []int{fi15}
        var v17 int64 = 9223372036854775807
        v17 = 914
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 uint8 = 2
    switch "a" {
    case "a":
        if false {
            v19 := GenS1{F0: v18}
            var v20 uint16 = 255
            v3[4294967295] = 2
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 19; fi21++ {
                fmt.Println(v14, v14, v6)
                _ = fi21
            }
        }
    case "hello":
        var v22 float32 = ((3.14159 * (float32((float64(0.1) / 3.14159)) / float32(1e10))) / v5)
        _ = v22
    case "foo bar":
        v23 := GenS1{F0: 0}
        _ = v23
    }
    v24 := GenS1{F0: v18}
    _ = len(v1)
    var v25 int8 = 0
    v26 := GenS2{F0: 4294967295, F1: v14, F2: 42}
    var v27 int8 = v25
    fmt.Println(v27)
    var v28 int32 = v4
    if true {
        var v29 int64 = 291
        _ = v29
    }
    if (!((v18 - v18) < 81)) {
        for fi30 := 0; fi30 < 8; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                v18 -= (v18 >> ((756) & 63))
                _ = fi31
            }
            _ = fi30
        }
        switch 683 {
        case 1:
            v32 := []int{805, 1, 0, 4294967295}
            _ = v32
        }
    }
    v14 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
