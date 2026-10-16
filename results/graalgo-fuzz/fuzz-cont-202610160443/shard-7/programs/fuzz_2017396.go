package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 bool }
type GenS2 struct { F0 int64; F1 bool }
func genFn0(p0 uint8, p1 uint, p2 int16) uint32 {
    switch 2147483647 {
    case 128:
        var v7 uint16 = 255
        _ = v7
    }
    v8 := []int{127, 2147483647}
    for fi9 := 0; fi9 < 1; fi9++ {
        v8 = append(v8, 4294967295)
        if len(v8) > 0 {
            v8[0] = (fi9 >> (((p1 / (((p1 * (p1 >> ((4294967295) & 63)))) | 1))) & 63))
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
    }
    if (2147483647 <= 155) {
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(p2)
            _ = fi10
        }
        var v11 string = ("a" + ("hello" + "\n\t"))
        _ = v11
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 557
}

func main() {
    var v1 int64 = -1
    var v2 int8 = 1
    v3 := make(map[string]uint32)
    var v4 uint = 1
    v5 := make(map[int]bool)
    var v6 int16 = 32767
    v1 -= 1
    var v12 uint64 = 0
    v4 += (v4 % ((v4) | 1))
    v4 += (v4 % ((v4) | 1))
    v13 := GenS1{F0: (v2 >> ((v4) & 63)), F1: (v1 & (^(v1))), F2: ("foo bar" < "foo bar")}
    v1 = (v1 + v1)
    if false {
        switch ("foo bar" + ("a" + "hello")) {
        case "a":
            var v14 uint = (v4 / ((v4) | 1))
            _ = v14
        default:
            for fi15 := 0; fi15 < 11; fi15++ {
                _ = len(v3)
                v1 = v1
                if (!(false && false)) {
                    v16 := v3[(("\n\t" + "a") + "hello")]
                    _ = v16
                } else {
                    v2 = v2
                    v1 -= v1
                }
                _ = fi15
            }
        }
        v4 -= (v4 >> ((v4) & 63))
        var v17 float64 = ((((0.1 + 0.0) + 0.0) + (3.14159 - 1e10)) - (0.1 / 3.14159))
        _ = v17
    }
    delete(v5, 4294967295)
    fmt.Println(v6, v4)
    var v18 float64 = (-2.25 * (1.5 + 0.0))
    v19 := GenS2{F0: v1, F1: true}
    var v20 int64 = (-(v1))
    v1 += v20
    var v21 uint = v4
    var v22 uint32 = 2147483648
    var v23 uint8 = 1
    switch ("a" + "a") {
    case "a":
        var v24 uint8 = 2
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
