package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int }
type GenS2 struct { F0 bool; F1 int; F2 string }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool; F1 uint16 }
func genFn0(p0 bool, p1 string) string {
    var v4 int8 = -128
    var v5 int32 = 949
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{fi6, (fi6 / (((^(fi6))) | 1)), (fi6 << ((0) & 63)), (-(fi6))}
        p1 = "hello"
        for fi8 := 0; fi8 < 1; fi8++ {
            switch p1 {
            case "\n\t":
                switch ((p1 + (((p1 + p1) + p1) + "")) + "") {
                case "a":
                    if len(v7) > 0 {
                        v7[0] = fi6
                    }
                case "foo bar":
                    var v9 bool = (p0 && (521 <= 2))
                    _ = v9
                default:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            default:
                p1 = "foo bar"
            }
            v10 := GenS1{F0: (^(fi6)), F1: -554, F2: fi6}
            _ = fi8
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    v11 := GenS2{F0: false, F1: 4294967295, F2: p1}
    _ = p0
    _ = p1
    _ = v11
    _ = v4
    _ = v5
    return ((p1 + (("a" + "foo bar") + p1)) + p1)
}
func genFn1(p0 bool, p1 string) int16 {
    fmt.Println(p1)
    p0 = ((2 <= 119) && p0)
    if p0 {
        var v12 int64 = 457
        _ = v12
    } else {
        if (0.1 == 1.5) {
            switch 278 {
            case -2147483648:
                switch ((p1 + "\n\t") + "\n\t") {
                case "foo bar":
                    for fi13 := 0; fi13 < 15; fi13++ {
                        for fi14 := 0; fi14 < 19; fi14++ {
                            v15 := []int{(fi13 ^ fi13), -643}
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 1; fi16++ {
                            fmt.Println(fi16)
                            _ = fi16
                        }
                        _ = fi13
                    }
                case "a":
                    p0 = (1 > 127)
                case "\n\t":
                    fmt.Println(p0)
                default:
                    var v17 int16 = -32768
                    _ = v17
                }
            }
        } else {
            v18 := []int{127}
            p0 = (!true)
            _ = v18
        }
        v19 := GenS3{F0: 2, F1: 42}
        v20 := GenS4{F0: 127}
        _ = v19
        _ = v20
    }
    switch p1 {
    case "a":
        var v21 bool = ((1 != 32767) && (!(-128 > 1)))
        _ = v21
    case "\n\t":
        p1 = "hello"
    default:
        if true {
            fmt.Println(p1, p1)
            if p0 {
                v22 := GenS3{F0: 18446744073709551615, F1: 85}
                _ = v22
            }
            var v23 int = -1
            _ = v23
        } else {
            switch 9223372036854775807 {
            case -1:
                if (!p0) {
                    var v24 float32 = 0.1
                    v25 := []int{-233, -9223372036854775808}
                    _ = v24
                    _ = v25
                }
            case -9223372036854775808:
                p1 = p1
            case 4294967295:
                v26 := []int{32768, 760, 32767, 127}
                _ = v26
            }
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 uint32 = 551
                _ = fi27
                _ = v28
            }
            if (1 != -604) {
                var v29 int16 = -32768
                for fi30 := 0; fi30 < 19; fi30++ {
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 int32 = -243
                        v33 := GenS5{F0: p1}
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    v34 := GenS6{F0: 2147483647, F1: -78}
                    _ = fi30
                    _ = v34
                }
                _ = v29
            } else {
                switch p1 {
                case "":
                    fmt.Println(p1)
                }
                for fi35 := 0; fi35 < 9; fi35++ {
                    if p0 {
                        var v36 int64 = 9223372036854775807
                        _ = v36
                    } else {
                        if (!(!false)) {
                            for fi37 := 0; fi37 < 20; fi37++ {
                                p0 = (p1 < ("hello" + p1))
                                _ = fi37
                            }
                        } else {
                            fmt.Println(p0, p0, p0)
                            for fi38 := 0; fi38 < 16; fi38++ {
                                var v39 uint64 = 42
                                v40 := []int{-64, 42, fi38, fi38}
                                switch ("" + "\n\t") {
                                case "hello":
                                    fmt.Println(fi35)
                                case "":
                                    if len(v40) > 0 {
                                        v40[0] = fi35
                                    }
                                default:
                                    v41 := GenS7{F0: -1}
                                    _ = v41
                                }
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                            if (!p0) {
                                var v42 uint64 = 42
                                _ = v42
                            }
                        }
                        p1 = "hello"
                    }
                    _ = fi35
                }
                if true {
                    for fi43 := 0; fi43 < 15; fi43++ {
                        v44 := []int{}
                        var v45 uint16 = 270
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                } else {
                    for fi46 := 0; fi46 < 13; fi46++ {
                        p1 = (p1 + "foo bar")
                        fmt.Println(p1, p0)
                        switch fi46 {
                        case -130:
                            v47 := GenS4{F0: 255}
                            _ = v47
                        case 1:
                            if (!false) {
                                v48 := GenS6{F0: -1, F1: 127}
                                _ = v48
                            } else {
                                p1 = "foo bar"
                            }
                        default:
                            fmt.Println(p1)
                        }
                        _ = fi46
                    }
                    if (!p0) {
                        for fi49 := 0; fi49 < 3; fi49++ {
                            for fi50 := 0; fi50 < 3; fi50++ {
                                v51 := []int{fi50, (fi49 | (-(fi49)))}
                                _ = fi50
                                _ = v51
                            }
                            v52 := []int{-533, fi49, fi49}
                            _ = fi49
                            _ = v52
                        }
                        var v53 uint = 65535
                        _ = v53
                    }
                    var v54 int16 = -662
                    _ = v54
                }
            }
        }
    }
    switch 127 {
    case 32768:
        switch 9223372036854775807 {
        case -9223372036854775808:
            v55 := GenS8{F0: -618}
            _ = v55
        case 42:
            var v56 uint = 278
            _ = v56
        case 32768:
            p0 = false
        default:
            var v57 int16 = -32768
            _ = v57
        }
    case 42:
        v58 := GenS9{F0: (739 < 4294967295), F1: 42}
        _ = v58
    case 0:
        p0 = p0
    }
    _ = p0
    _ = p1
    return -32768
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{""}
    var v3 uint64 = 18446744073709551615
    switch 4294967295 {
    case 4294967295:
        _ = len(v1)
    case 127:
        var v59 int = 128
        _ = v59
    case -9223372036854775808:
        _ = len(v2)
    }
    v3 = (v3 << ((2) & 63))
    switch 32768 {
    case 9223372036854775807:
        for fi60 := 0; fi60 < 8; fi60++ {
            _ = len(v1)
            _ = fi60
        }
    case 127:
        fmt.Println(v3)
    case 32768:
        v3 = v3
    }
    fmt.Println(v3)
    for fi61 := 0; fi61 < 2; fi61++ {
        for fi62 := 0; fi62 < 13; fi62++ {
            fmt.Println(fi61, v3, fi61)
            _ = len(v2)
            _ = fi62
        }
        var v63 uint64 = v3
        var v64 int8 = 127
        _ = fi61
        _ = v63
        _ = v64
    }
    v2 = append(v2, ("hello" + ("hello" + "\n\t")))
    v3 += v3
    v3 -= (v3 | (v3 - (^(v3))))
    v3 += (v3 % ((2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
