package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int64) uint64 {
    p0 += 0
    v4 := GenS1{F0: 4294967295}
    _ = p0
    _ = v4
    return 37
}

func main() {
    v1 := []uint32{}
    v2 := []int{}
    var v3 int = -9223372036854775808
    switch 2147483647 {
    case 9223372036854775807:
        if (-1 == 0) {
            var v5 int16 = 638
            _ = len(v1)
            _ = v5
        } else {
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 float32 = (-2.25 / 1.0)
                var v8 float64 = -2.25
                _ = fi6
                _ = v7
                _ = v8
            }
            switch ("hello" + "a") {
            case "foo bar":
                v1 = append(v1, 2147483647)
            case "\n\t":
                for fi9 := 0; fi9 < 10; fi9++ {
                    for fi10 := 0; fi10 < 9; fi10++ {
                        v3 += 2147483647
                        switch fi9 {
                        case -1:
                            fmt.Println(fi9, fi10)
                        }
                        _ = fi10
                    }
                    v11 := GenS2{F0: 0, F1: "hello"}
                    if (true || false) {
                        switch 709 {
                        case -9223372036854775808:
                            var v12 uint = 1
                            _ = v12
                        default:
                            switch (v3 ^ (v3 / (((fi9 / ((fi9) | 1))) | 1))) {
                            case 32767:
                                v3 = (fi9 | (-(fi9)))
                            case -743:
                                v3 += (fi9 >> ((18446744073709551615) & 63))
                            case 90:
                                var v13 string = (("hello" + "hello") + ((("\n\t" + ("" + "hello")) + "a") + (("hello" + "") + (("foo bar" + "hello") + ""))))
                                _ = v13
                            }
                        }
                        var v14 int16 = 42
                        _ = v14
                    } else {
                        fmt.Println(v3)
                    }
                    _ = fi9
                    _ = v11
                }
            }
            v3 += v3
        }
    case 428:
        _ = len(v2)
    case 127:
        v3 = 947
    }
    switch v3 {
    case 32768:
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := GenS1{F0: 854}
            _ = fi15
            _ = v16
        }
    case 312:
        v17 := GenS3{F0: 716, F1: 1}
        _ = v17
    case -897:
        v3 += v3
    }
    if (65535 < 653) {
        var v18 int = 21
        var v19 uint = 1
        _ = v18
        _ = v19
    }
    var v20 uint32 = 484
    var v21 uint = 2
    if false {
        if len(v1) > 0 {
            v1[0] = (^(v20))
        }
        var v22 string = ""
        var v23 float32 = (0.1 * 3.14159)
        _ = v22
        _ = v23
    } else {
        var v24 float64 = float64(3.14159)
        v24 -= float64((-2.25 / 1.5))
        _ = v24
    }
    fmt.Println(v3, v3, v3)
    var v25 bool = true
    fmt.Println(v20)
    if len(v2) > 0 {
        v2[0] = (v3 << ((v21) & 63))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
}
