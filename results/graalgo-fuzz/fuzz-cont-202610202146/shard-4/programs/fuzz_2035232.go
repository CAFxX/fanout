package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int; F1 uint32; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 int16; F2 uint16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 int16; F1 string; F2 string }
type GenS10 struct { F0 int16 }
func genFn0(p0 string) uint8 {
    for fi6 := 0; fi6 < 12; fi6++ {
        if (true || false) {
            p0 = ("" + "a")
            for fi7 := 0; fi7 < 17; fi7++ {
                v8 := GenS1{F0: 42, F1: "a", F2: -39}
                var v9 uint32 = 116
                _ = fi7
                _ = v8
                _ = v9
            }
            var v10 string = ("hello" + "foo bar")
            _ = v10
        } else {
            fmt.Println(p0, fi6)
            v11 := GenS2{F0: 108, F1: true}
            var v12 uint8 = 255
            _ = v11
            _ = v12
        }
        p0 = (p0 + p0)
        v13 := GenS3{F0: 0, F1: 672, F2: 2}
        _ = fi6
        _ = v13
    }
    if true {
        if (625 > 4294967295) {
            p0 = (p0 + "")
            var v14 int8 = -1
            _ = v14
        }
        p0 = ("a" + (p0 + ("" + ("\n\t" + ""))))
    }
    if true {
        for fi15 := 0; fi15 < 10; fi15++ {
            fmt.Println(p0, fi15)
            var v16 int = (fi15 / (((fi15 / ((fi15) | 1))) | 1))
            v17 := GenS3{F0: (v16 / ((32767) | 1)), F1: 1, F2: 754}
            _ = fi15
            _ = v16
            _ = v17
        }
        if (true || (128 > 0)) {
            for fi18 := 0; fi18 < 10; fi18++ {
                switch "foo bar" {
                case "":
                    if false {
                        var v19 int8 = 127
                        if ((fi18 >> ((18446744073709551615) & 63)) != 32767) {
                            v20 := GenS4{F0: 244}
                            v21 := GenS5{F0: 60}
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    }
                case "hello":
                    fmt.Println(fi18, p0)
                case "foo bar":
                    var v22 int16 = 884
                    _ = v22
                }
                _ = fi18
            }
            var v23 int16 = 32767
            p0 = "hello"
            _ = v23
        }
        fmt.Println(p0)
    }
    var v24 uint64 = 1
    var v25 float64 = -2.25
    _ = p0
    _ = v24
    _ = v25
    return 0
}
func genFn1(p0 int32, p1 bool) int {
    switch -1 {
    case 32768:
        p1 = ((!false) || (!(false && (1 != 127))))
    case 2147483647:
        v26 := []int{1}
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 13; fi28++ {
            p1 = p1
            p1 = p1
            switch "a" {
            case "":
                var v29 uint16 = 1
                _ = v29
            case "hello":
                if (true && p1) {
                    if p1 {
                        fmt.Println(p1, p1, p0)
                    }
                    switch (fi28 + -426) {
                    case -1:
                        fmt.Println(p0, p1, fi27)
                    case 32768:
                        if false {
                            fmt.Println(fi27, fi28)
                            p0 = (p0 & p0)
                            fmt.Println(p1)
                        }
                        var v30 int32 = 746
                        var v31 int64 = -1
                        var v32 uint32 = 2147483648
                        _ = v30
                        _ = v31
                        _ = v32
                    default:
                        switch "" {
                        case "a":
                            v33 := GenS4{F0: 65535}
                            _ = v33
                        case "foo bar":
                            var v34 uint = 1
                            _ = v34
                        case "":
                            if ((float32(float64((3.14159 - 0.0))) >= (1e10 * (1e10 * 0.1))) || (!true)) {
                                p1 = true
                                p1 = p1
                            } else {
                                switch "a" {
                                case "\n\t":
                                    v35 := GenS6{F0: 255, F1: 1, F2: 42}
                                    _ = v35
                                default:
                                    v36 := GenS3{F0: fi27, F1: 20, F2: 752}
                                    _ = v36
                                }
                                v37 := []int{(fi27 / (((-(fi28))) | 1))}
                                v37 = append(v37, 42)
                                if (-32768 <= 77) {
                                    if len(v37) > 0 {
                                        v37[0] = 4294967295
                                    }
                                    p1 = p1
                                    _ = len(v37)
                                } else {
                                    if len(v37) > 0 {
                                        v37[0] = 32768
                                    }
                                    var v38 uint8 = genFn0(("foo bar" + "\n\t"))
                                    _ = v38
                                }
                                _ = v37
                            }
                        default:
                            var v39 float32 = ((0.1 * -2.25) - 0.0)
                            _ = v39
                        }
                    }
                    if true {
                        v40 := []int{fi27, -478}
                        var v41 uint64 = 939
                        fmt.Println(p1, p1)
                        _ = v40
                        _ = v41
                    }
                }
            case "foo bar":
                v42 := []int{-9223372036854775808, fi28, (fi27 / ((fi27) | 1)), (fi28 - 1)}
                _ = v42
            }
            _ = fi28
        }
        for fi43 := 0; fi43 < 3; fi43++ {
            switch "\n\t" {
            case "a":
                var v44 int64 = 999
                _ = v44
            case "\n\t":
                if p1 {
                    fmt.Println(fi43, p0, p1)
                    v45 := []int{127, fi27}
                    if p1 {
                        var v46 int = fi43
                        if p1 {
                            switch "\n\t" {
                            case "":
                                var v47 int8 = 127
                                _ = v47
                            default:
                                var v48 uint32 = 0
                                _ = v48
                            }
                            if p1 {
                                var v49 int16 = 742
                                v50 := GenS7{F0: 65535}
                                v46 -= int(v49)
                                _ = v49
                                _ = v50
                            } else {
                                v46 -= (fi27 / (((-(fi27))) | 1))
                            }
                            p1 = p1
                            fmt.Println(p1, p1)
                            if (p1 && (0 > 42)) {
                                v46 += (v46 * v46)
                                _ = len(v45)
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                            } else {
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                                var v51 int32 = (p0 / ((1) | 1))
                                _ = v51
                            }
                        } else {
                            v52 := GenS8{F0: 32767, F1: 255}
                            v46 -= 42
                            _ = v52
                        }
                        if (!((fi27 == 2147483647) && (!(18446744073709551615 >= 4294967295)))) {
                            if true {
                                var v53 uint = 255
                                _ = v53
                            } else {
                                if (42 == 2) {
                                    var v54 float64 = (1.5 * -2.25)
                                    _ = v54
                                } else {
                                    p1 = p1
                                }
                            }
                        } else {
                            if true {
                                var v55 uint = 42
                                switch fi27 {
                                case 1:
                                    _ = len(v45)
                                case 32767:
                                    p0 -= 609
                                case 127:
                                    p1 = (-9223372036854775808 < 41)
                                }
                                _ = v55
                            }
                            fmt.Println(v46, p1, fi43)
                            var v56 uint64 = 0
                            _ = v56
                        }
                        _ = v46
                    }
                    p0 = p0
                    _ = v45
                }
            }
            var v57 string = "foo bar"
            var v58 uint32 = 2
            _ = fi43
            _ = v57
            _ = v58
        }
        v59 := GenS9{F0: -32768, F1: ("a" + "\n\t"), F2: ("hello" + "hello")}
        _ = fi27
        _ = v59
    }
    _ = p0
    _ = p1
    return -858
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 42
    var v3 uint32 = 2147483648
    v4 := make(map[int]uint32)
    var v5 int16 = -1
    var v60 int32 = 2147483647
    var v61 bool = false
    v60 = (v60 | (v60 - (v60 >> ((1) & 63))))
    v62 := GenS10{F0: (v5 - (v5 * (-(v5))))}
    v63 := []int{128}
    _ = len(v63)
    v5 = -1
    v64 := GenS10{F0: v5}
    var v65 int = -1
    v60 += (v60 | v60)
    for fi66 := 0; fi66 < 16; fi66++ {
        v67 := GenS6{F0: 824, F1: (-(v5)), F2: 742}
        _ = fi66
        _ = v67
    }
    if len(v63) > 0 {
        v63[0] = (v65 << ((42) & 63))
    }
    v60 -= (v60 % ((v60) | 1))
    var v68 uint = 890
    var v69 uint8 = (v2 >> ((v68) & 63))
    for fi70 := 0; fi70 < 10; fi70++ {
        fmt.Println(v61, v60, v5)
        _ = fi70
    }
    var v71 uint = v68
    v60 = (^(v60))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
    _ = v68
    _ = v69
    _ = v71
}
