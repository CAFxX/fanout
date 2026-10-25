package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS7 struct { F0 string; F1 uint16; F2 int64 }
func genFn0(p0 int16, p1 int, p2 int32) int32 {
    fmt.Println(p0, p0, p2)
    var v4 int16 = p0
    switch ("foo bar" + ("foo bar" + "a")) {
    case "\n\t":
        var v5 int = p1
        _ = v5
    case "a":
        var v6 int64 = 0
        _ = v6
    case "hello":
        p0 = (v4 % (((p0 & v4)) | 1))
    default:
        p1 += 9223372036854775807
    }
    if (false || (!(false || false))) {
        v7 := []int{p1, p1, (p1 + p1), (p1 - 2147483647)}
        _ = len(v7)
        _ = v7
    }
    v8 := GenS1{F0: (p1 - p1), F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return (-(p2))
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = -1
    var v3 int64 = v2
    var v9 string = (("foo bar" + ("foo bar" + "hello")) + "\n\t")
    v10 := GenS2{F0: -1, F1: -32768, F2: 65535}
    v9 = v9
    var v11 bool = true
    var v12 uint32 = 927
    v3 = (v3 ^ (v3 - (-(v3))))
    if (v3 < v2) {
        v13 := GenS2{F0: 1, F1: -162, F2: 0}
        if false {
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 19; fi15++ {
                    fmt.Println(v9)
                    var v16 uint64 = 136
                    fmt.Println(v2, v11, v12)
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
        } else {
            var v17 float64 = (1.5 * -2.25)
            var v18 uint = 4294967295
            _ = v17
            _ = v18
        }
        if ("foo bar" <= v9) {
            if v11 {
                switch 2147483647 {
                case 128:
                    v19 := GenS3{F0: -99, F1: 60, F2: 666}
                    _ = v19
                case -614:
                    if (!v11) {
                        if v11 {
                            for fi20 := 0; fi20 < 18; fi20++ {
                                var v21 uint64 = 10
                                if v11 {
                                    var v22 float64 = 1.5
                                    _ = v22
                                }
                                _ = fi20
                                _ = v21
                            }
                        } else {
                            if (((127 <= 101) && (767 >= 1)) && true) {
                                var v23 string = (v9 + ("hello" + (v9 + "")))
                                _ = v23
                            }
                            switch 32768 {
                            case 42:
                                if (!v11) {
                                    switch 128 {
                                    case 0:
                                        fmt.Println(v9)
                                    default:
                                        var v24 uint32 = (^(v12))
                                        _ = v24
                                    }
                                } else {
                                    v11 = (!false)
                                    switch 4294967295 {
                                    case 4294967295:
                                        var v25 float32 = (((float32((-2.25 + 0.1)) - 0.1) - 0.0) - 1.5)
                                        _ = v25
                                    case -9223372036854775808:
                                        v26 := GenS1{F0: 127, F1: (v2 - (v3 >> ((4294967295) & 63)))}
                                        _ = v26
                                    case 32767:
                                        var v27 uint = 2
                                        _ = v27
                                    }
                                    var v28 uint16 = 42
                                    _ = v28
                                }
                            case 127:
                                v9 = v9
                            case 758:
                                var v29 int32 = -641
                                _ = v29
                            default:
                                fmt.Println(v2, v11, v2)
                            }
                        }
                        v30 := GenS3{F0: -128, F1: 1, F2: 743}
                        var v31 uint64 = 472
                        _ = v30
                        _ = v31
                    } else {
                        for fi32 := 0; fi32 < 19; fi32++ {
                            var v33 int16 = 1
                            var v34 int32 = -1
                            for fi35 := 0; fi35 < 6; fi35++ {
                                if v11 {
                                    if (fi35 < fi35) {
                                        v33 -= v33
                                    } else {
                                        if (false && true) {
                                            var v36 int = 902
                                            var v37 uint64 = 580
                                            v38 := GenS2{F0: -589, F1: (v33 ^ int16(v3)), F2: 255}
                                            var v39 uint64 = v37
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        }
                                        var v40 uint16 = 1
                                        v11 = (!false)
                                        _ = v40
                                    }
                                    if (true || false) {
                                        v41 := GenS4{F0: 42, F1: v12, F2: 42}
                                        var v42 int8 = 80
                                        var v43 uint = 65535
                                        var v44 string = v9
                                        var v45 string = "hello"
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                }
                                var v46 string = (v9 + "\n\t")
                                v47 := GenS5{F0: (fi32 << ((0) & 63))}
                                _ = fi35
                                _ = v46
                                _ = v47
                            }
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        if false {
                            _ = len(v1)
                        }
                        var v48 uint64 = 1
                        _ = v48
                    }
                default:
                    for fi49 := 0; fi49 < 5; fi49++ {
                        v50 := []int{fi49, fi49, (fi49 / ((1) | 1)), 1}
                        _ = fi49
                        _ = v50
                    }
                }
            } else {
                fmt.Println(v2)
                fmt.Println(v12)
            }
            var v51 uint32 = (v12 / ((v12) | 1))
            for fi52 := 0; fi52 < 12; fi52++ {
                var v53 int16 = 263
                _ = fi52
                _ = v53
            }
            _ = v51
        }
        _ = v13
    }
    delete(v1, ("\n\t" + ""))
    var v54 uint32 = (v12 & (v12 | (v12 - (v12 | v12))))
    v55 := GenS6{F0: v12, F1: 0, F2: 83}
    var v56 uint32 = 369
    fmt.Println(v11, v2)
    var v57 int64 = -9223372036854775808
    var v58 int32 = 2147483647
    for fi59 := 0; fi59 < 17; fi59++ {
        if true {
            v60 := GenS1{F0: 107, F1: v2}
            v2 = (v2 / (((v57 | v3)) | 1))
            _ = v60
        }
        v61 := GenS7{F0: (("\n\t" + v9) + v9), F1: 32767, F2: (v2 << ((255) & 63))}
        switch "" {
        case "a":
            fmt.Println(v12)
        case "foo bar":
            var v62 uint64 = 61
            _ = v62
        }
        _ = fi59
        _ = v61
    }
    var v63 string = v9
    v57 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v63:", v63)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v63
    _ = v9
}
