package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int; F2 string }
type GenS4 struct { F0 uint16; F1 string; F2 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool }
func genFn0(p0 int16, p1 int16, p2 uint16) uint8 {
    p2 = p2
    v4 := GenS1{F0: ("hello" + ("a" + "")), F1: 870}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 127
}
func genFn1() int64 {
    var v5 uint = 679
    var v6 int32 = -597
    v5 = v5
    v6 = -864
    v6 += (-(v6))
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []int{-692, -2147483648, -9223372036854775808, -1}
    v2 := make(map[string]string)
    var v3 int8 = 100
    if true {
        v7 := GenS1{F0: (("hello" + (("a" + "") + ("\n\t" + ""))) + ""), F1: 2}
        for fi8 := 0; fi8 < 6; fi8++ {
            _ = len(v2)
            if false {
                switch -283 {
                case -624:
                    switch "a" {
                    case "\n\t":
                        var v9 string = "hello"
                        _ = v9
                    case "hello":
                        switch (fi8 << ((65535) & 63)) {
                        case -9223372036854775808:
                            if (!(false || true)) {
                                var v10 bool = true
                                v3 = (v3 >> ((0) & 63))
                                v11 := v2[(("a" + "\n\t") + "")]
                                _ = v10
                                _ = v11
                            } else {
                                v12 := GenS1{F0: "a", F1: 4294967295}
                                v13 := GenS1{F0: "\n\t", F1: 4294967295}
                                for fi14 := 0; fi14 < 4; fi14++ {
                                    var v15 bool = false
                                    _ = fi14
                                    _ = v15
                                }
                                _ = v12
                                _ = v13
                            }
                        case 32768:
                            v16 := GenS1{F0: "", F1: 255}
                            _ = v16
                        case 127:
                            if ((fi8 + (fi8 * fi8)) >= (fi8 & fi8)) {
                                for fi17 := 0; fi17 < 7; fi17++ {
                                    v18 := GenS2{F0: 1}
                                    var v19 int64 = -1
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            }
                        default:
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = fi8
                                }
                                v3 += (v3 & 100)
                                if false {
                                    switch (fi8 | fi8) {
                                    case 2147483647:
                                        var v20 uint32 = 2147483648
                                        _ = v20
                                    case 128:
                                        v21 := GenS3{F0: fi8, F1: -2147483648, F2: ""}
                                        _ = v21
                                    case -977:
                                        fmt.Println(v3, v3)
                                    default:
                                        fmt.Println(fi8)
                                    }
                                }
                            }
                        }
                    default:
                        v3 -= v3
                    }
                case -101:
                    switch fi8 {
                    case 32767:
                        v2["\n\t"] = "\n\t"
                    case -739:
                        v22 := GenS4{F0: 32767, F1: "\n\t", F2: 32767}
                        _ = v22
                    case 2147483647:
                        var v23 uint16 = 2
                        _ = v23
                    default:
                        if false {
                            if true {
                                v3 -= (v3 & (v3 / ((v3) | 1)))
                                for fi24 := 0; fi24 < 11; fi24++ {
                                    v1 = append(v1, (fi24 - -72))
                                    fmt.Println(v3)
                                    _ = fi24
                                }
                                fmt.Println(fi8, fi8)
                            }
                            switch (fi8 << ((370) & 63)) {
                            case -1:
                                v2["foo bar"] = "\n\t"
                            }
                            var v25 int64 = 9223372036854775807
                            _ = v25
                        } else {
                            fmt.Println(fi8, fi8)
                            v3 = (v3 + 1)
                        }
                    }
                }
                v3 = 52
            }
            var v26 uint = 42
            _ = fi8
            _ = v26
        }
        _ = len(v1)
        _ = v7
    } else {
        switch ("hello" + "") {
        case "\n\t":
            var v27 uint64 = 4294967295
            _ = v27
        case "hello":
            v28 := GenS2{F0: 32767}
            _ = v28
        default:
            v3 -= v3
        }
    }
    v3 = (v3 - v3)
    v3 -= (v3 | (v3 + -127))
    var v29 uint = 823
    v3 -= 0
    if false {
        if (1000 > 32767) {
            if false {
                if (((!true) || (true && false)) && ("hello" == "\n\t")) {
                    if (false || (int16(v3) == 1)) {
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                    } else {
                        fmt.Println(v29, v3)
                        var v30 int32 = -966
                        _ = v30
                    }
                }
                for fi31 := 0; fi31 < 7; fi31++ {
                    var v32 int64 = genFn1()
                    _ = fi31
                    _ = v32
                }
                switch "hello" {
                case "":
                    var v33 uint8 = 42
                    _ = v33
                }
            }
        } else {
            v29 += (v29 / ((v29) | 1))
            var v34 int8 = v3
            v35 := GenS5{F0: -369}
            _ = v34
            _ = v35
        }
    } else {
        _ = len(v2)
        fmt.Println(v29)
        v36 := v2["\n\t"]
        _ = v36
    }
    v37 := GenS5{F0: -457}
    if (793 >= 19) {
        var v38 uint32 = 4294967295
        _ = v38
    } else {
        v3 -= -84
        for fi39 := 0; fi39 < 5; fi39++ {
            if true {
                for fi40 := 0; fi40 < 5; fi40++ {
                    _ = len(v2)
                    var v41 int8 = (v3 | (v3 / ((v3) | 1)))
                    _ = fi40
                    _ = v41
                }
            }
            for fi42 := 0; fi42 < 1; fi42++ {
                var v43 uint8 = 255
                _ = fi42
                _ = v43
            }
            _ = fi39
        }
        var v44 uint16 = 255
        _ = v44
    }
    _ = len(v1)
    fmt.Println(v29, v29)
    for fi45 := 0; fi45 < 19; fi45++ {
        var v46 uint32 = 668
        if false {
            v46 += (v46 & 292)
            v47 := GenS6{F0: true}
            for fi48 := 0; fi48 < 1; fi48++ {
                var v49 int = (fi48 / ((fi48) | 1))
                var v50 bool = (-2.25 < 0.1)
                _ = fi48
                _ = v49
                _ = v50
            }
            _ = v47
        }
        v46 = 703
        _ = fi45
        _ = v46
    }
    var v51 int = int(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v51
}
