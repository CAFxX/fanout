package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int }
type GenS2 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 bool; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint8; F1 int16; F2 int16 }
type GenS6 struct { F0 int32; F1 int32 }
func genFn0(p0 int64, p1 uint64, p2 int32) uint16 {
    var v7 uint16 = 65535
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint8 = 1
        for fi10 := 0; fi10 < 6; fi10++ {
            if (fi8 <= fi10) {
                v11 := []int{fi10, (fi8 & 678)}
                _ = v11
            }
            fmt.Println(v7, fi10, p1)
            v9 = v9
            switch -2147483648 {
            case 907:
                fmt.Println(fi10)
            case -2147483648:
                p0 -= (p0 - (p0 - int64(p2)))
            case 4294967295:
                v12 := []int{}
                _ = v12
            default:
                v13 := []int{}
                _ = v13
            }
            v14 := []int{-1}
            _ = fi10
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    var v15 bool = false
    v16 := []int{int(p2), 127}
    fmt.Println(v7, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v7
    return (^(v7))
}
func genFn1(p0 int32, p1 int16, p2 uint64) int64 {
    p1 -= (p1 % ((-32768) | 1))
    switch -601 {
    case 42:
        p1 -= p1
    default:
        p0 = -248
    }
    var v17 uint16 = 42
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 0
}

func main() {
    var v1 int64 = 0
    var v2 int8 = 42
    var v3 float64 = (1.5 / (3.14159 * 0.1))
    var v4 int16 = 0
    v5 := make(map[uint32]bool)
    var v6 int64 = v1
    if (127 >= 0) {
        delete(v5, 2)
        if true {
            for fi18 := 0; fi18 < 8; fi18++ {
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 int16 = v4
                    fmt.Println(v3, v20, v6)
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            _ = len(v5)
        }
        var v21 int = -9223372036854775808
        _ = v21
    } else {
        v3 += -2.25
    }
    if true {
        if (v4 < 1) {
            v2 = (^(v2))
            v22 := GenS1{F0: 296, F1: v1, F2: 531}
            _ = v22
        } else {
            if ((v2 % ((v2) | 1)) > v2) {
                for fi23 := 0; fi23 < 1; fi23++ {
                    v6 = 1
                    _ = fi23
                }
                for fi24 := 0; fi24 < 18; fi24++ {
                    v25 := GenS2{F0: 136, F1: -704, F2: 42}
                    _ = fi24
                    _ = v25
                }
            } else {
                v26 := GenS3{F0: 238, F1: ""}
                _ = v26
            }
            fmt.Println(v2, v1)
            switch int(v1) {
            case 42:
                fmt.Println(v6, v1)
            case 127:
                for fi27 := 0; fi27 < 6; fi27++ {
                    v2 -= 0
                    fmt.Println(v1, fi27, v4)
                    _ = fi27
                }
            case 605:
                var v28 bool = (true && false)
                _ = v28
            }
        }
    } else {
        for fi29 := 0; fi29 < 17; fi29++ {
            v30 := GenS4{F0: false, F1: (!false), F2: 2}
            for fi31 := 0; fi31 < 2; fi31++ {
                delete(v5, 42)
                var v32 uint64 = 18446744073709551615
                v33 := GenS1{F0: 2, F1: 1, F2: fi29}
                var v34 float32 = 3.14159
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            _ = fi29
            _ = v30
        }
    }
    var v35 int = -9223372036854775808
    fmt.Println(v1)
    for fi36 := 0; fi36 < 18; fi36++ {
        v1 -= (^(v6))
        var v37 int = fi36
        _ = fi36
        _ = v37
    }
    v38 := v5[2147483647]
    var v39 int = 42
    v40 := []int{v39}
    v2 -= -128
    var v41 int16 = (v4 + (v4 % ((int16(v2)) | 1)))
    v2 += v2
    for fi42 := 0; fi42 < 13; fi42++ {
        fmt.Println(v35, v3, v35)
        v41 -= (v4 >> ((227) & 63))
        for fi43 := 0; fi43 < 16; fi43++ {
            v3 += 3.14159
            v44 := v5[470]
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    v45 := GenS1{F0: 65535, F1: genFn1(2147483647, (v4 / (((v4 * (v41 - (v41 % ((-1) | 1))))) | 1)), 18446744073709551615), F2: (v39 << ((65535) & 63))}
    for fi46 := 0; fi46 < 10; fi46++ {
        if v38 {
            var v47 uint16 = 32767
            _ = v47
        } else {
            for fi48 := 0; fi48 < 11; fi48++ {
                switch ("hello" + "hello") {
                case "foo bar":
                    if ((((!false) || v38) || v38) && true) {
                        fmt.Println(v1, v6)
                    } else {
                        v49 := GenS5{F0: 20, F1: v41, F2: (v41 % (((v4 >> ((855) & 63))) | 1))}
                        if ((2 > 10) && (v38 || v38)) {
                            _ = len(v5)
                        } else {
                            fmt.Println(v35, v35, v39)
                        }
                        _ = v49
                    }
                case "":
                    v38 = (42 != 255)
                default:
                    v39 = 128
                }
                switch 128 {
                case -9223372036854775808:
                    v2 = (v2 * (-(v2)))
                case 882:
                    v50 := GenS4{F0: v38, F1: (212 > 467), F2: 0}
                    _ = v50
                }
                if v38 {
                    _ = len(v5)
                    v2 = v2
                    if (!((!v38) && v38)) {
                        var v51 float32 = (0.0 - 3.14159)
                        _ = v51
                    } else {
                        var v52 int8 = v2
                        fmt.Println(v1, v3, v6)
                        _ = v52
                    }
                    var v53 int16 = v41
                    _ = v53
                }
                v54 := GenS1{F0: 405, F1: genFn1(int32(v2), int16(v1), 2), F2: v39}
                _ = fi48
                _ = v54
            }
        }
        _ = fi46
    }
    fmt.Println(v38, v6)
    for fi55 := 0; fi55 < 9; fi55++ {
        switch "a" {
        case "hello":
            for fi56 := 0; fi56 < 15; fi56++ {
                fmt.Println(v1, v39)
                v57 := v5[2147483648]
                _ = len(v40)
                _ = fi56
                _ = v57
            }
        case "\n\t":
            if (-32 <= 0) {
                _ = len(v5)
                if (!(v38 && v38)) {
                    var v58 bool = (836 > 42)
                    for fi59 := 0; fi59 < 7; fi59++ {
                        if len(v40) > 0 {
                            v40[0] = 4294967295
                        }
                        _ = fi59
                    }
                    var v60 float32 = float32((0.1 * (float64(0.1) - (float64(3.14159) / float64(0.1)))))
                    _ = v58
                    _ = v60
                }
            } else {
                v61 := GenS5{F0: 0, F1: v41, F2: (v4 & v4)}
                for fi62 := 0; fi62 < 9; fi62++ {
                    if v38 {
                        v41 = (v41 - (v41 - (v41 - (v41 % (((v4 / (((v4 * 1000)) | 1))) | 1)))))
                        v63 := GenS6{F0: -2147483648, F1: 1}
                        _ = v63
                    } else {
                        var v64 uint8 = 128
                        var v65 uint16 = 8
                        var v66 int = v35
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    var v67 uint = 2
                    v39 = fi55
                    v68 := v5[2147483647]
                    v41 += v4
                    _ = fi62
                    _ = v67
                    _ = v68
                }
                _ = v61
            }
        case "foo bar":
            v3 = v3
        default:
            for fi69 := 0; fi69 < 19; fi69++ {
                var v70 float32 = float32(v3)
                var v71 uint8 = 127
                v72 := GenS3{F0: 2, F1: "\n\t"}
                _ = fi69
                _ = v70
                _ = v71
                _ = v72
            }
        }
        _ = fi55
    }
    var v73 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v45
    _ = v5
    _ = v6
    _ = v73
}
