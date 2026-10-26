package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 bool; F2 int16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32; F1 uint32; F2 int }
type GenS6 struct { F0 uint64; F1 string; F2 int }
type GenS7 struct { F0 int64; F1 string; F2 uint }
func genFn0(p0 uint64, p1 uint, p2 uint64) string {
    p1 += (p1 - p1)
    fmt.Println(p0, p0)
    var v5 int64 = 0
    switch 9223372036854775807 {
    case 32767:
        p0 = (p0 / (((p2 / (((^(p0))) | 1))) | 1))
    case 128:
        p0 += (p0 & (^(p2)))
    case 2147483647:
        v5 -= -1
    default:
        p1 -= 18446744073709551615
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        p2 = uint64(p1)
        if false {
            switch -670 {
            case 483:
                if true {
                    v7 := GenS1{F0: 42, F1: -1, F2: (p0 % ((p2) | 1))}
                    _ = v7
                } else {
                    for fi8 := 0; fi8 < 4; fi8++ {
                        if true {
                            fmt.Println(p1, p0, p0)
                        } else {
                            switch (("" + "foo bar") + "hello") {
                            case "a":
                                var v9 float64 = float64(0.0)
                                _ = v9
                            case "\n\t":
                                v10 := GenS2{F0: -1}
                                _ = v10
                            default:
                                v11 := []int{}
                                _ = v11
                            }
                        }
                        _ = fi8
                    }
                }
            case -2147483648:
                var v12 uint = (p1 | (p1 | p1))
                _ = v12
            default:
                if false {
                    switch 42 {
                    case 128:
                        v13 := GenS1{F0: 2147483647, F1: 231, F2: 837}
                        _ = v13
                    case -1:
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v5 -= -1
                            var v15 int32 = 0
                            _ = fi14
                            _ = v15
                        }
                    case -2147483648:
                        if false {
                            if ((-(fi6)) != 1) {
                                v16 := GenS3{F0: 32768, F1: (-399 < v5), F2: 0}
                                p2 -= 898
                                _ = v16
                            }
                            v17 := GenS2{F0: 1}
                            _ = v17
                        } else {
                            v18 := GenS4{F0: (p2 - p2)}
                            for fi19 := 0; fi19 < 11; fi19++ {
                                v20 := GenS5{F0: -28, F1: 2147483648, F2: fi19}
                                var v21 string = ""
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            _ = v18
                        }
                    default:
                        switch fi6 {
                        case 9223372036854775807:
                            for fi22 := 0; fi22 < 16; fi22++ {
                                if true {
                                    if false {
                                        var v23 uint8 = 0
                                        if ("" >= "") {
                                            var v24 float64 = (0.0 / 1.0)
                                            var v25 int = (fi6 * int(v5))
                                            var v26 uint64 = (p2 | (p0 / ((p2) | 1)))
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        }
                                        var v27 int8 = -40
                                        _ = v23
                                        _ = v27
                                    }
                                }
                                _ = fi22
                            }
                        }
                    }
                    if (!(("foo bar" + (("foo bar" + "\n\t") + (("" + (("\n\t" + (("\n\t" + "\n\t") + ("hello" + "\n\t"))) + ("\n\t" + ("hello" + "a")))) + ("foo bar" + (("hello" + "foo bar") + ("hello" + (("a" + "\n\t") + "a"))))))) == ("\n\t" + "\n\t"))) {
                        switch 205 {
                        case 149:
                            v28 := []int{(fi6 + 0), fi6}
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 20; fi29++ {
                            p1 = (p1 | (p1 & (p1 + p1)))
                            if true {
                                var v30 uint8 = 0
                                v30 = v30
                                switch "\n\t" {
                                case "":
                                    v30 = v30
                                case "\n\t":
                                    if true {
                                        switch (fi6 % ((672) | 1)) {
                                        case -152:
                                            if (!true) {
                                                var v31 int16 = -1
                                                var v32 uint64 = (p0 >> ((p1) & 63))
                                                v33 := []int{}
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                            }
                                        case 2147483647:
                                            v34 := []int{fi29, -1, (^(fi6))}
                                            _ = v34
                                        }
                                        v35 := GenS5{F0: 628, F1: 65535, F2: fi6}
                                        var v36 uint32 = 4294967295
                                        var v37 int16 = -71
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                case "foo bar":
                                    v30 -= v30
                                }
                                v30 = (v30 + uint8(p2))
                                _ = v30
                            }
                            _ = fi29
                        }
                    }
                }
            }
        } else {
            var v38 int8 = 1
            var v39 int32 = -299
            var v40 int64 = (v5 ^ 42)
            _ = v38
            _ = v39
            _ = v40
        }
        if false {
            v41 := []int{}
            var v42 bool = false
            _ = v41
            _ = v42
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "hello"
}

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    var v3 int8 = -64
    var v4 int64 = -727
    if (true || true) {
        if (68 >= 1) {
            v3 -= 19
            if ((26 != 42) || false) {
                v4 = v4
                for fi43 := 0; fi43 < 18; fi43++ {
                    v44 := GenS6{F0: 2, F1: (("foo bar" + "foo bar") + ("hello" + "")), F2: 536}
                    v2["a"] = ""
                    _ = fi43
                    _ = v44
                }
                fmt.Println(v4, v4)
            } else {
                v45 := GenS2{F0: -1}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v45
            }
        }
    } else {
        var v46 bool = true
        fmt.Println(v46, v4, v46)
        _ = v46
    }
    var v47 float32 = 1.5
    switch 1 {
    case 127:
        v48 := GenS6{F0: 4294967295, F1: "a", F2: 4294967295}
        _ = v48
    }
    fmt.Println(v3, v3)
    var v49 uint16 = 2
    var v50 float32 = v47
    v3 = (v3 + (v3 - (v3 / (((v3 >> ((473) & 63))) | 1))))
    var v51 bool = true
    v52 := GenS7{F0: (v4 - v4), F1: "a", F2: 0}
    if (v50 == 0.0) {
        if (!true) {
            for fi53 := 0; fi53 < 5; fi53++ {
                if v51 {
                    var v54 float64 = (-2.25 * 0.1)
                    _ = v54
                }
                v4 += 1
                v1 = append(v1, fi53)
                _ = fi53
            }
        }
        var v55 bool = (652 < 9)
        if v51 {
            var v56 string = "a"
            var v57 uint16 = v49
            _ = v56
            _ = v57
        } else {
            v47 += v50
            fmt.Println(v55, v50)
        }
        _ = v55
    }
    if (!v51) {
        v3 += -102
    }
    fmt.Println(v3)
    var v58 int16 = -1
    for fi59 := 0; fi59 < 20; fi59++ {
        for fi60 := 0; fi60 < 4; fi60++ {
            var v61 float32 = (v47 * ((float32((3.14159 + 0.0)) * v50) + ((3.14159 - 1e10) / ((v47 / v50) * 1.5))))
            v51 = false
            _ = fi60
            _ = v61
        }
        _ = len(v2)
        _ = fi59
    }
    switch 906 {
    case 1:
        if true {
            var v62 float32 = (((v50 / (v47 * v47)) + v50) * v47)
            _ = v62
        }
    }
    if v51 {
        for fi63 := 0; fi63 < 6; fi63++ {
            v3 = (v3 >> ((65535) & 63))
            _ = fi63
        }
        for fi64 := 0; fi64 < 4; fi64++ {
            v65 := v2[(("" + ("a" + "\n\t")) + "hello")]
            v2[((((v65 + ("\n\t" + v65)) + v65) + "foo bar") + "foo bar")] = ("\n\t" + v65)
            _ = fi64
            _ = v65
        }
        v58 = (^(v58))
    } else {
        for fi66 := 0; fi66 < 10; fi66++ {
            v58 += v58
            _ = len(v1)
            _ = fi66
        }
        v67 := GenS3{F0: 128, F1: false, F2: (-(v58))}
        _ = v67
    }
    if ((v58 ^ (-(v58))) < v58) {
        v58 += 32767
        v68 := GenS3{F0: -9223372036854775808, F1: (213 <= 587), F2: int16(v4)}
        _ = v68
    }
    var v69 int64 = (v4 >> ((1) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v58:", v58)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v58
    _ = v69
}
