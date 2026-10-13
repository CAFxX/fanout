package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint8) uint64 {
    var v7 bool = true
    var v8 uint = 1
    _ = p0
    _ = v7
    _ = v8
    return uint64(v8)
}
func genFn1(p0 uint16, p1 uint8) int32 {
    v9 := []int{1}
    fmt.Println(p0, p1)
    var v10 float32 = float32((0.0 / -2.25))
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: p1, F1: false, F2: true}
        var v13 string = (("a" + "hello") + "a")
        v10 = (v10 * v10)
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: p1, F1: (847 >= p0), F2: (false && (129 > 463))}
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    _ = v9
    return 0
}

func main() {
    var v1 int8 = 0
    var v2 string = (("a" + ((("" + ((("\n\t" + "hello") + ("a" + ("foo bar" + "hello"))) + "foo bar")) + "foo bar") + "foo bar")) + "")
    var v3 uint32 = 596
    var v4 int8 = -30
    v5 := []string{("foo bar" + ((("" + "\n\t") + (v2 + v2)) + v2)), v2}
    v6 := make(map[int]bool)
    if (!(1 > -32768)) {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint32 = (^(v3))
            v17 := GenS1{F0: 169, F1: false, F2: (!true)}
            v2 = "\n\t"
            _ = fi15
            _ = v16
            _ = v17
        }
        v2 = "a"
    }
    if false {
        v6[0] = ((-1 >= (-(v1))) || ((!(180 > 1)) && false))
        var v18 float64 = 1e10
        v2 = v2
        _ = v18
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 = v2
    _ = len(v5)
    v19 := GenS1{F0: 2, F1: ((!(int8(v3) != 42)) && (732 != 42)), F2: (!(-2.25 >= 0.1))}
    for fi20 := 0; fi20 < 9; fi20++ {
        v2 = "hello"
        if false {
            var v21 int32 = 42
            _ = v21
        } else {
            if (!false) {
                if (-724 == 42) {
                    if false {
                        for fi22 := 0; fi22 < 6; fi22++ {
                            if false {
                                if true {
                                    var v23 uint32 = 0
                                    _ = len(v6)
                                    _ = v23
                                } else {
                                    _ = len(v5)
                                    if false {
                                        fmt.Println(v1, v1, fi20)
                                        var v24 uint32 = v3
                                        var v25 uint64 = 18446744073709551615
                                        _ = v24
                                        _ = v25
                                    }
                                }
                                v2 = ("\n\t" + "foo bar")
                            } else {
                                v26 := GenS1{F0: 2, F1: false, F2: ((255 <= 127) || false)}
                                var v27 uint8 = 2
                                _ = v26
                                _ = v27
                            }
                            if (1 >= 18446744073709551615) {
                                v1 += v4
                            } else {
                                fmt.Println(v2)
                            }
                            _ = fi22
                        }
                        v28 := GenS2{F0: 255, F1: (v1 << ((2) & 63))}
                        _ = v28
                    }
                    v29 := GenS1{F0: 1, F1: true, F2: ((45 == 4294967295) || true)}
                    _ = v29
                }
                if true {
                    v30 := GenS3{F0: 2, F1: 42}
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v3 = 1
                        _ = fi31
                    }
                    _ = v30
                } else {
                    v3 = v3
                    v5 = append(v5, v2)
                }
                _ = len(v5)
            }
            var v32 bool = false
            _ = v32
        }
        _ = fi20
    }
    var v33 int8 = (v4 ^ (^(v4)))
    v4 += int8(v3)
    if true {
        var v34 uint8 = 0
        switch "foo bar" {
        case "hello":
            v34 -= v34
        case "\n\t":
            var v35 uint16 = 801
            _ = v35
        default:
            var v36 uint16 = 32767
            _ = v36
        }
        _ = v34
    }
    fmt.Println(v1, v4)
    for fi37 := 0; fi37 < 11; fi37++ {
        switch (fi37 >> ((1) & 63)) {
        case 629:
            for fi38 := 0; fi38 < 20; fi38++ {
                var v39 bool = false
                if v39 {
                    v6[-2147483648] = (v39 && (!v39))
                    var v40 string = ((("\n\t" + (v2 + v2)) + v2) + v2)
                    if (v3 >= (v3 % ((41) | 1))) {
                        fmt.Println(v3)
                    }
                    _ = v40
                }
                _ = fi38
                _ = v39
            }
        default:
            fmt.Println(v3)
        }
        _ = fi37
    }
    v41 := v6[1]
    switch v2 {
    case "\n\t":
        switch v2 {
        case "a":
            var v42 uint32 = v3
            _ = v42
        default:
            delete(v6, 42)
        }
    case "":
        if v41 {
            v43 := GenS4{F0: 2}
            v4 -= (v33 * v33)
            var v44 string = ("foo bar" + v2)
            _ = v43
            _ = v44
        }
    default:
        var v45 float64 = ((0.1 * 0.1) + 0.1)
        _ = v45
    }
    if len(v5) > 0 {
        v5[0] = ("foo bar" + "")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
