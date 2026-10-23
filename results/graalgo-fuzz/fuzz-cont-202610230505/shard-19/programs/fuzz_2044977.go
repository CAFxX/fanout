package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := []int{42, -1}
    var v2 int8 = 0
    var v3 uint8 = 127
    v4 := make(map[int]bool)
    var v5 uint32 = 2147483648
    var v6 uint16 = 340
    fmt.Println(v2, v2)
    var v7 int8 = 127
    for fi8 := 0; fi8 < 13; fi8++ {
        switch "\n\t" {
        case "\n\t":
            v4[-106] = (false || ((32767 >= -727) && true))
        case "foo bar":
            var v9 int64 = 42
            _ = v9
        }
        v2 = 100
        switch fi8 {
        case 127:
            fmt.Println(v2, fi8)
        default:
            v7 = 111
        }
        _ = fi8
    }
    v2 += 0
    delete(v4, 525)
    if true {
        for fi10 := 0; fi10 < 17; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(fi11)
                _ = fi11
            }
            _ = fi10
        }
    } else {
        var v12 int64 = 164
        _ = v12
    }
    v5 = (^(v5))
    v3 = 1
    fmt.Println(v3)
    var v13 uint8 = (v3 >> ((126) & 63))
    _ = len(v1)
    var v14 string = "foo bar"
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v2, v13)
        v16 := GenS1{F0: 32767}
        if ((0.0 * (0.1 / 1.0)) != ((3.14159 * float64(((float32(float64(((1e10 - 0.0) / 1.0))) + (1.5 + 1.5)) / ((-2.25 / 1.0) * ((1.5 * 3.14159) - 1.5))))) - 1.5)) {
            v17 := GenS2{F0: 42}
            _ = v17
        } else {
            v18 := GenS3{F0: (v13 ^ (v13 ^ v3)), F1: 0}
            var v19 bool = ((-2.25 + (((0.1 - (1e10 - 3.14159)) - (0.1 * ((3.14159 + 1e10) + (-2.25 / (3.14159 - ((3.14159 + 1e10) / (1.5 * 1.5))))))) + 0.1)) >= float64((1.5 / float32(((3.14159 - 1e10) / 3.14159)))))
            var v20 uint16 = 42
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
