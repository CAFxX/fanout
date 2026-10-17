package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int32 }
type GenS2 struct { F0 uint }

func main() {
    var v1 bool = (!true)
    v2 := []uint32{2147483647}
    v3 := []string{"hello", "hello", ""}
    v4 := []int{6, 2147483647, 128, 32767}
    var v5 uint = 0
    switch "foo bar" {
    case "":
        var v6 float32 = 0.1
        _ = v6
    case "hello":
        var v7 int = 32768
        _ = v7
    }
    switch 128 {
    case -250:
        var v8 uint32 = 2
        _ = v8
    default:
        fmt.Println(v5, v1)
    }
    if (18446744073709551615 <= 2) {
        var v9 float64 = 1e10
        _ = v9
    }
    var v10 uint32 = 2147483647
    v10 += (v10 + v10)
    v11 := GenS1{F0: 127, F1: -1, F2: 1}
    for fi12 := 0; fi12 < 19; fi12++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi12
    }
    v13 := GenS1{F0: 42, F1: 1000, F2: -2147483648}
    v10 = 468
    v5 -= 4294967295
    var v14 float64 = (0.1 - 1e10)
    var v15 uint64 = 0
    var v16 int8 = 1
    for fi17 := 0; fi17 < 19; fi17++ {
        switch fi17 {
        case 960:
            var v18 string = "\n\t"
            _ = v18
        case 128:
            v19 := GenS2{F0: 255}
            _ = v19
        }
        var v20 int8 = (v16 >> (((^(v5))) & 63))
        v14 += float64(((float32(v14) + ((((0.1 - (((1.5 * -2.25) - (((0.1 + 3.14159) + 3.14159) - 1e10)) + (-2.25 * (0.1 / ((0.1 + 3.14159) + 1e10))))) / ((1.5 + 1e10) - 1.5)) + float32(1e10)) - 1.5)) * 1.5))
        _ = fi17
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
