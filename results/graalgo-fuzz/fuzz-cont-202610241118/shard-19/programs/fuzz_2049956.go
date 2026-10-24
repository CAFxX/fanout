package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }

func main() {
    var v1 uint16 = 0
    var v2 uint = 65535
    var v3 float32 = 1.5
    var v4 float64 = 0.0
    var v5 float32 = float32((((0.1 - (v4 * -2.25)) / v4) + (1.5 - 0.0)))
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            if ((v3 + 0.1) != 1.5) {
                var v8 float32 = (v5 / 1.5)
                _ = v8
            }
            var v9 string = "foo bar"
            v10 := []int{fi7}
            var v11 float64 = v4
            _ = fi7
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v2, fi12, v1)
        var v13 bool = true
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := []int{(fi14 >> (((v2 - 149)) & 63)), fi14, 127, int(v1)}
            _ = len(v15)
            if len(v15) > 0 {
                v15[0] = -835
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    var v16 float64 = v4
    switch -238 {
    case 42:
        v16 += v4
    case 259:
        v17 := GenS1{F0: 65535, F1: -1}
        _ = v17
    default:
        v18 := []int{}
        _ = v18
    }
    var v19 int8 = -13
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS1{F0: 792, F1: int64(v2)}
        var v22 string = "\n\t"
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := []int{-9223372036854775808}
    v4 -= 0.0
    v1 += v1
    var v24 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
