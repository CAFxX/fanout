package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 bool }
func genFn0(p0 uint, p1 uint8, p2 int64) uint {
    v4 := []int{128, -1}
    if (1e10 < float32(float64(0.0))) {
        v5 := GenS1{F0: 42, F1: 42, F2: ((!false) && (!(!(0 <= 4294967295))))}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v5
    } else {
        p2 -= 42
        p2 += p2
    }
    fmt.Println(p0, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 + p0)
}

func main() {
    v1 := []int{-979}
    v2 := []string{("a" + "foo bar"), "hello", "\n\t"}
    v3 := []string{}
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint32 = 2147483648
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 bool = true
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 bool = (genFn0(1, 187, -9223372036854775808) <= 2)
    var v11 float64 = 1.5
    switch ("foo bar" + (((((("hello" + "hello") + "hello") + "") + ("" + "")) + ("foo bar" + ("foo bar" + ""))) + (("a" + (((("\n\t" + "foo bar") + "foo bar") + "\n\t") + "hello")) + ("foo bar" + (((("a" + "\n\t") + "foo bar") + "foo bar") + "foo bar"))))) {
    case "hello":
        fmt.Println(v10)
    case "":
        fmt.Println(v11, v11)
    }
    if (genFn0(4294967295, 1, -9223372036854775808) <= 1) {
        v10 = (815 >= 4294967295)
        switch 520 {
        case 2147483647:
            v11 = ((v11 / v11) / (v11 + v11))
        }
        if len(v1) > 0 {
            v1[0] = -774
        }
    }
    var v12 float32 = (1e10 * ((0.1 * -2.25) + ((1.5 * (float32(v11) + -2.25)) + ((-2.25 / (1e10 * float32(float64(0.1)))) + 1.5))))
    if v10 {
        var v13 uint16 = 121
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if v10 {
                    v10 = (v10 && v10)
                    var v16 uint = 0
                    if (((v13 | (v13 << (((v16 * v16)) & 63))) >= v13) && (1 != 0)) {
                        fmt.Println(v12, v11, fi14)
                        _ = len(v3)
                    }
                    _ = v16
                }
                _ = fi15
            }
            _ = fi14
        }
        _ = v13
    }
    var v17 float32 = 0.1
    v18 := GenS1{F0: 255, F1: 827, F2: true}
    v17 += v17
    v19 := GenS1{F0: 1, F1: genFn0(genFn0(482, 255, -1), 128, -9223372036854775808), F2: v10}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
