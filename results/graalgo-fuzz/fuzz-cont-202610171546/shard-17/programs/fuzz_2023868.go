package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{-2147483648, 32767, 9223372036854775807}
    v3 := make(map[int]int)
    var v4 int32 = -1
    v5 := make(map[int]bool)
    v6 := GenS1{F0: 773, F1: 162, F2: 765}
    v4 = (v4 ^ v4)
    if true {
        fmt.Println(v4)
        _ = len(v5)
        var v7 uint = 149
        _ = v7
    } else {
        var v8 int8 = -48
        var v9 int32 = v4
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 962, F1: 18446744073709551615, F2: 2}
    v4 += (v4 * v4)
    var v11 uint16 = 1
    if (!false) {
        fmt.Println(v4)
        for fi12 := 0; fi12 < 19; fi12++ {
            v4 = v4
            switch (fi12 * fi12) {
            case 42:
                switch (("a" + "foo bar") + "a") {
                case "hello":
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v11 += v11
                        _ = fi13
                    }
                case "a":
                    v5[(fi12 * 4294967295)] = true
                default:
                    v4 -= v4
                }
            case 127:
                switch (-(fi12)) {
                case 42:
                    v14 := GenS2{F0: "a"}
                    _ = v14
                case -61:
                    v4 = 880
                }
            default:
                if (false || true) {
                    v4 = (^(v4))
                    switch (fi12 & fi12) {
                    case 9223372036854775807:
                        var v15 int8 = int8(v4)
                        _ = v15
                    case 314:
                        v16 := GenS3{F0: "a", F1: (fi12 << ((65535) & 63)), F2: 238}
                        _ = v16
                    case -1:
                        if true {
                            v2 = append(v2, (fi12 % ((128) | 1)))
                            v11 += v11
                            v11 = v11
                        }
                    default:
                        var v17 uint8 = 127
                        _ = v17
                    }
                    var v18 uint64 = 4294967295
                    _ = v18
                }
            }
            v4 = 2147483647
            _ = fi12
        }
    } else {
        v4 -= v4
    }
    var v19 uint8 = 0
    if len(v2) > 0 {
        v2[0] = -418
    }
    v20 := v1["foo bar"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
