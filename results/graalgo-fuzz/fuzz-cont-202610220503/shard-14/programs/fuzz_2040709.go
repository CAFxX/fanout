package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 string; F2 bool }
type GenS7 struct { F0 int32; F1 int; F2 int }
func genFn0(p0 string, p1 uint) uint32 {
    p1 += (p1 % ((p1) | 1))
    fmt.Println(p1)
    var v4 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v4
    return 65535
}
func genFn1(p0 uint16, p1 uint8, p2 uint) int8 {
    var v5 uint = p2
    p2 += v5
    var v6 float32 = ((0.1 + 0.0) * 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 127
}

func main() {
    v1 := make(map[string]string)
    var v2 uint = 2
    var v3 int64 = 42
    if false {
        var v7 int64 = v3
        _ = v7
    } else {
        v2 = v2
        var v8 float64 = -2.25
        for fi9 := 0; fi9 < 16; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 uint32 = uint32(v2)
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        _ = v8
    }
    v2 = v2
    v2 = v2
    if false {
        v12 := []int{32768, 501, -2147483648}
        var v13 float64 = float64(float32(1.5))
        var v14 int8 = 127
        _ = v12
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = 964
        _ = v15
    }
    if true {
        if true {
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 float64 = (float64(-2.25) / 1.5)
                v17 -= v17
                v3 += (v3 >> ((v2) & 63))
                _ = fi16
                _ = v17
            }
            v3 += 9223372036854775807
            v3 -= (v3 + (v3 ^ (v3 % ((v3) | 1))))
        }
        v3 = v3
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 float64 = (((float64(((3.14159 - -2.25) - 0.1)) / (float64(3.14159) / 1e10)) + 3.14159) + (((1.5 / -2.25) + 3.14159) * 0.1))
            for fi20 := 0; fi20 < 5; fi20++ {
                if true {
                    if true {
                        var v21 string = ""
                        v19 -= (((v19 + float64(1e10)) + v19) + v19)
                        var v22 int = (fi18 << ((v2) & 63))
                        v22 = 42
                        v3 -= 9223372036854775807
                        v2 = 255
                        _ = v21
                        _ = v22
                    } else {
                        var v23 uint = 165
                        if (!true) {
                            switch ("hello" + ("\n\t" + "\n\t")) {
                            case "\n\t":
                                if true {
                                    v19 -= 0.0
                                    var v24 int64 = 9223372036854775807
                                    v25 := GenS1{F0: 22, F1: "\n\t", F2: 255}
                                    _ = v24
                                    _ = v25
                                }
                            case "a":
                                v19 = (v19 * v19)
                            }
                            v26 := []int{}
                            _ = v26
                        } else {
                            v3 += (^(v3))
                            switch fi20 {
                            case 32768:
                                switch "a" {
                                case "hello":
                                    var v27 uint8 = 1
                                    var v28 uint64 = 18446744073709551615
                                    v19 = v19
                                    if false {
                                        v3 = 9223372036854775807
                                        var v29 uint8 = v27
                                        _ = v29
                                    } else {
                                        switch "hello" {
                                        case "a":
                                            if (false || ("a" <= ((("\n\t" + "a") + "a") + ("a" + "hello")))) {
                                                var v30 uint32 = 744
                                                _ = v30
                                            } else {
                                                if true {
                                                    v31 := GenS1{F0: 478, F1: "foo bar", F2: (^(v27))}
                                                    if (!(9223372036854775807 > v3)) {
                                                        v32 := []int{-1, fi20, fi18, (fi20 % (((fi20 / (((^(fi20))) | 1))) | 1))}
                                                        _ = v32
                                                    }
                                                    _ = v31
                                                }
                                                delete(v1, ("\n\t" + "a"))
                                                v28 = (v28 << (((^(v23))) & 63))
                                            }
                                        case "foo bar":
                                            fmt.Println(v28, v23, v2)
                                        }
                                        v3 = 1
                                        if ((v3 % (((v3 << ((v2) & 63))) | 1)) < (v3 % (((v3 | 0)) | 1))) {
                                            v33 := GenS1{F0: 1, F1: "a", F2: uint8(v28)}
                                            if (!(false && false)) {
                                                v2 -= 42
                                                var v34 int64 = v3
                                                var v35 uint8 = (^(v27))
                                                _ = v34
                                                _ = v35
                                            }
                                            v28 = (v28 | (v28 | v28))
                                            _ = v33
                                        }
                                    }
                                    _ = v27
                                    _ = v28
                                case "\n\t":
                                    var v36 float32 = 0.1
                                    _ = v36
                                case "foo bar":
                                    fmt.Println(fi18)
                                }
                            default:
                                if (!(!true)) {
                                    var v37 int = fi20
                                    _ = v37
                                }
                            }
                            fmt.Println(v3, v3, fi20)
                        }
                        v1[("foo bar" + "\n\t")] = "foo bar"
                        _ = v23
                    }
                } else {
                    var v38 int64 = int64(v2)
                    v38 -= -552
                    v39 := GenS2{F0: uint16(v2), F1: v2}
                    _ = v38
                    _ = v39
                }
                v40 := GenS3{F0: v3, F1: 1, F2: 42}
                if (!(!true)) {
                    v41 := GenS1{F0: 462, F1: "", F2: 42}
                    switch "foo bar" {
                    case "hello":
                        v19 = v19
                    case "a":
                        v42 := GenS1{F0: 42, F1: "hello", F2: 128}
                        _ = v42
                    }
                    _ = v41
                } else {
                    v19 = v19
                }
                _ = fi20
                _ = v40
            }
            _ = fi18
            _ = v19
        }
    } else {
        for fi43 := 0; fi43 < 3; fi43++ {
            if false {
                if ((-(v3)) > (v3 ^ -9223372036854775808)) {
                    v3 -= v3
                    v3 += -9223372036854775808
                    v3 = 0
                }
                v44 := []int{}
                v2 = (v2 & 0)
                _ = v44
            } else {
                v45 := GenS4{F0: 2}
                v2 -= v2
                for fi46 := 0; fi46 < 10; fi46++ {
                    v47 := GenS1{F0: 255, F1: "", F2: 2}
                    _ = fi46
                    _ = v47
                }
                _ = v45
            }
            _ = fi43
        }
        switch "" {
        case "hello":
            v48 := []int{4294967295, 42, -1}
            _ = v48
        default:
            switch "a" {
            case "foo bar":
                var v49 bool = true
                _ = v49
            case "a":
                var v50 uint32 = 2147483647
                _ = v50
            case "":
                switch 42 {
                case 4294967295:
                    for fi51 := 0; fi51 < 13; fi51++ {
                        v52 := []int{}
                        fmt.Println(fi51, v3, v2)
                        _ = fi51
                        _ = v52
                    }
                }
            default:
                for fi53 := 0; fi53 < 6; fi53++ {
                    v3 = -1
                    delete(v1, "hello")
                    switch fi53 {
                    case -23:
                        fmt.Println(v3, v2)
                    case 2147483647:
                        v3 += 508
                    case 300:
                        var v54 uint = 2
                        _ = v54
                    default:
                        fmt.Println(v3)
                    }
                    _ = fi53
                }
            }
        }
    }
    if ((v3 - (^(v3))) == (v3 << ((65535) & 63))) {
        v55 := GenS5{F0: -2147483648}
        v2 = v2
        switch "hello" {
        case "hello":
            if (((("foo bar" == "hello") || false) && (!true)) && false) {
                switch ("" + "\n\t") {
                case "\n\t":
                    var v56 float64 = 1.5
                    _ = v56
                case "foo bar":
                    if true {
                        v3 = v3
                        for fi57 := 0; fi57 < 9; fi57++ {
                            switch 32767 {
                            case 32767:
                                v58 := GenS6{F0: 109, F1: ("foo bar" + "foo bar"), F2: true}
                                _ = v58
                            }
                            _ = fi57
                        }
                    }
                }
            } else {
                if (0.0 <= (float32(float64((0.1 + 0.1))) - 0.0)) {
                    v3 += v3
                    if ((v3 << ((v2) & 63)) != v3) {
                        var v59 uint16 = 1
                        var v60 uint32 = 65535
                        fmt.Println(v2, v3)
                        _ = v59
                        _ = v60
                    } else {
                        fmt.Println(v2)
                        var v61 uint16 = 32767
                        _ = v61
                    }
                    fmt.Println(v3, v3)
                }
            }
        case "a":
            switch -9223372036854775808 {
            case 0:
                v3 += 1
            case 4294967295:
                var v62 float32 = (-2.25 / float32((((0.1 + 0.1) / 1e10) + 1.5)))
                _ = v62
            }
        case "\n\t":
            var v63 uint16 = 986
            _ = v63
        default:
            var v64 int32 = -909
            _ = v64
        }
        _ = v55
    }
    var v65 uint32 = 2147483648
    var v66 string = "a"
    for fi67 := 0; fi67 < 1; fi67++ {
        v2 = (v2 / ((v2) | 1))
        _ = fi67
    }
    v68 := GenS7{F0: 1, F1: 42, F2: 4294967295}
    v3 += v3
    var v69 uint16 = 1
    fmt.Println(v65, v69, v69)
    fmt.Println(v65, v65, v65)
    var v70 bool = false
    for fi71 := 0; fi71 < 9; fi71++ {
        for fi72 := 0; fi72 < 5; fi72++ {
            v66 = ""
            v2 -= 805
            _ = fi72
        }
        _ = fi71
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v66
    _ = v68
    _ = v69
    _ = v70
}
