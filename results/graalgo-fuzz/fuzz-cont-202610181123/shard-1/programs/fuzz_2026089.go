package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 int; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 uint }

func main() {
    var v1 int = -2147483648
    var v2 string = "foo bar"
    var v3 int32 = 0
    var v4 int64 = -306
    var v5 string = (v2 + v2)
    switch (v1 + v1) {
    case -167:
        v5 = v5
    case 127:
        if (v3 >= (v3 * -1)) {
            var v6 int8 = 42
            _ = v6
        } else {
            var v7 int8 = 1
            if (true && true) {
                var v8 int64 = 9223372036854775807
                if (-2.25 == float64(-2.25)) {
                    if false {
                        var v9 int32 = -124
                        v7 += (v7 / ((v7) | 1))
                        if true {
                            var v10 bool = ((2147483648 != 4294967295) || (!(2 > 1)))
                            switch (v2 + (v2 + ((("a" + (v5 + (v5 + "foo bar"))) + "hello") + v5))) {
                            case "\n\t":
                                v11 := []int{127, v1, (v1 ^ v1), 2147483647}
                                _ = v11
                            case "a":
                                if (v2 >= ("a" + v2)) {
                                    if v10 {
                                        var v12 int16 = int16(v4)
                                        var v13 uint16 = 130
                                        if v10 {
                                            var v14 int16 = (v12 >> ((18446744073709551615) & 63))
                                            _ = v14
                                        } else {
                                            for fi15 := 0; fi15 < 17; fi15++ {
                                                switch "\n\t" {
                                                case "hello":
                                                    v16 := GenS1{F0: v13, F1: v13, F2: (-(v3))}
                                                    _ = v16
                                                default:
                                                    v4 -= -522
                                                }
                                                _ = fi15
                                            }
                                        }
                                        _ = v12
                                        _ = v13
                                    } else {
                                        var v17 uint8 = 0
                                        _ = v17
                                    }
                                } else {
                                    fmt.Println(v8, v5, v8)
                                }
                            case "hello":
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    v19 := GenS2{F0: v7}
                                    _ = fi18
                                    _ = v19
                                }
                            default:
                                for fi20 := 0; fi20 < 17; fi20++ {
                                    fmt.Println(v3, v8)
                                    var v21 int8 = 100
                                    for fi22 := 0; fi22 < 11; fi22++ {
                                        var v23 uint16 = 32767
                                        var v24 int32 = v9
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                            }
                            _ = v10
                        }
                        _ = v9
                    }
                    if (false || ((((0.0 + (0.0 / float32(0.1))) == (0.0 + ((1e10 - 0.1) - 3.14159))) && false) || (int32(v7) >= (v3 % ((681) | 1))))) {
                        var v25 int8 = v7
                        v5 = v2
                        _ = v25
                    } else {
                        v26 := GenS3{F0: 0, F1: 2147483647, F2: v5}
                        v2 = "\n\t"
                        switch "hello" {
                        case "a":
                            v27 := GenS3{F0: 1, F1: (^(v1)), F2: v2}
                            _ = v27
                        case "":
                            var v28 float32 = 0.1
                            _ = v28
                        case "foo bar":
                            var v29 int = 127
                            _ = v29
                        default:
                            if ((-2.25 + (((1.5 - float64(0.1)) * 0.0) / 0.1)) > float64((1.5 + 0.0))) {
                                for fi30 := 0; fi30 < 5; fi30++ {
                                    var v31 int8 = (v7 - v7)
                                    var v32 uint = 42
                                    if (2 >= 2) {
                                        for fi33 := 0; fi33 < 9; fi33++ {
                                            v4 += 42
                                            switch (fi33 / (((v1 >> (((v32 / ((v32) | 1))) & 63))) | 1)) {
                                            case 128:
                                                v7 += 100
                                            case -2147483648:
                                                var v34 float32 = ((1e10 / 1e10) - (1.5 * 0.0))
                                                _ = v34
                                            }
                                            _ = fi33
                                        }
                                        if false {
                                            v35 := []int{}
                                            _ = len(v35)
                                            _ = v35
                                        }
                                        if (false && (-2.25 < 3.14159)) {
                                            switch v5 {
                                            case "\n\t":
                                                for fi36 := 0; fi36 < 11; fi36++ {
                                                    v31 = (v31 >> (((v32 ^ 18446744073709551615)) & 63))
                                                    v5 = (("hello" + "hello") + v2)
                                                    var v37 uint64 = 676
                                                    _ = fi36
                                                    _ = v37
                                                }
                                            case "a":
                                                fmt.Println(v2, v32, v7)
                                            case "":
                                                var v38 int32 = int32(v31)
                                                _ = v38
                                            }
                                            v5 = "hello"
                                            if true {
                                                v4 = (v4 >> ((42) & 63))
                                                v7 -= v31
                                                if ((v1 >> (((v32 - 65535)) & 63)) >= (^(v1))) {
                                                    v39 := []int{int(v31), (v1 - (-(fi30))), 9223372036854775807, fi30}
                                                    v8 -= (v8 % ((v4) | 1))
                                                    _ = v39
                                                } else {
                                                    if false {
                                                        fmt.Println(v1)
                                                        fmt.Println(v8, v31, v8)
                                                        fmt.Println(v31, v1)
                                                    } else {
                                                        var v40 uint32 = 0
                                                        _ = v40
                                                    }
                                                    var v41 float64 = 0.0
                                                    fmt.Println(v7, v32)
                                                    _ = v41
                                                }
                                            } else {
                                                var v42 uint64 = 743
                                                _ = v42
                                            }
                                        }
                                    } else {
                                        var v43 int64 = v8
                                        _ = v43
                                    }
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                v2 = ((v2 + (v5 + ("hello" + "a"))) + v5)
                                var v44 bool = true
                                _ = v44
                            } else {
                                if (42 < 2) {
                                    for fi45 := 0; fi45 < 18; fi45++ {
                                        switch ((("" + "\n\t") + v5) + "\n\t") {
                                        case "\n\t":
                                            var v46 int32 = int32(v4)
                                            _ = v46
                                        case "hello":
                                            if (true || true) {
                                                if false {
                                                    fmt.Println(v1, v3)
                                                    v47 := []int{128, fi45}
                                                    _ = v47
                                                }
                                                v2 = ("\n\t" + v5)
                                            }
                                        default:
                                            v7 += v7
                                        }
                                        _ = fi45
                                    }
                                    v3 -= v3
                                    var v48 uint32 = 2147483648
                                    _ = v48
                                }
                                var v49 uint = 4294967295
                                _ = v49
                            }
                        }
                        _ = v26
                    }
                    v3 = v3
                }
                _ = v8
            } else {
                v4 += v4
                v2 = "a"
            }
            _ = v7
        }
    default:
        v5 = v5
    }
    var v50 uint = 584
    var v51 uint64 = 1
    switch (v1 << (((v50 ^ 255)) & 63)) {
    case 4294967295:
        v52 := []int{128}
        _ = v52
    }
    v53 := GenS4{F0: 128}
    var v54 int32 = (^(v3))
    v55 := []int{9223372036854775807, int(v3), -9223372036854775808, v1}
    if false {
        var v56 uint16 = 671
        _ = len(v55)
        var v57 int16 = 517
        _ = v56
        _ = v57
    } else {
        v58 := GenS1{F0: 255, F1: 255, F2: v3}
        v59 := GenS5{F0: v4, F1: 1}
        _ = v58
        _ = v59
    }
    v60 := GenS1{F0: 0, F1: 32767, F2: 2147483647}
    fmt.Println(v3, v54)
    if len(v55) > 0 {
        v55[0] = v1
    }
    fmt.Println(v3, v54)
    if false {
        _ = len(v55)
    }
    v61 := GenS6{F0: 169, F1: 4294967295}
    switch v1 {
    case -574:
        var v62 float64 = (1.5 / float64(1e10))
        _ = v62
    default:
        if false {
            switch "a" {
            case "hello":
                v50 = v50
            default:
                v3 += v3
            }
            if ((!true) && (0 >= (^(v4)))) {
                for fi63 := 0; fi63 < 10; fi63++ {
                    v64 := GenS7{F0: 42}
                    v1 += (v1 % ((int(v51)) | 1))
                    v2 = v2
                    _ = fi63
                    _ = v64
                }
                v50 -= uint(v51)
            }
        } else {
            if len(v55) > 0 {
                v55[0] = v1
            }
        }
    }
    if len(v55) > 0 {
        v55[0] = 0
    }
    for fi65 := 0; fi65 < 12; fi65++ {
        v51 -= 18446744073709551615
        var v66 int8 = 1
        _ = fi65
        _ = v66
    }
    v1 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v55:", len(v55), v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v53
    _ = v54
    _ = v55
    _ = v60
    _ = v61
}
