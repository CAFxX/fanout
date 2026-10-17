package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
type GenS3 struct { F0 int16; F1 uint16; F2 int8 }
func genFn0(p0 uint16) bool {
    var v4 uint32 = 2147483647
    if true {
        fmt.Println(p0)
    }
    var v5 int = -2147483648
    _ = p0
    _ = v4
    _ = v5
    return true
}

func main() {
    var v1 int8 = 100
    v2 := make(map[int]bool)
    var v3 uint = 18446744073709551615
    v2[0] = (false && false)
    var v6 float64 = -2.25
    var v7 int8 = v1
    if (false && true) {
        fmt.Println(v1)
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v6 = ((float64((0.0 / (-2.25 + -2.25))) + 3.14159) - float64(0.0))
        var v9 uint = 653
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: uint16(v3)}
    var v11 int = -83
    switch -402 {
    case 9223372036854775807:
        v2[(v11 | 181)] = (-987 < -512)
    default:
        var v12 string = ""
        _ = v12
    }
    v3 += v3
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 string = "a"
        v7 -= v7
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint32 = 2
            v7 = 42
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v7)
    v17 := []int{}
    switch ("foo bar" + "") {
    case "":
        v18 := GenS2{F0: v3, F1: 829, F2: (!((!false) && (false || false)))}
        _ = v18
    case "foo bar":
        var v19 uint8 = 25
        _ = v19
    }
    var v20 int = (v11 << ((v3) & 63))
    v21 := GenS3{F0: -1, F1: 2, F2: -128}
    var v22 int8 = -57
    v23 := GenS2{F0: (v3 + 18446744073709551615), F1: 336, F2: ((338 == 0) || false)}
    delete(v2, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v6
    _ = v7
}
