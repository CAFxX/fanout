package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint8 }
func genFn0(p0 int64) int8 {
    fmt.Println(p0)
    v6 := GenS1{F0: 128}
    v7 := GenS2{F0: -336}
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 18446744073709551615
    var v3 uint64 = 18446744073709551615
    var v4 int16 = 42
    v5 := make(map[int]bool)
    v3 -= v3
    v8 := []int{}
    switch 42 {
    case 32767:
        var v9 uint16 = 2
        _ = v9
    default:
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 uint32 = 1
            v3 = (v3 - (v3 * v3))
            var v12 float64 = 3.14159
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    fmt.Println(v2, v2, v4)
    v13 := GenS2{F0: (v4 & 42)}
    v4 = v4
    var v14 int64 = -887
    if (2147483647 >= 2147483648) {
        v2 -= (v2 ^ (v2 % (((v2 - (v2 | (^(v2))))) | 1)))
    } else {
        var v15 bool = true
        if (!(!(((v14 % ((v14) | 1)) < 9223372036854775807) && v15))) {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v3 -= (v3 << ((v2) & 63))
            v14 += (v14 | 863)
        }
        v16 := GenS1{F0: 186}
        _ = v15
        _ = v16
    }
    fmt.Println(v4, v4)
    v17 := GenS3{F0: (1 <= 1), F1: v3, F2: 0}
    var v18 uint32 = 0
    for fi19 := 0; fi19 < 13; fi19++ {
        if true {
            var v20 string = ""
            _ = v20
        }
        v21 := v5[1]
        _ = fi19
        _ = v21
    }
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
