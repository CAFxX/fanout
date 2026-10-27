package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0(p0 uint64, p1 uint, p2 int64) uint64 {
    if (true && (p2 >= 0)) {
        if (!true) {
            fmt.Println(p2, p0)
            var v7 uint = p1
            if false {
                p2 += (p2 & (p2 ^ p2))
            }
            _ = v7
        }
    }
    v8 := GenS1{F0: (p2 / (((p2 >> (((p1 >> (((^(p1))) & 63))) & 63))) | 1)), F1: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}
func genFn1(p0 int64, p1 uint16, p2 bool) uint8 {
    p0 += 429
    var v9 int8 = 100
    v10 := GenS1{F0: -9223372036854775808, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 uint = 0
    v2 := []string{""}
    var v3 int32 = -2147483648
    v4 := []int{32768}
    var v5 int64 = 1
    v6 := make(map[string]int)
    v5 += (v5 & v5)
    v11 := GenS2{F0: 1, F1: false}
    var v12 int64 = v5
    v13 := GenS2{F0: 255, F1: false}
    var v14 uint8 = 238
    v12 = (v12 * (v12 ^ v5))
    v12 = (v12 ^ v5)
    switch "hello" {
    case "":
        switch -2147483648 {
        case 32767:
            var v15 string = "\n\t"
            _ = v15
        case -2147483648:
            var v16 bool = true
            _ = v16
        case -1:
            v12 = (-(v5))
        default:
            v17 := GenS1{F0: v5, F1: 652}
            _ = v17
        }
    case "\n\t":
        v6[("a" + "a")] = -1
    case "a":
        for fi18 := 0; fi18 < 17; fi18++ {
            if ("" > ("" + "a")) {
                fmt.Println(fi18)
            } else {
                if false {
                    fmt.Println(fi18, v5, v14)
                    var v19 float32 = (((3.14159 - 1e10) * (3.14159 + ((float32(-2.25) + 1.5) * -2.25))) + 1e10)
                    if false {
                        var v20 float32 = v19
                        fmt.Println(v12)
                        _ = v20
                    }
                    _ = v19
                } else {
                    var v21 float32 = 1e10
                    if false {
                        v22 := GenS1{F0: (v12 + v5), F1: 236}
                        v1 = v1
                        v23 := GenS3{F0: v14, F1: 0, F2: -35}
                        _ = v22
                        _ = v23
                    }
                    _ = v21
                }
                fmt.Println(v3, v12, fi18)
                var v24 uint64 = 0
                _ = v24
            }
            _ = fi18
        }
    default:
        switch "a" {
        case "":
            fmt.Println(v3, v1, v5)
        case "foo bar":
            var v25 uint16 = 922
            _ = v25
        default:
            fmt.Println(v12, v14, v1)
        }
    }
    v26 := GenS1{F0: 1, F1: 1}
    fmt.Println(v3, v5, v14)
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
