package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
func genFn0() string {
    if true {
        v5 := GenS1{F0: 42, F1: 11}
        if (!true) {
            var v6 float32 = 3.14159
            for fi7 := 0; fi7 < 7; fi7++ {
                v6 += float32((float64((float32(3.14159) * -2.25)) + (((1e10 - 0.0) + 0.1) * ((3.14159 / 1.0) + 1e10))))
                fmt.Println(v6, fi7)
                v6 -= (float32((1.5 + 1e10)) + v6)
                _ = fi7
            }
            v6 = 1.5
            _ = v6
        } else {
            var v8 bool = (-2147483648 < 1)
            _ = v8
        }
        switch 127 {
        case 9223372036854775807:
            v9 := GenS1{F0: 617, F1: -1}
            _ = v9
        case -723:
            var v10 uint64 = 4294967295
            _ = v10
        case 2147483647:
            var v11 uint16 = 42
            _ = v11
        }
        _ = v5
    } else {
        v12 := GenS1{F0: 32767, F1: 127}
        var v13 uint32 = 65535
        if ((1.5 - (0.0 / 1.0)) > 0.0) {
            v13 = 1
        } else {
            v13 = (v13 << ((255) & 63))
        }
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 255, F1: -1}
    var v15 int8 = 42
    _ = v14
    _ = v15
    return "\n\t"
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = (0.0 / ((1e10 / 1e10) * -2.25))
    v3 := make(map[int]bool)
    var v4 uint16 = 0
    var v16 string = ("" + (v1 + "a"))
    var v17 int16 = 32767
    fmt.Println(v4, v2, v2)
    switch v16 {
    case "\n\t":
        fmt.Println(v17)
    case "hello":
        v18 := GenS1{F0: (v4 << ((255) & 63)), F1: -128}
        _ = v18
    case "foo bar":
        var v19 uint = 18446744073709551615
        _ = v19
    }
    fmt.Println(v17, v4, v17)
    var v20 int64 = 143
    switch v1 {
    case "":
        var v21 int8 = 100
        _ = v21
    case "a":
        v22 := GenS1{F0: (v4 | v4), F1: -128}
        _ = v22
    case "\n\t":
        var v23 string = (v1 + ((v1 + ("" + v16)) + v1))
        _ = v23
    }
    v1 = v1
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := []int{fi24, fi24}
        _ = fi24
        _ = v25
    }
    switch ((v1 + v16) + v16) {
    case "":
        v26 := v3[42]
        _ = v26
    }
    v27 := []int{2147483647}
    v20 = (v20 % ((42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
}
