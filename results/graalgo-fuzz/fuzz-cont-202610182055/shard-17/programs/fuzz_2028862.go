package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint8; F1 bool; F2 string }
func genFn0(p0 int8, p1 uint16) bool {
    p1 -= p1
    var v6 float32 = (0.1 / 1.0)
    var v7 uint16 = 32767
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (1e10 <= 0.0)
}
func genFn1() uint8 {
    var v8 float32 = 3.14159
    if (false || (v8 > 3.14159)) {
        v8 = (float32((1e10 + 1.5)) / (v8 - (1.5 + 0.0)))
        var v9 uint8 = 0
        var v10 uint64 = 18446744073709551615
        _ = v10
        _ = v9
    } else {
        v11 := []int{227, -9223372036854775808, 218, 42}
        if false {
            for fi12 := 0; fi12 < 7; fi12++ {
                switch "a" {
                case "hello":
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v8 = (-2.25 / float32(3.14159))
                        if false {
                            if len(v11) > 0 {
                                v11[0] = 0
                            }
                            var v14 uint32 = 65535
                            _ = v14
                        } else {
                            var v15 float32 = (v8 * 1e10)
                            v16 := GenS1{F0: -715}
                            if false {
                                var v17 uint = 2
                                switch "\n\t" {
                                case "foo bar":
                                    v15 += v8
                                }
                                _ = v17
                            } else {
                                v8 -= v15
                                if true {
                                    switch "\n\t" {
                                    case "foo bar":
                                        if len(v11) > 0 {
                                            v11[0] = 140
                                        }
                                    case "":
                                        switch ("a" + ("foo bar" + "")) {
                                        case "a":
                                            var v18 int16 = -396
                                            _ = v18
                                        case "hello":
                                            switch ("foo bar" + "") {
                                            case "hello":
                                                fmt.Println(v15, fi12, v8)
                                            }
                                        }
                                    }
                                    v19 := GenS1{F0: 621}
                                    switch "a" {
                                    case "a":
                                        var v20 uint32 = 282
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        v21 := GenS2{F0: -371, F1: -9223372036854775808}
                                        _ = v20
                                        _ = v21
                                    case "":
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                    case "foo bar":
                                        v15 = (1.5 + (v8 + float32(0.1)))
                                        switch ("foo bar" + "") {
                                        case "\n\t":
                                            var v22 uint16 = 2
                                            _ = v22
                                        case "a":
                                            switch (fi13 & fi13) {
                                            case 727:
                                                v15 = (1.5 + 1.5)
                                            case 1:
                                                v15 -= (float32(((float64(v8) / 0.1) - 1e10)) + 3.14159)
                                            default:
                                                var v23 string = "hello"
                                                var v24 uint32 = 2147483648
                                                v8 = -2.25
                                                _ = v23
                                                _ = v24
                                            }
                                        case "hello":
                                            if (42 > 32767) {
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                                if true {
                                                    v8 += 1.5
                                                    v8 += ((v8 - ((0.1 * v8) + 0.1)) - v8)
                                                } else {
                                                    v11 = append(v11, fi13)
                                                    var v25 uint8 = 21
                                                    _ = v25
                                                }
                                            }
                                        }
                                    default:
                                        fmt.Println(v15, v8)
                                    }
                                    _ = v19
                                }
                            }
                            _ = v15
                            _ = v16
                        }
                        _ = fi13
                    }
                case "foo bar":
                    v8 += v8
                }
                if (false || false) {
                    v8 = v8
                } else {
                    v26 := GenS2{F0: -1, F1: 42}
                    v11 = append(v11, 409)
                    fmt.Println(fi12, v8)
                    _ = v26
                }
                _ = fi12
            }
        }
        _ = v11
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        if (1 != 65535) {
            switch ("hello" + ("hello" + "foo bar")) {
            case "\n\t":
                switch (-(fi27)) {
                case 127:
                    v8 = v8
                default:
                    v28 := GenS3{F0: false, F1: 701, F2: 0}
                    _ = v28
                }
            case "hello":
                if false {
                    v8 += ((v8 / (-2.25 + v8)) * v8)
                    var v29 bool = false
                    v8 -= 1.5
                    _ = v29
                }
            default:
                var v30 int64 = -1
                _ = v30
            }
            if ((765 != 348) || (false && true)) {
                v31 := GenS1{F0: 1}
                _ = v31
            } else {
                v8 += float32(float64(((((v8 - ((v8 / v8) * v8)) - v8) * v8) / float32(1.5))))
            }
            v32 := []int{32768, fi27}
            _ = v32
        }
        v8 = ((0.0 - v8) - 0.1)
        for fi33 := 0; fi33 < 7; fi33++ {
            var v34 uint32 = 1
            _ = fi33
            _ = v34
        }
        _ = fi27
    }
    _ = v8
    return 255
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := []int{-293, 9223372036854775807, 4294967295}
    v4 := make(map[uint32]bool)
    var v5 uint16 = 1
    for fi35 := 0; fi35 < 1; fi35++ {
        v5 -= v5
        _ = fi35
    }
    v36 := GenS3{F0: false, F1: 0, F2: 127}
    v5 += 192
    for fi37 := 0; fi37 < 1; fi37++ {
        var v38 string = "hello"
        v38 = v38
        _ = fi37
        _ = v38
    }
    switch 9223372036854775807 {
    case 4294967295:
        if (42 == 116) {
            var v39 int32 = 308
            v5 -= v5
            _ = v39
        } else {
            var v40 int16 = 32767
            switch "foo bar" {
            case "a":
                switch -979 {
                case 32767:
                    v3 = append(v3, 1)
                case 1:
                    v40 += (v40 & 32767)
                case 540:
                    v5 -= 42
                default:
                    v40 -= v40
                }
            }
            _ = v40
        }
    case 812:
        v5 += (v5 + 2)
    case -97:
        v5 = 65535
    default:
        v41 := GenS3{F0: (-32768 < -563), F1: v5, F2: genFn1()}
        _ = v41
    }
    var v42 float32 = ((1.5 + ((float32(3.14159) - 1e10) + 0.0)) + (1.5 * 3.14159))
    v43 := GenS4{F0: 43, F1: true, F2: "hello"}
    if false {
        if ((v42 - v42) != 3.14159) {
            if (3.14159 > 1e10) {
                _ = len(v3)
                var v44 uint64 = 0
                fmt.Println(v44, v44, v42)
                _ = v44
            } else {
                var v45 int16 = 32767
                var v46 int16 = (v45 % (((v45 / ((0) | 1))) | 1))
                v42 += (((v42 * (v42 * 3.14159)) / 0.1) - (((-2.25 + ((v42 - v42) + v42)) / 1.0) / v42))
                _ = v45
                _ = v46
            }
        } else {
            v42 = v42
            var v47 string = "foo bar"
            _ = v47
        }
        var v48 float64 = (3.14159 - 1e10)
        for fi49 := 0; fi49 < 10; fi49++ {
            v42 -= (1e10 + (float32(v48) * v42))
            _ = fi49
        }
        _ = v48
    } else {
        var v50 int64 = 0
        if (9223372036854775807 < 156) {
            var v51 int = 32768
            _ = v51
        } else {
            var v52 uint32 = 682
            v50 = int64(v52)
            _ = v52
        }
        _ = v50
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
