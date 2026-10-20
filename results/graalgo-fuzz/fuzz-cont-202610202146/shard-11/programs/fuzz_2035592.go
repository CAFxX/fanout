package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint }
type GenS4 struct { F0 int; F1 int8; F2 int32 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 int; F2 int16 }
type GenS8 struct { F0 int8; F1 uint32; F2 int8 }
type GenS9 struct { F0 uint; F1 int16; F2 uint32 }
func genFn0() bool {
    switch 328 {
    case 127:
        if false {
            v7 := GenS1{F0: -2147483648, F1: "foo bar"}
            _ = v7
        } else {
            var v8 uint16 = 32767
            for fi9 := 0; fi9 < 3; fi9++ {
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := []int{2147483647, fi9, (fi9 % (((fi10 | (fi10 % ((42) | 1)))) | 1)), fi10}
                    var v12 int8 = 0
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 uint64 = uint64(v8)
                _ = fi9
                _ = v13
            }
            _ = v8
        }
    case 32768:
        var v14 int32 = -2147483648
        _ = v14
    case 2147483647:
        var v15 int8 = -1
        _ = v15
    }
    v16 := []int{32768}
    v17 := GenS2{F0: 42}
    _ = len(v16)
    var v18 float32 = float32((0.1 / (1e10 * (-2.25 / 0.1))))
    _ = v16
    _ = v17
    _ = v18
    return (727 > 0)
}
func genFn1(p0 int64, p1 uint) int64 {
    var v19 float64 = 1e10
    p0 += -9223372036854775808
    _ = p0
    _ = p1
    _ = v19
    return (p0 / ((p0) | 1))
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 642
    var v3 int = 985
    v4 := []string{("a" + "a"), (v1 + ("a" + "hello")), v1}
    v5 := make(map[int]bool)
    v6 := []string{v1, (v1 + ("" + "foo bar"))}
    var v20 int8 = 100
    if false {
        for fi21 := 0; fi21 < 15; fi21++ {
            delete(v5, (fi21 / (((^(fi21))) | 1)))
            if true {
                switch "a" {
                case "a":
                    if (42 <= 4294967295) {
                        fmt.Println(v20, v2, v1)
                    } else {
                        for fi22 := 0; fi22 < 12; fi22++ {
                            v3 = (fi21 % (((fi22 & fi21)) | 1))
                            delete(v5, (fi22 - fi22))
                            _ = fi22
                        }
                        if (0 >= v20) {
                            for fi23 := 0; fi23 < 1; fi23++ {
                                v24 := GenS3{F0: v20, F1: 967, F2: 720}
                                if true {
                                    fmt.Println(fi23, v1, fi23)
                                }
                                if true {
                                    v3 = (fi23 >> ((255) & 63))
                                    v25 := GenS4{F0: (v3 >> ((65535) & 63)), F1: v20, F2: -723}
                                    switch (("" + v1) + "") {
                                    case "\n\t":
                                        var v26 int16 = (v2 % ((v2) | 1))
                                        _ = v26
                                    case "hello":
                                        var v27 uint64 = 0
                                        _ = v27
                                    case "a":
                                        var v28 int32 = -409
                                        switch int(v20) {
                                        case 32767:
                                            var v29 int64 = -1
                                            _ = v29
                                        default:
                                            v28 = (v28 << ((2) & 63))
                                            v5[32767] = (v28 == -373)
                                            v20 = (^(v20))
                                        }
                                        _ = v28
                                    default:
                                        if ((!(v20 != v20)) || false) {
                                            if (649 < fi23) {
                                                delete(v5, fi21)
                                                var v30 bool = (-1 <= 1)
                                                _ = v30
                                            } else {
                                                v31 := GenS3{F0: v20, F1: 400, F2: 255}
                                                var v32 uint = 42
                                                v33 := GenS5{F0: 26, F1: (!(!(42 >= 2))), F2: 65535}
                                                if ((v32 | (^(v32))) == v32) {
                                                    delete(v5, -83)
                                                    if true {
                                                        if true {
                                                            var v34 int64 = -1
                                                            _ = v34
                                                        } else {
                                                            var v35 uint32 = 671
                                                            v36 := GenS6{F0: 100}
                                                            v1 = (v1 + ("hello" + "\n\t"))
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                        fmt.Println(v2)
                                                    }
                                                    var v37 int64 = 198
                                                    _ = v37
                                                } else {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    v3 = fi21
                                                }
                                                v38 := GenS7{F0: -1, F1: -754, F2: 368}
                                                v1 = v1
                                                switch 0 {
                                                case 42:
                                                    v3 -= v3
                                                default:
                                                    delete(v5, (fi21 / ((v3) | 1)))
                                                }
                                                _ = v31
                                                _ = v32
                                                _ = v33
                                                _ = v38
                                            }
                                            switch (fi21 | (-(fi21))) {
                                            case -307:
                                                v3 += v3
                                            default:
                                                v39 := GenS8{F0: (v20 % ((30) | 1)), F1: 0, F2: -69}
                                                fmt.Println(fi21, v20, fi21)
                                                v2 -= v2
                                                _ = v39
                                            }
                                        } else {
                                            var v40 float32 = 0.1
                                            _ = v40
                                        }
                                    }
                                    _ = v25
                                }
                                _ = fi23
                                _ = v24
                            }
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            fmt.Println(v2)
                        }
                        var v41 int16 = v2
                        _ = v41
                    }
                case "\n\t":
                    v20 -= (-(v20))
                case "hello":
                    var v42 float32 = 0.0
                    _ = v42
                default:
                    for fi43 := 0; fi43 < 13; fi43++ {
                        v3 -= (^(fi21))
                        _ = fi43
                    }
                }
                v3 = (fi21 - (v3 * fi21))
            } else {
                var v44 bool = (!(!(42 < 0)))
                if (1 <= 42) {
                    var v45 bool = (!(1e10 == (-2.25 / 1.0)))
                    _ = v45
                }
                _ = v44
            }
            v3 -= v3
            _ = fi21
        }
        var v46 int64 = 0
        _ = v46
    } else {
        v47 := GenS3{F0: v20, F1: 2147483647, F2: 2}
        for fi48 := 0; fi48 < 19; fi48++ {
            if (2 < 4294967295) {
                if (!(69 > 128)) {
                    if true {
                        if ((float64((0.0 * float32((0.0 * 3.14159)))) >= 1.5) && false) {
                            delete(v5, fi48)
                        }
                        v49 := GenS9{F0: 424, F1: (v2 & v2), F2: 2}
                        fmt.Println(v3, v2)
                        _ = v49
                    }
                    var v50 uint = 2
                    for fi51 := 0; fi51 < 3; fi51++ {
                        var v52 uint16 = 0
                        _ = fi51
                        _ = v52
                    }
                    _ = v50
                } else {
                    v53 := GenS5{F0: 18446744073709551615, F1: false, F2: 65535}
                    _ = v53
                }
            } else {
                for fi54 := 0; fi54 < 20; fi54++ {
                    v1 = ("hello" + (v1 + v1))
                    _ = fi54
                }
                v20 = (v20 ^ v20)
                v1 = "hello"
            }
            _ = fi48
        }
        v2 += (v2 % (((v2 & (v2 / ((-1) | 1)))) | 1))
        _ = v47
    }
    var v55 float32 = (3.14159 + 0.1)
    fmt.Println(v55)
    v2 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v56 uint = 4294967295
    var v57 bool = false
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v58 bool = v57
    v2 -= (v2 % ((v2) | 1))
    v59 := v5[v3]
    if (78 > 42) {
        v1 = (v1 + v1)
    } else {
        _ = len(v5)
        var v60 float64 = (0.0 - 1e10)
        var v61 string = v1
        _ = v60
        _ = v61
    }
    fmt.Println(v58, v3)
    v20 -= v20
    if (!v59) {
        for fi62 := 0; fi62 < 8; fi62++ {
            v59 = true
            if v59 {
                v1 = "\n\t"
                for fi63 := 0; fi63 < 14; fi63++ {
                    var v64 uint = (^(v56))
                    v58 = v58
                    v65 := GenS3{F0: 116, F1: 42, F2: v56}
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                v5[v3] = false
            }
            _ = fi62
        }
        var v66 string = v1
        _ = v66
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
}
