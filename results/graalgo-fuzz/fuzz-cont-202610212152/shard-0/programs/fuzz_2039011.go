package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int64 }
type GenS2 struct { F0 int; F1 string; F2 int64 }
type GenS3 struct { F0 int; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 bool }

func main() {
    var v1 int32 = -2147483648
    var v2 uint16 = 2
    var v3 float64 = (1.5 - (3.14159 * -2.25))
    v4 := []int{}
    v5 := GenS1{F0: v2, F1: 2, F2: -895}
    switch 1 {
    case -2147483648:
        v6 := GenS2{F0: 1, F1: ("" + "hello"), F2: 1}
        _ = v6
    case 9223372036854775807:
        var v7 int8 = 42
        _ = v7
    }
    if len(v4) > 0 {
        v4[0] = -33
    }
    if false {
        v2 = v2
        if (!(!true)) {
            v1 -= (-(v1))
            fmt.Println(v3, v3, v1)
            if true {
                switch "hello" {
                case "\n\t":
                    for fi8 := 0; fi8 < 2; fi8++ {
                        var v9 uint64 = 18446744073709551615
                        switch (fi8 & 0) {
                        case 128:
                            switch (("foo bar" + "hello") + "foo bar") {
                            case "\n\t":
                                var v10 uint = 552
                                _ = v10
                            default:
                                v9 = 867
                            }
                        case 9223372036854775807:
                            for fi11 := 0; fi11 < 18; fi11++ {
                                v12 := GenS1{F0: v2, F1: 4294967295, F2: 963}
                                if (!(!true)) {
                                    fmt.Println(v1)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v1 += v1
                                    var v13 int = 2147483647
                                    var v14 uint16 = 66
                                    _ = v13
                                    _ = v14
                                } else {
                                    v2 -= v2
                                    var v15 int16 = -32768
                                    _ = v15
                                }
                                _ = fi11
                                _ = v12
                            }
                        case -9223372036854775808:
                            v16 := GenS2{F0: (fi8 % (((fi8 >> ((42) & 63))) | 1)), F1: "hello", F2: -774}
                            _ = v16
                        }
                        _ = fi8
                        _ = v9
                    }
                case "":
                    fmt.Println(v2, v1)
                case "foo bar":
                    fmt.Println(v1, v3, v3)
                }
            }
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            fmt.Println(v1, fi17)
            if (((v2 / (((v2 / (((v2 >> ((18446744073709551615) & 63))) | 1))) | 1)) == v2) && true) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v18 := GenS2{F0: (^(fi17)), F1: "foo bar", F2: -9223372036854775808}
                _ = v18
            }
            _ = fi17
        }
    }
    var v19 uint16 = (v2 << ((4294967295) & 63))
    if (false || false) {
        fmt.Println(v3, v3, v3)
        var v20 uint8 = 94
        v2 -= 2
        _ = v20
    }
    if true {
        var v21 float32 = (float32(v3) * 1e10)
        _ = v21
    }
    var v22 uint16 = v19
    v19 -= v19
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "" {
    case "":
        if (!(-32768 <= int16(v22))) {
            for fi23 := 0; fi23 < 5; fi23++ {
                _ = len(v4)
                _ = fi23
            }
            for fi24 := 0; fi24 < 16; fi24++ {
                v1 += -989
                v2 = 2
                _ = fi24
            }
            if false {
                v3 -= v3
            } else {
                if len(v4) > 0 {
                    v4[0] = 9223372036854775807
                }
                v1 = (^(v1))
                var v25 uint32 = 626
                _ = v25
            }
        } else {
            v26 := GenS2{F0: 127, F1: "\n\t", F2: 1}
            if true {
                var v27 uint32 = 2147483647
                _ = v27
            } else {
                v28 := GenS1{F0: (v22 << ((0) & 63)), F1: 472, F2: 915}
                fmt.Println(v22, v1, v19)
                _ = v28
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                switch fi29 {
                case -1:
                    for fi30 := 0; fi30 < 19; fi30++ {
                        v4 = append(v4, fi29)
                        if true {
                            if false {
                                if true {
                                    v4 = append(v4, (fi30 << ((1) & 63)))
                                    var v31 float64 = (1e10 * (v3 / float64(1.5)))
                                    if true {
                                        _ = len(v4)
                                    } else {
                                        switch ((("foo bar" + "") + "foo bar") + "") {
                                        case "a":
                                            v3 += float64(1e10)
                                        case "":
                                            fmt.Println(v19, v2)
                                        case "hello":
                                            var v32 int = fi29
                                            _ = v32
                                        }
                                        if false {
                                            if (-1 <= 48) {
                                                v19 -= (^(v2))
                                                v2 = v22
                                                if (81 < 128) {
                                                    switch 0 {
                                                    case -976:
                                                        var v33 bool = false
                                                        _ = v33
                                                    case 127:
                                                        v2 = (v19 + v2)
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = fi29
                                                    }
                                                }
                                            } else {
                                                _ = len(v4)
                                                var v34 uint32 = 689
                                                _ = v34
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = fi30
                                            }
                                        }
                                        v1 = -1
                                    }
                                    _ = v31
                                }
                            } else {
                                switch "foo bar" {
                                case "":
                                    if (!(true || true)) {
                                        var v35 uint8 = 108
                                        v4 = append(v4, (^(fi29)))
                                        if true {
                                            v36 := GenS3{F0: (^(fi29)), F1: 42, F2: 100}
                                            v37 := GenS3{F0: (fi30 ^ (fi30 * 1)), F1: 1, F2: 42}
                                            if false {
                                                v22 = v19
                                                var v38 uint = 891
                                                v19 += v19
                                                _ = v38
                                            } else {
                                                if ((v22 ^ (v22 / ((v2) | 1))) >= (^(v2))) {
                                                    fmt.Println(v19)
                                                    v19 += 255
                                                }
                                            }
                                            _ = v36
                                            _ = v37
                                        }
                                        _ = v35
                                    }
                                case "a":
                                    v1 = v1
                                case "hello":
                                    var v39 bool = true
                                    _ = v39
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            var v40 int16 = 88
                            _ = v40
                        } else {
                            if false {
                                if false {
                                    v41 := GenS3{F0: 127, F1: 1, F2: 0}
                                    if true {
                                        v42 := GenS4{F0: 1, F1: 1}
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v42
                                    }
                                    var v43 uint64 = 18446744073709551615
                                    switch 9223372036854775807 {
                                    case 1:
                                        var v44 uint = 1
                                        _ = v44
                                    }
                                    _ = v41
                                    _ = v43
                                }
                                var v45 int8 = 42
                                if true {
                                    v22 -= (v2 / ((133) | 1))
                                    var v46 float32 = 0.1
                                    _ = v46
                                }
                                v22 = 1
                                _ = v45
                            } else {
                                switch "a" {
                                case "":
                                    v3 = 3.14159
                                case "foo bar":
                                    v19 -= (^(v2))
                                    fmt.Println(v19, fi30)
                                case "\n\t":
                                    v1 += (v1 % ((v1) | 1))
                                default:
                                    var v47 int16 = 106
                                    _ = v47
                                }
                            }
                            var v48 uint8 = 42
                            var v49 bool = false
                            fmt.Println(v49, v48)
                            switch "" {
                            case "":
                                fmt.Println(v48, v49)
                            }
                            _ = v48
                            _ = v49
                        }
                        var v50 int = fi29
                        _ = fi30
                        _ = v50
                    }
                case 11:
                    var v51 uint = 65535
                    _ = v51
                case 32767:
                    var v52 string = "hello"
                    _ = v52
                default:
                    v3 += v3
                }
                _ = fi29
            }
            _ = v26
        }
    default:
        if false {
            _ = len(v4)
            v19 -= (v2 * v19)
        } else {
            for fi53 := 0; fi53 < 14; fi53++ {
                for fi54 := 0; fi54 < 20; fi54++ {
                    v4 = append(v4, int(v19))
                    _ = fi54
                }
                _ = fi53
            }
            for fi55 := 0; fi55 < 15; fi55++ {
                var v56 uint64 = 2
                var v57 int16 = 42
                var v58 uint = 129
                _ = fi55
                _ = v56
                _ = v57
                _ = v58
            }
            if (!(65 > 127)) {
                var v59 bool = true
                v60 := GenS5{F0: v59}
                v2 = 2
                _ = v59
                _ = v60
            }
        }
    }
    v61 := GenS5{F0: (18446744073709551615 >= 796)}
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v61
}
