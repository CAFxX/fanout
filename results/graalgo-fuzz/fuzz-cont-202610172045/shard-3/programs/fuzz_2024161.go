package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 32767
    var v4 string = "a"
    var v5 int8 = -89
    var v6 uint32 = 50
    switch v4 {
    case "hello":
        v1 = v1
    case "":
        switch 32768 {
        case -1:
            if (!(!false)) {
                v7 := GenS1{F0: (^(v3)), F1: (42 <= 933)}
                _ = v7
            } else {
                var v8 int = (^(v1))
                _ = v8
            }
        case 42:
            fmt.Println(v6, v6)
        default:
            var v9 uint16 = (v3 + (v3 * 2))
            _ = v9
        }
    case "a":
        v1 += v1
    default:
        var v10 bool = false
        _ = v10
    }
    v3 += 65535
    fmt.Println(v4, v5)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 int32 = 1
        _ = fi11
        _ = v12
    }
    v13 := []int{v1, (v1 - v1), -1}
    v4 = v4
    var v14 int8 = 42
    v15 := GenS1{F0: v3, F1: true}
    v16 := GenS1{F0: 42, F1: (0.1 >= (0.0 - float32(1.5)))}
    v14 -= 1
    v4 = v4
    var v17 float32 = float32(1.5)
    v4 = v4
    fmt.Println(v4)
    switch v1 {
    case 1:
        if false {
            var v18 int32 = -1
            _ = v18
        } else {
            if false {
                fmt.Println(v6, v6)
                var v19 bool = true
                _ = v19
            }
        }
    case 127:
        v20 := v2[v6]
        _ = v20
    default:
        v21 := GenS1{F0: v3, F1: (!(false && false))}
        _ = v21
    }
    v13 = append(v13, v1)
    switch "" {
    case "hello":
        v17 = 3.14159
    case "":
        fmt.Println(v5, v17)
    default:
        switch (v1 << ((4294967295) & 63)) {
        case 127:
            var v22 uint8 = 255
            _ = v22
        case -527:
            v6 = (v6 + v6)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
