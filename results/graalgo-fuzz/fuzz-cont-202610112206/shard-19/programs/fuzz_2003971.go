package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 string; F2 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    var v3 float32 = ((1e10 + (float32(0.1) - 1e10)) + 0.1)
    var v4 int8 = 127
    v5 := make(map[string]int)
    v6 := make(map[uint32]uint32)
    var v7 int8 = v4
    v8 := GenS1{F0: -453, F1: 9}
    if (32767 != 128) {
        v9 := GenS1{F0: 42, F1: 37}
        v10 := GenS1{F0: -2147483648, F1: 42}
        _ = v10
        _ = v9
    } else {
        fmt.Println(v7, v3, v7)
        for fi11 := 0; fi11 < 1; fi11++ {
            if true {
                v2 = append(v2, "hello")
                v7 = v4
                var v12 int = fi11
                _ = v12
            }
            var v13 int = (fi11 << ((930) & 63))
            _ = fi11
            _ = v13
        }
        v7 -= 42
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if (!false) {
            v15 := GenS1{F0: 1, F1: 143}
            switch (fi14 * 0) {
            case -982:
                v3 = -2.25
            }
            _ = v15
        }
        v16 := GenS1{F0: 42, F1: 42}
        _ = fi14
        _ = v16
    }
    var v17 int = -594
    v7 += v7
    v4 = (v7 << ((42) & 63))
    var v18 uint8 = 219
    if false {
        v18 = (v18 + v18)
    }
    if false {
        var v19 uint16 = 1
        var v20 bool = false
        _ = v19
        _ = v20
    }
    if ((23 > 2) || false) {
        v21 := GenS1{F0: -36, F1: 128}
        v22 := GenS2{F0: true}
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 uint = 255
        switch (v17 - (fi23 - (-(fi23)))) {
        case 128:
            v18 = 0
        case -1:
            var v25 bool = ((false && ((fi23 < 42) && (fi23 >= fi23))) || (!(42 > 1)))
            _ = v25
        case 9223372036854775807:
            v26 := GenS3{F0: 1, F1: "", F2: 2}
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
