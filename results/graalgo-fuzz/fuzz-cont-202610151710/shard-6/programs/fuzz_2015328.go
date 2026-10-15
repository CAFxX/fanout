package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
type GenS2 struct { F0 int32; F1 bool; F2 string }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 bool; F1 int8; F2 uint64 }
func genFn0(p0 bool, p1 int) uint {
    var v4 int8 = 91
    for fi5 := 0; fi5 < 20; fi5++ {
        v4 += v4
        p1 = int(v4)
        fmt.Println(v4, p1)
        _ = fi5
    }
    v6 := GenS1{F0: 822, F1: "foo bar", F2: 721}
    var v7 uint = 816
    v4 = (v4 + 127)
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    _ = v7
    return (v7 + v7)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v8 int = 846
    var v9 uint64 = 382
    var v10 float64 = 3.14159
    fmt.Println(v8)
    switch "" {
    case "a":
        for fi11 := 0; fi11 < 8; fi11++ {
            v9 = v9
            _ = fi11
        }
    case "foo bar":
        v12 := []int{(v8 + (v8 + (v8 >> ((genFn0((!(0 >= 9223372036854775807)), 616)) & 63)))), 902}
        _ = v12
    case "":
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 uint32 = 1
            _ = fi13
            _ = v14
        }
    default:
        var v15 uint8 = 1
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        if true {
            v17 := GenS2{F0: -577, F1: (9223372036854775807 <= -1), F2: "foo bar"}
            v18 := GenS2{F0: 1, F1: false, F2: "foo bar"}
            delete(v1, (fi16 + (v8 >> ((genFn0((1.5 > -2.25), (fi16 * (v8 >> ((1) & 63))))) & 63))))
            _ = v17
            _ = v18
        } else {
            v9 += (v9 >> ((255) & 63))
        }
        _ = fi16
    }
    var v19 uint8 = 127
    v20 := GenS1{F0: -2147483648, F1: "hello", F2: 844}
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            switch "foo bar" {
            case "a":
                v10 = ((v10 / float64(1.5)) * v10)
            case "foo bar":
                v9 -= (v9 - (v9 + 473))
            }
            if (!false) {
                v23 := GenS3{F0: 42, F1: 424}
                _ = v23
            }
            fmt.Println(v19, v9, v8)
            _ = fi22
        }
        v24 := []int{fi21, (fi21 % (((fi21 | fi21)) | 1)), v8, -457}
        _ = fi21
        _ = v24
    }
    v25 := []int{(v8 * (v8 & (^(v8))))}
    v26 := GenS4{F0: true, F1: 35, F2: (^(v9))}
    if len(v25) > 0 {
        v25[0] = 32768
    }
    switch "\n\t" {
    case "a":
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 13; fi28++ {
                v19 -= (v19 << ((4294967295) & 63))
                _ = fi28
            }
            _ = fi27
        }
    case "":
        for fi29 := 0; fi29 < 3; fi29++ {
            v10 += v10
            fmt.Println(v10)
            fmt.Println(v10, v19, v9)
            _ = fi29
        }
    }
    if (float32(v10) != (float32((1.5 - 0.1)) * 1.5)) {
        for fi30 := 0; fi30 < 3; fi30++ {
            v9 = v9
            for fi31 := 0; fi31 < 16; fi31++ {
                v9 = v9
                _ = fi31
            }
            v10 = -2.25
            _ = fi30
        }
        var v32 uint = genFn0((!true), (v8 + v8))
        _ = v32
    }
    if true {
        fmt.Println(v9, v10, v10)
        var v33 float64 = (v10 + (v10 + v10))
        v8 = 350
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v3
    _ = v8
    _ = v9
}
