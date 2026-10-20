package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16; F1 int }
func genFn0(p0 uint32, p1 int16, p2 uint64) int16 {
    if true {
        fmt.Println(p1)
        var v6 int = 1
        var v7 uint32 = (p0 * (p0 >> ((65535) & 63)))
        _ = v6
        _ = v7
    }
    var v8 uint = 287
    fmt.Println(p0, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p1
}

func main() {
    var v1 int16 = 1
    var v2 uint = 599
    var v3 uint = (v2 % ((v2) | 1))
    var v4 string = ("hello" + ("" + "a"))
    var v5 float32 = float32(3.14159)
    v1 += v1
    v2 -= v2
    if (false || (true || true)) {
        v9 := []int{-188, 922}
        v10 := GenS1{F0: (9223372036854775807 == 433), F1: v4}
        switch ((v4 + v4) + v4) {
        case "foo bar":
            var v11 float64 = (0.0 + (3.14159 / float64(v5)))
            _ = v11
        case "":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        case "a":
            switch "foo bar" {
            case "hello":
                fmt.Println(v4, v5)
            }
        default:
            if len(v9) > 0 {
                v9[0] = -5
            }
        }
        _ = v10
        _ = v9
    }
    v12 := []int{-307, 1, -9223372036854775808, 128}
    switch (v4 + (v4 + "foo bar")) {
    case "foo bar":
        v2 = (v3 - 78)
    default:
        for fi13 := 0; fi13 < 1; fi13++ {
            if (!true) {
                v12 = append(v12, (fi13 * (fi13 | -1)))
                var v14 uint64 = 42
                _ = v14
            } else {
                var v15 int16 = v1
                v12 = append(v12, 128)
                v4 = v4
                _ = v15
            }
            var v16 int32 = -1
            var v17 int64 = -4
            _ = fi13
            _ = v16
            _ = v17
        }
    }
    v18 := GenS2{F0: (v1 | -699), F1: 9223372036854775807}
    fmt.Println(v3, v5, v4)
    v19 := GenS2{F0: v1, F1: -838}
    if len(v12) > 0 {
        v12[0] = -1
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        v4 = (v4 + "")
        v21 := GenS1{F0: true, F1: "\n\t"}
        fmt.Println(v3, fi20)
        _ = fi20
        _ = v21
    }
    var v22 bool = true
    v12 = append(v12, 351)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
