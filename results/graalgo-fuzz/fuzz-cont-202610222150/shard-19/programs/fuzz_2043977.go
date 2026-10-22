package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []uint32{65535, 1}
    v2 := []uint32{4294967295, 122}
    v3 := []int{0, 1, 32768}
    var v4 int16 = -32768
    v5 := []uint32{4294967295}
    if false {
        switch "foo bar" {
        case "hello":
            if false {
                var v6 float32 = (3.14159 * 0.0)
                _ = v6
            }
        case "\n\t":
            var v7 uint64 = 0
            _ = v7
        }
        switch ("hello" + ("a" + "a")) {
        case "\n\t":
            v8 := GenS1{F0: 1, F1: (^(v4))}
            _ = v8
        case "hello":
            fmt.Println(v4)
        case "":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := GenS1{F0: 42, F1: v4}
            _ = fi9
            _ = v10
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if false {
            var v11 uint16 = 2
            if ((-543 > v4) && (uint32(v11) <= 128)) {
                for fi12 := 0; fi12 < 6; fi12++ {
                    _ = len(v2)
                    v11 = v11
                    if false {
                        fmt.Println(v4, v4, v11)
                        switch fi12 {
                        case -9223372036854775808:
                            var v13 int32 = 491
                            _ = v13
                        case -376:
                            v11 += (v11 * v11)
                        }
                        v4 = (v4 % ((v4) | 1))
                    }
                    _ = fi12
                }
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 int64 = -1
                    v16 := GenS1{F0: 0, F1: (v4 * 205)}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                v3 = append(v3, 32767)
            } else {
                if (-9223372036854775808 != 42) {
                    fmt.Println(v11)
                    fmt.Println(v11, v4)
                    v17 := GenS1{F0: v11, F1: (v4 & v4)}
                    _ = v17
                } else {
                    var v18 uint64 = 1
                    switch "a" {
                    case "a":
                        var v19 uint64 = v18
                        _ = v19
                    case "foo bar":
                        v4 = v4
                    case "hello":
                        v20 := GenS2{F0: v18}
                        _ = v20
                    default:
                        v1 = append(v1, 995)
                    }
                    _ = len(v3)
                    _ = v18
                }
                fmt.Println(v4, v11)
            }
            _ = v11
        } else {
            var v21 float32 = float32(0.1)
            v21 += v21
            _ = v21
        }
        fmt.Println(v4)
    }
    fmt.Println(v4)
    for fi22 := 0; fi22 < 12; fi22++ {
        v4 -= v4
        _ = fi22
    }
    fmt.Println(v4)
    _ = len(v5)
    fmt.Println(v4)
    switch -406 {
    case 2147483647:
        v4 = 1
    case -650:
        fmt.Println(v4)
    case -895:
        v4 -= 418
    }
    v5 = append(v5, 818)
    var v23 float32 = (0.0 - 3.14159)
    v4 = (v4 | v4)
    var v24 int64 = 9223372036854775807
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v24 += v24
    v23 = v23
    if false {
        switch 212 {
        case -9223372036854775808:
            var v25 int = 0
            _ = v25
        case -395:
            _ = len(v5)
        default:
            for fi26 := 0; fi26 < 12; fi26++ {
                v23 = float32(0.0)
                var v27 float64 = 0.1
                _ = fi26
                _ = v27
            }
        }
        var v28 int64 = (v24 ^ 1)
        _ = v28
    } else {
        var v29 float64 = 0.0
        fmt.Println(v4, v24)
        var v30 uint8 = 1
        _ = v29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
