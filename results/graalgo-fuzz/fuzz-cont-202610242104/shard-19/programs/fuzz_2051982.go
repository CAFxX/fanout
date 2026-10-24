package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8; F2 int32 }
func genFn0() int {
    var v4 int32 = -1
    fmt.Println(v4)
    _ = v4
    return 128
}

func main() {
    var v1 uint64 = 1
    var v2 int16 = -1
    v3 := make(map[int]bool)
    if true {
        v5 := GenS1{F0: 441}
        if (genFn0() >= genFn0()) {
            if (!(false && (18446744073709551615 <= 886))) {
                v2 += 0
                v6 := GenS1{F0: 42}
                _ = v6
            }
            v7 := GenS1{F0: 2}
            _ = v7
        }
        _ = v5
    }
    var v8 float64 = 1e10
    if (1.5 == (1e10 - 1e10)) {
        var v9 int32 = 2147483647
        if false {
            fmt.Println(v9)
        }
        _ = v9
    }
    if false {
        if (false || (true || false)) {
            v10 := []int{-604, 32768}
            v11 := GenS2{F0: 127, F1: 105, F2: -2147483648}
            _ = v10
            _ = v11
        }
        var v12 string = ("foo bar" + ("" + ""))
        _ = v12
    }
    var v13 int32 = -177
    var v14 bool = (float32(((-2.25 * float64(0.0)) * -2.25)) >= ((3.14159 + -2.25) / ((((float32(v8) + (1e10 - (3.14159 / 1.5))) / (-2.25 + (0.1 / -2.25))) / (float32(v8) * (1e10 - ((0.1 + 3.14159) / 1e10)))) - -2.25)))
    var v15 string = "\n\t"
    var v16 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
