package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 string }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint64 }
func genFn0(p0 uint64) uint {
    p0 = p0
    p0 = 794
    var v5 uint8 = uint8(p0)
    _ = p0
    _ = v5
    return 2
}

func main() {
    var v1 string = "\n\t"
    var v2 int32 = -903
    v3 := make(map[string]string)
    var v4 bool = true
    if (54 >= 42) {
        if v4 {
            var v6 uint32 = 921
            _ = len(v3)
            _ = v6
        } else {
            v1 = (v1 + "\n\t")
            v7 := v3[v1]
            v1 = "a"
            _ = v7
        }
    }
    v8 := GenS1{F0: 9223372036854775807, F1: 42, F2: v1}
    if (v1 == v1) {
        fmt.Println(v1, v2)
        var v9 int16 = 42
        _ = v9
    }
    v4 = (128 != 127)
    fmt.Println(v4, v1, v1)
    var v10 float32 = 3.14159
    if false {
        var v11 string = v1
        delete(v3, v1)
        switch v11 {
        case "hello":
            v12 := []int{127, -2147483648}
            _ = v12
        case "foo bar":
            var v13 bool = true
            _ = v13
        case "":
            switch v1 {
            case "":
                fmt.Println(v1, v4)
            }
        }
        _ = v11
    } else {
        v10 = (((v10 / (-2.25 - (float32(0.0) * 1.5))) / (3.14159 * (float32(1.5) / v10))) - v10)
    }
    v14 := GenS2{F0: 42, F1: 2, F2: 355}
    v1 = v1
    v10 -= v10
    if ((v1 + "foo bar") != v1) {
        var v15 uint = 4294967295
        fmt.Println(v1, v2, v10)
        var v16 int16 = -507
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
