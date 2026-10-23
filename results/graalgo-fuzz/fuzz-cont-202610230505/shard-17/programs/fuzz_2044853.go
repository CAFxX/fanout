package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 int }

func main() {
    var v1 float64 = ((1e10 + (float64(((1e10 + 0.1) / 3.14159)) + float64(0.0))) + 1.5)
    var v2 float32 = (1e10 * ((0.1 / -2.25) - (1e10 + -2.25)))
    var v3 int16 = 764
    var v4 float64 = (v1 + ((0.0 / (v1 * v1)) * -2.25))
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: (646 == 1)}
    var v8 int8 = 100
    var v9 uint32 = 4294967295
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS1{F0: true}
        if true {
            if false {
                var v12 int64 = 0
                var v13 float32 = (3.14159 + ((v2 - 0.1) / v2))
                _ = v12
                _ = v13
            } else {
                v8 = (v8 * 127)
            }
        } else {
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(v2, v1)
                v8 = (-(v8))
                fmt.Println(v2, fi14)
                _ = fi14
            }
            var v15 uint16 = 2
            _ = v15
        }
        v16 := GenS1{F0: true}
        _ = fi10
        _ = v11
        _ = v16
    }
    v17 := GenS1{F0: (!(!((!true) || false)))}
    _ = len(v6)
    fmt.Println(v1, v4, v8)
    var v18 int8 = v8
    var v19 string = (("\n\t" + "a") + "a")
    v20 := v5[-552]
    v21 := GenS2{F0: -1, F1: 4294967295, F2: -9223372036854775808}
    var v22 int8 = v8
    v9 = (^(v20))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
