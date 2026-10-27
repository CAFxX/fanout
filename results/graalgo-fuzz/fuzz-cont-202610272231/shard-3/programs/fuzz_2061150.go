package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 string }

func main() {
    v1 := make(map[string]string)
    v2 := []int{9223372036854775807, 262}
    var v3 int8 = 1
    var v4 int64 = -9223372036854775808
    v4 = (^(v4))
    if len(v2) > 0 {
        v2[0] = 128
    }
    var v5 uint64 = 485
    v6 := GenS1{F0: -539}
    v7 := GenS2{F0: (^(v3)), F1: 2147483647}
    var v8 uint = 613
    switch -547 {
    case -2147483648:
        var v9 uint32 = 42
        _ = v9
    case 32768:
        var v10 int16 = 42
        _ = v10
    case 127:
        for fi11 := 0; fi11 < 14; fi11++ {
            v3 -= (v3 | (v3 & -1))
            v3 = v3
            if ((!true) || true) {
                for fi12 := 0; fi12 < 14; fi12++ {
                    var v13 uint64 = (v5 & v5)
                    fmt.Println(fi12, v3)
                    fmt.Println(fi12, v8)
                    if (0.1 <= 0.1) {
                        var v14 uint16 = uint16(v8)
                        _ = v14
                    }
                    _ = fi12
                    _ = v13
                }
                if true {
                    switch "\n\t" {
                    case "hello":
                        var v15 int = fi11
                        _ = v15
                    case "a":
                        var v16 uint64 = (^(v5))
                        _ = v16
                    }
                } else {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        var v18 int32 = -1
                        _ = fi17
                        _ = v18
                    }
                    v5 = (v5 & v5)
                }
            } else {
                if true {
                    delete(v1, "\n\t")
                    v3 -= v3
                    var v19 int16 = -988
                    _ = v19
                }
                v20 := GenS3{F0: (v8 + v8), F1: 1, F2: "a"}
                _ = v20
            }
            _ = fi11
        }
    default:
        var v21 uint = (v8 / ((4294967295) | 1))
        _ = v21
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
