package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 uint8 = 123
    var v2 int32 = 829
    var v3 string = "foo bar"
    var v4 int16 = -580
    v3 = v3
    switch v3 {
    case "foo bar":
        v2 = -304
    case "a":
        switch v3 {
        case "\n\t":
            v3 = "hello"
        case "hello":
            fmt.Println(v3, v2, v4)
        default:
            var v5 int = 1
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(v1)
        _ = fi6
    }
    v7 := []int{2147483647, 42}
    if (3.14159 >= 0.1) {
        switch v3 {
        case "a":
            for fi8 := 0; fi8 < 3; fi8++ {
                if true {
                    switch (("a" + "hello") + ("hello" + "foo bar")) {
                    case "hello":
                        var v9 int8 = 42
                        _ = v9
                    case "foo bar":
                        v3 = ""
                    }
                } else {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        v3 = (v3 + (v3 + "\n\t"))
                        v2 = (v2 / (((v2 + v2)) | 1))
                        switch fi10 {
                        case 127:
                            if false {
                                switch (fi8 / (((fi10 | 912)) | 1)) {
                                case -9223372036854775808:
                                    v4 = 1000
                                case -288:
                                    fmt.Println(v3, v1, v2)
                                default:
                                    var v11 int64 = -363
                                    _ = v11
                                }
                            }
                            v12 := GenS1{F0: 42}
                            if (true && false) {
                                fmt.Println(v1, v1, v4)
                                var v13 int = (fi8 & (fi8 & fi10))
                                _ = v13
                            }
                            _ = v12
                        case 2147483647:
                            v14 := GenS2{F0: true, F1: 42}
                            _ = v14
                        }
                        _ = fi10
                    }
                }
                v1 += v1
                var v15 float64 = 3.14159
                _ = fi8
                _ = v15
            }
        }
        v16 := GenS3{F0: 2}
        _ = v16
    } else {
        if (0 < 2147483648) {
            var v17 float64 = (((-2.25 / 3.14159) / 1e10) * -2.25)
            if (true && true) {
                for fi18 := 0; fi18 < 10; fi18++ {
                    switch (fi18 & (fi18 & (^(fi18)))) {
                    case -1:
                        var v19 uint = 65535
                        _ = v19
                    default:
                        switch "hello" {
                        case "hello":
                            var v20 float32 = 1e10
                            _ = v20
                        case "":
                            v4 -= 1000
                        }
                    }
                    if false {
                        if true {
                            v2 -= 0
                            fmt.Println(v4, v1, v3)
                        } else {
                            v21 := GenS4{F0: (fi18 - (^(fi18))), F1: 42}
                            _ = v21
                        }
                        v2 += -357
                    } else {
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 int64 = int64(v1)
                            v24 := GenS3{F0: 4294967295}
                            var v25 float64 = float64(0.1)
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        var v26 float32 = 1e10
                        fmt.Println(v1, v26)
                        _ = v26
                    }
                    if (v2 <= v2) {
                        v27 := GenS4{F0: fi18, F1: v4}
                        switch v3 {
                        case "foo bar":
                            for fi28 := 0; fi28 < 2; fi28++ {
                                switch (fi18 & fi18) {
                                case 0:
                                    if (739 == 42) {
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        var v29 int64 = 9223372036854775807
                                        var v30 int8 = 42
                                        var v31 int8 = v30
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    } else {
                                        var v32 uint8 = v1
                                        fmt.Println(v32, fi18)
                                        v2 = v2
                                        _ = v32
                                    }
                                case 4294967295:
                                    fmt.Println(v17, fi18)
                                }
                                fmt.Println(v1)
                                _ = fi28
                            }
                        case "a":
                            v4 += v4
                        case "\n\t":
                            if (v4 > (v4 ^ v4)) {
                                v33 := GenS3{F0: 4294967295}
                                _ = v33
                            } else {
                                if len(v7) > 0 {
                                    v7[0] = fi18
                                }
                                if false {
                                    switch ("\n\t" + v3) {
                                    case "hello":
                                        for fi34 := 0; fi34 < 19; fi34++ {
                                            var v35 uint32 = 1
                                            v36 := GenS1{F0: (v4 | v4)}
                                            var v37 uint16 = 931
                                            v35 -= (v35 << ((0) & 63))
                                            _ = fi34
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                    }
                                    if false {
                                        var v38 bool = (0.1 <= float64(float32(((v17 + 0.1) * 3.14159))))
                                        v17 += float64(-2.25)
                                        var v39 uint = 65535
                                        _ = v38
                                        _ = v39
                                    }
                                    fmt.Println(v1, v3)
                                }
                            }
                        }
                        v40 := GenS3{F0: 134}
                        _ = v27
                        _ = v40
                    } else {
                        v1 -= v1
                    }
                    _ = fi18
                }
                if true {
                    v1 -= (^(v1))
                    v7 = append(v7, 32767)
                    v41 := GenS5{F0: 1}
                    _ = v41
                } else {
                    v4 = v4
                }
                for fi42 := 0; fi42 < 1; fi42++ {
                    for fi43 := 0; fi43 < 18; fi43++ {
                        v4 = (v4 * (v4 ^ (v4 ^ (v4 * -1))))
                        _ = fi43
                    }
                    v44 := GenS1{F0: v4}
                    _ = fi42
                    _ = v44
                }
            } else {
                var v45 string = ((v3 + v3) + v3)
                _ = v45
            }
            if len(v7) > 0 {
                v7[0] = 804
            }
            _ = v17
        }
        for fi46 := 0; fi46 < 13; fi46++ {
            if (((true && (761 > 1)) || false) || false) {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                switch v3 {
                case "a":
                    if (!((1e10 - 0.1) == float64(3.14159))) {
                        v2 = v2
                        v4 = (v4 + (v4 >> ((65535) & 63)))
                        v4 += 0
                    }
                case "":
                    if true {
                        v1 -= (v1 - (^(v1)))
                        var v47 int32 = v2
                        _ = v47
                    }
                }
            } else {
                var v48 uint32 = 0
                _ = v48
            }
            switch "a" {
            case "foo bar":
                var v49 int64 = 42
                _ = v49
            case "hello":
                if true {
                    var v50 uint16 = uint16(v1)
                    var v51 int = (fi46 | fi46)
                    _ = v50
                    _ = v51
                }
            }
            switch (^(fi46)) {
            case 0:
                var v52 int = (fi46 / ((fi46) | 1))
                _ = v52
            }
            _ = fi46
        }
        var v53 string = v3
        _ = v53
    }
    if false {
        for fi54 := 0; fi54 < 10; fi54++ {
            v7 = append(v7, (fi54 / (((fi54 ^ 32767)) | 1)))
            v55 := GenS5{F0: 2}
            _ = fi54
            _ = v55
        }
        var v56 float64 = 0.1
        if ((v2 >> ((454) & 63)) <= (v2 << ((42) & 63))) {
            for fi57 := 0; fi57 < 20; fi57++ {
                var v58 float32 = 0.0
                for fi59 := 0; fi59 < 15; fi59++ {
                    v60 := GenS4{F0: (fi57 & fi59), F1: -1}
                    if ((-2.25 - 0.0) < 0.0) {
                        var v61 uint8 = v1
                        fmt.Println(fi57, v61)
                        _ = v61
                    } else {
                        if len(v7) > 0 {
                            v7[0] = fi57
                        }
                    }
                    _ = fi59
                    _ = v60
                }
                var v62 int32 = 2147483647
                _ = fi57
                _ = v58
                _ = v62
            }
            switch v3 {
            case "\n\t":
                fmt.Println(v4, v2)
            case "hello":
                _ = len(v7)
            case "foo bar":
                var v63 uint8 = v1
                _ = v63
            default:
                v64 := GenS5{F0: 2}
                _ = v64
            }
            if false {
                v4 -= v4
                if true {
                    v2 = v2
                    fmt.Println(v3, v4, v1)
                    switch -451 {
                    case -9223372036854775808:
                        fmt.Println(v3, v3, v1)
                    case 42:
                        var v65 float32 = float32(float64(0.1))
                        _ = v65
                    case 4294967295:
                        if ((v1 / ((v1) | 1)) <= 40) {
                            if (((((float32(1.5) + (1.5 - ((1.5 / float32(-2.25)) - -2.25))) + (((3.14159 + (3.14159 * 3.14159)) - (1e10 * 1e10)) + 0.0)) - 1e10) < 0.0) && ((((v4 / ((v4) | 1)) < (v4 - v4)) || false) && (!true))) {
                                v2 += (v2 >> ((0) & 63))
                                v3 = ""
                            }
                        }
                    default:
                        var v66 int32 = 42
                        _ = v66
                    }
                } else {
                    fmt.Println(v3, v1)
                }
                for fi67 := 0; fi67 < 13; fi67++ {
                    fmt.Println(v4, v4)
                    fmt.Println(v4, v56, v3)
                    v2 -= v2
                    _ = fi67
                }
            }
        }
        _ = v56
    }
    v3 = v3
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
