package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int8, p1 int64, p2 uint64) uint32 {
    p2 -= 1
    if false {
        p0 += p0
        p1 = p1
    }
    p0 += (^(p0))
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{128}
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 835
}

func main() {
    var v1 uint64 = 473
    var v2 uint64 = (v1 % (((v1 + (v1 / ((42) | 1)))) | 1))
    var v3 int16 = 42
    v4 := []string{"foo bar", ("a" + "\n\t"), ("a" + ""), "hello"}
    v7 := GenS1{F0: true, F1: 0, F2: -951}
    var v8 uint = 91
    v8 += v8
    v2 = (v1 >> ((65535) & 63))
    switch "a" {
    case "foo bar":
        _ = len(v4)
    case "\n\t":
        var v9 int8 = 42
        _ = v9
    default:
        var v10 int8 = 0
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS2{F0: 2, F1: 0}
        var v13 int16 = v3
        v2 += v1
        _ = fi11
        _ = v12
        _ = v13
    }
    if (2 != 65535) {
        v8 -= v8
        var v14 int32 = -2147483648
        _ = v14
    } else {
        v2 += v2
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v2 = v1
        fmt.Println(v2, v3, fi15)
        v4 = append(v4, "hello")
        _ = fi15
    }
    fmt.Println(v3, v8)
    var v16 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
