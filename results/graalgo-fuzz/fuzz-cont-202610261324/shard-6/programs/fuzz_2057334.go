package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int8; F2 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 int; F1 string; F2 uint }
type GenS8 struct { F0 uint8 }
func genFn0(p0 uint) int64 {
    if false {
        v5 := GenS1{F0: 219, F1: 2, F2: -648}
        _ = v5
    }
    p0 = p0
    if true {
        if false {
            switch -2147483648 {
            case 197:
                v6 := []int{-71, 32767, 42}
                _ = v6
            }
            v7 := GenS1{F0: 32767, F1: 225, F2: -461}
            _ = v7
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            if (1 >= 42) {
                for fi9 := 0; fi9 < 3; fi9++ {
                    fmt.Println(p0, fi9)
                    if true {
                        fmt.Println(fi8)
                        v10 := []int{(fi8 ^ fi9), (fi9 << ((p0) & 63)), fi9}
                        _ = v10
                    }
                    p0 = 255
                    var v11 float32 = 1e10
                    _ = fi9
                    _ = v11
                }
            }
            _ = fi8
        }
        fmt.Println(p0)
    } else {
        var v12 string = "foo bar"
        v13 := GenS1{F0: -1, F1: 0, F2: 32767}
        v14 := GenS2{F0: 431}
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    return 42
}
func genFn1() int8 {
    var v15 int = -981
    switch (v15 + v15) {
    case -2147483648:
        var v16 int8 = 42
        _ = v16
    case 42:
        v17 := []int{v15, 32767, (v15 + 2147483647), 1}
        _ = v17
    }
    fmt.Println(v15)
    switch v15 {
    case -2147483648:
        v18 := []int{}
        _ = v18
    }
    v19 := GenS2{F0: 42}
    _ = v15
    _ = v19
    return 127
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 504
    v3 := make(map[uint32]int)
    var v4 int16 = -32768
    switch ("hello" + "hello") {
    case "":
        var v20 bool = (!(false || (false || true)))
        _ = v20
    }
    v2 += 2
    if (true && (!true)) {
        var v21 uint = 0
        if (false || (!((!true) && (!false)))) {
            if (!false) {
                fmt.Println(v21)
                var v22 uint64 = (^(v2))
                for fi23 := 0; fi23 < 2; fi23++ {
                    delete(v1, fi23)
                    _ = fi23
                }
                _ = v22
            } else {
                var v24 float32 = 1.5
                if false {
                    if true {
                        var v25 uint16 = 42
                        v26 := v3[1]
                        v27 := []int{(v26 * -1)}
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 18; fi28++ {
                        for fi29 := 0; fi29 < 7; fi29++ {
                            v4 = 0
                            var v30 uint32 = 328
                            v30 += v30
                            _ = fi29
                            _ = v30
                        }
                        if true {
                            switch ((("\n\t" + "\n\t") + "a") + "a") {
                            case "\n\t":
                                for fi31 := 0; fi31 < 7; fi31++ {
                                    if false {
                                        delete(v1, fi31)
                                        switch 32767 {
                                        case 32767:
                                            delete(v3, 803)
                                        case 32768:
                                            if true {
                                                v4 -= (v4 - v4)
                                                var v32 int64 = 0
                                                var v33 bool = (true || true)
                                                _ = v32
                                                _ = v33
                                            }
                                        case 2147483647:
                                            fmt.Println(fi28)
                                        }
                                    } else {
                                        switch 4294967295 {
                                        case 0:
                                            _ = len(v1)
                                            delete(v1, fi31)
                                            switch 127 {
                                            case -988:
                                                var v34 int32 = 1
                                                _ = v34
                                            case 215:
                                                var v35 int32 = -908
                                                _ = v35
                                            }
                                        }
                                        v24 = v24
                                        var v36 string = ("hello" + "foo bar")
                                        _ = v36
                                    }
                                    _ = fi31
                                }
                            case "a":
                                for fi37 := 0; fi37 < 12; fi37++ {
                                    v2 = 20
                                    v38 := []int{fi37, (fi37 & (fi37 & fi37)), int(v4), (fi28 - -564)}
                                    _ = fi37
                                    _ = v38
                                }
                            case "":
                                v39 := GenS2{F0: v4}
                                _ = v39
                            default:
                                v40 := []int{-9223372036854775808, fi28, (fi28 & fi28), 127}
                                _ = v40
                            }
                        }
                        delete(v1, fi28)
                        _ = fi28
                    }
                } else {
                    v41 := GenS1{F0: 128, F1: 42, F2: 32767}
                    _ = v41
                }
                _ = v24
            }
            var v42 int8 = 42
            for fi43 := 0; fi43 < 20; fi43++ {
                v42 -= (v42 * (v42 >> ((42) & 63)))
                fmt.Println(v42, v2, v4)
                _ = fi43
            }
            _ = v42
        }
        _ = v21
    }
    if false {
        v2 = v2
        v4 = (v4 & (-(v4)))
    } else {
        switch 9223372036854775807 {
        case 2147483647:
            for fi44 := 0; fi44 < 11; fi44++ {
                if true {
                    for fi45 := 0; fi45 < 20; fi45++ {
                        var v46 uint16 = 255
                        if false {
                            switch (fi44 + fi45) {
                            case 32768:
                                var v47 int32 = 0
                                _ = v47
                            case 9223372036854775807:
                                v48 := v3[2147483648]
                                _ = v48
                            case 4294967295:
                                v2 = v2
                            }
                        } else {
                            if (((140 <= 925) && (false || (0 <= 2147483647))) && true) {
                                var v49 int64 = -1
                                var v50 float64 = 1e10
                                var v51 int = (fi44 ^ (fi44 ^ 9223372036854775807))
                                v52 := []int{fi45}
                                _ = v49
                                _ = v50
                                _ = v51
                                _ = v52
                            } else {
                                fmt.Println(v2, v2, v2)
                                switch "hello" {
                                case "a":
                                    var v53 int = fi44
                                    _ = v53
                                default:
                                    v4 += -131
                                }
                                v54 := GenS3{F0: fi45, F1: 127, F2: true}
                                _ = v54
                            }
                            v55 := []int{(fi44 & 544), (fi45 | fi45), (fi44 * 4294967295), 9223372036854775807}
                            _ = v55
                        }
                        v56 := GenS4{F0: 0}
                        _ = fi45
                        _ = v46
                        _ = v56
                    }
                    fmt.Println(fi44, v4)
                } else {
                    for fi57 := 0; fi57 < 4; fi57++ {
                        v4 -= v4
                        fmt.Println(fi44, fi44)
                        _ = fi57
                    }
                }
                _ = fi44
            }
        case 4294967295:
            var v58 uint = 18446744073709551615
            _ = v58
        case 32768:
            if (127 > 1) {
                var v59 int16 = (v4 & v4)
                _ = v59
            } else {
                for fi60 := 0; fi60 < 20; fi60++ {
                    v61 := []int{(-(fi60))}
                    if true {
                        for fi62 := 0; fi62 < 20; fi62++ {
                            v63 := GenS3{F0: (fi60 - -1), F1: 0, F2: true}
                            v2 = (v2 | v2)
                            _ = fi62
                            _ = v63
                        }
                        var v64 bool = true
                        var v65 uint = 18446744073709551615
                        _ = v64
                        _ = v65
                    } else {
                        switch (fi60 | fi60) {
                        case -9223372036854775808:
                            var v66 int32 = -2147483648
                            _ = v66
                        default:
                            var v67 bool = ((!false) && true)
                            _ = v67
                        }
                    }
                    switch (fi60 + fi60) {
                    case 32768:
                        switch 127 {
                        case 9223372036854775807:
                            v68 := GenS4{F0: 32767}
                            _ = v68
                        case 1:
                            var v69 uint32 = uint32(v2)
                            _ = v69
                        case 32767:
                            v70 := GenS1{F0: fi60, F1: 42, F2: (fi60 + fi60)}
                            _ = v70
                        default:
                            v71 := GenS2{F0: (v4 / (((v4 ^ (v4 ^ v4))) | 1))}
                            _ = v71
                        }
                    }
                    _ = fi60
                    _ = v61
                }
                v4 = v4
            }
        default:
            v4 += 42
        }
    }
    for fi72 := 0; fi72 < 19; fi72++ {
        v2 += v2
        fmt.Println(v4, v4)
        _ = fi72
    }
    var v73 string = "hello"
    for fi74 := 0; fi74 < 9; fi74++ {
        if ((!((-2.25 - (1.5 + 3.14159)) <= (0.0 * 0.1))) || false) {
            switch v73 {
            case "hello":
                if ((^(fi74)) < (fi74 + (fi74 >> ((86) & 63)))) {
                    var v75 uint32 = 2147483647
                    v4 = 1
                    _ = v75
                }
            case "a":
                v4 -= 32767
            case "":
                v4 = v4
            default:
                v73 = ""
            }
            v4 += v4
        }
        _ = fi74
    }
    switch 128 {
    case 1:
        fmt.Println(v4, v73)
    case -9223372036854775808:
        for fi76 := 0; fi76 < 16; fi76++ {
            fmt.Println(v73)
            _ = len(v1)
            _ = fi76
        }
    }
    v77 := []int{}
    var v78 uint32 = 362
    for fi79 := 0; fi79 < 5; fi79++ {
        v4 = (v4 | 67)
        v80 := GenS2{F0: (v4 * -32768)}
        _ = fi79
        _ = v80
    }
    v2 = v2
    if ((-525 <= 32768) || true) {
        switch ("\n\t" + (v73 + (v73 + (v73 + "")))) {
        case "\n\t":
            v73 = v73
        case "":
            v81 := GenS5{F0: 42, F1: 4294967295}
            _ = v81
        }
        v77 = append(v77, 9223372036854775807)
    }
    var v82 int8 = 0
    v82 = (^(v82))
    for fi83 := 0; fi83 < 13; fi83++ {
        var v84 uint64 = v2
        for fi85 := 0; fi85 < 12; fi85++ {
            v86 := GenS6{F0: 0, F1: 2}
            _ = fi85
            _ = v86
        }
        if (false && true) {
            for fi87 := 0; fi87 < 18; fi87++ {
                v88 := GenS7{F0: fi83, F1: "foo bar", F2: 255}
                if true {
                    switch -9223372036854775808 {
                    case -1:
                        var v89 int = int(v78)
                        _ = v89
                    case -9223372036854775808:
                        var v90 int64 = int64(v78)
                        _ = v90
                    case 1:
                        if true {
                            v91 := GenS4{F0: (v4 * (v4 >> ((42) & 63)))}
                            v73 = ((v73 + v73) + v73)
                            _ = v91
                        }
                    }
                } else {
                    v92 := GenS6{F0: 42, F1: 1}
                    var v93 uint8 = 0
                    v94 := GenS8{F0: 1}
                    fmt.Println(v4, fi87)
                    var v95 uint = 255
                    v4 = v4
                    var v96 uint64 = 42
                    _ = v92
                    _ = v93
                    _ = v94
                    _ = v95
                    _ = v96
                }
                _ = fi87
                _ = v88
            }
        } else {
            if (!(126 < 2)) {
                if false {
                    v3[(^(v78))] = (fi83 << ((1) & 63))
                    v78 = (v78 ^ v78)
                }
                switch fi83 {
                case -965:
                    switch "hello" {
                    case "a":
                        var v97 string = v73
                        _ = v97
                    default:
                        v98 := GenS6{F0: -1, F1: 220}
                        _ = v98
                    }
                case 32767:
                    var v99 float32 = 3.14159
                    _ = v99
                case -1:
                    _ = len(v1)
                }
            }
            v100 := GenS5{F0: -237, F1: v78}
            _ = v100
        }
        _ = fi83
        _ = v84
    }
    var v101 float64 = ((1.5 * float64(0.1)) + 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v73:", v73)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v73
    _ = v77
    _ = v78
    _ = v82
}
