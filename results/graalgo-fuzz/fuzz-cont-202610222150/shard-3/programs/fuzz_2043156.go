package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 int64; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 int64; F1 uint16 }

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 976
    var v3 uint8 = 1
    for fi4 := 0; fi4 < 18; fi4++ {
        v2 -= uint16(v1)
        var v5 uint8 = 0
        var v6 float32 = 0.1
        _ = fi4
        _ = v5
        _ = v6
    }
    if false {
        v7 := []int{1}
        _ = v7
    }
    fmt.Println(v3)
    v8 := GenS1{F0: 290}
    fmt.Println(v1)
    switch -383 {
    case 999:
        v9 := []int{-9223372036854775808}
        _ = v9
    case 32767:
        if (((((("" + "") + "") + "a") + "") + "hello") == ("a" + ("foo bar" + ""))) {
            if true {
                var v10 uint = 661
                _ = v10
            } else {
                var v11 int8 = 1
                v1 = (v1 % ((v1) | 1))
                var v12 float32 = ((1.5 * -2.25) - -2.25)
                _ = v11
                _ = v12
            }
            var v13 uint32 = 908
            if (((float32((0.0 / 1.5)) + ((1.5 - 3.14159) + 0.0)) + 0.0) != (1e10 * 1e10)) {
                switch "a" {
                case "a":
                    v14 := GenS2{F0: 186, F1: int64(v2), F2: (^(v3))}
                    _ = v14
                case "hello":
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v1)
                    case "hello":
                        var v15 float64 = 0.1
                        _ = v15
                    default:
                        fmt.Println(v13)
                    }
                default:
                    var v16 string = "\n\t"
                    _ = v16
                }
            } else {
                v13 -= v13
                switch 51 {
                case -935:
                    v17 := []int{int(v3), 9223372036854775807, 32767, 943}
                    _ = v17
                default:
                    var v18 int32 = 1
                    _ = v18
                }
            }
            _ = v13
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                v1 += v1
                _ = fi19
            }
            if (v3 <= 42) {
                v20 := GenS1{F0: -156}
                var v21 float32 = float32((0.0 - 3.14159))
                _ = v20
                _ = v21
            } else {
                switch 4294967295 {
                case -445:
                    var v22 uint8 = (v3 % (((v3 ^ 10)) | 1))
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 18; fi23++ {
                        switch 0 {
                        case 42:
                            v24 := GenS2{F0: fi23, F1: int64(fi23), F2: (v3 >> ((65535) & 63))}
                            _ = v24
                        }
                        var v25 uint16 = (v2 ^ 65535)
                        if (151 == 331) {
                            switch ("hello" + "hello") {
                            case "a":
                                var v26 uint64 = (v1 - 42)
                                _ = v26
                            default:
                                var v27 int32 = 42
                                _ = v27
                            }
                        }
                        _ = fi23
                        _ = v25
                    }
                }
                var v28 uint32 = 2147483648
                v2 += (v2 & (v2 >> ((255) & 63)))
                _ = v28
            }
        }
    case 0:
        fmt.Println(v2)
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        if false {
            if (!true) {
                for fi30 := 0; fi30 < 18; fi30++ {
                    if (v2 > v2) {
                        if true {
                            v3 = 127
                            var v31 uint64 = (v1 + 18446744073709551615)
                            if (!(!false)) {
                                fmt.Println(v2, v31, fi30)
                                var v32 int64 = -714
                                _ = v32
                            }
                            _ = v31
                        } else {
                            v33 := GenS3{F0: -1}
                            switch fi29 {
                            case -1:
                                v1 += 18446744073709551615
                            case 32768:
                                v3 += 2
                            case 4294967295:
                                v34 := GenS1{F0: -775}
                                _ = v34
                            default:
                                v2 = (v2 * v2)
                            }
                            _ = v33
                        }
                        var v35 int64 = 0
                        _ = v35
                    }
                    _ = fi30
                }
            }
        }
        fmt.Println(v2, v3)
        var v36 int64 = 276
        _ = fi29
        _ = v36
    }
    var v37 uint16 = 2
    switch (("foo bar" + "foo bar") + "\n\t") {
    case "hello":
        v1 += (v1 % ((v1) | 1))
    case "\n\t":
        for fi38 := 0; fi38 < 11; fi38++ {
            v39 := GenS1{F0: 42}
            for fi40 := 0; fi40 < 8; fi40++ {
                switch "\n\t" {
                case "":
                    var v41 float64 = float64(3.14159)
                    _ = v41
                }
                switch "" {
                case "\n\t":
                    v3 += v3
                    var v42 float64 = (3.14159 - 1e10)
                    var v43 int32 = 0
                    _ = v42
                    _ = v43
                }
                _ = fi40
            }
            var v44 float64 = 1e10
            _ = fi38
            _ = v39
            _ = v44
        }
    default:
        v45 := GenS4{F0: 1, F1: "\n\t"}
        _ = v45
    }
    if (9223372036854775807 < 42) {
        var v46 int16 = -918
        if true {
            var v47 int = 42
            v48 := GenS4{F0: (v47 * v47), F1: "foo bar"}
            _ = v47
            _ = v48
        }
        _ = v46
    } else {
        if true {
            if false {
                v3 -= (^(v3))
                if false {
                    var v49 uint16 = v37
                    for fi50 := 0; fi50 < 7; fi50++ {
                        var v51 uint16 = (v49 / ((v2) | 1))
                        v51 -= (v49 % (((v51 << ((0) & 63))) | 1))
                        var v52 int32 = -2147483648
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    _ = v49
                }
            } else {
                v53 := []int{921, 390, -9223372036854775808}
                var v54 float32 = -2.25
                _ = v53
                _ = v54
            }
            v55 := []int{}
            _ = v55
        } else {
            fmt.Println(v37, v3)
            if (-2147483648 != -2147483648) {
                switch "a" {
                case "\n\t":
                    v37 = (v2 % (((v37 << ((1) & 63))) | 1))
                }
                var v56 int = 128
                _ = v56
            }
        }
        if true {
            v57 := GenS5{F0: 127, F1: ("a" + (((("\n\t" + "") + (("a" + "foo bar") + "a")) + "hello") + "foo bar"))}
            for fi58 := 0; fi58 < 6; fi58++ {
                v59 := GenS6{F0: -9223372036854775808, F1: v37}
                _ = fi58
                _ = v59
            }
            _ = v57
        } else {
            v2 += (v37 / ((v2) | 1))
        }
    }
    var v60 uint64 = (v1 / ((v1) | 1))
    switch -802 {
    case 4294967295:
        for fi61 := 0; fi61 < 8; fi61++ {
            var v62 uint64 = (v60 >> ((782) & 63))
            v63 := []int{fi61, (fi61 >> ((65535) & 63)), 32768}
            _ = fi61
            _ = v62
            _ = v63
        }
    case -687:
        v60 -= 277
    case 2147483647:
        var v64 int = -9223372036854775808
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v60
    _ = v8
}
