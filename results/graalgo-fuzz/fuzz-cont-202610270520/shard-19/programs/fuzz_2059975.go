package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 int16; F1 string }
func genFn0(p0 uint32) int {
    var v5 int64 = 9223372036854775807
    if (!false) {
        v6 := []int{-9223372036854775808}
        v5 += (v5 / (((v5 << ((18446744073709551615) & 63))) | 1))
        _ = v6
    }
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 int64, p1 int32) bool {
    fmt.Println(p0, p1)
    v7 := GenS1{F0: -500}
    _ = p0
    _ = p1
    _ = v7
    return (true && false)
}

func main() {
    var v1 float64 = (-2.25 * 0.1)
    v2 := []uint32{65535, 642, 912}
    v3 := []int{-282, 37}
    v4 := make(map[string]bool)
    var v8 int = -1
    var v9 int16 = 386
    v10 := GenS1{F0: (v8 * 2147483647)}
    v1 = (v1 + v1)
    v11 := GenS1{F0: genFn0(2147483648)}
    v12 := GenS1{F0: v8}
    if len(v3) > 0 {
        v3[0] = (v8 / ((v8) | 1))
    }
    v1 += v1
    var v13 uint64 = 18446744073709551615
    switch v8 {
    case 32767:
        switch "a" {
        case "a":
            if (42 != 0) {
                if (!true) {
                    var v14 int64 = 1
                    switch "" {
                    case "a":
                        v15 := GenS1{F0: v8}
                        _ = v15
                    case "":
                        fmt.Println(v14)
                    default:
                        for fi16 := 0; fi16 < 13; fi16++ {
                            var v17 int8 = -128
                            for fi18 := 0; fi18 < 19; fi18++ {
                                v1 -= ((1e10 - (v1 / 1e10)) + (((v1 + (v1 - (float64(-2.25) + float64((1.5 / float32(0.1)))))) * v1) + (v1 / (v1 * float64(float32(1e10))))))
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi18
                            }
                            if true {
                                if true {
                                    for fi19 := 0; fi19 < 2; fi19++ {
                                        if (42 <= v9) {
                                            if true {
                                                v13 += (v13 & (v13 % ((18446744073709551615) | 1)))
                                            } else {
                                                if (1 == v9) {
                                                    v14 = (v14 >> ((42) & 63))
                                                }
                                            }
                                            var v20 int = v8
                                            v8 = (-(fi19))
                                            _ = v20
                                        } else {
                                            v21 := GenS1{F0: (-(v8))}
                                            if (((v17 << ((739) & 63)) == 0) || false) {
                                                switch fi19 {
                                                case -534:
                                                    var v22 int8 = (v17 - v17)
                                                    _ = v22
                                                case -359:
                                                    if (448 < 198) {
                                                        v23 := GenS2{F0: ""}
                                                        var v24 float32 = 1e10
                                                        switch "" {
                                                        case "\n\t":
                                                            if (true || false) {
                                                                v13 = 704
                                                            } else {
                                                                switch ("a" + "hello") {
                                                                case "hello":
                                                                    if true {
                                                                        switch (-(fi19)) {
                                                                        case 233:
                                                                            var v25 uint = 114
                                                                            _ = v25
                                                                        default:
                                                                            fmt.Println(fi19, v14)
                                                                        }
                                                                    }
                                                                    fmt.Println(v17, v24, fi16)
                                                                    var v26 string = "\n\t"
                                                                    _ = v26
                                                                default:
                                                                    fmt.Println(v14)
                                                                }
                                                                var v27 int16 = v9
                                                                _ = v27
                                                            }
                                                        default:
                                                            v4[""] = ((v14 / ((v14) | 1)) >= (-(v14)))
                                                            v28 := GenS3{F0: -288}
                                                            v29 := GenS4{F0: "foo bar", F1: (v13 - 42), F2: v13}
                                                            _ = v28
                                                            _ = v29
                                                        }
                                                        _ = v23
                                                        _ = v24
                                                    }
                                                    var v30 uint = 18446744073709551615
                                                    v14 = (-(v14))
                                                    _ = v30
                                                case 127:
                                                    v31 := GenS5{F0: v13, F1: false, F2: v13}
                                                    fmt.Println(v14)
                                                    _ = v31
                                                }
                                            }
                                            _ = v21
                                        }
                                        _ = fi19
                                    }
                                    fmt.Println(v14)
                                    for fi32 := 0; fi32 < 5; fi32++ {
                                        if true {
                                            if false {
                                                if true {
                                                    if true {
                                                        fmt.Println(v17)
                                                        _ = len(v4)
                                                    } else {
                                                        if true {
                                                            v33 := v4[("foo bar" + "\n\t")]
                                                            _ = v33
                                                        }
                                                        var v34 uint64 = v13
                                                        _ = v34
                                                    }
                                                } else {
                                                    if (((true && (!false)) && (v13 > 42)) && false) {
                                                        var v35 uint64 = (^(v13))
                                                        _ = v35
                                                    } else {
                                                        var v36 int = v8
                                                        var v37 bool = (0.1 == (1e10 + (1.5 / 0.1)))
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    var v38 float64 = v1
                                                    _ = v38
                                                }
                                            }
                                            v14 -= 42
                                            if (2147483647 != 535) {
                                                v39 := GenS6{F0: 756, F1: v17}
                                                if ((0 >= 1) || (true || true)) {
                                                    v40 := GenS2{F0: ("\n\t" + (("a" + "foo bar") + ""))}
                                                    var v41 uint32 = uint32(v13)
                                                    switch "" {
                                                    case "\n\t":
                                                        v1 = v1
                                                    case "a":
                                                        var v42 float64 = v1
                                                        _ = v42
                                                    }
                                                    _ = v40
                                                    _ = v41
                                                } else {
                                                    var v43 float32 = (0.1 * (float32(3.14159) * (1.5 + ((1e10 - -2.25) * float32(v1)))))
                                                    if (18446744073709551615 >= 65535) {
                                                        v44 := GenS7{F0: "hello"}
                                                        v45 := GenS8{F0: -1, F1: 127}
                                                        fmt.Println(fi32)
                                                        v17 += v17
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                    _ = v43
                                                }
                                                _ = v39
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = fi32
                                                }
                                                var v46 bool = (-1 >= -872)
                                                _ = v46
                                            }
                                            var v47 uint16 = 42
                                            v17 -= (v17 >> ((18446744073709551615) & 63))
                                            v8 = -9223372036854775808
                                            _ = len(v3)
                                            v14 -= v14
                                            var v48 int32 = 1
                                            _ = v47
                                            _ = v48
                                        } else {
                                            fmt.Println(fi16, v13, v14)
                                            _ = len(v2)
                                        }
                                        _ = fi32
                                    }
                                }
                                v49 := GenS9{F0: (v9 % ((v9) | 1)), F1: ""}
                                v14 += 42
                                _ = v49
                            }
                            _ = fi16
                            _ = v17
                        }
                    }
                    for fi50 := 0; fi50 < 9; fi50++ {
                        var v51 float64 = 1.5
                        _ = fi50
                        _ = v51
                    }
                    _ = v14
                } else {
                    var v52 bool = true
                    _ = v52
                }
            }
        default:
            v4["a"] = true
        }
    case 9223372036854775807:
        var v53 int64 = 42
        _ = v53
    default:
        v54 := GenS5{F0: (v13 ^ (^(v13))), F1: ((1e10 < -2.25) && (0 != 637)), F2: (v13 << ((379) & 63))}
        _ = v54
    }
    if false {
        switch v8 {
        case 32767:
            var v55 uint64 = v13
            _ = v55
        case 127:
            var v56 uint64 = v13
            _ = v56
        case 35:
            var v57 string = "a"
            _ = v57
        default:
            for fi58 := 0; fi58 < 12; fi58++ {
                var v59 uint8 = uint8(v13)
                v1 = (0.1 + ((float64(0.0) + float64(1e10)) - (3.14159 + 0.1)))
                _ = fi58
                _ = v59
            }
        }
        v60 := v4["hello"]
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
