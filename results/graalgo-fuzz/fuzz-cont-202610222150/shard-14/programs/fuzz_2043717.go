package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 bool; F1 uint64 }
func genFn0(p0 bool, p1 uint16, p2 uint16) uint16 {
    v5 := GenS1{F0: "\n\t", F1: 42}
    for fi6 := 0; fi6 < 17; fi6++ {
        p0 = p0
        _ = fi6
    }
    var v7 int = 1
    if p0 {
        v8 := []int{}
        for fi9 := 0; fi9 < 8; fi9++ {
            if false {
                if ((3.14159 * (0.1 / ((((float64(3.14159) / (0.1 - (3.14159 * float64(1e10)))) / (float64((0.1 / (0.1 / (((0.0 / 0.1) + 3.14159) * (0.1 + 0.1))))) * float64(((((0.1 / (1.5 - ((0.1 - 0.1) * 1e10))) - 0.0) - (float32(0.0) * 3.14159)) * ((1e10 * (0.1 / 0.1)) / (3.14159 + 3.14159)))))) + 1e10) * 1e10))) != 0.1) {
                    v7 -= (fi9 + v7)
                    for fi10 := 0; fi10 < 10; fi10++ {
                        if p0 {
                            p2 -= 0
                            var v11 uint32 = 2147483648
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            _ = v11
                        }
                        _ = fi10
                    }
                } else {
                    v12 := GenS1{F0: "hello", F1: 66}
                    p0 = (p0 && (p0 && (p0 || true)))
                    _ = v12
                }
                var v13 string = ("" + "\n\t")
                var v14 float32 = (0.0 / 1.0)
                _ = v13
                _ = v14
            }
            _ = len(v8)
            if false {
                for fi15 := 0; fi15 < 20; fi15++ {
                    fmt.Println(v7, p2, p0)
                    p1 = (^(p2))
                    _ = fi15
                }
                p1 -= (p2 + (p2 % (((^(p2))) | 1)))
            }
            _ = fi9
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return (^(p1))
}
func genFn1() int64 {
    for fi16 := 0; fi16 < 3; fi16++ {
        switch "\n\t" {
        case "a":
            var v17 uint32 = 189
            _ = v17
        case "":
            switch (("" + ("a" + "a")) + "\n\t") {
            case "a":
                v18 := []int{(fi16 & 620), fi16}
                _ = v18
            case "\n\t":
                v19 := GenS1{F0: "hello", F1: -128}
                _ = v19
            case "foo bar":
                for fi20 := 0; fi20 < 8; fi20++ {
                    switch "\n\t" {
                    case "":
                        fmt.Println(fi16, fi16)
                    case "foo bar":
                        v21 := []int{fi20, fi20, fi16}
                        _ = v21
                    case "\n\t":
                        v22 := []int{}
                        _ = v22
                    }
                    var v23 float32 = 1.5
                    var v24 bool = (0.1 < (1e10 * 1.5))
                    fmt.Println(v23)
                    v25 := []int{fi20, (-(fi20)), fi16}
                    _ = fi20
                    _ = v23
                    _ = v24
                    _ = v25
                }
            default:
                var v26 float64 = 0.0
                _ = v26
            }
        case "\n\t":
            switch fi16 {
            case 9223372036854775807:
                var v27 float32 = -2.25
                _ = v27
            case 42:
                switch ("a" + "a") {
                case "":
                    v28 := GenS1{F0: "a", F1: -17}
                    _ = v28
                case "foo bar":
                    fmt.Println(fi16)
                case "\n\t":
                    if (float64((float32((0.1 - -2.25)) / 3.14159)) > (-2.25 - (-2.25 * (1.5 + -2.25)))) {
                        var v29 int8 = 99
                        var v30 uint16 = 457
                        _ = v29
                        _ = v30
                    }
                }
            }
        }
        v31 := GenS1{F0: (((((("" + "a") + "\n\t") + (("a" + "hello") + ("hello" + "\n\t"))) + "\n\t") + "hello") + "a"), F1: 1}
        if (!true) {
            var v32 float64 = (1e10 * 1.5)
            _ = v32
        } else {
            if (false || true) {
                var v33 uint64 = 652
                if (true && false) {
                    switch (fi16 << ((1) & 63)) {
                    case -113:
                        for fi34 := 0; fi34 < 3; fi34++ {
                            switch ("" + ("" + "hello")) {
                            case "a":
                                v35 := GenS2{F0: 0}
                                _ = v35
                            case "foo bar":
                                v33 = 1
                            case "hello":
                                v36 := GenS3{F0: 255, F1: 880, F2: 100}
                                _ = v36
                            }
                            var v37 uint32 = 2147483647
                            v38 := GenS1{F0: "", F1: -1}
                            fmt.Println(v33)
                            v33 += v33
                            v39 := GenS4{F0: 0}
                            _ = fi34
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    }
                }
                _ = v33
            } else {
                var v40 int = -665
                _ = v40
            }
        }
        _ = fi16
        _ = v31
    }
    var v41 uint8 = 255
    var v42 uint16 = 32767
    var v43 uint16 = (v42 | uint16(v41))
    fmt.Println(v42, v43)
    _ = v41
    _ = v42
    _ = v43
    return 1
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]uint32)
    var v3 int16 = 42
    var v4 int8 = 42
    var v44 int64 = -506
    fmt.Println(v1)
    var v45 uint16 = genFn0(((3.14159 / ((1.5 + (0.0 / float32(1.5))) / (3.14159 - ((3.14159 + float32((-2.25 * ((v1 * 0.0) + v1)))) + (-2.25 * 3.14159))))) <= -2.25), 664, 892)
    var v46 string = ("\n\t" + (("\n\t" + "foo bar") + "foo bar"))
    fmt.Println(v3, v1)
    v47 := []int{}
    var v48 int64 = v44
    var v49 int32 = -823
    var v50 int16 = 42
    switch 32768 {
    case 32768:
        v45 = v45
    default:
        var v51 int8 = v4
        _ = v51
    }
    var v52 uint8 = 2
    for fi53 := 0; fi53 < 18; fi53++ {
        v45 = (v45 & v45)
        switch fi53 {
        case -695:
            v54 := GenS1{F0: ((("hello" + ("a" + (v46 + ("hello" + v46)))) + v46) + "hello"), F1: (^(v4))}
            _ = v54
        default:
            if ((^(v45)) > 2) {
                v55 := GenS4{F0: 4294967295}
                _ = len(v47)
                v48 -= -553
                _ = v55
            } else {
                switch fi53 {
                case 42:
                    for fi56 := 0; fi56 < 13; fi56++ {
                        v57 := GenS1{F0: "\n\t", F1: v4}
                        v44 += -9223372036854775808
                        _ = fi56
                        _ = v57
                    }
                case -182:
                    for fi58 := 0; fi58 < 13; fi58++ {
                        fmt.Println(v50, v50)
                        v48 = int64(v4)
                        v48 = v44
                        _ = fi58
                    }
                case 8:
                    v52 = v52
                default:
                    v44 = (^(v44))
                }
                var v59 int = (fi53 | fi53)
                _ = v59
            }
        }
        _ = fi53
    }
    v3 = v3
    var v60 uint8 = (v52 % ((v52) | 1))
    v61 := GenS5{F0: false, F1: 202}
    v62 := GenS6{F0: (genFn1() != (-(v44))), F1: 537}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v52
    _ = v60
    _ = v61
    _ = v62
}
