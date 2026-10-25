package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int; F1 string; F2 string }
func genFn0(p0 uint32, p1 uint16) uint16 {
    for fi4 := 0; fi4 < 6; fi4++ {
        if false {
            var v5 uint8 = 1
            _ = v5
        } else {
            v6 := GenS1{F0: -1}
            switch "a" {
            case "hello":
                if (42 == 52) {
                    if (1 > 3) {
                        for fi7 := 0; fi7 < 11; fi7++ {
                            var v8 float64 = 1e10
                            v9 := GenS1{F0: 42}
                            fmt.Println(v8, v8)
                            _ = fi7
                            _ = v8
                            _ = v9
                        }
                        for fi10 := 0; fi10 < 18; fi10++ {
                            fmt.Println(p1, fi10)
                            v11 := []int{232}
                            if ((1.5 + -2.25) == (0.1 - float64(3.14159))) {
                                switch ("" + "a") {
                                case "\n\t":
                                    p1 = 2
                                case "hello":
                                    v12 := GenS1{F0: -1}
                                    _ = v12
                                case "":
                                    fmt.Println(fi4)
                                default:
                                    p1 += p1
                                }
                                var v13 int16 = 0
                                if ((false || false) || true) {
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    if ("\n\t" < "hello") {
                                        v14 := GenS2{F0: uint16(p0), F1: 59}
                                        fmt.Println(p1, p1, p1)
                                        var v15 float64 = (-2.25 - (0.1 + (-2.25 + 0.1)))
                                        v16 := GenS2{F0: (p1 >> ((42) & 63)), F1: 100}
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    }
                                    v17 := GenS2{F0: 65535, F1: -20}
                                    _ = v17
                                }
                                p0 = (^(p0))
                                _ = v13
                            } else {
                                if ((!(!true)) || (0 > 2147483647)) {
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    fmt.Println(p1, p0, fi4)
                                    var v18 uint8 = 89
                                    if (false && (0 == (p0 % (((p0 & p0)) | 1)))) {
                                        var v19 bool = (60 < 127)
                                        v19 = false
                                        _ = v19
                                    }
                                    _ = v18
                                } else {
                                    if false {
                                        fmt.Println(fi10)
                                        if (0 <= p1) {
                                            var v20 int16 = -1
                                            if (!true) {
                                                var v21 int32 = 323
                                                v22 := GenS3{F0: (v21 << ((1) & 63))}
                                                _ = v21
                                                _ = v22
                                            }
                                            _ = v20
                                        } else {
                                            var v23 uint64 = 553
                                            _ = v23
                                        }
                                    }
                                    if (((0 <= int16(p1)) || ((fi10 * fi4) >= fi4)) || (!(!(!true)))) {
                                        p1 = (p1 >> ((255) & 63))
                                        var v24 float32 = (-2.25 * 0.1)
                                        _ = v24
                                    }
                                    _ = len(v11)
                                    v25 := GenS4{F0: -1, F1: -128}
                                    switch "a" {
                                    case "foo bar":
                                        v26 := GenS1{F0: -1}
                                        _ = v26
                                    }
                                    var v27 float32 = 0.1
                                    _ = v25
                                    _ = v27
                                }
                                p1 = (p1 * (^(p1)))
                            }
                            _ = fi10
                            _ = v11
                        }
                        for fi28 := 0; fi28 < 17; fi28++ {
                            v29 := []int{fi4, (fi4 >> ((2) & 63)), int(p1), (fi28 - (-(fi28)))}
                            fmt.Println(fi28, fi28, p1)
                            if true {
                                v30 := GenS5{F0: 42, F1: 0}
                                _ = v30
                            } else {
                                switch "" {
                                case "":
                                    if false {
                                        switch 128 {
                                        case -1:
                                            var v31 uint16 = (p1 ^ (p1 + p1))
                                            _ = v31
                                        }
                                    } else {
                                        fmt.Println(p1, fi28)
                                        p0 += (p0 * 2)
                                    }
                                case "\n\t":
                                    var v32 int16 = 1000
                                    _ = v32
                                }
                            }
                            var v33 uint64 = 1
                            v34 := GenS1{F0: 1000}
                            var v35 float64 = 1.5
                            v29 = append(v29, (fi28 & (fi4 % (((fi28 ^ (^(fi4)))) | 1))))
                            _ = fi28
                            _ = v29
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    } else {
                        var v36 int = (fi4 - (-(fi4)))
                        v37 := GenS5{F0: 128, F1: -1}
                        _ = v36
                        _ = v37
                    }
                    var v38 uint32 = p0
                    _ = v38
                } else {
                    if (!false) {
                        if true {
                            v39 := []int{}
                            p0 = p0
                            _ = v39
                        }
                        if true {
                            var v40 uint64 = 42
                            _ = v40
                        }
                    }
                }
            case "foo bar":
                var v41 uint16 = (^(p1))
                _ = v41
            case "a":
                v42 := GenS2{F0: (^(p1)), F1: 57}
                _ = v42
            default:
                p0 = 0
            }
            for fi43 := 0; fi43 < 19; fi43++ {
                p0 += p0
                var v44 uint16 = p1
                _ = fi43
                _ = v44
            }
            _ = v6
        }
        _ = fi4
    }
    var v45 uint64 = 2
    _ = p0
    _ = p1
    _ = v45
    return (p1 + (p1 << ((65535) & 63)))
}
func genFn1(p0 int8) int8 {
    if true {
        var v46 float64 = (3.14159 + 1.5)
        if (true && (!true)) {
            var v47 float64 = (v46 / v46)
            var v48 int16 = 42
            _ = v47
            _ = v48
        } else {
            if true {
                fmt.Println(v46, p0)
                var v49 float64 = 0.1
                _ = v49
            } else {
                fmt.Println(v46, v46)
                v50 := GenS2{F0: genFn0(526, 336), F1: p0}
                _ = v50
            }
        }
        _ = v46
    }
    v51 := []int{-2147483648, 155, -1, 419}
    v52 := GenS2{F0: 2, F1: (p0 / (((p0 - p0)) | 1))}
    p0 += 127
    p0 -= 0
    _ = p0
    _ = v51
    _ = v52
    return -128
}

func main() {
    v1 := []int{-496, 454}
    v2 := make(map[uint32]string)
    var v3 int8 = 1
    if (2 == 42) {
        if false {
            for fi53 := 0; fi53 < 17; fi53++ {
                for fi54 := 0; fi54 < 13; fi54++ {
                    var v55 uint16 = 65535
                    _ = fi54
                    _ = v55
                }
                for fi56 := 0; fi56 < 8; fi56++ {
                    if (1e10 < 3.14159) {
                        switch (fi56 % ((42) | 1)) {
                        case -1:
                            var v57 int8 = (v3 % ((82) | 1))
                            _ = v57
                        case 959:
                            var v58 float32 = (0.0 - float32(3.14159))
                            _ = v58
                        case 2147483647:
                            v3 = v3
                        }
                        if (4294967295 <= 600) {
                            if (false || false) {
                                switch "\n\t" {
                                case "hello":
                                    v3 = v3
                                }
                            }
                            var v59 bool = (257 > 2)
                            _ = v59
                        } else {
                            if false {
                                var v60 string = ""
                                var v61 int64 = 1
                                _ = v60
                                _ = v61
                            }
                        }
                    } else {
                        if true {
                            v62 := GenS5{F0: 255, F1: int16(fi56)}
                            var v63 int = (^(fi56))
                            v64 := GenS5{F0: 1, F1: -1}
                            _ = v62
                            _ = v63
                            _ = v64
                        } else {
                            var v65 int8 = (v3 * v3)
                            v66 := v2[2147483647]
                            _ = v65
                            _ = v66
                        }
                        var v67 float64 = 1e10
                        switch "foo bar" {
                        case "hello":
                            var v68 uint32 = 1
                            var v69 bool = true
                            _ = v68
                            _ = v69
                        case "":
                            var v70 uint8 = 2
                            switch ((("hello" + "a") + "") + ("foo bar" + "foo bar")) {
                            case "a":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "\n\t":
                                v3 = (v3 / ((genFn1(v3)) | 1))
                            case "foo bar":
                                v71 := v2[4294967295]
                                _ = v71
                            }
                            _ = v70
                        case "foo bar":
                            var v72 int8 = int8(fi53)
                            _ = v72
                        }
                        var v73 uint = 83
                        _ = v67
                        _ = v73
                    }
                    var v74 uint32 = 42
                    v75 := GenS6{F0: 128, F1: "\n\t", F2: ("foo bar" + "a")}
                    _ = fi56
                    _ = v74
                    _ = v75
                }
                var v76 int64 = -74
                _ = fi53
                _ = v76
            }
            delete(v2, 555)
        }
        v77 := v2[2147483648]
        _ = v77
    }
    for fi78 := 0; fi78 < 6; fi78++ {
        switch "\n\t" {
        case "a":
            var v79 uint32 = 0
            _ = v79
        case "hello":
            v80 := GenS6{F0: (-(fi78)), F1: "hello", F2: "a"}
            _ = v80
        case "\n\t":
            v3 = (v3 << ((675) & 63))
        }
        _ = fi78
    }
    switch 128 {
    case 32768:
        if (false || false) {
            v81 := GenS3{F0: 2147483647}
            _ = len(v1)
            _ = v81
        }
    default:
        v82 := GenS1{F0: 0}
        _ = v82
    }
    switch -396 {
    case 543:
        switch "\n\t" {
        case "\n\t":
            for fi83 := 0; fi83 < 14; fi83++ {
                var v84 uint32 = 2
                var v85 float32 = 1.5
                var v86 float64 = 1e10
                _ = fi83
                _ = v84
                _ = v85
                _ = v86
            }
        }
    case 1:
        var v87 bool = (!false)
        _ = v87
    case -549:
        v3 = 42
    }
    var v88 uint8 = 70
    if len(v1) > 0 {
        v1[0] = 127
    }
    v88 = v88
    v88 += 255
    var v89 float32 = float32(0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v88
    _ = v89
}
