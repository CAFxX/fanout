package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 bool, p1 string, p2 int64) uint {
    var v7 uint64 = 1
    var v8 bool = true
    var v9 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1() bool {
    var v10 int8 = 127
    switch 42 {
    case 32767:
        v11 := []int{128, 0}
        _ = v11
    case 128:
        v12 := GenS1{F0: 42}
        _ = v12
    case -1:
        var v13 float64 = float64(1e10)
        _ = v13
    default:
        var v14 int = 9223372036854775807
        _ = v14
    }
    switch "a" {
    case "hello":
        var v15 uint = 320
        _ = v15
    case "":
        v16 := GenS1{F0: 416}
        _ = v16
    }
    _ = v10
    return (!(2147483647 >= 0))
}

func main() {
    var v1 int = 127
    var v2 uint16 = 255
    v3 := make(map[int]uint32)
    var v4 uint8 = 255
    v5 := make(map[int]uint32)
    var v6 float64 = 0.0
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := []int{(-(v1)), (fi17 >> ((1) & 63)), 1, (v1 >> ((16) & 63))}
        _ = fi17
        _ = v18
    }
    if (!true) {
        if true {
            v6 -= v6
            v2 = 65535
        } else {
            for fi19 := 0; fi19 < 19; fi19++ {
                fmt.Println(fi19, v2)
                v2 += 995
                _ = fi19
            }
            v2 = (v2 ^ (v2 / (((v2 >> ((uint(v2)) & 63))) | 1)))
            var v20 uint8 = 235
            _ = v20
        }
    } else {
        switch (("a" + "a") + ((("" + "") + "foo bar") + "a")) {
        case "\n\t":
            var v21 int32 = -1
            _ = v21
        case "":
            v4 = 255
        default:
            if true {
                var v22 float32 = 1.5
                _ = v22
            } else {
                for fi23 := 0; fi23 < 14; fi23++ {
                    switch "a" {
                    case "\n\t":
                        if (false || (float32((0.1 - (v6 - (v6 - float64(-2.25))))) > float32(float64(0.0)))) {
                            for fi24 := 0; fi24 < 5; fi24++ {
                                if (!(true || false)) {
                                    v6 = (-2.25 / float64(-2.25))
                                    var v25 int32 = 0
                                    _ = v25
                                }
                                _ = len(v3)
                                _ = fi24
                            }
                            var v26 bool = true
                            _ = v26
                        } else {
                            v1 -= fi23
                            if true {
                                switch "a" {
                                case "\n\t":
                                    v1 -= 818
                                case "a":
                                    var v27 uint64 = 0
                                    _ = v27
                                }
                            } else {
                                for fi28 := 0; fi28 < 14; fi28++ {
                                    v5[(v1 << ((18446744073709551615) & 63))] = 4294967295
                                    _ = fi28
                                }
                                if true {
                                    for fi29 := 0; fi29 < 5; fi29++ {
                                        v30 := GenS2{F0: 0, F1: ("" + "\n\t")}
                                        v31 := GenS1{F0: 773}
                                        fmt.Println(v4)
                                        _ = fi29
                                        _ = v30
                                        _ = v31
                                    }
                                }
                                delete(v5, (v1 - -9223372036854775808))
                            }
                            var v32 int = (v1 % (((fi23 / (((^(fi23))) | 1))) | 1))
                            _ = v32
                        }
                    case "foo bar":
                        v33 := []int{}
                        _ = v33
                    case "hello":
                        if true {
                            if ((2147483647 >= -1) || true) {
                                for fi34 := 0; fi34 < 8; fi34++ {
                                    v35 := []int{32767, 32767}
                                    v6 = v6
                                    _ = fi34
                                    _ = v35
                                }
                            } else {
                                if (float64((0.1 - 1.5)) < (v6 - (v6 * v6))) {
                                    v1 = -9223372036854775808
                                    v2 = 1
                                    v36 := GenS1{F0: -1}
                                    _ = v36
                                } else {
                                    v4 -= v4
                                    for fi37 := 0; fi37 < 10; fi37++ {
                                        var v38 uint16 = v2
                                        _ = fi37
                                        _ = v38
                                    }
                                }
                                if (-1 < 1) {
                                    var v39 int16 = 1
                                    _ = v39
                                }
                            }
                        } else {
                            v40 := []int{}
                            _ = v40
                        }
                    default:
                        for fi41 := 0; fi41 < 19; fi41++ {
                            v42 := []int{0, 0, (fi23 >> ((190) & 63))}
                            v43 := GenS3{F0: 153}
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    }
                    var v44 bool = false
                    v45 := GenS3{F0: 65535}
                    _ = fi23
                    _ = v44
                    _ = v45
                }
                v2 = 42
            }
        }
        delete(v3, v1)
        for fi46 := 0; fi46 < 3; fi46++ {
            var v47 float64 = 0.1
            _ = fi46
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 12; fi48++ {
        var v49 uint32 = 428
        var v50 uint8 = uint8(v49)
        if true {
            v51 := GenS1{F0: -31}
            _ = v51
        }
        _ = fi48
        _ = v49
        _ = v50
    }
    v4 = (v4 | v4)
    switch "foo bar" {
    case "hello":
        fmt.Println(v4, v2, v4)
    case "":
        var v52 uint64 = 4294967295
        _ = v52
    }
    v2 -= (v2 - 255)
    for fi53 := 0; fi53 < 10; fi53++ {
        v54 := GenS3{F0: 292}
        _ = fi53
        _ = v54
    }
    v55 := GenS3{F0: v2}
    v6 = (float64(3.14159) - (v6 * 0.0))
    if true {
        if (-32768 > 32767) {
            for fi56 := 0; fi56 < 16; fi56++ {
                for fi57 := 0; fi57 < 18; fi57++ {
                    var v58 uint64 = 1
                    var v59 uint16 = (v2 % ((1) | 1))
                    v6 -= v6
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                var v60 uint = 1
                v61 := []int{(v1 | -65), (v1 - v1), (-(fi56)), (v1 & fi56)}
                _ = fi56
                _ = v60
                _ = v61
            }
        }
    } else {
        var v62 float32 = 0.1
        var v63 float64 = ((v6 + v6) - 1e10)
        v64 := GenS4{F0: v4}
        _ = v62
        _ = v63
        _ = v64
    }
    var v65 int16 = -403
    var v66 int8 = 46
    var v67 uint64 = 11
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v65
    _ = v66
    _ = v67
}
