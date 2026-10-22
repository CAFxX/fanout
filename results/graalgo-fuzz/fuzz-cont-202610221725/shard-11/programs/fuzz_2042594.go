package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint64; F1 string; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 bool }
func genFn0(p0 uint8, p1 int8, p2 int8) string {
    if false {
        p2 += 100
        var v6 int64 = -1
        _ = v6
    } else {
        if true {
            var v7 int8 = (p2 ^ (p2 / (((^(p2))) | 1)))
            var v8 uint64 = 2
            _ = v7
            _ = v8
        }
    }
    p1 = (p2 % ((-1) | 1))
    var v9 float32 = float32(float64(float32(0.1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return ("" + ("" + ("hello" + (("" + (("foo bar" + ("a" + ("\n\t" + ""))) + ((("foo bar" + "\n\t") + ("\n\t" + "foo bar")) + "hello"))) + "\n\t"))))
}
func genFn1(p0 int) int8 {
    v10 := []int{}
    v11 := GenS1{F0: 0, F1: 0, F2: (!(4294967295 < 4294967295))}
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS2{F0: 65535, F1: 0, F2: 704}
        p0 = p0
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 uint = 65535
        fmt.Println(p0, p0, v15)
        if (33 >= 4294967295) {
            for fi16 := 0; fi16 < 12; fi16++ {
                p0 -= (^(fi16))
                _ = fi16
            }
        }
        _ = fi14
        _ = v15
    }
    v17 := GenS3{F0: 117, F1: -98}
    _ = p0
    _ = v10
    _ = v11
    _ = v17
    return 42
}

func main() {
    v1 := []string{"foo bar", "hello", (("hello" + (("" + ("a" + "hello")) + "a")) + ""), "\n\t"}
    v2 := []int{-1}
    v3 := []uint32{2, 2}
    v4 := make(map[int]bool)
    var v5 string = "\n\t"
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 int16 = 32767
        v5 = "a"
        v19 = 1000
        _ = fi18
        _ = v19
    }
    switch "a" {
    case "hello":
        switch 183 {
        case 789:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 636:
            v4[-9223372036854775808] = true
        case -321:
            switch 127 {
            case 926:
                v20 := GenS2{F0: 683, F1: 2, F2: 454}
                _ = v20
            case 4294967295:
                for fi21 := 0; fi21 < 5; fi21++ {
                    v5 = v5
                    v22 := v4[fi21]
                    v22 = false
                    _ = fi21
                    _ = v22
                }
            case 127:
                switch 32767 {
                case 14:
                    v23 := GenS2{F0: 537, F1: 0, F2: 403}
                    _ = v23
                }
            }
        default:
            v5 = v5
        }
    case "foo bar":
        switch 1 {
        case -295:
            var v24 bool = false
            _ = v24
        case 157:
            var v25 float64 = ((0.0 - 0.0) / 1.0)
            _ = v25
        }
    case "a":
        v26 := GenS4{F0: 1, F1: 100}
        _ = v26
    }
    switch "\n\t" {
    case "a":
        switch "a" {
        case "\n\t":
            var v27 int8 = 5
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 int16 = 1
        var v30 float64 = 0.1
        if ((-106 > 127) && (573 > 0)) {
            var v31 int32 = -32
            _ = v31
        } else {
            var v32 bool = (65535 == 255)
            for fi33 := 0; fi33 < 18; fi33++ {
                fmt.Println(fi33)
                v30 = (1e10 * 1.5)
                if (((104 > 42) || (2 != 1)) || (v30 == v30)) {
                    v34 := v4[fi33]
                    if (42 >= genFn1(fi33)) {
                        v32 = v32
                        v5 = (v5 + "")
                    } else {
                        var v35 uint64 = 42
                        _ = v35
                    }
                    if v32 {
                        v36 := GenS1{F0: 255, F1: 112, F2: (!((v5 + v5) < v5))}
                        if false {
                            var v37 uint64 = 1
                            var v38 float64 = v30
                            var v39 bool = v32
                            _ = v37
                            _ = v38
                            _ = v39
                        } else {
                            switch v5 {
                            case "\n\t":
                                if v32 {
                                    if (42 < 1) {
                                        var v40 int = (fi28 ^ -2147483648)
                                        v34 = (true || v34)
                                        _ = v40
                                    }
                                    v4[(-(fi33))] = ((v34 && v34) && (2 <= 1))
                                }
                            }
                        }
                        var v41 float64 = v30
                        _ = v36
                        _ = v41
                    }
                    _ = v34
                }
                var v42 int8 = -127
                _ = fi33
                _ = v42
            }
            _ = v32
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    var v43 int64 = 9223372036854775807
    for fi44 := 0; fi44 < 9; fi44++ {
        v43 += 9223372036854775807
        if (true && (v43 < (-(v43)))) {
            v45 := GenS5{F0: 2, F1: ("a" + v5), F2: 4294967295}
            _ = v45
        }
        if (32767 > 65535) {
            for fi46 := 0; fi46 < 18; fi46++ {
                if (!false) {
                    if (true || (-2.25 != (0.1 / (0.1 + (float64(((3.14159 / ((float32(1e10) - ((-2.25 - float32(float64(1.5))) / 1.0)) / 1e10)) * 3.14159)) + (-2.25 + 3.14159)))))) {
                        v47 := GenS2{F0: 296, F1: 612, F2: 4294967295}
                        if true {
                            switch ((v5 + (v5 + ((((((("a" + (v5 + (((((v5 + v5) + "\n\t") + (v5 + "foo bar")) + (v5 + (("a" + (((("hello" + "") + ("a" + (v5 + ("a" + (v5 + (("hello" + "a") + v5)))))) + "a") + "")) + (v5 + (v5 + v5))))) + v5))) + v5) + "hello") + (v5 + v5)) + ("\n\t" + "\n\t")) + v5) + v5))) + v5) {
                            case "\n\t":
                                v4[32767] = false
                            }
                            v48 := GenS3{F0: 244, F1: -2147483648}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v48
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 559
                            }
                            v5 = "foo bar"
                            var v49 uint32 = 42
                            _ = v49
                        }
                        switch fi44 {
                        case -9223372036854775808:
                            delete(v4, fi44)
                        case 9223372036854775807:
                            var v50 float64 = 3.14159
                            _ = v50
                        case 1:
                            if (false || (!((false && (-196 < -303)) || (4294967295 <= 1)))) {
                                var v51 uint16 = 889
                                _ = v51
                            }
                        default:
                            if true {
                                v5 = v5
                                var v52 int16 = -570
                                _ = v52
                            }
                            v43 = -825
                        }
                        _ = v47
                    }
                    v53 := GenS6{F0: 2147483647}
                    _ = v53
                } else {
                    v3 = append(v3, 2147483648)
                }
                v43 = v43
                _ = fi46
            }
            v43 = 9223372036854775807
        } else {
            v43 = -1
            if (0 <= -2147483648) {
                v5 = ("\n\t" + v5)
            } else {
                fmt.Println(v5, v5)
            }
        }
        _ = fi44
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v54 string = ("foo bar" + v5)
    if true {
        if true {
            var v55 uint32 = 0
            for fi56 := 0; fi56 < 11; fi56++ {
                v55 -= (v55 - 445)
                v57 := GenS5{F0: 453, F1: v5, F2: fi56}
                fmt.Println(v54, fi56, v55)
                _ = fi56
                _ = v57
            }
            _ = v55
        } else {
            v58 := GenS7{F0: -670, F1: true}
            _ = v58
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v54
}
