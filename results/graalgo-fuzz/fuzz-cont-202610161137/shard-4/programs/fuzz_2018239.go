package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := []int{150}
    var v2 string = ("\n\t" + "a")
    var v3 int8 = -1
    var v4 int8 = (v3 - -128)
    var v5 string = ""
    v6 := GenS1{F0: 1000, F1: 4294967295}
    _ = len(v1)
    switch (v2 + v5) {
    case "a":
        var v7 uint16 = 1
        _ = v7
    case "":
        var v8 float64 = -2.25
        _ = v8
    }
    var v9 int8 = v3
    var v10 float64 = 1.5
    switch v2 {
    case "hello":
        for fi11 := 0; fi11 < 16; fi11++ {
            v5 = ((v5 + v5) + v5)
            _ = fi11
        }
    default:
        v1 = append(v1, -552)
    }
    var v12 string = (("\n\t" + v2) + "foo bar")
    fmt.Println(v2)
    v5 = v12
    switch 9223372036854775807 {
    case -11:
        if (false && (true && false)) {
            switch "hello" {
            case "hello":
                if false {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    for fi13 := 0; fi13 < 12; fi13++ {
                        _ = len(v1)
                        v14 := GenS1{F0: -32768, F1: fi13}
                        _ = fi13
                        _ = v14
                    }
                    fmt.Println(v5)
                } else {
                    if (!false) {
                        v15 := GenS1{F0: -32768, F1: 679}
                        if (true || true) {
                            fmt.Println(v12, v12, v3)
                            fmt.Println(v3, v3, v9)
                        } else {
                            var v16 int8 = v3
                            switch 9223372036854775807 {
                            case 127:
                                for fi17 := 0; fi17 < 9; fi17++ {
                                    switch ("foo bar" + (("a" + "a") + v5)) {
                                    case "hello":
                                        v2 = v5
                                    case "a":
                                        switch fi17 {
                                        case 635:
                                            switch "\n\t" {
                                            case "hello":
                                                v18 := GenS1{F0: 1000, F1: fi17}
                                                _ = v18
                                            case "\n\t":
                                                if (65535 == 2) {
                                                    if true {
                                                        var v19 uint32 = 2147483647
                                                        _ = v19
                                                    } else {
                                                        if false {
                                                            fmt.Println(v3, v12)
                                                            var v20 uint8 = 0
                                                            _ = v20
                                                        }
                                                    }
                                                } else {
                                                    v16 += v4
                                                }
                                            case "foo bar":
                                                var v21 float64 = v10
                                                _ = v21
                                            default:
                                                var v22 uint = 293
                                                _ = v22
                                            }
                                        case 306:
                                            fmt.Println(fi17)
                                        default:
                                            var v23 uint = 0
                                            _ = v23
                                        }
                                    case "foo bar":
                                        v3 -= v9
                                    default:
                                        v1 = append(v1, fi17)
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v9 += (v9 >> ((18446744073709551615) & 63))
                                    _ = fi17
                                }
                            default:
                                for fi24 := 0; fi24 < 7; fi24++ {
                                    v1 = append(v1, (fi24 / ((fi24) | 1)))
                                    var v25 int16 = 42
                                    v1 = append(v1, 9223372036854775807)
                                    _ = fi24
                                    _ = v25
                                }
                            }
                            _ = v16
                        }
                        if (!false) {
                            v10 -= (v10 / (v10 / (float64((1.5 + ((3.14159 - (0.0 / 1e10)) / 0.1))) + (float64(-2.25) - ((float64(0.0) * 0.1) / v10)))))
                            if (9223372036854775807 >= int(v3)) {
                                v9 -= (v9 * (v4 << ((0) & 63)))
                                for fi26 := 0; fi26 < 18; fi26++ {
                                    fmt.Println(v4, v4)
                                    v9 = 115
                                    if true {
                                        v27 := GenS2{F0: 118, F1: 758}
                                        switch v12 {
                                        case "a":
                                            _ = len(v1)
                                        case "":
                                            var v28 uint8 = 207
                                            _ = v28
                                        }
                                        _ = v27
                                    }
                                    _ = fi26
                                }
                                v1 = append(v1, 751)
                            } else {
                                v1 = append(v1, 32768)
                                switch 128 {
                                case 32767:
                                    v9 = (v4 % (((v9 | (-(v3)))) | 1))
                                }
                            }
                            v1 = append(v1, -898)
                        } else {
                            v1 = append(v1, 210)
                        }
                        _ = v15
                    }
                    v9 -= (v9 + -1)
                }
            case "foo bar":
                v2 = v2
            }
            v2 = (v12 + (("foo bar" + "hello") + v5))
            v29 := GenS1{F0: 32767, F1: -510}
            _ = v29
        }
    }
    if true {
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            if (65535 <= 2) {
                v31 := GenS3{F0: 0, F1: 65535}
                for fi32 := 0; fi32 < 18; fi32++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if true {
                        if false {
                            _ = len(v1)
                            _ = len(v1)
                            switch fi30 {
                            case 768:
                                if false {
                                    var v33 float32 = ((1.5 - -2.25) - 3.14159)
                                    switch 32768 {
                                    case -913:
                                        v34 := GenS4{F0: 18446744073709551615}
                                        _ = v34
                                    case -9223372036854775808:
                                        switch v2 {
                                        case "hello":
                                            fmt.Println(v33, v5)
                                        case "":
                                            v2 = v2
                                        }
                                    default:
                                        v1 = append(v1, (fi32 / (((fi30 + (fi30 & fi30))) | 1)))
                                    }
                                    switch 935 {
                                    case -428:
                                        switch v5 {
                                        case "":
                                            _ = len(v1)
                                        }
                                    case 127:
                                        fmt.Println(v2)
                                    case -171:
                                        v1 = append(v1, (fi32 + fi32))
                                        v3 = (v3 >> ((65535) & 63))
                                    }
                                    _ = v33
                                } else {
                                    v5 = v5
                                }
                            }
                        } else {
                            var v35 uint32 = 2147483647
                            v1 = append(v1, fi32)
                            v4 -= v9
                            _ = v35
                        }
                        fmt.Println(v4, v2)
                        switch fi32 {
                        case -9223372036854775808:
                            v36 := GenS3{F0: -32768, F1: 42}
                            _ = v36
                        }
                        v37 := GenS3{F0: 0, F1: 325}
                        v9 += v4
                        _ = v37
                    } else {
                        if len(v1) > 0 {
                            v1[0] = fi32
                        }
                        var v38 string = v5
                        _ = v38
                    }
                    _ = fi32
                }
                _ = v31
            } else {
                var v39 int = fi30
                v1 = append(v1, (v39 << ((179) & 63)))
                _ = v39
            }
            var v40 int64 = -1
            switch fi30 {
            case 1:
                v41 := GenS1{F0: 32767, F1: (fi30 / ((-855) | 1))}
                _ = v41
            case -336:
                var v42 int = fi30
                _ = v42
            case 32768:
                fmt.Println(v2)
            default:
                if (4294967295 < 2) {
                    v5 = "a"
                    if (!((0.0 - 0.1) > 1e10)) {
                        if false {
                            if (!(3.14159 < ((0.1 + 0.1) * -2.25))) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                fmt.Println(v5, v5, v9)
                                v43 := GenS2{F0: (fi30 + fi30), F1: 2147483647}
                                var v44 int = fi30
                                _ = v43
                                _ = v44
                            }
                            var v45 float32 = 0.1
                            _ = v45
                        } else {
                            fmt.Println(v5)
                            var v46 uint8 = 2
                            for fi47 := 0; fi47 < 10; fi47++ {
                                var v48 uint8 = 127
                                v46 += 2
                                _ = fi47
                                _ = v48
                            }
                            _ = v46
                        }
                    } else {
                        v9 = int8(v40)
                        v49 := GenS3{F0: 0, F1: 808}
                        _ = v49
                    }
                    if false {
                        v9 -= 42
                        if true {
                            switch v2 {
                            case "hello":
                                v3 += v3
                            case "foo bar":
                                _ = len(v1)
                            case "a":
                                v1 = append(v1, fi30)
                            }
                            fmt.Println(v12, v12, v10)
                            switch (fi30 + fi30) {
                            case 728:
                                for fi50 := 0; fi50 < 5; fi50++ {
                                    fmt.Println(fi30)
                                    _ = fi50
                                }
                            default:
                                if false {
                                    v40 -= (^(v40))
                                }
                            }
                        } else {
                            var v51 float32 = 1e10
                            v52 := GenS3{F0: -32768, F1: 65535}
                            _ = v51
                            _ = v52
                        }
                    } else {
                        if (v10 == (((float64(0.1) / v10) / ((v10 - 3.14159) / v10)) - (v10 + v10))) {
                            if true {
                                for fi53 := 0; fi53 < 18; fi53++ {
                                    v40 = v40
                                    _ = fi53
                                }
                                if true {
                                    fmt.Println(fi30, v12)
                                    v54 := GenS3{F0: -552, F1: 971}
                                    _ = v54
                                }
                            }
                            for fi55 := 0; fi55 < 13; fi55++ {
                                if true {
                                    v4 += v9
                                    if len(v1) > 0 {
                                        v1[0] = fi30
                                    }
                                    var v56 float32 = 3.14159
                                    _ = v56
                                }
                                v1 = append(v1, (fi55 % ((fi55) | 1)))
                                _ = fi55
                            }
                        } else {
                            if ((v9 != (v9 / ((v9) | 1))) || (true || (!false))) {
                                var v57 uint16 = 844
                                var v58 uint32 = 954
                                _ = v57
                                _ = v58
                            }
                            v10 += 1.5
                            if (399 < 18446744073709551615) {
                                for fi59 := 0; fi59 < 18; fi59++ {
                                    fmt.Println(v40, fi30, v4)
                                    v9 = v9
                                    v2 = ((v2 + "a") + v5)
                                    _ = fi59
                                }
                                var v60 uint64 = 0
                                _ = v60
                            } else {
                                var v61 int16 = 132
                                var v62 string = (v5 + v2)
                                _ = v61
                                _ = v62
                            }
                        }
                        if true {
                            v12 = v2
                        } else {
                            _ = len(v1)
                            switch (v5 + v12) {
                            case "":
                                v5 = ""
                            default:
                                v3 -= (v4 >> ((255) & 63))
                            }
                            var v63 float64 = (1e10 + 1e10)
                            _ = v63
                        }
                    }
                }
            }
            _ = fi30
            _ = v40
        }
    } else {
        if len(v1) > 0 {
            v1[0] = 9223372036854775807
        }
    }
    var v64 int8 = v9
    var v65 uint64 = 42
    v12 = "\n\t"
    var v66 uint8 = 0
    var v67 uint8 = (v66 | 137)
    v64 -= (v3 + v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v9
}
