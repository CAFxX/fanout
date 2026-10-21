package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]string)
    var v3 uint32 = 2147483648
    var v4 uint = 51
    switch ("a" + "a") {
    case "\n\t":
        if (!((32767 >= 42) && false)) {
            var v5 int = 32768
            var v6 uint = v4
            _ = v5
            _ = v6
        }
    }
    v7 := v2[42]
    if true {
        _ = len(v2)
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: v1, F1: 892}
            _ = fi8
            _ = v9
        }
    } else {
        v10 := GenS1{F0: v1, F1: 201}
        var v11 float64 = (1.5 - 0.1)
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := []int{}
            var v15 uint8 = 128
            v14 = append(v14, 737)
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := GenS1{F0: 0, F1: 1}
        _ = fi12
        _ = v16
    }
    if (9223372036854775807 >= -9223372036854775808) {
        v4 -= (^(v4))
    } else {
        for fi17 := 0; fi17 < 15; fi17++ {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 int32 = 42
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 10; fi20++ {
                v7 = "foo bar"
                if (!((!(v4 == (^(v4)))) && (65535 < 42))) {
                    if (9223372036854775807 != -9223372036854775808) {
                        var v21 bool = ((v3 & (^(v3))) >= 750)
                        if v21 {
                            v22 := GenS1{F0: int8(v4), F1: 915}
                            if (!false) {
                                fmt.Println(v1, fi17)
                                v3 = v3
                                delete(v2, v3)
                            }
                            var v23 string = "foo bar"
                            v2[v3] = "foo bar"
                            _ = v22
                            _ = v23
                        }
                        v21 = v21
                        fmt.Println(fi20, fi17)
                        _ = v21
                    } else {
                        switch 2147483647 {
                        case 898:
                            if ((false || true) || (4294967295 >= 0)) {
                                v7 = "a"
                                v2[v3] = ("hello" + "")
                            }
                        case 53:
                            switch 4294967295 {
                            case 0:
                                v2[v3] = v7
                            case 97:
                                var v24 int32 = -253
                                _ = v24
                            default:
                                v1 += v1
                            }
                        default:
                            if false {
                                v7 = v7
                                if ((v4 + v4) == (v4 & v4)) {
                                    switch v7 {
                                    case "\n\t":
                                        switch fi17 {
                                        case 127:
                                            var v25 int64 = 42
                                            _ = v25
                                        case 128:
                                            var v26 uint8 = 128
                                            var v27 int32 = -893
                                            v28 := GenS2{F0: 42, F1: 127, F2: 42}
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                    }
                                    v2[(^(v3))] = ((v7 + v7) + "a")
                                    switch 719 {
                                    case -769:
                                        switch 9223372036854775807 {
                                        case 32767:
                                            if false {
                                                v1 = (v1 ^ 42)
                                            } else {
                                                v29 := GenS3{F0: ("hello" + "a")}
                                                _ = v29
                                            }
                                            fmt.Println(v3)
                                            v30 := GenS4{F0: (v1 | (v1 | (v1 & 0))), F1: 42}
                                            _ = v30
                                        case 728:
                                            v7 = v7
                                        case 802:
                                            v3 = (v3 >> ((v4) & 63))
                                        default:
                                            if true {
                                                v31 := GenS1{F0: 127, F1: 4294967295}
                                                _ = v31
                                            } else {
                                                _ = len(v2)
                                                switch v7 {
                                                case "\n\t":
                                                    if ((fi20 >> ((v4) & 63)) >= (fi17 % (((-(fi17))) | 1))) {
                                                        v7 = v7
                                                        v32 := GenS1{F0: (v1 & (^(v1))), F1: 18446744073709551615}
                                                        switch -873 {
                                                        case 2147483647:
                                                            if ((((float64((0.0 + (float32(1.5) * (1.5 - float32(1e10))))) + 1e10) * 1.5) * 0.1) == float64(float32(3.14159))) {
                                                                v33 := GenS5{F0: 65535}
                                                                var v34 int64 = -373
                                                                if ((!true) && true) {
                                                                    v2[v3] = v7
                                                                } else {
                                                                    var v35 uint32 = v3
                                                                    v3 -= v3
                                                                    _ = v35
                                                                }
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                        case 4294967295:
                                                            var v36 int32 = 42
                                                            _ = v36
                                                        case 1:
                                                            switch fi17 {
                                                            case 0:
                                                                var v37 uint8 = 24
                                                                _ = v37
                                                            case 273:
                                                                switch (v7 + "\n\t") {
                                                                case "a":
                                                                    v1 -= (v1 << ((v4) & 63))
                                                                }
                                                            case 32767:
                                                                v38 := GenS1{F0: v1, F1: 42}
                                                                _ = v38
                                                            default:
                                                                if (!true) {
                                                                    fmt.Println(fi20, fi17)
                                                                }
                                                            }
                                                        }
                                                        _ = v32
                                                    }
                                                case "hello":
                                                    v39 := []int{(fi17 << ((v4) & 63)), int(v4)}
                                                    _ = v39
                                                }
                                                var v40 float32 = (0.1 - 0.0)
                                                _ = v40
                                            }
                                        }
                                    case 9223372036854775807:
                                        var v41 int8 = (v1 % ((v1) | 1))
                                        _ = v41
                                    default:
                                        v2[v3] = v7
                                    }
                                } else {
                                    v7 = v7
                                    var v42 float64 = 3.14159
                                    if false {
                                        fmt.Println(v42, v42)
                                    } else {
                                        v43 := GenS6{F0: 1, F1: 1000, F2: 42}
                                        v44 := GenS3{F0: "a"}
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = v42
                                }
                            } else {
                                v1 = int8(fi20)
                                if (!((-467 == 0) && true)) {
                                    fmt.Println(fi17, fi20)
                                    var v45 uint8 = uint8(v3)
                                    _ = v45
                                }
                            }
                        }
                        switch "hello" {
                        case "foo bar":
                            var v46 int = 4294967295
                            _ = v46
                        }
                    }
                }
                _ = fi20
            }
            _ = fi17
        }
        v1 -= -70
        fmt.Println(v4, v7, v7)
    }
    if (((-71 == -821) || false) && false) {
        v47 := GenS2{F0: 1, F1: 0, F2: 0}
        if false {
            if (!(!(!false))) {
                fmt.Println(v4)
                var v48 uint64 = 4294967295
                _ = v48
            } else {
                v4 -= (v4 - (v4 / (((v4 - v4)) | 1)))
            }
            v3 = (v3 / (((v3 - (^(v3)))) | 1))
            var v49 uint = (v4 >> (((v4 % ((65535) | 1))) & 63))
            _ = v49
        } else {
            var v50 int = -2147483648
            _ = v50
        }
        var v51 string = v7
        _ = v47
        _ = v51
    } else {
        fmt.Println(v1, v3)
        var v52 int8 = (v1 * (v1 & (v1 << ((2) & 63))))
        _ = len(v2)
        _ = v52
    }
    fmt.Println(v7)
    v53 := GenS7{F0: 1}
    if false {
        var v54 int32 = int32(v1)
        _ = v54
    } else {
        var v55 int32 = 0
        _ = v55
    }
    for fi56 := 0; fi56 < 8; fi56++ {
        if true {
            v57 := []int{fi56, (fi56 % ((128) | 1)), fi56}
            switch fi56 {
            case 42:
                v58 := GenS6{F0: 128, F1: (v3 & v3), F2: 2}
                _ = v58
            case 169:
                _ = len(v2)
            default:
                v7 = "hello"
            }
            _ = v57
        } else {
            if true {
                v59 := GenS6{F0: 111, F1: (v3 - v3), F2: 1}
                fmt.Println(v1)
                _ = v59
            } else {
                var v60 uint32 = (v3 << ((267) & 63))
                _ = v60
            }
            for fi61 := 0; fi61 < 5; fi61++ {
                v62 := GenS8{F0: 0, F1: 255}
                _ = len(v2)
                _ = fi61
                _ = v62
            }
        }
        _ = fi56
    }
    v63 := []int{515, -9223372036854775808, 683}
    fmt.Println(v3, v4)
    var v64 string = (("a" + v7) + v7)
    v7 = (v64 + v7)
    v4 -= 4294967295
    v2[v3] = v7
    var v65 bool = (32767 < -1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v63
    _ = v64
    _ = v65
    _ = v7
}
