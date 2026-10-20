package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int16; F1 int8; F2 uint }
type GenS5 struct { F0 uint64; F1 string }
func genFn0() uint16 {
    if true {
        switch -497 {
        case 42:
            var v6 uint = 527
            _ = v6
        case 2147483647:
            var v7 int16 = -32768
            _ = v7
        case 127:
            switch "a" {
            case "a":
                var v8 int = 128
                _ = v8
            case "foo bar":
                var v9 uint64 = 809
                _ = v9
            case "\n\t":
                var v10 int64 = 151
                _ = v10
            }
        default:
            v11 := GenS1{F0: 23}
            _ = v11
        }
        var v12 int64 = -814
        _ = v12
    }
    v13 := GenS2{F0: 1, F1: 222}
    _ = v13
    return 905
}
func genFn1() uint16 {
    for fi14 := 0; fi14 < 11; fi14++ {
        switch fi14 {
        case 32768:
            var v15 int8 = 1
            _ = v15
        }
        v16 := []int{4294967295, fi14, fi14, fi14}
        _ = fi14
        _ = v16
    }
    var v17 uint = 42
    if false {
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := GenS3{F0: 846, F1: 1}
            for fi20 := 0; fi20 < 13; fi20++ {
                v21 := GenS4{F0: 32767, F1: -28, F2: v17}
                if (false || true) {
                    v17 += v17
                }
                switch "a" {
                case "foo bar":
                    var v22 uint = (v17 << (((^(v17))) & 63))
                    _ = v22
                default:
                    v23 := []int{}
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            var v24 int16 = -1
            _ = fi18
            _ = v19
            _ = v24
        }
        if (!(v17 > (v17 - v17))) {
            if (584 == 2) {
                v17 += v17
                if false {
                    v25 := GenS1{F0: 18446744073709551615}
                    v26 := GenS1{F0: 42}
                    _ = v25
                    _ = v26
                }
                if true {
                    var v27 uint16 = 0
                    _ = v27
                } else {
                    if true {
                        var v28 bool = false
                        var v29 float64 = 1e10
                        _ = v28
                        _ = v29
                    } else {
                        fmt.Println(v17)
                    }
                }
            } else {
                for fi30 := 0; fi30 < 7; fi30++ {
                    v17 -= v17
                    switch "hello" {
                    case "a":
                        fmt.Println(fi30)
                    case "\n\t":
                        var v31 uint32 = 550
                        _ = v31
                    }
                    v17 += v17
                    _ = fi30
                }
            }
            fmt.Println(v17)
            for fi32 := 0; fi32 < 3; fi32++ {
                for fi33 := 0; fi33 < 3; fi33++ {
                    if (65535 > 464) {
                        if ((!((true || (!((18446744073709551615 <= 2) && (true && (-2147483648 == 717))))) || false)) && (-1 != 0)) {
                            var v34 int = (fi32 | (fi32 % (((fi32 / ((32768) | 1))) | 1)))
                            _ = v34
                        } else {
                            v35 := GenS3{F0: -32768, F1: 517}
                            v36 := []int{2147483647, (-(fi32)), 128}
                            var v37 uint16 = 1
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                    }
                    if (1 == 0) {
                        if (1 != v17) {
                            v38 := []int{fi32}
                            v17 = v17
                            _ = v38
                        }
                        v17 = 4294967295
                        v39 := []int{-18}
                        _ = v39
                    } else {
                        var v40 float32 = 3.14159
                        var v41 uint32 = 250
                        v42 := GenS4{F0: 1000, F1: int8(fi33), F2: v17}
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    v43 := GenS4{F0: 42, F1: -1, F2: (v17 & 65535)}
                    _ = fi33
                    _ = v43
                }
                v44 := []int{(fi32 ^ fi32), fi32, (fi32 & (fi32 | (fi32 + (fi32 - (^(fi32))))))}
                for fi45 := 0; fi45 < 3; fi45++ {
                    var v46 uint16 = 42
                    _ = fi45
                    _ = v46
                }
                _ = fi32
                _ = v44
            }
        }
    } else {
        var v47 float64 = (((float64(((1.5 / ((1.5 + (1e10 - (-2.25 / 0.1))) - (1.5 - -2.25))) - 0.1)) * ((1e10 * 1e10) - (((-2.25 * 1.5) - 1e10) + 0.1))) + (1.5 + 0.1)) - 3.14159)
        _ = v47
    }
    _ = v17
    return 255
}

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    var v4 uint = (v1 / ((v1) | 1))
    v5 := []uint32{0}
    var v48 float32 = 0.0
    v49 := GenS5{F0: 42, F1: "\n\t"}
    v1 = (v4 % ((1) | 1))
    v1 -= (v4 & 992)
    fmt.Println(v1, v4)
    fmt.Println(v4)
    for fi50 := 0; fi50 < 3; fi50++ {
        if (!false) {
            var v51 int = 1
            v52 := GenS2{F0: 678, F1: 964}
            _ = v51
            _ = v52
        }
        v53 := GenS3{F0: int16(fi50), F1: -1}
        _ = fi50
        _ = v53
    }
    var v54 int64 = -9223372036854775808
    fmt.Println(v4, v4, v1)
    if (!(!(42 < 1))) {
        v55 := GenS2{F0: 869, F1: 109}
        for fi56 := 0; fi56 < 14; fi56++ {
            v1 -= 351
            for fi57 := 0; fi57 < 6; fi57++ {
                delete(v2, 42)
                v58 := GenS4{F0: 32767, F1: 1, F2: (v1 / ((65535) | 1))}
                var v59 int16 = -991
                _ = fi57
                _ = v58
                _ = v59
            }
            v1 = 0
            _ = fi56
        }
        var v60 int8 = 107
        _ = v55
        _ = v60
    } else {
        if false {
            var v61 uint16 = 2
            _ = v61
        } else {
            for fi62 := 0; fi62 < 11; fi62++ {
                v4 = (v4 ^ (^(v1)))
                _ = fi62
            }
            v5 = append(v5, 0)
        }
        delete(v2, 42)
    }
    for fi63 := 0; fi63 < 3; fi63++ {
        v48 = 3.14159
        for fi64 := 0; fi64 < 10; fi64++ {
            var v65 int64 = (v54 / (((v54 ^ 181)) | 1))
            _ = fi64
            _ = v65
        }
        _ = fi63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v54
}
