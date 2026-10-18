package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 int; F1 int32; F2 int32 }
type GenS8 struct { F0 int32; F1 uint16 }
type GenS9 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    v4 := []int{}
    v5 := make(map[int]uint32)
    if true {
        var v6 uint = 65535
        var v7 uint = 4294967295
        _ = v6
        _ = v7
    } else {
        switch -410 {
        case 128:
            var v8 string = "a"
            _ = v8
        default:
            var v9 float32 = float32((3.14159 * (1e10 - 1e10)))
            _ = v9
        }
        if false {
            _ = len(v4)
            v10 := GenS1{F0: 608, F1: (v1 & 42)}
            _ = v10
        } else {
            v1 += (v1 ^ v1)
            var v11 uint = 65535
            _ = v11
        }
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v1)
    case "":
        var v12 int8 = 100
        _ = v12
    case "a":
        switch -9223372036854775808 {
        case -378:
            if true {
                var v13 uint = 777
                for fi14 := 0; fi14 < 5; fi14++ {
                    switch "hello" {
                    case "\n\t":
                        switch ("hello" + "a") {
                        case "foo bar":
                            delete(v5, -585)
                        default:
                            fmt.Println(v1)
                        }
                    case "a":
                        v15 := v5[(fi14 * fi14)]
                        _ = v15
                    }
                    _ = fi14
                }
                _ = v13
            } else {
                for fi16 := 0; fi16 < 15; fi16++ {
                    var v17 float64 = 0.1
                    for fi18 := 0; fi18 < 18; fi18++ {
                        if false {
                            v19 := GenS2{F0: v1}
                            var v20 float64 = v17
                            switch "a" {
                            case "a":
                                fmt.Println(fi18)
                            case "hello":
                                if (0 <= 24) {
                                    var v21 string = ("\n\t" + "")
                                    v22 := GenS1{F0: 42, F1: 2147483648}
                                    _ = v21
                                    _ = v22
                                } else {
                                    if false {
                                        v23 := GenS3{F0: 570}
                                        var v24 int32 = -995
                                        _ = v23
                                        _ = v24
                                    } else {
                                        if true {
                                            v25 := GenS3{F0: 1}
                                            v26 := GenS2{F0: v1}
                                            v1 += 0
                                            switch (((("a" + (("a" + (("a" + "hello") + "foo bar")) + "foo bar")) + ("\n\t" + "foo bar")) + "foo bar") + "") {
                                            case "":
                                                var v27 int64 = 1
                                                _ = v27
                                            case "\n\t":
                                                v2[(fi16 ^ -2147483648)] = (("" + ("hello" + "a")) + "a")
                                            }
                                            v20 -= float64((((((-2.25 / (0.1 / 0.1)) + 1.5) * -2.25) * 0.1) / 1.5))
                                            fmt.Println(fi16, v20, v20)
                                            _ = v25
                                            _ = v26
                                        } else {
                                            var v28 int64 = 9223372036854775807
                                            var v29 uint8 = 128
                                            var v30 int16 = 42
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                        var v31 int = fi18
                                        switch (v31 - -717) {
                                        case -752:
                                            var v32 uint = 65535
                                            _ = v32
                                        }
                                        _ = v31
                                    }
                                    fmt.Println(fi18, fi18, v1)
                                }
                            case "\n\t":
                                var v33 string = (("" + "\n\t") + "foo bar")
                                _ = v33
                            }
                            _ = v19
                            _ = v20
                        } else {
                            var v34 uint8 = 36
                            delete(v2, (^(fi16)))
                            fmt.Println(v1, fi16)
                            _ = v34
                        }
                        v35 := GenS1{F0: 17, F1: 1}
                        v36 := GenS4{F0: 0}
                        _ = fi18
                        _ = v35
                        _ = v36
                    }
                    _ = fi16
                    _ = v17
                }
                v1 -= (v1 | 494)
                for fi37 := 0; fi37 < 8; fi37++ {
                    v38 := GenS5{F0: -1}
                    if (false || (-706 != -32768)) {
                        var v39 uint16 = 110
                        v39 -= 255
                        for fi40 := 0; fi40 < 18; fi40++ {
                            if true {
                                if true {
                                    v1 += v1
                                    v1 = (v1 * v1)
                                    var v41 uint = 106
                                    switch (-(fi40)) {
                                    case 32767:
                                        if ((0 >= 127) && (!(133 == 33))) {
                                            if (3.14159 < -2.25) {
                                                v39 = (v39 % (((^(v39))) | 1))
                                                switch "" {
                                                case "":
                                                    var v42 int8 = 100
                                                    _ = v42
                                                case "a":
                                                    switch (("foo bar" + "foo bar") + "") {
                                                    case "\n\t":
                                                        switch (fi40 | fi40) {
                                                        case 42:
                                                            v1 -= (^(v1))
                                                        case 186:
                                                            v43 := GenS6{F0: "foo bar", F1: (^(v1))}
                                                            _ = v43
                                                        }
                                                        var v44 uint16 = v39
                                                        var v45 uint8 = 182
                                                        var v46 int8 = 0
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                case "hello":
                                                    fmt.Println(v39, fi40)
                                                    var v47 uint64 = 4294967295
                                                    v48 := GenS5{F0: int32(fi37)}
                                                    _ = v47
                                                    _ = v48
                                                default:
                                                    var v49 uint8 = 255
                                                    _ = v49
                                                }
                                            }
                                        }
                                    case 2147483647:
                                        fmt.Println(fi40, v1, v39)
                                    }
                                    _ = v41
                                } else {
                                    switch "" {
                                    case "a":
                                        var v50 int32 = 0
                                        _ = v50
                                    }
                                    v51 := GenS3{F0: -9223372036854775808}
                                    var v52 float32 = float32(0.0)
                                    v52 = float32(1.5)
                                    _ = v51
                                    _ = v52
                                }
                                var v53 uint64 = 981
                                var v54 uint16 = 152
                                _ = v53
                                _ = v54
                            } else {
                                var v55 float64 = 0.0
                                if (181 >= -2147483648) {
                                    var v56 bool = true
                                    if v56 {
                                        v57 := GenS4{F0: 487}
                                        var v58 uint = 282
                                        _ = v57
                                        _ = v58
                                    }
                                    _ = v56
                                } else {
                                    var v59 int8 = 114
                                    var v60 float32 = 0.1
                                    fmt.Println(v59)
                                    var v61 string = ("hello" + "\n\t")
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                                _ = v55
                            }
                            var v62 int8 = 100
                            _ = fi40
                            _ = v62
                        }
                        _ = v39
                    }
                    _ = fi37
                    _ = v38
                }
            }
        case -2147483648:
            v1 = 729
        case -363:
            v63 := GenS2{F0: v1}
            _ = v63
        }
    default:
        for fi64 := 0; fi64 < 3; fi64++ {
            v65 := GenS7{F0: 9223372036854775807, F1: 914, F2: -662}
            v4 = append(v4, (fi64 % ((fi64) | 1)))
            _ = fi64
            _ = v65
        }
    }
    if false {
        var v66 bool = true
        for fi67 := 0; fi67 < 16; fi67++ {
            v66 = (v66 || v66)
            if v66 {
                v68 := v2[(fi67 ^ 1)]
                var v69 uint32 = (v1 + (^(v1)))
                _ = v68
                _ = v69
            } else {
                var v70 uint8 = 52
                var v71 uint = 2
                var v72 float64 = float64(-2.25)
                _ = v70
                _ = v71
                _ = v72
            }
            var v73 uint8 = 215
            _ = fi67
            _ = v73
        }
        _ = v66
    }
    v74 := GenS8{F0: -2147483648, F1: 1}
    var v75 int32 = 1
    v76 := GenS9{F0: 128, F1: (!(-1 != -32768))}
    v77 := GenS4{F0: 1}
    delete(v2, int(v75))
    var v78 uint8 = 55
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v75:", v75)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
