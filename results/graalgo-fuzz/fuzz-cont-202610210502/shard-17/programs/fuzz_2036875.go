package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int16; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 uint8 }
func genFn0(p0 string, p1 int64) uint32 {
    var v5 bool = false
    var v6 int32 = 0
    v7 := GenS1{F0: 32767}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1() uint16 {
    var v8 int16 = -32768
    var v9 int8 = -1
    _ = v8
    _ = v9
    return 964
}

func main() {
    var v1 bool = (0 < 65535)
    v2 := make(map[int]bool)
    v3 := []string{}
    var v4 uint8 = 1
    switch 1 {
    case 128:
        v1 = (-93 > 48)
    default:
        var v10 uint16 = 255
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 uint16 = 190
        _ = fi11
        _ = v12
    }
    if ((v4 | (v4 | v4)) < v4) {
        v13 := GenS1{F0: 861}
        _ = v13
    }
    v14 := GenS1{F0: 951}
    var v15 float32 = 3.14159
    if (v1 && true) {
        fmt.Println(v4)
    } else {
        if ("hello" <= "hello") {
            var v16 string = ("a" + "")
            var v17 int16 = 32767
            v2[-697] = true
            _ = v16
            _ = v17
        }
    }
    v15 -= v15
    v18 := GenS1{F0: -51}
    switch 1 {
    case 32767:
        for fi19 := 0; fi19 < 14; fi19++ {
            if (1e10 == 1e10) {
                v20 := GenS2{F0: -9223372036854775808, F1: "\n\t"}
                var v21 string = (("hello" + (("\n\t" + "a") + ("" + "\n\t"))) + "a")
                if v1 {
                    for fi22 := 0; fi22 < 7; fi22++ {
                        v21 = ("\n\t" + ("" + v21))
                        if (!((-1 > 127) || (65535 != 388))) {
                            v21 = (("foo bar" + "hello") + "foo bar")
                            var v23 float32 = v15
                            v1 = (fi22 != fi22)
                            v4 = 2
                            fmt.Println(v21, v1)
                            _ = v23
                        } else {
                            v24 := GenS3{F0: -9223372036854775808}
                            var v25 uint64 = 1
                            delete(v2, fi22)
                            _ = v24
                            _ = v25
                        }
                        _ = fi22
                    }
                } else {
                    v26 := GenS1{F0: 32767}
                    v27 := GenS3{F0: 831}
                    switch (((v21 + (("\n\t" + ("a" + "foo bar")) + v21)) + v21) + v21) {
                    case "\n\t":
                        var v28 float64 = 1e10
                        _ = v28
                    }
                    _ = v26
                    _ = v27
                }
                _ = v20
                _ = v21
            }
            _ = fi19
        }
    case 2147483647:
        v2[-46] = v1
    default:
        var v29 uint = 4294967295
        _ = v29
    }
    v4 += v4
    v30 := GenS1{F0: 1000}
    fmt.Println(v4, v1)
    for fi31 := 0; fi31 < 9; fi31++ {
        switch "" {
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = (("hello" + "foo bar") + "foo bar")
            }
        case "a":
            delete(v2, (fi31 / (((fi31 - fi31)) | 1)))
        }
        if v1 {
            _ = len(v3)
        } else {
            v32 := v2[fi31]
            if v32 {
                if v32 {
                    var v33 bool = v32
                    _ = len(v2)
                    v2[(-(fi31))] = (-676 < 277)
                    _ = v33
                } else {
                    v34 := GenS4{F0: v4, F1: (v4 >> ((18446744073709551615) & 63))}
                    _ = v34
                }
                for fi35 := 0; fi35 < 3; fi35++ {
                    v32 = v32
                    var v36 uint64 = 232
                    delete(v2, fi35)
                    _ = fi35
                    _ = v36
                }
            } else {
                if (2147483647 != -38) {
                    if len(v3) > 0 {
                        v3[0] = ("\n\t" + ("hello" + ("foo bar" + ("" + "foo bar"))))
                    }
                    v1 = v1
                } else {
                    var v37 uint64 = 2
                    fmt.Println(v1, v15)
                    _ = v37
                }
            }
            v32 = v1
            _ = v32
        }
        _ = fi31
    }
    for fi38 := 0; fi38 < 5; fi38++ {
        v39 := v2[fi38]
        for fi40 := 0; fi40 < 8; fi40++ {
            v4 = 1
            switch ("" + "") {
            case "":
                if v1 {
                    if v1 {
                        v41 := GenS4{F0: v4, F1: (v4 / ((v4) | 1))}
                        switch (fi40 + fi38) {
                        case 4294967295:
                            v42 := GenS3{F0: -1}
                            _ = v42
                        default:
                            if (true || v1) {
                                if v1 {
                                    fmt.Println(v15, v39, v4)
                                    v43 := GenS5{F0: 42, F1: 2, F2: 0}
                                    var v44 string = ("hello" + "\n\t")
                                    _ = v43
                                    _ = v44
                                } else {
                                    if v1 {
                                        v2[(fi40 % ((fi40) | 1))] = true
                                        v4 = v4
                                    } else {
                                        var v45 uint32 = 4294967295
                                        if len(v3) > 0 {
                                            v3[0] = ""
                                        }
                                        if v1 {
                                            _ = len(v3)
                                            v45 = genFn0("a", int64(v4))
                                        } else {
                                            _ = len(v2)
                                            if (42 <= 877) {
                                                var v46 uint8 = v4
                                                v15 = (-2.25 - 1.5)
                                                var v47 uint16 = 65535
                                                _ = v46
                                                _ = v47
                                            } else {
                                                v4 -= (v4 ^ (v4 | (^(v4))))
                                            }
                                        }
                                        _ = v45
                                    }
                                }
                            }
                            var v48 uint8 = v4
                            _ = v48
                        }
                        v1 = true
                        _ = v41
                    } else {
                        fmt.Println(fi40, fi38)
                        fmt.Println(v15, v15)
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                    }
                    v49 := GenS4{F0: v4, F1: (v4 * v4)}
                    _ = v49
                } else {
                    v1 = (-128 >= 96)
                }
            case "foo bar":
                if v39 {
                    if (-128 <= 127) {
                        v50 := GenS6{F0: 282, F1: (fi38 % (((fi40 | (fi40 << ((42) & 63)))) | 1))}
                        fmt.Println(v1, v39)
                        _ = v50
                    }
                } else {
                    v51 := GenS1{F0: 1000}
                    _ = len(v3)
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    fmt.Println(v15, v4)
                    if (false || (65535 != genFn1())) {
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        v52 := GenS5{F0: -1, F1: 251, F2: int16(fi40)}
                        fmt.Println(fi38)
                        _ = v52
                    } else {
                        var v53 uint = 132
                        _ = v53
                    }
                    v54 := GenS4{F0: (v4 << ((18446744073709551615) & 63)), F1: v4}
                    _ = v51
                    _ = v54
                }
                var v55 string = ""
                fmt.Println(v15, v39)
                _ = v55
            default:
                v56 := GenS5{F0: -467, F1: 180, F2: -32768}
                _ = v56
            }
            v57 := GenS4{F0: 127, F1: v4}
            var v58 int64 = 42
            _ = fi40
            _ = v57
            _ = v58
        }
        v15 = -2.25
        _ = fi38
        _ = v39
    }
    _ = len(v3)
    var v59 int64 = 812
    switch 9223372036854775807 {
    case 128:
        var v60 uint32 = genFn0("", (v59 % ((v59) | 1)))
        _ = v60
    case 1:
        if v1 {
            if true {
                v4 = (v4 % ((v4) | 1))
            } else {
                v61 := GenS7{F0: (v4 << ((42) & 63))}
                if (0.1 != (-2.25 / (0.1 - (1.5 - (float64((v15 / (((v15 - v15) - v15) - ((v15 + v15) + v15)))) + float64((0.0 / 3.14159))))))) {
                    var v62 int8 = 76
                    _ = v62
                }
                _ = v61
            }
            for fi63 := 0; fi63 < 3; fi63++ {
                fmt.Println(v59)
                _ = fi63
            }
        }
    }
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v59
}
