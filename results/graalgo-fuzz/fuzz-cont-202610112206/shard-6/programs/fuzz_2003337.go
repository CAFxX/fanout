package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int }
type GenS2 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint8 = 55
    var v2 string = "foo bar"
    var v3 float64 = -2.25
    var v4 int64 = 0
    var v5 int = -9223372036854775808
    var v6 uint8 = v1
    var v7 uint32 = 0
    var v8 uint16 = 1
    var v9 bool = (!(false && (true || (-1 < int16(v8)))))
    v10 := []int{v5, v5}
    var v11 string = "\n\t"
    var v12 int8 = 100
    var v13 int = 0
    if len(v10) > 0 {
        v10[0] = v13
    }
    v14 := GenS1{F0: (v7 ^ 2147483647), F1: v5, F2: 1}
    if (!(1.5 < (float32(v3) - ((0.1 - 1.5) * (-2.25 / ((float32(v3) - (-2.25 + (3.14159 + 3.14159))) + ((-2.25 / (float32(float64(3.14159)) / (-2.25 / 3.14159))) + (1e10 * 0.0)))))))) {
        if len(v10) > 0 {
            v10[0] = (v5 >> ((42) & 63))
        }
    } else {
        v4 = (v4 | v4)
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int = 256
        v10 = append(v10, 1)
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 0, F1: v8}
    v5 -= 850
    v10 = append(v10, (v13 << ((4294967295) & 63)))
    fmt.Println(v13)
    var v18 uint16 = v8
    var v19 uint = 323
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
