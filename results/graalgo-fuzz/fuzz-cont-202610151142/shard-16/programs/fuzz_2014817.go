package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS5 struct { F0 string; F1 uint32; F2 uint16 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int8, p1 int16) string {
    p1 = p1
    var v4 int16 = p1
    switch 128 {
    case 32767:
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 7; fi6++ {
                v4 += (v4 + (v4 * (p1 - v4)))
                _ = fi6
            }
            _ = fi5
        }
    case 128:
        p0 = p0
    case 0:
        v7 := []int{-2147483648, -290}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return (("" + "\n\t") + "\n\t")
}
func genFn1() uint32 {
    var v8 int8 = 100
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := GenS1{F0: "a", F1: 147, F2: -9223372036854775808}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS1{F0: ("foo bar" + "\n\t"), F1: 2, F2: (^(fi12))}
        _ = fi12
        _ = v13
    }
    if (true && false) {
        var v14 int16 = -83
        if (false || ((42 < 81) || (("\n\t" + "hello") >= ""))) {
            v14 = (^(v14))
            v15 := []int{4294967295}
            switch 296 {
            case 4294967295:
                for fi16 := 0; fi16 < 15; fi16++ {
                    v8 = (v8 >> ((255) & 63))
                    _ = len(v15)
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v18 := GenS2{F0: true}
                        v19 := GenS3{F0: 42, F1: (v14 >> ((780) & 63)), F2: 128}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = fi16
                }
            case 32767:
                for fi20 := 0; fi20 < 17; fi20++ {
                    v15 = append(v15, 127)
                    switch 9223372036854775807 {
                    case 42:
                        if (18446744073709551615 <= 2) {
                            switch ("a" + "\n\t") {
                            case "":
                                for fi21 := 0; fi21 < 4; fi21++ {
                                    v22 := GenS2{F0: ((v14 / ((v14) | 1)) > v14)}
                                    _ = fi21
                                    _ = v22
                                }
                            case "hello":
                                for fi23 := 0; fi23 < 14; fi23++ {
                                    v8 = v8
                                    _ = fi23
                                }
                            case "\n\t":
                                v24 := GenS2{F0: ((fi20 >> ((0) & 63)) >= (fi20 | fi20))}
                                _ = v24
                            default:
                                var v25 int64 = 9223372036854775807
                                _ = v25
                            }
                            var v26 int32 = -1
                            _ = v26
                        } else {
                            if len(v15) > 0 {
                                v15[0] = (fi20 * fi20)
                            }
                            switch (fi20 - int(v14)) {
                            case 2147483647:
                                v14 = (v14 + (v14 * (v14 >> ((163) & 63))))
                            case -1:
                                var v27 uint16 = 712
                                _ = v27
                            default:
                                switch fi20 {
                                case 32767:
                                    v14 = v14
                                case 9223372036854775807:
                                    var v28 int64 = -1
                                    _ = v28
                                case 317:
                                    switch "foo bar" {
                                    case "a":
                                        v29 := GenS3{F0: 255, F1: (v14 | (-(v14))), F2: 30}
                                        _ = v29
                                    case "hello":
                                        v14 -= int16(fi20)
                                    default:
                                        v8 += v8
                                    }
                                default:
                                    var v30 int64 = 656
                                    _ = v30
                                }
                            }
                            if true {
                                var v31 float32 = -2.25
                                switch (fi20 * fi20) {
                                case 127:
                                    var v32 uint16 = 65535
                                    _ = v32
                                case 1:
                                    var v33 bool = (!true)
                                    _ = v33
                                }
                                _ = v31
                            } else {
                                if (false || true) {
                                    var v34 int = fi20
                                    _ = v34
                                } else {
                                    _ = len(v15)
                                }
                            }
                        }
                    case 9223372036854775807:
                        var v35 string = ("hello" + "foo bar")
                        _ = v35
                    default:
                        if false {
                            var v36 bool = (0.0 >= (((3.14159 * 0.1) / (0.1 / 0.1)) - 3.14159))
                            _ = v36
                        }
                    }
                    _ = fi20
                }
            default:
                v8 = v8
            }
            _ = v15
        }
        _ = v14
    } else {
        fmt.Println(v8)
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        var v38 bool = (32767 == 255)
        v38 = ((v8 ^ -96) != 127)
        fmt.Println(v38, fi37)
        _ = fi37
        _ = v38
    }
    _ = v8
    return 409
}

func main() {
    var v1 uint = 863
    v2 := []uint32{uint32(v1), 615}
    v3 := []string{""}
    for fi39 := 0; fi39 < 12; fi39++ {
        v1 += 124
        v40 := GenS3{F0: 255, F1: -208, F2: 230}
        if true {
            fmt.Println(fi39)
            var v41 int = 0
            _ = v41
        } else {
            var v42 int16 = 0
            var v43 string = "hello"
            _ = v42
            _ = v43
        }
        _ = fi39
        _ = v40
    }
    switch 146 {
    case 42:
        switch "" {
        case "a":
            v44 := GenS1{F0: "foo bar", F1: 2, F2: 999}
            _ = v44
        case "hello":
            var v45 int64 = 9223372036854775807
            _ = v45
        }
    case -975:
        v1 -= v1
    case 2147483647:
        var v46 uint32 = 876
        _ = v46
    }
    for fi47 := 0; fi47 < 13; fi47++ {
        fmt.Println(fi47, fi47)
        _ = fi47
    }
    switch "a" {
    case "a":
        var v48 int32 = 42
        _ = v48
    }
    _ = len(v2)
    for fi49 := 0; fi49 < 18; fi49++ {
        v3 = append(v3, "hello")
        v50 := GenS1{F0: "", F1: 66, F2: 2147483647}
        var v51 int64 = -1
        _ = fi49
        _ = v50
        _ = v51
    }
    if (((v1 & (v1 * v1)) != v1) || (false && (!false))) {
        v1 -= (^(v1))
    } else {
        fmt.Println(v1)
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        for fi52 := 0; fi52 < 8; fi52++ {
            for fi53 := 0; fi53 < 7; fi53++ {
                v1 += (v1 % ((v1) | 1))
                if true {
                    switch (-(fi53)) {
                    case 976:
                        fmt.Println(fi52)
                    }
                } else {
                    v54 := GenS4{F0: 2, F1: 0, F2: 593}
                    _ = v54
                }
                v55 := GenS2{F0: false}
                _ = fi53
                _ = v55
            }
            var v56 int32 = -2147483648
            v57 := GenS5{F0: "foo bar", F1: 65535, F2: 0}
            _ = fi52
            _ = v56
            _ = v57
        }
    }
    v58 := GenS2{F0: (-128 >= 1)}
    fmt.Println(v1)
    _ = len(v2)
    var v59 float64 = 0.0
    switch 32767 {
    case 42:
        for fi60 := 0; fi60 < 13; fi60++ {
            for fi61 := 0; fi61 < 16; fi61++ {
                switch int(v1) {
                case -1:
                    v59 = 0.1
                    var v62 bool = (false || true)
                    _ = len(v2)
                    _ = v62
                case 0:
                    v59 = (0.1 - 0.0)
                }
                _ = fi61
            }
            _ = fi60
        }
    default:
        switch ("foo bar" + "foo bar") {
        case "":
            v63 := GenS6{F0: 1}
            _ = v63
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v58
    _ = v59
}
