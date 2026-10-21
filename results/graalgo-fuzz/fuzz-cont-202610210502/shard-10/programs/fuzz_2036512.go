package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int, p2 int16) uint64 {
    switch "hello" {
    case "hello":
        v7 := []int{(p1 ^ (p1 % ((p1) | 1)))}
        _ = v7
    case "\n\t":
        v8 := []int{p1, p1, (p1 * 0)}
        _ = v8
    }
    v9 := []int{127, 0, 42, p1}
    p0 -= 1
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p0
}

func main() {
    var v1 bool = false
    var v2 string = "\n\t"
    v3 := make(map[uint32]int)
    v4 := []string{("a" + ""), "\n\t", "\n\t", v2}
    var v5 uint8 = 128
    var v6 float32 = ((-2.25 + float32((3.14159 * -2.25))) - 0.0)
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            if v1 {
                v12 := GenS1{F0: -1, F1: (fi11 - fi11)}
                _ = v12
            }
            v6 += float32((3.14159 * 0.1))
            _ = fi11
        }
        switch v2 {
        case "":
            v5 -= (v5 ^ v5)
        case "hello":
            v13 := GenS1{F0: -9223372036854775808, F1: (^(fi10))}
            _ = v13
        }
        _ = fi10
    }
    v14 := GenS2{F0: 255}
    v4 = append(v4, v2)
    switch -789 {
    case 127:
        if v1 {
            if v1 {
                v6 = ((float32(0.0) / ((3.14159 * float32(0.0)) * v6)) * v6)
                v15 := GenS2{F0: 2}
                _ = v15
            } else {
                v16 := GenS2{F0: 970}
                _ = v16
            }
            for fi17 := 0; fi17 < 8; fi17++ {
                delete(v3, 634)
                _ = fi17
            }
        }
    case -166:
        fmt.Println(v1, v5, v6)
    }
    if (v1 && (90 < 0)) {
        var v18 int64 = 9223372036854775807
        var v19 float32 = 0.0
        var v20 int = 4294967295
        _ = v18
        _ = v19
        _ = v20
    }
    _ = len(v3)
    v2 = (v2 + (v2 + v2))
    v5 = v5
    v21 := v3[2]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
