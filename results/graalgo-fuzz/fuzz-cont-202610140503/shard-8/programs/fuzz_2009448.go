package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS8 struct { F0 int64; F1 int8; F2 bool }
type GenS9 struct { F0 int32 }

func main() {
    var v1 uint8 = 163
    var v2 uint16 = 255
    v3 := make(map[uint32]bool)
    var v4 uint64 = 234
    v5 := make(map[string]string)
    v6 := GenS1{F0: 144, F1: 2147483647, F2: 0}
    v7 := GenS2{F0: ((v4 % ((760) | 1)) != 408)}
    fmt.Println(v1, v1)
    v2 = v2
    v8 := GenS2{F0: true}
    v9 := GenS2{F0: ((v4 < v4) || true)}
    if true {
        v10 := []int{}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        switch 696 {
        case -961:
            if (false && (false && false)) {
                fmt.Println(v1, v4, v1)
                _ = len(v10)
                for fi11 := 0; fi11 < 18; fi11++ {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v13 := GenS3{F0: 0, F1: 100}
                        _ = fi12
                        _ = v13
                    }
                    v14 := v5[("a" + "hello")]
                    switch "foo bar" {
                    case "":
                        _ = len(v5)
                    case "foo bar":
                        for fi15 := 0; fi15 < 13; fi15++ {
                            fmt.Println(v2)
                            var v16 int8 = 42
                            _ = fi15
                            _ = v16
                        }
                    case "\n\t":
                        if false {
                            switch "hello" {
                            case "\n\t":
                                v1 = (^(v1))
                            default:
                                if len(v10) > 0 {
                                    v10[0] = fi11
                                }
                            }
                            var v17 int = fi11
                            fmt.Println(v14, fi11, v4)
                            _ = v17
                        }
                    }
                    _ = fi11
                    _ = v14
                }
            } else {
                v18 := GenS3{F0: 2147483647, F1: 44}
                var v19 int8 = 1
                _ = v18
                _ = v19
            }
        }
        _ = v10
    } else {
        if (!true) {
            var v20 float64 = 1e10
            _ = v20
        } else {
            delete(v3, 42)
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS4{F0: fi21}
                var v23 int = -1
                var v24 int = v23
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 string = ("" + "")
            _ = v25
        }
        v3[1] = (0.0 >= (0.1 - 0.1))
    }
    if (2147483647 < 949) {
        fmt.Println(v2)
        v4 -= v4
        v4 = 238
    }
    var v26 uint32 = uint32(v4)
    if true {
        v27 := GenS3{F0: 42, F1: 1}
        _ = v27
    }
    if false {
        switch 32767 {
        case 32767:
            v28 := GenS1{F0: 1, F1: 1, F2: -128}
            _ = v28
        }
        var v29 float32 = (3.14159 * float32(1.5))
        _ = v29
    } else {
        if (!true) {
            for fi30 := 0; fi30 < 1; fi30++ {
                var v31 bool = true
                _ = fi30
                _ = v31
            }
            v32 := []int{649}
            _ = v32
        }
        fmt.Println(v1, v1)
    }
    if ((0.1 * ((1e10 + float32(1e10)) + 1.5)) > 0.0) {
        var v33 uint8 = v1
        var v34 float32 = 1e10
        v35 := GenS1{F0: -1, F1: -60, F2: -1}
        _ = v33
        _ = v34
        _ = v35
    } else {
        v1 += v1
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        v37 := GenS5{F0: (v4 * (v4 % (((^(v4))) | 1))), F1: 255}
        if (v2 <= uint16(v4)) {
            var v38 uint16 = (v2 | (v2 * 42))
            switch (fi36 & (fi36 + (fi36 ^ fi36))) {
            case 4294967295:
                v39 := GenS5{F0: v4, F1: 2}
                _ = v39
            default:
                var v40 uint8 = 127
                _ = v40
            }
            for fi41 := 0; fi41 < 20; fi41++ {
                v2 = v2
                if false {
                    v38 -= v2
                    if false {
                        if (false || (!(!true))) {
                            if (((true && false) || true) && (!(int16(v26) < 42))) {
                                var v42 int32 = 42
                                _ = v42
                            }
                        } else {
                            v43 := GenS5{F0: (v4 >> ((1) & 63)), F1: 0}
                            var v44 uint64 = 2
                            if true {
                                if (true && (!false)) {
                                    var v45 bool = true
                                    _ = v45
                                } else {
                                    var v46 int8 = 127
                                    if (((^(v44)) != 0) && true) {
                                        var v47 int = (-(fi41))
                                        fmt.Println(fi36, v44, v1)
                                        var v48 string = ("" + "a")
                                        _ = v47
                                        _ = v48
                                    } else {
                                        v49 := []int{fi41, (fi36 % ((fi36) | 1))}
                                        v50 := GenS6{F0: -1, F1: 1}
                                        v51 := GenS7{F0: (v1 + 255), F1: 0, F2: (^(v1))}
                                        var v52 uint32 = v26
                                        var v53 int = (^(fi41))
                                        v44 -= v44
                                        if (1e10 >= (((1e10 / 1.0) - 0.1) + ((0.0 * float32(0.0)) / float32(3.14159)))) {
                                            var v54 int64 = 9223372036854775807
                                            _ = v54
                                        } else {
                                            var v55 uint32 = (v26 / ((0) | 1))
                                            v56 := GenS1{F0: 9223372036854775807, F1: (fi36 | (fi41 >> ((0) & 63))), F2: v46}
                                            _ = v55
                                            _ = v56
                                        }
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                    fmt.Println(fi36)
                                    if false {
                                        fmt.Println(v26, v44)
                                    } else {
                                        if (!true) {
                                            v57 := GenS7{F0: (v1 | 130), F1: (v1 | 128), F2: v1}
                                            var v58 int32 = 42
                                            switch ("hello" + "hello") {
                                            case "hello":
                                                v44 += v4
                                            case "\n\t":
                                                var v59 uint16 = 32767
                                                _ = v59
                                            case "a":
                                                if (!((!(v26 > v26)) || (0 >= -2147483648))) {
                                                    v60 := v3[(v26 << ((uint(v1)) & 63))]
                                                    var v61 int8 = (v46 | v46)
                                                    var v62 uint8 = v1
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    v1 = v1
                                                    delete(v5, "foo bar")
                                                    var v63 string = "foo bar"
                                                    _ = v63
                                                }
                                            }
                                            var v64 uint8 = (v1 % ((128) | 1))
                                            v65 := GenS6{F0: -573, F1: 251}
                                            _ = v57
                                            _ = v58
                                            _ = v64
                                            _ = v65
                                        }
                                        v26 -= v26
                                        v66 := GenS8{F0: 1, F1: (v46 & v46), F2: (int64(fi41) != -16)}
                                        _ = v66
                                    }
                                    v26 = (^(v26))
                                    var v67 int = int(v46)
                                    _ = v46
                                    _ = v67
                                }
                            } else {
                                if false {
                                    fmt.Println(v4, v4)
                                    v68 := GenS3{F0: 42, F1: -29}
                                    v4 = v4
                                    _ = v68
                                }
                                v69 := GenS2{F0: false}
                                _ = v69
                            }
                            _ = v43
                            _ = v44
                        }
                    } else {
                        v2 = (v2 % ((v38) | 1))
                        fmt.Println(fi41, v38)
                        var v70 uint32 = (v26 * (v26 / (((v26 << ((42) & 63))) | 1)))
                        _ = v70
                    }
                }
                if (-32768 < 745) {
                    var v71 uint32 = 2147483647
                    _ = v71
                } else {
                    fmt.Println(v26, v1, v38)
                    switch (fi36 & -31) {
                    case 9223372036854775807:
                        var v72 uint8 = (v1 % ((128) | 1))
                        _ = v72
                    case 4294967295:
                        switch "" {
                        case "hello":
                            v2 = 255
                            fmt.Println(fi36)
                            v26 += (v26 >> ((144) & 63))
                        case "foo bar":
                            v73 := v3[(v26 - (v26 ^ 2147483648))]
                            _ = v73
                        case "":
                            var v74 int = 688
                            _ = v74
                        }
                    }
                    if (!(3.14159 != 0.1)) {
                        var v75 int64 = 42
                        _ = v75
                    }
                    v26 -= uint32(v4)
                    fmt.Println(v4, fi41)
                }
                _ = fi41
            }
            _ = v38
        } else {
            if (2147483648 != 402) {
                v1 = (v1 & v1)
            } else {
                var v76 uint64 = uint64(v26)
                var v77 float32 = 3.14159
                _ = v76
                _ = v77
            }
        }
        _ = fi36
        _ = v37
    }
    var v78 uint64 = v4
    v79 := GenS9{F0: -1}
    switch "hello" {
    case "\n\t":
        v4 += 18446744073709551615
    case "hello":
        var v80 uint16 = v2
        _ = v80
    case "":
        v81 := []int{}
        _ = v81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v78
    _ = v79
    _ = v8
    _ = v9
}
