package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 uint64; F1 uint16; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{}
    var v3 bool = (-651 != -801)
    v4 := make(map[uint32]int)
    var v5 int = 4294967295
    v6 := GenS1{F0: 677}
    var v7 int8 = 0
    v3 = ((v3 || (2 >= 0)) || ((0.0 / 1.0) != 3.14159))
    switch "a" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = v5
        }
    case "foo bar":
        _ = len(v1)
    case "\n\t":
        fmt.Println(v7, v7)
    }
    if true {
        fmt.Println(v3, v5)
    } else {
        var v8 string = "a"
        v2 = append(v2, 1)
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS2{F0: fi9, F1: 2}
        var v11 int = fi9
        _ = fi9
        _ = v10
        _ = v11
    }
    delete(v1, 2)
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int64 = 9223372036854775807
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 9
    if ((1.5 != (3.14159 - 3.14159)) && true) {
        if (261 > 934) {
            var v15 uint32 = 0
            fmt.Println(v15, v14)
            _ = v15
        } else {
            v5 = v5
            v5 -= v5
            var v16 uint = 1
            _ = v16
        }
        _ = len(v2)
    } else {
        v17 := GenS2{F0: (^(v5)), F1: 572}
        v14 = v14
        _ = len(v2)
        _ = v17
    }
    if false {
        switch "foo bar" {
        case "hello":
            fmt.Println(v14, v3)
        default:
            if (2 > 255) {
                var v18 int8 = v7
                _ = v18
            }
        }
        var v19 uint8 = 127
        v3 = true
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        switch "" {
        case "foo bar":
            fmt.Println(fi20, v14, fi20)
        case "a":
            v21 := GenS3{F0: 4294967295, F1: v3, F2: (!(!true))}
            _ = v21
        }
        _ = fi20
    }
    v22 := v1[65535]
    switch v5 {
    case 1:
        v23 := GenS1{F0: v14}
        _ = v23
    case -585:
        switch v5 {
        case 2147483647:
            for fi24 := 0; fi24 < 3; fi24++ {
                var v25 int32 = 0
                v3 = v3
                var v26 int8 = (-(v7))
                _ = fi24
                _ = v25
                _ = v26
            }
        case 207:
            for fi27 := 0; fi27 < 4; fi27++ {
                for fi28 := 0; fi28 < 7; fi28++ {
                    v5 += 2147483647
                    var v29 uint32 = 2147483647
                    _ = len(v4)
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
        default:
            fmt.Println(v7)
        }
    default:
        v30 := GenS3{F0: 255, F1: (1e10 == 1e10), F2: v22}
        _ = v30
    }
    v22 = (((!v22) && false) && false)
    for fi31 := 0; fi31 < 6; fi31++ {
        switch v5 {
        case 4294967295:
            fmt.Println(v3, v22, v3)
        case 0:
            if ((v14 ^ 32767) == v14) {
                _ = len(v4)
                var v32 bool = (!((0.1 != float64(float32(1e10))) || v3))
                var v33 uint8 = 1
                _ = v32
                _ = v33
            } else {
                switch (fi31 + v5) {
                case 127:
                    if (false && (2147483647 != 2147483647)) {
                        switch "hello" {
                        case "hello":
                            var v34 uint64 = 164
                            _ = v34
                        default:
                            v14 -= 65535
                        }
                    } else {
                        v3 = v22
                    }
                default:
                    if v22 {
                        for fi35 := 0; fi35 < 8; fi35++ {
                            v14 += v14
                            var v36 int64 = 9223372036854775807
                            v5 -= (fi35 | fi31)
                            v3 = v22
                            _ = fi35
                            _ = v36
                        }
                        switch ("hello" + "a") {
                        case "":
                            v37 := GenS2{F0: v5, F1: 2147483648}
                            _ = v37
                        case "a":
                            if (!true) {
                                v5 -= v5
                                if (1 >= 1) {
                                    switch "hello" {
                                    case "a":
                                        v38 := GenS3{F0: 18446744073709551615, F1: v3, F2: true}
                                        _ = v38
                                    case "hello":
                                        v4[952] = v5
                                    case "":
                                        var v39 int32 = int32(v7)
                                        _ = v39
                                    }
                                    if (2 <= 1) {
                                        if (42 <= 2) {
                                            var v40 int8 = (v7 << ((65535) & 63))
                                            switch 32768 {
                                            case 128:
                                                v3 = ("hello" >= "\n\t")
                                            case 2147483647:
                                                fmt.Println(v3)
                                            case 872:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            _ = v40
                                        }
                                        for fi41 := 0; fi41 < 13; fi41++ {
                                            v3 = (!(v22 && (!((2 != 42) && v3))))
                                            fmt.Println(v22)
                                            _ = fi41
                                        }
                                    }
                                } else {
                                    if false {
                                        var v42 int64 = -241
                                        var v43 bool = (9223372036854775807 < 127)
                                        _ = v42
                                        _ = v43
                                    }
                                    switch ("hello" + ((("" + ("hello" + "\n\t")) + "hello") + "\n\t")) {
                                    case "foo bar":
                                        var v44 int64 = 1
                                        _ = v44
                                    case "a":
                                        var v45 float32 = 1e10
                                        _ = v45
                                    case "":
                                        fmt.Println(v5, v14, v7)
                                    }
                                }
                                fmt.Println(v7, v14)
                            } else {
                                v1[2147483648] = (42 < 42)
                            }
                        case "hello":
                            v46 := GenS4{F0: v7, F1: 2147483647}
                            _ = v46
                        }
                        fmt.Println(fi31)
                    } else {
                        v47 := GenS2{F0: v5, F1: 42}
                        v3 = v22
                        _ = v47
                    }
                }
                fmt.Println(v7, v22)
                var v48 float32 = 0.0
                _ = v48
            }
        }
        v49 := GenS1{F0: (v14 - (v14 / (((v14 / ((v14) | 1))) | 1)))}
        v50 := v4[2147483648]
        _ = fi31
        _ = v49
        _ = v50
    }
    var v51 int32 = 0
    if ("a" < ("a" + "\n\t")) {
        for fi52 := 0; fi52 < 6; fi52++ {
            var v53 uint64 = 42
            for fi54 := 0; fi54 < 1; fi54++ {
                if v3 {
                    var v55 uint64 = (v53 - (v53 & (v53 * v53)))
                    var v56 uint64 = v55
                    _ = len(v2)
                    _ = v55
                    _ = v56
                } else {
                    var v57 int64 = -9223372036854775808
                    _ = v57
                }
                var v58 string = "foo bar"
                _ = fi54
                _ = v58
            }
            if ((v53 + v53) != 4294967295) {
                for fi59 := 0; fi59 < 13; fi59++ {
                    fmt.Println(v53, v22, v7)
                    _ = fi59
                }
                for fi60 := 0; fi60 < 16; fi60++ {
                    if false {
                        v2 = append(v2, (fi52 & v5))
                    } else {
                        v5 += 1
                        switch (fi52 ^ (fi60 ^ (fi60 | (v5 - (fi60 >> ((525) & 63)))))) {
                        case 9223372036854775807:
                            var v61 float64 = (3.14159 + 3.14159)
                            v7 = v7
                            v62 := GenS5{F0: 4294967295, F1: (v3 || v3)}
                            _ = v61
                            _ = v62
                        case 32768:
                            var v63 uint32 = 42
                            _ = v63
                        default:
                            switch (v5 / ((42) | 1)) {
                            case 127:
                                v3 = v3
                            case 128:
                                v51 = v51
                            }
                        }
                        switch (fi52 % ((int(v53)) | 1)) {
                        case 127:
                            v22 = false
                            v64 := GenS6{F0: 42, F1: 346}
                            _ = v64
                        case -9223372036854775808:
                            var v65 int32 = (v51 % ((-2147483648) | 1))
                            _ = v65
                        default:
                            v5 -= fi52
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v66 int16 = 42
                    _ = fi60
                    _ = v66
                }
                v1[884] = ((v3 && v3) || (true && (42 == 370)))
            }
            _ = fi52
            _ = v53
        }
        var v67 int64 = 42
        v68 := GenS7{F0: 2, F1: v14, F2: 2}
        _ = v67
        _ = v68
    } else {
        if len(v2) > 0 {
            v2[0] = (v5 & int(v51))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
}
