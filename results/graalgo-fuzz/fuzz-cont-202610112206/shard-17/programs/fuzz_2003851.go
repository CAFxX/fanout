package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint64 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 bool; F1 bool; F2 int }
type GenS4 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0(p0 int8, p1 uint16) string {
    for fi5 := 0; fi5 < 13; fi5++ {
        p0 += p0
        var v6 bool = (127 != 128)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 124, F1: (((("a" + "") + "") + "") + "foo bar"), F2: 2}
    _ = p0
    _ = p1
    _ = v7
    return ""
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -1
    var v3 uint8 = 127
    var v4 uint16 = 255
    v2 += 111
    v3 += (v3 - v3)
    if true {
        v8 := GenS2{F0: 32767, F1: 60}
        _ = v8
    }
    v9 := []int{-587, 1, -9223372036854775808, -9223372036854775808}
    var v10 uint32 = 0
    switch "hello" {
    case "hello":
        for fi11 := 0; fi11 < 14; fi11++ {
            _ = len(v1)
            switch "\n\t" {
            case "a":
                fmt.Println(v10)
            case "foo bar":
                v3 = 255
            case "hello":
                v4 -= uint16(v10)
            default:
                if len(v9) > 0 {
                    v9[0] = (fi11 % (((fi11 + fi11)) | 1))
                }
            }
            var v12 bool = (937 != (v4 & (v4 + v4)))
            _ = fi11
            _ = v12
        }
    case "a":
        var v13 int64 = -9223372036854775808
        _ = v13
    case "":
        v10 = v10
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        if true {
            _ = len(v1)
        } else {
            v2 -= (-(v2))
        }
        if false {
            fmt.Println(v10)
            v1[(v10 / ((v10) | 1))] = ""
            for fi15 := 0; fi15 < 3; fi15++ {
                _ = len(v1)
                _ = fi15
            }
        }
        if false {
            switch "hello" {
            case "a":
                if false {
                    v16 := GenS1{F0: 2, F1: (("" + "hello") + "hello"), F2: 2}
                    for fi17 := 0; fi17 < 4; fi17++ {
                        if len(v9) > 0 {
                            v9[0] = (fi14 ^ -716)
                        }
                        v18 := v1[(v10 % (((v10 << ((65535) & 63))) | 1))]
                        var v19 int8 = v2
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = v16
                }
            case "foo bar":
                delete(v1, (v10 / (((v10 - v10)) | 1)))
            case "":
                v2 = 1
            }
            v10 = v10
        } else {
            if (false && false) {
                var v20 bool = true
                v21 := v1[42]
                var v22 uint32 = (v10 / ((v10) | 1))
                _ = v20
                _ = v21
                _ = v22
            } else {
                v3 -= (v3 | v3)
                v23 := GenS1{F0: 297, F1: "foo bar", F2: 18446744073709551615}
                _ = v23
            }
            v24 := GenS2{F0: fi14, F1: 0}
            _ = v24
        }
        _ = fi14
    }
    v25 := GenS2{F0: 128, F1: 114}
    var v26 uint64 = 1
    var v27 uint64 = v26
    v28 := GenS3{F0: false, F1: true, F2: 42}
    v29 := GenS4{F0: -242, F1: v3, F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
