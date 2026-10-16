package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 int16 = -645
    var v4 uint8 = 119
    if (-2147483648 > 127) {
        v5 := []int{-2147483648, 385, -1}
        _ = v5
    } else {
        fmt.Println(v4, v4)
        switch ("a" + "a") {
        case "":
            for fi6 := 0; fi6 < 19; fi6++ {
                for fi7 := 0; fi7 < 5; fi7++ {
                    v8 := GenS1{F0: v4, F1: 65535}
                    _ = fi7
                    _ = v8
                }
                _ = fi6
            }
        case "hello":
            var v9 uint = 776
            _ = v9
        default:
            fmt.Println(v3, v4)
        }
    }
    var v10 int16 = (v3 << ((18446744073709551615) & 63))
    v11 := []int{1, 491}
    v12 := GenS1{F0: 128, F1: 812}
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 float32 = 1.5
            v4 -= v4
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    var v16 float32 = (float32(0.0) - (-2.25 / (3.14159 / -2.25)))
    if false {
        switch "a" {
        case "foo bar":
            v17 := GenS1{F0: v4, F1: 1}
            _ = v17
        case "\n\t":
            for fi18 := 0; fi18 < 13; fi18++ {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 int32 = -896
                    _ = fi19
                    _ = v20
                }
                v4 = (^(v4))
                _ = fi18
            }
        case "a":
            v21 := GenS1{F0: (v4 + v4), F1: 2}
            _ = v21
        default:
            v16 = float32((-2.25 / ((1.5 * 1.5) - (0.1 + 0.1))))
        }
        v4 += v4
        fmt.Println(v4, v10, v3)
    } else {
        fmt.Println(v10)
        var v22 string = "foo bar"
        _ = v22
    }
    if (393 > 2) {
        var v23 float32 = (v16 * (-2.25 + v16))
        var v24 float32 = (((0.0 * (v23 - 0.1)) + 3.14159) / v16)
        if true {
            v24 = 1.5
            var v25 uint32 = 107
            _ = v25
        } else {
            switch 2147483647 {
            case -493:
                v26 := v1[4294967295]
                _ = v26
            case 0:
                v4 = 120
            }
        }
        _ = v23
        _ = v24
    }
    fmt.Println(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
