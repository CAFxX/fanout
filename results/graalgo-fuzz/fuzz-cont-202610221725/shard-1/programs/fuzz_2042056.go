package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint16; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 uint8; F2 bool }
type GenS7 struct { F0 int }

func main() {
    var v1 int8 = 127
    var v2 int8 = (v1 % ((v1) | 1))
    var v3 uint32 = 4294967295
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: (v1 & (v1 + (^(v2)))), F1: true}
    v8 := GenS2{F0: 255, F1: 184}
    v9 := []int{9223372036854775807, -1, 9223372036854775807}
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            switch (-(fi11)) {
            case 32768:
                var v12 uint8 = 0
                _ = v12
            }
            v13 := v5[(v3 ^ (v3 % ((v6) | 1)))]
            v2 -= v2
            _ = fi11
            _ = v13
        }
        var v14 int = -1
        _ = fi10
        _ = v14
    }
    _ = len(v5)
    v15 := GenS2{F0: 120, F1: 22}
    for fi16 := 0; fi16 < 6; fi16++ {
        if true {
            v17 := GenS2{F0: 2, F1: 128}
            v2 += (v2 ^ (v2 % (((v1 >> ((0) & 63))) | 1)))
            switch (fi16 >> ((0) & 63)) {
            case 128:
                var v18 uint = 65535
                _ = v18
            default:
                if false {
                    v9 = append(v9, (-(fi16)))
                } else {
                    switch "hello" {
                    case "\n\t":
                        v2 = -128
                    case "hello":
                        fmt.Println(v1)
                    default:
                        if (0.1 > 1.5) {
                            if false {
                                if false {
                                    v19 := GenS3{F0: ((true || ((!(!(1e10 < (1.5 - (0.0 / 0.1))))) || ((0 == 1) && ((v1 + v1) == int8(fi16))))) && false)}
                                    if true {
                                        v20 := GenS3{F0: (true && false)}
                                        fmt.Println(v1)
                                        _ = v20
                                    }
                                    _ = v19
                                }
                                var v21 int64 = -9223372036854775808
                                switch fi16 {
                                case -2147483648:
                                    for fi22 := 0; fi22 < 14; fi22++ {
                                        var v23 int16 = 0
                                        fmt.Println(v23)
                                        _ = fi22
                                        _ = v23
                                    }
                                case 949:
                                    for fi24 := 0; fi24 < 15; fi24++ {
                                        v25 := GenS2{F0: 1, F1: 1}
                                        var v26 float64 = 3.14159
                                        if ("foo bar" != "hello") {
                                            v2 = v1
                                            if true {
                                                v27 := GenS4{F0: v1, F1: 42, F2: -2147483648}
                                                _ = v27
                                            } else {
                                                v28 := GenS1{F0: 100, F1: (!true)}
                                                var v29 int8 = 32
                                                var v30 uint32 = 774
                                                v4[fi24] = (!(!("" > "\n\t")))
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                        } else {
                                            var v31 uint = 1
                                            switch fi16 {
                                            case -445:
                                                if false {
                                                    v32 := GenS4{F0: v2, F1: 42, F2: -1}
                                                    switch fi24 {
                                                    case 32767:
                                                        v26 = v26
                                                    case 32768:
                                                        if true {
                                                            v26 += v26
                                                            fmt.Println(v3, fi16, fi16)
                                                        }
                                                        var v33 int16 = -700
                                                        var v34 uint = 912
                                                        _ = v33
                                                        _ = v34
                                                    case 2147483647:
                                                        v31 = 39
                                                    default:
                                                        var v35 uint32 = 189
                                                        _ = v35
                                                    }
                                                    _ = v32
                                                } else {
                                                    v5[v6] = true
                                                }
                                            case 1:
                                                switch "a" {
                                                case "hello":
                                                    v36 := GenS5{F0: 805}
                                                    _ = v36
                                                default:
                                                    if (65535 != (v31 - uint(v3))) {
                                                        switch 942 {
                                                        case 0:
                                                            fmt.Println(fi16, v26, v31)
                                                        }
                                                        switch "foo bar" {
                                                        case "":
                                                            switch -632 {
                                                            case 2147483647:
                                                                var v37 int8 = (v2 / (((v1 >> ((v31) & 63))) | 1))
                                                                _ = v37
                                                            default:
                                                                fmt.Println(fi16)
                                                            }
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                if len(v9) > 0 {
                                                                    v9[0] = (fi16 / ((int(v21)) | 1))
                                                                }
                                                            default:
                                                                fmt.Println(fi24, fi16, fi16)
                                                            }
                                                        }
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                    }
                                                    var v38 float32 = 1.5
                                                    fmt.Println(v26, v38)
                                                    v21 += v21
                                                    _ = v38
                                                }
                                            }
                                            v6 -= 42
                                            fmt.Println(v21, v6, v21)
                                            if (0.0 == ((float64(-2.25) - (v26 - v26)) + 1e10)) {
                                                var v39 uint8 = uint8(v31)
                                                var v40 string = "hello"
                                                switch 32767 {
                                                case 2147483647:
                                                    delete(v5, 2147483648)
                                                }
                                                _ = v39
                                                _ = v40
                                            }
                                            _ = v31
                                        }
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                case 127:
                                    if len(v9) > 0 {
                                        v9[0] = fi16
                                    }
                                }
                                _ = v21
                            } else {
                                for fi41 := 0; fi41 < 13; fi41++ {
                                    var v42 int32 = 2147483647
                                    _ = fi41
                                    _ = v42
                                }
                                var v43 int = (fi16 ^ fi16)
                                _ = v43
                            }
                            delete(v4, fi16)
                            if true {
                                for fi44 := 0; fi44 < 4; fi44++ {
                                    switch (("a" + ("a" + "")) + "hello") {
                                    case "hello":
                                        v45 := GenS6{F0: 255, F1: 235, F2: true}
                                        _ = v45
                                    case "a":
                                        if ((-(v2)) > 100) {
                                            v1 -= v2
                                            v9 = append(v9, (fi44 ^ (fi44 << ((2) & 63))))
                                            var v46 int32 = 1
                                            fmt.Println(v1, fi16)
                                            v47 := GenS7{F0: 0}
                                            _ = v46
                                            _ = v47
                                        } else {
                                            v6 = (v3 >> ((2) & 63))
                                            var v48 int32 = -2147483648
                                            v3 = (v3 + (v3 % ((v3) | 1)))
                                            _ = v48
                                        }
                                    default:
                                        delete(v5, (v6 + v3))
                                    }
                                    fmt.Println(fi16)
                                    if false {
                                        var v49 int32 = 0
                                        if (v49 < 1) {
                                            var v50 string = ("" + (("\n\t" + ("a" + "\n\t")) + ""))
                                            v50 = ("foo bar" + v50)
                                            v51 := GenS2{F0: 42, F1: 127}
                                            _ = v50
                                            _ = v51
                                        } else {
                                            var v52 uint8 = 146
                                            _ = v52
                                        }
                                        _ = v49
                                    } else {
                                        var v53 string = ""
                                        _ = v53
                                    }
                                    delete(v4, fi44)
                                    v1 -= 42
                                    _ = fi44
                                }
                                var v54 string = ""
                                _ = v54
                            } else {
                                fmt.Println(v6, v6, fi16)
                            }
                        }
                    }
                }
            }
            _ = v17
        }
        fmt.Println(v3)
        _ = fi16
    }
    var v55 uint8 = 255
    if (!false) {
        v3 = (v3 & (v3 * v6))
        var v56 int8 = 127
        var v57 float64 = (-2.25 + float64((-2.25 * -2.25)))
        _ = v56
        _ = v57
    }
    switch 9223372036854775807 {
    case 522:
        var v58 uint8 = v55
        _ = v58
    default:
        if len(v9) > 0 {
            v9[0] = -955
        }
    }
    var v59 float64 = (3.14159 + (-2.25 * (0.1 / ((1e10 / 3.14159) / (0.1 - (0.0 / 3.14159))))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v59
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
