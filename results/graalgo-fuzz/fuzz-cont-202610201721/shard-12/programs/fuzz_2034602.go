package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int32 }
func genFn0(p0 uint8) int32 {
    if false {
        var v6 int16 = 1000
        _ = v6
    } else {
        p0 = 2
    }
    p0 = (p0 - (p0 % ((2) | 1)))
    _ = p0
    return -1
}
func genFn1(p0 int32, p1 int8) uint8 {
    p0 += (p0 ^ p0)
    fmt.Println(p0, p0)
    p0 = (p0 << ((52) & 63))
    _ = p0
    _ = p1
    return 234
}

func main() {
    var v1 int32 = 42
    var v2 float32 = 1.5
    v3 := make(map[string]string)
    var v4 int32 = (v1 / ((v1) | 1))
    var v5 uint64 = 549
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint = 354
        v9 := v3["foo bar"]
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 string = v9
            if true {
                var v12 float32 = (1e10 - float32((float64((v2 + ((3.14159 + v2) / v2))) * 0.1)))
                if (1 <= -1) {
                    if (42 < 382) {
                        switch (fi7 - (fi10 + (fi7 ^ (^(fi7))))) {
                        case -131:
                            switch v9 {
                            case "a":
                                v13 := GenS1{F0: (!true), F1: 2147483647, F2: (v5 + 2)}
                                _ = v13
                            case "foo bar":
                                v8 += 255
                            default:
                                if true {
                                    v3[(v9 + v11)] = (v11 + (v9 + (v9 + v11)))
                                    if false {
                                        v14 := v3[""]
                                        _ = v14
                                    }
                                    switch fi7 {
                                    case -739:
                                        v15 := GenS2{F0: true, F1: int64(v1), F2: -9223372036854775808}
                                        _ = v15
                                    case -9223372036854775808:
                                        if false {
                                            v2 = (v2 + -2.25)
                                            switch (fi7 << (((v8 >> (((v8 | 255)) & 63))) & 63)) {
                                            case -437:
                                                v8 -= v8
                                            }
                                        }
                                    case 42:
                                        var v16 int64 = 1
                                        _ = v16
                                    default:
                                        switch v9 {
                                        case "foo bar":
                                            fmt.Println(v5, v5, v4)
                                        }
                                    }
                                    v17 := GenS1{F0: (v1 <= v4), F1: 2, F2: 31}
                                    v18 := []int{(fi7 - fi7), fi7, (-(fi7))}
                                    _ = v17
                                    _ = v18
                                } else {
                                    v19 := GenS1{F0: true, F1: 2, F2: (v5 + (v5 >> (((^(v8))) & 63)))}
                                    switch v11 {
                                    case "\n\t":
                                        fmt.Println(v12, v4)
                                    case "":
                                        v20 := GenS3{F0: "", F1: 2147483647}
                                        _ = v20
                                    case "a":
                                        if (float32(0.1) <= v12) {
                                            v21 := []int{(fi7 >> ((v8) & 63))}
                                            _ = v21
                                        } else {
                                            var v22 uint8 = genFn1((v4 & v1), -1)
                                            v1 += 2147483647
                                            fmt.Println(v11, v4)
                                            v23 := GenS2{F0: true, F1: -9223372036854775808, F2: int64(v8)}
                                            var v24 int8 = 42
                                            v5 = (^(v5))
                                            v4 -= -84
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                    default:
                                        fmt.Println(v9, v2)
                                    }
                                    var v25 int16 = 42
                                    _ = v19
                                    _ = v25
                                }
                                delete(v3, (v11 + v11))
                                fmt.Println(v12, v2, fi10)
                            }
                        case 359:
                            v11 = ("foo bar" + v11)
                        default:
                            v3["\n\t"] = ("" + (v11 + v11))
                        }
                        v26 := []int{}
                        if false {
                            fmt.Println(v1, v4)
                        }
                        _ = v26
                    }
                    var v27 string = "a"
                    _ = v27
                }
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        if (-1 != 32767) {
            var v28 int64 = 9223372036854775807
            fmt.Println(v4)
            _ = v28
        } else {
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 uint16 = 1
                v31 := GenS1{F0: false, F1: 4294967295, F2: 891}
                if (4294967295 >= 1) {
                    fmt.Println(v1)
                } else {
                    var v32 int = 32767
                    v4 -= (v4 / ((-48) | 1))
                    v33 := GenS3{F0: "foo bar", F1: 327}
                    _ = v32
                    _ = v33
                }
                _ = fi29
                _ = v30
                _ = v31
            }
            var v34 float64 = 1.5
            _ = v34
        }
        v4 = -400
        v5 = v5
    }
    var v35 float64 = 0.1
    fmt.Println(v4, v4, v5)
    v36 := v3[("a" + (("\n\t" + ("\n\t" + ("foo bar" + ""))) + "foo bar"))]
    var v37 int64 = -245
    v38 := v3[v36]
    for fi39 := 0; fi39 < 16; fi39++ {
        for fi40 := 0; fi40 < 12; fi40++ {
            if ((2147483648 <= 424) || true) {
                var v41 string = (("hello" + "foo bar") + v36)
                var v42 uint64 = 18446744073709551615
                v43 := []int{fi39, 116, (fi40 * (-(fi39)))}
                _ = v41
                _ = v42
                _ = v43
            }
            v44 := []int{fi39}
            _ = fi40
            _ = v44
        }
        v45 := GenS2{F0: (0.0 == (1.5 / 1e10)), F1: (^(v37)), F2: v37}
        _ = fi39
        _ = v45
    }
    var v46 float64 = (1.5 + 1e10)
    for fi47 := 0; fi47 < 14; fi47++ {
        for fi48 := 0; fi48 < 18; fi48++ {
            v2 = 1.5
            _ = fi48
        }
        for fi49 := 0; fi49 < 5; fi49++ {
            v50 := GenS4{F0: 42, F1: (v4 ^ (v4 * 2147483647))}
            var v51 int64 = (v37 ^ (v37 / ((0) | 1)))
            fmt.Println(fi47)
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = fi47
    }
    for fi52 := 0; fi52 < 11; fi52++ {
        var v53 bool = (!(-966 >= v37))
        v54 := GenS2{F0: true, F1: (v37 + (^(v37))), F2: v37}
        var v55 uint16 = 0
        _ = fi52
        _ = v53
        _ = v54
        _ = v55
    }
    for fi56 := 0; fi56 < 20; fi56++ {
        var v57 uint16 = 255
        _ = fi56
        _ = v57
    }
    v37 -= v37
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v46
    _ = v5
}
