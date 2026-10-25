package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 uint32; F1 bool; F2 int }
type GenS5 struct { F0 uint64; F1 int64 }

func main() {
    v1 := []uint32{2147483647}
    v2 := []uint32{415, 1}
    var v3 int8 = 42
    var v4 bool = true
    var v5 uint = 2
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: 4294967295}
    var v8 bool = (!v4)
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v8)
        v8 = v4
        _ = fi9
    }
    v10 := GenS1{F0: 41}
    if ((!v4) && ((2147483647 == 2) || false)) {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 bool = (!(false && false))
            _ = len(v1)
            _ = fi11
            _ = v12
        }
        if ((42 >= 0) || (!v4)) {
            for fi13 := 0; fi13 < 7; fi13++ {
                if (30 != 1) {
                    v14 := GenS1{F0: 0}
                    v8 = false
                    var v15 float32 = 3.14159
                    _ = v14
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        v5 = v5
                        var v17 uint = v5
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 6; fi18++ {
                        var v19 uint16 = 1
                        _ = fi18
                        _ = v19
                    }
                    var v20 uint8 = 0
                    _ = v20
                }
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 float64 = (3.14159 + float64(((1.5 / 1.0) / 1e10)))
                    if ((fi21 & -1) < (fi13 % (((fi21 | fi21)) | 1))) {
                        v5 = v5
                    } else {
                        if (!true) {
                            fmt.Println(fi13, v8)
                        } else {
                            if v8 {
                                var v23 int64 = int64(v3)
                                v24 := GenS1{F0: 2147483647}
                                if v8 {
                                    switch fi21 {
                                    case 32767:
                                        var v25 string = "foo bar"
                                        _ = v25
                                    default:
                                        v3 = v3
                                    }
                                    v26 := GenS1{F0: 200}
                                    _ = v26
                                }
                                _ = v23
                                _ = v24
                            } else {
                                v22 = -2.25
                                v27 := GenS2{F0: v3}
                                if (390 >= 1) {
                                    v3 += (v3 | (v3 ^ v3))
                                }
                                _ = v27
                            }
                        }
                        var v28 int8 = v3
                        if (((("" + "\n\t") + "") + ("\n\t" + "\n\t")) >= "\n\t") {
                            var v29 string = (("" + ("a" + "foo bar")) + ("foo bar" + ""))
                            _ = v29
                        }
                        v30 := GenS2{F0: (v28 << (((v5 >> ((817) & 63))) & 63))}
                        _ = v28
                        _ = v30
                    }
                    _ = fi21
                    _ = v22
                }
                fmt.Println(fi13, v3)
                _ = fi13
            }
        } else {
            for fi31 := 0; fi31 < 19; fi31++ {
                for fi32 := 0; fi32 < 14; fi32++ {
                    fmt.Println(v5, v3, v3)
                    switch "foo bar" {
                    case "":
                        if (9223372036854775807 < -567) {
                            var v33 string = ((("" + ("foo bar" + "a")) + "\n\t") + "hello")
                            _ = v33
                        } else {
                            if (18446744073709551615 >= 1) {
                                v34 := GenS2{F0: v3}
                                v35 := GenS1{F0: 1}
                                _ = v34
                                _ = v35
                            } else {
                                v36 := GenS2{F0: v3}
                                v37 := GenS1{F0: 2147483647}
                                v4 = v4
                                fmt.Println(v8)
                                _ = v36
                                _ = v37
                            }
                            v5 = (v5 - 993)
                        }
                    default:
                        _ = len(v2)
                    }
                    if (v8 || true) {
                        v5 = (v5 ^ (v5 / ((v5) | 1)))
                        fmt.Println(v8, v5)
                        if ((("" + ("hello" + "hello")) + (("foo bar" + "foo bar") + "hello")) != (("\n\t" + "foo bar") + ("\n\t" + (("a" + "") + "foo bar")))) {
                            if (18446744073709551615 != 1) {
                                var v38 float32 = ((1e10 * 0.0) - 1e10)
                                var v39 int = fi32
                                if len(v2) > 0 {
                                    v2[0] = 1
                                }
                                _ = v38
                                _ = v39
                            } else {
                                fmt.Println(v8, v4)
                                var v40 uint16 = 2
                                _ = v40
                            }
                        } else {
                            v41 := GenS2{F0: 42}
                            _ = v41
                        }
                    }
                    _ = fi32
                }
                for fi42 := 0; fi42 < 5; fi42++ {
                    if ((fi31 & fi31) == fi31) {
                        v43 := GenS2{F0: v3}
                        v44 := GenS3{F0: false, F1: (v5 << ((v5) & 63))}
                        switch fi31 {
                        case -612:
                            if (!v8) {
                                v3 = v3
                            } else {
                                v8 = v4
                            }
                        default:
                            v8 = (((-2.25 / 1.0) - 0.1) != (1e10 * 0.1))
                        }
                        _ = v43
                        _ = v44
                    }
                    var v45 bool = v8
                    v6[65535] = ""
                    v3 -= 42
                    _ = fi42
                    _ = v45
                }
                _ = fi31
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 = append(v1, uint32(v5))
        }
        v46 := GenS2{F0: v3}
        _ = v46
    } else {
        fmt.Println(v4, v4, v8)
        _ = len(v6)
    }
    var v47 float32 = (1e10 * 0.0)
    v6[568] = "\n\t"
    v48 := GenS4{F0: 863, F1: (!(v4 || (v5 < 773))), F2: 360}
    v47 += v47
    var v49 bool = v8
    var v50 string = "\n\t"
    for fi51 := 0; fi51 < 10; fi51++ {
        switch (fi51 ^ -85) {
        case 9223372036854775807:
            var v52 uint8 = 42
            _ = v52
        case -882:
            v53 := GenS1{F0: 948}
            _ = v53
        case 622:
            v8 = (v3 != v3)
        }
        v6[0] = v50
        _ = fi51
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        if (1 >= -548) {
            v55 := GenS3{F0: v49, F1: (v5 / ((v5) | 1))}
            switch (v50 + v50) {
            case "foo bar":
                var v56 string = (v50 + v50)
                _ = v56
            case "\n\t":
                var v57 uint32 = 2147483648
                _ = v57
            case "hello":
                for fi58 := 0; fi58 < 10; fi58++ {
                    v59 := GenS4{F0: 42, F1: v49, F2: (fi58 + fi58)}
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    _ = fi58
                    _ = v59
                }
            }
            _ = v55
        }
        v60 := GenS5{F0: 4294967295, F1: 1}
        _ = fi54
        _ = v60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
    _ = v8
}
