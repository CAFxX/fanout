package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0(p0 bool, p1 int64) uint8 {
    fmt.Println(p0, p0)
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 619, F1: 870}
        _ = fi7
        _ = v8
    }
    p1 += (^(p1))
    p1 = p1
    switch -2147483648 {
    case -699:
        switch ("hello" + "a") {
        case "\n\t":
            var v9 uint = 339
            _ = v9
        case "hello":
            v10 := GenS1{F0: 2147483647, F1: 4294967295}
            _ = v10
        default:
            var v11 uint = 255
            _ = v11
        }
    case -9223372036854775808:
        if (p0 || p0) {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int16 = -1
                v14 := GenS1{F0: 734, F1: 214}
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := []int{}
            var v16 int8 = 100
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: 770, F1: 2}
            _ = v17
        }
    case 2147483647:
        v18 := GenS1{F0: 2147483647, F1: 0}
        _ = v18
    default:
        var v19 int = 9223372036854775807
        _ = v19
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 int, p1 int32, p2 int16) int8 {
    p1 = (p1 / ((42) | 1))
    v20 := GenS2{F0: -25, F1: 589}
    v21 := []int{}
    fmt.Println(p0, p0)
    if len(v21) > 0 {
        v21[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return -1
}

func main() {
    var v1 int64 = 238
    var v2 uint8 = 42
    v3 := make(map[uint32]uint32)
    var v4 string = ""
    var v5 uint32 = 1
    var v6 int32 = 551
    fmt.Println(v5)
    v22 := []int{23, 4294967295, 127, 128}
    delete(v3, 2147483647)
    v2 = genFn0(true, v1)
    v1 += v1
    var v23 uint = 2
    var v24 float64 = (1.5 * ((0.0 + float64((1e10 / (1e10 * 0.1)))) / (1e10 + 0.1)))
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
