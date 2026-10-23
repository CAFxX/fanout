package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint16 }
func genFn0(p0 string) uint16 {
    v7 := GenS1{F0: 1000, F1: 2147483648, F2: 255}
    var v8 uint8 = 255
    v8 = (^(v8))
    fmt.Println(v8)
    _ = p0
    _ = v7
    _ = v8
    return 32767
}

func main() {
    var v1 uint16 = 0
    var v2 float64 = ((1.5 * 1e10) * 3.14159)
    v3 := make(map[int]uint32)
    var v4 uint = 0
    var v5 int16 = 1000
    var v6 bool = (true || false)
    for fi9 := 0; fi9 < 9; fi9++ {
        v2 += (1e10 * 1.5)
        var v10 uint = 0
        v11 := []int{2147483647, (fi9 / ((-845) | 1)), (^(fi9))}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        v4 += v4
        v2 -= 0.0
        _ = fi12
    }
    fmt.Println(v4, v5)
    switch int(v5) {
    case 487:
        v13 := GenS1{F0: v5, F1: 4294967295, F2: v1}
        _ = v13
    }
    fmt.Println(v4, v1)
    switch -9223372036854775808 {
    case -728:
        v2 -= float64(((1e10 / 0.1) * 3.14159))
    case 32767:
        v1 -= 2
    default:
        fmt.Println(v4)
    }
    v2 = 0.1
    var v14 int64 = -808
    switch "" {
    case "foo bar":
        for fi15 := 0; fi15 < 4; fi15++ {
            v1 = (v1 ^ genFn0("\n\t"))
            var v16 uint64 = 80
            _ = fi15
            _ = v16
        }
    case "\n\t":
        for fi17 := 0; fi17 < 19; fi17++ {
            v4 = 2
            _ = fi17
        }
    }
    var v18 float32 = 3.14159
    v18 -= -2.25
    v2 -= v2
    fmt.Println(v5, v6, v1)
    var v19 uint64 = 1
    var v20 uint = (v4 & (v4 + uint(v19)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
