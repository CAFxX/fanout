package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 bool }
type GenS2 struct { F0 bool; F1 int16; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint64; F2 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 uint8; F1 uint64 }
type GenS9 struct { F0 uint8; F1 string }
type GenS10 struct { F0 int32; F1 bool }
type GenS11 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint8, p1 int) int32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        if false {
            var v6 int16 = -32768
            var v7 int = fi5
            _ = v6
            _ = v7
        }
        _ = fi5
    }
    v8 := GenS1{F0: ("\n\t" + ("\n\t" + "")), F1: true, F2: true}
    v9 := GenS2{F0: true, F1: 0, F2: false}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint64) uint64 {
    var v10 int16 = 0
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint16 = 0
        if ((("" + ("" + (("a" + "a") + ("a" + ("a" + "hello"))))) + "foo bar") < "\n\t") {
            var v13 uint = 0
            var v14 int32 = 0
            _ = v13
            _ = v14
        } else {
            v15 := []int{int(p0), (fi11 << ((674) & 63))}
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            p0 = (p0 | uint64(v12))
            v17 := []int{(-(fi11)), fi16}
            if (1 >= 9223372036854775807) {
                if len(v17) > 0 {
                    v17[0] = fi16
                }
                if true {
                    switch "a" {
                    case "a":
                        if (false && (0 == 18446744073709551615)) {
                            fmt.Println(v10, p0, p0)
                            fmt.Println(fi11, p0, p0)
                        }
                    }
                    switch 811 {
                    case -256:
                        var v18 string = "foo bar"
                        _ = v18
                    case 0:
                        v19 := GenS3{F0: 88}
                        _ = v19
                    default:
                        switch fi16 {
                        case -9223372036854775808:
                            if ((float32(float64(0.1)) == (1e10 / (1e10 - ((0.1 * 1.5) - -2.25)))) && (1 <= -802)) {
                                p0 -= (p0 - 4294967295)
                                var v20 uint = 549
                                var v21 uint = uint(v12)
                                if false {
                                    var v22 string = ("foo bar" + ("" + ""))
                                    _ = v22
                                } else {
                                    v17 = append(v17, fi16)
                                    v10 = v10
                                    switch fi11 {
                                    case 230:
                                        switch (fi16 % ((4294967295) | 1)) {
                                        case 32767:
                                            var v23 float64 = 0.1
                                            _ = v23
                                        case 32768:
                                            v20 -= 459
                                        default:
                                            if (p0 > p0) {
                                                v21 += 871
                                                var v24 int32 = -1
                                                _ = v24
                                            } else {
                                                if true {
                                                    fmt.Println(v10)
                                                    if len(v17) > 0 {
                                                        v17[0] = (fi16 & fi11)
                                                    }
                                                    var v25 string = (("\n\t" + "hello") + "hello")
                                                    _ = v25
                                                }
                                                v12 = (v12 << (((v21 / ((1) | 1))) & 63))
                                            }
                                        }
                                    case 901:
                                        var v26 int32 = -804
                                        v12 -= v12
                                        var v27 uint8 = 2
                                        _ = v26
                                        _ = v27
                                    }
                                    v12 = v12
                                    switch "" {
                                    case "foo bar":
                                        var v28 uint = 42
                                        _ = v28
                                    default:
                                        _ = len(v17)
                                    }
                                    var v29 bool = true
                                    p0 = uint64(v20)
                                    _ = v29
                                }
                                if len(v17) > 0 {
                                    v17[0] = (fi16 + (fi11 | fi11))
                                }
                                _ = v20
                                _ = v21
                            }
                        case -2147483648:
                            if (!(2147483648 >= 2)) {
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    p0 += p0
                                case "foo bar":
                                    if false {
                                        v30 := GenS4{F0: fi16, F1: 607, F2: 0}
                                        _ = len(v17)
                                        var v31 int16 = -1
                                        _ = v30
                                        _ = v31
                                    }
                                case "":
                                    fmt.Println(fi16, fi11)
                                }
                            } else {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            }
                        default:
                            fmt.Println(p0, fi11, fi16)
                        }
                    }
                } else {
                    switch "a" {
                    case "a":
                        if len(v17) > 0 {
                            v17[0] = -842
                        }
                    case "foo bar":
                        var v32 int32 = genFn0(188, fi11)
                        _ = v32
                    default:
                        var v33 uint16 = 42
                        var v34 int64 = -703
                        _ = v33
                        _ = v34
                    }
                    fmt.Println(p0)
                    v10 = 42
                }
                fmt.Println(p0)
            }
            _ = fi16
            _ = v17
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 uint8 = 255
    v2 := []string{("a" + (("a" + ("foo bar" + "")) + ("a" + "a"))), ""}
    v3 := make(map[uint32]int)
    var v4 uint64 = 515
    for fi35 := 0; fi35 < 4; fi35++ {
        for fi36 := 0; fi36 < 12; fi36++ {
            switch fi36 {
            case -506:
                v37 := GenS1{F0: ("a" + "foo bar"), F1: (true && (true || (fi35 == (-(fi36))))), F2: false}
                _ = v37
            default:
                v38 := GenS5{F0: 0}
                switch ("" + "foo bar") {
                case "a":
                    v1 = (^(v1))
                }
                _ = v38
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if false {
                _ = len(v3)
                v39 := GenS6{F0: 0, F1: (2147483647 > 140)}
                _ = v39
            } else {
                var v40 int64 = 748
                _ = v40
            }
            _ = fi36
        }
        _ = fi35
    }
    v1 = (v1 | v1)
    var v41 uint8 = (v1 % ((v1) | 1))
    v1 -= (v1 + (v41 % ((2) | 1)))
    v42 := GenS7{F0: (!true), F1: -9223372036854775808}
    if (true && false) {
        v1 = (v1 % ((v1) | 1))
        switch "a" {
        case "\n\t":
            v1 = uint8(v4)
        case "a":
            v41 = v1
        default:
            for fi43 := 0; fi43 < 18; fi43++ {
                v4 += v4
                if false {
                    fmt.Println(v1)
                    fmt.Println(v4, v1, v4)
                    v4 -= v4
                } else {
                    v4 = genFn1((v4 % ((genFn1(2)) | 1)))
                    switch (-(fi43)) {
                    case -9223372036854775808:
                        _ = len(v2)
                    case 1:
                        for fi44 := 0; fi44 < 1; fi44++ {
                            v45 := GenS6{F0: -2147483648, F1: false}
                            if true {
                                _ = len(v3)
                                if (false && true) {
                                    fmt.Println(v41, v4)
                                    fmt.Println(fi44, v41, v41)
                                    var v46 float32 = 1.5
                                    _ = v46
                                } else {
                                    if false {
                                        switch ("a" + "hello") {
                                        case "foo bar":
                                            switch "hello" {
                                            case "":
                                                switch fi44 {
                                                case 9223372036854775807:
                                                    var v47 float32 = ((1.5 * 0.0) / 1.0)
                                                    _ = v47
                                                default:
                                                    v48 := GenS8{F0: 42, F1: (^(v4))}
                                                    _ = v48
                                                }
                                            case "a":
                                                v4 = 18446744073709551615
                                            case "hello":
                                                fmt.Println(v4, v1, fi44)
                                            default:
                                                delete(v3, 65535)
                                            }
                                        case "hello":
                                            if ((fi43 << ((1) & 63)) <= (fi43 + 4294967295)) {
                                                fmt.Println(fi44, fi44, fi44)
                                                fmt.Println(fi44, fi43, fi43)
                                                switch "\n\t" {
                                                case "a":
                                                    var v49 int16 = 955
                                                    _ = v49
                                                case "":
                                                    v1 -= v41
                                                default:
                                                    var v50 uint = 0
                                                    _ = v50
                                                }
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        default:
                                            v1 = (v41 * 127)
                                        }
                                        fmt.Println(v4, v1)
                                        var v51 uint8 = (v1 + (v41 - (v1 | v41)))
                                        _ = v51
                                    }
                                }
                            }
                            var v52 int8 = 42
                            _ = fi44
                            _ = v45
                            _ = v52
                        }
                    case -596:
                        v53 := GenS9{F0: v41, F1: (((("hello" + "") + "hello") + "\n\t") + "a")}
                        _ = v53
                    }
                }
                var v54 int16 = 1000
                _ = fi43
                _ = v54
            }
        }
    } else {
        v55 := GenS10{F0: 42, F1: false}
        v41 = (^(v41))
        _ = v55
    }
    fmt.Println(v41, v1, v41)
    v2 = append(v2, "")
    var v56 float32 = 1e10
    var v57 int = 867
    v3[0] = (-(v57))
    var v58 bool = (-2147483648 <= int32(v4))
    if v58 {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 += v1
        v59 := GenS11{F0: v41, F1: (v4 & (v4 >> ((749) & 63)))}
        _ = v59
    } else {
        var v60 int = (v57 ^ (v57 << ((4294967295) & 63)))
        _ = v60
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v56
    _ = v57
    _ = v58
}
