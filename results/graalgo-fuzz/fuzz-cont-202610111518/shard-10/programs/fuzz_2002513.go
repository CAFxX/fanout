package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 string; F2 int8 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int32, p1 int64) uint16 {
    switch -1 {
    case -1:
        if (!true) {
            var v6 int = -648
            for fi7 := 0; fi7 < 2; fi7++ {
                var v8 uint = 106
                var v9 int16 = -114
                _ = fi7
                _ = v8
                _ = v9
            }
            v10 := []int{(v6 % (((v6 / (((v6 | 4294967295)) | 1))) | 1)), (v6 ^ (v6 & (v6 + (v6 - int(p0))))), v6}
            _ = v10
            _ = v6
        } else {
            v11 := GenS1{F0: 42}
            _ = v11
        }
    case 32767:
        fmt.Println(p0, p0)
    }
    v12 := GenS1{F0: 1}
    if false {
        p1 = (p1 * (p1 & (-(p1))))
        var v13 float64 = 3.14159
        var v14 float64 = 1.5
        _ = v13
        _ = v14
    } else {
        v15 := GenS2{F0: 1, F1: "foo bar", F2: -1}
        v16 := []int{-457, -713, 42, 2147483647}
        for fi17 := 0; fi17 < 19; fi17++ {
            if (1.5 > 1e10) {
                var v18 int8 = 42
                for fi19 := 0; fi19 < 19; fi19++ {
                    p0 = (p0 | (p0 ^ (p0 * 1)))
                    if false {
                        var v20 uint8 = 128
                        if len(v16) > 0 {
                            v16[0] = 9223372036854775807
                        }
                        p0 = p0
                        _ = v20
                    } else {
                        var v21 string = ""
                        _ = v21
                    }
                    fmt.Println(fi19)
                    p0 -= p0
                    _ = fi19
                }
                _ = v18
            } else {
                if (!((false && (-101 >= 0)) && (p0 == (p0 >> ((47) & 63))))) {
                    p0 -= (p0 * (p0 & -1))
                    v16 = append(v16, fi17)
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                }
                _ = len(v16)
                var v22 uint64 = 4294967295
                _ = v22
            }
            p1 = p1
            _ = fi17
        }
        _ = v15
        _ = v16
    }
    var v23 float64 = 0.0
    var v24 uint8 = 23
    _ = p0
    _ = p1
    _ = v12
    _ = v23
    _ = v24
    return 65535
}
func genFn1() uint16 {
    var v25 int8 = 127
    var v26 bool = ((492 <= 1) && (844 <= 4294967295))
    if ((!((("hello" + ("a" + "foo bar")) + "\n\t") > "a")) && false) {
        if (0 < 0) {
            switch "foo bar" {
            case "a":
                if (244 > 65535) {
                    v27 := GenS3{F0: 65535, F1: true}
                    _ = v27
                }
            case "":
                v28 := GenS2{F0: 2, F1: "a", F2: (v25 >> ((42) & 63))}
                _ = v28
            case "hello":
                fmt.Println(v26, v25)
            }
            var v29 uint8 = 0
            for fi30 := 0; fi30 < 20; fi30++ {
                v26 = ((v25 <= v25) || false)
                for fi31 := 0; fi31 < 15; fi31++ {
                    var v32 int8 = (v25 | v25)
                    var v33 string = ("" + "")
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = fi30
            }
            _ = v29
        }
    }
    v34 := []int{-812, 127, 32767}
    _ = v25
    _ = v26
    _ = v34
    return 929
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    var v5 uint32 = 2
    var v35 string = "foo bar"
    v36 := []int{32768, -680}
    if true {
        if (("hello" + ((v35 + "a") + ("" + v35))) <= v35) {
            var v37 int16 = -606
            _ = v37
        }
    }
    var v38 uint16 = 65535
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 bool = true
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        var v42 int16 = -1
        if true {
            var v43 int8 = 100
            var v44 float32 = 1.5
            if ((v2 / ((v2) | 1)) > v2) {
                var v45 int64 = 202
                var v46 int64 = -9223372036854775808
                v47 := GenS2{F0: uint64(v5), F1: "\n\t", F2: v43}
                _ = v45
                _ = v46
                _ = v47
            } else {
                if (0 <= -2147483648) {
                    _ = len(v36)
                    switch v35 {
                    case "hello":
                        var v48 uint16 = (v38 % ((v38) | 1))
                        _ = v48
                    case "foo bar":
                        v49 := GenS2{F0: v2, F1: v35, F2: (v43 - (v43 * (v43 % ((v43) | 1))))}
                        _ = v49
                    default:
                        v4[v35] = (150 < 2147483647)
                    }
                    var v50 string = "\n\t"
                    _ = v50
                }
            }
            _ = v43
            _ = v44
        }
        _ = fi41
        _ = v42
    }
    v35 = (((("\n\t" + v35) + v35) + "") + v35)
    fmt.Println(v35, v35)
    v5 = (v5 - 333)
    for fi51 := 0; fi51 < 20; fi51++ {
        if (!((^(v5)) <= (v5 | 865))) {
            if (fi51 != (fi51 - (-(fi51)))) {
                if len(v36) > 0 {
                    v36[0] = (fi51 / (((fi51 / ((fi51) | 1))) | 1))
                }
            }
            v52 := GenS4{F0: 233, F1: -633}
            v35 = (((v35 + v35) + ("a" + (v35 + v35))) + (v35 + v35))
            _ = v52
        } else {
            var v53 int32 = -789
            _ = v53
        }
        if (false && true) {
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
            if (442 <= 2) {
                if false {
                    var v54 float64 = 3.14159
                    switch fi51 {
                    case -406:
                        v55 := GenS5{F0: -1}
                        _ = v55
                    case 1:
                        v35 = v35
                    case 805:
                        var v56 float64 = v54
                        _ = v56
                    default:
                        var v57 float32 = 1.5
                        _ = v57
                    }
                    fmt.Println(fi51)
                    _ = v54
                }
                fmt.Println(v2, v35, v2)
                v35 = ""
            } else {
                var v58 uint = 1
                _ = v58
            }
            v2 = v2
        } else {
            var v59 uint64 = 4294967295
            for fi60 := 0; fi60 < 18; fi60++ {
                fmt.Println(v5)
                _ = fi60
            }
            _ = v59
        }
        _ = fi51
    }
    var v61 int32 = 158
    var v62 float32 = 3.14159
    var v63 float32 = (v62 * v62)
    v64 := GenS1{F0: 42}
    fmt.Println(v63, v38)
    v63 = (v63 - v62)
    var v65 int16 = -1
    if ((!((v38 + v38) >= genFn0(v61, int64(v65)))) && true) {
        fmt.Println(v63, v61, v65)
        v66 := GenS1{F0: 340}
        _ = v66
    } else {
        var v67 uint8 = 191
        delete(v1, v35)
        _ = v67
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
