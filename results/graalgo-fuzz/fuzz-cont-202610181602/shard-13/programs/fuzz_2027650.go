package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int; F1 int8; F2 uint8 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -32768
    var v3 uint32 = 2147483647
    var v4 uint64 = 823
    var v5 int32 = -1
    var v6 int32 = (v5 - (^(v5)))
    for fi7 := 0; fi7 < 11; fi7++ {
        v5 -= -800
        var v8 bool = true
        delete(v1, fi7)
        _ = fi7
        _ = v8
    }
    v2 += (v2 - (v2 | v2))
    switch 143 {
    case -789:
        fmt.Println(v3)
    case 0:
        fmt.Println(v4, v2, v3)
    case 213:
        switch (("\n\t" + "foo bar") + ("foo bar" + "\n\t")) {
        case "\n\t":
            fmt.Println(v6)
        case "foo bar":
            v2 = v2
        case "hello":
            for fi9 := 0; fi9 < 3; fi9++ {
                for fi10 := 0; fi10 < 18; fi10++ {
                    if true {
                        var v11 uint64 = (v4 | (v4 | 4294967295))
                        _ = v11
                    } else {
                        _ = len(v1)
                        v12 := GenS1{F0: 2}
                        _ = v12
                    }
                    if (42 >= -9223372036854775808) {
                        fmt.Println(v4, fi9)
                        switch -9223372036854775808 {
                        case 2147483647:
                            v2 -= int16(v5)
                        case 32767:
                            v13 := []int{-2147483648}
                            _ = v13
                        case 962:
                            v4 = (^(v4))
                        }
                        v14 := GenS1{F0: 255}
                        _ = v14
                    }
                    _ = fi10
                }
                v3 -= 2147483648
                _ = fi9
            }
        default:
            v15 := []int{-599, 9223372036854775807}
            _ = v15
        }
    }
    var v16 float32 = -2.25
    v17 := GenS1{F0: 18446744073709551615}
    var v18 float64 = 0.1
    var v19 uint16 = 65535
    v20 := GenS2{F0: "a", F1: 9223372036854775807, F2: v19}
    var v21 int = 42
    if (!(!(!(0 == 255)))) {
        for fi22 := 0; fi22 < 2; fi22++ {
            v21 += int(v3)
            v19 += 545
            _ = fi22
        }
    } else {
        if false {
            v23 := []int{(v21 << ((42) & 63)), v21, 1}
            _ = v23
        } else {
            switch v21 {
            case 988:
                if true {
                    v21 = (-(v21))
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v4 = 42
                        var v25 int16 = int16(v4)
                        _ = fi24
                        _ = v25
                    }
                    v26 := GenS3{F0: (v4 & (v4 << ((688) & 63))), F1: v21}
                    _ = v26
                } else {
                    var v27 uint64 = (v4 % ((v4) | 1))
                    _ = v27
                }
            case 704:
                fmt.Println(v4, v2)
            default:
                v2 -= 42
            }
            switch "foo bar" {
            case "foo bar":
                if true {
                    v19 = (v19 * (^(v19)))
                } else {
                    v6 -= (v6 << ((510) & 63))
                    v3 = (v3 * (^(v3)))
                    v28 := GenS3{F0: v4, F1: (^(v21))}
                    _ = v28
                }
            }
            if true {
                var v29 uint32 = 4294967295
                v30 := GenS2{F0: ("hello" + "foo bar"), F1: 9223372036854775807, F2: v19}
                _ = v29
                _ = v30
            } else {
                v19 = v19
                if true {
                    for fi31 := 0; fi31 < 18; fi31++ {
                        v5 = 103
                        _ = fi31
                    }
                    v32 := GenS4{F0: (-(v21)), F1: -128, F2: 42}
                    if (((false || (0.1 < v18)) && (v3 != 0)) || false) {
                        for fi33 := 0; fi33 < 14; fi33++ {
                            if ((65535 < 18446744073709551615) || (1 <= 0)) {
                                var v34 uint32 = v3
                                var v35 int8 = 1
                                fmt.Println(fi33)
                                _ = v34
                                _ = v35
                            }
                            _ = fi33
                        }
                    }
                    _ = v32
                } else {
                    if false {
                        v36 := []int{(v21 + (v21 / (((v21 & v21)) | 1))), int(v3)}
                        v4 = 2
                        _ = v36
                    } else {
                        var v37 uint32 = 42
                        for fi38 := 0; fi38 < 11; fi38++ {
                            if (true || false) {
                                v4 += 384
                                var v39 int32 = -2147483648
                                var v40 uint = 42
                                _ = v39
                                _ = v40
                            }
                            switch 47 {
                            case 0:
                                var v41 uint8 = 128
                                _ = v41
                            case -402:
                                var v42 int16 = (v2 + (v2 / (((v2 + v2)) | 1)))
                                _ = v42
                            }
                            switch "\n\t" {
                            case "hello":
                                if (v16 == 1e10) {
                                    fmt.Println(v21)
                                    switch "a" {
                                    case "":
                                        var v43 float32 = 0.1
                                        _ = v43
                                    case "hello":
                                        var v44 float64 = (v18 + (0.0 * (v18 * v18)))
                                        _ = v44
                                    case "a":
                                        v45 := GenS3{F0: 4294967295, F1: fi38}
                                        _ = v45
                                    default:
                                        fmt.Println(fi38, v21)
                                    }
                                } else {
                                    v16 += (-2.25 / float32(v18))
                                    if false {
                                        v46 := GenS3{F0: 2, F1: 9223372036854775807}
                                        var v47 float64 = v18
                                        _ = v46
                                        _ = v47
                                    } else {
                                        var v48 int64 = 1
                                        var v49 uint = 4294967295
                                        _ = len(v1)
                                        _ = v48
                                        _ = v49
                                    }
                                }
                            case "a":
                                v5 -= int32(v2)
                            case "":
                                var v50 float32 = (v16 * ((0.0 - v16) + (((v16 + (1.5 * v16)) / (1e10 * float32((v18 - 1e10)))) + (0.0 / (v16 * -2.25)))))
                                _ = v50
                            default:
                                if false {
                                    var v51 string = "foo bar"
                                    if true {
                                        fmt.Println(v6)
                                        var v52 uint32 = v3
                                        switch v51 {
                                        case "\n\t":
                                            v4 = v4
                                        default:
                                            v21 = (v21 ^ 606)
                                        }
                                        _ = v52
                                    } else {
                                        fmt.Println(v21)
                                        v37 = (v3 | (v3 * v3))
                                    }
                                    v51 = v51
                                    _ = v51
                                } else {
                                    for fi53 := 0; fi53 < 10; fi53++ {
                                        var v54 bool = true
                                        if (!v54) {
                                            if v54 {
                                                var v55 int32 = (v6 + (v5 << ((18446744073709551615) & 63)))
                                                fmt.Println(v21, v5)
                                                v6 = (v6 * v55)
                                                _ = v55
                                            } else {
                                                fmt.Println(v21, fi38)
                                                v56 := GenS5{F0: (v21 ^ 127)}
                                                if (int8(v3) <= 127) {
                                                    v37 -= uint32(v4)
                                                    switch 0 {
                                                    case 128:
                                                        fmt.Println(v4, v6, v3)
                                                    case 32768:
                                                        var v57 int32 = 42
                                                        _ = v57
                                                    case 32767:
                                                        var v58 int16 = (v2 - v2)
                                                        _ = v58
                                                    default:
                                                        var v59 string = "foo bar"
                                                        _ = v59
                                                    }
                                                }
                                                _ = v56
                                            }
                                            var v60 int = 9223372036854775807
                                            if false {
                                                fmt.Println(v3, v37)
                                                v60 = 42
                                            } else {
                                                var v61 uint8 = 2
                                                _ = v61
                                            }
                                            fmt.Println(v37, v5, v16)
                                            v1[(^(v21))] = ""
                                            _ = v60
                                        } else {
                                            v62 := []int{2147483647, v21}
                                            fmt.Println(v21, v2)
                                            if ((18446744073709551615 != (v4 >> ((4294967295) & 63))) || v54) {
                                                var v63 uint16 = (^(v19))
                                                var v64 float32 = -2.25
                                                _ = v63
                                                _ = v64
                                            }
                                            v6 = (^(v6))
                                            _ = v62
                                        }
                                        _ = fi53
                                        _ = v54
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        v65 := []int{v21}
                                        _ = v65
                                    case "":
                                        if false {
                                            v2 += v2
                                            v66 := v1[(v21 ^ 127)]
                                            fmt.Println(v37, v37)
                                            _ = v66
                                        } else {
                                            fmt.Println(v37, v19, v18)
                                            for fi67 := 0; fi67 < 5; fi67++ {
                                                v18 = (v18 + (float64(1e10) / ((((float64((((v16 * v16) + float32(3.14159)) + ((v16 - -2.25) / (0.0 + v16)))) / v18) + (v18 * v18)) - float64(((v16 - (0.0 / v16)) - (v16 - ((v16 / -2.25) / (float32(0.1) * ((((v16 * v16) / float32(1e10)) + v16) / 3.14159))))))) * v18)))
                                                v18 = float64(float32(1e10))
                                                _ = fi67
                                            }
                                        }
                                    case "foo bar":
                                        if true {
                                            var v68 int32 = 42
                                            _ = v68
                                        }
                                    default:
                                        var v69 uint = 0
                                        _ = v69
                                    }
                                }
                            }
                            _ = fi38
                        }
                        if true {
                            if (65535 < (v19 + 42)) {
                                var v70 int = -2147483648
                                _ = v70
                            }
                            v18 -= (v18 + ((v18 + v18) + v18))
                        } else {
                            var v71 int8 = 0
                            _ = v71
                        }
                        _ = v37
                    }
                }
                fmt.Println(v21, v18, v2)
            }
        }
    }
    v72 := []int{(v21 - (v21 + 42)), (v21 % (((v21 ^ v21)) | 1))}
    v18 = -2.25
    var v73 string = ""
    v74 := GenS3{F0: 2, F1: (v21 | (v21 / (((v21 + v21)) | 1)))}
    if (!false) {
        v3 = (v3 | (v3 * v3))
    }
    var v75 int32 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v73
    _ = v74
    _ = v75
}
