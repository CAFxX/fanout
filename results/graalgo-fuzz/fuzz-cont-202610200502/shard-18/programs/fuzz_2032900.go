package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint; F2 int64 }
type GenS4 struct { F0 uint8; F1 string; F2 int }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 int16; F1 uint16; F2 int64 }
type GenS7 struct { F0 int32; F1 int16; F2 int }
type GenS8 struct { F0 int64; F1 uint32; F2 int64 }
type GenS9 struct { F0 uint64; F1 int; F2 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int64 }
func genFn0(p0 uint64) uint8 {
    v4 := GenS1{F0: -2147483648, F1: 42}
    v5 := GenS1{F0: -2147483648, F1: 42}
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1() int16 {
    if true {
        if (1e10 != 1e10) {
            v6 := []int{}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v6
        }
    } else {
        if true {
            var v7 int64 = 42
            _ = v7
        } else {
            switch "hello" {
            case "\n\t":
                if false {
                    if (-2147483648 != 42) {
                        var v8 uint16 = 154
                        v9 := []int{127, 2147483647}
                        v8 = (v8 * 65535)
                        _ = v8
                        _ = v9
                    }
                    v10 := []int{1}
                    var v11 int8 = 127
                    _ = v10
                    _ = v11
                } else {
                    v12 := GenS1{F0: 42, F1: 2}
                    for fi13 := 0; fi13 < 15; fi13++ {
                        for fi14 := 0; fi14 < 6; fi14++ {
                            var v15 string = "foo bar"
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 9; fi16++ {
                            fmt.Println(fi16)
                            if false {
                                if true {
                                    fmt.Println(fi13)
                                    v17 := GenS1{F0: -632, F1: 42}
                                    v18 := []int{}
                                    _ = v17
                                    _ = v18
                                }
                                switch (-(fi13)) {
                                case 32767:
                                    var v19 int = fi13
                                    _ = v19
                                }
                                v20 := GenS2{F0: 0}
                                _ = v20
                            }
                            fmt.Println(fi13, fi16)
                            _ = fi16
                        }
                        for fi21 := 0; fi21 < 11; fi21++ {
                            var v22 uint64 = 7
                            fmt.Println(v22, v22, fi21)
                            _ = fi21
                            _ = v22
                        }
                        _ = fi13
                    }
                    v23 := GenS3{F0: 688, F1: 865, F2: -975}
                    _ = v12
                    _ = v23
                }
            case "a":
                v24 := []int{2147483647, -2147483648, 32767}
                _ = v24
            default:
                v25 := []int{42, -9223372036854775808}
                _ = v25
            }
        }
        var v26 int8 = -91
        var v27 uint16 = 0
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        switch fi28 {
        case -2147483648:
            for fi29 := 0; fi29 < 13; fi29++ {
                var v30 float64 = ((0.1 / (1e10 + (0.0 / ((float64(((0.0 / float32(float64(((1.5 + 0.0) + 1.5)))) - (float32(3.14159) - 0.0))) + float64(1e10)) + ((0.0 - 1e10) + 1e10))))) / 0.1)
                var v31 int8 = int8(fi29)
                var v32 uint32 = 42
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        case 4294967295:
            var v33 float64 = 3.14159
            _ = v33
        default:
            v34 := GenS4{F0: genFn0(2), F1: "hello", F2: fi28}
            _ = v34
        }
        if false {
            var v35 string = "a"
            _ = v35
        } else {
            v36 := GenS1{F0: 78, F1: 2}
            switch (("foo bar" + "a") + "hello") {
            case "hello":
                var v37 uint64 = 4294967295
                _ = v37
            case "foo bar":
                var v38 int16 = 1
                _ = v38
            case "\n\t":
                if true {
                    switch ("\n\t" + "\n\t") {
                    case "a":
                        var v39 int64 = -9223372036854775808
                        _ = v39
                    }
                    switch (fi28 & fi28) {
                    case -1:
                        switch fi28 {
                        case -1:
                            if false {
                                fmt.Println(fi28)
                                v40 := []int{-2147483648, (^(fi28)), (^(fi28))}
                                var v41 uint8 = 95
                                _ = v40
                                _ = v41
                            }
                        }
                    case 9223372036854775807:
                        v42 := []int{-2147483648, fi28}
                        _ = v42
                    case -9223372036854775808:
                        var v43 uint32 = 2147483647
                        _ = v43
                    }
                    var v44 float64 = 3.14159
                    _ = v44
                }
            default:
                var v45 bool = false
                _ = v45
            }
            var v46 uint64 = 101
            _ = v36
            _ = v46
        }
        _ = fi28
    }
    switch "a" {
    case "":
        var v47 int64 = -1
        _ = v47
    case "hello":
        var v48 int16 = 965
        _ = v48
    case "a":
        if (1 == 255) {
            var v49 float32 = (float32(1e10) - (1e10 - 1e10))
            if false {
                switch "a" {
                case "\n\t":
                    switch "hello" {
                    case "\n\t":
                        v50 := GenS5{F0: 298, F1: -697}
                        _ = v50
                    case "a":
                        v51 := GenS6{F0: 42, F1: 1, F2: 1}
                        _ = v51
                    default:
                        if (((v49 / v49) + (v49 + v49)) == v49) {
                            v52 := []int{}
                            if ((445 != 263) || (false && (1 >= 0))) {
                                v53 := GenS2{F0: 2}
                                _ = v53
                            }
                            v52 = append(v52, 677)
                            _ = v52
                        } else {
                            v49 += v49
                        }
                    }
                case "a":
                    if true {
                        var v54 int = -9223372036854775808
                        _ = v54
                    } else {
                        v55 := GenS7{F0: -1, F1: 941, F2: 32767}
                        v56 := []int{}
                        _ = v55
                        _ = v56
                    }
                default:
                    v57 := []int{0, -54, -9223372036854775808}
                    _ = v57
                }
                fmt.Println(v49)
            } else {
                switch "\n\t" {
                case "foo bar":
                    if (v49 > 0.1) {
                        v58 := GenS8{F0: -9223372036854775808, F1: 912, F2: -9223372036854775808}
                        _ = v58
                    } else {
                        var v59 int64 = -1
                        _ = v59
                    }
                case "":
                    fmt.Println(v49)
                case "\n\t":
                    v49 = ((v49 * 1e10) - float32((((1.5 * 1e10) / 3.14159) / float64(v49))))
                }
                v60 := GenS5{F0: 42, F1: 42}
                _ = v60
            }
            _ = v49
        }
    }
    return 42
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]bool)
    var v3 uint = 0
    v61 := GenS1{F0: 0, F1: 982}
    v3 += 2
    for fi62 := 0; fi62 < 3; fi62++ {
        v63 := v2[("\n\t" + "hello")]
        _ = fi62
        _ = v63
    }
    delete(v2, ("hello" + "hello"))
    v1 += v1
    fmt.Println(v3)
    if true {
        fmt.Println(v1, v3)
        if (v1 <= (v1 / (0.1 / 3.14159))) {
            for fi64 := 0; fi64 < 1; fi64++ {
                v65 := GenS9{F0: 42, F1: (fi64 + 128), F2: ("hello" + ("hello" + ("hello" + (("\n\t" + ("\n\t" + ("foo bar" + ""))) + "\n\t"))))}
                v3 -= v3
                _ = fi64
                _ = v65
            }
            v1 -= 1e10
        } else {
            var v66 uint16 = 42
            _ = v66
        }
        fmt.Println(v3)
    } else {
        v67 := GenS9{F0: 42, F1: -446, F2: "foo bar"}
        switch "hello" {
        case "\n\t":
            v68 := GenS10{F0: 1}
            _ = v68
        case "a":
            for fi69 := 0; fi69 < 15; fi69++ {
                for fi70 := 0; fi70 < 11; fi70++ {
                    var v71 uint = (v3 % ((v3) | 1))
                    _ = fi70
                    _ = v71
                }
                _ = fi69
            }
        case "foo bar":
            if (205 >= 65535) {
                if (-121 != 9) {
                    v1 = float32((1e10 + 0.0))
                    var v72 int8 = -124
                    _ = v72
                } else {
                    var v73 uint16 = 0
                    _ = v73
                }
                var v74 bool = (332 != 7)
                fmt.Println(v1, v1, v3)
                _ = v74
            }
        default:
            v75 := GenS7{F0: 42, F1: -921, F2: 127}
            _ = v75
        }
        _ = v67
    }
    v2[("foo bar" + ((("foo bar" + "foo bar") + "") + "foo bar"))] = true
    v76 := GenS2{F0: 255}
    if (1.5 == 1e10) {
        fmt.Println(v3, v3)
        var v77 uint8 = 198
        _ = v77
    }
    fmt.Println(v3, v3)
    v1 = (v1 * 1.5)
    var v78 string = "\n\t"
    if (!(false && false)) {
        for fi79 := 0; fi79 < 8; fi79++ {
            fmt.Println(v1)
            _ = fi79
        }
    }
    for fi80 := 0; fi80 < 4; fi80++ {
        v3 = 2
        v81 := GenS11{F0: -456}
        _ = fi80
        _ = v81
    }
    v78 = (v78 + "hello")
    var v82 string = v78
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v78:", v78)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v76
    _ = v78
    _ = v82
}
