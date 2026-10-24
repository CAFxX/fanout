package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }
type GenS4 struct { F0 uint8 }
func genFn0() uint8 {
    var v6 uint8 = 0
    var v7 int16 = -1
    switch -657 {
    case 128:
        for fi8 := 0; fi8 < 11; fi8++ {
            v9 := GenS1{F0: 706}
            for fi10 := 0; fi10 < 15; fi10++ {
                fmt.Println(fi10, v7)
                v11 := []int{232, fi10, fi8}
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    case 987:
        v12 := []int{-9223372036854775808, -855, -850, 8}
        _ = v12
    case -9223372036854775808:
        v7 -= (v7 << ((0) & 63))
    default:
        v6 += v6
    }
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := []int{32767, 127, 0}
    var v2 uint64 = 42
    v3 := []uint32{2147483647, 217, 391}
    var v4 int16 = -1
    var v5 float64 = float64(((1e10 + 3.14159) - 1e10))
    var v13 uint = 18446744073709551615
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v5)
    var v14 uint64 = v2
    v3 = append(v3, 2147483648)
    v15 := GenS1{F0: 931}
    var v16 int32 = -1
    switch -9223372036854775808 {
    case -2147483648:
        switch (("foo bar" + "\n\t") + "hello") {
        case "":
            v2 = 358
        default:
            v17 := GenS1{F0: 546}
            _ = v17
        }
    case 128:
        v18 := GenS2{F0: 65535, F1: (v14 - v2)}
        _ = v18
    case -337:
        if (!true) {
            switch 88 {
            case -564:
                if ((!false) || (-877 > -608)) {
                    var v19 bool = false
                    v20 := GenS2{F0: 2, F1: v14}
                    if ((!false) && v19) {
                        fmt.Println(v19)
                        v3 = append(v3, 2147483648)
                        var v21 uint16 = 255
                        _ = v21
                    }
                    _ = v19
                    _ = v20
                }
            case 698:
                v22 := GenS1{F0: 2}
                _ = v22
            default:
                v4 -= v4
            }
        } else {
            if (true && true) {
                switch -9223372036854775808 {
                case 9223372036854775807:
                    v2 -= 42
                case -2147483648:
                    var v23 int64 = -9223372036854775808
                    _ = v23
                case 127:
                    fmt.Println(v16, v14)
                }
            }
            if false {
                var v24 uint32 = 2
                _ = v24
            }
        }
    }
    if true {
        var v25 uint16 = 32767
        _ = v25
    }
    if (!false) {
        switch 9223372036854775807 {
        case -1:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 1:
            if true {
                var v26 float64 = (1.5 * float64(1.5))
                fmt.Println(v2, v4)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v26
            } else {
                v27 := GenS1{F0: 810}
                var v28 int32 = (^(v16))
                for fi29 := 0; fi29 < 19; fi29++ {
                    var v30 float32 = 1.5
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            v14 -= 0
                            var v31 uint8 = 84
                            if false {
                                v31 = (^(v31))
                                fmt.Println(v5)
                                fmt.Println(v28, v14, v28)
                            }
                            _ = v31
                        } else {
                            if ((1 != (^(v28))) || ((3.14159 > v5) && true)) {
                                if (!(v5 <= 1.5)) {
                                    for fi32 := 0; fi32 < 2; fi32++ {
                                        v3 = append(v3, 2147483648)
                                        _ = fi32
                                    }
                                }
                            } else {
                                v5 = ((v5 * v5) / float64((0.1 - 0.0)))
                                switch "" {
                                case "hello":
                                    switch "a" {
                                    case "a":
                                        v13 += uint(v2)
                                    case "":
                                        var v33 uint16 = 32767
                                        _ = v33
                                    default:
                                        v2 += (v14 << (((v13 | v13)) & 63))
                                    }
                                case "":
                                    fmt.Println(v5, v16, v28)
                                }
                                v1 = append(v1, (-(fi29)))
                            }
                            for fi34 := 0; fi34 < 12; fi34++ {
                                v2 += v2
                                fmt.Println(v4, v28)
                                var v35 uint = (v13 ^ v13)
                                _ = fi34
                                _ = v35
                            }
                            v13 = v13
                        }
                    }
                    v36 := GenS3{F0: 36, F1: (fi29 | (fi29 - fi29)), F2: 18446744073709551615}
                    _ = fi29
                    _ = v30
                    _ = v36
                }
                _ = v27
                _ = v28
            }
        default:
            v37 := GenS3{F0: 42, F1: -9223372036854775808, F2: 4294967295}
            _ = v37
        }
        for fi38 := 0; fi38 < 8; fi38++ {
            if true {
                var v39 uint16 = 2
                fmt.Println(v13, v13, v5)
                if false {
                    if ((((false && (127 >= 255)) && (!((^(v39)) == (^(v39))))) && (false && (938 == -1))) || false) {
                        switch fi38 {
                        case 42:
                            v14 = (v14 << ((v13) & 63))
                        case 2147483647:
                            var v40 uint16 = (v39 & (v39 >> ((0) & 63)))
                            _ = v40
                        case -661:
                            var v41 int32 = (v16 & (v16 % ((v16) | 1)))
                            _ = v41
                        }
                        v5 = (0.1 * (v5 * 3.14159))
                    } else {
                        if false {
                            fmt.Println(v13)
                            if true {
                                var v42 int8 = -1
                                _ = len(v1)
                                _ = v42
                            }
                            var v43 int64 = 0
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 18; fi44++ {
                            var v45 float64 = 3.14159
                            _ = len(v3)
                            _ = fi44
                            _ = v45
                        }
                    }
                    if ((("\n\t" + "a") + "foo bar") >= "") {
                        var v46 float64 = 0.0
                        v4 += (v4 + (^(v4)))
                        _ = v46
                    }
                }
                _ = v39
            }
            fmt.Println(fi38, v5)
            v47 := GenS3{F0: 1, F1: fi38, F2: (v14 >> (((v13 - (v13 | v13))) & 63))}
            _ = fi38
            _ = v47
        }
        for fi48 := 0; fi48 < 19; fi48++ {
            for fi49 := 0; fi49 < 10; fi49++ {
                v50 := GenS4{F0: genFn0()}
                var v51 uint = v13
                v51 += (v13 | (^(v13)))
                _ = fi49
                _ = v50
                _ = v51
            }
            var v52 bool = false
            fmt.Println(v14)
            _ = fi48
            _ = v52
        }
    } else {
        var v53 uint8 = 2
        _ = len(v1)
        _ = v53
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
