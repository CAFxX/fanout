package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint16; F1 bool; F2 int64 }

func main() {
    v1 := []string{("" + (("hello" + ("foo bar" + ("\n\t" + "a"))) + ("foo bar" + ""))), "\n\t", "hello"}
    v2 := []int{-9223372036854775808, 128, 1}
    var v3 uint8 = 185
    var v4 float64 = (1.5 / 0.1)
    if true {
        v3 = v3
        if true {
            for fi5 := 0; fi5 < 8; fi5++ {
                var v6 uint16 = 1
                if false {
                    if true {
                        if (0 <= 743) {
                            for fi7 := 0; fi7 < 18; fi7++ {
                                v3 -= 72
                                _ = fi7
                            }
                            switch fi5 {
                            case 992:
                                v8 := GenS1{F0: 216}
                                _ = v8
                            case 127:
                                for fi9 := 0; fi9 < 3; fi9++ {
                                    v3 = (v3 * (v3 & (v3 << ((255) & 63))))
                                    v10 := GenS1{F0: 1}
                                    _ = fi9
                                    _ = v10
                                }
                            case -126:
                                var v11 float64 = (v4 + 1.5)
                                _ = v11
                            }
                        }
                        v6 -= (^(v6))
                    }
                    v12 := GenS1{F0: 469}
                    _ = v12
                }
                _ = fi5
                _ = v6
            }
            v3 += (^(v3))
            v13 := GenS2{F0: -2147483648, F1: 127, F2: 44}
            _ = v13
        }
        v4 -= ((3.14159 / -2.25) / (((float64((1.5 + 1.5)) - (float64(3.14159) - ((v4 * v4) - 0.1))) / 1.0) * v4))
    } else {
        v3 = 19
    }
    var v14 uint16 = 895
    if (777 < 1) {
        v15 := GenS1{F0: 668}
        if false {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 float64 = v4
                _ = fi16
                _ = v17
            }
        } else {
            fmt.Println(v14, v3, v14)
            v3 -= v3
        }
        v3 = v3
        _ = v15
    } else {
        v1 = append(v1, "hello")
        v18 := GenS1{F0: 727}
        if true {
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 bool = true
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 uint8 = 1
                    v23 := GenS3{F0: 4, F1: 42}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = fi19
                _ = v20
            }
        }
        _ = v18
    }
    v14 += v14
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS4{F0: (v14 - v14), F1: (fi24 != 32767), F2: 71}
        _ = fi24
        _ = v25
    }
    if (("\n\t" + ("" + "hello")) == "") {
        _ = len(v2)
    }
    if (3.14159 >= (3.14159 - 0.1)) {
        switch int(v14) {
        case 128:
            v14 = (v14 - 674)
        case -517:
            if false {
                var v26 uint16 = 2
                _ = v26
            } else {
                fmt.Println(v4, v14, v14)
                var v27 int = 32767
                for fi28 := 0; fi28 < 6; fi28++ {
                    fmt.Println(v4, v14, v4)
                    var v29 int = fi28
                    var v30 string = "a"
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = v27
            }
        case 32767:
            v3 += (^(v3))
        default:
            var v31 int16 = 32767
            _ = v31
        }
        var v32 float64 = 0.0
        _ = len(v1)
        _ = v32
    } else {
        for fi33 := 0; fi33 < 16; fi33++ {
            if len(v1) > 0 {
                v1[0] = "a"
            }
            _ = fi33
        }
        v34 := GenS1{F0: 0}
        _ = v34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        fmt.Println(fi35, v14)
        _ = fi35
    }
    v36 := GenS1{F0: 255}
    var v37 uint = 951
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 int16 = 405
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        for fi41 := 0; fi41 < 19; fi41++ {
            switch (("a" + ("a" + "foo bar")) + "foo bar") {
            case "":
                if true {
                    fmt.Println(v3, fi41)
                    if false {
                        var v42 int16 = -1
                        _ = v42
                    } else {
                        var v43 uint64 = 0
                        _ = v43
                    }
                    var v44 uint64 = 4294967295
                    _ = v44
                } else {
                    var v45 bool = true
                    _ = v45
                }
                v46 := GenS2{F0: fi41, F1: fi40, F2: uint64(v3)}
                switch ("a" + "\n\t") {
                case "a":
                    var v47 uint8 = (v3 / ((v3) | 1))
                    _ = v47
                default:
                    v14 += (^(v14))
                }
                var v48 float32 = (0.0 / ((1.5 * 1e10) / -2.25))
                v49 := GenS1{F0: (^(v37))}
                _ = v46
                _ = v48
                _ = v49
            case "a":
                _ = len(v2)
            case "foo bar":
                var v50 uint8 = v3
                _ = v50
            default:
                v2 = append(v2, -9223372036854775808)
            }
            _ = fi41
        }
        _ = fi40
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
