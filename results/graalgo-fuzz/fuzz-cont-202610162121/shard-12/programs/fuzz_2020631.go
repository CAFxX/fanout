package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 65535
    var v3 int32 = -158
    var v4 int8 = 127
    v5 := make(map[string]uint32)
    var v6 float32 = (0.1 / -2.25)
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v6, v6, v2)
    case "\n\t":
        for fi7 := 0; fi7 < 20; fi7++ {
            v6 += v6
            v8 := []int{(-(fi7)), 308}
            _ = fi7
            _ = v8
        }
    case "a":
        if true {
            v9 := GenS1{F0: 0, F1: 255, F2: 1}
            fmt.Println(v2, v4)
            _ = v9
        } else {
            var v10 uint16 = (v2 / ((65535) | 1))
            _ = v10
        }
    default:
        fmt.Println(v4)
    }
    var v11 int16 = 993
    var v12 int16 = v11
    v3 = (v3 / ((int32(v4)) | 1))
    v6 = float32(float64(v6))
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := []int{}
        _ = fi13
        _ = v14
    }
    if ((0.0 + 1e10) > 1e10) {
        for fi15 := 0; fi15 < 4; fi15++ {
            switch fi15 {
            case -64:
                v4 -= (v4 * v4)
            case 32767:
                _ = len(v5)
            case 42:
                v16 := GenS1{F0: 0, F1: 32767, F2: 0}
                _ = v16
            }
            v11 += (v11 | 32767)
            _ = fi15
        }
        var v17 int = -1
        _ = v17
    } else {
        switch 9223372036854775807 {
        case 900:
            fmt.Println(v4)
        case 42:
            if ((true && (true && (true && (1.5 > (0.0 * (3.14159 * 0.1)))))) && true) {
                v3 = 2147483647
                v2 -= 2
                v18 := GenS2{F0: (-(v12))}
                _ = v18
            }
        case -1:
            var v19 uint = 2
            _ = v19
        }
        switch int(v4) {
        case 731:
            var v20 uint32 = 42
            _ = v20
        }
        v21 := []int{9223372036854775807, 2147483647}
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
