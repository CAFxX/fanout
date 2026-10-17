package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 bool }
type GenS3 struct { F0 int8; F1 uint16; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 int; F1 string }
func genFn0(p0 uint8, p1 uint, p2 string) string {
    var v5 int = 976
    var v6 bool = (-32 >= 61)
    var v7 string = "foo bar"
    if (0 > 2147483648) {
        v8 := GenS1{F0: -1, F1: (((v7 + "foo bar") + "\n\t") + (v7 + v7)), F2: 1}
        v9 := GenS2{F0: 2, F1: 1, F2: (v7 >= (p2 + (p2 + p2)))}
        fmt.Println(p2)
        _ = v8
        _ = v9
    } else {
        var v10 int8 = 127
        var v11 uint8 = 0
        _ = v10
        _ = v11
    }
    var v12 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v5
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 uint16) uint64 {
    if true {
        for fi13 := 0; fi13 < 12; fi13++ {
            v14 := GenS3{F0: 42, F1: (^(p0)), F2: true}
            fmt.Println(p0, p0)
            _ = fi13
            _ = v14
        }
    } else {
        switch int(p0) {
        case 549:
            if false {
                p0 -= p0
            } else {
                var v15 int32 = 1
                var v16 int16 = -1
                _ = v15
                _ = v16
            }
        case -1:
            switch "foo bar" {
            case "":
                p0 = p0
            case "hello":
                v17 := GenS4{F0: 2147483647}
                _ = v17
            }
        case 127:
            p0 += (p0 >> ((21) & 63))
        default:
            v18 := GenS4{F0: -2147483648}
            _ = v18
        }
    }
    v19 := GenS5{F0: -208, F1: 2}
    _ = p0
    _ = v19
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 2147483648
    var v3 uint8 = 183
    var v4 bool = false
    switch 0 {
    case -1:
        var v20 int32 = 813
        _ = v20
    case 1:
        v21 := GenS6{F0: 1, F1: (v3 - 0)}
        _ = v21
    }
    var v22 bool = ((!false) && v4)
    fmt.Println(v3, v4, v4)
    v23 := []int{-702, 1, 208}
    fmt.Println(v22, v2)
    var v24 uint32 = (^(v2))
    v3 += v3
    v2 = (v24 << ((255) & 63))
    v25 := v1[2147483648]
    if true {
        v26 := GenS6{F0: -128, F1: v3}
        for fi27 := 0; fi27 < 2; fi27++ {
            fmt.Println(v4, fi27)
            for fi28 := 0; fi28 < 10; fi28++ {
                v1[v24] = (v25 + "foo bar")
                v25 = v25
                var v29 int32 = -14
                _ = fi28
                _ = v29
            }
            _ = fi27
        }
        var v30 int32 = 555
        _ = v26
        _ = v30
    }
    v22 = false
    v31 := GenS7{F0: 127, F1: v25}
    v2 -= 42
    fmt.Println(v22, v22)
    if false {
        for fi32 := 0; fi32 < 7; fi32++ {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            v3 = (v3 * v3)
            _ = fi32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
}
