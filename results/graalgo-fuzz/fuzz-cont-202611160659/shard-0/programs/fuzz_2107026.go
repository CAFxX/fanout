package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 bool }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int; F1 int64; F2 int8 }
func genFn0(p0 uint64, p1 int) uint16 {
    p1 += 2147483647
    if (((785 == 255) || (((0.0 * (((-2.25 / (((0.1 - -2.25) / 3.14159) * (0.0 + ((-2.25 * 1e10) - -2.25)))) + ((0.1 / (float32((1.5 * (0.1 * 3.14159))) / float32(((float64(1.5) * 0.1) - (3.14159 * ((1e10 / (1e10 * (-2.25 / 3.14159))) - 1.5)))))) + 1.5)) + float32(1e10))) > (-2.25 * 0.1)) || ((!((3.14159 * -2.25) < (float64((3.14159 * (((((1e10 + (0.0 * ((3.14159 / 1e10) + 1.5))) - ((float32((float64((0.0 - (1e10 / 1.0))) - 3.14159)) * 1e10) + 1e10)) + 1e10) / (1e10 - float32((3.14159 - 0.0)))) / 1.5))) / 1.0))) && (false || true)))) || false) {
        if false {
            for fi4 := 0; fi4 < 16; fi4++ {
                var v5 uint = 65535
                var v6 int32 = -2147483648
                _ = fi4
                _ = v5
                _ = v6
            }
        } else {
            fmt.Println(p0, p1)
            v7 := []int{}
            _ = v7
        }
        if (1 == 42) {
            p1 -= p1
            fmt.Println(p0, p1)
        }
    }
    _ = p0
    _ = p1
    return 775
}

func main() {
    var v1 uint64 = 42
    var v2 float64 = 3.14159
    var v3 bool = (((-81 != -128) && false) && (2 < 245))
    v8 := GenS1{F0: int8(v1), F1: 100, F2: v3}
    var v9 uint = 18446744073709551615
    v9 = v9
    v1 -= (^(v1))
    v10 := GenS2{F0: -52, F1: (-871 <= -198)}
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(v3, v2)
        _ = fi11
    }
    v3 = v3
    v12 := GenS3{F0: 32767, F1: -1, F2: -66}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
