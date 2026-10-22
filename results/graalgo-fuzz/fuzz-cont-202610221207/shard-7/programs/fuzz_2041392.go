package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 uint64; F2 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -2147483648
    var v3 int64 = 1
    var v4 uint8 = 2
    v5 := GenS1{F0: 0, F1: 2147483647, F2: 100}
    v3 = v3
    switch 890 {
    case 4294967295:
        var v6 uint8 = v4
        _ = v6
    default:
        if false {
            v7 := GenS2{F0: 1}
            var v8 uint32 = 640
            _ = v7
            _ = v8
        } else {
            v4 = 127
        }
    }
    v2 = (^(v2))
    fmt.Println(v4, v4, v3)
    v3 += v3
    v9 := []int{0}
    switch 32767 {
    case -488:
        if true {
            v2 = v2
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
    case 386:
        var v10 bool = false
        _ = v10
    }
    if true {
        var v11 uint = 42
        var v12 float64 = 0.0
        _ = v11
        _ = v12
    } else {
        switch "\n\t" {
        case "\n\t":
            if ("hello" <= "") {
                v4 = (^(v4))
            } else {
                v13 := GenS1{F0: 42, F1: 2, F2: -43}
                for fi14 := 0; fi14 < 12; fi14++ {
                    switch ("a" + (("" + "a") + "a")) {
                    case "hello":
                        switch ("hello" + ("hello" + "")) {
                        case "\n\t":
                            if false {
                                for fi15 := 0; fi15 < 4; fi15++ {
                                    v4 -= (v4 / ((v4) | 1))
                                    v3 = (v3 - -246)
                                    _ = fi15
                                }
                            } else {
                                switch (fi14 | (fi14 ^ fi14)) {
                                case 42:
                                    v2 = -2147483648
                                case 0:
                                    for fi16 := 0; fi16 < 13; fi16++ {
                                        switch 4294967295 {
                                        case 1:
                                            _ = len(v1)
                                        default:
                                            var v17 int16 = 1
                                            _ = v17
                                        }
                                        var v18 uint16 = 0
                                        v4 -= 255
                                        _ = fi16
                                        _ = v18
                                    }
                                }
                                var v19 int32 = v2
                                _ = v19
                            }
                        default:
                            v2 = v2
                        }
                    case "":
                        v1[42] = 1
                    case "foo bar":
                        v20 := GenS1{F0: 0, F1: 2147483647, F2: 127}
                        _ = v20
                    default:
                        v21 := GenS2{F0: (v2 - (-(v2)))}
                        _ = v21
                    }
                    fmt.Println(fi14, v4)
                    v9 = append(v9, (^(fi14)))
                    _ = fi14
                }
                if false {
                    switch "hello" {
                    case "":
                        _ = len(v9)
                    case "hello":
                        switch "foo bar" {
                        case "hello":
                            v3 += int64(v2)
                        case "foo bar":
                            switch (((("\n\t" + "foo bar") + "hello") + "") + (("" + "foo bar") + "foo bar")) {
                            case "\n\t":
                                v2 = v2
                            case "hello":
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    if (32767 < 607) {
                                        for fi23 := 0; fi23 < 10; fi23++ {
                                            _ = len(v1)
                                            _ = fi23
                                        }
                                        if ("" == "foo bar") {
                                            if len(v9) > 0 {
                                                v9[0] = int(v3)
                                            }
                                            v2 = (^(v2))
                                            var v24 bool = false
                                            _ = v24
                                        }
                                        fmt.Println(v4, v2)
                                    } else {
                                        fmt.Println(v2)
                                        var v25 int64 = (v3 % ((0) | 1))
                                        _ = v25
                                    }
                                    var v26 int16 = 475
                                    _ = fi22
                                    _ = v26
                                }
                            case "a":
                                v3 += v3
                            default:
                                var v27 uint8 = 128
                                _ = v27
                            }
                        default:
                            var v28 uint32 = 440
                            _ = v28
                        }
                    default:
                        fmt.Println(v4, v4)
                    }
                }
                _ = v13
            }
        case "":
            v29 := GenS1{F0: 65535, F1: 65535, F2: 0}
            _ = v29
        case "a":
            switch -701 {
            case -962:
                v30 := GenS1{F0: 463, F1: 2147483648, F2: 127}
                _ = v30
            }
        }
        if false {
            if (97 != 1) {
                fmt.Println(v3)
                v2 -= (v2 * (v2 * v2))
                for fi31 := 0; fi31 < 15; fi31++ {
                    fmt.Println(fi31, v3)
                    switch (fi31 | 32767) {
                    case 9223372036854775807:
                        v3 = v3
                    default:
                        for fi32 := 0; fi32 < 10; fi32++ {
                            v1[42] = 851
                            _ = fi32
                        }
                    }
                    _ = fi31
                }
            }
            v33 := GenS3{F0: 173, F1: 1, F2: -32768}
            var v34 int16 = 32767
            _ = v33
            _ = v34
        } else {
            if false {
                var v35 int32 = (v2 ^ (v2 >> ((781) & 63)))
                v36 := GenS4{F0: false, F1: 126}
                for fi37 := 0; fi37 < 9; fi37++ {
                    for fi38 := 0; fi38 < 20; fi38++ {
                        var v39 uint64 = 42
                        v40 := GenS5{F0: "\n\t"}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = fi37
                }
                _ = v35
                _ = v36
            } else {
                switch "\n\t" {
                case "":
                    v41 := GenS6{F0: -32768}
                    _ = v41
                case "foo bar":
                    for fi42 := 0; fi42 < 2; fi42++ {
                        for fi43 := 0; fi43 < 10; fi43++ {
                            switch (fi43 / ((fi42) | 1)) {
                            case -575:
                                var v44 uint8 = 128
                                _ = v44
                            case 674:
                                v3 = (v3 * (-(v3)))
                            case 32767:
                                switch "\n\t" {
                                case "hello":
                                    if (!true) {
                                        if true {
                                            if (!(false || (!true))) {
                                                _ = len(v9)
                                            } else {
                                                switch fi43 {
                                                case -879:
                                                    _ = len(v1)
                                                default:
                                                    fmt.Println(fi43, fi42)
                                                }
                                            }
                                            v1[65535] = 2147483648
                                            var v45 uint32 = 370
                                            _ = v45
                                        } else {
                                            fmt.Println(fi42, fi42, fi42)
                                            var v46 int16 = int16(v2)
                                            _ = v46
                                        }
                                        v3 = int64(v2)
                                        v3 = 0
                                    }
                                case "\n\t":
                                    v47 := GenS6{F0: 1}
                                    _ = v47
                                default:
                                    var v48 int32 = int32(v4)
                                    _ = v48
                                }
                            }
                            if false {
                                v3 -= (v3 / ((881) | 1))
                                _ = len(v1)
                            }
                            if (((1 < 1000) || true) || true) {
                                fmt.Println(v4, v4)
                                if false {
                                    var v49 float32 = 0.0
                                    v3 = v3
                                    _ = v49
                                }
                            } else {
                                if (18446744073709551615 != 42) {
                                    v3 = v3
                                    v50 := GenS5{F0: "foo bar"}
                                    v51 := GenS7{F0: 100, F1: 4294967295, F2: 1}
                                    _ = v50
                                    _ = v51
                                }
                            }
                            _ = fi43
                        }
                        _ = fi42
                    }
                case "\n\t":
                    _ = len(v9)
                default:
                    var v52 string = ""
                    _ = v52
                }
            }
            var v53 int64 = v3
            var v54 string = "a"
            _ = v53
            _ = v54
        }
        v3 = (v3 % (((v3 & v3)) | 1))
    }
    var v55 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v9
}
