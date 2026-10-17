package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64 }
func genFn0(p0 uint32, p1 int8, p2 uint8) int16 {
    p1 = p1
    v4 := GenS1{F0: 0, F1: p0}
    if (!(32767 < 4294967295)) {
        v5 := GenS1{F0: 763, F1: p0}
        var v6 int32 = 188
        v6 = -773
        _ = v5
        _ = v6
    } else {
        var v7 bool = (1 <= (p1 >> ((135) & 63)))
        p0 = p0
        _ = v7
    }
    p1 += (p1 * (p1 & p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32767
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    v3 := make(map[uint32]string)
    v8 := GenS1{F0: 1, F1: 149}
    var v9 uint64 = 18446744073709551615
    fmt.Println(v9)
    switch 4294967295 {
    case 42:
        if (-1 < 42) {
            v10 := GenS2{F0: 359, F1: 65535}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 6; fi11++ {
                for fi12 := 0; fi12 < 1; fi12++ {
                    var v13 string = "foo bar"
                    _ = fi12
                    _ = v13
                }
                var v14 int8 = 1
                _ = fi11
                _ = v14
            }
            var v15 uint = 42
            var v16 int16 = 42
            _ = v15
            _ = v16
        }
    case -469:
        v9 += 0
    case 1:
        fmt.Println(v9)
    }
    fmt.Println(v9)
    for fi17 := 0; fi17 < 9; fi17++ {
        delete(v3, 1)
        v18 := GenS3{F0: fi17}
        _ = fi17
        _ = v18
    }
    v9 += v9
    switch 2147483647 {
    case 128:
        v9 -= (v9 - v9)
    case 968:
        v19 := GenS1{F0: 2, F1: 2147483648}
        _ = v19
    case -1:
        if (!((0.0 / -2.25) > (0.0 / 0.1))) {
            v9 -= v9
        }
    }
    v20 := GenS4{F0: -1, F1: 2147483647}
    for fi21 := 0; fi21 < 3; fi21++ {
        switch (-(fi21)) {
        case 9223372036854775807:
            v9 = (v9 | (v9 - v9))
        case 32768:
            if (-9223372036854775808 <= 9223372036854775807) {
                v22 := GenS5{F0: true, F1: 32767}
                var v23 float64 = 3.14159
                if true {
                    switch ("hello" + ("" + (("\n\t" + "\n\t") + ("a" + "foo bar")))) {
                    case "foo bar":
                        switch fi21 {
                        case 234:
                            _ = len(v3)
                        case -895:
                            var v24 bool = false
                            _ = v24
                        }
                    default:
                        for fi25 := 0; fi25 < 2; fi25++ {
                            fmt.Println(v23, fi25)
                            _ = fi25
                        }
                    }
                    var v26 uint64 = (v9 ^ 42)
                    v9 += 440
                    _ = v26
                } else {
                    var v27 bool = (176 > 138)
                    _ = v27
                }
                _ = v22
                _ = v23
            } else {
                var v28 uint16 = 255
                var v29 uint8 = uint8(v28)
                v30 := GenS6{F0: true, F1: (v28 | 733)}
                _ = v28
                _ = v29
                _ = v30
            }
        case 127:
            var v31 uint8 = 127
            _ = v31
        default:
            var v32 int = fi21
            _ = v32
        }
        var v33 int32 = -2147483648
        _ = fi21
        _ = v33
    }
    if (true && true) {
        v9 = (v9 - (v9 ^ (v9 - v9)))
        v34 := GenS4{F0: -893, F1: 2147483647}
        var v35 uint = 571
        _ = v34
        _ = v35
    }
    var v36 uint64 = v9
    var v37 float32 = (-2.25 - 1.5)
    if false {
        fmt.Println(v9)
        for fi38 := 0; fi38 < 8; fi38++ {
            var v39 uint16 = 336
            _ = fi38
            _ = v39
        }
    } else {
        var v40 uint8 = 0
        v41 := GenS7{F0: 1, F1: 0}
        if true {
            var v42 float32 = ((v37 * v37) * v37)
            var v43 uint16 = 255
            _ = v42
            _ = v43
        } else {
            fmt.Println(v9)
            if false {
                v44 := GenS8{F0: 2147483648}
                if true {
                    var v45 uint32 = 65535
                    v3[(v45 << ((4294967295) & 63))] = (("a" + ("hello" + ("a" + ""))) + "a")
                    _ = v45
                } else {
                    v46 := GenS1{F0: 0, F1: 4294967295}
                    if (!(false || (!true))) {
                        var v47 int = -555
                        v9 = (v36 / ((42) | 1))
                        _ = v47
                    } else {
                        switch ("a" + ("" + "\n\t")) {
                        case "\n\t":
                            v37 -= v37
                        case "hello":
                            for fi48 := 0; fi48 < 12; fi48++ {
                                if ((v40 - v40) <= 2) {
                                    switch "a" {
                                    case "hello":
                                        if true {
                                            v37 = (v37 / v37)
                                            if false {
                                                fmt.Println(v36, v40)
                                                _ = len(v2)
                                                var v49 float32 = v37
                                                _ = v49
                                            }
                                            v36 += (v36 % ((42) | 1))
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                var v50 bool = (!(!true))
                                                _ = v50
                                            case "":
                                                var v51 uint32 = 375
                                                _ = v51
                                            default:
                                                var v52 uint = 18446744073709551615
                                                _ = v52
                                            }
                                        }
                                    case "a":
                                        switch (fi48 * 42) {
                                        case 127:
                                            fmt.Println(v40)
                                        case 4294967295:
                                            v53 := GenS1{F0: 255, F1: 65535}
                                            _ = v53
                                        default:
                                            v3[2] = "a"
                                        }
                                    default:
                                        var v54 uint32 = 2147483647
                                        _ = v54
                                    }
                                    v40 = 127
                                } else {
                                    _ = len(v1)
                                }
                                _ = fi48
                            }
                        default:
                            if (!true) {
                                for fi55 := 0; fi55 < 10; fi55++ {
                                    for fi56 := 0; fi56 < 8; fi56++ {
                                        switch "foo bar" {
                                        case "hello":
                                            var v57 bool = false
                                            _ = v57
                                        default:
                                            v58 := GenS4{F0: -9223372036854775808, F1: fi55}
                                            _ = v58
                                        }
                                        _ = fi56
                                    }
                                    v37 -= v37
                                    _ = fi55
                                }
                                v59 := GenS2{F0: 0, F1: 65535}
                                _ = v59
                            }
                        }
                        if (1 != -653) {
                            for fi60 := 0; fi60 < 9; fi60++ {
                                var v61 float32 = 1.5
                                for fi62 := 0; fi62 < 14; fi62++ {
                                    if (v36 != v9) {
                                        var v63 int8 = 106
                                        if ("" > ("\n\t" + "")) {
                                            var v64 uint16 = 914
                                            switch ("a" + ("foo bar" + "a")) {
                                            case "a":
                                                var v65 uint32 = 1
                                                _ = v65
                                            case "hello":
                                                if len(v2) > 0 {
                                                    v2[0] = "foo bar"
                                                }
                                            default:
                                                v66 := GenS5{F0: true, F1: (v64 - v64)}
                                                _ = v66
                                            }
                                            var v67 uint8 = uint8(v64)
                                            _ = v64
                                            _ = v67
                                        } else {
                                            var v68 uint16 = 32767
                                            if true {
                                                var v69 float64 = 3.14159
                                                fmt.Println(v68, v63)
                                                _ = v69
                                            } else {
                                                var v70 string = "foo bar"
                                                v9 -= v36
                                                _ = v70
                                            }
                                            var v71 bool = (true && true)
                                            var v72 bool = ((v61 <= float32(1e10)) && v71)
                                            _ = len(v2)
                                            _ = v68
                                            _ = v71
                                            _ = v72
                                        }
                                        var v73 uint8 = (^(v40))
                                        _ = v63
                                        _ = v73
                                    }
                                    _ = fi62
                                }
                                _ = fi60
                                _ = v61
                            }
                            var v74 uint16 = 1
                            _ = v74
                        } else {
                            delete(v1, 32767)
                            if true {
                                v40 = 127
                                v37 = v37
                            } else {
                                if (true || false) {
                                    if (1 == -1) {
                                        var v75 bool = (9223372036854775807 >= -9223372036854775808)
                                        _ = v75
                                    } else {
                                        v76 := GenS9{F0: 42}
                                        _ = v76
                                    }
                                }
                            }
                        }
                    }
                    _ = v46
                }
                fmt.Println(v9)
                _ = v44
            }
        }
        _ = v40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v37
    _ = v8
    _ = v9
}
