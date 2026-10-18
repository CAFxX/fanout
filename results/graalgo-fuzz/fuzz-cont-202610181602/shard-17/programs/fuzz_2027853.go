package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 bool; F2 uint8 }
type GenS8 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0(p0 int16, p1 int32, p2 uint16) uint32 {
    var v5 int64 = 0
    fmt.Println(p1, v5, p0)
    v5 += (^(v5))
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 42, F1: 42, F2: (p1 * (p1 ^ p1))}
        v8 := GenS2{F0: 991}
        var v9 int8 = 127
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 549
}
func genFn1(p0 int64, p1 bool) int64 {
    p0 = (^(p0))
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    var v3 uint16 = 65535
    v4 := make(map[uint32]string)
    _ = len(v4)
    var v10 uint16 = 0
    v3 += (v3 << ((18446744073709551615) & 63))
    var v11 uint = 4294967295
    switch ("a" + "hello") {
    case "hello":
        if false {
            v1 = (42 == 1)
            if (-39 > -40) {
                for fi12 := 0; fi12 < 2; fi12++ {
                    v13 := GenS3{F0: -1}
                    if false {
                        switch ("hello" + ("\n\t" + ("foo bar" + "\n\t"))) {
                        case "hello":
                            switch ((("foo bar" + "") + "a") + "foo bar") {
                            case "a":
                                v11 = (^(v11))
                            case "hello":
                                for fi14 := 0; fi14 < 4; fi14++ {
                                    var v15 float64 = 1.5
                                    v16 := GenS1{F0: 0, F1: genFn0(1, 2147483647, v10), F2: 42}
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                            case "":
                                switch (fi12 + fi12) {
                                case 42:
                                    switch 4294967295 {
                                    case 32768:
                                        v17 := []int{0, (fi12 >> (((v11 % ((349) | 1))) & 63)), fi12, int(v11)}
                                        _ = v17
                                    case 127:
                                        if ((!(-1 >= 1000)) && v1) {
                                            var v18 int8 = -1
                                            switch -1 {
                                            case 928:
                                                v19 := GenS4{F0: fi12}
                                                _ = v19
                                            case -706:
                                                fmt.Println(fi12, v1, v1)
                                            case 128:
                                                v1 = v1
                                            }
                                            _ = v18
                                        } else {
                                            v20 := GenS5{F0: 1}
                                            _ = v20
                                        }
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            var v21 int16 = -483
                                            _ = v21
                                        default:
                                            switch (fi12 % ((fi12) | 1)) {
                                            case 127:
                                                if v1 {
                                                    v22 := GenS2{F0: 18446744073709551615}
                                                    _ = v22
                                                } else {
                                                    v23 := []int{(fi12 << (((v11 * 1)) & 63)), (fi12 - fi12)}
                                                    var v24 bool = (fi12 == 32768)
                                                    _ = v23
                                                    _ = v24
                                                }
                                            case 817:
                                                fmt.Println(fi12)
                                            case 92:
                                                for fi25 := 0; fi25 < 18; fi25++ {
                                                    var v26 int64 = -9223372036854775808
                                                    var v27 uint64 = uint64(v3)
                                                    var v28 uint32 = 0
                                                    _ = fi25
                                                    _ = v26
                                                    _ = v27
                                                    _ = v28
                                                }
                                            default:
                                                for fi29 := 0; fi29 < 15; fi29++ {
                                                    switch 4294967295 {
                                                    case 245:
                                                        v30 := []int{fi12, (fi12 % ((fi12) | 1))}
                                                        _ = v30
                                                    case 961:
                                                        if (0 > 0) {
                                                            v11 = (v11 << ((1) & 63))
                                                        } else {
                                                            fmt.Println(v1, fi12, fi12)
                                                            v3 = v10
                                                            var v31 int = (^(fi29))
                                                            _ = v31
                                                        }
                                                    case 1:
                                                        var v32 uint64 = 0
                                                        _ = v32
                                                    default:
                                                        switch "\n\t" {
                                                        case "":
                                                            v11 = (v11 + (v11 % (((v11 % ((uint(v3)) | 1))) | 1)))
                                                        case "hello":
                                                            switch "a" {
                                                            case "foo bar":
                                                                v3 -= (v10 % (((v3 >> ((v11) & 63))) | 1))
                                                            case "a":
                                                                fmt.Println(fi29, fi12)
                                                            case "":
                                                                v3 = 0
                                                            }
                                                        default:
                                                            if true {
                                                                if v1 {
                                                                    var v33 int8 = 127
                                                                    v10 = v10
                                                                    var v34 float32 = -2.25
                                                                    _ = v33
                                                                    _ = v34
                                                                } else {
                                                                    var v35 float32 = -2.25
                                                                    fmt.Println(v10, v35)
                                                                    if false {
                                                                        var v36 float64 = -2.25
                                                                        v37 := []int{fi12}
                                                                        fmt.Println(fi12, v1, v11)
                                                                        _ = v36
                                                                        _ = v37
                                                                    } else {
                                                                        var v38 uint8 = 4
                                                                        if ((!v1) || v1) {
                                                                            var v39 int16 = 32767
                                                                            _ = v39
                                                                        } else {
                                                                            v1 = v1
                                                                            v40 := v4[uint32(v38)]
                                                                            _ = v40
                                                                        }
                                                                        _ = v38
                                                                    }
                                                                    _ = v35
                                                                }
                                                                if v1 {
                                                                    var v41 float32 = -2.25
                                                                    var v42 int16 = 1
                                                                    v41 = (float32(((0.0 / -2.25) / (1.5 * 1e10))) * 1.5)
                                                                    _ = v41
                                                                    _ = v42
                                                                } else {
                                                                    v43 := GenS3{F0: 265}
                                                                    _ = v43
                                                                }
                                                            } else {
                                                                switch 32767 {
                                                                case -746:
                                                                    fmt.Println(v10, fi12, fi29)
                                                                case -98:
                                                                    fmt.Println(v1, v11)
                                                                    var v44 uint16 = v3
                                                                    v44 = v10
                                                                    _ = v44
                                                                case -1:
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        _ = len(v2)
                                                                    }
                                                                }
                                                                fmt.Println(v11, v10)
                                                            }
                                                        }
                                                    }
                                                    v10 = (v10 * v3)
                                                    _ = fi29
                                                }
                                            }
                                        }
                                    }
                                case 9223372036854775807:
                                    var v45 string = "a"
                                    _ = v45
                                case 348:
                                    v1 = v1
                                }
                            default:
                                switch 417 {
                                case -9223372036854775808:
                                    if false {
                                        for fi46 := 0; fi46 < 15; fi46++ {
                                            var v47 uint = v11
                                            _ = fi46
                                            _ = v47
                                        }
                                        var v48 float32 = 3.14159
                                        _ = v48
                                    } else {
                                        var v49 int32 = -2147483648
                                        v11 -= (v11 + (v11 << ((v11) & 63)))
                                        var v50 int32 = v49
                                        _ = v49
                                        _ = v50
                                    }
                                }
                            }
                        }
                        v51 := GenS6{F0: 215}
                        _ = v51
                    }
                    _ = fi12
                    _ = v13
                }
            }
        }
    default:
        var v52 uint64 = 42
        _ = v52
    }
    for fi53 := 0; fi53 < 14; fi53++ {
        fmt.Println(v10, v10)
        v1 = v1
        v54 := []int{-2147483648, (fi53 & (-(fi53))), fi53}
        _ = fi53
        _ = v54
    }
    for fi55 := 0; fi55 < 7; fi55++ {
        v10 = v10
        for fi56 := 0; fi56 < 7; fi56++ {
            var v57 uint8 = 255
            _ = fi56
            _ = v57
        }
        _ = fi55
    }
    if true {
        switch -2147483648 {
        case 4294967295:
            v58 := GenS7{F0: -924, F1: (!true), F2: 225}
            _ = v58
        case 286:
            if v1 {
                v59 := []int{254, -9223372036854775808, 128}
                _ = v59
            }
        }
        if v1 {
            v60 := GenS8{F0: 785, F1: 1, F2: 199}
            v10 = 470
            _ = v60
        } else {
            v10 -= (v3 >> ((v11) & 63))
            fmt.Println(v10, v1, v10)
        }
    } else {
        v61 := []int{168, 32768}
        for fi62 := 0; fi62 < 19; fi62++ {
            _ = len(v61)
            _ = fi62
        }
        var v63 float64 = (0.0 / 1.0)
        _ = v61
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
