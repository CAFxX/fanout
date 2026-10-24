package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int32 }
func genFn0(p0 bool) int32 {
    p0 = (p0 || (p0 && (712 != 0)))
    var v7 int32 = -101
    fmt.Println(p0, p0)
    _ = p0
    _ = v7
    return v7
}
func genFn1(p0 uint32) int {
    p0 = p0
    var v8 int8 = 127
    v8 -= 0
    _ = p0
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -615
    var v3 int64 = -9223372036854775808
    v4 := make(map[int]bool)
    var v5 int32 = v2
    v6 := []uint32{42}
    switch "a" {
    case "hello":
        if false {
            for fi9 := 0; fi9 < 7; fi9++ {
                for fi10 := 0; fi10 < 1; fi10++ {
                    v11 := GenS1{F0: 42, F1: 18446744073709551615, F2: (v2 * genFn0(false))}
                    var v12 uint = 2
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v5 -= (v2 - v5)
                _ = fi9
            }
        } else {
            var v13 int = -590
            fmt.Println(v2)
            _ = v13
        }
    case "":
        switch 794 {
        case -752:
            delete(v4, 4294967295)
        case 32767:
            v4[42] = true
        }
    }
    v6 = append(v6, 65535)
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint = 18446744073709551615
        _ = fi14
        _ = v15
    }
    switch 824 {
    case 42:
        var v16 float64 = 1.5
        _ = v16
    }
    v6 = append(v6, 350)
    v3 -= (v3 + 9223372036854775807)
    v3 = v3
    switch int(v2) {
    case 127:
        v17 := GenS1{F0: 42, F1: 1, F2: (v2 ^ (v2 ^ (-(v2))))}
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
