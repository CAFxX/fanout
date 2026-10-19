package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int32 }
type GenS2 struct { F0 int64; F1 string; F2 int64 }
type GenS3 struct { F0 int64; F1 bool; F2 uint64 }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 string) uint8 {
    p0 = p0
    for fi5 := 0; fi5 < 15; fi5++ {
        p0 = "\n\t"
        v6 := []int{(fi5 << ((4294967295) & 63)), fi5, (fi5 | fi5), (fi5 - fi5)}
        var v7 uint16 = 802
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    return 127
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]int)
    var v3 uint = 42
    var v4 uint64 = 2
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint64 = 1
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: true, F2: -875}
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint16 = 2
        switch "foo bar" {
        case "\n\t":
            var v13 int32 = 42
            _ = v13
        case "foo bar":
            if (0.1 <= ((1.5 * (0.1 + 0.1)) * 1e10)) {
                var v14 uint32 = (^(v1))
                v15 := GenS2{F0: int64(v4), F1: "\n\t", F2: 618}
                v16 := []int{910}
                _ = v14
                _ = v15
                _ = v16
            }
        case "hello":
            v12 = (v12 + (^(v12)))
        }
        v17 := v2["foo bar"]
        _ = fi11
        _ = v12
        _ = v17
    }
    var v18 uint16 = 671
    v19 := GenS3{F0: -9223372036854775808, F1: false, F2: 1}
    v3 += (v3 % (((v3 % ((760) | 1))) | 1))
    var v20 int16 = -1
    v3 -= 255
    v21 := GenS4{F0: 382, F1: false}
    v4 += v4
    var v22 string = "\n\t"
    v23 := []int{}
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 uint = (v3 * uint(v18))
        fmt.Println(v4, v1)
        _ = fi24
        _ = v25
    }
    if true {
        for fi26 := 0; fi26 < 17; fi26++ {
            v3 = uint(v18)
            _ = fi26
        }
        var v27 int16 = v20
        for fi28 := 0; fi28 < 18; fi28++ {
            v27 = (^(v27))
            if len(v23) > 0 {
                v23[0] = (fi28 % ((fi28) | 1))
            }
            _ = fi28
        }
        _ = v27
    }
    v1 -= v1
    v29 := GenS4{F0: -245, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
}
