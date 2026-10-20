package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint; F2 int8 }
type GenS6 struct { F0 uint; F1 int; F2 int }
type GenS7 struct { F0 int64; F1 int; F2 int16 }
type GenS8 struct { F0 int32; F1 bool }
type GenS9 struct { F0 int64; F1 int16; F2 uint64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 int64; F2 int32 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint64; F1 int16; F2 uint64 }
func genFn0() string {
    var v4 float32 = (3.14159 * (-2.25 * (0.0 / ((0.0 + 0.1) - float32(0.0)))))
    v4 -= 0.0
    var v5 float32 = (v4 * v4)
    if true {
        var v6 string = ""
        v7 := []int{}
        var v8 uint64 = 2
        _ = v6
        _ = v7
        _ = v8
    }
    _ = v4
    _ = v5
    return "foo bar"
}
func genFn1(p0 uint16, p1 uint8, p2 uint) uint16 {
    var v9 uint = p2
    fmt.Println(v9, p0, v9)
    if true {
        var v10 int16 = -461
        v11 := GenS1{F0: -9223372036854775808, F1: 0}
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        if false {
            var v13 bool = false
            _ = v13
        } else {
            v14 := []int{(fi12 & fi12), fi12, (fi12 >> (((v9 | 77)) & 63))}
            if (v9 <= 42) {
                switch ("\n\t" + "a") {
                case "a":
                    v15 := GenS2{F0: uint8(p2)}
                    _ = v15
                case "hello":
                    if (565 >= 905) {
                        _ = len(v14)
                        switch (("foo bar" + "\n\t") + "a") {
                        case "foo bar":
                            fmt.Println(v9, p1)
                        case "a":
                            v16 := GenS3{F0: 127}
                            _ = v16
                        }
                    } else {
                        p0 = 65535
                        p1 = p1
                    }
                }
            }
            if (int32(p0) == 1) {
                p1 += 162
            } else {
                var v17 int16 = 1000
                _ = v17
            }
            _ = v14
        }
        _ = fi12
    }
    v18 := GenS4{F0: (("" + ("hello" + "foo bar")) == "\n\t")}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v9
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]int)
    v19 := []int{-463, -9223372036854775808}
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 uint = 4294967295
        v2 = 0
        v2 = v2
        _ = fi20
        _ = v21
    }
    v2 -= v2
    var v22 int16 = -64
    if true {
        v3[42] = 2147483647
        if false {
            fmt.Println(v2, v22)
        }
    }
    if (!true) {
        v22 = v22
        fmt.Println(v22, v2)
        v23 := GenS5{F0: 1, F1: 42, F2: 27}
        _ = v23
    }
    if false {
        for fi24 := 0; fi24 < 18; fi24++ {
            for fi25 := 0; fi25 < 2; fi25++ {
                if (true && (780 > 4294967295)) {
                    v19 = append(v19, 2147483647)
                    _ = len(v19)
                    if (float64(0.0) != float64(3.14159)) {
                        fmt.Println(fi24, fi25)
                        if false {
                            var v26 uint32 = 2
                            v2 -= (^(v2))
                            v27 := GenS6{F0: 255, F1: (fi24 * fi24), F2: (fi24 * fi25)}
                            _ = v26
                            _ = v27
                        }
                        v2 = (v2 & (-(v2)))
                        var v28 uint64 = 1
                        _ = v28
                    } else {
                        v2 = -31
                    }
                }
                var v29 int = int(v22)
                v30 := GenS7{F0: v2, F1: fi24, F2: 32767}
                _ = fi25
                _ = v29
                _ = v30
            }
            var v31 float32 = 0.0
            for fi32 := 0; fi32 < 12; fi32++ {
                v33 := GenS8{F0: 0, F1: false}
                v2 = v2
                v2 += 1
                _ = fi32
                _ = v33
            }
            _ = fi24
            _ = v31
        }
    }
    var v34 uint16 = 0
    if (true && (true && true)) {
        v34 = v34
        if (835 == -2147483648) {
            v22 = 42
            v35 := GenS9{F0: (v2 << ((2) & 63)), F1: 42, F2: 910}
            var v36 string = ("" + "foo bar")
            _ = v35
            _ = v36
        } else {
            for fi37 := 0; fi37 < 20; fi37++ {
                v38 := GenS10{F0: fi37}
                var v39 uint = 0
                _ = fi37
                _ = v38
                _ = v39
            }
            if true {
                fmt.Println(v22)
            }
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        }
    } else {
        v40 := v1[42]
        fmt.Println(v22)
        if (!false) {
            if false {
                v41 := GenS3{F0: -128}
                _ = v41
            } else {
                v34 -= (v34 | 32767)
                delete(v1, 2147483648)
                v42 := GenS7{F0: -480, F1: -1, F2: v22}
                _ = v42
            }
            v34 = genFn1((^(v34)), 0, 114)
        } else {
            fmt.Println(v2, v34)
            if (false || (false || (((1e10 / 0.1) + (float64(-2.25) + -2.25)) < 0.0))) {
                var v43 uint32 = 4294967295
                _ = v43
            }
            if (0.1 == (0.0 / 3.14159)) {
                for fi44 := 0; fi44 < 2; fi44++ {
                    fmt.Println(v34, v40)
                    _ = fi44
                }
            } else {
                v45 := v1[969]
                if (v40 > (v45 + "")) {
                    fmt.Println(v2, v2)
                    fmt.Println(v45, v45, v22)
                } else {
                    switch ((v40 + ((("a" + "hello") + v40) + "hello")) + "hello") {
                    case "\n\t":
                        fmt.Println(v22, v22, v22)
                    case "a":
                        v46 := GenS11{F0: true, F1: (v2 - (v2 | (-(v2)))), F2: 2147483647}
                        _ = v46
                    case "":
                        v1[42] = "a"
                    default:
                        for fi47 := 0; fi47 < 16; fi47++ {
                            var v48 int32 = -1
                            for fi49 := 0; fi49 < 7; fi49++ {
                                var v50 int8 = -128
                                v51 := GenS12{F0: (fi49 * (fi47 / ((fi47) | 1)))}
                                var v52 uint = 65535
                                _ = fi49
                                _ = v50
                                _ = v51
                                _ = v52
                            }
                            _ = fi47
                            _ = v48
                        }
                    }
                }
                if false {
                    v40 = "a"
                    v53 := GenS13{F0: false}
                    _ = v53
                } else {
                    if true {
                        v54 := GenS1{F0: -9223372036854775808, F1: 784}
                        v2 += 0
                        _ = v54
                    }
                    v45 = v45
                }
                _ = v45
            }
        }
        _ = v40
    }
    _ = len(v1)
    _ = len(v3)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    if false {
        fmt.Println(v22, v22, v22)
    }
    var v55 uint16 = 1
    switch "" {
    case "foo bar":
        fmt.Println(v2)
    case "":
        v55 = v34
    case "a":
        if true {
            v55 = v55
            switch "\n\t" {
            case "foo bar":
                if (true && false) {
                    fmt.Println(v22, v34)
                } else {
                    if (((495 < 2) && (!(false && true))) || true) {
                        var v56 float64 = -2.25
                        _ = v56
                    }
                }
            case "":
                if ((true && (656 >= 1)) || true) {
                    v34 = (v34 | (^(v55)))
                } else {
                    v55 = v34
                    if (genFn1((v34 % ((v55) | 1)), 2, 42) <= 466) {
                        fmt.Println(v22)
                    }
                }
            default:
                var v57 string = "hello"
                _ = v57
            }
            v58 := GenS14{F0: 233, F1: (v22 + 42), F2: 0}
            _ = v58
        } else {
            if false {
                fmt.Println(v55, v22, v2)
                v22 = 434
                v22 += 1000
            } else {
                if ("a" == ("" + (("a" + "") + "\n\t"))) {
                    switch ("a" + "foo bar") {
                    case "a":
                        delete(v3, -516)
                    case "hello":
                        var v59 float32 = ((-2.25 - (1.5 - 0.1)) / 1.0)
                        _ = v59
                    }
                    var v60 int8 = 0
                    if true {
                        var v61 uint64 = 122
                        v22 += v22
                        _ = v61
                    }
                    _ = v60
                } else {
                    fmt.Println(v55, v2)
                }
                if true {
                    for fi62 := 0; fi62 < 6; fi62++ {
                        for fi63 := 0; fi63 < 2; fi63++ {
                            switch (("hello" + "a") + "a") {
                            case "foo bar":
                                fmt.Println(fi63)
                            case "a":
                                var v64 int8 = 0
                                fmt.Println(v64, v22)
                                v65 := v1[109]
                                v65 = v65
                                _ = v64
                                _ = v65
                            }
                            if (((1.5 > -2.25) && (false && false)) || true) {
                                switch (fi63 & fi63) {
                                case 4294967295:
                                    v55 -= 255
                                case 127:
                                    var v66 bool = true
                                    v67 := GenS13{F0: v66}
                                    v22 -= 1
                                    _ = v66
                                    _ = v67
                                case 170:
                                    fmt.Println(v55, fi63)
                                default:
                                    fmt.Println(fi63, v22, v2)
                                }
                            } else {
                                var v68 string = ("" + "")
                                _ = v68
                            }
                            var v69 int64 = -12
                            v55 += 2
                            var v70 int16 = 1000
                            _ = fi63
                            _ = v69
                            _ = v70
                        }
                        var v71 bool = true
                        switch fi62 {
                        case -1:
                            var v72 int32 = 756
                            _ = v72
                        case -98:
                            if true {
                                v2 -= 42
                                switch (("foo bar" + ("foo bar" + ((("foo bar" + "foo bar") + "hello") + ("\n\t" + "hello")))) + "") {
                                case "\n\t":
                                    v71 = (false || v71)
                                }
                            } else {
                                v73 := GenS9{F0: -111, F1: v22, F2: 4294967295}
                                v22 += 326
                                for fi74 := 0; fi74 < 14; fi74++ {
                                    var v75 string = "\n\t"
                                    v76 := GenS12{F0: (-(fi62))}
                                    if (-1 == -2147483648) {
                                        v77 := GenS2{F0: 127}
                                        fmt.Println(v71, v34, v71)
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                        _ = v77
                                    }
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                }
                                _ = v73
                            }
                        }
                        _ = fi62
                        _ = v71
                    }
                } else {
                    if true {
                        var v78 uint8 = 137
                        _ = v78
                    }
                    fmt.Println(v22, v55, v2)
                }
            }
        }
    default:
        fmt.Println(v34, v22, v55)
    }
    for fi79 := 0; fi79 < 10; fi79++ {
        _ = len(v3)
        switch 1 {
        case 9223372036854775807:
            var v80 uint32 = 841
            _ = v80
        case 2147483647:
            v19 = append(v19, fi79)
        }
        var v81 float64 = (-2.25 / ((((((0.0 / (3.14159 / 1.5)) * 1e10) * 0.0) - float64((0.1 / 1.5))) / (-2.25 * 1e10)) + float64(1e10)))
        _ = fi79
        _ = v81
    }
    fmt.Println(v2, v22)
    v55 = (v34 + 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v55
}
