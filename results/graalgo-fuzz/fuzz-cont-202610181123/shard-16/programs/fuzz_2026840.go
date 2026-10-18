package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint16 }

func main() {
    var v1 int32 = -861
    var v2 uint = 2
    v3 := make(map[int]int)
    var v4 int8 = 42
    var v5 uint32 = 2
    var v6 bool = (((("foo bar" + "a") + "") + "") < (("a" + ("foo bar" + "a")) + "hello"))
    v7 := GenS1{F0: 1, F1: "", F2: 13}
    v1 = v1
    var v8 int = 32768
    v9 := GenS1{F0: (v4 ^ v4), F1: "a", F2: v4}
    v2 = (v2 / (((v2 & 956)) | 1))
    if (v6 || v6) {
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 18; fi11++ {
                v6 = v6
                _ = fi11
            }
            if (false && v6) {
                fmt.Println(fi10, v5, v1)
                v3[v8] = -68
                if true {
                    v5 = 203
                }
            } else {
                fmt.Println(v6, fi10)
            }
            var v12 float32 = -2.25
            _ = fi10
            _ = v12
        }
        v4 = 42
    } else {
        var v13 int64 = -520
        var v14 string = "foo bar"
        if (v13 < (v13 * (v13 + 42))) {
            v5 = 2147483647
        } else {
            var v15 int = (v8 / (((v8 << ((358) & 63))) | 1))
            _ = v15
        }
        _ = v13
        _ = v14
    }
    v16 := GenS2{F0: (v1 >> (((^(v2))) & 63)), F1: 255}
    var v17 int8 = (v4 | (v4 - v4))
    var v18 float64 = 0.0
    var v19 int64 = -1
    var v20 float32 = 3.14159
    for fi21 := 0; fi21 < 8; fi21++ {
        v4 += (v4 / ((127) | 1))
        _ = fi21
    }
    v22 := []int{32768}
    switch ("" + "\n\t") {
    case "foo bar":
        v1 += -1
    case "\n\t":
        v4 -= (v17 % (((v17 / ((100) | 1))) | 1))
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 int64 = v19
            v25 := GenS3{F0: 18446744073709551615, F1: 473, F2: 2}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
