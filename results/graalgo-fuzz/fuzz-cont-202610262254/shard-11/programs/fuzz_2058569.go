package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int64; F2 bool }
type GenS5 struct { F0 uint16; F1 uint32; F2 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32; F1 string; F2 uint }

func main() {
    v1 := []string{("\n\t" + "foo bar")}
    v2 := []uint32{2147483648, 4294967295, 2147483647, 42}
    var v3 string = ""
    var v4 int64 = 0
    var v5 bool = ((829 < 457) && true)
    if false {
        if (127 == -616) {
            v3 = v3
            v6 := GenS1{F0: 4294967295}
            _ = v6
        }
        var v7 int32 = 42
        v5 = (4294967295 <= 4294967295)
        _ = v7
    } else {
        v3 = v3
        v8 := GenS1{F0: 4294967295}
        var v9 uint32 = 0
        _ = v8
        _ = v9
    }
    v4 = v4
    if ("a" > v3) {
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := GenS2{F0: ((v4 | v4) >= v4), F1: 0}
                var v13 float32 = ((0.0 + (1.5 - (float32(float64((0.1 / (3.14159 + 0.0)))) + 1.5))) * (float32((0.1 * 3.14159)) / (3.14159 - (1e10 - (0.1 * 0.0)))))
                v14 := GenS3{F0: 346}
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            if v5 {
                v15 := GenS1{F0: 636}
                v16 := GenS2{F0: true, F1: -336}
                var v17 int = (fi10 ^ 617)
                _ = v15
                _ = v16
                _ = v17
            } else {
                v1 = append(v1, ("\n\t" + "foo bar"))
                if v5 {
                    if (328 > 65535) {
                        v2 = append(v2, 2147483648)
                        if len(v2) > 0 {
                            v2[0] = 4294967295
                        }
                        if true {
                            var v18 uint16 = 255
                            _ = v18
                        }
                    } else {
                        switch "foo bar" {
                        case "a":
                            if (!((1 != 1) && v5)) {
                                var v19 int32 = 2147483647
                                switch "a" {
                                case "":
                                    for fi20 := 0; fi20 < 18; fi20++ {
                                        v4 = v4
                                        var v21 uint8 = 97
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = fi20
                                        _ = v21
                                    }
                                case "hello":
                                    if v5 {
                                        var v22 float64 = 1e10
                                        _ = v22
                                    }
                                }
                                _ = v19
                            } else {
                                if true {
                                    if ((fi10 & fi10) < fi10) {
                                        var v23 uint32 = 2147483648
                                        _ = v23
                                    } else {
                                        for fi24 := 0; fi24 < 7; fi24++ {
                                            switch v3 {
                                            case "a":
                                                _ = len(v1)
                                            }
                                            _ = fi24
                                        }
                                        v3 = (v3 + (v3 + v3))
                                        v3 = "foo bar"
                                    }
                                    if (!false) {
                                        v25 := GenS1{F0: 848}
                                        _ = v25
                                    }
                                    fmt.Println(fi10, fi10, v4)
                                } else {
                                    var v26 int32 = 42
                                    fmt.Println(v26, v26)
                                    _ = v26
                                }
                                var v27 int16 = 0
                                var v28 float32 = -2.25
                                _ = v27
                                _ = v28
                            }
                        case "\n\t":
                            v29 := GenS3{F0: 65535}
                            _ = v29
                        case "hello":
                            var v30 uint32 = 318
                            _ = v30
                        default:
                            if v5 {
                                v31 := GenS2{F0: false, F1: -2147483648}
                                v3 = v3
                                _ = len(v1)
                                _ = v31
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                            }
                        }
                        v4 = (-(v4))
                        if (!(false && v5)) {
                            v32 := GenS1{F0: 42}
                            _ = v32
                        } else {
                            fmt.Println(v3, v4)
                        }
                    }
                    var v33 uint32 = 4294967295
                    _ = v33
                }
            }
            _ = fi10
        }
    } else {
        if true {
            var v34 uint16 = 32767
            _ = v34
        } else {
            v35 := GenS1{F0: 42}
            var v36 uint = 831
            if (!(v5 && false)) {
                var v37 int64 = -1
                for fi38 := 0; fi38 < 16; fi38++ {
                    for fi39 := 0; fi39 < 2; fi39++ {
                        v4 -= (^(v4))
                        _ = fi39
                    }
                    v4 -= (v4 ^ (v4 % ((9223372036854775807) | 1)))
                    v40 := GenS4{F0: 400, F1: (v4 >> ((820) & 63)), F2: v5}
                    _ = fi38
                    _ = v40
                }
                if v5 {
                    _ = len(v2)
                    switch (v3 + v3) {
                    case "\n\t":
                        if ((("a" + (((v3 + "a") + v3) + (v3 + "a"))) + v3) >= (v3 + (v3 + "hello"))) {
                            v1 = append(v1, "\n\t")
                        } else {
                            v2 = append(v2, 4294967295)
                            var v41 uint = (v36 - (v36 % (((v36 / ((v36) | 1))) | 1)))
                            for fi42 := 0; fi42 < 3; fi42++ {
                                v2 = append(v2, 2147483647)
                                var v43 int16 = 32767
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi42
                                _ = v43
                            }
                            _ = v41
                        }
                    case "":
                        var v44 float64 = 1e10
                        _ = v44
                    case "a":
                        v2 = append(v2, 65535)
                    default:
                        for fi45 := 0; fi45 < 15; fi45++ {
                            var v46 uint8 = 0
                            if v5 {
                                v47 := GenS5{F0: 255, F1: 2147483647, F2: 127}
                                v48 := GenS6{F0: 656}
                                _ = v47
                                _ = v48
                            } else {
                                v4 = (v4 ^ 0)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = fi45
                            _ = v46
                        }
                    }
                }
                _ = v37
            }
            _ = v35
            _ = v36
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (v3 + v3)
    var v49 float64 = float64(0.0)
    var v50 uint = 42
    if v5 {
        var v51 uint64 = 1
        fmt.Println(v3, v51)
        _ = v51
    } else {
        if (-1 != 38) {
            if len(v1) > 0 {
                v1[0] = ""
            }
        } else {
            v49 -= v49
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v3 = (v3 + "hello")
                var v52 uint64 = 0
                _ = v52
            } else {
                for fi53 := 0; fi53 < 19; fi53++ {
                    v54 := GenS6{F0: 1}
                    if true {
                        v3 = v3
                    } else {
                        v49 = v49
                    }
                    v55 := GenS1{F0: 18446744073709551615}
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                fmt.Println(v4)
            }
        }
        switch 4294967295 {
        case 128:
            for fi56 := 0; fi56 < 11; fi56++ {
                var v57 uint8 = 127
                _ = fi56
                _ = v57
            }
        case -2147483648:
            for fi58 := 0; fi58 < 6; fi58++ {
                fmt.Println(v49, v5)
                var v59 bool = (42 > 0)
                _ = fi58
                _ = v59
            }
        case 683:
            var v60 int8 = 30
            _ = v60
        default:
            switch -9223372036854775808 {
            case 42:
                if (1.5 > (float32(v49) - (3.14159 * 3.14159))) {
                    fmt.Println(v50, v5)
                }
            }
        }
    }
    v3 = "a"
    v61 := GenS7{F0: 1, F1: "foo bar", F2: v50}
    var v62 int32 = -1
    switch v3 {
    case "hello":
        v2 = append(v2, 2)
    case "foo bar":
        v5 = v5
    default:
        switch ("foo bar" + ((v3 + (v3 + "foo bar")) + v3)) {
        case "hello":
            _ = len(v2)
        default:
            v4 -= (v4 * 0)
        }
    }
    v5 = (!(v62 != v62))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v61
    _ = v62
}
