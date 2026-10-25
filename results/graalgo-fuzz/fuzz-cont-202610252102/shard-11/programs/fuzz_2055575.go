package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint16 }
func genFn0(p0 int64) uint64 {
    p0 = (p0 >> ((65535) & 63))
    if false {
        v5 := []int{0, 2147483647, 2147483647}
        var v6 uint = 631
        _ = v5
        _ = v6
    }
    _ = p0
    return 816
}
func genFn1(p0 int64) uint8 {
    v7 := []int{0, 32768, -780, 1}
    var v8 uint8 = 0
    var v9 int16 = -32768
    v10 := GenS1{F0: v8, F1: int64(v9), F2: 42}
    v7 = append(v7, 1)
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (v8 ^ 202)
}

func main() {
    var v1 float64 = 1e10
    var v2 uint16 = 32767
    v3 := make(map[string]bool)
    var v4 string = "hello"
    switch -463 {
    case -70:
        v2 += (v2 + (v2 - (v2 / ((v2) | 1))))
    case 1:
        v2 = (v2 ^ (v2 >> ((107) & 63)))
    }
    v3[(v4 + v4)] = (genFn1(795) > 2)
    switch ((v4 + v4) + v4) {
    case "\n\t":
        v11 := GenS1{F0: genFn1(0), F1: 670, F2: v2}
        _ = v11
    case "a":
        delete(v3, "a")
    case "foo bar":
        v2 -= 32767
    default:
        fmt.Println(v2)
    }
    if false {
        var v12 string = v4
        _ = v12
    }
    fmt.Println(v1, v2, v2)
    var v13 int = -1
    var v14 string = v4
    var v15 int = (v13 | (v13 + v13))
    fmt.Println(v15)
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint = 263
        _ = fi16
        _ = v17
    }
    switch v15 {
    case 32767:
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 bool = (0 == uint32(v2))
            v2 -= (v2 << ((1) & 63))
            _ = fi18
            _ = v19
        }
    case 9223372036854775807:
        var v20 uint = 0
        _ = v20
    }
    v21 := GenS1{F0: 128, F1: -9223372036854775808, F2: 65535}
    switch (v15 + v13) {
    case -1:
        var v22 int16 = 1
        _ = v22
    case 32767:
        switch 427 {
        case 514:
            fmt.Println(v14, v13)
        default:
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 uint64 = 0
                v13 -= v15
                _ = fi23
                _ = v24
            }
        }
    }
    v2 = v2
    v15 -= int(v2)
    v14 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
