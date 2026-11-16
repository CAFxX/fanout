package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 int32 = -1
    v2 := []string{("\n\t" + "foo bar"), ((("" + "foo bar") + "hello") + "foo bar")}
    var v3 uint = 18446744073709551615
    var v4 int64 = -364
    var v5 int8 = 127
    if true {
        var v6 uint16 = 42
        if true {
            for fi7 := 0; fi7 < 6; fi7++ {
                v2 = append(v2, "a")
                _ = fi7
            }
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 int64 = v4
                _ = fi8
                _ = v9
            }
        }
        v6 += 0
        _ = v6
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v3 = (^(v3))
        _ = fi10
    }
    switch ("" + (("hello" + "") + ("hello" + "a"))) {
    case "foo bar":
        switch 9223372036854775807 {
        case -314:
            for fi11 := 0; fi11 < 4; fi11++ {
                v1 = (v1 | (v1 / (((v1 << (((v3 << ((v3) & 63))) & 63))) | 1)))
                v12 := GenS1{F0: "hello"}
                _ = fi11
                _ = v12
            }
        default:
            var v13 float64 = 1.5
            _ = v13
        }
    case "a":
        for fi14 := 0; fi14 < 7; fi14++ {
            if true {
                var v15 int = -266
                _ = v15
            }
            _ = fi14
        }
    default:
        if (!("foo bar" > "a")) {
            var v16 float32 = float32(((-2.25 + (1.5 + 1e10)) / ((3.14159 / -2.25) - (3.14159 + 0.0))))
            v16 = v16
            switch "" {
            case "\n\t":
                v4 = v4
            case "":
                v17 := GenS2{F0: 2147483648}
                _ = v17
            case "hello":
                var v18 string = (("\n\t" + ("" + (("" + (("a" + "a") + "")) + ""))) + "foo bar")
                _ = v18
            default:
                if (false && true) {
                    v19 := GenS3{F0: 0}
                    _ = v19
                } else {
                    v2 = append(v2, ("a" + "a"))
                    v3 -= (v3 | 126)
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                }
            }
            _ = v16
        } else {
            if ((-9223372036854775808 != (v4 % ((1) | 1))) && ((v1 % ((v1) | 1)) < 0)) {
                for fi20 := 0; fi20 < 14; fi20++ {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v5 -= v5
                        if (32767 <= 42) {
                            if false {
                                switch fi20 {
                                case -1:
                                    v3 = v3
                                case 2147483647:
                                    fmt.Println(fi20, fi20, v3)
                                case -2147483648:
                                    v4 = v4
                                }
                                var v22 uint = v3
                                v4 = v4
                                _ = v22
                            } else {
                                v3 += 2
                                switch (fi20 % (((fi20 * (fi20 << ((v3) & 63)))) | 1)) {
                                case -527:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 127:
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(fi21)
                                    default:
                                        v23 := GenS4{F0: 482, F1: (fi21 ^ fi21)}
                                        _ = v23
                                    }
                                }
                                var v24 int = fi21
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                var v25 bool = false
                                _ = v24
                                _ = v25
                            }
                            if false {
                                v26 := GenS2{F0: 2147483647}
                                fmt.Println(fi21, v3, v3)
                                _ = v26
                            }
                            switch "\n\t" {
                            case "":
                                fmt.Println(fi20, v4)
                            case "hello":
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                            case "a":
                                fmt.Println(v1)
                            default:
                                v27 := GenS2{F0: 42}
                                _ = v27
                            }
                        }
                        _ = fi21
                    }
                    var v28 int = (fi20 << ((v3) & 63))
                    for fi29 := 0; fi29 < 14; fi29++ {
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                v1 = (-(v1))
                            }
                            v30 := GenS5{F0: fi29, F1: v4}
                            _ = v30
                        }
                        _ = fi29
                    }
                    _ = fi20
                    _ = v28
                }
                var v31 int16 = 896
                var v32 bool = true
                _ = v31
                _ = v32
            }
            switch -2147483648 {
            case -1:
                fmt.Println(v1, v1, v5)
            case 32767:
                var v33 float64 = 1e10
                _ = v33
            }
            fmt.Println(v5)
        }
    }
    var v34 int = -885
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v35 float32 = 1.5
    for fi36 := 0; fi36 < 14; fi36++ {
        fmt.Println(v4, v34)
        var v37 string = (((("foo bar" + ((((("hello" + "foo bar") + "\n\t") + (("\n\t" + (("" + (("\n\t" + "a") + "")) + ("\n\t" + "foo bar"))) + "hello")) + "\n\t") + "")) + "") + "a") + (("foo bar" + (("foo bar" + "a") + "a")) + "\n\t"))
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 2; fi38++ {
        var v39 uint = (v3 / ((2) | 1))
        if ((v4 % (((v4 >> (((v3 / (((v39 + v3)) | 1))) & 63))) | 1)) <= 661) {
            var v40 bool = false
            _ = v40
        } else {
            switch fi38 {
            case 0:
                var v41 float64 = (0.0 + ((0.1 - (1.5 * float64(v35))) - (3.14159 - (3.14159 + (((float64((float32(-2.25) - (v35 * v35))) * 3.14159) / 0.1) / 1.5)))))
                _ = v41
            case 790:
                if true {
                    fmt.Println(v34, v3)
                    v4 -= v4
                    fmt.Println(v1, v1)
                } else {
                    for fi42 := 0; fi42 < 10; fi42++ {
                        var v43 uint = v39
                        _ = fi42
                        _ = v43
                    }
                    if true {
                        v4 += 474
                        var v44 int64 = (v4 ^ (v4 ^ v4))
                        if true {
                            var v45 uint16 = 770
                            var v46 int32 = v1
                            if (!(42 != 978)) {
                                var v47 int = -1
                                _ = v47
                            } else {
                                var v48 int8 = -1
                                _ = v48
                            }
                            _ = v45
                            _ = v46
                        }
                        _ = v44
                    }
                    var v49 string = ""
                    _ = v49
                }
            default:
                var v50 int8 = 42
                _ = v50
            }
            var v51 float32 = ((1.5 / 1e10) - (float32(1.5) + ((v35 + v35) / v35)))
            for fi52 := 0; fi52 < 12; fi52++ {
                v3 -= v39
                v39 = (v39 - v3)
                var v53 uint16 = 255
                _ = fi52
                _ = v53
            }
            _ = v51
        }
        _ = fi38
        _ = v39
    }
    if false {
        var v54 bool = true
        if (463 != 4294967295) {
            v1 = 1
            switch v34 {
            case 9223372036854775807:
                v55 := GenS1{F0: "foo bar"}
                _ = v55
            case -300:
                v56 := GenS6{F0: 1}
                _ = v56
            case 127:
                var v57 int8 = 127
                _ = v57
            }
        }
        fmt.Println(v3, v3, v34)
        _ = v54
    } else {
        var v58 float64 = ((1e10 / ((1.5 - 0.0) / float64(3.14159))) / 0.1)
        v59 := GenS3{F0: 32767}
        _ = v58
        _ = v59
    }
    fmt.Println(v34, v34, v5)
    var v60 uint8 = 255
    v5 = (^(v5))
    v1 = (^(v1))
    v61 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v60
    _ = v61
}
