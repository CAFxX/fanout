package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
type GenS2 struct { F0 bool; F1 uint32; F2 bool }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 int64; F2 int }
func genFn0(p0 int64, p1 bool) int8 {
    fmt.Println(p1, p0)
    p1 = p1
    fmt.Println(p0)
    var v4 string = "foo bar"
    var v5 int = 794
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 127
}

func main() {
    var v1 uint8 = 0
    v2 := []int{-9223372036854775808, 127, 332, 1}
    v3 := []string{"\n\t", "foo bar"}
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            fmt.Println(fi6)
            v1 = v1
        } else {
            v3 = append(v3, "hello")
            var v7 int64 = -1
            for fi8 := 0; fi8 < 17; fi8++ {
                fmt.Println(v7)
                if len(v2) > 0 {
                    v2[0] = (fi8 >> ((506) & 63))
                }
                var v9 int64 = (v7 + (v7 ^ v7))
                _ = fi8
                _ = v9
            }
            _ = v7
        }
        var v10 int32 = 42
        var v11 int8 = 50
        _ = fi6
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        if (!(!true)) {
            if false {
                fmt.Println(v1, v1)
            }
            v1 = 128
            if false {
                switch -790 {
                case 0:
                    for fi13 := 0; fi13 < 14; fi13++ {
                        fmt.Println(v1, fi13)
                        switch "a" {
                        case "":
                            v1 = (v1 | v1)
                        case "foo bar":
                            v1 = v1
                        case "a":
                            v14 := GenS1{F0: 42, F1: 0, F2: -32768}
                            _ = v14
                        }
                        _ = fi13
                    }
                case 4294967295:
                    var v15 uint32 = 65535
                    _ = v15
                default:
                    if (false && ((v1 < 196) && false)) {
                        if true {
                            for fi16 := 0; fi16 < 12; fi16++ {
                                fmt.Println(v1)
                                if (((true && false) || (2 > 4294967295)) || (2 == 229)) {
                                    fmt.Println(fi16)
                                    var v17 uint = 4294967295
                                    var v18 uint = 18446744073709551615
                                    var v19 uint8 = (v1 >> ((v18) & 63))
                                    if (-32768 != 1) {
                                        var v20 int64 = 9223372036854775807
                                        _ = v20
                                    } else {
                                        var v21 int8 = -105
                                        if len(v2) > 0 {
                                            v2[0] = fi12
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v21
                                    }
                                    var v22 uint64 = 927
                                    switch "hello" {
                                    case "foo bar":
                                        if (!false) {
                                            v23 := GenS2{F0: true, F1: 0, F2: true}
                                            _ = v23
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    case "hello":
                                        v2 = append(v2, (fi16 + 587))
                                    default:
                                        if (false || false) {
                                            var v24 int32 = -2147483648
                                            _ = len(v2)
                                            _ = v24
                                        } else {
                                            if ((fi12 & -742) > 1) {
                                                var v25 bool = (-368 == 1000)
                                                if (-9223372036854775808 >= -9223372036854775808) {
                                                    v1 -= v19
                                                    switch "" {
                                                    case "hello":
                                                        v25 = v25
                                                    case "foo bar":
                                                        v26 := GenS3{F0: 42, F1: "a"}
                                                        _ = v26
                                                    case "a":
                                                        var v27 int8 = -1
                                                        _ = v27
                                                    }
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = "foo bar"
                                                    }
                                                    fmt.Println(v19, v1, v25)
                                                    switch ("\n\t" + "foo bar") {
                                                    case "\n\t":
                                                        if (255 == uint16(v17)) {
                                                            fmt.Println(v25)
                                                        }
                                                    case "hello":
                                                        switch "a" {
                                                        case "":
                                                            if ((1.5 >= float32(-2.25)) && (42 >= 9223372036854775807)) {
                                                                switch (fi12 >> ((255) & 63)) {
                                                                case 9223372036854775807:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                case -942:
                                                                    var v28 string = "\n\t"
                                                                    _ = v28
                                                                case 42:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                default:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = "foo bar"
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(fi12, v18)
                                                    if (fi16 <= 0) {
                                                        if len(v3) > 0 {
                                                            v3[0] = "a"
                                                        }
                                                        v1 -= 42
                                                        v19 += v19
                                                    } else {
                                                        v25 = (65535 > 277)
                                                        if true {
                                                            fmt.Println(fi12)
                                                        } else {
                                                            var v29 int64 = 9223372036854775807
                                                            var v30 int32 = 1
                                                            _ = v29
                                                            _ = v30
                                                        }
                                                    }
                                                }
                                                var v31 int64 = 334
                                                _ = v25
                                                _ = v31
                                            }
                                            fmt.Println(v18, v18)
                                        }
                                    }
                                    v1 = (v1 << ((v17) & 63))
                                    v32 := GenS3{F0: -1, F1: "foo bar"}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                    _ = v22
                                    _ = v32
                                }
                                _ = fi16
                            }
                            var v33 uint32 = 554
                            _ = v33
                        }
                        switch "" {
                        case "foo bar":
                            switch (fi12 - fi12) {
                            case -9223372036854775808:
                                var v34 uint64 = 206
                                _ = v34
                            case 128:
                                if ((1e10 / 0.1) < 0.1) {
                                    var v35 int16 = 1
                                    v35 = v35
                                    _ = v35
                                }
                            }
                        case "":
                            fmt.Println(fi12, v1)
                        default:
                            var v36 float64 = 3.14159
                            _ = v36
                        }
                        if ((1 > 619) || (!true)) {
                            _ = len(v3)
                        } else {
                            if (!true) {
                                if (656 == 0) {
                                    var v37 uint = 65535
                                    _ = len(v3)
                                    _ = v37
                                } else {
                                    var v38 uint8 = (v1 / (((^(v1))) | 1))
                                    _ = v38
                                }
                                v39 := GenS3{F0: -301, F1: ("a" + "a")}
                                v40 := GenS4{F0: -9223372036854775808, F1: 413}
                                _ = v39
                                _ = v40
                            }
                            for fi41 := 0; fi41 < 8; fi41++ {
                                v1 -= (v1 & (v1 / (((v1 % ((v1) | 1))) | 1)))
                                switch 128 {
                                case 32768:
                                    v3 = append(v3, "a")
                                case 4294967295:
                                    v42 := GenS3{F0: 42, F1: "hello"}
                                    _ = v42
                                }
                                v43 := GenS4{F0: -1, F1: -371}
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                _ = fi41
                                _ = v43
                            }
                            v44 := GenS4{F0: -324, F1: 42}
                            _ = v44
                        }
                    } else {
                        if ((!false) || ((361 >= -1) || ("\n\t" <= (("foo bar" + "foo bar") + "foo bar")))) {
                            v1 = (v1 - v1)
                        }
                    }
                }
            }
        }
        for fi45 := 0; fi45 < 4; fi45++ {
            v1 += (^(v1))
            var v46 int8 = genFn0(0, true)
            _ = fi45
            _ = v46
        }
        _ = fi12
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    v47 := GenS5{F0: 65535}
    var v48 uint32 = 4294967295
    v49 := GenS3{F0: -2147483648, F1: "foo bar"}
    v50 := GenS5{F0: v48}
    var v51 bool = ((false && false) && true)
    var v52 int64 = -63
    v53 := GenS6{F0: -128, F1: (^(v52)), F2: 2147483647}
    v54 := GenS5{F0: v48}
    if ((((float32(float64(1e10)) + (-2.25 * -2.25)) - (((float32(-2.25) + (1.5 * 0.1)) * -2.25) / 1.5)) - -2.25) >= 0.0) {
        var v55 int8 = genFn0(v52, v51)
        _ = v55
    } else {
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
