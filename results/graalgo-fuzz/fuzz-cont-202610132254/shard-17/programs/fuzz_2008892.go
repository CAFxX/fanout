package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }

func main() {
    var v1 int8 = 4
    v2 := []uint32{}
    var v3 float32 = ((((float32(3.14159) - ((0.1 + 0.0) * (0.0 + (1e10 + 1.5)))) / 1.5) / 1.5) - -2.25)
    var v4 uint = 18446744073709551615
    var v5 float32 = float32(float64(-2.25))
    v6 := make(map[int]uint32)
    v3 = (float32(0.0) * -2.25)
    v4 += v4
    v7 := GenS1{F0: (v1 * (v1 << ((v4) & 63))), F1: 1}
    var v8 int32 = -319
    v5 -= 1.5
    if false {
        v3 = (0.0 / (((float32(1e10) + v5) - float32((((((-2.25 + 0.0) - 3.14159) + (3.14159 + 3.14159)) / (0.1 - float64((v5 - ((v5 - (1e10 * v3)) + float32(1.5)))))) / (0.1 / (3.14159 * -2.25))))) / 3.14159))
        v9 := v6[32767]
        _ = v9
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: 42, F1: 2}
            var v12 uint64 = 18446744073709551615
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 float64 = (float64(v3) * -2.25)
    var v14 bool = false
    var v15 int8 = v1
    v15 += v15
    fmt.Println(v15)
    for fi16 := 0; fi16 < 1; fi16++ {
        if len(v2) > 0 {
            v2[0] = 836
        }
        var v17 int64 = 0
        _ = fi16
        _ = v17
    }
    v8 = 484
    var v18 uint8 = 127
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 bool = (!(!(!(2147483648 > 942))))
        _ = fi19
        _ = v20
    }
    fmt.Println(v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
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
