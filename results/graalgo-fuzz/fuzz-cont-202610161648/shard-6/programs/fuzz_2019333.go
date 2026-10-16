package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 bool; F2 string }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 uint32; F2 string }
func genFn0(p0 uint) uint64 {
    p0 += (p0 + p0)
    if true {
        v6 := []int{0}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v6
    } else {
        fmt.Println(p0)
    }
    if true {
        var v7 int16 = 1
        _ = v7
    }
    _ = p0
    return 581
}

func main() {
    var v1 int16 = -32768
    var v2 float32 = 3.14159
    v3 := make(map[string]bool)
    var v4 bool = (!true)
    var v5 uint8 = 255
    fmt.Println(v4, v2, v5)
    if (2 <= 2) {
        var v8 uint32 = 4294967295
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 int8 = 1
            _ = fi9
            _ = v10
        }
        var v11 uint = 1
        _ = v11
        _ = v8
    }
    var v12 float32 = v2
    switch 2147483647 {
    case -405:
        v2 += v2
    case 32768:
        v13 := GenS1{F0: 32768}
        _ = v13
    case 635:
        v14 := GenS1{F0: 0}
        _ = v14
    default:
        switch -377 {
        case 32767:
            v12 = 3.14159
        case 42:
            v15 := GenS2{F0: v1, F1: ((genFn0(2) >= 79) && v4), F2: ("" + "\n\t")}
            _ = v15
        case -406:
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 uint16 = 42
                _ = fi16
                _ = v17
            }
        }
    }
    var v18 uint16 = 65535
    if ((!v4) || true) {
        v18 -= 255
        if v4 {
            if (!true) {
                v12 += v2
                v19 := GenS1{F0: 9223372036854775807}
                if (v4 && true) {
                    v20 := GenS1{F0: 528}
                    _ = v20
                } else {
                    v2 = float32(-2.25)
                }
                _ = v19
            }
        } else {
            var v21 float64 = 0.0
            _ = v21
        }
        v12 -= -2.25
    }
    v22 := GenS3{F0: 9223372036854775807, F1: 1}
    v4 = (genFn0(42) == 255)
    v2 += v2
    fmt.Println(v1, v18)
    var v23 uint64 = 72
    v24 := GenS2{F0: 0, F1: false, F2: "a"}
    v25 := []int{0}
    var v26 int32 = -884
    if true {
        if len(v25) > 0 {
            v25[0] = 9223372036854775807
        }
        switch "hello" {
        case "":
            switch 329 {
            case 32768:
                if (1.5 > 3.14159) {
                    var v27 int8 = 0
                    v23 = 2
                    for fi28 := 0; fi28 < 6; fi28++ {
                        v5 = uint8(v23)
                        v29 := GenS4{F0: "\n\t"}
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                } else {
                    switch "" {
                    case "":
                        v18 -= (v18 | v18)
                    case "\n\t":
                        for fi30 := 0; fi30 < 9; fi30++ {
                            fmt.Println(v2)
                            _ = fi30
                        }
                    }
                    var v31 uint16 = (^(v18))
                    _ = v31
                }
            case 127:
                switch ("" + ("foo bar" + ("hello" + ""))) {
                case "hello":
                    for fi32 := 0; fi32 < 1; fi32++ {
                        var v33 int = fi32
                        v34 := GenS5{F0: -1, F1: uint32(v5), F2: "hello"}
                        v18 += v18
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                case "a":
                    var v35 bool = (1 < (v5 & 255))
                    _ = v35
                default:
                    v25 = append(v25, -2147483648)
                }
            default:
                v1 += v1
            }
        }
        v1 += (v1 ^ 32767)
    }
    if v4 {
        switch "a" {
        case "a":
            var v36 uint64 = (v23 / ((genFn0(18446744073709551615)) | 1))
            _ = v36
        case "hello":
            if (32767 > v1) {
                v18 += (^(v18))
                fmt.Println(v2, v23)
            }
        default:
            fmt.Println(v23, v26)
        }
        fmt.Println(v12, v2)
        var v37 uint8 = 2
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
