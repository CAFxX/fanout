package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 bool }
func genFn0(p0 uint64, p1 int8) uint16 {
    switch -358 {
    case -722:
        switch 0 {
        case 531:
            fmt.Println(p1, p1)
        case -910:
            p0 = (p0 + 42)
        case 9223372036854775807:
            fmt.Println(p0)
        }
    case -709:
        var v5 uint32 = 2147483647
        _ = v5
    }
    var v6 string = ("" + "")
    if (false || ((0.1 / (0.1 + 3.14159)) == -2.25)) {
        p0 += (^(p0))
        v7 := []int{680, 125, 128}
        _ = v7
    }
    if (255 == 135) {
        p1 = 127
        v8 := []int{-423, 4294967295}
        _ = v8
    } else {
        var v9 uint = 4294967295
        v10 := []int{4294967295}
        _ = v10
        _ = v9
    }
    if true {
        v11 := GenS1{F0: 42, F1: 1, F2: 598}
        v12 := GenS2{F0: 556}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 string = "\n\t"
    var v2 int = 127
    v3 := make(map[int]int)
    var v4 bool = false
    switch 657 {
    case 201:
        var v13 uint8 = 42
        _ = v13
    }
    var v14 int64 = 980
    v4 = (771 == 0)
    for fi15 := 0; fi15 < 16; fi15++ {
        v3[fi15] = 32767
        _ = fi15
    }
    var v16 int32 = 0
    if (-1 >= 118) {
        v4 = (998 != 42)
    }
    v17 := []int{v2, 802, (v2 - v2)}
    v18 := GenS3{F0: (!v4), F1: (false || (-973 > -266))}
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 int16 = -321
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v2 -= v2
    if v4 {
        var v22 int32 = (^(v16))
        var v23 int8 = -128
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = v22
        _ = v23
    } else {
        var v24 uint8 = 227
        _ = v24
    }
    var v25 float32 = float32((1e10 + (0.0 + -2.25)))
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 uint16 = 410
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 uint32 = 65535
            v25 = v25
            var v30 bool = (2147483647 < fi26)
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi26
        _ = v27
    }
    switch v2 {
    case 0:
        v4 = true
    default:
        if len(v17) > 0 {
            v17[0] = v2
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
