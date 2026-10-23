package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 float32 = (((0.1 - ((0.1 * 1e10) / (((1e10 + ((((3.14159 * 0.0) / -2.25) + -2.25) * ((0.1 * (-2.25 + 1.5)) + float32(0.0)))) / 1.5) / (0.0 + 1e10)))) - 3.14159) + 1.5)
    var v2 uint32 = 986
    var v3 int32 = 1
    var v4 int = 32768
    var v5 int16 = 0
    var v6 float64 = 3.14159
    v3 = v3
    fmt.Println(v1, v3)
    v7 := GenS1{F0: 2}
    if (v4 == v4) {
        switch 874 {
        case -970:
            v4 += (v4 - -2147483648)
        default:
            var v8 bool = false
            _ = v8
        }
        v3 = -1
    }
    var v9 uint8 = 128
    v10 := GenS1{F0: 32767}
    v6 = -2.25
    var v11 uint32 = v2
    v12 := []int{9223372036854775807}
    v13 := GenS1{F0: 1}
    v12 = append(v12, v4)
    switch (v4 - (v4 ^ v4)) {
    case 1:
        var v14 uint = 1
        _ = v14
    case 42:
        var v15 int = 589
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
