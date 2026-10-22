package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint) int32 {
    if true {
        switch 128 {
        case -1:
            switch "foo bar" {
            case "hello":
                if true {
                    var v6 int64 = 966
                    _ = v6
                }
            default:
                for fi7 := 0; fi7 < 16; fi7++ {
                    var v8 uint64 = 1
                    _ = fi7
                    _ = v8
                }
            }
        case -390:
            var v9 uint = p1
            _ = v9
        }
    }
    v10 := GenS1{F0: 42, F1: 888, F2: true}
    v11 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 42
}
func genFn1(p0 uint8) int64 {
    switch "" {
    case "":
        p0 += p0
    }
    var v12 float64 = ((float64(0.1) + 0.1) - (0.0 - -2.25))
    fmt.Println(p0)
    v13 := []int{}
    _ = p0
    _ = v12
    _ = v13
    return -952
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 222
    v3 := make(map[string]bool)
    var v4 uint8 = 1
    v5 := make(map[uint32]string)
    var v14 int8 = 127
    v4 = (v4 >> ((v2) & 63))
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 bool = (!(!true))
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: -529, F1: 178, F2: (62 < 32767)}
    var v18 int = 176
    v19 := []int{v18}
    if false {
        var v20 uint = (v2 - (^(v2)))
        _ = v20
    }
    v21 := GenS2{F0: 363}
    switch (v18 << (((v2 >> ((891) & 63))) & 63)) {
    case -325:
        var v22 uint8 = (^(v4))
        _ = v22
    case -240:
        v2 += v2
    default:
        var v23 float32 = ((1.5 * 0.1) - (-2.25 - -2.25))
        _ = v23
    }
    v2 = (v2 % ((v2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
