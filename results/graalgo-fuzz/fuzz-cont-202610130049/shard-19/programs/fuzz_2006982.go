package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 bool }
func genFn0() int16 {
    for fi5 := 0; fi5 < 13; fi5++ {
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 string = "foo bar"
                _ = fi6
                _ = v7
            }
        case "foo bar":
            v8 := []int{(-(fi5)), (-(fi5))}
            _ = v8
        }
        _ = fi5
    }
    v9 := []int{607, 0, 42}
    var v10 int16 = -32768
    _ = v10
    _ = v9
    return (-(v10))
}

func main() {
    var v1 int = 32768
    v2 := make(map[int]bool)
    var v3 int32 = -396
    var v4 uint = 4294967295
    var v11 float64 = ((0.0 * (3.14159 / 0.1)) * 1e10)
    var v12 uint8 = 42
    fmt.Println(v12, v1)
    v13 := GenS1{F0: 320, F1: 1, F2: false}
    v4 += v4
    v12 -= v12
    var v14 bool = false
    fmt.Println(v1)
    fmt.Println(v1, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
