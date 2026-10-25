package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int64 }
type GenS2 struct { F0 uint64; F1 int32; F2 int64 }
type GenS3 struct { F0 uint }
func genFn0() bool {
    var v6 int64 = 1
    v6 -= v6
    _ = v6
    return true
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    var v5 float64 = 1e10
    var v7 string = "hello"
    for fi8 := 0; fi8 < 20; fi8++ {
        _ = len(v1)
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS1{F0: 2, F1: 2, F2: -182}
            switch "\n\t" {
            case "a":
                v7 = "hello"
            default:
                var v11 uint8 = 42
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            v5 += v5
            switch (v7 + (("a" + "a") + "a")) {
            case "":
                v5 = v5
            case "\n\t":
                if false {
                    fmt.Println(v7, v5, fi12)
                    var v13 int32 = 42
                    var v14 uint32 = 2
                    var v15 uint32 = v14
                    v5 = float64(1e10)
                    _ = v13
                    _ = v14
                    _ = v15
                }
            }
            v5 += v5
            _ = fi12
        }
        _ = fi8
    }
    fmt.Println(v5, v5)
    v16 := []int{-680, 128, -700, 557}
    for fi17 := 0; fi17 < 12; fi17++ {
        if (!false) {
            for fi18 := 0; fi18 < 20; fi18++ {
                if len(v16) > 0 {
                    v16[0] = fi18
                }
                _ = fi18
            }
            var v19 uint64 = 1
            var v20 int64 = -1
            _ = v19
            _ = v20
        } else {
            var v21 uint16 = 0
            _ = v21
        }
        if (("\n\t" <= "") || false) {
            switch v7 {
            case "a":
                v22 := GenS2{F0: 2, F1: 2147483647, F2: 9223372036854775807}
                _ = v22
            case "\n\t":
                var v23 uint8 = 31
                _ = v23
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                fmt.Println(v5)
                v25 := v2[(fi17 + 127)]
                var v26 bool = false
                _ = fi24
                _ = v25
                _ = v26
            }
            v27 := v2[fi17]
            _ = v27
        } else {
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS2{F0: 1, F1: 2147483647, F2: 42}
                fmt.Println(v5)
                _ = fi28
                _ = v29
            }
        }
        _ = fi17
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v5 = v5
        v1[v7] = (fi30 / ((fi30) | 1))
        _ = fi30
    }
    v31 := GenS3{F0: 27}
    fmt.Println(v5, v7)
    v7 = (("a" + "hello") + v7)
    var v32 int16 = 311
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v7
}
