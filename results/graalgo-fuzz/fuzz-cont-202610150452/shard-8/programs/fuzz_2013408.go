package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int }
func genFn0() uint {
    if (false || false) {
        v4 := GenS1{F0: ("a" + "hello")}
        _ = v4
    } else {
        if (false && (("hello" == "a") && (-32768 > -105))) {
            var v5 uint = 65535
            v5 += v5
            _ = v5
        }
        v6 := GenS1{F0: "\n\t"}
        switch "a" {
        case "foo bar":
            v7 := GenS1{F0: "\n\t"}
            _ = v7
        case "":
            var v8 int16 = 0
            _ = v8
        }
        _ = v6
    }
    var v9 uint = 2
    var v10 int8 = -3
    var v11 uint64 = 18446744073709551615
    switch 0 {
    case -1:
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS2{F0: (v11 + uint64(v9)), F1: (^(v9))}
            v11 -= 2
            _ = fi12
            _ = v13
        }
    case -9223372036854775808:
        v11 -= (v11 << ((v9) & 63))
    case -945:
        if (false && (!true)) {
            var v14 int = 387
            v10 += (v10 * (v10 * 100))
            v15 := GenS2{F0: (v11 | (v11 / ((886) | 1))), F1: 351}
            _ = v14
            _ = v15
        } else {
            var v16 int16 = 1000
            v17 := GenS3{F0: 42}
            _ = v16
            _ = v17
        }
    default:
        var v18 bool = (true || true)
        _ = v18
    }
    _ = v10
    _ = v11
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 int32, p1 bool, p2 int64) uint32 {
    var v19 uint32 = 0
    if (-1 == 1) {
        switch 707 {
        case 299:
            if ((true && p1) || p1) {
                p1 = false
                v19 -= (v19 ^ 42)
                var v20 uint64 = 1
                _ = v20
            } else {
                v21 := GenS4{F0: 303}
                if ((("a" + "a") + ("" + "\n\t")) > "a") {
                    if p1 {
                        p0 -= (p0 + p0)
                    } else {
                        var v22 bool = p1
                        v23 := []int{522, -840, 364}
                        _ = v22
                        _ = v23
                    }
                    var v24 string = "hello"
                    var v25 int32 = (p0 % ((p0) | 1))
                    _ = v24
                    _ = v25
                } else {
                    p0 -= (p0 ^ (p0 / ((p0) | 1)))
                    fmt.Println(p2)
                }
                p1 = false
                _ = v21
            }
        case -2147483648:
            var v26 int8 = -1
            _ = v26
        case 503:
            var v27 string = "hello"
            _ = v27
        default:
            if (p0 != p0) {
                v28 := GenS4{F0: 753}
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 3; fi29++ {
            v30 := []int{}
            var v31 bool = (!p1)
            _ = fi29
            _ = v30
            _ = v31
        }
        fmt.Println(p2)
    } else {
        v32 := GenS5{F0: p0}
        _ = v32
    }
    v33 := []int{}
    var v34 int16 = 42
    v35 := GenS5{F0: (p0 * (p0 + -2147483648))}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v33
    _ = v34
    _ = v35
    return v19
}

func main() {
    var v1 int64 = 932
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    v36 := GenS3{F0: 0}
    v37 := []int{127, -6, 2147483647, 2147483647}
    if false {
        var v38 int16 = 853
        _ = v38
    }
    fmt.Println(v1)
    _ = len(v37)
    var v39 uint32 = 410
    v1 = (v1 & (v1 >> ((4294967295) & 63)))
    v1 += (v1 / ((0) | 1))
    var v40 uint32 = (^(v39))
    if (!false) {
        var v41 int32 = int32(v1)
        switch "\n\t" {
        case "foo bar":
            for fi42 := 0; fi42 < 14; fi42++ {
                var v43 int16 = 0
                v41 = v41
                _ = fi42
                _ = v43
            }
        default:
            for fi44 := 0; fi44 < 3; fi44++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(v40, v41, fi44)
                case "":
                    var v45 string = "a"
                    _ = v45
                case "\n\t":
                    fmt.Println(v39, v40, v40)
                }
                _ = fi44
            }
        }
        v41 += v41
        _ = v41
    }
    var v46 float64 = 3.14159
    for fi47 := 0; fi47 < 2; fi47++ {
        if false {
            if false {
                v39 = v39
                for fi48 := 0; fi48 < 20; fi48++ {
                    if false {
                        if true {
                            v2["foo bar"] = true
                            var v49 int8 = 127
                            if true {
                                _ = len(v2)
                            }
                            _ = v49
                        }
                    }
                    if true {
                        v39 = genFn1(1, (v40 == 288), v1)
                        var v50 int32 = 1
                        v37 = append(v37, fi48)
                        _ = v50
                    }
                    _ = fi48
                }
            }
            v3[(fi47 | fi47)] = false
        } else {
            v40 -= v40
            v51 := GenS3{F0: -28}
            v1 = v1
            _ = v51
        }
        for fi52 := 0; fi52 < 5; fi52++ {
            if len(v37) >= 2 {
                v37 = v37[1:len(v37)]
            }
            var v53 string = "\n\t"
            _ = fi52
            _ = v53
        }
        var v54 bool = true
        _ = fi47
        _ = v54
    }
    switch 9223372036854775807 {
    case 128:
        fmt.Println(v39)
    case 998:
        if false {
            if len(v37) > 0 {
                v37[0] = -2147483648
            }
            fmt.Println(v1, v1)
        } else {
            var v55 uint = 65535
            _ = v55
        }
    case 458:
        switch "a" {
        case "":
            v40 -= 580
        default:
            var v56 float32 = float32((float64((((-2.25 + (0.0 * (3.14159 - (1e10 / 1.5)))) * (3.14159 / ((1.5 / (1.5 / 1.5)) - (((0.1 + 1.5) / 0.1) - 0.0)))) * -2.25)) + ((((((v46 * (float64(float32(0.1)) / v46)) / v46) + ((v46 + v46) * ((0.1 - v46) + v46))) - (((((v46 / float64(((1e10 * float32((v46 / float64(3.14159)))) / 3.14159))) / v46) - v46) + float64(float32(v46))) + v46)) - (((v46 + float64(1.5)) * v46) / v46)) - v46)))
            _ = v56
        }
    default:
        v57 := GenS6{F0: 4294967295}
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v40
    _ = v46
}
