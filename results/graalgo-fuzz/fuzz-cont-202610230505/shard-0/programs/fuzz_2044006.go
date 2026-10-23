package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }
type GenS2 struct { F0 int; F1 string; F2 string }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int8 }
func genFn0() uint32 {
    switch ("" + "hello") {
    case "\n\t":
        v4 := GenS1{F0: (65535 < 255), F1: 794, F2: "hello"}
        _ = v4
    case "":
        var v5 int32 = -2147483648
        _ = v5
    }
    switch "hello" {
    case "foo bar":
        var v6 bool = (0 < 42)
        _ = v6
    case "a":
        if ((2 >= 18446744073709551615) || (983 <= 587)) {
            if false {
                if true {
                    if (!(-9223372036854775808 <= 1)) {
                        for fi7 := 0; fi7 < 5; fi7++ {
                            var v8 int = (-(fi7))
                            v9 := GenS2{F0: fi7, F1: "\n\t", F2: ""}
                            v8 += (fi7 << ((4294967295) & 63))
                            _ = fi7
                            _ = v8
                            _ = v9
                        }
                        v10 := GenS2{F0: 81, F1: "\n\t", F2: ""}
                        _ = v10
                    }
                } else {
                    var v11 bool = (!(0 > 623))
                    if true {
                        if (0.1 <= 0.1) {
                            for fi12 := 0; fi12 < 6; fi12++ {
                                v13 := []int{(fi12 - fi12), 2147483647, 4294967295}
                                if v11 {
                                    fmt.Println(v11, fi12)
                                }
                                _ = fi12
                                _ = v13
                            }
                        }
                        v11 = v11
                        var v14 int32 = 2147483647
                        _ = v14
                    } else {
                        var v15 bool = false
                        v15 = (119 == 1)
                        var v16 uint = 920
                        _ = v15
                        _ = v16
                    }
                    _ = v11
                }
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int = fi17
                    _ = fi17
                    _ = v18
                }
            } else {
                var v19 uint64 = 2
                _ = v19
            }
            v20 := []int{-111}
            _ = v20
        }
    case "":
        v21 := []int{2147483647, -490, 42}
        _ = v21
    default:
        for fi22 := 0; fi22 < 20; fi22++ {
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 string = ""
                var v25 string = v24
                _ = fi23
                _ = v24
                _ = v25
            }
            _ = fi22
        }
    }
    var v26 uint64 = 1
    _ = v26
    return 2
}

func main() {
    v1 := []uint32{0, 4294967295}
    var v2 uint16 = 32767
    v3 := []int{4294967295}
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 uint8 = uint8(v2)
        fmt.Println(v28, fi27, v28)
        for fi29 := 0; fi29 < 12; fi29++ {
            v2 += v2
            switch fi29 {
            case 32767:
                if true {
                    if false {
                        var v30 uint8 = (v28 % ((0) | 1))
                        v2 = (v2 - uint16(v30))
                        _ = v30
                    }
                } else {
                    if false {
                        if (!true) {
                            if len(v1) > 0 {
                                v1[0] = genFn0()
                            }
                        }
                        var v31 int = fi29
                        fmt.Println(v28, v31)
                        _ = v31
                    }
                    var v32 int8 = 41
                    v33 := GenS3{F0: 257, F1: 65535, F2: -28}
                    _ = v32
                    _ = v33
                }
            case 4294967295:
                v34 := GenS4{F0: 1}
                _ = v34
            case -907:
                var v35 uint32 = 2147483648
                if (822 >= 0) {
                    var v36 bool = false
                    v3 = append(v3, fi29)
                    if (3.14159 > 0.1) {
                        var v37 uint = 65535
                        v37 = v37
                        _ = len(v3)
                        if ((v28 | v28) < (v28 >> ((v37) & 63))) {
                            switch (fi27 / ((fi29) | 1)) {
                            case 9223372036854775807:
                                fmt.Println(v35, v36, fi29)
                            case 4294967295:
                                fmt.Println(v2, fi29, v28)
                            case -294:
                                fmt.Println(v2, v28)
                            }
                            fmt.Println(fi27, v2, v35)
                            v38 := GenS3{F0: 1, F1: 2147483648, F2: int16(fi27)}
                            _ = v38
                        }
                        v37 += 2
                        v1 = append(v1, 2147483648)
                        switch ("\n\t" + "foo bar") {
                        case "":
                            if (v37 == (v37 / ((v37) | 1))) {
                                v39 := GenS5{F0: 42, F1: 270}
                                _ = v39
                            } else {
                                v3 = append(v3, fi29)
                                var v40 uint16 = (v2 * (v2 - (v2 * (v2 - (^(v2))))))
                                var v41 uint16 = v40
                                fmt.Println(v2, v2)
                                v40 += (v41 & (v2 - (v2 ^ (v40 | 907))))
                                v42 := GenS4{F0: 42}
                                _ = len(v3)
                                var v43 int = int(v35)
                                if len(v3) > 0 {
                                    v3[0] = (fi27 ^ fi27)
                                }
                                fmt.Println(v41, v2)
                                switch fi27 {
                                case 0:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case -391:
                                    switch v43 {
                                    case 9223372036854775807:
                                        v43 += fi29
                                    }
                                case 2147483647:
                                    if true {
                                        var v44 uint64 = 42
                                        v3 = append(v3, (fi27 & v43))
                                        _ = v44
                                    } else {
                                        fmt.Println(v2)
                                    }
                                default:
                                    if v36 {
                                        v45 := GenS6{F0: 42, F1: ""}
                                        _ = v45
                                    } else {
                                        v46 := GenS4{F0: 1}
                                        v47 := GenS1{F0: false, F1: 970, F2: "hello"}
                                        _ = v46
                                        _ = v47
                                    }
                                }
                                v3 = append(v3, 9223372036854775807)
                                _ = v40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                        }
                        _ = v37
                    }
                    var v48 bool = (1.5 <= ((float64(float32(0.0)) / float64((1e10 + float32(0.0)))) - 3.14159))
                    _ = v36
                    _ = v48
                }
                var v49 int64 = 9223372036854775807
                if (!false) {
                    v50 := GenS6{F0: 0, F1: "\n\t"}
                    switch fi29 {
                    case 1:
                        if (!false) {
                            v35 -= genFn0()
                            v49 = v49
                            switch ("a" + "hello") {
                            case "a":
                                if true {
                                    if ((v2 + v2) > v2) {
                                        var v51 string = ("a" + "a")
                                        var v52 float64 = 3.14159
                                        v53 := GenS7{F0: 105}
                                        v52 -= v52
                                        var v54 int64 = (^(v49))
                                        var v55 float64 = 0.1
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if (0 >= -973) {
                                            if ("hello" != "\n\t") {
                                                if true {
                                                    v2 = (v2 % ((v2) | 1))
                                                    switch "foo bar" {
                                                    case "a":
                                                        v2 = 423
                                                    }
                                                }
                                                var v56 int8 = 73
                                                v28 = v28
                                                _ = v56
                                            }
                                            v1 = append(v1, (^(v35)))
                                        }
                                    }
                                }
                            default:
                                v2 += 65535
                            }
                        } else {
                            v35 += (v35 + v35)
                        }
                        v57 := GenS6{F0: 13, F1: "a"}
                        _ = v57
                    default:
                        if false {
                            var v58 float64 = 3.14159
                            _ = v58
                        }
                        fmt.Println(v49, fi27)
                    }
                    _ = v50
                }
                if false {
                    var v59 int16 = 32767
                    _ = v59
                }
                _ = v35
                _ = v49
            default:
                var v60 int16 = -32768
                _ = v60
            }
            _ = fi29
        }
        _ = fi27
        _ = v28
    }
    v2 = v2
    var v61 float32 = 0.1
    v61 = 3.14159
    v62 := GenS7{F0: 112}
    v61 = 1e10
    var v63 string = ((("hello" + "\n\t") + ("" + "\n\t")) + "a")
    var v64 int64 = 0
    var v65 uint16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
