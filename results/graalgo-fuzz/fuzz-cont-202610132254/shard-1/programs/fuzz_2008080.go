package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 string; F2 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int16; F1 int64 }
func genFn0(p0 uint, p1 int) int64 {
    for fi6 := 0; fi6 < 7; fi6++ {
        p0 -= p0
        var v7 uint64 = 2
        switch ("a" + (("" + "\n\t") + ("foo bar" + ""))) {
        case "hello":
            p1 += fi6
        case "\n\t":
            if false {
                v7 -= v7
                var v8 int32 = -1
                v9 := GenS1{F0: 127, F1: (("a" + ("hello" + "a")) + "\n\t"), F2: p1}
                _ = v8
                _ = v9
            }
        case "a":
            for fi10 := 0; fi10 < 11; fi10++ {
                p0 = p0
                _ = fi10
            }
        }
        _ = fi6
        _ = v7
    }
    p1 += 633
    var v11 int64 = -359
    v12 := []int{(-(p1)), p1, p1}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return (v11 - v11)
}
func genFn1() int16 {
    v13 := GenS1{F0: 42, F1: "foo bar", F2: 273}
    var v14 bool = ((151 >= 18446744073709551615) || true)
    var v15 uint = 4294967295
    for fi16 := 0; fi16 < 8; fi16++ {
        switch ((("a" + "\n\t") + "hello") + "hello") {
        case "a":
            if true {
                v15 -= (v15 + 355)
                var v17 string = ((("hello" + ("hello" + "")) + "foo bar") + "a")
                _ = v17
            } else {
                v14 = v14
                switch (fi16 + int(v15)) {
                case 32768:
                    v18 := GenS2{F0: 42, F1: 128}
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 2; fi19++ {
                        switch "foo bar" {
                        case "foo bar":
                            v15 = 1
                        case "":
                            v20 := GenS3{F0: genFn0(v15, (fi16 ^ 462))}
                            _ = v20
                        case "a":
                            v21 := []int{(fi19 << ((4294967295) & 63)), -9223372036854775808, (fi19 / (((fi16 >> ((v15) & 63))) | 1)), fi19}
                            _ = v21
                        default:
                            v15 = v15
                        }
                        _ = fi19
                    }
                }
                if v14 {
                    v22 := []int{-548, -988, fi16}
                    for fi23 := 0; fi23 < 16; fi23++ {
                        fmt.Println(v14, v14, fi23)
                        if (297 <= 2147483647) {
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            var v24 int = (-(fi16))
                            v25 := GenS4{F0: -1}
                            _ = v24
                            _ = v25
                        }
                        var v26 int32 = 0
                        _ = fi23
                        _ = v26
                    }
                    _ = v22
                } else {
                    var v27 uint = (v15 << ((42) & 63))
                    _ = v27
                }
            }
        case "hello":
            v28 := GenS4{F0: -32768}
            _ = v28
        default:
            v14 = (0 >= 756)
        }
        _ = fi16
    }
    _ = v13
    _ = v14
    _ = v15
    return -971
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 3.14159
    var v3 float32 = (1.5 / ((0.1 + 3.14159) - -2.25))
    var v4 int = -2147483648
    var v5 int64 = -9223372036854775808
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := GenS5{F0: genFn0(0, v4), F1: "foo bar", F2: 2}
        v3 = v3
        v5 = (-(v5))
        _ = fi29
        _ = v30
    }
    v31 := GenS4{F0: genFn1()}
    var v32 int8 = 42
    if (v4 > (v4 << ((255) & 63))) {
        if (2 != 18446744073709551615) {
            var v33 int = v4
            for fi34 := 0; fi34 < 17; fi34++ {
                for fi35 := 0; fi35 < 18; fi35++ {
                    var v36 uint64 = 4294967295
                    _ = fi35
                    _ = v36
                }
                v33 = v4
                _ = fi34
            }
            _ = v33
        } else {
            v4 -= (v4 + (v4 * v4))
        }
    }
    for fi37 := 0; fi37 < 10; fi37++ {
        for fi38 := 0; fi38 < 17; fi38++ {
            _ = len(v1)
            var v39 string = ("hello" + "a")
            switch "a" {
            case "":
                switch fi37 {
                case 0:
                    v40 := []int{(fi38 | fi37)}
                    _ = v40
                case -555:
                    fmt.Println(fi37, fi37)
                }
            case "\n\t":
                var v41 float64 = v2
                _ = v41
            case "foo bar":
                if (534 != 267) {
                    v32 -= v32
                }
                v32 -= v32
                v4 -= 32767
                var v42 uint = 4294967295
                _ = v42
            default:
                var v43 int16 = -32768
                _ = v43
            }
            if false {
                v5 += -9223372036854775808
                if (true || (2147483647 > 69)) {
                    var v44 uint32 = 2
                    switch 0 {
                    case 4294967295:
                        v45 := GenS5{F0: 9223372036854775807, F1: "hello", F2: 128}
                        _ = v45
                    case -9223372036854775808:
                        v46 := v1["foo bar"]
                        _ = v46
                    }
                    switch 32768 {
                    case 590:
                        var v47 uint = 280
                        _ = v47
                    }
                    _ = v44
                } else {
                    fmt.Println(v3, v2)
                    fmt.Println(v39, v2)
                }
                v48 := GenS4{F0: 0}
                _ = v48
            } else {
                v3 = (((v3 / (3.14159 + (0.0 - v3))) - v3) * 1e10)
                if ((v32 == v32) || false) {
                    fmt.Println(v2)
                }
                var v49 float64 = (0.0 * v2)
                var v50 string = ""
                _ = v49
                _ = v50
            }
            _ = fi38
            _ = v39
        }
        v51 := []int{v4, fi37, (^(fi37))}
        fmt.Println(v3)
        _ = fi37
        _ = v51
    }
    if (-2147483648 < -727) {
        v52 := GenS1{F0: 0, F1: "a", F2: (-(v4))}
        for fi53 := 0; fi53 < 18; fi53++ {
            for fi54 := 0; fi54 < 6; fi54++ {
                v3 += (v3 + v3)
                v55 := GenS6{F0: v32}
                _ = fi54
                _ = v55
            }
            var v56 bool = false
            _ = fi53
            _ = v56
        }
        _ = v52
    }
    v57 := GenS7{F0: int32(v32)}
    v58 := GenS7{F0: -137}
    fmt.Println(v4, v32)
    fmt.Println(v4, v32)
    if (97 < 127) {
        if (((0.0 / ((3.14159 + (v2 - v2)) + v2)) >= float64((0.0 * v3))) && false) {
            v3 = (v3 * -2.25)
            for fi59 := 0; fi59 < 3; fi59++ {
                v2 -= ((v2 / ((v2 + (0.0 * v2)) * v2)) - ((-2.25 / v2) - 0.1))
                if ((4294967295 >= 49) && true) {
                    v60 := []int{(^(v4)), fi59, v4}
                    _ = v60
                }
                _ = fi59
            }
            _ = len(v1)
        }
        v2 -= (v2 * v2)
        v4 = v4
    }
    if false {
        v61 := GenS8{F0: genFn1(), F1: (v5 / (((v5 << ((25) & 63))) | 1))}
        var v62 int64 = v5
        for fi63 := 0; fi63 < 7; fi63++ {
            if true {
                switch fi63 {
                case 4294967295:
                    _ = len(v1)
                default:
                    var v64 int8 = v32
                    _ = v64
                }
                var v65 int32 = 463
                _ = v65
            } else {
                v4 = (^(fi63))
                switch v4 {
                case 127:
                    if (212 == 42) {
                        v66 := GenS7{F0: 2147483647}
                        _ = v66
                    }
                }
            }
            for fi67 := 0; fi67 < 20; fi67++ {
                v2 -= 0.0
                _ = fi67
            }
            _ = fi63
        }
        _ = v61
        _ = v62
    }
    v68 := []int{(v4 / (((v4 + v4)) | 1)), 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v68:", len(v68), v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v68
}
