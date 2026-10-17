package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint64 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
func genFn0(p0 uint16) uint8 {
    var v7 uint32 = 4294967295
    p0 = (p0 >> ((255) & 63))
    _ = p0
    _ = v7
    return 2
}

func main() {
    var v1 uint32 = 42
    var v2 uint = 627
    v3 := make(map[uint32]bool)
    var v4 uint16 = 2
    v5 := make(map[int]int)
    var v6 float32 = 3.14159
    for fi8 := 0; fi8 < 7; fi8++ {
        v6 = (v6 + (v6 + float32(0.1)))
        v5[615] = (fi8 << (((^(v2))) & 63))
        _ = fi8
    }
    if (true && false) {
        for fi9 := 0; fi9 < 5; fi9++ {
            if true {
                v4 = (v4 / (((^(v4))) | 1))
                v10 := GenS1{F0: "\n\t", F1: (^(v2)), F2: 42}
                _ = v10
            }
            var v11 uint = (^(v2))
            _ = fi9
            _ = v11
        }
        v12 := []int{32768, 558, -9223372036854775808}
        _ = v12
    }
    v13 := []int{9223372036854775807, 397}
    switch ("foo bar" + "a") {
    case "foo bar":
        if true {
            if len(v13) > 0 {
                v13[0] = 32768
            }
        }
    case "\n\t":
        v2 += (v2 & v2)
    case "a":
        if (61 <= 242) {
            var v14 float64 = (((3.14159 + 3.14159) / -2.25) + (0.0 - 1e10))
            var v15 int32 = 0
            v16 := GenS1{F0: (((("foo bar" + "hello") + ("hello" + "hello")) + ("hello" + "hello")) + "foo bar"), F1: 4294967295, F2: 1}
            _ = v14
            _ = v15
            _ = v16
        } else {
            var v17 int64 = 111
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        if ((^(v2)) > v2) {
            if true {
                v4 += 65535
                v19 := GenS1{F0: "foo bar", F1: v2, F2: 42}
                v2 = v2
                _ = v19
            }
        } else {
            var v20 int32 = 2147483647
            _ = v20
        }
        _ = fi18
    }
    v21 := GenS1{F0: "foo bar", F1: (v2 + 4294967295), F2: 4294967295}
    var v22 int = 32767
    v4 = (v4 & v4)
    switch -25 {
    case 32767:
        v23 := GenS2{F0: ((380 < 1000) && ((v4 + (^(v4))) <= 32767)), F1: -969, F2: 2147483647}
        _ = v23
    case 2147483647:
        switch "" {
        case "a":
            fmt.Println(v1, v22)
        case "\n\t":
            v4 -= v4
        }
    }
    v24 := GenS2{F0: false, F1: (^(v22)), F2: -991}
    v13 = append(v13, (v22 ^ v22))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
