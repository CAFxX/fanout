package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 string; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int64; F2 int16 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 int64; F1 int32; F2 int64 }
type GenS10 struct { F0 int; F1 bool; F2 uint }
type GenS11 struct { F0 string; F1 bool }
type GenS12 struct { F0 int32 }
func genFn0(p0 uint16, p1 uint) int64 {
    v4 := GenS1{F0: 2, F1: 42}
    fmt.Println(p1, p0)
    p1 = 42
    switch (("hello" + "hello") + "\n\t") {
    case "a":
        var v5 int = 42
        _ = v5
    }
    if (63 == 0) {
        fmt.Println(p0, p0)
        v6 := GenS2{F0: 42}
        _ = v6
    } else {
        v7 := GenS1{F0: 127, F1: -67}
        p1 = p1
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return -484
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 0
    var v3 uint16 = 65535
    v8 := GenS3{F0: 0}
    var v9 bool = ((v2 << ((18446744073709551615) & 63)) <= v2)
    v1["a"] = 4294967295
    v2 = (v2 + (v2 << ((234) & 63)))
    for fi10 := 0; fi10 < 11; fi10++ {
        v9 = (!(!(!(!((fi10 % (((fi10 * v2)) | 1)) <= fi10)))))
        v11 := GenS4{F0: (fi10 ^ v2), F1: "a", F2: (v2 % ((1) | 1))}
        if v9 {
            v12 := GenS1{F0: 2, F1: -538}
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    if (!v9) {
        fmt.Println(v3)
        if v9 {
            v3 += v3
            if v9 {
                for fi13 := 0; fi13 < 15; fi13++ {
                    if (true || v9) {
                        var v14 float32 = 0.0
                        if v9 {
                            if v9 {
                                var v15 bool = v9
                                if false {
                                    fmt.Println(v3)
                                    var v16 string = ("\n\t" + "\n\t")
                                    _ = v16
                                } else {
                                    var v17 uint64 = 18446744073709551615
                                    _ = v17
                                }
                                _ = v15
                            }
                            v2 = v2
                        } else {
                            var v18 uint8 = 1
                            fmt.Println(fi13, v14, v2)
                            _ = v18
                        }
                        v19 := GenS3{F0: 449}
                        _ = v14
                        _ = v19
                    }
                    var v20 bool = false
                    _ = fi13
                    _ = v20
                }
            }
            v1[("\n\t" + ("a" + ((("foo bar" + ("" + ("\n\t" + "hello"))) + ("a" + "a")) + (("foo bar" + "a") + (("a" + "\n\t") + "a")))))] = 121
        }
    }
    if true {
        var v21 float64 = ((0.1 * 1.5) / 1.0)
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS3{F0: 1}
            switch ("" + "foo bar") {
            case "a":
                v2 = (v2 ^ (fi22 + (fi22 ^ (v2 * fi22))))
            case "hello":
                v24 := GenS5{F0: 214}
                _ = v24
            case "":
                var v25 int8 = 0
                _ = v25
            }
            if (42 > int16(fi22)) {
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := v1["foo bar"]
                    _ = fi26
                    _ = v27
                }
            } else {
                v28 := GenS4{F0: fi22, F1: "\n\t", F2: -9223372036854775808}
                if (42 < 0) {
                    switch "foo bar" {
                    case "foo bar":
                        switch ("a" + "hello") {
                        case "":
                            v3 -= 1
                        default:
                            if (200 == 0) {
                                switch 40 {
                                case 32767:
                                    v29 := GenS6{F0: 4294967295, F1: -1, F2: 1}
                                    _ = v29
                                case 42:
                                    var v30 int16 = -32768
                                    _ = v30
                                default:
                                    var v31 uint = 4294967295
                                    _ = v31
                                }
                                if true {
                                    v32 := GenS7{F0: 0, F1: (42 >= 48)}
                                    v33 := GenS3{F0: 255}
                                    v3 = (v3 & v3)
                                    _ = v32
                                    _ = v33
                                }
                            } else {
                                switch "a" {
                                case "":
                                    var v34 float64 = v21
                                    _ = v34
                                case "hello":
                                    _ = len(v1)
                                }
                                fmt.Println(fi22, v3, v21)
                                var v35 string = ("foo bar" + "\n\t")
                                _ = v35
                            }
                        }
                    case "a":
                        if false {
                            switch (fi22 / (((-(v2))) | 1)) {
                            case 32768:
                                var v36 uint16 = v3
                                _ = v36
                            case 32767:
                                for fi37 := 0; fi37 < 17; fi37++ {
                                    fmt.Println(fi22)
                                    _ = fi37
                                }
                            case -2147483648:
                                for fi38 := 0; fi38 < 2; fi38++ {
                                    _ = len(v1)
                                    v39 := GenS6{F0: 2, F1: 1, F2: -498}
                                    _ = fi38
                                    _ = v39
                                }
                            default:
                                v21 = float64(0.1)
                            }
                            v40 := GenS8{F0: v9, F1: 65535}
                            if (!(v9 && (1 != 57))) {
                                var v41 int = -1
                                v21 += 3.14159
                                _ = v41
                            }
                            _ = v40
                        } else {
                            v21 += float64(1e10)
                        }
                    case "":
                        v42 := GenS4{F0: (v2 % ((v2) | 1)), F1: "", F2: 128}
                        _ = v42
                    default:
                        fmt.Println(fi22, fi22)
                    }
                } else {
                    v43 := GenS1{F0: 42, F1: -624}
                    for fi44 := 0; fi44 < 8; fi44++ {
                        fmt.Println(v2, fi22, v3)
                        _ = fi44
                    }
                    _ = v43
                }
                _ = v28
            }
            _ = fi22
            _ = v23
        }
        for fi45 := 0; fi45 < 17; fi45++ {
            v9 = v9
            _ = fi45
        }
        _ = v21
    }
    v46 := GenS6{F0: 18446744073709551615, F1: 42, F2: -1}
    if false {
        var v47 uint8 = 255
        var v48 uint8 = v47
        if true {
            v49 := []int{(-(v2))}
            for fi50 := 0; fi50 < 9; fi50++ {
                if v9 {
                    if v9 {
                        for fi51 := 0; fi51 < 13; fi51++ {
                            v52 := GenS6{F0: 65535, F1: 42, F2: -1}
                            _ = fi51
                            _ = v52
                        }
                        v3 += (v3 * uint16(v47))
                        for fi53 := 0; fi53 < 3; fi53++ {
                            _ = len(v49)
                            v48 -= (v48 - (^(v48)))
                            _ = fi53
                        }
                    }
                    var v54 uint32 = 191
                    _ = v54
                }
                if (!v9) {
                    v2 = v2
                } else {
                    v2 = (v2 | -9223372036854775808)
                    var v55 float32 = 0.0
                    switch v2 {
                    case 198:
                        v47 -= (v48 & 255)
                    case -896:
                        var v56 uint8 = (v47 | v48)
                        _ = v56
                    }
                    _ = v55
                }
                if true {
                    _ = len(v1)
                    v3 = (v3 ^ (v3 | (v3 * v3)))
                    switch (v2 << ((18446744073709551615) & 63)) {
                    case 32767:
                        var v57 int16 = -32768
                        _ = v57
                    case 42:
                        switch v2 {
                        case 9223372036854775807:
                            v48 = 127
                        case 965:
                            if len(v49) >= 2 {
                                v49 = v49[1:len(v49)]
                            }
                        case 128:
                            switch "a" {
                            case "":
                                fmt.Println(v2, v48, fi50)
                            case "foo bar":
                                var v58 uint8 = 0
                                _ = v58
                            case "\n\t":
                                for fi59 := 0; fi59 < 15; fi59++ {
                                    v9 = (4294967295 >= 95)
                                    _ = fi59
                                }
                            }
                        default:
                            if v9 {
                                var v60 uint = 4294967295
                                var v61 int = fi50
                                _ = v60
                                _ = v61
                            }
                        }
                    case 767:
                        v62 := GenS3{F0: uint(v48)}
                        _ = v62
                    }
                } else {
                    if true {
                        v49 = append(v49, v2)
                        fmt.Println(v2, v47)
                    }
                    switch fi50 {
                    case 42:
                        for fi63 := 0; fi63 < 20; fi63++ {
                            if v9 {
                                if (v9 && (1 <= 1)) {
                                    v47 += v48
                                    _ = len(v1)
                                    v64 := GenS9{F0: -1, F1: 0, F2: -1}
                                    _ = v64
                                } else {
                                    switch ("\n\t" + "") {
                                    case "\n\t":
                                        fmt.Println(fi63, v2, v47)
                                    case "a":
                                        var v65 string = "hello"
                                        if false {
                                            var v66 uint8 = v48
                                            if len(v49) >= 2 {
                                                v49 = v49[1:len(v49)]
                                            }
                                            switch v2 {
                                            case 281:
                                                v66 = (v47 << ((uint(v48)) & 63))
                                            }
                                            _ = v66
                                        }
                                        var v67 uint8 = (v48 & v47)
                                        _ = v65
                                        _ = v67
                                    default:
                                        var v68 string = "a"
                                        _ = v68
                                    }
                                    v69 := GenS6{F0: 151, F1: 1, F2: 813}
                                    _ = v69
                                }
                            }
                            _ = len(v49)
                            fmt.Println(v48)
                            _ = fi63
                        }
                    case 128:
                        delete(v1, "a")
                    case 174:
                        fmt.Println(v48, v2, v48)
                    default:
                        var v70 uint32 = 2147483648
                        _ = v70
                    }
                    if (!v9) {
                        if true {
                            v71 := GenS10{F0: (fi50 >> ((2) & 63)), F1: v9, F2: 0}
                            v47 += v47
                            _ = v71
                        }
                        v48 += 34
                        if v9 {
                            var v72 int8 = 42
                            v47 = (^(v47))
                            _ = len(v1)
                            _ = v72
                        } else {
                            if len(v49) >= 2 {
                                v49 = v49[1:len(v49)]
                            }
                            fmt.Println(v47, v3, v47)
                            if v9 {
                                _ = len(v1)
                            }
                        }
                    } else {
                        v73 := GenS11{F0: (("" + "foo bar") + ("\n\t" + "a")), F1: v9}
                        v74 := GenS12{F0: 0}
                        for fi75 := 0; fi75 < 15; fi75++ {
                            if true {
                                switch "" {
                                case "hello":
                                    v3 -= (v3 << ((1) & 63))
                                case "":
                                    var v76 int = (v2 / (((v2 / (((v2 - v2)) | 1))) | 1))
                                    _ = v76
                                default:
                                    var v77 bool = v9
                                    _ = v77
                                }
                                fmt.Println(v47)
                            } else {
                                var v78 uint8 = (^(v48))
                                _ = v78
                            }
                            v3 = v3
                            if len(v49) > 0 {
                                v49[0] = (fi50 | (fi75 | fi50))
                            }
                            _ = fi75
                        }
                        _ = v73
                        _ = v74
                    }
                }
                _ = fi50
            }
            fmt.Println(v48, v2)
            _ = v49
        }
        _ = v47
        _ = v48
    }
    for fi79 := 0; fi79 < 20; fi79++ {
        v2 += (fi79 | (v2 & (-(fi79))))
        v2 = (v2 >> ((595) & 63))
        _ = fi79
    }
    if (!((-1 > -11) || (-1 <= -128))) {
        switch 4294967295 {
        case 32767:
            v3 -= 582
        case 0:
            for fi80 := 0; fi80 < 4; fi80++ {
                v81 := []int{(^(fi80)), 128, (v2 * 0), 286}
                v82 := GenS7{F0: 42, F1: true}
                switch fi80 {
                case -1:
                    var v83 uint32 = 1
                    _ = v83
                case 1:
                    var v84 int8 = -40
                    _ = v84
                default:
                    v3 -= (v3 + v3)
                }
                _ = fi80
                _ = v81
                _ = v82
            }
        case 2147483647:
            v85 := []int{-816, v2, v2}
            _ = v85
        }
    } else {
        var v86 uint64 = 978
        _ = v86
    }
    var v87 uint64 = 18446744073709551615
    for fi88 := 0; fi88 < 10; fi88++ {
        var v89 uint32 = 2
        delete(v1, "\n\t")
        _ = fi88
        _ = v89
    }
    if v9 {
        v90 := GenS9{F0: -373, F1: -535, F2: -1}
        _ = v90
    }
    v91 := GenS5{F0: 2}
    v9 = (!("foo bar" == ""))
    v2 -= (v2 * 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v87:", v87)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v8
    _ = v87
    _ = v9
    _ = v91
}
