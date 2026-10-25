package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint16; F1 int64 }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 int8 = 1
        v6 -= v6
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 float32 = (0.0 - 0.1)
            if true {
                v9 := GenS1{F0: (^(v6)), F1: -2147483648}
                v6 = (v6 % ((1) | 1))
                v10 := GenS1{F0: (-(v6)), F1: 1}
                v11 := []int{}
                v6 -= (v6 - v6)
                var v12 float64 = (0.0 + ((1.5 - 0.0) + ((0.0 - 0.0) * 0.1)))
                v11 = append(v11, -454)
                v8 += v8
                _ = v10
                _ = v11
                _ = v12
                _ = v9
            }
            switch (("foo bar" + ("hello" + "a")) + "") {
            case "foo bar":
                fmt.Println(fi7, fi7, v6)
            case "\n\t":
                switch fi7 {
                case 1:
                    switch "" {
                    case "hello":
                        var v13 uint = 65535
                        _ = v13
                    case "\n\t":
                        if ((3.14159 != (1e10 * float32((0.0 - -2.25)))) && true) {
                            v14 := GenS1{F0: v6, F1: -1}
                            fmt.Println(fi7)
                            _ = v14
                        } else {
                            switch "a" {
                            case "":
                                switch (fi5 | (fi5 ^ 127)) {
                                case 127:
                                    var v15 int16 = 0
                                    _ = v15
                                case 678:
                                    v16 := []int{9223372036854775807}
                                    _ = v16
                                case -2147483648:
                                    if false {
                                        v17 := GenS2{F0: 2147483648, F1: 65535, F2: 18446744073709551615}
                                        var v18 uint = 4294967295
                                        _ = v17
                                        _ = v18
                                    } else {
                                        var v19 int64 = 1
                                        v20 := GenS1{F0: v6, F1: -1}
                                        _ = v19
                                        _ = v20
                                    }
                                default:
                                    v21 := GenS2{F0: 2, F1: 543, F2: 18446744073709551615}
                                    _ = v21
                                }
                            case "\n\t":
                                v22 := GenS1{F0: v6, F1: 2147483647}
                                _ = v22
                            }
                            v23 := []int{fi7, 127, 32767, 9223372036854775807}
                            _ = v23
                        }
                    }
                }
            }
            if ((3.14159 + (0.1 + 3.14159)) != 1.5) {
                var v24 float32 = (v8 / ((v8 / 1.0) - ((v8 + (v8 / v8)) - v8)))
                if false {
                    v25 := GenS3{F0: fi5}
                    if true {
                        v26 := []int{}
                        v27 := GenS1{F0: v6, F1: -1}
                        _ = v26
                        _ = v27
                    }
                    switch ("" + (("hello" + "hello") + "")) {
                    case "hello":
                        if false {
                            v6 -= -69
                            var v28 float64 = -2.25
                            v8 += float32(v28)
                            v28 -= (v28 - v28)
                            _ = v28
                        } else {
                            fmt.Println(fi5)
                            fmt.Println(v8, fi5, fi5)
                            if false {
                                var v29 uint8 = 0
                                if (true || (4294967295 >= 42)) {
                                    v29 = (v29 % (((v29 + v29)) | 1))
                                } else {
                                    var v30 uint8 = (v29 & v29)
                                    _ = v30
                                }
                                v8 += (v8 + v8)
                                _ = v29
                            }
                        }
                    default:
                        if (-1 == (v6 + (v6 * (v6 & 42)))) {
                            var v31 int16 = -30
                            fmt.Println(v24, v31, v24)
                            v24 = (((v8 * (1.5 * (float32(float64(float32((1e10 * 0.0)))) + float32(1.5)))) / 1.0) - 1e10)
                            var v32 uint8 = 2
                            _ = v31
                            _ = v32
                        }
                    }
                    _ = v25
                } else {
                    v24 -= (((1.5 * v8) * -2.25) * (v8 * v24))
                    v33 := GenS4{F0: (fi5 << ((0) & 63)), F1: 2, F2: 2}
                    switch 32768 {
                    case 128:
                        v6 = v6
                    case -2147483648:
                        v34 := GenS2{F0: 443, F1: 2, F2: 42}
                        _ = v34
                    case -9223372036854775808:
                        switch fi5 {
                        case 128:
                            v35 := GenS4{F0: (fi5 ^ (fi5 * (fi5 + (-(fi7))))), F1: 4294967295, F2: 2147483648}
                            _ = v35
                        default:
                            var v36 float64 = (float64(v24) + (-2.25 / -2.25))
                            _ = v36
                        }
                    default:
                        v37 := []int{0, 128, fi5}
                        _ = v37
                    }
                    _ = v33
                }
                var v38 float32 = 3.14159
                _ = v24
                _ = v38
            } else {
                v39 := GenS2{F0: 0, F1: 18446744073709551615, F2: 4294967295}
                _ = v39
            }
            var v40 int = fi5
            _ = fi7
            _ = v40
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    var v41 bool = false
    v41 = (-9223372036854775808 < 889)
    v42 := GenS5{F0: 1, F1: 32768}
    _ = v41
    _ = v42
    return 255
}
func genFn1(p0 int, p1 string, p2 uint) int {
    p2 = (p2 - (p2 << ((820) & 63)))
    if (!false) {
        for fi43 := 0; fi43 < 17; fi43++ {
            p0 = 1
            _ = fi43
        }
    }
    v44 := GenS3{F0: (p0 - 128)}
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    return p0
}

func main() {
    v1 := []uint32{2147483648}
    var v2 int32 = -1
    v3 := []int{61}
    var v4 int16 = 42
    v45 := GenS4{F0: 4294967295, F1: 18446744073709551615, F2: 0}
    for fi46 := 0; fi46 < 10; fi46++ {
        var v47 float32 = 3.14159
        _ = fi46
        _ = v47
    }
    fmt.Println(v2, v4)
    switch -2147483648 {
    case -2147483648:
        var v48 uint8 = 128
        _ = v48
    case 32767:
        v2 = v2
    case 1:
        fmt.Println(v4, v2)
    }
    var v49 float32 = 0.0
    switch "a" {
    case "foo bar":
        var v50 float64 = (3.14159 * 0.1)
        _ = v50
    case "a":
        v2 = (v2 + v2)
    case "hello":
        var v51 int16 = v4
        _ = v51
    default:
        switch ("a" + ("foo bar" + "a")) {
        case "\n\t":
            v49 = 0.1
        case "a":
            fmt.Println(v2, v2, v49)
        default:
            v49 -= float32(float64(v49))
        }
    }
    switch 9223372036854775807 {
    case -310:
        v4 -= 0
    }
    fmt.Println(v4, v49, v2)
    v52 := GenS5{F0: 238, F1: 1}
    v49 = (float32((1e10 * 0.1)) / ((1.5 / (-2.25 * v49)) - v49))
    switch -718 {
    case 0:
        var v53 float64 = 0.1
        _ = v53
    }
    v54 := GenS6{F0: 858, F1: 9223372036854775807}
    if ((0.1 >= 0.0) || (!((float32(3.14159) / (1.5 + v49)) <= v49))) {
        if true {
            var v55 int8 = 127
            var v56 uint8 = 26
            fmt.Println(v56)
            _ = v55
            _ = v56
        }
        v57 := GenS3{F0: genFn1(32767, "a", 255)}
        switch -8 {
        case 0:
            v4 -= v4
        case 32768:
            var v58 float32 = v49
            _ = v58
        case 688:
            v2 = -2147483648
        default:
            v4 = (v4 & (v4 % ((v4) | 1)))
        }
        _ = v57
    }
    switch "foo bar" {
    case "\n\t":
        switch ("foo bar" + ("" + ("" + ("hello" + "hello")))) {
        case "foo bar":
            for fi59 := 0; fi59 < 5; fi59++ {
                v2 = (v2 % ((v2) | 1))
                _ = fi59
            }
        case "\n\t":
            for fi60 := 0; fi60 < 16; fi60++ {
                v49 -= float32(float64(v49))
                if false {
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        if (v49 >= 0.0) {
                            v49 = 1e10
                            switch "foo bar" {
                            case "\n\t":
                                v61 := GenS4{F0: (fi60 >> ((0) & 63)), F1: 30, F2: 300}
                                _ = v61
                            case "":
                                fmt.Println(v2)
                            }
                        } else {
                            v4 = (v4 & (v4 % ((v4) | 1)))
                        }
                    case "\n\t":
                        var v62 int32 = (v2 >> ((18446744073709551615) & 63))
                        _ = v62
                    }
                }
                v49 = (float32((0.0 - float64(0.0))) / -2.25)
                _ = fi60
            }
        case "hello":
            if len(v3) > 0 {
                v3[0] = 9223372036854775807
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case "a":
        v49 -= (-2.25 / (v49 + float32(0.1)))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v52
    _ = v54
}
