package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int16; F2 bool }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 bool; F1 int64; F2 uint }
type GenS8 struct { F0 int16 }
func genFn0(p0 string, p1 int, p2 uint64) uint16 {
    for fi5 := 0; fi5 < 2; fi5++ {
        if false {
            for fi6 := 0; fi6 < 18; fi6++ {
                if (!(!(1e10 == (1.5 - 3.14159)))) {
                    if false {
                        fmt.Println(p2)
                        switch 4294967295 {
                        case 128:
                            p1 = fi5
                        case -698:
                            var v7 uint32 = 537
                            _ = v7
                        }
                        fmt.Println(p2, fi5)
                    }
                    var v8 float32 = (float32(1e10) - float32(-2.25))
                    v9 := []int{fi6, (fi5 ^ fi6)}
                    _ = v8
                    _ = v9
                }
                switch p0 {
                case "a":
                    v10 := GenS1{F0: 127, F1: 213, F2: 269}
                    _ = v10
                }
                fmt.Println(p2)
                _ = fi6
            }
        } else {
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 uint = 2
                v13 := []int{4294967295, p1, (fi5 & fi5), int(v12)}
                fmt.Println(p0)
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        var v14 string = "foo bar"
        fmt.Println(p1, fi5)
        _ = fi5
        _ = v14
    }
    fmt.Println(p2, p1)
    p1 -= (p1 | int(p2))
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 126
}
func genFn1() string {
    var v15 uint16 = 42
    if ((!(1 <= v15)) || true) {
        if ((("" + "a") + ("\n\t" + "")) > "a") {
            fmt.Println(v15)
        } else {
            var v16 float32 = 0.0
            for fi17 := 0; fi17 < 19; fi17++ {
                v15 = (v15 * (v15 + v15))
                _ = fi17
            }
            v16 = v16
            _ = v16
        }
    } else {
        if false {
            v15 = v15
            v18 := GenS2{F0: "foo bar"}
            _ = v18
        } else {
            v15 = (v15 & (v15 / ((genFn0("", 0, 1)) | 1)))
            switch ("a" + "\n\t") {
            case "foo bar":
                var v19 int = -2147483648
                _ = v19
            case "":
                var v20 int32 = -2147483648
                _ = v20
            default:
                var v21 int64 = 1
                _ = v21
            }
            var v22 int32 = -2147483648
            _ = v22
        }
        v23 := GenS2{F0: ((("\n\t" + ("foo bar" + "a")) + ("a" + "a")) + ("hello" + ""))}
        fmt.Println(v15)
        _ = v23
    }
    _ = v15
    return ("\n\t" + "hello")
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{("\n\t" + "\n\t"), "", "hello", ((("" + ("\n\t" + "\n\t")) + "\n\t") + ("foo bar" + "foo bar"))}
    var v3 int16 = 1
    v4 := []int{-1, 4294967295, -2147483648, 98}
    switch 127 {
    case -1:
        if (1e10 >= (3.14159 + -2.25)) {
            v24 := GenS3{F0: 42, F1: (26 < 1), F2: 1}
            _ = v24
        } else {
            fmt.Println(v3)
            var v25 uint64 = 171
            _ = v25
        }
    }
    _ = len(v1)
    v3 -= v3
    var v26 string = (("hello" + ("" + "")) + "hello")
    if (!((697 >= 42) || (2147483647 > -1))) {
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint = 18446744073709551615
            if (3.14159 > 3.14159) {
                _ = len(v2)
                var v29 float32 = 3.14159
                _ = v29
            } else {
                v30 := GenS3{F0: -128, F1: false, F2: 1}
                v31 := GenS4{F0: 2}
                if false {
                    v28 = (v28 | v28)
                    if true {
                        for fi32 := 0; fi32 < 18; fi32++ {
                            if (((fi27 & (-(fi27))) == 32767) || true) {
                                v26 = v26
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            v3 -= 32767
                            _ = fi32
                        }
                    } else {
                        var v33 int64 = -1
                        _ = v33
                    }
                    var v34 uint64 = 1
                    _ = v34
                } else {
                    fmt.Println(v28, v28)
                    var v35 string = (("\n\t" + v26) + ((v26 + v26) + v26))
                    _ = v35
                }
                _ = v30
                _ = v31
            }
            var v36 int = fi27
            _ = fi27
            _ = v28
            _ = v36
        }
    }
    v37 := GenS4{F0: 0}
    v26 = (("foo bar" + ("hello" + v26)) + v26)
    var v38 uint64 = 2
    if true {
        v3 -= 185
        var v39 int8 = 42
        _ = v39
    }
    var v40 float64 = (float64((((0.0 / 1.0) - (1.5 * 0.1)) * (((-2.25 - (0.1 * 0.1)) - 1.5) / 1.0))) + -2.25)
    v3 = (v3 + v3)
    v3 += (v3 ^ v3)
    var v41 float32 = 0.1
    switch v26 {
    case "a":
        if ((float32(v40) - (v41 * v41)) == 0.1) {
            for fi42 := 0; fi42 < 15; fi42++ {
                if false {
                    fmt.Println(v3, v3, v26)
                }
                _ = fi42
            }
            v43 := GenS5{F0: 884, F1: (v3 << ((1) & 63)), F2: (2147483647 == -163)}
            _ = v43
        }
    }
    switch (v26 + v26) {
    case "a":
        v44 := v1[2147483647]
        _ = v44
    case "":
        var v45 float64 = float64(v41)
        _ = v45
    default:
        for fi46 := 0; fi46 < 5; fi46++ {
            for fi47 := 0; fi47 < 7; fi47++ {
                fmt.Println(v3, fi46)
                _ = fi47
            }
            _ = fi46
        }
    }
    if true {
        var v48 bool = (693 <= 0)
        var v49 uint16 = 217
        if v48 {
            v48 = (2147483647 <= 42)
            if v48 {
                for fi50 := 0; fi50 < 14; fi50++ {
                    v51 := GenS6{F0: "a", F1: 158}
                    _ = fi50
                    _ = v51
                }
                fmt.Println(v48)
                var v52 int8 = -1
                _ = v52
            }
        } else {
            for fi53 := 0; fi53 < 8; fi53++ {
                for fi54 := 0; fi54 < 3; fi54++ {
                    if ((-2.25 > (v40 * v40)) && v48) {
                        v49 = v49
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v1[fi54] = -643
                    if (875 != 818) {
                        if v48 {
                            if v48 {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v41 = v41
                                v41 = v41
                            } else {
                                v55 := GenS4{F0: (v38 + 18446744073709551615)}
                                if ((2147483647 == 683) && v48) {
                                    v56 := GenS5{F0: 42, F1: v3, F2: ("\n\t" == v26)}
                                    _ = v56
                                }
                                if v48 {
                                    var v57 float32 = 0.1
                                    _ = v57
                                } else {
                                    v38 = (v38 % ((v38) | 1))
                                    var v58 int16 = v3
                                    _ = v58
                                }
                                v59 := GenS7{F0: (v49 != v49), F1: -656, F2: 42}
                                v40 = (1.5 + v40)
                                if false {
                                    var v60 bool = v48
                                    _ = v60
                                }
                                _ = v55
                                _ = v59
                            }
                        }
                    }
                    _ = fi54
                }
                for fi61 := 0; fi61 < 17; fi61++ {
                    if true {
                        v62 := GenS8{F0: (v3 * 1000)}
                        _ = v62
                    }
                    v3 = (v3 ^ (v3 - v3))
                    _ = fi61
                }
                v63 := GenS2{F0: v26}
                _ = fi53
                _ = v63
            }
            if (v38 > 90) {
                v38 += v38
            } else {
                v26 = v26
                v3 = -1
            }
        }
        _ = v48
        _ = v49
    }
    v64 := GenS2{F0: "foo bar"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v64
}
