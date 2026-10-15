package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint }
func genFn0() string {
    v7 := []int{9223372036854775807, 9223372036854775807}
    if (654 >= -812) {
        var v8 int32 = -2147483648
        _ = v8
    } else {
        var v9 uint8 = 127
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 string = (((("" + "foo bar") + "hello") + "") + "foo bar")
            var v12 int64 = 956
            if true {
                switch (("" + v11) + "") {
                case "":
                    v9 = v9
                default:
                    if false {
                        var v13 float32 = (-2.25 - (float32(3.14159) - float32((0.0 + 0.1))))
                        fmt.Println(v12, v13)
                        v9 = v9
                        _ = v13
                    } else {
                        v9 = v9
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        v12 = 9223372036854775807
                    }
                }
            } else {
                var v14 string = ("\n\t" + "a")
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = v9
    }
    _ = v7
    return ""
}
func genFn1() int64 {
    var v15 int32 = -425
    v15 = (v15 << ((149) & 63))
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(v15, v15)
        v15 = (v15 % ((-663) | 1))
        var v17 int = fi16
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 2147483647
    _ = v15
    _ = v18
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1e10
    v3 := make(map[string]uint32)
    var v4 uint32 = 1
    var v5 uint = 255
    var v6 int64 = -165
    fmt.Println(v6, v5, v6)
    switch -1 {
    case 32767:
        switch 2147483647 {
        case 9223372036854775807:
            var v19 bool = (2 <= 76)
            _ = v19
        }
    case 385:
        v20 := GenS1{F0: (v6 >> ((65535) & 63)), F1: 104}
        _ = v20
    }
    switch "\n\t" {
    case "":
        fmt.Println(v4, v4)
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        v22 := GenS1{F0: (v6 - (^(v6))), F1: 0}
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: (v6 << ((v5) & 63)), F1: 42}
    if true {
        v5 += v5
    } else {
        fmt.Println(v6, v2)
        if false {
            var v24 int32 = 2147483647
            _ = v24
        }
        var v25 bool = (-2.25 > (3.14159 / (((0.1 * -2.25) + (float64(v2) + 1.5)) - 1.5)))
        _ = v25
    }
    v6 = 9223372036854775807
    v6 -= 0
    v26 := []int{4294967295}
    if (true || (true && (!true))) {
        if false {
            var v27 float64 = -2.25
            v5 = (v5 >> (((v5 >> (((v5 - v5)) & 63))) & 63))
            var v28 int32 = 2147483647
            _ = v27
            _ = v28
        }
        v29 := GenS2{F0: 660}
        _ = len(v26)
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
