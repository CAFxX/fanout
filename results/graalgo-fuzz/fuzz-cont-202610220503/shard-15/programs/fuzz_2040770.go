package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 string; F1 uint64; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 int; F1 int8 }
func genFn0(p0 int64) int {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: -382, F1: 292}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: 65535}
    v10 := []int{803, -809}
    _ = p0
    _ = v10
    _ = v9
    return 2147483647
}
func genFn1(p0 string, p1 int, p2 int64) int16 {
    v11 := GenS2{F0: 127, F1: 42}
    v12 := []int{p1}
    if len(v12) > 0 {
        v12[0] = (-(p1))
    }
    if true {
        p1 -= p1
        for fi13 := 0; fi13 < 8; fi13++ {
            if len(v12) > 0 {
                v12[0] = (fi13 - (fi13 % ((p1) | 1)))
            }
            _ = fi13
        }
        p1 = p1
    }
    if false {
        p2 += -200
        for fi14 := 0; fi14 < 8; fi14++ {
            if (p2 <= p2) {
                p0 = ((p0 + p0) + (("" + "foo bar") + p0))
            }
            var v15 int32 = 42
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(p0, p2)
                v15 -= (v15 * 0)
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 uint8 = 214
            _ = fi17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]string)
    var v3 int64 = 1
    var v4 uint64 = 0
    var v5 int16 = 42
    v6 := make(map[string]bool)
    if true {
        var v19 uint8 = (v1 + (^(v1)))
        _ = v19
    } else {
        v1 += v1
        if (((v4 >= (v4 ^ 2)) || (2 <= 1)) && true) {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 string = "foo bar"
                _ = fi20
                _ = v21
            }
        } else {
            delete(v6, ("foo bar" + "foo bar"))
            if (v4 == v4) {
                var v22 float64 = 1e10
                _ = v22
            }
        }
        fmt.Println(v3, v1, v3)
    }
    v1 -= 2
    v5 = -1
    v23 := GenS3{F0: "", F1: v4, F2: genFn0(9223372036854775807)}
    var v24 float64 = 3.14159
    if false {
        if (0.0 >= ((((1e10 / 0.1) + 3.14159) + ((-2.25 / float32(v24)) * (-2.25 * (1.5 + 0.0)))) * float32(((v24 + v24) + float64(-2.25))))) {
            v24 = 0.1
        }
    } else {
        v25 := GenS2{F0: 100, F1: 2}
        switch genFn0(v3) {
        case -2147483648:
            var v26 uint64 = 4294967295
            _ = v26
        case 46:
            var v27 uint64 = 2
            _ = v27
        case 128:
            v28 := GenS4{F0: v1}
            _ = v28
        }
        for fi29 := 0; fi29 < 6; fi29++ {
            _ = len(v6)
            for fi30 := 0; fi30 < 18; fi30++ {
                var v31 int64 = v3
                var v32 int8 = 1
                v3 = v31
                _ = fi30
                _ = v31
                _ = v32
            }
            switch 32768 {
            case 0:
                delete(v2, 4294967295)
            case 423:
                var v33 int64 = (v3 - (v3 + -1))
                _ = v33
            }
            _ = fi29
        }
        _ = v25
    }
    var v34 uint32 = 42
    for fi35 := 0; fi35 < 6; fi35++ {
        if true {
            v36 := GenS3{F0: "", F1: v4, F2: -9223372036854775808}
            delete(v6, "")
            _ = v36
        } else {
            switch "" {
            case "\n\t":
                for fi37 := 0; fi37 < 4; fi37++ {
                    v38 := GenS5{F0: 255, F1: v3}
                    fmt.Println(v1)
                    var v39 uint64 = (v4 * 42)
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            case "a":
                if (false || false) {
                    fmt.Println(v3)
                    if false {
                        if (true && (v5 != (v5 ^ (v5 % ((genFn1("hello", fi35, v3)) | 1))))) {
                            fmt.Println(v24, v3)
                        } else {
                            v40 := v2[v34]
                            _ = v40
                        }
                    } else {
                        v41 := []int{fi35}
                        if ((^(v5)) > v5) {
                            var v42 int64 = v3
                            v43 := GenS6{F0: v4, F1: 42, F2: 1}
                            _ = v42
                            _ = v43
                        } else {
                            for fi44 := 0; fi44 < 2; fi44++ {
                                v24 = v24
                                _ = fi44
                            }
                            v45 := GenS6{F0: 401, F1: 255, F2: 1}
                            _ = v45
                        }
                        _ = v41
                    }
                }
            }
            v24 -= v24
            var v46 float64 = v24
            _ = v46
        }
        if false {
            if false {
                v4 = v4
            }
            v24 = (0.0 * v24)
            for fi47 := 0; fi47 < 16; fi47++ {
                switch (fi47 + (fi35 % ((fi35) | 1))) {
                case 2147483647:
                    if (23 < 42) {
                        v34 += (v34 | v34)
                        var v48 int = fi47
                        if (!("hello" < "\n\t")) {
                            fmt.Println(v1)
                            v49 := GenS7{F0: (v48 + (v48 - (v48 + genFn0(v3)))), F1: (^(v3))}
                            _ = v49
                        } else {
                            delete(v2, v34)
                            v34 -= 2147483647
                            v50 := GenS4{F0: (v1 ^ 49)}
                            _ = v50
                        }
                        _ = v48
                    } else {
                        var v51 uint32 = 177
                        v34 -= 599
                        _ = v51
                    }
                default:
                    if true {
                        var v52 bool = true
                        if true {
                            v53 := v2[v34]
                            _ = v53
                        } else {
                            var v54 uint = 18446744073709551615
                            v55 := GenS8{F0: (fi35 << ((v54) & 63)), F1: int8(v5)}
                            var v56 uint = 18446744073709551615
                            v52 = v52
                            v57 := GenS4{F0: uint8(v4)}
                            _ = v54
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                        _ = v52
                    }
                }
                fmt.Println(v34)
                _ = fi47
            }
        } else {
            v58 := GenS1{F0: 2147483647, F1: 42}
            var v59 int = fi35
            _ = v58
            _ = v59
        }
        v6[(("hello" + "foo bar") + "\n\t")] = false
        _ = fi35
    }
    fmt.Println(v5, v24, v34)
    var v60 int16 = (v5 * (v5 | v5))
    v61 := []int{-2147483648, 9223372036854775807}
    for fi62 := 0; fi62 < 15; fi62++ {
        if false {
            var v63 int8 = -128
            v64 := GenS8{F0: (-(fi62)), F1: v63}
            _ = v63
            _ = v64
        } else {
            v65 := GenS7{F0: (^(fi62)), F1: -445}
            _ = v65
        }
        _ = fi62
    }
    fmt.Println(v24, v4, v34)
    var v66 string = ("a" + "a")
    v67 := GenS6{F0: 0, F1: 255, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v66
    _ = v67
}
