package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 string; F1 uint32 }

func main() {
    var v1 uint = 42
    v2 := make(map[string]string)
    v3 := []uint32{0, 2, 42}
    switch 4294967295 {
    case 1:
        var v4 int8 = 127
        _ = v4
    case -530:
        switch 4294967295 {
        case 32767:
            var v5 uint16 = 2
            _ = v5
        case 0:
            var v6 int16 = 1
            _ = v6
        default:
            switch -45 {
            case -645:
                v2["\n\t"] = ""
            }
        }
    case -2147483648:
        var v7 uint64 = 2
        _ = v7
    }
    var v8 uint64 = 708
    v1 = (v1 / ((v1) | 1))
    var v9 float64 = 3.14159
    var v10 float32 = 1.5
    v11 := GenS1{F0: -177, F1: 32767}
    v3 = append(v3, 316)
    if false {
        if (((3.14159 + ((0.0 - 0.1) + (v10 + (v10 - 1e10)))) / 1.0) <= v10) {
            fmt.Println(v1)
            v9 = 1.5
            fmt.Println(v9, v8)
        } else {
            fmt.Println(v10, v10)
        }
    }
    var v12 string = ((("" + "") + "") + ((("a" + "a") + "\n\t") + ("hello" + ("" + "foo bar"))))
    var v13 uint8 = 0
    v9 += ((float64(3.14159) / v9) * (v9 / (3.14159 - ((v9 * (float64(v10) - v9)) + v9))))
    v14 := GenS2{F0: "a", F1: 2147483647}
    v12 = ""
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 bool = false
        if (((1.5 / (v10 / (v10 / (v10 / 1.0)))) - (3.14159 / v10)) != ((float32(v9) + v10) - (float32(0.0) * float32(0.1)))) {
            v17 := GenS2{F0: (v12 + (v12 + (v12 + v12))), F1: 61}
            var v18 int8 = -1
            switch (^(fi15)) {
            case 376:
                var v19 uint16 = 65535
                _ = v19
            case 0:
                v20 := GenS2{F0: v12, F1: 2147483647}
                _ = v20
            case 9223372036854775807:
                var v21 int = fi15
                _ = v21
            }
            _ = v17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 bool = false
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
