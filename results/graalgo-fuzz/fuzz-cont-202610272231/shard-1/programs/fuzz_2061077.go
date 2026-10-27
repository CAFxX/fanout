package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int; F1 int8; F2 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64; F1 uint; F2 uint16 }

func main() {
    v1 := []string{("a" + (("a" + "\n\t") + ("hello" + ("foo bar" + ("a" + "foo bar"))))), "a"}
    v2 := []int{0, 127, 32767, 2147483647}
    v3 := make(map[int]uint32)
    var v4 int16 = -32768
    var v5 bool = true
    var v6 int = 128
    if ((((42 != 0) || (v5 || true)) || (0 != 145)) || ((!v5) && (v5 || v5))) {
        for fi7 := 0; fi7 < 9; fi7++ {
            switch (fi7 + v6) {
            case -635:
                if (!v5) {
                    for fi8 := 0; fi8 < 11; fi8++ {
                        v9 := GenS1{F0: 128, F1: 1}
                        _ = fi8
                        _ = v9
                    }
                }
            case -885:
                fmt.Println(v5, v5)
            }
            v4 += 42
            _ = fi7
        }
        v5 = (!v5)
        if v5 {
            var v10 uint8 = 255
            if v5 {
                v4 -= (v4 & (v4 | (v4 | v4)))
                if v5 {
                    v11 := GenS1{F0: 1, F1: -1}
                    v6 -= v6
                    v12 := GenS1{F0: (v10 & (v10 ^ v10)), F1: 1}
                    _ = v11
                    _ = v12
                } else {
                    v4 = 42
                    var v13 int32 = 0
                    v14 := GenS1{F0: (v10 >> ((1) & 63)), F1: 121}
                    _ = v13
                    _ = v14
                }
            }
            _ = v10
        } else {
            v15 := GenS1{F0: 158, F1: 42}
            _ = v15
        }
    }
    v5 = (659 < 496)
    if (!(!v5)) {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int = v6
            var v18 uint64 = 741
            var v19 int8 = -16
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
    } else {
        v20 := GenS2{F0: (!v5)}
        _ = v20
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS1{F0: 2, F1: 42}
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int64 = -605
            var v25 uint = 2
            fmt.Println(fi21)
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    v4 = (-(v4))
    var v26 int32 = -382
    switch "foo bar" {
    case "foo bar":
        v2 = append(v2, (v6 - -1))
    case "\n\t":
        v27 := GenS2{F0: v5}
        _ = v27
    case "a":
        v26 -= v26
    default:
        v28 := GenS3{F0: (v4 / (((v4 & v4)) | 1)), F1: 255}
        _ = v28
    }
    if ((!(!false)) || (-2.25 < (0.1 + ((0.0 - 3.14159) / 3.14159)))) {
        var v29 bool = (65535 >= 2147483648)
        _ = v29
    }
    if false {
        if v5 {
            var v30 float32 = 1.5
            switch ("\n\t" + "") {
            case "foo bar":
                if v5 {
                    var v31 string = ""
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 7; fi32++ {
                        switch "foo bar" {
                        case "\n\t":
                            var v33 uint32 = 2
                            _ = v33
                        case "hello":
                            switch (v6 ^ (v6 << ((65535) & 63))) {
                            case -989:
                                v34 := GenS4{F0: 1, F1: ("hello" + "hello")}
                                _ = v34
                            case -2147483648:
                                fmt.Println(v26, v5, v5)
                            case 4294967295:
                                for fi35 := 0; fi35 < 7; fi35++ {
                                    fmt.Println(v4)
                                    var v36 uint = 4294967295
                                    v5 = v5
                                    _ = fi35
                                    _ = v36
                                }
                            }
                        default:
                            var v37 uint16 = 65535
                            _ = v37
                        }
                        v6 = fi32
                        _ = fi32
                    }
                    switch ("a" + "a") {
                    case "":
                        var v38 uint32 = 2147483647
                        _ = v38
                    case "\n\t":
                        if len(v2) > 0 {
                            v2[0] = (v6 % ((v6) | 1))
                        }
                    }
                }
            case "hello":
                var v39 int32 = -68
                _ = v39
            }
            _ = v30
        } else {
            v40 := GenS5{F0: false, F1: v6}
            var v41 int64 = 1
            _ = v40
            _ = v41
        }
        var v42 float32 = (-2.25 * 0.0)
        for fi43 := 0; fi43 < 13; fi43++ {
            if (true || (117 < 42)) {
                var v44 uint8 = 127
                v26 += -2147483648
                _ = v44
            }
            var v45 int64 = 611
            var v46 int64 = 42
            _ = fi43
            _ = v45
            _ = v46
        }
        _ = v42
    } else {
        v5 = (!v5)
        v47 := GenS2{F0: (!(511 == 1))}
        _ = v47
    }
    switch (("hello" + "foo bar") + "") {
    case "hello":
        v48 := GenS6{F0: (v6 - 9223372036854775807), F1: 0, F2: v6}
        _ = v48
    case "":
        switch (v6 - (v6 & (v6 << ((18446744073709551615) & 63)))) {
        case 9223372036854775807:
            var v49 bool = ((1.5 <= -2.25) && true)
            _ = v49
        case -9223372036854775808:
            if (v26 < (v26 ^ (v26 / ((521) | 1)))) {
                fmt.Println(v26, v5, v4)
                v50 := GenS7{F0: 42}
                _ = v50
            } else {
                for fi51 := 0; fi51 < 4; fi51++ {
                    var v52 uint8 = 1
                    for fi53 := 0; fi53 < 4; fi53++ {
                        fmt.Println(v4, v52)
                        var v54 int = (fi51 / ((fi53) | 1))
                        _ = fi53
                        _ = v54
                    }
                    var v55 int = v6
                    _ = fi51
                    _ = v52
                    _ = v55
                }
                var v56 uint32 = 2147483647
                _ = v56
            }
        }
    case "a":
        switch "" {
        case "":
            v6 += v6
        case "hello":
            v6 = -329
        case "a":
            for fi57 := 0; fi57 < 4; fi57++ {
                fmt.Println(fi57, fi57)
                for fi58 := 0; fi58 < 8; fi58++ {
                    var v59 float32 = (0.0 + 0.1)
                    _ = fi58
                    _ = v59
                }
                _ = fi57
            }
        }
    default:
        var v60 float64 = -2.25
        _ = v60
    }
    v61 := GenS8{F0: 18446744073709551615, F1: 255, F2: 304}
    var v62 uint = 411
    fmt.Println(v62)
    delete(v3, v6)
    v63 := GenS7{F0: 0}
    if len(v2) > 0 {
        v2[0] = 128
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        var v65 bool = v5
        var v66 bool = (v65 && v65)
        _ = fi64
        _ = v65
        _ = v66
    }
    var v67 int = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v63
    _ = v67
}
