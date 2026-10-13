package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int, p1 uint32, p2 string) int16 {
    v5 := GenS1{F0: 0, F1: "", F2: 41}
    switch p2 {
    case "foo bar":
        v6 := GenS2{F0: (42 != 0)}
        _ = v6
    case "":
        var v7 uint64 = 559
        _ = v7
    default:
        v8 := GenS3{F0: 2}
        _ = v8
    }
    v9 := []int{(^(p0))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return -743
}

func main() {
    var v1 int16 = 711
    v2 := []int{-1, 127, -801, -9223372036854775808}
    v3 := []int{-834}
    v4 := make(map[string]bool)
    fmt.Println(v1)
    v10 := GenS3{F0: 163}
    if len(v2) > 0 {
        v2[0] = 256
    }
    var v11 uint8 = 34
    switch "hello" {
    case "foo bar":
        var v12 int64 = -447
        _ = v12
    case "hello":
        v11 = (v11 ^ v11)
    case "":
        v1 -= v1
    default:
        if (!true) {
            v13 := GenS3{F0: 2}
            v14 := v4[(("\n\t" + "hello") + ("a" + ("foo bar" + "foo bar")))]
            _ = v13
            _ = v14
        }
    }
    if false {
        switch "foo bar" {
        case "\n\t":
            if ((!(true && (true && false))) && true) {
                v11 += v11
                v1 -= (v1 % ((v1) | 1))
            } else {
                for fi15 := 0; fi15 < 4; fi15++ {
                    if (fi15 > 32767) {
                        fmt.Println(v1, fi15, v11)
                        v11 += (v11 - (v11 << ((605) & 63)))
                    } else {
                        fmt.Println(fi15, v11, v11)
                        if ((1 < 65535) || (18446744073709551615 > 65)) {
                            _ = len(v4)
                        }
                        var v16 int64 = 9223372036854775807
                        _ = v16
                    }
                    _ = fi15
                }
                var v17 uint32 = uint32(v11)
                _ = v17
            }
        case "":
            fmt.Println(v11, v11)
        default:
            fmt.Println(v1, v1)
        }
        if (18446744073709551615 == 42) {
            _ = len(v4)
            v18 := GenS4{F0: 0, F1: v11}
            var v19 uint8 = v11
            _ = v18
            _ = v19
        } else {
            v11 -= 255
            fmt.Println(v11, v11)
            v11 = (v11 & v11)
        }
    }
    var v20 int8 = -1
    v20 = -25
    switch "a" {
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        var v21 bool = true
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
