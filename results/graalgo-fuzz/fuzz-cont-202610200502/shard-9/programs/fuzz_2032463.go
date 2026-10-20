package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 bool }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16 }
func genFn0(p0 int64) int64 {
    v6 := GenS1{F0: 128, F1: -1, F2: (!(false && (!true)))}
    switch "hello" {
    case "foo bar":
        switch 767 {
        case 853:
            v7 := []int{-513, 1, 32767}
            _ = v7
        default:
            v8 := []int{int(p0), 481}
            _ = v8
        }
    case "\n\t":
        var v9 uint32 = 402
        _ = v9
    case "a":
        switch (("\n\t" + "hello") + "\n\t") {
        case "\n\t":
            for fi10 := 0; fi10 < 19; fi10++ {
                v11 := GenS2{F0: 53, F1: 255}
                v12 := []int{fi10}
                var v13 uint64 = 150
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        case "hello":
            p0 += 42
        }
    }
    if (4294967295 <= 97) {
        p0 = p0
    }
    var v14 uint = 42
    _ = p0
    _ = v14
    _ = v6
    return p0
}

func main() {
    var v1 uint8 = 170
    var v2 int16 = 42
    var v3 float64 = ((-2.25 + 1e10) / -2.25)
    var v4 uint = 42
    var v5 uint32 = 42
    v1 = v1
    v3 += ((v3 / 0.1) + v3)
    var v15 int16 = v2
    v4 -= 2
    var v16 uint = v4
    var v17 int32 = 1
    v2 += v15
    if false {
        v1 = v1
        for fi18 := 0; fi18 < 8; fi18++ {
            if (109 <= 18446744073709551615) {
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 int32 = (v17 ^ v17)
                    _ = fi19
                    _ = v20
                }
            }
            _ = fi18
        }
        var v21 int = -690
        _ = v21
    } else {
        var v22 int8 = 73
        _ = v22
    }
    if false {
        fmt.Println(v15, v2, v15)
        var v23 uint16 = 255
        var v24 int = -9223372036854775808
        _ = v23
        _ = v24
    } else {
        switch "foo bar" {
        case "a":
            v25 := GenS3{F0: -9223372036854775808, F1: v17}
            _ = v25
        case "hello":
            v26 := []int{9223372036854775807, int(v1), -9223372036854775808, 543}
            _ = v26
        case "":
            for fi27 := 0; fi27 < 8; fi27++ {
                if true {
                    v3 = v3
                    v15 = (-(v15))
                    v28 := GenS1{F0: v1, F1: 1, F2: ((v17 ^ v17) != (v17 ^ (v17 & v17)))}
                    _ = v28
                }
                if (true && false) {
                    switch fi27 {
                    case -1:
                        if true {
                            var v29 string = "foo bar"
                            v30 := []int{fi27, 509, (fi27 * fi27), fi27}
                            if false {
                                var v31 bool = (!false)
                                if ((v5 * v5) == (v5 % ((v5) | 1))) {
                                    switch ("a" + "\n\t") {
                                    case "\n\t":
                                        v1 -= 2
                                    }
                                } else {
                                    if ((v31 && (949 != 617)) || v31) {
                                        switch (v29 + "hello") {
                                        case "":
                                            var v32 int8 = 127
                                            _ = v32
                                        case "hello":
                                            if v31 {
                                                v15 = 1000
                                            }
                                        case "a":
                                            for fi33 := 0; fi33 < 19; fi33++ {
                                                if len(v30) > 0 {
                                                    v30[0] = fi27
                                                }
                                                _ = fi33
                                            }
                                        default:
                                            for fi34 := 0; fi34 < 7; fi34++ {
                                                if true {
                                                    if (fi34 <= fi34) {
                                                        if ((-(v2)) < 0) {
                                                            var v35 float64 = (3.14159 - float64(((0.0 - 3.14159) * float32(1e10))))
                                                            var v36 int64 = -1
                                                            var v37 int16 = v2
                                                            _ = v35
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            v31 = v31
                                                            v38 := GenS4{F0: 42, F1: v29}
                                                            if len(v30) > 0 {
                                                                v30[0] = (fi34 << ((v16) & 63))
                                                            }
                                                            _ = v38
                                                        }
                                                        v30 = append(v30, int(v17))
                                                        var v39 uint16 = 2
                                                        _ = v39
                                                    } else {
                                                        if len(v30) >= 2 {
                                                            v30 = v30[1:len(v30)]
                                                        }
                                                    }
                                                    if v31 {
                                                        var v40 float32 = (((float32(v3) + 1e10) * (1.5 * (1e10 / 1.5))) / 3.14159)
                                                        switch v29 {
                                                        case "hello":
                                                            if false {
                                                                switch "" {
                                                                case "hello":
                                                                    v4 = v16
                                                                case "":
                                                                    if (uint8(v4) == (v1 >> ((v16) & 63))) {
                                                                        var v41 bool = v31
                                                                        _ = v41
                                                                    }
                                                                case "foo bar":
                                                                    var v42 int8 = 42
                                                                    _ = v42
                                                                default:
                                                                    if len(v30) > 0 {
                                                                        v30[0] = -661
                                                                    }
                                                                }
                                                                fmt.Println(v31, v29)
                                                                switch (v29 + v29) {
                                                                case "hello":
                                                                    _ = len(v30)
                                                                case "\n\t":
                                                                    v43 := GenS4{F0: -1, F1: ""}
                                                                    _ = v43
                                                                }
                                                                fmt.Println(v1, v3, v40)
                                                            } else {
                                                                if v31 {
                                                                    switch fi27 {
                                                                    case -553:
                                                                        if (!true) {
                                                                            v3 = float64(v40)
                                                                            _ = len(v30)
                                                                        }
                                                                        v4 += (v16 / ((173) | 1))
                                                                        _ = len(v30)
                                                                    }
                                                                    fmt.Println(v4, v4)
                                                                    v29 = "a"
                                                                } else {
                                                                    v15 = (v2 % (((v2 - v2)) | 1))
                                                                    v15 = (v15 + v2)
                                                                    v29 = v29
                                                                    if false {
                                                                        var v44 int64 = genFn0(-885)
                                                                        _ = v44
                                                                    } else {
                                                                        switch (fi34 | 0) {
                                                                        case 1:
                                                                            var v45 bool = v31
                                                                            _ = v45
                                                                        case 2147483647:
                                                                            var v46 uint8 = v1
                                                                            _ = v46
                                                                        }
                                                                    }
                                                                    var v47 int = (-(fi34))
                                                                    var v48 int16 = 1000
                                                                    _ = v47
                                                                    _ = v48
                                                                }
                                                            }
                                                        case "":
                                                            fmt.Println(v40, v3)
                                                        case "a":
                                                            v1 -= uint8(v5)
                                                        default:
                                                            v5 = v5
                                                        }
                                                        v49 := GenS5{F0: 0}
                                                        _ = v40
                                                        _ = v49
                                                    } else {
                                                        fmt.Println(v31, v3)
                                                        fmt.Println(v29)
                                                    }
                                                    v30 = append(v30, (fi34 * fi34))
                                                } else {
                                                    v50 := GenS3{F0: 0, F1: (^(v17))}
                                                    _ = v50
                                                }
                                                _ = fi34
                                            }
                                        }
                                        if v31 {
                                            v4 = 470
                                        } else {
                                            if false {
                                                switch 32767 {
                                                case -132:
                                                    fmt.Println(v29, v16, v1)
                                                }
                                                if false {
                                                    fmt.Println(v5)
                                                } else {
                                                    v2 = 1000
                                                }
                                            } else {
                                                var v51 uint64 = 67
                                                for fi52 := 0; fi52 < 11; fi52++ {
                                                    switch (fi27 + (fi52 + -9223372036854775808)) {
                                                    case 32768:
                                                        v53 := GenS6{F0: 42}
                                                        _ = v53
                                                    }
                                                    _ = fi52
                                                }
                                                _ = v51
                                            }
                                            var v54 float32 = ((1.5 * 1e10) - 3.14159)
                                            _ = v54
                                        }
                                    }
                                    if (464 <= (v17 / (((-(v17))) | 1))) {
                                        v30 = append(v30, (fi27 << (((v16 ^ v16)) & 63)))
                                    } else {
                                        var v55 int16 = v2
                                        _ = v55
                                    }
                                }
                                _ = v31
                            }
                            _ = v29
                            _ = v30
                        } else {
                            for fi56 := 0; fi56 < 8; fi56++ {
                                if true {
                                    switch (fi27 + fi56) {
                                    case 0:
                                        v57 := []int{}
                                        v2 -= v15
                                        _ = v57
                                    default:
                                        var v58 int32 = v17
                                        _ = v58
                                    }
                                } else {
                                    var v59 int8 = 71
                                    v60 := GenS6{F0: 76}
                                    v15 -= 1000
                                    var v61 uint16 = 42
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                _ = fi56
                            }
                            fmt.Println(v15, v5)
                        }
                    default:
                        v62 := []int{}
                        _ = v62
                    }
                    var v63 int16 = v2
                    v4 += (v4 & 65535)
                    _ = v63
                }
                _ = fi27
            }
        }
    }
    var v64 bool = (802 > 0)
    var v65 bool = v64
    v1 = uint8(v5)
    var v66 string = "foo bar"
    v67 := []int{-602, 0, 576}
    var v68 int64 = -360
    fmt.Println(v5, v4)
    v69 := GenS1{F0: (v1 / ((255) | 1)), F1: 91, F2: (!("foo bar" == v66))}
    for fi70 := 0; fi70 < 16; fi70++ {
        v5 += (v5 << (((v16 % (((v16 / (((v4 << (((v4 & 0)) & 63))) | 1))) | 1))) & 63))
        var v71 int8 = 0
        _ = fi70
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
    _ = v69
}
