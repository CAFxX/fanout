package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int32 = -1
    var v2 float64 = 0.0
    var v3 int32 = v1
    var v4 uint16 = 2
    v5 := make(map[uint32]string)
    v3 += -1
    v2 = v2
    v3 += v1
    v1 = int32(v4)
    var v6 int16 = 0
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            switch 32767 {
            case 9223372036854775807:
                v8 := GenS1{F0: (v4 + v4), F1: (fi7 - (fi7 + fi7)), F2: (fi7 + (fi7 % ((fi7) | 1)))}
                _ = v8
            }
            v9 := GenS2{F0: 255}
            var v10 int8 = 100
            _ = v10
            _ = v9
        } else {
            fmt.Println(v4)
            v1 = v1
            delete(v5, 599)
        }
        fmt.Println(v2, v1)
        _ = fi7
    }
    var v11 uint64 = 0
    v6 -= (^(v6))
    switch -795 {
    case 0:
        if ((((v2 * (((v2 - (float64(float32(3.14159)) + v2)) * 1e10) * (0.0 - v2))) + v2) * v2) >= (v2 + 1.5)) {
            fmt.Println(v4, v2)
            v4 += v4
        } else {
            if true {
                v11 += (v11 >> ((18446744073709551615) & 63))
                var v12 uint64 = 18446744073709551615
                _ = v12
            }
            if true {
                switch -1 {
                case -989:
                    v13 := []int{-2147483648}
                    _ = v13
                case 42:
                    v1 = v3
                }
            } else {
                var v14 int = 579
                _ = v14
            }
            var v15 float64 = v2
            _ = v15
        }
    case 1:
        switch 654 {
        case -1:
            v16 := []int{9223372036854775807}
            _ = v16
        case -9223372036854775808:
            v17 := GenS3{F0: -9223372036854775808}
            _ = v17
        case 128:
            v6 = 1000
        }
    case -969:
        var v18 float64 = (((v2 * ((v2 / v2) - v2)) / 1.5) * v2)
        _ = v18
    default:
        fmt.Println(v2, v2)
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if true {
            var v20 int16 = (v6 >> ((65535) & 63))
            _ = v20
        } else {
            v5[2147483647] = "foo bar"
            v21 := GenS3{F0: 174}
            var v22 uint32 = 0
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
