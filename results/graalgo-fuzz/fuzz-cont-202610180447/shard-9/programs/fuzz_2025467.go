package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int64 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint32) uint32 {
    v5 := GenS1{F0: "\n\t", F1: 2147483647, F2: -9223372036854775808}
    switch "a" {
    case "foo bar":
        if ((0.0 - 0.1) <= (0.1 - 1.5)) {
            var v6 uint16 = 255
            var v7 int = 1
            _ = v6
            _ = v7
        }
    }
    switch 128 {
    case -9223372036854775808:
        var v8 float64 = float64((float32(-2.25) / float32(0.1)))
        _ = v8
    case 9223372036854775807:
        p0 = (p0 >> ((4294967295) & 63))
    case -843:
        v9 := GenS1{F0: "\n\t", F1: 0, F2: 42}
        _ = v9
    default:
        v10 := []int{}
        _ = v10
    }
    var v11 uint32 = 1
    _ = p0
    _ = v11
    _ = v5
    return v11
}
func genFn1(p0 int, p1 uint16) int {
    if (true && false) {
        var v12 uint8 = 243
        _ = v12
    }
    v13 := []int{(p0 - (p0 & (p0 ^ 4294967295))), (p0 / ((-571) | 1)), p0, p0}
    var v14 uint = 42
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return p0
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]string)
    var v3 int32 = 72
    var v4 int64 = 1
    v4 += (v4 | (v4 + int64(v3)))
    switch -2147483648 {
    case 42:
        if (!true) {
            fmt.Println(v3, v3, v4)
            var v15 int = 32767
            _ = v15
        } else {
            if true {
                fmt.Println(v3, v3)
                var v16 bool = true
                var v17 int32 = (v3 / ((-636) | 1))
                _ = v16
                _ = v17
            } else {
                switch "\n\t" {
                case "hello":
                    for fi18 := 0; fi18 < 20; fi18++ {
                        var v19 int16 = 452
                        _ = fi18
                        _ = v19
                    }
                case "a":
                    if false {
                        v20 := GenS1{F0: "hello", F1: (v3 / ((42) | 1)), F2: (^(v4))}
                        var v21 int8 = 0
                        _ = v20
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 16; fi22++ {
                            var v23 int = (^(fi22))
                            _ = fi22
                            _ = v23
                        }
                    }
                default:
                    if (false || true) {
                        delete(v2, ("a" + ""))
                        fmt.Println(v1, v4, v1)
                    }
                }
                v24 := GenS1{F0: "", F1: -241, F2: v4}
                var v25 uint64 = 0
                _ = v24
                _ = v25
            }
        }
    case -2147483648:
        var v26 float64 = (((v1 - 0.1) + 3.14159) * float64((1e10 / 1.5)))
        _ = v26
    case 1:
        switch 1 {
        case 32768:
            fmt.Println(v3, v3, v1)
        case -1:
            v4 = -1
        case -915:
            var v27 int16 = 1000
            _ = v27
        default:
            for fi28 := 0; fi28 < 19; fi28++ {
                for fi29 := 0; fi29 < 14; fi29++ {
                    if false {
                        if (!(!(18446744073709551615 > 943))) {
                            v4 += v4
                            fmt.Println(v3, fi28)
                            v3 = (^(v3))
                        }
                        fmt.Println(fi29, v1, v3)
                        v30 := []int{}
                        switch ("hello" + "a") {
                        case "hello":
                            if (!(-659 <= 309)) {
                                switch 128 {
                                case -1:
                                    var v31 float32 = (1e10 / 3.14159)
                                    _ = v31
                                case -832:
                                    _ = len(v30)
                                case 4294967295:
                                    switch ("" + ("hello" + "\n\t")) {
                                    case "foo bar":
                                        var v32 int32 = v3
                                        _ = v32
                                    case "hello":
                                        var v33 float64 = v1
                                        _ = v33
                                    default:
                                        if (!false) {
                                            delete(v2, ("\n\t" + ""))
                                        } else {
                                            v1 = (v1 + (1e10 + v1))
                                        }
                                        _ = len(v2)
                                    }
                                default:
                                    if false {
                                        fmt.Println(v3, fi29)
                                    } else {
                                        var v34 int8 = 127
                                        _ = v34
                                    }
                                }
                                fmt.Println(v1)
                            }
                        case "a":
                            v35 := GenS1{F0: "hello", F1: v3, F2: 930}
                            _ = v35
                        case "foo bar":
                            v36 := GenS1{F0: "hello", F1: (v3 + 1), F2: 614}
                            _ = v36
                        default:
                            _ = len(v2)
                        }
                        _ = v30
                    }
                    v4 = 42
                    switch fi29 {
                    case 1:
                        switch (fi29 + fi28) {
                        case 195:
                            v37 := []int{genFn1(fi29, 42), fi28, 127}
                            _ = v37
                        }
                    case 12:
                        v38 := v2["hello"]
                        _ = v38
                    }
                    v4 = v4
                    _ = fi29
                }
                _ = fi28
            }
        }
    }
    if true {
        switch 2147483647 {
        case 42:
            var v39 int32 = v3
            _ = v39
        }
        switch -137 {
        case -967:
            if true {
                v4 = v4
                v4 += (v4 - -308)
                for fi40 := 0; fi40 < 16; fi40++ {
                    for fi41 := 0; fi41 < 17; fi41++ {
                        if (v3 >= -180) {
                            switch "" {
                            case "a":
                                v42 := []int{9223372036854775807, (fi40 ^ (fi40 - fi40))}
                                _ = v42
                            case "":
                                v4 = (v4 | 648)
                            }
                        }
                        _ = fi41
                    }
                    _ = fi40
                }
            } else {
                v4 -= v4
                switch 1 {
                case -9223372036854775808:
                    for fi43 := 0; fi43 < 15; fi43++ {
                        v1 = v1
                        if ((fi43 | genFn1(1, 974)) >= (fi43 << ((749) & 63))) {
                            v44 := GenS2{F0: (v1 <= 3.14159), F1: 0}
                            v1 -= v1
                            _ = v44
                        } else {
                            fmt.Println(v4, v4, v4)
                        }
                        fmt.Println(v4)
                        _ = fi43
                    }
                case 32767:
                    delete(v2, ("foo bar" + ""))
                case 42:
                    fmt.Println(v4, v3)
                }
                v45 := GenS3{F0: false, F1: 42}
                _ = v45
            }
        case 275:
            var v46 uint8 = 42
            _ = v46
        }
    } else {
        var v47 uint32 = 4294967295
        _ = v47
    }
    switch 32767 {
    case 127:
        v48 := v2["hello"]
        _ = v48
    case 4294967295:
        if false {
            for fi49 := 0; fi49 < 1; fi49++ {
                v3 = v3
                delete(v2, ("foo bar" + ("hello" + "hello")))
                _ = fi49
            }
            var v50 int32 = v3
            for fi51 := 0; fi51 < 9; fi51++ {
                for fi52 := 0; fi52 < 20; fi52++ {
                    v4 += (v4 >> ((65535) & 63))
                    _ = fi52
                }
                _ = fi51
            }
            _ = v50
        } else {
            if true {
                v53 := GenS1{F0: "\n\t", F1: (^(v3)), F2: v4}
                switch "a" {
                case "\n\t":
                    var v54 uint = 42
                    _ = v54
                case "foo bar":
                    v55 := GenS3{F0: (!(255 == 1)), F1: v3}
                    _ = v55
                case "hello":
                    var v56 uint16 = 0
                    _ = v56
                }
                _ = v53
            } else {
                if true {
                    fmt.Println(v1, v1)
                }
            }
        }
    }
    var v57 int32 = (v3 ^ (v3 >> ((42) & 63)))
    v58 := GenS2{F0: ((2 <= 1) || true), F1: 100}
    v59 := GenS4{F0: 42, F1: (false || (!(true && false)))}
    v60 := []int{}
    for fi61 := 0; fi61 < 12; fi61++ {
        switch 32768 {
        case 128:
            if true {
                if ((((("foo bar" + "foo bar") != "a") || false) && (true && (!false))) || (2147483648 >= 399)) {
                    v62 := GenS5{F0: 9223372036854775807, F1: (42 < 176)}
                    var v63 bool = (!true)
                    _ = v62
                    _ = v63
                } else {
                    v57 = v3
                    for fi64 := 0; fi64 < 3; fi64++ {
                        v4 = -288
                        _ = fi64
                    }
                }
                switch "foo bar" {
                case "":
                    fmt.Println(v57)
                case "\n\t":
                    if true {
                        var v65 int32 = (v57 - (v57 % (((v3 >> ((1) & 63))) | 1)))
                        v3 = v57
                        _ = v65
                    }
                default:
                    v66 := GenS2{F0: false, F1: -1}
                    _ = v66
                }
                var v67 float32 = (-2.25 * 1e10)
                _ = v67
            }
        case -9223372036854775808:
            for fi68 := 0; fi68 < 20; fi68++ {
                v69 := GenS6{F0: 32767}
                v70 := GenS5{F0: (-(v4)), F1: false}
                _ = len(v2)
                _ = fi68
                _ = v69
                _ = v70
            }
        case 1:
            v2[("\n\t" + ("foo bar" + ("\n\t" + "hello")))] = "hello"
        default:
            v71 := GenS5{F0: (v4 + 0), F1: (18446744073709551615 != 42)}
            _ = v71
        }
        if true {
            for fi72 := 0; fi72 < 19; fi72++ {
                var v73 float32 = ((float32(v1) * (-2.25 - (0.0 / 1.0))) + 1.5)
                v74 := GenS5{F0: -9223372036854775808, F1: true}
                v75 := v2["hello"]
                _ = fi72
                _ = v73
                _ = v74
                _ = v75
            }
        } else {
            v76 := GenS2{F0: false, F1: 1}
            if ((v4 / (((v4 / (((-(v4))) | 1))) | 1)) != v4) {
                v57 = (v57 * 0)
                fmt.Println(v4, v4)
            }
            _ = v76
        }
        var v77 bool = false
        _ = fi61
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v59
    _ = v60
}
