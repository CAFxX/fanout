package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS7 struct { F0 string; F1 uint8; F2 uint8 }
type GenS8 struct { F0 uint8; F1 int32 }
func genFn0(p0 int16) uint {
    v6 := GenS1{F0: 530}
    v7 := []int{587, 32768}
    _ = p0
    _ = v6
    _ = v7
    return 851
}

func main() {
    var v1 bool = false
    v2 := []int{552}
    v3 := []uint32{0, 378, 865, 0}
    var v4 string = ((((("a" + ("" + "\n\t")) + "foo bar") + "hello") + "a") + "")
    v5 := make(map[string]int)
    var v8 uint32 = 2147483648
    v4 = ((((v4 + "foo bar") + v4) + "") + ((v4 + (v4 + "a")) + "a"))
    var v9 uint16 = 1
    v5[v4] = -1
    v10 := GenS2{F0: v1, F1: v8}
    switch "foo bar" {
    case "":
        if true {
            if (-9223372036854775808 >= 9223372036854775807) {
                v1 = v1
                v9 += v9
            }
        } else {
            var v11 string = v4
            _ = len(v2)
            if v1 {
                v4 = v4
            } else {
                for fi12 := 0; fi12 < 2; fi12++ {
                    if v1 {
                        fmt.Println(v11, fi12, v4)
                        for fi13 := 0; fi13 < 4; fi13++ {
                            v11 = (v4 + ("" + v4))
                            v8 += 65535
                            fmt.Println(fi13, v11, v4)
                            v14 := GenS2{F0: (uint16(v8) != v9), F1: v8}
                            _ = fi13
                            _ = v14
                        }
                    }
                    v11 = v11
                    var v15 uint8 = 0
                    _ = fi12
                    _ = v15
                }
                v9 = v9
            }
            _ = v11
        }
    case "a":
        v16 := GenS3{F0: 917, F1: "\n\t"}
        _ = v16
    case "foo bar":
        if (((true || (!true)) && (1.5 >= (1.5 + (3.14159 * -2.25)))) || (!(!true))) {
            var v17 int64 = -988
            var v18 uint16 = v9
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 int32 = 42
                v8 -= v8
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        }
    default:
        v4 = ("a" + "\n\t")
    }
    switch v4 {
    case "":
        for fi21 := 0; fi21 < 18; fi21++ {
            delete(v5, v4)
            delete(v5, v4)
            _ = fi21
        }
    }
    fmt.Println(v8)
    var v22 int16 = 32767
    if (false || true) {
        v4 = ("\n\t" + v4)
        for fi23 := 0; fi23 < 14; fi23++ {
            for fi24 := 0; fi24 < 18; fi24++ {
                v4 = "foo bar"
                fmt.Println(v22, v8)
                var v25 int64 = 534
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 1; fi26++ {
                v1 = ((^(v9)) != 299)
                if (false && false) {
                    v4 = (((v4 + "") + v4) + v4)
                    if v1 {
                        v27 := GenS4{F0: 0}
                        var v28 int64 = 0
                        var v29 int8 = 42
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS2{F0: v1, F1: uint32(v9)}
                        _ = v30
                    }
                    var v31 uint32 = (^(v8))
                    _ = v31
                } else {
                    v32 := GenS5{F0: (true && (!v1))}
                    v33 := v5[v4]
                    _ = v32
                    _ = v33
                }
                _ = fi26
            }
            for fi34 := 0; fi34 < 8; fi34++ {
                var v35 float32 = (3.14159 + -2.25)
                fmt.Println(fi34, fi34, v4)
                v9 += 2
                _ = fi34
                _ = v35
            }
            _ = fi23
        }
        var v36 float32 = (0.1 * -2.25)
        _ = v36
    } else {
        var v37 int16 = (v22 & v22)
        switch 830 {
        case 0:
            for fi38 := 0; fi38 < 7; fi38++ {
                for fi39 := 0; fi39 < 16; fi39++ {
                    fmt.Println(fi39, v8, v37)
                    var v40 bool = ((0.0 - (float64(((0.1 * (1.5 - 0.1)) + (0.0 * 1e10))) * 0.1)) == (float64(0.0) / 1.5))
                    v1 = v1
                    _ = fi39
                    _ = v40
                }
                switch "hello" {
                case "\n\t":
                    var v41 float32 = (1.5 * (0.1 / 1.0))
                    _ = v41
                case "a":
                    v22 = v22
                case "hello":
                    var v42 uint = 1
                    _ = v42
                }
                v9 += (v9 << ((2) & 63))
                _ = fi38
            }
        }
        var v43 int64 = -600
        _ = v37
        _ = v43
    }
    v22 = -137
    if len(v2) > 0 {
        v2[0] = 128
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        for fi45 := 0; fi45 < 13; fi45++ {
            if v1 {
                if len(v2) > 0 {
                    v2[0] = fi45
                }
                if len(v2) > 0 {
                    v2[0] = 127
                }
            } else {
                if ((0.1 - 0.0) >= 0.0) {
                    var v46 int64 = 1
                    var v47 int64 = v46
                    fmt.Println(fi45, fi44, v8)
                    fmt.Println(v1, fi45)
                    if true {
                        v48 := GenS5{F0: ((fi44 / ((fi44) | 1)) != (fi45 % (((^(fi44))) | 1)))}
                        if true {
                            v46 = 1
                        } else {
                            if (v9 > 65535) {
                                switch (v4 + v4) {
                                case "foo bar":
                                    v49 := GenS6{F0: 65535, F1: 42, F2: 2}
                                    _ = v49
                                }
                                switch "foo bar" {
                                case "hello":
                                    if true {
                                        delete(v5, ("" + v4))
                                        v50 := GenS7{F0: ("a" + (((("" + ((v4 + v4) + v4)) + v4) + ("foo bar" + (("foo bar" + "\n\t") + v4))) + "")), F1: 255, F2: 128}
                                        v22 = (v22 >> ((genFn0(-1)) & 63))
                                        switch (fi44 >> ((734) & 63)) {
                                        case 32767:
                                            v51 := GenS3{F0: 0, F1: ((v4 + "hello") + "foo bar")}
                                            _ = v51
                                        default:
                                            v52 := v5[(v4 + v4)]
                                            _ = v52
                                        }
                                        _ = v50
                                    } else {
                                        v53 := v5[v4]
                                        _ = v53
                                    }
                                }
                                v1 = v1
                                fmt.Println(v9, v47, fi44)
                            }
                            if len(v3) > 0 {
                                v3[0] = (v8 >> ((506) & 63))
                            }
                            if (v22 >= v22) {
                                if v1 {
                                    v22 = v22
                                    v54 := GenS8{F0: 1, F1: 42}
                                    _ = v54
                                }
                            } else {
                                var v55 float64 = ((((((0.1 * -2.25) - 0.0) + -2.25) + 0.1) * 1.5) / 1e10)
                                v1 = ((v22 | 1000) == v22)
                                _ = v55
                            }
                        }
                        _ = v48
                    }
                    _ = v46
                    _ = v47
                }
                if true {
                    var v56 int32 = -570
                    var v57 bool = v1
                    v57 = true
                    _ = v56
                    _ = v57
                }
                v1 = false
            }
            _ = fi45
        }
        for fi58 := 0; fi58 < 16; fi58++ {
            fmt.Println(v1)
            var v59 int64 = 493
            _ = fi58
            _ = v59
        }
        _ = fi44
    }
    var v60 int16 = 42
    v61 := GenS5{F0: ((!v1) && (!(v1 || (v8 > 2))))}
    if v1 {
        var v62 bool = ((!(-489 > 194)) && v1)
        var v63 int = 1
        _ = v62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v60:", v60)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
    _ = v8
    _ = v9
}
