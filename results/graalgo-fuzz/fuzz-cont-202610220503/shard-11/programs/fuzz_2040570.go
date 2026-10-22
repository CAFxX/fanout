package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint; F2 int16 }
func genFn0(p0 int8, p1 int16, p2 string) string {
    v6 := []int{128}
    if len(v6) > 0 {
        v6[0] = 150
    }
    p2 = p2
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint32 = 65535
        _ = fi7
        _ = v8
    }
    switch -180 {
    case 0:
        var v9 uint8 = 3
        _ = v9
    case 9223372036854775807:
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint16 = 2
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}

func main() {
    var v1 string = "hello"
    var v2 int64 = 393
    v3 := []string{v1, "hello"}
    v4 := []int{-519, 127, 161, 386}
    v5 := []uint32{42}
    if true {
        v12 := GenS1{F0: true}
        _ = v12
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    if false {
        for fi13 := 0; fi13 < 13; fi13++ {
            var v14 int16 = 644
            fmt.Println(fi13)
            _ = fi13
            _ = v14
        }
        var v15 int8 = -73
        v3 = append(v3, ((((("a" + "a") + "\n\t") + "hello") + ("foo bar" + (v1 + v1))) + v1))
        _ = v15
    }
    v2 = 9223372036854775807
    var v16 uint = 2
    v2 -= (v2 << ((255) & 63))
    _ = len(v3)
    v4 = append(v4, -78)
    for fi17 := 0; fi17 < 14; fi17++ {
        _ = len(v3)
        if true {
            v16 = (v16 / ((0) | 1))
            v18 := GenS1{F0: true}
            var v19 uint8 = 42
            _ = v18
            _ = v19
        }
        _ = fi17
    }
    v20 := GenS2{F0: v16, F1: (v16 | 866), F2: 1}
    if true {
        fmt.Println(v1, v16)
    }
    v1 = "\n\t"
    _ = len(v3)
    switch -932 {
    case -978:
        var v21 int16 = 1000
        _ = v21
    default:
        fmt.Println(v16, v2)
    }
    v2 = (v2 - v2)
    switch "" {
    case "foo bar":
        v3 = append(v3, v1)
    case "\n\t":
        var v22 float32 = (float32((0.1 / 1.5)) - ((1.5 * ((float32(3.14159) / 0.1) - 0.1)) * -2.25))
        _ = v22
    }
    v23 := GenS2{F0: (v16 | 42), F1: (v16 | 65535), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
