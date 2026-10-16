package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int64 = -137
    var v3 int64 = (^(v2))
    var v4 uint32 = 217
    v5 := []int{42}
    v4 -= 471
    switch ((("\n\t" + "\n\t") + "\n\t") + "a") {
    case "a":
        v3 -= int64(v4)
    }
    fmt.Println(v4, v4, v2)
    var v6 int8 = 100
    v6 = v6
    v6 -= v6
    var v7 int16 = -1
    fmt.Println(v7, v4, v4)
    switch -9223372036854775808 {
    case 32768:
        v4 = (v4 + 2)
    case 0:
        if (false || false) {
            v7 -= int16(v1)
        } else {
            var v8 uint32 = v4
            v9 := GenS1{F0: (^(v6)), F1: -32768}
            _ = v8
            _ = v9
        }
    case -9223372036854775808:
        if false {
            v10 := GenS1{F0: -77, F1: 42}
            _ = v10
        }
    default:
        v1 = v1
    }
    v5 = append(v5, 128)
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            switch 819 {
            case 703:
                v13 := GenS2{F0: v3}
                _ = v13
            }
            _ = fi12
        }
        _ = fi11
    }
    var v14 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
