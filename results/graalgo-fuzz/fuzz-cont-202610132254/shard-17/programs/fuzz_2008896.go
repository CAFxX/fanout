package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint; F2 uint32 }
type GenS7 struct { F0 uint; F1 int16; F2 string }
type GenS8 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS9 struct { F0 int16 }

func main() {
    v1 := []string{"a", (("hello" + "\n\t") + "foo bar")}
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    v4 := GenS1{F0: "a", F1: true}
    fmt.Println(v3)
    var v5 int16 = 42
    v3 -= v3
    v6 := GenS1{F0: ("a" + ("" + "foo bar")), F1: (-870 > 42)}
    v5 = v5
    fmt.Println(v5)
    if ((v5 | v5) < (v5 << ((2) & 63))) {
        fmt.Println(v3)
    } else {
        var v7 int = -348
        var v8 int16 = 42
        for fi9 := 0; fi9 < 4; fi9++ {
            if false {
                for fi10 := 0; fi10 < 19; fi10++ {
                    var v11 uint16 = v3
                    var v12 uint32 = 42
                    v13 := GenS2{F0: true, F1: 2, F2: uint64(v12)}
                    var v14 float32 = 1.5
                    v8 = (v5 & v8)
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 string = ("foo bar" + ("hello" + (("\n\t" + "foo bar") + "hello")))
                var v16 int8 = 100
                _ = v15
                _ = v16
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 float64 = 0.1
        _ = fi17
        _ = v18
    }
    fmt.Println(v3, v3)
    if false {
        var v19 int32 = 1
        _ = v19
    } else {
        for fi20 := 0; fi20 < 14; fi20++ {
            if false {
                var v21 uint8 = 132
                switch ((("a" + ("a" + "hello")) + "a") + "") {
                case "\n\t":
                    var v22 uint = 1
                    _ = v22
                case "foo bar":
                    _ = len(v2)
                case "a":
                    for fi23 := 0; fi23 < 17; fi23++ {
                        var v24 int16 = -136
                        fmt.Println(v3, fi23)
                        v21 = v21
                        if (-1 > 414) {
                            _ = len(v2)
                            switch ("\n\t" + "hello") {
                            case "\n\t":
                                if true {
                                    _ = len(v2)
                                    var v25 int16 = v24
                                    _ = v25
                                }
                            case "a":
                                v1 = append(v1, "foo bar")
                            }
                            _ = len(v1)
                            v3 += v3
                        }
                        _ = fi23
                        _ = v24
                    }
                default:
                    v26 := GenS2{F0: (!(-1 < 1)), F1: 232, F2: 18446744073709551615}
                    _ = v26
                }
                _ = v21
            } else {
                if ((0 > 2147483647) && ((fi20 << ((807) & 63)) > 350)) {
                    for fi27 := 0; fi27 < 5; fi27++ {
                        var v28 uint16 = (^(v3))
                        var v29 float64 = 3.14159
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                }
                for fi30 := 0; fi30 < 7; fi30++ {
                    v31 := GenS3{F0: 2147483647, F1: 1}
                    v32 := GenS4{F0: (true || (true || (42 >= int32(v3)))), F1: -783}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                var v33 uint32 = 42
                _ = v33
            }
            var v34 float32 = 0.0
            var v35 uint16 = 980
            _ = fi20
            _ = v34
            _ = v35
        }
        v3 = 42
        switch -9223372036854775808 {
        case 32767:
            for fi36 := 0; fi36 < 11; fi36++ {
                v5 -= (-(v5))
                _ = fi36
            }
        case -9223372036854775808:
            switch "\n\t" {
            case "":
                if false {
                    _ = len(v1)
                } else {
                    switch "hello" {
                    case "foo bar":
                        var v37 uint64 = 42
                        _ = v37
                    case "a":
                        var v38 int64 = 600
                        _ = v38
                    case "":
                        switch "" {
                        case "hello":
                            if (4294967295 >= 1) {
                                fmt.Println(v3, v3)
                                if true {
                                    var v39 uint64 = 75
                                    _ = v39
                                }
                                v40 := GenS1{F0: "a", F1: false}
                                _ = v40
                            } else {
                                var v41 uint16 = (v3 & (v3 >> ((65535) & 63)))
                                for fi42 := 0; fi42 < 12; fi42++ {
                                    v43 := GenS1{F0: "a", F1: (!((2 <= v3) && ((1 >= 128) || (500 <= 0))))}
                                    switch "\n\t" {
                                    case "a":
                                        v5 -= v5
                                    case "\n\t":
                                        fmt.Println(v5, v41, v3)
                                    case "hello":
                                        var v44 uint = 65535
                                        _ = v44
                                    default:
                                        if true {
                                            var v45 bool = (-2147483648 >= 1)
                                            if v45 {
                                                v3 = v41
                                            } else {
                                                v5 = v5
                                                fmt.Println(v45, v3)
                                                v46 := GenS5{F0: v41}
                                                _ = v46
                                            }
                                            for fi47 := 0; fi47 < 20; fi47++ {
                                                fmt.Println(fi47, v45)
                                                v5 += v5
                                                v1 = append(v1, "foo bar")
                                                var v48 uint = 65535
                                                var v49 int8 = -17
                                                _ = fi47
                                                _ = v48
                                                _ = v49
                                            }
                                            _ = v45
                                        } else {
                                            v41 = (v41 + v41)
                                            fmt.Println(v41, fi42)
                                        }
                                    }
                                    v3 += 32767
                                    _ = fi42
                                    _ = v43
                                }
                                var v50 bool = false
                                _ = v41
                                _ = v50
                            }
                        case "":
                            v1 = append(v1, ("\n\t" + "a"))
                        case "a":
                            if false {
                                v3 = v3
                                v2[2147483647] = 90
                                if ((((("" + "a") + "a") + "hello") + "a") < (((("foo bar" + "hello") + "a") + "") + "")) {
                                    delete(v2, 4294967295)
                                    v51 := GenS6{F0: 0, F1: 0, F2: 33}
                                    var v52 uint8 = 127
                                    _ = v51
                                    _ = v52
                                } else {
                                    v53 := GenS7{F0: 4294967295, F1: (v5 / ((v5) | 1)), F2: ""}
                                    if true {
                                        switch -9223372036854775808 {
                                        case -520:
                                            switch "\n\t" {
                                            case "a":
                                                switch 32767 {
                                                case 32768:
                                                    _ = len(v2)
                                                default:
                                                    var v54 int8 = 100
                                                    _ = v54
                                                }
                                            case "\n\t":
                                                var v55 uint64 = 4294967295
                                                _ = v55
                                            default:
                                                fmt.Println(v3, v5)
                                            }
                                        case 806:
                                            switch 2147483647 {
                                            case 32768:
                                                v5 += -1
                                            case 42:
                                                v3 = v3
                                            case -1:
                                                v56 := GenS8{F0: 1, F1: 128, F2: 0}
                                                _ = v56
                                            }
                                        default:
                                            v5 = v5
                                        }
                                        v57 := v2[918]
                                        _ = v57
                                    }
                                    v3 += (v3 >> ((0) & 63))
                                    _ = v53
                                }
                            } else {
                                v3 += v3
                                switch "" {
                                case "hello":
                                    fmt.Println(v5, v3)
                                case "\n\t":
                                    _ = len(v1)
                                }
                                var v58 uint32 = uint32(v3)
                                _ = v58
                            }
                        }
                    }
                    for fi59 := 0; fi59 < 1; fi59++ {
                        var v60 bool = false
                        v61 := GenS8{F0: 2147483648, F1: 127, F2: 42}
                        switch "\n\t" {
                        case "":
                            var v62 string = "a"
                            _ = v62
                        case "a":
                            fmt.Println(v60, v60, v3)
                        default:
                            v1 = append(v1, ("\n\t" + (("hello" + "hello") + (("\n\t" + "\n\t") + "foo bar"))))
                        }
                        _ = fi59
                        _ = v60
                        _ = v61
                    }
                    v63 := GenS9{F0: -32768}
                    _ = v63
                }
            }
        case 127:
            if true {
                var v64 int8 = 27
                v65 := v2[927]
                var v66 int32 = -35
                _ = v64
                _ = v65
                _ = v66
            }
        }
    }
    for fi67 := 0; fi67 < 11; fi67++ {
        v5 = v5
        if true {
            if (1.5 == (1.5 * (0.1 + (1e10 + float64(((-2.25 - float32(-2.25)) * 0.1)))))) {
                v68 := v2[0]
                v69 := GenS6{F0: 0, F1: 121, F2: 4294967295}
                _ = v68
                _ = v69
            }
        } else {
            fmt.Println(fi67, fi67)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v5 += (^(v5))
        }
        _ = fi67
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
