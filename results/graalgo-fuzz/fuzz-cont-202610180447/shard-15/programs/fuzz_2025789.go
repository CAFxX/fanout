package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS7 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS8 struct { F0 int }

func main() {
    var v1 uint16 = 991
    v2 := make(map[int]bool)
    var v3 uint16 = (^(v1))
    v4 := []string{"hello", ""}
    v5 := []int{-9223372036854775808}
    v6 := []string{}
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v1 = (v3 >> ((188) & 63))
            v9 := GenS1{F0: 131, F1: 65535, F2: (fi7 % ((fi8) | 1))}
            v10 := GenS2{F0: ("a" + "a"), F1: (^(fi8))}
            _ = fi8
            _ = v10
            _ = v9
        }
        switch "" {
        case "a":
            var v11 int8 = 36
            _ = v11
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            switch -794 {
            case 903:
                v13 := GenS2{F0: "foo bar", F1: 468}
                _ = v13
            case 875:
                if (fi12 <= fi7) {
                    v14 := GenS2{F0: "\n\t", F1: fi12}
                    fmt.Println(v3, fi12)
                    var v15 uint16 = v1
                    fmt.Println(v15, v15, v3)
                    _ = v14
                    _ = v15
                } else {
                    v16 := GenS3{F0: 2}
                    _ = v16
                }
            case 1:
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
            default:
                v17 := GenS2{F0: "\n\t", F1: fi12}
                _ = v17
            }
            if true {
                v1 = v3
                fmt.Println(v3, fi12)
                switch ("foo bar" + "a") {
                case "hello":
                    if false {
                        v1 -= (v3 / (((v1 & v3)) | 1))
                    }
                case "":
                    var v18 uint16 = v1
                    _ = v18
                case "a":
                    var v19 uint64 = 4294967295
                    _ = v19
                }
            } else {
                var v20 int16 = 963
                v21 := v2[-9223372036854775808]
                v22 := GenS3{F0: 384}
                v23 := GenS4{F0: v20}
                _ = v20
                _ = v21
                _ = v22
                _ = v23
            }
            var v24 int = (fi7 % ((fi7) | 1))
            v25 := v2[154]
            _ = fi12
            _ = v24
            _ = v25
        }
        _ = fi7
    }
    fmt.Println(v1, v1)
    var v26 int16 = 1000
    for fi27 := 0; fi27 < 17; fi27++ {
        v28 := GenS5{F0: (^(fi27))}
        if (true && true) {
            if (3.14159 >= ((-2.25 * (3.14159 * 0.0)) + -2.25)) {
                v1 -= (v1 * (v3 / (((v3 % (((v1 % ((v3) | 1))) | 1))) | 1)))
            } else {
                for fi29 := 0; fi29 < 14; fi29++ {
                    v30 := GenS6{F0: v1, F1: (("hello" + ("" + "foo bar")) < ""), F2: 2}
                    _ = fi29
                    _ = v30
                }
                fmt.Println(v1, v26, fi27)
                var v31 int64 = -39
                _ = v31
            }
            for fi32 := 0; fi32 < 14; fi32++ {
                v2[int(v1)] = ((fi32 | fi32) <= fi27)
                _ = fi32
            }
        }
        _ = fi27
        _ = v28
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        v34 := GenS6{F0: v3, F1: (4294967295 != 2), F2: 174}
        v2[-625] = false
        _ = fi33
        _ = v34
    }
    var v35 uint = 1
    switch "foo bar" {
    case "a":
        switch "" {
        case "a":
            v36 := GenS5{F0: -418}
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        switch ("" + "a") {
        case "a":
            var v38 int8 = 42
            _ = v38
        case "":
            var v39 int8 = 0
            _ = v39
        case "\n\t":
            delete(v2, 128)
        }
        var v40 int = -762
        v40 = v40
        _ = fi37
        _ = v40
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    if ((!false) || (!true)) {
        if false {
            fmt.Println(v3)
            var v41 uint64 = 1
            _ = v41
        } else {
            v35 = (v35 + 348)
        }
    }
    var v42 float32 = 1.5
    var v43 int32 = 876
    v35 += (v35 * (v35 >> (((v35 * (v35 & (v35 >> (((v35 + v35)) & 63))))) & 63)))
    for fi44 := 0; fi44 < 4; fi44++ {
        var v45 int32 = v43
        if false {
            switch fi44 {
            case 32767:
                switch 9223372036854775807 {
                case -9223372036854775808:
                    var v46 float32 = (1.5 / (v42 - ((((v42 / 0.1) - v42) / v42) / v42)))
                    _ = v46
                }
            case 4294967295:
                v47 := GenS6{F0: uint16(v35), F1: ((42 != 2) || (-9223372036854775808 <= 9223372036854775807)), F2: 42}
                _ = v47
            case 1:
                if false {
                    switch "" {
                    case "":
                        fmt.Println(v26, fi44)
                    default:
                        var v48 int64 = -1
                        _ = v48
                    }
                }
            }
            for fi49 := 0; fi49 < 16; fi49++ {
                var v50 uint = (^(v35))
                v51 := GenS7{F0: 127, F1: (false && (!false)), F2: (v1 << (((v50 / ((18446744073709551615) | 1))) & 63))}
                _ = fi49
                _ = v50
                _ = v51
            }
            if true {
                v43 -= v43
                v3 = v3
            } else {
                for fi52 := 0; fi52 < 16; fi52++ {
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v53 uint32 = 1
                        _ = v53
                    } else {
                        var v54 uint64 = 4294967295
                        v55 := GenS8{F0: 32768}
                        if (("\n\t" + "foo bar") <= "hello") {
                            v1 = (v3 | v1)
                            if true {
                                var v56 int64 = int64(v3)
                                v54 += 4294967295
                                v56 += 0
                                _ = v56
                            } else {
                                var v57 float32 = (0.0 + 3.14159)
                                _ = v57
                            }
                            var v58 int64 = 347
                            if (!false) {
                                fmt.Println(v26, v45, v35)
                                var v59 uint32 = uint32(v3)
                                if false {
                                    v42 += v42
                                    _ = len(v5)
                                } else {
                                    fmt.Println(v58, v43, v26)
                                    v60 := v2[int(v45)]
                                    _ = v60
                                }
                                _ = v59
                            } else {
                                fmt.Println(v1, v42, fi44)
                            }
                            if ((32767 <= 471) || false) {
                                if false {
                                    fmt.Println(fi52)
                                    if true {
                                        if false {
                                            _ = len(v2)
                                            if (476 == (^(v3))) {
                                                if (!(118 >= 127)) {
                                                    var v61 int = fi52
                                                    v26 += int16(v35)
                                                    var v62 int16 = v26
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    v63 := v2[(fi52 << (((^(v35))) & 63))]
                                                    _ = v63
                                                }
                                            } else {
                                                fmt.Println(fi52)
                                            }
                                            if true {
                                                v64 := GenS4{F0: v26}
                                                var v65 int32 = v45
                                                _ = v64
                                                _ = v65
                                            }
                                        }
                                        v42 = (v42 + (-2.25 + v42))
                                    }
                                } else {
                                    _ = len(v4)
                                    v66 := v2[fi44]
                                    _ = v66
                                }
                            } else {
                                v45 = -1
                                v54 = (^(v54))
                            }
                            _ = len(v2)
                            if (false && (fi52 != -2147483648)) {
                                v58 -= 42
                            }
                            _ = v58
                        } else {
                            fmt.Println(fi44)
                        }
                        _ = v54
                        _ = v55
                    }
                    _ = fi52
                }
                v3 = v1
                v42 = -2.25
            }
        }
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
