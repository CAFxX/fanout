package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint }

func main() {
    var v1 int8 = 1
    var v2 string = "\n\t"
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    var v5 float32 = 1e10
    v6 := []int{}
    for fi7 := 0; fi7 < 2; fi7++ {
        if (!((0 >= 4294967295) && true)) {
            v8 := GenS1{F0: 0, F1: 42, F2: 65535}
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint32 = 137
                switch "hello" {
                case "\n\t":
                    var v11 float64 = -2.25
                    _ = v11
                case "a":
                    switch "\n\t" {
                    case "\n\t":
                        var v12 float64 = float64((v5 * v5))
                        _ = v12
                    case "foo bar":
                        if false {
                            var v13 int8 = v1
                            _ = v13
                        }
                    case "a":
                        _ = len(v3)
                    }
                case "":
                    v1 += v1
                default:
                    var v14 int32 = -2147483648
                    _ = v14
                }
                _ = fi9
                _ = v10
            }
            v15 := GenS1{F0: 1, F1: -994, F2: 852}
            _ = v15
            _ = v8
        }
        var v16 bool = (137 < 55)
        _ = fi7
        _ = v16
    }
    v17 := GenS2{F0: 0, F1: 255, F2: 0}
    v18 := GenS2{F0: 1, F1: 2, F2: 2147483648}
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            delete(v3, -339)
            _ = fi20
        }
        switch fi19 {
        case -2147483648:
            for fi21 := 0; fi21 < 16; fi21++ {
                v1 -= v1
                _ = fi21
            }
        case 32767:
            v22 := GenS3{F0: (v2 + ""), F1: 2147483648}
            _ = v22
        case -9223372036854775808:
            var v23 uint = 255
            _ = v23
        default:
            switch 128 {
            case 42:
                v24 := GenS3{F0: ("hello" + v2), F1: 492}
                _ = v24
            case 2147483647:
                fmt.Println(v2, v1)
            case 1:
                v25 := GenS1{F0: 42, F1: -32768, F2: 0}
                _ = v25
            }
        }
        fmt.Println(v5, fi19)
        _ = fi19
    }
    v5 = v5
    v2 = ("" + "")
    v5 = v5
    v6 = append(v6, -2147483648)
    var v26 uint32 = 65535
    for fi27 := 0; fi27 < 3; fi27++ {
        v1 = v1
        _ = fi27
    }
    v26 += v26
    for fi28 := 0; fi28 < 10; fi28++ {
        fmt.Println(v1)
        v29 := GenS4{F0: -1}
        switch v2 {
        case "a":
            fmt.Println(v26, v1, v5)
        }
        _ = fi28
        _ = v29
    }
    v1 = 127
    v6 = append(v6, -9223372036854775808)
    switch v2 {
    case "foo bar":
        if (true || (1 < 1)) {
            v30 := GenS4{F0: -1}
            _ = v30
        } else {
            v31 := GenS2{F0: 1, F1: 85, F2: (v26 / (((v26 | v26)) | 1))}
            _ = v31
        }
    case "hello":
        if (4294967295 > 458) {
            if true {
                var v32 string = ((v2 + "a") + v2)
                v5 = float32(1.5)
                _ = v32
            } else {
                v26 = v26
                for fi33 := 0; fi33 < 12; fi33++ {
                    for fi34 := 0; fi34 < 16; fi34++ {
                        v35 := GenS4{F0: 1}
                        v36 := GenS2{F0: 32767, F1: 42, F2: v26}
                        v37 := GenS3{F0: v2, F1: 1}
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    v4[127] = (!(!(!(1 < 2))))
                    var v38 uint32 = (^(v26))
                    _ = fi33
                    _ = v38
                }
            }
            fmt.Println(v5)
        } else {
            switch "hello" {
            case "":
                fmt.Println(v26, v2, v2)
            case "hello":
                if false {
                    for fi39 := 0; fi39 < 1; fi39++ {
                        var v40 int32 = int32(v1)
                        _ = fi39
                        _ = v40
                    }
                    _ = len(v4)
                }
            case "a":
                v5 += ((v5 * v5) / -2.25)
            }
            var v41 int = -616
            var v42 int64 = -1
            _ = v41
            _ = v42
        }
    default:
        switch 1 {
        case 939:
            if true {
                switch int(v26) {
                case 128:
                    switch (v2 + "a") {
                    case "":
                        fmt.Println(v5, v5)
                    default:
                        for fi43 := 0; fi43 < 19; fi43++ {
                            switch fi43 {
                            case 32767:
                                var v44 uint32 = (v26 & v26)
                                _ = v44
                            default:
                                for fi45 := 0; fi45 < 3; fi45++ {
                                    fmt.Println(fi43, v5, v1)
                                    v46 := GenS2{F0: 2, F1: 255, F2: (v26 % ((v26) | 1))}
                                    v5 += v5
                                    v26 += v26
                                    switch (fi45 % (((fi43 / (((-(fi45))) | 1))) | 1)) {
                                    case 211:
                                        if (!(32767 < 2)) {
                                            switch -799 {
                                            case 2147483647:
                                                switch (-(fi43)) {
                                                case 195:
                                                    v1 = (v1 % ((v1) | 1))
                                                    fmt.Println(fi45)
                                                    var v47 uint8 = 1
                                                    _ = v47
                                                case -627:
                                                    v48 := GenS5{F0: 4294967295, F1: 331}
                                                    _ = v48
                                                }
                                            case -19:
                                                var v49 uint = 0
                                                _ = v49
                                            default:
                                                var v50 string = (v2 + ((v2 + "hello") + "hello"))
                                                _ = v50
                                            }
                                            fmt.Println(fi45, fi45, v5)
                                            var v51 int8 = (v1 | int8(fi45))
                                            _ = v51
                                        }
                                    case -1:
                                        fmt.Println(fi43)
                                    default:
                                        if true {
                                            switch v2 {
                                            case "":
                                                switch v2 {
                                                case "":
                                                    var v52 float64 = 0.0
                                                    _ = v52
                                                }
                                            case "\n\t":
                                                var v53 float32 = v5
                                                _ = v53
                                            case "hello":
                                                var v54 float64 = -2.25
                                                _ = v54
                                            default:
                                                switch v2 {
                                                case "foo bar":
                                                    var v55 int8 = 127
                                                    _ = v55
                                                case "\n\t":
                                                    switch fi45 {
                                                    case 32768:
                                                        _ = len(v4)
                                                    case 514:
                                                        switch "" {
                                                        case "hello":
                                                            var v56 uint = 18446744073709551615
                                                            _ = v56
                                                        case "a":
                                                            v57 := GenS1{F0: 822, F1: -32768, F2: 32767}
                                                            _ = v57
                                                        case "foo bar":
                                                            fmt.Println(v1, v2, v26)
                                                        }
                                                    case 1:
                                                        delete(v3, fi45)
                                                    }
                                                case "":
                                                    v26 += v26
                                                }
                                            }
                                            var v58 int32 = 2147483647
                                            switch (^(fi45)) {
                                            case -9223372036854775808:
                                                v59 := GenS6{F0: 255}
                                                _ = v59
                                            default:
                                                var v60 string = "\n\t"
                                                v26 = v26
                                                var v61 int16 = -32768
                                                _ = v60
                                                _ = v61
                                            }
                                            _ = v58
                                        }
                                    }
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            _ = fi43
                        }
                    }
                default:
                    var v62 int8 = (v1 ^ (v1 >> ((340) & 63)))
                    _ = v62
                }
            }
        case -1:
            var v63 uint8 = 100
            _ = v63
        case -370:
            v2 = ("foo bar" + v2)
        }
    }
    if ("" == "a") {
        if ((-2.25 < 3.14159) && false) {
            if (199 != 1) {
                var v64 string = v2
                _ = v64
            }
            v65 := v3[42]
            _ = v65
        }
    }
    if true {
        _ = len(v6)
    }
    fmt.Println(v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
