package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int16; F1 int; F2 uint64 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 uint; F1 uint8; F2 int64 }
type GenS8 struct { F0 int }

func main() {
    var v1 uint32 = 4294967295
    var v2 float32 = 3.14159
    v3 := make(map[int]bool)
    v2 = 1e10
    v1 = (v1 & v1)
    delete(v3, -9223372036854775808)
    v2 -= v2
    switch "\n\t" {
    case "\n\t":
        v4 := GenS1{F0: 325, F1: (v1 % ((235) | 1))}
        _ = v4
    case "hello":
        if (1 <= 398) {
            v5 := GenS1{F0: 255, F1: (v1 % ((v1) | 1))}
            _ = v5
        } else {
            var v6 float64 = (0.1 * 0.1)
            v7 := GenS1{F0: 0, F1: (v1 + (^(v1)))}
            _ = v6
            _ = v7
        }
    default:
        switch ("" + ((("\n\t" + "hello") + "") + ("a" + "a"))) {
        case "foo bar":
            if (true || false) {
                v8 := v3[1]
                if (true && v8) {
                    v9 := GenS2{F0: 32767, F1: 748}
                    v10 := GenS1{F0: 348, F1: v1}
                    v11 := GenS2{F0: 0, F1: 1}
                    _ = v10
                    _ = v11
                    _ = v9
                }
                v12 := GenS2{F0: 0, F1: 0}
                _ = v12
                _ = v8
            } else {
                var v13 bool = true
                switch 0 {
                case 2147483647:
                    v1 = (v1 | v1)
                default:
                    if (uint(v1) != 255) {
                        var v14 int = 2147483647
                        var v15 int64 = -406
                        var v16 string = "foo bar"
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
                _ = v13
            }
        }
    }
    if false {
        v17 := v3[2147483647]
        v18 := GenS3{F0: 1, F1: 32767, F2: 302}
        v17 = (59 > 127)
        _ = v17
        _ = v18
    }
    if true {
        var v19 uint8 = 128
        _ = v19
    } else {
        v1 += 1
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 int32 = 344
            var v22 string = "a"
            _ = fi20
            _ = v21
            _ = v22
        }
        if (false && false) {
            v23 := GenS4{F0: 53, F1: 0}
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v1 -= v1
        _ = fi24
    }
    if ((0.1 + 0.1) >= -2.25) {
        var v25 uint64 = 42
        v26 := GenS5{F0: v1, F1: 176, F2: (v1 - v1)}
        _ = v25
        _ = v26
    }
    v1 = v1
    for fi27 := 0; fi27 < 7; fi27++ {
        v2 = (v2 - float32(1e10))
        for fi28 := 0; fi28 < 3; fi28++ {
            if ((v2 / 1e10) != (3.14159 * v2)) {
                var v29 int8 = -1
                if false {
                    var v30 float64 = (0.1 - 3.14159)
                    var v31 string = (((("a" + ("" + "foo bar")) + (("\n\t" + "a") + "a")) + "") + (("" + ("" + ("\n\t" + ("" + "foo bar")))) + "a"))
                    v2 = (1.5 - v2)
                    _ = v30
                    _ = v31
                } else {
                    var v32 uint8 = 1
                    v29 = v29
                    _ = v32
                }
                v3[(fi28 ^ 128)] = true
                _ = v29
            } else {
                v2 = 1.5
                if false {
                    switch (^(fi28)) {
                    case 32767:
                        var v33 int32 = 364
                        _ = v33
                    case -9223372036854775808:
                        var v34 uint64 = 18446744073709551615
                        _ = v34
                    default:
                        v35 := GenS6{F0: (594 != (fi27 * (fi27 ^ (fi28 + (fi28 | fi27))))), F1: fi27}
                        _ = v35
                    }
                    if false {
                        var v36 int64 = -1
                        switch fi28 {
                        case 1:
                            _ = len(v3)
                        default:
                            if (923 < 255) {
                                if false {
                                    v36 = v36
                                    if (-9223372036854775808 == v36) {
                                        if false {
                                            fmt.Println(v2, fi27)
                                            if false {
                                                v1 = v1
                                                v37 := []int{fi28, fi27}
                                                if false {
                                                    v38 := GenS4{F0: 18, F1: fi27}
                                                    if (4294967295 == 0) {
                                                        var v39 uint64 = 287
                                                        var v40 float32 = (v2 - (float32((-2.25 - 0.0)) * float32((float64(((-2.25 - (1.5 - 1e10)) + 1.5)) * ((0.0 + float64(v2)) + ((3.14159 / 3.14159) - (((float64(-2.25) * (float64((v2 - -2.25)) + 3.14159)) - (0.1 * (3.14159 * 0.1))) * ((1.5 + (1.5 + 3.14159)) + ((3.14159 / ((0.1 + (((float64(1e10) + -2.25) + 0.1) + 0.1)) + (-2.25 * (-2.25 + (1e10 - (-2.25 - ((1e10 + float64(v2)) / 1.0))))))) + 1e10)))))))))
                                                        _ = v39
                                                        _ = v40
                                                    } else {
                                                        delete(v3, (fi28 << ((352) & 63)))
                                                        if (18446744073709551615 != 2) {
                                                            var v41 int16 = 1000
                                                            _ = v41
                                                        }
                                                    }
                                                    fmt.Println(v1, v2, v36)
                                                    _ = v38
                                                }
                                                _ = v37
                                            } else {
                                                v42 := []int{(fi27 << ((583) & 63))}
                                                if (false || true) {
                                                    var v43 uint64 = 0
                                                    _ = v43
                                                }
                                                _ = v42
                                            }
                                            v44 := []int{562}
                                            _ = v44
                                        }
                                        v45 := GenS5{F0: v1, F1: 164, F2: (v1 % ((v1) | 1))}
                                        v46 := GenS7{F0: 596, F1: 2, F2: (-(v36))}
                                        _ = v45
                                        _ = v46
                                    }
                                } else {
                                    var v47 int64 = 9223372036854775807
                                    v48 := GenS8{F0: -1}
                                    _ = v47
                                    _ = v48
                                }
                            }
                        }
                        v49 := GenS6{F0: (true || (42 >= 0)), F1: (fi27 ^ fi28)}
                        _ = v36
                        _ = v49
                    }
                    if (!false) {
                        v50 := GenS8{F0: (fi27 % ((fi27) | 1))}
                        _ = v50
                    }
                }
            }
            _ = fi28
        }
        _ = fi27
    }
    if (1.5 == -2.25) {
        var v51 uint16 = 42
        _ = v51
    }
    v1 += v1
    fmt.Println(v1, v1)
    var v52 uint32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
}
