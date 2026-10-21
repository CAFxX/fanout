package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]bool)
    var v3 int = 42
    var v4 int16 = 42
    var v5 float64 = 3.14159
    v3 = (v3 / (((-(v3))) | 1))
    v1 -= v1
    switch "\n\t" {
    case "\n\t":
        v6 := GenS1{F0: 1}
        _ = v6
    case "hello":
        v1 += v1
    }
    var v7 bool = true
    fmt.Println(v4)
    var v8 int16 = 0
    for fi9 := 0; fi9 < 16; fi9++ {
        if (!(v1 < 127)) {
            fmt.Println(v7)
            v7 = (!(0.0 == -2.25))
        }
        var v10 int16 = (v4 - v4)
        v11 := GenS2{F0: v7}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 bool = v7
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint64 = 815
            var v17 int = 32767
            switch (v17 + fi14) {
            case -201:
                v7 = v7
            case 4294967295:
                var v18 int = -975
                _ = v18
            case -9223372036854775808:
                var v19 int16 = v4
                _ = v19
            default:
                switch "foo bar" {
                case "":
                    v20 := GenS3{F0: -449}
                    _ = v20
                default:
                    v21 := []int{v3}
                    v1 -= (v1 ^ 1)
                    _ = v21
                }
            }
            var v22 uint32 = 934
            _ = fi15
            _ = v16
            _ = v17
            _ = v22
        }
        v8 = v8
        _ = fi14
    }
    var v23 uint32 = 826
    var v24 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
