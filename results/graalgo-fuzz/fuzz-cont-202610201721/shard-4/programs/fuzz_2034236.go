package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }

func main() {
    v1 := []string{}
    v2 := make(map[string]int)
    v3 := []int{32768}
    var v4 uint16 = 42
    var v5 uint8 = 0
    v6 := make(map[int]bool)
    for fi7 := 0; fi7 < 11; fi7++ {
        v1 = append(v1, "hello")
        _ = fi7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        _ = len(v6)
        _ = fi8
    }
    fmt.Println(v4, v5)
    if true {
        v9 := GenS1{F0: 2147483647, F1: -633}
        var v10 uint8 = uint8(v4)
        var v11 float32 = (3.14159 / ((-2.25 * 0.1) / 1.0))
        _ = v10
        _ = v11
        _ = v9
    }
    if true {
        fmt.Println(v4, v5)
        var v12 int32 = -921
        _ = v12
    } else {
        fmt.Println(v4, v4)
    }
    var v13 int8 = 81
    if ((1.5 - 1.5) < ((0.1 / 3.14159) * ((1e10 - float64((3.14159 / 1.5))) / 3.14159))) {
        fmt.Println(v5, v5, v13)
        var v14 int32 = -382
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int32 = -352
        switch ("\n\t" + "") {
        case "hello":
            if len(v1) > 0 {
                v1[0] = "a"
            }
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
