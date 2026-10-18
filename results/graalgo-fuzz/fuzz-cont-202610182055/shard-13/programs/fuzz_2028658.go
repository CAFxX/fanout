package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint32; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 string; F2 uint64 }
type GenS6 struct { F0 uint8; F1 int32; F2 int64 }
func genFn0(p0 string, p1 string) uint32 {
    var v6 int16 = -32768
    v7 := []int{32768, 413, 4294967295}
    v7 = append(v7, 32767)
    fmt.Println(v6, p0)
    switch 0 {
    case 9223372036854775807:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case 0:
        var v8 bool = (2147483647 == -80)
        _ = v8
    default:
        _ = len(v7)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int32, p1 uint32, p2 uint) uint64 {
    var v9 int = -2147483648
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: 815, F1: 42}
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 382
    if (-1 > 1) {
        v13 := []int{0}
        _ = v13
    } else {
        var v14 int16 = 1
        switch (v9 + (v9 ^ v9)) {
        case 0:
            var v15 int32 = -2147483648
            _ = v15
        }
        fmt.Println(p2)
        _ = v14
    }
    switch (v9 & (^(v9))) {
    case -126:
        if (uint64(v12) >= 2) {
            fmt.Println(v9, p1)
            var v16 int16 = -1
            v16 = v16
            _ = v16
        } else {
            switch "\n\t" {
            case "hello":
                for fi17 := 0; fi17 < 3; fi17++ {
                    fmt.Println(v12)
                    for fi18 := 0; fi18 < 7; fi18++ {
                        if false {
                            if (("\n\t" + "hello") != ("hello" + "a")) {
                                p1 = p1
                                switch "" {
                                case "":
                                    v19 := []int{-746}
                                    _ = v19
                                }
                                switch v9 {
                                case 32767:
                                    if ((true && (2 == 456)) || (false || false)) {
                                        v20 := []int{(-(fi17))}
                                        var v21 float32 = (3.14159 * 0.0)
                                        _ = v20
                                        _ = v21
                                    } else {
                                        fmt.Println(fi18, fi17, v9)
                                        v9 += fi17
                                        var v22 string = "hello"
                                        _ = v22
                                    }
                                case 128:
                                    fmt.Println(fi17, fi17)
                                case -696:
                                    if (false && true) {
                                        var v23 int16 = 800
                                        if (1 == 224) {
                                            switch (fi17 - fi18) {
                                            case 32768:
                                                switch v9 {
                                                case 327:
                                                    switch (fi17 & (fi18 - 42)) {
                                                    case 32768:
                                                        switch (fi17 | 32768) {
                                                        case -229:
                                                            var v24 int64 = -9223372036854775808
                                                            _ = v24
                                                        case 127:
                                                            var v25 int16 = 1000
                                                            _ = v25
                                                        }
                                                    }
                                                case -1:
                                                    if true {
                                                        v23 += (v23 & (v23 ^ (v23 + 42)))
                                                        if true {
                                                            var v26 uint16 = (v12 | (v12 + 1))
                                                            _ = v26
                                                        } else {
                                                            switch (v9 ^ (fi17 << ((p2) & 63))) {
                                                            case -851:
                                                                p1 = genFn0("hello", "\n\t")
                                                                v27 := []int{}
                                                                _ = v27
                                                            case -765:
                                                                switch 9223372036854775807 {
                                                                case 689:
                                                                    v28 := GenS2{F0: v9}
                                                                    _ = v28
                                                                }
                                                            case 42:
                                                                v29 := []int{(fi18 >> ((uint(v12)) & 63)), v9, int(p2)}
                                                                _ = v29
                                                            }
                                                            var v30 int8 = -1
                                                            v31 := []int{int(v12), (fi17 | (^(fi17)))}
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        var v32 float32 = 3.14159
                                                        _ = v32
                                                    }
                                                default:
                                                    if ((p1 / ((p1) | 1)) == (p1 & p1)) {
                                                        var v33 uint32 = p1
                                                        _ = v33
                                                    }
                                                }
                                            case 1:
                                                var v34 float64 = (0.0 * float64(float32((-2.25 * 0.0))))
                                                _ = v34
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                var v35 string = ("\n\t" + "hello")
                                                _ = v35
                                            case "a":
                                                switch ("a" + ("hello" + "")) {
                                                case "\n\t":
                                                    fmt.Println(p2, p2)
                                                case "a":
                                                    v36 := []int{v9, fi17}
                                                    _ = v36
                                                }
                                            case "hello":
                                                v9 = 434
                                            }
                                        }
                                        fmt.Println(p0, v9, p1)
                                        v37 := []int{(v9 | (v9 / ((fi17) | 1))), (fi18 * v9), (fi18 ^ 32767), (v9 + fi17)}
                                        var v38 int = (fi17 / (((fi18 ^ (v9 % (((fi18 << ((18446744073709551615) & 63))) | 1)))) | 1))
                                        fmt.Println(v38)
                                        _ = v23
                                        _ = v37
                                        _ = v38
                                    } else {
                                        var v39 int = (v9 >> ((p2) & 63))
                                        _ = v39
                                    }
                                }
                            }
                            v40 := GenS2{F0: (v9 >> (((p2 / ((p2) | 1))) & 63))}
                            _ = v40
                        }
                        _ = fi18
                    }
                    _ = fi17
                }
            case "":
                v41 := GenS3{F0: 126, F1: 65535, F2: int(p1)}
                _ = v41
            case "\n\t":
                var v42 int = v9
                _ = v42
            }
        }
    case 927:
        for fi43 := 0; fi43 < 6; fi43++ {
            var v44 int16 = -369
            v45 := []int{4294967295, v9, (v9 - (v9 / ((32768) | 1))), fi43}
            _ = fi43
            _ = v44
            _ = v45
        }
    default:
        v46 := GenS4{F0: v12}
        _ = v46
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 0.0
    v2 := []int{-1, 824, -9223372036854775808}
    var v3 int32 = 2147483647
    v4 := []int{0, -9223372036854775808, -9223372036854775808}
    var v5 float64 = 0.1
    if false {
        v2 = append(v2, 128)
        switch "a" {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            for fi47 := 0; fi47 < 8; fi47++ {
                v48 := GenS2{F0: fi47}
                _ = fi47
                _ = v48
            }
        }
        fmt.Println(v3, v3)
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi49 := 0; fi49 < 6; fi49++ {
            if ("foo bar" == "a") {
                fmt.Println(fi49, v5, fi49)
                v5 += 0.1
            } else {
                for fi50 := 0; fi50 < 5; fi50++ {
                    v1 = v1
                    if true {
                        var v51 uint64 = 42
                        if ((true && (127 < 1)) && ((!false) || (!true))) {
                            var v52 bool = false
                            if (347 != -9223372036854775808) {
                                var v53 uint32 = 1
                                if (v1 >= (1e10 / float32((v5 - (v5 + v5))))) {
                                    var v54 float64 = ((v5 + 1e10) + v5)
                                    _ = v54
                                } else {
                                    v3 = (^(v3))
                                    v1 += (v1 + v1)
                                    switch int(v51) {
                                    case 127:
                                        switch (fi50 | -554) {
                                        case 750:
                                            var v55 uint64 = v51
                                            _ = v55
                                        default:
                                            var v56 int64 = 0
                                            _ = v56
                                        }
                                    case 9223372036854775807:
                                        if v52 {
                                            v57 := GenS2{F0: (fi50 - fi50)}
                                            v2 = append(v2, fi49)
                                            if (!(v52 || (v53 <= (v53 | (^(v53)))))) {
                                                var v58 int8 = 79
                                                if ((-2.25 - v1) >= -2.25) {
                                                    fmt.Println(v51, v53, v58)
                                                    var v59 int8 = 1
                                                    _ = v59
                                                }
                                                _ = v58
                                            }
                                            _ = v57
                                        }
                                    }
                                }
                                _ = v53
                            }
                            _ = v52
                        } else {
                            var v60 uint16 = 2
                            v3 = 2147483647
                            _ = v60
                        }
                        var v61 bool = true
                        _ = v51
                        _ = v61
                    }
                    v3 -= (v3 | v3)
                    _ = fi50
                }
                var v62 int = (fi49 % (((fi49 * fi49)) | 1))
                fmt.Println(v5, fi49, v1)
                _ = v62
            }
            _ = fi49
        }
        v63 := GenS4{F0: 0}
        _ = v63
    }
    if len(v4) > 0 {
        v4[0] = 41
    }
    v64 := GenS5{F0: -921, F1: "\n\t", F2: 0}
    switch 470 {
    case 4294967295:
        v65 := GenS6{F0: 1, F1: -2147483648, F2: 9223372036854775807}
        _ = v65
    case 32767:
        v3 += v3
    case 32768:
        v5 = v5
    }
    switch "" {
    case "\n\t":
        fmt.Println(v3, v5)
    case "a":
        fmt.Println(v5, v1)
    case "":
        if len(v4) > 0 {
            v4[0] = 32768
        }
    default:
        v2 = append(v2, 0)
    }
    if true {
        v1 = 1e10
        for fi66 := 0; fi66 < 19; fi66++ {
            for fi67 := 0; fi67 < 11; fi67++ {
                v5 -= (float64(((1e10 * v1) - 1e10)) + 0.1)
                _ = fi67
            }
            v3 = v3
            _ = fi66
        }
        switch 32767 {
        case 110:
            var v68 uint8 = 189
            _ = v68
        case 9223372036854775807:
            var v69 uint8 = 128
            _ = v69
        }
    }
    var v70 uint = 42
    if (506 >= 1) {
        v5 = v5
        var v71 string = (("foo bar" + "foo bar") + ("hello" + "\n\t"))
        var v72 float32 = v1
        _ = v71
        _ = v72
    } else {
        v1 += 1e10
    }
    fmt.Println(v70)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v70
}
