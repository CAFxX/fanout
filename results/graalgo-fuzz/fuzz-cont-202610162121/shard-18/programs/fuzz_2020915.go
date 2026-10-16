package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint; F2 uint64 }
func genFn0(p0 int) int16 {
    p0 -= (p0 % ((p0) | 1))
    p0 += -2147483648
    var v5 int64 = 1
    _ = p0
    _ = v5
    return -32768
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]int)
    var v3 float64 = (float64(3.14159) - 0.0)
    var v4 bool = (6 == 42)
    var v6 uint = v1
    v6 = (v6 / (((^(v6))) | 1))
    if true {
        v3 -= ((v3 / 0.1) / (((float64(0.1) - float64(0.0)) / 1.0) / ((v3 + (v3 / -2.25)) - v3)))
    }
    v7 := []int{654, 4294967295, 128}
    if len(v7) > 0 {
        v7[0] = 42
    }
    fmt.Println(v4, v3, v4)
    v6 += v6
    v8 := GenS1{F0: 42}
    v6 = (v1 % (((v6 - v1)) | 1))
    var v9 int16 = -14
    if (v4 && ((true && v4) && v4)) {
        switch "a" {
        case "a":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case "foo bar":
            v1 = 42
        case "\n\t":
            switch "\n\t" {
            case "hello":
                for fi10 := 0; fi10 < 2; fi10++ {
                    switch "" {
                    case "foo bar":
                        v3 = (v3 - 1e10)
                    case "\n\t":
                        switch ("\n\t" + "") {
                        case "":
                            if (("foo bar" == ("\n\t" + (("\n\t" + "foo bar") + ("hello" + "foo bar")))) || v4) {
                                for fi11 := 0; fi11 < 2; fi11++ {
                                    switch fi11 {
                                    case 2147483647:
                                        v9 += (^(v9))
                                    case 42:
                                        fmt.Println(fi10, v4, v1)
                                    }
                                    if v4 {
                                        v9 += genFn0(-649)
                                        v2[fi10] = fi11
                                        v12 := GenS2{F0: (fi11 | int(v9)), F1: 1}
                                        var v13 int8 = -40
                                        _ = v12
                                        _ = v13
                                    } else {
                                        switch (fi11 * (fi10 | 61)) {
                                        case 32767:
                                            var v14 float64 = v3
                                            _ = v14
                                        case 847:
                                            if false {
                                                v15 := GenS3{F0: 2}
                                                _ = v15
                                            } else {
                                                _ = len(v2)
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                            }
                                        }
                                    }
                                    _ = fi11
                                }
                                if true {
                                    delete(v2, (^(fi10)))
                                } else {
                                    var v16 uint64 = 0
                                    _ = v16
                                }
                                v17 := v2[(^(fi10))]
                                _ = v17
                            } else {
                                var v18 int32 = 1
                                if ((fi10 >= (fi10 + 2147483647)) || (255 >= 95)) {
                                    if v4 {
                                        delete(v2, fi10)
                                        var v19 float64 = 0.1
                                        switch -2147483648 {
                                        case -569:
                                            var v20 uint16 = 0
                                            _ = v20
                                        default:
                                            v9 += (-(v9))
                                        }
                                        _ = v19
                                    }
                                    if (v4 || v4) {
                                        for fi21 := 0; fi21 < 13; fi21++ {
                                            var v22 uint = v6
                                            v7 = append(v7, -9223372036854775808)
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            _ = fi21
                                            _ = v22
                                        }
                                        for fi23 := 0; fi23 < 8; fi23++ {
                                            switch "" {
                                            case "foo bar":
                                                var v24 uint = (v6 << (((v1 << (((v6 * (v1 & v1))) & 63))) & 63))
                                                _ = v24
                                            case "hello":
                                                fmt.Println(fi10)
                                            case "":
                                                v25 := v2[fi23]
                                                _ = v25
                                            }
                                            fmt.Println(v4, v4)
                                            _ = fi23
                                        }
                                    } else {
                                        delete(v2, (fi10 % ((1) | 1)))
                                        if v4 {
                                            fmt.Println(v18, v4, v6)
                                        } else {
                                            switch (fi10 / ((fi10) | 1)) {
                                            case 527:
                                                v3 = v3
                                            case 32767:
                                                if v4 {
                                                    v9 = int16(fi10)
                                                }
                                            default:
                                                v4 = false
                                            }
                                            v26 := v2[-1]
                                            _ = v26
                                        }
                                    }
                                    v27 := GenS1{F0: 0}
                                    _ = v27
                                }
                                _ = v18
                            }
                        default:
                            fmt.Println(fi10, v3)
                        }
                    case "":
                        fmt.Println(fi10)
                    }
                    _ = fi10
                }
            default:
                var v28 float32 = -2.25
                _ = v28
            }
        default:
            if len(v7) > 0 {
                v7[0] = 0
            }
        }
        fmt.Println(v3, v3)
        v29 := GenS4{F0: 9223372036854775807}
        _ = v29
    }
    var v30 uint16 = 1
    switch "a" {
    case "":
        var v31 float32 = 1e10
        _ = v31
    case "a":
        v32 := GenS5{F0: "hello", F1: (v6 | uint(v30)), F2: 2}
        _ = v32
    case "hello":
        v4 = v4
    }
    if false {
        for fi33 := 0; fi33 < 6; fi33++ {
            for fi34 := 0; fi34 < 17; fi34++ {
                v1 = v1
                switch (("a" + "hello") + (("" + "hello") + ("foo bar" + "a"))) {
                case "foo bar":
                    switch (("a" + "hello") + "") {
                    case "\n\t":
                        v6 -= (v1 >> ((255) & 63))
                    case "":
                        fmt.Println(v30, v4, fi34)
                    case "foo bar":
                        fmt.Println(fi33, v6, v6)
                    default:
                        v1 = v1
                    }
                case "":
                    fmt.Println(fi34, fi33)
                case "hello":
                    var v35 int8 = -41
                    _ = v35
                }
                v2[(fi33 & (fi33 - 32767))] = (-(fi33))
                _ = fi34
            }
            var v36 int64 = -9223372036854775808
            if false {
                if (2 < 2) {
                    v3 -= v3
                }
                v37 := GenS1{F0: 42}
                if v4 {
                    v9 -= v9
                }
                _ = v37
            } else {
                var v38 int32 = 1
                _ = v38
            }
            _ = fi33
            _ = v36
        }
        var v39 int32 = -2147483648
        _ = v39
    }
    var v40 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
