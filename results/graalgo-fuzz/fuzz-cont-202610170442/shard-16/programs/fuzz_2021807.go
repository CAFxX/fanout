package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]int)
    var v3 bool = false
    var v4 float32 = -2.25
    v5 := GenS1{F0: 42, F1: -126}
    if v3 {
        _ = len(v2)
        v4 -= 0.1
        for fi6 := 0; fi6 < 20; fi6++ {
            if v3 {
                v3 = (!(v3 && ((2147483647 > 1) && (((v3 && v3) && v3) && v3))))
            } else {
                var v7 uint16 = 65535
                _ = v7
            }
            _ = fi6
        }
    }
    switch "" {
    case "\n\t":
        fmt.Println(v3, v1, v1)
    default:
        fmt.Println(v3, v1)
    }
    v3 = (v1 >= 1000)
    if v3 {
        fmt.Println(v3, v3)
    } else {
        v8 := GenS2{F0: (v1 | (v1 % ((v1) | 1))), F1: 253}
        _ = v8
    }
    if (v3 && v3) {
        fmt.Println(v4)
        fmt.Println(v4, v4, v4)
        switch 32767 {
        case 128:
            v3 = false
        case -897:
            v4 += (((float32((float64(0.1) / 1.0)) * v4) / 1.5) / (1e10 + float32((0.0 + (0.1 + 1e10)))))
        case 967:
            v9 := GenS3{F0: 42, F1: 18446744073709551615}
            _ = v9
        default:
            var v10 float64 = ((0.1 + 0.1) - (3.14159 + (0.0 / (0.1 - -2.25))))
            _ = v10
        }
    }
    v1 = v1
    v11 := GenS3{F0: 2147483647, F1: 1}
    var v12 float32 = float32((((3.14159 + (3.14159 * 0.0)) * -2.25) + float64(-2.25)))
    switch -1 {
    case -117:
        for fi13 := 0; fi13 < 15; fi13++ {
            if (!(false || (v3 || v3))) {
                v14 := GenS2{F0: 0, F1: 255}
                _ = v14
            }
            _ = len(v2)
            v15 := []int{fi13, 9223372036854775807}
            _ = fi13
            _ = v15
        }
    }
    var v16 int = 1
    var v17 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
