package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint64 = 2
    var v2 bool = (42 > 169)
    var v3 bool = (true || ((0 >= -32768) && (65535 >= 65535)))
    var v4 uint32 = 410
    v5 := make(map[int]uint32)
    switch "a" {
    case "foo bar":
        var v6 string = (("\n\t" + "a") + (("a" + "a") + (("hello" + ("hello" + (("\n\t" + (("a" + "foo bar") + "hello")) + "foo bar"))) + "\n\t")))
        _ = v6
    default:
        var v7 float64 = float64(((1.5 / (-2.25 / 1.0)) - ((0.0 + (1e10 / float32(-2.25))) / 1.5)))
        _ = v7
    }
    var v8 uint8 = 1
    if v3 {
        if ((v1 % (((v1 << ((825) & 63))) | 1)) < 4294967295) {
            var v9 float32 = 3.14159
            _ = v9
        } else {
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: 2147483647, F1: v4, F2: (v8 << ((0) & 63))}
                switch ((("hello" + "hello") + ("" + "foo bar")) + ("foo bar" + "")) {
                case "hello":
                    v3 = (0 != 0)
                case "\n\t":
                    v8 = v8
                case "":
                    v2 = true
                }
                v5[(fi10 + fi10)] = v4
                _ = fi10
                _ = v11
            }
            var v12 int16 = 0
            _ = v12
        }
        v13 := GenS1{F0: 2, F1: 42, F2: v8}
        _ = v13
    }
    var v14 int = 2147483647
    var v15 int32 = int32(v4)
    v1 = v1
    v16 := GenS2{F0: -435}
    for fi17 := 0; fi17 < 1; fi17++ {
        v14 -= 603
        switch "a" {
        case "foo bar":
            var v18 string = ""
            _ = v18
        case "":
            v5[v14] = v4
        }
        _ = fi17
    }
    var v19 float64 = 1.5
    v14 = 127
    var v20 int8 = 31
    v19 = (1e10 + (v19 * 3.14159))
    v21 := v5[v14]
    v22 := GenS1{F0: 634, F1: 329, F2: uint8(v4)}
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(v21, v15)
        if true {
            v24 := []int{v14, (v14 >> ((uint(v8)) & 63)), v14, (v14 - (fi23 * (v14 - 2147483647)))}
            _ = len(v24)
            _ = v24
        } else {
            fmt.Println(v1, v2)
        }
        for fi25 := 0; fi25 < 6; fi25++ {
            v19 = v19
            v26 := GenS2{F0: fi23}
            v20 -= v20
            _ = fi25
            _ = v26
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
