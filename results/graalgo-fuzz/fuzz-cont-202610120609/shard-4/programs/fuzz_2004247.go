package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 string; F1 string; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 uint }
func genFn0() uint {
    if true {
        var v5 float64 = 1e10
        var v6 string = (((("hello" + ("\n\t" + "hello")) + "\n\t") + ("\n\t" + ("\n\t" + "a"))) + "foo bar")
        if (!(0 < 867)) {
            v5 = v5
            v7 := GenS1{F0: 932, F1: 286}
            _ = v7
        } else {
            var v8 uint64 = 921
            if (!true) {
                if true {
                    switch -1 {
                    case 204:
                        if false {
                            if (true && true) {
                                v6 = v6
                            } else {
                                fmt.Println(v6, v6, v6)
                            }
                        }
                    case -12:
                        for fi9 := 0; fi9 < 8; fi9++ {
                            for fi10 := 0; fi10 < 19; fi10++ {
                                v5 = v5
                                v11 := GenS1{F0: (v8 >> ((2) & 63)), F1: 0}
                                if true {
                                    if true {
                                        var v12 uint8 = 127
                                        _ = v12
                                    } else {
                                        v6 = ""
                                        switch ((v6 + ("\n\t" + "foo bar")) + v6) {
                                        case "\n\t":
                                            var v13 uint64 = v8
                                            _ = v13
                                        case "hello":
                                            fmt.Println(v6)
                                        case "":
                                            fmt.Println(v5, v5, v5)
                                        default:
                                            var v14 int16 = 307
                                            _ = v14
                                        }
                                    }
                                    v15 := GenS2{F0: (v8 ^ (v8 & v8))}
                                    _ = v15
                                } else {
                                    v16 := GenS2{F0: (v8 ^ 2)}
                                    _ = v16
                                }
                                _ = fi10
                                _ = v11
                            }
                            _ = fi9
                        }
                    default:
                        v17 := []int{552}
                        _ = v17
                    }
                    v8 += (v8 >> ((1) & 63))
                }
            } else {
                for fi18 := 0; fi18 < 19; fi18++ {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v20 := GenS2{F0: (^(v8))}
                        _ = fi19
                        _ = v20
                    }
                    var v21 bool = false
                    _ = fi18
                    _ = v21
                }
            }
            var v22 uint = 255
            _ = v22
            _ = v8
        }
        _ = v5
        _ = v6
    } else {
        if (!true) {
            var v23 int = -1
            _ = v23
        } else {
            v24 := []int{582, 32768, 128}
            _ = v24
        }
    }
    var v25 uint8 = 1
    _ = v25
    return 255
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]int)
    var v3 int32 = 42
    var v4 int = 897
    var v26 uint8 = 127
    v27 := []int{(v4 & v4), v4, v4}
    v1 += 0.0
    var v28 uint32 = 710
    switch (("foo bar" + "a") + "hello") {
    case "hello":
        var v29 int16 = 32767
        _ = v29
    case "\n\t":
        if true {
            v30 := GenS1{F0: 42, F1: (v28 << ((2) & 63))}
            var v31 int = v4
            var v32 uint16 = 32767
            _ = v30
            _ = v31
            _ = v32
        }
    case "foo bar":
        v28 = (v28 | v28)
    default:
        v27 = append(v27, 866)
    }
    var v33 int32 = v3
    var v34 uint = 858
    switch 42 {
    case 4294967295:
        switch (^(v4)) {
        case 32768:
            var v35 float64 = 1.5
            _ = v35
        case -2147483648:
            switch (v4 & (v4 | (^(v4)))) {
            case -918:
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
            case 9223372036854775807:
                v4 += 32767
            case -9223372036854775808:
                v28 -= v28
            default:
                fmt.Println(v33, v4)
            }
        case -746:
            if true {
                switch (v4 | (v4 - 127)) {
                case 127:
                    fmt.Println(v3)
                case 32768:
                    switch -664 {
                    case -966:
                        var v36 float64 = 1e10
                        _ = v36
                    }
                case 42:
                    switch v4 {
                    case 4294967295:
                        for fi37 := 0; fi37 < 12; fi37++ {
                            v34 -= 65535
                            _ = fi37
                        }
                    default:
                        v4 += v4
                    }
                }
            }
        default:
            var v38 uint = uint(v26)
            _ = v38
        }
    case -595:
        switch 210 {
        case 42:
            for fi39 := 0; fi39 < 17; fi39++ {
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
                fmt.Println(v26, v26, fi39)
                v40 := GenS1{F0: 0, F1: v28}
                _ = fi39
                _ = v40
            }
        default:
            if (v28 == 541) {
                v3 = v3
            } else {
                for fi41 := 0; fi41 < 14; fi41++ {
                    for fi42 := 0; fi42 < 9; fi42++ {
                        _ = len(v27)
                        fmt.Println(v26, v3, v3)
                        _ = fi42
                    }
                    _ = fi41
                }
                for fi43 := 0; fi43 < 6; fi43++ {
                    v4 -= (fi43 << ((826) & 63))
                    _ = fi43
                }
            }
        }
    case -454:
        v44 := GenS3{F0: 18446744073709551615, F1: (v28 * v28)}
        _ = v44
    default:
        if true {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v34)
            case "hello":
                if len(v27) >= 2 {
                    v27 = v27[1:len(v27)]
                }
            default:
                v2[(v4 - v4)] = v4
            }
            if (2147483648 > (v28 % ((uint32(v26)) | 1))) {
                var v45 int64 = int64(v34)
                var v46 bool = true
                fmt.Println(v28)
                _ = v45
                _ = v46
            } else {
                v4 += 1
            }
            if (false && false) {
                if (87 != 42) {
                    for fi47 := 0; fi47 < 3; fi47++ {
                        v48 := GenS3{F0: 757, F1: (v28 ^ 65535)}
                        _ = fi47
                        _ = v48
                    }
                    v2[v4] = (v4 + v4)
                    v49 := GenS4{F0: "a", F1: "\n\t", F2: -337}
                    _ = v49
                } else {
                    if true {
                        var v50 uint8 = 185
                        v2[v4] = v4
                        _ = v50
                    }
                    switch (^(v4)) {
                    case -726:
                        v51 := GenS4{F0: "", F1: "foo bar", F2: -1}
                        _ = v51
                    case 9223372036854775807:
                        _ = len(v27)
                    case 2147483647:
                        v52 := GenS1{F0: 0, F1: v28}
                        _ = v52
                    default:
                        v27 = append(v27, (v4 - (v4 % ((v4) | 1))))
                    }
                }
            } else {
                _ = len(v2)
                fmt.Println(v26, v1, v1)
                v53 := GenS2{F0: 4294967295}
                _ = v53
            }
        } else {
            v54 := GenS5{F0: (^(v28))}
            v55 := v2[42]
            var v56 bool = (1 > 1)
            _ = v54
            _ = v55
            _ = v56
        }
    }
    var v57 float32 = 3.14159
    var v58 int8 = 1
    fmt.Println(v58, v33, v1)
    v3 += int32(v26)
    v4 += (v4 >> ((genFn0()) & 63))
    if false {
        v59 := GenS1{F0: 1, F1: (v28 + uint32(v34))}
        _ = v59
    }
    if len(v27) > 0 {
        v27[0] = 9223372036854775807
    }
    fmt.Println(v4)
    v60 := GenS5{F0: (v28 ^ v28)}
    for fi61 := 0; fi61 < 1; fi61++ {
        v62 := GenS6{F0: (v28 - v28), F1: (v34 & v34)}
        _ = fi61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v57
    _ = v58
    _ = v60
}
