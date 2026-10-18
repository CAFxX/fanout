package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }

func main() {
    var v1 float32 = 3.14159
    var v2 string = ("" + "\n\t")
    v3 := []uint32{0, 4294967295, 42, 523}
    v4 := make(map[string]string)
    v1 = (float32(0.1) + 0.1)
    fmt.Println(v1, v1)
    switch 32767 {
    case 9223372036854775807:
        v1 = (3.14159 - (1e10 * (v1 - v1)))
    case -947:
        v2 = v2
    }
    v2 = "a"
    var v5 bool = (false && true)
    if v5 {
        _ = len(v4)
    } else {
        switch 0 {
        case 471:
            v6 := GenS1{F0: 234}
            _ = v6
        case 2147483647:
            if true {
                var v7 float32 = v1
                v1 -= 0.0
                switch v2 {
                case "":
                    if v5 {
                        delete(v4, v2)
                    }
                case "hello":
                    var v8 bool = false
                    _ = v8
                case "\n\t":
                    if v5 {
                        v7 = (v1 * v1)
                    }
                }
                _ = v7
            }
        case 497:
            fmt.Println(v2)
        default:
            fmt.Println(v5, v5, v1)
        }
        v1 = 1.5
        v5 = v5
    }
    v1 -= 0.1
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS2{F0: 876, F1: 255, F2: 2147483647}
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(v1, v2, v2)
            v12 := GenS1{F0: 1}
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    switch v2 {
    case "":
        v2 = (((((v2 + ("a" + v2)) + "\n\t") + "hello") + v2) + "\n\t")
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 int32 = -188
    switch 32768 {
    case 128:
        fmt.Println(v5, v2)
    case -1:
        v14 := GenS2{F0: 1, F1: 2, F2: v13}
        _ = v14
    default:
        var v15 bool = ((v13 << ((0) & 63)) < (v13 ^ (-(v13))))
        _ = v15
    }
    fmt.Println(v2, v13, v5)
    fmt.Println(v13, v13)
    var v16 int64 = -806
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
