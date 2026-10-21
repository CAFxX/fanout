package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 uint8 }

func main() {
    var v1 bool = (-1 >= -1)
    var v2 string = ("" + "\n\t")
    var v3 int32 = -2147483648
    var v4 float64 = -2.25
    v5 := []int{int(v3), 0, -1}
    if v1 {
        v6 := GenS1{F0: 42, F1: 866}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v1 = true
        _ = v6
    } else {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint16 = 42
            _ = fi7
            _ = v8
        }
        if (!false) {
            v1 = (true || (!true))
        }
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v1, v1, fi9)
        v10 := GenS2{F0: "a"}
        _ = fi9
        _ = v10
    }
    if (86 <= 100) {
        v1 = ("\n\t" != ((((((v2 + "hello") + ((("foo bar" + "") + ("" + (((((((((v2 + v2) + "a") + "\n\t") + ("\n\t" + (v2 + ""))) + (v2 + (v2 + (v2 + v2)))) + (v2 + (("foo bar" + "hello") + v2))) + (v2 + v2)) + "\n\t") + "hello"))) + v2)) + v2) + v2) + ("" + (v2 + (v2 + v2)))) + v2))
        v11 := GenS3{F0: 42, F1: -11}
        if false {
            v2 = v2
            var v12 string = (v2 + v2)
            _ = v12
        } else {
            if len(v5) > 0 {
                v5[0] = 905
            }
        }
        _ = v11
    }
    v13 := GenS1{F0: 842, F1: 0}
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(v2, v3)
        _ = fi14
    }
    var v15 int32 = v3
    var v16 string = "hello"
    fmt.Println(v3)
    var v17 float64 = v4
    v2 = v16
    if v1 {
        if true {
            if false {
                for fi18 := 0; fi18 < 10; fi18++ {
                    v4 = 0.1
                    _ = fi18
                }
                switch v16 {
                case "hello":
                    switch v2 {
                    case "\n\t":
                        for fi19 := 0; fi19 < 1; fi19++ {
                            var v20 string = v2
                            _ = fi19
                            _ = v20
                        }
                    }
                case "a":
                    if len(v5) > 0 {
                        v5[0] = -550
                    }
                }
                var v21 int32 = -136
                _ = v21
            }
            if false {
                var v22 uint32 = 2147483647
                v1 = v1
                _ = v22
            } else {
                switch ("" + "") {
                case "hello":
                    v16 = v2
                default:
                    var v23 int16 = 32767
                    _ = v23
                }
                if false {
                    v17 -= 0.1
                }
                if ((v1 || true) || false) {
                    switch (("\n\t" + v2) + (v2 + v16)) {
                    case "a":
                        fmt.Println(v3, v15, v3)
                    case "hello":
                        if (!true) {
                            v4 = 0.1
                        }
                    case "foo bar":
                        if false {
                            var v24 float64 = v17
                            if (-1 != -7) {
                                switch -843 {
                                case 127:
                                    v3 = -2147483648
                                case 32767:
                                    for fi25 := 0; fi25 < 16; fi25++ {
                                        var v26 bool = ((!v1) || (((!((1 <= 42) || true)) || (((42 >= 0) && false) || (-296 == -9223372036854775808))) && (fi25 == (fi25 % (((fi25 ^ (fi25 / (((fi25 | 32768)) | 1)))) | 1)))))
                                        _ = fi25
                                        _ = v26
                                    }
                                case 946:
                                    fmt.Println(v1, v24, v16)
                                }
                            }
                            _ = v24
                        } else {
                            for fi27 := 0; fi27 < 4; fi27++ {
                                for fi28 := 0; fi28 < 5; fi28++ {
                                    var v29 int64 = -9223372036854775808
                                    var v30 uint32 = 1
                                    if len(v5) > 0 {
                                        v5[0] = fi27
                                    }
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                v2 = v16
                                _ = fi27
                            }
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
        } else {
            v16 = "\n\t"
            v3 -= v15
            var v31 bool = (255 < 0)
            _ = v31
        }
        var v32 uint8 = 128
        v15 += -1
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS4{F0: 255}
        var v35 float32 = (3.14159 + (((((-2.25 + ((1.5 + float32(0.1)) - 0.0)) - 0.1) / (0.1 * 1.5)) / (float32(1e10) / float32(1.5))) * 0.1))
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 string = ((v16 + (v16 + v2)) + "\n\t")
    v37 := GenS1{F0: 636, F1: 727}
    if v1 {
        for fi38 := 0; fi38 < 13; fi38++ {
            v39 := GenS5{F0: v1, F1: 508}
            _ = len(v5)
            v1 = true
            _ = fi38
            _ = v39
        }
        switch 0 {
        case 1:
            for fi40 := 0; fi40 < 2; fi40++ {
                var v41 bool = (!(255 != 0))
                var v42 float64 = 0.1
                _ = fi40
                _ = v41
                _ = v42
            }
        case 976:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            for fi43 := 0; fi43 < 16; fi43++ {
                var v44 int8 = -128
                v45 := GenS4{F0: 42}
                for fi46 := 0; fi46 < 7; fi46++ {
                    v47 := GenS6{F0: v15}
                    _ = fi46
                    _ = v47
                }
                _ = fi43
                _ = v44
                _ = v45
            }
        }
        for fi48 := 0; fi48 < 19; fi48++ {
            fmt.Println(v16, v3, v1)
            for fi49 := 0; fi49 < 1; fi49++ {
                if v1 {
                    var v50 uint8 = 42
                    v1 = v1
                    _ = v50
                }
                if ((^(fi49)) < -278) {
                    v51 := GenS7{F0: -1, F1: 231}
                    _ = v51
                }
                v16 = v2
                _ = fi49
            }
            switch "" {
            case "a":
                if ((fi48 < fi48) && false) {
                    var v52 float32 = 3.14159
                    switch v2 {
                    case "foo bar":
                        v36 = v36
                    case "a":
                        switch 2147483647 {
                        case -399:
                            var v53 int8 = -128
                            _ = v53
                        case 32767:
                            v5 = append(v5, fi48)
                        default:
                            v36 = ((("foo bar" + "") + v2) + v2)
                        }
                    default:
                        if (634 == 358) {
                            v5 = append(v5, 9223372036854775807)
                            var v54 float64 = (((1.5 / float64(-2.25)) - ((-2.25 * v4) / v4)) + v4)
                            _ = v54
                        }
                    }
                    _ = v52
                } else {
                    v15 += v3
                    v55 := GenS4{F0: 1}
                    if false {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    _ = v55
                }
            case "hello":
                for fi56 := 0; fi56 < 15; fi56++ {
                    var v57 uint = 4294967295
                    var v58 uint16 = 1
                    _ = fi56
                    _ = v57
                    _ = v58
                }
            case "\n\t":
                fmt.Println(v3, v16)
            default:
                var v59 uint64 = 2
                _ = v59
            }
            _ = fi48
        }
    } else {
        _ = len(v5)
    }
    var v60 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v60
}
