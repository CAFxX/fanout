package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 bool; F1 int; F2 int64 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 uint16 }
func genFn0() uint16 {
    v5 := []int{128, -2147483648, 9223372036854775807}
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(fi6)
        v7 := GenS1{F0: 18446744073709551615}
        var v8 uint8 = 128
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int = (-(fi10))
            var v12 int8 = -53
            if (!false) {
                if (2 == 42) {
                    fmt.Println(fi10)
                    v13 := GenS2{F0: 1, F1: 0}
                    if (32767 != 737) {
                        if true {
                            var v14 uint8 = 127
                            fmt.Println(v11)
                            fmt.Println(v12)
                            _ = v14
                        }
                        v15 := GenS3{F0: (false && false), F1: (fi10 + -93), F2: -1}
                        _ = v15
                    }
                    _ = v13
                }
                switch ("hello" + "\n\t") {
                case "a":
                    var v16 uint16 = 32767
                    _ = v16
                default:
                    fmt.Println(v11, v12, v12)
                }
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 uint = 0
            if (!(!(((2147483648 == 2147483648) || (-348 < -32768)) || (818 <= 255)))) {
                v20 := GenS1{F0: 65535}
                _ = len(v5)
                if len(v5) > 0 {
                    v5[0] = 85
                }
                switch (^(fi18)) {
                case -610:
                    v19 = 255
                case 4294967295:
                    var v21 int8 = 42
                    _ = v21
                }
                v19 -= v19
                var v22 bool = false
                switch (^(fi17)) {
                case 127:
                    v22 = ((3.14159 - (0.0 / 1.5)) < 1.5)
                }
                fmt.Println(v22)
                _ = v20
                _ = v22
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    _ = v5
    return 67
}
func genFn1(p0 int) int16 {
    if (!false) {
        var v23 int64 = 1
        var v24 bool = (true || false)
        _ = v23
        _ = v24
    } else {
        if (false && true) {
            if true {
                v25 := GenS4{F0: (p0 < (p0 ^ p0)), F1: 2}
                var v26 uint32 = 42
                p0 -= (p0 * p0)
                _ = v25
                _ = v26
            } else {
                fmt.Println(p0)
                var v27 uint64 = 0
                _ = v27
            }
            p0 += (p0 % (((p0 * -869)) | 1))
            v28 := []int{p0, p0, p0, p0}
            _ = v28
        }
        p0 = (p0 / (((p0 % (((p0 >> ((65535) & 63))) | 1))) | 1))
    }
    fmt.Println(p0)
    if ("foo bar" > "hello") {
        v29 := GenS2{F0: 1, F1: 2}
        if true {
            p0 = p0
            p0 += (-(p0))
            p0 += p0
        }
        p0 = (p0 << ((18446744073709551615) & 63))
        _ = v29
    }
    switch "" {
    case "foo bar":
        v30 := GenS5{F0: 0, F1: 42}
        _ = v30
    case "":
        var v31 int16 = 42
        _ = v31
    case "hello":
        var v32 uint8 = 127
        _ = v32
    default:
        var v33 int16 = 32767
        _ = v33
    }
    _ = p0
    return -693
}

func main() {
    var v1 bool = false
    var v2 uint16 = 659
    v3 := make(map[int]int)
    var v4 uint64 = 880
    v34 := GenS6{F0: 9223372036854775807, F1: 127}
    if false {
        v4 += (v4 / (((v4 + (v4 << ((255) & 63)))) | 1))
    }
    v3[9223372036854775807] = 127
    v4 = 10
    switch 4294967295 {
    case 42:
        if true {
            for fi35 := 0; fi35 < 7; fi35++ {
                var v36 int64 = -42
                _ = fi35
                _ = v36
            }
            var v37 int = -365
            v4 = 0
            _ = v37
        }
    case 529:
        for fi38 := 0; fi38 < 20; fi38++ {
            v39 := []int{(fi38 & (fi38 / (((fi38 + fi38)) | 1))), fi38, -1}
            v40 := GenS7{F0: uint16(v4)}
            _ = fi38
            _ = v39
            _ = v40
        }
    case 2147483647:
        var v41 float32 = 0.1
        _ = v41
    }
    var v42 uint8 = 127
    var v43 float32 = 0.0
    fmt.Println(v1, v1, v2)
    v44 := GenS2{F0: -288, F1: v4}
    var v45 int = 487
    v46 := []int{v45, (v45 | v45)}
    _ = len(v46)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
