package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS5 struct { F0 int16; F1 int8; F2 bool }
type GenS6 struct { F0 string; F1 uint; F2 int32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int }

func main() {
    v1 := []string{"a", ("hello" + "a"), "a", "\n\t"}
    var v2 int = 534
    var v3 int16 = -32768
    if (0 >= 4294967295) {
        if false {
            v3 += (-(v3))
        } else {
            var v4 string = "a"
            v5 := GenS1{F0: 230}
            _ = v4
            _ = v5
        }
    } else {
        v3 = v3
        if (true && false) {
            var v6 float64 = 3.14159
            _ = v6
        } else {
            var v7 float64 = (0.0 + (1e10 * 1e10))
            _ = v7
        }
        _ = len(v1)
    }
    var v8 int8 = 0
    var v9 int16 = v3
    switch ("\n\t" + "a") {
    case "foo bar":
        v3 = v3
    case "a":
        v1 = append(v1, ("\n\t" + "\n\t"))
    default:
        if false {
            if (!(true || (!true))) {
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 int8 = v8
                    if false {
                        if false {
                            var v12 float64 = 1.5
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 7; fi13++ {
                                var v14 uint8 = 52
                                v14 -= 128
                                v11 = (^(v11))
                                _ = fi13
                                _ = v14
                            }
                        }
                        v11 += int8(fi10)
                    } else {
                        v8 = v11
                    }
                    if true {
                        var v15 string = (("" + "a") + "\n\t")
                        if (!false) {
                            v11 = v11
                            fmt.Println(v8, v2)
                            v16 := GenS2{F0: 2}
                            _ = v16
                        }
                        _ = v15
                    }
                    _ = fi10
                    _ = v11
                }
                for fi17 := 0; fi17 < 9; fi17++ {
                    v3 += v3
                    _ = fi17
                }
            } else {
                v3 = 32767
                for fi18 := 0; fi18 < 18; fi18++ {
                    var v19 string = "\n\t"
                    if false {
                        var v20 uint16 = 42
                        var v21 uint8 = 15
                        var v22 int64 = -1
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    _ = fi18
                    _ = v19
                }
            }
            fmt.Println(v9, v8)
            if true {
                switch "hello" {
                case "foo bar":
                    v23 := GenS1{F0: 2}
                    _ = v23
                case "a":
                    _ = len(v1)
                case "\n\t":
                    if ((^(v2)) >= v2) {
                        if (-1 > 1) {
                            if (v3 >= v3) {
                                v24 := GenS2{F0: 4294967295}
                                _ = v24
                            } else {
                                for fi25 := 0; fi25 < 7; fi25++ {
                                    if ((v8 - v8) > (v8 >> ((2) & 63))) {
                                        v2 -= 32768
                                        if (false || false) {
                                            var v26 uint = 2
                                            _ = v26
                                        }
                                        if (false || false) {
                                            var v27 int = (fi25 * v2)
                                            if false {
                                                v1 = append(v1, "\n\t")
                                                fmt.Println(v3, v8)
                                            } else {
                                                for fi28 := 0; fi28 < 4; fi28++ {
                                                    v1 = append(v1, ((("" + ("\n\t" + "")) + ("\n\t" + "\n\t")) + "hello"))
                                                    switch 0 {
                                                    case -134:
                                                        var v29 uint64 = 418
                                                        _ = v29
                                                    case 32768:
                                                        if ((1 >= 1) && (!true)) {
                                                            if ((0 == 2) && true) {
                                                                switch v2 {
                                                                case 42:
                                                                    v30 := GenS1{F0: 2}
                                                                    _ = v30
                                                                case 1:
                                                                    v8 += v8
                                                                default:
                                                                    v1 = append(v1, "a")
                                                                    var v31 int32 = int32(v8)
                                                                    _ = v31
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = ""
                                                                }
                                                                if (75 > 255) {
                                                                    v32 := GenS3{F0: (v9 | v3), F1: 915}
                                                                    _ = v32
                                                                } else {
                                                                    v8 = -1
                                                                    v2 += (fi25 & fi28)
                                                                }
                                                            } else {
                                                                var v33 string = ((("" + ("a" + ("foo bar" + "foo bar"))) + "hello") + "")
                                                                _ = v33
                                                            }
                                                            switch v27 {
                                                            case 409:
                                                                fmt.Println(fi28, v9)
                                                            }
                                                            fmt.Println(v8)
                                                        }
                                                    }
                                                    v34 := GenS1{F0: 291}
                                                    _ = len(v1)
                                                    _ = fi28
                                                    _ = v34
                                                }
                                                for fi35 := 0; fi35 < 6; fi35++ {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    fmt.Println(v8, fi35)
                                                    _ = fi35
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            var v36 float64 = 0.1
                                            _ = v27
                                            _ = v36
                                        } else {
                                            v37 := GenS2{F0: 1}
                                            v38 := GenS4{F0: 750, F1: 956, F2: 1}
                                            _ = v37
                                            _ = v38
                                        }
                                    }
                                    _ = fi25
                                }
                                v2 += v2
                                for fi39 := 0; fi39 < 2; fi39++ {
                                    v1 = append(v1, ("foo bar" + ("a" + "foo bar")))
                                    if false {
                                        for fi40 := 0; fi40 < 9; fi40++ {
                                            fmt.Println(v8, v9, v9)
                                            _ = len(v1)
                                            _ = fi40
                                        }
                                        fmt.Println(v3, v8)
                                    }
                                    _ = fi39
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            var v41 int = v2
                            _ = v41
                        }
                    }
                default:
                    if false {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                    } else {
                        if (2 < 18446744073709551615) {
                            var v42 uint = 0
                            fmt.Println(v8)
                            for fi43 := 0; fi43 < 20; fi43++ {
                                fmt.Println(v2)
                                _ = len(v1)
                                _ = fi43
                            }
                            _ = v42
                        } else {
                            v44 := GenS5{F0: (^(v3)), F1: v8, F2: ((1.5 + 0.1) < -2.25)}
                            v45 := GenS6{F0: "foo bar", F1: 65535, F2: -2147483648}
                            _ = v44
                            _ = v45
                        }
                        fmt.Println(v3, v9)
                        if (127 >= 210) {
                            v9 -= (v3 | v3)
                            v1 = append(v1, "")
                        } else {
                            if (1 < 2147483647) {
                                fmt.Println(v9, v2)
                                var v46 float64 = -2.25
                                fmt.Println(v8, v3, v2)
                                _ = v46
                            }
                            v47 := GenS4{F0: -2147483648, F1: 1, F2: 2147483648}
                            _ = v47
                        }
                    }
                }
                _ = len(v1)
                v48 := GenS7{F0: 1}
                _ = v48
            } else {
                switch v2 {
                case -224:
                    v2 -= v2
                case 2147483647:
                    var v49 int = 128
                    _ = v49
                default:
                    var v50 uint64 = 33
                    _ = v50
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        } else {
            for fi51 := 0; fi51 < 13; fi51++ {
                fmt.Println(fi51)
                v1 = append(v1, "hello")
                if false {
                    for fi52 := 0; fi52 < 7; fi52++ {
                        var v53 bool = true
                        _ = fi52
                        _ = v53
                    }
                    if (454 > 423) {
                        fmt.Println(v8, fi51)
                    }
                }
                _ = fi51
            }
        }
    }
    var v54 int64 = 348
    var v55 int16 = 832
    var v56 int64 = (v54 * v54)
    var v57 uint16 = 2
    v8 -= (v8 & (-(v8)))
    var v58 int64 = (v54 % ((v56) | 1))
    for fi59 := 0; fi59 < 3; fi59++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v56, v58)
        _ = fi59
    }
    v60 := GenS8{F0: v2}
    if true {
        v55 = v9
    } else {
        for fi61 := 0; fi61 < 12; fi61++ {
            var v62 int64 = 1
            _ = fi61
            _ = v62
        }
    }
    var v63 uint64 = 18446744073709551615
    v57 += v57
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v63:", v63)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v60
    _ = v63
    _ = v8
    _ = v9
}
