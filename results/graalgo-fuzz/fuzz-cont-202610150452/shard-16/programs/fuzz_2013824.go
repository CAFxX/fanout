package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint; F1 uint; F2 string }
type GenS7 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    var v3 uint32 = 65535
    var v4 int32 = -875
    v5 := make(map[int]string)
    v6 := []string{(("foo bar" + ("foo bar" + "hello")) + ""), ""}
    var v7 int64 = 336
    if v2 {
        if false {
            v8 := GenS1{F0: 0}
            v9 := GenS2{F0: v7}
            if true {
                v2 = v2
                v5[32767] = "hello"
                v10 := GenS3{F0: 2, F1: 304}
                _ = v10
            }
            _ = v8
            _ = v9
        } else {
            if false {
                if v2 {
                    _ = len(v6)
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 float32 = 0.0
                        _ = fi11
                        _ = v12
                    }
                    var v13 float32 = 1.5
                    _ = v13
                }
            } else {
                v14 := GenS3{F0: 816, F1: 32767}
                if true {
                    v15 := GenS3{F0: 1, F1: 65535}
                    var v16 int = 127
                    _ = v15
                    _ = v16
                }
                _ = v14
            }
            var v17 int8 = -1
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 uint = 778
                var v20 int8 = 42
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        }
    } else {
        for fi21 := 0; fi21 < 12; fi21++ {
            switch fi21 {
            case 128:
                for fi22 := 0; fi22 < 2; fi22++ {
                    var v23 int8 = -128
                    if true {
                        var v24 int = (fi21 >> ((65535) & 63))
                        var v25 int8 = (^(v23))
                        _ = v24
                        _ = v25
                    } else {
                        if (0 >= 308) {
                            if v2 {
                                if ((float32(0.1) + (3.14159 * (0.0 - float32(1.5)))) <= 0.1) {
                                    if v2 {
                                        var v26 int32 = v4
                                        v4 = (v26 * -36)
                                        fmt.Println(v7)
                                        _ = v26
                                    }
                                } else {
                                    v1[v3] = ((0.0 / ((-2.25 + float32((-2.25 / float64(1.5)))) - 1e10)) > 1e10)
                                }
                            } else {
                                if true {
                                    if (true && v2) {
                                        fmt.Println(v4, v4)
                                    } else {
                                        _ = len(v5)
                                        var v27 uint32 = (v3 | v3)
                                        var v28 int8 = (^(v23))
                                        _ = v27
                                        _ = v28
                                    }
                                    delete(v5, 52)
                                }
                            }
                            v23 += v23
                        }
                        switch fi22 {
                        case 355:
                            v7 += 9223372036854775807
                            v2 = false
                        default:
                            _ = len(v6)
                        }
                        if (18446744073709551615 > 1) {
                            switch fi22 {
                            case 32767:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            v23 -= 111
                            fmt.Println(v4, fi21, v23)
                        }
                    }
                    var v29 string = "hello"
                    _ = fi22
                    _ = v23
                    _ = v29
                }
            }
            if len(v6) > 0 {
                v6[0] = "hello"
            }
            _ = fi21
        }
        var v30 int16 = 1
        for fi31 := 0; fi31 < 5; fi31++ {
            if (("\n\t" + "\n\t") >= "a") {
                switch fi31 {
                case 32767:
                    v3 += 2147483648
                }
                v32 := GenS4{F0: 255}
                _ = v32
            } else {
                fmt.Println(fi31, v2, v7)
            }
            _ = fi31
        }
        _ = v30
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        v4 += (v4 % ((-641) | 1))
        for fi34 := 0; fi34 < 5; fi34++ {
            fmt.Println(v4, v4)
            _ = fi34
        }
        _ = fi33
    }
    switch 94 {
    case -2147483648:
        switch 4294967295 {
        case 32767:
            if (!(!(!((!(0.1 < 3.14159)) || false)))) {
                var v35 uint16 = 65535
                switch 298 {
                case 4294967295:
                    for fi36 := 0; fi36 < 12; fi36++ {
                        var v37 uint = 659
                        _ = fi36
                        _ = v37
                    }
                default:
                    var v38 float64 = 1e10
                    _ = v38
                }
                _ = v35
            }
        case 32768:
            for fi39 := 0; fi39 < 7; fi39++ {
                fmt.Println(v7)
                for fi40 := 0; fi40 < 8; fi40++ {
                    if v2 {
                        v41 := GenS2{F0: (v7 & 0)}
                        v42 := GenS5{F0: 65535, F1: fi40}
                        v43 := GenS4{F0: 92}
                        _ = v41
                        _ = v42
                        _ = v43
                    } else {
                        fmt.Println(v4, v7, v3)
                        var v44 float32 = (1e10 / 3.14159)
                        fmt.Println(v4, v3, v2)
                        _ = v44
                    }
                    if ((((!(!v2)) || (1 < 0)) && v2) && v2) {
                        if (!true) {
                            _ = len(v6)
                            switch 2147483647 {
                            case -90:
                                var v45 string = ("foo bar" + ("hello" + ((((("\n\t" + ("a" + ("a" + "foo bar"))) + "foo bar") + "") + "a") + "\n\t")))
                                _ = v45
                            case 32767:
                                switch (("hello" + "foo bar") + "foo bar") {
                                case "a":
                                    var v46 bool = v2
                                    _ = v46
                                case "\n\t":
                                    v4 = (v4 >> ((800) & 63))
                                case "foo bar":
                                    var v47 float32 = -2.25
                                    _ = v47
                                default:
                                    v2 = (!v2)
                                }
                            default:
                                switch 42 {
                                case 360:
                                    v7 += (v7 - (v7 >> ((726) & 63)))
                                case -2147483648:
                                    v2 = ("\n\t" > "hello")
                                }
                            }
                            switch int(v3) {
                            case 2147483647:
                                var v48 uint64 = 18446744073709551615
                                _ = v48
                            case 503:
                                var v49 int16 = 277
                                _ = v49
                            case 4294967295:
                                v2 = v2
                            }
                        } else {
                            fmt.Println(v4)
                            v50 := GenS4{F0: 255}
                            switch fi39 {
                            case 74:
                                v4 = v4
                            case 0:
                                fmt.Println(v2, v2)
                            default:
                                var v51 uint32 = v3
                                fmt.Println(v3, v7, v2)
                                fmt.Println(v3, v3, v7)
                                _ = v51
                            }
                            _ = v50
                        }
                        v3 += v3
                    }
                    v52 := GenS5{F0: 11, F1: -2147483648}
                    _ = fi40
                    _ = v52
                }
                _ = fi39
            }
        case 0:
            fmt.Println(v7, v3)
        }
    }
    for fi53 := 0; fi53 < 15; fi53++ {
        if true {
            v7 -= v7
            for fi54 := 0; fi54 < 2; fi54++ {
                fmt.Println(fi53, v4, v2)
                switch 9223372036854775807 {
                case 127:
                    fmt.Println(v7, v2)
                case -233:
                    var v55 bool = v2
                    _ = v55
                case -9223372036854775808:
                    if (4294967295 > 42) {
                        var v56 uint = 42
                        v57 := GenS5{F0: (^(v56)), F1: 146}
                        fmt.Println(v4, v4, v4)
                        _ = v56
                        _ = v57
                    }
                default:
                    fmt.Println(fi53, v3)
                }
                var v58 float32 = ((1.5 - float32(0.1)) / float32((float64(0.0) - ((3.14159 * 3.14159) - 1e10))))
                _ = fi54
                _ = v58
            }
            switch "\n\t" {
            case "\n\t":
                var v59 float32 = 1e10
                _ = v59
            case "hello":
                var v60 string = "foo bar"
                _ = v60
            case "a":
                var v61 int32 = (v4 & v4)
                _ = v61
            }
        }
        switch (fi53 + fi53) {
        case 0:
            var v62 string = ("" + "hello")
            _ = v62
        case -614:
            var v63 uint8 = 2
            _ = v63
        default:
            fmt.Println(v3, v7)
        }
        v64 := GenS6{F0: 958, F1: 4294967295, F2: (("" + "\n\t") + (("hello" + "hello") + "hello"))}
        _ = fi53
        _ = v64
    }
    v7 -= 0
    v65 := GenS7{F0: 1, F1: 1}
    v2 = false
    v66 := GenS4{F0: 127}
    for fi67 := 0; fi67 < 19; fi67++ {
        var v68 int16 = 0
        switch 743 {
        case -2147483648:
            var v69 bool = (true || v2)
            _ = v69
        case 128:
            for fi70 := 0; fi70 < 9; fi70++ {
                v3 -= v3
                _ = fi70
            }
        }
        v68 = (v68 - 1000)
        _ = fi67
        _ = v68
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v7
}
