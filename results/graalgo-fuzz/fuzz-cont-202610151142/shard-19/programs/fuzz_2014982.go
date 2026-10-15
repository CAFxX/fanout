package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int32; F1 int; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint64; F1 int; F2 bool }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int }

func main() {
    var v1 bool = (1.5 != 3.14159)
    var v2 int8 = 50
    var v3 uint16 = 2
    if true {
        if (368 < 9223372036854775807) {
            fmt.Println(v3)
        } else {
            v4 := []int{-708}
            fmt.Println(v1, v2, v3)
            switch 128 {
            case -9223372036854775808:
                v5 := GenS1{F0: 396}
                _ = v5
            default:
                v4 = append(v4, 2147483647)
            }
            _ = v4
        }
        v2 += v2
        v6 := GenS2{F0: "hello", F1: 42}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 6; fi7++ {
            v8 := []int{257}
            var v9 float32 = 1.5
            _ = fi7
            _ = v8
            _ = v9
        }
        v2 += (v2 | 100)
        for fi10 := 0; fi10 < 16; fi10++ {
            v3 = v3
            var v11 int8 = v2
            _ = fi10
            _ = v11
        }
    }
    var v12 int = 170
    for fi13 := 0; fi13 < 15; fi13++ {
        v2 += v2
        if (!v1) {
            if (fi13 == fi13) {
                v12 = v12
                var v14 uint64 = 1
                v1 = (v1 || v1)
                _ = v14
            } else {
                for fi15 := 0; fi15 < 16; fi15++ {
                    var v16 int16 = -1
                    switch "foo bar" {
                    case "foo bar":
                        if v1 {
                            var v17 float64 = 1e10
                            v16 = (v16 + int16(fi13))
                            var v18 int16 = (v16 * (v16 ^ 0))
                            _ = v17
                            _ = v18
                        } else {
                            var v19 int8 = v2
                            _ = v19
                        }
                    case "a":
                        fmt.Println(v3, v2)
                    }
                    _ = fi15
                    _ = v16
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    v21 := []int{fi20, (fi20 + (^(fi13))), fi13}
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    if v1 {
                        var v22 int16 = 32767
                        _ = v22
                    } else {
                        if true {
                            v21 = append(v21, 128)
                            fmt.Println(fi20, fi13, v2)
                            _ = len(v21)
                        } else {
                            v12 += (^(fi20))
                            fmt.Println(v12, v2)
                            switch 2147483647 {
                            case 128:
                                fmt.Println(fi20)
                            default:
                                var v23 int16 = 200
                                _ = v23
                            }
                        }
                        _ = len(v21)
                        fmt.Println(v12, v2, v1)
                        v1 = ((-(v2)) < (v2 * v2))
                    }
                    v24 := GenS1{F0: 0}
                    if false {
                        v1 = (-32768 == 1000)
                        if ((-(v2)) != 25) {
                            switch fi13 {
                            case 9223372036854775807:
                                var v25 int16 = 1
                                _ = v25
                            case 32767:
                                var v26 uint16 = 1
                                _ = v26
                            }
                        } else {
                            var v27 uint64 = 697
                            v27 = 302
                            v2 -= (v2 | (v2 * -1))
                            switch fi20 {
                            case 4294967295:
                                if false {
                                    if true {
                                        var v28 uint32 = 1
                                        var v29 int = fi13
                                        if ((1.5 - 0.0) == 1e10) {
                                            v30 := GenS3{F0: 1, F1: v29, F2: (!true)}
                                            var v31 uint = 875
                                            v29 = v29
                                            _ = v30
                                            _ = v31
                                        }
                                        v32 := GenS4{F0: ""}
                                        _ = v28
                                        _ = v29
                                        _ = v32
                                    } else {
                                        v3 = (v3 | v3)
                                        v27 -= v27
                                    }
                                    v33 := GenS5{F0: int16(v3)}
                                    v34 := GenS3{F0: 42, F1: -1, F2: false}
                                    var v35 string = "a"
                                    var v36 bool = (0.1 == float32(0.1))
                                    v3 = 2
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                            }
                            _ = v27
                        }
                    }
                    if true {
                        if (true && (2147483647 < 2)) {
                            var v37 float64 = 0.0
                            fmt.Println(v12, v3, fi13)
                            switch (fi13 % ((-1) | 1)) {
                            case 0:
                                if (true || true) {
                                    switch "a" {
                                    case "":
                                        var v38 uint8 = 128
                                        _ = v38
                                    }
                                    v39 := GenS6{F0: -9223372036854775808}
                                    fmt.Println(v12, fi20, v3)
                                    _ = v39
                                } else {
                                    switch 32768 {
                                    case 4294967295:
                                        if v1 {
                                            switch (fi13 & 18) {
                                            case -844:
                                                if v1 {
                                                    v40 := GenS7{F0: (-(v12)), F1: v3}
                                                    if len(v21) >= 2 {
                                                        v21 = v21[1:len(v21)]
                                                    }
                                                    _ = v40
                                                } else {
                                                    v41 := GenS7{F0: (v12 / (((v12 + fi20)) | 1)), F1: (v3 ^ (^(v3)))}
                                                    _ = v41
                                                }
                                            }
                                            var v42 string = ""
                                            _ = v42
                                        }
                                    case 127:
                                        switch (fi20 & v12) {
                                        case 1:
                                            v3 = (v3 & (v3 | v3))
                                        }
                                    case -9223372036854775808:
                                        v21 = append(v21, (fi13 * (fi13 - (^(fi20)))))
                                    default:
                                        var v43 uint8 = 0
                                        _ = v43
                                    }
                                    v44 := GenS8{F0: 42, F1: v12, F2: ((v1 && v1) || (v1 || (759 != 418)))}
                                    if v1 {
                                        var v45 uint64 = 360
                                        v46 := GenS9{F0: (v45 | v45)}
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = v44
                                }
                            case 9223372036854775807:
                                if (-32768 <= -650) {
                                    v37 = 1.5
                                    var v47 int8 = v2
                                    _ = v47
                                } else {
                                    var v48 float32 = (1.5 - ((1.5 - 1.5) + 1.5))
                                    v37 = ((((float64(v48) / 1.5) * (0.0 * v37)) / v37) * -2.25)
                                    _ = v48
                                }
                            case -2147483648:
                                v21 = append(v21, (fi20 % ((fi13) | 1)))
                            }
                            _ = v37
                        } else {
                            v1 = v1
                            var v49 uint64 = 0
                            var v50 int32 = 42
                            _ = v49
                            _ = v50
                        }
                        var v51 bool = (float64(-2.25) != 1.5)
                        _ = v51
                    }
                    _ = fi20
                    _ = v21
                    _ = v24
                }
                v1 = v1
            }
        } else {
            if true {
                v2 = v2
            }
            v52 := GenS7{F0: v12, F1: v3}
            _ = v52
        }
        _ = fi13
    }
    v53 := GenS8{F0: 0, F1: v12, F2: (!(false || (!((!(!(v1 || ((false || true) && v1)))) && v1))))}
    switch v12 {
    case -2147483648:
        for fi54 := 0; fi54 < 15; fi54++ {
            fmt.Println(v2, v2)
            var v55 int64 = 42
            _ = fi54
            _ = v55
        }
    case 742:
        for fi56 := 0; fi56 < 13; fi56++ {
            v57 := GenS10{F0: v12, F1: 624}
            var v58 int8 = -128
            v12 += (v12 | -9223372036854775808)
            _ = fi56
            _ = v57
            _ = v58
        }
    case 0:
        var v59 uint8 = 127
        _ = v59
    default:
        if ((v3 * (^(v3))) >= (v3 + v3)) {
            v12 = 303
            if (184 != 2147483648) {
                v3 += (v3 / ((42) | 1))
            }
            v60 := []int{}
            _ = v60
        } else {
            if v1 {
                v12 -= 656
                for fi61 := 0; fi61 < 14; fi61++ {
                    switch "" {
                    case "hello":
                        var v62 int32 = -259
                        _ = v62
                    case "":
                        v63 := GenS6{F0: 42}
                        _ = v63
                    case "a":
                        for fi64 := 0; fi64 < 11; fi64++ {
                            if false {
                                var v65 string = "\n\t"
                                _ = v65
                            } else {
                                fmt.Println(v3, v3)
                            }
                            if ((false || v1) || (-166 >= 32767)) {
                                if (-2147483648 > -1) {
                                    switch fi61 {
                                    case 2147483647:
                                        var v66 int16 = -32768
                                        if v1 {
                                            v3 += (v3 ^ 65535)
                                            var v67 bool = v1
                                            _ = v67
                                        }
                                        var v68 uint = 65535
                                        _ = v66
                                        _ = v68
                                    case -9223372036854775808:
                                        var v69 uint = 42
                                        _ = v69
                                    }
                                    v70 := []int{fi64}
                                    _ = v70
                                }
                            }
                            fmt.Println(v12)
                            _ = fi64
                        }
                    default:
                        v1 = ((128 > 42) || (713 == 2))
                    }
                    if (127 == 1) {
                        v71 := GenS8{F0: 1, F1: 128, F2: v1}
                        var v72 int16 = -1
                        v73 := []int{-814, v12}
                        _ = v71
                        _ = v72
                        _ = v73
                    } else {
                        var v74 bool = ((1 >= 65535) || false)
                        for fi75 := 0; fi75 < 17; fi75++ {
                            v74 = v1
                            _ = fi75
                        }
                        _ = v74
                    }
                    var v76 float64 = 1e10
                    _ = fi61
                    _ = v76
                }
                fmt.Println(v1)
            }
        }
    }
    for fi77 := 0; fi77 < 18; fi77++ {
        switch (fi77 ^ (v12 / (((v12 << ((42) & 63))) | 1))) {
        case 42:
            v78 := GenS11{F0: (!v1)}
            _ = v78
        case 4294967295:
            fmt.Println(v1, v2, fi77)
        case -640:
            var v79 float32 = 1.5
            _ = v79
        }
        for fi80 := 0; fi80 < 10; fi80++ {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v3, v3)
            case "\n\t":
                v3 = (^(v3))
            case "":
                v81 := []int{(fi80 >> ((4294967295) & 63)), 538, fi77, fi77}
                _ = v81
            }
            var v82 uint8 = 0
            _ = fi80
            _ = v82
        }
        v83 := GenS12{F0: 42}
        _ = fi77
        _ = v83
    }
    v2 -= v2
    switch ("hello" + "a") {
    case "foo bar":
        var v84 int16 = 1
        _ = v84
    case "\n\t":
        v2 -= 1
    default:
        v85 := GenS11{F0: (!v1)}
        _ = v85
    }
    switch v12 {
    case -167:
        switch "hello" {
        case "hello":
            switch v12 {
            case -200:
                switch "a" {
                case "\n\t":
                    v12 += v12
                case "":
                    for fi86 := 0; fi86 < 13; fi86++ {
                        v87 := GenS5{F0: -907}
                        if ((2147483647 == 1) || v1) {
                            v1 = false
                        } else {
                            v88 := GenS1{F0: uint(v3)}
                            _ = v88
                        }
                        var v89 int = (v12 - fi86)
                        _ = fi86
                        _ = v87
                        _ = v89
                    }
                default:
                    switch ("a" + "") {
                    case "\n\t":
                        if (127 == v2) {
                            v2 = v2
                        } else {
                            v3 = (v3 / ((v3) | 1))
                            var v90 string = (("" + "foo bar") + "")
                            fmt.Println(v12, v3)
                            _ = v90
                        }
                    }
                }
            default:
                v91 := []int{int(v2), (v12 << ((18446744073709551615) & 63)), -187, 32767}
                _ = v91
            }
        case "foo bar":
            fmt.Println(v3)
        }
    default:
        fmt.Println(v2)
    }
    v1 = v1
    v2 = int8(v12)
    v3 = v3
    for fi92 := 0; fi92 < 18; fi92++ {
        v93 := GenS1{F0: 682}
        _ = fi92
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v53
}
