package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint64 }
func genFn0(p0 uint16, p1 int64) int {
    v5 := []int{}
    switch "\n\t" {
    case "hello":
        p1 = (^(p1))
    case "a":
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 float64 = 1.5
            for fi8 := 0; fi8 < 19; fi8++ {
                if true {
                    switch "foo bar" {
                    case "a":
                        p0 = (^(p0))
                    default:
                        p0 += 491
                    }
                    p0 -= p0
                    switch (fi6 % ((fi6) | 1)) {
                    case 890:
                        if true {
                            if (((1.5 * 0.1) * (3.14159 * 1.5)) < (1.5 * 0.0)) {
                                var v9 uint16 = (^(p0))
                                _ = len(v5)
                                var v10 uint64 = 4294967295
                                _ = v10
                                _ = v9
                            }
                        } else {
                            if len(v5) > 0 {
                                v5[0] = 128
                            }
                            switch (-(fi6)) {
                            case 2147483647:
                                if false {
                                    v7 += 3.14159
                                } else {
                                    v7 += 1e10
                                    var v11 uint = 560
                                    if false {
                                        switch fi6 {
                                        case 646:
                                            var v12 int64 = -1
                                            _ = v12
                                        case 128:
                                            v11 = (^(v11))
                                        case 405:
                                            var v13 int64 = int64(p0)
                                            if false {
                                                v14 := GenS1{F0: 441, F1: 2147483648, F2: 2}
                                                _ = v14
                                            }
                                            if true {
                                                var v15 int16 = 1000
                                                switch (-(fi6)) {
                                                case 822:
                                                    switch "foo bar" {
                                                    case "":
                                                        v7 = 1.5
                                                    case "foo bar":
                                                        fmt.Println(v11)
                                                    }
                                                }
                                                var v16 int16 = v15
                                                _ = v15
                                                _ = v16
                                            }
                                            v11 -= v11
                                            _ = v13
                                        }
                                        if (float64(1.5) < 1.5) {
                                            v17 := GenS2{F0: 100}
                                            p0 -= p0
                                            _ = v17
                                        } else {
                                            fmt.Println(v7)
                                        }
                                        fmt.Println(fi8, p1, p0)
                                    } else {
                                        p0 = 65535
                                        fmt.Println(p0)
                                    }
                                    _ = v11
                                }
                            case 42:
                                p1 = (p1 & -1)
                            default:
                                _ = len(v5)
                            }
                        }
                    case 1:
                        p0 += p0
                    case -52:
                        switch fi8 {
                        case 127:
                            var v18 string = "foo bar"
                            _ = v18
                        case 2147483647:
                            var v19 int = fi8
                            _ = v19
                        case 32768:
                            var v20 int8 = -109
                            _ = v20
                        default:
                            switch "hello" {
                            case "a":
                                v21 := GenS1{F0: 42, F1: 4294967295, F2: 4294967295}
                                _ = v21
                            case "foo bar":
                                v7 = float64((0.0 * 0.0))
                            case "":
                                v7 = 3.14159
                            default:
                                p1 = (-(p1))
                            }
                            v22 := GenS3{F0: 1000}
                            _ = v22
                        }
                    default:
                        p1 = (p1 | (p1 / ((int64(fi8)) | 1)))
                    }
                    p1 = (p1 ^ 0)
                } else {
                    var v23 uint8 = 0
                    fmt.Println(fi8, fi8)
                    var v24 uint64 = 665
                    v23 += v23
                    _ = v23
                    _ = v24
                }
                fmt.Println(fi6)
                var v25 int64 = p1
                var v26 float32 = 0.1
                fmt.Println(p0)
                _ = fi8
                _ = v25
                _ = v26
            }
            _ = fi6
            _ = v7
        }
    }
    var v27 int32 = 0
    var v28 float32 = 0.1
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    _ = v5
    return 4294967295
}
func genFn1() uint64 {
    v29 := []int{genFn0(302, -207)}
    var v30 float32 = 0.1
    v30 = (0.1 + (0.1 * 1.5))
    v31 := GenS4{F0: 18446744073709551615, F1: -1}
    _ = v29
    _ = v30
    _ = v31
    return 18446744073709551615
}

func main() {
    var v1 float64 = (float64(3.14159) / 3.14159)
    var v2 uint32 = 980
    var v3 uint8 = 128
    var v4 int16 = -446
    for fi32 := 0; fi32 < 13; fi32++ {
        if (789 <= 0) {
            fmt.Println(v3)
            v33 := []int{genFn0(487, 105), (fi32 + (fi32 >> ((65535) & 63))), fi32}
            v1 += v1
            _ = v33
        } else {
            v34 := GenS5{F0: 32767, F1: v3, F2: genFn1()}
            var v35 int64 = 9223372036854775807
            _ = v34
            _ = v35
        }
        v36 := []int{genFn0(255, int64(v3)), (-(fi32)), fi32}
        _ = fi32
        _ = v36
    }
    fmt.Println(v2, v4, v1)
    var v37 string = ""
    if true {
        fmt.Println(v1)
        v38 := []int{127}
        fmt.Println(v3, v1)
        _ = v38
    } else {
        if (1.5 > (1e10 / 3.14159)) {
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 int32 = 1
                v41 := []int{(fi39 - -846)}
                _ = fi39
                _ = v40
                _ = v41
            }
        } else {
            var v42 int16 = (v4 - -32768)
            var v43 uint32 = 65535
            _ = v42
            _ = v43
        }
    }
    v4 = v4
    v44 := GenS1{F0: 0, F1: (v2 >> ((4294967295) & 63)), F2: v2}
    v4 = v4
    if false {
        var v45 int8 = 0
        var v46 uint16 = 933
        v47 := []int{-256, 32767}
        _ = v45
        _ = v46
        _ = v47
    } else {
        for fi48 := 0; fi48 < 19; fi48++ {
            v49 := []int{}
            _ = fi48
            _ = v49
        }
        for fi50 := 0; fi50 < 3; fi50++ {
            if (false || true) {
                for fi51 := 0; fi51 < 6; fi51++ {
                    if true {
                        v37 = (v37 + (v37 + v37))
                        v52 := []int{fi51, 9223372036854775807, 32767, fi50}
                        switch fi50 {
                        case 32768:
                            fmt.Println(fi50, fi51)
                        case -9223372036854775808:
                            v2 = v2
                        default:
                            var v53 uint64 = 18446744073709551615
                            _ = v53
                        }
                        _ = v52
                    } else {
                        var v54 uint32 = (^(v2))
                        _ = v54
                    }
                    v1 = v1
                    if ((fi50 * genFn0(65535, -449)) >= fi50) {
                        v4 = int16(v3)
                        v55 := []int{genFn0(32767, -1)}
                        var v56 float64 = ((v1 + 1e10) + -2.25)
                        _ = v55
                        _ = v56
                    } else {
                        v3 = 127
                        switch (fi51 / ((fi50) | 1)) {
                        case 42:
                            var v57 int8 = 36
                            _ = v57
                        case 127:
                            v58 := GenS4{F0: 432, F1: 9223372036854775807}
                            _ = v58
                        }
                        var v59 uint = 255
                        v60 := GenS3{F0: (v4 ^ 1000)}
                        _ = v59
                        _ = v60
                    }
                    _ = fi51
                }
            }
            v1 = 1.5
            v61 := GenS1{F0: 42, F1: (v2 << ((892) & 63)), F2: v2}
            _ = fi50
            _ = v61
        }
        var v62 float32 = 0.0
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v44
}
