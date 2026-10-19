package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 uint; F2 uint }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8; F1 int }

func main() {
    var v1 string = "a"
    var v2 float32 = 0.0
    var v3 float64 = 1e10
    var v4 float32 = float32(v3)
    v5 := []string{v1, "foo bar"}
    fmt.Println(v3, v2)
    if true {
        var v6 string = "a"
        if (-1 >= 0) {
            if len(v5) > 0 {
                v5[0] = (v6 + (v6 + (v6 + "a")))
            }
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 float64 = v3
                if ((!true) && ((!(!(!(2147483648 <= 65535)))) && ((^(fi7)) <= fi7))) {
                    for fi9 := 0; fi9 < 15; fi9++ {
                        switch (^(fi9)) {
                        case 42:
                            if true {
                                var v10 float32 = float32(v8)
                                v11 := GenS1{F0: 2147483647, F1: fi7}
                                switch "foo bar" {
                                case "":
                                    if false {
                                        var v12 float64 = 0.0
                                        v5 = append(v5, "hello")
                                        var v13 uint32 = 65535
                                        _ = v12
                                        _ = v13
                                    }
                                }
                                fmt.Println(fi9, v4)
                                if false {
                                    switch v6 {
                                    case "hello":
                                        fmt.Println(fi9, v10)
                                    case "":
                                        if true {
                                            v2 = 0.0
                                            if false {
                                                v14 := GenS1{F0: -1, F1: 0}
                                                if len(v5) > 0 {
                                                    v5[0] = v1
                                                }
                                                _ = v14
                                            } else {
                                                v2 += ((1.5 / (v4 + (v4 - -2.25))) / float32((v8 * v8)))
                                                if false {
                                                    var v15 int32 = 0
                                                    if true {
                                                        v15 = -1
                                                        v4 -= (v4 / 1.0)
                                                    }
                                                    v1 = (v6 + v6)
                                                    var v16 uint64 = 778
                                                    v17 := GenS1{F0: 873, F1: 838}
                                                    switch (^(fi7)) {
                                                    case -624:
                                                        switch (fi9 / (((fi7 * fi9)) | 1)) {
                                                        case 128:
                                                            v18 := GenS1{F0: 2147483647, F1: (fi7 % ((fi9) | 1))}
                                                            _ = v18
                                                        case -963:
                                                            v19 := GenS2{F0: -2147483648}
                                                            _ = v19
                                                        }
                                                    default:
                                                        var v20 uint = 1
                                                        _ = v20
                                                    }
                                                    _ = v15
                                                    _ = v16
                                                    _ = v17
                                                } else {
                                                    v21 := GenS3{F0: 5}
                                                    _ = v21
                                                }
                                            }
                                            v6 = (v6 + v1)
                                        }
                                    case "\n\t":
                                        if (2 == 0) {
                                            switch fi9 {
                                            case 9223372036854775807:
                                                v22 := GenS2{F0: -385}
                                                _ = v22
                                            case 32767:
                                                v23 := GenS1{F0: 2147483647, F1: 32768}
                                                v5 = append(v5, v6)
                                                _ = v23
                                            case 1:
                                                var v24 uint8 = 0
                                                _ = v24
                                            default:
                                                v25 := GenS4{F0: int32(fi9), F1: 18446744073709551615}
                                                _ = v25
                                            }
                                            switch "a" {
                                            case "hello":
                                                switch 127 {
                                                case -9223372036854775808:
                                                    v26 := GenS5{F0: 284, F1: 644}
                                                    _ = v26
                                                }
                                            case "a":
                                                v6 = v1
                                            default:
                                                v27 := GenS6{F0: -128}
                                                _ = v27
                                            }
                                            if (0 < 6) {
                                                fmt.Println(v10)
                                            }
                                        } else {
                                            switch v1 {
                                            case "hello":
                                                v5 = append(v5, v6)
                                            case "\n\t":
                                                if false {
                                                    var v28 string = (v1 + (v1 + (("\n\t" + "hello") + v6)))
                                                    fmt.Println(fi9, v4, v4)
                                                    switch ("\n\t" + "hello") {
                                                    case "":
                                                        fmt.Println(v4, v10, v4)
                                                    case "foo bar":
                                                        v29 := GenS7{F0: v6, F1: 42, F2: 42}
                                                        _ = v29
                                                    default:
                                                        var v30 float64 = v8
                                                        _ = v30
                                                    }
                                                    _ = v28
                                                } else {
                                                    var v31 float64 = v3
                                                    _ = v31
                                                }
                                            case "a":
                                                v6 = (v1 + v1)
                                            }
                                        }
                                    }
                                    fmt.Println(v3, v6)
                                    if true {
                                        if (553 != 730) {
                                            var v32 uint64 = 42
                                            _ = v32
                                        }
                                        v6 = v1
                                        v33 := GenS8{F0: 973, F1: -14}
                                        _ = v33
                                    } else {
                                        if (2147483647 == 2147483647) {
                                            var v34 uint16 = 0
                                            if true {
                                                var v35 int = -428
                                                var v36 uint64 = 2
                                                _ = len(v5)
                                                _ = v35
                                                _ = v36
                                            }
                                            _ = v34
                                        }
                                        v37 := GenS7{F0: "", F1: 968, F2: 65535}
                                        switch v6 {
                                        case "\n\t":
                                            v3 += ((float64(v10) + (1e10 * v8)) * v3)
                                        case "hello":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        _ = v37
                                    }
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = ("hello" + ("\n\t" + v1))
                                    }
                                    v6 = v1
                                }
                                _ = v10
                                _ = v11
                            } else {
                                v38 := GenS5{F0: 0, F1: 2}
                                v1 = ("a" + v1)
                                _ = v38
                            }
                        case 4294967295:
                            fmt.Println(fi7, v1, v6)
                        case 128:
                            var v39 float32 = 1e10
                            _ = v39
                        default:
                            var v40 float32 = (1.5 - -2.25)
                            _ = v40
                        }
                        switch ((("a" + v1) + (v6 + v1)) + v6) {
                        case "hello":
                            var v41 int = fi9
                            _ = v41
                        case "foo bar":
                            v42 := GenS9{F0: 227}
                            _ = v42
                        case "\n\t":
                            var v43 uint16 = 773
                            v3 -= (v8 / v8)
                            _ = v43
                        default:
                            v44 := GenS3{F0: 18446744073709551615}
                            _ = v44
                        }
                        _ = fi9
                    }
                } else {
                    v3 = v8
                    v8 += (v3 * 1.5)
                }
                _ = fi7
                _ = v8
            }
        }
        _ = v6
    } else {
        if true {
            v4 = (((v2 * (v2 + v4)) - v4) * 1e10)
            fmt.Println(v4, v2, v4)
            switch 53 {
            case -2147483648:
                v45 := GenS2{F0: -475}
                _ = v45
            default:
                v2 += -2.25
            }
        }
    }
    fmt.Println(v2)
    for fi46 := 0; fi46 < 5; fi46++ {
        if (true && false) {
            v3 = float64(float32((v3 - (v3 - 1e10))))
            v1 = (v1 + "a")
            v47 := GenS3{F0: 753}
            _ = v47
        } else {
            switch -951 {
            case -9223372036854775808:
                for fi48 := 0; fi48 < 16; fi48++ {
                    v3 -= 1e10
                    fmt.Println(v1)
                    v2 = v4
                    _ = fi48
                }
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            v49 := GenS8{F0: 255, F1: 100}
            switch (v1 + v1) {
            case "":
                var v50 int32 = -2147483648
                _ = v50
            }
            _ = v49
        }
        var v51 uint32 = 0
        _ = fi46
        _ = v51
    }
    fmt.Println(v3, v4)
    _ = len(v5)
    for fi52 := 0; fi52 < 11; fi52++ {
        var v53 int = (fi52 & 1)
        v54 := GenS7{F0: v1, F1: 0, F2: 0}
        _ = fi52
        _ = v53
        _ = v54
    }
    var v55 int8 = -18
    var v56 int32 = 42
    fmt.Println(v2)
    if (float32(0.1) >= 1.5) {
        if (false && (32767 == 1000)) {
            v4 += 0.0
            var v57 uint = 42
            v3 = (3.14159 - v3)
            _ = v57
        } else {
            fmt.Println(v3)
            if (((v3 * 3.14159) - ((((0.1 - 0.1) * v3) - (v3 / v3)) + (v3 + (v3 - 0.1)))) != v3) {
                for fi58 := 0; fi58 < 20; fi58++ {
                    v59 := GenS10{F0: 255, F1: -319}
                    _ = fi58
                    _ = v59
                }
                var v60 uint = 65535
                _ = v60
            }
        }
        if false {
            _ = len(v5)
            v56 -= (^(v56))
            for fi61 := 0; fi61 < 10; fi61++ {
                _ = len(v5)
                _ = fi61
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
}
