package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int16; F1 bool; F2 int32 }
type GenS9 struct { F0 uint32 }
func genFn0(p0 int16, p1 uint64, p2 uint8) uint8 {
    var v5 uint8 = (p2 / ((p2) | 1))
    var v6 int = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p2
}
func genFn1(p0 string, p1 uint32) uint32 {
    for fi7 := 0; fi7 < 2; fi7++ {
        p1 -= p1
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := []int{(fi9 << ((4294967295) & 63)), (fi8 | -2147483648), (fi8 / (((fi8 - fi9)) | 1)), fi9}
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 0}
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    return (p1 >> ((359) & 63))
}

func main() {
    var v1 uint = 0
    var v2 int64 = 0
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    v2 = v2
    v12 := []int{9223372036854775807, 0}
    if false {
        v1 += 2
        fmt.Println(v1, v2)
        for fi13 := 0; fi13 < 11; fi13++ {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                v1 = 42
            }
            if (207 <= 149) {
                fmt.Println(v2)
                switch (fi13 + (fi13 | fi13)) {
                case 128:
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v2 -= 42
                        if true {
                            var v15 uint32 = 2147483648
                            fmt.Println(v1, v15)
                            _ = v15
                        } else {
                            _ = len(v4)
                            fmt.Println(v2, v2, fi14)
                            v16 := GenS2{F0: ((^(v1)) >= v1), F1: (^(v1))}
                            _ = v16
                        }
                        var v17 int32 = 0
                        _ = fi14
                        _ = v17
                    }
                case 32768:
                    v2 = v2
                case -2147483648:
                    var v18 uint8 = genFn0(0, 4294967295, 0)
                    _ = v18
                }
                var v19 int64 = -1
                _ = v19
            } else {
                var v20 int8 = 127
                switch "a" {
                case "":
                    fmt.Println(fi13, v2)
                case "a":
                    switch fi13 {
                    case 32767:
                        var v21 float32 = 3.14159
                        _ = v21
                    case 2147483647:
                        switch (fi13 / ((int(v2)) | 1)) {
                        case 4294967295:
                            v1 = v1
                        }
                    default:
                        for fi22 := 0; fi22 < 18; fi22++ {
                            if false {
                                var v23 int16 = -32768
                                if true {
                                    v24 := GenS2{F0: ((v2 % ((v2) | 1)) != v2), F1: (^(v1))}
                                    v25 := GenS1{F0: 2}
                                    v12 = append(v12, -343)
                                    _ = v24
                                    _ = v25
                                } else {
                                    _ = len(v3)
                                }
                                _ = v23
                            } else {
                                v26 := GenS3{F0: 219}
                                fmt.Println(v20)
                                _ = v26
                            }
                            _ = fi22
                        }
                    }
                case "foo bar":
                    v20 -= v20
                }
                v27 := GenS3{F0: (v1 | (v1 | v1))}
                _ = v20
                _ = v27
            }
            _ = fi13
        }
    }
    var v28 uint8 = genFn0(0, 509, 127)
    v28 = v28
    v1 -= (v1 & v1)
    v29 := GenS4{F0: -604, F1: -1, F2: uint8(v1)}
    if len(v12) > 0 {
        v12[0] = 0
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        if ((("hello" + "") + "\n\t") > ("a" + "foo bar")) {
            for fi31 := 0; fi31 < 5; fi31++ {
                switch "hello" {
                case "\n\t":
                    var v32 uint32 = genFn1((("foo bar" + "") + "hello"), 4294967295)
                    fmt.Println(fi31)
                    v2 = (v2 & 409)
                    _ = v32
                }
                _ = fi31
            }
        } else {
            for fi33 := 0; fi33 < 13; fi33++ {
                if true {
                    delete(v4, fi30)
                    if false {
                        v28 += 59
                    }
                    var v34 uint16 = 255
                    _ = v34
                }
                _ = fi33
            }
            v35 := GenS4{F0: 1, F1: 0, F2: genFn0(-983, 884, v28)}
            _ = v35
        }
        _ = fi30
    }
    var v36 uint16 = 2
    for fi37 := 0; fi37 < 18; fi37++ {
        if (!((-2.25 <= 3.14159) || true)) {
            v36 = (^(v36))
            v2 = v2
            if len(v12) > 0 {
                v12[0] = (fi37 % (((fi37 + (-(fi37)))) | 1))
            }
        } else {
            switch ("a" + "hello") {
            case "hello":
                v38 := GenS3{F0: (v1 % (((v1 & v1)) | 1))}
                _ = v38
            case "foo bar":
                v39 := v3[127]
                _ = v39
            case "":
                for fi40 := 0; fi40 < 11; fi40++ {
                    v2 -= -9223372036854775808
                    _ = fi40
                }
            }
            for fi41 := 0; fi41 < 3; fi41++ {
                v1 += v1
                switch "foo bar" {
                case "foo bar":
                    v12 = append(v12, (fi37 & (fi37 - (-(fi41)))))
                case "\n\t":
                    v12 = append(v12, 0)
                    var v42 int8 = 0
                    _ = v42
                case "":
                    v43 := GenS4{F0: -32768, F1: 6, F2: 127}
                    _ = v43
                default:
                    switch 736 {
                    case 127:
                        v44 := GenS5{F0: (^(v36))}
                        _ = v44
                    case 1:
                        v45 := GenS3{F0: v1}
                        _ = v45
                    case -912:
                        fmt.Println(v36, v28)
                    default:
                        v46 := GenS6{F0: ((42 > 319) || false), F1: 32767}
                        _ = v46
                    }
                }
                v3[int(v2)] = ""
                switch "a" {
                case "hello":
                    switch fi41 {
                    case 32768:
                        _ = len(v12)
                    case -981:
                        if (-32768 < 32767) {
                            var v47 bool = ((false && true) || (514 == 18446744073709551615))
                            _ = v47
                        } else {
                            v48 := GenS7{F0: 1, F1: -716}
                            _ = v48
                        }
                    }
                }
                if true {
                    fmt.Println(v2, fi37)
                }
                _ = fi41
            }
        }
        _ = fi37
    }
    for fi49 := 0; fi49 < 9; fi49++ {
        for fi50 := 0; fi50 < 20; fi50++ {
            if (v1 < (v1 | v1)) {
                v51 := GenS8{F0: 0, F1: false, F2: -1}
                if ((!true) && true) {
                    fmt.Println(v1, v1)
                } else {
                    var v52 float64 = (0.0 + 1.5)
                    _ = v52
                }
                _ = v51
            }
            _ = fi50
        }
        _ = fi49
    }
    switch -402 {
    case -1:
        for fi53 := 0; fi53 < 10; fi53++ {
            v36 = (v36 / ((0) | 1))
            var v54 uint64 = 42
            _ = fi53
            _ = v54
        }
    case 1:
        v36 = (v36 | (v36 ^ (v36 % (((^(v36))) | 1))))
    default:
        v55 := GenS9{F0: uint32(v1)}
        _ = v55
    }
    for fi56 := 0; fi56 < 16; fi56++ {
        for fi57 := 0; fi57 < 15; fi57++ {
            switch 42 {
            case 127:
                switch ("" + (("a" + "") + "hello")) {
                case "a":
                    if true {
                        v28 -= (v28 + v28)
                    } else {
                        fmt.Println(fi57, v2)
                    }
                case "foo bar":
                    var v58 uint64 = 393
                    _ = v58
                default:
                    var v59 int32 = 1
                    _ = v59
                }
                fmt.Println(v28)
            case 32767:
                v60 := GenS2{F0: (v36 <= (v36 & v36)), F1: v1}
                _ = v60
            case -811:
                if (false && false) {
                    v2 = -1
                }
            }
            v3[fi57] = ((("\n\t" + "foo bar") + ("hello" + "hello")) + (("\n\t" + ("" + ("" + ("a" + ("hello" + "foo bar"))))) + ""))
            var v61 uint = 18446744073709551615
            _ = fi57
            _ = v61
        }
        for fi62 := 0; fi62 < 19; fi62++ {
            _ = len(v12)
            _ = len(v12)
            var v63 int16 = 1
            _ = fi62
            _ = v63
        }
        var v64 float32 = float32((1.5 * 3.14159))
        _ = fi56
        _ = v64
    }
    for fi65 := 0; fi65 < 20; fi65++ {
        if (-2.25 <= 0.1) {
            fmt.Println(v36, fi65, v2)
        }
        var v66 int = fi65
        _ = fi65
        _ = v66
    }
    v36 = v36
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
}
