package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 string }
type GenS2 struct { F0 int16; F1 int }
func genFn0(p0 uint8) bool {
    var v5 float32 = float32(3.14159)
    if true {
        p0 -= p0
    }
    fmt.Println(p0, v5)
    _ = p0
    _ = v5
    return false
}
func genFn1(p0 bool, p1 uint, p2 int64) uint64 {
    fmt.Println(p1, p2)
    v6 := []int{32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    var v1 int = 25
    v2 := []uint32{888, 0}
    v3 := make(map[int]string)
    v4 := []string{""}
    switch v1 {
    case 127:
        v1 -= (v1 & v1)
    case 1:
        _ = len(v4)
    }
    if (!(763 > 2)) {
        v1 = 4294967295
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(fi7, v1)
            _ = fi7
        }
        v1 = v1
    }
    switch (("\n\t" + "") + ("foo bar" + ("" + "\n\t"))) {
    case "a":
        var v8 int64 = 1
        _ = v8
    case "":
        fmt.Println(v1)
    case "\n\t":
        var v9 int64 = 956
        _ = v9
    default:
        for fi10 := 0; fi10 < 1; fi10++ {
            v1 += v1
            v11 := GenS1{F0: 65, F1: (("foo bar" <= ("foo bar" + ("\n\t" + "\n\t"))) && false), F2: ""}
            _ = fi10
            _ = v11
        }
    }
    if (2 < 0) {
        for fi12 := 0; fi12 < 17; fi12++ {
            v1 -= (-(fi12))
            var v13 float32 = 0.1
            _ = fi12
            _ = v13
        }
    } else {
        fmt.Println(v1)
    }
    if false {
        var v14 float64 = (-2.25 * 0.1)
        _ = v14
    }
    var v15 uint16 = 293
    if false {
        switch (v1 & (^(v1))) {
        case 339:
            _ = len(v4)
        case -589:
            var v16 int16 = 0
            _ = v16
        }
    } else {
        for fi17 := 0; fi17 < 20; fi17++ {
            delete(v3, 32768)
            var v18 string = "foo bar"
            fmt.Println(v15, v18, fi17)
            _ = fi17
            _ = v18
        }
        _ = len(v3)
        fmt.Println(v15)
    }
    v19 := GenS2{F0: 32767, F1: -528}
    var v20 int16 = -1
    if false {
        v20 = v20
        if (v20 <= (v20 << ((4294967295) & 63))) {
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 bool = false
                if false {
                    switch "\n\t" {
                    case "foo bar":
                        _ = len(v3)
                    }
                    for fi23 := 0; fi23 < 19; fi23++ {
                        switch (v1 ^ (v1 / ((-269) | 1))) {
                        case 2147483647:
                            var v24 uint64 = 42
                            _ = v24
                        case 1:
                            v15 += (v15 | v15)
                        default:
                            fmt.Println(v20)
                        }
                        v20 -= (v20 + v20)
                        _ = fi23
                    }
                } else {
                    for fi25 := 0; fi25 < 16; fi25++ {
                        var v26 int = 9223372036854775807
                        v1 = (v26 & fi21)
                        v22 = v22
                        _ = fi25
                        _ = v26
                    }
                }
                var v27 float32 = 3.14159
                _ = fi21
                _ = v22
                _ = v27
            }
        } else {
            var v28 int64 = int64(v15)
            _ = v28
        }
    }
    var v29 uint16 = (v15 - (v15 ^ 636))
    v4 = append(v4, "\n\t")
    v2 = append(v2, 1)
    var v30 uint16 = (v29 / ((507) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
