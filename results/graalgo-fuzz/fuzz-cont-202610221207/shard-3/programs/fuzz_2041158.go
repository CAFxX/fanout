package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int }
func genFn0(p0 uint32) string {
    var v5 int8 = 0
    v6 := []int{32767, -461}
    _ = p0
    _ = v5
    _ = v6
    return "a"
}
func genFn1(p0 int32) int8 {
    v7 := GenS1{F0: 80, F1: 39}
    fmt.Println(p0)
    var v8 int8 = 0
    v8 = 0
    for fi9 := 0; fi9 < 20; fi9++ {
        if (595 == 18446744073709551615) {
            switch "" {
            case "a":
                var v10 uint16 = 32767
                _ = v10
            default:
                v11 := GenS2{F0: 1, F1: 2}
                _ = v11
            }
            v12 := []int{(fi9 / ((343) | 1)), int(p0), fi9}
            var v13 int32 = 1
            _ = v12
            _ = v13
        } else {
            v14 := GenS3{F0: 821}
            fmt.Println(v8, fi9)
            _ = v14
        }
        switch 0 {
        case 814:
            v15 := GenS2{F0: 2, F1: 1}
            _ = v15
        case -1:
            for fi16 := 0; fi16 < 3; fi16++ {
                v8 += (v8 & (v8 ^ (v8 - (v8 & v8))))
                switch (fi16 / ((fi16) | 1)) {
                case 0:
                    v17 := []int{0, (^(fi9)), fi16, 4294967295}
                    _ = v17
                case 9223372036854775807:
                    v18 := GenS2{F0: 2, F1: 282}
                    _ = v18
                case -2147483648:
                    v19 := []int{int(v8), (fi16 / (((fi16 % ((fi9) | 1))) | 1)), 912, -9223372036854775808}
                    _ = v19
                default:
                    if true {
                        v20 := GenS3{F0: 18446744073709551615}
                        if (-1 < -1) {
                            v21 := []int{fi16, (fi16 / ((2147483647) | 1))}
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            v8 = v8
                            _ = v21
                        }
                        _ = v20
                    }
                }
                fmt.Println(v8)
                _ = fi16
            }
        case 32767:
            v22 := []int{1}
            _ = v22
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := []int{-568, 4294967295, -315, -2147483648}
    v2 := []string{((((("" + "foo bar") + "foo bar") + "a") + "") + "\n\t"), "\n\t"}
    var v3 bool = (955 >= 1)
    var v4 uint8 = 255
    fmt.Println(v4, v3)
    var v23 float64 = (1.5 * 1e10)
    switch "\n\t" {
    case "\n\t":
        var v24 uint = 65535
        _ = v24
    }
    fmt.Println(v3, v3)
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := GenS3{F0: 42}
        var v27 int32 = -689
        _ = fi25
        _ = v26
        _ = v27
    }
    if (v3 && (18446744073709551615 != 757)) {
        v28 := GenS2{F0: 1, F1: 4294967295}
        for fi29 := 0; fi29 < 9; fi29++ {
            v3 = false
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 int64 = 1
                v23 += 0.1
                _ = fi30
                _ = v31
            }
            v4 = (v4 + 2)
            _ = fi29
        }
        switch "hello" {
        case "hello":
            var v32 float32 = 1e10
            _ = v32
        case "\n\t":
            var v33 string = ("hello" + "hello")
            _ = v33
        case "":
            v23 -= v23
        }
        _ = v28
    } else {
        for fi34 := 0; fi34 < 1; fi34++ {
            switch (fi34 | 9223372036854775807) {
            case 38:
                for fi35 := 0; fi35 < 12; fi35++ {
                    v36 := GenS4{F0: 255}
                    v37 := GenS1{F0: genFn1(-899), F1: (^(v4))}
                    switch "a" {
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        if (65535 == 0) {
                            fmt.Println(fi34)
                            v38 := GenS5{F0: 42}
                            _ = v38
                        }
                    }
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        fmt.Println(v4, v3, v23)
                        var v39 uint16 = 65535
                        v40 := GenS6{F0: -1}
                        _ = v39
                        _ = v40
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            case 589:
                v4 = v4
            case 298:
                fmt.Println(v23, fi34, fi34)
            default:
                v41 := GenS7{F0: 206}
                _ = v41
            }
            _ = fi34
        }
    }
    fmt.Println(v23, v23, v4)
    v4 += (v4 | 137)
    fmt.Println(v23, v4)
    v42 := GenS5{F0: 127}
    v3 = (42 >= 1)
    if (193 == 2) {
        var v43 float32 = (0.0 * 1.5)
        v1 = append(v1, 9223372036854775807)
        _ = v43
    } else {
        if (genFn1(-2147483648) <= genFn1(0)) {
            var v44 uint8 = (v4 << ((646) & 63))
            if true {
                var v45 uint64 = 130
                fmt.Println(v44, v45, v3)
                _ = v45
            } else {
                var v46 int32 = -334
                _ = v46
            }
            _ = v44
        } else {
            for fi47 := 0; fi47 < 12; fi47++ {
                for fi48 := 0; fi48 < 3; fi48++ {
                    if (-2.25 <= 3.14159) {
                        switch fi47 {
                        case -118:
                            if v3 {
                                v4 -= v4
                                v3 = v3
                            } else {
                                v4 -= 2
                                var v49 uint8 = v4
                                if true {
                                    v4 = (v49 >> ((65535) & 63))
                                    var v50 uint8 = 0
                                    fmt.Println(v50)
                                    _ = v50
                                } else {
                                    v51 := GenS6{F0: fi47}
                                    var v52 uint8 = (^(v4))
                                    if len(v1) > 0 {
                                        v1[0] = 210
                                    }
                                    v52 -= v4
                                    _ = v51
                                    _ = v52
                                }
                                _ = v49
                            }
                        default:
                            v1 = append(v1, (fi47 % ((fi47) | 1)))
                        }
                    } else {
                        _ = len(v2)
                        if false {
                            switch fi48 {
                            case -1:
                                if false {
                                    v53 := GenS8{F0: 42, F1: v4}
                                    if v3 {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    _ = v53
                                }
                            }
                            v4 = (v4 >> ((1) & 63))
                            if true {
                                if (v3 && v3) {
                                    if (!(74 <= -106)) {
                                        v54 := GenS9{F0: 1}
                                        if true {
                                            var v55 bool = v3
                                            v56 := GenS10{F0: 0}
                                            _ = v55
                                            _ = v56
                                        } else {
                                            var v57 float64 = -2.25
                                            _ = v57
                                        }
                                        _ = v54
                                    }
                                } else {
                                    var v58 uint32 = 4294967295
                                    var v59 uint = 42
                                    _ = v58
                                    _ = v59
                                }
                            }
                        }
                        if v3 {
                            v3 = (v3 && ((((("\n\t" + ("hello" + "")) + "a") + "") <= "foo bar") || (((42 != 575) && (v3 || (3.14159 <= 0.1))) || (v23 < (0.0 / (v23 * -2.25))))))
                        }
                    }
                    _ = fi48
                }
                fmt.Println(v4, fi47)
                _ = fi47
            }
        }
        var v60 uint8 = v4
        v23 -= v23
        _ = v60
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    for fi61 := 0; fi61 < 15; fi61++ {
        v62 := GenS11{F0: 128}
        v4 = 2
        _ = fi61
        _ = v62
    }
    switch 128 {
    case -1:
        v23 -= float64(3.14159)
    case 128:
        if len(v1) > 0 {
            v1[0] = 32768
        }
    default:
        var v63 uint8 = 1
        _ = v63
    }
    var v64 uint8 = v4
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v42
    _ = v64
}
