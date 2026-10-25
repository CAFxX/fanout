package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 int8; F1 uint32; F2 int }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }

func main() {
    var v1 bool = true
    var v2 uint = 745
    var v3 int8 = 0
    var v4 uint32 = 474
    v5 := []int{-9223372036854775808, 0, 900, 32767}
    v6 := []string{"a", "a"}
    v6 = append(v6, ((("" + "foo bar") + "a") + "a"))
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 2}
            switch (fi7 | (fi8 >> ((4294967295) & 63))) {
            case -757:
                var v10 uint32 = 65535
                _ = v10
            case 32767:
                var v11 float32 = -2.25
                _ = v11
            default:
                v4 = 831
            }
            var v12 int64 = -9223372036854775808
            var v13 int64 = v12
            _ = len(v6)
            var v14 uint32 = v4
            var v15 uint32 = 195
            _ = fi8
            _ = v12
            _ = v13
            _ = v14
            _ = v15
            _ = v9
        }
        switch "hello" {
        case "foo bar":
            v16 := GenS1{F0: 2}
            _ = v16
        default:
            var v17 int16 = -1
            _ = v17
        }
        _ = fi7
    }
    v2 = v2
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int32 = 0
        v3 -= v3
        _ = fi18
        _ = v19
    }
    v5 = append(v5, 345)
    var v20 uint8 = 143
    if v1 {
        v3 = v3
        if v1 {
            v3 += 56
        } else {
            _ = len(v6)
            v20 = (v20 ^ v20)
        }
    } else {
        v21 := GenS1{F0: 1}
        _ = v21
    }
    switch 32767 {
    case 128:
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 uint64 = 18446744073709551615
            _ = fi22
            _ = v23
        }
    case 1:
        v6 = append(v6, "")
    case 741:
        v24 := GenS2{F0: 2, F1: 34, F2: (v3 << (((v2 | (v2 >> (((v2 & 956)) & 63)))) & 63))}
        _ = v24
    }
    switch "hello" {
    case "foo bar":
        v25 := GenS3{F0: (-(v3)), F1: 65535, F2: 0}
        _ = v25
    case "\n\t":
        if v1 {
            switch -582 {
            case -932:
                fmt.Println(v1, v2)
            case 1:
                var v26 bool = true
                _ = v26
            default:
                for fi27 := 0; fi27 < 13; fi27++ {
                    if false {
                        if true {
                            var v28 string = ("\n\t" + "a")
                            var v29 int32 = -2147483648
                            _ = v28
                            _ = v29
                        }
                        v30 := GenS2{F0: 2, F1: (fi27 - 32768), F2: 127}
                        if false {
                            v20 += (v20 | (v20 % ((v20) | 1)))
                        } else {
                            fmt.Println(v3, v1, v20)
                            v31 := GenS1{F0: 394}
                            var v32 string = ((("" + (("hello" + "a") + ("" + "\n\t"))) + (("\n\t" + "") + ("\n\t" + "\n\t"))) + "\n\t")
                            _ = v31
                            _ = v32
                        }
                        _ = v30
                    }
                    _ = fi27
                }
            }
            var v33 uint = v2
            if v1 {
                switch -960 {
                case 32768:
                    _ = len(v5)
                default:
                    var v34 string = ("" + "")
                    _ = v34
                }
            } else {
                switch 1 {
                case 32768:
                    if len(v5) > 0 {
                        v5[0] = 32767
                    }
                case -92:
                    switch "hello" {
                    case "a":
                        for fi35 := 0; fi35 < 2; fi35++ {
                            if v1 {
                                v33 = (^(v2))
                                if v1 {
                                    for fi36 := 0; fi36 < 17; fi36++ {
                                        var v37 string = "\n\t"
                                        _ = fi36
                                        _ = v37
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = (-(fi35))
                                    }
                                }
                            } else {
                                v38 := GenS4{F0: 0, F1: 42}
                                if v1 {
                                    if (fi35 > fi35) {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                } else {
                                    for fi39 := 0; fi39 < 13; fi39++ {
                                        v40 := GenS5{F0: v3, F1: "hello"}
                                        _ = fi39
                                        _ = v40
                                    }
                                    for fi41 := 0; fi41 < 19; fi41++ {
                                        if len(v5) > 0 {
                                            v5[0] = 696
                                        }
                                        v20 = 64
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v3 = v3
                                        v2 = (v33 ^ v2)
                                        _ = fi41
                                    }
                                }
                                if false {
                                    v42 := GenS1{F0: 18446744073709551615}
                                    v2 -= 255
                                    switch -9223372036854775808 {
                                    case 9223372036854775807:
                                        var v43 uint64 = 1
                                        _ = v43
                                    case 32768:
                                        fmt.Println(v33, v4)
                                    }
                                    _ = v42
                                } else {
                                    v20 += v20
                                    if (true && v1) {
                                        for fi44 := 0; fi44 < 15; fi44++ {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            var v45 string = "\n\t"
                                            v46 := GenS3{F0: (v3 / (((v3 - (^(v3)))) | 1)), F1: (v4 % (((v4 ^ v4)) | 1)), F2: (fi44 & fi44)}
                                            _ = fi44
                                            _ = v45
                                            _ = v46
                                        }
                                        v4 += (^(v4))
                                        for fi47 := 0; fi47 < 15; fi47++ {
                                            switch "" {
                                            case "\n\t":
                                                switch 4294967295 {
                                                case 761:
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                case -159:
                                                    if (-2147483648 != 1) {
                                                        fmt.Println(v1, fi35)
                                                        v20 -= v20
                                                    }
                                                }
                                            default:
                                                if ((1 < 1) || (0 > 75)) {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v3 += (v3 | -78)
                                                    case "hello":
                                                        fmt.Println(v20)
                                                    case "foo bar":
                                                        if (v33 >= 507) {
                                                            fmt.Println(fi35, v20, fi47)
                                                            var v48 string = "a"
                                                            _ = v48
                                                        }
                                                        if false {
                                                            var v49 int8 = 100
                                                            v50 := GenS5{F0: 0, F1: (("\n\t" + (("a" + ("" + "\n\t")) + "")) + "\n\t")}
                                                            var v51 int = fi47
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                        if false {
                                                            if (427 >= 42) {
                                                                var v52 uint32 = 526
                                                                fmt.Println(fi47)
                                                                v20 += 136
                                                                _ = v52
                                                            } else {
                                                                v4 = 18
                                                            }
                                                            var v53 int = fi47
                                                            var v54 int = 0
                                                            _ = v53
                                                            _ = v54
                                                        } else {
                                                            switch ("\n\t" + "foo bar") {
                                                            case "":
                                                                fmt.Println(fi47, v1)
                                                            }
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if (1 < 265) {
                                                                    _ = len(v5)
                                                                    var v55 int8 = v3
                                                                    v2 -= v33
                                                                    _ = v55
                                                                }
                                                                fmt.Println(v33, v1, v1)
                                                            }
                                                            v5 = append(v5, int(v2))
                                                        }
                                                    default:
                                                        if v1 {
                                                            v56 := GenS5{F0: -71, F1: ""}
                                                            v57 := GenS3{F0: v3, F1: 2, F2: fi35}
                                                            var v58 float64 = ((1.5 + ((float64(1.5) + 3.14159) / ((0.1 + 0.1) / 3.14159))) / -2.25)
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                        }
                                                    }
                                                    v59 := GenS6{F0: 344}
                                                    if v1 {
                                                        var v60 int64 = -1
                                                        var v61 int64 = (-(v60))
                                                        _ = len(v5)
                                                        switch (fi47 - fi47) {
                                                        case 9223372036854775807:
                                                            if (!v1) {
                                                                fmt.Println(v60)
                                                            }
                                                        }
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                    _ = v59
                                                } else {
                                                    v4 += v4
                                                }
                                            }
                                            _ = fi47
                                        }
                                    } else {
                                        switch fi35 {
                                        case 1:
                                            var v62 int32 = int32(v4)
                                            _ = v62
                                        }
                                    }
                                    var v63 string = "hello"
                                    _ = v63
                                }
                                _ = v38
                            }
                            switch fi35 {
                            case -1:
                                v64 := GenS6{F0: 32767}
                                _ = v64
                            case -9223372036854775808:
                                if len(v6) > 0 {
                                    v6[0] = ""
                                }
                            case 1:
                                fmt.Println(v1, v2)
                            }
                            var v65 uint8 = (v20 << ((v33) & 63))
                            _ = fi35
                            _ = v65
                        }
                    case "\n\t":
                        var v66 int16 = 42
                        _ = v66
                    case "foo bar":
                        fmt.Println(v2, v4, v33)
                    }
                }
                if len(v5) > 0 {
                    v5[0] = -9223372036854775808
                }
                for fi67 := 0; fi67 < 1; fi67++ {
                    v68 := GenS7{F0: -32768}
                    for fi69 := 0; fi69 < 4; fi69++ {
                        var v70 int16 = -808
                        var v71 uint64 = 1
                        v70 += v70
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    _ = fi67
                    _ = v68
                }
            }
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
