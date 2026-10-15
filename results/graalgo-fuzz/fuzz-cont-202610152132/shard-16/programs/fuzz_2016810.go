package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint, p1 uint32, p2 uint32) uint {
    switch 9223372036854775807 {
    case 2147483647:
        p2 = 737
    }
    v5 := []int{-9223372036854775808}
    p2 = p2
    _ = len(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p0 << (((p0 % ((p0) | 1))) & 63))
}
func genFn1(p0 int8) int64 {
    var v6 uint = 18446744073709551615
    if false {
        var v7 int8 = p0
        _ = v7
    }
    _ = p0
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 518
    var v3 int16 = 1000
    v4 := make(map[string]bool)
    fmt.Println(v3, v1, v2)
    var v8 uint32 = 4294967295
    if (144 == 42) {
        fmt.Println(v1, v1)
        v4[(v1 + v1)] = false
        v1 = v1
    } else {
        if true {
            v3 -= v3
        } else {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint32 = 107
                delete(v4, v1)
                fmt.Println(v2)
                _ = fi9
                _ = v10
            }
        }
    }
    if true {
        fmt.Println(v3)
        v11 := []int{32767, 0, 2147483647, 32768}
        var v12 uint64 = 671
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 255, F1: 0}
    v1 = "hello"
    v4[v1] = false
    var v14 int8 = 0
    _ = len(v4)
    switch "foo bar" {
    case "":
        fmt.Println(v8, v14)
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v4[("" + "")] = false
        if true {
            var v16 uint64 = 4294967295
            var v17 float64 = ((0.1 - 1e10) + (0.1 / 3.14159))
            var v18 string = v1
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    fmt.Println(v2, v3, v3)
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint64 = 42
        v3 = v3
        var v21 float64 = 0.0
        _ = fi19
        _ = v20
        _ = v21
    }
    delete(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
