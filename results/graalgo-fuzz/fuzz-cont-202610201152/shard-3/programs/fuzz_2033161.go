package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int64; F1 int64; F2 uint }
type GenS6 struct { F0 int; F1 uint16 }
func genFn0(p0 int, p1 uint, p2 uint32) int32 {
    v5 := []int{p0}
    p0 = p0
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 int = 32767
    v2 := make(map[int]uint32)
    var v3 int8 = 0
    v4 := make(map[uint32]string)
    if true {
        v1 = (^(v1))
        if false {
            v6 := GenS1{F0: 42}
            _ = v6
        }
    }
    if false {
        v7 := GenS1{F0: 702}
        var v8 float32 = -2.25
        _ = v7
        _ = v8
    } else {
        fmt.Println(v3)
        for fi9 := 0; fi9 < 12; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                if true {
                    v11 := GenS1{F0: 9223372036854775807}
                    v3 = (v3 >> ((65535) & 63))
                    switch (fi9 - (fi9 | 9223372036854775807)) {
                    case -227:
                        v12 := GenS1{F0: 1}
                        _ = v12
                    default:
                        switch "foo bar" {
                        case "\n\t":
                            var v13 uint64 = 4294967295
                            _ = v13
                        }
                    }
                    _ = v11
                } else {
                    if ((2 > 2) && true) {
                        if (!true) {
                            var v14 int8 = v3
                            v14 = v3
                            switch v1 {
                            case 968:
                                v15 := GenS1{F0: 0}
                                _ = v15
                            case -57:
                                switch "a" {
                                case "hello":
                                    v16 := v4[0]
                                    _ = v16
                                }
                            default:
                                switch "hello" {
                                case "\n\t":
                                    var v17 uint32 = 65535
                                    _ = v17
                                case "hello":
                                    var v18 int = fi10
                                    v19 := GenS2{F0: 612, F1: 4294967295, F2: 454}
                                    switch "a" {
                                    case "":
                                        fmt.Println(fi9, v14)
                                    case "hello":
                                        switch v1 {
                                        case 127:
                                            v20 := GenS1{F0: 1}
                                            _ = v20
                                        }
                                    case "foo bar":
                                        v21 := GenS1{F0: -1}
                                        _ = v21
                                    default:
                                        var v22 uint16 = 341
                                        _ = v22
                                    }
                                    if (913 <= (v1 + v1)) {
                                        fmt.Println(v1, v3, v3)
                                        var v23 uint = 1
                                        _ = v23
                                    } else {
                                        v24 := GenS1{F0: 1}
                                        if true {
                                            var v25 int16 = 1000
                                            _ = v25
                                        } else {
                                            var v26 int = -1
                                            fmt.Println(v14, v18)
                                            var v27 int16 = 0
                                            delete(v4, 4294967295)
                                            var v28 uint = 65535
                                            var v29 bool = true
                                            v30 := GenS1{F0: 9223372036854775807}
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                        _ = v24
                                    }
                                    if true {
                                        switch fi9 {
                                        case -520:
                                            fmt.Println(v14, v18)
                                        case 127:
                                            v31 := []int{(v18 ^ fi9), fi9}
                                            _ = v31
                                        case -86:
                                            var v32 int64 = 0
                                            fmt.Println(v3)
                                            var v33 uint = 255
                                            _ = v32
                                            _ = v33
                                        default:
                                            v14 = (v3 | v14)
                                            var v34 bool = (0.1 < 0.1)
                                            v1 = (fi9 * int(v3))
                                            _ = v34
                                        }
                                    } else {
                                        v35 := []int{(fi10 / ((fi10) | 1)), 9223372036854775807, v18}
                                        var v36 float64 = 1e10
                                        fmt.Println(v1)
                                        _ = v35
                                        _ = v36
                                    }
                                    v18 = 9223372036854775807
                                    switch fi10 {
                                    case 4294967295:
                                        v37 := GenS3{F0: (v3 >> ((65535) & 63))}
                                        _ = v37
                                    case 473:
                                        v38 := GenS1{F0: -1}
                                        fmt.Println(v3, v3, v18)
                                        _ = v38
                                    }
                                    _ = len(v2)
                                    var v39 int = (fi9 << ((877) & 63))
                                    v40 := GenS4{F0: 42, F1: 255}
                                    if (1 <= 18446744073709551615) {
                                        v1 -= (-(v1))
                                        if true {
                                            var v41 float32 = ((1.5 / 1.5) - float32(((0.1 * -2.25) - 0.1)))
                                            var v42 uint = 4294967295
                                            _ = v41
                                            _ = v42
                                        } else {
                                            var v43 int16 = -32768
                                            _ = v43
                                        }
                                        v3 = 127
                                        v44 := []int{(v39 % (((v39 - (v1 & fi10))) | 1)), (v1 | (-(fi9))), v18, fi10}
                                        switch "" {
                                        case "":
                                            if len(v44) >= 2 {
                                                v44 = v44[1:len(v44)]
                                            }
                                        case "\n\t":
                                            v1 += fi10
                                        default:
                                            if (-2.25 > 1e10) {
                                                v3 = -1
                                            } else {
                                                var v45 string = "\n\t"
                                                switch ("\n\t" + ("\n\t" + v45)) {
                                                case "":
                                                    v3 = (^(v3))
                                                case "hello":
                                                    v46 := GenS5{F0: 567, F1: 1, F2: 497}
                                                    _ = v46
                                                }
                                                v4[0] = v45
                                                _ = v45
                                            }
                                        }
                                        _ = v44
                                    } else {
                                        v47 := []int{}
                                        v3 += v14
                                        var v48 int8 = v14
                                        if (!false) {
                                            var v49 uint64 = 18446744073709551615
                                            v3 += (v3 % ((v48) | 1))
                                            v18 -= fi10
                                            _ = v49
                                        } else {
                                            if false {
                                                fmt.Println(v48)
                                                v2[(v18 ^ fi9)] = 4294967295
                                            }
                                            switch (-(v18)) {
                                            case 32767:
                                                v4[65535] = ((("\n\t" + "foo bar") + "a") + "")
                                            default:
                                                switch ("" + ("hello" + "hello")) {
                                                case "\n\t":
                                                    v3 += 42
                                                case "foo bar":
                                                    var v50 uint = 255
                                                    _ = v50
                                                case "":
                                                    v39 += v1
                                                default:
                                                    if false {
                                                        var v51 uint = 65535
                                                        var v52 uint = v51
                                                        v47 = append(v47, (fi9 / ((v18) | 1)))
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                }
                                            }
                                            v18 = (^(v1))
                                        }
                                        _ = v47
                                        _ = v48
                                    }
                                    _ = v18
                                    _ = v19
                                    _ = v39
                                    _ = v40
                                case "":
                                    switch (v1 / ((v1) | 1)) {
                                    case -1:
                                        if false {
                                            var v53 int8 = (v3 << ((65535) & 63))
                                            _ = v53
                                        }
                                    case -2147483648:
                                        switch fi10 {
                                        case 128:
                                            v1 -= fi10
                                            var v54 uint64 = 4294967295
                                            _ = v54
                                        case 32768:
                                            v55 := GenS5{F0: -1, F1: 357, F2: 619}
                                            _ = v55
                                        case 127:
                                            if (42 >= 853) {
                                                v1 = fi9
                                            }
                                        }
                                    default:
                                        v14 += 127
                                    }
                                }
                                v1 = (^(fi10))
                            }
                            _ = v14
                        }
                        switch (fi10 * -2147483648) {
                        case 127:
                            v1 -= 42
                        case 352:
                            v56 := v2[v1]
                            _ = v56
                        default:
                            v57 := v2[(fi10 - (fi9 / (((fi10 | v1)) | 1)))]
                            _ = v57
                        }
                        switch 9223372036854775807 {
                        case -920:
                            switch "" {
                            case "":
                                fmt.Println(v1)
                            }
                        case 0:
                            v3 = v3
                        case 127:
                            v58 := GenS2{F0: 65535, F1: 18446744073709551615, F2: 65535}
                            _ = v58
                        }
                        if true {
                            v1 = v1
                            v3 -= -1
                        } else {
                            switch (v1 ^ fi10) {
                            case 9223372036854775807:
                                fmt.Println(v1)
                            case 32767:
                                v3 = v3
                            case -2147483648:
                                v1 = 145
                                v1 = fi10
                                var v59 uint8 = 42
                                _ = v59
                            }
                        }
                    } else {
                        v60 := GenS6{F0: v1, F1: 65535}
                        _ = v60
                    }
                }
                _ = fi10
            }
            v61 := GenS1{F0: 1}
            var v62 uint32 = 0
            _ = fi9
            _ = v61
            _ = v62
        }
    }
    for fi63 := 0; fi63 < 4; fi63++ {
        v3 = (v3 >> ((65535) & 63))
        v64 := GenS1{F0: 1}
        _ = fi63
        _ = v64
    }
    v65 := GenS3{F0: (v3 | 42)}
    v66 := GenS1{F0: -9223372036854775808}
    if false {
        v67 := GenS4{F0: -1, F1: 947}
        fmt.Println(v3, v1)
        _ = v67
    }
    v68 := []int{}
    v1 = (-(v1))
    var v69 int16 = 1
    var v70 int32 = genFn0(1, 18446744073709551615, 838)
    v70 -= (^(v70))
    var v71 bool = true
    v72 := GenS4{F0: -1, F1: 172}
    v2[(v1 & (^(v1)))] = 965
    v3 = -128
    var v73 int32 = 1
    v74 := GenS2{F0: 32767, F1: 966, F2: 944}
    v69 = (-(v69))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v68:", len(v68), v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v66
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v74
}
