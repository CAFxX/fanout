package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int) int8 {
    if false {
        fmt.Println(p0)
    }
    p0 += p0
    p0 -= p0
    p0 += 127
    if (!false) {
        if true {
            p0 = p0
            var v6 int32 = 16
            _ = v6
        } else {
            switch (p0 / ((p0) | 1)) {
            case -271:
                if (0.1 > ((0.0 - 1.5) + 1.5)) {
                    v7 := []int{}
                    var v8 int64 = 9223372036854775807
                    _ = v7
                    _ = v8
                }
            }
            var v9 int8 = -1
            _ = v9
        }
    }
    _ = p0
    return 1
}
func genFn1() uint16 {
    v10 := GenS1{F0: 1000}
    v11 := []int{42}
    for fi12 := 0; fi12 < 8; fi12++ {
        v13 := GenS1{F0: -32768}
        _ = fi12
        _ = v13
    }
    var v14 uint64 = 2
    _ = v10
    _ = v11
    _ = v14
    return 255
}

func main() {
    var v1 float32 = 1.5
    v2 := []string{}
    var v3 int64 = -1
    v4 := []string{"\n\t", "foo bar"}
    v5 := make(map[int]string)
    if (!true) {
        v15 := GenS2{F0: 255, F1: 4294967295}
        v16 := v5[9223372036854775807]
        _ = len(v2)
        _ = v15
        _ = v16
    } else {
        var v17 int = 111
        _ = v17
    }
    switch (("hello" + "a") + "\n\t") {
    case "\n\t":
        var v18 bool = (float32((float64((1e10 - v1)) / 1e10)) < 0.1)
        _ = v18
    }
    fmt.Println(v1, v3)
    fmt.Println(v3)
    if true {
        if true {
            v1 = v1
            var v19 int64 = 0
            var v20 uint = 87
            _ = v19
            _ = v20
        }
        v21 := v5[995]
        _ = v21
    }
    if (42 == 128) {
        _ = len(v5)
    } else {
        v22 := v5[-2147483648]
        if (1 < 255) {
            fmt.Println(v3)
            for fi23 := 0; fi23 < 11; fi23++ {
                v22 = ((v22 + (v22 + "hello")) + ("" + ""))
                _ = fi23
            }
            fmt.Println(v1, v22)
        } else {
            if ("\n\t" <= "\n\t") {
                v2 = append(v2, v22)
                if (true || true) {
                    v22 = v22
                }
            } else {
                _ = len(v2)
            }
            v24 := GenS2{F0: 2, F1: 4294967295}
            _ = v24
        }
        _ = v22
    }
    switch 128 {
    case 0:
        for fi25 := 0; fi25 < 8; fi25++ {
            v1 += v1
            _ = fi25
        }
    default:
        fmt.Println(v1)
    }
    var v26 int8 = 42
    fmt.Println(v26, v1, v1)
    switch 9223372036854775807 {
    case 4294967295:
        fmt.Println(v26, v1)
    case 9223372036854775807:
        v3 += (v3 & -698)
    case -9223372036854775808:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v2 = append(v2, "a")
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
