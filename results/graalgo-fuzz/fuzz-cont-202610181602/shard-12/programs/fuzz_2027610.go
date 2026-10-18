package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 string; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint8 }
func genFn0() int64 {
    var v7 uint = 65535
    v8 := GenS1{F0: -961, F1: 42, F2: 32767}
    var v9 float32 = -2.25
    switch -114 {
    case 1:
        fmt.Println(v7)
    case -1:
        fmt.Println(v7)
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint16 = 42
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() int {
    v12 := []int{-9223372036854775808, 2147483647, 32767}
    var v13 int32 = -1
    var v14 int = -9223372036854775808
    _ = v12
    _ = v13
    _ = v14
    return 290
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -1
    var v3 uint = 255
    var v4 float64 = (3.14159 * 0.1)
    v5 := make(map[int]bool)
    var v6 uint16 = 1
    var v15 int64 = genFn0()
    v16 := []int{genFn1(), 42, genFn1(), 32767}
    var v17 float32 = 0.1
    v15 += (^(v15))
    _ = len(v1)
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS1{F0: 311, F1: 65535, F2: v6}
        v20 := v5[4294967295]
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int = int(v6)
        v17 = v17
        _ = fi21
        _ = v22
    }
    v15 = 42
    if true {
        if (2 >= 138) {
            var v23 bool = ((v2 & v2) == v2)
            v17 = v17
            for fi24 := 0; fi24 < 15; fi24++ {
                v17 = 1.5
                for fi25 := 0; fi25 < 7; fi25++ {
                    var v26 string = ("a" + ("\n\t" + ("" + ("foo bar" + "\n\t"))))
                    _ = len(v5)
                    var v27 float32 = (((v17 - v17) - v17) / (v17 / float32((((0.1 + ((((v4 / 1.0) / v4) * v4) - v4)) + (((0.1 * v4) + ((v4 / (0.1 * ((v4 + v4) / v4))) + v4)) - (float64((v17 * (3.14159 / float32(1.5)))) * v4))) * (1e10 + v4)))))
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = fi24
            }
            _ = v23
        }
        var v28 bool = (!false)
        if (!v28) {
            switch 42 {
            case -2147483648:
                v6 = 65535
            }
        } else {
            for fi29 := 0; fi29 < 7; fi29++ {
                var v30 bool = (v28 && v28)
                _ = fi29
                _ = v30
            }
            if (v15 > 385) {
                v6 = 542
                v31 := v1[901]
                switch 683 {
                case 42:
                    for fi32 := 0; fi32 < 6; fi32++ {
                        v33 := GenS1{F0: (v2 - int32(v6)), F1: (v6 >> (((v3 - (v3 & v3))) & 63)), F2: v6}
                        if false {
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        } else {
                            for fi34 := 0; fi34 < 10; fi34++ {
                                v17 = v17
                                v1[2] = v31
                                v17 -= (v17 - ((3.14159 * v17) - (1e10 / float32(v4))))
                                _ = fi34
                            }
                            v35 := GenS2{F0: v28}
                            _ = v35
                        }
                        _ = fi32
                        _ = v33
                    }
                case -2147483648:
                    v15 = (-(v15))
                }
                _ = v31
            } else {
                if true {
                    for fi36 := 0; fi36 < 9; fi36++ {
                        v37 := GenS2{F0: (v28 || (!((-(v2)) == v2)))}
                        if true {
                            v17 -= v17
                        }
                        if (v28 && (("foo bar" + "\n\t") <= (("hello" + (("hello" + "a") + ("a" + ""))) + "\n\t"))) {
                            for fi38 := 0; fi38 < 12; fi38++ {
                                fmt.Println(v4, v17)
                                if ("\n\t" != "") {
                                    v6 = v6
                                    v15 = genFn0()
                                    if false {
                                        v6 = v6
                                        v28 = true
                                        var v39 int16 = -426
                                        switch genFn1() {
                                        case 323:
                                            if v28 {
                                                var v40 string = "a"
                                                var v41 int = (fi38 - (-(fi36)))
                                                _ = v40
                                                _ = v41
                                            }
                                        }
                                        var v42 uint64 = 4294967295
                                        if (((((v4 * v4) + v4) + v4) <= (1.5 * ((0.0 / 1.0) / v4))) && v28) {
                                            v28 = (v28 && v28)
                                        }
                                        if len(v16) > 0 {
                                            v16[0] = (fi36 - (^(fi38)))
                                        }
                                        _ = v39
                                        _ = v42
                                    } else {
                                        if ((v6 & 303) == v6) {
                                            v4 += 0.1
                                        }
                                    }
                                    switch (fi38 & (fi38 / (((fi36 / ((genFn1()) | 1))) | 1))) {
                                    case 1:
                                        if v28 {
                                            var v43 float64 = (float64((v17 / (0.1 / (((3.14159 - 0.1) + (1.5 + float32(float64((float32(0.0) * float32(1.5)))))) - 0.0)))) * (v4 - v4))
                                            v44 := GenS2{F0: v28}
                                            v28 = true
                                            _ = v43
                                            _ = v44
                                        } else {
                                            v2 += -1
                                            var v45 string = "hello"
                                            _ = v45
                                        }
                                    case 42:
                                        v6 += 32767
                                    case 600:
                                        _ = len(v16)
                                        v46 := GenS2{F0: (v4 == (v4 - v4))}
                                        if v28 {
                                            fmt.Println(v6)
                                            v4 -= v4
                                            var v47 uint = (v3 ^ v3)
                                            v48 := GenS2{F0: (!(!v28))}
                                            if (v15 != v15) {
                                                v28 = v28
                                                var v49 float32 = 1.5
                                                _ = v49
                                            } else {
                                                _ = len(v5)
                                                v50 := GenS2{F0: (-1 == (^(v2)))}
                                                _ = v50
                                            }
                                            _ = v47
                                            _ = v48
                                        }
                                        _ = v46
                                    }
                                }
                                if v28 {
                                    if v28 {
                                        v6 = 42
                                        v4 += float64(0.1)
                                        v51 := GenS1{F0: 42, F1: (v6 | v6), F2: v6}
                                        v15 = v15
                                        _ = v51
                                    } else {
                                        var v52 int64 = 0
                                        _ = v52
                                    }
                                } else {
                                    v4 = v4
                                    var v53 float64 = float64(v17)
                                    fmt.Println(fi36, v4)
                                    _ = v53
                                }
                                switch (fi38 >> ((v3) & 63)) {
                                case 2147483647:
                                    var v54 uint16 = 1
                                    _ = v54
                                case 1:
                                    delete(v5, genFn1())
                                case 9223372036854775807:
                                    switch fi36 {
                                    case 0:
                                        fmt.Println(v6)
                                    }
                                    var v55 uint16 = 15
                                    _ = v55
                                default:
                                    v2 -= v2
                                }
                                var v56 int = fi38
                                _ = fi38
                                _ = v56
                            }
                            if (v28 || v28) {
                                v6 = v6
                                switch genFn1() {
                                case -455:
                                    if (-539 != 996) {
                                        for fi57 := 0; fi57 < 9; fi57++ {
                                            var v58 float64 = v4
                                            v17 = (((0.0 - (1e10 - 3.14159)) + v17) / v17)
                                            var v59 int64 = v15
                                            var v60 uint = uint(v6)
                                            var v61 uint64 = 0
                                            _ = fi57
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                            _ = v61
                                        }
                                        if v28 {
                                            _ = len(v5)
                                            var v62 bool = false
                                            _ = v62
                                        } else {
                                            v63 := GenS3{F0: 1, F1: "", F2: 772}
                                            fmt.Println(v4, v17)
                                            var v64 uint = v3
                                            _ = v63
                                            _ = v64
                                        }
                                    }
                                case 32768:
                                    fmt.Println(v15)
                                case 972:
                                    switch fi36 {
                                    case 4294967295:
                                        switch fi36 {
                                        case 0:
                                            v1[65535] = ("hello" + ("\n\t" + ("" + ("a" + "\n\t"))))
                                        }
                                    case 42:
                                        if (2 == 325) {
                                            v17 = -2.25
                                            var v65 float32 = 1.5
                                            if (-9223372036854775808 > genFn0()) {
                                                v66 := GenS3{F0: v3, F1: "", F2: 1}
                                                v17 = 3.14159
                                                var v67 uint = v3
                                                _ = v66
                                                _ = v67
                                            } else {
                                                _ = len(v16)
                                                var v68 int16 = 32767
                                                v2 += (v2 << (((v3 ^ 255)) & 63))
                                                _ = v68
                                            }
                                            _ = v65
                                        } else {
                                            _ = len(v16)
                                            var v69 uint32 = 4294967295
                                            _ = v69
                                        }
                                    case -21:
                                        var v70 uint8 = 37
                                        _ = v70
                                    }
                                default:
                                    var v71 uint32 = uint32(v3)
                                    _ = v71
                                }
                            }
                        } else {
                            for fi72 := 0; fi72 < 9; fi72++ {
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                fmt.Println(v3)
                                _ = fi72
                            }
                            for fi73 := 0; fi73 < 15; fi73++ {
                                var v74 int = (fi36 + 32768)
                                _ = fi73
                                _ = v74
                            }
                            v1[2147483647] = "foo bar"
                        }
                        _ = fi36
                        _ = v37
                    }
                } else {
                    v75 := GenS4{F0: v3, F1: 124}
                    _ = v75
                }
            }
            v76 := v5[127]
            _ = v76
        }
        _ = v28
    }
    delete(v5, -1)
    var v77 uint16 = (^(v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
}
