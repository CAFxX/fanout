package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{328, 2147483647}
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    var v5 uint = 698
    if (42 <= -105) {
        v6 := GenS1{F0: 2, F1: 32767, F2: -1}
        for fi7 := 0; fi7 < 6; fi7++ {
            if false {
                v8 := GenS1{F0: 4294967295, F1: -861, F2: 42}
                fmt.Println(fi7, v1, fi7)
                switch fi7 {
                case 219:
                    for fi9 := 0; fi9 < 7; fi9++ {
                        var v10 int8 = 100
                        v5 = v5
                        var v11 int = -45
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case 128:
                    v1 += v1
                }
                _ = v8
            }
            _ = fi7
        }
        _ = v6
    }
    if true {
        var v12 float32 = (1.5 + (1e10 / 1.5))
        _ = v12
    }
    v1 += v1
    v4["foo bar"] = "hello"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v13 := GenS2{F0: -762, F1: -1}
    var v14 float64 = (0.1 * 1.5)
    v15 := GenS3{F0: int32(v1)}
    var v16 float64 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
