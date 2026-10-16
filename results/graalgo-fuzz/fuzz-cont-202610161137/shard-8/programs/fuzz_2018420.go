package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int; F1 int; F2 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int32; F2 uint32 }
type GenS5 struct { F0 int; F1 int8; F2 int32 }
type GenS6 struct { F0 string; F1 uint; F2 uint }
func genFn0(p0 uint32) int8 {
    if false {
        v7 := []int{128, -277}
        p0 += 0
        p0 = p0
        _ = v7
    }
    p0 -= p0
    for fi8 := 0; fi8 < 1; fi8++ {
        switch "foo bar" {
        case "hello":
            v9 := []int{}
            _ = v9
        case "":
            fmt.Println(fi8)
        case "a":
            for fi10 := 0; fi10 < 10; fi10++ {
                fmt.Println(fi8, p0, fi10)
                _ = fi10
            }
        default:
            p0 += p0
        }
        if true {
            if false {
                if (4294967295 >= 259) {
                    var v11 int32 = -2147483648
                    for fi12 := 0; fi12 < 19; fi12++ {
                        var v13 uint8 = 0
                        fmt.Println(p0, fi12)
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                } else {
                    var v14 uint16 = 65535
                    v15 := []int{}
                    var v16 uint = 4294967295
                    _ = v14
                    _ = v15
                    _ = v16
                }
                p0 = (p0 << ((4294967295) & 63))
            } else {
                p0 = p0
            }
            p0 -= p0
        } else {
            for fi17 := 0; fi17 < 3; fi17++ {
                p0 -= p0
                var v18 int = 1
                _ = fi17
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := GenS1{F0: 249, F1: 42}
            p0 -= 42
            _ = fi19
            _ = v20
        }
        _ = fi8
    }
    var v21 uint = 436
    _ = p0
    _ = v21
    return 33
}
func genFn1(p0 int64, p1 uint32) uint8 {
    var v22 uint8 = 255
    v23 := []int{18, -66}
    var v24 int32 = -1
    var v25 bool = true
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    return (v22 % (((v22 - v22)) | 1))
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 49
    v3 := make(map[string]string)
    v4 := []string{"foo bar", ((("hello" + "\n\t") + "a") + (("a" + "hello") + "\n\t"))}
    v5 := make(map[string]uint32)
    var v6 int8 = 42
    v26 := GenS2{F0: 2147483647, F1: int(v6), F2: 578}
    _ = len(v1)
    v27 := GenS3{F0: 102}
    v28 := GenS3{F0: -1}
    var v29 bool = (!false)
    if v29 {
        v30 := GenS4{F0: -9223372036854775808, F1: -2147483648, F2: 2147483648}
        _ = v30
    }
    v6 += v6
    if (!((-1 != -683) || false)) {
        fmt.Println(v2, v6, v2)
        var v31 float32 = float32((1e10 - 1e10))
        _ = v31
    }
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 uint64 = 0
        var v34 uint64 = 537
        _ = fi32
        _ = v33
        _ = v34
    }
    switch "\n\t" {
    case "foo bar":
        v6 = (v6 << (((v2 % ((0) | 1))) & 63))
    case "":
        if v29 {
            v4 = append(v4, "foo bar")
            if len(v4) > 0 {
                v4[0] = "hello"
            }
        }
    case "a":
        if v29 {
            v6 = 1
            if true {
                var v35 int32 = 1
                if v29 {
                    var v36 uint16 = 74
                    v6 -= genFn0(193)
                    v37 := GenS1{F0: 2, F1: 2}
                    _ = v36
                    _ = v37
                } else {
                    var v38 string = "a"
                    switch ("hello" + ((v38 + "hello") + ("" + v38))) {
                    case "a":
                        v38 = v38
                    case "":
                        v39 := GenS1{F0: 210, F1: 177}
                        _ = v39
                    case "foo bar":
                        if false {
                            fmt.Println(v29, v6)
                            if (65535 != 1) {
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                            }
                            if len(v4) > 0 {
                                v4[0] = "a"
                            }
                        }
                    default:
                        var v40 int = 9223372036854775807
                        _ = v40
                    }
                    _ = v38
                }
                var v41 bool = (v29 && true)
                _ = v35
                _ = v41
            } else {
                var v42 int64 = 0
                v29 = true
                _ = v42
            }
            for fi43 := 0; fi43 < 12; fi43++ {
                v3["foo bar"] = ("foo bar" + (("foo bar" + "hello") + "hello"))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 = v2
                _ = fi43
            }
        } else {
            switch 0 {
            case 127:
                for fi44 := 0; fi44 < 5; fi44++ {
                    var v45 uint = (^(v2))
                    if (2 == 364) {
                        var v46 int = (fi44 % ((fi44) | 1))
                        var v47 bool = (!v29)
                        _ = v46
                        _ = v47
                    } else {
                        if v29 {
                            v29 = (v29 || (!((v6 | (^(v6))) >= v6)))
                        } else {
                            if true {
                                var v48 float32 = 3.14159
                                _ = v48
                            } else {
                                if (65535 < 42) {
                                    v29 = v29
                                    v29 = (!v29)
                                }
                                v49 := GenS2{F0: fi44, F1: fi44, F2: 0}
                                v45 -= 65535
                                _ = v49
                            }
                            for fi50 := 0; fi50 < 4; fi50++ {
                                var v51 uint = (v45 * (v45 * (v2 >> (((v45 / (((^(v2))) | 1))) & 63))))
                                v52 := v1[(^(fi50))]
                                if ((!(42 <= 4294967295)) || (true && false)) {
                                    var v53 string = ""
                                    switch int(v6) {
                                    case 2147483647:
                                        if v29 {
                                            if ((2 != genFn1(-1, 2147483647)) && (!(v29 && v52))) {
                                                if true {
                                                    if (v53 <= v53) {
                                                        fmt.Println(v2, v52)
                                                    }
                                                    var v54 int8 = v6
                                                    var v55 uint16 = 255
                                                    var v56 uint32 = 256
                                                    _ = v54
                                                    _ = v55
                                                    _ = v56
                                                }
                                                var v57 float32 = (0.0 - ((1.5 / 0.1) * 0.1))
                                                v6 -= (v6 / ((v6) | 1))
                                                _ = v57
                                            } else {
                                                switch (^(fi44)) {
                                                case -786:
                                                    switch v53 {
                                                    case "hello":
                                                        var v58 uint8 = 24
                                                        _ = v58
                                                    case "a":
                                                        if v29 {
                                                            fmt.Println(v2, v45)
                                                            delete(v3, v53)
                                                        } else {
                                                            fmt.Println(fi44, v51)
                                                            fmt.Println(fi50, v29)
                                                            if len(v4) > 0 {
                                                                v4[0] = v53
                                                            }
                                                        }
                                                    case "":
                                                        v59 := GenS2{F0: (fi44 % ((819) | 1)), F1: fi44, F2: (v2 % ((18446744073709551615) | 1))}
                                                        _ = v59
                                                    }
                                                    var v60 uint32 = 402
                                                    _ = v60
                                                case 128:
                                                    fmt.Println(v6, v6, v53)
                                                default:
                                                    v2 = 120
                                                }
                                            }
                                        } else {
                                            if (-9223372036854775808 > 4294967295) {
                                                v61 := GenS5{F0: 4294967295, F1: (v6 - 77), F2: -1}
                                                _ = v61
                                            }
                                            _ = len(v4)
                                            if false {
                                                v5[v53] = 679
                                                if (!(208 <= 0)) {
                                                    fmt.Println(v6, v45)
                                                    if len(v4) > 0 {
                                                        v4[0] = ("hello" + "\n\t")
                                                    }
                                                }
                                            }
                                            var v62 uint64 = 2
                                            _ = v62
                                        }
                                    case -178:
                                        var v63 uint16 = 834
                                        _ = v63
                                    default:
                                        v5["hello"] = 4294967295
                                    }
                                    _ = v53
                                }
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            for fi64 := 0; fi64 < 14; fi64++ {
                                if v29 {
                                    v65 := GenS2{F0: fi44, F1: 1, F2: v2}
                                    switch 215 {
                                    case 127:
                                        fmt.Println(v29, v6)
                                    case -9223372036854775808:
                                        v6 = -21
                                    }
                                    v4 = append(v4, ("" + (("foo bar" + "a") + ((("a" + ("foo bar" + "")) + "") + "foo bar"))))
                                    _ = v65
                                }
                                switch (fi44 | (fi44 ^ 4294967295)) {
                                case 840:
                                    var v66 bool = (!(-32768 == -32768))
                                    if (2 <= 65535) {
                                        v66 = (0 > 297)
                                        v67 := GenS5{F0: (^(fi44)), F1: v6, F2: -1}
                                        if v66 {
                                            switch -842 {
                                            case -195:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case 466:
                                                var v68 int = fi64
                                                _ = v68
                                            }
                                        }
                                        _ = v67
                                    }
                                    _ = v66
                                case -2147483648:
                                    if (42 != 42) {
                                        v2 += 65535
                                        var v69 float32 = -2.25
                                        _ = v69
                                    }
                                }
                                _ = fi64
                            }
                        }
                    }
                    _ = fi44
                    _ = v45
                }
            }
            var v70 uint16 = 42
            _ = v70
        }
    default:
        v71 := GenS6{F0: "foo bar", F1: (v2 >> ((42) & 63)), F2: (v2 ^ v2)}
        _ = v71
    }
    if true {
        fmt.Println(v6, v6)
        v72 := GenS5{F0: -1, F1: v6, F2: 0}
        for fi73 := 0; fi73 < 14; fi73++ {
            fmt.Println(v29)
            var v74 float64 = (3.14159 / 0.1)
            var v75 uint64 = 0
            _ = fi73
            _ = v74
            _ = v75
        }
        _ = v72
    }
    var v76 float64 = ((0.0 * 1e10) * 0.0)
    fmt.Println(v76, v6)
    var v77 int64 = -512
    v29 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
}
