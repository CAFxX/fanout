package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 string; F1 uint; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 uint; F2 int }
type GenS7 struct { F0 uint16; F1 string }

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]int)
    var v3 int64 = 42
    v4 := []string{"foo bar", ""}
    var v5 uint64 = 2
    v5 -= (v5 % ((4294967295) | 1))
    var v6 uint32 = 0
    if false {
        if true {
            delete(v2, ("foo bar" + ("foo bar" + ("foo bar" + "a"))))
            for fi7 := 0; fi7 < 2; fi7++ {
                v8 := v2["hello"]
                var v9 int32 = -724
                _ = fi7
                _ = v8
                _ = v9
            }
        } else {
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2["foo bar"] = 9223372036854775807
                if true {
                    if true {
                        var v10 uint64 = v5
                        switch 2147483647 {
                        case 568:
                            if len(v4) > 0 {
                                v4[0] = ("hello" + "a")
                            }
                        }
                        _ = v10
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    for fi11 := 0; fi11 < 11; fi11++ {
                        if true {
                            v3 = v3
                        }
                        v1 += (1.5 + v1)
                        switch ("foo bar" + "") {
                        case "a":
                            fmt.Println(fi11, v1)
                        case "\n\t":
                            if false {
                                var v12 int16 = 0
                                v1 = float64(1e10)
                                fmt.Println(v12, v12, v3)
                                _ = v12
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v13 bool = true
                                v5 += (^(v5))
                                _ = v13
                            }
                        default:
                            var v14 int32 = -2147483648
                            _ = v14
                        }
                        _ = fi11
                    }
                    fmt.Println(v5, v5)
                    if len(v4) > 0 {
                        v4[0] = ((((((("a" + ("foo bar" + ("\n\t" + "\n\t"))) + ("\n\t" + ("" + "a"))) + "foo bar") + ("\n\t" + "foo bar")) + "foo bar") + "foo bar") + "hello")
                    }
                }
            } else {
                var v15 int32 = 705
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 float32 = -2.25
                    for fi18 := 0; fi18 < 12; fi18++ {
                        if false {
                            fmt.Println(v5, v17, v17)
                        }
                        _ = fi18
                    }
                    _ = fi16
                    _ = v17
                }
                for fi19 := 0; fi19 < 11; fi19++ {
                    v15 -= int32(v5)
                    switch (fi19 << ((1) & 63)) {
                    case -490:
                        var v20 uint8 = 255
                        _ = v20
                    default:
                        if (true && true) {
                            if true {
                                fmt.Println(v1)
                                switch "" {
                                case "a":
                                    v5 = 95
                                }
                                v21 := GenS1{F0: (v5 - (v5 << ((18446744073709551615) & 63))), F1: -2147483648, F2: 128}
                                _ = v21
                            }
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v22 := GenS2{F0: 1, F1: 378}
                            _ = v22
                        }
                    }
                    if true {
                        for fi23 := 0; fi23 < 13; fi23++ {
                            switch "" {
                            case "hello":
                                v24 := GenS3{F0: "hello", F1: 0, F2: 128}
                                _ = v24
                            case "\n\t":
                                fmt.Println(v6)
                            case "a":
                                if false {
                                    v3 = (v3 ^ (-(v3)))
                                    v25 := GenS2{F0: 1, F1: 0}
                                    _ = v25
                                }
                            }
                            _ = fi23
                        }
                        if (true && true) {
                            var v26 uint32 = (v6 * 2147483647)
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 14; fi27++ {
                            if (42 <= 1) {
                                var v28 bool = (32767 > 32767)
                                v3 = (^(v3))
                                v3 -= (v3 | v3)
                                _ = v28
                            }
                            v29 := GenS4{F0: false, F1: 42}
                            var v30 int8 = 121
                            delete(v2, (("a" + "hello") + ""))
                            switch ("a" + "hello") {
                            case "":
                                v31 := GenS3{F0: "a", F1: 42, F2: 2}
                                _ = v31
                            case "foo bar":
                                var v32 uint64 = (v5 - (v5 % ((2) | 1)))
                                _ = v32
                            case "hello":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            _ = fi27
                            _ = v29
                            _ = v30
                        }
                    }
                    _ = fi19
                }
                _ = v15
            }
            if false {
                for fi33 := 0; fi33 < 10; fi33++ {
                    var v34 int32 = 996
                    v6 = 721
                    _ = fi33
                    _ = v34
                }
            } else {
                v35 := GenS3{F0: ("" + "hello"), F1: 458, F2: 128}
                var v36 uint8 = 0
                _ = v35
                _ = v36
            }
        }
        fmt.Println(v3, v1)
    }
    switch "a" {
    case "hello":
        if true {
            fmt.Println(v1)
            var v37 string = "\n\t"
            _ = v37
        } else {
            var v38 int8 = 100
            v4 = append(v4, ("hello" + "a"))
            var v39 int32 = int32(v38)
            _ = v38
            _ = v39
        }
    case "a":
        for fi40 := 0; fi40 < 14; fi40++ {
            v5 = v5
            delete(v2, "\n\t")
            _ = fi40
        }
    }
    var v41 int64 = v3
    var v42 int16 = -1
    fmt.Println(v42, v41)
    var v43 int32 = 0
    if (v41 != v41) {
        if false {
            var v44 int32 = (v43 * (v43 | int32(v3)))
            v45 := GenS5{F0: 255}
            _ = v44
            _ = v45
        }
        switch "foo bar" {
        case "\n\t":
            switch ("\n\t" + ("\n\t" + "hello")) {
            case "hello":
                var v46 int8 = int8(v5)
                _ = v46
            case "foo bar":
                v5 += (v5 % ((v5) | 1))
            }
        }
    } else {
        if true {
            if (-128 <= 100) {
                var v47 int64 = v3
                v48 := GenS6{F0: 127, F1: 65535, F2: 552}
                v47 = (v41 | (v41 | v41))
                _ = v47
                _ = v48
            } else {
                if (("" + "") >= "a") {
                    fmt.Println(v42, v6, v5)
                    var v49 uint8 = 127
                    _ = v49
                } else {
                    _ = len(v4)
                }
                var v50 uint64 = (v5 - v5)
                _ = v50
            }
            var v51 int = -9223372036854775808
            _ = v51
        }
        if false {
            _ = len(v4)
            if true {
                if ((2 > 32767) && (false && (!true))) {
                    var v52 float64 = ((float64(0.1) - -2.25) / v1)
                    v42 = v42
                    _ = v52
                }
            }
        } else {
            if false {
                var v53 int8 = 42
                switch "a" {
                case "hello":
                    var v54 int16 = int16(v43)
                    _ = v54
                case "\n\t":
                    switch 9223372036854775807 {
                    case 1:
                        if true {
                            for fi55 := 0; fi55 < 6; fi55++ {
                                v53 = v53
                                v1 = (v1 + 0.0)
                                if false {
                                    if (false && true) {
                                        var v56 int8 = v53
                                        _ = v56
                                    }
                                    var v57 int8 = (v53 % (((v53 >> ((442) & 63))) | 1))
                                    _ = v57
                                }
                                _ = fi55
                            }
                        }
                    default:
                        v58 := GenS7{F0: 255, F1: "foo bar"}
                        _ = v58
                    }
                }
                _ = v53
            }
            fmt.Println(v43)
            switch -512 {
            case 1:
                v59 := GenS5{F0: 18446744073709551615}
                _ = v59
            case 532:
                fmt.Println(v43, v3, v43)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
