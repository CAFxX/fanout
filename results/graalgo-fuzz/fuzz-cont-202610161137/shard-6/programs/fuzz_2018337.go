package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int16; F2 int8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 uint64; F2 int32 }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int8; F1 int16; F2 int32 }
type GenS13 struct { F0 uint }
func genFn0() uint8 {
    v6 := []int{-1, 127, -32, 42}
    switch "hello" {
    case "a":
        for fi7 := 0; fi7 < 12; fi7++ {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(fi7)
            case "hello":
                for fi8 := 0; fi8 < 20; fi8++ {
                    switch ("hello" + (("a" + "foo bar") + ("a" + "\n\t"))) {
                    case "":
                        if true {
                            var v9 uint32 = 65535
                            v10 := GenS1{F0: 201}
                            _ = v10
                            _ = v9
                        } else {
                            var v11 bool = (fi8 <= fi8)
                            var v12 uint16 = 497
                            fmt.Println(fi8)
                            _ = v11
                            _ = v12
                        }
                        if false {
                            fmt.Println(fi7)
                            var v13 uint16 = 255
                            var v14 int16 = 32767
                            v15 := GenS2{F0: 128, F1: v14, F2: -84}
                            var v16 string = ("foo bar" + "")
                            _ = v13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    }
                    _ = fi8
                }
            }
            var v17 int = fi7
            fmt.Println(v17)
            _ = fi7
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 18; fi18++ {
            if (1.5 > 1.5) {
                fmt.Println(fi18)
                var v19 bool = ((1 < 42) || (532 != 42))
                switch "\n\t" {
                case "a":
                    v6 = append(v6, (fi18 - 2147483647))
                default:
                    fmt.Println(v19, v19)
                }
                _ = v19
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint8 = 41
                _ = fi20
                _ = v21
            }
            _ = fi18
        }
    }
    var v22 uint16 = 42
    if len(v6) > 0 {
        v6[0] = 42
    }
    if true {
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 bool = false
            if ("a" == ("hello" + "hello")) {
                v25 := GenS3{F0: 1, F1: (fi23 & fi23)}
                _ = v25
            } else {
                var v26 float32 = 0.1
                if v24 {
                    v27 := GenS4{F0: 1, F1: (v22 % (((^(v22))) | 1))}
                    for fi28 := 0; fi28 < 13; fi28++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v29 float32 = (v26 - float32((1.5 - (float64((((v26 - (v26 / v26)) * v26) - float32(((-2.25 - (-2.25 + 0.1)) - 0.0)))) / (0.1 / float64(1.5))))))
                        _ = fi28
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 9; fi30++ {
                        fmt.Println(fi30)
                        if true {
                            v31 := GenS5{F0: 1, F1: fi30}
                            var v32 string = ((("foo bar" + "\n\t") + "foo bar") + "\n\t")
                            v24 = (v24 || true)
                            _ = v31
                            _ = v32
                        }
                        _ = fi30
                    }
                    _ = v27
                }
                _ = v26
            }
            _ = len(v6)
            _ = fi23
            _ = v24
        }
        for fi33 := 0; fi33 < 18; fi33++ {
            var v34 uint = 4294967295
            _ = fi33
            _ = v34
        }
        v6 = append(v6, 10)
    } else {
        var v35 float64 = -2.25
        v36 := GenS3{F0: 1, F1: 42}
        _ = len(v6)
        _ = v35
        _ = v36
    }
    _ = v22
    _ = v6
    return 42
}
func genFn1(p0 int32, p1 uint32, p2 uint8) int8 {
    for fi37 := 0; fi37 < 4; fi37++ {
        switch -45 {
        case 814:
            p2 -= p2
        case 32767:
            v38 := []int{(fi37 | fi37), fi37, fi37}
            _ = v38
        default:
            fmt.Println(p1, p1, fi37)
        }
        var v39 int = (fi37 & (fi37 << ((1) & 63)))
        for fi40 := 0; fi40 < 2; fi40++ {
            fmt.Println(p0, v39)
            var v41 int8 = 88
            v42 := GenS6{F0: 780, F1: "a"}
            v43 := []int{v39, (fi37 | fi37)}
            var v44 int = int(p2)
            if len(v43) > 0 {
                v43[0] = fi40
            }
            v45 := GenS7{F0: "hello", F1: false}
            v43 = append(v43, fi40)
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
            _ = v44
            _ = v45
        }
        _ = fi37
        _ = v39
    }
    for fi46 := 0; fi46 < 9; fi46++ {
        for fi47 := 0; fi47 < 11; fi47++ {
            v48 := GenS5{F0: 18446744073709551615, F1: fi46}
            var v49 int = (fi46 ^ 226)
            p0 = (-(p0))
            v50 := GenS8{F0: 160}
            if true {
                if false {
                    var v51 uint8 = p2
                    p0 -= p0
                    v52 := []int{}
                    if len(v52) >= 2 {
                        v52 = v52[1:len(v52)]
                    }
                    _ = v51
                    _ = v52
                }
                v53 := GenS9{F0: 412, F1: 925, F2: p0}
                _ = v53
            }
            v54 := GenS9{F0: 0, F1: 18446744073709551615, F2: (p0 % (((-(p0))) | 1))}
            _ = fi47
            _ = v48
            _ = v49
            _ = v50
            _ = v54
        }
        v55 := []int{-9223372036854775808, fi46}
        _ = fi46
        _ = v55
    }
    if false {
        p1 = (p1 << ((377) & 63))
        for fi56 := 0; fi56 < 11; fi56++ {
            if false {
                var v57 float64 = ((0.1 * (0.0 + 1.5)) - 1.5)
                var v58 float32 = (((0.1 + 1.5) * (1e10 + -2.25)) - 0.1)
                v57 = 0.0
                _ = v57
                _ = v58
            }
            _ = fi56
        }
    }
    switch "hello" {
    case "":
        for fi59 := 0; fi59 < 10; fi59++ {
            for fi60 := 0; fi60 < 17; fi60++ {
                p2 -= 42
                _ = fi60
            }
            var v61 string = "a"
            for fi62 := 0; fi62 < 18; fi62++ {
                if false {
                    v63 := GenS4{F0: 403, F1: 42}
                    _ = v63
                } else {
                    v64 := []int{fi62}
                    p2 = genFn0()
                    v65 := GenS10{F0: 288, F1: 4294967295}
                    _ = v64
                    _ = v65
                }
                switch (fi62 * 795) {
                case -595:
                    var v66 uint = 65535
                    _ = v66
                case 99:
                    v67 := []int{4294967295, fi62, (fi62 % ((fi59) | 1)), fi59}
                    _ = v67
                case 42:
                    var v68 uint8 = (p2 & (p2 >> ((2) & 63)))
                    _ = v68
                }
                v61 = ((("hello" + v61) + v61) + "")
                _ = fi62
            }
            _ = fi59
            _ = v61
        }
    default:
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint16 = 379
    var v2 string = "\n\t"
    var v3 uint16 = (v1 >> ((42) & 63))
    var v4 uint64 = 998
    v5 := []string{v2, v2, v2, (v2 + v2)}
    switch (v2 + "a") {
    case "\n\t":
        _ = len(v5)
    case "":
        var v69 bool = false
        _ = v69
    case "hello":
        var v70 int8 = 1
        _ = v70
    }
    v71 := GenS2{F0: 2, F1: 352, F2: 0}
    v72 := GenS11{F0: v4}
    v73 := GenS9{F0: 1, F1: 930, F2: 42}
    for fi74 := 0; fi74 < 18; fi74++ {
        var v75 float64 = ((0.1 / (1e10 - float64((float32(1e10) + 1.5)))) + 3.14159)
        var v76 uint64 = v4
        var v77 float64 = v75
        _ = fi74
        _ = v75
        _ = v76
        _ = v77
    }
    switch 42 {
    case 1:
        fmt.Println(v1, v4, v4)
    case 42:
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v5 = append(v5, "foo bar")
            for fi78 := 0; fi78 < 14; fi78++ {
                var v79 bool = (!("a" < "a"))
                _ = fi78
                _ = v79
            }
        } else {
            var v80 int8 = 42
            var v81 float64 = (3.14159 + (-2.25 - (-2.25 * -2.25)))
            v4 -= v4
            _ = v80
            _ = v81
        }
    case 4294967295:
        var v82 uint64 = 960
        _ = v82
    default:
        if false {
            fmt.Println(v2, v3)
            switch (v2 + ("\n\t" + ("" + (((v2 + "a") + v2) + "hello")))) {
            case "a":
                for fi83 := 0; fi83 < 15; fi83++ {
                    v2 = ("\n\t" + (((v2 + (v2 + "\n\t")) + v2) + v2))
                    fmt.Println(v2, v1)
                    _ = fi83
                }
            }
        } else {
            v84 := GenS12{F0: 0, F1: 1, F2: -2147483648}
            _ = v84
        }
    }
    v85 := GenS13{F0: 18446744073709551615}
    if len(v5) > 0 {
        v5[0] = (v2 + "hello")
    }
    var v86 uint = 255
    var v87 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v71
    _ = v72
    _ = v73
    _ = v85
    _ = v86
    _ = v87
}
