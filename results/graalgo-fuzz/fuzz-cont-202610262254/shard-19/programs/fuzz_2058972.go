package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS2 struct { F0 string; F1 int64; F2 int32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int32; F1 bool }
func genFn0(p0 uint, p1 int16) int16 {
    fmt.Println(p0, p1)
    v6 := []int{0, 128, 2147483647, -878}
    _ = p0
    _ = p1
    _ = v6
    return (p1 ^ -1)
}
func genFn1() int {
    v7 := []int{-149}
    switch -665 {
    case -1:
        var v8 int16 = 1
        _ = v8
    case 2147483647:
        switch 461 {
        case 0:
            var v9 int64 = -49
            _ = v9
        default:
            var v10 uint = 1
            _ = v10
        }
    default:
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 float32 = -2.25
            _ = fi11
            _ = v12
        }
    }
    _ = v7
    return -1
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]bool)
    v3 := []uint32{34, 0, 354}
    v4 := []int{0}
    var v5 int32 = v1
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(fi13, v5)
        var v14 uint8 = 125
        v5 = v5
        _ = fi13
        _ = v14
    }
    v1 -= (-(v1))
    for fi15 := 0; fi15 < 3; fi15++ {
        v1 = 462
        _ = fi15
    }
    fmt.Println(v1)
    v16 := GenS1{F0: 33, F1: 42, F2: 9223372036854775807}
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 float64 = (1e10 * 0.0)
        if (!true) {
            if true {
                var v19 int64 = -606
                _ = v19
            }
        } else {
            var v20 uint = 467
            v21 := GenS2{F0: "\n\t", F1: -811, F2: v5}
            var v22 uint8 = 0
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi17
        _ = v18
    }
    var v23 float64 = 0.0
    var v24 int8 = -1
    v25 := GenS1{F0: 0, F1: 2, F2: -9223372036854775808}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v26 string = "\n\t"
    fmt.Println(v24, v1)
    for fi27 := 0; fi27 < 15; fi27++ {
        switch v26 {
        case "":
            v26 = ((v26 + v26) + ((v26 + ("hello" + (v26 + v26))) + (v26 + ("" + (v26 + "\n\t")))))
        case "foo bar":
            if (1 == 4294967295) {
                var v28 int = fi27
                v24 -= (v24 % (((v24 * v24)) | 1))
                for fi29 := 0; fi29 < 7; fi29++ {
                    var v30 uint8 = 204
                    switch v26 {
                    case "foo bar":
                        var v31 string = v26
                        switch (v26 + ((((v31 + "a") + "foo bar") + v31) + v31)) {
                        case "foo bar":
                            var v32 int64 = 1
                            _ = v32
                        case "":
                            _ = len(v2)
                        case "hello":
                            var v33 uint32 = 567
                            _ = v33
                        default:
                            v5 = v1
                        }
                        switch (v28 * 4294967295) {
                        case -1:
                            switch 42 {
                            case -1:
                                var v34 float64 = v23
                                _ = v34
                            default:
                                if true {
                                    var v35 bool = (true && false)
                                    v3 = append(v3, 378)
                                    v36 := GenS2{F0: v26, F1: 42, F2: v1}
                                    delete(v2, 1)
                                    fmt.Println(v24, v31, v31)
                                    fmt.Println(fi27, v26, v23)
                                    if (1e10 != (v23 - float64(((float32(1.5) + float32((1.5 + v23))) / -2.25)))) {
                                        var v37 float64 = (v23 * 0.0)
                                        if len(v3) > 0 {
                                            v3[0] = 2
                                        }
                                        v38 := GenS3{F0: (v30 + v30), F1: int64(v28)}
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = len(v3)
                                    _ = v35
                                    _ = v36
                                } else {
                                    var v39 float64 = 1.5
                                    fmt.Println(v1, v31, fi27)
                                    v40 := GenS2{F0: (v31 + v26), F1: 9223372036854775807, F2: (v5 >> ((42) & 63))}
                                    _ = v39
                                    _ = v40
                                }
                            }
                        }
                        _ = v31
                    }
                    var v41 int64 = -125
                    var v42 uint32 = 2147483648
                    var v43 string = "\n\t"
                    _ = fi29
                    _ = v30
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = v28
            } else {
                var v44 int32 = 1
                v45 := GenS4{F0: (v5 & 736), F1: false}
                v26 = ((v26 + v26) + ((((((v26 + v26) + "a") + v26) + v26) + ("" + "hello")) + "foo bar"))
                _ = v44
                _ = v45
            }
        default:
            var v46 uint64 = 18446744073709551615
            _ = v46
        }
        for fi47 := 0; fi47 < 5; fi47++ {
            v48 := GenS4{F0: (v1 + v5), F1: true}
            delete(v2, 1)
            if (827 >= 32767) {
                _ = len(v2)
                switch (fi27 << ((17) & 63)) {
                case 245:
                    var v49 int = fi27
                    if false {
                        if (4294967295 != 1) {
                            var v50 int64 = 42
                            v51 := GenS1{F0: 2, F1: 42, F2: 0}
                            _ = v50
                            _ = v51
                        }
                    }
                    _ = len(v3)
                    _ = v49
                case 294:
                    switch 864 {
                    case 527:
                        fmt.Println(fi27)
                    case 128:
                        v23 -= 1e10
                    default:
                        switch fi27 {
                        case 32768:
                            if (false || ((float32(((v23 - (v23 - v23)) - float64(3.14159))) != 1.5) || (13 <= 1))) {
                                var v52 int = (fi47 - (fi27 << ((4294967295) & 63)))
                                var v53 int32 = (v1 * (v1 - 1))
                                _ = v52
                                _ = v53
                            }
                        }
                    }
                default:
                    var v54 float64 = v23
                    _ = v54
                }
                v26 = ("hello" + "a")
            }
            _ = fi47
            _ = v48
        }
        v26 = "\n\t"
        _ = fi27
    }
    var v55 int = 32768
    if false {
        if (v5 < 0) {
            var v56 float32 = (3.14159 + 1e10)
            if false {
                v23 += v23
            }
            var v57 int64 = -855
            _ = v56
            _ = v57
        } else {
            var v58 uint16 = 2
            var v59 float32 = (0.1 + ((((1.5 * ((1e10 * (0.1 / 1e10)) - ((float32(v23) - 0.1) / 1e10))) / (3.14159 / (3.14159 + (0.1 / -2.25)))) / 1.0) / 1.5))
            switch ("hello" + v26) {
            case "\n\t":
                v58 -= v58
            case "hello":
                var v60 uint16 = (^(v58))
                _ = v60
            }
            _ = v58
            _ = v59
        }
        if (v23 >= (float64((1e10 + ((1.5 / 1.0) + (3.14159 - float32(v23))))) / 1.0)) {
            for fi61 := 0; fi61 < 11; fi61++ {
                fmt.Println(v24, v5, v24)
                _ = fi61
            }
            v62 := GenS2{F0: v26, F1: -113, F2: v5}
            _ = v62
        } else {
            var v63 float32 = float32(v23)
            fmt.Println(v55, v23)
            _ = v63
        }
        switch genFn1() {
        case 33:
            if false {
                if len(v4) > 0 {
                    v4[0] = (v55 & (v55 << ((42) & 63)))
                }
                delete(v2, 2147483648)
                var v64 uint8 = 173
                _ = v64
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi65 := 0; fi65 < 14; fi65++ {
        v66 := GenS1{F0: 2, F1: 512, F2: 1}
        _ = fi65
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v55
}
