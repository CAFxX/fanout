package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint64; F1 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 uint }

func main() {
    var v1 int32 = 42
    v2 := make(map[string]int)
    var v3 uint8 = 128
    var v4 uint8 = (^(v3))
    var v5 int = 2147483647
    var v6 float32 = ((0.0 + 1.5) / 0.1)
    v1 = v1
    for fi7 := 0; fi7 < 16; fi7++ {
        switch "foo bar" {
        case "\n\t":
            v8 := GenS1{F0: true}
            _ = v8
        case "foo bar":
            v2[("" + "\n\t")] = fi7
        }
        fmt.Println(v4, v5)
        _ = fi7
    }
    switch (("a" + "hello") + "a") {
    case "\n\t":
        var v9 int32 = v1
        _ = v9
    }
    if true {
        switch (("a" + "a") + "\n\t") {
        case "":
            v10 := []int{(v5 ^ -633), (^(v5))}
            _ = v10
        case "foo bar":
            if false {
                delete(v2, ((("\n\t" + ("" + "")) + "foo bar") + (((("\n\t" + ("a" + ("hello" + ""))) + "") + (("a" + "\n\t") + "")) + "")))
            } else {
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 uint64 = 159
                    _ = fi11
                    _ = v12
                }
                if (-358 >= 32768) {
                    delete(v2, ("\n\t" + "a"))
                    v5 = v5
                }
            }
        default:
            var v13 int64 = 42
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := []int{v5, -980, fi14, fi14}
            switch ("" + (("" + "a") + (("foo bar" + "") + "foo bar"))) {
            case "\n\t":
                v6 -= 1.5
            case "a":
                v5 = v5
            }
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS1{F0: (v5 == (v5 - (^(v5))))}
    var v17 int16 = 769
    switch 940 {
    case -404:
        v6 += float32(1e10)
    case 1:
        if (false || (!true)) {
            var v18 uint32 = 2
            if true {
                v19 := GenS1{F0: true}
                _ = v19
            }
            _ = v18
        } else {
            var v20 uint16 = uint16(v3)
            if (true || (!(true || false))) {
                v21 := GenS1{F0: (!true)}
                if true {
                    fmt.Println(v3, v17, v4)
                    v17 = (v17 / (((v17 * (v17 & v17))) | 1))
                } else {
                    v22 := GenS1{F0: (!(false || false))}
                    var v23 int16 = -448
                    v24 := []int{int(v23)}
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = v21
            } else {
                var v25 int = v5
                for fi26 := 0; fi26 < 10; fi26++ {
                    v27 := GenS2{F0: v1, F1: 4294967295, F2: (v1 & (v1 >> ((1) & 63)))}
                    if true {
                        if true {
                            if (true && false) {
                                fmt.Println(v6)
                                v1 = (v1 | -569)
                                if true {
                                    var v28 uint8 = 108
                                    v4 -= (v3 / ((v3) | 1))
                                    _ = v28
                                } else {
                                    var v29 string = "a"
                                    _ = v29
                                }
                            } else {
                                var v30 string = "hello"
                                v31 := []int{(v25 % ((int(v1)) | 1))}
                                _ = v30
                                _ = v31
                            }
                        } else {
                            v5 = (fi26 - 9223372036854775807)
                        }
                        _ = len(v2)
                    }
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
            _ = v20
        }
    case 726:
        if (18446744073709551615 != 42) {
            v32 := GenS3{F0: 760, F1: 42}
            _ = v32
        }
    }
    switch 505 {
    case 128:
        var v33 int8 = int8(v5)
        _ = v33
    }
    v3 += (v4 * 53)
    var v34 float32 = (0.0 - v6)
    if false {
        for fi35 := 0; fi35 < 9; fi35++ {
            switch "foo bar" {
            case "\n\t":
                var v36 uint = uint(v4)
                _ = v36
            default:
                switch "" {
                case "":
                    fmt.Println(v1)
                case "foo bar":
                    v17 = 42
                }
            }
            v37 := GenS4{F0: 0, F1: (!false)}
            _ = fi35
            _ = v37
        }
    } else {
        v38 := GenS1{F0: ((1.5 != 0.1) || (!false))}
        if false {
            for fi39 := 0; fi39 < 19; fi39++ {
                for fi40 := 0; fi40 < 7; fi40++ {
                    v41 := GenS5{F0: 972, F1: 583}
                    v42 := GenS2{F0: -1, F1: 65535, F2: v1}
                    v43 := []int{}
                    _ = fi40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                var v44 int8 = -1
                v45 := []int{-362, (fi39 << ((1) & 63))}
                _ = fi39
                _ = v44
                _ = v45
            }
        } else {
            if (false && false) {
                v4 += v3
                if (("\n\t" + "") != "\n\t") {
                    var v46 float64 = 0.0
                    var v47 float64 = -2.25
                    _ = v46
                    _ = v47
                }
            } else {
                var v48 int = -2147483648
                _ = v48
            }
            v49 := GenS6{F0: (65535 >= 676)}
            _ = v49
        }
        _ = v38
    }
    for fi50 := 0; fi50 < 9; fi50++ {
        if ((v5 & (v5 % ((128) | 1))) >= v5) {
            if false {
                v2["\n\t"] = (^(v5))
                v51 := GenS6{F0: ((((float64((v6 / v6)) - 1.5) / 1.0) + 0.1) < (0.1 * -2.25))}
                _ = v51
            } else {
                v52 := GenS1{F0: (-215 == (v1 * -1))}
                _ = v52
            }
        }
        _ = fi50
    }
    v53 := GenS7{F0: 2, F1: 18446744073709551615}
    v54 := []int{v5}
    if ("foo bar" > "hello") {
        for fi55 := 0; fi55 < 11; fi55++ {
            var v56 int32 = -2147483648
            _ = fi55
            _ = v56
        }
        for fi57 := 0; fi57 < 2; fi57++ {
            switch "foo bar" {
            case "foo bar":
                v2["hello"] = 128
            case "":
                v58 := GenS3{F0: 521, F1: 375}
                _ = v58
            }
            _ = fi57
        }
    }
    for fi59 := 0; fi59 < 18; fi59++ {
        v2["hello"] = 1
        _ = fi59
    }
    if (!false) {
        if ("" <= (("hello" + (("" + "hello") + "foo bar")) + "a")) {
            v3 -= v4
        } else {
            v1 = (v1 ^ (v1 % ((v1) | 1)))
            fmt.Println(v4)
            v60 := GenS6{F0: true}
            _ = v60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
