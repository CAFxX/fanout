package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int16; F2 bool }

func main() {
    var v1 int8 = -55
    var v2 uint = 4294967295
    var v3 int = 32768
    v4 := []int{v3}
    var v5 int16 = -1
    if false {
        v6 := GenS1{F0: 1, F1: 305}
        _ = v6
    } else {
        v1 -= (v1 ^ (v1 ^ 127))
        if (false && true) {
            v5 -= (v5 % ((v5) | 1))
        }
    }
    var v7 uint32 = 1
    fmt.Println(v1, v2, v5)
    for fi8 := 0; fi8 < 4; fi8++ {
        v2 -= 726
        v9 := GenS1{F0: 618, F1: v7}
        _ = fi8
        _ = v9
    }
    if false {
        v2 -= 65535
    }
    var v10 float64 = 0.1
    var v11 float64 = (((v10 * v10) / 1.0) / v10)
    switch "a" {
    case "foo bar":
        fmt.Println(v3)
    }
    v12 := GenS2{F0: v7}
    var v13 float32 = 0.0
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 float32 = ((1.5 / (((((v13 * ((1.5 - (v13 + 3.14159)) * (v13 * v13))) * v13) - v13) * 3.14159) + 0.1)) + 3.14159)
        v16 := GenS1{F0: 42, F1: v7}
        v13 -= ((0.0 + v13) + float32((v10 * float64(v13))))
        _ = fi14
        _ = v15
        _ = v16
    }
    if (0 != 42) {
        v10 -= float64(float32((v11 / v11)))
        _ = len(v4)
    } else {
        v4 = append(v4, v3)
        v4 = append(v4, v3)
        v7 += v7
    }
    v13 += v13
    if (false && true) {
        for fi17 := 0; fi17 < 13; fi17++ {
            v7 = 1
            switch "hello" {
            case "hello":
                fmt.Println(fi17, v11, v11)
            }
            fmt.Println(v13, v10)
            _ = fi17
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            if ((v2 - (v2 >> (((v2 ^ (v2 - (^(v2))))) & 63))) >= (v2 + (v2 << ((721) & 63)))) {
                var v19 int = fi18
                _ = v19
            }
            _ = fi18
        }
        v20 := GenS3{F0: (v7 / ((945) | 1)), F1: 163, F2: (127 >= 42)}
        _ = v20
    }
    var v21 int32 = int32(v5)
    fmt.Println(v2, v13, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
