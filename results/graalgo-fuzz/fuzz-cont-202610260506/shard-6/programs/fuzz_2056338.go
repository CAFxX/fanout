package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 string; F1 int8; F2 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 float64 = 0.0
    if true {
        switch 42 {
        case 2147483647:
            var v4 uint32 = 0
            _ = v4
        case 859:
            v3 -= float64((-2.25 + float32(v3)))
        default:
            v5 := GenS1{F0: 361, F1: 0, F2: 316}
            _ = v5
        }
        v3 -= v3
        for fi6 := 0; fi6 < 13; fi6++ {
            switch (fi6 - 0) {
            case 32767:
                v3 = v3
            case 509:
                var v7 bool = false
                _ = v7
            case 754:
                v3 = 1.5
            }
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 int64 = 9223372036854775807
                _ = len(v1)
                v9 = v9
                _ = fi8
                _ = v9
            }
            v10 := GenS1{F0: 260, F1: -2147483648, F2: 42}
            _ = fi6
            _ = v10
        }
    }
    v11 := GenS2{F0: ("a" + "foo bar"), F1: 127, F2: 2147483647}
    v3 = v3
    fmt.Println(v3)
    switch ("foo bar" + "foo bar") {
    case "hello":
        var v12 uint = 999
        _ = v12
    default:
        v13 := GenS1{F0: 1, F1: -1, F2: 1}
        _ = v13
    }
    v14 := GenS1{F0: 872, F1: 42, F2: 2147483647}
    fmt.Println(v3)
    fmt.Println(v3)
    v15 := GenS2{F0: ("a" + ("a" + "")), F1: -51, F2: 571}
    v16 := GenS1{F0: 961, F1: 2147483647, F2: 470}
    var v17 uint32 = 142
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
