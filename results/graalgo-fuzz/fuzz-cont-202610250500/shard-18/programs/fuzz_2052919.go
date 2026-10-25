package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int32; F2 uint32 }
func genFn0(p0 int8, p1 int32) uint32 {
    fmt.Println(p1, p1)
    v5 := []int{9223372036854775807}
    var v6 uint32 = 989
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint16) uint64 {
    v7 := []int{-1, 4294967295, 2147483647}
    switch ((("\n\t" + "") + "a") + "\n\t") {
    case "hello":
        var v8 uint8 = 249
        _ = v8
    case "foo bar":
        switch 905 {
        case 4294967295:
            p0 = (p0 - p0)
        case -9223372036854775808:
            v9 := GenS1{F0: -1}
            _ = v9
        case 907:
            p0 = (p0 - (p0 << ((4294967295) & 63)))
        default:
            var v10 uint32 = 753
            _ = v10
        }
    default:
        p0 -= p0
    }
    _ = p0
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 string = ("foo bar" + "hello")
    v3 := make(map[int]string)
    v4 := []int{-486, 42, 1}
    var v11 float32 = (0.0 / 1e10)
    _ = len(v4)
    if false {
        v11 = (float32(1e10) - (v11 * v11))
        v2 = v2
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint64 = 187
        _ = fi12
        _ = v13
    }
    v11 += (((0.1 * v11) - v11) / ((v11 * v11) * v11))
    var v14 int8 = -1
    var v15 float64 = 1e10
    var v16 float64 = (1.5 / v15)
    switch v2 {
    case "hello":
        if (-144 < -32768) {
            fmt.Println(v11, v16)
        } else {
            for fi17 := 0; fi17 < 5; fi17++ {
                v15 -= ((v16 + -2.25) - ((v15 / 1.5) / v15))
                var v18 float64 = v15
                for fi19 := 0; fi19 < 4; fi19++ {
                    fmt.Println(v2, fi17)
                    _ = fi19
                }
                _ = fi17
                _ = v18
            }
        }
    case "\n\t":
        var v20 string = v2
        _ = v20
    case "":
        var v21 uint64 = 42
        _ = v21
    default:
        v22 := GenS2{F0: -235, F1: 2147483647, F2: 635}
        _ = v22
    }
    v11 -= (0.1 * v11)
    switch -9223372036854775808 {
    case 329:
        var v23 string = v2
        _ = v23
    case 1:
        var v24 int = -2147483648
        _ = v24
    case 361:
        var v25 uint64 = 983
        _ = v25
    }
    v2 = v2
    switch ((v2 + ("" + v2)) + ("foo bar" + v2)) {
    case "hello":
        v15 += v15
    case "a":
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    default:
        v14 -= (v14 - v14)
    }
    fmt.Println(v11, v16, v2)
    delete(v3, -790)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
