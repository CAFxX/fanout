package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint; F1 int; F2 int }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int16; F1 uint64; F2 int }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 int16; F1 int }
func genFn0(p0 uint64) uint16 {
    v6 := []int{-1, 128, 595}
    var v7 int = -346
    var v8 uint = 2
    for fi9 := 0; fi9 < 8; fi9++ {
        if false {
            switch "a" {
            case "":
                if ("foo bar" <= "a") {
                    if len(v6) > 0 {
                        v6[0] = (fi9 % ((-9223372036854775808) | 1))
                    }
                    var v10 int = v7
                    for fi11 := 0; fi11 < 6; fi11++ {
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            fmt.Println(v10)
                        default:
                            p0 = (p0 ^ p0)
                        }
                        var v12 int64 = 42
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                } else {
                    var v13 bool = (2 != 18446744073709551615)
                    var v14 uint16 = 0
                    _ = v13
                    _ = v14
                }
            default:
                if true {
                    var v15 bool = (1e10 != (1e10 - 0.1))
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v6 = append(v6, fi16)
                        if (v15 && (fi16 > (v7 | (v7 + 32)))) {
                            fmt.Println(v8, fi16)
                            if (0 <= 475) {
                                v8 -= (^(v8))
                            } else {
                                v8 += 2
                                if ("a" >= "a") {
                                    v7 = fi9
                                    if v15 {
                                        var v17 int = (fi9 >> ((v8) & 63))
                                        if len(v6) > 0 {
                                            v6[0] = v7
                                        }
                                        _ = v17
                                    }
                                } else {
                                    if true {
                                        var v18 int = v7
                                        v15 = ((p0 / ((p0) | 1)) >= (p0 % (((p0 - p0)) | 1)))
                                        if false {
                                            p0 += 235
                                        }
                                        _ = v18
                                    }
                                    var v19 int = -2147483648
                                    var v20 string = "\n\t"
                                    var v21 int16 = 1000
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                            }
                            v22 := GenS1{F0: 0, F1: 127}
                            _ = v22
                        } else {
                            if v15 {
                                var v23 uint8 = 255
                                var v24 uint32 = 511
                                v25 := GenS2{F0: p0}
                                _ = v23
                                _ = v24
                                _ = v25
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = (fi16 * (fi16 + (^(fi9))))
                                }
                                v26 := GenS3{F0: 127, F1: -32768}
                                var v27 bool = (true && false)
                                _ = v26
                                _ = v27
                            }
                            v15 = (1 != 9223372036854775807)
                            var v28 uint = (v8 & v8)
                            _ = v28
                        }
                        if false {
                            if v15 {
                                var v29 uint32 = 42
                                fmt.Println(v8, v29)
                                _ = v29
                            }
                            v30 := GenS1{F0: -1, F1: 100}
                            _ = v30
                        }
                        _ = fi16
                    }
                    v31 := GenS4{F0: (v8 * v8), F1: -918, F2: (^(fi9))}
                    _ = v15
                    _ = v31
                }
            }
            var v32 int = v7
            var v33 float64 = float64((1e10 * -2.25))
            _ = v32
            _ = v33
        }
        _ = fi9
    }
    var v34 bool = true
    _ = p0
    _ = v34
    _ = v6
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 int = 128
    v2 := []string{}
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    var v5 uint64 = 42
    delete(v3, ("hello" + (((("foo bar" + (((("foo bar" + "a") + ("a" + ("" + ""))) + ("hello" + "hello")) + ("" + (("foo bar" + ("\n\t" + "foo bar")) + "")))) + ("a" + "a")) + "foo bar") + "\n\t")))
    v5 -= v5
    if (true && (true && ((v5 * (v5 / ((v5) | 1))) > (v5 - (^(v5)))))) {
        for fi35 := 0; fi35 < 11; fi35++ {
            switch (v1 << ((18446744073709551615) & 63)) {
            case -2147483648:
                fmt.Println(v1, v1)
            case 42:
                var v36 uint64 = v5
                _ = v36
            }
            var v37 uint64 = v5
            v5 = 2
            _ = fi35
            _ = v37
        }
    } else {
        if (-28 > 100) {
            if (65535 == 42) {
                for fi38 := 0; fi38 < 5; fi38++ {
                    for fi39 := 0; fi39 < 20; fi39++ {
                        switch (fi39 / ((fi38) | 1)) {
                        case 127:
                            v40 := GenS5{F0: 124, F1: (("" + "a") + "foo bar")}
                            v41 := GenS5{F0: 65535, F1: "a"}
                            fmt.Println(v5, v1, fi38)
                            _ = v40
                            _ = v41
                        case 1:
                            if (1 >= 583) {
                                if len(v2) > 0 {
                                    v2[0] = "\n\t"
                                }
                                if (false && false) {
                                    v1 = v1
                                }
                                var v42 int16 = 32767
                                var v43 uint8 = 102
                                _ = v42
                                _ = v43
                            }
                        }
                        v2 = append(v2, ("a" + "a"))
                        _ = fi39
                    }
                    v44 := GenS5{F0: 42, F1: ("\n\t" + ("a" + "hello"))}
                    _ = fi38
                    _ = v44
                }
                fmt.Println(v1)
                if ((!((v5 & v5) == 2)) && (true && (!((v1 == (v1 * v1)) || (1 < 42))))) {
                    if true {
                        var v45 int32 = 578
                        var v46 uint = 18446744073709551615
                        _ = v45
                        _ = v46
                    }
                }
            } else {
                v5 = 1
            }
            v3["a"] = v1
            v5 = 42
        }
        if (174 < 2) {
            v47 := GenS2{F0: (v5 ^ 42)}
            var v48 uint8 = 255
            var v49 int64 = 312
            _ = v47
            _ = v48
            _ = v49
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v50 uint64 = v5
            _ = v50
        }
        var v51 uint8 = 231
        _ = v51
    }
    v52 := GenS6{F0: -242, F1: 1, F2: v1}
    if false {
        for fi53 := 0; fi53 < 16; fi53++ {
            var v54 float32 = float32(float64(0.0))
            for fi55 := 0; fi55 < 18; fi55++ {
                if false {
                    var v56 int = 42
                    v57 := GenS7{F0: 69, F1: 281}
                    var v58 int32 = 2147483647
                    v54 += 1.5
                    if (false || true) {
                        v5 = (v5 >> ((1) & 63))
                    }
                    fmt.Println(v54, v56)
                    _ = v56
                    _ = v57
                    _ = v58
                } else {
                    v59 := GenS6{F0: 1, F1: (v5 ^ v5), F2: -641}
                    if false {
                        fmt.Println(v54, fi53, fi53)
                        var v60 uint64 = 42
                        _ = v60
                    } else {
                        if ((130 <= 1000) && (!false)) {
                            v61 := GenS5{F0: 65535, F1: ("foo bar" + "")}
                            v4[534] = 0
                            v5 -= (v5 & v5)
                            fmt.Println(fi53)
                            _ = v61
                        }
                        var v62 uint8 = 42
                        v1 = (v1 ^ fi55)
                        _ = v62
                    }
                    _ = v59
                }
                if false {
                    var v63 int16 = -32768
                    _ = v63
                }
                var v64 uint = 826
                _ = fi55
                _ = v64
            }
            if ((2 != 42) && (!false)) {
                var v65 uint16 = 1
                var v66 float32 = ((1.5 - v54) - v54)
                _ = v65
                _ = v66
            }
            _ = fi53
            _ = v54
        }
        v5 -= v5
        fmt.Println(v1, v5)
    }
    v1 -= v1
    v1 = (v1 - v1)
    var v67 int64 = 0
    for fi68 := 0; fi68 < 15; fi68++ {
        v69 := GenS8{F0: 687, F1: -589}
        delete(v3, "a")
        _ = fi68
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v67
}
