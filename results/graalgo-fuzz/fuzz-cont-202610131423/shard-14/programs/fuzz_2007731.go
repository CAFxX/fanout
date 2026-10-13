package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int; F2 int32 }
type GenS5 struct { F0 uint64; F1 int8; F2 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32 }
func genFn0() bool {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: -128}
        v7 := []int{fi5, 127}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int16 = 1
        _ = fi8
        _ = v9
    }
    var v10 uint = 65535
    if (32767 <= 442) {
        if (2147483647 == -1) {
            v11 := []int{1, 42, -921}
            if (69 > -67) {
                switch 836 {
                case -2147483648:
                    v10 = 65535
                default:
                    v10 -= 349
                }
                var v12 uint8 = 255
                var v13 bool = (0.0 >= (3.14159 * 1.5))
                _ = v12
                _ = v13
            }
            var v14 float32 = 0.1
            _ = v11
            _ = v14
        }
    }
    _ = v10
    return false
}

func main() {
    var v1 float64 = ((0.0 - 0.1) + -2.25)
    var v2 int = 0
    v3 := []string{(("\n\t" + ("a" + "")) + "\n\t")}
    var v4 int16 = 42
    var v15 bool = (!true)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        v15 = v15
        _ = fi16
    }
    var v17 uint16 = 2
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int32 = 877
        v15 = (v15 && (0 == 4294967295))
        if true {
            v20 := GenS2{F0: ("\n\t" + ("a" + "foo bar")), F1: 680, F2: (v17 & v17)}
            switch (fi18 + v2) {
            case -1:
                if ((v19 == 2147483647) && (255 > (v17 + 255))) {
                    v19 = v19
                    v21 := GenS2{F0: ("hello" + "\n\t"), F1: 42, F2: v17}
                    _ = v21
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v1 = (float64(-2.25) / (1e10 - ((v1 / v1) / float64((3.14159 / 1.0)))))
                    v15 = (v15 || (v15 && (100 != -38)))
                }
            case -526:
                if (!false) {
                    var v22 uint = 1
                    var v23 uint64 = 494
                    _ = v22
                    _ = v23
                }
            case 0:
                if len(v3) > 0 {
                    v3[0] = ("foo bar" + "\n\t")
                }
            default:
                switch fi18 {
                case 2147483647:
                    for fi24 := 0; fi24 < 12; fi24++ {
                        var v25 string = "\n\t"
                        if (-484 > (v19 - v19)) {
                            var v26 int8 = 42
                            var v27 int8 = -53
                            _ = v26
                            _ = v27
                        } else {
                            var v28 int16 = 42
                            _ = v28
                        }
                        _ = fi24
                        _ = v25
                    }
                default:
                    v2 -= fi18
                }
            }
            _ = v20
        } else {
            var v29 int8 = -1
            switch ("foo bar" + ("a" + "hello")) {
            case "\n\t":
                v30 := GenS2{F0: "hello", F1: 2, F2: 65535}
                _ = v30
            }
            _ = v29
        }
        _ = fi18
        _ = v19
    }
    var v31 uint64 = 4294967295
    for fi32 := 0; fi32 < 8; fi32++ {
        for fi33 := 0; fi33 < 4; fi33++ {
            if ("foo bar" >= "a") {
                if ((!v15) && (false || true)) {
                    var v34 int64 = 42
                    _ = v34
                } else {
                    fmt.Println(v31)
                    if false {
                        var v35 int16 = v4
                        fmt.Println(fi32, fi33)
                        _ = v35
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            switch ("foo bar" + ("a" + "\n\t")) {
                            case "\n\t":
                                var v36 int8 = -104
                                _ = v36
                            case "foo bar":
                                var v37 int8 = 127
                                v38 := GenS2{F0: "", F1: (^(v31)), F2: (v17 << ((496) & 63))}
                                _ = v37
                                _ = v38
                            }
                        }
                        var v39 uint16 = (v17 * v17)
                        var v40 int8 = -34
                        var v41 int = fi33
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                }
                v42 := GenS3{F0: (int8(v4) <= 127)}
                _ = v42
            }
            _ = fi33
        }
        v31 += (v31 | (v31 / (((^(v31))) | 1)))
        _ = fi32
    }
    if v15 {
        v43 := GenS4{F0: 0, F1: v2, F2: 11}
        _ = v43
    }
    v44 := GenS5{F0: (^(v31)), F1: 42, F2: 20}
    if v15 {
        v17 = v17
    } else {
        v31 = v31
        fmt.Println(v4, v15)
        if (!false) {
            switch (-(v2)) {
            case 0:
                v1 -= (v1 / v1)
            case 9223372036854775807:
                v17 = 496
            default:
                var v45 float32 = (0.0 - (0.0 + (3.14159 + 1.5)))
                _ = v45
            }
            switch "\n\t" {
            case "\n\t":
                v17 += v17
            }
        } else {
            for fi46 := 0; fi46 < 3; fi46++ {
                fmt.Println(v17, v1, v15)
                _ = fi46
            }
            v47 := GenS3{F0: (false && ("\n\t" != "hello"))}
            v48 := GenS2{F0: "a", F1: (v31 | (^(v31))), F2: v17}
            _ = v47
            _ = v48
        }
    }
    v49 := GenS4{F0: v17, F1: v2, F2: -134}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if v15 {
        if ((1 < 42) && (-1 > 1)) {
            if true {
                if ("hello" <= ((("a" + "hello") + ("a" + "hello")) + "")) {
                    var v50 int32 = 2147483647
                    var v51 int = v2
                    _ = v50
                    _ = v51
                } else {
                    var v52 float32 = 1.5
                    for fi53 := 0; fi53 < 9; fi53++ {
                        fmt.Println(v2)
                        _ = fi53
                    }
                    switch -1 {
                    case 877:
                        if false {
                            switch (v2 & 128) {
                            case 1:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case -553:
                                v54 := GenS6{F0: 65535}
                                _ = v54
                            case -9223372036854775808:
                                v55 := GenS7{F0: -1}
                                _ = v55
                            }
                        }
                    case -20:
                        var v56 int32 = -1
                        _ = v56
                    }
                    _ = v52
                }
            }
        } else {
            if v15 {
                v2 = 128
            }
            for fi57 := 0; fi57 < 17; fi57++ {
                for fi58 := 0; fi58 < 2; fi58++ {
                    var v59 uint = 2
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v3 = append(v3, (("\n\t" + ("hello" + ((("a" + ("hello" + "hello")) + (("hello" + "") + "hello")) + ("foo bar" + (("" + "a") + (("" + "\n\t") + "foo bar")))))) + ""))
                    _ = fi58
                    _ = v59
                }
                _ = fi57
            }
            if len(v3) > 0 {
                v3[0] = "a"
            }
        }
    } else {
        if (1 == 42) {
            if v15 {
                fmt.Println(v2, v2, v17)
            }
            var v60 int16 = 0
            var v61 int8 = 100
            _ = v60
            _ = v61
        }
        var v62 int8 = 70
        fmt.Println(v15)
        _ = v62
    }
    var v63 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v44
    _ = v49
    _ = v63
}
