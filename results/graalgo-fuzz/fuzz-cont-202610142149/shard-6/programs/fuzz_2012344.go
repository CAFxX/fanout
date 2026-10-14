package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{437, 2147483647}
    v3 := make(map[int]bool)
    v4 := GenS1{F0: 0, F1: 128}
    var v5 uint8 = 58
    v6 := v1[-2147483648]
    v5 = v5
    for fi7 := 0; fi7 < 8; fi7++ {
        delete(v1, (fi7 >> ((31) & 63)))
        _ = fi7
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int8 = 0
            if (false || true) {
                fmt.Println(v6, v5, v5)
            } else {
                if false {
                    fmt.Println(v6)
                } else {
                    v11 := GenS2{F0: v10, F1: v10}
                    fmt.Println(v5, v5, v5)
                    var v12 int8 = (v10 << ((18446744073709551615) & 63))
                    _ = v11
                    _ = v12
                }
            }
            if true {
                if false {
                    v10 -= v10
                    var v13 int8 = (v10 ^ v10)
                    _ = v13
                }
                if false {
                    var v14 uint16 = 32767
                    fmt.Println(v5, v6, fi8)
                    _ = v14
                }
            }
            switch v6 {
            case "a":
                v5 = 0
            case "":
                v15 := GenS2{F0: (v10 | 1), F1: v10}
                _ = v15
            case "foo bar":
                var v16 uint32 = 42
                _ = v16
            }
            fmt.Println(fi8, v5, v5)
            _ = fi9
            _ = v10
        }
        fmt.Println(v5, fi8, v5)
        var v17 int = (fi8 >> ((65535) & 63))
        _ = fi8
        _ = v17
    }
    v6 = v6
    for fi18 := 0; fi18 < 6; fi18++ {
        v1[(fi18 / ((fi18) | 1))] = "a"
        v19 := GenS2{F0: 1, F1: 0}
        _ = len(v2)
        _ = fi18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
