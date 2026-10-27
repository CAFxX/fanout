package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    v3 := make(map[string]uint32)
    if (163 != 0) {
        switch "" {
        case "hello":
            v4 := GenS1{F0: 1, F1: 42}
            _ = v4
        case "":
            v2 = (957 == 1)
        default:
            fmt.Println(v2)
        }
    } else {
        v2 = v2
    }
    var v5 uint16 = 65535
    fmt.Println(v2, v2)
    fmt.Println(v5, v5)
    var v6 int16 = 978
    v7 := GenS1{F0: -128, F1: 1000}
    fmt.Println(v2, v6, v6)
    switch "hello" {
    case "\n\t":
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                _ = len(v3)
                _ = fi9
            }
            var v10 int8 = 42
            v2 = ((1.5 >= ((3.14159 - float32(1.5)) - (0.1 + (3.14159 - 1e10)))) && v2)
            _ = fi8
            _ = v10
        }
    default:
        for fi11 := 0; fi11 < 10; fi11++ {
            if (v2 || false) {
                v12 := GenS1{F0: -93, F1: v6}
                var v13 float64 = ((0.0 + 0.1) / 0.1)
                _ = v12
                _ = v13
            } else {
                fmt.Println(v5)
            }
            var v14 bool = (fi11 <= (fi11 / ((fi11) | 1)))
            _ = fi11
            _ = v14
        }
    }
    v5 -= v5
    fmt.Println(v6, v2)
    v15 := []int{-987}
    if (2 > 255) {
        v16 := GenS1{F0: 42, F1: (v6 / ((v6) | 1))}
        v5 += v5
        _ = v16
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        if v2 {
            if (v2 || (fi17 > (fi17 % ((2147483647) | 1)))) {
                v2 = (!v2)
            } else {
                v6 -= v6
            }
            v2 = v2
        }
        v2 = v2
        v5 += v5
        _ = fi17
    }
    var v18 float64 = 3.14159
    var v19 uint = 18446744073709551615
    switch 0 {
    case -17:
        fmt.Println(v18)
    case 795:
        if (32767 == 889) {
            var v20 int64 = 1
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
            _ = v20
        }
    case 4294967295:
        var v21 int8 = 5
        _ = v21
    }
    var v22 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
