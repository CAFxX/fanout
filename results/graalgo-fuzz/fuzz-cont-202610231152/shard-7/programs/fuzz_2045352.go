package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8) int {
    p0 = (p0 / (((p0 << ((943) & 63))) | 1))
    var v6 bool = false
    _ = p0
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    var v7 uint16 = 1
    v7 -= (^(v7))
    switch 2147483647 {
    case 32768:
        for fi8 := 0; fi8 < 1; fi8++ {
            v7 += v7
            for fi9 := 0; fi9 < 5; fi9++ {
                v7 += (v7 % ((v7) | 1))
                _ = fi9
            }
            _ = fi8
        }
    case 127:
        var v10 int32 = -466
        _ = v10
    }
    v7 = (^(v7))
    if ((611 <= 193) && (!true)) {
        var v11 uint32 = 42
        if true {
            v12 := GenS1{F0: 32767, F1: false}
            _ = v12
        }
        if (v11 > v11) {
            v13 := GenS2{F0: -105}
            _ = v13
        } else {
            switch 127 {
            case 42:
                v7 += (^(v7))
            }
            fmt.Println(v7, v7)
        }
        _ = v11
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := []int{fi14, 32768, genFn0(-49), (-(fi14))}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        if ((!(-893 > -1)) && false) {
            v17 := []int{230, (fi16 + (fi16 - 2147483647))}
            var v18 int16 = -1
            _ = v17
            _ = v18
        }
        var v19 string = ("foo bar" + "hello")
        _ = fi16
        _ = v19
    }
    v7 -= 2
    v4["a"] = 2
    var v20 uint8 = 128
    v21 := GenS1{F0: 767, F1: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
