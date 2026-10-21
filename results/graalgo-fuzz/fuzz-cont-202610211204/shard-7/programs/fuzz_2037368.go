package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int16; F2 bool }
func genFn0(p0 uint16) int32 {
    switch "a" {
    case "":
        v4 := GenS1{F0: 178}
        _ = v4
    case "hello":
        if false {
            var v5 bool = true
            v5 = (4294967295 == 65535)
            _ = v5
        } else {
            v6 := []int{2147483647}
            p0 -= p0
            _ = v6
        }
    default:
        for fi7 := 0; fi7 < 12; fi7++ {
            p0 += 1
            _ = fi7
        }
    }
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 736
}

func main() {
    var v1 int = 127
    var v2 uint = 42
    var v3 uint16 = 32767
    fmt.Println(v1, v1)
    if (!(0.1 >= ((3.14159 / (0.1 - (float32(1.5) / float32(-2.25)))) / -2.25))) {
        switch v1 {
        case 0:
            v1 = v1
        }
        var v8 string = "\n\t"
        _ = v8
    }
    v9 := []int{}
    v1 += (v1 - v1)
    var v10 uint32 = 799
    _ = len(v9)
    fmt.Println(v10, v1)
    v1 = int(v10)
    v11 := GenS1{F0: 1}
    v3 -= v3
    var v12 uint8 = 135
    v10 = (^(v10))
    var v13 string = "\n\t"
    if (0 >= 556) {
        var v14 string = v13
        v9 = append(v9, v1)
        _ = v14
    } else {
        v15 := GenS2{F0: false, F1: -1, F2: false}
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
