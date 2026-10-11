package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int; F2 uint8 }
func genFn0(p0 int8) uint64 {
    if (0.0 != 1e10) {
        switch -642 {
        case -9223372036854775808:
            switch ("\n\t" + "") {
            case "hello":
                p0 = (-(p0))
            case "a":
                p0 = p0
            case "foo bar":
                if true {
                    v5 := GenS1{F0: 143, F1: 1}
                    v6 := GenS1{F0: 168, F1: 65535}
                    _ = v5
                    _ = v6
                } else {
                    fmt.Println(p0)
                }
            }
        case 9223372036854775807:
            p0 = (p0 | 0)
        default:
            p0 += p0
        }
    }
    var v7 uint = 1
    v7 = 42
    v8 := []int{-2147483648}
    if true {
        var v9 int32 = -2147483648
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = len(v8)
        _ = v9
    }
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int32 = -1
    var v2 uint64 = 42
    var v3 float64 = 1e10
    var v4 bool = true
    v2 -= 912
    var v10 float64 = ((v3 / v3) - v3)
    v11 := GenS1{F0: 42, F1: 2}
    for fi12 := 0; fi12 < 17; fi12++ {
        if ((v10 - 1.5) > (v3 + v10)) {
            v13 := []int{(fi12 | (fi12 & fi12)), (^(fi12)), 2147483647, fi12}
            _ = v13
        } else {
            v1 += v1
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 uint64 = v2
            _ = fi14
            _ = v15
        }
        _ = fi12
    }
    var v16 int64 = 709
    if ("\n\t" == "a") {
        switch "\n\t" {
        case "hello":
            v17 := []int{42, 32767, -584}
            _ = v17
        case "":
            v10 += float64(3.14159)
        }
        fmt.Println(v1)
        fmt.Println(v1, v4)
    }
    var v18 uint16 = 789
    if v4 {
        var v19 float64 = v10
        v20 := []int{128, -357}
        if len(v20) > 0 {
            v20[0] = -1
        }
        _ = v19
        _ = v20
    } else {
        var v21 uint = 746
        if ((-9223372036854775808 >= 32768) && false) {
            var v22 int8 = -128
            var v23 int16 = 1
            if v4 {
                v24 := GenS2{F0: (v18 + 59), F1: 2147483647, F2: 209}
                _ = v24
            } else {
                switch "" {
                case "a":
                    v25 := GenS1{F0: 255, F1: v18}
                    _ = v25
                case "":
                    v22 = 17
                case "hello":
                    v23 = (v23 & v23)
                default:
                    v26 := GenS1{F0: 0, F1: v18}
                    _ = v26
                }
            }
            _ = v22
            _ = v23
        } else {
            v1 += (v1 << ((v21) & 63))
            var v27 uint = (v21 << ((v21) & 63))
            v27 = v27
            _ = v27
        }
        var v28 int16 = 1
        _ = v21
        _ = v28
    }
    fmt.Println(v1, v18, v4)
    v10 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
