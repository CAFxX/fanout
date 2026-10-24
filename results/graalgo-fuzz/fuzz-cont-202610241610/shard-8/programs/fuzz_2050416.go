package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint = 255
    v2 := make(map[uint32]int)
    var v3 float32 = ((1e10 + 1e10) * 0.1)
    var v4 int16 = 0
    var v5 float32 = v3
    v1 = v1
    var v6 float64 = 0.1
    if true {
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := []int{32768}
            _ = fi7
            _ = v8
        }
    } else {
        v6 = 0.0
        var v9 uint32 = 2147483648
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: 49, F1: 152}
            if false {
                fmt.Println(v3, fi10)
                var v12 uint = 988
                var v13 uint32 = 42
                _ = v12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    fmt.Println(v3, v4)
    v14 := GenS2{F0: 198, F1: 127, F2: 0}
    v15 := []int{-9223372036854775808, 9223372036854775807}
    var v16 uint = v1
    v1 += (v1 % ((2) | 1))
    switch 4294967295 {
    case 128:
        fmt.Println(v6)
    case 127:
        if false {
            v17 := GenS3{F0: 0}
            var v18 uint16 = 843
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 18; fi19++ {
                fmt.Println(v16, v1, v1)
                v1 = (^(v16))
                v20 := GenS2{F0: 128, F1: 127, F2: 2}
                _ = fi19
                _ = v20
            }
        }
    }
    fmt.Println(v1)
    v2[65535] = 42
    v4 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
