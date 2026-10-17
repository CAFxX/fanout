package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint8; F1 uint; F2 uint64 }
func genFn0() int64 {
    var v4 float32 = (((0.1 * (0.1 + 0.1)) * (float32(1e10) / 3.14159)) + float32(1e10))
    v4 += (((v4 / v4) + v4) - ((v4 - float32(((((1e10 - 1.5) / (3.14159 - (0.1 + 1.5))) + 3.14159) - float64((3.14159 * -2.25))))) * v4))
    switch 32767 {
    case -1:
        fmt.Println(v4)
    default:
        v5 := []int{598, 9223372036854775807, -351}
        _ = v5
    }
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 810
    var v2 int = 32767
    var v3 int64 = -308
    var v6 uint = 837
    v6 = uint(v1)
    for fi7 := 0; fi7 < 13; fi7++ {
        switch "a" {
        case "a":
            v8 := GenS1{F0: ((42 == -128) && true), F1: (fi7 * v2)}
            _ = v8
        case "hello":
            var v9 int = 0
            _ = v9
        }
        _ = fi7
    }
    v3 = genFn0()
    fmt.Println(v6)
    var v10 int16 = 379
    if (v3 > (v3 ^ (v3 + genFn0()))) {
        if (false || false) {
            v11 := []int{v2, (-(v2))}
            v10 = 445
            _ = v11
        } else {
            for fi12 := 0; fi12 < 1; fi12++ {
                fmt.Println(v2, v2, v6)
                _ = fi12
            }
            if true {
                v13 := GenS1{F0: (156 <= 128), F1: 32768}
                v14 := []int{32768, (v2 & v2)}
                for fi15 := 0; fi15 < 20; fi15++ {
                    fmt.Println(v2, v2, v2)
                    v16 := GenS1{F0: (false && true), F1: (fi15 % (((fi15 ^ fi15)) | 1))}
                    _ = fi15
                    _ = v16
                }
                _ = v13
                _ = v14
            }
            v17 := []int{v2, (v2 + (v2 | (v2 - (^(v2))))), (^(v2))}
            _ = v17
        }
    }
    var v18 uint64 = 0
    v6 = (v6 + 0)
    v1 = v1
    for fi19 := 0; fi19 < 3; fi19++ {
        if true {
            v20 := GenS2{F0: 42, F1: 0}
            v21 := GenS3{F0: 127, F1: v6, F2: (^(v18))}
            _ = v20
            _ = v21
        } else {
            v10 = (-(v10))
            switch (v2 | 523) {
            case -1:
                fmt.Println(v10, v6, v6)
            }
        }
        _ = fi19
    }
    v6 += (v6 / ((v6) | 1))
    v2 = 917
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v6
}
