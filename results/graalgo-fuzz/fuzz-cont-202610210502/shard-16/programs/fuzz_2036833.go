package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0() int {
    v6 := []int{}
    for fi7 := 0; fi7 < 5; fi7++ {
        if true {
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 uint16 = 249
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: 0, F1: 0}
                switch "a" {
                case "hello":
                    if len(v6) > 0 {
                        v6[0] = (^(fi10))
                    }
                case "":
                    var v12 uint64 = 1
                    _ = v12
                }
                var v13 int32 = 657
                fmt.Println(v13, fi7, fi10)
                _ = fi10
                _ = v11
                _ = v13
            }
            var v14 int64 = -9223372036854775808
            _ = v14
        }
        v15 := GenS1{F0: 4294967295, F1: 42}
        if true {
            if false {
                v6 = append(v6, fi7)
                var v16 int32 = -830
                _ = v16
            }
        }
        _ = fi7
        _ = v15
    }
    var v17 int8 = -49
    _ = v17
    _ = v6
    return 963
}

func main() {
    var v1 uint64 = 42
    var v2 float64 = float64((0.0 / ((0.1 + 0.0) + float32((0.1 / 1.0)))))
    var v3 int = -1
    v4 := make(map[int]bool)
    v5 := make(map[string]bool)
    v3 = (v3 << ((65535) & 63))
    v3 = (v3 % ((1) | 1))
    if (!(1 <= 27)) {
        v2 += float64(0.0)
    } else {
        if true {
            fmt.Println(v1, v3, v3)
            var v18 uint8 = 42
            _ = v18
        }
    }
    fmt.Println(v3, v2)
    v19 := GenS1{F0: 765, F1: -1}
    switch (-(v3)) {
    case -9223372036854775808:
        v2 += v2
    case 9223372036854775807:
        fmt.Println(v2, v3)
    case 127:
        var v20 uint16 = 2
        _ = v20
    default:
        for fi21 := 0; fi21 < 9; fi21++ {
            switch v3 {
            case -1:
                delete(v5, "foo bar")
            }
            _ = fi21
        }
    }
    v22 := GenS1{F0: 1, F1: -939}
    v23 := GenS1{F0: 42, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
