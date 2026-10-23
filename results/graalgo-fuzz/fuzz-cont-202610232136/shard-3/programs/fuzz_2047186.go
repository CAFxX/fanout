package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint64; F2 uint }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int }
func genFn0() uint64 {
    var v4 int8 = 85
    v5 := GenS1{F0: 676}
    v6 := []int{-636, 32767, 4294967295}
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int8, p1 int8) uint32 {
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(p1)
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        if (0 < 2) {
            for fi10 := 0; fi10 < 15; fi10++ {
                switch "hello" {
                case "\n\t":
                    var v11 int16 = -289
                    _ = v11
                case "a":
                    switch ("" + "\n\t") {
                    case "foo bar":
                        if (1 > 42) {
                            var v12 uint8 = 1
                            p0 += 127
                            _ = v12
                        } else {
                            v13 := GenS2{F0: 821}
                            var v14 float32 = 3.14159
                            var v15 int16 = 42
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                    default:
                        fmt.Println(p1)
                    }
                case "foo bar":
                    v16 := []int{fi10, fi9}
                    _ = v16
                }
                fmt.Println(fi9, fi9)
                _ = fi10
            }
        }
        p0 += -128
        v17 := GenS3{F0: 42, F1: -270}
        _ = fi9
        _ = v17
    }
    var v18 int16 = 0
    v19 := GenS4{F0: 0}
    p1 = p0
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 65535
}

func main() {
    v1 := []int{2147483647, 32767, 2147483647}
    var v2 int16 = 0
    var v3 int8 = 100
    v3 = v3
    switch 293 {
    case 127:
        v20 := GenS5{F0: ((("hello" + "foo bar") + "\n\t") + "hello"), F1: 42, F2: 4294967295}
        _ = v20
    }
    if (false && (505 == 390)) {
        if (!false) {
            v3 = -128
        } else {
            fmt.Println(v2, v3)
            if (false && (true && true)) {
                v3 = 42
                if len(v1) > 0 {
                    v1[0] = 2147483647
                }
            }
            fmt.Println(v2, v2)
        }
    } else {
        _ = len(v1)
        v3 = (v3 ^ v3)
    }
    v21 := GenS5{F0: (("" + ("" + "\n\t")) + "foo bar"), F1: 788, F2: 42}
    if (!true) {
        v3 = v3
    }
    var v22 int = -1
    v1 = append(v1, -9223372036854775808)
    switch -1 {
    case -9223372036854775808:
        switch "\n\t" {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            fmt.Println(v22)
        case "foo bar":
            for fi23 := 0; fi23 < 9; fi23++ {
                for fi24 := 0; fi24 < 20; fi24++ {
                    if len(v1) > 0 {
                        v1[0] = fi23
                    }
                    v25 := GenS6{F0: (v22 * v22), F1: "\n\t"}
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 8; fi26++ {
                    if len(v1) > 0 {
                        v1[0] = (fi23 | (^(fi23)))
                    }
                    v27 := GenS2{F0: 65535}
                    var v28 bool = true
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = fi23
            }
        default:
            v29 := GenS5{F0: "hello", F1: 2, F2: 42}
            _ = v29
        }
    case -793:
        if true {
            fmt.Println(v3)
            for fi30 := 0; fi30 < 13; fi30++ {
                fmt.Println(v3)
                _ = fi30
            }
        }
    default:
        v22 = v22
    }
    _ = len(v1)
    var v31 uint8 = 255
    v3 = (v3 << ((65535) & 63))
    var v32 int8 = (v3 & (v3 * (v3 | v3)))
    v2 = (v2 + v2)
    v33 := GenS3{F0: 155, F1: 196}
    v34 := GenS7{F0: int(v31)}
    fmt.Println(v22)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
