package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 string; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{32768, -170, 0}
    var v3 int8 = 127
    var v4 uint64 = 42
    var v5 int16 = -287
    v6 := make(map[string]bool)
    var v7 string = "\n\t"
    v3 -= -40
    var v8 int16 = (v5 << ((626) & 63))
    v9 := GenS1{F0: 2147483647}
    for fi10 := 0; fi10 < 10; fi10++ {
        v8 += v8
        _ = fi10
    }
    var v11 float64 = ((-2.25 - 0.0) / ((0.0 + 3.14159) + 3.14159))
    if false {
        if ((^(v4)) == v4) {
            var v12 int64 = 397
            _ = v12
        }
    } else {
        v13 := GenS1{F0: 127}
        _ = v13
    }
    if (!true) {
        switch "hello" {
        case "a":
            fmt.Println(v11)
        case "":
            fmt.Println(v8, v3)
        case "foo bar":
            v5 -= (-(v8))
        default:
            var v14 int32 = 0
            _ = v14
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            switch ("foo bar" + v7) {
            case "a":
                v11 = (float64(1.5) * ((float64(((3.14159 / (0.0 - float32((-2.25 + v11)))) / 1.5)) * (v11 / (v11 + (v11 / (((float64(((0.1 * 0.0) - (-2.25 + (3.14159 + 1e10)))) - v11) - 0.0) + v11))))) + 1e10))
            case "hello":
                fmt.Println(v7, fi15)
            case "":
                var v16 uint32 = 2
                _ = v16
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                v4 += v4
                _ = fi17
            }
            v11 = ((v11 + v11) * v11)
            _ = fi15
        }
        v18 := GenS2{F0: 2147483648, F1: v7, F2: 537}
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
