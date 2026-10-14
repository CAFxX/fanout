package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int16; F2 int32 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int) uint16 {
    v7 := GenS1{F0: (!true), F1: 1, F2: "\n\t"}
    v8 := GenS1{F0: false, F1: 153, F2: "\n\t"}
    _ = p0
    _ = v7
    _ = v8
    return 65535
}

func main() {
    v1 := []string{"a"}
    var v2 int = -339
    var v3 string = "\n\t"
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    var v6 uint = 126
    var v9 int32 = -725
    var v10 int64 = 0
    fmt.Println(v9, v2)
    var v11 uint32 = 667
    if true {
        var v12 float64 = 0.1
        var v13 int32 = (v9 * v9)
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: (int64(v11) <= 0), F1: 1, F2: v3}
            _ = fi14
            _ = v15
        }
        if (460 != 1) {
            if ((v10 & (v10 << ((65535) & 63))) == (v10 / ((v10) | 1))) {
                var v16 uint16 = 255
                v11 = 2147483647
                v17 := GenS2{F0: 42}
                _ = v16
                _ = v17
            } else {
                var v18 float32 = float32((-2.25 + -2.25))
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch ("" + v3) {
                    case "a":
                        if (v3 < "\n\t") {
                            if false {
                                if ((v3 + "hello") != "a") {
                                    v10 = v10
                                    var v20 uint16 = 2
                                    fmt.Println(v10)
                                    _ = v20
                                }
                            }
                            for fi21 := 0; fi21 < 3; fi21++ {
                                v9 += (^(v9))
                                v22 := GenS1{F0: true, F1: 2, F2: v3}
                                _ = fi21
                                _ = v22
                            }
                        }
                    case "hello":
                        var v23 bool = true
                        _ = v23
                    case "\n\t":
                        var v24 uint32 = v11
                        _ = v24
                    }
                    if false {
                        fmt.Println(v10, v9)
                        fmt.Println(v2, v10, v2)
                    } else {
                        if ((255 >= 42) && false) {
                            v6 = (^(v6))
                            if (true && false) {
                                v10 += 1
                                if ((^(v11)) < v11) {
                                    v10 += v10
                                    for fi25 := 0; fi25 < 9; fi25++ {
                                        if (0 >= 128) {
                                            var v26 uint = (v6 / ((910) | 1))
                                            var v27 uint16 = genFn0((fi25 >> ((0) & 63)))
                                            v1 = append(v1, "hello")
                                            _ = v26
                                            _ = v27
                                        } else {
                                            v9 = v9
                                            fmt.Println(fi19, v6)
                                        }
                                        var v28 int16 = -1
                                        v29 := v5[(fi25 >> ((4294967295) & 63))]
                                        _ = fi25
                                        _ = v28
                                        _ = v29
                                    }
                                }
                                var v30 uint16 = 32767
                                _ = v30
                            }
                            var v31 uint16 = 0
                            _ = v31
                        } else {
                            if true {
                                var v32 uint64 = 1
                                v33 := GenS2{F0: (v32 & 42)}
                                for fi34 := 0; fi34 < 7; fi34++ {
                                    if len(v1) > 0 {
                                        v1[0] = v3
                                    }
                                    _ = fi34
                                }
                                _ = v32
                                _ = v33
                            } else {
                                if ((v10 > (v10 * v10)) || (!true)) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    if ((false && false) || (false || false)) {
                                        for fi35 := 0; fi35 < 14; fi35++ {
                                            var v36 float64 = 1e10
                                            var v37 uint32 = v11
                                            _ = fi35
                                            _ = v36
                                            _ = v37
                                        }
                                        var v38 int8 = 42
                                        _ = v38
                                    }
                                    v2 += 32767
                                    v9 = (v9 & (v9 >> ((65535) & 63)))
                                }
                                if ((v6 << ((604) & 63)) != v6) {
                                    var v39 uint16 = genFn0((v2 >> (((v6 % ((v6) | 1))) & 63)))
                                    _ = v39
                                }
                            }
                        }
                        for fi40 := 0; fi40 < 11; fi40++ {
                            v11 -= v11
                            _ = fi40
                        }
                    }
                    _ = fi19
                }
                switch v2 {
                case 9223372036854775807:
                    v41 := GenS2{F0: 18446744073709551615}
                    _ = v41
                case -570:
                    var v42 int8 = -1
                    _ = v42
                case -1:
                    v10 += v10
                }
                _ = v18
            }
            var v43 int64 = v10
            v43 -= v10
            _ = v43
        } else {
            var v44 int32 = -2147483648
            v6 += v6
            v6 = v6
            _ = v44
        }
        if (-82 > 1) {
            v1 = append(v1, (v3 + (((("a" + v3) + "\n\t") + v3) + "foo bar")))
            if len(v1) > 0 {
                v1[0] = v3
            }
            for fi45 := 0; fi45 < 3; fi45++ {
                v46 := GenS3{F0: 255, F1: -1, F2: (v9 / (((^(v9))) | 1))}
                v47 := GenS4{F0: 1, F1: v9}
                if (0 != 65535) {
                    v3 = v3
                    v10 += 0
                    for fi48 := 0; fi48 < 4; fi48++ {
                        v2 += (fi45 / ((v2) | 1))
                        var v49 int32 = (^(v9))
                        v50 := GenS5{F0: (v11 - v11)}
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                } else {
                    var v51 uint16 = 42
                    var v52 uint64 = 18446744073709551615
                    var v53 uint8 = 128
                    _ = v51
                    _ = v52
                    _ = v53
                }
                _ = fi45
                _ = v46
                _ = v47
            }
        }
    }
    if true {
        var v54 uint = 42
        _ = v54
    }
    if false {
        var v55 int64 = v10
        _ = v55
    }
    var v56 bool = (!true)
    for fi57 := 0; fi57 < 10; fi57++ {
        var v58 uint64 = 47
        v59 := GenS1{F0: v56, F1: 128, F2: "hello"}
        _ = fi57
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 1; fi60++ {
        var v61 int32 = (v9 | (^(v9)))
        v3 = (("a" + (v3 + v3)) + "")
        v11 += (^(v11))
        _ = fi60
        _ = v61
    }
    v56 = (1.5 != ((1e10 - 0.1) + (-2.25 * 3.14159)))
    for fi62 := 0; fi62 < 16; fi62++ {
        var v63 uint16 = 1
        if len(v1) > 0 {
            v1[0] = ((v3 + "a") + v3)
        }
        _ = fi62
        _ = v63
    }
    v6 = (v6 | (v6 ^ (^(v6))))
    fmt.Println(v6, v11)
    v56 = ((1.5 / 1e10) < 0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v9
}
