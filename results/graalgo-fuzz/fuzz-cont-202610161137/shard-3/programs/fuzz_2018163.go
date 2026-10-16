package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }

func main() {
    v1 := []int{1, 9223372036854775807, 32768}
    var v2 int8 = 0
    var v3 uint = 2
    var v4 float32 = ((float32(0.1) + 0.0) + 0.0)
    var v5 float64 = (0.0 / (((float64(1.5) - 1.5) - float64((1.5 - (v4 - v4)))) - 3.14159))
    var v6 uint64 = 42
    switch -1 {
    case -9223372036854775808:
        if len(v1) > 0 {
            v1[0] = 128
        }
    case 9223372036854775807:
        switch 127 {
        case 4294967295:
            var v7 float32 = v4
            _ = v7
        case 897:
            fmt.Println(v6)
        case 0:
            var v8 string = ""
            _ = v8
        default:
            switch 221 {
            case 2147483647:
                fmt.Println(v5, v4, v5)
            case 32768:
                fmt.Println(v5, v5)
            case 127:
                v9 := GenS1{F0: (("" + "hello") + "\n\t"), F1: 382}
                _ = v9
            }
        }
    default:
        v4 += float32(v5)
    }
    if (1.5 != v4) {
        v10 := GenS1{F0: "a", F1: 2147483647}
        fmt.Println(v6)
        _ = v10
    }
    v3 -= uint(v6)
    for fi11 := 0; fi11 < 12; fi11++ {
        v6 += 42
        v2 = (v2 / ((v2) | 1))
        _ = fi11
    }
    switch 717 {
    case -9223372036854775808:
        for fi12 := 0; fi12 < 15; fi12++ {
            fmt.Println(fi12)
            _ = fi12
        }
    case 128:
        v4 -= 1.5
    default:
        var v13 float64 = -2.25
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        _ = len(v1)
        _ = fi14
    }
    if (true && (!(!true))) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v15 := GenS1{F0: "a", F1: 42}
    v6 += v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
