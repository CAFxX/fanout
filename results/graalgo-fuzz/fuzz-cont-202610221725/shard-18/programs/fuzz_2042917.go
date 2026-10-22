package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 bool; F1 uint16; F2 int32 }
type GenS4 struct { F0 int; F1 uint16; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 bool; F2 uint64 }
func genFn0(p0 bool, p1 uint8, p2 string) uint8 {
    p2 = p2
    v6 := []int{841, -388, 334}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p1 ^ (p1 & 42))
}
func genFn1(p0 int) int32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 += -1
        v8 := []int{182, -9223372036854775808, p0}
        var v9 bool = (0 <= 832)
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        if false {
            p0 = (p0 << ((2) & 63))
            for fi10 := 0; fi10 < 11; fi10++ {
                v11 := GenS1{F0: 182}
                _ = fi10
                _ = v11
            }
        }
        fmt.Println(p0)
        v12 := GenS1{F0: 2}
        _ = v12
    } else {
        var v13 bool = ((!false) || (((!true) || (2147483647 != 1)) && true))
        v14 := []int{(p0 + 32768)}
        if ((p0 + p0) == p0) {
            if false {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                if v13 {
                    var v15 bool = false
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        if (1 > 2) {
                            v17 := GenS2{F0: -73, F1: 42}
                            _ = len(v14)
                            var v18 string = ""
                            _ = v17
                            _ = v18
                        }
                        var v19 int8 = 127
                        _ = fi16
                        _ = v19
                    }
                    var v20 float32 = 1e10
                    if len(v14) > 0 {
                        v14[0] = (p0 + p0)
                    }
                    _ = v20
                }
                p0 += 463
            } else {
                v21 := GenS1{F0: 4294967295}
                _ = v21
            }
        }
        _ = v13
        _ = v14
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        fmt.Println(fi22)
        _ = fi22
    }
    switch p0 {
    case 32767:
        var v23 uint = 395
        _ = v23
    case 9223372036854775807:
        if false {
            v24 := GenS3{F0: (!("foo bar" == (("foo bar" + "hello") + ("foo bar" + ((("a" + ("" + "")) + "foo bar") + ("" + "foo bar")))))), F1: 857, F2: 935}
            p0 = p0
            _ = v24
        } else {
            p0 -= 11
            v25 := GenS4{F0: (p0 / ((p0) | 1)), F1: 0, F2: -47}
            _ = v25
        }
    case -9223372036854775808:
        if true {
            for fi26 := 0; fi26 < 17; fi26++ {
                for fi27 := 0; fi27 < 15; fi27++ {
                    v28 := []int{}
                    p0 = fi27
                    p0 = (fi26 % (((p0 + fi26)) | 1))
                    _ = fi27
                    _ = v28
                }
                p0 += fi26
                _ = fi26
            }
            p0 = p0
            for fi29 := 0; fi29 < 12; fi29++ {
                for fi30 := 0; fi30 < 1; fi30++ {
                    v31 := GenS3{F0: (32768 > (^(p0))), F1: 1, F2: 962}
                    v32 := []int{fi30, (fi30 >> ((255) & 63)), (-(fi30))}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = fi29
            }
        } else {
            var v33 float64 = ((float64(0.0) - 0.1) * (0.0 - 3.14159))
            for fi34 := 0; fi34 < 15; fi34++ {
                switch (^(p0)) {
                case -2147483648:
                    v35 := GenS4{F0: (p0 << ((246) & 63)), F1: 0, F2: -116}
                    _ = v35
                case 60:
                    var v36 float64 = (float64(float32(3.14159)) - v33)
                    _ = v36
                default:
                    v33 -= v33
                }
                _ = fi34
            }
            _ = v33
        }
    }
    _ = p0
    return 594
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{("foo bar" + ((("hello" + "\n\t") + ("" + ("foo bar" + "hello"))) + "a")), "hello", ("\n\t" + "foo bar"), "foo bar"}
    v3 := []int{32768, -2147483648}
    v4 := make(map[string]bool)
    v5 := []uint32{1, 346, 4294967295, 2147483648}
    switch ("" + (("\n\t" + "\n\t") + "hello")) {
    case "hello":
        var v37 int = 32767
        _ = v37
    case "foo bar":
        v38 := GenS3{F0: false, F1: 32767, F2: 1}
        _ = v38
    default:
        if true {
            var v39 bool = false
            _ = v39
        } else {
            var v40 bool = false
            for fi41 := 0; fi41 < 13; fi41++ {
                v42 := GenS3{F0: true, F1: 408, F2: genFn1((fi41 >> ((255) & 63)))}
                _ = fi41
                _ = v42
            }
            _ = v40
        }
    }
    var v43 uint64 = 1
    v44 := GenS5{F0: (!(!true))}
    var v45 float32 = 0.1
    v46 := GenS3{F0: ("hello" > "hello"), F1: 871, F2: -1}
    for fi47 := 0; fi47 < 6; fi47++ {
        var v48 float32 = v45
        _ = len(v1)
        v49 := GenS3{F0: false, F1: 399, F2: 740}
        _ = fi47
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 5; fi50++ {
        v51 := GenS2{F0: 14, F1: 0}
        v43 += (v43 | (v43 & v43))
        _ = fi50
        _ = v51
    }
    switch 2147483647 {
    case -170:
        for fi52 := 0; fi52 < 8; fi52++ {
            v45 = (v45 + (1.5 / v45))
            v53 := GenS4{F0: 0, F1: 66, F2: 42}
            _ = fi52
            _ = v53
        }
    case -2147483648:
        if true {
            var v54 uint8 = 59
            _ = v54
        } else {
            v45 = v45
        }
    case 123:
        delete(v1, 0)
    default:
        if (!false) {
            v55 := v4["foo bar"]
            var v56 int8 = -1
            _ = v55
            _ = v56
        } else {
            delete(v1, 4294967295)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if false {
                for fi57 := 0; fi57 < 4; fi57++ {
                    if false {
                        v45 += v45
                        fmt.Println(v43, v43)
                        fmt.Println(v45, v43, fi57)
                    } else {
                        v43 -= v43
                        for fi58 := 0; fi58 < 18; fi58++ {
                            delete(v4, ("hello" + (((((("" + "a") + "") + "") + "hello") + "foo bar") + "foo bar")))
                            if (4294967295 == 42) {
                                v59 := v1[fi57]
                                _ = v59
                            } else {
                                v60 := GenS6{F0: 112, F1: 65535}
                                _ = v60
                            }
                            if (((469 != (v43 / (((v43 / (((v43 >> ((4294967295) & 63))) | 1))) | 1))) && false) || (42 >= genFn1(2147483647))) {
                                _ = len(v1)
                                v61 := v4["\n\t"]
                                _ = v61
                            } else {
                                if true {
                                    v62 := GenS7{F0: 1000}
                                    _ = v62
                                } else {
                                    if true {
                                        var v63 string = ("" + "\n\t")
                                        _ = v63
                                    } else {
                                        v64 := GenS7{F0: 0}
                                        _ = v64
                                    }
                                    var v65 uint32 = 2
                                    var v66 string = "\n\t"
                                    _ = v65
                                    _ = v66
                                }
                            }
                            _ = fi58
                        }
                        v67 := GenS7{F0: -1}
                        _ = v67
                    }
                    v5 = append(v5, 2)
                    v45 = (float32(1e10) + (v45 - 0.0))
                    _ = fi57
                }
                switch (("hello" + ("\n\t" + ("" + "foo bar"))) + (("" + "\n\t") + ((("a" + "a") + ("\n\t" + "foo bar")) + ("a" + "foo bar")))) {
                case "":
                    v68 := GenS8{F0: 42, F1: (42 > -482), F2: (v43 * (v43 * (^(v43))))}
                    _ = v68
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
    }
    switch "\n\t" {
    case "hello":
        var v69 string = ("" + "foo bar")
        _ = v69
    case "\n\t":
        var v70 int16 = 75
        _ = v70
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
