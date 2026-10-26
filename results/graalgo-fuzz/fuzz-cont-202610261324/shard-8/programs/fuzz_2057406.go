package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
func genFn0() uint16 {
    v7 := []int{-9223372036854775808, -58, -1, 128}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v8 int8 = 100
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint8 = 193
    var v2 uint64 = 42
    var v3 int8 = -30
    v4 := make(map[int]uint32)
    var v5 string = ("a" + "")
    var v6 float32 = 0.0
    v2 -= (v2 >> ((455) & 63))
    v9 := []int{2147483647, 42, -781, 42}
    if (-1 > -1) {
        v10 := GenS1{F0: 2147483647, F1: 32767, F2: (v5 + v5)}
        _ = v10
    }
    if false {
        if len(v9) > 0 {
            v9[0] = 127
        }
        if true {
            v3 = v3
            fmt.Println(v3, v6, v1)
            fmt.Println(v1, v2, v1)
        } else {
            v2 = (v2 ^ 0)
            v6 = float32(float64(-2.25))
        }
    }
    v6 -= float32(1e10)
    v11 := GenS1{F0: -803, F1: 32768, F2: (("" + (("foo bar" + "a") + (v5 + v5))) + "a")}
    v12 := v4[128]
    delete(v4, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
