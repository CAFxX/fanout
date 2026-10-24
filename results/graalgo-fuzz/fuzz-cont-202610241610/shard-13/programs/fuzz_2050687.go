package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint; F1 int32; F2 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint; F2 int }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 bool }

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    var v4 int16 = (v1 + (v1 ^ v1))
    v5 := make(map[uint32]int)
    var v6 uint16 = 65535
    if (4294967295 < 0) {
        v1 = v1
    } else {
        var v7 int8 = -128
        v7 = (v7 * (v7 | (^(v7))))
        _ = v7
    }
    var v8 float64 = 3.14159
    v4 = v1
    v9 := []int{-803, 128, -118, 4294967295}
    v10 := GenS1{F0: 997, F1: 4294967295}
    v6 += 185
    fmt.Println(v4, v8)
    var v11 int = 2147483647
    v1 = v4
    v1 = (v4 & v1)
    switch v11 {
    case 650:
        delete(v5, 553)
    case -9223372036854775808:
        switch ("hello" + ("hello" + (("hello" + "foo bar") + ""))) {
        case "foo bar":
            if true {
                v11 -= v11
                v12 := GenS1{F0: 2147483647, F1: 795}
                v13 := GenS1{F0: 42, F1: 42}
                _ = v12
                _ = v13
            } else {
                v11 -= v11
                v8 += (float64(1e10) - v8)
                v4 += 0
            }
        case "hello":
            var v14 uint16 = 939
            _ = v14
        case "a":
            v11 = (v11 >> ((623) & 63))
        default:
            var v15 uint16 = v6
            _ = v15
        }
    case -801:
        for fi16 := 0; fi16 < 10; fi16++ {
            v8 = ((v8 - v8) + 0.0)
            if false {
                v3["foo bar"] = -28
                _ = len(v9)
                switch "" {
                case "foo bar":
                    v9 = append(v9, v11)
                case "\n\t":
                    var v17 bool = false
                    _ = v17
                }
            } else {
                v18 := GenS1{F0: -1, F1: 1}
                if (!((^(v6)) == v6)) {
                    var v19 float32 = 1e10
                    _ = v19
                }
                _ = v18
            }
            fmt.Println(v1, fi16, v11)
            _ = fi16
        }
    default:
        v6 += (v6 / (((v6 * 0)) | 1))
    }
    var v20 int8 = 42
    v8 = v8
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS1{F0: -2147483648, F1: 2147483648}
        fmt.Println(v20, v6, v4)
        for fi23 := 0; fi23 < 18; fi23++ {
            if (611 >= 22) {
                if true {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    switch fi21 {
                    case 9223372036854775807:
                        var v24 int = (fi23 - 2147483647)
                        _ = v24
                    case 4294967295:
                        switch v11 {
                        case 2147483647:
                            if (-2.25 == 3.14159) {
                                _ = len(v2)
                                var v25 float32 = (3.14159 + ((1.5 * ((3.14159 + 0.0) - 1e10)) / 1e10))
                                v8 = v8
                                _ = v25
                            }
                        case -2147483648:
                            fmt.Println(v1, v6)
                        default:
                            v26 := GenS1{F0: -1, F1: 65535}
                            _ = v26
                        }
                    case 32767:
                        v27 := GenS2{F0: 9223372036854775807, F1: 18446744073709551615}
                        _ = v27
                    }
                }
                v28 := GenS1{F0: 42, F1: 65535}
                v8 -= v8
                _ = v28
            }
            delete(v3, (("hello" + "hello") + "hello"))
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    var v29 int = 856
    for fi30 := 0; fi30 < 7; fi30++ {
        if (1 < 9223372036854775807) {
            v31 := GenS3{F0: 22, F1: 0, F2: true}
            _ = v31
        } else {
            v32 := GenS2{F0: (v11 ^ (v29 - 128)), F1: 546}
            v33 := GenS1{F0: -2147483648, F1: 42}
            v8 -= -2.25
            _ = v32
            _ = v33
        }
        v6 = (v6 >> ((1) & 63))
        var v34 bool = ((!false) || ("hello" > (("\n\t" + ("\n\t" + "hello")) + "a")))
        _ = fi30
        _ = v34
    }
    fmt.Println(v8, v29)
    for fi35 := 0; fi35 < 1; fi35++ {
        for fi36 := 0; fi36 < 12; fi36++ {
            var v37 string = "hello"
            var v38 uint16 = (v6 % ((v6) | 1))
            fmt.Println(v6)
            v39 := GenS4{F0: 42}
            v40 := GenS1{F0: 2147483647, F1: 1}
            v41 := GenS5{F0: (425 == (v1 & -741)), F1: 553, F2: v11}
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
            _ = v40
            _ = v41
        }
        if (!(!false)) {
            v42 := GenS6{F0: 2, F1: "a"}
            var v43 int32 = 2147483647
            switch ("a" + "a") {
            case "":
                switch "\n\t" {
                case "foo bar":
                    if (true || true) {
                        if false {
                            if (true || false) {
                                var v44 uint = 1
                                _ = v44
                            }
                        }
                        if false {
                            switch "" {
                            case "hello":
                                for fi45 := 0; fi45 < 15; fi45++ {
                                    fmt.Println(v4)
                                    var v46 float64 = (1.5 + 3.14159)
                                    v29 = fi35
                                    _ = fi45
                                    _ = v46
                                }
                            case "":
                                v47 := GenS4{F0: 0}
                                _ = v47
                            case "a":
                                var v48 uint64 = 4294967295
                                _ = v48
                            }
                            v49 := GenS1{F0: (-(v43)), F1: 2147483648}
                            fmt.Println(v11, fi35)
                            _ = v49
                        }
                        delete(v3, "")
                    } else {
                        v50 := GenS7{F0: false}
                        _ = v50
                    }
                }
            case "hello":
                if (!false) {
                    if (false || (!(-2.25 >= 0.0))) {
                        var v51 int64 = 478
                        switch "foo bar" {
                        case "foo bar":
                            var v52 int8 = (v20 >> ((uint(v6)) & 63))
                            _ = v52
                        }
                        _ = v51
                    }
                }
            default:
                if (927 == 2) {
                    var v53 uint8 = 0
                    var v54 bool = false
                    _ = v53
                    _ = v54
                } else {
                    for fi55 := 0; fi55 < 17; fi55++ {
                        v11 += 638
                        v6 = (^(v6))
                        var v56 float64 = v8
                        _ = fi55
                        _ = v56
                    }
                    switch (("foo bar" + "\n\t") + ("" + ("\n\t" + ("" + "a")))) {
                    case "foo bar":
                        v43 = int32(v4)
                    case "hello":
                        v1 = (v1 + 0)
                    case "":
                        if true {
                            switch -1 {
                            case 32768:
                                if (-2.25 >= 1e10) {
                                    v57 := v5[2147483647]
                                    _ = v57
                                }
                            case 2147483647:
                                var v58 string = "a"
                                _ = v58
                            case 32767:
                                var v59 uint16 = 42
                                _ = v59
                            }
                            v2["\n\t"] = 508
                        } else {
                            var v60 string = "\n\t"
                            fmt.Println(v60, v4)
                            var v61 float32 = 0.0
                            _ = v60
                            _ = v61
                        }
                    }
                }
            }
            _ = v42
            _ = v43
        }
        _ = fi35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
