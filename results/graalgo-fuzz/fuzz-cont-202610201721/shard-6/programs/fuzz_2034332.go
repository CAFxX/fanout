package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 bool; F1 string }

func main() {
    var v1 int16 = -32768
    v2 := []uint32{1, 2147483647, 521}
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    switch 1 {
    case 32767:
        v4[32767] = "a"
    default:
        v1 -= 383
    }
    var v5 string = "\n\t"
    switch (v5 + v5) {
    case "\n\t":
        var v6 int16 = (v1 & (v1 | (v1 * v1)))
        _ = v6
    case "hello":
        v1 += v1
    default:
        var v7 float64 = ((((1.5 / 1e10) / 1.0) / 1e10) / (0.0 + (float64(-2.25) * 1.5)))
        _ = v7
    }
    v1 += 32767
    fmt.Println(v1, v5)
    var v8 uint16 = 0
    switch -310 {
    case 2147483647:
        fmt.Println(v8)
    case 1:
        delete(v4, 0)
    case 42:
        var v9 int32 = 42
        _ = v9
    }
    v1 -= v1
    var v10 bool = ((true || (42 > 0)) || ((!(((!false) || true) && (v1 != v1))) || false))
    v5 = (((v5 + (v5 + v5)) + (((v5 + v5) + ("\n\t" + "foo bar")) + "foo bar")) + v5)
    switch v5 {
    case "a":
        if (!(((-2.25 - (float64((1e10 + 1e10)) + -2.25)) == 0.1) && v10)) {
            fmt.Println(v1, v10)
        }
    }
    v11 := GenS1{F0: (v8 & v8), F1: ""}
    switch -9223372036854775808 {
    case 32768:
        var v12 int8 = 100
        _ = v12
    case -937:
        v8 = 699
    default:
        v1 = v1
    }
    if true {
        v13 := v4[621]
        _ = v13
    } else {
        fmt.Println(v10)
        var v14 uint = 756
        _ = v14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        if (!v10) {
            var v16 int8 = 100
            if false {
                if len(v2) > 0 {
                    v2[0] = 48
                }
                var v17 uint16 = v8
                v16 += (v16 << ((42) & 63))
                _ = v17
            }
            _ = v16
        }
        v8 = (^(v8))
        _ = fi15
    }
    switch v5 {
    case "foo bar":
        v18 := GenS1{F0: v8, F1: v5}
        _ = v18
    case "a":
        v1 += (v1 / (((^(v1))) | 1))
    case "hello":
        v19 := GenS2{F0: v10, F1: (("foo bar" + v5) + "a")}
        _ = v19
    }
    var v20 uint32 = 2147483647
    delete(v3, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
