package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0() int8 {
    v6 := []int{-2147483648}
    var v7 float32 = (((0.1 + float32(0.0)) - 0.1) * ((((((3.14159 + (0.0 - 0.0)) / ((float32((0.1 + 0.0)) + (1.5 / 3.14159)) * 1e10)) + 0.1) - float32(1.5)) / 1.5) + (0.0 * 1e10)))
    if false {
        var v8 uint16 = 32767
        if false {
            v9 := GenS1{F0: -1, F1: 1000}
            var v10 int64 = -434
            var v11 string = "hello"
            _ = v10
            _ = v11
            _ = v9
        } else {
            fmt.Println(v8)
            var v12 uint8 = uint8(v8)
            _ = v12
        }
        _ = v8
    }
    var v13 int64 = 1
    _ = v13
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    var v3 bool = true
    var v4 string = "\n\t"
    v5 := make(map[int]int)
    var v14 uint64 = 42
    var v15 uint16 = 2
    switch 4294967295 {
    case 42:
        var v16 string = ("\n\t" + ((v4 + v4) + (v4 + v4)))
        _ = v16
    }
    v3 = (true || (!v3))
    switch (("" + v4) + "") {
    case "a":
        fmt.Println(v3, v15, v3)
    case "":
        v15 -= v15
    case "\n\t":
        v17 := GenS2{F0: 772, F1: 42}
        _ = v17
    }
    var v18 float32 = (((float32(-2.25) / (3.14159 / ((0.1 / 3.14159) + 1e10))) - 0.0) * (((float32(3.14159) - float32((1e10 * 1e10))) * (1e10 - float32(1.5))) * 0.0))
    v18 += -2.25
    fmt.Println(v15)
    var v19 uint = uint(v15)
    var v20 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
