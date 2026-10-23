package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int }

func main() {
    var v1 int8 = 108
    var v2 float32 = (0.0 + 1.5)
    var v3 uint = 42
    var v4 int16 = -32768
    var v5 uint8 = 42
    var v6 bool = true
    if (v5 >= 134) {
        v1 -= int8(v5)
        v7 := []int{4294967295, -127, 32768}
        if (((v6 && (((0.1 * 0.0) < 3.14159) && (uint16(v5) < 0))) && false) || ((1.5 + 0.0) > (v2 * (v2 / v2)))) {
            var v8 float32 = ((v2 / ((float32((1e10 - -2.25)) - 1e10) + (1e10 - 0.0))) - 0.1)
            var v9 uint32 = 2
            _ = len(v7)
            _ = v8
            _ = v9
        } else {
            if (!(42 >= uint16(v3))) {
                var v10 uint64 = 1
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v10
            }
        }
        _ = v7
    } else {
        if v6 {
            v11 := []int{}
            switch int(v3) {
            case -1:
                var v12 int = 481
                _ = v12
            default:
                if len(v11) > 0 {
                    v11[0] = 895
                }
            }
            _ = v11
        } else {
            v6 = v6
            v6 = true
        }
    }
    var v13 bool = ("foo bar" < "")
    var v14 uint8 = (v5 * 204)
    var v15 int = 0
    if v13 {
        v14 -= 148
    }
    v15 += v15
    v16 := []int{v15}
    v17 := GenS1{F0: -1, F1: ((v4 - (v4 + v4)) != (v4 ^ (v4 * 1)))}
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 uint32 = 213
        for fi20 := 0; fi20 < 6; fi20++ {
            fmt.Println(v13, v19, v3)
            var v21 float32 = (v2 * 3.14159)
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    var v22 string = (("foo bar" + (("\n\t" + "") + "foo bar")) + "a")
    v23 := GenS1{F0: v4, F1: (v3 == v3)}
    v24 := GenS2{F0: v15}
    v6 = (((v3 | v3) == v3) && (!(450 > 2147483647)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
