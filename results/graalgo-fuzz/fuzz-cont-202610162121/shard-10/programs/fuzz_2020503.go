package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 string }
func genFn0() bool {
    v5 := []int{225}
    var v6 int = 32767
    v7 := GenS1{F0: -348, F1: 65535}
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 821, F1: 0}
        if len(v5) > 0 {
            v5[0] = (v6 ^ -2147483648)
        }
        v10 := GenS2{F0: 454}
        _ = fi8
        _ = v10
        _ = v9
    }
    switch (v6 & (v6 ^ v6)) {
    case -131:
        v6 = 2147483647
    default:
        switch (-(v6)) {
        case 1:
            var v11 int16 = 42
            _ = v11
        case 128:
            v6 += 425
        default:
            v5 = append(v5, (v6 % ((508) | 1)))
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return (true && ("" > "hello"))
}
func genFn1(p0 int64, p1 uint8, p2 uint64) int {
    p2 += 510
    var v12 int = -2147483648
    var v13 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return v12
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 204
    v3 := []string{"\n\t"}
    var v4 uint64 = (v2 * (v2 / (((v2 ^ 1)) | 1)))
    var v14 int32 = 42
    _ = len(v1)
    var v15 int32 = (v14 - (-(v14)))
    var v16 int = 454
    if true {
        v17 := GenS3{F0: 42, F1: "hello"}
        v2 = 4294967295
        _ = v17
    }
    var v18 uint = 255
    var v19 uint32 = 2147483648
    var v20 float32 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
