package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint8, p1 int16) uint {
    v6 := GenS1{F0: 2, F1: (p0 >> ((65) & 63))}
    var v7 bool = true
    var v8 string = ((("foo bar" + "\n\t") + "a") + "hello")
    if v7 {
        var v9 float64 = 0.1
        _ = v9
    }
    v8 = ((v8 + ("" + v8)) + "")
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 float64 = ((3.14159 / 1.0) / 3.14159)
    v2 := make(map[string]string)
    var v3 float32 = float32(v1)
    var v4 int64 = 9223372036854775807
    v5 := make(map[int]int)
    if true {
        v3 = v3
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := []int{}
            v12 := GenS1{F0: 1, F1: 240}
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        var v13 string = "\n\t"
        v14 := []int{-278}
        _ = v13
        _ = v14
    }
    if (!(162 > 1)) {
        v1 = v1
        v1 = -2.25
    }
    fmt.Println(v3, v1)
    v1 += (1.5 + (v1 / v1))
    var v15 string = "\n\t"
    fmt.Println(v15, v1, v4)
    v15 = ""
    fmt.Println(v3, v15, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
