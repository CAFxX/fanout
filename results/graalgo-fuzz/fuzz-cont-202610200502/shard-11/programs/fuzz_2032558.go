package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int8 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int32; F1 string; F2 bool }
func genFn0(p0 bool, p1 string) uint16 {
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(p1)
        var v8 uint64 = 844
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 bool = p0
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int8 = 93
        p1 = "hello"
        v13 := []int{(-(fi11))}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int = 259
    _ = p0
    _ = p1
    _ = v14
    return 42
}
func genFn1(p0 uint, p1 int16) int32 {
    var v15 int = 2147483647
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 float64 = (((1.5 - 3.14159) * 1.5) - 0.0)
        v18 := []int{0, (^(fi16))}
        for fi19 := 0; fi19 < 2; fi19++ {
            if false {
                if (!true) {
                    if true {
                        if (42 > 937) {
                            v15 = (fi16 ^ fi19)
                        } else {
                            if ((v15 ^ v15) != fi19) {
                                fmt.Println(fi16, fi16)
                                if false {
                                    var v20 uint64 = 558
                                    var v21 uint64 = 1
                                    _ = v20
                                    _ = v21
                                }
                            }
                            v18 = append(v18, fi16)
                        }
                        fmt.Println(p0, fi19)
                        if true {
                            var v22 int64 = -677
                            _ = v22
                        }
                    }
                    p0 += (p0 >> (((p0 / (((p0 / (((p0 / (((p0 << (((^(p0))) & 63))) | 1))) | 1))) | 1))) & 63))
                }
                v18 = append(v18, fi19)
            }
            if len(v18) > 0 {
                v18[0] = fi16
            }
            _ = fi19
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    var v23 uint64 = 42
    v24 := GenS1{F0: 2147483647, F1: 42, F2: -128}
    _ = p0
    _ = p1
    _ = v15
    _ = v23
    _ = v24
    return 105
}

func main() {
    var v1 string = ("hello" + "")
    v2 := []string{v1, v1, ""}
    var v3 uint = 560
    var v4 int32 = 886
    var v5 int = 605
    var v6 uint8 = 149
    for fi25 := 0; fi25 < 15; fi25++ {
        switch 32768 {
        case 651:
            _ = len(v2)
        case -1:
            for fi26 := 0; fi26 < 16; fi26++ {
                var v27 uint16 = 60
                _ = fi26
                _ = v27
            }
        case 2147483647:
            switch fi25 {
            case 127:
                switch v1 {
                case "a":
                    v2 = append(v2, "\n\t")
                case "\n\t":
                    v28 := GenS2{F0: (^(v6)), F1: genFn1(135, 32767), F2: 255}
                    _ = v28
                case "foo bar":
                    for fi29 := 0; fi29 < 5; fi29++ {
                        v30 := GenS2{F0: 2, F1: int32(fi25), F2: 32767}
                        if len(v2) > 0 {
                            v2[0] = (("\n\t" + v1) + ("foo bar" + v1))
                        }
                        var v31 uint8 = (v6 ^ uint8(v3))
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
            case 42:
                fmt.Println(fi25, v6, fi25)
            }
        default:
            v4 += genFn1((^(v3)), 788)
        }
        if false {
            var v32 int16 = 0
            _ = v32
        } else {
            if len(v2) > 0 {
                v2[0] = v1
            }
            v4 += (v4 * 42)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi25
    }
    var v33 float64 = float64(1.5)
    if true {
        v34 := GenS1{F0: 1, F1: -9223372036854775808, F2: 0}
        _ = v34
    } else {
        var v35 string = v1
        _ = v35
    }
    v5 = v5
    var v36 bool = true
    var v37 uint64 = 399
    if len(v2) > 0 {
        v2[0] = v1
    }
    v37 = v37
    var v38 uint16 = genFn0((0.0 >= float64(1e10)), v1)
    _ = len(v2)
    for fi39 := 0; fi39 < 20; fi39++ {
        if (v6 != v6) {
            for fi40 := 0; fi40 < 16; fi40++ {
                fmt.Println(v37)
                _ = fi40
            }
            switch (v1 + v1) {
            case "a":
                v4 += 2147483647
            }
            if v36 {
                switch 0 {
                case 32768:
                    for fi41 := 0; fi41 < 19; fi41++ {
                        var v42 int8 = 16
                        var v43 float64 = v33
                        if false {
                            switch v5 {
                            case -470:
                                var v44 string = v1
                                _ = v44
                            case 4294967295:
                                var v45 float64 = (3.14159 * v33)
                                _ = v45
                            case 42:
                                v46 := GenS3{F0: (v42 % (((v42 * 42)) | 1)), F1: 32767}
                                _ = v46
                            }
                        }
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                default:
                    if true {
                        v37 += 367
                    }
                }
                switch ("" + (v1 + v1)) {
                case "\n\t":
                    var v47 bool = ((("" + "hello") > v1) && v36)
                    _ = v47
                }
                switch 9223372036854775807 {
                case 174:
                    fmt.Println(v5, v33)
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if v36 {
                    fmt.Println(v1, v33, v38)
                    var v48 uint = (v3 + 759)
                    for fi49 := 0; fi49 < 5; fi49++ {
                        v38 += (v38 << ((v48) & 63))
                        if ((uint32(v6) != 4294967295) && ((false && (v36 && true)) || v36)) {
                            if (!v36) {
                                v1 = (v1 + v1)
                            }
                            fmt.Println(v5, fi49)
                            v33 -= (float64(-2.25) / ((0.0 * v33) + v33))
                        }
                        if (!true) {
                            if (!v36) {
                                switch (fi39 - 42) {
                                case -863:
                                    var v50 float64 = 1.5
                                    _ = v50
                                default:
                                    fmt.Println(v37, v38)
                                    if v36 {
                                        var v51 uint8 = (v6 ^ v6)
                                        fmt.Println(v3, v48)
                                        _ = v51
                                    }
                                }
                            } else {
                                var v52 int8 = 127
                                _ = v52
                            }
                            var v53 bool = (fi39 < fi39)
                            if (v3 <= v48) {
                                v54 := GenS3{F0: 42, F1: 1000}
                                _ = v54
                            } else {
                                fmt.Println(v53)
                                var v55 uint8 = v6
                                _ = v55
                            }
                            v33 -= ((((v33 / (v33 + (0.1 * v33))) - 0.0) / float64(3.14159)) - 3.14159)
                            var v56 int = (fi49 ^ fi39)
                            _ = v53
                            _ = v56
                        }
                        _ = fi49
                    }
                    _ = v48
                } else {
                    fmt.Println(v6)
                    var v57 float64 = v33
                    _ = v57
                }
                v58 := GenS4{F0: (v4 | (v4 / (((v4 - v4)) | 1))), F1: v1, F2: v36}
                _ = v58
            }
        }
        _ = fi39
    }
    var v59 int = (v5 << ((v3) & 63))
    fmt.Println(v38)
    fmt.Println(v38)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
