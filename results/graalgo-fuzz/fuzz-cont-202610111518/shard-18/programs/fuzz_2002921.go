package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint32; F1 int64; F2 string }
type GenS5 struct { F0 int64; F1 int32; F2 uint64 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS7 struct { F0 string; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint; F1 int; F2 int8 }
type GenS9 struct { F0 int16; F1 int8; F2 uint }
type GenS10 struct { F0 uint16; F1 int8 }
type GenS11 struct { F0 bool }
func genFn0(p0 uint64, p1 bool) uint32 {
    v6 := []int{127, -1, -47}
    v7 := GenS1{F0: -91, F1: -142, F2: 175}
    switch "hello" {
    case "a":
        if ((745 < 971) && true) {
            var v8 uint = 0
            switch "foo bar" {
            case "foo bar":
                p0 = (^(p0))
            case "a":
                var v9 uint = (v8 >> ((v8) & 63))
                _ = v9
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v8
        } else {
            fmt.Println(p1)
        }
    case "foo bar":
        p1 = p1
    default:
        if len(v6) > 0 {
            v6[0] = 127
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return uint32(p0)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = -2.25
    var v3 uint = 0
    v4 := []string{}
    v5 := []int{}
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int32 = 883
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        if (0.0 < 0.0) {
            fmt.Println(v3)
            switch 127 {
            case 42:
                var v13 uint = (v3 + v3)
                _ = v13
            }
            v14 := GenS1{F0: 42, F1: -2147483648, F2: 1}
            _ = v14
        } else {
            if ((2147483648 > 694) || true) {
                v15 := GenS2{F0: -1}
                _ = v15
            }
            var v16 int64 = 447
            _ = v16
        }
        v2 = v2
        _ = fi12
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v18 := GenS1{F0: 100, F1: -471, F2: 1}
        fmt.Println(fi17, v2)
        _ = fi17
        _ = v18
    }
    if (877 < 32768) {
        v3 = v3
        for fi19 := 0; fi19 < 18; fi19++ {
            v3 = 255
            _ = fi19
        }
    } else {
        for fi20 := 0; fi20 < 16; fi20++ {
            v2 += v2
            v2 = v2
            switch ("hello" + "\n\t") {
            case "foo bar":
                var v21 int32 = -2147483648
                _ = v21
            case "":
                fmt.Println(v3, v2, v3)
            case "\n\t":
                var v22 uint8 = 42
                _ = v22
            }
            _ = fi20
        }
        for fi23 := 0; fi23 < 16; fi23++ {
            v4 = append(v4, "\n\t")
            v3 -= 1
            _ = fi23
        }
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        if (false || false) {
            _ = len(v5)
            if (!true) {
                for fi25 := 0; fi25 < 14; fi25++ {
                    var v26 float64 = v2
                    delete(v1, (^(fi25)))
                    _ = fi25
                    _ = v26
                }
                var v27 uint8 = 159
                _ = v27
            } else {
                v28 := GenS2{F0: 0}
                if (!true) {
                    v29 := GenS3{F0: 453, F1: ("" + ("a" + "hello"))}
                    for fi30 := 0; fi30 < 17; fi30++ {
                        fmt.Println(fi24)
                        _ = fi30
                    }
                    for fi31 := 0; fi31 < 19; fi31++ {
                        if (!(!true)) {
                            var v32 uint64 = 4294967295
                            _ = v32
                        } else {
                            if true {
                                if (41 != 2) {
                                    var v33 string = ("foo bar" + "foo bar")
                                    var v34 int32 = -2147483648
                                    v1[127] = 4294967295
                                    _ = v33
                                    _ = v34
                                }
                                var v35 uint16 = 0
                                _ = v35
                            } else {
                                v36 := GenS1{F0: 100, F1: -1, F2: 111}
                                v2 = (0.0 * -2.25)
                                _ = v36
                            }
                        }
                        var v37 int64 = -9223372036854775808
                        v38 := GenS2{F0: 0}
                        _ = fi31
                        _ = v37
                        _ = v38
                    }
                    _ = v29
                }
                _ = v28
            }
        }
        for fi39 := 0; fi39 < 2; fi39++ {
            if false {
                switch "foo bar" {
                case "":
                    v3 += v3
                case "a":
                    v3 = (^(v3))
                }
                var v40 int8 = 42
                _ = v40
            }
            _ = fi39
        }
        if true {
            switch fi24 {
            case 32768:
                v2 -= v2
            default:
                v1[fi24] = genFn0(0, true)
            }
            for fi41 := 0; fi41 < 15; fi41++ {
                if len(v4) > 0 {
                    v4[0] = "a"
                }
                v42 := GenS2{F0: 819}
                v43 := v1[fi24]
                _ = fi41
                _ = v42
                _ = v43
            }
        }
        _ = fi24
    }
    v2 = v2
    fmt.Println(v3, v2)
    v44 := GenS4{F0: 2, F1: 42, F2: "a"}
    if false {
        fmt.Println(v2)
        if true {
            if false {
                v45 := GenS3{F0: 1, F1: ""}
                _ = v45
            } else {
                if (-122 > -585) {
                    v46 := GenS2{F0: 1000}
                    if true {
                        v2 += v2
                    } else {
                        var v47 int = 11
                        v48 := GenS5{F0: -1, F1: 1, F2: 4294967295}
                        if (!true) {
                            fmt.Println(v2, v3)
                            v3 += (v3 % (((v3 << ((v3) & 63))) | 1))
                        }
                        _ = v47
                        _ = v48
                    }
                    switch -2147483648 {
                    case 468:
                        v2 -= 0.0
                    case 227:
                        switch (((("\n\t" + (((("a" + "foo bar") + "foo bar") + "") + "")) + "") + "hello") + "hello") {
                        case "hello":
                            if false {
                                _ = len(v4)
                            }
                        case "foo bar":
                            v2 += (3.14159 / 0.1)
                        default:
                            var v49 int32 = 0
                            _ = v49
                        }
                    default:
                        fmt.Println(v2, v3)
                    }
                    _ = v46
                } else {
                    var v50 int32 = 0
                    _ = v50
                }
                fmt.Println(v3, v3)
            }
            v5 = append(v5, 32767)
            if (-2147483648 > 42) {
                v51 := GenS6{F0: 947, F1: 65535, F2: 2}
                _ = v51
            } else {
                switch 9223372036854775807 {
                case -1:
                    var v52 uint8 = 128
                    _ = v52
                case -2147483648:
                    v3 = (v3 >> ((42) & 63))
                case 0:
                    v53 := GenS7{F0: "\n\t", F1: -1, F2: 669}
                    _ = v53
                default:
                    for fi54 := 0; fi54 < 7; fi54++ {
                        switch ("foo bar" + "hello") {
                        case "hello":
                            for fi55 := 0; fi55 < 20; fi55++ {
                                var v56 uint16 = 2
                                _ = fi55
                                _ = v56
                            }
                        case "a":
                            v57 := GenS2{F0: 23}
                            _ = v57
                        }
                        var v58 uint16 = 2
                        if (false && (-32768 >= 1000)) {
                            switch 9223372036854775807 {
                            case -741:
                                if (true && (true && true)) {
                                    v59 := GenS4{F0: 2, F1: -1, F2: "hello"}
                                    switch "a" {
                                    case "":
                                        fmt.Println(fi54)
                                    default:
                                        v2 += v2
                                    }
                                    _ = v59
                                }
                            }
                            fmt.Println(v2)
                            for fi60 := 0; fi60 < 2; fi60++ {
                                v2 += (float64(0.1) + v2)
                                v2 += v2
                                _ = fi60
                            }
                        } else {
                            v61 := GenS8{F0: (v3 & v3), F1: (fi54 & fi54), F2: 127}
                            _ = v61
                        }
                        _ = fi54
                        _ = v58
                    }
                }
            }
        }
    } else {
        var v62 string = "a"
        var v63 float32 = -2.25
        _ = v62
        _ = v63
    }
    if true {
        var v64 float32 = (-2.25 + 0.1)
        var v65 int16 = -541
        _ = v64
        _ = v65
    } else {
        v4 = append(v4, "")
    }
    v66 := GenS9{F0: -876, F1: 100, F2: 228}
    var v67 int16 = 1000
    switch "foo bar" {
    case "\n\t":
        v68 := GenS10{F0: 32767, F1: -65}
        _ = v68
    case "foo bar":
        if true {
            v69 := GenS11{F0: ((42 < 1) || (!(!((-(v67)) <= (v67 * -32768)))))}
            v67 = (v67 + (^(v67)))
            _ = v69
        }
    default:
        for fi70 := 0; fi70 < 7; fi70++ {
            var v71 int16 = 416
            var v72 int = fi70
            _ = fi70
            _ = v71
            _ = v72
        }
    }
    v67 = v67
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v66
    _ = v67
}
