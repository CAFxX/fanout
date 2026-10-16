package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 int = -9223372036854775808
    var v4 int64 = 268
    v4 = v4
    for fi5 := 0; fi5 < 16; fi5++ {
        v4 = (v4 % ((-9223372036854775808) | 1))
        _ = fi5
    }
    v6 := GenS1{F0: 0, F1: true}
    var v7 int32 = 1
    if (!false) {
        var v8 string = (((("a" + "hello") + (("foo bar" + "\n\t") + "hello")) + "foo bar") + "hello")
        v9 := GenS1{F0: 0, F1: (520 == 670)}
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 bool = (1 > 127)
            if ((!v11) || v11) {
                var v12 uint16 = 411
                for fi13 := 0; fi13 < 7; fi13++ {
                    delete(v2, 2147483647)
                    if (2 == 255) {
                        var v14 float64 = ((0.0 + 1.5) * 1e10)
                        _ = v14
                    }
                    switch (-(v3)) {
                    case 0:
                        if v11 {
                            switch 4294967295 {
                            case 42:
                                var v15 string = (v8 + v8)
                                _ = v15
                            }
                            fmt.Println(v3)
                            fmt.Println(v4)
                        }
                    case 42:
                        v8 = (v8 + "")
                    }
                    var v16 bool = v11
                    _ = fi13
                    _ = v16
                }
                _ = v12
            }
            switch ((("" + "") + v8) + v8) {
            case "hello":
                var v17 int32 = v7
                _ = v17
            case "\n\t":
                var v18 float32 = -2.25
                _ = v18
            default:
                var v19 bool = v11
                _ = v19
            }
            _ = fi10
            _ = v11
        }
        _ = v8
        _ = v9
    }
    v7 = 1
    fmt.Println(v4, v4)
    switch "a" {
    case "a":
        v20 := GenS1{F0: -1, F1: (v3 != (-(v3)))}
        _ = v20
    case "hello":
        for fi21 := 0; fi21 < 14; fi21++ {
            v7 -= 0
            if (1.5 == 0.1) {
                var v22 string = "hello"
                v4 -= 746
                v3 -= 32768
                _ = v22
            }
            _ = fi21
        }
    }
    v7 -= (v7 - 42)
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint16 = 255
        if false {
            fmt.Println(fi23, v24)
        } else {
            switch (fi23 * (^(fi23))) {
            case 2147483647:
                v25 := GenS1{F0: v7, F1: true}
                _ = v25
            default:
                var v26 uint8 = 147
                _ = v26
            }
        }
        _ = fi23
        _ = v24
    }
    v7 = (v7 & (v7 << ((9) & 63)))
    fmt.Println(v3)
    v3 -= 619
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
