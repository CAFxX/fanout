package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
type GenS2 struct { F0 uint32; F1 string; F2 uint8 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = -2147483648
    v3 := []uint32{65535}
    var v4 int8 = 1
    var v5 uint = 42
    v6 := make(map[string]string)
    v3 = append(v3, 276)
    v5 = (v5 % ((v5) | 1))
    var v7 string = ""
    if false {
        v8 := GenS1{F0: 1, F1: "hello", F2: 127}
        v1 = (v1 | (v1 - (v1 / ((v1) | 1))))
        _ = len(v6)
        _ = v8
    }
    fmt.Println(v1)
    delete(v6, ("hello" + ("a" + ((v7 + (v7 + (((v7 + (v7 + v7)) + ("a" + v7)) + (v7 + v7)))) + "a"))))
    fmt.Println(v1)
    switch 32768 {
    case 799:
        v5 += v5
    case 127:
        if (((1 < 1) || (199 >= 27)) && true) {
            v2 += (v2 + (v2 >> ((255) & 63)))
            _ = len(v3)
            v9 := GenS2{F0: 1, F1: "\n\t", F2: 255}
            _ = v9
        } else {
            v1 = 49
            fmt.Println(v4)
        }
    case -602:
        v10 := GenS2{F0: 4294967295, F1: v7, F2: 106}
        _ = v10
    }
    v11 := GenS1{F0: v1, F1: v7, F2: 42}
    var v12 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
