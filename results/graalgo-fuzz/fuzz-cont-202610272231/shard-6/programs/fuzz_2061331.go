package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint8, p1 uint32, p2 uint16) int32 {
    v4 := GenS1{F0: p1, F1: (p1 * (p1 * p1)), F2: int16(p0)}
    for fi5 := 0; fi5 < 4; fi5++ {
        p0 = 208
        v6 := GenS1{F0: (^(p1)), F1: (p1 | (p1 + p1)), F2: -722}
        for fi7 := 0; fi7 < 3; fi7++ {
            p2 = p2
            p1 = (p1 + 4294967295)
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    p0 = (p0 + 95)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return int32(p0)
}
func genFn1(p0 string, p1 uint16) bool {
    switch ("hello" + "") {
    case "foo bar":
        switch 782 {
        case -2147483648:
            var v8 float64 = 1e10
            _ = v8
        case 9223372036854775807:
            v9 := GenS1{F0: 2, F1: 194, F2: 75}
            _ = v9
        }
    case "\n\t":
        switch -280 {
        case 4294967295:
            switch 9223372036854775807 {
            case 32768:
                var v10 uint32 = 65535
                _ = v10
            case 9223372036854775807:
                if true {
                    for fi11 := 0; fi11 < 13; fi11++ {
                        v12 := []int{}
                        for fi13 := 0; fi13 < 7; fi13++ {
                            var v14 float64 = (1.5 / (1e10 + float64(float32(1.5))))
                            fmt.Println(fi11)
                            fmt.Println(p1, fi13, v14)
                            var v15 uint16 = 65535
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        var v16 float32 = 0.0
                        _ = fi11
                        _ = v12
                        _ = v16
                    }
                    v17 := GenS1{F0: 2147483648, F1: 4294967295, F2: 1}
                    fmt.Println(p0, p0)
                    _ = v17
                }
            case -9223372036854775808:
                if (false && ((((p1 >> ((4294967295) & 63)) < p1) || (-1 >= 0)) && (false && ((p1 ^ p1) <= 110)))) {
                    var v18 uint16 = p1
                    fmt.Println(v18, p1, p1)
                    if (18446744073709551615 != 2) {
                        var v19 int64 = -102
                        _ = v19
                    } else {
                        v20 := []int{-41, 287, 9223372036854775807, 32767}
                        if false {
                            if len(v20) > 0 {
                                v20[0] = 42
                            }
                            fmt.Println(p0, p1)
                            var v21 int16 = 42
                            _ = v21
                        }
                        v18 += 255
                        _ = v20
                    }
                    _ = v18
                }
            }
        case -2147483648:
            if true {
                fmt.Println(p1, p0)
                v22 := GenS2{F0: 0}
                _ = v22
            } else {
                var v23 int32 = -2147483648
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 uint = 371
                    switch ("foo bar" + p0) {
                    case "hello":
                        fmt.Println(v23, v23)
                    case "foo bar":
                        if (!false) {
                            if (115 <= int8(v25)) {
                                for fi26 := 0; fi26 < 5; fi26++ {
                                    fmt.Println(p1, fi26, fi26)
                                    if (-9223372036854775808 <= 0) {
                                        p0 = (p0 + "\n\t")
                                        fmt.Println(p0)
                                        p0 = p0
                                        var v27 uint16 = 32767
                                        _ = v27
                                    }
                                    _ = fi26
                                }
                                var v28 uint8 = 1
                                v29 := []int{}
                                _ = v28
                                _ = v29
                            }
                            v30 := []int{fi24, -21, -461}
                            _ = v30
                        }
                    default:
                        for fi31 := 0; fi31 < 1; fi31++ {
                            v32 := GenS2{F0: 37}
                            var v33 string = p0
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                    v34 := GenS1{F0: 2, F1: 0, F2: 181}
                    _ = fi24
                    _ = v25
                    _ = v34
                }
                _ = v23
            }
        case 762:
            if (p1 >= (p1 ^ p1)) {
                var v35 uint = 4294967295
                for fi36 := 0; fi36 < 9; fi36++ {
                    v37 := []int{144, fi36, (fi36 - (fi36 + (fi36 / ((-9223372036854775808) | 1))))}
                    if true {
                        for fi38 := 0; fi38 < 17; fi38++ {
                            var v39 int = (fi36 - fi38)
                            v40 := GenS3{F0: 128, F1: p0}
                            var v41 uint32 = 397
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        v42 := GenS1{F0: 4294967295, F1: 591, F2: -695}
                        _ = v42
                    } else {
                        _ = len(v37)
                        if (p1 > 255) {
                            var v43 string = (("" + "\n\t") + (("\n\t" + ("foo bar" + p0)) + (p0 + ("" + "hello"))))
                            _ = v43
                        } else {
                            var v44 uint16 = (p1 ^ (^(p1)))
                            v44 = (^(p1))
                            v45 := GenS3{F0: 127, F1: ((p0 + "\n\t") + p0)}
                            _ = v44
                            _ = v45
                        }
                    }
                    _ = fi36
                    _ = v37
                }
                _ = v35
            }
        }
    case "":
        v46 := GenS1{F0: 957, F1: 2147483647, F2: 1}
        _ = v46
    }
    var v47 int8 = 100
    v48 := []int{1}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v47
    _ = v48
    return false
}

func main() {
    var v1 uint64 = 42
    var v2 float32 = 0.0
    var v3 uint64 = (^(v1))
    for fi49 := 0; fi49 < 19; fi49++ {
        v50 := []int{}
        _ = fi49
        _ = v50
    }
    switch ("a" + "hello") {
    case "a":
        v51 := GenS4{F0: 276, F1: 485}
        _ = v51
    case "":
        if (9223372036854775807 < 9223372036854775807) {
            v3 = 892
            v2 += float32((0.1 + 0.1))
        } else {
            var v52 int8 = -101
            v53 := []int{128, -1}
            _ = v52
            _ = v53
        }
    case "hello":
        if (((!(65535 <= 447)) || true) || (false || false)) {
            var v54 int = 32768
            if false {
                switch ("" + "a") {
                case "a":
                    fmt.Println(v2, v2, v54)
                default:
                    v55 := GenS3{F0: 1, F1: "a"}
                    _ = v55
                }
            }
            _ = v54
        }
    }
    v56 := GenS5{F0: -936}
    if true {
        if (v3 == 1) {
            for fi57 := 0; fi57 < 8; fi57++ {
                switch fi57 {
                case 9223372036854775807:
                    var v58 float32 = (0.0 * 3.14159)
                    _ = v58
                case 1:
                    fmt.Println(v2)
                default:
                    v1 = (v3 | 2)
                }
                v59 := []int{844}
                _ = fi57
                _ = v59
            }
            fmt.Println(v1)
            v3 = v3
        } else {
            v60 := []int{42, 2147483647}
            _ = v60
        }
    }
    var v61 uint8 = 128
    v1 = 4294967295
    fmt.Println(v2, v3)
    switch "a" {
    case "\n\t":
        fmt.Println(v1, v2)
    case "hello":
        fmt.Println(v3, v61)
    case "a":
        v62 := GenS6{F0: 108, F1: 42}
        _ = v62
    default:
        v63 := []int{}
        _ = v63
    }
    v1 = v1
    var v64 int64 = 9223372036854775807
    v65 := []int{}
    v61 = 46
    v65 = append(v65, 32767)
    v1 = 0
    if len(v65) >= 2 {
        v65 = v65[1:len(v65)]
    }
    if false {
        v3 -= (v3 >> ((944) & 63))
    } else {
        fmt.Println(v61, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v61:", v61)
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v61
    _ = v64
    _ = v65
}
