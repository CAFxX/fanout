package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 bool) bool {
    switch -2147483648 {
    case 32:
        p0 = (255 == 74)
    case 42:
        for fi5 := 0; fi5 < 20; fi5++ {
            p0 = true
            for fi6 := 0; fi6 < 7; fi6++ {
                fmt.Println(p0, p0)
                fmt.Println(fi5, fi6, fi5)
                _ = fi6
            }
            _ = fi5
        }
    case 591:
        var v7 uint64 = 1
        _ = v7
    }
    v8 := GenS1{F0: 161, F1: p0}
    v9 := GenS2{F0: (!p0)}
    var v10 int32 = -963
    var v11 uint64 = 0
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return ((!(p0 || (-32768 >= 608))) || true)
}
func genFn1(p0 int, p1 uint, p2 bool) bool {
    v12 := GenS2{F0: p2}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return ((0.1 - 1e10) >= ((0.0 / (0.1 - float32(1e10))) - 3.14159))
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 9223372036854775807
    var v3 bool = false
    var v4 int8 = 111
    v2 = -1
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int8 = 38
        v2 = v2
        for fi15 := 0; fi15 < 6; fi15++ {
            v14 = v14
            if v3 {
                switch "hello" {
                case "\n\t":
                    var v16 float32 = 0.0
                    _ = v16
                case "a":
                    var v17 float64 = 1e10
                    _ = v17
                }
            }
            fmt.Println(v4, v2)
            v1[162] = "foo bar"
            v14 -= v14
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    v18 := []int{32767}
    switch 1 {
    case 1:
        v3 = v3
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v2 += (-(v2))
        fmt.Println(v4, fi19, v3)
        _ = fi19
    }
    v4 = (v4 | 0)
    var v20 bool = ((float32(1.5) == (((((float32(0.0) / 1.0) * 0.0) - float32((float64(0.0) - -2.25))) + (-2.25 / 1.0)) / (3.14159 - 0.1))) || true)
    v21 := GenS3{F0: 42, F1: 371}
    var v22 float64 = 1e10
    var v23 float32 = 1e10
    for fi24 := 0; fi24 < 17; fi24++ {
        fmt.Println(v3)
        switch ("foo bar" + ("a" + ("hello" + "\n\t"))) {
        case "hello":
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 int16 = 1000
                var v27 float64 = v22
                v28 := GenS1{F0: 255, F1: (((159 > 193) || v20) && (false || (v3 && v20)))}
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        default:
            v2 -= (v2 % ((v2) | 1))
        }
        _ = fi24
    }
    if len(v18) > 0 {
        v18[0] = 32767
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v29 float64 = float64(3.14159)
    for fi30 := 0; fi30 < 20; fi30++ {
        fmt.Println(v20)
        _ = fi30
    }
    var v31 int64 = (v2 >> ((42) & 63))
    var v32 int = 9223372036854775807
    var v33 float32 = (v23 + 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
