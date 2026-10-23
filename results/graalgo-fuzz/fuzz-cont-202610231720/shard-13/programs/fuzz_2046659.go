package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS5 struct { F0 int32; F1 uint64 }
func genFn0(p0 string) bool {
    p0 = ((p0 + "\n\t") + p0)
    p0 = "\n\t"
    var v7 int16 = -171
    _ = p0
    _ = v7
    return false
}

func main() {
    var v1 string = ("\n\t" + ("hello" + ((((("hello" + (("foo bar" + "") + (("" + "") + ("foo bar" + ("" + "foo bar"))))) + "foo bar") + "a") + "\n\t") + "foo bar")))
    var v2 int8 = 127
    var v3 float64 = (float64(((0.0 + 1e10) - 1.5)) + -2.25)
    var v4 uint8 = 255
    v5 := make(map[string]int)
    var v6 float64 = v3
    v8 := []int{32767, 9223372036854775807}
    var v9 float32 = 1.5
    if ((v9 / (1.5 * (v9 / v9))) <= v9) {
        if (!false) {
            var v10 float64 = ((1e10 + 1e10) - (3.14159 - (v6 / float64(v9))))
            _ = v10
        } else {
            switch "a" {
            case "foo bar":
                _ = len(v8)
            case "a":
                for fi11 := 0; fi11 < 2; fi11++ {
                    v12 := GenS1{F0: v4}
                    _ = fi11
                    _ = v12
                }
            case "hello":
                v6 = 1e10
            }
            if (false || (-181 <= -409)) {
                if true {
                    v13 := GenS1{F0: 128}
                    if len(v8) > 0 {
                        v8[0] = 32767
                    }
                    if true {
                        v14 := GenS1{F0: 1}
                        v1 = (v1 + v1)
                        fmt.Println(v6, v6)
                        _ = v14
                    } else {
                        v6 = v6
                        fmt.Println(v4, v6)
                        if (float32((v3 - ((float64((float32(0.1) - v9)) + -2.25) + v3))) != v9) {
                            for fi15 := 0; fi15 < 13; fi15++ {
                                if (1.5 != 1.5) {
                                    for fi16 := 0; fi16 < 9; fi16++ {
                                        var v17 bool = false
                                        _ = fi16
                                        _ = v17
                                    }
                                } else {
                                    fmt.Println(v6)
                                }
                                if (!true) {
                                    var v18 int64 = 42
                                    switch (^(fi15)) {
                                    case 558:
                                        v19 := GenS1{F0: v4}
                                        _ = v19
                                    case 42:
                                        for fi20 := 0; fi20 < 6; fi20++ {
                                            fmt.Println(v2)
                                            v6 = 3.14159
                                            switch fi20 {
                                            case -9223372036854775808:
                                                v21 := GenS1{F0: (v4 & (v4 & (v4 & v4)))}
                                                _ = v21
                                            }
                                            _ = fi20
                                        }
                                    default:
                                        for fi22 := 0; fi22 < 1; fi22++ {
                                            if (42 == 42) {
                                                v23 := GenS2{F0: 9223372036854775807, F1: v4}
                                                v18 -= (-(v18))
                                                _ = v23
                                            }
                                            if (1 < 763) {
                                                var v24 uint16 = 834
                                                if ((525 <= 0) && true) {
                                                    var v25 int8 = 110
                                                    if true {
                                                        fmt.Println(v9)
                                                        if len(v8) > 0 {
                                                            v8[0] = (fi15 >> ((65535) & 63))
                                                        }
                                                    }
                                                    v26 := GenS2{F0: v18, F1: v4}
                                                    _ = v25
                                                    _ = v26
                                                } else {
                                                    if false {
                                                        var v27 uint8 = (^(v4))
                                                        v5["\n\t"] = 127
                                                        _ = v27
                                                    } else {
                                                        var v28 uint16 = (v24 * v24)
                                                        _ = len(v5)
                                                        _ = v28
                                                    }
                                                }
                                                v18 = int64(fi15)
                                                _ = v24
                                            }
                                            _ = fi22
                                        }
                                    }
                                    fmt.Println(v9)
                                    _ = v18
                                }
                                switch ((v1 + "foo bar") + ("a" + (v1 + v1))) {
                                case "foo bar":
                                    v29 := GenS3{F0: v2, F1: -1, F2: (v4 - v4)}
                                    _ = v29
                                case "":
                                    v9 = 0.0
                                }
                                _ = fi15
                            }
                            var v30 string = ("hello" + "\n\t")
                            _ = v30
                        } else {
                            if (false && (1 >= 0)) {
                                fmt.Println(v6, v6)
                                v31 := GenS3{F0: (v2 & 127), F1: 1, F2: v4}
                                fmt.Println(v1, v6)
                                _ = v31
                            }
                        }
                    }
                    _ = v13
                } else {
                    v4 += (^(v4))
                }
            }
        }
    } else {
        v4 = v4
        v32 := v5[v1]
        v4 = 255
        _ = v32
    }
    if true {
        fmt.Println(v9, v4)
    }
    v4 = (v4 * v4)
    if (v4 < 0) {
        if (true || false) {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v9 = (v9 / 0.1)
        } else {
            v33 := GenS2{F0: 0, F1: 205}
            for fi34 := 0; fi34 < 13; fi34++ {
                for fi35 := 0; fi35 < 5; fi35++ {
                    if true {
                        v36 := GenS3{F0: (v2 << ((2) & 63)), F1: 249, F2: 128}
                        v4 -= (v4 ^ v4)
                        if true {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                        _ = v36
                    }
                    var v37 uint32 = 2147483648
                    _ = len(v8)
                    _ = fi35
                    _ = v37
                }
                if (true && false) {
                    if (!((1 != 1) && false)) {
                        v38 := v5[""]
                        var v39 float64 = v6
                        _ = v38
                        _ = v39
                    } else {
                        var v40 uint8 = v4
                        fmt.Println(fi34)
                        _ = v40
                    }
                    for fi41 := 0; fi41 < 16; fi41++ {
                        fmt.Println(v1)
                        v8 = append(v8, fi41)
                        _ = fi41
                    }
                } else {
                    var v42 string = v1
                    for fi43 := 0; fi43 < 1; fi43++ {
                        var v44 string = "\n\t"
                        v45 := GenS3{F0: 1, F1: 0, F2: v4}
                        v9 += 0.1
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                    switch v1 {
                    case "\n\t":
                        v46 := GenS1{F0: 127}
                        _ = v46
                    default:
                        var v47 uint32 = 0
                        _ = v47
                    }
                    _ = v42
                }
                _ = fi34
            }
            _ = v33
        }
        switch -9223372036854775808 {
        case 32767:
            if (!(!true)) {
                v6 -= v6
            } else {
                fmt.Println(v3, v4)
                v48 := GenS3{F0: (-(v2)), F1: -1, F2: v4}
                _ = v48
            }
        case 9223372036854775807:
            if (-1 >= 0) {
                var v49 int8 = v2
                for fi50 := 0; fi50 < 8; fi50++ {
                    if (true && ("\n\t" <= "foo bar")) {
                        for fi51 := 0; fi51 < 5; fi51++ {
                            v3 += v3
                            var v52 float32 = (0.0 - (v9 + float32(((v6 + (v3 * 3.14159)) - (v3 / float64((v9 * v9)))))))
                            _ = fi51
                            _ = v52
                        }
                        v4 = (v4 % ((49) | 1))
                        var v53 int8 = 121
                        _ = v53
                    } else {
                        var v54 string = v1
                        _ = v54
                    }
                    _ = fi50
                }
                if false {
                    fmt.Println(v4, v3, v1)
                }
                _ = v49
            }
        case 2147483647:
            var v55 uint = 4294967295
            _ = v55
        }
        var v56 uint64 = 384
        _ = v56
    }
    delete(v5, v1)
    var v57 uint64 = 4294967295
    _ = len(v8)
    v58 := GenS4{F0: 42, F1: (v57 >> ((42) & 63)), F2: -788}
    v1 = ("hello" + (((v1 + v1) + v1) + v1))
    v1 = "\n\t"
    if ((v6 + ((1.5 - v3) + v3)) != (v3 / v3)) {
        if (0 == 767) {
            var v59 uint32 = 534
            _ = v59
        } else {
            var v60 string = v1
            switch "" {
            case "hello":
                if true {
                    fmt.Println(v6)
                    if (true && (0 > -2147483648)) {
                        v4 = v4
                        v1 = (((v60 + ("a" + (v60 + v1))) + "") + v60)
                    } else {
                        var v61 int16 = 591
                        _ = len(v5)
                        _ = v61
                    }
                    v62 := GenS5{F0: 0, F1: 18446744073709551615}
                    _ = v62
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        v57 = v57
                    }
                    if len(v8) > 0 {
                        v8[0] = 32768
                    }
                }
            case "foo bar":
                v63 := GenS2{F0: 341, F1: (v4 & v4)}
                _ = v63
            case "a":
                fmt.Println(v6, v3)
            default:
                v4 -= 127
            }
            var v64 int8 = (v2 << ((996) & 63))
            _ = v60
            _ = v64
        }
        v65 := GenS5{F0: 503, F1: (v57 * v57)}
        _ = v65
    }
    v66 := GenS4{F0: 0, F1: 453, F2: 803}
    var v67 uint32 = 42
    var v68 uint = 548
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v8
    _ = v9
}
