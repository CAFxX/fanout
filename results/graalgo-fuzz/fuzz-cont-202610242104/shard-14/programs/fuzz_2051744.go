package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint }

func main() {
    v1 := []string{("hello" + "a"), "", "hello"}
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    switch "hello" {
    case "a":
        v4 := GenS1{F0: 32767, F1: 42}
        _ = v4
    case "":
        var v5 uint8 = 2
        _ = v5
    case "hello":
        var v6 bool = false
        _ = v6
    default:
        v7 := GenS1{F0: 2, F1: 255}
        _ = v7
    }
    if ("\n\t" != "\n\t") {
        var v8 string = (("foo bar" + ((("" + ("foo bar" + "a")) + "") + "\n\t")) + "")
        _ = v8
    } else {
        if len(v1) > 0 {
            v1[0] = ("a" + ("hello" + ((("a" + "a") + ("hello" + "\n\t")) + "hello")))
        }
        var v9 uint16 = 255
        _ = v9
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (-27 <= 122) {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint64 = 878
            fmt.Println(v11)
            delete(v2, 2147483648)
            _ = fi10
            _ = v11
        }
        if false {
            var v12 float64 = 1e10
            v13 := GenS1{F0: 65535, F1: 252}
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 15; fi14++ {
                if (((-2.25 / (0.1 / 0.1)) > 1e10) || (((!true) && false) || (94 >= 4294967295))) {
                    var v15 int = (fi14 << ((2) & 63))
                    _ = v15
                }
                _ = fi14
            }
        }
        var v16 uint16 = 1
        _ = v16
    } else {
        var v17 string = "\n\t"
        switch (v17 + ((v17 + (v17 + (v17 + ("a" + v17)))) + v17)) {
        case "hello":
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS1{F0: 255, F1: 920}
                v17 = v17
                _ = fi18
                _ = v19
            }
        case "":
            if ("a" != v17) {
                for fi20 := 0; fi20 < 15; fi20++ {
                    v21 := v2[345]
                    v17 = v17
                    v17 = v17
                    _ = fi20
                    _ = v21
                }
            }
        case "\n\t":
            for fi22 := 0; fi22 < 14; fi22++ {
                switch "\n\t" {
                case "hello":
                    delete(v3, "hello")
                default:
                    if (31 <= 42) {
                        for fi23 := 0; fi23 < 19; fi23++ {
                            var v24 float64 = 0.1
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            _ = fi23
                            _ = v24
                        }
                        var v25 bool = true
                        _ = v25
                    }
                }
                v2[42] = (^(fi22))
                _ = fi22
            }
        default:
            delete(v2, 19)
        }
        var v26 int64 = 1
        _ = v17
        _ = v26
    }
    var v27 float32 = (0.1 + (float32(((1e10 * 3.14159) * ((0.0 + float64(0.0)) - ((float64((-2.25 + (0.0 / 1.0))) * 3.14159) / (3.14159 - -2.25))))) * -2.25))
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := v2[2]
        _ = fi28
        _ = v29
    }
    if true {
        v27 = 3.14159
        if ((-105 >= 0) || (128 != 1)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v27)
            if (879 < -148) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v30 := GenS2{F0: 0}
                if (1.5 <= (0.0 * 0.1)) {
                    v31 := GenS2{F0: -630}
                    var v32 float64 = (-2.25 + 1.5)
                    _ = v31
                    _ = v32
                } else {
                    v33 := GenS1{F0: 65535, F1: 57}
                    v34 := GenS1{F0: 133, F1: 1}
                    fmt.Println(v27)
                    _ = v33
                    _ = v34
                }
                var v35 string = "foo bar"
                _ = v30
                _ = v35
            }
        }
        delete(v3, "")
    } else {
        var v36 int64 = 0
        fmt.Println(v36)
        v27 = v27
        _ = v36
    }
    v27 = v27
    v27 = v27
    switch "" {
    case "a":
        v27 -= v27
    default:
        var v37 float64 = -2.25
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 uint8 = 240
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 15; fi40++ {
        var v41 string = "a"
        v42 := GenS2{F0: 42}
        _ = fi40
        _ = v41
        _ = v42
    }
    if true {
        if true {
            v27 += (v27 / v27)
        } else {
            switch 127 {
            case 1:
                switch -820 {
                case 4294967295:
                    if true {
                        v43 := GenS2{F0: -32768}
                        var v44 int64 = 9223372036854775807
                        var v45 int = -233
                        _ = v43
                        _ = v44
                        _ = v45
                    } else {
                        fmt.Println(v27)
                    }
                case 32767:
                    switch 32768 {
                    case -2147483648:
                        var v46 int64 = 1
                        _ = v46
                    case 2147483647:
                        fmt.Println(v27)
                    case 4294967295:
                        if true {
                            for fi47 := 0; fi47 < 10; fi47++ {
                                if false {
                                    var v48 float64 = ((1.5 - -2.25) * (1.5 - -2.25))
                                    v27 -= float32(0.1)
                                    _ = v48
                                } else {
                                    v49 := GenS2{F0: 42}
                                    v27 += float32(float64(-2.25))
                                    v27 = (v27 + (((v27 * (v27 * 0.0)) - float32(-2.25)) * v27))
                                    _ = v49
                                }
                                var v50 int16 = 1
                                var v51 uint16 = 451
                                _ = fi47
                                _ = v50
                                _ = v51
                            }
                            if (42 == 37) {
                                for fi52 := 0; fi52 < 1; fi52++ {
                                    v53 := GenS3{F0: ("hello" + ("a" + "\n\t")), F1: 42}
                                    var v54 string = "foo bar"
                                    fmt.Println(v54)
                                    _ = fi52
                                    _ = v53
                                    _ = v54
                                }
                                var v55 uint32 = 2147483647
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                _ = v55
                            }
                            for fi56 := 0; fi56 < 10; fi56++ {
                                v1 = append(v1, "a")
                                fmt.Println(fi56, v27)
                                _ = fi56
                            }
                        }
                    default:
                        if (true && true) {
                            if true {
                                v57 := GenS1{F0: 255, F1: 1}
                                switch 9223372036854775807 {
                                case -1:
                                    var v58 uint = 18446744073709551615
                                    _ = v58
                                case 137:
                                    v59 := GenS1{F0: 147, F1: 0}
                                    _ = v59
                                case -9223372036854775808:
                                    v60 := GenS2{F0: 662}
                                    _ = v60
                                }
                                _ = v57
                            } else {
                                var v61 int64 = 963
                                v27 = 1e10
                                _ = v61
                            }
                            v62 := GenS4{F0: 255}
                            _ = v62
                        }
                    }
                case -2147483648:
                    var v63 bool = true
                    _ = v63
                }
            }
            fmt.Println(v27)
            fmt.Println(v27)
        }
        v27 = ((v27 + (v27 - 1e10)) - (v27 * (3.14159 / (v27 + 0.1))))
        v27 = v27
    } else {
        v27 = (1.5 - v27)
    }
    v3["a"] = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
}
