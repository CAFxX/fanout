package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 151
    var v3 uint32 = 618
    var v4 int = -564
    var v5 float64 = 0.1
    v6 := []int{127}
    var v7 int32 = v2
    var v8 float64 = ((1.5 * float64((0.0 + (((1.5 * 1e10) / 1.5) * (-2.25 / ((0.1 - -2.25) / 3.14159)))))) * v5)
    switch "hello" {
    case "a":
        if true {
            var v9 uint8 = 2
            _ = v9
        }
    }
    _ = len(v6)
    v10 := GenS1{F0: 255, F1: "", F2: (142 < 629)}
    for fi11 := 0; fi11 < 13; fi11++ {
        v1[""] = ""
        _ = fi11
    }
    _ = len(v1)
    if (((!(int16(v4) != 1000)) && (!(42 == 18446744073709551615))) || (-128 >= -1)) {
        var v12 uint8 = 128
        _ = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
