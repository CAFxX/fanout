package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{2147483648, 4294967295, 574}
    var v3 float32 = -2.25
    v4 := make(map[int]string)
    fmt.Println(v3)
    fmt.Println(v3)
    fmt.Println(v3)
    for fi5 := 0; fi5 < 4; fi5++ {
        switch "" {
        case "foo bar":
            v6 := GenS1{F0: 42, F1: 0, F2: 42}
            _ = v6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        v3 = ((1e10 * float32(1e10)) * v3)
        v3 += (float32((3.14159 * (0.1 * float64((1.5 * (float32(1e10) + v3)))))) - 1.5)
        v3 += 1e10
        _ = fi7
    }
    fmt.Println(v3)
    if (539 <= 2) {
        for fi8 := 0; fi8 < 16; fi8++ {
            fmt.Println(v3, v3)
            _ = fi8
        }
        v3 = float32((-2.25 - 0.1))
    }
    _ = len(v2)
    for fi9 := 0; fi9 < 11; fi9++ {
        v3 = 1.5
        _ = fi9
    }
    switch ("hello" + "a") {
    case "\n\t":
        var v10 int32 = 1
        _ = v10
    }
    fmt.Println(v3)
    switch 42 {
    case 201:
        var v11 uint = 65535
        _ = v11
    case 128:
        if false {
            var v12 int64 = 1
            switch "" {
            case "":
                v12 += v12
            case "\n\t":
                v12 = (^(v12))
            default:
                v13 := GenS1{F0: 42, F1: 1, F2: 32767}
                _ = v13
            }
            _ = v12
        } else {
            v3 -= 0.1
            v3 += 3.14159
        }
    case 42:
        v14 := GenS1{F0: 0, F1: 42, F2: 2}
        _ = v14
    }
    v3 += -2.25
    v3 += 0.0
    var v15 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
