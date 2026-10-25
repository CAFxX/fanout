package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int8; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 int32; F1 uint8; F2 string }
type GenS7 struct { F0 uint64; F1 int8; F2 uint16 }

func main() {
    var v1 uint64 = 823
    v2 := make(map[uint32]int)
    var v3 uint = 179
    var v4 uint = 0
    var v5 uint16 = 1
    switch (("foo bar" + (("" + "\n\t") + "foo bar")) + "a") {
    case "a":
        v4 += (^(v3))
    case "\n\t":
        delete(v2, 42)
    }
    switch 0 {
    case 4294967295:
        v6 := GenS1{F0: 4294967295, F1: 1}
        _ = v6
    case -2147483648:
        v5 += v5
    case -725:
        var v7 uint16 = v5
        _ = v7
    default:
        var v8 uint = (v4 * (v4 >> ((723) & 63)))
        _ = v8
    }
    var v9 uint32 = 0
    v1 = uint64(v9)
    v4 -= 1
    v9 = (^(v9))
    if (v1 >= (v1 | (v1 + v1))) {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := []int{fi10, fi10, fi10}
            v4 = v3
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 int64 = 1
                fmt.Println(v1)
                fmt.Println(v3, v4)
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        var v14 float64 = 3.14159
        _ = v14
    }
    v15 := GenS2{F0: 1, F1: (^(v1)), F2: 9223372036854775807}
    var v16 int32 = -1
    v3 = v3
    if false {
        v17 := GenS3{F0: (v5 + 799)}
        v5 = v5
        _ = v17
    } else {
        var v18 uint32 = 65535
        var v19 float32 = float32(float64(0.0))
        _ = v18
        _ = v19
    }
    v16 = v16
    v20 := GenS3{F0: 0}
    switch "\n\t" {
    case "foo bar":
        if (!((v16 - v16) <= v16)) {
            if ((^(v3)) != (^(v3))) {
                fmt.Println(v3, v9)
                switch 82 {
                case 128:
                    v21 := []int{0, 989, -114}
                    _ = v21
                default:
                    v22 := GenS1{F0: 2147483647, F1: 32767}
                    _ = v22
                }
                v23 := GenS4{F0: "hello", F1: 150}
                _ = v23
            }
            if true {
                v24 := GenS5{F0: 32767, F1: (v3 ^ 0)}
                v25 := GenS6{F0: v16, F1: 49, F2: "hello"}
                _ = v24
                _ = v25
            }
            switch 127 {
            case 32768:
                var v26 float32 = (0.1 / 0.1)
                _ = v26
            case -2147483648:
                switch 1 {
                case 127:
                    switch "a" {
                    case "hello":
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            switch "\n\t" {
                            case "foo bar":
                                var v27 uint64 = (v1 << (((v3 >> ((v3) & 63))) & 63))
                                _ = v27
                            case "hello":
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    v29 := GenS2{F0: 100, F1: v1, F2: -9223372036854775808}
                                    for fi30 := 0; fi30 < 11; fi30++ {
                                        v31 := []int{}
                                        switch "" {
                                        case "":
                                            var v32 int32 = v16
                                            _ = v32
                                        case "hello":
                                            if false {
                                                v31 = append(v31, fi28)
                                                fmt.Println(fi30, v5, fi28)
                                            }
                                        default:
                                            var v33 float64 = 1.5
                                            _ = v33
                                        }
                                        _ = fi30
                                        _ = v31
                                    }
                                    v34 := GenS6{F0: 2147483647, F1: 255, F2: ("hello" + "hello")}
                                    _ = fi28
                                    _ = v29
                                    _ = v34
                                }
                            case "":
                                v35 := GenS3{F0: uint16(v1)}
                                _ = v35
                            default:
                                v5 = v5
                            }
                        case "a":
                            switch -2147483648 {
                            case 127:
                                var v36 int8 = 100
                                _ = v36
                            }
                        default:
                            v37 := GenS2{F0: 127, F1: (v1 & (v1 / (((^(v1))) | 1))), F2: 9223372036854775807}
                            _ = v37
                        }
                    case "\n\t":
                        if false {
                            v9 -= (v9 + (v9 >> (((v4 % (((v3 << ((255) & 63))) | 1))) & 63)))
                            for fi38 := 0; fi38 < 9; fi38++ {
                                fmt.Println(v5, v1)
                                if ((0.0 / float64((float32(((0.1 - ((0.0 * 1e10) * 1.5)) + (-2.25 / 1.0))) * (((1e10 / float32(1e10)) * 1e10) + 0.1)))) != -2.25) {
                                    v1 = (v1 / (((v1 % (((v1 % ((42) | 1))) | 1))) | 1))
                                    fmt.Println(v1, v16, fi38)
                                    var v39 uint64 = (v1 & 236)
                                    _ = v39
                                } else {
                                    switch int(v9) {
                                    case 32767:
                                        v40 := GenS1{F0: (^(v9)), F1: 32767}
                                        _ = v40
                                    case 128:
                                        var v41 uint64 = v1
                                        _ = v41
                                    case -2147483648:
                                        if true {
                                            var v42 uint32 = v9
                                            v43 := GenS1{F0: v42, F1: 1000}
                                            _ = v42
                                            _ = v43
                                        } else {
                                            if true {
                                                switch ("foo bar" + "") {
                                                case "foo bar":
                                                    var v44 uint64 = (v1 - 46)
                                                    _ = v44
                                                case "a":
                                                    if ((v4 <= 817) && (!true)) {
                                                        if true {
                                                            v45 := GenS5{F0: 42, F1: v3}
                                                            _ = v45
                                                        } else {
                                                            if (v5 == 567) {
                                                                v4 -= (v3 / ((v3) | 1))
                                                            }
                                                            var v46 uint8 = 127
                                                            fmt.Println(v16, v4)
                                                            _ = v46
                                                        }
                                                    }
                                                case "":
                                                    v47 := []int{}
                                                    _ = v47
                                                }
                                                if (("a" + ("\n\t" + ("foo bar" + ("" + "")))) > "foo bar") {
                                                    if false {
                                                        var v48 int32 = v16
                                                        _ = v48
                                                    }
                                                    var v49 uint = v4
                                                    for fi50 := 0; fi50 < 9; fi50++ {
                                                        v51 := []int{fi50}
                                                        fmt.Println(v16, v49, v3)
                                                        fmt.Println(fi38)
                                                        _ = fi50
                                                        _ = v51
                                                    }
                                                    _ = v49
                                                } else {
                                                    v1 = 1
                                                }
                                            }
                                        }
                                    default:
                                        _ = len(v2)
                                    }
                                    v4 -= uint(v1)
                                }
                                var v52 int16 = 1000
                                _ = fi38
                                _ = v52
                            }
                        } else {
                            v3 -= (^(v4))
                            v9 += v9
                        }
                    case "a":
                        v9 += v9
                    }
                }
            case 128:
                if (false || false) {
                    for fi53 := 0; fi53 < 8; fi53++ {
                        fmt.Println(v5, v16, v5)
                        v16 -= int32(v5)
                        fmt.Println(v4, v5, fi53)
                        _ = fi53
                    }
                    var v54 uint = uint(v1)
                    var v55 uint32 = v9
                    _ = v54
                    _ = v55
                } else {
                    var v56 uint32 = 42
                    fmt.Println(v16)
                    var v57 float32 = 1.5
                    _ = v56
                    _ = v57
                }
            default:
                var v58 int16 = -469
                _ = v58
            }
        }
    }
    for fi59 := 0; fi59 < 1; fi59++ {
        v3 = 550
        if false {
            v60 := []int{0}
            _ = v60
        } else {
            for fi61 := 0; fi61 < 20; fi61++ {
                v9 += 72
                v3 -= (v4 % ((v3) | 1))
                delete(v2, (v9 % (((v9 | v9)) | 1)))
                v62 := GenS2{F0: 42, F1: (v1 / ((uint64(v5)) | 1)), F2: -9223372036854775808}
                _ = fi61
                _ = v62
            }
            var v63 int16 = 42
            if false {
                if (((v1 >> ((v3) & 63)) == (v1 / ((v1) | 1))) || (uint16(v4) == v5)) {
                    v1 = 4294967295
                    v4 = v4
                } else {
                    for fi64 := 0; fi64 < 19; fi64++ {
                        v65 := GenS2{F0: 42, F1: v1, F2: 0}
                        var v66 uint32 = v9
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    switch (fi59 ^ (-(fi59))) {
                    case 9223372036854775807:
                        for fi67 := 0; fi67 < 14; fi67++ {
                            var v68 uint32 = (^(v9))
                            v69 := GenS5{F0: 1, F1: (^(v3))}
                            v3 = v3
                            _ = fi67
                            _ = v68
                            _ = v69
                        }
                    default:
                        if false {
                            var v70 int16 = v63
                            switch "" {
                            case "a":
                                v71 := []int{(fi59 % ((128) | 1)), fi59}
                                _ = v71
                            default:
                                for fi72 := 0; fi72 < 18; fi72++ {
                                    var v73 uint32 = uint32(v3)
                                    v5 = v5
                                    _ = fi72
                                    _ = v73
                                }
                            }
                            v74 := []int{fi59, (-(fi59)), (fi59 << ((v4) & 63)), fi59}
                            _ = v70
                            _ = v74
                        }
                    }
                    v75 := GenS7{F0: 4294967295, F1: 100, F2: (^(v5))}
                    _ = v75
                }
                switch "a" {
                case "\n\t":
                    var v76 string = ""
                    _ = v76
                default:
                    v63 = 1
                }
                if (596 != v1) {
                    var v77 bool = true
                    _ = len(v2)
                    _ = v77
                } else {
                    v3 -= (v4 | v3)
                }
            }
            _ = v63
        }
        _ = fi59
    }
    v78 := GenS4{F0: "\n\t", F1: uint8(v1)}
    switch "" {
    case "hello":
        v79 := []int{}
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v9
}
