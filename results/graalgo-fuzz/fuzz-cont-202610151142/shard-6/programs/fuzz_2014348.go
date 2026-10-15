package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int8; F1 int64; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint64; F2 int }
type GenS5 struct { F0 bool; F1 string; F2 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 2
    v3 := make(map[int]string)
    var v4 int64 = 0
    var v5 uint = uint(v2)
    v6 := make(map[int]bool)
    var v7 uint32 = 42
    var v8 uint16 = 255
    v9 := GenS1{F0: -1, F1: -1}
    v7 = uint32(v8)
    for fi10 := 0; fi10 < 1; fi10++ {
        delete(v3, fi10)
        _ = fi10
    }
    delete(v1, "foo bar")
    v11 := GenS1{F0: 127, F1: 328}
    switch ("" + (("hello" + "\n\t") + "hello")) {
    case "a":
        v12 := []int{}
        _ = v12
    case "\n\t":
        v2 = v2
    case "foo bar":
        v2 = (^(v2))
    }
    v13 := GenS2{F0: -128, F1: v4, F2: (!(32767 == -9223372036854775808))}
    var v14 float64 = (0.0 / (3.14159 + float64(1.5)))
    fmt.Println(v5, v2, v14)
    for fi15 := 0; fi15 < 5; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 int8 = 100
            v14 = (v14 * v14)
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    if false {
        switch "\n\t" {
        case "hello":
            if ((!false) || (!(v4 > (v4 - 1)))) {
                if (true && false) {
                    var v18 uint64 = 19
                    var v19 int32 = 2147483647
                    _ = v18
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        for fi21 := 0; fi21 < 14; fi21++ {
                            if true {
                                v8 = (^(v8))
                            } else {
                                delete(v1, "a")
                                v22 := GenS3{F0: int64(v8)}
                                var v23 uint8 = 1
                                _ = v22
                                _ = v23
                            }
                            _ = fi21
                        }
                        for fi24 := 0; fi24 < 9; fi24++ {
                            v14 -= v14
                            var v25 uint32 = 2147483648
                            fmt.Println(fi24, v2, fi20)
                            _ = fi24
                            _ = v25
                        }
                        _ = fi20
                    }
                    var v26 int8 = 42
                    var v27 uint = (v5 | (^(v5)))
                    _ = v26
                    _ = v27
                }
                var v28 float32 = 1e10
                _ = v28
            } else {
                v29 := GenS4{F0: 786, F1: 184, F2: 9223372036854775807}
                var v30 uint16 = (v8 << (((v5 - (v5 + (v5 - (v5 | (v5 % ((uint(v8)) | 1))))))) & 63))
                _ = v29
                _ = v30
            }
        default:
            v8 += 1
        }
        v7 = uint32(v5)
    } else {
        fmt.Println(v7)
    }
    delete(v6, 0)
    switch -669 {
    case 6:
        if false {
            v8 = 255
        } else {
            v31 := GenS5{F0: (!true), F1: "foo bar", F2: -527}
            var v32 uint64 = 539
            _ = v31
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
