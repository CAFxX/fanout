package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int32; F2 int }
type GenS3 struct { F0 int16 }
func genFn0() uint {
    var v6 float64 = (0.1 - (1.5 / 3.14159))
    for fi7 := 0; fi7 < 19; fi7++ {
        v6 = (((v6 * (3.14159 * (-2.25 * (((v6 - v6) / (((v6 * 3.14159) + v6) - float64(3.14159))) - v6)))) / v6) / 1.0)
        _ = fi7
    }
    switch -418 {
    case 128:
        v6 = v6
    }
    v6 = 3.14159
    v8 := GenS1{F0: true}
    _ = v6
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]uint32)
    var v3 float32 = 3.14159
    var v4 int = 29
    v5 := []uint32{(v1 % ((v1) | 1))}
    v3 = v3
    switch "foo bar" {
    case "hello":
        v9 := GenS2{F0: 42, F1: -2147483648, F2: -741}
        _ = v9
    }
    if true {
        fmt.Println(v3)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if ((v3 / v3) != 1e10) {
        switch 4294967295 {
        case 354:
            if (true && (1 >= 42)) {
                v10 := GenS3{F0: 42}
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 uint64 = 1
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                v13 := v2[2147483648]
                _ = v13
            }
        case 1:
            switch "a" {
            case "\n\t":
                _ = len(v5)
            }
        case 313:
            var v14 int64 = -1
            _ = v14
        }
        var v15 uint64 = 4294967295
        var v16 int64 = -724
        _ = v15
        _ = v16
    }
    var v17 int64 = 1
    v18 := GenS2{F0: (v1 - 2), F1: 0, F2: -2147483648}
    var v19 string = "a"
    switch v4 {
    case 478:
        _ = len(v2)
    case -176:
        var v20 float32 = float32(3.14159)
        _ = v20
    case 32768:
        v17 -= 9223372036854775807
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
