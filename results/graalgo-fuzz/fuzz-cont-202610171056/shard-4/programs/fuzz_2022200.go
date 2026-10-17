package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32, p1 int64, p2 string) int {
    v7 := []int{192, -1}
    if (428 > 42) {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 int64 = 1
            switch fi8 {
            case 1:
                v10 := GenS1{F0: v9, F1: (p0 + p0), F2: 708}
                _ = v10
            default:
                fmt.Println(p2, p0)
            }
            _ = fi8
            _ = v9
        }
        p0 = p0
        v7 = append(v7, 2147483647)
    }
    if true {
        fmt.Println(p2, p1)
        v11 := GenS2{F0: "foo bar"}
        var v12 int8 = -89
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32768
}
func genFn1() string {
    v13 := GenS1{F0: 1, F1: 42, F2: 0}
    var v14 int16 = 1000
    var v15 int16 = v14
    switch "foo bar" {
    case "\n\t":
        v15 = -866
    case "":
        var v16 uint = 980
        _ = v16
    case "a":
        v14 = (v14 % ((42) | 1))
    default:
        var v17 int16 = (v15 << ((921) & 63))
        _ = v17
    }
    _ = v13
    _ = v14
    _ = v15
    return "hello"
}

func main() {
    var v1 string = ""
    var v2 uint8 = 42
    v3 := []uint32{}
    v4 := make(map[int]uint32)
    var v5 uint32 = 2
    var v6 uint16 = 0
    var v18 uint16 = v6
    var v19 uint16 = (v6 / ((v6) | 1))
    switch (v1 + "a") {
    case "foo bar":
        var v20 float64 = float64((1e10 * float32((1e10 - 3.14159))))
        _ = v20
    }
    var v21 uint8 = v2
    fmt.Println(v19, v2)
    var v22 float32 = 1e10
    var v23 int = 32768
    for fi24 := 0; fi24 < 10; fi24++ {
        for fi25 := 0; fi25 < 16; fi25++ {
            delete(v4, genFn0(v5, -1, ("hello" + (v1 + "\n\t"))))
            _ = fi25
        }
        v26 := GenS3{F0: ((v19 >= v6) || (v21 <= 97))}
        _ = fi24
        _ = v26
    }
    v27 := GenS1{F0: -79, F1: (^(v5)), F2: 255}
    var v28 uint64 = uint64(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
