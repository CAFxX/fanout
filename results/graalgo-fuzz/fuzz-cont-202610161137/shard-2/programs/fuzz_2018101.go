package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int; F1 int16; F2 int32 }
type GenS6 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint16; F1 string; F2 bool }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 bool; F1 int16 }
type GenS10 struct { F0 uint; F1 string }
type GenS11 struct { F0 uint; F1 int16 }
type GenS12 struct { F0 uint }
func genFn0(p0 uint) uint32 {
    if false {
        if false {
            for fi7 := 0; fi7 < 4; fi7++ {
                fmt.Println(p0)
                _ = fi7
            }
            var v8 int8 = 112
            p0 -= 2
            _ = v8
        } else {
            var v9 uint64 = 2
            _ = v9
        }
        var v10 int64 = 0
        _ = v10
    } else {
        if true {
            var v11 int64 = 1
            p0 -= (p0 + p0)
            var v12 uint = 541
            _ = v11
            _ = v12
        } else {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(p0)
            default:
                var v13 string = ""
                _ = v13
            }
            v14 := GenS1{F0: 127, F1: 42, F2: 42}
            _ = v14
        }
        var v15 float32 = 1.5
        for fi16 := 0; fi16 < 10; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := []int{657}
                var v19 int8 = 1
                switch "foo bar" {
                case "":
                    if ((^(p0)) == p0) {
                        switch (^(fi16)) {
                        case 910:
                            if (1 > 792) {
                                v20 := GenS2{F0: -2147483648}
                                v15 += v15
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                _ = v20
                            } else {
                                switch ("hello" + "") {
                                case "hello":
                                    p0 += 697
                                case "foo bar":
                                    v18 = append(v18, 403)
                                case "":
                                    if ((1.5 + (0.1 + v15)) == v15) {
                                        v15 += 1e10
                                        switch (^(fi17)) {
                                        case 635:
                                            var v21 uint32 = 4294967295
                                            _ = v21
                                        }
                                    } else {
                                        var v22 int64 = 9223372036854775807
                                        if len(v18) > 0 {
                                            v18[0] = 127
                                        }
                                        _ = v22
                                    }
                                default:
                                    var v23 uint64 = 4294967295
                                    _ = v23
                                }
                                fmt.Println(fi17, fi17)
                                var v24 uint32 = 733
                                _ = v24
                            }
                        case -69:
                            v25 := GenS3{F0: false}
                            _ = v25
                        case 127:
                            v19 = v19
                        default:
                            v26 := GenS4{F0: 4294967295, F1: -1}
                            _ = v26
                        }
                        if true {
                            if len(v18) > 0 {
                                v18[0] = fi16
                            }
                        } else {
                            var v27 float64 = 0.0
                            _ = v27
                        }
                    } else {
                        var v28 float64 = 0.0
                        switch "foo bar" {
                        case "":
                            if (v19 != (v19 << ((p0) & 63))) {
                                v18 = append(v18, (fi16 & (fi16 ^ (-(fi16)))))
                            } else {
                                fmt.Println(v19)
                                var v29 int64 = 9223372036854775807
                                var v30 uint16 = 65535
                                fmt.Println(p0)
                                _ = v29
                                _ = v30
                            }
                        case "a":
                            if (104 > 0) {
                                v15 = 3.14159
                                fmt.Println(v28, p0)
                            }
                        }
                        if ((!false) || (v28 == float64(v15))) {
                            var v31 bool = true
                            _ = v31
                        } else {
                            v32 := GenS5{F0: fi16, F1: -32768, F2: 2147483647}
                            fmt.Println(fi17, fi17)
                            _ = v32
                        }
                        var v33 bool = (!(1e10 < v15))
                        if true {
                            if len(v18) > 0 {
                                v18[0] = (fi17 | (^(fi17)))
                            }
                        } else {
                            switch "a" {
                            case "":
                                v34 := GenS6{F0: 255, F1: 507, F2: 2}
                                _ = v34
                            default:
                                var v35 int64 = 0
                                _ = v35
                            }
                            switch "foo bar" {
                            case "foo bar":
                                v19 = v19
                            case "\n\t":
                                fmt.Println(v19, fi17, v33)
                                var v36 float32 = (((v15 - float32(v28)) + float32(v28)) / float32(v28))
                                var v37 int16 = -32768
                                _ = v36
                                _ = v37
                            }
                        }
                        _ = v28
                        _ = v33
                    }
                default:
                    var v38 bool = true
                    _ = v38
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            v39 := []int{fi16}
            var v40 uint16 = 32767
            _ = fi16
            _ = v39
            _ = v40
        }
        _ = v15
    }
    v41 := GenS7{F0: 32767, F1: "", F2: (811 > -421)}
    v42 := []int{}
    _ = p0
    _ = v41
    _ = v42
    return 42
}
func genFn1(p0 int) uint16 {
    if true {
        var v43 int32 = 2147483647
        var v44 uint = 201
        fmt.Println(v44, v43, p0)
        _ = v43
        _ = v44
    } else {
        var v45 bool = true
        var v46 bool = true
        _ = v45
        _ = v46
    }
    v47 := []int{2147483647, (p0 % ((-1) | 1)), 844, p0}
    if false {
        fmt.Println(p0)
        p0 = p0
    } else {
        v48 := GenS6{F0: 134, F1: 2147483647, F2: 42}
        switch p0 {
        case 32767:
            v49 := GenS8{F0: 42, F1: 492}
            _ = v49
        case -1:
            p0 += p0
        case -371:
            fmt.Println(p0)
        default:
            for fi50 := 0; fi50 < 11; fi50++ {
                switch (p0 * (fi50 ^ -385)) {
                case -9223372036854775808:
                    switch (p0 / ((127) | 1)) {
                    case 375:
                        fmt.Println(fi50, fi50)
                    case -9223372036854775808:
                        p0 -= fi50
                    default:
                        v51 := GenS9{F0: ((!false) && false), F1: 0}
                        _ = v51
                    }
                case -776:
                    for fi52 := 0; fi52 < 13; fi52++ {
                        if (42 >= 2147483647) {
                            v53 := GenS4{F0: 2, F1: -2147483648}
                            var v54 uint32 = 2
                            v54 = genFn0(277)
                            _ = v53
                            _ = v54
                        } else {
                            var v55 string = ((("foo bar" + "a") + (("foo bar" + "foo bar") + ("hello" + ("a" + ((((("" + ("a" + "hello")) + "\n\t") + "\n\t") + "foo bar") + "foo bar"))))) + "foo bar")
                            var v56 int64 = 0
                            if (true || (0 < -1)) {
                                v56 -= (v56 << ((0) & 63))
                                _ = len(v47)
                            }
                            _ = v55
                            _ = v56
                        }
                        switch ("" + ("" + "\n\t")) {
                        case "foo bar":
                            p0 = fi50
                            v57 := GenS10{F0: 195, F1: ("a" + "a")}
                            p0 -= -537
                            _ = v57
                        case "":
                            var v58 uint64 = 312
                            _ = v58
                        case "hello":
                            v59 := GenS11{F0: 18446744073709551615, F1: 0}
                            _ = v59
                        default:
                            v60 := GenS9{F0: true, F1: -629}
                            _ = v60
                        }
                        if len(v47) >= 2 {
                            v47 = v47[1:len(v47)]
                        }
                        _ = fi52
                    }
                case -1:
                    if len(v47) > 0 {
                        v47[0] = (p0 | p0)
                    }
                }
                var v61 int = (p0 + 127)
                _ = fi50
                _ = v61
            }
        }
        _ = v48
    }
    var v62 float64 = 3.14159
    fmt.Println(v62, p0)
    _ = p0
    _ = v47
    _ = v62
    return 931
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 247
    var v3 int16 = -32768
    var v4 uint16 = 65535
    var v5 uint64 = 1
    var v6 int32 = 905
    if true {
        v2 -= 630
        switch "\n\t" {
        case "":
            delete(v1, "a")
        }
    } else {
        fmt.Println(v2, v3)
        var v63 float64 = 3.14159
        v64 := GenS9{F0: ((1.5 * ((3.14159 + float32(v63)) / ((0.1 + 0.0) - -2.25))) > (-2.25 + 3.14159)), F1: v3}
        _ = v63
        _ = v64
    }
    _ = len(v1)
    var v65 uint32 = 1
    v66 := []int{}
    var v67 float32 = float32(0.1)
    var v68 uint8 = 2
    v65 -= genFn0(0)
    v69 := GenS1{F0: v68, F1: (v5 - v5), F2: int64(v5)}
    v70 := GenS12{F0: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
