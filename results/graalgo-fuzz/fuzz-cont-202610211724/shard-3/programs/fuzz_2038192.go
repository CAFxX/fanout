package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 int32; F2 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 uint; F1 uint64 }
type GenS10 struct { F0 int8 }

func main() {
    var v1 uint8 = 0
    v2 := []uint32{2}
    var v3 uint = 18446744073709551615
    var v4 int64 = 1
    var v5 int16 = -73
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = v5
        _ = fi6
        _ = v7
    }
    var v8 uint8 = v1
    var v9 bool = true
    var v10 uint8 = 234
    v1 = v10
    v4 = (v4 | (-(v4)))
    v11 := GenS1{F0: v10, F1: 18446744073709551615, F2: (("hello" + "\n\t") + "hello")}
    for fi12 := 0; fi12 < 5; fi12++ {
        v10 -= 128
        if (true && (!false)) {
            var v13 string = ""
            var v14 int64 = v4
            _ = v13
            _ = v14
        }
        var v15 float64 = 1e10
        _ = fi12
        _ = v15
    }
    if len(v2) > 0 {
        v2[0] = uint32(v8)
    }
    if v9 {
        var v16 int64 = v4
        _ = v16
    }
    var v17 string = ""
    fmt.Println(v10)
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 float64 = 0.0
        if false {
            for fi20 := 0; fi20 < 3; fi20++ {
                if v9 {
                    if false {
                        v3 += (v3 & (v3 >> ((1) & 63)))
                        _ = len(v2)
                        if true {
                            if (!(32767 > -1)) {
                                var v21 uint32 = 117
                                _ = v21
                            } else {
                                switch ("" + v17) {
                                case "\n\t":
                                    v22 := GenS1{F0: (v10 ^ (v8 << (((v3 | 100)) & 63))), F1: 784, F2: (((("hello" + (v17 + v17)) + ((v17 + v17) + "")) + v17) + "foo bar")}
                                    _ = v22
                                case "a":
                                    v3 += v3
                                default:
                                    v23 := GenS1{F0: v10, F1: 1, F2: v17}
                                    _ = v23
                                }
                            }
                            var v24 int = (fi18 - (fi18 >> (((v3 + uint(v8))) & 63)))
                            v25 := GenS2{F0: 42}
                            _ = v24
                            _ = v25
                        } else {
                            v9 = false
                        }
                    } else {
                        v5 = (v5 - (-(v5)))
                        switch fi18 {
                        case 9223372036854775807:
                            _ = len(v2)
                        }
                        v2 = append(v2, 957)
                        fmt.Println(v3)
                        v2 = append(v2, 65535)
                    }
                    fmt.Println(v5, fi18, v19)
                    v10 -= v8
                    if (!v9) {
                        var v26 float32 = 0.1
                        if (false && (!true)) {
                            fmt.Println(v5)
                            if true {
                                v19 = 0.1
                                var v27 uint32 = 869
                                _ = v27
                            } else {
                                if v9 {
                                    v28 := GenS1{F0: (v8 << ((v3) & 63)), F1: 1, F2: (v17 + "\n\t")}
                                    fmt.Println(v1)
                                    fmt.Println(v19, fi20, v5)
                                    _ = v28
                                } else {
                                    v3 += (v3 >> ((10) & 63))
                                    if v9 {
                                        fmt.Println(v10, v5)
                                        _ = len(v2)
                                        var v29 bool = v9
                                        fmt.Println(v10)
                                        var v30 uint32 = 2
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                v9 = (18446744073709551615 > 4294967295)
                                if (42 > (v10 >> ((1) & 63))) {
                                    fmt.Println(v19)
                                    if v9 {
                                        if true {
                                            _ = len(v2)
                                            var v31 bool = v9
                                            if v9 {
                                                v3 += (v3 % ((v3) | 1))
                                            }
                                            _ = v31
                                        }
                                    }
                                }
                                v1 = v10
                            }
                        } else {
                            var v32 uint8 = v8
                            fmt.Println(v3, fi18, v9)
                            v33 := GenS2{F0: 255}
                            _ = v32
                            _ = v33
                        }
                        _ = v26
                    } else {
                        if true {
                            if true {
                                if true {
                                    v34 := GenS1{F0: v8, F1: 2, F2: v17}
                                    var v35 uint32 = 0
                                    _ = v34
                                    _ = v35
                                } else {
                                    v9 = (255 > 1)
                                    var v36 float64 = v19
                                    _ = v36
                                }
                                v8 = 2
                                v4 = -1
                                switch v17 {
                                case "hello":
                                    v17 = "a"
                                default:
                                    v37 := GenS3{F0: 32767}
                                    _ = v37
                                }
                            } else {
                                v38 := GenS4{F0: fi18, F1: 930}
                                if (v9 && v9) {
                                    var v39 float64 = v19
                                    var v40 string = (v17 + v17)
                                    v5 = (v5 & v5)
                                    _ = v39
                                    _ = v40
                                }
                                v2 = append(v2, 421)
                                _ = v38
                            }
                        }
                        v17 = ""
                    }
                    var v41 float32 = 0.1
                    switch fi18 {
                    case 4294967295:
                        v1 = v8
                    case 387:
                        fmt.Println(v4, v41, v1)
                    default:
                        v42 := GenS2{F0: 32767}
                        _ = v42
                    }
                    _ = v41
                }
                _ = fi20
            }
            var v43 float64 = v19
            _ = v43
        }
        v5 = (v5 + -1)
        _ = fi18
        _ = v19
    }
    var v44 uint = (v3 >> (((v3 % ((v3) | 1))) & 63))
    if v9 {
        if v9 {
            if ((1e10 * 1e10) < -2.25) {
                v17 = "foo bar"
            }
            var v45 int64 = (v4 + 1)
            v46 := GenS5{F0: 4294967295}
            _ = v45
            _ = v46
        } else {
            v47 := GenS1{F0: v10, F1: 0, F2: v17}
            if v9 {
                for fi48 := 0; fi48 < 11; fi48++ {
                    v4 -= 9223372036854775807
                    var v49 uint64 = 18446744073709551615
                    if (1 <= 0) {
                        v10 = v1
                        v1 += v8
                        _ = len(v2)
                    } else {
                        fmt.Println(v4)
                        v50 := GenS5{F0: 2147483648}
                        _ = v50
                    }
                    _ = fi48
                    _ = v49
                }
                if v9 {
                    v1 = (v1 + (v1 + (v1 >> ((65535) & 63))))
                    switch (v17 + ("a" + "")) {
                    case "a":
                        var v51 uint64 = 5
                        _ = v51
                    case "hello":
                        var v52 uint32 = 446
                        _ = v52
                    case "\n\t":
                        if (!(!(-125 > 9223372036854775807))) {
                            var v53 bool = true
                            v3 -= (v44 ^ (^(v44)))
                            _ = v53
                        } else {
                            switch "\n\t" {
                            case "":
                                if (100 >= 0) {
                                    var v54 uint32 = 780
                                    var v55 int32 = -1
                                    _ = v54
                                    _ = v55
                                }
                            case "\n\t":
                                v10 = 242
                            default:
                                v3 = 18446744073709551615
                            }
                            if (100 != -128) {
                                if (-104 != -3) {
                                    v2 = append(v2, 747)
                                } else {
                                    fmt.Println(v4)
                                    fmt.Println(v10, v17)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                        }
                    default:
                        v10 = v8
                    }
                    for fi56 := 0; fi56 < 14; fi56++ {
                        for fi57 := 0; fi57 < 14; fi57++ {
                            if (((fi56 & (fi57 % ((int(v3)) | 1))) > -9223372036854775808) || v9) {
                                if v9 {
                                    switch (^(fi56)) {
                                    case -230:
                                        v8 = (v1 & 205)
                                    default:
                                        v4 -= v4
                                    }
                                    v58 := GenS2{F0: 0}
                                    if ((4294967295 <= 366) || v9) {
                                        v44 = (v3 * (v3 << ((v44) & 63)))
                                    } else {
                                        if v9 {
                                            v8 += v10
                                            var v59 uint32 = 333
                                            v60 := GenS4{F0: (^(fi57)), F1: (v5 ^ 1000)}
                                            fmt.Println(v17, v9, v5)
                                            if len(v2) > 0 {
                                                v2[0] = v59
                                            }
                                            _ = v59
                                            _ = v60
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            fmt.Println(v17)
                                            fmt.Println(v3)
                                            switch (fi56 + 591) {
                                            case 4294967295:
                                                if false {
                                                    switch v17 {
                                                    case "hello":
                                                        _ = len(v2)
                                                    }
                                                    v61 := GenS6{F0: (v17 + v17), F1: -1, F2: 707}
                                                    _ = v61
                                                }
                                            case -280:
                                                var v62 float32 = 0.1
                                                _ = v62
                                            case -2147483648:
                                                fmt.Println(fi56, v1, v10)
                                            }
                                            switch v17 {
                                            case "":
                                                var v63 uint16 = 32767
                                                _ = v63
                                            }
                                        }
                                        v1 = (v10 % ((v10) | 1))
                                    }
                                    _ = v58
                                } else {
                                    fmt.Println(v17)
                                    switch "foo bar" {
                                    case "\n\t":
                                        v64 := GenS7{F0: 0}
                                        _ = v64
                                    case "a":
                                        v65 := GenS8{F0: 2147483647, F1: 42}
                                        _ = v65
                                    case "":
                                        var v66 int8 = 42
                                        _ = v66
                                    }
                                    if true {
                                        if v9 {
                                            v3 = v44
                                            v44 -= v44
                                            var v67 uint16 = 1
                                            _ = v67
                                        }
                                    }
                                }
                                switch (v17 + "a") {
                                case "hello":
                                    switch (v17 + "") {
                                    case "":
                                        v9 = false
                                    case "a":
                                        var v68 float64 = 3.14159
                                        _ = v68
                                    }
                                case "\n\t":
                                    v69 := GenS4{F0: fi56, F1: (v5 - -1)}
                                    _ = v69
                                }
                                if ((v5 / (((-(v5))) | 1)) == v5) {
                                    v9 = false
                                }
                            }
                            if true {
                                v70 := GenS9{F0: v3, F1: 0}
                                _ = v70
                            }
                            _ = fi57
                        }
                        var v71 uint16 = 65535
                        _ = fi56
                        _ = v71
                    }
                }
                if v9 {
                    v5 -= (^(v5))
                    fmt.Println(v4, v10, v10)
                    var v72 int8 = 60
                    _ = v72
                }
            } else {
                if v9 {
                    var v73 int = 127
                    v5 -= (v5 | v5)
                    _ = v73
                } else {
                    fmt.Println(v3, v17, v5)
                    if (-2.25 > (0.0 - 3.14159)) {
                        fmt.Println(v17)
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (false || false) {
                            if (32767 >= 657) {
                                if len(v2) > 0 {
                                    v2[0] = 65535
                                }
                                v8 = 0
                                var v74 uint = (v3 - 2)
                                _ = v74
                            }
                            v75 := GenS10{F0: -1}
                            _ = v75
                        }
                    }
                    var v76 uint32 = 2147483647
                    _ = v76
                }
            }
            v77 := GenS6{F0: v17, F1: -2147483648, F2: 4294967295}
            _ = v47
            _ = v77
        }
        for fi78 := 0; fi78 < 8; fi78++ {
            v5 += (^(v5))
            _ = fi78
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v8
    _ = v9
}
