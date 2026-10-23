package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int64; F2 int }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 bool; F1 uint8; F2 int64 }
func genFn0(p0 uint8, p1 uint, p2 string) uint16 {
    p0 = p0
    fmt.Println(p2, p0)
    p0 = p0
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{(fi7 | -1), fi7, fi7, (-(fi7))}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = len(v8)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 uint32, p1 string, p2 uint64) uint32 {
    p2 = p2
    switch p1 {
    case "":
        v9 := GenS1{F0: true, F1: p2}
        _ = v9
    case "a":
        for fi10 := 0; fi10 < 9; fi10++ {
            p0 = 42
            fmt.Println(fi10)
            _ = fi10
        }
    case "foo bar":
        if (!true) {
            fmt.Println(p1)
            v11 := []int{}
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                if false {
                    switch fi12 {
                    case 136:
                        var v14 uint16 = 65535
                        _ = v14
                    case 325:
                        var v15 int32 = -692
                        _ = v15
                    default:
                        if ((!true) && true) {
                            var v16 string = "foo bar"
                            v17 := []int{(^(fi12)), (fi13 / ((1) | 1)), 9223372036854775807}
                            fmt.Println(p0)
                            _ = v16
                            _ = v17
                        } else {
                            switch "\n\t" {
                            case "foo bar":
                                v18 := GenS1{F0: (!false), F1: 18446744073709551615}
                                _ = v18
                            case "":
                                if true {
                                    v19 := []int{fi13, fi13, 42}
                                    if true {
                                        v20 := GenS2{F0: (1e10 <= 1.5)}
                                        _ = v20
                                    } else {
                                        fmt.Println(p0, fi13, fi12)
                                        if (true && (-849 <= 2147483647)) {
                                            p1 = "\n\t"
                                        }
                                        fmt.Println(p1)
                                    }
                                    v21 := GenS3{F0: 127, F1: 0, F2: (fi12 >> ((18446744073709551615) & 63))}
                                    _ = v19
                                    _ = v21
                                } else {
                                    fmt.Println(fi12, p0)
                                    var v22 int = 9223372036854775807
                                    var v23 uint = 65535
                                    _ = v22
                                    _ = v23
                                }
                                v24 := GenS4{F0: (fi12 ^ fi13), F1: -680}
                                p2 = (^(p2))
                                switch fi12 {
                                case -947:
                                    v25 := GenS2{F0: (true || false)}
                                    _ = v25
                                case -9223372036854775808:
                                    p0 -= p0
                                default:
                                    var v26 int = fi12
                                    p1 = (("a" + ((p1 + p1) + "foo bar")) + "foo bar")
                                    if (!false) {
                                        p0 = p0
                                        if (921 == 96) {
                                            p0 = p0
                                        }
                                    }
                                    _ = v26
                                }
                                if true {
                                    p1 = p1
                                }
                                if false {
                                    v27 := []int{fi12, (fi13 * fi12), fi12}
                                    var v28 uint32 = (^(p0))
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                _ = v24
                            case "a":
                                if (p1 <= p1) {
                                    var v29 string = p1
                                    _ = v29
                                }
                            default:
                                if true {
                                    var v30 float64 = 0.0
                                    _ = v30
                                }
                            }
                            var v31 string = p1
                            _ = v31
                        }
                    }
                    switch p1 {
                    case "":
                        var v32 float32 = float32((0.1 - 1e10))
                        _ = v32
                    default:
                        if (-2.25 >= (1.5 + (3.14159 * (0.0 + -2.25)))) {
                            v33 := []int{int(p2), (^(fi12))}
                            var v34 uint16 = 1
                            if false {
                                var v35 bool = false
                                v36 := GenS2{F0: (880 <= uint(p2))}
                                fmt.Println(fi13)
                                _ = v35
                                _ = v36
                            } else {
                                var v37 float64 = 3.14159
                                _ = v37
                            }
                            _ = v33
                            _ = v34
                        } else {
                            v38 := []int{}
                            v39 := GenS2{F0: false}
                            var v40 uint16 = 65535
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                    }
                    v41 := []int{}
                    _ = v41
                } else {
                    v42 := GenS2{F0: ((!((((0.1 * float64(-2.25)) / (-2.25 + 1e10)) * float64(3.14159)) > 3.14159)) || true)}
                    p1 = p1
                    _ = v42
                }
                v43 := []int{fi13, fi13, (-(fi12))}
                _ = fi13
                _ = v43
            }
            for fi44 := 0; fi44 < 10; fi44++ {
                var v45 int8 = 127
                _ = fi44
                _ = v45
            }
            _ = fi12
        }
    }
    v46 := []int{}
    v47 := GenS2{F0: false}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v46
    _ = v47
    return 2
}

func main() {
    var v1 bool = true
    var v2 uint16 = 545
    var v3 int64 = 1
    v4 := make(map[uint32]bool)
    var v5 string = "hello"
    var v6 string = v5
    v48 := GenS5{F0: v1, F1: 0, F2: (v3 * v3)}
    v1 = (1 < int16(v2))
    v49 := GenS4{F0: 339, F1: -651}
    v1 = v1
    var v50 int8 = 42
    v51 := []int{-470, -1}
    v2 -= (v2 / ((v2) | 1))
    v3 = v3
    if (127 >= 127) {
        for fi52 := 0; fi52 < 14; fi52++ {
            fmt.Println(fi52, v6)
            if false {
                var v53 float64 = 3.14159
                _ = v53
            } else {
                delete(v4, genFn1(genFn1(42, v6, 1), "\n\t", 869))
            }
            var v54 uint16 = (^(v2))
            _ = fi52
            _ = v54
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
