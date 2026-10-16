package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1.5
    v3 := make(map[int]bool)
    var v4 int32 = 469
    v5 := []int{-282}
    fmt.Println(v4, v4)
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v4 -= (v4 >> ((0) & 63))
            _ = fi7
        }
        _ = fi6
    }
    switch ("hello" + "foo bar") {
    case "foo bar":
        v2 = v2
    case "\n\t":
        v8 := GenS1{F0: -751, F1: 129, F2: -101}
        _ = v8
    case "a":
        v4 = v4
    }
    if false {
        v4 = 0
        switch "foo bar" {
        case "":
            v4 -= (^(v4))
        }
    }
    var v9 int = 2147483647
    if false {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint32 = 65535
            _ = fi10
            _ = v11
        }
        fmt.Println(v4)
    } else {
        v2 += ((v2 / v2) - 0.1)
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        if (((v4 % ((-2147483648) | 1)) == (v4 + (v4 ^ v4))) && (247 == 0)) {
            v2 += ((0.1 / v2) * v2)
        } else {
            v1[1] = 2147483648
        }
        _ = fi12
    }
    fmt.Println(v9, v2, v2)
    fmt.Println(v9)
    v9 = (v9 | v9)
    if false {
        var v13 float64 = (1e10 * -2.25)
        _ = v13
    }
    var v14 uint64 = 256
    _ = len(v5)
    for fi15 := 0; fi15 < 18; fi15++ {
        v4 += (v4 << ((65535) & 63))
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
