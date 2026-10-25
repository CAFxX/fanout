package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := GenS1{F0: 4294967295, F1: (((65535 >= 2) || true) && (!(true || false))), F2: 275}
        fmt.Println(fi4)
        var v6 float64 = (1e10 / 1e10)
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: ((((285 >= 0) && false) && (395 < -505)) || (577 != 850)), F2: -119}
    if (42 > 338) {
        var v8 string = ""
        _ = v8
    } else {
        if true {
            switch 138 {
            case 127:
                var v9 int64 = 709
                _ = v9
            default:
                var v10 string = "\n\t"
                _ = v10
            }
        }
        v3["\n\t"] = 317
        var v11 uint = 1
        _ = v11
    }
    var v12 int32 = -1
    v12 -= (^(v12))
    switch 0 {
    case 32768:
        v12 = (v12 | v12)
    case 452:
        var v13 uint32 = 2
        _ = v13
    case 42:
        for fi14 := 0; fi14 < 9; fi14++ {
            v12 += (^(v12))
            v12 += (v12 | v12)
            v12 += 860
            _ = fi14
        }
    }
    if true {
        if (!(!(65535 <= 2))) {
            v15 := GenS2{F0: 128}
            switch 493 {
            case -389:
                var v16 string = "hello"
                _ = v16
            case -9223372036854775808:
                if (true || false) {
                    fmt.Println(v12)
                    v12 -= v12
                    v12 -= (v12 | (v12 | v12))
                }
            }
            _ = v15
        }
        v12 += 379
        var v17 string = "a"
        _ = v17
    }
    if (15 > 809) {
        v18 := []int{32767}
        v12 -= (v12 | -1)
        if (false && true) {
            if false {
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 uint32 = 2147483647
                    var v21 int = fi19
                    var v22 float32 = 1.5
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
            } else {
                fmt.Println(v12)
                v12 += (v12 / (((v12 * -1)) | 1))
                switch (("foo bar" + ((("\n\t" + (("" + "a") + "foo bar")) + ("" + "\n\t")) + "\n\t")) + "\n\t") {
                case "foo bar":
                    if (42 < 0) {
                        if (v12 > v12) {
                            v23 := GenS3{F0: 0, F1: 1}
                            _ = v23
                        } else {
                            fmt.Println(v12)
                            v12 = 0
                            fmt.Println(v12)
                        }
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v25 := GenS4{F0: 1, F1: 688}
                            if true {
                                v26 := GenS5{F0: int32(fi24)}
                                v27 := GenS5{F0: v12}
                                switch "\n\t" {
                                case "hello":
                                    var v28 float64 = 0.0
                                    _ = v28
                                case "":
                                    if (3.14159 < 0.1) {
                                        switch (fi24 - 32767) {
                                        case 560:
                                            for fi29 := 0; fi29 < 1; fi29++ {
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                                var v30 uint = 4294967295
                                                _ = fi29
                                                _ = v30
                                            }
                                        default:
                                            var v31 int64 = -266
                                            _ = v31
                                        }
                                    }
                                default:
                                    var v32 uint64 = 4294967295
                                    _ = v32
                                }
                                _ = v26
                                _ = v27
                            } else {
                                var v33 string = "a"
                                _ = v33
                            }
                            v12 = (^(v12))
                            _ = fi24
                            _ = v25
                        }
                    }
                case "hello":
                    if (32767 <= 1000) {
                        var v34 float32 = ((0.0 * 3.14159) - (1e10 * 3.14159))
                        delete(v1, -412)
                        var v35 int = 127
                        _ = v34
                        _ = v35
                    } else {
                        var v36 int16 = 1
                        _ = v36
                    }
                }
            }
            v12 -= v12
        } else {
            switch "a" {
            case "hello":
                v12 -= (v12 + v12)
            }
        }
        _ = v18
    } else {
        v37 := GenS1{F0: 0, F1: (!(-305 < 2147483647)), F2: 9223372036854775807}
        _ = v37
    }
    if (true && ((3.14159 <= (-2.25 * 3.14159)) && (320 < 65535))) {
        v12 -= 0
        if false {
            switch 57 {
            case 32768:
                fmt.Println(v12)
            case 935:
                switch "hello" {
                case "a":
                    switch "" {
                    case "foo bar":
                        var v38 uint64 = 4294967295
                        _ = v38
                    }
                }
            }
        } else {
            switch 128 {
            case -1:
                switch ("foo bar" + ("a" + "foo bar")) {
                case "a":
                    v39 := GenS5{F0: v12}
                    _ = v39
                }
            }
        }
        var v40 int = 32768
        _ = v40
    } else {
        switch "a" {
        case "a":
            switch 32767 {
            case 32768:
                var v41 float32 = 0.0
                _ = v41
            case 42:
                v42 := GenS1{F0: 325, F1: false, F2: 345}
                _ = v42
            case -148:
                var v43 string = ("\n\t" + "\n\t")
                _ = v43
            default:
                if (!((!false) || (0 > 4294967295))) {
                    for fi44 := 0; fi44 < 10; fi44++ {
                        v2["hello"] = 674
                        v12 -= v12
                        v45 := GenS5{F0: (v12 ^ v12)}
                        _ = fi44
                        _ = v45
                    }
                }
            }
        case "foo bar":
            var v46 bool = false
            _ = v46
        default:
            if true {
                v47 := GenS5{F0: (v12 & (v12 & v12))}
                v12 = 774
                _ = v47
            } else {
                fmt.Println(v12)
                var v48 string = "a"
                v12 += v12
                _ = v48
            }
        }
        var v49 int32 = (-(v12))
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
}
