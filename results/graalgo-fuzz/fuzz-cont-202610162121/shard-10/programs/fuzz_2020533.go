package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 bool; F2 int }
type GenS9 struct { F0 string; F1 int64 }
type GenS10 struct { F0 uint; F1 uint; F2 uint64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 0
    v3 := make(map[uint32]int)
    var v4 int32 = -1
    switch "\n\t" {
    case "foo bar":
        if true {
            for fi5 := 0; fi5 < 20; fi5++ {
                v2 = 4294967295
                _ = len(v3)
                _ = fi5
            }
            v3[2147483647] = 127
            v4 = v4
        } else {
            v6 := []int{}
            var v7 int32 = (v4 | v4)
            switch "" {
            case "hello":
                v3[561] = 0
            case "foo bar":
                var v8 int = int(v2)
                _ = v8
            case "":
                v9 := GenS1{F0: 255, F1: 0}
                _ = v9
            }
            _ = v6
            _ = v7
        }
    case "hello":
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := []int{}
            if true {
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 int8 = 100
                    _ = fi12
                    _ = v13
                }
                var v14 string = "foo bar"
                _ = v14
            }
            switch "a" {
            case "foo bar":
                var v15 uint8 = 42
                _ = v15
            case "a":
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 float64 = 1e10
                    if ((!false) || (32767 <= 1)) {
                        v18 := GenS1{F0: v2, F1: 364}
                        if (!(false && (0 < 663))) {
                            var v19 int = (fi10 % ((227) | 1))
                            var v20 int64 = 283
                            switch fi10 {
                            case 32768:
                                var v21 int = 32767
                                _ = v21
                            case -291:
                                if len(v11) > 0 {
                                    v11[0] = (fi10 >> ((v2) & 63))
                                }
                            case 127:
                                v22 := v1[("\n\t" + (("\n\t" + ("foo bar" + "")) + ""))]
                                _ = v22
                            }
                            _ = v19
                            _ = v20
                        }
                        var v23 uint64 = 992
                        _ = v18
                        _ = v23
                    }
                    v24 := GenS1{F0: (v2 / ((v2) | 1)), F1: 255}
                    v2 = 4294967295
                    if false {
                        v17 += 0.0
                        var v25 int8 = -128
                        switch 0 {
                        case -1:
                            var v26 float64 = v17
                            _ = v26
                        case -332:
                            var v27 int8 = v25
                            _ = v27
                        }
                        _ = len(v1)
                        var v28 string = "hello"
                        fmt.Println(v4, v25, v2)
                        var v29 int16 = -32768
                        switch (fi10 + fi16) {
                        case 32767:
                            if (!true) {
                                if true {
                                    v17 -= v17
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch 644 {
                                            case 615:
                                                v3[256] = 0
                                            case 1:
                                                v30 := GenS2{F0: (v4 ^ (v4 - -2147483648)), F1: v25}
                                                _ = v30
                                            default:
                                                fmt.Println(fi10, v29, fi10)
                                                if false {
                                                    switch fi10 {
                                                    case 0:
                                                        v31 := GenS2{F0: -1, F1: v25}
                                                        _ = v31
                                                    case -1:
                                                        v32 := GenS3{F0: 32767}
                                                        _ = v32
                                                    }
                                                } else {
                                                    v33 := GenS3{F0: 2}
                                                    var v34 uint = v2
                                                    v17 -= v17
                                                    v11 = append(v11, fi10)
                                                    fmt.Println(v29)
                                                    if false {
                                                        if true {
                                                            if false {
                                                                var v35 float32 = (1e10 * 1.5)
                                                                var v36 int64 = 0
                                                                _ = v35
                                                                _ = v36
                                                            }
                                                        } else {
                                                            v29 -= (-(v29))
                                                        }
                                                        var v37 uint64 = 719
                                                        _ = v37
                                                    } else {
                                                        switch (^(fi10)) {
                                                        case 9223372036854775807:
                                                            var v38 int32 = 2147483647
                                                            _ = v38
                                                        case 2147483647:
                                                            fmt.Println(v2, v34)
                                                        }
                                                        v29 = (-(v29))
                                                    }
                                                    switch fi10 {
                                                    case 9223372036854775807:
                                                        v11 = append(v11, (fi16 << (((v2 & v34)) & 63)))
                                                    }
                                                    if (1 != 533) {
                                                        var v39 uint32 = 37
                                                        if false {
                                                            v39 = (v39 << ((v34) & 63))
                                                        } else {
                                                            var v40 uint = v34
                                                            _ = v40
                                                        }
                                                        v41 := GenS3{F0: 65535}
                                                        _ = v39
                                                        _ = v41
                                                    } else {
                                                        v29 += 1000
                                                        v42 := GenS2{F0: v4, F1: (v25 % (((v25 ^ v25)) | 1))}
                                                        var v43 float32 = (float32(1e10) / 1.5)
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    _ = v33
                                                    _ = v34
                                                }
                                            }
                                        default:
                                            v29 = (v29 & v29)
                                        }
                                    case "":
                                        v28 = v28
                                        var v44 uint8 = 42
                                        var v45 uint16 = 255
                                        _ = v44
                                        _ = v45
                                    case "a":
                                        v25 = 6
                                    default:
                                        v25 -= (v25 >> ((514) & 63))
                                    }
                                } else {
                                    var v46 int8 = (v25 + (v25 ^ (v25 + (v25 & 100))))
                                    _ = v46
                                }
                            }
                            if false {
                                if (!(!true)) {
                                    var v47 string = v28
                                    var v48 float32 = 1e10
                                    _ = v47
                                    _ = v48
                                } else {
                                    v49 := GenS4{F0: (^(v4))}
                                    v50 := GenS3{F0: 42}
                                    switch (fi16 / ((fi16) | 1)) {
                                    case 4294967295:
                                        var v51 uint16 = 32767
                                        _ = v51
                                    default:
                                        v52 := GenS4{F0: (-(v4))}
                                        _ = v52
                                    }
                                    _ = v49
                                    _ = v50
                                }
                                v28 = (v28 + "")
                                if len(v11) > 0 {
                                    v11[0] = 32768
                                }
                            }
                        case 32768:
                            var v53 int32 = -2147483648
                            _ = v53
                        }
                        _ = v25
                        _ = v28
                        _ = v29
                    } else {
                        fmt.Println(v17, v17, v17)
                    }
                    _ = fi16
                    _ = v17
                    _ = v24
                }
            default:
                v54 := GenS1{F0: v2, F1: 65535}
                _ = v54
            }
            _ = fi10
            _ = v11
        }
    case "\n\t":
        for fi55 := 0; fi55 < 7; fi55++ {
            if (false || (-1 >= 0)) {
                for fi56 := 0; fi56 < 13; fi56++ {
                    v2 = 255
                    var v57 string = ""
                    _ = fi56
                    _ = v57
                }
                v58 := []int{}
                _ = v58
            }
            _ = fi55
        }
    }
    switch 1 {
    case -9223372036854775808:
        fmt.Println(v2, v4)
    case -341:
        if ((1 <= 492) || (!true)) {
            var v59 uint8 = 2
            v60 := GenS5{F0: 127}
            if false {
                for fi61 := 0; fi61 < 1; fi61++ {
                    v2 = v2
                    if (!false) {
                        var v62 uint32 = 2147483647
                        v63 := GenS6{F0: (^(v62)), F1: v62}
                        v59 += v59
                        _ = v62
                        _ = v63
                    } else {
                        var v64 bool = false
                        var v65 int32 = 508
                        for fi66 := 0; fi66 < 19; fi66++ {
                            switch ("" + ((((("" + "foo bar") + "a") + "a") + "") + "hello")) {
                            case "\n\t":
                                v67 := GenS7{F0: v59}
                                _ = v67
                            }
                            if v64 {
                                if (!false) {
                                    v68 := []int{fi66, (^(fi61)), fi61}
                                    v3[1] = fi61
                                    v69 := v3[927]
                                    _ = v68
                                    _ = v69
                                }
                                if ((fi61 >> ((v2) & 63)) >= (fi66 | fi66)) {
                                    var v70 int32 = v65
                                    _ = v70
                                }
                                var v71 uint16 = 32767
                                _ = v71
                            } else {
                                var v72 float32 = 3.14159
                                _ = v72
                            }
                            v73 := GenS5{F0: v59}
                            if (32767 != 570) {
                                fmt.Println(v59, v2, v59)
                            }
                            v4 = v4
                            _ = fi66
                            _ = v73
                        }
                        _ = v64
                        _ = v65
                    }
                    v4 = 1
                    _ = fi61
                }
                v74 := GenS8{F0: ((-2.25 * 3.14159) >= 0.1), F1: true, F2: 127}
                v4 -= v4
                _ = v74
            }
            _ = v59
            _ = v60
        }
    default:
        var v75 int8 = 100
        _ = v75
    }
    var v76 uint64 = 234
    if (!((v76 * v76) > (v76 * v76))) {
        v4 = 0
    }
    var v77 int32 = (v4 + (v4 + -1))
    var v78 uint8 = 128
    if true {
        for fi79 := 0; fi79 < 9; fi79++ {
            for fi80 := 0; fi80 < 10; fi80++ {
                v81 := []int{}
                switch "a" {
                case "\n\t":
                    v77 -= 42
                default:
                    if ((^(v77)) != (v77 & (-(v77)))) {
                        var v82 uint32 = 2147483648
                        switch 32768 {
                        case 985:
                            var v83 uint32 = (^(v82))
                            _ = v83
                        case 705:
                            delete(v1, "\n\t")
                        case 717:
                            v84 := GenS2{F0: v4, F1: 127}
                            _ = v84
                        default:
                            v2 = v2
                            _ = len(v81)
                            v85 := GenS3{F0: 67}
                            _ = v85
                        }
                        _ = v82
                    }
                }
                _ = fi80
                _ = v81
            }
            for fi86 := 0; fi86 < 1; fi86++ {
                v87 := v1[("" + "\n\t")]
                v88 := v1["foo bar"]
                _ = fi86
                _ = v87
                _ = v88
            }
            _ = fi79
        }
        var v89 int8 = 34
        switch 32768 {
        case 801:
            for fi90 := 0; fi90 < 16; fi90++ {
                fmt.Println(v77, v76)
                var v91 uint32 = 65535
                _ = fi90
                _ = v91
            }
        default:
            v92 := GenS9{F0: "", F1: -1}
            _ = v92
        }
        _ = v89
    }
    var v93 bool = false
    for fi94 := 0; fi94 < 20; fi94++ {
        for fi95 := 0; fi95 < 6; fi95++ {
            if true {
                v96 := []int{fi95, 9223372036854775807, (fi95 / ((fi94) | 1)), fi94}
                _ = v96
            } else {
                delete(v1, "foo bar")
            }
            v97 := []int{}
            var v98 uint = (v2 | (v2 / (((v2 & (v2 / ((v2) | 1)))) | 1)))
            _ = fi95
            _ = v97
            _ = v98
        }
        v2 = (v2 >> ((uint(v76)) & 63))
        if v93 {
            v99 := GenS10{F0: (v2 | v2), F1: (^(v2)), F2: (v76 ^ v76)}
            _ = v99
        }
        _ = fi94
    }
    v4 -= -1
    fmt.Println(v4)
    v100 := GenS11{F0: 1000}
    v101 := GenS12{F0: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v76
    _ = v77
    _ = v78
    _ = v93
}
