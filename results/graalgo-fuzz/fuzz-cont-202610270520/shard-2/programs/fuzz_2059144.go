package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 string, p1 int16) uint8 {
    for fi5 := 0; fi5 < 14; fi5++ {
        switch "a" {
        case "hello":
            v6 := []int{}
            _ = v6
        default:
            for fi7 := 0; fi7 < 6; fi7++ {
                switch p0 {
                case "\n\t":
                    v8 := []int{(fi7 | -986)}
                    _ = v8
                case "":
                    var v9 uint32 = 208
                    _ = v9
                case "a":
                    var v10 int16 = p1
                    _ = v10
                }
                _ = fi7
            }
        }
        _ = fi5
    }
    p0 = "\n\t"
    var v11 int32 = -1
    _ = p0
    _ = p1
    _ = v11
    return 218
}
func genFn1(p0 int64, p1 bool) uint64 {
    p0 -= 477
    if p1 {
        var v12 int64 = (p0 ^ p0)
        _ = v12
    }
    v13 := []int{-9223372036854775808}
    var v14 uint8 = 127
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 4294967295
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    v15 := GenS1{F0: 152}
    v1 += v1
    v16 := []int{-2147483648}
    var v17 int = 0
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 bool = (false && false)
        v16 = append(v16, 2147483647)
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 int64 = 0
            v1 += v1
            v21 = 0
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    var v22 string = "a"
    v23 := GenS2{F0: v1, F1: -1}
    v22 = v22
    var v24 float32 = (1e10 + (3.14159 / 1.0))
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint = 255
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
