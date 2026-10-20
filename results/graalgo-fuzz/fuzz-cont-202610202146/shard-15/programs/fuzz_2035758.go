package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0() uint8 {
    var v5 int32 = -1
    v6 := []int{798, 128, -210, -279}
    v5 = -2147483648
    v5 -= -1
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int64 = 0
    var v2 uint16 = 0
    var v3 int8 = -128
    var v4 uint16 = v2
    v7 := []int{299, 128, 127}
    var v8 string = (("foo bar" + "\n\t") + (("" + "a") + "foo bar"))
    v4 += 0
    switch 996 {
    case 2:
        for fi9 := 0; fi9 < 17; fi9++ {
            if len(v7) > 0 {
                v7[0] = fi9
            }
            _ = fi9
        }
    default:
        var v10 int64 = v1
        _ = v10
    }
    v11 := GenS1{F0: -572, F1: 382, F2: v3}
    if (((v2 | v4) != (v4 % (((v2 / ((v4) | 1))) | 1))) && (v8 != "")) {
        v3 -= -128
        if len(v7) > 0 {
            v7[0] = 324
        }
        fmt.Println(v4, v3, v4)
    } else {
        _ = len(v7)
        v7 = append(v7, 127)
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS1{F0: 42, F1: 1, F2: (^(v3))}
        v14 := GenS2{F0: (v3 % ((114) | 1)), F1: 1}
        if false {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int32 = 42
                var v17 float64 = (0.0 * 3.14159)
                var v18 int8 = (v3 % ((45) | 1))
                var v19 uint16 = (^(v2))
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
            }
            v1 = v1
        } else {
            var v20 int32 = 42
            v20 -= v20
            v20 = (^(v20))
            _ = v20
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    var v21 uint8 = 0
    var v22 uint = 0
    var v23 int8 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
