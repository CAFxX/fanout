package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint8 = 128
    var v2 uint = 63
    v3 := make(map[int]string)
    v4 := []int{9223372036854775807, -449, 42}
    var v5 uint = v2
    v6 := GenS1{F0: 0, F1: 643}
    switch "" {
    case "a":
        var v7 string = ("\n\t" + ("foo bar" + "a"))
        _ = v7
    case "hello":
        v4 = append(v4, 42)
    case "":
        var v8 uint8 = v1
        _ = v8
    }
    v1 = 220
    var v9 int32 = -302
    for fi10 := 0; fi10 < 15; fi10++ {
        if (v9 != -215) {
            var v11 uint = v2
            v12 := GenS1{F0: 42, F1: 718}
            _ = v11
            _ = v12
        }
        var v13 int8 = int8(v5)
        v14 := v3[(fi10 & (fi10 + fi10))]
        _ = fi10
        _ = v13
        _ = v14
    }
    var v15 float32 = (0.1 + 1e10)
    var v16 uint = (v5 / ((2) | 1))
    v17 := GenS1{F0: -1, F1: 160}
    if true {
        var v18 bool = (true || (((v5 - 146) == v5) || true))
        switch 32767 {
        case 42:
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 float64 = (1.5 + 1.5)
                _ = fi19
                _ = v20
            }
        case 9223372036854775807:
            var v21 int8 = 100
            _ = v21
        case 32768:
            v18 = (32767 < 255)
        }
        _ = v18
    }
    var v22 uint64 = 42
    var v23 uint64 = (^(v22))
    v3[-2147483648] = "hello"
    v24 := v3[127]
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
