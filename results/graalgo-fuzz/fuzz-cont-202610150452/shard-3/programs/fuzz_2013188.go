package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint8 }
func genFn0(p0 int8) uint8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: (fi8 + 2147483647)}
            fmt.Println(p0)
            fmt.Println(fi7)
            _ = fi8
            _ = v9
        }
        v10 := []int{fi7, fi7, 1, fi7}
        _ = fi7
        _ = v10
    }
    var v11 uint32 = 653
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := []int{(fi12 % ((fi12) | 1))}
        _ = fi12
        _ = v13
    }
    if false {
        v14 := []int{32768, 2147483647}
        _ = v14
    } else {
        var v15 int8 = p0
        if true {
            var v16 uint8 = 127
            _ = v16
        } else {
            v17 := []int{-74, 2147483647, 42}
            _ = v17
        }
        v18 := []int{968, -348}
        _ = v15
        _ = v18
    }
    _ = p0
    _ = v11
    return 128
}
func genFn1() int {
    if true {
        var v19 float32 = 1.5
        var v20 uint64 = 4294967295
        v21 := GenS1{F0: 239}
        _ = v19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := []int{fi22, fi22, (fi22 >> ((255) & 63))}
        if (-1 <= -1) {
            var v24 int32 = -1
            switch "foo bar" {
            case "foo bar":
                v24 += (v24 << ((42) & 63))
            case "a":
                var v25 uint16 = 65535
                _ = v25
            default:
                v24 -= v24
            }
            switch (fi22 & fi22) {
            case -1:
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
            default:
                if (32767 < 470) {
                    v24 = 2147483647
                    var v26 bool = false
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        fmt.Println(v24, v26)
                    case "a":
                        v24 = v24
                    case "foo bar":
                        var v27 float32 = 0.0
                        _ = v27
                    default:
                        v26 = (173 >= 128)
                    }
                    _ = v26
                }
            }
            _ = v24
        } else {
            v28 := GenS1{F0: fi22}
            if len(v23) > 0 {
                v23[0] = -985
            }
            var v29 int16 = 0
            _ = v28
            _ = v29
        }
        _ = fi22
        _ = v23
    }
    if (false && (!true)) {
        for fi30 := 0; fi30 < 19; fi30++ {
            v31 := []int{fi30}
            _ = fi30
            _ = v31
        }
        switch ("foo bar" + "") {
        case "\n\t":
            switch "" {
            case "":
                var v32 float64 = 0.0
                _ = v32
            }
        case "a":
            switch 42 {
            case -318:
                var v33 uint32 = 2147483647
                _ = v33
            }
        case "hello":
            var v34 float32 = (1e10 * float32(1.5))
            _ = v34
        default:
            if true {
                var v35 int = 59
                _ = v35
            }
        }
    } else {
        v36 := []int{32768}
        _ = v36
    }
    return -1
}

func main() {
    var v1 int64 = 74
    var v2 float64 = 1e10
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    var v6 int64 = (v1 / ((-1) | 1))
    var v37 uint8 = 180
    v38 := GenS1{F0: -9223372036854775808}
    var v39 int64 = (^(v1))
    if (127 >= -128) {
        if (!true) {
            if false {
                switch -280 {
                case 9223372036854775807:
                    var v40 string = "hello"
                    _ = v40
                }
                v41 := []int{2147483647, 4294967295}
                _ = v41
            } else {
                if (genFn0(-128) > (v37 / ((v37) | 1))) {
                    if false {
                        delete(v4, "a")
                        var v42 int64 = v6
                        _ = v42
                    }
                }
                for fi43 := 0; fi43 < 4; fi43++ {
                    v2 = float64((0.1 + 0.0))
                    var v44 uint64 = 4294967295
                    _ = fi43
                    _ = v44
                }
            }
        }
        if (127 >= -837) {
            var v45 int16 = 32767
            if true {
                for fi46 := 0; fi46 < 12; fi46++ {
                    if false {
                        var v47 int16 = 1
                        var v48 int8 = 42
                        _ = v47
                        _ = v48
                    } else {
                        v49 := GenS1{F0: genFn1()}
                        if false {
                            v50 := GenS1{F0: 891}
                            _ = v50
                        }
                        _ = v49
                    }
                    _ = fi46
                }
                var v51 int = 32768
                for fi52 := 0; fi52 < 19; fi52++ {
                    var v53 int32 = 0
                    _ = fi52
                    _ = v53
                }
                _ = v51
            } else {
                for fi54 := 0; fi54 < 2; fi54++ {
                    if (true || false) {
                        v55 := GenS2{F0: 1, F1: (fi54 + (^(fi54)))}
                        v56 := []int{(fi54 | (fi54 ^ fi54)), fi54, fi54, (fi54 * -9223372036854775808)}
                        _ = v55
                        _ = v56
                    } else {
                        var v57 float32 = 3.14159
                        _ = v57
                    }
                    _ = fi54
                }
                v2 += 1e10
                switch 4294967295 {
                case 1:
                    switch -1 {
                    case 32767:
                        v2 -= 0.0
                    case -2147483648:
                        v2 = (v2 - v2)
                    case 127:
                        for fi58 := 0; fi58 < 10; fi58++ {
                            if true {
                                for fi59 := 0; fi59 < 19; fi59++ {
                                    if true {
                                        switch -9223372036854775808 {
                                        case -886:
                                            v39 -= v6
                                        default:
                                            var v60 string = "a"
                                            v60 = ""
                                            _ = v60
                                        }
                                        fmt.Println(v39, v39, fi59)
                                        v61 := GenS2{F0: 2, F1: fi58}
                                        var v62 int64 = (v1 - (v39 | -93))
                                        _ = v61
                                        _ = v62
                                    } else {
                                        v63 := GenS2{F0: 32767, F1: (fi58 << ((42) & 63))}
                                        delete(v5, "")
                                        if false {
                                            v64 := GenS1{F0: fi58}
                                            if (!false) {
                                                fmt.Println(v2, v37, v39)
                                            } else {
                                                v65 := GenS3{F0: 42}
                                                var v66 uint = 18446744073709551615
                                                fmt.Println(v39, v6)
                                                _ = v65
                                                _ = v66
                                            }
                                            _ = v64
                                        } else {
                                            v67 := GenS3{F0: 32767}
                                            v1 = -1
                                            var v68 float64 = v2
                                            _ = v67
                                            _ = v68
                                        }
                                        _ = v63
                                    }
                                    var v69 float32 = 0.1
                                    _ = fi59
                                    _ = v69
                                }
                                v39 -= 42
                                v70 := []int{-743, fi58, fi58, (fi58 % (((^(fi58))) | 1))}
                                _ = v70
                            } else {
                                v71 := GenS2{F0: 899, F1: fi58}
                                v6 = (v6 * int64(v37))
                                _ = v71
                            }
                            var v72 int32 = -2147483648
                            _ = fi58
                            _ = v72
                        }
                    default:
                        if (-1 == (v6 / ((v6) | 1))) {
                            var v73 uint64 = 774
                            _ = v73
                        } else {
                            v1 -= v6
                            for fi74 := 0; fi74 < 1; fi74++ {
                                v37 += (^(v37))
                                if true {
                                    fmt.Println(v39)
                                    fmt.Println(v6, v6, v2)
                                    v75 := v5[""]
                                    _ = v75
                                } else {
                                    if false {
                                        fmt.Println(v39, v1)
                                    }
                                    if false {
                                        v37 = (^(v37))
                                        fmt.Println(v1, v39, v1)
                                        v76 := GenS2{F0: 0, F1: fi74}
                                        _ = v76
                                    } else {
                                        switch ("a" + "hello") {
                                        case "\n\t":
                                            if true {
                                                var v77 uint64 = 18446744073709551615
                                                if (false && (1 > 1)) {
                                                    var v78 string = "\n\t"
                                                    _ = v78
                                                }
                                                v2 -= (0.1 + v2)
                                                _ = v77
                                            }
                                        case "a":
                                            switch "hello" {
                                            case "hello":
                                                v37 += genFn0(-4)
                                            case "\n\t":
                                                var v79 int16 = v45
                                                _ = v79
                                            case "a":
                                                v80 := []int{fi74}
                                                _ = v80
                                            default:
                                                for fi81 := 0; fi81 < 2; fi81++ {
                                                    var v82 bool = (!false)
                                                    if (v82 || (1.5 == 0.1)) {
                                                        var v83 uint32 = 240
                                                        var v84 uint16 = 255
                                                        _ = v83
                                                        _ = v84
                                                    } else {
                                                        fmt.Println(fi81)
                                                        var v85 int8 = 1
                                                        _ = v85
                                                    }
                                                    _ = fi81
                                                    _ = v82
                                                }
                                            }
                                        default:
                                            for fi86 := 0; fi86 < 5; fi86++ {
                                                if true {
                                                    var v87 int16 = (v45 - (v45 | -32768))
                                                    v87 -= (v45 << ((1) & 63))
                                                    v1 += int64(v37)
                                                    _ = v87
                                                }
                                                switch "hello" {
                                                case "foo bar":
                                                    v1 = int64(fi86)
                                                case "":
                                                    v88 := GenS4{F0: int64(v45), F1: genFn0(127)}
                                                    _ = v88
                                                default:
                                                    v89 := GenS4{F0: (v39 << ((4294967295) & 63)), F1: v37}
                                                    _ = v89
                                                }
                                                _ = fi86
                                            }
                                        }
                                        fmt.Println(v6, v1)
                                    }
                                    v90 := GenS4{F0: -9223372036854775808, F1: genFn0(100)}
                                    _ = v90
                                }
                                _ = fi74
                            }
                        }
                    }
                case 577:
                    var v91 uint = 65535
                    _ = v91
                default:
                    delete(v3, 25)
                }
            }
            _ = v45
        }
    }
    for fi92 := 0; fi92 < 19; fi92++ {
        v37 = (v37 ^ (v37 | genFn0(-90)))
        _ = fi92
    }
    v1 += 1
    fmt.Println(v39)
    if (!true) {
        v93 := []int{42, -21}
        var v94 int64 = v6
        v95 := v4["a"]
        _ = v93
        _ = v94
        _ = v95
    } else {
        for fi96 := 0; fi96 < 8; fi96++ {
            v39 = -480
            if false {
                v97 := GenS4{F0: 213, F1: v37}
                var v98 int = (fi96 * 128)
                _ = v97
                _ = v98
            }
            switch -159 {
            case -2147483648:
                var v99 float32 = ((0.1 / float32(((v2 + 1e10) - 0.0))) / 3.14159)
                _ = v99
            default:
                delete(v4, ("a" + ("foo bar" + "foo bar")))
            }
            _ = fi96
        }
        for fi100 := 0; fi100 < 8; fi100++ {
            switch fi100 {
            case 9223372036854775807:
                if true {
                    switch ((("foo bar" + "hello") + "\n\t") + "foo bar") {
                    case "a":
                        v6 -= (v39 % ((v39) | 1))
                    default:
                        var v101 uint = 1
                        _ = v101
                    }
                    var v102 uint32 = 575
                    _ = v102
                }
            case 0:
                for fi103 := 0; fi103 < 18; fi103++ {
                    var v104 int32 = 238
                    var v105 float32 = ((((1e10 * 3.14159) / 1e10) + (((3.14159 - -2.25) + (3.14159 - (((-2.25 / 1.5) / (-2.25 / 0.1)) + 0.0))) + 1e10)) - 0.0)
                    _ = fi103
                    _ = v104
                    _ = v105
                }
            }
            var v106 uint16 = 32767
            v107 := GenS4{F0: (v6 + v39), F1: (^(v37))}
            _ = fi100
            _ = v106
            _ = v107
        }
    }
    v37 -= (v37 - v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
