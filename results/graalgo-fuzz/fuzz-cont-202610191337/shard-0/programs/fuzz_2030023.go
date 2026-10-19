package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 int8; F2 bool }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 bool }
func genFn0(p0 uint8) int {
    v6 := GenS1{F0: 100}
    v7 := []int{1, 42, -653, 32767}
    if (2 > 2) {
        v8 := GenS2{F0: 851, F1: 285}
        _ = v8
    }
    fmt.Println(p0)
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 float32 = 1.5
        fmt.Println(fi9, p0, fi9)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int16 = 116
    v2 := make(map[uint32]bool)
    v3 := make(map[string]bool)
    var v4 float32 = 1.5
    v5 := []int{9223372036854775807, 2147483647}
    v4 += 0.1
    var v11 float64 = -2.25
    for fi12 := 0; fi12 < 1; fi12++ {
        if true {
            _ = len(v2)
            v2[1] = false
        }
        _ = fi12
    }
    _ = len(v5)
    fmt.Println(v1, v1, v1)
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := GenS3{F0: 1, F1: false}
        switch fi13 {
        case 42:
            fmt.Println(fi13, v1)
        case -438:
            for fi15 := 0; fi15 < 11; fi15++ {
                v2[610] = (genFn0(127) <= 32767)
                v4 -= 0.1
                var v16 string = ("foo bar" + ("" + ("foo bar" + "")))
                var v17 int8 = -48
                var v18 float64 = 1e10
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        case -492:
            var v19 int32 = 1
            _ = v19
        default:
            if true {
                for fi20 := 0; fi20 < 19; fi20++ {
                    switch ("a" + "hello") {
                    case "a":
                        fmt.Println(v11, v4, v1)
                    case "foo bar":
                        if true {
                            delete(v3, "a")
                            v4 += v4
                        } else {
                            switch "\n\t" {
                            case "a":
                                if true {
                                    switch 603 {
                                    case 128:
                                        fmt.Println(v1)
                                        var v21 bool = true
                                        var v22 float32 = 0.1
                                        if (65535 != 2147483648) {
                                            _ = len(v3)
                                            switch fi20 {
                                            case 127:
                                                var v23 float64 = v11
                                                _ = v23
                                            case 42:
                                                var v24 uint64 = 1
                                                _ = v24
                                            }
                                            v1 = v1
                                        }
                                        _ = v21
                                        _ = v22
                                    case 992:
                                        if true {
                                            v25 := GenS2{F0: 1, F1: -365}
                                            switch ("\n\t" + "a") {
                                            case "a":
                                                var v26 uint64 = 4294967295
                                                _ = v26
                                            case "\n\t":
                                                var v27 uint64 = 18446744073709551615
                                                _ = v27
                                            case "hello":
                                                v28 := GenS4{F0: 1000}
                                                _ = v28
                                            default:
                                                v1 -= v1
                                            }
                                            _ = v25
                                        } else {
                                            if false {
                                                v29 := GenS1{F0: -128}
                                                _ = v29
                                            } else {
                                                var v30 int = -9223372036854775808
                                                v31 := v3[(("" + "") + ("foo bar" + ("a" + "a")))]
                                                var v32 uint32 = 2147483648
                                                v33 := GenS5{F0: 1, F1: -1, F2: (!v31)}
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            }
                                            var v34 float32 = 1e10
                                            _ = v34
                                        }
                                    case 2147483647:
                                        var v35 int32 = 42
                                        _ = v35
                                    }
                                    v36 := GenS6{F0: 1, F1: "a"}
                                    v4 -= ((float32(v11) / v4) * 0.1)
                                    fmt.Println(v11)
                                    _ = v36
                                } else {
                                    _ = len(v2)
                                }
                            default:
                                v37 := GenS7{F0: -9223372036854775808}
                                _ = v37
                            }
                            if (true || (115 > 65535)) {
                                switch fi20 {
                                case 32767:
                                    fmt.Println(fi13, v4, fi20)
                                case -9223372036854775808:
                                    if (300 == 42) {
                                        v38 := GenS8{F0: true, F1: (-7 > -128)}
                                        _ = v38
                                    } else {
                                        v4 = -2.25
                                        var v39 int16 = v1
                                        v1 = (v1 / ((0) | 1))
                                        _ = v39
                                    }
                                default:
                                    switch "a" {
                                    case "a":
                                        v4 -= (v4 / v4)
                                    }
                                }
                                fmt.Println(v4, v1, v11)
                                fmt.Println(fi20)
                            }
                            switch (fi20 ^ (fi20 - -810)) {
                            case 9223372036854775807:
                                if (0 <= 127) {
                                    delete(v3, "")
                                    var v40 uint = 18446744073709551615
                                    v41 := GenS4{F0: 42}
                                    _ = v40
                                    _ = v41
                                }
                            }
                        }
                    }
                    var v42 int64 = -1
                    v4 = ((v4 / -2.25) + v4)
                    _ = fi20
                    _ = v42
                }
            }
        }
        _ = fi13
        _ = v14
    }
    _ = len(v5)
    v4 -= v4
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 uint16 = 195
        var v45 uint32 = 1
        var v46 int16 = v1
        _ = fi43
        _ = v44
        _ = v45
        _ = v46
    }
    v4 -= 3.14159
    v47 := GenS3{F0: 601, F1: (58 <= 1)}
    for fi48 := 0; fi48 < 6; fi48++ {
        var v49 int16 = -515
        fmt.Println(v1, v11, v49)
        v4 = float32(v11)
        _ = fi48
        _ = v49
    }
    v2[65535] = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
