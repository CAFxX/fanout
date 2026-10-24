package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 string }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 uint; F2 int }
type GenS7 struct { F0 uint8; F1 string; F2 int64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int32; F1 int16 }
type GenS10 struct { F0 uint8 }
func genFn0(p0 int16, p1 bool, p2 uint8) int {
    var v5 uint16 = 42
    switch ((("" + "a") + "a") + "foo bar") {
    case "hello":
        var v6 uint = 65535
        _ = v6
    case "foo bar":
        if (!p1) {
            if p1 {
                switch ("a" + "foo bar") {
                case "a":
                    var v7 int = 4294967295
                    _ = v7
                case "hello":
                    v8 := []int{42, 127}
                    _ = v8
                case "\n\t":
                    p1 = false
                default:
                    var v9 uint = 634
                    _ = v9
                }
                p0 += (p0 + int16(v5))
                for fi10 := 0; fi10 < 2; fi10++ {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v12 := []int{}
                        p1 = (42 != 127)
                        _ = fi11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 5; fi13++ {
                        var v14 bool = (v5 == uint16(p2))
                        if ((p2 % ((2) | 1)) <= (p2 - (p2 >> ((2) & 63)))) {
                            v15 := GenS1{F0: 2147483648, F1: 42}
                            _ = v15
                        } else {
                            if (p1 || (uint32(p2) == 1)) {
                                var v16 uint = 255
                                p0 = (p0 - (p0 + (-(p0))))
                                _ = v16
                            }
                            if true {
                                if ((0.0 == float32(3.14159)) && (("\n\t" >= "foo bar") && (3.14159 == 1e10))) {
                                    v17 := GenS1{F0: 133, F1: 311}
                                    var v18 string = ("hello" + "foo bar")
                                    v19 := GenS2{F0: 0}
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                }
                                v20 := GenS2{F0: (v5 & v5)}
                                _ = v20
                            } else {
                                v21 := GenS1{F0: 2, F1: int64(v5)}
                                v22 := GenS3{F0: 1, F1: uint16(p2), F2: (("a" + ((("hello" + ("a" + ("foo bar" + ("" + "a")))) + ("foo bar" + ("foo bar" + ("" + ("" + "a"))))) + "a")) + "")}
                                _ = v21
                                _ = v22
                            }
                            v23 := []int{fi10, (fi10 * (fi13 * (^(fi10)))), fi13, (-(fi13))}
                            _ = v23
                        }
                        _ = fi13
                        _ = v14
                    }
                    _ = fi10
                }
            }
            v24 := []int{-699}
            var v25 uint8 = (p2 >> ((2) & 63))
            _ = v24
            _ = v25
        } else {
            fmt.Println(v5, p2, p2)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}
func genFn1(p0 int32, p1 uint64, p2 uint64) uint8 {
    if ((!("" < "\n\t")) && ((2147483648 >= 42) && true)) {
        if (!true) {
            v26 := []int{32767, 9223372036854775807}
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 string = "foo bar"
            _ = fi27
            _ = v28
        }
        v29 := GenS4{F0: p0, F1: 10}
        _ = v29
    }
    fmt.Println(p2, p1, p2)
    if false {
        v30 := []int{32768, 9223372036854775807}
        for fi31 := 0; fi31 < 17; fi31++ {
            _ = len(v30)
            var v32 uint32 = 1
            _ = fi31
            _ = v32
        }
        _ = v30
    } else {
        p1 = p2
    }
    p1 -= (p2 >> ((4294967295) & 63))
    for fi33 := 0; fi33 < 1; fi33++ {
        v34 := []int{genFn0(0, true, 0), fi33, (^(fi33))}
        if len(v34) >= 2 {
            v34 = v34[1:len(v34)]
        }
        _ = fi33
        _ = v34
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 bool = (!false)
    var v4 float32 = 0.0
    if v3 {
        delete(v2, 840)
    }
    if v3 {
        v4 -= 0.1
        v4 = ((0.0 + v4) - v4)
        v4 += v4
    }
    v35 := []int{860, 32767}
    for fi36 := 0; fi36 < 10; fi36++ {
        v37 := v1[926]
        var v38 uint8 = 0
        for fi39 := 0; fi39 < 4; fi39++ {
            fmt.Println(v38, v4, v3)
            if (!v3) {
                fmt.Println(fi36, v4)
            }
            var v40 int16 = -125
            var v41 int16 = v40
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    switch 2147483647 {
    case 0:
        v42 := v1[32767]
        _ = v42
    }
    var v43 uint = 4294967295
    var v44 uint16 = 2
    v45 := GenS5{F0: -970}
    var v46 uint32 = 2147483648
    for fi47 := 0; fi47 < 7; fi47++ {
        if false {
            switch fi47 {
            case 4294967295:
                fmt.Println(v44, v46, v3)
            case -1:
                var v48 uint64 = 2
                _ = v48
            case 42:
                if len(v35) > 0 {
                    v35[0] = fi47
                }
            }
            var v49 float64 = 1.5
            if false {
                for fi50 := 0; fi50 < 20; fi50++ {
                    v43 -= 49
                    v51 := GenS6{F0: "", F1: (v43 | 65535), F2: genFn0(42, v3, 42)}
                    v52 := GenS7{F0: 128, F1: "\n\t", F2: 1}
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                v3 = true
                if v3 {
                    v35 = append(v35, fi47)
                    if v3 {
                        fmt.Println(v4, v43, v49)
                        for fi53 := 0; fi53 < 2; fi53++ {
                            fmt.Println(v44, v43, v46)
                            var v54 int64 = 0
                            switch genFn0(32767, v3, 128) {
                            case -124:
                                switch "foo bar" {
                                case "foo bar":
                                    _ = len(v35)
                                case "":
                                    if v3 {
                                        var v55 float64 = (-2.25 + (1.5 / v49))
                                        v44 -= (v44 >> ((v43) & 63))
                                        var v56 uint32 = v46
                                        _ = v55
                                        _ = v56
                                    }
                                default:
                                    switch "a" {
                                    case "foo bar":
                                        v49 -= float64(float32(v49))
                                    case "a":
                                        var v57 uint8 = 0
                                        _ = v57
                                    }
                                }
                            case 646:
                                switch (fi53 ^ -601) {
                                case 32767:
                                    if (v46 < v46) {
                                        switch "a" {
                                        case "\n\t":
                                            var v58 uint = 0
                                            _ = v58
                                        case "foo bar":
                                            var v59 uint = v43
                                            _ = v59
                                        case "":
                                            if (((float32(((v49 + (float64(float32(v49)) / float64(v4))) - 1.5)) * v4) - 1e10) <= v4) {
                                                switch ("" + ("\n\t" + ("a" + "hello"))) {
                                                case "foo bar":
                                                    fmt.Println(v49, fi47, v3)
                                                default:
                                                    fmt.Println(fi53, v46)
                                                }
                                                switch 836 {
                                                case 42:
                                                    fmt.Println(v4)
                                                case 4294967295:
                                                    var v60 float64 = 1e10
                                                    var v61 int32 = 459
                                                    v62 := GenS6{F0: "foo bar", F1: v43, F2: fi53}
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                case -1:
                                                    v44 = 42
                                                default:
                                                    var v63 bool = (!true)
                                                    _ = v63
                                                }
                                                var v64 int16 = 42
                                                _ = v64
                                            }
                                        default:
                                            v44 = (v44 >> ((v43) & 63))
                                        }
                                        if v3 {
                                            if (uint64(v43) <= 381) {
                                                v1[genFn0(976, (!(!v3)), 255)] = "hello"
                                                _ = len(v2)
                                                switch "foo bar" {
                                                case "a":
                                                    if ((v3 || v3) || (uint64(v46) > 42)) {
                                                        v65 := GenS8{F0: v43}
                                                        var v66 string = ("\n\t" + "\n\t")
                                                        var v67 string = v66
                                                        _ = v65
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                case "\n\t":
                                                    _ = len(v2)
                                                case "hello":
                                                    var v68 int8 = 100
                                                    _ = v68
                                                }
                                            }
                                        } else {
                                            v35 = append(v35, 1)
                                            var v69 uint8 = genFn1(-973, 2, 18446744073709551615)
                                            _ = v69
                                        }
                                        switch ("" + "hello") {
                                        case "a":
                                            v70 := GenS3{F0: -1, F1: v44, F2: "a"}
                                            _ = v70
                                        }
                                        var v71 float32 = v4
                                        _ = len(v35)
                                        var v72 uint = v43
                                        _ = v71
                                        _ = v72
                                    }
                                }
                            default:
                                var v73 float32 = -2.25
                                _ = v73
                            }
                            _ = fi53
                            _ = v54
                        }
                        if ((v43 < (^(v43))) && (!v3)) {
                            var v74 uint = v43
                            _ = v74
                        } else {
                            if v3 {
                                var v75 bool = (!((true || (true || (!(-2.25 > 3.14159)))) || true))
                                if len(v35) > 0 {
                                    v35[0] = -887
                                }
                                _ = v75
                            }
                            v76 := GenS9{F0: 42, F1: 0}
                            var v77 uint64 = 4294967295
                            _ = v76
                            _ = v77
                        }
                    } else {
                        _ = len(v1)
                        v78 := GenS10{F0: uint8(v43)}
                        _ = v78
                    }
                    v35 = append(v35, fi47)
                }
            }
            _ = v49
        }
        if (!true) {
            fmt.Println(fi47, fi47, v44)
        }
        fmt.Println(v44, v46)
        _ = fi47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
