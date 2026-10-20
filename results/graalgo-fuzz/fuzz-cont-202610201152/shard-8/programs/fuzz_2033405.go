package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
func genFn0(p0 int64, p1 uint32, p2 uint16) bool {
    v6 := GenS1{F0: uint16(p1), F1: p2, F2: 42}
    v7 := GenS1{F0: p2, F1: (p2 - 714), F2: -105}
    v8 := []int{4294967295, -31}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return ((p0 << ((255) & 63)) != p0)
}

func main() {
    var v1 uint = 2
    var v2 uint = (^(v1))
    var v3 uint8 = 127
    var v4 float32 = 0.0
    var v5 float64 = float64((v4 - (float32(((1.5 + float64(float32((1.5 * ((((float64(((float32((((0.1 / 1.5) + -2.25) / (0.1 / (-2.25 / 1e10)))) / v4) - 0.1)) / (((1.5 - (0.1 * -2.25)) * 1e10) + (3.14159 * 0.1))) / 1.0) + float64(v4)) + (float64(float32(3.14159)) + (1.5 * 1.5))))))) - (1e10 - 3.14159))) - v4)))
    v4 += (v4 / 3.14159)
    var v9 int16 = -19
    v10 := GenS1{F0: 255, F1: 782, F2: -968}
    var v11 uint = 2
    v12 := []int{32767, 32768, -814, 32768}
    var v13 int16 = -1
    v4 = -2.25
    fmt.Println(v5, v13)
    var v14 string = ""
    if len(v12) > 0 {
        v12[0] = 32767
    }
    var v15 float32 = (v4 / ((v4 / float32((v5 - v5))) + 3.14159))
    for fi16 := 0; fi16 < 16; fi16++ {
        if true {
            switch v14 {
            case "":
                var v17 uint64 = 0
                _ = v17
            case "\n\t":
                v3 += v3
            }
        } else {
            var v18 uint8 = 117
            _ = v18
        }
        if len(v12) > 0 {
            v12[0] = fi16
        }
        fmt.Println(v5)
        _ = fi16
    }
    v19 := GenS1{F0: 0, F1: 32767, F2: 471}
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 float64 = 1e10
        fmt.Println(v21, v2)
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
