package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 4294967295
    var v3 string = "a"
    v4 := []int{-794, 9223372036854775807, 4294967295, 127}
    v5 := make(map[int]bool)
    var v6 string = (("a" + "hello") + (("hello" + ("hello" + v3)) + (v3 + v3)))
    fmt.Println(v6)
    if true {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 int = fi7
            _ = fi7
            _ = v8
        }
        var v9 uint32 = 1
        _ = v9
    }
    var v10 uint8 = 1
    var v11 bool = (0 < 42)
    for fi12 := 0; fi12 < 14; fi12++ {
        switch v6 {
        case "hello":
            v10 = 2
        default:
            var v13 uint32 = 42
            _ = v13
        }
        if (v10 == v10) {
            switch fi12 {
            case -95:
                v14 := GenS1{F0: uint16(v10), F1: 4294967295, F2: v6}
                _ = v14
            default:
                var v15 int8 = -1
                _ = v15
            }
        }
        var v16 uint = v2
        _ = fi12
        _ = v16
    }
    var v17 float32 = 1e10
    var v18 uint32 = 2
    v3 = ""
    _ = len(v4)
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 uint64 = 458
        v17 = ((3.14159 / 1.5) * v17)
        _ = fi19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
