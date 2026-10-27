package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 0
    var v3 string = "hello"
    v4 := []string{((((v3 + "hello") + ((v3 + v3) + v3)) + "a") + (v3 + ("\n\t" + v3)))}
    var v5 int = 255
    var v6 float64 = (((3.14159 + -2.25) - float64((1.5 - 0.1))) * 3.14159)
    fmt.Println(v6, v3, v2)
    var v7 int16 = -357
    var v8 float64 = ((1.5 * (v6 + 1e10)) / v6)
    var v9 int8 = 100
    switch 4294967295 {
    case 9223372036854775807:
        if ((("hello" + (v3 + v3)) + (v3 + (v3 + v3))) != v3) {
            var v10 int8 = v9
            var v11 uint64 = 272
            _ = v10
            _ = v11
        }
    }
    v12 := GenS1{F0: 1, F1: (v9 * (v9 % (((-(v9))) | 1)))}
    var v13 bool = true
    var v14 string = v3
    v15 := v1[2]
    fmt.Println(v13, v9, v3)
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(v5, v14, v2)
        v4 = append(v4, "")
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint8 = 127
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    var v19 float32 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
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
