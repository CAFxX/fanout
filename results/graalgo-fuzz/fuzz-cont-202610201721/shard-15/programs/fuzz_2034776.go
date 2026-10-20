package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    var v3 int8 = -1
    var v4 uint64 = 1
    fmt.Println(v4, v4)
    var v5 string = "a"
    v4 = v4
    if (true && (255 != 2)) {
        var v6 uint8 = 0
        v7 := GenS1{F0: 0, F1: (v6 / ((v6) | 1)), F2: v6}
        _ = v6
        _ = v7
    } else {
        v8 := []int{-1, 837, 13}
        _ = v8
    }
    if false {
        v4 = 0
    } else {
        v9 := GenS1{F0: 1, F1: 127, F2: 255}
        _ = v9
    }
    if true {
        v4 = (^(v4))
        var v10 int32 = 0
        if true {
            fmt.Println(v3, v5, v3)
            if false {
                var v11 bool = ((0 != 128) && true)
                v5 = ((v5 + v5) + v5)
                _ = v11
            } else {
                v12 := GenS2{F0: 130, F1: 42}
                v5 = (v5 + "\n\t")
                _ = v12
            }
            v13 := GenS3{F0: v5, F1: v4, F2: v4}
            _ = v13
        }
        _ = v10
    }
    var v14 int = 128
    var v15 uint = 0
    var v16 int32 = -32
    v17 := GenS4{F0: (false || (int32(v4) < (^(v16))))}
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 int16 = int16(v4)
        _ = fi18
        _ = v19
    }
    var v20 int64 = 42
    switch (v14 | (v14 ^ v14)) {
    case -2147483648:
        if false {
            v14 += v14
            v21 := []int{32768, (v14 << ((v15) & 63)), 2147483647, (v14 * -949)}
            _ = v21
        } else {
            var v22 int64 = (v20 & (v20 / ((42) | 1)))
            var v23 float64 = 0.1
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := []int{v14, -587, (v14 ^ 32767), (fi24 / (((-(v14))) | 1))}
                if (!(!false)) {
                    var v26 int16 = 0
                    _ = v26
                } else {
                    _ = len(v1)
                    var v27 float32 = 1.5
                    var v28 uint64 = (v4 | (v4 % (((v4 % ((v4) | 1))) | 1)))
                    _ = v27
                    _ = v28
                }
                switch "foo bar" {
                case "hello":
                    v14 = v14
                }
                _ = fi24
                _ = v25
            }
            _ = v22
            _ = v23
        }
    }
    var v29 int64 = (v20 % ((v20) | 1))
    v5 = v5
    v3 = (v3 | (v3 % ((v3) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
