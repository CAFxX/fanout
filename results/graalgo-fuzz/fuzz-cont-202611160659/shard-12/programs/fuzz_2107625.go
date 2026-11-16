package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 uint32 }
func genFn0(p0 uint64, p1 uint8, p2 bool) uint64 {
    switch -1 {
    case 2147483647:
        if (0.1 == 1.5) {
            v4 := []int{32767, 32768}
            for fi5 := 0; fi5 < 8; fi5++ {
                p2 = true
                v6 := GenS1{F0: -487, F1: 4294967295}
                _ = fi5
                _ = v6
            }
            p2 = (1 <= 18446744073709551615)
            _ = v4
        }
    case -655:
        var v7 uint = 598
        _ = v7
    case 966:
        switch "foo bar" {
        case "a":
            p0 -= 4294967295
        default:
            var v8 int8 = 70
            _ = v8
        }
    default:
        var v9 uint64 = (^(p0))
        _ = v9
    }
    p0 = p0
    fmt.Println(p2, p2)
    if (!p2) {
        if p2 {
            p0 = p0
            fmt.Println(p2, p1, p2)
        } else {
            fmt.Println(p0, p1, p1)
            for fi10 := 0; fi10 < 19; fi10++ {
                if p2 {
                    p0 += 131
                    v11 := GenS2{F0: -32768, F1: 165, F2: 1}
                    _ = v11
                }
                fmt.Println(fi10, p0)
                _ = fi10
            }
            switch (("a" + "a") + "\n\t") {
            case "hello":
                v12 := GenS2{F0: 1000, F1: uint16(p1), F2: 65535}
                _ = v12
            }
        }
        p0 += 2
    } else {
        switch -954 {
        case 107:
            v13 := []int{127, 629}
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            p1 = (p1 % (((p1 % ((p1) | 1))) | 1))
            _ = fi14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (^(p0))
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]bool)
    var v3 int64 = -1
    fmt.Println(v1, v3)
    v15 := []int{9223372036854775807, 9223372036854775807, 9223372036854775807, -2147483648}
    v3 -= (v3 ^ v3)
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(v1, fi16, v1)
        _ = fi16
    }
    var v17 int = -590
    var v18 int = (v17 & v17)
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float32 = 0.1
        v20 = (v20 / -2.25)
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi19
        _ = v20
    }
    var v21 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
}
