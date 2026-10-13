package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint16 = 691
    var v2 int8 = 127
    var v3 float32 = 1e10
    var v4 int = -2147483648
    var v5 int8 = (v2 % (((v2 * (v2 & v2))) | 1))
    var v6 float64 = 1e10
    v5 = v2
    fmt.Println(v2, v5)
    if (true || ((1e10 * (float32(-2.25) - -2.25)) >= float32((1.5 / float64(1.5))))) {
        var v7 uint32 = 2147483647
        for fi8 := 0; fi8 < 14; fi8++ {
            v7 += v7
            _ = fi8
        }
        v9 := []int{v4, 9223372036854775807}
        _ = v7
        _ = v9
    }
    v10 := GenS1{F0: true}
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint8 = 127
        if false {
            fmt.Println(v6)
            v13 := []int{fi11, v4}
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    var v14 int16 = 1000
    var v15 int16 = v14
    v5 = 40
    v3 = float32(((float64(v3) - (v6 * (0.1 + float64(v3)))) / 1.5))
    var v16 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
