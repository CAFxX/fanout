package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 string }
type GenS3 struct { F0 uint32; F1 uint16 }

func main() {
    var v1 int32 = -989
    var v2 int32 = 42
    var v3 uint32 = 2
    v4 := make(map[int]bool)
    var v5 int64 = 9223372036854775807
    var v6 int8 = -67
    v5 += (v5 & 9223372036854775807)
    var v7 uint = 2
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := []int{fi8, (fi8 | 128), fi8}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        if ((!((42 != (-(v5))) && true)) && true) {
            v11 := []int{-2147483648, 896}
            fmt.Println(v6, v7)
            for fi12 := 0; fi12 < 11; fi12++ {
                if false {
                    var v13 int = (fi12 * fi12)
                    _ = v13
                } else {
                    v14 := GenS1{F0: 65535, F1: (v3 & 2147483648), F2: -1}
                    var v15 bool = true
                    v3 += (v3 & (v3 + (v3 >> (((v7 ^ v7)) & 63))))
                    _ = len(v11)
                    _ = v14
                    _ = v15
                }
                fmt.Println(fi12, fi12, v1)
                _ = fi12
            }
            _ = v11
        }
        v16 := GenS1{F0: 42, F1: 1, F2: (-(v6))}
        _ = fi10
        _ = v16
    }
    v7 = v7
    var v17 float64 = (3.14159 - 3.14159)
    for fi18 := 0; fi18 < 15; fi18++ {
        if true {
            v19 := GenS2{F0: 4294967295, F1: 691, F2: ""}
            _ = v19
        } else {
            v20 := GenS1{F0: 2, F1: (v3 / (((v3 - 0)) | 1)), F2: v6}
            v5 = v5
            fmt.Println(v6, v3)
            _ = v20
        }
        _ = fi18
    }
    if false {
        switch "\n\t" {
        case "\n\t":
            v2 += v1
        case "a":
            v21 := GenS1{F0: 42, F1: 495, F2: v6}
            _ = v21
        }
    }
    var v22 uint8 = 127
    if (3.14159 > v17) {
        v23 := GenS1{F0: 255, F1: v3, F2: 1}
        v22 += (v22 & v22)
        v24 := GenS2{F0: 2, F1: 32767, F2: ("foo bar" + "\n\t")}
        _ = v23
        _ = v24
    } else {
        if (1 < 42) {
            var v25 int64 = 1
            fmt.Println(v3, v2)
            switch 0 {
            case 26:
                v26 := v4[32767]
                _ = v26
            }
            _ = v25
        }
        v17 = 0.1
        switch 42 {
        case -2147483648:
            v27 := GenS3{F0: (v3 + v3), F1: 32767}
            _ = v27
        case 1:
            v28 := []int{128, 9223372036854775807, 2147483647, 632}
            _ = v28
        case 4294967295:
            var v29 float32 = 1e10
            _ = v29
        }
    }
    fmt.Println(v3, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
