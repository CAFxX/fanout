package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 string; F1 int; F2 string }
func genFn0(p0 string) bool {
    var v5 int8 = 127
    v5 = (v5 & v5)
    for fi6 := 0; fi6 < 5; fi6++ {
        switch 2147483647 {
        case 2147483647:
            var v7 int16 = -1
            _ = v7
        case 0:
            v8 := GenS1{F0: 42}
            _ = v8
        case 32768:
            v9 := GenS2{F0: 1, F1: 82}
            _ = v9
        default:
            if true {
                v5 = v5
                fmt.Println(fi6, v5, v5)
            }
        }
        v10 := GenS2{F0: 42, F1: 678}
        if (true && (42 == 0)) {
            if false {
                p0 = p0
                switch ("a" + p0) {
                case "a":
                    v11 := []int{(^(fi6)), (fi6 - (fi6 % ((20) | 1))), (fi6 ^ (fi6 * (-(fi6))))}
                    _ = v11
                case "":
                    var v12 bool = true
                    _ = v12
                case "\n\t":
                    switch (fi6 ^ int(v5)) {
                    case -52:
                        v5 += (v5 / (((v5 - v5)) | 1))
                    case 820:
                        v5 = 43
                    default:
                        v5 = (v5 | (^(v5)))
                    }
                default:
                    v13 := GenS1{F0: 567}
                    _ = v13
                }
                if true {
                    p0 = p0
                    var v14 int32 = -1
                    _ = v14
                } else {
                    fmt.Println(p0)
                    if (!true) {
                        switch p0 {
                        case "\n\t":
                            var v15 int = fi6
                            _ = v15
                        }
                        v5 = (^(v5))
                    }
                    v16 := GenS3{F0: "hello", F1: "\n\t"}
                    _ = v16
                }
            }
        } else {
            var v17 uint8 = 127
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 float64 = (float64((1e10 * (1e10 / float32(3.14159)))) * (float64((-2.25 + 0.1)) * (-2.25 * ((((3.14159 * -2.25) + 3.14159) / 3.14159) / 3.14159))))
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        _ = fi6
        _ = v10
    }
    fmt.Println(p0, p0)
    var v20 uint32 = 520
    _ = p0
    _ = v20
    _ = v5
    return false
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint32 = v1
    var v3 uint32 = 176
    var v4 int32 = 0
    var v21 int64 = int64(v4)
    v22 := []int{-626, 128, 32767}
    for fi23 := 0; fi23 < 6; fi23++ {
        v3 = 357
        if (int16(fi23) <= -1) {
            fmt.Println(fi23, v4, v1)
            var v24 int64 = 0
            if false {
                var v25 bool = (362 < 0)
                var v26 int64 = (-(v21))
                v27 := GenS4{F0: "", F1: fi23, F2: ("hello" + "hello")}
                _ = v25
                _ = v26
                _ = v27
            }
            _ = v24
        }
        v28 := GenS4{F0: (((("hello" + "") + ("" + (("\n\t" + "\n\t") + ""))) + "a") + "\n\t"), F1: (fi23 / ((9223372036854775807) | 1)), F2: ("foo bar" + "")}
        _ = fi23
        _ = v28
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 uint8 = 111
        var v31 uint64 = 1
        _ = fi29
        _ = v30
        _ = v31
    }
    switch 2147483647 {
    case 2147483647:
        v22 = append(v22, 32767)
    case 0:
        v1 = v1
    case 32767:
        fmt.Println(v2, v1, v1)
    }
    fmt.Println(v21, v4)
    _ = len(v22)
    var v32 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v4
}
