package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint }
func genFn0(p0 string, p1 uint32, p2 int) int {
    switch -9223372036854775808 {
    case 42:
        p1 = p1
    default:
        fmt.Println(p2, p2, p0)
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 bool = (false && ((true && ((255 < 79) && true)) || false))
        _ = fi7
        _ = v8
    }
    var v9 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 4294967295
    v3 := []int{-2147483648, 460, 4294967295, 42}
    var v4 uint64 = 42
    v5 := make(map[uint32]string)
    v6 := []int{-51, 32768, 2147483647, 128}
    v2 -= (^(v2))
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: (v4 << ((v2) & 63))}
    v13 := GenS1{F0: v4}
    v4 = (v4 << (((v2 * (v2 % ((v2) | 1)))) & 63))
    var v14 uint = 2
    for fi15 := 0; fi15 < 14; fi15++ {
        if false {
            var v16 string = ("foo bar" + "a")
            fmt.Println(fi15, fi15)
            _ = v16
        } else {
            v17 := GenS2{F0: 212, F1: 100, F2: 579}
            _ = v17
        }
        _ = fi15
    }
    v3 = append(v3, 127)
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 uint32 = 2147483647
        _ = fi18
        _ = v19
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
