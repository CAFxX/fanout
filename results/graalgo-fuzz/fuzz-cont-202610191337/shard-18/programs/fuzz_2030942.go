package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 int; F1 uint8; F2 string }
func genFn0(p0 int, p1 string) int32 {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := []int{fi4}
        fmt.Println(p1)
        var v6 int8 = -128
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int8 = 42
    var v8 uint32 = 312
    for fi9 := 0; fi9 < 17; fi9++ {
        if false {
            fmt.Println(v8, v7)
            v10 := GenS1{F0: p1}
            v11 := []int{}
            _ = v10
            _ = v11
        }
        v8 = v8
        if (100 != (^(v7))) {
            for fi12 := 0; fi12 < 10; fi12++ {
                if (1 == 1) {
                    fmt.Println(v8, v8, p1)
                    switch p0 {
                    case 127:
                        switch p1 {
                        case "hello":
                            switch p1 {
                            case "hello":
                                if (2147483647 <= 862) {
                                    v8 += (v8 & v8)
                                }
                            case "\n\t":
                                v13 := GenS1{F0: ((("a" + p1) + p1) + "hello")}
                                _ = v13
                            case "":
                                if (-676 != 42) {
                                    var v14 int = (p0 + 2147483647)
                                    _ = v14
                                }
                            default:
                                v15 := GenS2{F0: (p1 + "hello")}
                                _ = v15
                            }
                            var v16 int = (fi12 - 9223372036854775807)
                            fmt.Println(v8, v7)
                            _ = v16
                        case "":
                            v17 := GenS2{F0: p1}
                            _ = v17
                        default:
                            if (!false) {
                                switch "\n\t" {
                                case "hello":
                                    if (v7 > v7) {
                                        v18 := GenS3{F0: 42}
                                        var v19 uint16 = 898
                                        fmt.Println(fi12, v8, v8)
                                        _ = v18
                                        _ = v19
                                    }
                                case "\n\t":
                                    switch (fi9 % ((fi9) | 1)) {
                                    case -2147483648:
                                        var v20 uint32 = (v8 % ((v8) | 1))
                                        _ = v20
                                    default:
                                        v21 := GenS4{F0: (p1 + ("a" + "\n\t"))}
                                        _ = v21
                                    }
                                }
                                p1 = "a"
                            } else {
                                var v22 int8 = 127
                                v23 := GenS5{F0: (fi9 >= p0), F1: v7}
                                _ = v22
                                _ = v23
                            }
                        }
                    }
                    v24 := []int{fi12}
                    _ = v24
                } else {
                    p0 += fi12
                    v25 := GenS3{F0: 1}
                    _ = v25
                }
                if false {
                    v26 := []int{p0}
                    _ = v26
                }
                if (((!false) && ((false && false) || (p1 <= (("hello" + p1) + p1)))) && false) {
                    fmt.Println(p1, v7)
                }
                _ = fi12
            }
            if true {
                v27 := GenS6{F0: 2147483647, F1: fi9}
                for fi28 := 0; fi28 < 10; fi28++ {
                    var v29 uint32 = (v8 & v8)
                    _ = fi28
                    _ = v29
                }
                _ = v27
            } else {
                p1 = (p1 + "foo bar")
                v30 := []int{-9223372036854775808, p0}
                v31 := GenS5{F0: true, F1: (v7 + v7)}
                _ = v30
                _ = v31
            }
            v32 := []int{}
            _ = v32
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 820
}

func main() {
    var v1 int64 = 42
    var v2 bool = (2147483648 < 950)
    var v3 int16 = 0
    fmt.Println(v1, v1, v3)
    v33 := []int{2147483647, 9223372036854775807}
    v3 = v3
    v34 := GenS1{F0: "\n\t"}
    v1 += v1
    v35 := GenS4{F0: "a"}
    fmt.Println(v1)
    v1 = (v1 & v1)
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 bool = false
        for fi38 := 0; fi38 < 10; fi38++ {
            if v37 {
                v37 = v37
            } else {
                var v39 float64 = -2.25
                if (!((!v2) && (2147483647 != genFn0(fi36, "\n\t")))) {
                    switch (fi36 * -1) {
                    case 96:
                        if len(v33) > 0 {
                            v33[0] = -436
                        }
                    }
                    if (v2 || ("foo bar" != "")) {
                        if (("hello" + "") <= "hello") {
                            if ((-9223372036854775808 == fi36) && true) {
                                switch "hello" {
                                case "":
                                    var v40 int = (^(fi36))
                                    _ = v40
                                case "hello":
                                    var v41 int8 = 76
                                    _ = v41
                                case "\n\t":
                                    v42 := GenS7{F0: (fi36 | 127), F1: 255, F2: "hello"}
                                    v33 = append(v33, fi36)
                                    var v43 uint16 = 255
                                    _ = v42
                                    _ = v43
                                }
                                fmt.Println(v2, v37, v1)
                                v2 = (!false)
                            } else {
                                _ = len(v33)
                            }
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                        }
                    } else {
                        if len(v33) > 0 {
                            v33[0] = (-(fi38))
                        }
                        var v44 float64 = (v39 - 1.5)
                        _ = v44
                    }
                }
                var v45 uint8 = 21
                _ = v39
                _ = v45
            }
            if (100 != -117) {
                var v46 int64 = (^(v1))
                var v47 string = "\n\t"
                _ = v46
                _ = v47
            } else {
                var v48 int = (fi36 * 127)
                fmt.Println(fi36)
                v33 = append(v33, fi36)
                _ = v48
            }
            v49 := GenS4{F0: ("" + "hello")}
            _ = fi38
            _ = v49
        }
        if false {
            v1 -= v1
            for fi50 := 0; fi50 < 10; fi50++ {
                _ = len(v33)
                v51 := GenS3{F0: 9223372036854775807}
                var v52 float32 = (1.5 + 0.0)
                _ = fi50
                _ = v51
                _ = v52
            }
        }
        _ = fi36
        _ = v37
    }
    if len(v33) > 0 {
        v33[0] = 32768
    }
    var v53 uint32 = 453
    var v54 uint = 4294967295
    v53 -= v53
    fmt.Println(v3)
    if (false || v2) {
        v53 = uint32(v54)
        v54 = v54
    } else {
        if len(v33) > 0 {
            v33[0] = -1
        }
        v53 += v53
        v55 := GenS5{F0: v2, F1: 0}
        _ = v55
    }
    v33 = append(v33, 128)
    v53 = 42
    switch 42 {
    case -523:
        if (-143 != int(v53)) {
            if true {
                v2 = (0 <= 1)
                if (true && (!((v53 / ((v53) | 1)) != (v53 - 0)))) {
                    for fi56 := 0; fi56 < 4; fi56++ {
                        var v57 uint32 = (v53 ^ (v53 * (^(v53))))
                        v58 := GenS2{F0: "foo bar"}
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    if len(v33) > 0 {
                        v33[0] = int(v53)
                    }
                }
                v59 := GenS2{F0: ("a" + ("a" + ((("" + "foo bar") + "\n\t") + ("foo bar" + "foo bar"))))}
                _ = v59
            }
            v2 = ((-(v3)) <= v3)
        } else {
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            var v60 float32 = 0.1
            _ = v60
        }
    case 32768:
        v33 = append(v33, 0)
    default:
        var v61 int16 = (v3 * v3)
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v53
    _ = v54
}
