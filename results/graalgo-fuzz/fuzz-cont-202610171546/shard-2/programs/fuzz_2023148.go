package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint64) uint {
    p0 = p0
    if (0 > 255) {
        v4 := []int{-101, 32768, -963}
        fmt.Println(p0)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v4
    } else {
        p0 -= (p0 << ((1) & 63))
    }
    _ = p0
    return 1
}

func main() {
    v1 := []int{32768}
    v2 := make(map[string]string)
    var v3 uint16 = 65535
    v3 += v3
    v3 += (v3 - v3)
    v3 -= v3
    v3 += (v3 + (v3 & v3))
    v5 := GenS1{F0: int32(v3), F1: v3}
    v3 = (v3 + v3)
    if len(v1) > 0 {
        v1[0] = 32767
    }
    v6 := GenS1{F0: 853, F1: 32767}
    for fi7 := 0; fi7 < 4; fi7++ {
        v3 += v3
        _ = fi7
    }
    v8 := GenS2{F0: 42}
    v3 -= v3
    v9 := GenS1{F0: 0, F1: v3}
    v3 -= (^(v3))
    switch 0 {
    case 127:
        v3 += (^(v3))
    case 4294967295:
        var v10 uint64 = 18446744073709551615
        _ = v10
    case 195:
        v2["foo bar"] = "a"
    }
    v3 -= (v3 * 643)
    if (2 > 1) {
        switch "foo bar" {
        case "hello":
            v11 := v2["a"]
            _ = v11
        case "\n\t":
            var v12 string = "hello"
            _ = v12
        case "":
            v3 = (v3 + (v3 & 65535))
        }
        v3 -= v3
        v1 = append(v1, -1)
    } else {
        var v13 uint8 = 209
        _ = v13
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
