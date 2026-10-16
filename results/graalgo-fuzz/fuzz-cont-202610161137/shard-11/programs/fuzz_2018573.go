package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 uint }
func genFn0(p0 uint, p1 uint16, p2 uint32) bool {
    v5 := GenS1{F0: 0, F1: 38, F2: 848}
    if (1000 == int16(p0)) {
        switch 128 {
        case -913:
            var v6 int8 = -48
            _ = v6
        case -2147483648:
            p2 = p2
        default:
            if (("\n\t" + (("" + "") + (("a" + "foo bar") + "\n\t"))) >= "foo bar") {
                var v7 int32 = -2147483648
                p1 -= p1
                _ = v7
            } else {
                var v8 uint16 = (p1 * p1)
                _ = v8
            }
        }
        v9 := GenS1{F0: -1, F1: 1, F2: p0}
        v10 := GenS1{F0: -587, F1: -128, F2: 1}
        _ = v10
        _ = v9
    }
    var v11 int32 = 1
    if (!(true && false)) {
        v12 := GenS1{F0: 42, F1: 1, F2: 1}
        if true {
            p2 -= p2
            switch 2147483647 {
            case 42:
                v11 += (v11 - (v11 << ((p0) & 63)))
            default:
                p0 = p0
            }
            var v13 int = -9223372036854775808
            _ = v13
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 string = "hello"
                for fi16 := 0; fi16 < 16; fi16++ {
                    v17 := []int{fi16, (fi16 & 313), 42}
                    var v18 int8 = 42
                    var v19 int16 = 42
                    fmt.Println(v11, v18)
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v11 += v11
                _ = fi14
                _ = v15
            }
            var v20 float32 = (float32(0.1) + 0.0)
            p2 -= (p2 % ((793) | 1))
            _ = v20
        }
        p1 -= (p1 / ((p1) | 1))
        _ = v12
    }
    switch ("" + "hello") {
    case "foo bar":
        p2 += p2
    case "a":
        v21 := []int{127, 32767, -892, int(p1)}
        _ = v21
    case "\n\t":
        v22 := GenS1{F0: int16(p1), F1: -26, F2: (^(p0))}
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v5
    return (false && true)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 1
    var v3 float64 = (3.14159 + (1.5 * 3.14159))
    var v4 float64 = v3
    v4 += v4
    fmt.Println(v4)
    if false {
        v4 = (0.1 + -2.25)
        switch ("foo bar" + "a") {
        case "foo bar":
            fmt.Println(v2, v3)
        case "hello":
            v3 -= v3
        }
        v2 = (v2 - 18446744073709551615)
    } else {
        var v23 int32 = -1
        v24 := GenS1{F0: -32768, F1: 99, F2: v2}
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 1000, F1: 5, F2: (v2 << (((v2 & 641)) & 63))}
    _ = len(v1)
    v26 := GenS1{F0: -802, F1: -108, F2: (^(v2))}
    var v27 int = -662
    if (!true) {
        v28 := []int{42}
        if (v27 < (v27 ^ v27)) {
            v2 = (v2 << ((18446744073709551615) & 63))
        }
        if true {
            v2 -= 814
        }
        _ = v28
    }
    v29 := GenS1{F0: 0, F1: 100, F2: (v2 >> (((v2 ^ v2)) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v4
}
