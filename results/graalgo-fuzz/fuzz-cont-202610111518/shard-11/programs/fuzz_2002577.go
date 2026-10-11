package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint64 = 0
    v2 := []uint32{464, 438}
    var v3 float64 = (3.14159 * (0.0 * 0.0))
    v4 := []int{-522, -703, 2147483647, 892}
    v5 := make(map[uint32]uint32)
    v6 := []int{}
    fmt.Println(v1)
    for fi7 := 0; fi7 < 2; fi7++ {
        v2 = append(v2, 65535)
        fmt.Println(v1, v1)
        _ = fi7
    }
    var v8 uint32 = 4294967295
    fmt.Println(v8, v1, v1)
    if true {
        fmt.Println(v1, v3, v8)
        v1 += (v1 - 2)
        v9 := GenS1{F0: 42}
        _ = v9
    } else {
        _ = len(v4)
    }
    if (v8 != v8) {
        fmt.Println(v3, v8)
        v10 := GenS2{F0: -434}
        _ = v10
    } else {
        _ = len(v6)
    }
    switch 127 {
    case 32767:
        fmt.Println(v1)
    case 9223372036854775807:
        switch ("\n\t" + "foo bar") {
        case "a":
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := v5[(v8 / ((v8) | 1))]
                _ = fi11
                _ = v12
            }
        }
    case 897:
        v2 = append(v2, (v8 & 2147483648))
    }
    v6 = append(v6, 32767)
    var v13 float64 = 3.14159
    v13 += (1.5 / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
