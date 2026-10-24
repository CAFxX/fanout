package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint; F1 int16; F2 int32 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 int; F1 int32; F2 int32 }
func genFn0(p0 int64) uint8 {
    v5 := GenS1{F0: 42}
    switch "\n\t" {
    case "hello":
        switch 32767 {
        case 51:
            var v6 uint8 = 155
            _ = v6
        }
    case "foo bar":
        switch 32768 {
        case 0:
            for fi7 := 0; fi7 < 20; fi7++ {
                v8 := GenS1{F0: 42}
                _ = fi7
                _ = v8
            }
        case 32768:
            p0 = -719
        }
    case "":
        v9 := GenS2{F0: 0}
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(p0, fi10)
        var v11 int16 = 1000
        for fi12 := 0; fi12 < 4; fi12++ {
            switch "\n\t" {
            case "":
                v13 := []int{-2147483648, 32768}
                _ = v13
            }
            var v14 uint = 255
            if false {
                v14 -= v14
                v15 := GenS3{F0: 68, F1: 937}
                _ = v15
            }
            _ = fi12
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        if ((fi16 + (fi16 & -860)) == (^(fi16))) {
            for fi17 := 0; fi17 < 19; fi17++ {
                p0 = p0
                v18 := GenS4{F0: 2, F1: 184, F2: 1}
                p0 = 0
                _ = fi17
                _ = v18
            }
            p0 -= (p0 - p0)
            if true {
                var v19 float32 = 1.5
                _ = v19
            } else {
                var v20 int32 = -492
                _ = v20
            }
        } else {
            for fi21 := 0; fi21 < 2; fi21++ {
                switch fi21 {
                case 127:
                    var v22 int32 = 42
                    _ = v22
                case -240:
                    var v23 int8 = -128
                    _ = v23
                default:
                    v24 := GenS3{F0: 45, F1: 556}
                    _ = v24
                }
                _ = fi21
            }
        }
        switch "\n\t" {
        case "a":
            v25 := GenS4{F0: 2, F1: 1, F2: -32768}
            _ = v25
        case "foo bar":
            var v26 uint64 = 0
            _ = v26
        default:
            var v27 int = (fi16 / (((fi16 << ((0) & 63))) | 1))
            _ = v27
        }
        _ = fi16
    }
    _ = p0
    _ = v5
    return 127
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -509
    var v3 uint = 18446744073709551615
    var v4 int64 = 386
    v1[42] = int(v2)
    v2 += v2
    v4 += v4
    v28 := GenS2{F0: 2}
    var v29 int32 = 2147483647
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 int16 = 42
        if ((v29 | (v29 - v29)) <= (v29 ^ 0)) {
            v2 = (v4 + 1)
            v32 := GenS1{F0: 1}
            v4 -= v2
            _ = v32
        } else {
            if true {
                var v33 uint16 = 0
                switch ("\n\t" + "a") {
                case "":
                    v34 := GenS5{F0: v3, F1: (v31 - 1000), F2: (v29 + 763)}
                    _ = v34
                case "a":
                    if ((v3 - 255) > v3) {
                        v2 = -9223372036854775808
                        var v35 uint32 = 0
                        for fi36 := 0; fi36 < 9; fi36++ {
                            var v37 float64 = (float64(float32(0.0)) + (3.14159 - -2.25))
                            v38 := GenS4{F0: v3, F1: 0, F2: 1}
                            fmt.Println(fi30, v29)
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = v35
                    } else {
                        for fi39 := 0; fi39 < 19; fi39++ {
                            var v40 float64 = -2.25
                            var v41 float32 = (1e10 - (0.1 * 0.1))
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                    }
                case "foo bar":
                    if (0.0 == 0.0) {
                        v3 += (v3 * v3)
                        v2 += v4
                    } else {
                        v29 -= (v29 % (((-(v29))) | 1))
                    }
                }
                _ = v33
            } else {
                if false {
                    v31 += 948
                    v42 := GenS6{F0: -1, F1: 0}
                    _ = v42
                } else {
                    if (false || false) {
                        fmt.Println(fi30, v31, v29)
                        var v43 int16 = (v31 << (((v3 ^ v3)) & 63))
                        switch fi30 {
                        case -538:
                            for fi44 := 0; fi44 < 9; fi44++ {
                                var v45 int64 = 9223372036854775807
                                _ = fi44
                                _ = v45
                            }
                        case 0:
                            v2 += -623
                        default:
                            var v46 int8 = -128
                            _ = v46
                        }
                        _ = v43
                    }
                    fmt.Println(v3)
                }
            }
            v1[-404] = fi30
            for fi47 := 0; fi47 < 3; fi47++ {
                var v48 uint8 = 127
                v48 += (v48 / ((255) | 1))
                v3 = uint(v48)
                _ = fi47
                _ = v48
            }
        }
        if true {
            v49 := GenS7{F0: 9223372036854775807, F1: 129, F2: (v29 | 42)}
            _ = v49
        } else {
            var v50 bool = (4294967295 <= 2147483648)
            _ = v50
        }
        _ = fi30
        _ = v31
    }
    v51 := GenS2{F0: 4294967295}
    v52 := GenS6{F0: 95, F1: 463}
    if (true && true) {
        var v53 uint32 = 431
        _ = v53
    }
    switch 32768 {
    case -3:
        v54 := []int{147}
        _ = v54
    case 42:
        var v55 uint16 = 1
        _ = v55
    case 9223372036854775807:
        if (127 >= 42) {
            v56 := []int{}
            fmt.Println(v29)
            _ = v56
        }
    default:
        var v57 int = 32768
        _ = v57
    }
    if false {
        switch ("foo bar" + "foo bar") {
        case "a":
            var v58 string = ""
            _ = v58
        case "hello":
            _ = len(v1)
        case "":
            var v59 int32 = (^(v29))
            _ = v59
        default:
            v4 += v2
        }
        v1[-9223372036854775808] = int(v29)
    } else {
        var v60 uint64 = uint64(v3)
        _ = v60
    }
    var v61 uint32 = 390
    v1[9223372036854775807] = int(v29)
    var v62 uint8 = 2
    v62 = (v62 | v62)
    v2 -= (v4 | 9223372036854775807)
    var v63 uint = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v61
    _ = v62
    _ = v63
}
