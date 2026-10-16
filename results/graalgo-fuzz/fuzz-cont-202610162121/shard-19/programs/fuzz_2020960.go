package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 string }
func genFn0(p0 int16) uint64 {
    var v7 uint16 = 255
    var v8 int32 = -376
    v7 = v7
    v7 = v7
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int8 = -73
    var v2 float64 = 0.1
    var v3 string = ""
    var v4 bool = (v1 >= (-(v1)))
    var v5 string = v3
    v6 := make(map[string]int)
    v1 = (v1 & v1)
    v3 = v3
    switch -53 {
    case 127:
        v9 := []int{555}
        _ = v9
    default:
        var v10 bool = (1 != 255)
        _ = v10
    }
    var v11 int = 2147483647
    if true {
        if ((((0.0 + 0.1) * -2.25) + float32(float64(0.1))) <= 0.0) {
            switch v11 {
            case 128:
                fmt.Println(v4, v2, v3)
            }
        } else {
            var v12 uint16 = 42
            _ = v12
        }
        fmt.Println(v2, v1)
        v5 = (v3 + v3)
    }
    var v13 int32 = -603
    var v14 string = ("hello" + "a")
    v15 := GenS1{F0: 338, F1: 766, F2: (("a" + "foo bar") + ("" + (v5 + v3)))}
    v16 := GenS1{F0: genFn0(1), F1: -20, F2: "hello"}
    v17 := GenS1{F0: genFn0(-148), F1: 32767, F2: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
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
