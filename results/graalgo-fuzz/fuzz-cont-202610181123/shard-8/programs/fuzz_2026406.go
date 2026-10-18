package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int8 {
    var v4 uint32 = 1
    var v5 float32 = float32(-2.25)
    var v6 uint16 = 255
    v5 -= ((float32((((1e10 - 0.1) + -2.25) / (((3.14159 / 0.1) * (((3.14159 * -2.25) / (0.0 - -2.25)) - 3.14159)) - 3.14159))) / v5) / (((v5 - ((-2.25 / v5) + ((v5 - ((v5 / v5) + 0.1)) - v5))) / v5) / v5))
    _ = v4
    _ = v5
    _ = v6
    return 100
}
func genFn1(p0 int64, p1 int) uint {
    v7 := []int{}
    var v8 uint8 = 144
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := []string{"a"}
    var v2 int64 = 696
    var v3 int64 = (v2 << ((255) & 63))
    var v9 float32 = -2.25
    fmt.Println(v3, v2, v9)
    v10 := GenS1{F0: 474}
    var v11 bool = ((0 == -32768) || (((!true) && (!((32767 > -32768) && false))) || false))
    v2 += (v2 & v3)
    v1 = append(v1, "")
    var v12 int = -796
    if ((v12 / (((-(v12))) | 1)) == 32767) {
        _ = len(v1)
    } else {
        var v13 int8 = -128
        v13 += (v13 >> ((796) & 63))
        _ = v13
    }
    var v14 int64 = 9223372036854775807
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
