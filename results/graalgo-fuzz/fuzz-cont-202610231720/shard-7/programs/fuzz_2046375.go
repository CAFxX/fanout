package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint }

func main() {
    var v1 string = "\n\t"
    var v2 float32 = 0.0
    v3 := make(map[int]int)
    var v4 int32 = 42
    var v5 float32 = 1.5
    var v6 uint16 = 0
    v4 = v4
    var v7 int16 = -261
    if (-747 > 114) {
        v8 := GenS1{F0: 127, F1: -336, F2: int64(v7)}
        v1 = (v1 + "")
        v5 -= ((float32((0.1 - ((0.0 * 0.1) - (1.5 * 3.14159)))) + v2) + v5)
        _ = v8
    }
    switch -1 {
    case -73:
        v1 = v1
    case -878:
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint8 = 128
            v11 := GenS1{F0: 42, F1: (fi9 * 2147483647), F2: 9223372036854775807}
            v12 := GenS2{F0: 42, F1: v10, F2: 302}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    case 32768:
        if ((true && false) || (1.5 >= ((1.5 * 1e10) - 0.1))) {
            _ = len(v3)
            v6 -= 0
        }
    default:
        for fi13 := 0; fi13 < 2; fi13++ {
            if (!true) {
                v14 := []int{fi13}
                var v15 int32 = v4
                _ = len(v14)
                _ = v14
                _ = v15
            }
            _ = fi13
        }
    }
    v1 = v1
    var v16 float64 = 0.1
    if false {
        switch ("foo bar" + "") {
        case "\n\t":
            var v17 uint8 = 128
            _ = v17
        default:
            fmt.Println(v6, v1, v6)
        }
    } else {
        v18 := []int{}
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        v20 := GenS2{F0: 1, F1: 0, F2: 936}
        _ = fi19
        _ = v20
    }
    var v21 float32 = (((float32(((v16 / (v16 - v16)) - v16)) + v2) / float32(v16)) / v5)
    v16 -= float64(0.0)
    v1 = v1
    var v22 uint16 = (v6 & (v6 - (v6 + (^(v6)))))
    switch v1 {
    case "":
        fmt.Println(v5, v4)
    }
    var v23 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
