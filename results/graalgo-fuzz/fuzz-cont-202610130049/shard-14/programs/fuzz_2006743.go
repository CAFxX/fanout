package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int64 }
func genFn0(p0 uint8) int32 {
    fmt.Println(p0)
    p0 = p0
    v7 := []int{-1, 9223372036854775807, 0, -405}
    var v8 bool = true
    _ = p0
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 127
    var v3 int64 = 47
    var v4 float64 = (1e10 + (1.5 * 1e10))
    var v5 int32 = 1
    var v6 int32 = v5
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(v4, fi9, v3)
        var v10 uint64 = 0
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 0
    v4 += -2.25
    var v12 uint64 = 2
    if ((83 < 4294967295) || (false && false)) {
        v5 -= (-(v5))
    } else {
        v4 = ((((v4 + (v4 - v4)) - (v4 / (v4 / 1e10))) * v4) / float64(float32(v4)))
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        if false {
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 bool = true
                var v16 int8 = 127
                var v17 float32 = 1e10
                v3 += v3
                _ = len(v1)
                var v18 uint16 = v11
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 float64 = 0.1
            fmt.Println(v4)
            _ = fi19
            _ = v20
        }
        switch ((("" + "\n\t") + "hello") + "a") {
        case "hello":
            switch ("a" + ("foo bar" + "a")) {
            case "a":
                fmt.Println(v5, v2, v11)
            case "":
                var v21 string = "hello"
                _ = v21
            }
        }
        _ = fi13
    }
    switch "" {
    case "foo bar":
        var v22 int = -911
        _ = v22
    case "hello":
        v4 -= (0.0 / v4)
    default:
        fmt.Println(v6, v4)
    }
    var v23 bool = (32767 >= -1)
    v12 = v12
    delete(v1, 520)
    if (!(127 > -16)) {
        switch ("" + "foo bar") {
        case "hello":
            v24 := GenS1{F0: v11, F1: -736, F2: v3}
            _ = v24
        default:
            fmt.Println(v11, v4)
        }
        v25 := GenS1{F0: 255, F1: int(v3), F2: v3}
        _ = v25
    } else {
        var v26 float32 = 0.1
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
