package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int32; F1 bool; F2 int32 }

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]uint32)
    var v3 bool = true
    var v4 uint = 4294967295
    v1 += 255
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            v4 += v1
            if (v3 || (!((!v3) && (!(true && v3))))) {
                if v3 {
                    fmt.Println(fi5)
                    if ((v4 % ((v1) | 1)) >= (v4 & (^(v1)))) {
                        fmt.Println(fi5, fi5, fi6)
                        fmt.Println(fi5)
                        v7 := GenS1{F0: 42, F1: 0}
                        _ = v7
                    }
                } else {
                    v8 := GenS2{F0: 552, F1: 744, F2: 42}
                    _ = v8
                }
            }
            var v9 uint16 = 65535
            var v10 float64 = 1e10
            _ = fi6
            _ = v10
            _ = v9
        }
        if ((v1 | v1) > (v4 << ((v4) & 63))) {
            var v11 int64 = 42
            var v12 uint = v4
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                v2[2] = 600
                v1 = 4294967295
                fmt.Println(v3, fi5)
                v3 = v3
                if v3 {
                    v3 = (128 <= fi5)
                    fmt.Println(v1, v4)
                } else {
                    if true {
                        switch ("a" + "a") {
                        case "hello":
                            if true {
                                var v14 uint8 = 128
                                _ = v14
                            } else {
                                v3 = v3
                                var v15 bool = (!(!v3))
                                if v3 {
                                    v15 = (v15 || v3)
                                }
                                _ = v15
                            }
                        case "":
                            v1 = 1
                        case "foo bar":
                            if ((!v3) && v3) {
                                fmt.Println(fi5, fi5)
                                if v3 {
                                    v1 += v1
                                    v16 := GenS2{F0: 234, F1: -136, F2: 1}
                                    fmt.Println(fi13, v4)
                                    _ = v16
                                }
                            }
                        }
                        switch fi5 {
                        case 42:
                            v17 := []int{}
                            _ = v17
                        case -1:
                            var v18 uint16 = 2
                            _ = v18
                        default:
                            if (!v3) {
                                v19 := GenS1{F0: 57, F1: 2}
                                v3 = v3
                                if (4294967295 < 46) {
                                    fmt.Println(fi5, v4, v3)
                                    var v20 bool = (1 > 2147483647)
                                    switch fi13 {
                                    case -9223372036854775808:
                                        var v21 uint16 = 667
                                        _ = v21
                                    default:
                                        var v22 int16 = 1000
                                        _ = v22
                                    }
                                    _ = v20
                                } else {
                                    if (-160 > (^(fi5))) {
                                        if (!v3) {
                                            v23 := GenS2{F0: 245, F1: -156, F2: 767}
                                            var v24 int8 = 127
                                            var v25 uint8 = 1
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        var v26 int = fi13
                                        if (!v3) {
                                            fmt.Println(v1)
                                            var v27 int = -788
                                            var v28 uint64 = 18446744073709551615
                                            _ = v27
                                            _ = v28
                                        }
                                        _ = v26
                                    } else {
                                        if (!(v3 || (4294967295 <= 0))) {
                                            var v29 int64 = 1
                                            switch fi13 {
                                            case 0:
                                                v30 := []int{(fi13 - fi5)}
                                                _ = v30
                                            }
                                            _ = v29
                                        }
                                        v31 := GenS1{F0: -1, F1: 77}
                                        var v32 uint64 = 0
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                _ = v19
                            }
                        }
                        v1 = 999
                    }
                    var v33 int = (fi5 | 32767)
                    var v34 uint64 = 4294967295
                    var v35 uint64 = v34
                    _ = v33
                    _ = v34
                    _ = v35
                }
                _ = fi13
            }
            if v3 {
                var v36 uint8 = 0
                var v37 uint32 = 0
                var v38 uint8 = (v36 << ((744) & 63))
                _ = v36
                _ = v37
                _ = v38
            } else {
                var v39 uint64 = 4294967295
                fmt.Println(v4, fi5)
                switch (-(fi5)) {
                case 4294967295:
                    v40 := GenS1{F0: -1, F1: 127}
                    _ = v40
                }
                _ = v39
            }
            v41 := []int{fi5}
            _ = v41
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            switch (fi42 >> ((v1) & 63)) {
            case 32768:
                v43 := []int{}
                _ = v43
            default:
                v44 := GenS3{F0: -225}
                _ = v44
            }
            var v45 bool = v3
            _ = fi42
            _ = v45
        }
        _ = fi5
    }
    for fi46 := 0; fi46 < 11; fi46++ {
        v47 := []int{(^(fi46)), fi46}
        _ = fi46
        _ = v47
    }
    v48 := GenS4{F0: 32767}
    v1 = v1
    switch 884 {
    case 9223372036854775807:
        for fi49 := 0; fi49 < 4; fi49++ {
            v50 := GenS5{F0: 4294967295, F1: v1}
            _ = fi49
            _ = v50
        }
    default:
        var v51 int8 = 42
        _ = v51
    }
    switch 0 {
    case 0:
        var v52 uint16 = 2
        _ = v52
    }
    switch "a" {
    case "foo bar":
        v53 := []int{781, 1, 32767}
        _ = v53
    case "":
        var v54 int = -1
        _ = v54
    }
    fmt.Println(v4)
    for fi55 := 0; fi55 < 7; fi55++ {
        var v56 uint8 = 2
        if (552 == -1) {
            var v57 bool = true
            v58 := GenS3{F0: 42}
            _ = v57
            _ = v58
        }
        _ = fi55
        _ = v56
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        if (18446744073709551615 > 1) {
            v60 := GenS5{F0: 0, F1: v1}
            var v61 uint8 = 42
            for fi62 := 0; fi62 < 11; fi62++ {
                v4 -= (^(v1))
                _ = fi62
            }
            _ = v60
            _ = v61
        } else {
            v63 := []int{9223372036854775807, fi59}
            _ = v63
        }
        for fi64 := 0; fi64 < 6; fi64++ {
            v65 := []int{}
            v1 += (^(v4))
            var v66 int8 = 100
            _ = fi64
            _ = v65
            _ = v66
        }
        var v67 float32 = 3.14159
        _ = fi59
        _ = v67
    }
    v3 = v3
    switch (("hello" + ("foo bar" + ("\n\t" + "\n\t"))) + "hello") {
    case "foo bar":
        var v68 string = "foo bar"
        _ = v68
    case "\n\t":
        for fi69 := 0; fi69 < 1; fi69++ {
            fmt.Println(v1, fi69)
            _ = fi69
        }
    default:
        for fi70 := 0; fi70 < 2; fi70++ {
            switch (-(fi70)) {
            case 127:
                var v71 uint64 = 4294967295
                _ = v71
            }
            var v72 int64 = 438
            if true {
                switch 32767 {
                case 4294967295:
                    var v73 int32 = 0
                    _ = v73
                case 2147483647:
                    v4 += 2
                case 782:
                    if (!(!true)) {
                        fmt.Println(v4, v3, v4)
                        var v74 uint = v4
                        _ = v74
                    } else {
                        v72 = v72
                        if v3 {
                            var v75 float32 = ((1.5 / 1.5) * 1.5)
                            _ = v75
                        } else {
                            if v3 {
                                for fi76 := 0; fi76 < 18; fi76++ {
                                    fmt.Println(v1, fi70)
                                    v77 := []int{(fi70 / ((32767) | 1)), (fi76 & (fi76 / ((fi70) | 1))), (fi70 % ((fi70) | 1)), int(v72)}
                                    _ = fi76
                                    _ = v77
                                }
                            }
                            fmt.Println(v4, v1)
                        }
                    }
                default:
                    switch fi70 {
                    case -9223372036854775808:
                        var v78 bool = false
                        _ = v78
                    }
                }
            }
            _ = fi70
            _ = v72
        }
    }
    var v79 float64 = (-2.25 + 0.0)
    switch "" {
    case "":
        var v80 uint8 = 255
        _ = v80
    case "hello":
        switch "a" {
        case "\n\t":
            fmt.Println(v1, v4)
        case "foo bar":
            for fi81 := 0; fi81 < 7; fi81++ {
                if true {
                    if false {
                        if (573 >= 0) {
                            if v3 {
                                v82 := []int{(-(fi81)), fi81, (fi81 & (fi81 ^ (^(fi81))))}
                                v83 := GenS3{F0: 42}
                                var v84 bool = (127 <= 37)
                                _ = v82
                                _ = v83
                                _ = v84
                            } else {
                                for fi85 := 0; fi85 < 20; fi85++ {
                                    fmt.Println(v3)
                                    v86 := GenS2{F0: 42, F1: 1, F2: 0}
                                    fmt.Println(v4)
                                    v87 := GenS3{F0: 0}
                                    _ = fi85
                                    _ = v86
                                    _ = v87
                                }
                                if v3 {
                                    var v88 int64 = -1
                                    if v3 {
                                        delete(v2, uint32(v1))
                                        if (2 != uint8(v1)) {
                                            v89 := GenS1{F0: 1, F1: 52}
                                            if true {
                                                v1 -= 42
                                                if v3 {
                                                    var v90 uint = (^(v4))
                                                    delete(v2, 4294967295)
                                                    _ = v90
                                                }
                                            } else {
                                                switch (^(fi81)) {
                                                case -2147483648:
                                                    var v91 float32 = 1.5
                                                    _ = v91
                                                }
                                            }
                                            var v92 int8 = -128
                                            _ = v89
                                            _ = v92
                                        } else {
                                            var v93 int64 = (v88 | v88)
                                            _ = v93
                                        }
                                    }
                                    fmt.Println(v4)
                                    _ = v88
                                } else {
                                    v94 := []int{}
                                    for fi95 := 0; fi95 < 9; fi95++ {
                                        if (v3 || false) {
                                            var v96 uint16 = 65535
                                            _ = v96
                                        } else {
                                            v97 := GenS2{F0: 255, F1: -9223372036854775808, F2: 1}
                                            v98 := GenS6{F0: -795, F1: (0 == 873), F2: 1}
                                            _ = v97
                                            _ = v98
                                        }
                                        _ = fi95
                                    }
                                    _ = v94
                                }
                                if v3 {
                                    v99 := v2[1]
                                    _ = v99
                                } else {
                                    if v3 {
                                        v100 := GenS4{F0: 32767}
                                        var v101 uint = v4
                                        _ = v100
                                        _ = v101
                                    } else {
                                        var v102 uint16 = 0
                                        _ = v102
                                    }
                                }
                            }
                            fmt.Println(v1, v4)
                            var v103 uint8 = 42
                            _ = v103
                        } else {
                            v4 = 544
                        }
                        if (-128 != -128) {
                            v4 -= 643
                            v79 += (v79 + (v79 + 0.0))
                            v104 := GenS4{F0: 0}
                            _ = v104
                        }
                        if v3 {
                            v1 = (^(v1))
                            for fi105 := 0; fi105 < 3; fi105++ {
                                v106 := GenS2{F0: 797, F1: -1, F2: 2147483647}
                                _ = fi105
                                _ = v106
                            }
                            var v107 uint16 = 789
                            _ = v107
                        }
                    }
                    var v108 uint16 = 42
                    v3 = (842 <= 225)
                    _ = v108
                }
                var v109 uint8 = 165
                _ = fi81
                _ = v109
            }
        case "hello":
            switch "\n\t" {
            case "foo bar":
                switch "foo bar" {
                case "\n\t":
                    var v110 int32 = 0
                    _ = v110
                default:
                    v1 = v1
                }
            case "a":
                var v111 uint8 = 124
                _ = v111
            case "\n\t":
                v79 -= v79
            }
        default:
            v112 := GenS2{F0: 255, F1: 9223372036854775807, F2: 2147483647}
            _ = v112
        }
    case "foo bar":
        fmt.Println(v3)
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v79
}
