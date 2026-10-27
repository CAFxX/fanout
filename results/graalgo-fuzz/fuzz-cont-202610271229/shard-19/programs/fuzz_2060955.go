package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0(p0 uint32, p1 uint64, p2 int64) int64 {
    if false {
        fmt.Println(p0, p2)
        var v6 string = "foo bar"
        _ = v6
    } else {
        p2 = (-(p2))
    }
    if false {
        v7 := GenS1{F0: -9223372036854775808, F1: p1, F2: 32767}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 / (((p2 + (p2 * (p2 << ((42) & 63))))) | 1))
}

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 2
    var v3 float32 = float32(0.0)
    v4 := make(map[int]string)
    var v5 int32 = 0
    if (127 == 42) {
        v3 += v3
        v8 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: 32767}
        _ = v8
    } else {
        v9 := GenS1{F0: 4294967295, F1: 1, F2: -418}
        v3 = (v3 + v3)
        if (830 >= 761) {
            fmt.Println(v5)
            for fi10 := 0; fi10 < 13; fi10++ {
                fmt.Println(v1, v5)
                for fi11 := 0; fi11 < 8; fi11++ {
                    v2 = v2
                    _ = fi11
                }
                _ = fi10
            }
            switch 32767 {
            case 1:
                v12 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615, F2: 128}
                _ = v12
            }
        }
        _ = v9
    }
    delete(v4, 170)
    var v13 uint32 = v2
    v1 += (v1 - -2.25)
    switch ("" + ("a" + "hello")) {
    case "\n\t":
        v14 := GenS2{F0: 127, F1: ((4294967295 >= 18446744073709551615) || (false && true))}
        _ = v14
    case "foo bar":
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 uint64 = 0
            var v17 int32 = 1
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    fmt.Println(v13, v2, v3)
    v18 := []int{}
    var v19 int = -273
    v20 := GenS2{F0: 127, F1: false}
    v19 = int(v5)
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
