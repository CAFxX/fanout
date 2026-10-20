package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{1, 0}
    v3 := []uint32{69}
    var v4 int32 = -952
    v5 := make(map[string]int)
    var v6 uint = 330
    switch "" {
    case "a":
        v7 := GenS1{F0: ((501 < 1) && (251 > 18446744073709551615)), F1: 42}
        _ = v7
    }
    v8 := GenS1{F0: (!(!false)), F1: 855}
    var v9 uint16 = 840
    if ((!false) || ((v4 * (v4 * (v4 % ((235) | 1)))) == 1)) {
        fmt.Println(v4, v6, v6)
    } else {
        var v10 uint16 = v9
        _ = v10
    }
    v9 = v9
    v4 = v4
    var v11 float32 = 1.5
    var v12 float32 = (3.14159 / 1.5)
    _ = len(v2)
    if false {
        var v13 uint = (v6 % ((18446744073709551615) | 1))
        var v14 uint = 255
        _ = v13
        _ = v14
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    _ = len(v2)
    var v15 int8 = -1
    if true {
        if false {
            var v16 uint8 = 141
            _ = v16
        } else {
            v15 = v15
            v11 = v11
        }
        v11 += v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
