package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int64; F1 string; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS5 struct { F0 int16 }

func main() {
    v1 := []uint32{1, 65535, 0, 654}
    v2 := []int{-779, 265}
    var v3 uint64 = 352
    v3 -= 659
    v4 := GenS1{F0: 32768, F1: 42}
    if ((!(1 < 133)) || false) {
        v3 -= v3
    }
    v5 := GenS2{F0: -1, F1: "hello", F2: 0}
    for fi6 := 0; fi6 < 14; fi6++ {
        _ = len(v1)
        v3 = 42
        _ = fi6
    }
    if len(v2) > 0 {
        v2[0] = 128
    }
    switch 127 {
    case 1:
        var v7 float32 = 0.1
        _ = v7
    }
    var v8 int32 = 0
    switch "hello" {
    case "":
        if (9223372036854775807 <= 1) {
            v9 := GenS2{F0: 1, F1: ((("" + (("\n\t" + "hello") + "\n\t")) + "hello") + "foo bar"), F2: 1}
            var v10 uint8 = 42
            for fi11 := 0; fi11 < 4; fi11++ {
                if (2 >= 930) {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        v10 += v10
                        v8 -= 2147483647
                        switch ("a" + "foo bar") {
                        case "a":
                            var v13 uint = uint(v3)
                            _ = v13
                        default:
                            switch (fi11 >> ((255) & 63)) {
                            case -9223372036854775808:
                                var v14 uint8 = (v10 * (v10 * 1))
                                _ = v14
                            default:
                                _ = len(v1)
                            }
                        }
                        v15 := GenS2{F0: -1, F1: "hello", F2: 65535}
                        v3 = (^(v3))
                        var v16 uint64 = 0
                        _ = fi12
                        _ = v15
                        _ = v16
                    }
                    if true {
                        if (!("hello" == "hello")) {
                            var v17 uint8 = 128
                            _ = v17
                        } else {
                            if true {
                                var v18 uint64 = 4294967295
                                var v19 int16 = 0
                                _ = v18
                                _ = v19
                            }
                        }
                        v20 := GenS3{F0: 40}
                        var v21 uint8 = v10
                        _ = v20
                        _ = v21
                    }
                    v8 -= 1
                } else {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        fmt.Println(fi11, v10)
                        v23 := GenS3{F0: 1}
                        _ = fi22
                        _ = v23
                    }
                    v10 += v10
                    fmt.Println(fi11, v8)
                }
                if (false && true) {
                    switch fi11 {
                    case 2147483647:
                        var v24 uint64 = v3
                        _ = v24
                    case 9223372036854775807:
                        var v25 bool = false
                        _ = v25
                    }
                } else {
                    v1 = append(v1, 65535)
                    fmt.Println(v3, v10)
                    for fi26 := 0; fi26 < 2; fi26++ {
                        v27 := GenS4{F0: v3, F1: 42, F2: (v3 * v3)}
                        v3 -= v3
                        var v28 uint64 = v3
                        fmt.Println(fi11, v28)
                        fmt.Println(v8)
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                }
                _ = fi11
            }
            _ = v10
            _ = v9
        }
    default:
        if true {
            v8 += (v8 % (((v8 & v8)) | 1))
            var v29 int32 = v8
            v29 += v29
            _ = v29
        } else {
            for fi30 := 0; fi30 < 13; fi30++ {
                v3 -= (v3 * v3)
                v31 := GenS5{F0: 32767}
                _ = fi30
                _ = v31
            }
            v2 = append(v2, -756)
            v3 = v3
        }
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 int = 1
        _ = fi32
        _ = v33
    }
    fmt.Println(v8, v3)
    _ = len(v2)
    v34 := GenS3{F0: 1}
    if len(v2) > 0 {
        v2[0] = 128
    }
    switch ("foo bar" + "a") {
    case "foo bar":
        for fi35 := 0; fi35 < 3; fi35++ {
            if false {
                v8 = v8
            } else {
                v3 = 1
                v2 = append(v2, (fi35 | (fi35 ^ fi35)))
            }
            _ = fi35
        }
    case "\n\t":
        var v36 int64 = -9223372036854775808
        _ = v36
    case "a":
        switch -560 {
        case 0:
            var v37 int = -632
            _ = v37
        case 32768:
            v3 = 357
        }
    default:
        fmt.Println(v8, v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v8
}
