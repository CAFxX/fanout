package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int8; F2 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32; F1 string; F2 uint32 }
type GenS9 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS10 struct { F0 int8; F1 int8; F2 uint32 }
type GenS11 struct { F0 uint64; F1 int8 }
type GenS12 struct { F0 string; F1 int; F2 uint32 }
type GenS13 struct { F0 int; F1 uint; F2 int32 }
type GenS14 struct { F0 uint16; F1 uint64 }
func genFn0(p0 bool) uint {
    p0 = p0
    if (394 != 0) {
        var v5 uint32 = 437
        _ = v5
    } else {
        v6 := []int{}
        var v7 int64 = -1
        _ = v6
        _ = v7
    }
    var v8 uint64 = 0
    if (!(-997 >= 1)) {
        if (4294967295 < 850) {
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 int64 = 42
                fmt.Println(fi9)
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                if p0 {
                    if p0 {
                        var v12 string = "foo bar"
                        _ = v12
                    } else {
                        var v13 uint16 = 255
                        p0 = (!(!true))
                        switch "hello" {
                        case "foo bar":
                            var v14 float64 = 0.1
                            _ = v14
                        case "":
                            v15 := GenS1{F0: 124}
                            _ = v15
                        case "hello":
                            var v16 uint8 = 170
                            _ = v16
                        default:
                            for fi17 := 0; fi17 < 9; fi17++ {
                                v18 := GenS2{F0: fi11}
                                _ = fi17
                                _ = v18
                            }
                        }
                        _ = v13
                    }
                    var v19 uint32 = 0
                    _ = v19
                } else {
                    fmt.Println(v8)
                    if true {
                        for fi20 := 0; fi20 < 1; fi20++ {
                            if (float64(1e10) < (float64(1.5) - 1.5)) {
                                v8 = 2
                            } else {
                                fmt.Println(fi20)
                                v8 -= (v8 >> ((2) & 63))
                                p0 = (127 < 1)
                                if true {
                                    v21 := GenS3{F0: -1, F1: 42, F2: 42}
                                    p0 = (1e10 > 0.0)
                                    _ = v21
                                } else {
                                    fmt.Println(fi11, p0)
                                    v22 := []int{}
                                    v23 := GenS2{F0: fi11}
                                    _ = v22
                                    _ = v23
                                }
                            }
                            v24 := []int{}
                            _ = fi20
                            _ = v24
                        }
                        p0 = false
                    } else {
                        fmt.Println(fi11, fi11, fi11)
                        v25 := []int{fi11, fi11}
                        _ = v25
                    }
                }
                if true {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v27 := GenS4{F0: -9223372036854775808, F1: v8}
                        v8 += v8
                        if p0 {
                            switch "foo bar" {
                            case "hello":
                                var v28 uint32 = 60
                                _ = v28
                            case "a":
                                v8 = v8
                                var v29 string = ""
                                var v30 uint32 = 650
                                if p0 {
                                    if (uint16(v8) <= 757) {
                                        p0 = p0
                                        v31 := []int{}
                                        switch v29 {
                                        case "":
                                            v32 := GenS5{F0: 42, F1: 42, F2: 1}
                                            _ = v32
                                        case "hello":
                                            if false {
                                                p0 = (!(255 > uint8(v30)))
                                                fmt.Println(v30, v30, v30)
                                            }
                                        default:
                                            if len(v31) >= 2 {
                                                v31 = v31[1:len(v31)]
                                            }
                                        }
                                        _ = v31
                                    } else {
                                        v33 := GenS6{F0: 1}
                                        switch 1 {
                                        case -2147483648:
                                            var v34 uint8 = 128
                                            _ = v34
                                        }
                                        _ = v33
                                    }
                                }
                                v35 := GenS7{F0: -1}
                                var v36 int64 = 0
                                _ = v29
                                _ = v30
                                _ = v35
                                _ = v36
                            case "":
                                if p0 {
                                    switch ("\n\t" + ("" + "")) {
                                    case "foo bar":
                                        p0 = (-32768 != 1)
                                    case "\n\t":
                                        v37 := GenS7{F0: 1}
                                        p0 = p0
                                        _ = v37
                                    case "":
                                        v38 := []int{}
                                        _ = v38
                                    default:
                                        switch (fi26 - 0) {
                                        case 1:
                                            v39 := GenS8{F0: 72, F1: "foo bar", F2: 312}
                                            _ = v39
                                        case -440:
                                            v8 -= v8
                                        default:
                                            v40 := GenS9{F0: 273, F1: 0, F2: 2147483647}
                                            _ = v40
                                        }
                                    }
                                    v8 = v8
                                    if (205 <= 128) {
                                        if p0 {
                                            v41 := GenS2{F0: 42}
                                            v8 = (v8 ^ v8)
                                            _ = v41
                                        } else {
                                            v42 := GenS10{F0: -72, F1: -1, F2: 1}
                                            var v43 int32 = 2147483647
                                            var v44 bool = true
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                    } else {
                                        v8 = 597
                                        v45 := GenS11{F0: v8, F1: 100}
                                        _ = v45
                                    }
                                }
                            }
                        }
                        switch "hello" {
                        case "hello":
                            v8 += v8
                        case "":
                            v46 := GenS2{F0: 2147483647}
                            _ = v46
                        default:
                            v47 := GenS12{F0: ("hello" + "foo bar"), F1: fi26, F2: 65}
                            _ = v47
                        }
                        _ = fi26
                        _ = v27
                    }
                    v48 := []int{fi11}
                    _ = v48
                } else {
                    if p0 {
                        v49 := []int{(-(fi11)), fi11}
                        var v50 int64 = 616
                        _ = v49
                        _ = v50
                    }
                    v51 := GenS2{F0: 9223372036854775807}
                    fmt.Println(v8)
                    _ = v51
                }
                if false {
                    v52 := []int{(fi11 + int(v8))}
                    v53 := GenS13{F0: fi11, F1: 803, F2: 42}
                    _ = v52
                    _ = v53
                } else {
                    for fi54 := 0; fi54 < 8; fi54++ {
                        var v55 int = fi11
                        var v56 uint = 65535
                        v55 = (-(fi11))
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    v57 := GenS14{F0: 851, F1: v8}
                    v58 := GenS5{F0: 128, F1: -1, F2: 0}
                    _ = v57
                    _ = v58
                }
                _ = fi11
            }
            fmt.Println(p0)
        } else {
            var v59 int32 = 42
            for fi60 := 0; fi60 < 15; fi60++ {
                v61 := GenS9{F0: 200, F1: v59, F2: 47}
                _ = fi60
                _ = v61
            }
            _ = v59
        }
        if (-2147483648 < 2147483647) {
            p0 = true
        } else {
            var v62 int64 = -9223372036854775808
            var v63 uint64 = (^(v8))
            _ = v62
            _ = v63
        }
    }
    p0 = (p0 && false)
    _ = p0
    _ = v8
    return 430
}
func genFn1() int32 {
    var v64 bool = (false && false)
    fmt.Println(v64)
    _ = v64
    return -861
}

func main() {
    v1 := []uint32{2147483647, 2, 2147483648, 238}
    var v2 uint32 = 43
    var v3 uint32 = 65535
    v4 := []string{"\n\t", "hello", "foo bar"}
    fmt.Println(v2)
    v2 -= v3
    var v65 int8 = -128
    if false {
        var v66 uint = 2
        _ = v66
    }
    var v67 bool = true
    for fi68 := 0; fi68 < 12; fi68++ {
        var v69 int8 = v65
        fmt.Println(v67)
        _ = fi68
        _ = v69
    }
    for fi70 := 0; fi70 < 8; fi70++ {
        v65 += v65
        _ = fi70
    }
    var v71 int16 = 32767
    for fi72 := 0; fi72 < 17; fi72++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi72
    }
    switch -9223372036854775808 {
    case 127:
        var v73 uint64 = 772
        _ = v73
    default:
        v67 = (18446744073709551615 > 986)
    }
    var v74 int16 = (v71 & (v71 & v71))
    var v75 int32 = 0
    for fi76 := 0; fi76 < 6; fi76++ {
        v67 = v67
        if false {
            var v77 string = (("" + "hello") + "hello")
            v74 = (v74 << ((456) & 63))
            _ = v77
        } else {
            if v67 {
                v67 = ((((v2 * v3) == 2147483648) || (!false)) || false)
                if (v74 <= (v71 + v71)) {
                    v2 = 65535
                }
                fmt.Println(v67, v74, v3)
            } else {
                v78 := GenS11{F0: 777, F1: 1}
                v79 := GenS1{F0: 42}
                _ = v78
                _ = v79
            }
        }
        _ = fi76
    }
    for fi80 := 0; fi80 < 12; fi80++ {
        fmt.Println(v67, v65, v65)
        v75 = (v75 / ((0) | 1))
        v67 = v67
        _ = fi80
    }
    v3 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v65:", v65)
    fmt.Println("v67:", v67)
    fmt.Println("v71:", v71)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v67
    _ = v71
    _ = v74
    _ = v75
}
