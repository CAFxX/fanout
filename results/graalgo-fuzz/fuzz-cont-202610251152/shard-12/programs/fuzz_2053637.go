package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int; F2 bool }
type GenS3 struct { F0 int; F1 int8; F2 int64 }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 uint64; F1 int64 }
type GenS8 struct { F0 uint; F1 string; F2 uint8 }
func genFn0() uint8 {
    var v4 int = 9223372036854775807
    fmt.Println(v4)
    _ = v4
    return 154
}
func genFn1() uint {
    v5 := GenS1{F0: 42}
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 bool = (true && true)
            _ = fi7
            _ = v8
        }
        var v9 uint32 = 2
        v9 -= (v9 << ((255) & 63))
        _ = fi6
        _ = v9
    }
    var v10 uint8 = 128
    v10 = v10
    _ = v10
    _ = v5
    return 1
}

func main() {
    var v1 int64 = 42
    var v2 int64 = (-(v1))
    var v3 float64 = 0.1
    var v11 float64 = (((1e10 / -2.25) - v3) + v3)
    if true {
        v12 := GenS2{F0: genFn0(), F1: 32768, F2: (!false)}
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint16 = 898
            v2 += v2
            v15 := []int{(fi13 + 1), (-(fi13))}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    } else {
        switch (("hello" + ("" + "hello")) + ((("a" + "") + "\n\t") + "hello")) {
        case "a":
            if false {
                var v16 int64 = v2
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 uint = 513
                    v19 := GenS3{F0: fi17, F1: 1, F2: (v16 & v16)}
                    for fi20 := 0; fi20 < 2; fi20++ {
                        var v21 int16 = 0
                        _ = fi20
                        _ = v21
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            }
        case "":
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := []int{(fi22 >> ((733) & 63)), (fi22 - fi22)}
                _ = fi22
                _ = v23
            }
        }
        v24 := []int{127, 645, -2147483648, 4294967295}
        v25 := GenS3{F0: 127, F1: 1, F2: 1}
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 uint8 = 0
        for fi28 := 0; fi28 < 5; fi28++ {
            fmt.Println(v2, fi26, v1)
            var v29 uint64 = 2
            fmt.Println(v2, v27, v2)
            _ = fi28
            _ = v29
        }
        if ((fi26 * (^(fi26))) >= (fi26 >> ((977) & 63))) {
            v11 = -2.25
            v3 = (0.0 / ((1e10 * v3) - (1e10 * v11)))
            fmt.Println(v11, fi26)
        } else {
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 float64 = float64(((0.0 - (0.1 * float32(v3))) / 1.0))
                fmt.Println(fi30, fi26)
                _ = fi30
                _ = v31
            }
            v32 := GenS3{F0: (fi26 ^ (fi26 | 42)), F1: 1, F2: v1}
            _ = v32
        }
        _ = fi26
        _ = v27
    }
    if true {
        fmt.Println(v2, v11, v2)
        v33 := []int{-112, 9223372036854775807}
        if false {
            for fi34 := 0; fi34 < 14; fi34++ {
                fmt.Println(v1)
                if ((32768 < (fi34 << ((1) & 63))) || (!(false || true))) {
                    v35 := GenS4{F0: 2147483648, F1: 42}
                    switch (("a" + "hello") + ((("\n\t" + ("a" + "")) + ("\n\t" + "\n\t")) + "a")) {
                    case "a":
                        v3 += (1.5 + float64(3.14159))
                    case "":
                        v33 = append(v33, (fi34 - (fi34 | (-(fi34)))))
                    case "foo bar":
                        var v36 bool = false
                        _ = v36
                    default:
                        var v37 int32 = 2147483647
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 7; fi38++ {
                        switch ((("\n\t" + (("" + "a") + "")) + ("a" + ("\n\t" + "\n\t"))) + "a") {
                        case "hello":
                            fmt.Println(v1, fi34)
                        case "\n\t":
                            if false {
                                var v39 int = fi38
                                v3 += v11
                                fmt.Println(v39)
                                _ = v39
                            }
                        }
                        v3 -= v3
                        _ = fi38
                    }
                    _ = v35
                }
                _ = fi34
            }
            v40 := GenS1{F0: -125}
            for fi41 := 0; fi41 < 16; fi41++ {
                switch "foo bar" {
                case "\n\t":
                    for fi42 := 0; fi42 < 11; fi42++ {
                        var v43 float32 = 1e10
                        v33 = append(v33, fi42)
                        _ = fi42
                        _ = v43
                    }
                default:
                    var v44 int8 = 1
                    _ = v44
                }
                _ = fi41
            }
            _ = v40
        } else {
            var v45 bool = true
            _ = v45
        }
        _ = v33
    } else {
        v2 = 0
    }
    if true {
        v46 := GenS1{F0: 32}
        v47 := []int{9223372036854775807, -2147483648, 861, -1}
        var v48 float64 = v3
        _ = v46
        _ = v47
        _ = v48
    }
    v2 += 42
    var v49 string = ((("\n\t" + ("foo bar" + "a")) + ("a" + "hello")) + (("\n\t" + "foo bar") + ("hello" + "")))
    var v50 int16 = 0
    v3 -= (0.0 + (1e10 - (float64(0.0) * (v3 + v11))))
    var v51 uint16 = 662
    var v52 uint16 = v51
    fmt.Println(v3)
    v1 = -649
    if false {
        v3 += v3
    } else {
        if true {
            for fi53 := 0; fi53 < 19; fi53++ {
                var v54 string = v49
                v55 := []int{fi53, 4294967295, (^(fi53))}
                _ = fi53
                _ = v54
                _ = v55
            }
            v50 = (v50 + v50)
            v3 = v3
        } else {
            for fi56 := 0; fi56 < 8; fi56++ {
                var v57 bool = (-1 < 53)
                _ = fi56
                _ = v57
            }
            v50 += (-(v50))
            var v58 int8 = 0
            _ = v58
        }
        v49 = v49
        if true {
            var v59 uint32 = 920
            v50 -= v50
            v60 := GenS5{F0: v51}
            _ = v59
            _ = v60
        } else {
            for fi61 := 0; fi61 < 13; fi61++ {
                v62 := GenS6{F0: 0, F1: 2}
                v63 := GenS7{F0: 722, F1: 1}
                if true {
                    v64 := GenS8{F0: 251, F1: "\n\t", F2: 2}
                    _ = v64
                } else {
                    switch v49 {
                    case "a":
                        v65 := GenS6{F0: int8(v52), F1: 42}
                        _ = v65
                    case "hello":
                        var v66 uint16 = v51
                        _ = v66
                    default:
                        v67 := GenS2{F0: 252, F1: fi61, F2: (false || ((v49 < (v49 + (v49 + v49))) || true))}
                        _ = v67
                    }
                    fmt.Println(fi61, v11)
                    var v68 uint8 = 128
                    _ = v68
                }
                _ = fi61
                _ = v62
                _ = v63
            }
            fmt.Println(v50, v49)
            var v69 int64 = v1
            _ = v69
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v51
    _ = v52
}
