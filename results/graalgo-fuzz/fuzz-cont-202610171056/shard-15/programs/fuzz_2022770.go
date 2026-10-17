package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 bool }

func main() {
    var v1 int32 = -1
    var v2 int64 = 0
    v3 := []uint32{0, 2147483648, 0}
    var v4 uint8 = 128
    var v5 string = ("\n\t" + "")
    var v6 float64 = (((0.0 - 1e10) * ((3.14159 / 1.0) + 3.14159)) / 1.0)
    if (!true) {
        if len(v3) > 0 {
            v3[0] = 105
        }
    } else {
        v4 -= 0
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v7 int8 = -128
    var v8 int8 = int8(v4)
    fmt.Println(v6)
    _ = len(v3)
    fmt.Println(v5)
    v9 := GenS1{F0: -403, F1: 2147483647, F2: true}
    var v10 int32 = -737
    v11 := GenS1{F0: 42, F1: 498, F2: (!(((v4 + v4) != 127) || true))}
    v12 := GenS1{F0: 1, F1: 2147483647, F2: (!false)}
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
