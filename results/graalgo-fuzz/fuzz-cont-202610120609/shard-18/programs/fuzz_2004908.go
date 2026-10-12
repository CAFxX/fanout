package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int64; F1 uint; F2 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 int64; F1 int8 }
func genFn0(p0 string, p1 uint8) bool {
    p1 -= 2
    v4 := GenS1{F0: 127}
    var v5 uint8 = p1
    fmt.Println(v5, v5)
    if (282 < 42) {
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(v5, p0)
            v5 = 100
            _ = fi6
        }
        p1 -= (v5 % (((p1 - (v5 | 0))) | 1))
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (0 < 2)
}

func main() {
    v1 := []string{("a" + "foo bar"), ("hello" + (("hello" + "foo bar") + "\n\t")), "foo bar"}
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int16 = 446
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        switch fi9 {
        case 32767:
            if false {
                var v10 bool = true
                _ = v10
            } else {
                var v11 bool = true
                var v12 float64 = 0.1
                _ = v11
                _ = v12
            }
        case 9223372036854775807:
            var v13 bool = (!false)
            _ = v13
        default:
            if (!(321 != 9223372036854775807)) {
                var v14 uint32 = 2
                _ = v14
            }
        }
        var v15 int64 = 9223372036854775807
        _ = fi9
        _ = v15
    }
    switch "hello" {
    case "hello":
        var v16 uint = 1
        _ = v16
    case "":
        if (623 == 1) {
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 uint8 = 225
                var v19 uint8 = (v18 ^ (^(v18)))
                v20 := GenS2{F0: 418, F1: 0}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            v1 = append(v1, "a")
        } else {
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 int = fi21
                v22 = v22
                v23 := GenS3{F0: 1}
                _ = fi21
                _ = v22
                _ = v23
            }
            v24 := GenS3{F0: 2}
            _ = v24
        }
    case "\n\t":
        var v25 int32 = 0
        _ = v25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        if false {
            var v27 uint8 = 1
            _ = v27
        } else {
            var v28 string = ("hello" + "hello")
            for fi29 := 0; fi29 < 16; fi29++ {
                v30 := GenS4{F0: v28, F1: 2}
                v28 = ((v28 + v28) + v28)
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        for fi31 := 0; fi31 < 3; fi31++ {
            fmt.Println(fi31, fi31)
            fmt.Println(fi26, fi31)
            _ = len(v1)
            fmt.Println(fi26, fi31)
            var v32 float64 = 1.5
            v33 := GenS5{F0: 309, F1: 41, F2: 1}
            var v34 int32 = 2147483647
            v35 := GenS3{F0: 656}
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
            _ = v35
        }
        _ = fi26
    }
    var v36 int = -9223372036854775808
    v36 -= (v36 % ((1) | 1))
    var v37 uint = 18446744073709551615
    var v38 uint64 = 4294967295
    v36 += v36
    fmt.Println(v37)
    var v39 float64 = 3.14159
    var v40 float32 = 1e10
    if (251 <= 386) {
        v39 += ((float64(0.0) + 1.5) - (v39 * v39))
    } else {
        v1 = append(v1, "foo bar")
        for fi41 := 0; fi41 < 16; fi41++ {
            v3[int(v37)] = ("" + "")
            fmt.Println(v38, fi41, v36)
            fmt.Println(v38, fi41, v36)
            _ = fi41
        }
    }
    switch ("" + "foo bar") {
    case "a":
        fmt.Println(v37, v39)
    }
    var v42 uint16 = 321
    if false {
        for fi43 := 0; fi43 < 1; fi43++ {
            for fi44 := 0; fi44 < 10; fi44++ {
                v40 = float32(0.0)
                var v45 float32 = 3.14159
                _ = fi44
                _ = v45
            }
            _ = fi43
        }
        if true {
            v42 -= v42
            fmt.Println(v36, v38, v39)
            v3[(v36 << (((v37 ^ uint(v38))) & 63))] = ("foo bar" + "foo bar")
        } else {
            switch ("foo bar" + ("" + "hello")) {
            case "a":
                for fi46 := 0; fi46 < 17; fi46++ {
                    if (!(9223372036854775807 <= 1)) {
                        switch fi46 {
                        case 32768:
                            _ = len(v2)
                        case -1:
                            for fi47 := 0; fi47 < 20; fi47++ {
                                v48 := v3[fi46]
                                _ = fi47
                                _ = v48
                            }
                        case -219:
                            var v49 float32 = (-2.25 * (float32(-2.25) - (v40 - (3.14159 / ((-2.25 - (v40 - v40)) - v40)))))
                            _ = v49
                        default:
                            fmt.Println(v37)
                        }
                        var v50 int32 = -2147483648
                        switch "\n\t" {
                        case "foo bar":
                            var v51 int64 = 9223372036854775807
                            _ = v51
                        case "hello":
                            for fi52 := 0; fi52 < 9; fi52++ {
                                if false {
                                    v38 = (v38 ^ (v38 & 4294967295))
                                    var v53 uint64 = v38
                                    v54 := GenS6{F0: 6}
                                    _ = v53
                                    _ = v54
                                } else {
                                    fmt.Println(v37)
                                    var v55 int16 = 32767
                                    switch -2147483648 {
                                    case -2147483648:
                                        fmt.Println(v55)
                                    case 42:
                                        switch "foo bar" {
                                        case "hello":
                                            v55 -= (v55 * v55)
                                            v56 := GenS7{F0: -374, F1: ((fi52 != v36) && true)}
                                            _ = v56
                                        case "foo bar":
                                            fmt.Println(v40, v40)
                                        case "\n\t":
                                            if (!false) {
                                                v55 += 325
                                            }
                                        }
                                    case 0:
                                        switch fi52 {
                                        case 9223372036854775807:
                                            var v57 bool = ("" > (("foo bar" + ("hello" + "hello")) + (("hello" + "") + "")))
                                            _ = v57
                                        case 692:
                                            v58 := GenS8{F0: 0, F1: 0}
                                            var v59 int32 = -207
                                            _ = v58
                                            _ = v59
                                        case 947:
                                            var v60 uint16 = v42
                                            _ = v60
                                        default:
                                            v39 += v39
                                        }
                                    default:
                                        v36 = fi52
                                    }
                                    _ = v55
                                }
                                fmt.Println(v39, v50, v50)
                                _ = fi52
                            }
                        }
                        _ = v50
                    }
                    switch fi46 {
                    case 9223372036854775807:
                        if ((v40 - 0.0) > v40) {
                            v38 += (v38 / ((8) | 1))
                            fmt.Println(v36, v40, fi46)
                        }
                    default:
                        var v61 bool = false
                        _ = v61
                    }
                    _ = fi46
                }
            case "hello":
                var v62 float32 = v40
                _ = v62
            }
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        }
        v38 = v38
    } else {
        v39 += 0.1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v42
}
