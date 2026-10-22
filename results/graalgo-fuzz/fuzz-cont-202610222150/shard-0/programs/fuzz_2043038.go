package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint; F2 int }
type GenS3 struct { F0 uint64; F1 uint; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[string]string)
    v2 := []int{32767, -9223372036854775808, 128, 42}
    var v3 uint64 = 4294967295
    var v4 uint32 = 2
    v4 -= (v4 & v4)
    var v5 uint32 = (v4 + (^(v4)))
    v1[("foo bar" + "\n\t")] = "foo bar"
    if true {
        switch -40 {
        case 429:
            v2 = append(v2, -635)
        case -819:
            var v6 uint16 = 1
            _ = v6
        }
        for fi7 := 0; fi7 < 2; fi7++ {
            v4 += 42
            fmt.Println(fi7, v5)
            for fi8 := 0; fi8 < 3; fi8++ {
                v5 -= 1
                fmt.Println(fi8, v5)
                switch fi7 {
                case 2147483647:
                    v9 := GenS1{F0: 42}
                    _ = v9
                }
                _ = fi8
            }
            _ = fi7
        }
        _ = len(v1)
    }
    fmt.Println(v4, v4, v4)
    for fi10 := 0; fi10 < 12; fi10++ {
        if true {
            if (!(("a" + "hello") >= "foo bar")) {
                if false {
                    fmt.Println(fi10, v5, fi10)
                } else {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        var v12 bool = true
                        _ = fi11
                        _ = v12
                    }
                    v3 -= 4294967295
                    for fi13 := 0; fi13 < 16; fi13++ {
                        var v14 int64 = 1
                        var v15 int64 = v14
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    if true {
                        if true {
                            switch "\n\t" {
                            case "hello":
                                var v17 float32 = 3.14159
                                _ = v17
                            default:
                                if (18446744073709551615 > uint(v5)) {
                                    if (true && true) {
                                        v3 = (v3 * 907)
                                        var v18 bool = true
                                        if v18 {
                                            var v19 uint32 = 142
                                            _ = v19
                                        }
                                        _ = len(v2)
                                        _ = v18
                                    }
                                } else {
                                    _ = len(v1)
                                }
                            }
                        } else {
                            fmt.Println(fi10, v4)
                            var v20 uint16 = 42
                            _ = v20
                        }
                    } else {
                        v5 -= (^(v5))
                        if true {
                            v3 -= 2
                            var v21 uint8 = 127
                            _ = v21
                        } else {
                            var v22 int32 = -1
                            var v23 uint64 = (v3 >> ((837) & 63))
                            v3 = 42
                            _ = v22
                            _ = v23
                        }
                        switch (fi10 | fi10) {
                        case 0:
                            if true {
                                var v24 uint8 = 255
                                if (-45 > 5) {
                                    var v25 uint = 255
                                    _ = v25
                                }
                                _ = v24
                            }
                        case 9223372036854775807:
                            if (-2147483648 != -377) {
                                if false {
                                    fmt.Println(fi16, fi16, v5)
                                } else {
                                    switch fi10 {
                                    case 9223372036854775807:
                                        var v26 int8 = 42
                                        _ = v26
                                    case 32768:
                                        v3 = (v3 & 42)
                                    case 2147483647:
                                        var v27 int64 = 680
                                        if ((18446744073709551615 >= 42) || (v3 <= 18446744073709551615)) {
                                            if false {
                                                _ = len(v1)
                                                v4 = (v4 ^ (^(v4)))
                                                if (41 <= (v27 % ((-9223372036854775808) | 1))) {
                                                    _ = len(v2)
                                                    var v28 uint64 = v3
                                                    _ = v28
                                                }
                                            } else {
                                                fmt.Println(v4, fi10, fi16)
                                                if true {
                                                    v29 := GenS1{F0: 903}
                                                    _ = v29
                                                }
                                                var v30 uint32 = 1
                                                fmt.Println(fi16, v5)
                                                var v31 int64 = (v27 << ((313) & 63))
                                                if false {
                                                    fmt.Println(v3)
                                                    v3 = v3
                                                    v5 -= 65535
                                                } else {
                                                    switch fi10 {
                                                    case -1:
                                                        if true {
                                                            var v32 string = ("a" + "\n\t")
                                                            _ = v32
                                                        }
                                                    default:
                                                        v3 += (v3 + (^(v3)))
                                                    }
                                                }
                                                _ = v30
                                                _ = v31
                                            }
                                            if true {
                                                switch fi16 {
                                                case 0:
                                                    switch "hello" {
                                                    case "hello":
                                                        v33 := GenS1{F0: 892}
                                                        _ = v33
                                                    default:
                                                        var v34 int32 = 42
                                                        _ = v34
                                                    }
                                                case 4294967295:
                                                    var v35 uint64 = v3
                                                    _ = v35
                                                case 32767:
                                                    if (!false) {
                                                        v3 = v3
                                                    } else {
                                                        var v36 uint32 = (^(v5))
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v37 uint8 = 128
                                                            _ = v37
                                                        case "\n\t":
                                                            v1["\n\t"] = "\n\t"
                                                        case "foo bar":
                                                            if true {
                                                                if true {
                                                                    var v38 int64 = v27
                                                                    if true {
                                                                        v36 = 4294967295
                                                                        v38 -= v38
                                                                    } else {
                                                                        fmt.Println(v5, v27)
                                                                        fmt.Println(fi10, v3, v36)
                                                                    }
                                                                    v39 := GenS2{F0: 156, F1: 115, F2: fi10}
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    v27 -= -1
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        var v40 int16 = -32768
                                                                        _ = v40
                                                                    case "":
                                                                        v41 := GenS2{F0: 4294967295, F1: 4294967295, F2: (fi16 % (((-(fi10))) | 1))}
                                                                        var v42 float32 = 1e10
                                                                        _ = len(v2)
                                                                        _ = v41
                                                                        _ = v42
                                                                    default:
                                                                        fmt.Println(v4, v3, v4)
                                                                    }
                                                                    if ((!true) || ((((((("\n\t" + ("\n\t" + "foo bar")) + (("hello" + "foo bar") + "hello")) + "foo bar") + "a") + "a") + "hello") < "a")) {
                                                                        v4 += v36
                                                                        v43 := GenS2{F0: 0, F1: 2, F2: (fi10 / (((fi16 ^ fi10)) | 1))}
                                                                        _ = v43
                                                                    } else {
                                                                        if ((!true) && (465 < 0)) {
                                                                            switch "" {
                                                                            case "":
                                                                                fmt.Println(v36, v36, v36)
                                                                            case "\n\t":
                                                                                fmt.Println(v27, v5, v3)
                                                                            }
                                                                            if (!(128 == 128)) {
                                                                                v44 := GenS1{F0: 4294967295}
                                                                                _ = len(v1)
                                                                                v36 += (v36 * (v36 * (^(v4))))
                                                                                _ = v44
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                fmt.Println(fi16, fi16)
                                                                                v3 = (v3 << ((719) & 63))
                                                                            }
                                                                        }
                                                                        v45 := GenS3{F0: (^(v3)), F1: 42, F2: 0}
                                                                        var v46 uint8 = 126
                                                                        _ = v45
                                                                        _ = v46
                                                                    }
                                                                }
                                                                v36 = (v5 % (((v5 ^ (v5 % (((v36 | 42)) | 1)))) | 1))
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            } else {
                                                                var v47 int8 = -128
                                                                if len(v2) > 0 {
                                                                    v2[0] = 910
                                                                }
                                                                v36 -= v5
                                                                _ = v47
                                                            }
                                                        default:
                                                            v48 := GenS4{F0: -2147483648}
                                                            _ = v48
                                                        }
                                                        _ = v36
                                                    }
                                                default:
                                                    var v49 uint64 = (^(v3))
                                                    _ = v49
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                if (!true) {
                                                    v27 += v27
                                                    var v50 uint = 65535
                                                    if false {
                                                        v5 += v4
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v51 := GenS3{F0: (v3 / ((uint64(v50)) | 1)), F1: v50, F2: -1}
                                                        v52 := GenS5{F0: -128}
                                                        v53 := GenS4{F0: 742}
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    _ = v50
                                                }
                                            } else {
                                                fmt.Println(v5)
                                                v54 := GenS4{F0: 42}
                                                _ = v54
                                            }
                                        } else {
                                            var v55 bool = true
                                            _ = v55
                                        }
                                        _ = v27
                                    default:
                                        v56 := GenS6{F0: 127}
                                        _ = v56
                                    }
                                    fmt.Println(v3, fi16)
                                }
                            }
                        case 661:
                            fmt.Println(v3, fi16)
                        }
                    }
                    if len(v2) > 0 {
                        v2[0] = (fi16 + fi16)
                    }
                    _ = fi16
                }
            } else {
                v5 = 2147483648
            }
        }
        v3 = (v3 - 0)
        _ = fi10
    }
    v57 := GenS7{F0: 4294967295}
    v3 = (v3 & v3)
    var v58 int8 = -1
    for fi59 := 0; fi59 < 2; fi59++ {
        v60 := GenS5{F0: v58}
        var v61 uint8 = 1
        switch (fi59 << ((18446744073709551615) & 63)) {
        case 187:
            v3 = 1
        default:
            var v62 int = fi59
            _ = v62
        }
        _ = fi59
        _ = v60
        _ = v61
    }
    v63 := GenS5{F0: v58}
    var v64 int8 = v58
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v63
    _ = v64
}
