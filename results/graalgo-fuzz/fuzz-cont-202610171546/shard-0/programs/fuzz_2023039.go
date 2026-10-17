package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 int16, p1 uint8, p2 int32) int64 {
    if true {
        var v5 int32 = p2
        p1 = p1
        _ = v5
    } else {
        switch "a" {
        case "":
            v6 := GenS1{F0: 32767, F1: p1}
            _ = v6
        case "a":
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 uint16 = 42
                var v9 int32 = 1
                v9 = (p2 % (((p2 - v9)) | 1))
                _ = fi7
                _ = v8
                _ = v9
            }
        case "hello":
            fmt.Println(p0, p1, p0)
        }
        p1 -= p1
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 131
}

func main() {
    v1 := []int{}
    v2 := []uint32{42, 895, 1}
    var v3 uint16 = 1
    v4 := []uint32{2147483647, 0, 65535}
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: 42, F1: 66}
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if (0.0 > (0.1 - 1.5)) {
        v3 = (v3 & (v3 ^ v3))
    }
    v13 := GenS1{F0: 195, F1: 1}
    fmt.Println(v3)
    var v14 uint = 2
    if (1e10 < float64(0.0)) {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v15 uint8 = 42
    var v16 uint = uint(v15)
    v17 := GenS2{F0: 1, F1: 717}
    var v18 string = "\n\t"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
