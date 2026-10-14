package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int16; F2 bool }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint; F1 bool; F2 int }
type GenS5 struct { F0 uint64; F1 string; F2 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool }
func genFn0(p0 bool, p1 bool, p2 int16) uint {
    fmt.Println(p2)
    var v7 int16 = (p2 % ((-32768) | 1))
    var v8 int32 = -1
    v9 := []int{32767, -2147483648, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := []string{}
    var v2 string = (("foo bar" + (("\n\t" + "foo bar") + ("" + ("foo bar" + "foo bar")))) + "")
    var v3 bool = false
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    var v6 int64 = 648
    for fi10 := 0; fi10 < 4; fi10++ {
        _ = len(v1)
        if ((-2.25 == 3.14159) && (3.14159 <= (0.0 * float64(3.14159)))) {
            if (!(!v3)) {
                var v11 int8 = -20
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 int32 = 2147483647
                    _ = fi12
                    _ = v13
                }
                _ = v11
            } else {
                v6 += -582
                var v14 int32 = 53
                var v15 int32 = 676
                _ = v14
                _ = v15
            }
            switch (fi10 ^ fi10) {
            case -9223372036854775808:
                for fi16 := 0; fi16 < 17; fi16++ {
                    switch fi16 {
                    case 351:
                        var v17 uint32 = 2
                        _ = v17
                    case 42:
                        if (!v3) {
                            var v18 float64 = 1e10
                            v18 += ((v18 + v18) + (v18 / (((0.0 * ((v18 - -2.25) + (0.0 + v18))) / float64(-2.25)) / -2.25)))
                            _ = v18
                        } else {
                            fmt.Println(v3, v6)
                        }
                    default:
                        var v19 uint8 = 42
                        v19 = (^(v19))
                        _ = len(v4)
                        _ = v19
                    }
                    fmt.Println(fi10)
                    v20 := GenS1{F0: -62}
                    _ = fi16
                    _ = v20
                }
            case 42:
                v3 = v3
            default:
                var v21 int32 = -1
                _ = v21
            }
        } else {
            v2 = ((("foo bar" + v2) + ((v2 + (v2 + v2)) + "hello")) + ((((v2 + "hello") + v2) + (("a" + "a") + ((v2 + "\n\t") + v2))) + v2))
            var v22 uint16 = 0
            var v23 int = (fi10 | 2147483647)
            _ = v22
            _ = v23
        }
        _ = fi10
    }
    v4[("hello" + "")] = 436
    v2 = ((("\n\t" + "hello") + (v2 + (v2 + ("\n\t" + "hello")))) + v2)
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(fi24, v2)
        var v25 float64 = ((0.1 + -2.25) * 1e10)
        _ = fi24
        _ = v25
    }
    v3 = true
    if (v3 || false) {
        v26 := v4[(v2 + v2)]
        var v27 int32 = 42
        _ = v26
        _ = v27
    } else {
        delete(v5, ("hello" + v2))
        var v28 uint16 = 481
        v3 = true
        _ = v28
    }
    _ = len(v1)
    v3 = (42 != 0)
    var v29 bool = v3
    _ = len(v1)
    if true {
        var v30 int64 = (v6 / ((v6) | 1))
        var v31 string = "a"
        v32 := GenS2{F0: 32768, F1: 0, F2: v29}
        _ = v30
        _ = v31
        _ = v32
    } else {
        if ((2147483647 <= -303) || v3) {
            var v33 int32 = 2147483647
            var v34 float32 = -2.25
            for fi35 := 0; fi35 < 11; fi35++ {
                v36 := GenS3{F0: true, F1: 127}
                delete(v5, "")
                for fi37 := 0; fi37 < 9; fi37++ {
                    var v38 bool = v29
                    if v29 {
                        if (2 != 4294967295) {
                            fmt.Println(v38)
                            var v39 string = "a"
                            _ = v39
                        } else {
                            v40 := GenS4{F0: 344, F1: v3, F2: fi37}
                            var v41 uint8 = 128
                            fmt.Println(v2, v34, v29)
                            switch (fi35 ^ (fi35 % ((fi35) | 1))) {
                            case 32767:
                                if false {
                                    v42 := GenS5{F0: 2, F1: v2, F2: 0}
                                    _ = v42
                                }
                            case -627:
                                var v43 int32 = v33
                                _ = v43
                            case 0:
                                delete(v4, (v2 + "foo bar"))
                            }
                            _ = v40
                            _ = v41
                        }
                        var v44 uint16 = 1
                        _ = v44
                    }
                    var v45 uint64 = 785
                    _ = fi37
                    _ = v38
                    _ = v45
                }
                _ = fi35
                _ = v36
            }
            _ = v33
            _ = v34
        } else {
            if (false || v3) {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            } else {
                v46 := GenS6{F0: 0}
                v3 = false
                _ = v46
            }
        }
    }
    switch 4294967295 {
    case 9223372036854775807:
        v4[v2] = 65535
    case 128:
        if (9223372036854775807 < (v6 | (^(v6)))) {
            switch v2 {
            case "hello":
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
            case "":
                fmt.Println(v29, v3, v6)
            case "foo bar":
                var v47 int8 = 0
                _ = v47
            }
        }
    default:
        v2 = ((v2 + (v2 + v2)) + v2)
    }
    if ((-2.25 / -2.25) != 0.1) {
        var v48 int32 = 44
        _ = v48
    }
    for fi49 := 0; fi49 < 3; fi49++ {
        if true {
            if (v3 || v29) {
                v50 := GenS5{F0: 2, F1: v2, F2: 588}
                delete(v5, ((v2 + v2) + "foo bar"))
                _ = v50
            }
            if (42 <= 42) {
                v2 = "foo bar"
                v51 := GenS6{F0: 199}
                _ = v51
            }
            v52 := GenS6{F0: 1}
            _ = v52
        }
        if (!false) {
            v6 += v6
        } else {
            v2 = v2
            if (fi49 > (-(fi49))) {
                v29 = false
            } else {
                var v53 string = v2
                v54 := GenS3{F0: (0.1 > 1e10), F1: -26}
                _ = v53
                _ = v54
            }
            v55 := GenS1{F0: 88}
            _ = v55
        }
        var v56 bool = v3
        _ = fi49
        _ = v56
    }
    switch 910 {
    case -9223372036854775808:
        if v29 {
            var v57 uint = genFn0((v3 && (4294967295 > 59)), v29, 1)
            v2 = v2
            _ = len(v1)
            _ = v57
        } else {
            delete(v5, v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case -756:
        _ = len(v4)
    default:
        v2 = ("a" + "hello")
    }
    v3 = ((-488 > 1) && (4294967295 != 2147483648))
    v58 := GenS4{F0: 4294967295, F1: v29, F2: 2147483647}
    if v3 {
        fmt.Println(v6, v29)
        var v59 int8 = -128
        for fi60 := 0; fi60 < 16; fi60++ {
            switch (fi60 - (fi60 % ((495) | 1))) {
            case -2147483648:
                if (true && (((v59 <= -128) && v3) || v29)) {
                    var v61 int64 = v6
                    _ = v61
                }
            case 56:
                v62 := GenS2{F0: (fi60 % ((fi60) | 1)), F1: -293, F2: (((0.0 / (-2.25 / 1.0)) + 1.5) < -2.25)}
                _ = v62
            case -1:
                fmt.Println(v6, v6)
            default:
                var v63 uint = 18446744073709551615
                _ = v63
            }
            v64 := v4[((("a" + (v2 + (v2 + ((("foo bar" + (((v2 + v2) + "a") + (v2 + (((v2 + (("hello" + v2) + v2)) + "hello") + v2)))) + v2) + v2)))) + v2) + ("\n\t" + "\n\t"))]
            v65 := GenS5{F0: 42, F1: (("a" + v2) + ((v2 + v2) + ("hello" + v2))), F2: 319}
            _ = fi60
            _ = v64
            _ = v65
        }
        _ = v59
    } else {
        v29 = ((v29 || v29) || true)
        v66 := GenS7{F0: v29}
        _ = v66
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
