package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int = 0
    v2 := make(map[int]int)
    v3 := []int{-2147483648, (v1 + 1), (v1 & -9223372036854775808)}
    v4 := make(map[int]bool)
    v5 := []string{"\n\t", "", ""}
    v6 := make(map[uint32]uint32)
    v1 = (v1 % ((-537) | 1))
    var v7 int = (v1 * 127)
    _ = len(v2)
    var v8 string = ("foo bar" + "")
    if false {
        if false {
            v9 := GenS1{F0: v8}
            var v10 uint32 = 4294967295
            _ = v10
            _ = v9
        }
    } else {
        fmt.Println(v1, v7)
        v1 += (v1 * v7)
    }
    if ((1.5 + -2.25) <= 0.0) {
        for fi11 := 0; fi11 < 19; fi11++ {
            _ = len(v2)
            _ = fi11
        }
        v12 := GenS1{F0: v8}
        _ = v12
    } else {
        v7 += -1
        switch v7 {
        case 9223372036854775807:
            v1 += v1
        }
        v8 = ("a" + v8)
    }
    fmt.Println(v7, v7, v1)
    _ = len(v5)
    fmt.Println(v8, v1, v1)
    var v13 int = (v1 << ((65535) & 63))
    var v14 int64 = -9223372036854775808
    v15 := GenS1{F0: (("hello" + v8) + v8)}
    switch int(v14) {
    case -1:
        fmt.Println(v14)
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        switch fi16 {
        case 32768:
            v13 -= v1
        case -518:
            fmt.Println(v7)
        default:
            var v17 string = v8
            _ = v17
        }
        var v18 float64 = (1.5 + (-2.25 / 1.5))
        fmt.Println(v1, v8, v13)
        _ = fi16
        _ = v18
    }
    var v19 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
