package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 bool }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 int32; F2 int32 }
type GenS9 struct { F0 int32; F1 uint32 }
type GenS10 struct { F0 int16; F1 uint64; F2 int }
type GenS11 struct { F0 int; F1 uint64; F2 uint16 }
func genFn0(p0 int32) int16 {
    var v5 uint64 = 4294967295
    v6 := []int{2147483647}
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int32, p1 uint16, p2 int) string {
    v7 := []int{(-(p2)), p2}
    v8 := GenS1{F0: -120, F1: 0, F2: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    v1 := []int{}
    var v2 int16 = -295
    v3 := make(map[uint32]bool)
    var v4 int32 = 689
    if (0 <= (^(v2))) {
        v9 := GenS1{F0: 31, F1: 0, F2: (42 != 2147483647)}
        if false {
            switch ("hello" + ("a" + "hello")) {
            case "foo bar":
                v10 := GenS1{F0: -128, F1: 65535, F2: ("\n\t" <= "")}
                _ = v10
            }
            if false {
                switch "" {
                case "a":
                    v3[0] = false
                case "hello":
                    var v11 int16 = (v2 >> ((760) & 63))
                    _ = v11
                }
            } else {
                var v12 int64 = 9223372036854775807
                switch "\n\t" {
                case "foo bar":
                    v13 := GenS2{F0: -9223372036854775808, F1: 133}
                    _ = v13
                case "a":
                    v12 = (v12 | v12)
                case "\n\t":
                    if false {
                        if true {
                            var v14 float64 = -2.25
                            _ = v14
                        }
                        for fi15 := 0; fi15 < 9; fi15++ {
                            for fi16 := 0; fi16 < 19; fi16++ {
                                v17 := GenS1{F0: 42, F1: 4294967295, F2: false}
                                _ = fi16
                                _ = v17
                            }
                            v2 += 32767
                            v18 := v3[2]
                            _ = fi15
                            _ = v18
                        }
                        v4 += v4
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch 0 {
                        case 0:
                            v19 := GenS3{F0: 127}
                            _ = v19
                        case 32767:
                            var v20 uint8 = 128
                            _ = v20
                        }
                    }
                default:
                    v1 = append(v1, 252)
                }
                _ = v12
            }
        } else {
            if (3.14159 == float32(0.0)) {
                if false {
                    v2 += v2
                    v21 := GenS4{F0: 0, F1: 241, F2: 4294967295}
                    v22 := GenS5{F0: (v2 ^ genFn0((-(v4)))), F1: -128}
                    _ = v21
                    _ = v22
                } else {
                    fmt.Println(v2)
                    _ = len(v3)
                    _ = len(v3)
                }
            } else {
                switch "\n\t" {
                case "a":
                    for fi23 := 0; fi23 < 10; fi23++ {
                        _ = len(v3)
                        switch "hello" {
                        case "":
                            _ = len(v3)
                        default:
                            var v24 uint32 = 2147483647
                            _ = v24
                        }
                        _ = fi23
                    }
                case "":
                    v25 := v3[148]
                    _ = v25
                case "foo bar":
                    if false {
                        var v26 int64 = 9223372036854775807
                        v26 = 42
                        for fi27 := 0; fi27 < 14; fi27++ {
                            var v28 string = ("\n\t" + "a")
                            v29 := GenS5{F0: 1000, F1: 127}
                            switch "a" {
                            case "hello":
                                fmt.Println(v4, fi27, v28)
                            case "foo bar":
                                v2 -= genFn0(v4)
                            default:
                                switch (fi27 & int(v26)) {
                                case 1:
                                    v2 = v2
                                default:
                                    fmt.Println(v28, v26)
                                }
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        _ = v26
                    } else {
                        var v30 uint64 = 0
                        if false {
                            for fi31 := 0; fi31 < 11; fi31++ {
                                var v32 uint8 = 128
                                var v33 int16 = (v2 % ((genFn0(173)) | 1))
                                if false {
                                    var v34 uint64 = v30
                                    v34 = (^(v34))
                                    _ = v34
                                } else {
                                    var v35 int16 = (v33 >> ((4294967295) & 63))
                                    v3[4294967295] = true
                                    var v36 uint64 = 2
                                    _ = v35
                                    _ = v36
                                }
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        } else {
                            var v37 uint16 = 65535
                            if false {
                                var v38 float32 = (1.5 - 1.5)
                                var v39 uint8 = 63
                                var v40 string = ("" + "a")
                                _ = v38
                                _ = v39
                                _ = v40
                            } else {
                                if false {
                                    v41 := GenS4{F0: 9223372036854775807, F1: 42, F2: 2147483648}
                                    _ = v41
                                } else {
                                    for fi42 := 0; fi42 < 6; fi42++ {
                                        v37 = v37
                                        switch 32768 {
                                        case -408:
                                            fmt.Println(v4, v37, v2)
                                        case -2147483648:
                                            v43 := GenS6{F0: (!false)}
                                            _ = v43
                                        case 32767:
                                            var v44 string = ""
                                            _ = v44
                                        }
                                        fmt.Println(v30, v4)
                                        _ = fi42
                                    }
                                }
                                v30 += (^(v30))
                            }
                            _ = v37
                        }
                        v4 -= 42
                        _ = v30
                    }
                }
            }
            for fi45 := 0; fi45 < 2; fi45++ {
                v2 += v2
                switch fi45 {
                case 220:
                    if (42 >= 42) {
                        var v46 int64 = 1
                        var v47 uint32 = 527
                        _ = v46
                        _ = v47
                    }
                case 32768:
                    var v48 uint32 = 2
                    _ = v48
                }
                _ = fi45
            }
            for fi49 := 0; fi49 < 11; fi49++ {
                switch ("" + (("" + (("a" + (("hello" + "foo bar") + "\n\t")) + ("\n\t" + "a"))) + "hello")) {
                case "hello":
                    fmt.Println(fi49)
                }
                v50 := GenS7{F0: -9223372036854775808}
                _ = fi49
                _ = v50
            }
        }
        v2 += genFn0(-1)
        _ = v9
    } else {
        fmt.Println(v4)
        v4 = (v4 | -65)
        var v51 int16 = v2
        _ = v51
    }
    for fi52 := 0; fi52 < 9; fi52++ {
        v53 := GenS8{F0: -1, F1: 0, F2: (v4 - 42)}
        _ = len(v1)
        _ = fi52
        _ = v53
    }
    var v54 int32 = v4
    fmt.Println(v54)
    var v55 uint = 278
    v4 = (v4 / ((int32(v2)) | 1))
    v56 := GenS7{F0: -1}
    v57 := GenS9{F0: v4, F1: 4294967295}
    var v58 int8 = 1
    v58 += -128
    fmt.Println(v54, v55, v4)
    v59 := GenS10{F0: (v2 ^ genFn0(int32(v58))), F1: 18446744073709551615, F2: -787}
    v60 := GenS9{F0: 2147483647, F1: 4294967295}
    for fi61 := 0; fi61 < 20; fi61++ {
        v2 = 1
        _ = fi61
    }
    v4 = (v4 << ((0) & 63))
    v62 := GenS11{F0: 32768, F1: 18446744073709551615, F2: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v62
}
