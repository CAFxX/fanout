package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint }

func main() {
    var v1 uint64 = 0
    var v2 uint = 937
    v3 := []int{-596, 9223372036854775807, 2147483647, 165}
    v1 += 42
    fmt.Println(v2, v2)
    var v4 uint8 = 0
    for fi5 := 0; fi5 < 18; fi5++ {
        switch -371 {
        case 32767:
            switch "hello" {
            case "":
                fmt.Println(v2, v2)
            }
        }
        v2 += v2
        if true {
            var v6 uint16 = 1
            if (!false) {
                var v7 int = (fi5 / ((-9223372036854775808) | 1))
                _ = v7
            }
            _ = v6
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int32 = -513
        _ = fi8
        _ = v9
    }
    var v10 float32 = (0.1 * 1.5)
    var v11 bool = (false && false)
    if v11 {
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 float64 = -2.25
            v3 = append(v3, -2147483648)
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 992, F1: 1, F2: v2}
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
