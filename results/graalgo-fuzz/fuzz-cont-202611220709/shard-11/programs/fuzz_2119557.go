package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0(p0 int64) int16 {
    v5 := []int{-1, 316, 128}
    fmt.Println(p0)
    var v6 int = -9223372036854775808
    v7 := GenS1{F0: 0, F1: 42}
    v5 = append(v5, -878)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint64, p1 uint64) int32 {
    p0 -= 91
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := []int{(fi8 | (^(fi8))), 42}
        if false {
            if false {
                for fi10 := 0; fi10 < 10; fi10++ {
                    var v11 uint = 0
                    v12 := GenS1{F0: (^(fi10)), F1: 2147483647}
                    p0 += p1
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                p0 = (^(p1))
            }
            p0 += (^(p1))
            var v13 int64 = 9223372036854775807
            _ = v13
        } else {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        if true {
            var v14 float32 = float32((1e10 + 0.0))
            fmt.Println(p1)
            v15 := GenS2{F0: true, F1: -66}
            _ = v14
            _ = v15
        } else {
            var v16 uint64 = p1
            _ = v16
        }
        _ = fi8
        _ = v9
    }
    v17 := []int{1, 42}
    v17 = append(v17, 165)
    _ = p0
    _ = p1
    _ = v17
    return 284
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]uint32)
    var v3 int8 = 127
    var v4 string = "a"
    _ = len(v2)
    fmt.Println(v4)
    v3 = v3
    var v18 string = "a"
    v19 := v1[127]
    v4 = v18
    _ = len(v1)
    if true {
        switch v4 {
        case "a":
            v20 := []int{}
            _ = v20
        case "\n\t":
            if (false && v19) {
                var v21 uint64 = 1
                _ = v21
            } else {
                var v22 uint8 = 2
                fmt.Println(v22, v4, v3)
                var v23 bool = v19
                _ = v22
                _ = v23
            }
        case "hello":
            switch 32767 {
            case -743:
                v24 := v1[32767]
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 int16 = 76
                v27 := GenS3{F0: 2147483648, F1: 65535, F2: (v26 + genFn0(-9223372036854775808))}
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        v28 := []int{}
        _ = v28
    }
    if false {
        switch -264 {
        case 32768:
            v19 = false
        case -1:
            for fi29 := 0; fi29 < 14; fi29++ {
                fmt.Println(fi29, v3, v18)
                _ = fi29
            }
        default:
            v19 = (!v19)
        }
    }
    fmt.Println(v3, v4)
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
