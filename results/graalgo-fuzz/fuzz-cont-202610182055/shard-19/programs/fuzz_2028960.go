package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }

func main() {
    var v1 uint16 = 32767
    var v2 float64 = 0.0
    var v3 float64 = -2.25
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 bool = (4294967295 >= 270)
        _ = fi4
        _ = v5
    }
    v2 += v2
    v3 = 0.1
    if false {
        switch 1 {
        case -1:
            if (180 < 2) {
                v6 := GenS1{F0: 42, F1: 42, F2: 1000}
                v7 := []int{4294967295, 4294967295}
                _ = v6
                _ = v7
            } else {
                var v8 uint16 = v1
                v9 := []int{-9223372036854775808, -378, -1}
                _ = v8
                _ = v9
            }
        case 9223372036854775807:
            v1 -= v1
        case 811:
            switch "a" {
            case "foo bar":
                for fi10 := 0; fi10 < 14; fi10++ {
                    var v11 uint8 = 42
                    _ = fi10
                    _ = v11
                }
            case "\n\t":
                v2 += (v2 / (((3.14159 * (float64((-2.25 - 0.0)) * v2)) / ((float64(-2.25) / 1.0) * (v3 + -2.25))) - (1.5 / float64((1e10 - 0.1)))))
            }
        default:
            v12 := GenS2{F0: "hello", F1: 0, F2: v1}
            _ = v12
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int16 = 42
            _ = fi13
            _ = v14
        }
    }
    v2 = float64(((3.14159 + 0.1) - (0.1 + -2.25)))
    var v15 uint = 255
    if (-1 != -42) {
        v3 -= (v2 * (float64((float32((v3 * v3)) + 1.5)) * v2))
        var v16 int16 = 42
        var v17 int = -348
        _ = v16
        _ = v17
    }
    v18 := []int{}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if (true || true) {
            var v20 uint32 = 605
            _ = v20
        }
        fmt.Println(v2)
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
}
