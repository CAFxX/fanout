package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int; F1 int64; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint, p1 uint8, p2 uint32) uint {
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            if ((p0 << ((uint(p1)) & 63)) >= p0) {
                var v8 float32 = 0.1
                if false {
                    if false {
                        var v9 int = 900
                        var v10 int = fi6
                        if true {
                            v11 := []int{v10, 0, (fi6 << ((p0) & 63)), fi6}
                            fmt.Println(v10, p0)
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            _ = v11
                        } else {
                            p0 -= p0
                        }
                        _ = v10
                        _ = v9
                    }
                    p0 -= p0
                } else {
                    var v12 string = ("" + "")
                    v13 := GenS1{F0: ((-323 >= 42) || false), F1: 42}
                    _ = v12
                    _ = v13
                }
                v14 := GenS1{F0: (100 > 127), F1: 445}
                _ = v14
                _ = v8
            }
            v15 := GenS2{F0: 42, F1: ("\n\t" + "a")}
            var v16 bool = true
            v17 := []int{(-(fi7)), -1, 127}
            fmt.Println(p1, fi7, fi6)
            _ = fi7
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := []int{(^(fi6))}
        _ = fi6
        _ = v18
    }
    var v19 float64 = (0.0 / float64(1.5))
    p1 = (p1 ^ (p1 << ((0) & 63)))
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return p0
}
func genFn1(p0 int64, p1 int8) bool {
    for fi20 := 0; fi20 < 8; fi20++ {
        if false {
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := GenS1{F0: (true || false), F1: 0}
                v23 := []int{fi20, fi21}
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 float32 = 1e10
            _ = v24
        }
        _ = fi20
    }
    v25 := GenS1{F0: true, F1: 0}
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS1{F0: ((true && (324 == -32768)) && ("foo bar" > ("foo bar" + ""))), F1: 0}
        _ = fi26
        _ = v27
    }
    switch 581 {
    case -9223372036854775808:
        if true {
            if false {
                var v28 float32 = (-2.25 / 1.5)
                v29 := GenS2{F0: p0, F1: "\n\t"}
                _ = v28
                _ = v29
            }
            if (1 > 463) {
                v30 := GenS2{F0: 166, F1: ("" + "\n\t")}
                _ = v30
            }
            var v31 uint32 = 50
            _ = v31
        }
    case 4294967295:
        p1 = -95
    case 32768:
        var v32 int32 = 1
        _ = v32
    default:
        var v33 float32 = 0.0
        _ = v33
    }
    v34 := []int{42, 32768, -1, 32768}
    _ = p0
    _ = p1
    _ = v25
    _ = v34
    return false
}

func main() {
    var v1 uint64 = 5
    var v2 float32 = (1.5 / 1e10)
    v3 := []uint32{}
    var v4 int = 1
    v5 := make(map[int]int)
    if true {
        v2 = v2
        if true {
            var v35 bool = (!false)
            _ = v35
        } else {
            for fi36 := 0; fi36 < 5; fi36++ {
                v37 := GenS1{F0: false, F1: (v1 | (^(v1)))}
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 int8 = 127
                    _ = fi38
                    _ = v39
                }
                _ = fi36
                _ = v37
            }
            var v40 float64 = 1e10
            _ = v40
        }
        for fi41 := 0; fi41 < 2; fi41++ {
            var v42 int64 = 1
            v43 := GenS2{F0: -550, F1: ""}
            _ = fi41
            _ = v42
            _ = v43
        }
    } else {
        var v44 int32 = 42
        switch (v4 ^ 2147483647) {
        case 32768:
            if false {
                v45 := v5[2147483647]
                _ = v45
            } else {
                for fi46 := 0; fi46 < 4; fi46++ {
                    for fi47 := 0; fi47 < 6; fi47++ {
                        v1 -= (v1 % ((1) | 1))
                        switch ("a" + ("foo bar" + "a")) {
                        case "a":
                            fmt.Println(fi47, fi46, v44)
                            v48 := GenS2{F0: -1, F1: ""}
                            _ = v48
                        case "foo bar":
                            if (4294967295 <= 1) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                if true {
                                    var v49 string = (("a" + "") + ("" + "\n\t"))
                                    _ = v49
                                } else {
                                    var v50 int8 = -54
                                    _ = v50
                                }
                            }
                        case "":
                            v2 -= (v2 / v2)
                        default:
                            v4 -= (^(v4))
                        }
                        var v51 float32 = v2
                        _ = fi47
                        _ = v51
                    }
                    _ = fi46
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v1, v1, v2)
            }
        case -1:
            switch v4 {
            case 9223372036854775807:
                for fi52 := 0; fi52 < 8; fi52++ {
                    if true {
                        var v53 int32 = (v44 % ((-578) | 1))
                        v44 = (-(v44))
                        for fi54 := 0; fi54 < 2; fi54++ {
                            v4 = 9223372036854775807
                            v1 = 1
                            fmt.Println(v2)
                            _ = fi54
                        }
                        _ = v53
                    }
                    _ = fi52
                }
            case 1:
                var v55 bool = false
                _ = v55
            default:
                if len(v3) > 0 {
                    v3[0] = 42
                }
            }
        default:
            v56 := GenS3{F0: 1, F1: 1, F2: int(v1)}
            _ = v56
        }
        var v57 uint = 255
        _ = v44
        _ = v57
    }
    v3 = append(v3, 2147483647)
    v58 := GenS2{F0: 1, F1: ("hello" + "")}
    v59 := GenS4{F0: (!false)}
    v60 := GenS5{F0: 2147483647}
    v61 := v5[v4]
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    delete(v5, v61)
    v62 := GenS6{F0: -128, F1: 81}
    v61 += v4
    var v63 uint32 = 42
    v63 = v63
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
