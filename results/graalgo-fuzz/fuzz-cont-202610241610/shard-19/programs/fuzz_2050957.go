package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint64; F1 bool; F2 string }
type GenS3 struct { F0 string; F1 uint8; F2 int }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint) string {
    p0 = 42
    switch ((("a" + "\n\t") + (((("a" + "foo bar") + "a") + "hello") + "\n\t")) + "\n\t") {
    case "a":
        var v7 float32 = 0.1
        _ = v7
    case "hello":
        if false {
            p0 -= p0
            var v8 int8 = 1
            _ = v8
        }
    default:
        p0 += 14
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 uint64 = 42
        switch (("hello" + "a") + "\n\t") {
        case "a":
            var v11 uint16 = 2
            _ = v11
        }
        v12 := GenS1{F0: 42, F1: fi9}
        _ = fi9
        _ = v10
        _ = v12
    }
    var v13 int32 = 0
    _ = p0
    _ = p1
    _ = v13
    return ""
}
func genFn1(p0 int8) uint16 {
    fmt.Println(p0)
    var v14 int = 0
    var v15 bool = false
    _ = p0
    _ = v14
    _ = v15
    return 2
}

func main() {
    v1 := []int{-9223372036854775808}
    var v2 string = "hello"
    var v3 int64 = -628
    v4 := make(map[uint32]string)
    var v5 int16 = -32768
    v6 := []string{(v2 + "\n\t"), "hello"}
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(v5, v5)
        _ = fi16
    }
    switch 32767 {
    case 9223372036854775807:
        for fi17 := 0; fi17 < 6; fi17++ {
            switch (fi17 * fi17) {
            case -944:
                fmt.Println(v5, v3)
            default:
                var v18 uint = 535
                _ = v18
            }
            v3 = v3
            _ = fi17
        }
    case 686:
        if false {
            if false {
                var v19 bool = true
                v3 = 799
                _ = v19
            }
        } else {
            var v20 uint64 = 918
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(v3)
        for fi22 := 0; fi22 < 5; fi22++ {
            v4[2] = "hello"
            v2 = v2
            var v23 string = ""
            if false {
                if ((0.0 > 0.1) && false) {
                    _ = len(v4)
                    v5 = v5
                    v24 := GenS1{F0: v5, F1: fi21}
                    _ = v24
                } else {
                    fmt.Println(v23)
                    switch 128 {
                    case -171:
                        v25 := GenS1{F0: int16(v3), F1: 9223372036854775807}
                        _ = v25
                    default:
                        var v26 int32 = 220
                        _ = v26
                    }
                    if (840 != 1) {
                        if (!true) {
                            if (((v23 < v23) || (32767 != 2)) && (116 <= 128)) {
                                var v27 float64 = (0.1 + 0.1)
                                if ((!false) && (false || true)) {
                                    switch "a" {
                                    case "hello":
                                        var v28 float32 = 1e10
                                        _ = v28
                                    case "foo bar":
                                        var v29 uint16 = 542
                                        _ = v29
                                    case "a":
                                        v3 += v3
                                    default:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    }
                                    switch v2 {
                                    case "hello":
                                        _ = len(v4)
                                    case "":
                                        v30 := GenS2{F0: 628, F1: true, F2: (v23 + ((("hello" + v23) + v2) + v23))}
                                        _ = v30
                                    }
                                } else {
                                    fmt.Println(v5)
                                }
                                _ = v27
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v31 int32 = 0
                                v32 := GenS1{F0: 0, F1: (-(fi21))}
                                _ = v31
                                _ = v32
                            }
                            v23 = (v2 + v2)
                        }
                        v33 := GenS3{F0: "\n\t", F1: 86, F2: 509}
                        _ = v33
                    }
                }
                v23 = ("\n\t" + (v2 + v23))
            } else {
                if false {
                    switch v23 {
                    case "hello":
                        fmt.Println(fi22, v3)
                    }
                    fmt.Println(fi22)
                } else {
                    var v34 uint32 = 0
                    fmt.Println(fi21, v23)
                    _ = v34
                }
                v35 := GenS3{F0: v23, F1: 2, F2: (fi21 << ((2) & 63))}
                _ = v35
            }
            if ((("\n\t" + v23) + (((v23 + (v23 + (v2 + v23))) + ("hello" + v2)) + (v2 + v2))) >= (v23 + "hello")) {
                if (true && true) {
                    if (509 < 2) {
                        v36 := GenS2{F0: 1, F1: (4294967295 != 765), F2: v23}
                        _ = v36
                    }
                    delete(v4, 2147483648)
                    if false {
                        if (false || false) {
                            v37 := GenS3{F0: v2, F1: 127, F2: (fi21 ^ (fi21 / ((fi21) | 1)))}
                            fmt.Println(v23, fi22, v5)
                            v38 := GenS2{F0: 1, F1: (0 == 42), F2: (("" + ((v2 + "a") + ("hello" + "hello"))) + v2)}
                            _ = v37
                            _ = v38
                        } else {
                            if (((0.1 != 1e10) && true) && (2147483647 != 1)) {
                                v39 := GenS1{F0: 0, F1: fi22}
                                _ = v39
                            }
                        }
                        _ = len(v1)
                        var v40 float32 = (0.1 * (1e10 * 1.5))
                        _ = v40
                    } else {
                        var v41 int64 = v3
                        _ = v41
                    }
                } else {
                    switch 4294967295 {
                    case -9223372036854775808:
                        var v42 int = 587
                        _ = v42
                    case -775:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        var v43 int8 = 0
                        _ = v43
                    }
                }
                if true {
                    var v44 int32 = 2147483647
                    var v45 uint16 = genFn1(111)
                    v23 = (v2 + v23)
                    _ = v44
                    _ = v45
                } else {
                    var v46 uint8 = 1
                    if false {
                        var v47 bool = (!true)
                        _ = v47
                    }
                    v2 = "a"
                    fmt.Println(v5, v2)
                    _ = v46
                }
            } else {
                fmt.Println(v23)
                var v48 int32 = 1
                v3 += -1
                _ = v48
            }
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    v4[2] = "foo bar"
    v3 = v3
    v3 = v3
    var v49 int64 = 9223372036854775807
    var v50 uint = 4294967295
    for fi51 := 0; fi51 < 20; fi51++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v52 uint64 = 1
        v2 = v2
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 10; fi53++ {
        v54 := GenS4{F0: 453, F1: -588}
        v49 += -29
        _ = fi53
        _ = v54
    }
    fmt.Println(v3, v50, v50)
    v55 := GenS5{F0: genFn1(1)}
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v55
    _ = v6
}
