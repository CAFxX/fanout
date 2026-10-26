package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 uint16 }

func main() {
    var v1 uint16 = 0
    var v2 float64 = ((-2.25 / float64(1e10)) - 1.5)
    var v3 string = "a"
    var v4 uint8 = 2
    v5 := []int{}
    var v6 string = ("a" + (v3 + v3))
    for fi7 := 0; fi7 < 14; fi7++ {
        switch 30 {
        case 42:
            if (!(true && (42 < 42))) {
                var v8 uint32 = 555
                for fi9 := 0; fi9 < 16; fi9++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi9
                }
                v1 -= (^(v1))
                _ = v8
            } else {
                v1 += v1
            }
        case 1:
            for fi10 := 0; fi10 < 6; fi10++ {
                if (1 >= 18446744073709551615) {
                    fmt.Println(v3)
                    v1 += 352
                }
                if len(v5) > 0 {
                    v5[0] = fi7
                }
                fmt.Println(v2)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi10
            }
        default:
            switch fi7 {
            case 1:
                v6 = (((v6 + "a") + "foo bar") + "a")
            case 42:
                switch "\n\t" {
                case "a":
                    switch fi7 {
                    case 32768:
                        if false {
                            fmt.Println(v1, v6, v6)
                            var v11 uint16 = v1
                            _ = v11
                        }
                    case 4294967295:
                        v5 = append(v5, fi7)
                    case 276:
                        v5 = append(v5, fi7)
                    default:
                        v12 := GenS1{F0: 42, F1: 127, F2: 18446744073709551615}
                        _ = v12
                    }
                case "foo bar":
                    if false {
                        v13 := GenS2{F0: 1}
                        var v14 int8 = 0
                        _ = v13
                        _ = v14
                    } else {
                        if true {
                            v5 = append(v5, fi7)
                        } else {
                            switch v6 {
                            case "":
                                fmt.Println(fi7, v4)
                            case "a":
                                v1 = 2
                            case "foo bar":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            default:
                                var v15 int16 = 32767
                                _ = v15
                            }
                            switch fi7 {
                            case -652:
                                v3 = "hello"
                            case 1:
                                v16 := GenS3{F0: (("" + ((((((v6 + v3) + (v3 + (v6 + v3))) + v3) + v3) + v3) + "\n\t")) + "foo bar")}
                                _ = v16
                            case 9223372036854775807:
                                v17 := GenS4{F0: 69}
                                _ = v17
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            if (!false) {
                                switch v6 {
                                case "foo bar":
                                    v1 = (^(v1))
                                case "a":
                                    v3 = v6
                                case "":
                                    if true {
                                        v2 -= 1e10
                                        v3 = v6
                                    } else {
                                        v4 = 2
                                        var v18 string = v3
                                        v6 = v3
                                        _ = v18
                                    }
                                }
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                var v19 int16 = 279
                                _ = v19
                            } else {
                                for fi20 := 0; fi20 < 9; fi20++ {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi20
                                }
                                v2 -= 3.14159
                            }
                        }
                    }
                default:
                    fmt.Println(v2, v6, fi7)
                }
            case 613:
                switch (("a" + ("foo bar" + v6)) + v6) {
                case "":
                    for fi21 := 0; fi21 < 5; fi21++ {
                        fmt.Println(v4, v6, v2)
                        if false {
                            var v22 string = ""
                            _ = v22
                        }
                        _ = fi21
                    }
                case "a":
                    if (0 <= -127) {
                        for fi23 := 0; fi23 < 14; fi23++ {
                            var v24 uint32 = 2
                            fmt.Println(v1, v3, v6)
                            fmt.Println(fi7, fi23, v6)
                            _ = fi23
                            _ = v24
                        }
                    } else {
                        fmt.Println(v6, v4)
                        var v25 float32 = 1e10
                        v26 := GenS3{F0: (v6 + ((v3 + v6) + "a"))}
                        _ = v25
                        _ = v26
                    }
                case "\n\t":
                    v2 -= 1.5
                default:
                    var v27 int64 = 553
                    _ = v27
                }
            default:
                if true {
                    var v28 uint64 = 1
                    var v29 uint = 4294967295
                    _ = v28
                    _ = v29
                }
            }
        }
        v3 = (v6 + ("hello" + v3))
        _ = fi7
    }
    if len(v5) > 0 {
        v5[0] = 839
    }
    fmt.Println(v1, v4, v3)
    switch (v6 + v3) {
    case "\n\t":
        v4 -= v4
    case "hello":
        var v30 int = 730
        _ = v30
    case "":
        for fi31 := 0; fi31 < 19; fi31++ {
            v5 = append(v5, fi31)
            _ = fi31
        }
    default:
        fmt.Println(v3, v4, v3)
    }
    var v32 float64 = v2
    switch -2147483648 {
    case 4294967295:
        if false {
            var v33 float32 = 0.0
            _ = v33
        }
    case 127:
        for fi34 := 0; fi34 < 4; fi34++ {
            var v35 uint32 = 2147483647
            _ = fi34
            _ = v35
        }
    default:
        v2 += v2
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 int = fi36
        switch v6 {
        case "a":
            v38 := GenS5{F0: 4294967295}
            _ = v38
        case "hello":
            fmt.Println(v32, v1)
        }
        _ = fi36
        _ = v37
    }
    v39 := GenS6{F0: v3, F1: v1}
    if (0 >= 678) {
        switch -2147483648 {
        case -1:
            var v40 uint64 = 1
            _ = v40
        }
        v2 += 0.0
        var v41 bool = true
        _ = v41
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
