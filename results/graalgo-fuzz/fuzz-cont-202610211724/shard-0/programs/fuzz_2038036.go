package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint }
func genFn0(p0 int8, p1 bool, p2 int32) int {
    p1 = p1
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int32 = p2
        var v7 int = fi5
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 32768}
            _ = fi8
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1() bool {
    var v10 int32 = 2147483647
    switch 404 {
    case 127:
        var v11 float64 = float64(-2.25)
        _ = v11
    case 0:
        v10 -= 0
    default:
        var v12 int16 = -377
        _ = v12
    }
    _ = v10
    return ((1 >= 69) && false)
}

func main() {
    var v1 int64 = 0
    var v2 uint = 255
    v3 := make(map[string]bool)
    var v4 int32 = -2147483648
    v13 := GenS1{F0: 128}
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int16 = -423
        v15 = (v15 + 42)
        v16 := GenS2{F0: 2}
        _ = fi14
        _ = v15
        _ = v16
    }
    v1 = (v1 << ((v2) & 63))
    var v17 int8 = int8(v4)
    var v18 uint64 = 671
    var v19 uint = 18446744073709551615
    switch 128 {
    case 439:
        fmt.Println(v1)
    case -214:
        var v20 bool = true
        _ = v20
    case 720:
        fmt.Println(v17, v2)
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        v17 = v17
        fmt.Println(v4, v2)
        _ = fi21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 string = "foo bar"
        var v24 uint8 = 2
        v25 := GenS3{F0: 67, F1: 9, F2: v2}
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
