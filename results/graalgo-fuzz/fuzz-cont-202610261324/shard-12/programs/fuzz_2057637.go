package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 int16, p1 int16) uint16 {
    p0 += p1
    var v4 int = -333
    switch "hello" {
    case "a":
        var v5 int16 = p1
        _ = v5
    case "\n\t":
        v6 := []int{(v4 >> ((65535) & 63)), v4, v4, v4}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]bool)
    var v3 int8 = 30
    for fi7 := 0; fi7 < 6; fi7++ {
        v3 += (v3 - -15)
        _ = fi7
    }
    v1 += (v1 & v1)
    switch "\n\t" {
    case "a":
        fmt.Println(v3)
    case "hello":
        if true {
            v8 := GenS1{F0: 18446744073709551615}
            var v9 int32 = 2147483647
            if (2 != 987) {
                if (255 >= 2) {
                    v10 := GenS2{F0: 128, F1: 42}
                    _ = v10
                } else {
                    if false {
                        v1 = 32767
                    }
                    for fi11 := 0; fi11 < 14; fi11++ {
                        var v12 uint32 = 0
                        _ = fi11
                        _ = v12
                    }
                }
                v9 += (^(v9))
            }
            _ = v8
            _ = v9
        }
    default:
        v13 := []int{32767, 128, 127}
        _ = v13
    }
    v3 = -128
    v14 := []int{-2147483648}
    var v15 uint32 = 152
    v16 := v2[2147483648]
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    switch "hello" {
    case "\n\t":
        switch -2147483648 {
        case 32767:
            v17 := v2[65535]
            _ = v17
        case 127:
            var v18 bool = v16
            _ = v18
        case 23:
            var v19 int64 = 9223372036854775807
            _ = v19
        default:
            var v20 float32 = float32(1e10)
            _ = v20
        }
    case "a":
        v21 := GenS3{F0: true, F1: v15}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
