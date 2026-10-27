package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0() int32 {
    switch 198 {
    case 128:
        v6 := []int{}
        _ = v6
    default:
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint64 = 1
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(fi7)
                if true {
                    switch "" {
                    case "\n\t":
                        fmt.Println(fi9, fi9)
                    case "":
                        fmt.Println(fi9)
                    case "hello":
                        v8 += (v8 - (v8 << ((65535) & 63)))
                    default:
                        var v10 int = int(v8)
                        _ = v10
                    }
                    var v11 uint16 = 2
                    fmt.Println(v8, v8)
                    if false {
                        v11 = 1
                        v12 := GenS1{F0: -2147483648}
                        var v13 int64 = -309
                        _ = v12
                        _ = v13
                    }
                    v8 += v8
                    _ = v11
                }
                v14 := GenS1{F0: fi7}
                if (!true) {
                    v8 -= (v8 * v8)
                }
                _ = fi9
                _ = v14
            }
            _ = fi7
            _ = v8
        }
    }
    if (!true) {
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 uint32 = 4294967295
            v16 += (v16 | (v16 * v16))
            _ = fi15
            _ = v16
        }
        v17 := []int{-473, -190, 0, 9223372036854775807}
        _ = v17
    }
    var v18 float64 = float64((3.14159 * 0.0))
    v19 := GenS2{F0: ("a" + "\n\t")}
    _ = v18
    _ = v19
    return -533
}

func main() {
    var v1 int64 = 42
    v2 := []int{42, -9223372036854775808}
    var v3 uint16 = 2
    v4 := make(map[string]bool)
    var v5 uint8 = 195
    fmt.Println(v3, v1, v5)
    fmt.Println(v5, v5)
    var v20 int32 = 42
    v5 -= (v5 << ((255) & 63))
    for fi21 := 0; fi21 < 11; fi21++ {
        v20 = (v20 - (-(v20)))
        var v22 int32 = -538
        _ = fi21
        _ = v22
    }
    if (int8(v5) != 127) {
        v23 := GenS2{F0: ("hello" + "")}
        fmt.Println(v1, v20, v3)
        _ = v23
    }
    _ = len(v2)
    v24 := GenS3{F0: (^(v1)), F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
