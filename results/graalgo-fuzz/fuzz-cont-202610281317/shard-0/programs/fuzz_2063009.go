package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    var v1 int64 = 42
    v2 := []string{"", "hello"}
    v3 := make(map[int]string)
    v1 -= v1
    v4 := GenS1{F0: 255, F1: 22, F2: 148}
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: 128, F1: 1, F2: 0}
        for fi7 := 0; fi7 < 14; fi7++ {
            if true {
                if (!false) {
                    if (!(!(!(true && false)))) {
                        var v8 int = (fi7 << ((0) & 63))
                        _ = v8
                    }
                    var v9 uint16 = 255
                    fmt.Println(v1, v1)
                    _ = v9
                }
                v10 := GenS1{F0: 1, F1: 42, F2: 4294967295}
                v11 := GenS2{F0: (361 == 255), F1: 42}
                _ = v10
                _ = v11
            }
            fmt.Println(fi7, fi7, fi5)
            _ = fi7
        }
        var v12 uint64 = 372
        _ = fi5
        _ = v12
        _ = v6
    }
    var v13 float32 = 1e10
    switch "" {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    case "a":
        v1 = -709
    }
    v13 = (v13 + 0.1)
    v1 = v1
    var v14 float64 = 0.0
    fmt.Println(v1, v13, v13)
    _ = len(v2)
    v15 := GenS1{F0: 0, F1: 402, F2: 42}
    v16 := GenS3{F0: 42, F1: 1}
    var v17 float64 = (v14 * (v14 / float64(0.1)))
    v13 -= ((1e10 - v13) / (((-2.25 - v13) - float32(v17)) * v13))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
