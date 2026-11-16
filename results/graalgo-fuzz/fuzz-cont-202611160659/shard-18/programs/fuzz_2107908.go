package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 uint64, p1 uint32) int16 {
    var v7 int16 = -32768
    p1 -= (p1 * 2147483648)
    v8 := GenS1{F0: 255}
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return v7
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    v3 := []uint32{42, 2147483647, 245}
    var v4 uint32 = 720
    var v5 bool = true
    var v6 uint16 = 65535
    delete(v1, 127)
    v9 := GenS1{F0: 42}
    var v10 float64 = ((0.1 / (0.1 / ((3.14159 + ((0.0 * float64(1e10)) * (1e10 + 0.0))) + 0.1))) + 1.5)
    var v11 uint8 = 155
    v11 = 166
    v5 = true
    switch (("" + "hello") + "foo bar") {
    case "":
        v12 := GenS2{F0: genFn0(442, (v4 % ((0) | 1))), F1: -128}
        _ = v12
    case "a":
        v10 += float64(0.1)
    case "\n\t":
        for fi13 := 0; fi13 < 18; fi13++ {
            switch fi13 {
            case 32767:
                switch "hello" {
                case "hello":
                    var v14 int8 = 46
                    _ = v14
                case "a":
                    v5 = v5
                }
            case -2147483648:
                v15 := GenS1{F0: uint(v11)}
                _ = v15
            default:
                v10 = (float64((-2.25 / float32(v10))) + ((v10 + v10) * v10))
            }
            v11 = v11
            _ = fi13
        }
    }
    var v16 int16 = -1
    v17 := GenS2{F0: 32767, F1: 21}
    var v18 int64 = 694
    for fi19 := 0; fi19 < 16; fi19++ {
        if ((^(v4)) < (v4 * (v4 >> ((255) & 63)))) {
            v20 := GenS2{F0: v16, F1: -13}
            _ = v20
        }
        _ = fi19
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        fmt.Println(v11, v16)
        v11 = 127
        v10 = (v10 * v10)
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
