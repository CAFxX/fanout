package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint16; F2 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int8 }
func genFn0(p0 bool, p1 uint16) uint16 {
    p1 -= (p1 - (p1 & (p1 % (((p1 * p1)) | 1))))
    v5 := GenS1{F0: (p1 * p1), F1: 4294967295}
    p1 += p1
    v6 := GenS1{F0: 638, F1: -474}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 168
}

func main() {
    v1 := []string{""}
    var v2 int8 = 100
    var v3 bool = (0 <= 255)
    var v4 float64 = 1.5
    var v7 int8 = v2
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint64 = 717
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 0}
    for fi11 := 0; fi11 < 15; fi11++ {
        if v3 {
            for fi12 := 0; fi12 < 13; fi12++ {
                _ = len(v1)
                v13 := GenS1{F0: 940, F1: fi11}
                v14 := GenS3{F0: 1000, F1: 135}
                if (2 > 0) {
                    fmt.Println(fi11, v7)
                    switch "foo bar" {
                    case "foo bar":
                        v2 -= (v7 ^ v2)
                        v2 -= -42
                    }
                }
                switch "\n\t" {
                case "hello":
                    v15 := GenS2{F0: 337}
                    _ = v15
                default:
                    v3 = (!(true || (((-2.25 - 1e10) / float32(v4)) != (-2.25 * ((0.0 + 0.0) - (1e10 + (0.0 - (3.14159 / ((float32(0.0) + (-2.25 / -2.25)) / 3.14159)))))))))
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            if true {
                v2 = v7
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                if true {
                    v16 := GenS4{F0: 42, F1: -1}
                    if v3 {
                        var v17 int16 = 1
                        if (((float32(0.0) + 1.5) + float32(float64(0.0))) <= 0.0) {
                            var v18 string = ("foo bar" + ("\n\t" + ""))
                            var v19 int32 = 0
                            if (v3 && false) {
                                v20 := GenS3{F0: int16(v7), F1: 2}
                                _ = v20
                            }
                            _ = v18
                            _ = v19
                        } else {
                            v21 := GenS4{F0: -1, F1: -17}
                            fmt.Println(v7, v2)
                            var v22 int = 819
                            _ = v21
                            _ = v22
                        }
                        fmt.Println(v17, fi11)
                        _ = v17
                    }
                    v23 := GenS5{F0: 0, F1: 2, F2: (1 < 9)}
                    _ = v16
                    _ = v23
                } else {
                    var v24 uint8 = 1
                    v3 = (v2 >= (v2 / ((v7) | 1)))
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(v2, fi11)
                v26 := GenS6{F0: "\n\t"}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi25
                _ = v26
            }
        } else {
            if v3 {
                v27 := GenS7{F0: (872 > 18446744073709551615)}
                v28 := GenS8{F0: v2}
                _ = v27
                _ = v28
            }
        }
        var v29 int8 = (v7 >> ((4294967295) & 63))
        var v30 int8 = -79
        _ = fi11
        _ = v29
        _ = v30
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v31 uint = 2
    for fi32 := 0; fi32 < 20; fi32++ {
        for fi33 := 0; fi33 < 3; fi33++ {
            v7 = (v2 & 0)
            fmt.Println(fi33)
            _ = fi33
        }
        _ = fi32
    }
    if true {
        v1 = append(v1, "hello")
        if (319 >= 4294967295) {
            v2 = (v2 % ((-128) | 1))
            if v3 {
                if true {
                    if (42 <= 279) {
                        for fi34 := 0; fi34 < 5; fi34++ {
                            var v35 int64 = 1
                            for fi36 := 0; fi36 < 7; fi36++ {
                                var v37 uint64 = 18446744073709551615
                                v35 = -120
                                _ = fi36
                                _ = v37
                            }
                            _ = fi34
                            _ = v35
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if (v3 && true) {
                        var v38 int32 = -1
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        _ = v38
                    }
                    var v39 int16 = 112
                    _ = v39
                }
            }
        }
        switch 32768 {
        case 32767:
            if ((!(42 != genFn0(((("a" > ("" + "\n\t")) || v3) || v3), 647))) && true) {
                _ = len(v1)
            } else {
                for fi40 := 0; fi40 < 16; fi40++ {
                    v3 = v3
                    switch "a" {
                    case "\n\t":
                        if (v3 || false) {
                            v41 := GenS6{F0: ("hello" + "foo bar")}
                            _ = v41
                        } else {
                            for fi42 := 0; fi42 < 1; fi42++ {
                                if v3 {
                                    v7 = v7
                                    fmt.Println(fi40, fi42)
                                    var v43 uint16 = 42
                                    _ = v43
                                } else {
                                    var v44 int64 = 42
                                    _ = v44
                                }
                                var v45 uint8 = uint8(v31)
                                v46 := GenS9{F0: v7, F1: 42}
                                _ = fi42
                                _ = v45
                                _ = v46
                            }
                            fmt.Println(v3, v7, v3)
                            for fi47 := 0; fi47 < 1; fi47++ {
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                var v48 uint32 = 0
                                _ = fi47
                                _ = v48
                            }
                        }
                    case "foo bar":
                        var v49 int16 = 1
                        _ = v49
                    case "":
                        if (!v3) {
                            for fi50 := 0; fi50 < 7; fi50++ {
                                if v3 {
                                    _ = len(v1)
                                    v1 = append(v1, ("foo bar" + "\n\t"))
                                    v31 += (v31 * 961)
                                }
                                _ = fi50
                            }
                            switch (fi40 >> ((353) & 63)) {
                            case 127:
                                var v51 float32 = float32(((v4 + v4) / float64(0.1)))
                                _ = v51
                            case -2147483648:
                                v3 = false
                            }
                        }
                    default:
                        var v52 uint = 255
                        _ = v52
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi40
                }
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            }
        case 61:
            v3 = true
        default:
            if (!(1 == 864)) {
                for fi53 := 0; fi53 < 5; fi53++ {
                    v4 = v4
                    _ = fi53
                }
                if (false || v3) {
                    v54 := GenS4{F0: 1, F1: 154}
                    var v55 int8 = v2
                    _ = v54
                    _ = v55
                } else {
                    v2 = v7
                    switch (("foo bar" + "hello") + "hello") {
                    case "hello":
                        v56 := GenS10{F0: v3}
                        _ = v56
                    }
                    switch 486 {
                    case 2147483647:
                        switch ("" + "") {
                        case "hello":
                            fmt.Println(v3)
                        case "foo bar":
                            v2 = v7
                        case "a":
                            v2 = 35
                        }
                    case 32768:
                        v3 = true
                    default:
                        v57 := GenS1{F0: 158, F1: 9223372036854775807}
                        _ = v57
                    }
                }
                v31 += 18446744073709551615
            } else {
                for fi58 := 0; fi58 < 13; fi58++ {
                    switch -407 {
                    case 274:
                        v1 = append(v1, ("\n\t" + "\n\t"))
                    case -458:
                        if v3 {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v59 := GenS11{F0: 2}
                            _ = v59
                        }
                    default:
                        for fi60 := 0; fi60 < 12; fi60++ {
                            if (v4 < v4) {
                                if true {
                                    var v61 int32 = 0
                                    v2 += (v7 & 127)
                                    _ = v61
                                }
                            }
                            var v62 uint64 = 42
                            v63 := GenS8{F0: (v7 * 1)}
                            _ = fi60
                            _ = v62
                            _ = v63
                        }
                    }
                    var v64 int32 = 231
                    _ = fi58
                    _ = v64
                }
                v65 := GenS12{F0: v7}
                _ = v65
            }
        }
    }
    v66 := GenS12{F0: v7}
    fmt.Println(v4)
    v31 = v31
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v66
    _ = v7
}
