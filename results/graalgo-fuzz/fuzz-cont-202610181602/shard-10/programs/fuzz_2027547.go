package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int8; F1 bool; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16; F1 int32; F2 uint32 }
func genFn0() int {
    var v7 int = 4294967295
    fmt.Println(v7)
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 -= v7
        if (!false) {
            fmt.Println(v7, fi8)
            v9 := []int{-9223372036854775808, v7, 0}
            _ = v9
        }
        v7 = (^(fi8))
        _ = fi8
    }
    _ = v7
    return 32768
}
func genFn1(p0 uint8, p1 int32) uint8 {
    p0 -= 54
    fmt.Println(p0)
    v10 := []int{-17, 127, 2147483647}
    _ = p0
    _ = p1
    _ = v10
    return (p0 - 127)
}

func main() {
    var v1 uint64 = 558
    var v2 int16 = 42
    var v3 float32 = 1.5
    var v4 float32 = float32(0.0)
    v5 := make(map[int]uint32)
    var v6 int32 = -434
    v6 = (v6 + v6)
    if (genFn0() > 266) {
        v11 := GenS1{F0: 0, F1: v2}
        var v12 int64 = -922
        for fi13 := 0; fi13 < 2; fi13++ {
            if ("" <= "\n\t") {
                if (float32((float64((float32(1.5) + v4)) + 0.1)) < ((3.14159 * 1e10) / 1e10)) {
                    var v14 uint64 = (v1 / ((2) | 1))
                    var v15 string = "foo bar"
                    _ = v14
                    _ = v15
                } else {
                    v16 := []int{}
                    v17 := v5[(fi13 - genFn0())]
                    _ = v16
                    _ = v17
                }
                if ((849 != 42) && (42 >= 65535)) {
                    v18 := GenS1{F0: 62, F1: (v2 << ((255) & 63))}
                    for fi19 := 0; fi19 < 2; fi19++ {
                        var v20 uint = 4294967295
                        var v21 uint = (v20 | v20)
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = v18
                } else {
                    var v22 float64 = 0.0
                    _ = v22
                }
                v23 := GenS1{F0: v12, F1: (^(v2))}
                _ = v23
            } else {
                var v24 int8 = int8(v2)
                if (true && true) {
                    for fi25 := 0; fi25 < 11; fi25++ {
                        v26 := GenS2{F0: 4294967295}
                        fmt.Println(fi25, v6, v24)
                        _ = fi25
                        _ = v26
                    }
                    if false {
                        fmt.Println(v1, v3, v1)
                        v12 = v12
                    } else {
                        for fi27 := 0; fi27 < 1; fi27++ {
                            if true {
                                v12 -= 367
                                var v28 int8 = v24
                                var v29 string = ("" + ("hello" + "foo bar"))
                                var v30 uint8 = 224
                                _ = v28
                                _ = v29
                                _ = v30
                            } else {
                                var v31 int64 = -1
                                v32 := v5[genFn0()]
                                _ = v31
                                _ = v32
                            }
                            _ = fi27
                        }
                    }
                    if false {
                        v33 := v5[genFn0()]
                        v4 -= float32(1.5)
                        for fi34 := 0; fi34 < 17; fi34++ {
                            v35 := []int{(fi13 - (fi34 % ((genFn0()) | 1))), (fi34 * fi13), (fi13 + (fi13 & -2147483648))}
                            v36 := GenS2{F0: 499}
                            v37 := GenS1{F0: 430, F1: (v2 | v2)}
                            _ = fi34
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        _ = v33
                    } else {
                        var v38 uint8 = 218
                        if ((2147483647 >= 65535) || (-32768 <= 42)) {
                            var v39 uint8 = 245
                            for fi40 := 0; fi40 < 2; fi40++ {
                                var v41 string = "a"
                                var v42 int64 = (v12 << ((4294967295) & 63))
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                            _ = v39
                        } else {
                            v43 := v5[(^(fi13))]
                            _ = v43
                        }
                        _ = v38
                    }
                } else {
                    if false {
                        v44 := GenS2{F0: 65535}
                        for fi45 := 0; fi45 < 4; fi45++ {
                            if false {
                                v1 -= (v1 | v1)
                            } else {
                                fmt.Println(fi13, v4, v24)
                                if true {
                                    v46 := []int{fi45}
                                    v2 -= v2
                                    _ = v46
                                }
                            }
                            v12 -= v12
                            v6 -= 42
                            fmt.Println(v3, v3)
                            _ = fi45
                        }
                        switch genFn0() {
                        case -2147483648:
                            fmt.Println(v4)
                        case 127:
                            if (v6 <= -2147483648) {
                                switch (fi13 >> ((606) & 63)) {
                                case 9223372036854775807:
                                    v1 -= v1
                                }
                                if (!(1 <= (v6 ^ v6))) {
                                    v47 := v5[(-(fi13))]
                                    v24 -= v24
                                    var v48 uint32 = (v47 | (v47 << ((18446744073709551615) & 63)))
                                    _ = v47
                                    _ = v48
                                }
                                var v49 bool = true
                                _ = v49
                            }
                        default:
                            if (4294967295 == 2) {
                                for fi50 := 0; fi50 < 18; fi50++ {
                                    v51 := GenS3{F0: (v1 / ((v1) | 1)), F1: 420}
                                    v52 := GenS2{F0: 324}
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                }
                            } else {
                                fmt.Println(v6, v4)
                                v5[fi13] = 300
                            }
                        }
                        _ = v44
                    } else {
                        v1 += v1
                        for fi53 := 0; fi53 < 3; fi53++ {
                            if false {
                                fmt.Println(v12, v3, v6)
                            }
                            v54 := GenS4{F0: (v24 << ((954) & 63)), F1: true, F2: 2}
                            v55 := []int{(fi13 & fi13), -281, fi53}
                            v1 -= v1
                            fmt.Println(v24, v12)
                            switch fi13 {
                            case 23:
                                if false {
                                    switch fi53 {
                                    case 101:
                                        v55 = append(v55, 127)
                                    case 2147483647:
                                        _ = len(v55)
                                    case 42:
                                        fmt.Println(v12)
                                    }
                                    fmt.Println(v24, v2, v24)
                                } else {
                                    if true {
                                        v56 := GenS5{F0: "foo bar"}
                                        switch ("foo bar" + "a") {
                                        case "\n\t":
                                            switch (("foo bar" + "a") + (("\n\t" + "") + "hello")) {
                                            case "foo bar":
                                                var v57 uint16 = 65535
                                                _ = v57
                                            case "hello":
                                                if len(v55) >= 2 {
                                                    v55 = v55[1:len(v55)]
                                                }
                                            default:
                                                v3 = (0.0 + v4)
                                            }
                                        case "":
                                            switch "hello" {
                                            case "\n\t":
                                                var v58 uint64 = (v1 & 758)
                                                _ = v58
                                            default:
                                                var v59 float32 = (float32(0.1) / v4)
                                                _ = v59
                                            }
                                        }
                                        _ = v56
                                    }
                                    v60 := GenS6{F0: v2, F1: v6, F2: 2147483647}
                                    switch "foo bar" {
                                    case "a":
                                        v61 := GenS5{F0: ""}
                                        _ = v61
                                    case "hello":
                                        v4 -= float32(float64((float32(3.14159) + -2.25)))
                                    default:
                                        v55 = append(v55, (-(fi13)))
                                    }
                                    _ = v60
                                }
                            case 860:
                                _ = len(v5)
                            case 2147483647:
                                var v62 uint64 = (^(v1))
                                _ = v62
                            default:
                                if false {
                                    v1 -= v1
                                }
                            }
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                    }
                }
                v63 := []int{}
                _ = v24
                _ = v63
            }
            v2 = v2
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    for fi64 := 0; fi64 < 15; fi64++ {
        if true {
            v65 := GenS5{F0: "a"}
            var v66 int = (fi64 >> ((52) & 63))
            _ = v65
            _ = v66
        }
        _ = fi64
    }
    var v67 float32 = 1e10
    v68 := []int{128, 2147483647, 32768, -9223372036854775808}
    for fi69 := 0; fi69 < 13; fi69++ {
        v6 -= (v6 / ((42) | 1))
        _ = fi69
    }
    if (v2 >= v2) {
        var v70 uint32 = 42
        _ = v70
    } else {
        var v71 int16 = (v2 >> ((15) & 63))
        v72 := GenS1{F0: int64(v6), F1: 1}
        _ = v71
        _ = v72
    }
    if true {
        var v73 string = "hello"
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("slicelen v68:", len(v68), v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
}
