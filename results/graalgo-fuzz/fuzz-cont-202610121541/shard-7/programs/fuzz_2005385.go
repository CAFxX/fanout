package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int16, p1 int16, p2 int16) uint16 {
    p2 += p2
    fmt.Println(p1, p2)
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int8 = -128
            p2 = p1
            p1 = 1
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: "hello", F1: 42, F2: fi7}
        fmt.Println(p1)
        _ = fi7
        _ = v10
    }
    v11 := GenS1{F0: "\n\t", F1: 168, F2: 1}
    fmt.Println(p0, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 255
}
func genFn1(p0 uint16, p1 uint16) string {
    p0 += (p1 / (((p1 & (p0 * p0))) | 1))
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int64 = 402
        if (127 >= 42) {
            p1 -= p1
        } else {
            v14 := GenS2{F0: (!(genFn0(-685, -32768, -549) < 65535)), F1: 50}
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    if false {
        p0 -= 32767
    }
    var v15 string = "a"
    v16 := []int{32768}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return v15
}

func main() {
    var v1 int = -1
    v2 := make(map[string]uint32)
    var v3 uint32 = 65535
    v4 := make(map[int]bool)
    var v5 int = v1
    v6 := []uint32{}
    var v17 int = v1
    v18 := GenS3{F0: 32767}
    var v19 string = "hello"
    if (-32768 == 567) {
        if ((0 > 0) && (42 != 368)) {
            if len(v6) > 0 {
                v6[0] = v3
            }
            var v20 int64 = -1
            var v21 int8 = -61
            _ = v20
            _ = v21
        } else {
            if (0 == 329) {
                v22 := GenS3{F0: 255}
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS3{F0: genFn0(32767, -430, -861)}
            _ = fi23
            _ = v24
        }
    } else {
        var v25 int8 = 127
        if true {
            var v26 uint8 = 255
            _ = v26
        }
        _ = v25
    }
    fmt.Println(v3)
    fmt.Println(v19)
    var v27 uint = 438
    var v28 int16 = 1000
    var v29 int32 = 780
    if len(v6) > 0 {
        v6[0] = 0
    }
    switch "\n\t" {
    case "hello":
        v28 += v28
    case "\n\t":
        var v30 uint8 = 42
        _ = v30
    }
    var v31 uint32 = 2147483648
    v28 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
