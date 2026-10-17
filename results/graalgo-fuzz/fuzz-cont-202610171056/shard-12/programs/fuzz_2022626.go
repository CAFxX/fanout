package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int; F1 bool; F2 string }
func genFn0(p0 int16, p1 int16) string {
    v7 := GenS1{F0: (true || (2147483648 > 4294967295)), F1: 2, F2: 42}
    v8 := []int{127, 4294967295, 0}
    var v9 int64 = -652
    var v10 uint16 = 715
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return "a"
}
func genFn1(p0 int32, p1 uint32, p2 int8) int8 {
    var v11 float64 = ((-2.25 + 1.5) / -2.25)
    fmt.Println(p0, p1)
    v12 := GenS1{F0: false, F1: 2, F2: -816}
    var v13 int8 = p2
    var v14 string = ((("hello" + ("foo bar" + "foo bar")) + "a") + "")
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return -128
}

func main() {
    v1 := []string{("" + ("\n\t" + "foo bar")), "\n\t"}
    var v2 int16 = -348
    v3 := []uint32{777, 269}
    var v4 float64 = (0.0 / -2.25)
    v5 := []string{"foo bar"}
    var v6 int16 = v2
    for fi15 := 0; fi15 < 9; fi15++ {
        if true {
            var v16 uint8 = 42
            v2 = -992
            _ = v16
        }
        _ = fi15
    }
    v17 := GenS1{F0: (1 == 255), F1: 42, F2: 32768}
    v18 := GenS1{F0: (-256 >= 32), F1: 127, F2: 273}
    v19 := GenS2{F0: -9223372036854775808, F1: 2, F2: 42}
    var v20 int32 = 0
    v6 -= (v6 / ((v6) | 1))
    if len(v3) > 0 {
        v3[0] = 165
    }
    if (false && (32768 != 1)) {
        for fi21 := 0; fi21 < 18; fi21++ {
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 int64 = -1
                _ = fi22
                _ = v23
            }
            var v24 uint = 18446744073709551615
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
            _ = fi21
            _ = v24
        }
        v3 = append(v3, 138)
    } else {
        fmt.Println(v4)
        var v25 uint16 = 0
        for fi26 := 0; fi26 < 19; fi26++ {
            for fi27 := 0; fi27 < 8; fi27++ {
                fmt.Println(v2, v6)
                _ = fi27
            }
            v1 = append(v1, "a")
            if true {
                fmt.Println(v4, v4)
            } else {
                var v28 uint32 = 897
                var v29 bool = true
                _ = v28
                _ = v29
            }
            _ = fi26
        }
        _ = v25
    }
    fmt.Println(v2, v20, v6)
    for fi30 := 0; fi30 < 7; fi30++ {
        v6 = (-(v2))
        _ = fi30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        fmt.Println(v6)
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        if true {
            switch "a" {
            case "":
                switch "hello" {
                case "\n\t":
                    switch "foo bar" {
                    case "":
                        if (265 > 42) {
                            v32 := GenS3{F0: 75, F1: 4294967295}
                            for fi33 := 0; fi33 < 11; fi33++ {
                                switch "\n\t" {
                                case "":
                                    if (1 != 42) {
                                        if true {
                                            fmt.Println(v20, v2)
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                            var v34 float32 = 3.14159
                                            _ = v34
                                        }
                                        v6 += v2
                                    }
                                case "hello":
                                    v35 := GenS2{F0: 9223372036854775807, F1: 1, F2: 18446744073709551615}
                                    _ = v35
                                case "a":
                                    switch fi33 {
                                    case -903:
                                        if true {
                                            if (true || (!true)) {
                                                var v36 uint8 = 255
                                                fmt.Println(v6, v36, fi33)
                                                v37 := GenS4{F0: (fi31 * (fi33 | (fi33 << ((74) & 63)))), F1: (v36 == (v36 - (v36 ^ 238))), F2: "hello"}
                                                if (!false) {
                                                    var v38 uint32 = 291
                                                    _ = v38
                                                }
                                                _ = v36
                                                _ = v37
                                            } else {
                                                v39 := GenS4{F0: (fi33 | fi33), F1: (!((!false) && (!("a" >= "\n\t")))), F2: "hello"}
                                                var v40 uint64 = 42
                                                fmt.Println(fi33, v4)
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                    case -930:
                                        var v41 uint16 = 1
                                        _ = v41
                                    case 9223372036854775807:
                                        fmt.Println(v4, v20, v2)
                                    default:
                                        fmt.Println(fi31, fi31, v2)
                                    }
                                }
                                _ = fi33
                            }
                            _ = v32
                        } else {
                            for fi42 := 0; fi42 < 3; fi42++ {
                                switch "" {
                                case "\n\t":
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                case "hello":
                                    switch "a" {
                                    case "":
                                        if false {
                                            if (18446744073709551615 != 2) {
                                                v4 = -2.25
                                            }
                                            fmt.Println(fi42, v4, fi42)
                                        }
                                    case "\n\t":
                                        v5 = append(v5, ("" + "hello"))
                                    default:
                                        v5 = append(v5, "a")
                                    }
                                }
                                var v43 float64 = v4
                                _ = len(v1)
                                _ = fi42
                                _ = v43
                            }
                            var v44 int8 = -37
                            _ = v44
                        }
                    case "a":
                        v1 = append(v1, "hello")
                    }
                case "a":
                    var v45 uint32 = 635
                    _ = v45
                case "foo bar":
                    fmt.Println(v2, v6, v2)
                }
            case "\n\t":
                if (127 > 42) {
                    var v46 int8 = 44
                    var v47 int64 = 1
                    for fi48 := 0; fi48 < 9; fi48++ {
                        v3 = append(v3, 42)
                        v49 := GenS2{F0: v47, F1: 850, F2: 42}
                        _ = fi48
                        _ = v49
                    }
                    _ = v46
                    _ = v47
                }
            case "foo bar":
                switch ("hello" + "a") {
                case "":
                    v3 = append(v3, 42)
                default:
                    _ = len(v5)
                }
            }
            var v50 int64 = -264
            _ = v50
        }
        _ = fi31
    }
    for fi51 := 0; fi51 < 15; fi51++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = len(v5)
        v20 += (v20 % ((-2147483648) | 1))
        _ = fi51
    }
    if (0 == 255) {
        var v52 int16 = (v6 & v6)
        _ = v52
    }
    v6 = 1000
    v20 = (v20 - v20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
