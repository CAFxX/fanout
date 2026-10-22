package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 string; F2 uint }
type GenS5 struct { F0 uint64 }
func genFn0() uint16 {
    var v4 uint = 255
    v5 := GenS1{F0: 2, F1: 127, F2: 0}
    v6 := []int{}
    for fi7 := 0; fi7 < 3; fi7++ {
        if (32767 > 32767) {
            var v8 float64 = 3.14159
            _ = v8
        }
        _ = fi7
    }
    _ = v4
    _ = v5
    _ = v6
    return 374
}
func genFn1(p0 bool, p1 string) int64 {
    if false {
        var v9 int8 = -1
        fmt.Println(p0)
        p1 = "foo bar"
        _ = v9
    } else {
        var v10 int16 = 1000
        var v11 int32 = 0
        fmt.Println(v10, v11)
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 18446744073709551615, F1: 738, F2: 0}
    var v13 int32 = 0
    v14 := GenS2{F0: 1, F1: 2, F2: p1}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 651
    v3 := make(map[int]uint32)
    switch 4294967295 {
    case 127:
        var v15 string = ("hello" + ("hello" + ("foo bar" + ((("a" + ("hello" + "foo bar")) + (("foo bar" + "") + "foo bar")) + ((("hello" + "foo bar") + "\n\t") + "hello")))))
        _ = v15
    case 9223372036854775807:
        for fi16 := 0; fi16 < 8; fi16++ {
            delete(v3, (-(fi16)))
            _ = fi16
        }
    case 42:
        if ("\n\t" <= ("hello" + (("" + "foo bar") + "\n\t"))) {
            if true {
                v17 := GenS2{F0: 255, F1: 0, F2: "hello"}
                _ = v17
            }
        } else {
            v2 = -2147483648
        }
    default:
        v18 := GenS1{F0: 0, F1: 127, F2: 1}
        _ = v18
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v19 uint8 = 46
        _ = v19
    default:
        for fi20 := 0; fi20 < 11; fi20++ {
            v2 -= v2
            v1[4294967295] = (!(((((0.1 - 1e10) / 0.1) * float64(0.1)) + 0.1) == float64(((0.0 * 0.1) * (float32((3.14159 / (-2.25 - ((((0.0 - (float64((((1e10 - 3.14159) - 0.1) - (1.5 - 1.5))) * 1.5)) - 0.0) * 0.0) + (1.5 - 1e10))))) + float32(1.5))))))
            _ = fi20
        }
    }
    var v21 uint32 = 0
    v22 := []int{-587, 128}
    if (1 > 42) {
        v2 = v2
        fmt.Println(v21, v2)
        for fi23 := 0; fi23 < 3; fi23++ {
            var v24 string = ""
            _ = fi23
            _ = v24
        }
    } else {
        if true {
            var v25 int16 = 662
            _ = v25
        } else {
            switch 32768 {
            case 2147483647:
                for fi26 := 0; fi26 < 12; fi26++ {
                    v2 = v2
                    switch (("hello" + "\n\t") + ("" + "foo bar")) {
                    case "a":
                        if (true || true) {
                            if (2 > 0) {
                                if ((!true) && (true || false)) {
                                    if false {
                                        v27 := GenS1{F0: 380, F1: -282, F2: 332}
                                        _ = v27
                                    }
                                }
                                var v28 uint8 = 239
                                _ = v28
                            }
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                        }
                    default:
                        for fi29 := 0; fi29 < 6; fi29++ {
                            v2 = (v2 | v2)
                            fmt.Println(v2, v2)
                            delete(v1, v21)
                            v30 := GenS1{F0: 42, F1: (-(fi26)), F2: int64(fi26)}
                            switch fi26 {
                            case -9223372036854775808:
                                var v31 uint = 347
                                _ = v31
                            case 0:
                                fmt.Println(v21)
                            case -70:
                                v32 := GenS1{F0: 1, F1: fi29, F2: 704}
                                _ = v32
                            default:
                                var v33 int32 = -578
                                _ = v33
                            }
                            switch (^(fi29)) {
                            case 128:
                                var v34 int64 = 508
                                _ = v34
                            case -703:
                                if (0 == 18446744073709551615) {
                                    var v35 int8 = -117
                                    _ = v35
                                }
                                v2 += (v2 | 0)
                                v21 = (v21 | (v21 % ((v21) | 1)))
                            }
                            if false {
                                delete(v1, v21)
                            }
                            v36 := GenS1{F0: 2, F1: int(v21), F2: genFn1((4294967295 != 295), "a")}
                            fmt.Println(fi26, fi26)
                            v37 := GenS3{F0: ("a" + "")}
                            v38 := v3[fi29]
                            if ((1 >= 42) && true) {
                                v39 := GenS2{F0: 255, F1: 4294967295, F2: (("a" + "foo bar") + "foo bar")}
                                _ = v39
                            }
                            var v40 uint64 = 885
                            fmt.Println(v40, fi26, v38)
                            _ = fi29
                            _ = v30
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v40
                        }
                    }
                    switch "foo bar" {
                    case "hello":
                        v21 = v21
                    case "foo bar":
                        var v41 string = ("\n\t" + "foo bar")
                        _ = v41
                    }
                    _ = fi26
                }
            default:
                for fi42 := 0; fi42 < 6; fi42++ {
                    for fi43 := 0; fi43 < 17; fi43++ {
                        v2 -= -32
                        v3[fi43] = 65535
                        var v44 float64 = (float64((1.5 * 3.14159)) + 1.5)
                        _ = fi43
                        _ = v44
                    }
                    _ = fi42
                }
            }
        }
    }
    fmt.Println(v2, v21)
    if (!false) {
        fmt.Println(v21)
        v45 := GenS2{F0: 2, F1: 1, F2: ("foo bar" + "hello")}
        for fi46 := 0; fi46 < 19; fi46++ {
            if false {
                for fi47 := 0; fi47 < 5; fi47++ {
                    var v48 uint8 = 1
                    _ = fi47
                    _ = v48
                }
            }
            v21 += (v21 + 32)
            _ = fi46
        }
        _ = v45
    }
    delete(v1, v21)
    v49 := GenS4{F0: 0, F1: "foo bar", F2: 97}
    switch ("foo bar" + "") {
    case "foo bar":
        _ = len(v1)
    case "\n\t":
        for fi50 := 0; fi50 < 11; fi50++ {
            var v51 int = fi50
            _ = fi50
            _ = v51
        }
    default:
        v21 -= (v21 * (v21 / ((v21) | 1)))
    }
    v52 := GenS5{F0: 710}
    v21 -= (v21 % ((v21) | 1))
    v2 = (v2 / ((-2147483648) | 1))
    for fi53 := 0; fi53 < 9; fi53++ {
        for fi54 := 0; fi54 < 8; fi54++ {
            var v55 float32 = 1e10
            var v56 uint16 = 32767
            v56 += (v56 ^ (v56 * (v56 ^ v56)))
            _ = fi54
            _ = v55
            _ = v56
        }
        for fi57 := 0; fi57 < 2; fi57++ {
            var v58 int = 32767
            if ((!(42 <= 1)) && true) {
                v2 -= (-(v2))
                v21 = (v21 + v21)
            }
            if true {
                v21 = v21
            } else {
                v2 = v2
                var v59 int = v58
                _ = v59
            }
            _ = fi57
            _ = v58
        }
        _ = fi53
    }
    v21 = (v21 % ((1) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v49
    _ = v52
}
