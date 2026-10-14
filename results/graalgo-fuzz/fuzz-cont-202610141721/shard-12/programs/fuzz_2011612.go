package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 int32; F2 int16 }
type GenS7 struct { F0 int32; F1 uint }
type GenS8 struct { F0 uint64; F1 bool; F2 uint8 }
func genFn0(p0 uint16) int32 {
    p0 -= p0
    fmt.Println(p0)
    var v4 uint16 = 42
    var v5 float64 = ((1.5 / 1.0) * float64((1.5 - 3.14159)))
    switch 336 {
    case 8:
        for fi6 := 0; fi6 < 13; fi6++ {
            p0 += p0
            var v7 uint8 = 127
            var v8 int = fi6
            _ = fi6
            _ = v7
            _ = v8
        }
    default:
        v9 := GenS1{F0: 114}
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v5
    return 211
}
func genFn1(p0 uint16) bool {
    p0 -= p0
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(p0, fi10)
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS2{F0: 708}
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            if true {
                p0 -= (^(p0))
                switch fi13 {
                case 0:
                    fmt.Println(p0)
                case 318:
                    var v14 int32 = 1
                    _ = v14
                }
            }
            var v15 uint8 = 216
            _ = fi13
            _ = v15
        }
        _ = fi10
    }
    _ = p0
    return (-9223372036854775808 > 713)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 string = ((("foo bar" + "") + "a") + "hello")
    if true {
        if false {
            var v16 int = 444
            v3 = v3
            _ = v16
        } else {
            if (812 > 2) {
                var v17 float64 = float64(((((-2.25 - 1e10) + -2.25) * ((0.1 * 3.14159) - ((1.5 * float32(1.5)) / 1.5))) + (-2.25 + 0.0)))
                v18 := GenS3{F0: 793}
                v19 := GenS4{F0: 0}
                _ = v17
                _ = v18
                _ = v19
            }
            if (2 > 42) {
                switch 346 {
                case 128:
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 string = (v3 + v3)
                        for fi22 := 0; fi22 < 4; fi22++ {
                            v23 := GenS2{F0: 42}
                            _ = fi22
                            _ = v23
                        }
                        fmt.Println(fi20, v3)
                        _ = fi20
                        _ = v21
                    }
                }
                v24 := GenS5{F0: 9}
                _ = v24
            }
            v3 = v3
        }
        v25 := v2[572]
        var v26 int16 = 42
        _ = v25
        _ = v26
    } else {
        fmt.Println(v3)
        if true {
            for fi27 := 0; fi27 < 8; fi27++ {
                for fi28 := 0; fi28 < 16; fi28++ {
                    v3 = v3
                    fmt.Println(fi27, fi27)
                    var v29 int8 = 0
                    _ = fi28
                    _ = v29
                }
                var v30 float32 = 1e10
                v30 = v30
                _ = fi27
                _ = v30
            }
            var v31 int32 = -1
            _ = len(v2)
            _ = v31
        } else {
            for fi32 := 0; fi32 < 16; fi32++ {
                fmt.Println(fi32, fi32)
                _ = fi32
            }
        }
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        for fi34 := 0; fi34 < 18; fi34++ {
            if (((!true) && false) || false) {
                v35 := GenS2{F0: 42}
                if true {
                    var v36 int = 4294967295
                    _ = v36
                }
                _ = v35
            }
            fmt.Println(v3)
            _ = fi34
        }
        for fi37 := 0; fi37 < 16; fi37++ {
            switch ("hello" + "hello") {
            case "foo bar":
                v38 := GenS6{F0: 0, F1: 122, F2: 32767}
                _ = v38
            }
            var v39 bool = true
            _ = fi37
            _ = v39
        }
        _ = fi33
    }
    v3 = (v3 + v3)
    for fi40 := 0; fi40 < 13; fi40++ {
        switch (v3 + (v3 + "")) {
        case "foo bar":
            var v41 int32 = 42
            _ = v41
        case "a":
            if false {
                var v42 float64 = (0.0 / (((1e10 * float64(-2.25)) * 0.0) - 1.5))
                if true {
                    var v43 int = 4294967295
                    if (32767 == fi40) {
                        for fi44 := 0; fi44 < 3; fi44++ {
                            v3 = (v3 + v3)
                            fmt.Println(v43)
                            switch v43 {
                            case -1:
                                v42 = (v42 - ((0.1 - float64(1.5)) + (0.1 - v42)))
                            case 4294967295:
                                if true {
                                    v43 -= 32767
                                    switch ("hello" + ((v3 + "foo bar") + "hello")) {
                                    case "foo bar":
                                        v45 := []int{fi44}
                                        _ = v45
                                    case "hello":
                                        if (32767 == 255) {
                                            var v46 uint = 65535
                                            _ = v46
                                        }
                                    case "a":
                                        fmt.Println(v42)
                                    }
                                    var v47 uint32 = 65535
                                    _ = v47
                                }
                            }
                            _ = fi44
                        }
                        v42 -= float64(0.1)
                    }
                    for fi48 := 0; fi48 < 1; fi48++ {
                        _ = len(v1)
                        v49 := GenS4{F0: 255}
                        _ = fi48
                        _ = v49
                    }
                    _ = v43
                } else {
                    v50 := GenS1{F0: fi40}
                    if (429 != 913) {
                        for fi51 := 0; fi51 < 15; fi51++ {
                            v3 = v3
                            if false {
                                delete(v2, 1)
                                v3 = ((v3 + (v3 + ("" + ("\n\t" + ("\n\t" + v3))))) + v3)
                            } else {
                                var v52 uint = 918
                                v52 = (v52 / ((0) | 1))
                                _ = v52
                            }
                            switch fi40 {
                            case 1:
                                v1[v3] = 0
                            case -588:
                                v53 := GenS4{F0: 1}
                                _ = v53
                            default:
                                var v54 float32 = 1e10
                                _ = v54
                            }
                            _ = fi51
                        }
                    }
                    _ = v50
                }
                _ = v42
            }
        default:
            v55 := GenS7{F0: 2147483647, F1: 72}
            _ = v55
        }
        var v56 uint = 255
        _ = fi40
        _ = v56
    }
    var v57 uint64 = 18446744073709551615
    var v58 uint16 = 65535
    for fi59 := 0; fi59 < 7; fi59++ {
        fmt.Println(v57, v58)
        v58 -= (v58 % ((uint16(v57)) | 1))
        var v60 uint64 = v57
        _ = fi59
        _ = v60
    }
    v61 := GenS8{F0: 4294967295, F1: (42 <= 0), F2: 1}
    for fi62 := 0; fi62 < 5; fi62++ {
        v63 := []int{}
        _ = fi62
        _ = v63
    }
    if false {
        var v64 uint64 = v57
        fmt.Println(v58, v3, v64)
        _ = len(v2)
        _ = v64
    }
    fmt.Println(v57, v58)
    var v65 bool = false
    if v65 {
        if (-21 != -128) {
            if true {
                var v66 uint64 = (v57 / ((42) | 1))
                _ = v66
            } else {
                var v67 uint8 = 127
                v3 = v3
                if v65 {
                    var v68 int = 1
                    v68 -= (v68 | (-(v68)))
                    _ = v68
                }
                _ = v67
            }
            if (-9223372036854775808 > -26) {
                v69 := GenS7{F0: 1, F1: 354}
                if (0 < -32768) {
                    switch -1 {
                    case -982:
                        if (4294967295 >= 0) {
                            var v70 int16 = int16(v57)
                            if (100 < -5) {
                                v71 := GenS3{F0: 1}
                                _ = v71
                            }
                            var v72 string = (v3 + (v3 + "\n\t"))
                            _ = v70
                            _ = v72
                        } else {
                            v57 = (v57 / (((v57 % ((1) | 1))) | 1))
                            v3 = ("a" + "\n\t")
                        }
                    case 32768:
                        fmt.Println(v58, v65)
                    case -1:
                        v73 := GenS6{F0: v58, F1: 547, F2: 42}
                        _ = v73
                    default:
                        v3 = (v3 + "hello")
                    }
                    v57 = v57
                } else {
                    v74 := []int{}
                    _ = v74
                }
                for fi75 := 0; fi75 < 8; fi75++ {
                    var v76 uint = 2
                    switch 340 {
                    case 42:
                        v77 := v1["hello"]
                        _ = v77
                    default:
                        var v78 uint32 = 1
                        _ = v78
                    }
                    fmt.Println(v57, v3, v3)
                    _ = fi75
                    _ = v76
                }
                _ = v69
            }
            v57 += (v57 * 42)
        }
        switch 1 {
        case -243:
            var v79 uint16 = (^(v58))
            _ = v79
        default:
            var v80 int16 = 32767
            _ = v80
        }
    }
    var v81 bool = v65
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v65:", v65)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v57
    _ = v58
    _ = v61
    _ = v65
    _ = v81
}
