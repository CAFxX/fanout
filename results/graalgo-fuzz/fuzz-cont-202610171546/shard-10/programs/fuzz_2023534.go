package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint32; F1 string; F2 uint }
type GenS4 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS5 struct { F0 string; F1 uint }
func genFn0(p0 uint64, p1 bool, p2 string) uint64 {
    v4 := GenS1{F0: 32767, F1: false}
    p1 = ((^(p0)) < p0)
    v5 := []int{818}
    switch "\n\t" {
    case "a":
        for fi6 := 0; fi6 < 17; fi6++ {
            p0 = 0
            for fi7 := 0; fi7 < 3; fi7++ {
                fmt.Println(fi6, p0, p2)
                p1 = (1.5 < 0.1)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi7
            }
            _ = fi6
        }
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            p2 = p2
            p2 = p2
            if p1 {
                v10 := GenS2{F0: 534, F1: 44, F2: 0}
                var v11 int32 = 0
                _ = v10
                _ = v11
            }
            var v12 uint16 = 357
            switch (p2 + p2) {
            case "foo bar":
                p0 += 477
            case "hello":
                fmt.Println(p1, p2, p2)
            case "":
                var v13 uint32 = 605
                _ = v13
            default:
                switch (fi9 + fi8) {
                case 0:
                    _ = len(v5)
                case -847:
                    var v14 uint = 65535
                    _ = v14
                default:
                    p0 += p0
                }
            }
            _ = fi9
            _ = v12
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p0 | p0)
}
func genFn1(p0 string, p1 uint) uint64 {
    if false {
        fmt.Println(p0, p1)
        if false {
            if (true && true) {
                v15 := []int{4294967295, 1, 0, -9223372036854775808}
                var v16 uint = p1
                p1 -= (p1 | 398)
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 6; fi17++ {
                    p0 = "foo bar"
                    _ = fi17
                }
            }
        } else {
            if true {
                p0 = ("a" + ((p0 + (("a" + ("a" + "hello")) + ("" + "\n\t"))) + p0))
            } else {
                var v18 int = 1
                var v19 float64 = (((-2.25 / 1.0) / float64(3.14159)) + 1.5)
                _ = v18
                _ = v19
            }
        }
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := GenS1{F0: int16(p1), F1: false}
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    return 953
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 340
    v3 := make(map[uint32]bool)
    var v22 uint8 = 3
    v23 := []int{127, 4294967295, 940, -1}
    v24 := GenS3{F0: 65535, F1: "hello", F2: 79}
    var v25 float32 = 0.1
    switch 109 {
    case 9223372036854775807:
        if false {
            v22 = (v22 & 66)
            v26 := GenS4{F0: 127, F1: 808, F2: 618}
            if (!(false || (42 == 100))) {
                v23 = append(v23, int(v2))
                for fi27 := 0; fi27 < 6; fi27++ {
                    var v28 bool = false
                    if (42 <= uint64(v2)) {
                        switch "" {
                        case "a":
                            v28 = v28
                        case "hello":
                            switch (-(fi27)) {
                            case -1:
                                if true {
                                    v2 = v2
                                    var v29 uint8 = 255
                                    if false {
                                        var v30 uint32 = 1
                                        _ = v30
                                    } else {
                                        fmt.Println(fi27, v28, v25)
                                    }
                                    _ = v29
                                } else {
                                    switch (fi27 ^ fi27) {
                                    case 1:
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                    }
                                }
                            case 128:
                                var v31 float32 = (v25 / -2.25)
                                _ = v31
                            }
                        }
                    }
                    _ = fi27
                    _ = v28
                }
            } else {
                v2 = v2
                v22 = (v22 + v22)
            }
            _ = v26
        } else {
            v32 := v1[uint32(v22)]
            _ = v32
        }
    case 238:
        delete(v1, 65535)
    case -533:
        v25 = 1e10
    }
    if false {
        v33 := GenS1{F0: 1, F1: false}
        for fi34 := 0; fi34 < 7; fi34++ {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            if false {
                v35 := GenS1{F0: -118, F1: true}
                var v36 uint64 = 4294967295
                if true {
                    fmt.Println(v2)
                } else {
                    v36 -= (^(v36))
                }
                _ = v35
                _ = v36
            } else {
                fmt.Println(v25, v25)
            }
            _ = len(v3)
            _ = fi34
        }
        v37 := v1[720]
        _ = v33
        _ = v37
    } else {
        if len(v23) > 0 {
            v23[0] = -9223372036854775808
        }
        v38 := GenS5{F0: "\n\t", F1: 2}
        _ = v38
    }
    delete(v3, 2)
    v25 = 3.14159
    v1[42] = 781
    for fi39 := 0; fi39 < 14; fi39++ {
        switch 32767 {
        case 1:
            v2 += v2
        case 2147483647:
            delete(v3, 42)
        case 58:
            var v40 uint8 = 8
            _ = v40
        default:
            var v41 int8 = 127
            _ = v41
        }
        if false {
            if (1 != 42) {
                var v42 int32 = 1
                v2 = (v2 - 670)
                _ = v42
            } else {
                for fi43 := 0; fi43 < 9; fi43++ {
                    var v44 uint8 = 2
                    _ = fi43
                    _ = v44
                }
            }
            fmt.Println(v2, fi39, v22)
        } else {
            var v45 int64 = 9223372036854775807
            _ = v45
        }
        switch ("" + ((((("foo bar" + "\n\t") + ("foo bar" + ("\n\t" + ("foo bar" + "a")))) + "hello") + "hello") + "a")) {
        case "foo bar":
            v46 := v1[42]
            _ = v46
        case "":
            fmt.Println(fi39, v2, v25)
        case "hello":
            v23 = append(v23, (fi39 << ((0) & 63)))
        }
        _ = fi39
    }
    var v47 int32 = 2147483647
    v25 = v25
    v25 -= 0.0
    if true {
        v23 = append(v23, 913)
        for fi48 := 0; fi48 < 8; fi48++ {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = fi48
        }
    } else {
        var v49 int = 1
        v22 -= (v22 >> ((0) & 63))
        v50 := v3[4294967295]
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 17; fi51++ {
        fmt.Println(v25, v47, v47)
        v25 = v25
        _ = fi51
    }
    v22 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v47
}
