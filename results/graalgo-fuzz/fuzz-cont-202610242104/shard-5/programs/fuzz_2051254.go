package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 string; F2 int8 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint16 }
func genFn0() uint64 {
    var v7 float32 = (((3.14159 - ((0.0 / 3.14159) - 1e10)) - (0.1 - 0.1)) - ((1e10 / 1.0) + (1e10 - 0.1)))
    v7 = ((v7 * float32(-2.25)) + v7)
    _ = v7
    return 1
}

func main() {
    var v1 int16 = 852
    v2 := []int{-184, 9223372036854775807, 466}
    var v3 uint16 = 858
    var v4 int64 = -9223372036854775808
    v5 := []int{42}
    var v6 uint8 = 2
    switch -9223372036854775808 {
    case -9223372036854775808:
        switch "a" {
        case "foo bar":
            fmt.Println(v1, v3, v6)
        case "":
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v4, v6, v4)
            }
        case "hello":
            var v8 int16 = -954
            _ = v8
        default:
            v9 := GenS1{F0: (-(v4)), F1: 0, F2: v6}
            _ = v9
        }
    }
    v1 += 32767
    fmt.Println(v1, v6)
    fmt.Println(v4, v1, v4)
    var v10 int8 = -95
    var v11 bool = ((v6 | (v6 & (v6 ^ v6))) == v6)
    if ("" == "\n\t") {
        v12 := GenS2{F0: v6, F1: 109}
        _ = v12
    }
    v1 -= (v1 % ((v1) | 1))
    switch ((("\n\t" + (("hello" + "") + "foo bar")) + "\n\t") + "hello") {
    case "":
        v11 = (!v11)
    default:
        for fi13 := 0; fi13 < 4; fi13++ {
            switch "\n\t" {
            case "hello":
                v14 := GenS3{F0: ("" + "hello")}
                _ = v14
            case "":
                v3 -= v3
            case "a":
                v15 := GenS3{F0: ""}
                _ = v15
            }
            if (114 > 229) {
                v16 := GenS4{F0: (fi13 | 0), F1: "\n\t", F2: v10}
                _ = v16
            } else {
                var v17 float32 = (-2.25 + 0.1)
                var v18 float64 = 1e10
                _ = v17
                _ = v18
            }
            _ = fi13
        }
    }
    var v19 uint = 0
    var v20 uint16 = v3
    fmt.Println(v11, v6)
    if false {
        v4 += 942
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v6 = 42
        fmt.Println(v10, v20)
        _ = fi21
    }
    v22 := GenS5{F0: v10, F1: -1}
    v23 := GenS6{F0: 0}
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    switch -9223372036854775808 {
    case 1:
        var v24 int64 = 9223372036854775807
        _ = v24
    case -9223372036854775808:
        if (!v11) {
            v25 := GenS5{F0: (v10 % (((v10 | 0)) | 1)), F1: -548}
            var v26 float32 = 0.1
            fmt.Println(v19, v20)
            _ = v25
            _ = v26
        } else {
            switch (("foo bar" + ("hello" + ("\n\t" + (((("hello" + ("a" + "")) + "\n\t") + "foo bar") + "a")))) + "") {
            case "foo bar":
                var v27 bool = true
                _ = v27
            }
            switch "foo bar" {
            case "hello":
                var v28 uint64 = 403
                _ = v28
            case "foo bar":
                fmt.Println(v11, v6)
            case "\n\t":
                v1 += v1
            }
            var v29 bool = false
            _ = v29
        }
    case -305:
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
