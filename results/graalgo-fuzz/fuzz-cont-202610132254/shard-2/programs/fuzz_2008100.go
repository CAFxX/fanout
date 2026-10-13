package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 bool }
type GenS5 struct { F0 uint16; F1 int32; F2 uint }
func genFn0(p0 int16, p1 string) uint8 {
    v7 := GenS1{F0: 128}
    if false {
        v8 := []int{}
        v9 := GenS1{F0: 557}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 250
}

func main() {
    var v1 int8 = -15
    var v2 float64 = 1e10
    v3 := make(map[string]int)
    var v4 uint = 42
    v5 := make(map[uint32]int)
    var v6 int32 = 42
    v1 -= 52
    if (("a" + "a") != "") {
        var v10 uint64 = 1
        _ = v10
    }
    var v11 uint32 = 2
    v2 = (0.1 / v2)
    if false {
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: (fi12 << (((v4 | 0)) & 63))}
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 string = "foo bar"
                var v16 uint = 42
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        v17 := GenS2{F0: 42}
        _ = v17
    case "hello":
        if (-1 > 757) {
            v18 := []int{}
            switch "a" {
            case "hello":
                switch "\n\t" {
                case "\n\t":
                    switch -9223372036854775808 {
                    case 127:
                        if true {
                            _ = len(v5)
                            for fi19 := 0; fi19 < 1; fi19++ {
                                v18 = append(v18, fi19)
                                v20 := GenS2{F0: 661}
                                v6 = v6
                                _ = fi19
                                _ = v20
                            }
                        } else {
                            var v21 uint32 = v11
                            var v22 string = ""
                            _ = v21
                            _ = v22
                        }
                    }
                default:
                    if true {
                        fmt.Println(v2)
                    } else {
                        v23 := GenS2{F0: 1}
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 int8 = 127
                            _ = fi24
                            _ = v25
                        }
                        v11 += 2147483647
                        _ = v23
                    }
                }
            case "\n\t":
                switch ("foo bar" + "hello") {
                case "":
                    v18 = append(v18, 2147483647)
                }
            }
            _ = v18
        } else {
            var v26 string = "hello"
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 uint16 = uint16(v11)
                _ = fi27
                _ = v28
            }
            v4 -= (v4 >> ((18446744073709551615) & 63))
            _ = v26
        }
    case "":
        v29 := GenS1{F0: 32767}
        _ = v29
    }
    fmt.Println(v2, v1, v6)
    switch "" {
    case "\n\t":
        if (true && (104 <= 530)) {
            v30 := v5[uint32(v4)]
            _ = v30
        }
    case "hello":
        v1 = int8(v4)
    }
    fmt.Println(v4)
    v3[("hello" + "")] = 1
    v11 -= v11
    v31 := GenS2{F0: 0}
    v4 += v4
    if false {
        if true {
            switch "hello" {
            case "a":
                switch ("\n\t" + ("hello" + "a")) {
                case "foo bar":
                    v32 := []int{32768, -171, 32767}
                    _ = v32
                case "\n\t":
                    v1 += (-(v1))
                case "hello":
                    v33 := []int{493, 9223372036854775807, -600, -1}
                    _ = v33
                }
            }
            for fi34 := 0; fi34 < 12; fi34++ {
                v11 -= uint32(v4)
                v6 -= v6
                v35 := GenS3{F0: 42}
                _ = fi34
                _ = v35
            }
            if (v2 < 3.14159) {
                v2 -= v2
                v11 = 2147483648
            } else {
                switch 659 {
                case 207:
                    for fi36 := 0; fi36 < 8; fi36++ {
                        if (-1 >= -1) {
                            delete(v3, "foo bar")
                            if true {
                                var v37 float32 = 1e10
                                _ = v37
                            } else {
                                switch (^(fi36)) {
                                case -9223372036854775808:
                                    v6 = -1
                                case 32767:
                                    var v38 uint = 18446744073709551615
                                    _ = v38
                                case -2147483648:
                                    for fi39 := 0; fi39 < 16; fi39++ {
                                        v40 := GenS1{F0: 4294967295}
                                        v41 := GenS2{F0: -32768}
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                default:
                                    if (901 != 0) {
                                        switch ("foo bar" + ("foo bar" + ("foo bar" + "\n\t"))) {
                                        case "foo bar":
                                            var v42 float64 = (float64(((0.1 - 1e10) * 0.0)) / v2)
                                            _ = v42
                                        case "":
                                            v43 := v5[v11]
                                            _ = v43
                                        }
                                        v4 = v4
                                    } else {
                                        for fi44 := 0; fi44 < 11; fi44++ {
                                            switch "" {
                                            case "a":
                                                var v45 uint16 = 65535
                                                _ = v45
                                            case "\n\t":
                                                v4 = (^(v4))
                                            case "":
                                                v6 = v6
                                            default:
                                                v46 := GenS1{F0: (-(fi36))}
                                                _ = v46
                                            }
                                            _ = fi44
                                        }
                                        if false {
                                            if (!true) {
                                                v47 := []int{(-(fi36)), (fi36 - (fi36 - fi36)), fi36}
                                                _ = v47
                                            } else {
                                                if false {
                                                    var v48 int8 = (v1 >> (((v4 % ((v4) | 1))) & 63))
                                                    var v49 bool = false
                                                    v2 = float64(1e10)
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    fmt.Println(v4, v4, fi36)
                                                    v2 = (float64(0.0) / ((((0.1 + ((v2 / (v2 - v2)) * -2.25)) / ((float64(0.1) + v2) + v2)) * float64((1.5 / 1.0))) * ((1e10 * v2) / v2)))
                                                    fmt.Println(v2, v2)
                                                }
                                                for fi50 := 0; fi50 < 12; fi50++ {
                                                    v2 = v2
                                                    if false {
                                                        var v51 int64 = -9223372036854775808
                                                        _ = v51
                                                    }
                                                    v4 = (v4 << (((v4 / ((v4) | 1))) & 63))
                                                    _ = fi50
                                                }
                                            }
                                            if false {
                                                for fi52 := 0; fi52 < 7; fi52++ {
                                                    v53 := GenS2{F0: 42}
                                                    v6 -= (^(v6))
                                                    var v54 uint32 = (^(v11))
                                                    _ = fi52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                v4 -= v4
                                            }
                                        } else {
                                            switch int(v6) {
                                            case 127:
                                                v11 = v11
                                            case 469:
                                                v2 = v2
                                            case 32767:
                                                v1 = (v1 >> ((v4) & 63))
                                            default:
                                                var v55 string = "a"
                                                _ = v55
                                            }
                                        }
                                    }
                                }
                                v56 := GenS2{F0: 1}
                                _ = v56
                            }
                            var v57 int16 = -491
                            _ = v57
                        } else {
                            if (true || false) {
                                if (!true) {
                                    v6 += int32(v11)
                                    v58 := GenS3{F0: 1}
                                    v6 = (v6 % (((v6 * v6)) | 1))
                                    _ = v58
                                } else {
                                    v59 := []int{int(v6), fi36}
                                    _ = v59
                                }
                                var v60 int = 705
                                for fi61 := 0; fi61 < 16; fi61++ {
                                    var v62 int16 = 0
                                    v63 := GenS4{F0: 1, F1: v11, F2: (!(41 < 18446744073709551615))}
                                    _ = fi61
                                    _ = v62
                                    _ = v63
                                }
                                _ = v60
                            }
                            v11 -= (v11 | (v11 | 456))
                            v1 = v1
                        }
                        var v64 int = fi36
                        _ = fi36
                        _ = v64
                    }
                case -2147483648:
                    v65 := v3[("" + (("a" + (("\n\t" + "a") + ("a" + "\n\t"))) + "foo bar"))]
                    _ = v65
                case 32768:
                    v1 += 1
                }
                var v66 int8 = int8(v6)
                for fi67 := 0; fi67 < 8; fi67++ {
                    v68 := GenS2{F0: 42}
                    for fi69 := 0; fi69 < 20; fi69++ {
                        v1 = v66
                        _ = fi69
                    }
                    _ = fi67
                    _ = v68
                }
                _ = v66
            }
        } else {
            for fi70 := 0; fi70 < 20; fi70++ {
                v3["a"] = -9223372036854775808
                var v71 int32 = (v6 / ((int32(v11)) | 1))
                var v72 float64 = (v2 + v2)
                _ = fi70
                _ = v71
                _ = v72
            }
        }
    } else {
        v73 := GenS3{F0: -9223372036854775808}
        v74 := []int{347, 2147483647, 662}
        var v75 float32 = 0.0
        _ = v73
        _ = v74
        _ = v75
    }
    var v76 float64 = 3.14159
    switch int(v4) {
    case 2147483647:
        for fi77 := 0; fi77 < 17; fi77++ {
            var v78 float32 = 0.0
            _ = fi77
            _ = v78
        }
    case -2147483648:
        if ((((-2.25 - 1e10) - 3.14159) == 0.1) || false) {
            v1 += (v1 ^ (v1 - 42))
            var v79 float64 = (3.14159 / float64(-2.25))
            for fi80 := 0; fi80 < 11; fi80++ {
                var v81 int8 = (v1 ^ v1)
                _ = fi80
                _ = v81
            }
            _ = v79
        } else {
            v82 := GenS5{F0: 65535, F1: 1, F2: 42}
            v83 := []int{0, -2147483648, -9223372036854775808, 2147483647}
            _ = v82
            _ = v83
        }
    case -9223372036854775808:
        for fi84 := 0; fi84 < 16; fi84++ {
            fmt.Println(fi84, v11)
            for fi85 := 0; fi85 < 11; fi85++ {
                fmt.Println(v11)
                if true {
                    _ = len(v5)
                }
                v86 := GenS3{F0: 1}
                _ = fi85
                _ = v86
            }
            _ = fi84
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
