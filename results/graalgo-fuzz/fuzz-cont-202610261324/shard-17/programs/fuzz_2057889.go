package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int32; F2 uint }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int16; F1 bool }
func genFn0(p0 uint, p1 int8) uint {
    var v7 uint32 = 1
    fmt.Println(p0)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint16 = 2
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return (p0 & (p0 + (p0 & 319)))
}
func genFn1(p0 int, p1 uint32, p2 uint16) uint64 {
    v10 := GenS1{F0: 0}
    p2 = (p2 ^ (p2 ^ p2))
    var v11 string = "a"
    if (true || (!false)) {
        fmt.Println(p1, p0)
        v12 := []int{(p0 << ((21) & 63)), (p0 + p0), p0}
        v13 := GenS1{F0: -1}
        _ = v12
        _ = v13
    } else {
        for fi14 := 0; fi14 < 12; fi14++ {
            p0 = (p0 | 9223372036854775807)
            _ = fi14
        }
    }
    var v15 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v15
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]int)
    var v3 float64 = ((((0.1 - ((3.14159 / 1.0) / 1.0)) / 0.1) / 1.0) / 1.0)
    var v4 int8 = 0
    var v5 uint16 = 449
    v6 := make(map[uint32]int)
    switch ("hello" + "\n\t") {
    case "a":
        v16 := GenS2{F0: (v4 & (v4 / (((v4 * v4)) | 1))), F1: -1, F2: 18446744073709551615}
        _ = v16
    }
    var v17 uint = 2
    switch 128 {
    case 127:
        var v18 int32 = -2147483648
        _ = v18
    case -594:
        for fi19 := 0; fi19 < 8; fi19++ {
            switch (-(fi19)) {
            case -1:
                v3 = (v3 - v3)
            case 9223372036854775807:
                v20 := GenS3{F0: 643, F1: uint16(v17), F2: 65535}
                _ = v20
            case 1:
                v21 := []int{371}
                _ = v21
            }
            v22 := GenS2{F0: v4, F1: 42, F2: (v17 ^ v17)}
            if false {
                switch "hello" {
                case "a":
                    var v23 int64 = 254
                    _ = v23
                case "":
                    var v24 uint = genFn0(v17, (v4 >> ((v17) & 63)))
                    _ = v24
                case "foo bar":
                    if true {
                        if false {
                            v25 := GenS3{F0: uint32(v5), F1: (v5 << (((v17 >> ((v17) & 63))) & 63)), F2: 2}
                            for fi26 := 0; fi26 < 6; fi26++ {
                                if false {
                                    v27 := GenS4{F0: v5, F1: 224, F2: 51}
                                    var v28 int8 = v4
                                    _ = v27
                                    _ = v28
                                }
                                _ = fi26
                            }
                            var v29 uint64 = 18446744073709551615
                            _ = v25
                            _ = v29
                        } else {
                            v17 += (v17 & v17)
                            if (!(1e10 != (0.0 * 1.5))) {
                                for fi30 := 0; fi30 < 7; fi30++ {
                                    v4 -= (v4 - 127)
                                    _ = fi30
                                }
                            } else {
                                for fi31 := 0; fi31 < 3; fi31++ {
                                    v32 := GenS2{F0: (v4 << ((4294967295) & 63)), F1: 2147483647, F2: v17}
                                    if false {
                                        delete(v6, 2147483648)
                                        switch int(v17) {
                                        case 990:
                                            if (!false) {
                                                v33 := GenS3{F0: 467, F1: v5, F2: 341}
                                                _ = v33
                                            } else {
                                                fmt.Println(v5)
                                            }
                                        case 42:
                                            fmt.Println(v4)
                                        }
                                    } else {
                                        fmt.Println(v3, fi31)
                                        v2[4294967295] = (fi19 ^ fi31)
                                        fmt.Println(fi31, v3, fi19)
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                                for fi34 := 0; fi34 < 7; fi34++ {
                                    var v35 uint = v17
                                    v5 -= (v5 & (v5 << (((^(v17))) & 63)))
                                    _ = fi34
                                    _ = v35
                                }
                            }
                            v17 = (v17 * (v17 >> ((genFn0(v17, v4)) & 63)))
                        }
                        var v36 uint32 = 2147483648
                        var v37 uint8 = 1
                        _ = v36
                        _ = v37
                    }
                default:
                    var v38 int16 = 488
                    _ = v38
                }
            } else {
                _ = len(v2)
            }
            _ = fi19
            _ = v22
        }
    case 995:
        for fi39 := 0; fi39 < 6; fi39++ {
            v40 := GenS2{F0: (v4 * 100), F1: -2147483648, F2: (v17 * 65535)}
            for fi41 := 0; fi41 < 10; fi41++ {
                var v42 bool = (297 == v5)
                fmt.Println(v3, v4)
                v43 := []int{fi41, (fi41 + -2147483648), (fi41 & fi39), fi39}
                if len(v43) > 0 {
                    v43[0] = int(v5)
                }
                _ = fi41
                _ = v42
                _ = v43
            }
            for fi44 := 0; fi44 < 17; fi44++ {
                if (!true) {
                    var v45 int32 = 583
                    _ = v45
                }
                _ = fi44
            }
            _ = fi39
            _ = v40
        }
    default:
        for fi46 := 0; fi46 < 15; fi46++ {
            v4 = 42
            _ = fi46
        }
    }
    v17 += (^(v17))
    for fi47 := 0; fi47 < 6; fi47++ {
        for fi48 := 0; fi48 < 17; fi48++ {
            var v49 uint8 = 104
            if false {
                v50 := GenS3{F0: 42, F1: (v5 + (v5 / (((v5 ^ v5)) | 1))), F2: 65535}
                _ = v50
            }
            _ = fi48
            _ = v49
        }
        for fi51 := 0; fi51 < 4; fi51++ {
            var v52 float64 = v3
            v53 := []int{(-(fi51)), fi51, (fi47 ^ fi51), (fi47 | int(v4))}
            _ = fi51
            _ = v52
            _ = v53
        }
        if false {
            for fi54 := 0; fi54 < 6; fi54++ {
                if false {
                    var v55 uint64 = 0
                    _ = v55
                } else {
                    fmt.Println(v17, v5, v3)
                }
                _ = fi54
            }
            var v56 uint32 = 2
            var v57 uint8 = 1
            _ = v56
            _ = v57
        }
        _ = fi47
    }
    for fi58 := 0; fi58 < 12; fi58++ {
        for fi59 := 0; fi59 < 11; fi59++ {
            v60 := []int{}
            var v61 uint32 = 65535
            _ = fi59
            _ = v60
            _ = v61
        }
        v62 := GenS5{F0: 1000, F1: false}
        _ = fi58
        _ = v62
    }
    fmt.Println(v4)
    var v63 uint64 = 2
    var v64 string = "\n\t"
    if (!true) {
        for fi65 := 0; fi65 < 6; fi65++ {
            var v66 uint = (v17 ^ 65535)
            _ = fi65
            _ = v66
        }
    }
    v3 += float64(-2.25)
    var v67 uint = (^(v17))
    if (true && false) {
        var v68 uint = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 15; fi69++ {
        var v70 int64 = 36
        v71 := GenS2{F0: v4, F1: 1, F2: genFn0((v67 + v67), v4)}
        _ = fi69
        _ = v70
        _ = v71
    }
    fmt.Println(v64)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v67
}
