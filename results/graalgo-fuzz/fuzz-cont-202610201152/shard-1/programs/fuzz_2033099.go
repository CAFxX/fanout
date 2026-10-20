package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 int }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[int]string)
    v3 := make(map[string]int)
    var v4 uint64 = 42
    var v5 uint16 = 513
    v6 := GenS1{F0: ("a" + "hello"), F1: 1000, F2: 477}
    var v7 int32 = -2147483648
    v8 := []int{9223372036854775807, -943}
    fmt.Println(v7)
    v1 += float64(1e10)
    switch ("a" + "hello") {
    case "\n\t":
        v1 -= v1
    default:
        for fi9 := 0; fi9 < 4; fi9++ {
            v1 = v1
            for fi10 := 0; fi10 < 13; fi10++ {
                v11 := GenS1{F0: "", F1: -535, F2: 892}
                var v12 int8 = 1
                v12 = v12
                switch 2147483647 {
                case 0:
                    switch ("\n\t" + ("foo bar" + "a")) {
                    case "":
                        switch 4294967295 {
                        case -2147483648:
                            v13 := GenS1{F0: "\n\t", F1: -32768, F2: int64(v5)}
                            _ = v13
                        case 1:
                            var v14 int64 = -864
                            _ = v14
                        case 4294967295:
                            v15 := GenS1{F0: "", F1: 100, F2: 1}
                            _ = v15
                        default:
                            v1 -= (1.5 * ((1e10 - v1) + (float64((((3.14159 / (0.0 + float32(v1))) / 3.14159) * (1.5 - 3.14159))) * ((((0.0 * v1) + (0.1 + v1)) / -2.25) - v1))))
                        }
                    case "\n\t":
                        var v16 int8 = (v12 % ((v12) | 1))
                        _ = v16
                    case "foo bar":
                        var v17 float64 = (((float64((0.0 - (float32(-2.25) + 3.14159))) * ((v1 - v1) * v1)) * 1e10) * 1e10)
                        _ = v17
                    }
                    v1 = v1
                    v18 := GenS1{F0: "", F1: 1000, F2: -326}
                    _ = v18
                }
                v4 -= v4
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    }
    if false {
        v19 := GenS2{F0: v4, F1: -32768}
        v5 += (v5 - v5)
        _ = v19
    }
    if false {
        var v20 string = "hello"
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS3{F0: fi21}
        var v23 int32 = (v7 ^ (^(v7)))
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = len(v2)
    v1 -= 0.1
    v24 := GenS1{F0: ("\n\t" + "a"), F1: 0, F2: 1}
    v1 = v1
    if false {
        _ = len(v8)
        var v25 uint16 = 538
        _ = v25
    } else {
        v5 = uint16(v4)
        switch -2147483648 {
        case 127:
            v26 := GenS3{F0: int(v5)}
            _ = v26
        }
    }
    var v27 float32 = ((float32(v1) * ((0.1 + 1e10) - ((((0.1 / (1e10 + 1e10)) - 1e10) + 3.14159) - 0.0))) - (0.1 - 1.5))
    if ("hello" <= ("" + "hello")) {
        for fi28 := 0; fi28 < 12; fi28++ {
            fmt.Println(fi28)
            _ = fi28
        }
    } else {
        switch -9223372036854775808 {
        case 9223372036854775807:
            delete(v2, 128)
        }
        v7 = 2147483647
        fmt.Println(v4, v4, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
