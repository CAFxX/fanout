package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int; F2 uint16 }
func genFn0() int8 {
    v6 := GenS1{F0: 255}
    var v7 uint64 = 18446744073709551615
    v7 = (v7 % (((v7 / (((v7 | (v7 - 2))) | 1))) | 1))
    v8 := GenS2{F0: 372, F1: 32768, F2: uint16(v7)}
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float64 = ((3.14159 * (0.0 - 1e10)) - (1.5 * 0.1))
    var v2 float64 = float64(0.0)
    var v3 float32 = float32((1.5 + 3.14159))
    v4 := make(map[string]bool)
    var v5 uint32 = 1
    v9 := GenS1{F0: 2}
    var v10 float64 = v1
    var v11 int = 32767
    if (!(!false)) {
        var v12 uint16 = 32767
        _ = v12
    }
    v3 += -2.25
    var v13 float32 = 0.1
    var v14 float64 = 1.5
    if false {
        v15 := []int{v11}
        _ = len(v4)
        _ = v15
    } else {
        v16 := []int{-9223372036854775808, (v11 >> ((18446744073709551615) & 63))}
        v17 := GenS1{F0: 475}
        var v18 float32 = (v13 * v3)
        _ = v16
        _ = v17
        _ = v18
    }
    delete(v4, "")
    for fi19 := 0; fi19 < 6; fi19++ {
        switch "hello" {
        case "":
            delete(v4, "\n\t")
        default:
            if (!true) {
                var v20 uint64 = 42
                _ = v20
            } else {
                switch (v11 ^ v11) {
                case 32767:
                    var v21 int64 = 1
                    _ = v21
                case -187:
                    delete(v4, ("hello" + "\n\t"))
                }
            }
        }
        v11 += (^(fi19))
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
