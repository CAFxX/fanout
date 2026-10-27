package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint; F1 uint8; F2 uint }
type GenS5 struct { F0 uint32; F1 string; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint; F2 string }
type GenS7 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS8 struct { F0 bool; F1 uint8 }
type GenS9 struct { F0 bool; F1 string; F2 uint }
type GenS10 struct { F0 int; F1 uint32 }
func genFn0() int {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{fi6, fi6, (fi6 ^ (^(fi6))), fi6}
        if true {
            var v8 int64 = 0
            var v9 uint = 65535
            v9 = v9
            _ = v8
            _ = v9
        } else {
            var v10 uint = 4294967295
            switch "foo bar" {
            case "a":
                var v11 uint = 2
                _ = v11
            default:
                var v12 int = 42
                _ = v12
            }
            _ = v10
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint8 = 56
            var v15 uint64 = 0
            if (!((-71 <= -1) || (fi6 < (fi6 | -610)))) {
                fmt.Println(v15)
                fmt.Println(v15)
                switch 32768 {
                case -93:
                    switch 32767 {
                    case 32768:
                        switch (fi6 << ((18446744073709551615) & 63)) {
                        case 32768:
                            if true {
                                var v16 int = 127
                                v17 := GenS1{F0: 212}
                                var v18 int = (fi13 | (^(v16)))
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                        }
                    default:
                        var v19 int16 = 32767
                        _ = v19
                    }
                case -658:
                    fmt.Println(fi13, v15, v15)
                    var v20 bool = (!(0.1 >= -2.25))
                    switch (fi6 + fi6) {
                    case 127:
                        if (0 >= -2147483648) {
                            v15 -= (v15 % (((v15 & (v15 % (((v15 % (((v15 % ((248) | 1))) | 1))) | 1)))) | 1))
                        } else {
                            var v21 float64 = (0.0 * 1e10)
                            v21 += float64((-2.25 / (1e10 + 1e10)))
                            if len(v7) > 0 {
                                v7[0] = fi6
                            }
                            _ = v21
                        }
                    case 4294967295:
                        var v22 int = fi6
                        _ = v22
                    case 0:
                        var v23 int32 = 0
                        _ = v23
                    }
                    _ = v20
                }
                v24 := GenS1{F0: 2147483647}
                var v25 uint16 = 255
                v26 := GenS2{F0: -32768}
                _ = v24
                _ = v25
                _ = v26
            } else {
                v14 = v14
                v27 := GenS3{F0: "foo bar", F1: (-32768 < -88)}
                if (((1 < 42) && (!((("\n\t" + "a") + "hello") == ""))) && (42 <= 186)) {
                    v14 -= v14
                    v7 = append(v7, fi6)
                    var v28 float32 = (float32(1e10) / -2.25)
                    _ = v28
                }
                v15 += 92
                v29 := GenS3{F0: "foo bar", F1: (-1 >= 32767)}
                if ((!false) && (0 < 2)) {
                    var v30 int8 = -60
                    var v31 uint8 = (v14 - (v14 % ((v14) | 1)))
                    _ = v30
                    _ = v31
                } else {
                    v32 := GenS4{F0: 4294967295, F1: (^(v14)), F2: 984}
                    var v33 int16 = 160
                    if len(v7) > 0 {
                        v7[0] = fi13
                    }
                    _ = v32
                    _ = v33
                }
                _ = v27
                _ = v29
            }
            v34 := GenS1{F0: 0}
            _ = fi13
            _ = v14
            _ = v15
            _ = v34
        }
        _ = fi6
        _ = v7
    }
    var v35 uint32 = 891
    switch ("a" + ("" + "foo bar")) {
    case "a":
        if true {
            for fi36 := 0; fi36 < 2; fi36++ {
                var v37 uint16 = 42
                _ = fi36
                _ = v37
            }
            var v38 uint8 = 42
            var v39 int16 = 1
            _ = v38
            _ = v39
        }
    case "hello":
        switch 32767 {
        case 855:
            v40 := []int{}
            _ = v40
        case 32767:
            var v41 uint64 = 401
            _ = v41
        }
    case "\n\t":
        if true {
            v42 := GenS3{F0: "hello", F1: true}
            fmt.Println(v35)
            if (true || false) {
                fmt.Println(v35)
                var v43 uint = 18446744073709551615
                _ = v43
            } else {
                var v44 uint = 886
                for fi45 := 0; fi45 < 1; fi45++ {
                    v46 := []int{780, 32768, (fi45 ^ -698)}
                    var v47 uint32 = v35
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                v48 := GenS1{F0: 0}
                _ = v44
                _ = v48
            }
            _ = v42
        }
    }
    v49 := GenS2{F0: -32768}
    fmt.Println(v35)
    _ = v35
    _ = v49
    return 786
}
func genFn1(p0 int16, p1 uint16) uint {
    var v50 uint64 = 42
    var v51 int = 1
    p1 = 852
    _ = p0
    _ = p1
    _ = v50
    _ = v51
    return 255
}

func main() {
    var v1 int = -337
    var v2 int32 = 1
    var v3 uint8 = 2
    var v4 string = (("\n\t" + "") + ("\n\t" + "a"))
    v5 := []string{v4, v4, ((v4 + v4) + v4)}
    switch v1 {
    case -96:
        if false {
            v52 := GenS5{F0: 4294967295, F1: v4, F2: 32767}
            v3 = v3
            _ = len(v5)
            _ = v52
        }
    }
    var v53 int8 = 25
    v54 := GenS6{F0: 216, F1: 65535, F2: ""}
    v55 := GenS6{F0: 2, F1: 1, F2: (("\n\t" + "\n\t") + "a")}
    switch ((("a" + "a") + (("a" + "") + "hello")) + ("hello" + ((v4 + v4) + "a"))) {
    case "hello":
        if (false && ((((float32(-2.25) * 1.5) + 0.1) * (((-2.25 + float32(float64(1e10))) - -2.25) + ((0.0 - float32(1.5)) + (-2.25 / -2.25)))) > (1e10 - 1e10))) {
            _ = len(v5)
            v56 := GenS7{F0: 65535, F1: false, F2: 4294967295}
            if (42 <= 18446744073709551615) {
                if ((!true) && false) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(v53, v1, v2)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                for fi57 := 0; fi57 < 17; fi57++ {
                    v58 := GenS8{F0: false, F1: 53}
                    fmt.Println(v4, v2, fi57)
                    v59 := GenS9{F0: true, F1: "", F2: 4294967295}
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                if (genFn0() >= -9223372036854775808) {
                    v60 := GenS10{F0: (v1 * (^(v1))), F1: 820}
                    _ = v60
                }
            }
            _ = v56
        }
    case "a":
        var v61 uint16 = 65535
        _ = v61
    default:
        for fi62 := 0; fi62 < 6; fi62++ {
            for fi63 := 0; fi63 < 1; fi63++ {
                fmt.Println(fi63, v4, v4)
                _ = fi63
            }
            _ = fi62
        }
    }
    if (v3 <= (v3 & (v3 >> ((genFn1(-32768, 1)) & 63)))) {
        switch "hello" {
        case "hello":
            if (true || (v4 == (("a" + v4) + v4))) {
                v3 -= 255
                v64 := GenS3{F0: v4, F1: (!(42 > 2147483648))}
                _ = v64
            }
        case "\n\t":
            if true {
                if true {
                    var v65 uint32 = 1
                    _ = v65
                }
                var v66 int16 = int16(v3)
                var v67 int8 = 1
                _ = v66
                _ = v67
            } else {
                fmt.Println(v2)
            }
        case "a":
            v68 := GenS7{F0: 0, F1: true, F2: 730}
            _ = v68
        }
        fmt.Println(v53, v3, v2)
    } else {
        for fi69 := 0; fi69 < 9; fi69++ {
            v53 += (v53 & v53)
            if (-32768 != 42) {
                var v70 string = v4
                fmt.Println(v53, v1)
                for fi71 := 0; fi71 < 6; fi71++ {
                    v1 = (fi69 * (fi71 ^ 128))
                    fmt.Println(v1)
                    _ = len(v5)
                    switch int(v2) {
                    case 309:
                        v3 = 128
                        var v72 bool = false
                        _ = v72
                    case 32768:
                        var v73 float32 = 3.14159
                        _ = v73
                    }
                    var v74 uint8 = v3
                    _ = fi71
                    _ = v74
                }
                _ = v70
            } else {
                if false {
                    v1 -= 4294967295
                    switch (fi69 / ((198) | 1)) {
                    case -9223372036854775808:
                        if len(v5) > 0 {
                            v5[0] = (v4 + v4)
                        }
                    case 1:
                        var v75 int16 = 42
                        _ = v75
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    var v76 int64 = 9223372036854775807
                    _ = v76
                }
            }
            _ = fi69
        }
        v77 := GenS10{F0: v1, F1: 4294967295}
        _ = v77
    }
    var v78 float32 = ((0.1 * 3.14159) * 0.0)
    var v79 uint = 508
    var v80 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v78
    _ = v79
    _ = v80
}
