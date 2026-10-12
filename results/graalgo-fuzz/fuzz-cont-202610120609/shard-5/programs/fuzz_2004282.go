package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    var v1 int = 0
    var v2 int8 = 42
    v3 := make(map[string]string)
    if false {
        if true {
            switch (v1 - (^(v1))) {
            case 128:
                v4 := []int{(-(v1)), v1}
                _ = v4
            case -2147483648:
                if (2147483647 <= 1) {
                    for fi5 := 0; fi5 < 11; fi5++ {
                        v3[""] = ("foo bar" + "a")
                        _ = fi5
                    }
                    v1 -= (v1 & v1)
                    fmt.Println(v2, v2)
                }
            case 146:
                switch -696 {
                case -9223372036854775808:
                    fmt.Println(v2, v1)
                }
            default:
                if false {
                    var v6 uint64 = 23
                    var v7 uint64 = (v6 >> ((18446744073709551615) & 63))
                    _ = v6
                    _ = v7
                } else {
                    var v8 float32 = -2.25
                    v1 -= v1
                    _ = v8
                }
            }
            v9 := []int{32768}
            _ = v9
        } else {
            v10 := GenS1{F0: 9223372036854775807, F1: (-2.25 < 1.5)}
            v2 = v2
            v2 -= (-(v2))
            _ = v10
        }
    }
    v2 = 1
    if false {
        for fi11 := 0; fi11 < 6; fi11++ {
            v2 -= -1
            _ = fi11
        }
    } else {
        fmt.Println(v2, v2)
    }
    v12 := []int{v1, 128, v1, v1}
    var v13 float64 = float64((0.0 * 1.5))
    for fi14 := 0; fi14 < 18; fi14++ {
        v1 = 4294967295
        _ = fi14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        if (2 >= 1) {
            var v16 string = "hello"
            var v17 float32 = (1.5 - (1e10 / (float32(v13) + (float32(float64((1.5 - 0.1))) - 0.1))))
            v16 = "\n\t"
            _ = v16
            _ = v17
        } else {
            v13 -= 3.14159
            v13 -= float64(3.14159)
            v1 = (fi15 | fi15)
        }
        var v18 int64 = -1
        _ = fi15
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        v13 += v13
        switch "" {
        case "a":
            v20 := GenS1{F0: 1, F1: true}
            _ = v20
        default:
            if ((true || ((!(65535 == 2147483648)) || false)) && false) {
                delete(v3, ("foo bar" + (("" + "hello") + ("a" + (("hello" + "") + "\n\t")))))
                if (18446744073709551615 != 0) {
                    if len(v12) > 0 {
                        v12[0] = v1
                    }
                }
                v1 += fi19
            }
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v1, v13, v13)
        if true {
            _ = len(v3)
        }
        v12 = append(v12, fi21)
        _ = fi21
    }
    var v22 int16 = 0
    _ = len(v12)
    v23 := GenS1{F0: 9223372036854775807, F1: true}
    var v24 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
