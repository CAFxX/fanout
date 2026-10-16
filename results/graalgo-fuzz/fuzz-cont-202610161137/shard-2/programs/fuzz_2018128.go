package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int32; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint, p2 bool) int8 {
    p0 -= p1
    v5 := []int{42, 1}
    for fi6 := 0; fi6 < 14; fi6++ {
        p1 += p0
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return int8(p0)
}

func main() {
    v1 := []int{-960, 4294967295, 42}
    v2 := []string{}
    var v3 uint64 = 505
    var v4 string = ("foo bar" + ("foo bar" + ("hello" + "\n\t")))
    for fi7 := 0; fi7 < 17; fi7++ {
        if len(v2) > 0 {
            v2[0] = v4
        }
        _ = fi7
    }
    _ = len(v1)
    fmt.Println(v4)
    v3 += 797
    var v8 int = -293
    if false {
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                if false {
                    v4 = v4
                    v8 -= fi9
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v11 int = 9223372036854775807
                    _ = v11
                }
                var v12 int32 = -2147483648
                if false {
                    if true {
                        if false {
                            fmt.Println(v8, v3)
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = v8
                    }
                    v8 = fi10
                }
                _ = fi10
                _ = v12
            }
            var v13 uint64 = 1
            _ = fi9
            _ = v13
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            var v14 string = v4
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 5; fi15++ {
                for fi16 := 0; fi16 < 15; fi16++ {
                    if true {
                        var v17 bool = (((1.5 - 3.14159) * 0.1) == 1.5)
                        var v18 uint32 = 73
                        var v19 uint16 = 32767
                        var v20 int = (v8 - 128)
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint64 = (^(v3))
                        v22 := GenS1{F0: 1, F1: 42}
                        _ = v21
                        _ = v22
                    }
                    var v23 int = v8
                    _ = fi16
                    _ = v23
                }
                if (1 == (fi15 ^ fi15)) {
                    v8 -= fi15
                    var v24 int8 = -101
                    v25 := GenS2{F0: 1, F1: 1, F2: -128}
                    _ = v24
                    _ = v25
                }
                v3 = 0
                _ = fi15
            }
        }
        var v26 bool = (float32(float64(1.5)) == 1e10)
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 string = "foo bar"
            _ = fi27
            _ = v28
        }
        _ = v26
    }
    switch ("foo bar" + v4) {
    case "hello":
        v29 := GenS3{F0: 1}
        _ = v29
    case "foo bar":
        var v30 int16 = 849
        _ = v30
    default:
        var v31 int = v8
        _ = v31
    }
    var v32 int64 = 847
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v8
}
