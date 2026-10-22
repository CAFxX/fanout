package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint) int8 {
    if (!(true && (127 <= 2))) {
        for fi5 := 0; fi5 < 1; fi5++ {
            p0 = 1
            var v6 float64 = float64((-2.25 - 0.1))
            _ = fi5
            _ = v6
        }
        var v7 float64 = -2.25
        var v8 int = 42
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint64 = 1
            v10 -= v10
            v10 += v10
            _ = fi9
            _ = v10
        }
        var v11 float32 = (1e10 - (((1e10 + 1.5) * 1e10) * float32(1e10)))
        _ = v11
    }
    var v12 int64 = 142
    v12 -= (v12 + 9223372036854775807)
    switch "" {
    case "\n\t":
        for fi13 := 0; fi13 < 14; fi13++ {
            if (v12 != (v12 * (v12 / ((v12) | 1)))) {
                v12 += (v12 >> ((p0) & 63))
            }
            if false {
                v14 := GenS1{F0: 128}
                fmt.Println(fi13)
                v12 = -234
                _ = v14
            }
            _ = fi13
        }
    }
    _ = p0
    _ = v12
    return 127
}
func genFn1(p0 int16, p1 uint, p2 int16) int16 {
    p1 += (p1 % ((782) | 1))
    p2 = p2
    if (!(true || false)) {
        if ((true || (-1 >= 0)) && (!false)) {
            var v15 string = ("\n\t" + "foo bar")
            _ = v15
        } else {
            p2 -= -149
            v16 := []int{}
            _ = v16
        }
        var v17 float32 = 0.0
        p1 = 42
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 & p2)
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    var v3 float64 = ((0.0 * ((0.0 * 0.1) * 3.14159)) / 1.0)
    v4 := []uint32{2147483648, 2147483647, 0, 462}
    v3 = 1e10
    v3 += (v3 / v3)
    if false {
        var v18 uint8 = 1
        var v19 int64 = 0
        fmt.Println(v19, v19, v18)
        _ = v18
        _ = v19
    } else {
        var v20 float32 = 0.1
        _ = v20
    }
    switch 4294967295 {
    case 2147483647:
        v3 += v3
    case 0:
        v3 += (((v3 + (v3 - v3)) * (v3 * v3)) + 1e10)
    default:
        for fi21 := 0; fi21 < 3; fi21++ {
            if (false && true) {
                fmt.Println(fi21)
            }
            _ = fi21
        }
    }
    var v22 uint64 = 0
    var v23 uint16 = 891
    var v24 int16 = -505
    fmt.Println(v24, v3, v22)
    fmt.Println(v23, v24, v3)
    if len(v1) > 0 {
        v1[0] = (("a" + "a") + "")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
