package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0(p0 uint8, p1 int16) uint64 {
    if true {
        var v7 float64 = (1.5 * 1e10)
        _ = v7
    }
    fmt.Println(p1, p0)
    if false {
        switch ("hello" + "\n\t") {
        case "hello":
            p1 = p1
        case "foo bar":
            if true {
                fmt.Println(p0)
                if true {
                    var v8 int = 695
                    var v9 uint32 = 2147483648
                    _ = v8
                    _ = v9
                }
                p1 -= (p1 >> ((61) & 63))
            }
        default:
            v10 := []int{}
            _ = v10
        }
    } else {
        switch 400 {
        case 127:
            v11 := GenS1{F0: 32767, F1: 255}
            _ = v11
        default:
            p0 = 127
        }
        var v12 float32 = (float32(1e10) * ((float32(float64(0.0)) / 3.14159) * (1.5 - (float32(1.5) - 0.0))))
        _ = v12
    }
    v13 := []int{}
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 float64 = (0.0 * 0.1)
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v13
    return 0
}

func main() {
    var v1 int = -1
    var v2 float64 = 1.5
    v3 := []int{-1, v1, 128}
    var v4 int8 = 8
    var v5 float32 = 0.1
    var v6 float32 = 1e10
    v1 = v1
    v16 := GenS2{F0: "", F1: 0}
    var v17 uint = 4294967295
    v1 -= 127
    var v18 string = "hello"
    v1 = (v1 / ((v1) | 1))
    var v19 float64 = (3.14159 + v2)
    if len(v3) > 0 {
        v3[0] = (v1 % ((v1) | 1))
    }
    _ = len(v3)
    fmt.Println(v2, v1)
    fmt.Println(v17, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
