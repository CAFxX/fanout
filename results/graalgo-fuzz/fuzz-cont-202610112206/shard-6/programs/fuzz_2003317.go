package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 uint32; F1 int; F2 int32 }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 int; F1 uint; F2 int16 }

func main() {
    v1 := []int{-589, 4294967295, 1, -569}
    v2 := []string{(("hello" + "\n\t") + "a"), "a", ("foo bar" + "a")}
    var v3 uint8 = 127
    v4 := make(map[uint32]bool)
    var v5 float32 = 0.0
    var v6 int64 = -9223372036854775808
    for fi7 := 0; fi7 < 13; fi7++ {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint16 = 42
            v10 := GenS1{F0: -32768}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint32 = 527
        if (int(v6) < fi11) {
            _ = len(v1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            switch (("" + "\n\t") + "foo bar") {
            case "":
                v13 := GenS1{F0: 156}
                _ = v13
            case "hello":
                v14 := GenS1{F0: 857}
                _ = v14
            }
            fmt.Println(fi11, fi11)
        }
        v15 := GenS1{F0: 180}
        _ = fi11
        _ = v12
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint16 = 674
        v5 = v5
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        fmt.Println(v6)
        for fi19 := 0; fi19 < 18; fi19++ {
            v6 -= v6
            v20 := GenS2{F0: 1, F1: 702, F2: 775}
            var v21 uint64 = 2
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    v6 -= v6
    fmt.Println(v6, v6)
    v22 := GenS3{F0: 0, F1: 598}
    if false {
        if true {
            var v23 uint16 = 1
            var v24 int = 0
            _ = v23
            _ = v24
        } else {
            fmt.Println(v3)
        }
    } else {
        fmt.Println(v6)
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := GenS4{F0: 0, F1: 1}
        v1 = append(v1, 2147483647)
        if (true || (-1 == 1)) {
            var v27 int32 = 1
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    if false {
        for fi28 := 0; fi28 < 7; fi28++ {
            fmt.Println(fi28, v6)
            for fi29 := 0; fi29 < 8; fi29++ {
                switch ("foo bar" + "") {
                case "a":
                    v30 := GenS1{F0: 0}
                    _ = v30
                default:
                    var v31 uint = 65535
                    _ = v31
                }
                var v32 int = fi28
                v33 := GenS3{F0: 1, F1: 65535}
                _ = fi29
                _ = v32
                _ = v33
            }
            if (((255 > (v3 ^ v3)) && false) && false) {
                if true {
                    if true {
                        switch fi28 {
                        case 54:
                            fmt.Println(v6, v6)
                        case -650:
                            var v34 int32 = int32(fi28)
                            _ = v34
                        default:
                            if (!true) {
                                fmt.Println(fi28, v3)
                            } else {
                                var v35 int64 = (^(v6))
                                _ = v35
                            }
                        }
                    } else {
                        v5 += (v5 + 1e10)
                    }
                    for fi36 := 0; fi36 < 14; fi36++ {
                        v37 := GenS2{F0: 32767, F1: 255, F2: 44}
                        _ = fi36
                        _ = v37
                    }
                    var v38 uint16 = 2
                    _ = v38
                } else {
                    var v39 string = ("hello" + "a")
                    _ = v39
                }
                v3 = v3
            } else {
                v40 := GenS4{F0: 955, F1: 2}
                if false {
                    v6 -= -79
                    v5 = v5
                }
                _ = v40
            }
            _ = fi28
        }
        if (v6 <= v6) {
            fmt.Println(v3, v5)
            v3 = (v3 >> ((1) & 63))
            if false {
                for fi41 := 0; fi41 < 14; fi41++ {
                    v5 -= v5
                    _ = fi41
                }
            }
        }
        _ = len(v2)
    }
    var v42 int32 = 1
    var v43 float64 = (1.5 / (float64((v5 * -2.25)) + float64(1e10)))
    var v44 string = "foo bar"
    switch "\n\t" {
    case "\n\t":
        v3 = 44
    case "a":
        var v45 int16 = -350
        _ = v45
    }
    if false {
        if (1 <= 1) {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            v44 = v44
            if (407 >= -460) {
                v46 := GenS1{F0: 1}
                switch -2147483648 {
                case 4294967295:
                    var v47 uint16 = 42
                    _ = v47
                case -9223372036854775808:
                    v48 := GenS5{F0: 439}
                    _ = v48
                case 0:
                    if ((v3 << ((0) & 63)) > 127) {
                        v5 = v5
                        var v49 int16 = 32767
                        _ = v49
                    } else {
                        for fi50 := 0; fi50 < 1; fi50++ {
                            v2 = append(v2, "foo bar")
                            v51 := GenS6{F0: (fi50 * -211), F1: (fi50 * 32767)}
                            v52 := GenS4{F0: 0, F1: 710}
                            _ = fi50
                            _ = v51
                            _ = v52
                        }
                        v44 = ""
                    }
                }
                switch v44 {
                case "foo bar":
                    v5 = v5
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "a":
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        v42 += v42
                    case "a":
                        v42 = (v42 + -1)
                    case "\n\t":
                        for fi53 := 0; fi53 < 14; fi53++ {
                            v44 = "foo bar"
                            if (((false || false) || (0 == 18446744073709551615)) && (2 < 2147483647)) {
                                if (true && (940 > 32767)) {
                                    var v54 int = fi53
                                    _ = v54
                                } else {
                                    v4[2147483648] = ((127 >= -43) && (v5 > v5))
                                }
                            }
                            var v55 int64 = 42
                            _ = fi53
                            _ = v55
                        }
                    default:
                        v56 := GenS5{F0: -9223372036854775808}
                        _ = v56
                    }
                }
                _ = v46
            }
        } else {
            for fi57 := 0; fi57 < 4; fi57++ {
                delete(v4, 1)
                var v58 uint64 = 1
                _ = fi57
                _ = v58
            }
            v59 := GenS7{F0: 814, F1: 4294967295, F2: (v42 + v42)}
            var v60 uint = 383
            _ = v59
            _ = v60
        }
        var v61 int16 = 32767
        _ = v61
    }
    for fi62 := 0; fi62 < 7; fi62++ {
        var v63 int8 = 0
        if (v3 > v3) {
            var v64 uint64 = 99
            v65 := GenS3{F0: 32767, F1: 0}
            v64 -= (^(v64))
            _ = v64
            _ = v65
        } else {
            for fi66 := 0; fi66 < 14; fi66++ {
                var v67 float64 = 1e10
                _ = fi66
                _ = v67
            }
        }
        if true {
            var v68 string = v44
            if (32767 <= 391) {
                v43 -= (0.0 - (((-2.25 / 0.1) * v43) / v43))
                for fi69 := 0; fi69 < 13; fi69++ {
                    var v70 uint = 65535
                    v71 := GenS3{F0: -1, F1: 2147483648}
                    switch v68 {
                    case "hello":
                        if true {
                            v5 = 1.5
                            switch (fi69 ^ (fi69 << (((^(v70))) & 63))) {
                            case 2147483647:
                                var v72 int16 = -724
                                _ = v72
                            }
                            v73 := v4[2147483647]
                            v73 = false
                            v74 := GenS8{F0: 0, F1: v6}
                            _ = v73
                            _ = v74
                        }
                        _ = len(v4)
                    }
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            }
            var v75 uint32 = 2147483648
            _ = v68
            _ = v75
        }
        _ = fi62
        _ = v63
    }
    v76 := GenS9{F0: 6, F1: 2, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v76
}
