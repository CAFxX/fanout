package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int32 = 797
    var v2 string = ("hello" + "hello")
    var v3 uint16 = 2
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    v6 := GenS1{F0: v3, F1: 42}
    var v7 uint32 = 1
    if (float64(0.1) == 1e10) {
        if ((!(!false)) || ((2147483647 <= (^(v7))) || true)) {
            fmt.Println(v3, v7)
        }
    }
    var v8 float64 = 0.0
    switch v2 {
    case "foo bar":
        if (133 >= -430) {
            if (1e10 > (1e10 * -2.25)) {
                var v9 float32 = float32(v8)
                _ = v9
            }
        }
    case "hello":
        for fi10 := 0; fi10 < 11; fi10++ {
            if ((0.1 - 1.5) < ((0.1 * -2.25) + -2.25)) {
                v11 := GenS1{F0: v3, F1: v7}
                _ = v11
            } else {
                fmt.Println(v7)
                var v12 uint16 = v3
                _ = v12
            }
            var v13 int32 = (v1 >> ((209) & 63))
            v2 = (("foo bar" + v2) + (v2 + ((v2 + ((v2 + (v2 + v2)) + v2)) + v2)))
            _ = fi10
            _ = v13
        }
    default:
        var v14 int8 = -1
        _ = v14
    }
    var v15 int16 = -375
    v15 -= 913
    v3 = v3
    v15 = v15
    var v16 uint8 = 127
    if true {
        v8 = 1e10
        v17 := GenS1{F0: (v3 & 255), F1: uint32(v3)}
        v18 := []int{32768, 4294967295}
        _ = v17
        _ = v18
    } else {
        v16 -= (v16 + 127)
        if (!false) {
            switch "" {
            case "\n\t":
                fmt.Println(v8, v7, v3)
            case "hello":
                v19 := []int{}
                _ = v19
            case "a":
                v16 = (^(v16))
            default:
                var v20 int16 = v15
                _ = v20
            }
        } else {
            for fi21 := 0; fi21 < 14; fi21++ {
                if (v16 <= (v16 >> ((18446744073709551615) & 63))) {
                    fmt.Println(v7, v8)
                    v22 := []int{fi21, (fi21 >> ((42) & 63)), (fi21 / ((fi21) | 1))}
                    _ = v22
                } else {
                    v15 = v15
                }
                for fi23 := 0; fi23 < 13; fi23++ {
                    v16 -= (v16 | v16)
                    v15 = v15
                    _ = fi23
                }
                _ = fi21
            }
            var v24 int16 = (v15 * v15)
            _ = v24
        }
        v25 := GenS2{F0: 2}
        _ = v25
    }
    if (v16 == 0) {
        fmt.Println(v8, v16, v15)
    }
    switch 352 {
    case 1:
        fmt.Println(v1, v7)
    case 322:
        if (!(v8 == float64(0.0))) {
            var v26 uint64 = 2
            _ = v26
        } else {
            v16 = (v16 / (((v16 + (v16 + 255))) | 1))
            v4[-2147483648] = ((v2 + (v2 + "foo bar")) < ((v2 + ("" + "foo bar")) + (("" + "hello") + "\n\t")))
        }
    case 390:
        v15 = v15
    default:
        v27 := []int{127}
        _ = v27
    }
    var v28 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
