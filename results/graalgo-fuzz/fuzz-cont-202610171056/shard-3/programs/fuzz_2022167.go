package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 int32, p1 uint8) int8 {
    v7 := []int{0, -249, 186}
    switch 58 {
    case 0:
        p1 = p1
    case 128:
        var v8 int = 2147483647
        _ = v8
    default:
        p1 = 255
    }
    _ = p0
    _ = p1
    _ = v7
    return 127
}
func genFn1(p0 int) int {
    switch "foo bar" {
    case "\n\t":
        if (false && false) {
            v9 := []int{(^(p0)), (p0 * (p0 >> ((107) & 63)))}
            _ = v9
        }
    case "":
        p0 += p0
    }
    if true {
        fmt.Println(p0)
        var v10 uint32 = 330
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 string = "foo bar"
            var v13 int32 = 874
            var v14 int32 = 0
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = v10
    } else {
        switch (p0 << ((630) & 63)) {
        case 32768:
            v15 := []int{(p0 + (p0 + (^(p0))))}
            _ = v15
        }
    }
    _ = p0
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 991
    v3 := []string{}
    var v4 uint = 1
    var v5 float32 = -2.25
    var v6 bool = true
    var v16 int8 = -1
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    switch "foo bar" {
    case "a":
        var v17 uint8 = 128
        _ = v17
    case "hello":
        switch "a" {
        case "hello":
            v4 = v4
        case "":
            var v18 float32 = 0.1
            _ = v18
        default:
            var v19 int64 = -1
            _ = v19
        }
    }
    var v20 int32 = -261
    fmt.Println(v16)
    if (!false) {
        for fi21 := 0; fi21 < 14; fi21++ {
            var v22 uint16 = 42
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 uint64 = 42
                _ = fi23
                _ = v24
            }
            v16 = v16
            _ = fi21
            _ = v22
        }
        v20 += 0
        v5 += ((v5 * (v5 - v5)) / (v5 - (-2.25 / v5)))
    } else {
        v25 := GenS1{F0: 42}
        _ = v25
    }
    v5 = v5
    v26 := GenS1{F0: 9223372036854775807}
    fmt.Println(v2)
    fmt.Println(v6, v6, v5)
    var v27 float64 = 1.5
    switch 316 {
    case -432:
        v2 += v2
    }
    fmt.Println(v20)
    v20 = (v2 - v2)
    v5 = 3.14159
    var v28 uint64 = 18446744073709551615
    for fi29 := 0; fi29 < 16; fi29++ {
        v30 := GenS2{F0: (v4 / (((^(v4))) | 1)), F1: 0}
        _ = fi29
        _ = v30
    }
    var v31 int8 = int8(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
