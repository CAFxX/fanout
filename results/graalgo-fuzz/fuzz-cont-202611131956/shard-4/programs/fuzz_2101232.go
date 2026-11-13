package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint16 = 940
    var v2 string = ""
    var v3 uint8 = 128
    var v4 int32 = -812
    var v5 int32 = (v4 % ((284) | 1))
    if ((true && true) || false) {
        var v6 uint32 = 2147483648
        v2 = (v2 + v2)
        _ = v6
    } else {
        var v7 uint16 = 0
        for fi8 := 0; fi8 < 8; fi8++ {
            if (44 == (v3 << ((931) & 63))) {
                var v9 bool = false
                var v10 int = (fi8 & -742)
                var v11 uint8 = 98
                _ = v10
                _ = v11
                _ = v9
            }
            v1 = v1
            var v12 uint8 = v3
            _ = fi8
            _ = v12
        }
        _ = v7
    }
    v4 = (v5 | v4)
    var v13 string = v2
    var v14 float32 = 0.1
    var v15 int16 = -359
    var v16 uint8 = v3
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 float32 = v14
        _ = fi17
        _ = v18
    }
    switch 1 {
    case 0:
        switch -921 {
        case 0:
            for fi19 := 0; fi19 < 10; fi19++ {
                fmt.Println(v16)
                v20 := GenS1{F0: 519, F1: 42, F2: uint(v1)}
                for fi21 := 0; fi21 < 2; fi21++ {
                    if false {
                        switch v13 {
                        case "hello":
                            if (!(((((fi19 == (fi19 / (((fi21 - (fi19 >> ((42) & 63)))) | 1))) && true) || (((500 > 117) || (573 > fi21)) || true)) || false) || true)) {
                                switch (fi19 % ((int(v4)) | 1)) {
                                case -585:
                                    var v22 uint8 = (v3 | v16)
                                    _ = v22
                                case 127:
                                    v3 = (v16 + (v3 << ((1) & 63)))
                                case 128:
                                    v23 := []int{283, 4294967295}
                                    _ = v23
                                }
                                var v24 string = "a"
                                _ = v24
                            } else {
                                v3 = (v3 / ((225) | 1))
                                if false {
                                    v25 := GenS2{F0: -168}
                                    _ = v25
                                }
                                fmt.Println(v13)
                            }
                        case "a":
                            v5 += 152
                        case "\n\t":
                            switch (fi21 | fi19) {
                            case 883:
                                var v26 uint32 = 4294967295
                                _ = v26
                            case 9223372036854775807:
                                v27 := GenS3{F0: v4, F1: (^(fi19))}
                                var v28 bool = (false || false)
                                if false {
                                    var v29 int16 = v15
                                    if (true || v28) {
                                        var v30 uint8 = 2
                                        v15 = -191
                                        _ = v30
                                    } else {
                                        if false {
                                            var v31 int16 = (v15 << ((0) & 63))
                                            v2 = v2
                                            _ = v31
                                        } else {
                                            var v32 bool = ((v1 / (((^(v1))) | 1)) != 143)
                                            v33 := GenS3{F0: v4, F1: 127}
                                            _ = v32
                                            _ = v33
                                        }
                                        v34 := GenS4{F0: (127 > -6)}
                                        v35 := []int{fi21, fi19, fi19, 32768}
                                        _ = v34
                                        _ = v35
                                    }
                                    v36 := GenS4{F0: false}
                                    _ = v29
                                    _ = v36
                                } else {
                                    var v37 int16 = v15
                                    _ = v37
                                }
                                var v38 uint16 = 2
                                switch ("foo bar" + "\n\t") {
                                case "foo bar":
                                    v5 = v5
                                case "hello":
                                    v15 = -1
                                default:
                                    fmt.Println(fi19, v5)
                                }
                                v4 -= (v4 | v5)
                                v39 := []int{int(v15), -76, fi21, fi21}
                                _ = v27
                                _ = v28
                                _ = v38
                                _ = v39
                            case 680:
                                v40 := GenS5{F0: (v2 + v13), F1: -128}
                                _ = v40
                            }
                        default:
                            var v41 int16 = v15
                            _ = v41
                        }
                        var v42 uint64 = 706
                        _ = v42
                    } else {
                        var v43 uint = 255
                        _ = v43
                    }
                    if true {
                        switch v13 {
                        case "\n\t":
                            switch v2 {
                            case "a":
                                v2 = (("foo bar" + (((v13 + ((v13 + v2) + v13)) + "") + "hello")) + (v2 + v13))
                            case "":
                                v44 := GenS3{F0: v4, F1: fi21}
                                var v45 uint16 = (v1 + (v1 % (((v1 ^ v1)) | 1)))
                                _ = v44
                                _ = v45
                            case "foo bar":
                                v13 = "foo bar"
                            }
                        case "foo bar":
                            fmt.Println(v5)
                        case "a":
                            var v46 float64 = 0.1
                            _ = v46
                        }
                    } else {
                        v47 := []int{fi21, (fi19 * 127), (-(fi19))}
                        _ = v47
                    }
                    v48 := []int{fi19, (fi21 << ((43) & 63))}
                    _ = fi21
                    _ = v48
                }
                _ = fi19
                _ = v20
            }
        default:
            var v49 uint16 = 0
            _ = v49
        }
    case 352:
        switch "hello" {
        case "foo bar":
            for fi50 := 0; fi50 < 3; fi50++ {
                switch int(v4) {
                case 32767:
                    if (!(((v5 / (((^(v4))) | 1)) > (v5 | (v4 >> ((0) & 63)))) && (true && false))) {
                        if true {
                            var v51 uint64 = 18446744073709551615
                            _ = v51
                        } else {
                            v4 = v4
                            v52 := []int{32767, int(v1)}
                            v1 += (v1 >> ((65535) & 63))
                            _ = v52
                        }
                    } else {
                        v5 = 42
                        var v53 string = v2
                        _ = v53
                    }
                case 128:
                    var v54 float64 = 1e10
                    _ = v54
                case 4294967295:
                    if false {
                        var v55 uint = 42
                        var v56 string = v13
                        _ = v55
                        _ = v56
                    } else {
                        var v57 float64 = 1e10
                        switch (v13 + v2) {
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                var v58 float64 = v57
                                _ = v58
                            case "a":
                                for fi59 := 0; fi59 < 14; fi59++ {
                                    var v60 int = 0
                                    v61 := []int{}
                                    v62 := GenS1{F0: 42, F1: -1, F2: 679}
                                    _ = fi59
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                            }
                        case "a":
                            if (v57 != (v57 + v57)) {
                                v2 = "\n\t"
                            }
                        default:
                            v63 := GenS6{F0: fi50, F1: "foo bar"}
                            _ = v63
                        }
                        _ = v57
                    }
                default:
                    v64 := GenS7{F0: -118}
                    _ = v64
                }
                var v65 uint16 = v1
                _ = fi50
                _ = v65
            }
        case "\n\t":
            v66 := []int{-1}
            _ = v66
        case "hello":
            v15 = -32768
        }
    case 4294967295:
        for fi67 := 0; fi67 < 14; fi67++ {
            v15 += -891
            _ = fi67
        }
    }
    var v68 string = ("\n\t" + ("foo bar" + v13))
    v3 -= 128
    fmt.Println(v13, v15, v15)
    switch 2147483647 {
    case 127:
        for fi69 := 0; fi69 < 1; fi69++ {
            for fi70 := 0; fi70 < 18; fi70++ {
                v71 := []int{fi70, fi70, (fi70 | -497)}
                v16 = (^(v3))
                _ = fi70
                _ = v71
            }
            v72 := GenS5{F0: v68, F1: 100}
            _ = fi69
            _ = v72
        }
    default:
        fmt.Println(v3, v1, v1)
    }
    v73 := []int{-205, 4294967295, -1}
    fmt.Println(v1, v2)
    for fi74 := 0; fi74 < 15; fi74++ {
        v16 = v16
        _ = fi74
    }
    for fi75 := 0; fi75 < 4; fi75++ {
        var v76 int16 = (v15 << ((321) & 63))
        _ = fi75
        _ = v76
    }
    var v77 float32 = (v14 - v14)
    for fi78 := 0; fi78 < 18; fi78++ {
        v79 := GenS8{F0: v1, F1: -1}
        v2 = ""
        _ = fi78
        _ = v79
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v73:", len(v73), v73)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v73
    _ = v77
}
