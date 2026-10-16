package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int; F1 uint32; F2 int }
type GenS3 struct { F0 int32; F1 int16; F2 string }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 int; F1 int; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint = 2
    var v2 int16 = -1
    var v3 bool = (0 >= 65535)
    var v4 bool = true
    v5 := []int{32767}
    v6 := GenS1{F0: 32768, F1: -1}
    v3 = true
    switch "a" {
    case "":
        for fi7 := 0; fi7 < 19; fi7++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v3 = v3
            _ = fi7
        }
    case "hello":
        v8 := GenS1{F0: 0, F1: -2147483648}
        _ = v8
    case "a":
        switch "a" {
        case "":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "foo bar":
            var v9 int32 = 744
            _ = v9
        case "hello":
            for fi10 := 0; fi10 < 4; fi10++ {
                if v3 {
                    if true {
                        for fi11 := 0; fi11 < 12; fi11++ {
                            fmt.Println(v2, v2)
                            var v12 uint8 = 127
                            switch "\n\t" {
                            case "hello":
                                v1 += (v1 | (v1 & v1))
                            case "":
                                v3 = ((-989 <= -983) || ((v3 && false) && (!v4)))
                            case "\n\t":
                                v13 := GenS1{F0: fi10, F1: -2147483648}
                                _ = v13
                            default:
                                if v4 {
                                    if v4 {
                                        if (!(!false)) {
                                            v14 := GenS1{F0: (fi10 % ((42) | 1)), F1: 0}
                                            var v15 int32 = 42
                                            v16 := GenS1{F0: (fi10 | (-(fi11))), F1: (v15 % (((v15 << (((v1 * (v1 - v1))) & 63))) | 1))}
                                            _ = v14
                                            _ = v15
                                            _ = v16
                                        }
                                        fmt.Println(v1)
                                    }
                                    v17 := GenS1{F0: 0, F1: -715}
                                    _ = v17
                                }
                            }
                            _ = fi11
                            _ = v12
                        }
                    }
                    var v18 int32 = 24
                    switch "\n\t" {
                    case "a":
                        _ = len(v5)
                    default:
                        if (549 > 227) {
                            for fi19 := 0; fi19 < 10; fi19++ {
                                var v20 float32 = float32(0.1)
                                switch "hello" {
                                case "\n\t":
                                    v20 = -2.25
                                case "":
                                    var v21 int8 = 127
                                    _ = v21
                                default:
                                    if true {
                                        var v22 int16 = v2
                                        if true {
                                            switch "" {
                                            case "hello":
                                                var v23 uint32 = 442
                                                _ = v23
                                            case "\n\t":
                                                var v24 float32 = 0.1
                                                _ = v24
                                            }
                                        } else {
                                            var v25 uint64 = 113
                                            v22 = (v2 ^ (v2 | (^(v22))))
                                            v26 := GenS1{F0: 128, F1: v18}
                                            _ = v25
                                            _ = v26
                                        }
                                        v3 = v3
                                        _ = v22
                                    } else {
                                        var v27 uint8 = 255
                                        _ = v27
                                    }
                                }
                                var v28 int16 = 1000
                                _ = fi19
                                _ = v20
                                _ = v28
                            }
                            v29 := GenS2{F0: 42, F1: 768, F2: -1}
                            _ = v29
                        }
                    }
                    _ = v18
                }
                switch 3 {
                case -445:
                    var v30 int32 = -1
                    _ = v30
                default:
                    var v31 float32 = 1.5
                    _ = v31
                }
                _ = fi10
            }
        default:
            var v32 int8 = 1
            _ = v32
        }
    default:
        v4 = (2 < 2)
    }
    var v33 string = ("hello" + "a")
    v4 = (v2 != (v2 * (^(v2))))
    var v34 int64 = 0
    if (2 < 65535) {
        for fi35 := 0; fi35 < 1; fi35++ {
            v4 = (v4 || (!(127 != 127)))
            fmt.Println(v4, v4, v4)
            var v36 uint8 = 0
            _ = fi35
            _ = v36
        }
    } else {
        if v3 {
            v37 := GenS2{F0: 128, F1: 42, F2: 128}
            _ = v37
        } else {
            if len(v5) > 0 {
                v5[0] = 0
            }
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            if (36 != 42) {
                var v39 string = v33
                _ = v39
            } else {
                v5 = append(v5, fi38)
                if len(v5) > 0 {
                    v5[0] = 32767
                }
                if (!(!((1.5 == 1e10) && false))) {
                    if v4 {
                        v40 := GenS3{F0: 40, F1: (v2 / ((v2) | 1)), F2: v33}
                        var v41 uint64 = 42
                        if true {
                            var v42 int16 = (v2 + v2)
                            for fi43 := 0; fi43 < 10; fi43++ {
                                if v4 {
                                    _ = len(v5)
                                    v44 := GenS4{F0: 65535, F1: (v1 - (v1 * v1))}
                                    _ = v44
                                }
                                _ = fi43
                            }
                            v33 = ("" + "\n\t")
                            _ = v42
                        }
                        _ = v40
                        _ = v41
                    } else {
                        var v45 uint64 = 2
                        if v3 {
                            var v46 string = v33
                            v46 = ""
                            _ = v46
                        } else {
                            _ = len(v5)
                            switch v33 {
                            case "\n\t":
                                switch (fi38 + 32768) {
                                case -1:
                                    v2 += 808
                                case 128:
                                    if v4 {
                                        if v4 {
                                            switch 32768 {
                                            case 9223372036854775807:
                                                switch fi38 {
                                                case 128:
                                                    v33 = v33
                                                case 32767:
                                                    if (-2.25 != 1.5) {
                                                        if ((0 > -1) || (int32(v34) >= -800)) {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            switch -1 {
                                                            case -65:
                                                                var v47 uint = 255
                                                                _ = v47
                                                            default:
                                                                for fi48 := 0; fi48 < 9; fi48++ {
                                                                    switch (v33 + (v33 + v33)) {
                                                                    case "a":
                                                                        v49 := GenS5{F0: (-(fi38)), F1: (fi38 / (((fi38 + fi48)) | 1)), F2: 9223372036854775807}
                                                                        v4 = v3
                                                                        var v50 int = (fi38 + (fi48 << ((v1) & 63)))
                                                                        _ = v49
                                                                        _ = v50
                                                                    default:
                                                                        if ((v34 ^ (^(v34))) != v34) {
                                                                            var v51 float32 = 3.14159
                                                                            if len(v5) > 0 {
                                                                                v5[0] = (fi38 & (fi48 & 2147483647))
                                                                            }
                                                                            _ = v51
                                                                        }
                                                                    }
                                                                    _ = fi48
                                                                }
                                                            }
                                                            var v52 int8 = 118
                                                            _ = v52
                                                        } else {
                                                            if v3 {
                                                                if true {
                                                                    v53 := GenS6{F0: false}
                                                                    v54 := GenS7{F0: 42, F1: true}
                                                                    v1 = 389
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    if (-9223372036854775808 == int(v45)) {
                                                                        var v55 int32 = 42
                                                                        v56 := GenS2{F0: fi38, F1: 42, F2: 326}
                                                                        _ = v55
                                                                        _ = v56
                                                                    }
                                                                }
                                                                var v57 float32 = 0.1
                                                                _ = v57
                                                            } else {
                                                                var v58 int64 = (v34 << ((65535) & 63))
                                                                _ = v58
                                                            }
                                                        }
                                                        var v59 uint8 = 96
                                                        _ = v59
                                                    }
                                                }
                                            default:
                                                var v60 uint = 2
                                                _ = v60
                                            }
                                        } else {
                                            v3 = true
                                        }
                                        for fi61 := 0; fi61 < 2; fi61++ {
                                            var v62 int = fi61
                                            _ = fi61
                                            _ = v62
                                        }
                                        var v63 int16 = (v2 % (((v2 - (-(v2)))) | 1))
                                        _ = v63
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = 128
                                        }
                                        v45 = v45
                                    }
                                case -6:
                                    v1 -= v1
                                }
                            default:
                                for fi64 := 0; fi64 < 6; fi64++ {
                                    fmt.Println(v33, v34, v2)
                                    if len(v5) > 0 {
                                        v5[0] = -1
                                    }
                                    var v65 int8 = -2
                                    _ = fi64
                                    _ = v65
                                }
                            }
                            v66 := GenS5{F0: (-(fi38)), F1: (fi38 >> (((v1 - 2)) & 63)), F2: (fi38 ^ int(v34))}
                            _ = v66
                        }
                        for fi67 := 0; fi67 < 8; fi67++ {
                            if len(v5) > 0 {
                                v5[0] = (^(fi67))
                            }
                            var v68 bool = (-1 > 47)
                            var v69 int16 = (v2 >> (((v1 & v1)) & 63))
                            _ = fi67
                            _ = v68
                            _ = v69
                        }
                        _ = v45
                    }
                } else {
                    var v70 uint8 = 29
                    if len(v5) > 0 {
                        v5[0] = fi38
                    }
                    if (!((v3 && v4) && v4)) {
                        v3 = v3
                    }
                    _ = v70
                }
            }
            _ = fi38
        }
        v5 = append(v5, 2147483647)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
