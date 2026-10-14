package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int64; F1 uint8 }
func genFn0(p0 string, p1 int8, p2 int32) int64 {
    if (42 <= 65535) {
        p0 = p0
        fmt.Println(p1, p2)
    } else {
        p1 += p1
    }
    v6 := []int{713, -9223372036854775808, 1, -843}
    fmt.Println(p0, p1, p2)
    v7 := GenS1{F0: (p0 + p0)}
    _ = len(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -618
}
func genFn1() int16 {
    v8 := []int{-651, 495, 9223372036854775807, 810}
    var v9 int64 = -1
    v10 := GenS2{F0: 9223372036854775807, F1: 128, F2: 80}
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int64 = 42
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 int16 = 32767
    var v2 uint = 18446744073709551615
    v3 := []int{1}
    var v4 int64 = -9223372036854775808
    var v5 int32 = 1
    var v13 int8 = int8(v5)
    var v14 int = 4294967295
    switch "\n\t" {
    case "":
        v15 := GenS2{F0: v14, F1: 127, F2: 654}
        _ = v15
    case "foo bar":
        v16 := GenS1{F0: "\n\t"}
        _ = v16
    case "\n\t":
        var v17 bool = false
        _ = v17
    default:
        v14 -= (-(v14))
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS3{F0: "", F1: v1}
        if (!true) {
            var v20 uint8 = 111
            var v21 uint = (v2 | 1)
            v22 := GenS4{F0: 42, F1: 42}
            _ = v20
            _ = v21
            _ = v22
        }
        var v23 int64 = genFn0("", (v13 + (v13 << ((4294967295) & 63))), v5)
        _ = fi18
        _ = v19
        _ = v23
    }
    var v24 uint64 = 2
    var v25 string = ("\n\t" + "hello")
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v4, v2, v2)
    case "hello":
        switch "" {
        case "\n\t":
            v26 := GenS4{F0: v4, F1: 79}
            _ = v26
        case "hello":
            v27 := GenS4{F0: (v4 - genFn0(v25, (v13 / ((v13) | 1)), -1)), F1: 0}
            _ = v27
        default:
            var v28 int64 = v4
            _ = v28
        }
    case "\n\t":
        var v29 int16 = 42
        _ = v29
    }
    switch (("hello" + "hello") + v25) {
    case "\n\t":
        fmt.Println(v1)
    case "":
        v13 = v13
    case "hello":
        v3 = append(v3, v14)
    }
    _ = len(v3)
    var v30 int32 = (v5 << (((v2 << (((v2 ^ (v2 % (((v2 * v2)) | 1)))) & 63))) & 63))
    _ = len(v3)
    for fi31 := 0; fi31 < 12; fi31++ {
        for fi32 := 0; fi32 < 2; fi32++ {
            v3 = append(v3, (fi32 | 461))
            _ = fi32
        }
        if len(v3) > 0 {
            v3[0] = (^(v14))
        }
        _ = fi31
    }
    if true {
        for fi33 := 0; fi33 < 12; fi33++ {
            if true {
                v13 = v13
                v34 := GenS3{F0: v25, F1: v1}
                _ = v34
            }
            _ = fi33
        }
        var v35 int16 = 924
        _ = v35
    }
    v3 = append(v3, v14)
    var v36 string = "\n\t"
    if false {
        if (0 >= v14) {
            var v37 int8 = v13
            _ = v37
        } else {
            for fi38 := 0; fi38 < 7; fi38++ {
                if true {
                    v1 = v1
                    var v39 uint8 = 68
                    _ = v39
                }
                _ = fi38
            }
        }
        var v40 uint32 = 4294967295
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
}
