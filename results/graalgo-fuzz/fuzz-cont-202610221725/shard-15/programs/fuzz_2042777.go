package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint32; F1 string; F2 uint16 }
type GenS7 struct { F0 uint; F1 int; F2 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int; F1 uint }
type GenS10 struct { F0 uint8; F1 int64; F2 string }

func main() {
    var v1 int64 = 42
    var v2 int32 = 2147483647
    v3 := make(map[int]int)
    var v4 float64 = 0.0
    var v5 float32 = 1e10
    if true {
        v2 += 2147483647
        var v6 float64 = (v4 * 3.14159)
        _ = v6
    } else {
        for fi7 := 0; fi7 < 19; fi7++ {
            switch "hello" {
            case "foo bar":
                if (v1 >= v1) {
                    v5 -= ((v5 - 1e10) * 3.14159)
                    fmt.Println(fi7, v5, v4)
                }
            case "a":
                v8 := GenS1{F0: (v1 & v1), F1: 1, F2: 2}
                _ = v8
            case "\n\t":
                var v9 uint16 = 338
                _ = v9
            default:
                v1 += (v1 - (v1 >> ((65535) & 63)))
            }
            _ = fi7
        }
        v10 := GenS1{F0: (v1 + v1), F1: 65535, F2: 1}
        _ = v10
    }
    switch 817 {
    case -330:
        v4 += v4
    case -1:
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS2{F0: 84}
            _ = fi11
            _ = v12
        }
    case 42:
        var v13 int = -591
        _ = v13
    default:
        fmt.Println(v1, v5, v2)
    }
    switch 0 {
    case 0:
        v14 := []int{42, 128}
        _ = v14
    case -527:
        switch 32768 {
        case 32767:
            var v15 uint = 618
            _ = v15
        case 128:
            var v16 uint8 = 42
            _ = v16
        case -9223372036854775808:
            v1 -= (v1 - int64(v2))
        default:
            if false {
                fmt.Println(v5, v1)
                if false {
                    v1 -= v1
                    var v17 int32 = v2
                    _ = v17
                } else {
                    v1 = v1
                }
                switch 32767 {
                case 42:
                    for fi18 := 0; fi18 < 4; fi18++ {
                        var v19 float32 = (v5 / (0.1 - v5))
                        switch (fi18 ^ (fi18 ^ fi18)) {
                        case 1:
                            var v20 int64 = 9223372036854775807
                            _ = v20
                        case 0:
                            v21 := GenS3{F0: 991, F1: 65535, F2: -1}
                            _ = v21
                        case -2147483648:
                            var v22 int8 = -128
                            _ = v22
                        }
                        var v23 float64 = (v4 - ((v4 - ((v4 * 0.1) + (v4 * v4))) - (1e10 * 1.5)))
                        _ = fi18
                        _ = v19
                        _ = v23
                    }
                case 4294967295:
                    var v24 int16 = 1
                    _ = v24
                case -1:
                    if false {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            fmt.Println(v5, v1, v2)
                            switch "hello" {
                            case "\n\t":
                                v2 = 2147483647
                            }
                            v26 := []int{}
                            _ = fi25
                            _ = v26
                        }
                        if (false || true) {
                            _ = len(v3)
                            fmt.Println(v1, v5)
                            if false {
                                v27 := GenS2{F0: 0}
                                _ = v27
                            }
                        } else {
                            for fi28 := 0; fi28 < 4; fi28++ {
                                if ((true || (true && (true && true))) && true) {
                                    var v29 uint64 = 18446744073709551615
                                    var v30 int = 87
                                    _ = v29
                                    _ = v30
                                } else {
                                    if false {
                                        var v31 int = fi28
                                        if (v31 != -2147483648) {
                                            v32 := GenS4{F0: (v1 << ((448) & 63))}
                                            fmt.Println(v2, v4)
                                            _ = v32
                                        } else {
                                            switch ("a" + "\n\t") {
                                            case "a":
                                                fmt.Println(v31)
                                            }
                                        }
                                        var v33 bool = false
                                        _ = v31
                                        _ = v33
                                    }
                                }
                                v1 -= v1
                                if false {
                                    var v34 uint = 49
                                    _ = v34
                                }
                                _ = fi28
                            }
                            v35 := GenS3{F0: -9223372036854775808, F1: 4294967295, F2: 1}
                            _ = v35
                        }
                    } else {
                        var v36 int32 = (v2 ^ 42)
                        v37 := GenS3{F0: 0, F1: 2, F2: 42}
                        v4 = v4
                        _ = v36
                        _ = v37
                    }
                default:
                    v38 := v3[42]
                    _ = v38
                }
            } else {
                for fi39 := 0; fi39 < 3; fi39++ {
                    var v40 int64 = int64(v2)
                    fmt.Println(v40)
                    _ = fi39
                    _ = v40
                }
                v41 := []int{int(v1), 128}
                _ = v41
            }
        }
    default:
        var v42 uint16 = 65535
        _ = v42
    }
    for fi43 := 0; fi43 < 16; fi43++ {
        v44 := []int{2147483647, -9223372036854775808}
        var v45 int64 = (v1 / (((v1 - v1)) | 1))
        _ = fi43
        _ = v44
        _ = v45
    }
    delete(v3, 128)
    if (true || false) {
        v46 := GenS1{F0: (^(v1)), F1: 992, F2: 4294967295}
        switch "foo bar" {
        case "hello":
            for fi47 := 0; fi47 < 17; fi47++ {
                v4 += v4
                v4 = (v4 / float64(1.5))
                _ = fi47
            }
        case "foo bar":
            v48 := GenS2{F0: 1}
            _ = v48
        }
        v49 := GenS3{F0: -1, F1: 458, F2: -854}
        _ = v46
        _ = v49
    } else {
        for fi50 := 0; fi50 < 10; fi50++ {
            switch "foo bar" {
            case "a":
                switch ("foo bar" + "foo bar") {
                case "":
                    v51 := GenS5{F0: 65535, F1: 128}
                    _ = v51
                }
            case "foo bar":
                if (("foo bar" + ("foo bar" + "hello")) > "\n\t") {
                    v1 = v1
                } else {
                    v5 = 1.5
                }
            case "":
                switch "\n\t" {
                case "a":
                    v52 := []int{(fi50 - 32768)}
                    _ = v52
                case "":
                    v53 := GenS3{F0: (v1 * -376), F1: 42, F2: 42}
                    _ = v53
                case "foo bar":
                    v54 := v3[fi50]
                    _ = v54
                default:
                    fmt.Println(v2)
                }
            default:
                v55 := GenS6{F0: 665, F1: "", F2: 0}
                _ = v55
            }
            _ = fi50
        }
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            var v56 uint = 65535
            _ = v56
        case "a":
            v57 := GenS7{F0: 1, F1: -1, F2: -3}
            _ = v57
        }
    }
    if true {
        v1 = 9223372036854775807
        if true {
            for fi58 := 0; fi58 < 3; fi58++ {
                var v59 uint32 = 4294967295
                var v60 string = "\n\t"
                _ = fi58
                _ = v59
                _ = v60
            }
        } else {
            fmt.Println(v1)
            var v61 uint32 = 1
            var v62 bool = true
            _ = v61
            _ = v62
        }
    } else {
        var v63 uint64 = 42
        _ = v63
    }
    switch (("" + "") + "\n\t") {
    case "hello":
        delete(v3, 1)
    }
    switch "" {
    case "hello":
        var v64 string = "a"
        _ = v64
    }
    if ((!(617 < 604)) && (127 < 0)) {
        for fi65 := 0; fi65 < 12; fi65++ {
            v66 := []int{(fi65 % ((-747) | 1)), (fi65 / ((-1) | 1)), fi65}
            if false {
                v67 := GenS8{F0: (("foo bar" + ("\n\t" + "a")) + "a")}
                if false {
                    switch fi65 {
                    case 308:
                        if false {
                            _ = len(v3)
                            var v68 uint16 = 32767
                            fmt.Println(v5)
                            _ = v68
                        } else {
                            v69 := GenS1{F0: v1, F1: 232, F2: 763}
                            v70 := GenS2{F0: 59}
                            _ = v69
                            _ = v70
                        }
                    case 32767:
                        fmt.Println(v5, fi65, v2)
                    }
                }
                v3[fi65] = -363
                _ = v67
            } else {
                if (("a" + "") >= "foo bar") {
                    var v71 float32 = v5
                    v72 := GenS9{F0: fi65, F1: 2}
                    _ = v71
                    _ = v72
                } else {
                    var v73 uint32 = 232
                    _ = v73
                }
                switch "a" {
                case "\n\t":
                    v4 = float64(0.0)
                case "hello":
                    if false {
                        if len(v66) >= 2 {
                            v66 = v66[1:len(v66)]
                        }
                    }
                case "":
                    if (114 != -128) {
                        v74 := GenS4{F0: 851}
                        _ = len(v3)
                        _ = v74
                    } else {
                        var v75 uint16 = 928
                        v1 = v1
                        if (2 != 0) {
                            if false {
                                var v76 uint16 = (v75 - v75)
                                switch fi65 {
                                case 404:
                                    v77 := GenS7{F0: 18446744073709551615, F1: fi65, F2: 0}
                                    _ = v77
                                default:
                                    var v78 int32 = -1
                                    _ = v78
                                }
                                if (!false) {
                                    var v79 int = fi65
                                    for fi80 := 0; fi80 < 5; fi80++ {
                                        var v81 float64 = ((v4 / v4) + 1e10)
                                        _ = fi80
                                        _ = v81
                                    }
                                    _ = v79
                                } else {
                                    switch (fi65 * 32767) {
                                    case -494:
                                        var v82 bool = (!(false && false))
                                        _ = v82
                                    case 127:
                                        var v83 string = "\n\t"
                                        _ = v83
                                    case 969:
                                        if (((0.0 / v4) * v4) == float64(v5)) {
                                            v84 := GenS7{F0: 0, F1: fi65, F2: 42}
                                            for fi85 := 0; fi85 < 1; fi85++ {
                                                v86 := GenS10{F0: 128, F1: v1, F2: (("\n\t" + "a") + "a")}
                                                _ = fi85
                                                _ = v86
                                            }
                                            var v87 string = "hello"
                                            _ = v84
                                            _ = v87
                                        }
                                    default:
                                        v2 = -2147483648
                                    }
                                    v1 -= (v1 - 1)
                                }
                                _ = v76
                            }
                            if len(v66) > 0 {
                                v66[0] = (^(fi65))
                            }
                            switch (("foo bar" + ("foo bar" + "foo bar")) + "a") {
                            case "foo bar":
                                var v88 uint = 307
                                _ = v88
                            case "":
                                fmt.Println(v4, v2, v4)
                            case "hello":
                                fmt.Println(v2, v5)
                            default:
                                switch ("a" + ("hello" + "hello")) {
                                case "hello":
                                    v1 = v1
                                case "\n\t":
                                    v66 = append(v66, (fi65 - fi65))
                                }
                            }
                        } else {
                            fmt.Println(v5, v5, v4)
                            v2 = v2
                        }
                        _ = v75
                    }
                default:
                    v5 = v5
                }
            }
            _ = fi65
            _ = v66
        }
        var v89 uint32 = 9
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
