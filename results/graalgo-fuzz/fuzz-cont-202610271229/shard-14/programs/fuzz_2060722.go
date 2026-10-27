package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
func genFn0(p0 bool) int8 {
    fmt.Println(p0)
    p0 = true
    switch "hello" {
    case "foo bar":
        var v6 float64 = 0.0
        _ = v6
    default:
        if ((3.14159 - 0.1) != (1e10 - (1e10 - float64((float32((1e10 - 1.5)) / 3.14159))))) {
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := GenS1{F0: 681, F1: -1}
                _ = fi7
                _ = v8
            }
            var v9 uint16 = 255
            _ = v9
        } else {
            p0 = p0
        }
    }
    if (-952 == -1) {
        v10 := []int{-2147483648, 0, -2147483648}
        _ = v10
    }
    _ = p0
    return -9
}

func main() {
    var v1 uint64 = 0
    var v2 int8 = 20
    var v3 uint32 = 880
    var v4 int32 = 1
    v5 := []uint32{(v3 + (v3 + 1))}
    var v11 int16 = int16(v1)
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(v3)
        _ = fi12
    }
    switch 1 {
    case 32767:
        v3 += (^(v3))
    }
    var v13 int16 = int16(v3)
    var v14 uint16 = uint16(v1)
    var v15 string = "foo bar"
    v3 = 2147483648
    var v16 uint8 = 2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v17 uint64 = uint64(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
