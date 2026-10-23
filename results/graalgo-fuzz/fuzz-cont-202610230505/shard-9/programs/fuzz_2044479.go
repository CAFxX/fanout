package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
func genFn0(p0 int8) int64 {
    v6 := []int{1}
    var v7 string = ""
    _ = p0
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int32 = 0
    var v2 uint8 = 255
    var v3 bool = true
    v4 := []int{127, 127, 32767}
    var v5 string = "\n\t"
    var v8 float64 = 0.0
    v9 := GenS1{F0: 532, F1: 2, F2: 255}
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v1)
        _ = fi10
    }
    v4 = append(v4, -9223372036854775808)
    v3 = false
    switch v5 {
    case "":
        var v11 uint64 = 2
        _ = v11
    case "hello":
        var v12 int64 = -568
        _ = v12
    default:
        v4 = append(v4, 0)
    }
    v13 := GenS2{F0: 1, F1: 278, F2: 147}
    if true {
        var v14 int64 = 42
        if (-284 <= 883) {
            v8 = 1.5
            for fi15 := 0; fi15 < 13; fi15++ {
                switch int(v2) {
                case -1:
                    v1 -= -2147483648
                }
                _ = fi15
            }
            var v16 int8 = 75
            _ = v16
        }
        v2 = (^(v2))
        _ = v14
    }
    fmt.Println(v5, v2, v1)
    v5 = "hello"
    var v17 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
