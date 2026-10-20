package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "\n\t"
    var v3 uint16 = 952
    var v4 float32 = ((float32(((((0.0 / 0.1) * ((float64(-2.25) * ((0.0 * 1.5) - 1.5)) + float64(1e10))) - 3.14159) / ((((1e10 - (((0.1 * ((1.5 + -2.25) / 1.5)) * (float64(3.14159) / 1.5)) / 1.5)) * ((1.5 + 3.14159) / (1e10 + (((1e10 - -2.25) * 0.0) * 3.14159)))) - (1.5 / (1.5 + (1e10 / ((3.14159 * 0.1) - (-2.25 / 1.5)))))) * 1.5))) / 0.1) / (1.5 / 1e10))
    v5 := v1[65535]
    var v6 uint32 = (v5 + 2147483647)
    fmt.Println(v6, v4, v2)
    fmt.Println(v2, v3)
    if (!(0 > 127)) {
        var v7 int64 = -9223372036854775808
        v7 = (v7 & v7)
        _ = v7
    } else {
        v4 -= (v4 * 1e10)
        if ((^(v5)) == 155) {
            v8 := GenS1{F0: -1}
            _ = v8
        }
    }
    if true {
        _ = len(v1)
        var v9 float32 = v4
        var v10 int32 = 0
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: -32768}
    delete(v1, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
