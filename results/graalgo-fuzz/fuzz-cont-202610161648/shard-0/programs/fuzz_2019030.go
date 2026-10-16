package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{"\n\t", (("\n\t" + "hello") + "\n\t"), "hello"}
    var v3 uint16 = 32767
    var v4 float32 = -2.25
    var v5 int8 = 82
    var v6 uint64 = 4294967295
    var v7 float64 = (((-2.25 * float64((1e10 / 1.5))) + float64(0.0)) - 3.14159)
    var v8 bool = (-32768 == -1)
    var v9 string = "a"
    _ = len(v1)
    switch "" {
    case "\n\t":
        v10 := GenS1{F0: false, F1: v3}
        _ = v10
    default:
        var v11 float32 = v4
        _ = v11
    }
    switch 0 {
    case -46:
        v4 = 1e10
    }
    v4 += float32((0.0 / ((float64((v4 * -2.25)) + ((((v7 + v7) / (v7 / (0.0 + (((3.14159 - (v7 - (((v7 + v7) / v7) * v7))) * v7) / float64(v4))))) / v7) + ((float64(v4) - v7) / float64(((0.0 / (1.5 / float32(v7))) * 1.5))))) / (((1e10 - v7) + 1e10) * ((0.1 - ((3.14159 / v7) + 1e10)) - 3.14159)))))
    v8 = (1 != 18446744073709551615)
    v6 -= 18446744073709551615
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 uint = 18446744073709551615
        _ = fi12
        _ = v13
    }
    fmt.Println(v4, v6)
    var v14 uint16 = v3
    if (-124 >= -9223372036854775808) {
        switch "\n\t" {
        case "":
            var v15 int8 = -1
            _ = v15
        case "hello":
            fmt.Println(v14, v3)
        case "foo bar":
            v14 -= v3
        }
        v16 := GenS1{F0: (v7 != ((-2.25 - (v7 + v7)) * v7)), F1: (v3 + v14)}
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
