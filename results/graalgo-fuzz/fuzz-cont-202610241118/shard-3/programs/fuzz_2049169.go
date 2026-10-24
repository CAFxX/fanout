package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 uint16, p1 string, p2 int16) string {
    fmt.Println(p1)
    var v7 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "foo bar"
}

func main() {
    var v1 uint8 = 129
    var v2 float32 = ((-2.25 - (1e10 + 1.5)) + (float32(0.1) * 1e10))
    var v3 bool = false
    var v4 float64 = (0.0 + -2.25)
    v5 := []string{"foo bar", "", "hello"}
    var v6 float32 = float32((v4 / 0.1))
    var v8 uint32 = 2147483648
    for fi9 := 0; fi9 < 13; fi9++ {
        if (0 >= 2) {
            if len(v5) > 0 {
                v5[0] = "a"
            }
            v1 = 80
        }
        var v10 uint64 = 42
        _ = fi9
        _ = v10
    }
    v1 += v1
    var v11 string = "a"
    if true {
        for fi12 := 0; fi12 < 15; fi12++ {
            v5 = append(v5, v11)
            _ = fi12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v4, v6)
            v4 -= ((v4 / v4) - 0.0)
            _ = fi13
        }
        fmt.Println(v2, v4, v8)
    } else {
        v1 = (^(v1))
        var v14 int8 = -128
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int32 = 0
            var v17 int16 = 0
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    switch -40 {
    case -1:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case -991:
        if false {
            var v18 int = -926
            _ = v18
        } else {
            var v19 uint8 = (v1 >> ((313) & 63))
            _ = v19
        }
    case 32768:
        var v20 uint = 300
        _ = v20
    }
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    var v21 uint = 4294967295
    var v22 int = -1
    var v23 uint = (v21 | (v21 * 18446744073709551615))
    switch v11 {
    case "a":
        v24 := GenS1{F0: -1, F1: -168}
        _ = v24
    default:
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 uint = (v21 << ((v21) & 63))
            _ = fi25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 float64 = float64((v2 / ((-2.25 + v6) - 3.14159)))
        var v29 uint64 = 2
        _ = fi27
        _ = v28
        _ = v29
    }
    v4 -= (float64(v2) - 1e10)
    fmt.Println(v23, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
