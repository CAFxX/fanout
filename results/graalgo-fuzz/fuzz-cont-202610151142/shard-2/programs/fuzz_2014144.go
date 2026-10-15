package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint64; F1 uint8; F2 bool }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint16; F1 int32; F2 int }
type GenS7 struct { F0 uint64; F1 bool; F2 int64 }
type GenS8 struct { F0 int16; F1 uint64 }
func genFn0() int8 {
    var v6 bool = false
    v7 := GenS1{F0: 42}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 int32 = 42
    v2 := []int{}
    v3 := make(map[int]int)
    var v4 uint8 = 255
    var v5 int16 = 1
    v4 -= (v4 * v4)
    v2 = append(v2, 398)
    v8 := GenS1{F0: 9223372036854775807}
    v9 := GenS1{F0: 9223372036854775807}
    fmt.Println(v5, v4, v5)
    var v10 int16 = v5
    if (-107 >= -2147483648) {
        for fi11 := 0; fi11 < 12; fi11++ {
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 uint8 = (^(v4))
                v1 = v1
                if (!false) {
                    var v14 int = fi11
                    if (!(true && (4294967295 > 1))) {
                        var v15 int = int(v1)
                        _ = v15
                    }
                    if (42 >= 1) {
                        if true {
                            v4 = v13
                            v16 := GenS1{F0: 809}
                            v1 -= (v1 << ((2) & 63))
                            v1 += (v1 - v1)
                            _ = v16
                        }
                        v17 := GenS2{F0: 0, F1: 2}
                        var v18 int32 = (v1 % ((v1) | 1))
                        v2 = append(v2, 4294967295)
                        _ = v17
                        _ = v18
                    } else {
                        var v19 uint = 109
                        v19 -= v19
                        if false {
                            if true {
                                fmt.Println(v5, v4, v14)
                                fmt.Println(fi12)
                            }
                            v13 -= 1
                        } else {
                            var v20 uint8 = (v13 + v4)
                            v21 := GenS2{F0: 18446744073709551615, F1: (^(v19))}
                            _ = v20
                            _ = v21
                        }
                        _ = v19
                    }
                    _ = v14
                } else {
                    _ = len(v3)
                    if ((v13 >> ((42) & 63)) >= 255) {
                        var v22 int32 = v1
                        v2 = append(v2, -321)
                        switch fi11 {
                        case 4294967295:
                            if ((fi12 < (fi12 % (((-(fi11))) | 1))) && true) {
                                var v23 uint8 = (v4 / ((v13) | 1))
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v5 += (v10 << ((18446744073709551615) & 63))
                                    v5 = (^(v5))
                                    v24 := GenS3{F0: 18446744073709551615, F1: 127, F2: (false || true)}
                                    _ = v24
                                }
                                _ = v23
                            } else {
                                var v25 string = ""
                                var v26 uint16 = 65535
                                v27 := GenS4{F0: 58, F1: 2}
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        case 127:
                            switch (fi12 ^ 4294967295) {
                            case -742:
                                delete(v3, (fi11 * fi12))
                            default:
                                _ = len(v2)
                            }
                        case 42:
                            var v28 string = "\n\t"
                            _ = v28
                        }
                        _ = v22
                    }
                    if false {
                        if true {
                            var v29 uint = 255
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v30 := GenS2{F0: 42, F1: 4294967295}
                            _ = v29
                            _ = v30
                        }
                        fmt.Println(v5, v1)
                        var v31 uint8 = (v13 << ((65535) & 63))
                        _ = v31
                    }
                }
                v1 += v1
                if ((false || ("hello" <= ("a" + ""))) && (!false)) {
                    v32 := GenS5{F0: 1, F1: 2147483647, F2: 0}
                    v1 += (v1 % ((int32(v4)) | 1))
                    _ = v32
                }
                _ = fi12
                _ = v13
            }
            v5 = 0
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 int = 32767
                v35 := GenS2{F0: 427, F1: 1}
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi11
        }
        if true {
            switch -9223372036854775808 {
            case 4294967295:
                for fi36 := 0; fi36 < 14; fi36++ {
                    var v37 float32 = (3.14159 / 1.5)
                    _ = fi36
                    _ = v37
                }
            case 577:
                var v38 uint64 = 822
                _ = v38
            case 1:
                var v39 uint = 0
                _ = v39
            }
            for fi40 := 0; fi40 < 14; fi40++ {
                fmt.Println(v1, v5)
                if (!true) {
                    v41 := GenS3{F0: 1, F1: (v4 ^ v4), F2: false}
                    fmt.Println(v4)
                    _ = v41
                }
                var v42 bool = (118 != 303)
                _ = fi40
                _ = v42
            }
        }
        if len(v2) > 0 {
            v2[0] = 32768
        }
    } else {
        var v43 int16 = (^(v5))
        var v44 int64 = 0
        v4 = (v4 | 42)
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        var v46 string = ("\n\t" + "")
        _ = fi45
        _ = v46
    }
    if (32767 >= v10) {
        v47 := GenS4{F0: 42, F1: 0}
        v5 = v10
        for fi48 := 0; fi48 < 20; fi48++ {
            v49 := GenS6{F0: 42, F1: v1, F2: int(v1)}
            for fi50 := 0; fi50 < 3; fi50++ {
                v2 = append(v2, (fi48 / ((fi50) | 1)))
                v10 += v10
                v51 := GenS1{F0: -9223372036854775808}
                var v52 float64 = 0.1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi50
                _ = v51
                _ = v52
            }
            v53 := GenS7{F0: 442, F1: false, F2: -658}
            _ = fi48
            _ = v49
            _ = v53
        }
        _ = v47
    }
    v54 := v3[127]
    for fi55 := 0; fi55 < 20; fi55++ {
        v54 = (fi55 << ((1) & 63))
        _ = fi55
    }
    if (false || ((v1 / (((v1 - v1)) | 1)) <= v1)) {
        for fi56 := 0; fi56 < 14; fi56++ {
            for fi57 := 0; fi57 < 7; fi57++ {
                v3[fi57] = (fi56 - -9223372036854775808)
                v54 += (fi56 | 896)
                v5 += int16(fi56)
                _ = fi57
            }
            v58 := GenS5{F0: 65535, F1: 0, F2: -9223372036854775808}
            for fi59 := 0; fi59 < 4; fi59++ {
                if true {
                    v10 = (v5 << ((4294967295) & 63))
                } else {
                    var v60 string = ("a" + ("a" + ""))
                    _ = v60
                }
                var v61 int16 = (-(v10))
                switch 453 {
                case 39:
                    var v62 int = fi56
                    _ = v62
                case -908:
                    if ((!(!false)) && (!(v10 >= (v5 * (v10 >> ((42) & 63)))))) {
                        v1 -= v1
                        if false {
                            _ = len(v2)
                            var v63 int8 = 0
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if true {
                                fmt.Println(fi59, v61, fi59)
                                var v64 bool = (!(uint32(v4) > 0))
                                if false {
                                    if v64 {
                                        var v65 int = v54
                                        _ = v65
                                    }
                                    v10 = -32768
                                    delete(v3, (-(v54)))
                                }
                                _ = v64
                            }
                            var v66 uint8 = 42
                            _ = v63
                            _ = v66
                        }
                    }
                default:
                    v67 := GenS8{F0: v61, F1: 378}
                    _ = v67
                }
                _ = fi59
                _ = v61
            }
            _ = fi56
            _ = v58
        }
        if len(v2) > 0 {
            v2[0] = (v54 % ((2147483647) | 1))
        }
    } else {
        v54 = v54
        var v68 int = -400
        for fi69 := 0; fi69 < 1; fi69++ {
            v10 = (-(v5))
            _ = fi69
        }
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v8
    _ = v9
}
