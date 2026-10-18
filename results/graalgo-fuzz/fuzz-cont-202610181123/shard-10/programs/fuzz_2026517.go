package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 string; F2 int64 }
func genFn0(p0 bool, p1 int16) uint16 {
    p0 = p0
    v7 := GenS1{F0: -346}
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 uint64) bool {
    p0 = 2
    v8 := GenS1{F0: 541}
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: 0}
        _ = fi9
        _ = v10
    }
    if (true && (2147483648 != 1)) {
        p0 = p0
        p0 -= (p0 ^ (p0 ^ 2))
    } else {
        var v11 float32 = 0.0
        v11 += 3.14159
        _ = v11
    }
    _ = p0
    _ = v8
    return (-2147483648 <= 2147483647)
}

func main() {
    v1 := make(map[int]string)
    var v2 float32 = -2.25
    var v3 int16 = 42
    var v4 int8 = 1
    var v5 float64 = 0.0
    var v6 float64 = (((v5 - v5) / ((float64(float32((((v5 + v5) * (0.0 / v5)) / -2.25))) * -2.25) - ((v5 + 0.0) + float64(v2)))) * ((3.14159 - 0.0) / float64((float32((0.1 - (float64(((((v2 / (v2 / -2.25)) * ((1.5 - v2) * 0.0)) + v2) - ((float32(((v5 * float64((((v2 / 1.5) * v2) * 3.14159))) * ((float64((((3.14159 * 0.0) + v2) + ((v2 + v2) * 0.0))) / float64(v2)) - float64((float32((1.5 / float64((((v2 * float32((float64((v2 * ((v2 + 1.5) / v2))) / (-2.25 / (v5 / -2.25))))) / 1.0) - (float32(((v5 + v5) / 3.14159)) * v2))))) * ((v2 + float32((v5 - 0.1))) + v2)))))) / (-2.25 + v2)) + 0.0))) - v5))) / v2))))
    delete(v1, 32767)
    v12 := GenS2{F0: 2147483647, F1: "\n\t", F2: 0}
    v4 -= v4
    fmt.Println(v5, v3, v4)
    var v13 int8 = 100
    var v14 uint8 = 215
    var v15 int8 = v4
    var v16 int = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
