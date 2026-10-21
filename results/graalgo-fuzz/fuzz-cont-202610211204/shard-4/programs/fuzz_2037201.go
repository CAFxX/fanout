package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint; F2 uint }

func main() {
    var v1 float64 = ((0.1 - 1e10) / float64(float32(1.5)))
    var v2 uint8 = 255
    v3 := make(map[uint32]int)
    var v4 string = (("\n\t" + (("\n\t" + ("a" + "hello")) + ("\n\t" + ("foo bar" + "\n\t")))) + (("\n\t" + "foo bar") + ("hello" + "hello")))
    fmt.Println(v4, v2, v1)
    switch 0 {
    case 0:
        v5 := []int{0, -2147483648}
        _ = v5
    case -141:
        for fi6 := 0; fi6 < 12; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 float64 = ((float64(0.0) - float64(1e10)) - float64(((1.5 * 3.14159) * -2.25)))
                var v9 uint16 = 1
                _ = fi7
                _ = v8
                _ = v9
            }
            v10 := v3[678]
            v3[0] = (fi6 + (-(v10)))
            _ = fi6
            _ = v10
        }
    default:
        switch -945 {
        case -2147483648:
            var v11 int16 = 974
            _ = v11
        case 42:
            if ((((1e10 * 3.14159) + 0.0) / (1.5 / 1.0)) > -2.25) {
                v12 := GenS1{F0: uint(v2), F1: (!false)}
                _ = v12
            } else {
                v13 := GenS2{F0: false}
                _ = v13
            }
        default:
            v14 := []int{-2147483648}
            _ = v14
        }
    }
    v4 = v4
    v15 := GenS2{F0: ((0.1 / (1e10 - float32(v1))) >= 1.5)}
    v16 := GenS2{F0: false}
    v4 = v4
    _ = len(v3)
    if true {
        v4 = v4
        var v17 uint8 = 145
        _ = v17
    }
    delete(v3, 392)
    if false {
        v4 = v4
        v18 := GenS3{F0: 42, F1: uint(v2), F2: 18446744073709551615}
        v19 := []int{9223372036854775807}
        _ = v18
        _ = v19
    }
    v2 -= (v2 * (^(v2)))
    v20 := GenS2{F0: (!false)}
    if (0.1 < (-2.25 - 0.0)) {
        var v21 uint16 = 42
        v21 -= v21
        v4 = "foo bar"
        _ = v21
    }
    v4 = v4
    var v22 int8 = 1
    v23 := GenS2{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
