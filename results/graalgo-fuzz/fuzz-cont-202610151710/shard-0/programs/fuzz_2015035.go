package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16; F1 int8; F2 int }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 uint }
func genFn0(p0 bool, p1 uint32) bool {
    if (-823 <= -1) {
        var v7 uint64 = 2
        p1 += p1
        _ = v7
    } else {
        v8 := []int{0, -492, 439}
        _ = v8
    }
    v9 := []int{2147483647, 32767}
    switch "hello" {
    case "":
        p1 -= p1
    case "a":
        for fi10 := 0; fi10 < 1; fi10++ {
            switch "foo bar" {
            case "a":
                p0 = ((255 > 94) && p0)
            case "foo bar":
                v9 = append(v9, fi10)
            case "":
                var v11 int64 = 9223372036854775807
                _ = v11
            }
            _ = fi10
        }
    case "\n\t":
        if len(v9) > 0 {
            v9[0] = 32768
        }
    default:
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint16 = 32767
            v9 = append(v9, 0)
            p0 = true
            _ = fi12
            _ = v13
        }
    }
    if (!(true || (-1 <= 2147483647))) {
        var v14 int16 = -32768
        switch (((("\n\t" + ("hello" + "foo bar")) + "hello") + "\n\t") + "") {
        case "hello":
            v15 := GenS1{F0: (-964 != -270), F1: 612}
            _ = v15
        case "":
            var v16 int8 = -128
            _ = v16
        case "\n\t":
            switch "\n\t" {
            case "a":
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS1{F0: (true && (!(185 <= 2))), F1: v14}
                    for fi19 := 0; fi19 < 6; fi19++ {
                        var v20 float32 = 0.1
                        _ = fi19
                        _ = v20
                    }
                    _ = fi17
                    _ = v18
                }
            default:
                p1 = p1
            }
        }
        p0 = (v14 > (v14 << ((2) & 63)))
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v9
    return p0
}
func genFn1(p0 int64) int16 {
    if false {
        v21 := []int{127, -2147483648}
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int64 = p0
            v24 := GenS1{F0: (true || false), F1: 460}
            if len(v21) > 0 {
                v21[0] = -1
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        p0 = (p0 | (p0 / (((p0 ^ p0)) | 1)))
        _ = v21
    }
    var v25 float64 = float64(1e10)
    if true {
        var v26 string = ""
        v26 = v26
        var v27 string = "hello"
        _ = v26
        _ = v27
    } else {
        v28 := GenS1{F0: true, F1: 1000}
        _ = v28
    }
    var v29 uint64 = 1
    for fi30 := 0; fi30 < 4; fi30++ {
        v31 := []int{(^(fi30)), (fi30 >> ((172) & 63)), -999, (fi30 / ((fi30) | 1))}
        _ = fi30
        _ = v31
    }
    _ = p0
    _ = v25
    _ = v29
    return 32767
}

func main() {
    var v1 int16 = -618
    v2 := []int{}
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 1
    for fi32 := 0; fi32 < 8; fi32++ {
        if true {
            switch (fi32 % (((fi32 / ((fi32) | 1))) | 1)) {
            case 42:
                fmt.Println(v6, v6, fi32)
            case -2147483648:
                var v33 float64 = 0.1
                _ = v33
            }
            v1 = (^(v1))
        } else {
            for fi34 := 0; fi34 < 19; fi34++ {
                v2 = append(v2, fi32)
                v3[fi34] = "hello"
                _ = fi34
            }
            for fi35 := 0; fi35 < 2; fi35++ {
                switch "hello" {
                case "hello":
                    switch (^(fi32)) {
                    case 32767:
                        if (56 != 2) {
                            if true {
                                var v36 float32 = 3.14159
                                _ = v36
                            } else {
                                v37 := GenS2{F0: -1, F1: -59, F2: (fi35 + fi32)}
                                var v38 uint8 = 127
                                fmt.Println(fi35)
                                _ = v37
                                _ = v38
                            }
                            v1 = (-(v1))
                        } else {
                            v39 := GenS2{F0: v1, F1: -1, F2: (fi35 & -1)}
                            _ = v39
                        }
                    case -534:
                        v6 = 0
                    default:
                        v40 := GenS2{F0: 0, F1: 20, F2: fi35}
                        _ = v40
                    }
                default:
                    v6 = 32767
                }
                v6 = 42
                if (!true) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v41 uint32 = 878
                    switch "foo bar" {
                    case "\n\t":
                        v42 := GenS2{F0: genFn1(-984), F1: -1, F2: (fi35 | fi35)}
                        _ = v42
                    case "foo bar":
                        v43 := GenS3{F0: 4294967295, F1: 769, F2: 547}
                        _ = v43
                    default:
                        v5[2] = 30
                    }
                    _ = v41
                }
                _ = fi35
            }
            v6 += (v6 + (^(v6)))
        }
        switch (("foo bar" + "") + "hello") {
        case "a":
            if true {
                fmt.Println(v6, v1, v1)
            }
        default:
            v4[fi32] = 32768
        }
        _ = fi32
    }
    v44 := GenS2{F0: 32767, F1: 0, F2: 42}
    for fi45 := 0; fi45 < 5; fi45++ {
        switch fi45 {
        case 42:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            fmt.Println(fi45)
        }
        for fi46 := 0; fi46 < 3; fi46++ {
            v6 -= v6
            v3[fi46] = ""
            _ = fi46
        }
        _ = fi45
    }
    var v47 int16 = v1
    v48 := GenS4{F0: 65535}
    v4[32767] = -2147483648
    var v49 float64 = 1e10
    fmt.Println(v47, v49)
    if true {
        v49 -= 1e10
        var v50 int = 127
        var v51 bool = true
        _ = v50
        _ = v51
    }
    var v52 bool = true
    switch "\n\t" {
    case "":
        if ("a" <= "\n\t") {
            for fi53 := 0; fi53 < 6; fi53++ {
                if (1000 != (v1 * genFn1(-9223372036854775808))) {
                    v49 += v49
                }
                v54 := GenS5{F0: 0, F1: 2}
                if v52 {
                    if (3.14159 != 0.0) {
                        _ = len(v2)
                    }
                    v49 += v49
                } else {
                    fmt.Println(v1, v49, fi53)
                }
                _ = fi53
                _ = v54
            }
            if ((129 != 171) && v52) {
                v2 = append(v2, -2147483648)
                v55 := GenS3{F0: 18446744073709551615, F1: 2, F2: 45}
                _ = v55
            }
        } else {
            var v56 float32 = 1e10
            v57 := GenS2{F0: v47, F1: 118, F2: 4294967295}
            _ = v56
            _ = v57
        }
    }
    for fi58 := 0; fi58 < 14; fi58++ {
        for fi59 := 0; fi59 < 8; fi59++ {
            if (0.1 > 1.5) {
                var v60 uint = 196
                var v61 int8 = -41
                v62 := GenS1{F0: v52, F1: int16(fi58)}
                _ = v60
                _ = v61
                _ = v62
            }
            v52 = false
            _ = fi59
        }
        for fi63 := 0; fi63 < 2; fi63++ {
            var v64 int32 = 720
            _ = fi63
            _ = v64
        }
        _ = fi58
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v6
}
