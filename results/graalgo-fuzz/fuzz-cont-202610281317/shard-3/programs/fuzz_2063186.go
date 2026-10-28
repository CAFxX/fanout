package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint16 = 151
    var v2 uint32 = 2147483648
    var v3 int8 = -128
    v3 -= 100
    var v4 int8 = -128
    v5 := GenS1{F0: (v2 & 926), F1: true}
    var v6 int = -411
    var v7 int = int(v2)
    v8 := []int{v7, (v7 * (v7 - (v7 >> ((42) & 63))))}
    switch "" {
    case "":
        var v9 bool = false
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if true {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        } else {
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 bool = ((((0.1 + ((3.14159 / 1.0) - (-2.25 / -2.25))) / 1e10) + -2.25) <= -2.25)
                var v13 float32 = (((3.14159 * (((1.5 / float32(3.14159)) * (3.14159 * 0.0)) / 1.0)) * ((float32(3.14159) + 1e10) + (0.1 * 0.0))) - (1e10 * 0.0))
                var v14 uint64 = 0
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
        }
        _ = fi10
    }
    v1 += (v1 / ((v1) | 1))
    var v15 bool = (127 >= 0)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
