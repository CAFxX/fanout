package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int8; F1 int16; F2 bool }
func genFn0(p0 uint64) uint {
    switch "hello" {
    case "":
        v6 := []int{127, -264, -9223372036854775808}
        _ = v6
    }
    switch "a" {
    case "\n\t":
        p0 += p0
    }
    var v7 float32 = 0.1
    _ = p0
    _ = v7
    return 2
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "hello"
    v3 := []uint32{42}
    v4 := []string{v2, v2, v2}
    v5 := make(map[string]uint32)
    var v8 int32 = 42
    var v9 string = "hello"
    _ = len(v3)
    v10 := GenS1{F0: v9, F1: 751, F2: 128}
    if (false || (genFn0(202) < 42)) {
        if ((v8 - (-(v8))) <= (v8 << ((4294967295) & 63))) {
            _ = len(v4)
            for fi11 := 0; fi11 < 12; fi11++ {
                v9 = "hello"
                _ = fi11
            }
            fmt.Println(v8, v8, v8)
        } else {
            fmt.Println(v9)
            v8 += (v8 / ((v8) | 1))
            switch "\n\t" {
            case "foo bar":
                var v12 string = v2
                _ = v12
            default:
                v13 := GenS2{F0: true, F1: -981}
                _ = v13
            }
        }
        v14 := GenS1{F0: "\n\t", F1: 32767, F2: 1}
        _ = v14
    } else {
        if false {
            v8 -= v8
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := GenS1{F0: v9, F1: -2147483648, F2: fi15}
                if (!true) {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v8 += 0
                        var v18 int8 = -119
                        if (!false) {
                            var v19 uint8 = 16
                            _ = len(v3)
                            switch v9 {
                            case "a":
                                v20 := GenS2{F0: (v8 == (v8 + (v8 & (^(v8))))), F1: 1}
                                _ = v20
                            case "\n\t":
                                var v21 int = (-(fi17))
                                var v22 uint32 = 2147483647
                                fmt.Println(fi17, v21, fi15)
                                _ = v21
                                _ = v22
                            }
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                }
                if ((0.1 + (0.0 + 3.14159)) <= 1.5) {
                    var v23 bool = false
                    var v24 uint = 255
                    fmt.Println(v2, v2)
                    _ = v23
                    _ = v24
                } else {
                    v25 := GenS2{F0: true, F1: -9223372036854775808}
                    v26 := GenS3{F0: 42, F1: 32768}
                    if false {
                        for fi27 := 0; fi27 < 13; fi27++ {
                            if len(v3) > 0 {
                                v3[0] = 4294967295
                            }
                            v8 -= v8
                            var v28 string = "hello"
                            _ = len(v3)
                            if ((656 == 767) && true) {
                                v8 += v8
                                fmt.Println(fi15, v2, v9)
                                v29 := GenS1{F0: "foo bar", F1: fi27, F2: fi27}
                                _ = v29
                            }
                            _ = fi27
                            _ = v28
                        }
                        fmt.Println(fi15, v2, fi15)
                        _ = len(v5)
                    }
                    _ = v25
                    _ = v26
                }
                _ = fi15
                _ = v16
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = 905
    }
    v9 = (v2 + ("foo bar" + "a"))
    v30 := GenS4{F0: 100, F1: int16(v8), F2: (false || true)}
    fmt.Println(v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
