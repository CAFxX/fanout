package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint, p1 int32) int8 {
    var v7 float64 = float64((3.14159 / 0.1))
    var v8 float32 = 1.5
    fmt.Println(p0)
    fmt.Println(v8)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -128
}

func main() {
    v1 := []uint32{65535, 751, 150}
    var v2 uint8 = 255
    var v3 bool = true
    var v4 int16 = -32768
    var v5 int = -1
    var v6 uint64 = 1
    fmt.Println(v2, v2)
    var v9 int64 = -1
    v10 := GenS1{F0: ""}
    v2 -= v2
    var v11 uint32 = 2147483648
    if v3 {
        var v12 int64 = (v9 | 872)
        _ = v12
    }
    v5 = v5
    v1 = append(v1, (v11 >> ((967) & 63)))
    for fi13 := 0; fi13 < 1; fi13++ {
        switch (-(fi13)) {
        case 1:
            v4 += v4
        case -2147483648:
            v1 = append(v1, 42)
        case 128:
            v14 := GenS2{F0: (v2 - (v2 >> ((1) & 63))), F1: 128}
            _ = v14
        default:
            var v15 string = "\n\t"
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS1{F0: ""}
            _ = fi16
            _ = v17
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
