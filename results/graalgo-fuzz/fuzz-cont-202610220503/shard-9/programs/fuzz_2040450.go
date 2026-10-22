package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int64 }
func genFn0(p0 bool) uint8 {
    v7 := GenS1{F0: -128}
    fmt.Println(p0)
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 161
}
func genFn1(p0 uint64, p1 uint16, p2 uint16) uint8 {
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int8 = -128
        fmt.Println(v9, p0)
        var v10 int64 = 1
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{307, 128, 4294967295}
    var v12 float64 = ((1.5 * 0.0) - 1e10)
    switch 2147483647 {
    case 1:
        p1 -= (p1 >> ((42) & 63))
    case 32768:
        var v13 int32 = -2147483648
        _ = v13
    case 127:
        p2 -= (p1 - p2)
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS2{F0: -128, F1: -32}
        p1 -= 32767
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 127
    var v3 float64 = 3.14159
    var v4 uint64 = 1
    v5 := []string{"hello", ("hello" + (("a" + (("a" + "hello") + ("" + "hello"))) + ("" + ("hello" + "foo bar")))), "hello", "a"}
    v6 := make(map[int]bool)
    for fi16 := 0; fi16 < 5; fi16++ {
        switch (fi16 & 819) {
        case 128:
            var v17 uint32 = 786
            _ = v17
        case 2147483647:
            v4 -= v4
        case -2147483648:
            for fi18 := 0; fi18 < 20; fi18++ {
                switch ("a" + "foo bar") {
                case "a":
                    v19 := GenS1{F0: (^(v2))}
                    _ = v19
                case "hello":
                    fmt.Println(v3, v3)
                }
                fmt.Println(v4, fi16, v4)
                var v20 int16 = 0
                _ = fi18
                _ = v20
            }
        default:
            v3 = (0.1 * v3)
        }
        var v21 string = "a"
        _ = fi16
        _ = v21
    }
    switch 613 {
    case -400:
        if false {
            fmt.Println(v4, v2)
            fmt.Println(v4, v4, v2)
            switch 36 {
            case -594:
                v5 = append(v5, ("foo bar" + "\n\t"))
            case -1:
                fmt.Println(v3, v4)
            default:
                var v22 int = -9223372036854775808
                _ = v22
            }
        }
    case 2147483647:
        fmt.Println(v4)
    default:
        for fi23 := 0; fi23 < 8; fi23++ {
            _ = len(v5)
            _ = fi23
        }
    }
    v5 = append(v5, "\n\t")
    switch "hello" {
    case "a":
        if true {
            var v24 uint64 = (v4 * (v4 << ((65535) & 63)))
            v2 -= (v2 * v2)
            _ = v24
        } else {
            switch 400 {
            case 756:
                for fi25 := 0; fi25 < 9; fi25++ {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v3 = -2.25
                        fmt.Println(v3, v3)
                        if (true && (1.5 >= v3)) {
                            var v27 uint32 = 54
                            switch "\n\t" {
                            case "":
                                v28 := GenS2{F0: 100, F1: v2}
                                _ = v28
                            case "\n\t":
                                v29 := GenS3{F0: int(v27), F1: -566}
                                _ = v29
                            case "a":
                                var v30 uint64 = (^(v4))
                                _ = v30
                            default:
                                var v31 int32 = -1
                                _ = v31
                            }
                            var v32 float32 = 0.1
                            _ = v27
                            _ = v32
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if true {
                                v33 := GenS2{F0: v2, F1: 1}
                                _ = v33
                            }
                        }
                        _ = fi26
                    }
                    v34 := GenS2{F0: v2, F1: (v2 + (v2 >> ((0) & 63)))}
                    v2 = (v2 ^ (-(v2)))
                    _ = fi25
                    _ = v34
                }
            case 127:
                v35 := GenS1{F0: (v2 ^ v2)}
                _ = v35
            }
        }
    case "hello":
        var v36 uint16 = 505
        _ = v36
    case "":
        v37 := GenS4{F0: v4, F1: genFn0(true)}
        _ = v37
    default:
        v38 := v1[1]
        _ = v38
    }
    if true {
        switch -206 {
        case 42:
            switch "a" {
            case "hello":
                var v39 uint = 2
                _ = v39
            case "foo bar":
                if true {
                    switch "" {
                    case "a":
                        v4 = 42
                    case "":
                        switch 42 {
                        case -628:
                            v40 := GenS3{F0: 2147483647, F1: 42}
                            _ = v40
                        default:
                            switch -141 {
                            case 127:
                                v2 -= (v2 % ((v2) | 1))
                            case 9223372036854775807:
                                if (!false) {
                                    if true {
                                        v2 += (v2 ^ 0)
                                        if (3.14159 > float32(v3)) {
                                            switch 128 {
                                            case 2147483647:
                                                delete(v6, -1)
                                            case 684:
                                                v4 = (v4 % ((v4) | 1))
                                            default:
                                                v41 := GenS1{F0: (v2 >> ((479) & 63))}
                                                _ = v41
                                            }
                                            fmt.Println(v2, v2, v2)
                                        }
                                    } else {
                                        for fi42 := 0; fi42 < 3; fi42++ {
                                            v4 = 1
                                            _ = fi42
                                        }
                                        v43 := GenS5{F0: 2147483647}
                                        for fi44 := 0; fi44 < 8; fi44++ {
                                            v45 := GenS3{F0: (-(fi44)), F1: -1}
                                            _ = fi44
                                            _ = v45
                                        }
                                        _ = v43
                                    }
                                } else {
                                    v1[2147483647] = (794 <= 0)
                                }
                            }
                        }
                    case "foo bar":
                        v4 = (v4 / ((246) | 1))
                    default:
                        _ = len(v5)
                    }
                    if (!(!(0 > 32767))) {
                        v46 := v1[0]
                        v5 = append(v5, "foo bar")
                        _ = v46
                    }
                }
            case "\n\t":
                if false {
                    switch "\n\t" {
                    case "\n\t":
                        delete(v1, 837)
                    }
                } else {
                    v2 = v2
                    _ = len(v5)
                }
            default:
                if true {
                    for fi47 := 0; fi47 < 18; fi47++ {
                        v48 := GenS4{F0: v4, F1: genFn0(true)}
                        var v49 int64 = -693
                        for fi50 := 0; fi50 < 15; fi50++ {
                            v51 := GenS5{F0: 406}
                            var v52 uint16 = 65535
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    switch -305 {
                    case 0:
                        v4 = (v4 ^ (v4 + (v4 + v4)))
                    case 428:
                        if (603 < 1) {
                            for fi53 := 0; fi53 < 20; fi53++ {
                                for fi54 := 0; fi54 < 14; fi54++ {
                                    v5 = append(v5, ("hello" + ""))
                                    _ = fi54
                                }
                                _ = fi53
                            }
                        }
                    }
                    v55 := GenS4{F0: (^(v4)), F1: 42}
                    _ = v55
                }
            }
        case 492:
            if true {
                v2 -= 0
            } else {
                v56 := v6[0]
                _ = v56
            }
        }
        var v57 float64 = 3.14159
        if false {
            delete(v1, 1)
            var v58 bool = ((false && true) || true)
            _ = v58
        } else {
            v57 -= v57
            v59 := GenS6{F0: (v2 ^ (v2 - v2)), F1: 265}
            _ = v59
        }
        _ = v57
    }
    if false {
        delete(v6, 4294967295)
    }
    fmt.Println(v3, v2, v3)
    if false {
        fmt.Println(v2)
    } else {
        switch ("foo bar" + "foo bar") {
        case "hello":
            v3 += v3
        }
        v5 = append(v5, "\n\t")
        v60 := GenS6{F0: v2, F1: 0}
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
