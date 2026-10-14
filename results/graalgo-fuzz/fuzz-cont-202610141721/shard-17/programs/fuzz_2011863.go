package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint32) uint16 {
    var v6 int16 = 32767
    v7 := GenS1{F0: -52, F1: 9}
    _ = p0
    _ = v6
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[string]int)
    var v2 int = -127
    v3 := make(map[uint32]int)
    var v4 int8 = -1
    var v5 uint32 = 4294967295
    var v8 float32 = 1.5
    fmt.Println(v5, v5, v4)
    if true {
        v9 := []int{v2, v2, v2, (^(v2))}
        if false {
            v8 -= v8
            v2 = (v2 ^ v2)
            _ = len(v3)
        }
        if (!(false && true)) {
            fmt.Println(v8, v4, v4)
        }
        _ = v9
    } else {
        var v10 uint8 = 64
        _ = v10
    }
    fmt.Println(v4)
    v5 -= (v5 * v5)
    var v11 int64 = 593
    var v12 string = "a"
    v5 -= (v5 % ((v5) | 1))
    var v13 string = "\n\t"
    var v14 int8 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
