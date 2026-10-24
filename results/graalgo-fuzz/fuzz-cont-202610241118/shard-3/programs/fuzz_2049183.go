package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0() int {
    var v7 float64 = float64(((1.5 / 1.5) - (0.0 - float32((1.5 - ((3.14159 + 1e10) / 1.0))))))
    fmt.Println(v7)
    _ = v7
    return 2147483647
}
func genFn1() uint64 {
    var v8 string = "hello"
    for fi9 := 0; fi9 < 19; fi9++ {
        v8 = (("foo bar" + v8) + v8)
        fmt.Println(fi9, v8)
        var v10 uint32 = 2147483647
        _ = fi9
        _ = v10
    }
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint64 = 138
    var v2 int16 = -1
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    var v5 int64 = 1
    var v6 int16 = v2
    var v11 float64 = float64((0.0 / 1.0))
    var v12 uint = 401
    var v13 uint = (v12 << (((v12 & (v12 ^ (v12 >> (((v12 % (((v12 & (^(v12)))) | 1))) & 63))))) & 63))
    var v14 int32 = -1
    fmt.Println(v13, v14, v1)
    var v15 uint8 = 127
    var v16 int32 = (v14 / ((-1) | 1))
    fmt.Println(v14)
    v17 := GenS1{F0: 1, F1: ("hello" + ("\n\t" + "\n\t"))}
    v13 = (v12 | v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
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
