package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 float64 = (3.14159 - 0.0)
    var v2 int8 = 0
    v3 := make(map[int]int)
    var v4 uint64 = 1
    v5 := []int{94, 9223372036854775807, 128, 698}
    v6 := []uint32{0, 0, 430}
    var v7 float64 = v1
    if false {
        fmt.Println(v1, v7)
        v1 = ((v1 * v7) - ((0.1 + ((float64(float32(v7)) - v1) / (float64(3.14159) / v7))) + 1.5))
        var v8 float64 = 3.14159
        _ = v8
    }
    v4 += 2
    v9 := GenS1{F0: 951}
    delete(v3, 1)
    v7 += (v1 * 0.1)
    var v10 uint64 = (v4 + (v4 * v4))
    if false {
        var v11 uint8 = 127
        _ = v11
    } else {
        v7 = (v1 / (((v1 + v7) * (1.5 + 3.14159)) / float64(float32(((v7 / ((v7 + float64(3.14159)) - float64(3.14159))) + (v7 + (1e10 / (v7 + ((1e10 + (v7 / float64((1.5 * ((3.14159 - 0.1) - (1.5 * ((float32(float64(-2.25)) * 0.1) * (1e10 / 3.14159)))))))) * v1)))))))))
    }
    switch 32768 {
    case 0:
        switch 2147483647 {
        case 60:
            var v12 uint64 = v4
            _ = v12
        case 128:
            _ = len(v3)
        default:
            if false {
                _ = len(v3)
            }
        }
    }
    fmt.Println(v10, v1)
    var v13 uint = 0
    var v14 bool = false
    switch "\n\t" {
    case "\n\t":
        var v15 int64 = int64(v13)
        _ = v15
    case "foo bar":
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS1{F0: uint16(v10)}
            _ = fi16
            _ = v17
        }
    }
    v3[0] = 128
    var v18 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
