package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 uint64; F1 bool; F2 string }
func genFn0(p0 uint) uint16 {
    var v6 uint16 = 65535
    v7 := GenS1{F0: "\n\t", F1: 127}
    v6 -= v6
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: "hello", F1: -9223372036854775808}
        var v10 uint16 = v6
        if false {
            p0 -= 4294967295
            p0 = p0
            var v11 uint64 = 0
            _ = v11
        } else {
            v12 := GenS2{F0: (fi8 % (((fi8 & (fi8 << (((p0 / ((18446744073709551615) | 1))) & 63)))) | 1))}
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 uint64, p1 int64, p2 string) uint32 {
    if true {
        p1 = -9223372036854775808
    }
    if false {
        for fi13 := 0; fi13 < 17; fi13++ {
            v14 := []int{}
            if (!false) {
                var v15 float64 = 3.14159
                _ = v15
            }
            _ = fi13
            _ = v14
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            switch fi16 {
            case 32767:
                for fi17 := 0; fi17 < 16; fi17++ {
                    v18 := GenS3{F0: 9223372036854775807}
                    _ = fi17
                    _ = v18
                }
            case 1:
                fmt.Println(fi16, p2, p1)
            case 4294967295:
                fmt.Println(fi16, fi16)
            }
            p0 -= (^(p0))
            _ = fi16
        }
        p0 += (p0 << ((2) & 63))
    } else {
        if (p1 == (p1 + (-(p1)))) {
            if false {
                var v19 int64 = p1
                fmt.Println(v19, p0)
                for fi20 := 0; fi20 < 10; fi20++ {
                    p2 = "hello"
                    _ = fi20
                }
                _ = v19
            }
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                if (false || false) {
                    fmt.Println(fi21, p2, p0)
                    v22 := GenS4{F0: true}
                    p2 = p2
                    _ = v22
                }
                v23 := GenS5{F0: 0}
                _ = fi21
                _ = v23
            }
            if (false && ((!false) || (-1 > 1))) {
                p1 += -921
                p1 += p1
            }
            if (!(1 <= -484)) {
                if ((382 == 0) || (p0 > 45)) {
                    p2 = (("foo bar" + "a") + p2)
                    for fi24 := 0; fi24 < 17; fi24++ {
                        var v25 uint64 = p0
                        v25 -= 2
                        _ = fi24
                        _ = v25
                    }
                } else {
                    for fi26 := 0; fi26 < 11; fi26++ {
                        switch "a" {
                        case "foo bar":
                            p1 += p1
                        default:
                            v27 := GenS4{F0: false}
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 7; fi28++ {
                            var v29 uint = 0
                            fmt.Println(fi28, fi26)
                            var v30 uint32 = 1
                            v31 := []int{(fi26 & 4294967295), (fi26 * -1)}
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        v32 := GenS5{F0: genFn0(65535)}
                        _ = fi26
                        _ = v32
                    }
                    var v33 int = 824
                    _ = v33
                }
            }
        }
    }
    var v34 int32 = -1
    v34 = (v34 >> ((231) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    return 2147483648
}

func main() {
    var v1 uint = 839
    var v2 float64 = 0.0
    var v3 string = (((("a" + ("hello" + "hello")) + ("a" + ("a" + "a"))) + "a") + "\n\t")
    v4 := make(map[string]int)
    var v5 uint32 = 363
    for fi35 := 0; fi35 < 17; fi35++ {
        var v36 int = 692
        _ = fi35
        _ = v36
    }
    for fi37 := 0; fi37 < 2; fi37++ {
        fmt.Println(fi37)
        v38 := []int{128, fi37, (^(fi37)), (fi37 | (fi37 * (fi37 % ((fi37) | 1))))}
        v1 += 42
        _ = fi37
        _ = v38
    }
    var v39 int8 = 127
    fmt.Println(v5, v5, v5)
    v40 := v4[v3]
    var v41 int8 = (v39 - (v39 / (((-(v39))) | 1)))
    fmt.Println(v39, v41)
    if false {
        var v42 uint8 = 0
        _ = v42
    }
    v3 = "\n\t"
    if true {
        var v43 string = "\n\t"
        v4[(("" + (v43 + ((v43 + (v3 + (v3 + v3))) + v43))) + v43)] = v40
        _ = v43
    } else {
        v44 := GenS6{F0: v40, F1: -1}
        v2 += 0.0
        _ = v44
    }
    var v45 uint32 = genFn1(356, 0, v3)
    if true {
        if ((0.1 == ((0.0 + v2) - ((v2 * v2) + 0.0))) || (true || false)) {
            var v46 float32 = (3.14159 / (1e10 + (1.5 - 0.0)))
            for fi47 := 0; fi47 < 18; fi47++ {
                switch (v40 % ((2147483647) | 1)) {
                case -925:
                    v48 := []int{fi47, 458, 0}
                    _ = v48
                }
                switch ("hello" + "foo bar") {
                case "foo bar":
                    switch (v40 ^ v40) {
                    case 2147483647:
                        v45 = 2147483648
                    default:
                        if false {
                            for fi49 := 0; fi49 < 12; fi49++ {
                                var v50 int8 = v39
                                var v51 int = fi47
                                v52 := GenS6{F0: fi49, F1: 42}
                                var v53 int8 = -106
                                _ = fi49
                                _ = v50
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                        }
                    }
                case "\n\t":
                    fmt.Println(v46, v45, v41)
                case "a":
                    fmt.Println(v3, v45, v39)
                default:
                    delete(v4, (v3 + ("" + "")))
                }
                v39 += v39
                _ = fi47
            }
            v54 := []int{(v40 ^ (v40 >> ((v1) & 63))), (v40 ^ v40), (v40 * (v40 / (((v40 * int(v1))) | 1)))}
            _ = v46
            _ = v54
        } else {
            if true {
                switch 638 {
                case 4294967295:
                    switch "\n\t" {
                    case "\n\t":
                        var v55 bool = (3.14159 >= (3.14159 - v2))
                        _ = v55
                    case "a":
                        for fi56 := 0; fi56 < 3; fi56++ {
                            var v57 string = v3
                            _ = fi56
                            _ = v57
                        }
                    }
                default:
                    switch v40 {
                    case -465:
                        if (4294967295 == 2) {
                            for fi58 := 0; fi58 < 7; fi58++ {
                                if ((!(1 == 1)) && (((true || true) || true) || true)) {
                                    delete(v4, v3)
                                    var v59 uint8 = 2
                                    switch (v3 + v3) {
                                    case "a":
                                        v1 = (v1 << (((v1 * 255)) & 63))
                                    }
                                    _ = v59
                                } else {
                                    v60 := v4[("\n\t" + ("\n\t" + "hello"))]
                                    v61 := []int{-1, 4294967295, (fi58 << ((v1) & 63)), -1}
                                    v4[(v3 + ("\n\t" + v3))] = -1
                                    _ = v60
                                    _ = v61
                                }
                                var v62 int64 = 1
                                v63 := GenS7{F0: 509, F1: ((v3 + "") < v3), F2: v3}
                                _ = fi58
                                _ = v62
                                _ = v63
                            }
                            var v64 int32 = 42
                            _ = v64
                        } else {
                            v5 = v5
                            var v65 uint8 = 42
                            _ = v65
                        }
                    case 4294967295:
                        for fi66 := 0; fi66 < 15; fi66++ {
                            v41 = v39
                            _ = fi66
                        }
                    }
                }
                v41 -= v41
                var v67 int32 = 594
                _ = v67
            }
            var v68 uint32 = 2147483648
            _ = v68
        }
    } else {
        var v69 uint = v1
        v4[("a" + ((v3 + v3) + v3))] = v40
        switch "hello" {
        case "a":
            fmt.Println(v45, v69)
        }
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v5
}
