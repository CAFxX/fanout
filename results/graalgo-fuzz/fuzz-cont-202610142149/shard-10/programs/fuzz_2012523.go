package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint = 1
    v2 := make(map[string]uint32)
    v3 := []int{}
    var v4 float32 = (-2.25 + float32(float64(0.0)))
    if true {
        fmt.Println(v1, v4)
    }
    var v5 uint32 = 606
    for fi6 := 0; fi6 < 6; fi6++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v7 float64 = (1.5 / 1e10)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        _ = len(v2)
        v2["\n\t"] = (v5 ^ 853)
        _ = fi8
    }
    var v9 uint64 = 13
    var v10 uint8 = 127
    fmt.Println(v1, v1)
    v9 = (^(v9))
    v4 -= v4
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            v10 -= v10
            if ("foo bar" > "foo bar") {
                fmt.Println(v10, v5, fi12)
                fmt.Println(v5, fi12, fi11)
            }
            v4 = float32(1e10)
            _ = fi12
        }
        if (false && (v1 >= uint(v10))) {
            _ = len(v3)
            v2[""] = (v5 + v5)
        } else {
            fmt.Println(v4, v5)
            var v13 int64 = -391
            _ = v13
        }
        _ = fi11
    }
    if true {
        fmt.Println(v1, v4)
    } else {
        switch -2147483648 {
        case 42:
            v14 := GenS1{F0: (^(v5))}
            _ = v14
        default:
            var v15 int16 = 1
            _ = v15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
