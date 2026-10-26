package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 int16; F1 uint8; F2 int }

func main() {
    var v1 uint32 = 0
    var v2 float64 = ((0.0 + 1.5) - float64(1.5))
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2, v1}
    v5 := []string{"hello", "hello"}
    var v6 bool = (-2.25 == float32(-2.25))
    if true {
        fmt.Println(v1, v6)
    }
    if len(v5) > 0 {
        v5[0] = ("\n\t" + ("hello" + "a"))
    }
    v1 = 42
    v1 -= (v1 >> ((507) & 63))
    var v7 int = 127
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 0, F1: 42, F2: 4294967295}
        v6 = (!((((3.14159 * float32((float64(-2.25) - v2))) / 1.0) * 0.0) < 0.1))
        v4 = append(v4, v1)
        _ = fi8
        _ = v9
    }
    fmt.Println(v2, v7)
    switch ("a" + "foo bar") {
    case "hello":
        if len(v5) > 0 {
            v5[0] = "a"
        }
    }
    _ = len(v5)
    v2 = v2
    var v10 float64 = ((3.14159 - v2) - 3.14159)
    v11 := GenS2{F0: -250, F1: 128, F2: (v7 | int(v1))}
    v6 = false
    if (v10 >= (v2 / (-2.25 - (-2.25 + v2)))) {
        var v12 uint8 = 42
        for fi13 := 0; fi13 < 6; fi13++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi13
        }
        for fi14 := 0; fi14 < 8; fi14++ {
            switch "a" {
            case "foo bar":
                var v15 int8 = -91
                _ = v15
            case "\n\t":
                var v16 int16 = 389
                _ = v16
            default:
                fmt.Println(fi14, fi14)
            }
            if (-2.25 != 1.5) {
                var v17 uint = 4294967295
                _ = v17
            }
            var v18 int64 = -806
            _ = fi14
            _ = v18
        }
        _ = v12
    }
    var v19 bool = v6
    for fi20 := 0; fi20 < 15; fi20++ {
        v10 = v10
        v6 = (v19 && true)
        v19 = (0.0 != ((float32(float64(0.1)) / ((3.14159 * 3.14159) * -2.25)) - 0.0))
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
