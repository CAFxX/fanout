package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int32) int8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(p0)
        v8 := []int{fi7, (-(fi7)), (fi7 << ((255) & 63)), 42}
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 string = "hello"
            switch (fi9 * fi7) {
            case 9223372036854775807:
                switch ((v10 + (v10 + "a")) + (("hello" + v10) + ((((("a" + v10) + "hello") + v10) + "a") + v10))) {
                case "\n\t":
                    v10 = v10
                default:
                    v11 := GenS1{F0: 42, F1: ("hello" + "\n\t")}
                    _ = v11
                }
            case 36:
                switch "\n\t" {
                case "foo bar":
                    if true {
                        if (v10 <= ((("foo bar" + (v10 + ((v10 + "\n\t") + v10))) + "foo bar") + ((v10 + "foo bar") + v10))) {
                            _ = len(v8)
                            if (v10 <= "\n\t") {
                                var v12 uint8 = 42
                                var v13 string = v10
                                _ = v12
                                _ = v13
                            }
                            v10 = "foo bar"
                        } else {
                            p0 = p0
                        }
                        fmt.Println(v10, fi9, fi9)
                    } else {
                        if false {
                            var v14 int = fi9
                            v10 = (v10 + v10)
                            _ = v14
                        } else {
                            p0 = 0
                        }
                        v10 = ""
                        if len(v8) > 0 {
                            v8[0] = fi9
                        }
                    }
                case "hello":
                    fmt.Println(v10, v10, fi7)
                }
            default:
                var v15 int16 = 0
                _ = v15
            }
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    switch "a" {
    case "a":
        fmt.Println(p0)
    case "\n\t":
        v16 := []int{0, -1}
        _ = v16
    case "hello":
        var v17 float64 = (((0.1 * float64(float32(float64((1.5 + (((1e10 + 1.5) + (-2.25 + 1.5)) - 1e10)))))) + (-2.25 / (1e10 + 0.1))) / -2.25)
        _ = v17
    default:
        var v18 int32 = 2147483647
        _ = v18
    }
    _ = p0
    return -56
}

func main() {
    var v1 uint = 2
    var v2 uint = (v1 + v1)
    v3 := make(map[uint32]int)
    var v4 int32 = -469
    var v5 float64 = (3.14159 / ((1.5 * float64((0.1 - -2.25))) + 0.1))
    v6 := make(map[string]int)
    v4 += (v4 % ((v4) | 1))
    v19 := GenS2{F0: true, F1: 82}
    v20 := GenS2{F0: true, F1: 237}
    v4 = v4
    v2 = (v1 * v1)
    fmt.Println(v2, v2, v4)
    if true {
        var v21 int16 = 1
        _ = v21
    }
    v22 := []int{492, 0, 42, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
