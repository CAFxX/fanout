package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint32; F1 string; F2 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 bool; F2 string }
func genFn0(p0 int8, p1 int, p2 int8) uint {
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        v7 := GenS1{F0: 432}
        _ = v7
    }
    v8 := GenS2{F0: (true || false), F1: 2}
    var v9 string = "hello"
    var v10 int = p1
    var v11 int = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 65535
}
func genFn1(p0 uint8, p1 uint, p2 uint32) string {
    p0 -= p0
    v12 := GenS2{F0: false, F1: 255}
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            fmt.Println(p0, p0, p1)
            fmt.Println(p1, p0, p1)
            if true {
                var v15 int16 = int16(p2)
                var v16 float64 = ((0.0 - -2.25) / 1.5)
                if true {
                    var v17 uint16 = 938
                    var v18 float32 = (float32(v16) / float32((((-2.25 * (v16 - 0.1)) * (v16 + (((1e10 + float64(((((3.14159 + (0.1 - float32(-2.25))) + (-2.25 + ((0.1 / -2.25) - (0.0 * -2.25)))) / ((1e10 * 0.1) / (1.5 - (0.1 + 3.14159)))) / float32(v16)))) / v16) / v16))) + float64(float32(v16)))))
                    v19 := GenS2{F0: true, F1: (v17 - 1)}
                    v20 := GenS1{F0: 42}
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    if (true || false) {
                        var v21 uint = (p1 - 0)
                        _ = v21
                    } else {
                        v22 := []int{-788, (^(fi13))}
                        if ((!(!(!((v15 | (-(v15))) > v15)))) || ((p0 | (^(p0))) <= (p0 + 6))) {
                            var v23 uint32 = p2
                            v24 := GenS2{F0: (-128 >= int8(p2)), F1: 2}
                            _ = v23
                            _ = v24
                        } else {
                            if (0 > uint64(p2)) {
                                v25 := GenS2{F0: (false && true), F1: 2}
                                p2 -= 65535
                                _ = v25
                            }
                            fmt.Println(fi14)
                            v26 := GenS2{F0: false, F1: 255}
                            _ = v26
                        }
                        _ = v22
                    }
                    var v27 uint64 = 0
                    _ = v27
                }
                _ = v15
                _ = v16
            } else {
                v28 := GenS3{F0: 4294967295, F1: int(p2)}
                p2 += p2
                _ = v28
            }
            _ = fi14
        }
        _ = fi13
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        if (1 >= 1) {
            var v30 uint = genFn0(42, fi29, -128)
            if (42 >= 0) {
                v31 := []int{(fi29 >> ((p1) & 63)), fi29, fi29}
                fmt.Println(fi29, p2)
                var v32 uint32 = (p2 ^ (p2 * p2))
                _ = v31
                _ = v32
            } else {
                p0 -= p0
                v30 += p1
                for fi33 := 0; fi33 < 8; fi33++ {
                    var v34 float64 = 1e10
                    if true {
                        v35 := []int{fi29}
                        _ = v35
                    }
                    _ = fi33
                    _ = v34
                }
            }
            v36 := GenS2{F0: ("\n\t" == "a"), F1: uint16(p1)}
            _ = v30
            _ = v36
        }
        for fi37 := 0; fi37 < 19; fi37++ {
            var v38 uint = genFn0(100, 4294967295, 42)
            _ = fi37
            _ = v38
        }
        v39 := []int{fi29, (fi29 & 32768), fi29}
        _ = fi29
        _ = v39
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return "hello"
}

func main() {
    var v1 int8 = -52
    var v2 float64 = 0.0
    var v3 int = 0
    var v4 string = (("\n\t" + "a") + "a")
    var v5 float64 = (v2 - 1e10)
    v6 := make(map[int]uint32)
    v3 = int(v1)
    v3 = v3
    var v40 uint8 = 35
    switch 179 {
    case 1:
        switch v4 {
        case "\n\t":
            if true {
                var v41 uint = 1
                _ = v41
            }
        default:
            v40 = 255
        }
    case 4294967295:
        if (65535 >= 2) {
            var v42 int16 = -32768
            _ = v42
        } else {
            v40 += v40
            for fi43 := 0; fi43 < 5; fi43++ {
                var v44 int32 = -2147483648
                _ = fi43
                _ = v44
            }
            var v45 uint64 = 2
            _ = v45
        }
    }
    v46 := []int{}
    switch 42 {
    case 32767:
        for fi47 := 0; fi47 < 9; fi47++ {
            v4 = v4
            _ = fi47
        }
    case 1:
        if ((v4 + "a") != "hello") {
            if ((!(743 < 2)) && ((v4 + "\n\t") > ("foo bar" + (v4 + "hello")))) {
                if ((152 == 0) && (!true)) {
                    _ = len(v46)
                    v4 = (((((v4 + v4) + v4) + ("foo bar" + v4)) + (v4 + (v4 + ""))) + "")
                    switch v3 {
                    case 127:
                        var v48 string = (v4 + v4)
                        _ = v48
                    case 87:
                        switch (^(v3)) {
                        case 598:
                            var v49 int8 = v1
                            _ = v49
                        case -2147483648:
                            fmt.Println(v4, v5)
                        default:
                            var v50 uint16 = 1
                            _ = v50
                        }
                    case -1:
                        v51 := GenS2{F0: ((false || (-1 <= 719)) || (v3 == v3)), F1: 2}
                        _ = v51
                    }
                } else {
                    fmt.Println(v40)
                    v52 := GenS4{F0: 65535, F1: (("a" + ("\n\t" + (v4 + v4))) + v4), F2: v1}
                    _ = v52
                }
            } else {
                v4 = ((v4 + (("" + ((v4 + v4) + ("hello" + (v4 + v4)))) + ("\n\t" + (v4 + v4)))) + v4)
                switch v4 {
                case "a":
                    if ((!(!(-423 != 0))) || false) {
                        v53 := GenS5{F0: "\n\t"}
                        _ = v53
                    }
                case "hello":
                    v5 = 3.14159
                default:
                    var v54 int = (v3 * (^(v3)))
                    _ = v54
                }
                v55 := GenS2{F0: false, F1: 32767}
                _ = v55
            }
        }
    case -9223372036854775808:
        switch int(v40) {
        case 42:
            fmt.Println(v2)
        case -172:
            switch 128 {
            case 2147483647:
                v56 := v6[(-(v3))]
                _ = v56
            default:
                v3 = -1
            }
        case 681:
            v57 := GenS6{F0: v4}
            _ = v57
        }
    }
    var v58 uint64 = 1
    v1 += -59
    for fi59 := 0; fi59 < 18; fi59++ {
        if len(v46) >= 2 {
            v46 = v46[1:len(v46)]
        }
        var v60 uint16 = 935
        _ = fi59
        _ = v60
    }
    if ((((v4 + "") + v4) + v4) >= (((v4 + v4) + v4) + (v4 + v4))) {
        if false {
            if (((-(v1)) >= (v1 % (((-(v1))) | 1))) && true) {
                var v61 int8 = (v1 % ((v1) | 1))
                switch v4 {
                case "":
                    v1 = -128
                case "\n\t":
                    fmt.Println(v3, v1)
                }
                _ = v61
            }
            v40 += 42
        } else {
            if false {
                var v62 int32 = 42
                _ = v62
            } else {
                var v63 uint = 18446744073709551615
                _ = v63
            }
            var v64 int16 = -32768
            _ = v64
        }
    }
    v5 -= (v5 * float64(1.5))
    for fi65 := 0; fi65 < 1; fi65++ {
        v46 = append(v46, (fi65 ^ fi65))
        v4 = ""
        _ = fi65
    }
    if (!false) {
        fmt.Println(v4)
        v1 += v1
    } else {
        v40 += v40
        if (65535 < 1) {
            switch v3 {
            case 9223372036854775807:
                for fi66 := 0; fi66 < 13; fi66++ {
                    v40 = (v40 + (v40 + 245))
                    v46 = append(v46, (fi66 << ((genFn0((v1 & v1), 9223372036854775807, v1)) & 63)))
                    var v67 uint8 = (v40 & (v40 << ((65535) & 63)))
                    _ = fi66
                    _ = v67
                }
            }
        }
        for fi68 := 0; fi68 < 6; fi68++ {
            var v69 uint64 = (^(v58))
            if (false || false) {
                switch ("hello" + "") {
                case "foo bar":
                    v6[(fi68 + fi68)] = uint32(v69)
                default:
                    var v70 int16 = 42
                    _ = v70
                }
            }
            for fi71 := 0; fi71 < 15; fi71++ {
                fmt.Println(fi68, v1, v5)
                v40 = (v40 | (^(v40)))
                v40 -= v40
                _ = fi71
            }
            _ = fi68
            _ = v69
        }
    }
    v72 := GenS7{F0: true, F1: (genFn0(int8(v40), (-(v3)), (v1 << ((65535) & 63))) == 1), F2: "\n\t"}
    if true {
        if true {
            var v73 int = (^(v3))
            v73 += v73
            _ = v73
        }
        var v74 float64 = v5
        v1 += v1
        _ = v74
    } else {
        v75 := GenS5{F0: "a"}
        if (!(42 > 1)) {
            _ = len(v46)
        }
        fmt.Println(v3, v2, v3)
        _ = v75
    }
    if (false || true) {
        _ = len(v6)
        fmt.Println(v40, v5)
    } else {
        var v76 uint8 = v40
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v46
    _ = v5
    _ = v58
    _ = v6
    _ = v72
}
