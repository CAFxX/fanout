package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 bool; F1 int64; F2 int64 }
type GenS6 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint16; F1 int; F2 uint }
func genFn0() int {
    var v5 uint = 4294967295
    fmt.Println(v5)
    v5 = 999
    var v6 int32 = -972
    v7 := GenS1{F0: 4294967295}
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []int{2147483647}
    v2 := make(map[uint32]int)
    var v3 uint = 2
    var v4 uint = v3
    v4 = (v4 >> (((^(v3))) & 63))
    fmt.Println(v4, v4)
    if (42 < 65535) {
        v3 += 65535
        var v8 float64 = 1.5
        _ = len(v2)
        _ = v8
    } else {
        v9 := GenS1{F0: 2}
        var v10 uint16 = 0
        _ = v10
        _ = v9
    }
    v3 = v3
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 uint8 = 255
        for fi13 := 0; fi13 < 4; fi13++ {
            switch (fi11 >> ((0) & 63)) {
            case 32768:
                var v14 uint64 = 777
                _ = v14
            case 0:
                fmt.Println(fi11, fi11)
            case 42:
                _ = len(v2)
                var v15 float32 = 1e10
                v15 = 3.14159
                _ = v15
            }
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v3 = v4
        _ = fi16
    }
    v3 += v3
    fmt.Println(v4, v3)
    v4 -= v4
    var v17 uint64 = 18446744073709551615
    if false {
        for fi18 := 0; fi18 < 16; fi18++ {
            if (true && (true || (!(uint32(v4) == 192)))) {
                _ = len(v1)
                var v19 int = (fi18 / ((genFn0()) | 1))
                if (42 > 127) {
                    switch ("\n\t" + ((("" + "hello") + "foo bar") + ((("hello" + "\n\t") + "\n\t") + "foo bar"))) {
                    case "":
                        fmt.Println(v17, v19)
                    case "\n\t":
                        var v20 uint32 = 254
                        _ = v20
                    default:
                        if (!false) {
                            switch (v19 << (((v4 | v3)) & 63)) {
                            case -628:
                                if (582 == -1) {
                                    v19 -= (v19 | -1)
                                }
                            case -1:
                                v21 := v2[uint32(v17)]
                                _ = v21
                            }
                            v22 := GenS1{F0: (^(v17))}
                            switch (("\n\t" + "") + "a") {
                            case "a":
                                v17 -= 1
                            case "foo bar":
                                var v23 int = v19
                                _ = v23
                            }
                            _ = v22
                        }
                    }
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v3 = v3
                        if (42 > -2147483648) {
                            if true {
                                v4 = (v3 / (((v4 % ((v3) | 1))) | 1))
                                if len(v1) > 0 {
                                    v1[0] = 538
                                }
                            }
                        }
                        if (42 >= 0) {
                            if true {
                                v17 = (v17 / ((v17) | 1))
                                if true {
                                    var v25 float32 = ((1e10 * (0.1 / -2.25)) * 1e10)
                                    _ = v25
                                } else {
                                    var v26 int64 = 692
                                    switch ("" + "\n\t") {
                                    case "\n\t":
                                        if (false && (false || (false || (!(-32768 != -846))))) {
                                            if ("foo bar" > "") {
                                                if false {
                                                    if len(v1) > 0 {
                                                        v1[0] = 4294967295
                                                    }
                                                    v27 := GenS2{F0: (((!true) && ((((("" + "foo bar") + "a") + "") + "a") > "hello")) && false)}
                                                    _ = v27
                                                }
                                            }
                                        }
                                    case "a":
                                        if ((v26 / ((int64(v3)) | 1)) < (^(v26))) {
                                            var v28 float32 = float32(0.1)
                                            var v29 int8 = 1
                                            v26 = v26
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 uint32 = 2
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v31 int32 = 1
                                            v4 = (v3 & (^(v3)))
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                    _ = v26
                                }
                                if (647 <= 0) {
                                    v32 := GenS3{F0: v19, F1: 65535}
                                    var v33 uint8 = 109
                                    _ = v32
                                    _ = v33
                                }
                            }
                            var v34 int64 = 0
                            var v35 uint16 = uint16(v3)
                            _ = v34
                            _ = v35
                        }
                        _ = fi24
                    }
                    fmt.Println(v4, fi18, v19)
                } else {
                    var v36 float64 = 0.1
                    v37 := GenS4{F0: 1, F1: 688}
                    v38 := GenS5{F0: false, F1: -1, F2: 1}
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = v19
            }
            if true {
                var v39 int16 = -1
                if true {
                    v40 := v2[2]
                    if (0.1 <= 0.0) {
                        v41 := GenS5{F0: false, F1: 0, F2: 42}
                        if (v4 < (^(v4))) {
                            if false {
                                if false {
                                    var v42 int = (^(fi18))
                                    v39 -= -1
                                    v43 := GenS6{F0: 0, F1: v17, F2: 1}
                                    _ = v42
                                    _ = v43
                                } else {
                                    var v44 uint = v3
                                    _ = v44
                                }
                                if (395 > -897) {
                                    v39 = (v39 ^ (v39 >> ((511) & 63)))
                                    _ = len(v1)
                                } else {
                                    v45 := GenS4{F0: 0, F1: v39}
                                    v39 = v39
                                    var v46 uint32 = 200
                                    _ = v45
                                    _ = v46
                                }
                                switch (v40 ^ (fi18 | v40)) {
                                case 128:
                                    fmt.Println(v40)
                                case -212:
                                    v3 += (v3 ^ v3)
                                case 4294967295:
                                    if true {
                                        fmt.Println(v39, v40, v3)
                                    } else {
                                        var v47 bool = false
                                        var v48 int32 = 42
                                        var v49 int = (-(v40))
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                default:
                                    var v50 int64 = 436
                                    _ = v50
                                }
                            }
                            if false {
                                if (1e10 <= (3.14159 * -2.25)) {
                                    var v51 int16 = 0
                                    var v52 uint64 = (v17 & 237)
                                    v53 := v2[2]
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                } else {
                                    v39 = (v39 << ((v4) & 63))
                                    var v54 uint8 = 255
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v54
                                }
                            }
                        }
                        for fi55 := 0; fi55 < 15; fi55++ {
                            switch ("" + "\n\t") {
                            case "\n\t":
                                if (!true) {
                                    switch genFn0() {
                                    case 4294967295:
                                        v40 += (fi55 + (v40 / ((int(v17)) | 1)))
                                    default:
                                        switch fi55 {
                                        case 127:
                                            delete(v2, 4294967295)
                                        case 32767:
                                            var v56 uint32 = 1
                                            _ = v56
                                        case 4294967295:
                                            v2[581] = (v40 - (fi18 >> ((v3) & 63)))
                                        }
                                    }
                                    v57 := GenS2{F0: true}
                                    v17 += (v17 ^ (^(v17)))
                                    var v58 int64 = -9223372036854775808
                                    _ = v57
                                    _ = v58
                                }
                            }
                            _ = fi55
                        }
                        _ = v41
                    }
                    if false {
                        var v59 int = (v40 >> (((v4 & (v4 & (v3 | v3)))) & 63))
                        _ = len(v1)
                        if true {
                            v59 = genFn0()
                        } else {
                            var v60 uint8 = 42
                            var v61 uint = v4
                            _ = v60
                            _ = v61
                        }
                        _ = v59
                    } else {
                        v62 := GenS2{F0: true}
                        _ = v62
                    }
                    _ = v40
                }
                for fi63 := 0; fi63 < 16; fi63++ {
                    switch "\n\t" {
                    case "":
                        var v64 uint = 2
                        _ = v64
                    case "foo bar":
                        if true {
                            if true {
                                v65 := GenS3{F0: fi63, F1: 42}
                                _ = v65
                            } else {
                                var v66 string = ("foo bar" + ("hello" + "hello"))
                                if true {
                                    fmt.Println(v39, v4, fi18)
                                    if false {
                                        v3 += v4
                                        delete(v2, 65535)
                                        v4 -= v3
                                    }
                                } else {
                                    v66 = "\n\t"
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if (!false) {
                                    if true {
                                        v67 := v2[115]
                                        _ = v67
                                    } else {
                                        var v68 float32 = (1.5 - (0.1 - -2.25))
                                        delete(v2, 42)
                                        v69 := GenS1{F0: v17}
                                        fmt.Println(v4, v4)
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 uint16 = 255
                                    _ = v70
                                }
                                _ = v66
                            }
                            if len(v1) > 0 {
                                v1[0] = (fi18 ^ fi18)
                            }
                            var v71 uint32 = uint32(v4)
                            _ = v71
                        } else {
                            v3 = 948
                        }
                    case "hello":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v17 = v17
                    var v72 bool = true
                    var v73 uint32 = 65535
                    var v74 bool = v72
                    _ = fi63
                    _ = v72
                    _ = v73
                    _ = v74
                }
                _ = v39
            } else {
                v3 = uint(v17)
                var v75 int16 = 1
                var v76 int16 = (-(v75))
                _ = v75
                _ = v76
            }
            switch "foo bar" {
            case "":
                if true {
                    if false {
                        v17 = (v17 & 2)
                        switch fi18 {
                        case 2147483647:
                            if (43 == 250) {
                                v77 := v2[2147483648]
                                _ = v77
                            }
                        case -9223372036854775808:
                            var v78 int32 = int32(v17)
                            _ = v78
                        }
                        if false {
                            var v79 uint32 = 613
                            _ = v79
                        } else {
                            v17 = (^(v17))
                            var v80 float32 = ((-2.25 * 1e10) + -2.25)
                            _ = v80
                        }
                    } else {
                        var v81 uint16 = 980
                        var v82 int64 = -1
                        _ = v81
                        _ = v82
                    }
                }
            case "foo bar":
                var v83 uint8 = 1
                _ = v83
            case "a":
                if false {
                    var v84 uint16 = 42
                    var v85 uint16 = 42
                    if (false || false) {
                        v4 += (v3 - (^(v4)))
                        switch fi18 {
                        case -172:
                            var v86 int8 = 97
                            _ = v86
                        case 1:
                            if ("foo bar" >= "foo bar") {
                                fmt.Println(v17)
                                var v87 uint64 = (v17 >> ((v3) & 63))
                                _ = v87
                            } else {
                                var v88 int8 = 127
                                v17 = v17
                                _ = v88
                            }
                        }
                        switch -985 {
                        case 593:
                            if len(v1) > 0 {
                                v1[0] = genFn0()
                            }
                        case 2147483647:
                            var v89 int16 = 433
                            _ = v89
                        case 127:
                            var v90 string = ("a" + ((("\n\t" + "hello") + "a") + (((("a" + "\n\t") + ("a" + "\n\t")) + ("" + "")) + "")))
                            _ = v90
                        default:
                            var v91 uint32 = 906
                            _ = v91
                        }
                    }
                    _ = v84
                    _ = v85
                } else {
                    switch "" {
                    case "":
                        var v92 uint8 = 2
                        _ = v92
                    case "hello":
                        for fi93 := 0; fi93 < 15; fi93++ {
                            v4 += (^(v4))
                            v94 := GenS7{F0: 42, F1: 4294967295, F2: 65535}
                            _ = fi93
                            _ = v94
                        }
                    }
                    fmt.Println(fi18, v3)
                }
            }
            _ = fi18
        }
        v17 = v17
        v3 = (^(v3))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
