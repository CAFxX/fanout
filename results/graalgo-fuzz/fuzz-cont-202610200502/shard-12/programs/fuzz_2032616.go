package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int32 }
type GenS2 struct { F0 string; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int8; F2 uint64 }
func genFn0(p0 uint) int {
    fmt.Println(p0)
    var v6 uint8 = 166
    _ = p0
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 bool, p2 bool) int {
    p1 = false
    if p2 {
        for fi7 := 0; fi7 < 19; fi7++ {
            p0 = p0
            _ = fi7
        }
    } else {
        v8 := []int{}
        p2 = p1
        _ = v8
    }
    var v9 string = "foo bar"
    for fi10 := 0; fi10 < 14; fi10++ {
        p0 = 207
        v9 = "hello"
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: p1, F1: 1000, F2: -2147483648}
            v13 := []int{genFn0(2)}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    v9 = ("a" + v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2147483647
}

func main() {
    var v1 float32 = 1.5
    var v2 uint = 943
    v3 := make(map[string]int)
    var v4 int64 = -1
    var v5 float64 = (0.0 * float64((1e10 / (float32(3.14159) + -2.25))))
    switch "\n\t" {
    case "\n\t":
        if false {
            switch ("" + "hello") {
            case "foo bar":
                v2 += (v2 + 373)
            case "hello":
                var v14 float64 = 0.0
                _ = v14
            case "":
                fmt.Println(v4)
            default:
                delete(v3, (("hello" + "hello") + ("a" + ("" + "\n\t"))))
            }
        }
    case "":
        for fi15 := 0; fi15 < 6; fi15++ {
            if true {
                v16 := []int{fi15, genFn0(v2), (fi15 % ((fi15) | 1))}
                _ = v16
            }
            var v17 int64 = 1
            v18 := v3["a"]
            _ = fi15
            _ = v17
            _ = v18
        }
    }
    var v19 int32 = -181
    v19 = 0
    for fi20 := 0; fi20 < 18; fi20++ {
        if false {
            v21 := []int{}
            switch "\n\t" {
            case "foo bar":
                var v22 float32 = v1
                _ = v22
            case "a":
                for fi23 := 0; fi23 < 7; fi23++ {
                    v1 = 3.14159
                    _ = fi23
                }
            case "hello":
                var v24 bool = true
                _ = v24
            default:
                var v25 bool = true
                _ = v25
            }
            switch fi20 {
            case 127:
                for fi26 := 0; fi26 < 15; fi26++ {
                    fmt.Println(v1, fi26, v5)
                    if true {
                        v27 := GenS1{F0: true, F1: -1, F2: (v19 >> (((v2 ^ v2)) & 63))}
                        _ = len(v3)
                        v3[("foo bar" + "")] = -468
                        _ = v27
                    }
                    _ = fi26
                }
            default:
                v28 := GenS2{F0: "\n\t", F1: -128, F2: 0}
                _ = v28
            }
            _ = v21
        } else {
            if true {
                if (false && (-32768 < 1000)) {
                    delete(v3, "foo bar")
                    v29 := v3["a"]
                    _ = v29
                } else {
                    v30 := []int{(fi20 << (((v2 << (((^(v2))) & 63))) & 63)), -9223372036854775808, (fi20 << (((v2 & (v2 * (v2 / ((v2) | 1))))) & 63)), (fi20 ^ genFn1(12, ("\n\t" < ""), (false || ((v19 >> ((v2) & 63)) == v19))))}
                    v4 = v4
                    if (!false) {
                        if len(v30) >= 2 {
                            v30 = v30[1:len(v30)]
                        }
                        if false {
                            switch ("hello" + ((("a" + "hello") + "foo bar") + "foo bar")) {
                            case "\n\t":
                                v19 = 473
                            case "hello":
                                switch -9223372036854775808 {
                                case 833:
                                    for fi31 := 0; fi31 < 20; fi31++ {
                                        v2 += v2
                                        _ = fi31
                                    }
                                case 32767:
                                    v5 -= float64(1e10)
                                }
                            case "foo bar":
                                _ = len(v30)
                            }
                            v19 = 42
                            if true {
                                var v32 float64 = 1.5
                                v33 := GenS3{F0: 80}
                                _ = v32
                                _ = v33
                            }
                        } else {
                            _ = len(v30)
                            var v34 uint32 = 988
                            _ = v34
                        }
                        v35 := GenS4{F0: 1, F1: 100, F2: 299}
                        _ = v35
                    }
                    _ = v30
                }
                switch "\n\t" {
                case "\n\t":
                    for fi36 := 0; fi36 < 18; fi36++ {
                        var v37 int64 = (v4 & (v4 - (v4 | (v4 * (v4 + (v4 % (((v4 & v4)) | 1)))))))
                        if (!false) {
                            v38 := GenS4{F0: 1, F1: 42, F2: 0}
                            _ = v38
                        }
                        _ = fi36
                        _ = v37
                    }
                case "a":
                    v4 = -1
                }
                v4 += (v4 << ((v2) & 63))
            }
            var v39 float64 = (v5 / 3.14159)
            v39 = 3.14159
            _ = v39
        }
        _ = fi20
    }
    var v40 uint64 = 4294967295
    v5 = (1.5 + 0.1)
    var v41 uint = v2
    var v42 uint8 = 255
    v5 = float64(v1)
    switch 128 {
    case -1:
        v41 -= (^(v2))
    case 609:
        var v43 int8 = -4
        _ = v43
    default:
        if true {
            switch "foo bar" {
            case "":
                if false {
                    fmt.Println(v5, v1, v42)
                } else {
                    var v44 int64 = (v4 >> ((584) & 63))
                    var v45 int8 = -128
                    fmt.Println(v1, v19, v41)
                    _ = v44
                    _ = v45
                }
            case "a":
                v46 := GenS4{F0: 91, F1: 127, F2: v40}
                _ = v46
            default:
                v41 += v2
            }
            for fi47 := 0; fi47 < 3; fi47++ {
                fmt.Println(v40)
                _ = fi47
            }
        }
    }
    v48 := GenS2{F0: "", F1: -1, F2: v40}
    var v49 int8 = 42
    v41 = (v2 >> (((v41 - v2)) & 63))
    v50 := GenS2{F0: "hello", F1: (-(v49)), F2: v40}
    if (0 == 2) {
        _ = len(v3)
        for fi51 := 0; fi51 < 16; fi51++ {
            v1 -= 3.14159
            switch "" {
            case "a":
                var v52 bool = true
                _ = v52
            case "":
                for fi53 := 0; fi53 < 8; fi53++ {
                    var v54 uint = v41
                    v55 := []int{(fi51 << ((v2) & 63)), (-(fi53))}
                    switch (fi51 >> ((565) & 63)) {
                    case 32767:
                        if true {
                            var v56 bool = ((fi51 + 128) < fi51)
                            _ = v56
                        }
                    case 330:
                        v3[("a" + ("foo bar" + ""))] = (fi51 - (fi51 ^ (fi51 + (fi53 | fi51))))
                    case 2147483647:
                        switch fi53 {
                        case 1:
                            v41 -= 4294967295
                        case 2147483647:
                            v4 += (v4 / ((-253) | 1))
                            switch (fi53 - (fi53 | fi51)) {
                            case -1:
                                v54 = 416
                                var v57 uint16 = 397
                                _ = v57
                            case 1:
                                delete(v3, (((("foo bar" + "") + "\n\t") + "") + (("hello" + ((("a" + "foo bar") + "") + (("foo bar" + "hello") + "hello"))) + "")))
                            case -2147483648:
                                if true {
                                    var v58 int32 = (v19 % ((v19) | 1))
                                    _ = v58
                                }
                            default:
                                v59 := v3[("foo bar" + (("foo bar" + ((("a" + "") + "hello") + "\n\t")) + "foo bar"))]
                                _ = v59
                            }
                        case 127:
                            v60 := GenS4{F0: (v49 % ((v49) | 1)), F1: 127, F2: (v40 / ((2) | 1))}
                            _ = v60
                        }
                    default:
                        var v61 int = (fi51 + (fi53 / ((1) | 1)))
                        _ = v61
                    }
                    v42 += 13
                    v3["foo bar"] = (^(fi51))
                    v41 += (v54 - (v41 >> (((^(v54))) & 63)))
                    _ = fi53
                    _ = v54
                    _ = v55
                }
            default:
                if (("a" + "hello") <= "") {
                    var v62 int8 = (^(v49))
                    v63 := []int{fi51, (fi51 - (-(fi51)))}
                    _ = v62
                    _ = v63
                }
            }
            _ = fi51
        }
        var v64 uint8 = (v42 ^ v42)
        _ = v64
    }
    v2 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
