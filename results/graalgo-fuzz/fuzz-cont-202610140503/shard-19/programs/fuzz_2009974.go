package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int; F1 string; F2 uint64 }
type GenS6 struct { F0 uint32; F1 int8; F2 int }
func genFn0(p0 string, p1 int32, p2 string) uint32 {
    p1 -= p1
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 bool = true
        _ = fi6
        _ = v7
    }
    var v8 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483648
}
func genFn1(p0 string, p1 string, p2 int64) int8 {
    v9 := []int{32767, 0}
    var v10 uint = 255
    var v11 string = p0
    p1 = p0
    v11 = ("a" + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 100
}

func main() {
    v1 := []uint32{42}
    v2 := []string{("a" + (("\n\t" + ("\n\t" + "foo bar")) + (("foo bar" + "a") + "a"))), (("a" + (("" + ((("\n\t" + "foo bar") + "\n\t") + "hello")) + "")) + "\n\t"), ("a" + "")}
    var v3 uint64 = 1
    var v4 float32 = 0.1
    var v5 uint = 65535
    v1 = append(v1, 1)
    for fi12 := 0; fi12 < 20; fi12++ {
        if false {
            for fi13 := 0; fi13 < 18; fi13++ {
                v5 = (v5 & 42)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi13
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi12
    }
    if (128 > 127) {
        fmt.Println(v5)
        var v14 uint16 = 65535
        _ = v14
    } else {
        v15 := GenS1{F0: 2, F1: 1, F2: 93}
        v16 := GenS2{F0: 4294967295, F1: "hello"}
        _ = v15
        _ = v16
    }
    if true {
        switch "a" {
        case "":
            v17 := GenS1{F0: 2147483647, F1: -1, F2: -128}
            _ = v17
        case "a":
            v1 = append(v1, 2147483648)
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(fi19, fi19)
                v4 -= v4
                v3 -= (v3 + (v3 % (((v3 / ((v3) | 1))) | 1)))
                _ = fi19
            }
            _ = fi18
        }
    } else {
        switch 1 {
        case 32768:
            if (32768 > 32767) {
                if false {
                    var v20 float64 = -2.25
                    for fi21 := 0; fi21 < 17; fi21++ {
                        fmt.Println(v20, v4, v3)
                        _ = fi21
                    }
                    _ = v20
                }
                v5 = v5
            } else {
                v3 = v3
                fmt.Println(v5, v3, v5)
            }
        case 18:
            if len(v1) > 0 {
                v1[0] = 65535
            }
        case 42:
            v22 := GenS2{F0: 552, F1: ""}
            _ = v22
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            fmt.Println(v4, v3)
            _ = fi23
        }
    }
    _ = len(v2)
    if false {
        v24 := GenS1{F0: 4294967295, F1: 0, F2: 12}
        var v25 float64 = 0.0
        _ = len(v1)
        _ = v24
        _ = v25
    }
    if ((((268 <= genFn0(("" + "a"), 42, "a")) || (float64((((v4 + (-2.25 - (((v4 - 0.1) / -2.25) - ((v4 + v4) + float32(3.14159))))) + (float32((0.0 / 0.1)) - (((0.0 * float32(3.14159)) - 1e10) + 0.1))) * (v4 / float32((-2.25 - 0.1))))) <= 0.1)) || true) || true) {
        if false {
            if false {
                v5 -= 255
            } else {
                var v26 uint8 = 1
                _ = v26
            }
            v3 = v3
            v5 = (v5 - v5)
        }
        if false {
            var v27 uint8 = 7
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        switch (^(fi28)) {
        case -913:
            var v29 uint = v5
            _ = v29
        default:
            if false {
                var v30 int16 = -229
                switch ("foo bar" + "foo bar") {
                case "a":
                    if (-515 != 1) {
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            var v31 uint32 = genFn0("a", 42, "hello")
                            _ = v31
                        case "foo bar":
                            var v32 int32 = 0
                            _ = v32
                        case "":
                            v33 := GenS3{F0: -1}
                            _ = v33
                        default:
                            for fi34 := 0; fi34 < 6; fi34++ {
                                var v35 string = "\n\t"
                                _ = fi34
                                _ = v35
                            }
                        }
                        fmt.Println(v4, fi28, v5)
                        fmt.Println(v30, v30, v4)
                    } else {
                        var v36 uint = (^(v5))
                        switch "" {
                        case "":
                            var v37 uint64 = (v3 * (v3 % ((2) | 1)))
                            _ = v37
                        case "foo bar":
                            v3 = (v3 | 18446744073709551615)
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        switch fi28 {
                        case 32768:
                            fmt.Println(v5, v3)
                        default:
                            fmt.Println(fi28, fi28)
                        }
                        _ = v36
                    }
                case "hello":
                    for fi38 := 0; fi38 < 8; fi38++ {
                        v39 := GenS4{F0: 2147483648, F1: (fi28 | 4294967295)}
                        _ = fi38
                        _ = v39
                    }
                default:
                    var v40 int8 = 1
                    _ = v40
                }
                _ = v30
            }
        }
        v41 := GenS5{F0: fi28, F1: "\n\t", F2: (v3 + v3)}
        _ = fi28
        _ = v41
    }
    for fi42 := 0; fi42 < 5; fi42++ {
        fmt.Println(fi42)
        if (true || true) {
            for fi43 := 0; fi43 < 18; fi43++ {
                v5 = 11
                _ = fi43
            }
            if true {
                fmt.Println(v5, fi42, v5)
                fmt.Println(v3)
            }
            _ = len(v2)
        } else {
            v44 := GenS3{F0: 641}
            switch fi42 {
            case -9223372036854775808:
                v45 := GenS3{F0: 174}
                _ = v45
            case 2147483647:
                var v46 int = (fi42 * -2147483648)
                _ = v46
            default:
                if true {
                    v5 += (v5 * 65535)
                    v4 += float32(float64(v4))
                    v1 = append(v1, 0)
                }
            }
            for fi47 := 0; fi47 < 1; fi47++ {
                var v48 float64 = -2.25
                _ = fi47
                _ = v48
            }
            _ = v44
        }
        _ = fi42
    }
    var v49 uint64 = v3
    var v50 bool = true
    _ = len(v2)
    _ = len(v2)
    var v51 int16 = 1000
    fmt.Println(v50)
    switch 127 {
    case 102:
        var v52 uint = uint(v49)
        _ = v52
    case -235:
        switch ("hello" + "foo bar") {
        case "":
            switch 128 {
            case 127:
                var v53 float32 = (0.1 - v4)
                _ = v53
            case 0:
                v4 = v4
            case 128:
                v51 = 32767
            default:
                v54 := GenS2{F0: (v49 + (v3 >> (((^(v5))) & 63))), F1: (("" + "foo bar") + "foo bar")}
                _ = v54
            }
        case "a":
            v3 += (v3 / ((v3) | 1))
        case "\n\t":
            v50 = (true || v50)
        default:
            var v55 uint = (v5 + v5)
            _ = v55
        }
    case 1:
        v50 = (1 < 786)
    }
    switch 32767 {
    case -9223372036854775808:
        for fi56 := 0; fi56 < 18; fi56++ {
            if v50 {
                var v57 int16 = 32767
                var v58 uint16 = 65535
                _ = v57
                _ = v58
            }
            _ = fi56
        }
    default:
        for fi59 := 0; fi59 < 17; fi59++ {
            for fi60 := 0; fi60 < 9; fi60++ {
                if (v50 || true) {
                    switch fi60 {
                    case 32767:
                        switch fi59 {
                        case 42:
                            v61 := GenS2{F0: (v49 / ((v3) | 1)), F1: "hello"}
                            _ = v61
                        default:
                            if ((v5 <= (v5 >> ((0) & 63))) || true) {
                                _ = len(v1)
                            }
                        }
                    case 32768:
                        v62 := GenS5{F0: (fi60 % ((fi59) | 1)), F1: "foo bar", F2: (v3 % (((v3 % ((456) | 1))) | 1))}
                        _ = v62
                    case 128:
                        switch (fi60 + fi60) {
                        case 2147483647:
                            if len(v2) > 0 {
                                v2[0] = (("\n\t" + "a") + (("a" + "") + ("hello" + "a")))
                            }
                        case 239:
                            var v63 int16 = v51
                            _ = v63
                        case -138:
                            switch ("" + "\n\t") {
                            case "":
                                v50 = ((v50 && (-122 == 127)) || false)
                            }
                        }
                    }
                }
                _ = fi60
            }
            switch 0 {
            case 32768:
                fmt.Println(v5)
            default:
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
            }
            for fi64 := 0; fi64 < 4; fi64++ {
                v50 = (!v50)
                fmt.Println(v49, fi59)
                if (("\n\t" + ((("foo bar" + "") + "a") + "")) == (("a" + ("foo bar" + "foo bar")) + (("\n\t" + "foo bar") + ""))) {
                    v5 -= (v5 % ((v5) | 1))
                    if (-1 <= 1) {
                        if v50 {
                            v4 = (0.1 - 1.5)
                            v49 += uint64(v5)
                            if (!v50) {
                                switch 4294967295 {
                                case 2147483647:
                                    fmt.Println(fi64)
                                case 845:
                                    v65 := GenS1{F0: genFn0("", -227, "hello"), F1: 1, F2: genFn1((("foo bar" + "foo bar") + ("" + "foo bar")), "foo bar", -9223372036854775808)}
                                    _ = v65
                                case 44:
                                    var v66 int32 = -2147483648
                                    _ = v66
                                }
                                if v50 {
                                    if (0 >= int8(v5)) {
                                        if (uint(v3) >= 4294967295) {
                                            v50 = false
                                        } else {
                                            switch ("\n\t" + "hello") {
                                            case "foo bar":
                                                fmt.Println(v50, v3, v3)
                                            case "":
                                                v67 := GenS4{F0: 2147483647, F1: fi64}
                                                _ = v67
                                            case "a":
                                                fmt.Println(v49, v51, v3)
                                                _ = len(v2)
                                            default:
                                                v5 = v5
                                            }
                                        }
                                    } else {
                                        var v68 bool = true
                                        _ = v68
                                    }
                                    v69 := GenS3{F0: 9223372036854775807}
                                    _ = v69
                                } else {
                                    if (v50 || true) {
                                        fmt.Println(v49, v4, v4)
                                    }
                                }
                            }
                        } else {
                            switch (fi59 ^ -75) {
                            case 4294967295:
                                v5 += v5
                            case 748:
                                v70 := GenS5{F0: (fi59 | (^(fi59))), F1: "hello", F2: (^(v3))}
                                _ = v70
                            case 42:
                                var v71 float64 = float64(v4)
                                _ = v71
                            default:
                                var v72 string = "foo bar"
                                _ = v72
                            }
                            v73 := GenS6{F0: 2147483647, F1: 0, F2: fi59}
                            v50 = ((2147483648 >= genFn0("a", 2147483647, (("foo bar" + "hello") + "a"))) || (v5 != 0))
                            _ = v73
                        }
                    }
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v74 uint = 18446744073709551615
                        _ = v74
                    case "foo bar":
                        var v75 uint = (v5 ^ (v5 % ((v5) | 1)))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v75
                    }
                    fmt.Println(v4)
                }
                _ = fi64
            }
            _ = fi59
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
