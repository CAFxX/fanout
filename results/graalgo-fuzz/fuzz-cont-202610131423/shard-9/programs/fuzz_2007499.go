package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 int64, p1 int16, p2 int32) bool {
    if false {
        var v4 bool = false
        v5 := GenS1{F0: true, F1: "foo bar"}
        var v6 float64 = 3.14159
        _ = v4
        _ = v5
        _ = v6
    } else {
        p2 -= 42
        switch 42 {
        case 9223372036854775807:
            v7 := GenS1{F0: (!(!false)), F1: "foo bar"}
            _ = v7
        default:
            v8 := GenS2{F0: 100, F1: 525}
            _ = v8
        }
        p0 += 42
    }
    v9 := GenS1{F0: (true && false), F1: ""}
    var v10 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return true
}
func genFn1(p0 int16) int {
    p0 = p0
    fmt.Println(p0)
    var v11 uint32 = 480
    _ = p0
    _ = v11
    return 0
}

func main() {
    var v1 int16 = 42
    var v2 int16 = v1
    var v3 int64 = 0
    for fi12 := 0; fi12 < 7; fi12++ {
        v1 += (-(v2))
        _ = fi12
    }
    var v13 int32 = int32(v2)
    var v14 uint16 = 65535
    var v15 uint16 = (v14 / ((255) | 1))
    v16 := []int{-9223372036854775808, 128, 955}
    var v17 uint64 = 642
    var v18 uint = 255
    for fi19 := 0; fi19 < 6; fi19++ {
        _ = len(v16)
        _ = fi19
    }
    switch -2147483648 {
    case -1:
        v20 := GenS3{F0: 2147483647}
        _ = v20
    case 32768:
        switch 32767 {
        case 0:
            v17 += (v17 * v17)
        case 1:
            v17 -= v17
        case -1:
            v21 := GenS2{F0: 20, F1: (v13 % ((v13) | 1))}
            _ = v21
        default:
            v18 = v18
        }
    case 32767:
        if len(v16) > 0 {
            v16[0] = 32768
        }
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint8 = 128
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    var v25 int16 = -47
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
}
