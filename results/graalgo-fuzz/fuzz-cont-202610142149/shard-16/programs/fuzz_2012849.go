package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 911
    var v3 int16 = 719
    var v4 float64 = float64((((0.0 * -2.25) - v1) / 1.5))
    var v5 float64 = 3.14159
    v6 := make(map[uint32]string)
    var v7 float32 = 0.1
    var v8 string = ("\n\t" + "")
    if true {
        var v9 int16 = (-(v3))
        v10 := GenS1{F0: -784}
        _ = v10
        _ = v9
    }
    fmt.Println(v4, v8)
    var v11 uint = 65535
    v5 += v5
    v12 := GenS1{F0: -2147483648}
    var v13 int = -1
    v14 := []int{32767, v13, v13}
    v15 := GenS1{F0: (v13 * v13)}
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS1{F0: (v13 % ((42) | 1))}
        fmt.Println(v11, v8, v8)
        v3 = v3
        _ = fi16
        _ = v17
    }
    var v18 bool = (false && (true || (715 < (v11 << (((^(v2))) & 63)))))
    v7 = v7
    switch 42 {
    case 1:
        v19 := GenS1{F0: (v13 ^ v13)}
        _ = v19
    case 32768:
        for fi20 := 0; fi20 < 1; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 float32 = (((0.1 * v7) - ((-2.25 * v7) - ((v7 - (((((0.1 + (v7 * v7)) - 1e10) + -2.25) / v1) + (v1 * ((float32(3.14159) * (float32(v5) / (v7 / v1))) * 1e10)))) + v1))) / 3.14159)
                var v23 string = v8
                v24 := GenS1{F0: v13}
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 int64 = 0
            for fi26 := 0; fi26 < 9; fi26++ {
                v13 = (-(fi26))
                _ = fi26
            }
            _ = fi20
            _ = v25
        }
    case -212:
        v3 -= v3
    default:
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    }
    v3 = (v3 % ((1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
