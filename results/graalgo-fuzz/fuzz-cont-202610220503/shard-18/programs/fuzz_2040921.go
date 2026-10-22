package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint; F1 bool; F2 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8; F1 int32; F2 uint8 }
type GenS7 struct { F0 int16; F1 int16 }

func main() {
    var v1 string = "\n\t"
    var v2 int = -1
    v3 := make(map[uint32]int)
    var v4 uint = 2
    var v5 uint8 = 54
    v6 := make(map[int]bool)
    if (-2147483648 >= -2147483648) {
        if (0 == 0) {
            v7 := GenS1{F0: (v2 >> ((v4) & 63))}
            _ = v7
        }
        var v8 uint = v4
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v5 += v5
        _ = fi9
    }
    v1 = v1
    if true {
        var v10 uint8 = 42
        _ = v10
    } else {
        switch v1 {
        case "a":
            var v11 uint8 = (v5 & 225)
            _ = v11
        case "":
            for fi12 := 0; fi12 < 6; fi12++ {
                if false {
                    v3[1] = (-(fi12))
                    var v13 bool = ((!(1 != 32767)) || (91 < 4294967295))
                    var v14 int8 = -2
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS2{F0: 617}
                    for fi16 := 0; fi16 < 3; fi16++ {
                        var v17 uint8 = 222
                        var v18 uint = v4
                        var v19 bool = true
                        v20 := GenS1{F0: (fi16 | (^(fi12)))}
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 13; fi21++ {
                        var v22 int32 = 250
                        v4 -= 255
                        _ = fi21
                        _ = v22
                    }
                    _ = v15
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    if (432 <= 0) {
                        if true {
                            v24 := []int{-257, 32767, -2147483648}
                            var v25 float32 = float32(1e10)
                            _ = v24
                            _ = v25
                        }
                    } else {
                        v5 -= 42
                        fmt.Println(v2, v1, fi12)
                        var v26 float32 = -2.25
                        _ = v26
                    }
                    var v27 int8 = 100
                    _ = fi23
                    _ = v27
                }
                _ = fi12
            }
        case "hello":
            var v28 uint32 = 918
            _ = v28
        default:
            fmt.Println(v4, v2)
        }
        var v29 float32 = 1e10
        v4 = (^(v4))
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        fmt.Println(v2, v1, fi30)
        _ = fi30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 uint16 = 255
        _ = fi31
        _ = v32
    }
    var v33 uint32 = 4294967295
    v34 := []int{-811, v2, (v2 >> (((v4 % ((v4) | 1))) & 63))}
    if false {
        switch -2147483648 {
        case -2147483648:
            var v35 uint = 285
            _ = v35
        case 1:
            var v36 float32 = 3.14159
            _ = v36
        }
    } else {
        fmt.Println(v1, v1)
        switch -830 {
        case 1:
            for fi37 := 0; fi37 < 15; fi37++ {
                var v38 uint64 = 0
                v33 += 860
                v39 := GenS3{F0: int64(v5), F1: 0}
                _ = fi37
                _ = v38
                _ = v39
            }
        case -9223372036854775808:
            v5 = v5
        }
    }
    var v40 uint64 = 184
    v3[1] = v2
    if (!true) {
        if (v40 < v40) {
            v41 := GenS4{F0: v4, F1: false, F2: 42}
            var v42 int = (v2 & v2)
            _ = v41
            _ = v42
        }
        for fi43 := 0; fi43 < 2; fi43++ {
            for fi44 := 0; fi44 < 11; fi44++ {
                var v45 int = v2
                var v46 uint8 = 1
                _ = fi44
                _ = v45
                _ = v46
            }
            for fi47 := 0; fi47 < 14; fi47++ {
                if len(v34) >= 2 {
                    v34 = v34[1:len(v34)]
                }
                var v48 int8 = -62
                _ = fi47
                _ = v48
            }
            if (false && (v4 != v4)) {
                if ((v40 * v40) == (v40 % (((v40 * 2)) | 1))) {
                    var v49 uint16 = 42
                    _ = v49
                }
            }
            _ = fi43
        }
        var v50 float32 = -2.25
        _ = v50
    } else {
        var v51 bool = ((v33 >= v33) || (uint(v40) <= (v4 / ((65535) | 1))))
        var v52 uint32 = 1
        _ = v51
        _ = v52
    }
    for fi53 := 0; fi53 < 16; fi53++ {
        v6[v2] = true
        for fi54 := 0; fi54 < 3; fi54++ {
            if ((fi53 << ((18446744073709551615) & 63)) != (v2 / (((v2 << (((v4 << ((v4) & 63))) & 63))) | 1))) {
                var v55 uint8 = 128
                switch ("" + ((v1 + "") + v1)) {
                case "\n\t":
                    v56 := GenS1{F0: (fi53 - v2)}
                    _ = v56
                case "":
                    fmt.Println(v40, v5, fi53)
                case "a":
                    v57 := GenS5{F0: (^(v55))}
                    _ = v57
                }
                v33 = 4294967295
                _ = v55
            } else {
                v4 = 255
                if true {
                    if (v2 <= -494) {
                        v1 = v1
                        if true {
                            var v58 uint8 = v5
                            _ = v58
                        } else {
                            _ = len(v6)
                            v59 := v3[v33]
                            _ = v59
                        }
                        var v60 uint16 = 65535
                        var v61 float64 = (float64((1.5 * 1e10)) - (1.5 * float64(1.5)))
                        if ((!(((v61 + v61) < v61) || false)) && (v4 == v4)) {
                            v1 = "hello"
                        } else {
                            fmt.Println(v2, v60, v1)
                            v62 := GenS6{F0: 92, F1: -1, F2: 127}
                            var v63 float32 = 1e10
                            _ = v62
                            _ = v63
                        }
                        _ = v60
                        _ = v61
                    }
                }
                v64 := GenS7{F0: 0, F1: -1}
                _ = v64
            }
            var v65 int16 = 75
            v66 := GenS5{F0: v5}
            _ = len(v34)
            v6[486] = true
            if false {
                v3[(^(v33))] = int(v33)
            }
            switch (fi54 % ((-2147483648) | 1)) {
            case 84:
                var v67 int8 = 31
                _ = v67
            default:
                fmt.Println(v1, v2, v1)
            }
            var v68 bool = (v5 <= 39)
            v68 = v68
            if (v68 && v68) {
                var v69 uint32 = (^(v33))
                if ((v4 | 0) > 0) {
                    v70 := GenS2{F0: 9223372036854775807}
                    _ = v70
                }
                v71 := v6[-9223372036854775808]
                _ = v69
                _ = v71
            }
            v2 += (fi53 ^ (fi54 << ((uint(v33)) & 63)))
            _ = fi54
            _ = v65
            _ = v66
            _ = v68
        }
        _ = fi53
    }
    if len(v34) >= 2 {
        v34 = v34[1:len(v34)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
