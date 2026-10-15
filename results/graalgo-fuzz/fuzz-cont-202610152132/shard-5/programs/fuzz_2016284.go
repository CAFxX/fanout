package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 string = "hello"
    v4 := make(map[string]int)
    var v5 int8 = -1
    v5 += (v5 ^ v5)
    var v6 int32 = 1
    var v7 string = (("a" + "a") + v3)
    v8 := GenS1{F0: v7, F1: 42, F2: 519}
    for fi9 := 0; fi9 < 1; fi9++ {
        switch (fi9 - fi9) {
        case 128:
            v3 = "hello"
        default:
            v6 -= v6
        }
        fmt.Println(v6)
        var v10 uint = 481
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{(fi11 | 127), fi11, (fi11 / ((127) | 1))}
        _ = fi11
        _ = v12
    }
    var v13 float64 = (3.14159 - (3.14159 - 3.14159))
    for fi14 := 0; fi14 < 15; fi14++ {
        switch fi14 {
        case -2147483648:
            var v15 int32 = v6
            _ = v15
        case 32767:
            v3 = ""
        default:
            delete(v2, 418)
        }
        v7 = v3
        _ = fi14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
