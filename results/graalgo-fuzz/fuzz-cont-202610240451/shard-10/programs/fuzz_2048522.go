package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 bool; F1 bool; F2 uint32 }

func main() {
    var v1 int8 = 0
    var v2 uint = 4294967295
    v3 := make(map[string]int)
    var v4 uint16 = 855
    v2 += v2
    fmt.Println(v1, v2, v4)
    var v5 int8 = -128
    switch -2147483648 {
    case 9223372036854775807:
        if true {
            var v6 uint32 = 2
            _ = v6
        }
    case 32767:
        v7 := []int{4294967295, 4294967295}
        _ = v7
    default:
        var v8 uint64 = 42
        _ = v8
    }
    if (1 > 415) {
        var v9 float32 = (float32(float64(float32(0.0))) + 0.0)
        _ = v9
    }
    var v10 string = "a"
    var v11 int8 = v1
    for fi12 := 0; fi12 < 5; fi12++ {
        switch v10 {
        case "\n\t":
            var v13 float32 = 1.5
            _ = v13
        case "hello":
            var v14 string = v10
            _ = v14
        default:
            if true {
                fmt.Println(v2, v1)
                v10 = (v10 + v10)
                v15 := GenS1{F0: 109, F1: -1, F2: 0}
                _ = v15
            } else {
                var v16 bool = true
                _ = v16
            }
        }
        if false {
            if false {
                v17 := []int{(fi12 - fi12), (fi12 << (((v2 % (((v2 & 601)) | 1))) & 63)), fi12}
                fmt.Println(fi12, v1)
                _ = v17
            }
        } else {
            v18 := GenS2{F0: (-(fi12)), F1: (^(v4))}
            for fi19 := 0; fi19 < 12; fi19++ {
                fmt.Println(v2, v5)
                if (42 > 4294967295) {
                    fmt.Println(fi12, v11)
                }
                v20 := GenS1{F0: 65535, F1: -173, F2: 2147483647}
                _ = fi19
                _ = v20
            }
            if false {
                v21 := []int{}
                _ = v21
            } else {
                var v22 string = v10
                v23 := GenS3{F0: true, F1: (134 <= 168), F2: 65535}
                var v24 uint16 = (v4 % ((v4) | 1))
                _ = v22
                _ = v23
                _ = v24
            }
            _ = v18
        }
        _ = fi12
    }
    v4 -= 498
    v11 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
