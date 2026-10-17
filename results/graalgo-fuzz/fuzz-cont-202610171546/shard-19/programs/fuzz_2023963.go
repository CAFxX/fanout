package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 string }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint16, p1 uint16, p2 int16) int8 {
    v6 := []int{-37}
    switch 9223372036854775807 {
    case 42:
        switch "\n\t" {
        case "\n\t":
            if false {
                if true {
                    fmt.Println(p1)
                    if (2147483648 != 0) {
                        fmt.Println(p0)
                    } else {
                        if (("a" + "\n\t") != "a") {
                            if len(v6) > 0 {
                                v6[0] = 128
                            }
                        }
                    }
                }
            }
        case "":
            switch 4294967295 {
            case -2147483648:
                v6 = append(v6, 32768)
            case -1:
                var v7 bool = false
                _ = v7
            case 32767:
                fmt.Println(p0, p2)
            }
        case "hello":
            if false {
                var v8 float32 = float32(1.5)
                var v9 uint8 = 2
                _ = v8
                _ = v9
            }
        default:
            p2 -= (p2 & p2)
        }
    case 781:
        fmt.Println(p1, p1)
    case -901:
        if (!true) {
            switch ("" + "a") {
            case "":
                p1 = (p0 << ((1) & 63))
            case "a":
                v10 := GenS1{F0: 0, F1: -1, F2: "\n\t"}
                _ = v10
            case "\n\t":
                switch -9223372036854775808 {
                case -1:
                    var v11 uint = 313
                    _ = v11
                default:
                    p2 += (p2 + p2)
                }
            }
            var v12 int32 = -2147483648
            _ = v12
        } else {
            var v13 int = 128
            _ = len(v6)
            fmt.Println(p1, v13, p2)
            _ = v13
        }
    }
    if (!true) {
        p0 = (p1 / ((p1) | 1))
        switch "" {
        case "\n\t":
            fmt.Println(p1)
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            p2 += (-(p2))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            for fi15 := 0; fi15 < 11; fi15++ {
                p0 += 430
                _ = fi15
            }
            _ = fi14
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 float32 = 1.5
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 43
}
func genFn1(p0 uint8, p1 uint8) uint64 {
    for fi18 := 0; fi18 < 10; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            p0 -= p1
            var v20 uint = 2
            _ = fi19
            _ = v20
        }
        v21 := []int{}
        _ = fi18
        _ = v21
    }
    var v22 int16 = 0
    p0 -= p1
    p1 = 127
    if (656 != 42) {
        v23 := GenS1{F0: 0, F1: -2147483648, F2: "hello"}
        p0 = 127
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v22
    return 394
}

func main() {
    var v1 float32 = 0.1
    v2 := []int{}
    var v3 bool = true
    var v4 uint64 = 1
    var v5 uint = 4294967295
    v1 = 1e10
    v24 := GenS2{F0: (^(v4)), F1: 42}
    var v25 int8 = -1
    fmt.Println(v1, v5, v1)
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v25, v3)
    for fi26 := 0; fi26 < 15; fi26++ {
        if v3 {
            v4 -= (v4 | (^(v4)))
            if ("hello" >= "foo bar") {
                v2 = append(v2, int(v5))
            }
        }
        var v27 int = (^(fi26))
        _ = fi26
        _ = v27
    }
    if (v3 || false) {
        fmt.Println(v1, v1, v3)
        switch "" {
        case "hello":
            v3 = (0 > 191)
        default:
            var v28 string = ""
            _ = v28
        }
        fmt.Println(v5)
    }
    if v3 {
        if len(v2) > 0 {
            v2[0] = 32768
        }
    }
    v29 := GenS3{F0: (v25 + genFn0(255, 65535, -32768))}
    if len(v2) > 0 {
        v2[0] = 128
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
