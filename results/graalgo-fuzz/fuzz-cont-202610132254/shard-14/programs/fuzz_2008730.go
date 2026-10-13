package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0(p0 bool, p1 uint8) uint32 {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float64 = -2.25
        v8 = (((v8 * v8) * v8) - v8)
        _ = fi7
        _ = v8
    }
    var v9 bool = (!(!p0))
    p1 += (^(p1))
    var v10 uint32 = 65535
    v11 := GenS1{F0: 400, F1: -1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return v10
}
func genFn1() string {
    var v12 int8 = -43
    fmt.Println(v12)
    _ = v12
    return ""
}

func main() {
    var v1 bool = false
    v2 := make(map[string]bool)
    var v3 uint32 = 65535
    v4 := make(map[int]bool)
    var v5 uint16 = 65535
    var v6 uint8 = 127
    var v13 uint = 2
    v6 += uint8(v5)
    fmt.Println(v3, v6, v1)
    v5 -= (v5 >> (((v13 & (v13 - 42))) & 63))
    var v14 int32 = -660
    switch "" {
    case "":
        switch 4294967295 {
        case -760:
            v6 += (^(v6))
        case 0:
            v13 -= (v13 - (v13 % ((v13) | 1)))
        }
    case "foo bar":
        fmt.Println(v3, v14, v3)
    }
    var v15 int8 = -56
    for fi16 := 0; fi16 < 10; fi16++ {
        v3 = (v3 >> ((uint(v3)) & 63))
        var v17 int8 = (-(v15))
        fmt.Println(fi16, v1, v6)
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
