package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
func genFn0(p0 bool, p1 uint8, p2 int64) uint32 {
    v7 := GenS1{F0: -275, F1: 2147483647}
    v8 := GenS1{F0: -2147483648, F1: 42}
    p2 = (-(p2))
    if (p0 && (!false)) {
        var v9 string = ((("hello" + (("foo bar" + "\n\t") + "a")) + "a") + ("a" + (("a" + "\n\t") + "")))
        v10 := GenS1{F0: 4294967295, F1: 2147483647}
        _ = v10
        _ = v9
    }
    var v11 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    _ = v8
    return 758
}

func main() {
    var v1 int = -949
    var v2 int8 = 127
    var v3 uint16 = 667
    var v4 float64 = (float64((3.14159 - 1e10)) - 3.14159)
    var v5 uint32 = 65535
    var v6 bool = false
    v2 += -47
    switch "foo bar" {
    case "a":
        v12 := GenS1{F0: (v1 | (v1 ^ v1)), F1: 42}
        _ = v12
    case "\n\t":
        var v13 float64 = v4
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v5 = v5
        if (1 > 142) {
            switch 9223372036854775807 {
            case 2147483647:
                var v15 int64 = -9223372036854775808
                _ = v15
            case -9223372036854775808:
                v16 := GenS1{F0: fi14, F1: -2147483648}
                _ = v16
            }
        } else {
            if (v1 >= -751) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 float32 = -2.25
                    _ = fi17
                    _ = v18
                }
            } else {
                v19 := []int{32768}
                var v20 bool = v6
                _ = v19
                _ = v20
            }
            if (!false) {
                v2 = (v2 << ((4294967295) & 63))
                var v21 int16 = 577
                _ = v21
            }
        }
        v22 := []int{-2147483648, (v1 ^ (v1 & fi14)), 188, 9223372036854775807}
        _ = fi14
        _ = v22
    }
    v23 := GenS2{F0: 0, F1: (v2 / ((v2) | 1))}
    v4 = 1.5
    var v24 uint8 = 22
    if false {
        var v25 int16 = 163
        if v6 {
            switch 32768 {
            case 2147483647:
                if v6 {
                    var v26 float64 = (((0.0 * 0.1) + ((3.14159 + (3.14159 - float64((-2.25 * (1e10 / 3.14159))))) - 3.14159)) + (1.5 + -2.25))
                    _ = v26
                }
            }
        }
        _ = v25
    } else {
        if (genFn0(true, v24, 9223372036854775807) <= genFn0(false, (v24 << ((2) & 63)), -9223372036854775808)) {
            for fi27 := 0; fi27 < 13; fi27++ {
                v3 = (v3 << ((0) & 63))
                v28 := []int{fi27, (^(fi27))}
                v1 -= (fi27 + fi27)
                _ = fi27
                _ = v28
            }
            v29 := GenS1{F0: 0, F1: -1}
            _ = v29
        }
    }
    v24 += (v24 >> ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
