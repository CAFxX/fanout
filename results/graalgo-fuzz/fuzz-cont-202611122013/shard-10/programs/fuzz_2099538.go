package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }

func main() {
    var v1 uint16 = 607
    var v2 float64 = (0.1 / 1.0)
    var v3 uint32 = 494
    var v4 int8 = 94
    v5 := make(map[string]uint32)
    var v6 int32 = 2147483647
    v5["foo bar"] = (v3 - 61)
    v1 -= 938
    v1 = (v1 % ((v1) | 1))
    var v7 uint = 0
    for fi8 := 0; fi8 < 16; fi8++ {
        switch (("hello" + "hello") + "hello") {
        case "\n\t":
            switch "\n\t" {
            case "\n\t":
                v9 := GenS1{F0: -1}
                _ = v9
            case "":
                if (0.0 <= 0.0) {
                    v10 := GenS1{F0: 42}
                    var v11 uint16 = 0
                    _ = v10
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 7; fi12++ {
                        if (false || true) {
                            var v13 uint32 = v3
                            v14 := GenS1{F0: 0}
                            if false {
                                v15 := []int{(fi12 | fi8), (-(fi12)), fi8}
                                _ = v15
                            } else {
                                v1 = (v1 << (((v7 | 0)) & 63))
                            }
                            _ = v13
                            _ = v14
                        }
                        var v16 uint16 = (v1 % (((^(v1))) | 1))
                        fmt.Println(v16, v6)
                        if true {
                            var v17 uint32 = (v3 ^ 655)
                            fmt.Println(fi8, fi8, v3)
                            var v18 int8 = v4
                            _ = v17
                            _ = v18
                        }
                        v19 := GenS1{F0: -1}
                        _ = fi12
                        _ = v16
                        _ = v19
                    }
                    var v20 float64 = v2
                    _ = v20
                }
            case "a":
                var v21 uint8 = 1
                _ = v21
            default:
                if false {
                    switch "hello" {
                    case "\n\t":
                        var v22 uint8 = 42
                        _ = v22
                    case "":
                        fmt.Println(v7)
                    }
                    switch fi8 {
                    case 4294967295:
                        v5["foo bar"] = (v3 << ((uint(v3)) & 63))
                    case 128:
                        fmt.Println(v4, v4, v6)
                    case 425:
                        var v23 int16 = 1
                        _ = v23
                    }
                }
            }
        default:
            if false {
                v2 -= (v2 / float64(float32(1e10)))
                v3 += v3
            } else {
                v3 -= (v3 / ((v3) | 1))
            }
        }
        v24 := GenS1{F0: -230}
        for fi25 := 0; fi25 < 3; fi25++ {
            v7 -= (v7 % (((^(v7))) | 1))
            _ = fi25
        }
        _ = fi8
        _ = v24
    }
    v26 := GenS2{F0: uint32(v7), F1: 18446744073709551615}
    v27 := []int{}
    v28 := GenS1{F0: -32768}
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := GenS1{F0: int16(fi29)}
        if false {
            var v31 int = fi29
            _ = v31
        } else {
            v7 = (v7 | (v7 + v7))
        }
        v1 -= 42
        _ = fi29
        _ = v30
    }
    switch 128 {
    case 42:
        fmt.Println(v4, v1, v4)
    case -52:
        for fi32 := 0; fi32 < 11; fi32++ {
            fmt.Println(v1, v2, v3)
            _ = fi32
        }
    case 1:
        switch -1 {
        case 32768:
            _ = len(v27)
        case 4294967295:
            v33 := v5["hello"]
            _ = v33
        default:
            switch 966 {
            case -9223372036854775808:
                v7 += v7
            }
        }
    default:
        v4 -= v4
    }
    fmt.Println(v3, v3)
    switch "hello" {
    case "foo bar":
        for fi34 := 0; fi34 < 9; fi34++ {
            var v35 int8 = (v4 + (^(v4)))
            var v36 float32 = 0.0
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    v37 := GenS3{F0: v4, F1: v3, F2: ((v3 / (((v3 << ((65535) & 63))) | 1)) < v3)}
    v1 -= 65535
    for fi38 := 0; fi38 < 12; fi38++ {
        for fi39 := 0; fi39 < 5; fi39++ {
            v2 = v2
            v1 += v1
            var v40 int32 = v6
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    var v41 uint = (v7 & uint(v3))
    var v42 int16 = -32768
    switch ("a" + "a") {
    case "a":
        var v43 uint8 = 1
        _ = v43
    case "\n\t":
        v7 -= (v41 * v41)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
