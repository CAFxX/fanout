package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 bool }
func genFn0(p0 uint, p1 uint64, p2 uint8) uint32 {
    v6 := []int{42, 4294967295}
    v7 := GenS1{F0: false, F1: (!(422 <= 32767))}
    var v8 int64 = -9223372036854775808
    var v9 float32 = 1.5
    var v10 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := []string{"\n\t", ("hello" + "\n\t"), "foo bar"}
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    var v4 uint = 42
    var v5 uint16 = 255
    _ = len(v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 := GenS2{F0: 65535}
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS3{F0: genFn0((v4 * 255), (v3 - 2), 247), F1: -532, F2: (^(v3))}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v1 = append(v1, "foo bar")
        if (3.14159 > 3.14159) {
            if false {
                switch "a" {
                case "foo bar":
                    v15 := GenS2{F0: 65535}
                    _ = v15
                case "\n\t":
                    fmt.Println(v3, fi14)
                case "a":
                    v4 += (v4 % ((v4) | 1))
                }
            } else {
                if true {
                    switch "hello" {
                    case "hello":
                        var v16 int64 = 760
                        _ = v16
                    case "\n\t":
                        var v17 int = (fi14 % ((-9223372036854775808) | 1))
                        _ = v17
                    case "":
                        if (v5 > v5) {
                            v2[fi14] = (^(fi14))
                        }
                    default:
                        v18 := GenS1{F0: ((127 <= 0) || (366 == -1)), F1: (!true)}
                        _ = v18
                    }
                    if false {
                        if false {
                            _ = len(v1)
                        }
                        var v19 uint32 = 4294967295
                        _ = v19
                    } else {
                        if ((true && ((v4 & 42) >= v4)) || (!false)) {
                            var v20 uint8 = 226
                            var v21 string = "foo bar"
                            _ = v20
                            _ = v21
                        } else {
                            var v22 uint16 = (^(v5))
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            _ = v22
                        }
                        if (255 > 96) {
                            var v23 uint = 18446744073709551615
                            v5 = v5
                            switch int(v5) {
                            case -617:
                                if true {
                                    for fi24 := 0; fi24 < 16; fi24++ {
                                        var v25 int64 = 1
                                        var v26 uint8 = 128
                                        v26 -= v26
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                } else {
                                    v2[fi14] = fi14
                                    v27 := v2[-336]
                                    v28 := GenS1{F0: true, F1: (!true)}
                                    _ = v27
                                    _ = v28
                                }
                            case 128:
                                switch "" {
                                case "foo bar":
                                    if ((1 <= 100) && false) {
                                        v4 = (v23 % ((v23) | 1))
                                        v29 := GenS4{F0: 2, F1: false}
                                        _ = v29
                                    } else {
                                        if (true && ((2 == 158) || false)) {
                                            v30 := GenS5{F0: "", F1: "a"}
                                            var v31 uint64 = v3
                                            var v32 string = (("\n\t" + "\n\t") + "a")
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        } else {
                                            var v33 int = (fi14 / (((^(fi14))) | 1))
                                            v34 := GenS2{F0: 0}
                                            var v35 float32 = (0.0 / 0.1)
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        fmt.Println(v3, v23)
                                        v36 := GenS6{F0: genFn0(v23, 762, uint8(v3)), F1: 1000, F2: v3}
                                        _ = v36
                                    }
                                case "":
                                    for fi37 := 0; fi37 < 5; fi37++ {
                                        if false {
                                            var v38 int8 = 42
                                            if (v3 == v3) {
                                                v39 := GenS1{F0: (false && false), F1: (!(42 >= 1))}
                                                v1 = append(v1, (("foo bar" + (("a" + "\n\t") + "foo bar")) + "hello"))
                                                v5 = (v5 ^ 255)
                                                v23 -= v23
                                                _ = v39
                                            }
                                            if (49 < 87) {
                                                switch (fi37 << ((v4) & 63)) {
                                                case 9223372036854775807:
                                                    if ((v38 - (v38 << ((v4) & 63))) > 54) {
                                                        var v40 float64 = 1.5
                                                        v5 -= (v5 ^ (v5 & (v5 << (((^(v4))) & 63))))
                                                        v41 := GenS1{F0: (v40 == (1.5 - v40)), F1: true}
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                    fmt.Println(v38)
                                                    if (-1 == v38) {
                                                        v38 += (v38 << (((v4 >> ((v23) & 63))) & 63))
                                                        v42 := GenS3{F0: 733, F1: 2147483647, F2: v3}
                                                        v5 -= 142
                                                        _ = v42
                                                    } else {
                                                        v38 = int8(v3)
                                                        v38 += (v38 - 112)
                                                    }
                                                    if (v3 > (v3 >> ((v23) & 63))) {
                                                        var v43 string = ("hello" + "")
                                                        _ = v43
                                                    } else {
                                                        v3 += uint64(v4)
                                                    }
                                                default:
                                                    var v44 uint64 = v3
                                                    _ = v44
                                                }
                                                var v45 uint16 = 876
                                                _ = v45
                                            }
                                            _ = v38
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v3 += (^(v3))
                                            v23 = v4
                                        }
                                        v23 = v4
                                        _ = fi37
                                    }
                                case "a":
                                    if (0.1 <= 3.14159) {
                                        var v46 uint64 = 798
                                        var v47 string = "hello"
                                        var v48 float32 = ((0.1 * 3.14159) + (-2.25 - float32((0.1 * 0.0))))
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                            }
                            _ = v23
                        }
                        v49 := GenS7{F0: 2, F1: (fi14 * (fi14 & 127))}
                        _ = v49
                    }
                    fmt.Println(v4, v4)
                } else {
                    v50 := GenS8{F0: 9223372036854775807}
                    _ = v50
                }
                v51 := v2[fi14]
                var v52 int8 = 42
                _ = v51
                _ = v52
            }
        }
        _ = fi14
    }
    var v53 float32 = 1.5
    if ((0 != (^(v4))) || true) {
        if ((-2.25 != v53) || (186 <= -2147483648)) {
            v1 = append(v1, ("hello" + "a"))
        } else {
            v54 := GenS9{F0: 32767, F1: -610}
            _ = v54
        }
        var v55 int = -257
        _ = v55
    } else {
        v56 := GenS10{F0: ((float32(float64(v53)) - (v53 + (float32(0.1) - 1e10))) != float32(1.5))}
        v4 += (v4 % ((v4) | 1))
        v53 = 3.14159
        _ = v56
    }
    if (false || true) {
        fmt.Println(v5, v3)
        if (!(42 <= -358)) {
            if true {
                var v57 uint8 = 128
                if true {
                    var v58 bool = false
                    _ = len(v1)
                    v4 = v4
                    _ = v58
                } else {
                    var v59 float64 = 3.14159
                    _ = v59
                }
                _ = v57
            } else {
                v60 := GenS8{F0: 424}
                _ = v60
            }
        }
    }
    switch "\n\t" {
    case "a":
        var v61 int64 = -1
        _ = v61
    case "\n\t":
        v5 = v5
    case "":
        if (!(-2.25 > (3.14159 - (float64(3.14159) + 0.0)))) {
            if true {
                _ = len(v2)
                fmt.Println(v5, v4, v4)
            } else {
                delete(v2, 9223372036854775807)
                var v62 uint32 = 65535
                v3 = v3
                _ = v62
            }
        }
    }
    v63 := v2[-2147483648]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v63
}
