package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint; F2 string }
type GenS5 struct { F0 bool; F1 int8; F2 int }
type GenS6 struct { F0 int8; F1 int8; F2 int }
type GenS7 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 0
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 1
    v2 = (^(v2))
    var v5 int = 4294967295
    for fi6 := 0; fi6 < 11; fi6++ {
        v4 = (v4 | 4294967295)
        v7 := v1[(("a" + "") + "foo bar")]
        v8 := v1[(v7 + "\n\t")]
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v4 = v4
        if false {
            var v10 int = v5
            v10 += 332
            v11 := GenS1{F0: (v2 - v2), F1: -101}
            _ = v10
            _ = v11
        } else {
            fmt.Println(v5, v2, v2)
        }
        var v12 bool = true
        _ = fi9
        _ = v12
    }
    switch ("foo bar" + "") {
    case "\n\t":
        for fi13 := 0; fi13 < 4; fi13++ {
            v5 += int(v2)
            var v14 string = "a"
            if false {
                if true {
                    if false {
                        if false {
                            switch (-(fi13)) {
                            case 127:
                                switch ((("" + ("\n\t" + v14)) + v14) + v14) {
                                case "":
                                    v5 += v5
                                case "\n\t":
                                    if (1 < 42) {
                                        var v15 bool = false
                                        var v16 uint32 = 493
                                        _ = v15
                                        _ = v16
                                    }
                                }
                            case 42:
                                v17 := GenS2{F0: ((fi13 == v5) || ((0 != 109) && true)), F1: (v2 ^ v2)}
                                _ = v17
                            default:
                                for fi18 := 0; fi18 < 3; fi18++ {
                                    v2 -= 0
                                    switch ("\n\t" + v14) {
                                    case "foo bar":
                                        var v19 string = v14
                                        _ = v19
                                    case "a":
                                        v20 := GenS3{F0: 2, F1: -9223372036854775808, F2: 2}
                                        _ = v20
                                    case "hello":
                                        v2 = -826
                                        var v21 float64 = (float64(-2.25) - (3.14159 * 1e10))
                                        _ = v21
                                    }
                                    v5 -= 2147483647
                                    _ = fi18
                                }
                            }
                            v2 = 1
                        } else {
                            fmt.Println(v4, v4, v14)
                            if false {
                                var v22 int16 = 1
                                for fi23 := 0; fi23 < 5; fi23++ {
                                    v24 := GenS2{F0: (-26 > 42), F1: -440}
                                    v4 -= 4294967295
                                    if true {
                                        fmt.Println(v22, fi23, fi13)
                                        v4 = (v4 | (v4 >> ((42) & 63)))
                                    } else {
                                        if ((-(v2)) != 42) {
                                            if (733 < int16(v5)) {
                                                v2 = (v2 >> ((0) & 63))
                                                var v25 uint = 0
                                                switch -444 {
                                                case 2147483647:
                                                    fmt.Println(v25)
                                                case 127:
                                                    fmt.Println(v22, v14)
                                                default:
                                                    v22 -= (v22 >> ((v25) & 63))
                                                }
                                                _ = v25
                                            } else {
                                                v14 = v14
                                            }
                                        } else {
                                            var v26 bool = (0.0 > 0.0)
                                            if false {
                                                v27 := GenS3{F0: 1, F1: 0, F2: 2}
                                                _ = v27
                                            } else {
                                                var v28 int32 = (-(v2))
                                                v29 := v3[65535]
                                                _ = v28
                                                _ = v29
                                            }
                                            v14 = ("a" + "a")
                                            if (!(!(v26 && (2 > 65535)))) {
                                                if (0 >= 0) {
                                                    switch (v14 + v14) {
                                                    case "hello":
                                                        var v30 uint32 = 42
                                                        var v31 uint32 = (v30 / (((^(v30))) | 1))
                                                        fmt.Println(v14)
                                                        _ = v30
                                                        _ = v31
                                                    case "a":
                                                        if v26 {
                                                            fmt.Println(v5, v26)
                                                            switch (("\n\t" + v14) + v14) {
                                                            case "a":
                                                                v32 := GenS1{F0: 2147483647, F1: 100}
                                                                _ = v32
                                                            case "hello":
                                                                switch v14 {
                                                                case "foo bar":
                                                                    fmt.Println(v5, v14, v2)
                                                                case "\n\t":
                                                                    v33 := GenS4{F0: ((v22 | v22) >= v22), F1: 970, F2: ("foo bar" + v14)}
                                                                    _ = v33
                                                                default:
                                                                    var v34 uint16 = 32767
                                                                    _ = v34
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v22)
                                                            }
                                                            fmt.Println(v14, v26)
                                                            var v35 uint32 = 2147483647
                                                            _ = v35
                                                        } else {
                                                            if false {
                                                                v36 := GenS3{F0: 2, F1: 9223372036854775807, F2: 2}
                                                                _ = v36
                                                            } else {
                                                                if v26 {
                                                                    v5 = (fi13 ^ -2147483648)
                                                                    var v37 uint16 = 0
                                                                    _ = v37
                                                                }
                                                                v38 := []int{fi23, (^(fi13)), 128}
                                                                _ = v38
                                                            }
                                                        }
                                                    default:
                                                        if v26 {
                                                            var v39 uint = 273
                                                            _ = v39
                                                        }
                                                    }
                                                } else {
                                                    var v40 float64 = 0.0
                                                    if v26 {
                                                        var v41 float32 = (1.5 / 3.14159)
                                                        _ = v41
                                                    }
                                                    _ = v40
                                                }
                                                fmt.Println(v5, v2)
                                            }
                                            var v42 uint8 = 1
                                            _ = v26
                                            _ = v42
                                        }
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                                _ = v22
                            } else {
                                var v43 uint32 = 2147483648
                                var v44 float32 = 3.14159
                                v5 = (v5 | v5)
                                _ = v43
                                _ = v44
                            }
                            for fi45 := 0; fi45 < 6; fi45++ {
                                fmt.Println(fi13)
                                if (9223372036854775807 >= 0) {
                                    var v46 int8 = 119
                                    _ = v46
                                } else {
                                    if ((798 == -9223372036854775808) || (1.5 >= -2.25)) {
                                        v5 = (-(fi13))
                                    }
                                }
                                v47 := GenS1{F0: v2, F1: 1}
                                _ = fi45
                                _ = v47
                            }
                        }
                        v48 := GenS2{F0: (true || true), F1: 42}
                        var v49 uint64 = v4
                        _ = v48
                        _ = v49
                    }
                    v50 := GenS5{F0: true, F1: 42, F2: int(v2)}
                    fmt.Println(v4)
                    _ = v50
                }
                v51 := v3[914]
                var v52 uint = 1
                _ = v51
                _ = v52
            } else {
                v53 := []int{}
                if ((!(1 == -700)) && (42 > 181)) {
                    var v54 int8 = 100
                    _ = v54
                }
                _ = v53
            }
            _ = fi13
            _ = v14
        }
    case "foo bar":
        if (false && false) {
            var v55 uint16 = 2
            fmt.Println(v4)
            if false {
                v56 := v3[0]
                _ = v56
            }
            _ = v55
        }
    }
    v4 += (v4 | 1)
    switch "a" {
    case "":
        for fi57 := 0; fi57 < 15; fi57++ {
            v58 := GenS6{F0: 42, F1: -128, F2: 32768}
            v5 = 949
            _ = fi57
            _ = v58
        }
    case "a":
        var v59 int16 = -1
        _ = v59
    }
    v4 += (v4 + v4)
    var v60 int8 = -65
    if (!false) {
        v61 := GenS3{F0: 1, F1: 0, F2: 2}
        var v62 int = v5
        fmt.Println(v4, v4, v60)
        _ = v61
        _ = v62
    } else {
        v63 := GenS7{F0: 32767, F1: 327}
        _ = v63
    }
    var v64 int8 = (^(v60))
    v64 = v64
    v64 = 1
    v65 := []int{(v5 & (v5 ^ -797))}
    _ = len(v65)
    for fi66 := 0; fi66 < 4; fi66++ {
        v5 = fi66
        v4 += 4294967295
        var v67 int8 = (v60 | v60)
        _ = fi66
        _ = v67
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v64
    _ = v65
}
