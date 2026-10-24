package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    v1 := []int{}
    v2 := make(map[string]bool)
    var v3 int32 = 42
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 3; fi5++ {
        fmt.Println(v3, v3)
        _ = fi5
    }
    fmt.Println(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 := GenS1{F0: 255, F1: 32767, F2: false}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int64 = 588
        _ = fi7
        _ = v8
    }
    var v9 int = -881
    v10 := GenS1{F0: 1, F1: 1, F2: (57 == 1)}
    v11 := GenS1{F0: 42, F1: 255, F2: ((v3 | 211) >= 191)}
    v3 = 42
    if (!(3.14159 < (1e10 * float64(1.5)))) {
        v9 = v9
        if false {
            delete(v2, (("hello" + "") + ""))
            for fi12 := 0; fi12 < 11; fi12++ {
                v3 -= 1
                v9 += fi12
                var v13 uint8 = 220
                _ = fi12
                _ = v13
            }
            switch v9 {
            case 32768:
                delete(v2, "\n\t")
            case -9223372036854775808:
                var v14 bool = true
                _ = v14
            default:
                fmt.Println(v3, v9)
            }
        }
    }
    v15 := GenS2{F0: 32767, F1: 158}
    switch v9 {
    case 263:
        if len(v1) > 0 {
            v1[0] = v9
        }
    case 325:
        var v16 int16 = -753
        _ = v16
    }
    if true {
        v9 += v9
        v9 = 128
    }
    v1 = append(v1, (v9 - (^(v9))))
    if (!false) {
        var v17 float64 = 1.5
        v18 := v4[v9]
        v19 := GenS2{F0: 860, F1: 151}
        _ = v17
        _ = v18
        _ = v19
    } else {
        if true {
            v3 = v3
        }
        v9 -= (v9 & (^(v9)))
        var v20 int32 = -1
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
