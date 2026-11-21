package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 int32; F2 int64 }
func genFn0(p0 uint8, p1 int64) uint8 {
    v6 := []int{32768, -9223372036854775808}
    if len(v6) > 0 {
        v6[0] = 128
    }
    var v7 int32 = -2147483648
    for fi8 := 0; fi8 < 15; fi8++ {
        v7 += 1
        switch ("a" + "hello") {
        case "\n\t":
            _ = len(v6)
        case "":
            if ((p1 + (-(p1))) < p1) {
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 string = "hello"
                    v6 = append(v6, 4294967295)
                    if true {
                        if true {
                            var v11 string = (v10 + v10)
                            v10 = "foo bar"
                            v12 := GenS1{F0: 4294967295, F1: -1, F2: 42}
                            _ = v11
                            _ = v12
                        } else {
                            switch "" {
                            case "\n\t":
                                switch fi8 {
                                case -594:
                                    if (((false || (false && (42 > 329))) || (992 >= 1)) && ((3.14159 != 0.0) || false)) {
                                        fmt.Println(fi8)
                                    }
                                }
                            case "a":
                                var v13 uint16 = 42
                                _ = v13
                            case "":
                                switch v10 {
                                case "a":
                                    var v14 string = (((("foo bar" + v10) + v10) + ("" + v10)) + ((("foo bar" + v10) + v10) + ((("" + v10) + (v10 + "")) + "hello")))
                                    _ = v14
                                case "\n\t":
                                    p0 = 255
                                default:
                                    var v15 float64 = (((1.5 - 0.1) * 0.1) / 1.0)
                                    _ = v15
                                }
                            }
                        }
                    }
                    _ = fi9
                    _ = v10
                }
            } else {
                if false {
                    var v16 float32 = 0.1
                    _ = v16
                } else {
                    fmt.Println(p0, p0, p0)
                }
                var v17 uint64 = 18446744073709551615
                fmt.Println(v17, p0, p0)
                _ = v17
            }
        case "a":
            switch fi8 {
            case 69:
                if (-2.25 > 1.5) {
                    v18 := GenS1{F0: 2, F1: -355, F2: 0}
                    _ = v18
                }
            case 9223372036854775807:
                var v19 uint64 = 722
                _ = v19
            }
        }
        v20 := GenS2{F0: 0}
        _ = fi8
        _ = v20
    }
    var v21 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v21
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int8, p1 uint, p2 int8) string {
    v22 := []int{940, 1}
    p0 = p0
    fmt.Println(p1, p2, p0)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return "foo bar"
}

func main() {
    v1 := []string{"hello"}
    var v2 int16 = 1000
    var v3 bool = true
    v4 := make(map[string]uint32)
    v5 := []string{"a"}
    var v23 bool = false
    _ = len(v5)
    var v24 bool = v3
    var v25 int = 1
    var v26 uint16 = 0
    if ((0 >= 703) && v3) {
        var v27 int64 = 0
        var v28 int8 = -1
        if (241 == 7) {
            if v3 {
                fmt.Println(v24)
                v27 -= (v27 + int64(v28))
            } else {
                var v29 int64 = 245
                var v30 uint16 = 167
                _ = v29
                _ = v30
            }
        } else {
            v31 := GenS3{F0: 214}
            _ = v31
        }
        _ = v27
        _ = v28
    } else {
        v32 := GenS4{F0: -1}
        var v33 int64 = 0
        _ = v32
        _ = v33
    }
    v2 -= 298
    if false {
        v34 := GenS5{F0: 839}
        delete(v4, ("hello" + ("\n\t" + "\n\t")))
        fmt.Println(v24, v2, v3)
        _ = v34
    } else {
        v35 := v4[("a" + "\n\t")]
        if true {
            var v36 int16 = 1000
            _ = v36
        }
        v35 += 922
        _ = v35
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        var v38 uint = 4294967295
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (100 <= -124) {
            fmt.Println(v38)
            v39 := GenS5{F0: 18446744073709551615}
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    var v40 float32 = (0.1 / (-2.25 + 1e10))
    _ = len(v4)
    var v41 float64 = 1e10
    fmt.Println(v25, v40)
    for fi42 := 0; fi42 < 6; fi42++ {
        fmt.Println(fi42)
        for fi43 := 0; fi43 < 20; fi43++ {
            v44 := GenS1{F0: 255, F1: (v2 - (v2 ^ v2)), F2: 32767}
            v45 := GenS5{F0: 786}
            v41 -= (3.14159 / (v41 + ((0.1 + (((((v41 - float64(3.14159)) / ((v41 + v41) / v41)) + ((v41 / (-2.25 + float64((v40 + (1e10 * v40))))) * float64(v40))) * (((((1.5 / (3.14159 + v41)) / float64((v40 + -2.25))) * -2.25) + (v41 * float64(0.0))) / v41)) * (v41 - v41))) / (v41 - (v41 * 0.1)))))
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = fi42
    }
    if (0 == 711) {
        v25 = (v25 / ((v25) | 1))
        v3 = v23
    }
    v46 := GenS5{F0: 1}
    v47 := GenS6{F0: (v25 >> ((464) & 63)), F1: -325, F2: 1}
    for fi48 := 0; fi48 < 8; fi48++ {
        for fi49 := 0; fi49 < 7; fi49++ {
            fmt.Println(v2, v41)
            _ = fi49
        }
        for fi50 := 0; fi50 < 15; fi50++ {
            var v51 float64 = -2.25
            var v52 bool = ((!(false && false)) && (!(!v24)))
            var v53 uint16 = (v26 & (^(v26)))
            if v24 {
                v54 := GenS4{F0: 42}
                var v55 int64 = 1
                _ = v54
                _ = v55
            } else {
                v3 = ((1 >= 0) && (18446744073709551615 == 18446744073709551615))
                v51 += 1e10
                fmt.Println(fi50, v24)
            }
            v56 := GenS4{F0: 444}
            _ = fi50
            _ = v51
            _ = v52
            _ = v53
            _ = v56
        }
        for fi57 := 0; fi57 < 3; fi57++ {
            v4["hello"] = 63
            fmt.Println(v2, v3, fi48)
            _ = fi57
        }
        _ = fi48
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v46
    _ = v47
    _ = v5
}
