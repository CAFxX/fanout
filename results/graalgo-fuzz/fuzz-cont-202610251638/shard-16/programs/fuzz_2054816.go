package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 int; F1 int; F2 int64 }
type GenS3 struct { F0 bool; F1 int16; F2 int }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint16 }
func genFn0() uint16 {
    var v4 int16 = 32767
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := []int{(fi6 - -238), fi5, -1}
            v4 -= v4
            switch fi6 {
            case 127:
                var v8 uint64 = 0
                _ = v8
            case -141:
                v4 = v4
            default:
                _ = len(v7)
            }
            v4 += 688
            _ = fi6
            _ = v7
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 float64 = 1e10
            switch (fi5 << ((18446744073709551615) & 63)) {
            case 324:
                v11 := GenS1{F0: 4294967295, F1: (fi9 ^ (fi9 / ((fi9) | 1))), F2: 2147483647}
                _ = v11
            case 576:
                if (false || (fi5 < (fi5 / ((fi5) | 1)))) {
                    v12 := []int{(fi9 ^ fi5), (fi9 << ((65535) & 63))}
                    var v13 int = fi9
                    _ = v12
                    _ = v13
                } else {
                    v14 := []int{fi9, 4294967295, fi9, -9223372036854775808}
                    _ = v14
                }
            }
            var v15 uint16 = 2
            _ = fi9
            _ = v10
            _ = v15
        }
        _ = fi5
    }
    _ = v4
    return 1
}
func genFn1(p0 int64, p1 uint64, p2 int16) uint32 {
    p1 = 780
    fmt.Println(p0, p1)
    if (!(3.14159 == float32((0.1 + 0.1)))) {
        p2 += (p2 * -1)
    }
    p2 = (-(p2))
    var v16 uint64 = (p1 << ((65535) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 4294967295
    v3 := []string{("foo bar" + ("a" + "\n\t"))}
    fmt.Println(v2)
    if (-2.25 > 0.1) {
        var v17 uint = 65535
        _ = v17
    } else {
        v2 = 0
    }
    var v18 int64 = -9223372036854775808
    v18 -= (v18 | (v18 ^ (v18 / ((-81) | 1))))
    if false {
        v19 := GenS2{F0: 2147483647, F1: 956, F2: v18}
        v18 = v18
        if (("a" + "hello") != ("hello" + "foo bar")) {
            var v20 uint64 = 0
            switch "" {
            case "a":
                for fi21 := 0; fi21 < 20; fi21++ {
                    fmt.Println(v18)
                    fmt.Println(v2)
                    _ = fi21
                }
            }
            _ = v20
        }
        _ = v19
    } else {
        var v22 int = 508
        _ = v22
    }
    v1[2147483648] = ""
    v1[954] = ""
    v3 = append(v3, "foo bar")
    for fi23 := 0; fi23 < 11; fi23++ {
        v24 := GenS1{F0: 2, F1: 4294967295, F2: (-(fi23))}
        if true {
            v25 := GenS3{F0: true, F1: 32767, F2: (fi23 - (fi23 % ((4294967295) | 1)))}
            var v26 uint64 = 35
            fmt.Println(v2, fi23)
            _ = v25
            _ = v26
        } else {
            switch (("a" + "hello") + "\n\t") {
            case "foo bar":
                var v27 bool = ((1e10 < 1e10) || (0 > 0))
                _ = v27
            case "a":
                for fi28 := 0; fi28 < 18; fi28++ {
                    if (((float64(-2.25) / (-2.25 * (1.5 - (((1e10 * (float64(((((3.14159 + 3.14159) + (1e10 * (0.1 + float32(float64(1e10))))) - (float32(((-2.25 * (float64(float32((-2.25 + (((((-2.25 / 1.0) * (-2.25 * 0.0)) + 0.0) * 1.5) * 1.5)))) + 3.14159)) * 1e10)) / float32(-2.25))) * 0.0)) * (1.5 * (-2.25 * 0.1)))) - 1e10) + 1e10)))) > 1e10) && false) {
                        var v29 float32 = 1.5
                        _ = v29
                    } else {
                        switch (("a" + "\n\t") + "foo bar") {
                        case "a":
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            var v30 int32 = 2147483647
                            fmt.Println(v30)
                            v31 := GenS3{F0: false, F1: 1000, F2: 0}
                            _ = v30
                            _ = v31
                        case "hello":
                            if (!false) {
                                delete(v1, 2147483648)
                            }
                        }
                    }
                    switch "\n\t" {
                    case "foo bar":
                        if false {
                            v3 = append(v3, "\n\t")
                        }
                    default:
                        v18 += v18
                        switch fi23 {
                        case 32767:
                            if (!(!(265 > 452))) {
                                v2 = v2
                                fmt.Println(fi23)
                                var v32 bool = false
                                _ = v32
                            }
                        case -2147483648:
                            var v33 bool = (!false)
                            _ = v33
                        case 921:
                            var v34 int64 = (-(v18))
                            _ = v34
                        }
                        if false {
                            fmt.Println(v18, fi23, fi28)
                            var v35 int8 = 101
                            if ((1 != 3) || true) {
                                if (!(18446744073709551615 >= 42)) {
                                    v18 += v18
                                    v35 += (v35 | 127)
                                } else {
                                    v2 = (v2 + (v2 % ((v2) | 1)))
                                    delete(v1, genFn1((v18 / ((9223372036854775807) | 1)), 1, 0))
                                    var v36 uint8 = 255
                                    _ = v36
                                }
                                var v37 int32 = 42
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = "hello"
                                    }
                                    v18 = 1
                                    v35 -= (v35 / ((v35) | 1))
                                    if ((fi23 | -1) < (fi28 << ((715) & 63))) {
                                        v37 = v37
                                        v38 := GenS4{F0: 42, F1: fi23}
                                        switch fi23 {
                                        case -88:
                                            if (!(42 > 1)) {
                                                var v39 uint = (v2 & (v2 >> ((255) & 63)))
                                                _ = v39
                                            }
                                        default:
                                            fmt.Println(v18, v2, v18)
                                        }
                                        if ((0.1 + float32((0.0 / -2.25))) != (float32(0.1) - float32(float64(((0.1 / 1e10) - 0.1))))) {
                                            var v40 int = (fi23 | fi28)
                                            v41 := v1[2]
                                            v42 := v1[731]
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                        }
                                        _ = v38
                                    } else {
                                        switch (fi23 / ((fi28) | 1)) {
                                        case -9223372036854775808:
                                            v43 := GenS2{F0: 4294967295, F1: int(v2), F2: v18}
                                            _ = v43
                                        }
                                        _ = len(v1)
                                        v18 -= (v18 & (v18 & (v18 + (v18 ^ v18))))
                                    }
                                    v44 := GenS1{F0: 42, F1: fi23, F2: (-(fi28))}
                                    _ = v44
                                }
                                v37 = (v37 & (-(v37)))
                                if false {
                                    var v45 int8 = v35
                                    var v46 uint32 = 2147483647
                                    _ = v45
                                    _ = v46
                                }
                                _ = v37
                            }
                            v47 := GenS5{F0: 65535}
                            if (0 != 577) {
                                v18 = (^(v18))
                            }
                            _ = v35
                            _ = v47
                        } else {
                            var v48 float64 = ((((1e10 - (0.0 - 3.14159)) / (1e10 * 3.14159)) - ((-2.25 * ((3.14159 / (-2.25 / (1.5 / 3.14159))) + 0.1)) * (0.1 * (-2.25 - 0.1)))) * (1.5 + 0.1))
                            var v49 int64 = v18
                            v1[0] = "\n\t"
                            _ = len(v3)
                            fmt.Println(v2)
                            _ = v48
                            _ = v49
                        }
                    }
                    switch "foo bar" {
                    case "hello":
                        v2 = 65535
                    case "":
                        switch fi28 {
                        case -583:
                            v50 := GenS5{F0: 0}
                            _ = v50
                        }
                        fmt.Println(fi28, fi28, v18)
                        var v51 uint16 = 65535
                        fmt.Println(fi23, v2)
                        _ = v51
                    case "foo bar":
                        fmt.Println(v18, v18, fi28)
                    }
                    _ = fi28
                }
            default:
                var v52 string = ("a" + "foo bar")
                _ = v52
            }
            if (2 != 0) {
                v53 := v1[184]
                _ = v53
            }
        }
        _ = fi23
        _ = v24
    }
    if true {
        var v54 string = "hello"
        _ = v54
    }
    if true {
        v2 = v2
    }
    v2 = (v2 & (v2 >> ((802) & 63)))
    var v55 bool = (!false)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v55
}
