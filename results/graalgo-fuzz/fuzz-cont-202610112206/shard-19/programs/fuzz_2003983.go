package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint; F1 uint; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS10 struct { F0 int32 }
func genFn0(p0 int32, p1 bool) bool {
    v4 := GenS1{F0: 0, F1: p0, F2: 25}
    var v5 uint = 93
    fmt.Println(p1)
    var v6 float64 = 1.5
    p0 -= (p0 + p0)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return false
}
func genFn1() string {
    v7 := GenS2{F0: 4294967295, F1: 42}
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint64 = 0
        v10 := []int{(fi8 & fi8), 32768, -2147483648}
        var v11 float32 = (1e10 - -2.25)
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    _ = v7
    return ("a" + "\n\t")
}

func main() {
    var v1 bool = (-9223372036854775808 <= 531)
    var v2 int16 = 42
    var v3 uint16 = 228
    v1 = (((!(81 != 1)) && v1) && (!(!(109 == 1))))
    for fi12 := 0; fi12 < 9; fi12++ {
        v3 = v3
        _ = fi12
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        fmt.Println(v1, v2, v3)
    case "a":
        if v1 {
            fmt.Println(v1, v3, v2)
            v3 = (v3 / ((v3) | 1))
        }
    case "hello":
        v13 := GenS2{F0: 42, F1: (v2 / ((v2) | 1))}
        _ = v13
    default:
        if true {
            switch -2147483648 {
            case 547:
                var v14 int64 = 841
                _ = v14
            case 1:
                if (!v1) {
                    if (-346 < 697) {
                        v15 := GenS3{F0: 1, F1: 4294967295, F2: true}
                        fmt.Println(v1, v2, v2)
                        _ = v15
                    }
                }
            }
            switch int(v3) {
            case -1:
                v16 := []int{}
                _ = v16
            case -9223372036854775808:
                v17 := GenS2{F0: 42, F1: (v2 + v2)}
                _ = v17
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := []int{fi18, fi18, fi18, 224}
                v1 = ((!(0 > -2147483648)) && (!((100 <= 0) && true)))
                _ = fi18
                _ = v19
            }
        }
    }
    if true {
        var v20 uint = 2
        var v21 uint8 = 41
        v3 -= uint16(v20)
        _ = v20
        _ = v21
    }
    v3 = v3
    v1 = v1
    if true {
        v3 = (^(v3))
    }
    if (!(-9223372036854775808 < 127)) {
        var v22 int = -9223372036854775808
        var v23 bool = v1
        switch 1 {
        case 32768:
            v1 = true
        default:
            v24 := GenS3{F0: 781, F1: 42, F2: v23}
            _ = v24
        }
        _ = v22
        _ = v23
    } else {
        fmt.Println(v1, v2, v1)
    }
    if ("foo bar" > "") {
        for fi25 := 0; fi25 < 15; fi25++ {
            fmt.Println(v1, fi25)
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(v3, v2, fi25)
                v27 := GenS3{F0: 4294967295, F1: 255, F2: v1}
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    } else {
        switch int(v3) {
        case -9223372036854775808:
            var v28 uint16 = v3
            _ = v28
        }
        if v1 {
            v29 := []int{}
            for fi30 := 0; fi30 < 17; fi30++ {
                if v1 {
                    switch fi30 {
                    case 9223372036854775807:
                        v3 = 200
                    }
                } else {
                    var v31 float64 = ((-2.25 / 1.5) - ((1e10 / 1.0) + 1.5))
                    if v1 {
                        var v32 uint = 600
                        if (0.1 == v31) {
                            v32 -= v32
                            v32 = (v32 & v32)
                            v31 = 0.1
                        }
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 17; fi33++ {
                            v34 := GenS2{F0: 1, F1: v2}
                            if (!(0 < 4294967295)) {
                                v35 := GenS4{F0: v3}
                                fmt.Println(v31, v2, fi30)
                                if len(v29) >= 2 {
                                    v29 = v29[1:len(v29)]
                                }
                                _ = v35
                            }
                            switch "\n\t" {
                            case "":
                                v36 := GenS4{F0: v3}
                                _ = v36
                            case "foo bar":
                                fmt.Println(fi33, v1, fi30)
                            case "\n\t":
                                v37 := GenS5{F0: "hello"}
                                _ = v37
                            default:
                                var v38 int32 = 1
                                _ = v38
                            }
                            _ = fi33
                            _ = v34
                        }
                    }
                    v29 = append(v29, 2147483647)
                    _ = v31
                }
                _ = fi30
            }
            switch "hello" {
            case "hello":
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
            case "":
                var v39 uint8 = 218
                _ = v39
            }
            _ = v29
        } else {
            for fi40 := 0; fi40 < 11; fi40++ {
                for fi41 := 0; fi41 < 13; fi41++ {
                    fmt.Println(v1, v3)
                    if false {
                        fmt.Println(v3, v2, fi40)
                        v42 := []int{}
                        _ = v42
                    }
                    switch 127 {
                    case -9223372036854775808:
                        v43 := GenS1{F0: v3, F1: -502, F2: 28}
                        _ = v43
                    }
                    _ = fi41
                }
                for fi44 := 0; fi44 < 17; fi44++ {
                    switch "a" {
                    case "a":
                        v45 := []int{fi40, (fi44 ^ (fi40 * 938))}
                        _ = v45
                    case "hello":
                        v46 := []int{}
                        v3 -= v3
                        switch ("a" + "a") {
                        case "a":
                            v47 := GenS1{F0: v3, F1: 42, F2: 255}
                            _ = v47
                        case "":
                            fmt.Println(fi40, fi40, fi44)
                        case "foo bar":
                            if (100 >= 127) {
                                if true {
                                    v48 := GenS6{F0: v1}
                                    _ = len(v46)
                                    if len(v46) >= 2 {
                                        v46 = v46[1:len(v46)]
                                    }
                                    v49 := GenS3{F0: 2, F1: 65535, F2: v1}
                                    v2 = 32767
                                    _ = v48
                                    _ = v49
                                } else {
                                    v50 := GenS7{F0: "", F1: (!true)}
                                    _ = v50
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    v1 = true
                                }
                            } else {
                                var v51 float64 = (1e10 / -2.25)
                                v2 -= v2
                                _ = v51
                            }
                        }
                        _ = v46
                    }
                    v52 := GenS8{F0: 1}
                    _ = fi44
                    _ = v52
                }
                v53 := []int{fi40, (^(fi40))}
                _ = fi40
                _ = v53
            }
        }
    }
    v1 = (((-2.25 / -2.25) - 3.14159) > float32(float64(1e10)))
    if (2147483647 != -2147483648) {
        if (!(!v1)) {
            if ((v3 % (((v3 / ((32767) | 1))) | 1)) != v3) {
                if true {
                    if true {
                        switch "a" {
                        case "a":
                            if (v1 && (true && v1)) {
                                switch -2147483648 {
                                case -1:
                                    v2 += 32767
                                case 127:
                                    var v54 int32 = -2147483648
                                    _ = v54
                                default:
                                    v55 := []int{}
                                    _ = v55
                                }
                            } else {
                                for fi56 := 0; fi56 < 10; fi56++ {
                                    switch ("hello" + "foo bar") {
                                    case "a":
                                        if ((true || v1) || true) {
                                            for fi57 := 0; fi57 < 8; fi57++ {
                                                v2 = (v2 % ((v2) | 1))
                                                var v58 int8 = 127
                                                v1 = (int(v3) > fi57)
                                                _ = fi57
                                                _ = v58
                                            }
                                            v2 += v2
                                        } else {
                                            v1 = (12 < 100)
                                            var v59 int16 = 883
                                            v60 := []int{fi56, 32768}
                                            _ = v59
                                            _ = v60
                                        }
                                    case "hello":
                                        switch fi56 {
                                        case 32767:
                                            var v61 string = ("a" + "\n\t")
                                            _ = v61
                                        case 2147483647:
                                            fmt.Println(v1, fi56, v2)
                                        case 4294967295:
                                            v1 = false
                                        }
                                    }
                                    _ = fi56
                                }
                                v62 := GenS9{F0: 2, F1: v3, F2: v3}
                                v63 := []int{9223372036854775807, 621}
                                _ = v62
                                _ = v63
                            }
                        }
                    }
                    v64 := GenS9{F0: 65535, F1: (^(v3)), F2: v3}
                    v65 := []int{}
                    _ = v64
                    _ = v65
                }
            } else {
                var v66 uint16 = v3
                _ = v66
            }
            v67 := []int{554}
            fmt.Println(v1, v1, v2)
            _ = v67
        }
    } else {
        v68 := GenS3{F0: 255, F1: 42, F2: v1}
        v69 := GenS10{F0: 42}
        for fi70 := 0; fi70 < 11; fi70++ {
            v1 = (!v1)
            v71 := []int{}
            _ = fi70
            _ = v71
        }
        _ = v68
        _ = v69
    }
    for fi72 := 0; fi72 < 20; fi72++ {
        v73 := []int{(fi72 - (fi72 << ((65535) & 63))), fi72}
        v74 := GenS3{F0: 504, F1: 18446744073709551615, F2: true}
        _ = fi72
        _ = v73
        _ = v74
    }
    fmt.Println(v2, v2, v1)
    fmt.Println(v1, v2, v1)
    v2 -= 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
