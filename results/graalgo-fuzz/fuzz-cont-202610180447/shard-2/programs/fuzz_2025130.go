package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 string }
func genFn0() uint64 {
    var v5 int16 = 217
    v5 -= -32768
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 int32, p2 string) string {
    if ((!false) && false) {
        fmt.Println(p1, p0)
        p0 += (p0 | (p1 >> ((0) & 63)))
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint32 = 536
        v8 := GenS1{F0: 65535, F1: (p2 + ("foo bar" + "")), F2: v7}
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(p2, p1, p0)
    p1 += (p0 % ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    return p2
}

func main() {
    var v1 uint = 823
    var v2 uint = (v1 % ((v1) | 1))
    v3 := []int{-1, 42, 127, 128}
    v4 := make(map[uint32]uint32)
    v2 += (v2 * (v1 >> (((v2 | v2)) & 63)))
    if true {
        fmt.Println(v1, v1)
    } else {
        if true {
            if false {
                fmt.Println(v2)
            }
            var v9 int32 = 42
            v1 = 815
            _ = v9
        } else {
            switch 1 {
            case 537:
                switch "foo bar" {
                case "":
                    if true {
                        v10 := GenS2{F0: 1, F1: 2}
                        _ = v10
                    } else {
                        v2 -= v2
                    }
                case "a":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "hello":
                    switch 0 {
                    case -9223372036854775808:
                        fmt.Println(v2)
                    case 32768:
                        var v11 uint32 = 65535
                        _ = v11
                    default:
                        var v12 string = ("foo bar" + "hello")
                        _ = v12
                    }
                default:
                    v4[4294967295] = 2147483647
                }
            case -1:
                if len(v3) > 0 {
                    v3[0] = -304
                }
            }
            var v13 int16 = 1
            _ = v13
        }
    }
    v4[65535] = 42
    switch "\n\t" {
    case "a":
        v14 := GenS2{F0: 0, F1: 203}
        _ = v14
    case "\n\t":
        var v15 int64 = -475
        _ = v15
    case "hello":
        v1 = 565
    default:
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS2{F0: 1, F1: 1}
            var v18 int16 = 281
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    _ = len(v4)
    v19 := GenS1{F0: (^(v2)), F1: "foo bar", F2: 568}
    switch ("foo bar" + ("a" + "")) {
    case "a":
        v4[2147483648] = 4294967295
    case "\n\t":
        var v20 string = "hello"
        _ = v20
    case "foo bar":
        switch "" {
        case "\n\t":
            var v21 float32 = 1.5
            _ = v21
        case "hello":
            fmt.Println(v2, v1)
        }
    default:
        delete(v4, 2147483648)
    }
    if true {
        switch "" {
        case "\n\t":
            var v22 int16 = -839
            _ = v22
        }
        fmt.Println(v2, v1)
    } else {
        delete(v4, 301)
        fmt.Println(v1)
        if false {
            var v23 bool = (v1 > (v2 + v1))
            _ = v23
        } else {
            if (-2.25 <= 1e10) {
                var v24 uint8 = 112
                _ = v24
            }
            if (2 != 684) {
                fmt.Println(v2, v2)
                v25 := GenS3{F0: -19}
                v1 += 463
                _ = v25
            }
            var v26 uint = v2
            _ = v26
        }
    }
    if true {
        var v27 bool = false
        var v28 uint = (v2 - v1)
        _ = v27
        _ = v28
    } else {
        var v29 float32 = 1.5
        switch "\n\t" {
        case "\n\t":
            switch ("a" + "a") {
            case "":
                var v30 uint32 = 65535
                _ = v30
            }
        case "foo bar":
            fmt.Println(v2)
        case "a":
            fmt.Println(v2, v29, v29)
        }
        v31 := GenS3{F0: 32768}
        _ = v29
        _ = v31
    }
    var v32 int = -9223372036854775808
    v33 := GenS3{F0: int(v2)}
    for fi34 := 0; fi34 < 17; fi34++ {
        v1 = (v2 / ((v1) | 1))
        if true {
            fmt.Println(v32)
            v35 := GenS4{F0: 132, F1: "foo bar"}
            if (2 != 0) {
                var v36 uint8 = 127
                v1 += 255
                _ = v36
            }
            _ = v35
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
