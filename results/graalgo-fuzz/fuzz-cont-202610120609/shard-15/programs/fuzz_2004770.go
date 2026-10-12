package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint16; F1 uint8; F2 int64 }
func genFn0() int16 {
    var v5 int32 = 2147483647
    for fi6 := 0; fi6 < 17; fi6++ {
        if false {
            if ((fi6 > fi6) || true) {
                v5 = v5
                if false {
                    for fi7 := 0; fi7 < 5; fi7++ {
                        v8 := GenS1{F0: 1, F1: 32767, F2: 0}
                        _ = fi7
                        _ = v8
                    }
                } else {
                    var v9 int8 = 121
                    var v10 bool = false
                    _ = v10
                    _ = v9
                }
                v11 := GenS1{F0: 42, F1: -32768, F2: 20}
                _ = v11
            } else {
                fmt.Println(v5, fi6)
            }
        }
        v5 -= (^(v5))
        for fi12 := 0; fi12 < 4; fi12++ {
            v5 = (v5 + (^(v5)))
            _ = fi12
        }
        _ = fi6
    }
    switch (("hello" + ("\n\t" + "a")) + "\n\t") {
    case "":
        fmt.Println(v5)
    case "a":
        switch "a" {
        case "":
            var v13 bool = true
            _ = v13
        case "hello":
            var v14 uint = 577
            _ = v14
        default:
            var v15 int8 = 100
            _ = v15
        }
    case "hello":
        v16 := []int{127}
        _ = v16
    }
    _ = v5
    return -123
}
func genFn1(p0 uint64, p1 int64, p2 uint64) uint {
    var v17 uint64 = 2
    var v18 int32 = -454
    for fi19 := 0; fi19 < 5; fi19++ {
        switch fi19 {
        case 32768:
            v20 := GenS1{F0: -196, F1: -307, F2: 720}
            _ = v20
        case 1:
            v21 := []int{fi19, fi19}
            _ = v21
        default:
            fmt.Println(v17, p1)
        }
        if ("foo bar" != ("\n\t" + "\n\t")) {
            if true {
                p2 = (v17 % (((v17 * p0)) | 1))
                var v22 float64 = (0.1 + (1e10 / float64(3.14159)))
                _ = v22
            }
            v23 := []int{(fi19 * fi19)}
            _ = v23
        } else {
            var v24 uint8 = 127
            switch int(v18) {
            case 42:
                v18 += (v18 | v18)
            case 9223372036854775807:
                if false {
                    v24 = (^(v24))
                    for fi25 := 0; fi25 < 4; fi25++ {
                        v26 := []int{0, (fi25 >> ((18446744073709551615) & 63)), fi19}
                        switch (fi25 & -1) {
                        case -2147483648:
                            fmt.Println(v17, v17)
                        default:
                            _ = len(v26)
                        }
                        _ = fi25
                        _ = v26
                    }
                    v27 := []int{fi19}
                    _ = v27
                }
            case 32767:
                fmt.Println(v24, p1, v17)
            }
            v28 := []int{(fi19 + fi19), (fi19 % ((fi19) | 1))}
            _ = v24
            _ = v28
        }
        _ = fi19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 255
}

func main() {
    var v1 float64 = float64((3.14159 / 0.1))
    var v2 int = -2147483648
    v3 := make(map[uint32]int)
    var v4 int32 = 1
    var v29 int16 = -220
    v30 := GenS2{F0: (^(v29))}
    for fi31 := 0; fi31 < 19; fi31++ {
        if false {
            v32 := GenS2{F0: (-(v29))}
            if (-128 >= 42) {
                var v33 uint16 = 340
                if false {
                    for fi34 := 0; fi34 < 14; fi34++ {
                        switch ("foo bar" + "hello") {
                        case "hello":
                            v1 = v1
                        case "foo bar":
                            fmt.Println(v4)
                        case "":
                            delete(v3, 2147483647)
                        }
                        v29 += (v29 % (((v29 ^ v29)) | 1))
                        v4 -= -2147483648
                        _ = fi34
                    }
                    var v35 uint8 = 128
                    var v36 uint8 = v35
                    _ = v35
                    _ = v36
                } else {
                    var v37 uint32 = 2147483647
                    var v38 int32 = (v4 / ((-315) | 1))
                    _ = v37
                    _ = v38
                }
                var v39 int = v2
                _ = v33
                _ = v39
            }
            _ = v32
        } else {
            var v40 bool = (((float64(0.1) * 0.1) >= v1) || false)
            if (834 != 1) {
                v4 += v4
            } else {
                var v41 int = 42
                v41 = 32767
                _ = v41
            }
            for fi42 := 0; fi42 < 19; fi42++ {
                var v43 uint32 = 0
                _ = fi42
                _ = v43
            }
            _ = v40
        }
        _ = fi31
    }
    v2 = 9223372036854775807
    v44 := v3[4294967295]
    v44 = 9223372036854775807
    var v45 int8 = -1
    v46 := v3[0]
    for fi47 := 0; fi47 < 6; fi47++ {
        fmt.Println(v29, v46)
        _ = fi47
    }
    v1 += v1
    if true {
        v48 := GenS3{F0: -728}
        switch -190 {
        case -1:
            v2 += v2
        default:
            v49 := GenS4{F0: 42, F1: 180, F2: (v29 * (v29 * (v29 % ((-641) | 1))))}
            _ = v49
        }
        if false {
            switch v2 {
            case 262:
                v3[107] = (v46 ^ 127)
            }
        }
        _ = v48
    } else {
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            var v50 int = v2
            _ = v50
        case "a":
            if (0 > 472) {
                v51 := GenS5{F0: 128, F1: 255, F2: v29}
                v52 := v3[842]
                if (("" >= "hello") || (float64(float32((float64(float32(float64(0.1))) + 1e10))) > v1)) {
                    if ((int(v4) != 0) || false) {
                        v53 := []int{v2, -9223372036854775808, 665, (v46 >> ((4294967295) & 63))}
                        v54 := GenS4{F0: 729, F1: 468, F2: int16(v4)}
                        switch "a" {
                        case "\n\t":
                            v55 := GenS1{F0: -1, F1: (-(v29)), F2: 341}
                            _ = v55
                        }
                        _ = v53
                        _ = v54
                    } else {
                        var v56 int64 = int64(v4)
                        _ = v56
                    }
                    if (18446744073709551615 < 679) {
                        fmt.Println(v46)
                        if false {
                            var v57 int16 = v29
                            fmt.Println(v46, v29)
                            _ = v57
                        }
                    } else {
                        switch (^(v44)) {
                        case 32767:
                            for fi58 := 0; fi58 < 7; fi58++ {
                                var v59 uint = 403
                                if (!true) {
                                    v4 = v4
                                } else {
                                    fmt.Println(v59)
                                }
                                switch ("hello" + ("foo bar" + "")) {
                                case "":
                                    v60 := GenS6{F0: 263, F1: 255, F2: 42}
                                    _ = v60
                                case "hello":
                                    fmt.Println(v46, v4)
                                default:
                                    fmt.Println(v46, v4)
                                }
                                _ = fi58
                                _ = v59
                            }
                        }
                    }
                } else {
                    fmt.Println(v2, v4, v52)
                    for fi61 := 0; fi61 < 19; fi61++ {
                        v62 := GenS2{F0: (v29 | v29)}
                        v44 -= -2147483648
                        _ = fi61
                        _ = v62
                    }
                }
                _ = v51
                _ = v52
            } else {
                var v63 int16 = 1000
                var v64 uint16 = 32767
                _ = v63
                _ = v64
            }
        case "\n\t":
            v45 = (v45 | (v45 - -50))
        default:
            var v65 uint = 42
            _ = v65
        }
        v4 = (v4 + (^(v4)))
        v66 := []int{}
        _ = v66
    }
    for fi67 := 0; fi67 < 17; fi67++ {
        v44 -= v44
        fmt.Println(v4, v46)
        _ = fi67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
