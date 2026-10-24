package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 uint16 = 414
    var v2 uint8 = 2
    var v3 string = ""
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    v1 -= 1
    if true {
        for fi6 := 0; fi6 < 10; fi6++ {
            var v7 uint64 = 507
            var v8 float32 = 0.1
            if ((false && true) && false) {
                var v9 uint8 = 128
                v10 := v4["a"]
                _ = v10
                _ = v9
            } else {
                v11 := GenS1{F0: 1, F1: "foo bar"}
                _ = v11
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        var v12 float32 = float32(float64((1e10 * -2.25)))
        _ = v12
    } else {
        v3 = v3
        if true {
            v13 := GenS1{F0: 42, F1: v3}
            v2 = 255
            _ = v13
        }
        var v14 uint = 1
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint = 42
        v2 = (v2 % ((v2) | 1))
        delete(v4, v3)
        _ = fi15
        _ = v16
    }
    _ = len(v5)
    _ = len(v5)
    switch 42 {
    case 127:
        v2 -= (v2 % ((v2) | 1))
    case 32768:
        delete(v5, 1)
    case 2147483647:
        v3 = v3
    }
    switch v3 {
    case "a":
        v3 = ((v3 + (v3 + ((v3 + v3) + v3))) + "hello")
    case "foo bar":
        var v17 uint32 = 2147483648
        _ = v17
    case "hello":
        var v18 int = 721
        _ = v18
    }
    fmt.Println(v1, v1, v2)
    v19 := []int{-9223372036854775808, 2147483647, 128}
    _ = len(v19)
    var v20 float32 = 1.5
    fmt.Println(v3, v1, v1)
    v21 := GenS2{F0: 0, F1: v3}
    if false {
        v20 = 0.1
        v2 = (^(v2))
        v22 := GenS2{F0: 0, F1: v3}
        _ = v22
    }
    var v23 uint16 = v1
    var v24 bool = (!(-2147483648 <= 32768))
    v25 := v4[("hello" + "a")]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
