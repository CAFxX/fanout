package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 uint64; F1 uint32; F2 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64 }

func main() {
    var v1 int32 = 1
    var v2 string = ""
    v3 := []int{-524}
    v4 := make(map[int]string)
    var v5 uint8 = 21
    var v6 float32 = 0.1
    var v7 float64 = (0.0 + 0.1)
    if len(v3) > 0 {
        v3[0] = -70
    }
    v8 := GenS1{F0: 4294967295, F1: 42, F2: v2}
    fmt.Println(v1)
    var v9 float64 = v7
    v10 := GenS2{F0: 0}
    _ = len(v4)
    switch "a" {
    case "foo bar":
        switch 42 {
        case 9223372036854775807:
            v4[740] = ("" + ("\n\t" + ("foo bar" + (v2 + "foo bar"))))
        case -152:
            v1 -= (v1 >> ((42) & 63))
        default:
            switch ((("a" + "a") + v2) + ("a" + v2)) {
            case "a":
                switch v2 {
                case "\n\t":
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v12 := GenS3{F0: true}
                        _ = fi11
                        _ = v12
                    }
                }
            default:
                v7 += 0.1
            }
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (2 <= 686) {
        if false {
            var v13 int16 = 0
            v14 := GenS2{F0: 32767}
            var v15 int64 = -108
            _ = v13
            _ = v14
            _ = v15
        }
        switch 127 {
        case 0:
            var v16 int16 = 32767
            _ = v16
        case -41:
            if (v1 >= v1) {
                var v17 uint64 = 18446744073709551615
                fmt.Println(v17, v1, v5)
                if false {
                    for fi18 := 0; fi18 < 12; fi18++ {
                        for fi19 := 0; fi19 < 5; fi19++ {
                            switch v2 {
                            case "":
                                var v20 int16 = -619
                                _ = v20
                            case "\n\t":
                                var v21 uint64 = 42
                                _ = v21
                            case "foo bar":
                                if ((!true) && (!false)) {
                                    var v22 float32 = float32(1e10)
                                    var v23 uint64 = (v17 << ((448) & 63))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v22
                                    _ = v23
                                } else {
                                    if false {
                                        switch "hello" {
                                        case "hello":
                                            if (true || (35 == 127)) {
                                                fmt.Println(v1)
                                            } else {
                                                _ = len(v3)
                                            }
                                        default:
                                            switch -611 {
                                            case 32767:
                                                var v24 uint32 = 65535
                                                _ = v24
                                            case 4294967295:
                                                if (v1 <= 2147483647) {
                                                    v7 -= v7
                                                    switch fi19 {
                                                    case -2147483648:
                                                        var v25 float64 = (v9 + ((0.0 / (v7 / (v7 - v9))) + 0.0))
                                                        _ = v25
                                                    default:
                                                        var v26 int8 = int8(fi19)
                                                        _ = v26
                                                    }
                                                    if (true && (true && (16 == 65535))) {
                                                        if ((1.5 * (1.5 + v6)) <= (v6 * (-2.25 / 3.14159))) {
                                                            if len(v3) > 0 {
                                                                v3[0] = fi18
                                                            }
                                                        } else {
                                                            v27 := GenS4{F0: (v17 % ((2) | 1)), F1: -128}
                                                            _ = v27
                                                        }
                                                        v1 = -1
                                                    } else {
                                                        switch v2 {
                                                        case "foo bar":
                                                            switch 618 {
                                                            case -1:
                                                                if true {
                                                                    var v28 float32 = (3.14159 / (1e10 - v6))
                                                                    _ = v28
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v29 := GenS5{F0: -103}
                                                            _ = v29
                                                        }
                                                        v30 := GenS5{F0: 42}
                                                        _ = v30
                                                    }
                                                    v31 := GenS2{F0: 592}
                                                    switch fi19 {
                                                    case -1:
                                                        switch (fi19 ^ -1) {
                                                        case 106:
                                                            v5 -= v5
                                                        case 42:
                                                            v1 += v1
                                                        }
                                                    case 874:
                                                        var v32 uint = 2
                                                        var v33 float32 = 0.0
                                                        v34 := GenS1{F0: v32, F1: -1, F2: v2}
                                                        var v35 int16 = 0
                                                        _ = v32
                                                        _ = v33
                                                        _ = v34
                                                        _ = v35
                                                    case 0:
                                                        if false {
                                                            v1 = -1
                                                            _ = len(v3)
                                                        }
                                                    }
                                                    v4[fi18] = "foo bar"
                                                    _ = v31
                                                } else {
                                                    v2 = (v2 + ("\n\t" + (((v2 + "\n\t") + (v2 + "\n\t")) + "\n\t")))
                                                    switch v2 {
                                                    case "\n\t":
                                                        switch fi18 {
                                                        case 32767:
                                                            v2 = v2
                                                        case 428:
                                                            var v36 uint32 = 2147483648
                                                            _ = v36
                                                        case 9223372036854775807:
                                                            var v37 uint16 = 32767
                                                            _ = v37
                                                        default:
                                                            v9 += (1e10 - v7)
                                                        }
                                                    case "a":
                                                        var v38 float32 = ((((3.14159 / v6) / (float32((((1e10 / v7) - (-2.25 / ((1e10 - -2.25) / v9))) * (3.14159 + v9))) * (1e10 - (((v6 - -2.25) - (v6 - ((((v6 * (1e10 - (((float32(-2.25) + 1.5) - 3.14159) + -2.25))) / v6) + (v6 + (1e10 - float32((-2.25 - v9))))) + v6))) * (v6 / v6))))) - v6) / 1e10)
                                                        _ = v38
                                                    case "hello":
                                                        var v39 int8 = -128
                                                        _ = v39
                                                    default:
                                                        switch (v2 + v2) {
                                                        case "a":
                                                            v17 -= 2
                                                        case "foo bar":
                                                            v5 = (v5 - v5)
                                                        case "":
                                                            var v40 uint = uint(v5)
                                                            _ = v40
                                                        }
                                                        _ = len(v4)
                                                    }
                                                }
                                                if ((fi18 < 446) || true) {
                                                    fmt.Println(v1, fi19)
                                                } else {
                                                    v41 := GenS6{F0: 1, F1: 127}
                                                    _ = v41
                                                }
                                            case 42:
                                                v9 = (v7 * float64((0.1 + v6)))
                                            default:
                                                if (true || (!true)) {
                                                    v5 -= (v5 - (v5 << ((619) & 63)))
                                                }
                                            }
                                        }
                                    } else {
                                        var v42 uint16 = 255
                                        _ = len(v4)
                                        if false {
                                            switch (v2 + "\n\t") {
                                            case "foo bar":
                                                fmt.Println(v2, v1)
                                            case "":
                                                var v43 uint8 = v5
                                                _ = v43
                                            case "hello":
                                                v6 = v6
                                            default:
                                                v44 := GenS5{F0: 1000}
                                                _ = v44
                                            }
                                            if true {
                                                v7 = v9
                                                v45 := GenS7{F0: (v17 - 0), F1: 0, F2: true}
                                                var v46 int16 = 1000
                                                _ = v45
                                                _ = v46
                                            }
                                            if ((3.14159 < v7) && true) {
                                                v47 := v4[-99]
                                                var v48 int16 = -32768
                                                if (false && ((v9 - v7) != (-2.25 * float64(v6)))) {
                                                    var v49 uint = 255
                                                    delete(v4, fi18)
                                                    fmt.Println(v42, v5)
                                                    v50 := GenS8{F0: (v2 + (v47 + v47))}
                                                    switch (^(fi19)) {
                                                    case -2147483648:
                                                        var v51 uint16 = (^(v42))
                                                        _ = v51
                                                    case 770:
                                                        if len(v3) > 0 {
                                                            v3[0] = 4294967295
                                                        }
                                                        fmt.Println(v5)
                                                        fmt.Println(v17, fi19, v9)
                                                    }
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    if false {
                                                        v52 := GenS2{F0: v42}
                                                        var v53 float32 = v6
                                                        switch "a" {
                                                        case "foo bar":
                                                            fmt.Println(v9, v2, v5)
                                                        default:
                                                            var v54 uint32 = 2147483647
                                                            _ = v54
                                                        }
                                                        v48 -= v48
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        v55 := GenS6{F0: (v42 % ((v42) | 1)), F1: fi18}
                                                        v47 = (v47 + (((v47 + (v47 + (("" + v47) + "a"))) + "\n\t") + (v47 + v47)))
                                                        v5 = (^(v5))
                                                        _ = v55
                                                    }
                                                    v56 := GenS1{F0: 683, F1: -271, F2: ""}
                                                    if true {
                                                        if (fi18 == (fi19 - fi19)) {
                                                            if false {
                                                                v42 = 255
                                                                var v57 int8 = 0
                                                                fmt.Println(v42)
                                                                _ = v57
                                                            }
                                                            fmt.Println(v2, fi18)
                                                            var v58 int = fi19
                                                            _ = v58
                                                        } else {
                                                            var v59 int64 = 42
                                                            v60 := GenS9{F0: (v59 << ((4294967295) & 63))}
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                    _ = v56
                                                }
                                                var v61 uint = 875
                                                _ = v47
                                                _ = v48
                                                _ = v61
                                            }
                                            switch fi18 {
                                            case 2147483647:
                                                v4[-2147483648] = v2
                                                switch fi19 {
                                                case 32768:
                                                    if true {
                                                        v7 += 0.0
                                                        fmt.Println(v7, v1, fi19)
                                                    }
                                                case -881:
                                                    var v62 string = ""
                                                    _ = v62
                                                default:
                                                    v63 := GenS6{F0: (v42 - uint16(v17)), F1: 9223372036854775807}
                                                    _ = v63
                                                }
                                                fmt.Println(v6, v6)
                                            default:
                                                var v64 int32 = v1
                                                _ = v64
                                            }
                                            v7 = (0.1 - ((v7 * (float64((0.1 / (((v6 + -2.25) - float32(v7)) * 1.5))) + (v7 + 0.1))) - 0.0))
                                            v9 = (v7 / 0.1)
                                        } else {
                                            var v65 uint32 = 2
                                            switch fi18 {
                                            case 127:
                                                var v66 string = v2
                                                _ = v66
                                            case 32768:
                                                if true {
                                                    var v67 int16 = int16(v17)
                                                    if len(v3) > 0 {
                                                        v3[0] = -405
                                                    }
                                                    var v68 bool = false
                                                    _ = v67
                                                    _ = v68
                                                }
                                            case -9223372036854775808:
                                                switch fi19 {
                                                case 743:
                                                    v1 += v1
                                                default:
                                                    switch (fi18 & fi18) {
                                                    case 0:
                                                        fmt.Println(v5, fi19)
                                                    case -9223372036854775808:
                                                        if len(v3) > 0 {
                                                            v3[0] = -9223372036854775808
                                                        }
                                                    default:
                                                        var v69 uint32 = v65
                                                        _ = v69
                                                    }
                                                }
                                            }
                                            v9 -= (float64((v6 * float32(float64(v6)))) - float64((v6 * (v6 * 1e10))))
                                            _ = v65
                                        }
                                        var v70 int8 = -128
                                        var v71 uint64 = v17
                                        _ = v42
                                        _ = v70
                                        _ = v71
                                    }
                                }
                            }
                            var v72 uint16 = 981
                            _ = fi19
                            _ = v72
                        }
                        v7 = float64(0.0)
                        if len(v3) > 0 {
                            v3[0] = fi18
                        }
                        _ = fi18
                    }
                    var v73 int64 = -1
                    _ = v73
                }
                _ = v17
            } else {
                var v74 int8 = 100
                if (((0 != 284) || false) && false) {
                    var v75 int = 128
                    var v76 float64 = v7
                    _ = v75
                    _ = v76
                }
                _ = v74
            }
        case 32768:
            for fi77 := 0; fi77 < 2; fi77++ {
                var v78 uint8 = (v5 >> ((255) & 63))
                _ = len(v4)
                for fi79 := 0; fi79 < 11; fi79++ {
                    v3 = append(v3, fi77)
                    _ = fi79
                }
                _ = fi77
                _ = v78
            }
        default:
            v2 = "\n\t"
        }
    }
    if false {
        var v80 int8 = -1
        v80 -= (v80 - int8(v1))
        v2 = (v2 + v2)
        _ = v80
    }
    v9 += 1e10
    if false {
        switch 2147483647 {
        case 32768:
            if (-9223372036854775808 != 638) {
                v81 := GenS2{F0: 462}
                for fi82 := 0; fi82 < 8; fi82++ {
                    fmt.Println(v9)
                    _ = fi82
                }
                _ = v81
            }
        case 389:
            switch v2 {
            case "a":
                _ = len(v3)
            default:
                if (((9223372036854775807 == 9223372036854775807) && (!(-2147483648 == v1))) && (-488 > -760)) {
                    var v83 string = ""
                    _ = v83
                }
            }
        case -446:
            var v84 uint64 = 282
            _ = v84
        }
    }
    for fi85 := 0; fi85 < 1; fi85++ {
        if ((v1 / (((^(v1))) | 1)) == -973) {
            var v86 string = "\n\t"
            _ = v86
        }
        _ = fi85
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
