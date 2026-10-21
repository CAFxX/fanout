package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint32; F1 uint32; F2 int }
type GenS5 struct { F0 uint64; F1 int }
func genFn0(p0 int32, p1 int16, p2 int64) int64 {
    switch 4294967295 {
    case -548:
        fmt.Println(p2, p2, p2)
    }
    p0 = (p0 * (-(p0)))
    switch 817 {
    case 9223372036854775807:
        var v5 int = 127
        _ = v5
    case 32767:
        switch -9223372036854775808 {
        case 128:
            if true {
                for fi6 := 0; fi6 < 16; fi6++ {
                    fmt.Println(p0, p1, p1)
                    var v7 int64 = (p2 | (-(p2)))
                    _ = fi6
                    _ = v7
                }
                if (42 <= 42) {
                    v8 := []int{-9223372036854775808}
                    p0 = p0
                    _ = v8
                }
            }
        case 9223372036854775807:
            v9 := GenS1{F0: 0}
            _ = v9
        case -9223372036854775808:
            if false {
                p2 = 257
                var v10 uint64 = 1
                _ = v10
            } else {
                var v11 int8 = -96
                var v12 uint16 = 255
                _ = v11
                _ = v12
            }
        default:
            v13 := GenS2{F0: -1}
            _ = v13
        }
    }
    v14 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (p2 * int64(p1))
}

func main() {
    var v1 int16 = -573
    var v2 float32 = 1e10
    v3 := []string{("foo bar" + "hello")}
    var v4 string = "a"
    var v15 string = (v4 + "hello")
    if len(v3) > 0 {
        v3[0] = "a"
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        v15 = "foo bar"
        var v17 uint16 = 65535
        v18 := GenS3{F0: (v17 * (^(v17))), F1: (v15 + v4)}
        _ = fi16
        _ = v17
        _ = v18
    }
    if (0 > -1) {
        v2 = v2
    }
    v19 := GenS2{F0: 0}
    v20 := GenS4{F0: 4294967295, F1: 42, F2: 2147483647}
    v4 = "\n\t"
    var v21 float32 = 0.0
    for fi22 := 0; fi22 < 10; fi22++ {
        if false {
            switch "foo bar" {
            case "hello":
                fmt.Println(v2, v15, v21)
            case "\n\t":
                var v23 uint32 = 539
                _ = v23
            case "":
                v15 = (v4 + (v15 + v15))
            }
        } else {
            v2 += float32(1e10)
        }
        _ = fi22
    }
    switch -977 {
    case 127:
        if true {
            fmt.Println(v4)
        } else {
            fmt.Println(v2, v4)
            if false {
                var v24 uint = 4294967295
                v2 = v2
                _ = v24
            }
            fmt.Println(v2)
        }
    }
    var v25 string = (v15 + ("" + v4))
    v25 = v4
    _ = len(v3)
    if ((v21 <= 3.14159) || false) {
        var v26 uint16 = 576
        v27 := GenS5{F0: 1, F1: 2147483647}
        fmt.Println(v26)
        _ = v26
        _ = v27
    } else {
        fmt.Println(v15, v21, v25)
    }
    v28 := GenS5{F0: 4294967295, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
