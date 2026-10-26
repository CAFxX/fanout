package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 bool; F2 int }
type GenS5 struct { F0 int16; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint; F1 uint8; F2 int }
func genFn0(p0 uint8, p1 int16) uint16 {
    p0 -= (^(p0))
    if (true || false) {
        v5 := []int{}
        for fi6 := 0; fi6 < 1; fi6++ {
            if false {
                if true {
                    p1 = (-(p1))
                    _ = len(v5)
                } else {
                    p1 -= p1
                }
            } else {
                fmt.Println(p1)
            }
            for fi7 := 0; fi7 < 15; fi7++ {
                p1 = p1
                p1 = p1
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi7
            }
            var v8 uint = 18446744073709551615
            _ = fi6
            _ = v8
        }
        _ = v5
    } else {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int8 = 1
            switch fi9 {
            case -854:
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 uint = 18446744073709551615
                    _ = fi11
                    _ = v12
                }
            case 0:
                for fi13 := 0; fi13 < 12; fi13++ {
                    v10 -= -1
                    fmt.Println(fi13)
                    var v14 uint = 492
                    _ = fi13
                    _ = v14
                }
            case 562:
                var v15 uint64 = 0
                _ = v15
            }
            _ = fi9
            _ = v10
        }
        if (3.14159 <= (1.5 / (1.5 + 0.1))) {
            var v16 float32 = float32(-2.25)
            var v17 int16 = (p1 % ((p1) | 1))
            _ = v16
            _ = v17
        }
    }
    fmt.Println(p1, p0)
    var v18 float64 = (-2.25 - 0.0)
    if (255 == 65535) {
        var v19 int32 = -2147483648
        switch -9223372036854775808 {
        case -9223372036854775808:
            fmt.Println(p1, v18, p1)
        case -2147483648:
            var v20 int8 = -127
            _ = v20
        case 127:
            switch 127 {
            case 0:
                v21 := GenS1{F0: 763, F1: (p1 & p1), F2: -1}
                _ = v21
            case -2147483648:
                if false {
                    for fi22 := 0; fi22 < 5; fi22++ {
                        if false {
                            v23 := []int{(^(fi22)), (fi22 - (fi22 >> ((2) & 63))), -2147483648, int(p0)}
                            switch fi22 {
                            case -542:
                                v19 += v19
                            case 4294967295:
                                p1 = (p1 + p1)
                            case 127:
                                if true {
                                    fmt.Println(v18, p1)
                                    switch ("\n\t" + "\n\t") {
                                    case "hello":
                                        var v24 int8 = 42
                                        _ = v24
                                    default:
                                        if len(v23) > 0 {
                                            v23[0] = fi22
                                        }
                                    }
                                    v25 := GenS1{F0: uint(p0), F1: 1000, F2: fi22}
                                    _ = v25
                                }
                            }
                            _ = v23
                        }
                        p0 = (p0 >> ((118) & 63))
                        _ = fi22
                    }
                    if true {
                        switch 100 {
                        case 212:
                            for fi26 := 0; fi26 < 15; fi26++ {
                                var v27 int = -9223372036854775808
                                _ = fi26
                                _ = v27
                            }
                        case 2147483647:
                            v18 += ((v18 * 3.14159) / v18)
                        case -2147483648:
                            v28 := []int{0}
                            _ = v28
                        default:
                            var v29 int8 = 1
                            _ = v29
                        }
                        if true {
                            if (false || (1 >= 42)) {
                                v30 := []int{128, 4294967295, -1}
                                var v31 uint16 = 32767
                                _ = v30
                                _ = v31
                            } else {
                                v18 = (0.0 * v18)
                                var v32 bool = false
                                _ = v32
                            }
                            v18 += 0.0
                            fmt.Println(p1)
                        }
                    }
                    p0 = 1
                }
            }
        default:
            var v33 float32 = 3.14159
            _ = v33
        }
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v18
    return 1
}
func genFn1() string {
    var v34 string = ""
    switch -469 {
    case 144:
        v35 := GenS1{F0: 67, F1: 1, F2: -774}
        _ = v35
    default:
        switch ("\n\t" + "hello") {
        case "foo bar":
            v34 = ("hello" + (("\n\t" + v34) + ((((v34 + v34) + (((v34 + v34) + v34) + "a")) + (v34 + "\n\t")) + (v34 + "a"))))
        case "a":
            v36 := GenS2{F0: 255}
            _ = v36
        }
    }
    _ = v34
    return (("\n\t" + "foo bar") + "\n\t")
}

func main() {
    var v1 int64 = 327
    v2 := []uint32{65535}
    v3 := make(map[int]int)
    var v4 uint = 823
    for fi37 := 0; fi37 < 16; fi37++ {
        if true {
            v4 = (v4 & v4)
        }
        _ = fi37
    }
    fmt.Println(v4, v1)
    var v38 float32 = (float32(((-2.25 + 3.14159) + (1.5 + 3.14159))) + 0.1)
    if false {
        var v39 int = 32767
        _ = v39
    }
    v40 := GenS3{F0: v1}
    v1 = -9223372036854775808
    v1 = (v1 % ((v1) | 1))
    for fi41 := 0; fi41 < 7; fi41++ {
        var v42 int = fi41
        v1 += -1
        if true {
            switch "foo bar" {
            case "hello":
                for fi43 := 0; fi43 < 18; fi43++ {
                    _ = len(v2)
                    _ = fi43
                }
            case "a":
                v42 = 954
            case "":
                for fi44 := 0; fi44 < 11; fi44++ {
                    v4 = (^(v4))
                    switch "hello" {
                    case "hello":
                        if (false || false) {
                            switch fi41 {
                            case 127:
                                if true {
                                    v45 := GenS3{F0: (^(v1))}
                                    v4 += 970
                                    delete(v3, fi41)
                                    _ = v45
                                }
                            case 283:
                                v4 = v4
                            }
                            var v46 float32 = float32(1.5)
                            var v47 float64 = (-2.25 + 0.1)
                            _ = v46
                            _ = v47
                        }
                    case "":
                        v38 = (float32(((float64(((v38 * (0.0 / v38)) / v38)) - (float64((0.0 - 0.0)) / 1e10)) / (float64(v38) / float64(0.1)))) / float32(3.14159))
                    case "\n\t":
                        v38 -= (((0.0 + ((v38 / v38) * 3.14159)) - v38) + v38)
                    }
                    _ = fi44
                }
            }
            v48 := GenS1{F0: 819, F1: -339, F2: (fi41 * (v42 | v42))}
            _ = v48
        }
        _ = fi41
        _ = v42
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        for fi50 := 0; fi50 < 17; fi50++ {
            switch int(v1) {
            case 127:
                switch 747 {
                case 0:
                    fmt.Println(v1)
                default:
                    fmt.Println(fi49, fi49)
                }
            case -309:
                v38 -= (v38 / v38)
                v38 += v38
                var v51 uint32 = 1
                switch "foo bar" {
                case "foo bar":
                    v4 = 390
                }
                _ = v51
            default:
                switch -741 {
                case 127:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
            v52 := GenS4{F0: ("a" + ("hello" + ((("hello" + "foo bar") + "foo bar") + "\n\t"))), F1: (42 > 1), F2: 210}
            _ = fi50
            _ = v52
        }
        if (!((!true) || (((("hello" + ((("hello" + (("a" + ("" + (("\n\t" + ("a" + "hello")) + "hello"))) + ("hello" + "hello"))) + ("" + ("foo bar" + ""))) + "\n\t")) >= "a") || (genFn0(4, 32767) < 42)) || (!("foo bar" != (("a" + "hello") + "hello")))))) {
            var v53 float64 = 0.1
            v4 -= 4294967295
            _ = v53
        } else {
            if false {
                v54 := GenS5{F0: -1, F1: (v4 >> ((v4) & 63)), F2: 127}
                _ = v54
            }
        }
        v55 := GenS6{F0: (v4 / (((^(v4))) | 1)), F1: 128, F2: (fi49 / ((fi49) | 1))}
        _ = fi49
        _ = v55
    }
    if true {
        v1 = 9223372036854775807
        v38 += (((((float32(float64(((v38 + 1.5) - v38))) - v38) / (-2.25 / 1.5)) * (v38 - 0.1)) + v38) * v38)
        var v56 uint8 = 123
        _ = v56
    }
    var v57 int16 = 42
    if false {
        var v58 uint = 500
        _ = v58
    } else {
        var v59 string = "hello"
        _ = v59
    }
    v4 = v4
    if true {
        v2 = append(v2, 4294967295)
        fmt.Println(v4, v4, v1)
        fmt.Println(v4, v38, v1)
    }
    v60 := GenS4{F0: "\n\t", F1: false, F2: 4294967295}
    v61 := GenS4{F0: (("\n\t" + ("hello" + (("foo bar" + "") + ("a" + "\n\t")))) + ((("\n\t" + "foo bar") + ("" + "hello")) + "foo bar")), F1: false, F2: -476}
    var v62 uint64 = 31
    fmt.Println(v38, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v57
    _ = v60
    _ = v61
    _ = v62
}
