package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint; F1 int64; F2 int32 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 int32; F1 bool; F2 uint8 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 bool; F1 string; F2 int32 }
func genFn0() int64 {
    var v4 int64 = 688
    if false {
        v4 -= v4
        switch "foo bar" {
        case "hello":
            for fi5 := 0; fi5 < 15; fi5++ {
                switch ("foo bar" + ("foo bar" + (("" + "foo bar") + "hello"))) {
                case "":
                    switch fi5 {
                    case 713:
                        var v6 uint16 = 32767
                        _ = v6
                    case 42:
                        v7 := []int{fi5, 32768, 32768, fi5}
                        _ = v7
                    case 0:
                        if (((((0.1 * 1.5) * (0.1 / -2.25)) - 0.0) >= (float32(3.14159) * float32((0.1 / -2.25)))) && (34 != 1)) {
                            var v8 int8 = -128
                            _ = v8
                        } else {
                            var v9 string = ("foo bar" + ("\n\t" + "hello"))
                            _ = v9
                        }
                    }
                case "hello":
                    var v10 uint8 = 255
                    _ = v10
                }
                var v11 uint64 = 0
                var v12 int = -840
                _ = fi5
                _ = v11
                _ = v12
            }
        case "":
            for fi13 := 0; fi13 < 8; fi13++ {
                for fi14 := 0; fi14 < 6; fi14++ {
                    v15 := GenS1{F0: -9223372036854775808, F1: 353}
                    v16 := []int{(fi13 % (((fi14 << ((0) & 63))) | 1)), -716, fi14}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = fi13
            }
        default:
            v4 += v4
        }
    }
    switch "a" {
    case "\n\t":
        v17 := GenS1{F0: 0, F1: 0}
        _ = v17
    case "a":
        v18 := []int{1, 32767, 42, -40}
        _ = v18
    case "":
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := []int{32768, (fi19 >> ((804) & 63)), -846}
            if ((false && true) || ("" != "\n\t")) {
                for fi21 := 0; fi21 < 4; fi21++ {
                    v22 := GenS1{F0: fi19, F1: 2}
                    v4 = -1
                    v4 = (v4 / (((^(v4))) | 1))
                    _ = fi21
                    _ = v22
                }
            } else {
                v23 := GenS2{F0: 0, F1: 14, F2: 0}
                _ = v23
            }
            var v24 uint16 = 0
            _ = fi19
            _ = v20
            _ = v24
        }
    default:
        if true {
            var v25 float64 = 0.1
            v4 = 1
            _ = v25
        }
    }
    _ = v4
    return v4
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 127
    v26 := GenS3{F0: 4294967295, F1: 970, F2: -2147483648}
    v27 := GenS4{F0: 0, F1: false}
    var v28 int16 = 1
    for fi29 := 0; fi29 < 17; fi29++ {
        v28 = (v28 % (((v28 + v28)) | 1))
        v30 := v2[2147483648]
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        switch fi31 {
        case -757:
            v28 -= v28
        case 4294967295:
            var v32 uint32 = 2
            _ = v32
        }
        _ = fi31
    }
    if false {
        for fi33 := 0; fi33 < 3; fi33++ {
            var v34 int = -9223372036854775808
            if (genFn0() == 172) {
                var v35 uint16 = 325
                _ = v35
            } else {
                var v36 uint16 = 0
                var v37 float64 = (0.0 / (3.14159 + 0.1))
                _ = v36
                _ = v37
            }
            v28 += (v28 - (v28 * v28))
            _ = fi33
            _ = v34
        }
        v28 = (^(v28))
        switch ("hello" + "hello") {
        case "hello":
            var v38 uint32 = 42
            _ = v38
        case "a":
            v3 = 42
        case "":
            fmt.Println(v28, v3)
        }
    } else {
        switch "foo bar" {
        case "hello":
            v39 := GenS5{F0: 42, F1: (genFn0() < 1), F2: 1}
            _ = v39
        case "\n\t":
            for fi40 := 0; fi40 < 17; fi40++ {
                var v41 uint16 = 492
                _ = fi40
                _ = v41
            }
        }
    }
    v3 = 30
    v42 := GenS1{F0: 2147483647, F1: 1}
    v43 := GenS3{F0: 171, F1: 9223372036854775807, F2: 42}
    var v44 uint8 = v3
    _ = len(v1)
    v45 := GenS6{F0: -9223372036854775808, F1: (-(v28))}
    for fi46 := 0; fi46 < 8; fi46++ {
        var v47 float64 = 1e10
        switch (fi46 - fi46) {
        case -1:
            for fi48 := 0; fi48 < 11; fi48++ {
                var v49 int32 = -622
                if true {
                    var v50 uint8 = (^(v3))
                    fmt.Println(v28)
                    _ = v50
                }
                _ = fi48
                _ = v49
            }
        case 42:
            v51 := GenS7{F0: 36, F1: (v44 / (((v3 >> ((255) & 63))) | 1))}
            _ = v51
        default:
            v1[(fi46 + (fi46 & (^(fi46))))] = false
        }
        _ = fi46
        _ = v47
    }
    for fi52 := 0; fi52 < 14; fi52++ {
        var v53 uint = 1
        fmt.Println(v44, v28, v53)
        v54 := GenS8{F0: -1, F1: -116}
        _ = fi52
        _ = v53
        _ = v54
    }
    var v55 uint32 = 2
    for fi56 := 0; fi56 < 1; fi56++ {
        for fi57 := 0; fi57 < 3; fi57++ {
            var v58 uint = 4294967295
            switch fi57 {
            case 127:
                var v59 float64 = ((1e10 + 1e10) + -2.25)
                _ = v59
            case -346:
                if (v44 < 0) {
                    fmt.Println(v58)
                } else {
                    fmt.Println(v3, fi56)
                    v60 := GenS9{F0: false, F1: ("foo bar" + "hello"), F2: 42}
                    v55 = v55
                    _ = v60
                }
            case -506:
                switch -716 {
                case 1:
                    v28 -= 1
                case 42:
                    switch "hello" {
                    case "a":
                        v61 := []int{}
                        _ = v61
                    case "hello":
                        var v62 uint32 = v55
                        _ = v62
                    case "foo bar":
                        fmt.Println(v3, v58)
                    default:
                        v63 := GenS3{F0: 42, F1: -887, F2: 1}
                        _ = v63
                    }
                }
            default:
                v28 -= (v28 << ((v58) & 63))
            }
            fmt.Println(v58, v55)
            _ = fi57
            _ = v58
        }
        _ = fi56
    }
    v64 := v1[-850]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v55:", v55)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v55
    _ = v64
}
