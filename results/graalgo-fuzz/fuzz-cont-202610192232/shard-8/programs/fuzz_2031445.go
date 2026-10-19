package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int32 }

func main() {
    var v1 uint8 = 251
    var v2 uint64 = 4294967295
    var v3 float64 = float64((1e10 + 0.1))
    var v4 uint32 = 190
    v5 := make(map[string]string)
    if (4294967295 == (v4 * 0)) {
        v6 := GenS1{F0: 0}
        v7 := []int{847, 128, -2147483648}
        _ = v6
        _ = v7
    }
    if true {
        var v8 int32 = 2147483647
        _ = v8
    } else {
        var v9 uint64 = v2
        _ = v9
    }
    var v10 float32 = 1.5
    if true {
        var v11 int64 = -829
        _ = v11
    } else {
        var v12 int16 = 1
        var v13 int = 32767
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        switch fi14 {
        case -417:
            v10 -= v10
        case 32767:
            for fi15 := 0; fi15 < 20; fi15++ {
                switch fi14 {
                case 224:
                    var v16 int8 = -101
                    _ = v16
                case 0:
                    var v17 float32 = 1.5
                    _ = v17
                default:
                    var v18 int64 = 42
                    _ = v18
                }
                switch fi15 {
                case 9223372036854775807:
                    if (!(!true)) {
                        var v19 bool = false
                        v20 := []int{}
                        _ = v19
                        _ = v20
                    } else {
                        v3 -= v3
                        var v21 uint64 = v2
                        var v22 uint16 = uint16(v4)
                        if false {
                            if ("" > ("foo bar" + ((("\n\t" + "\n\t") + ("a" + ((("a" + "a") + "foo bar") + "foo bar"))) + (("foo bar" + "foo bar") + "foo bar")))) {
                                v23 := GenS1{F0: int8(v22)}
                                var v24 int8 = 42
                                fmt.Println(v2, v10, v10)
                                _ = v23
                                _ = v24
                            }
                            v2 = (v2 % ((4294967295) | 1))
                            var v25 string = ""
                            var v26 uint64 = (^(v2))
                            v27 := GenS1{F0: -1}
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            v28 := GenS1{F0: 42}
                            if true {
                                fmt.Println(v2)
                                v1 = 128
                            }
                            _ = v28
                        }
                        v29 := GenS1{F0: -128}
                        var v30 uint = 892
                        v1 = v1
                        _ = v21
                        _ = v22
                        _ = v29
                        _ = v30
                    }
                }
                _ = fi15
            }
        case -579:
            v31 := GenS2{F0: -2147483648}
            _ = v31
        }
        switch (fi14 % (((-(fi14))) | 1)) {
        case 872:
            v32 := v5[""]
            _ = v32
        default:
            for fi33 := 0; fi33 < 1; fi33++ {
                var v34 int = 301
                _ = fi33
                _ = v34
            }
        }
        v35 := GenS2{F0: 307}
        _ = fi14
        _ = v35
    }
    var v36 uint16 = 255
    var v37 int32 = 1
    delete(v5, "\n\t")
    if false {
        var v38 uint16 = v36
        v4 += v4
        v39 := GenS1{F0: 42}
        _ = v38
        _ = v39
    } else {
        switch 824 {
        case 32767:
            var v40 uint16 = v36
            _ = v40
        }
        switch -9223372036854775808 {
        case 127:
            if (255 <= 65535) {
                var v41 string = "hello"
                for fi42 := 0; fi42 < 14; fi42++ {
                    fmt.Println(v2, v4, v3)
                    var v43 uint = 281
                    delete(v5, "a")
                    _ = fi42
                    _ = v43
                }
                v44 := GenS3{F0: (v37 >> ((255) & 63)), F1: (v37 << ((2) & 63)), F2: 42}
                _ = v41
                _ = v44
            } else {
                for fi45 := 0; fi45 < 17; fi45++ {
                    var v46 uint16 = 738
                    var v47 int64 = 0
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                var v48 float32 = -2.25
                _ = v48
            }
        case -197:
            v36 = 32767
        }
        fmt.Println(v4, v2)
    }
    for fi49 := 0; fi49 < 7; fi49++ {
        v50 := GenS2{F0: -184}
        _ = fi49
        _ = v50
    }
    for fi51 := 0; fi51 < 8; fi51++ {
        var v52 uint8 = 2
        switch -227 {
        case -9223372036854775808:
            if true {
                switch (("foo bar" + ("" + "\n\t")) + ("hello" + "foo bar")) {
                case "":
                    var v53 uint16 = (v36 % ((v36) | 1))
                    _ = v53
                case "\n\t":
                    var v54 uint32 = (v4 * v4)
                    _ = v54
                case "a":
                    var v55 int8 = -128
                    _ = v55
                default:
                    v56 := GenS1{F0: 100}
                    _ = v56
                }
                fmt.Println(v2, v52, v36)
                var v57 uint64 = 18446744073709551615
                _ = v57
            } else {
                for fi58 := 0; fi58 < 7; fi58++ {
                    switch "a" {
                    case "a":
                        if true {
                            var v59 string = ""
                            v3 -= (v3 - v3)
                            if true {
                                if true {
                                    v5[v59] = ("a" + "hello")
                                    if (((v37 << ((18446744073709551615) & 63)) < v37) && (!false)) {
                                        v60 := GenS1{F0: -8}
                                        _ = v60
                                    }
                                    if false {
                                        _ = len(v5)
                                        fmt.Println(v10, fi58)
                                        switch v59 {
                                        case "a":
                                            v61 := []int{}
                                            _ = v61
                                        case "hello":
                                            fmt.Println(v2, v52)
                                            v52 += v52
                                            v62 := []int{(fi58 % (((fi58 * fi58)) | 1)), -252}
                                            var v63 int32 = v37
                                            _ = v62
                                            _ = v63
                                        case "\n\t":
                                            switch (v59 + "a") {
                                            case "":
                                                v1 = (v1 + v1)
                                            case "a":
                                                var v64 uint16 = uint16(v4)
                                                _ = v64
                                            }
                                        default:
                                            var v65 bool = true
                                            _ = v65
                                        }
                                        var v66 int32 = 2147483647
                                        v67 := []int{int(v4), fi51}
                                        _ = v66
                                        _ = v67
                                    } else {
                                        if false {
                                            var v68 string = ("\n\t" + "")
                                            _ = v68
                                        } else {
                                            v3 = v3
                                            var v69 bool = (uint8(v36) < (v1 >> ((2) & 63)))
                                            _ = v69
                                        }
                                        var v70 int64 = 402
                                        if false {
                                            switch -1 {
                                            case -759:
                                                var v71 float64 = float64(0.0)
                                                _ = v71
                                            default:
                                                v70 = 9223372036854775807
                                            }
                                        }
                                        _ = v70
                                    }
                                    var v72 int16 = 42
                                    _ = v72
                                }
                                if false {
                                    _ = len(v5)
                                    var v73 uint32 = 42
                                    v74 := GenS3{F0: v37, F1: -1, F2: v2}
                                    _ = v73
                                    _ = v74
                                }
                                fmt.Println(v2, v3)
                            } else {
                                var v75 float32 = v10
                                if (!false) {
                                    switch "\n\t" {
                                    case "":
                                        var v76 float64 = v3
                                        _ = v76
                                    case "foo bar":
                                        if (!true) {
                                            var v77 string = ("hello" + v59)
                                            v78 := GenS3{F0: 227, F1: (^(v37)), F2: v2}
                                            _ = v77
                                            _ = v78
                                        }
                                    case "hello":
                                        if (true || true) {
                                            v79 := GenS2{F0: v37}
                                            v80 := v5["foo bar"]
                                            switch 42 {
                                            case 9223372036854775807:
                                                var v81 uint8 = v52
                                                _ = v81
                                            case 42:
                                                delete(v5, "")
                                            case 653:
                                                fmt.Println(v10)
                                            }
                                            var v82 uint = 65535
                                            fmt.Println(fi51, v82, v4)
                                            _ = v79
                                            _ = v80
                                            _ = v82
                                        }
                                    }
                                    if true {
                                        v83 := []int{}
                                        var v84 int32 = (v37 + (v37 | (v37 | v37)))
                                        v84 = v84
                                        _ = v83
                                        _ = v84
                                    }
                                } else {
                                    var v85 int64 = -1
                                    var v86 float32 = float32(v3)
                                    v87 := GenS4{F0: 272, F1: (v37 - 14)}
                                    _ = v85
                                    _ = v86
                                    _ = v87
                                }
                                _ = v75
                            }
                            v2 -= (v2 & v2)
                            _ = v59
                        }
                    case "\n\t":
                        v88 := []int{(fi58 ^ (fi58 | fi51)), (fi51 - (fi51 - (fi51 / (((fi58 / ((127) | 1))) | 1))))}
                        _ = v88
                    case "hello":
                        fmt.Println(v52)
                    default:
                        fmt.Println(fi51, v2)
                    }
                    var v89 uint = 42
                    v90 := GenS1{F0: 41}
                    v91 := []int{(fi58 | fi58)}
                    switch (fi58 ^ fi58) {
                    case -876:
                        var v92 uint32 = uint32(v1)
                        _ = v92
                    case 413:
                        var v93 int = fi58
                        if true {
                            var v94 uint8 = v1
                            v52 = 127
                            v1 = v52
                            _ = len(v91)
                            _ = v94
                        } else {
                            var v95 int32 = (v37 - (v37 * 0))
                            _ = v95
                        }
                        _ = v93
                    }
                    v96 := GenS3{F0: int32(v52), F1: (^(v37)), F2: (v2 + (v2 * (v2 + v2)))}
                    _ = len(v5)
                    _ = fi58
                    _ = v89
                    _ = v90
                    _ = v91
                    _ = v96
                }
                fmt.Println(v37, v52)
                v37 += v37
            }
        case -167:
            fmt.Println(v10, v36, v1)
        case 2147483647:
            var v97 int64 = -479
            _ = v97
        }
        _ = fi51
        _ = v52
    }
    switch "" {
    case "a":
        if ((!(1 < (^(v36)))) && (0 == -533)) {
            for fi98 := 0; fi98 < 5; fi98++ {
                v2 = (^(v2))
                _ = fi98
            }
            for fi99 := 0; fi99 < 7; fi99++ {
                v100 := GenS2{F0: (v37 * 471)}
                fmt.Println(v3, v37, v10)
                _ = fi99
                _ = v100
            }
            if (1 >= uint(v1)) {
                v101 := []int{127, 625}
                _ = v101
            }
        } else {
            var v102 int8 = -113
            _ = v102
        }
    case "foo bar":
        v36 -= v36
    case "\n\t":
        v103 := []int{-9223372036854775808}
        _ = v103
    }
    var v104 float32 = v10
    if true {
        if (42 == -9223372036854775808) {
            v104 -= (0.1 + float32(v3))
            v105 := []int{}
            if ((v37 >> ((65535) & 63)) > v37) {
                v106 := v5["foo bar"]
                var v107 int32 = 1
                v4 = uint32(v2)
                _ = v106
                _ = v107
            }
            _ = v105
        } else {
            v108 := GenS1{F0: -118}
            if ((0.0 - v3) > v3) {
                switch int(v36) {
                case 2147483647:
                    v36 += (^(v36))
                case 128:
                    v104 = v104
                case -1:
                    v2 -= v2
                }
                var v109 string = "foo bar"
                _ = v109
            }
            v1 = 127
            _ = v108
        }
        v3 = 0.1
    }
    v110 := v5["a"]
    switch (v110 + (v110 + "a")) {
    case "foo bar":
        fmt.Println(v110, v104, v104)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v104:", v104)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v104
    _ = v110
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
