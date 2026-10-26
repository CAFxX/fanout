package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]string)
    var v3 uint32 = 80
    v4 := make(map[string]int)
    var v5 int8 = 100
    switch "\n\t" {
    case "hello":
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint64 = 484
            v8 := GenS1{F0: (!true), F1: 112}
            var v9 uint16 = 0
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    default:
        v5 = 42
    }
    v10 := GenS1{F0: false, F1: 255}
    v5 += (v5 ^ v5)
    fmt.Println(v5, v1)
    if false {
        switch ("foo bar" + (v1 + "")) {
        case "hello":
            v1 = ("\n\t" + v1)
        }
        v1 = ("hello" + ("a" + v1))
    } else {
        var v11 float32 = (0.1 * ((1e10 * ((1e10 + -2.25) * 3.14159)) / (((3.14159 / 1.5) - float32(((float64((0.0 / 3.14159)) * -2.25) * (0.0 * (-2.25 * (1.5 - float64((3.14159 * 1e10)))))))) + 0.1)))
        _ = v11
    }
    v12 := GenS1{F0: ((-486 >= 132) && false), F1: 1}
    v1 = v1
    v3 += v3
    switch -124 {
    case -2147483648:
        v13 := []int{237, 4294967295, -2147483648}
        _ = v13
    case -446:
        v1 = (((v1 + (v1 + ("" + (((v1 + (v1 + v1)) + v1) + "a")))) + v1) + v1)
    }
    fmt.Println(v1)
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: false, F1: 205}
        _ = fi14
        _ = v15
    }
    v16 := []int{2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
