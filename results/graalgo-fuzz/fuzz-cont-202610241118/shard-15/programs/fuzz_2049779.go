package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint; F1 int16; F2 uint }

func main() {
    var v1 uint = 65535
    var v2 bool = (float64((-2.25 / 3.14159)) <= 1.5)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    if (float64(0.1) == 1e10) {
        v6 := []int{}
        _ = v6
    } else {
        fmt.Println(v2, v2)
    }
    var v7 float64 = (0.0 - (-2.25 + ((float64((1.5 + (float32((1e10 + float64((0.1 / float32(((1.5 / 1.0) + 3.14159)))))) / 1e10))) * ((1.5 + 1.5) + ((-2.25 - (1.5 * 1e10)) + float64(1.5)))) + 0.1)))
    var v8 uint64 = 351
    var v9 float32 = 1.5
    v10 := GenS1{F0: uint16(v1), F1: 1}
    if (!(v2 || (42 <= -91))) {
        if ((("" + "foo bar") != "\n\t") || false) {
            var v11 int32 = -2147483648
            var v12 int8 = 0
            _ = v11
            _ = v12
        } else {
            v8 = uint64(v1)
            v13 := GenS2{F0: uint(v8), F1: 32767, F2: v1}
            var v14 uint16 = 32767
            _ = v13
            _ = v14
        }
        v8 = (v8 << ((v1) & 63))
    } else {
        v15 := []int{9223372036854775807, 1, 624}
        var v16 float64 = (v7 - float64(((v9 / v9) * 3.14159)))
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int64 = -738
            _ = fi17
            _ = v18
        }
        _ = v15
        _ = v16
    }
    v1 += (v1 + v1)
    v19 := GenS2{F0: (v1 >> (((v1 * 4294967295)) & 63)), F1: 108, F2: v1}
    for fi20 := 0; fi20 < 20; fi20++ {
        v2 = (129 >= 87)
        _ = fi20
    }
    v8 += (v8 ^ 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
