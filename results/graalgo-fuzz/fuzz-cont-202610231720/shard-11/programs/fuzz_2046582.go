package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int8; F1 int16 }

func main() {
    v1 := []uint32{4294967295, 1, 1, 4294967295}
    v2 := make(map[string]string)
    var v3 string = "a"
    v4 := make(map[int]bool)
    var v5 float64 = 0.0
    v6 := []int{-9223372036854775808}
    for fi7 := 0; fi7 < 11; fi7++ {
        if true {
            var v8 uint8 = 186
            _ = v8
        }
        v3 = ((v3 + "a") + "hello")
        _ = fi7
    }
    v5 += v5
    fmt.Println(v5, v3)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if true {
        v3 = (v3 + ((v3 + "hello") + v3))
        v9 := GenS1{F0: 42, F1: 100}
        _ = v9
    } else {
        if true {
            switch 1 {
            case 32767:
                var v10 float32 = (1.5 - -2.25)
                _ = v10
            default:
                v1 = append(v1, 713)
            }
        }
        switch -618 {
        case -9223372036854775808:
            fmt.Println(v5)
        case 141:
            if (false && (0 == 0)) {
                var v11 float32 = ((((1.5 * (float32((-2.25 / (v5 / (v5 / 1.0)))) + 3.14159)) + float32(v5)) + 3.14159) / (3.14159 - 0.0))
                _ = v11
            }
        }
        var v12 string = "\n\t"
        _ = v12
    }
    switch (v3 + (v3 + v3)) {
    case "hello":
        var v13 uint8 = 229
        _ = v13
    case "\n\t":
        v14 := v2[(v3 + (v3 + (v3 + "\n\t")))]
        _ = v14
    case "":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        v15 := GenS1{F0: -32768, F1: 100}
        _ = v15
    }
    v5 = v5
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint32 = 42
        _ = fi16
        _ = v17
    }
    v3 = ("\n\t" + ("" + v3))
    _ = len(v4)
    fmt.Println(v3, v5)
    v5 = 1.5
    var v18 uint64 = 529
    _ = len(v4)
    v19 := GenS1{F0: -777, F1: 22}
    v20 := GenS2{F0: 3, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
