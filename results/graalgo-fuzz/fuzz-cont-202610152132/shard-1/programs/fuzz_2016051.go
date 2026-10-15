package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int16; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint64; F2 string }
type GenS5 struct { F0 int64; F1 uint8; F2 bool }
type GenS6 struct { F0 int64; F1 string }
func genFn0() uint8 {
    if false {
        var v4 bool = true
        _ = v4
    } else {
        var v5 uint64 = 4294967295
        v5 -= (v5 ^ v5)
        var v6 int64 = 42
        _ = v5
        _ = v6
    }
    var v7 int8 = 0
    v8 := []int{-701}
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 string, p1 string) uint8 {
    v9 := []int{-884, 32768, -2147483648, 747}
    if false {
        p0 = p0
        var v10 uint8 = 127
        switch 154 {
        case 128:
            var v11 float64 = (float64((1.5 * 1.5)) / 1.5)
            _ = v11
        case 985:
            fmt.Println(v10, p0, p0)
        default:
            p1 = p1
        }
        _ = v10
    } else {
        v9 = append(v9, -471)
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(p0)
            var v13 float32 = 1.5
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    v1 := []int{}
    var v2 uint = 255
    var v3 int8 = 35
    for fi14 := 0; fi14 < 2; fi14++ {
        fmt.Println(fi14, v3)
        if (!(1e10 == 3.14159)) {
            v3 += 100
            for fi15 := 0; fi15 < 19; fi15++ {
                if true {
                    v16 := GenS1{F0: -128, F1: (!(float32(-2.25) > 1.5))}
                    v2 -= v2
                    var v17 uint8 = 1
                    var v18 uint32 = 2147483647
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    v3 = v3
                }
                v3 = 2
                if (!(true && ("hello" > "hello"))) {
                    v2 -= 42
                } else {
                    v19 := GenS1{F0: -128, F1: true}
                    var v20 uint = 678
                    var v21 int64 = 715
                    _ = len(v1)
                    _ = v19
                    _ = v20
                    _ = v21
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi15
            }
        } else {
            var v22 int = -895
            if (746 >= 1) {
                v23 := GenS1{F0: (v3 << ((v2) & 63)), F1: true}
                for fi24 := 0; fi24 < 14; fi24++ {
                    if false {
                        var v25 int = (v22 & fi24)
                        _ = v25
                    } else {
                        if true {
                            var v26 int8 = int8(v2)
                            var v27 float32 = -2.25
                            _ = v26
                            _ = v27
                        }
                        var v28 uint8 = 170
                        _ = v28
                    }
                    _ = fi24
                }
                var v29 int8 = (v3 | (v3 + 39))
                _ = v23
                _ = v29
            }
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 int64 = 42
                _ = fi30
                _ = v31
            }
            _ = v22
        }
        _ = fi14
    }
    v32 := GenS1{F0: (v3 | v3), F1: true}
    v3 -= v3
    switch ("" + ("hello" + ("hello" + "foo bar"))) {
    case "":
        var v33 bool = (false || true)
        _ = v33
    case "\n\t":
        if ((("\n\t" + "a") < "\n\t") && (0 >= 84)) {
            v3 = 100
            for fi34 := 0; fi34 < 13; fi34++ {
                switch "" {
                case "":
                    v35 := GenS1{F0: (-(v3)), F1: false}
                    _ = v35
                case "hello":
                    var v36 float64 = float64(3.14159)
                    _ = v36
                case "foo bar":
                    if false {
                        fmt.Println(v3)
                        for fi37 := 0; fi37 < 1; fi37++ {
                            v3 = -63
                            _ = fi37
                        }
                        v2 -= (v2 % ((v2) | 1))
                    }
                }
                var v38 int = (fi34 & (fi34 + fi34))
                _ = fi34
                _ = v38
            }
        } else {
            for fi39 := 0; fi39 < 10; fi39++ {
                if len(v1) > 0 {
                    v1[0] = fi39
                }
                if false {
                    var v40 uint64 = 927
                    var v41 int = fi39
                    for fi42 := 0; fi42 < 19; fi42++ {
                        var v43 float64 = float64(1e10)
                        _ = fi42
                        _ = v43
                    }
                    _ = v40
                    _ = v41
                }
                _ = fi39
            }
            v3 += (v3 | v3)
            if len(v1) > 0 {
                v1[0] = 9223372036854775807
            }
        }
    case "foo bar":
        var v44 bool = (-2147483648 < 4294967295)
        _ = v44
    }
    if ((((0.1 - (1.5 - 3.14159)) - 1e10) / 1.0) < 1.5) {
        if true {
            var v45 int8 = v3
            switch "foo bar" {
            case "a":
                switch "" {
                case "\n\t":
                    var v46 int8 = (v3 % (((v3 & 127)) | 1))
                    _ = v46
                case "hello":
                    var v47 float64 = 3.14159
                    _ = v47
                default:
                    var v48 int32 = 0
                    _ = v48
                }
            case "\n\t":
                v45 = (-(v3))
            case "foo bar":
                switch ("hello" + "hello") {
                case "a":
                    v3 = (^(v3))
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "":
                    if true {
                        for fi49 := 0; fi49 < 6; fi49++ {
                            switch "" {
                            case "hello":
                                v2 += 42
                            default:
                                for fi50 := 0; fi50 < 2; fi50++ {
                                    if (!(6 >= 80)) {
                                        if len(v1) > 0 {
                                            v1[0] = (^(fi49))
                                        }
                                        switch ((("\n\t" + "\n\t") + "") + "") {
                                        case "\n\t":
                                            _ = len(v1)
                                        default:
                                            if (207 <= int64(fi49)) {
                                                var v51 int64 = 42
                                                v52 := GenS1{F0: (v3 + (v3 % (((v3 - 42)) | 1))), F1: false}
                                                v53 := GenS1{F0: (v3 | (v3 ^ (v45 & v3))), F1: false}
                                                var v54 int = fi50
                                                v2 = (v2 / (((v2 % ((v2) | 1))) | 1))
                                                v55 := GenS2{F0: -32768, F1: v45, F2: 0}
                                                _ = v51
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                        }
                                        if (int(v2) == (fi49 >> ((4294967295) & 63))) {
                                            v56 := GenS1{F0: (v3 >> ((v2) & 63)), F1: (((^(v2)) >= v2) || (42 == 2147483648))}
                                            _ = v56
                                        } else {
                                            v3 += v3
                                            var v57 float64 = 0.0
                                            v1 = append(v1, (^(fi49)))
                                            _ = v57
                                        }
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi50
                                }
                            }
                            v58 := GenS3{F0: 65535}
                            v45 -= -1
                            _ = fi49
                            _ = v58
                        }
                        v3 += v3
                    }
                default:
                    v3 = (v3 / (((v3 - v3)) | 1))
                }
            default:
                fmt.Println(v3)
            }
            _ = v45
        }
        switch 903 {
        case 128:
            v59 := GenS2{F0: 1000, F1: 42, F2: 1}
            _ = v59
        default:
            switch 443 {
            case 127:
                v60 := GenS1{F0: (v3 >> ((v2) & 63)), F1: true}
                _ = v60
            case 735:
                var v61 uint8 = 128
                _ = v61
            default:
                fmt.Println(v2, v2)
            }
        }
    } else {
        v62 := GenS4{F0: 337, F1: 4294967295, F2: "\n\t"}
        var v63 uint = (^(v2))
        _ = v62
        _ = v63
    }
    for fi64 := 0; fi64 < 5; fi64++ {
        if (!((42 <= -298) || ((v3 << ((255) & 63)) < v3))) {
            fmt.Println(v2)
            fmt.Println(fi64, v2, v2)
        }
        v65 := GenS5{F0: -140, F1: 185, F2: (1 >= 2)}
        var v66 string = (("foo bar" + "\n\t") + "foo bar")
        _ = fi64
        _ = v65
        _ = v66
    }
    for fi67 := 0; fi67 < 19; fi67++ {
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
        _ = fi67
    }
    for fi68 := 0; fi68 < 4; fi68++ {
        v69 := GenS4{F0: 1, F1: 60, F2: "hello"}
        _ = fi68
        _ = v69
    }
    v70 := GenS1{F0: -1, F1: false}
    var v71 int64 = 271
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v72 := GenS6{F0: (v71 / ((v71) | 1)), F1: "hello"}
    if false {
        var v73 uint16 = 0
        _ = v73
    } else {
        if false {
            var v74 int = 32767
            fmt.Println(v3)
            _ = v74
        } else {
            switch "a" {
            case "hello":
                var v75 int16 = 32767
                _ = v75
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    }
    var v76 int8 = -128
    switch (("hello" + ("" + "\n\t")) + "hello") {
    case "hello":
        v76 = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v70
    _ = v71
    _ = v72
    _ = v76
}
