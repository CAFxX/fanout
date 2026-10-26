package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int16; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32; F1 uint }

func main() {
    var v1 int = 128
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    v1 = 32767
    v1 -= 42
    v1 = (v1 >> ((0) & 63))
    _ = len(v4)
    var v6 bool = ((float64((1.5 * 0.0)) == 0.0) && (true && ((!(!false)) || false)))
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            if (v6 && (!v6)) {
                if (34 > 1) {
                    fmt.Println(fi7)
                } else {
                    v1 += (v1 & (fi7 * (fi7 << ((2) & 63))))
                }
                var v9 float64 = (float64(-2.25) * 3.14159)
                var v10 float32 = (3.14159 + 3.14159)
                _ = v10
                _ = v9
            }
            v1 = 0
            _ = fi8
        }
        v11 := []int{32767, 0}
        _ = fi7
        _ = v11
    }
    fmt.Println(v1)
    if true {
        v6 = (72 >= 42)
    }
    if true {
        switch v1 {
        case 488:
            v6 = true
        case 42:
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint = 4294967295
                var v14 int8 = -128
                var v15 int = fi12
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        case -9223372036854775808:
            switch (v1 / (((v1 + (v1 >> ((65535) & 63)))) | 1)) {
            case 891:
                var v16 int16 = -1
                _ = v16
            case 1:
                v1 -= v1
            default:
                var v17 uint = 119
                _ = v17
            }
        default:
            if (!v6) {
                v6 = (false || v6)
            } else {
                v6 = (3.14159 != 1e10)
            }
        }
        fmt.Println(v6, v1)
        for fi18 := 0; fi18 < 4; fi18++ {
            switch "hello" {
            case "":
                if (float32(1.5) < -2.25) {
                    v19 := []int{9223372036854775807, (-(v1))}
                    fmt.Println(fi18)
                    _ = v19
                }
            case "foo bar":
                fmt.Println(v6, fi18)
            case "\n\t":
                v2[("\n\t" + ("hello" + (("" + "") + "foo bar")))] = (^(v1))
            }
            v1 = (fi18 + v1)
            _ = fi18
        }
    }
    var v20 uint8 = 130
    v21 := GenS1{F0: 1}
    var v22 int = v1
    v23 := GenS2{F0: v6, F1: 2}
    v24 := v4["hello"]
    for fi25 := 0; fi25 < 1; fi25++ {
        if (0.1 != (1.5 / 3.14159)) {
            var v26 uint8 = v20
            switch "\n\t" {
            case "hello":
                var v27 int = (v24 % ((v24) | 1))
                _ = v27
            }
            if true {
                v28 := GenS3{F0: -32768, F1: 42, F2: -1}
                if true {
                    v20 -= v26
                    switch v24 {
                    case 4294967295:
                        switch "a" {
                        case "\n\t":
                            if v6 {
                                var v29 string = "\n\t"
                                _ = v29
                            }
                        case "hello":
                            v24 -= (^(v22))
                        default:
                            var v30 int = (v1 | (^(fi25)))
                            _ = v30
                        }
                    case -1:
                        v4["hello"] = (v24 * (fi25 * 765))
                    default:
                        switch (fi25 / ((v24) | 1)) {
                        case -9223372036854775808:
                            if false {
                                for fi31 := 0; fi31 < 18; fi31++ {
                                    if true {
                                        if ((v6 || (1e10 != (3.14159 * 0.0))) || (32767 < 65535)) {
                                            switch v22 {
                                            case 42:
                                                v32 := GenS4{F0: 2}
                                                _ = v32
                                            case 475:
                                                v33 := GenS2{F0: false, F1: 0}
                                                var v34 int16 = -1
                                                _ = v33
                                                _ = v34
                                            default:
                                                v26 -= 167
                                            }
                                            v35 := GenS5{F0: (v1 + (v1 ^ fi31)), F1: 32767}
                                            var v36 int = 4294967295
                                            _ = v35
                                            _ = v36
                                        }
                                        v24 -= fi25
                                        switch "hello" {
                                        case "":
                                            var v37 uint = 65535
                                            _ = v37
                                        case "\n\t":
                                            v22 -= (fi31 >> ((4294967295) & 63))
                                        default:
                                            var v38 uint16 = 247
                                            _ = v38
                                        }
                                    } else {
                                        var v39 uint8 = 127
                                        if (4294967295 >= 0) {
                                            v26 -= (v39 & (v20 * 127))
                                            v20 = 42
                                        }
                                        _ = v39
                                    }
                                    v40 := v5[4294967295]
                                    fmt.Println(fi25)
                                    _ = fi31
                                    _ = v40
                                }
                            } else {
                                if (723 < 385) {
                                    v6 = false
                                    var v41 int64 = 1
                                    _ = v41
                                } else {
                                    _ = len(v4)
                                    for fi42 := 0; fi42 < 15; fi42++ {
                                        if (0 == -26) {
                                            v43 := GenS3{F0: -32768, F1: 909, F2: 100}
                                            _ = v43
                                        }
                                        fmt.Println(fi42)
                                        _ = fi42
                                    }
                                    var v44 string = "hello"
                                    _ = v44
                                }
                            }
                        default:
                            var v45 uint64 = 4294967295
                            _ = v45
                        }
                    }
                }
                _ = v28
            }
            _ = v26
        } else {
            v6 = v6
            v46 := GenS5{F0: (^(v1)), F1: 32767}
            for fi47 := 0; fi47 < 7; fi47++ {
                switch (v22 + (-(fi25))) {
                case 42:
                    switch 393 {
                    case 496:
                        delete(v5, 708)
                    case -426:
                        v1 = (-(fi25))
                    default:
                        v48 := []int{}
                        _ = v48
                    }
                    if ((830 < 2) || (467 > 65535)) {
                        v49 := []int{320, v1, fi47, (v1 >> ((42) & 63))}
                        if (v6 && (-536 != 42)) {
                            if (!(!v6)) {
                                var v50 float32 = (((0.0 + -2.25) - (float32(float64((0.1 + 3.14159))) * -2.25)) - float32(0.1))
                                _ = v50
                            }
                        }
                        _ = v49
                    }
                }
                v20 = v20
                if true {
                    v22 = (^(v22))
                }
                v51 := GenS6{F0: 32767}
                if v6 {
                    v52 := GenS4{F0: 65535}
                    v53 := []int{}
                    _ = v52
                    _ = v53
                } else {
                    var v54 uint32 = uint32(v20)
                    switch (-(v22)) {
                    case 1:
                        if true {
                            v20 = v20
                            v55 := GenS7{F0: 3, F1: 400}
                            fmt.Println(fi47)
                            _ = v55
                        } else {
                            fmt.Println(fi47)
                            fmt.Println(v20, v20, fi47)
                            v56 := []int{(v24 ^ 2147483647)}
                            _ = v56
                        }
                    case 128:
                        v24 += 1
                    }
                    v22 -= int(v54)
                    _ = v54
                }
                _ = fi47
                _ = v51
            }
            _ = v46
        }
        for fi57 := 0; fi57 < 13; fi57++ {
            var v58 uint16 = 65535
            _ = fi57
            _ = v58
        }
        _ = fi25
    }
    v59 := GenS6{F0: -855}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
