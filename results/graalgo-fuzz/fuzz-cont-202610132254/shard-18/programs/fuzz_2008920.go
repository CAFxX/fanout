package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[string]string)
    var v2 string = ((("hello" + ("hello" + ("\n\t" + ("hello" + ("" + "a"))))) + "a") + "hello")
    v3 := make(map[uint32]string)
    var v4 string = "hello"
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 uint64 = 452
        v7 := GenS1{F0: 111, F1: 255}
        v1[v4] = ("a" + v2)
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        var v8 uint16 = 32767
        v9 := GenS1{F0: 2147483648, F1: v8}
        v10 := GenS2{F0: 22, F1: (v8 ^ (v8 * v8))}
        _ = v10
        _ = v8
        _ = v9
    }
    v2 = "hello"
    var v11 float32 = (1e10 * (((((((0.0 + 0.0) - (1.5 + 0.0)) / 1.5) + 0.1) / float32(3.14159)) * 0.1) * (0.1 + -2.25)))
    v12 := GenS2{F0: 1, F1: 255}
    var v13 float32 = 1e10
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 float32 = v13
        v15 -= 0.0
        _ = fi14
        _ = v15
    }
    var v16 float64 = float64(3.14159)
    v13 = v11
    fmt.Println(v16)
    for fi17 := 0; fi17 < 9; fi17++ {
        v18 := GenS3{F0: 100}
        _ = fi17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
