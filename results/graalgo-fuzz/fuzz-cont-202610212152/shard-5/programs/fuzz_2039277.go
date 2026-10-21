package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 int64 }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int32 }
func genFn0(p0 int) bool {
    switch 1 {
    case 32768:
        v6 := GenS1{F0: 18446744073709551615, F1: "hello", F2: 2147483647}
        _ = v6
    case 2147483647:
        v7 := []int{(p0 ^ 426), p0}
        _ = v7
    case -2147483648:
        v8 := GenS2{F0: 774, F1: 0, F2: 0}
        _ = v8
    default:
        var v9 uint32 = 703
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        switch "" {
        case "":
            p0 -= (fi10 >> ((856) & 63))
        case "a":
            p0 -= (fi10 % ((fi10) | 1))
        }
        v11 := []int{fi10, 4294967295, 127, (p0 - fi10)}
        _ = fi10
        _ = v11
    }
    if false {
        switch "" {
        case "\n\t":
            fmt.Println(p0)
        case "hello":
            p0 -= p0
        case "foo bar":
            switch "" {
            case "\n\t":
                if (-2.25 == 1e10) {
                    p0 = -2147483648
                } else {
                    v12 := GenS2{F0: -1, F1: 124, F2: 9223372036854775807}
                    _ = v12
                }
            }
        }
        if ((1.5 == (1.5 * 0.0)) && (936 == 65535)) {
            p0 -= p0
            v13 := []int{p0, p0, p0, p0}
            _ = v13
        } else {
            if (-1 <= 100) {
                var v14 bool = false
                _ = v14
            }
            v15 := GenS1{F0: 686, F1: "a", F2: 308}
            _ = v15
        }
        fmt.Println(p0)
    }
    _ = p0
    return (((3.14159 + 0.1) * float64(1.5)) <= 0.1)
}
func genFn1() int8 {
    if ((false && true) && (!false)) {
        var v16 int8 = 0
        _ = v16
    }
    var v17 float64 = 1.5
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 int = 42
        fmt.Println(fi18, v19)
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        if (42 > 90) {
            v21 := GenS1{F0: 410, F1: "foo bar", F2: 42}
            _ = v21
        }
        if (false || (!false)) {
            v22 := GenS2{F0: -1, F1: 1, F2: 42}
            if ((("hello" + "\n\t") + ("foo bar" + ("" + "foo bar"))) != "\n\t") {
                v23 := []int{(fi20 % (((fi20 + 426)) | 1)), (-(fi20))}
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                switch "hello" {
                case "\n\t":
                    v17 = v17
                default:
                    v23 = append(v23, 128)
                }
                _ = v23
            } else {
                var v24 uint16 = 2
                fmt.Println(fi20)
                switch "hello" {
                case "hello":
                    var v25 int16 = 808
                    _ = v25
                case "\n\t":
                    v26 := GenS3{F0: ("\n\t" + ("a" + ("foo bar" + "foo bar"))), F1: 1, F2: 42}
                    _ = v26
                case "":
                    var v27 bool = false
                    _ = v27
                default:
                    var v28 bool = true
                    _ = v28
                }
                _ = v24
            }
            v29 := GenS4{F0: 127, F1: 42, F2: 2}
            _ = v22
            _ = v29
        } else {
            switch (fi20 >> ((507) & 63)) {
            case 0:
                var v30 int = (-(fi20))
                _ = v30
            default:
                var v31 uint16 = 0
                _ = v31
            }
            for fi32 := 0; fi32 < 18; fi32++ {
                v17 = ((-2.25 / 1.0) + (v17 * 3.14159))
                v17 = v17
                _ = fi32
            }
        }
        for fi33 := 0; fi33 < 10; fi33++ {
            switch ("" + "a") {
            case "":
                if true {
                    v17 -= (v17 + v17)
                    v17 -= (v17 - (v17 * v17))
                    if (((fi33 / (((fi20 << ((758) & 63))) | 1)) != (-(fi33))) || (false && true)) {
                        var v34 int64 = -1
                        if (0 > 2147483647) {
                            var v35 float64 = (v17 * 0.1)
                            v36 := []int{(fi33 * fi20)}
                            if false {
                                v37 := GenS4{F0: 128, F1: (v34 + 42), F2: 0}
                                var v38 bool = ((((0.1 / 1.0) + ((float64(0.0) + float64(0.1)) - ((v35 * v35) / v17))) + ((v35 - ((v17 + 1e10) / float64((((0.1 * 1e10) * ((0.0 + (3.14159 - 1e10)) - float32((float64(1.5) + 1e10)))) + (0.0 + (0.1 * 0.1)))))) - (1e10 / (v17 * (1.5 * v35))))) < 1.5)
                                var v39 uint = 255
                                _ = v37
                                _ = v38
                                _ = v39
                            } else {
                                v17 -= v35
                            }
                            _ = len(v36)
                            _ = v35
                            _ = v36
                        } else {
                            fmt.Println(v17, v34)
                            fmt.Println(v17, fi33, v17)
                            v40 := []int{(fi20 | fi33), fi33}
                            _ = v40
                        }
                        if true {
                            fmt.Println(v34, v17)
                        } else {
                            if true {
                                v34 = 42
                            } else {
                                v41 := []int{fi33, (-(fi20)), 4294967295, (fi33 / (((fi33 ^ 127)) | 1))}
                                v42 := GenS5{F0: 255, F1: 0, F2: -2147483648}
                                _ = len(v41)
                                _ = v41
                                _ = v42
                            }
                            v43 := []int{}
                            var v44 float64 = float64(((-2.25 + (1e10 - -2.25)) + 0.0))
                            v45 := GenS5{F0: 177, F1: 621, F2: -1}
                            fmt.Println(fi20, fi20)
                            var v46 float64 = ((0.0 / 1.5) + float64((1e10 - 0.1)))
                            _ = v43
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        _ = v34
                    } else {
                        v17 = 0.1
                    }
                } else {
                    if false {
                        var v47 float32 = 0.1
                        _ = v47
                    }
                }
            }
            _ = fi33
        }
        _ = fi20
    }
    _ = v17
    return 1
}

func main() {
    var v1 int16 = 770
    var v2 uint = 255
    var v3 uint = (v2 & 888)
    var v4 int16 = v1
    v5 := make(map[string]string)
    switch 4294967295 {
    case -851:
        var v48 uint8 = 128
        _ = v48
    case 32768:
        v49 := GenS5{F0: 2, F1: 65535, F2: -2147483648}
        _ = v49
    case 0:
        v2 -= (v3 / ((611) | 1))
    default:
        switch -454 {
        case 32768:
            if true {
                for fi50 := 0; fi50 < 12; fi50++ {
                    fmt.Println(v2, v3, v3)
                    v4 = -410
                    _ = fi50
                }
            }
        default:
            v2 = 4294967295
        }
    }
    v51 := GenS5{F0: 127, F1: 2147483648, F2: 969}
    var v52 float32 = (1.5 * (0.0 / 3.14159))
    var v53 uint32 = 2147483647
    v4 = (-(v1))
    if false {
        var v54 uint32 = v53
        _ = v54
    } else {
        for fi55 := 0; fi55 < 20; fi55++ {
            switch fi55 {
            case -101:
                var v56 uint16 = 711
                _ = v56
            case -9223372036854775808:
                v3 += 65535
            case -468:
                if (42 == 349) {
                    for fi57 := 0; fi57 < 17; fi57++ {
                        v52 = 1.5
                        _ = fi57
                    }
                    v58 := GenS4{F0: 255, F1: 9223372036854775807, F2: 42}
                    v4 += (v1 ^ v1)
                    _ = v58
                } else {
                    var v59 bool = (uint64(v2) != 18446744073709551615)
                    var v60 uint8 = 205
                    _ = v59
                    _ = v60
                }
            }
            if (!(true && (!(2 <= 255)))) {
                switch (("foo bar" + "") + "hello") {
                case "foo bar":
                    v2 = 615
                case "hello":
                    var v61 float64 = 1e10
                    _ = v61
                case "a":
                    v52 -= v52
                default:
                    if false {
                        if (-759 < -1) {
                            var v62 bool = true
                            _ = v62
                        } else {
                            v63 := GenS4{F0: 0, F1: 42, F2: 91}
                            _ = v63
                        }
                        var v64 uint8 = 25
                        _ = v64
                    }
                }
                v5["\n\t"] = ""
                v65 := v5["foo bar"]
                _ = v65
            } else {
                for fi66 := 0; fi66 < 3; fi66++ {
                    fmt.Println(v52, fi66, v52)
                    switch "foo bar" {
                    case "hello":
                        v2 = (v3 % ((v2) | 1))
                    case "\n\t":
                        var v67 int64 = 0
                        _ = v67
                    default:
                        delete(v5, "foo bar")
                    }
                    fmt.Println(fi66, fi55, fi55)
                    v68 := v5["a"]
                    var v69 float64 = 1e10
                    _ = fi66
                    _ = v68
                    _ = v69
                }
                for fi70 := 0; fi70 < 12; fi70++ {
                    var v71 int = fi55
                    _ = fi70
                    _ = v71
                }
                v2 = v3
            }
            _ = fi55
        }
    }
    for fi72 := 0; fi72 < 14; fi72++ {
        v73 := v5["hello"]
        _ = fi72
        _ = v73
    }
    var v74 uint8 = 1
    var v75 uint8 = v74
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v74
    _ = v75
}
