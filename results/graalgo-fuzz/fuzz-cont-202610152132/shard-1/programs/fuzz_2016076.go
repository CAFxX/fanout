package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 uint64 = 4294967295
    var v2 bool = true
    var v3 bool = (42 > -128)
    v4 := []int{2147483647, -2147483648}
    v5 := GenS1{F0: 2147483648}
    v2 = (!true)
    fmt.Println(v1, v2)
    var v6 uint = 1
    var v7 uint32 = 4294967295
    v2 = v3
    if true {
        var v8 int64 = 1
        switch 78 {
        case 0:
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(v6)
                v8 = (-(v8))
                if ((v2 || (v8 <= 0)) || (!v3)) {
                    v10 := GenS2{F0: 1, F1: (!("a" < "hello")), F2: 127}
                    fmt.Println(v3, v6, v8)
                    if false {
                        v11 := GenS1{F0: 4294967295}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi12 := 0; fi12 < 13; fi12++ {
                            var v13 uint64 = (v1 | v1)
                            var v14 float32 = float32(3.14159)
                            var v15 uint16 = 1
                            _ = fi12
                            _ = v13
                            _ = v14
                            _ = v15
                        }
                        _ = v11
                    }
                    _ = v10
                }
                _ = fi9
            }
        case 32768:
            var v16 uint8 = 127
            _ = v16
        case 1:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            for fi17 := 0; fi17 < 17; fi17++ {
                v18 := GenS1{F0: (v7 & v7)}
                _ = fi17
                _ = v18
            }
        }
        switch "\n\t" {
        case "":
            if v3 {
                if v2 {
                    v3 = v2
                    if true {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            var v20 int8 = 108
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        if len(v4) > 0 {
                            v4[0] = -948
                        }
                        for fi21 := 0; fi21 < 11; fi21++ {
                            var v22 uint64 = uint64(v6)
                            if (((v8 << ((4294967295) & 63)) >= 0) && (v3 || (int(v6) > fi21))) {
                                if true {
                                    v23 := GenS1{F0: v7}
                                    switch ("hello" + "a") {
                                    case "\n\t":
                                        if v3 {
                                            var v24 uint = v6
                                            if v2 {
                                                var v25 int8 = int8(v24)
                                                var v26 uint8 = 191
                                                if len(v4) > 0 {
                                                    v4[0] = fi21
                                                }
                                                _ = v25
                                                _ = v26
                                            } else {
                                                var v27 uint32 = 44
                                                _ = v27
                                            }
                                            v28 := GenS1{F0: (v7 & v7)}
                                            _ = v24
                                            _ = v28
                                        }
                                    }
                                    _ = v23
                                }
                                var v29 int8 = 100
                                _ = v29
                            } else {
                                fmt.Println(v1, v8, v3)
                                var v30 float32 = 3.14159
                                v22 = v1
                                _ = v30
                            }
                            _ = fi21
                            _ = v22
                        }
                    }
                } else {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        switch (-(fi31)) {
                        case -1:
                            switch (^(fi31)) {
                            case 305:
                                var v32 int = fi31
                                _ = v32
                            case -1000:
                                for fi33 := 0; fi33 < 19; fi33++ {
                                    v4 = append(v4, fi31)
                                    _ = fi33
                                }
                            }
                        case 2147483647:
                            v3 = true
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        _ = fi31
                    }
                }
            }
        case "hello":
            v7 += v7
        default:
            var v34 uint64 = 42
            _ = v34
        }
        _ = v8
    } else {
        _ = len(v4)
    }
    _ = len(v4)
    var v35 int8 = int8(v1)
    v36 := GenS2{F0: (v35 % (((v35 - v35)) | 1)), F1: (1 < 65535), F2: v35}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (93 > (^(v7))) {
        var v37 float32 = (1.5 * 0.1)
        v38 := GenS3{F0: "", F1: -4}
        _ = v37
        _ = v38
    } else {
        fmt.Println(v7, v6, v35)
        fmt.Println(v7, v2)
    }
    if v3 {
        switch 4294967295 {
        case -1:
            var v39 uint64 = v1
            _ = v39
        default:
            var v40 float64 = 1e10
            _ = v40
        }
        v4 = append(v4, -17)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        v41 := GenS2{F0: int8(v1), F1: (v3 || (!v2)), F2: int8(v6)}
        var v42 float32 = (0.1 + float32(1.5))
        if v3 {
            if v2 {
                for fi43 := 0; fi43 < 13; fi43++ {
                    for fi44 := 0; fi44 < 2; fi44++ {
                        var v45 int64 = 9223372036854775807
                        v3 = (42 > 0)
                        _ = fi44
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 1; fi46++ {
                        var v47 int32 = 0
                        var v48 int = (fi43 % (((fi46 - 2147483647)) | 1))
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    _ = fi43
                }
                v6 = v6
            }
        }
        _ = v41
        _ = v42
    }
    if false {
        if false {
            for fi49 := 0; fi49 < 4; fi49++ {
                if v3 {
                    var v50 float64 = -2.25
                    var v51 int32 = -502
                    _ = v50
                    _ = v51
                } else {
                    var v52 uint = 654
                    v53 := GenS3{F0: "foo bar", F1: v35}
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v52
                    _ = v53
                }
                _ = fi49
            }
            v54 := GenS2{F0: v35, F1: true, F2: v35}
            v3 = v2
            _ = v54
        } else {
            var v55 uint32 = 2
            v56 := GenS3{F0: "foo bar", F1: (v35 | (v35 + 1))}
            _ = v55
            _ = v56
        }
    }
    var v57 int16 = 459
    if (221 != -1) {
        v58 := GenS3{F0: "a", F1: v35}
        v35 += 60
        for fi59 := 0; fi59 < 2; fi59++ {
            v35 += -6
            _ = fi59
        }
        _ = v58
    }
    if ((v7 - v7) == 197) {
        v60 := GenS1{F0: (v7 / ((uint32(v1)) | 1))}
        if false {
            if true {
                fmt.Println(v35)
                var v61 uint = v6
                _ = len(v4)
                _ = v61
            } else {
                if false {
                    fmt.Println(v7, v1, v7)
                } else {
                    if v2 {
                        if ((v35 << ((4294967295) & 63)) < (-(v35))) {
                            var v62 uint = uint(v7)
                            _ = v62
                        } else {
                            var v63 uint32 = (v7 ^ (v7 >> ((v6) & 63)))
                            _ = v63
                        }
                        v64 := GenS3{F0: "a", F1: (v35 % ((v35) | 1))}
                        if v3 {
                            var v65 int8 = (v35 << ((v6) & 63))
                            var v66 uint64 = (v1 & 459)
                            _ = v65
                            _ = v66
                        }
                        _ = v64
                    }
                }
                if ((!((v57 / (((v57 >> (((v6 | v6)) & 63))) | 1)) >= v57)) || (!v3)) {
                    v67 := GenS4{F0: (^(v1)), F1: 42}
                    if (true || v3) {
                        var v68 int8 = v35
                        if len(v4) > 0 {
                            v4[0] = -1
                        }
                        _ = v68
                    }
                    v69 := GenS4{F0: v1, F1: 15}
                    _ = v67
                    _ = v69
                }
            }
            for fi70 := 0; fi70 < 9; fi70++ {
                v7 = (v7 % ((2) | 1))
                v57 = -716
                if ((^(v7)) > (v7 % ((v7) | 1))) {
                    v71 := GenS5{F0: uint64(v7)}
                    v72 := GenS2{F0: v35, F1: (false && false), F2: v35}
                    _ = v71
                    _ = v72
                } else {
                    var v73 uint32 = (v7 / ((511) | 1))
                    _ = v73
                }
                _ = fi70
            }
            v3 = (!(v2 || (2 <= 1)))
        }
        _ = v60
    } else {
        var v74 bool = (-111 > 100)
        for fi75 := 0; fi75 < 15; fi75++ {
            fmt.Println(fi75, v35)
            _ = fi75
        }
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v57
    _ = v6
    _ = v7
}
