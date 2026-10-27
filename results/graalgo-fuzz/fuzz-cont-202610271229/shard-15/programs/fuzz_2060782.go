package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int; F2 int8 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 bool; F1 int16; F2 int8 }

func main() {
    var v1 int16 = 1
    var v2 int16 = v1
    v3 := []string{"foo bar"}
    v2 = v1
    if (false || false) {
        _ = len(v3)
    }
    if true {
        fmt.Println(v2)
        var v4 float64 = 0.0
        _ = v4
    } else {
        var v5 uint64 = 860
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v5
    }
    if true {
        for fi6 := 0; fi6 < 2; fi6++ {
            fmt.Println(v2)
            v2 += v1
            _ = fi6
        }
        v3 = append(v3, ("\n\t" + (("hello" + "a") + (("a" + ("hello" + "a")) + "a"))))
        v7 := GenS1{F0: 18446744073709551615, F1: 42}
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 8; fi8++ {
            v2 -= (-(v1))
            v1 = v1
            v9 := GenS2{F0: -1}
            _ = fi8
            _ = v9
        }
        if true {
            v10 := GenS3{F0: 609, F1: 2147483647, F2: 100}
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int64 = 9223372036854775807
        _ = fi11
        _ = v12
    }
    fmt.Println(v2, v2)
    if false {
        v3 = append(v3, ((((("" + "") + ("\n\t" + "")) + "a") + "foo bar") + "a"))
    } else {
        var v13 float32 = 1e10
        var v14 uint32 = 0
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int = 32767
        if (!(384 > 4294967295)) {
            var v17 int = fi15
            _ = v17
        } else {
            if len(v3) > 0 {
                v3[0] = (("\n\t" + ("hello" + ("a" + "\n\t"))) + "\n\t")
            }
        }
        _ = fi15
        _ = v16
    }
    var v18 uint32 = 0
    if false {
        if true {
            v3 = append(v3, "\n\t")
            for fi19 := 0; fi19 < 18; fi19++ {
                if ((v1 / ((v1) | 1)) == 32767) {
                    if (0 > 1) {
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v2 = v1
                            if (42 > 70) {
                                if (127 == 0) {
                                    v3 = append(v3, "foo bar")
                                    v3 = append(v3, "foo bar")
                                    v21 := GenS2{F0: 42}
                                    _ = v21
                                } else {
                                    var v22 string = ("" + "hello")
                                    _ = v22
                                }
                                if (!false) {
                                    if false {
                                        _ = len(v3)
                                        switch "\n\t" {
                                        case "":
                                            switch -127 {
                                            case 42:
                                                fmt.Println(v1)
                                            case -9223372036854775808:
                                                var v23 float32 = (float32(0.0) * (1e10 / (1e10 * 1e10)))
                                                _ = v23
                                            }
                                        }
                                        if false {
                                            if false {
                                                var v24 float64 = 1e10
                                                _ = v24
                                            } else {
                                                _ = len(v3)
                                                v2 = 1000
                                                var v25 int16 = v1
                                                _ = v25
                                            }
                                            if (!(2147483647 <= -2147483648)) {
                                                var v26 uint64 = 0
                                                v1 -= -32768
                                                _ = v26
                                            }
                                            v3 = append(v3, "hello")
                                        }
                                    } else {
                                        if (0 <= 127) {
                                            if false {
                                                var v27 uint32 = (^(v18))
                                                _ = v27
                                            }
                                            var v28 string = "foo bar"
                                            var v29 int16 = (v2 - -1)
                                            var v30 string = ("\n\t" + (v28 + "foo bar"))
                                            var v31 int64 = 1
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                            _ = v31
                                        }
                                        if (!(fi20 > fi20)) {
                                            v3 = append(v3, "")
                                            switch ("" + "hello") {
                                            case "hello":
                                                v1 = -1
                                            case "":
                                                if len(v3) > 0 {
                                                    v3[0] = ""
                                                }
                                            }
                                            v1 += v2
                                        } else {
                                            if true {
                                                v18 = v18
                                            } else {
                                                var v32 uint8 = 42
                                                _ = v32
                                            }
                                            switch "" {
                                            case "":
                                                v2 = v1
                                            }
                                            v33 := GenS2{F0: 1}
                                            _ = v33
                                        }
                                        fmt.Println(fi20, v2)
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = ""
                                    }
                                } else {
                                    fmt.Println(v18, v1)
                                }
                            } else {
                                fmt.Println(v18)
                            }
                            v1 += v1
                            fmt.Println(fi20, v2, v1)
                            v1 = 32767
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            _ = fi20
                        }
                        if (v18 < v18) {
                            v18 = (v18 + 2147483647)
                        }
                        v1 += 1
                    } else {
                        _ = len(v3)
                    }
                    for fi34 := 0; fi34 < 15; fi34++ {
                        var v35 int32 = 0
                        v36 := GenS1{F0: 618, F1: 821}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    if (false || true) {
                        fmt.Println(v1)
                        var v37 uint16 = 1
                        _ = v37
                    } else {
                        v38 := GenS4{F0: (("" + "hello") + ("\n\t" + "foo bar")), F1: v1}
                        _ = v38
                    }
                }
                if false {
                    var v39 float64 = (-2.25 / float64((0.0 + (0.1 * -2.25))))
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 18; fi40++ {
                        switch "\n\t" {
                        case "a":
                            v1 = int16(fi19)
                        case "hello":
                            fmt.Println(v1, fi19, fi19)
                        }
                        v2 = v1
                        _ = fi40
                    }
                }
                var v41 uint64 = 42
                _ = fi19
                _ = v41
            }
        }
    }
    fmt.Println(v18, v18)
    v18 -= v18
    for fi42 := 0; fi42 < 19; fi42++ {
        if (((0.0 * -2.25) / 1.0) >= 1.5) {
            v1 -= (v2 % ((v2) | 1))
            for fi43 := 0; fi43 < 18; fi43++ {
                v1 += v1
                var v44 string = ""
                fmt.Println(v44, fi43, v44)
                _ = fi43
                _ = v44
            }
        }
        v45 := GenS5{F0: (fi42 <= (fi42 | (-(fi42)))), F1: (v2 | v2), F2: 1}
        var v46 int = (fi42 & (fi42 | 42))
        _ = fi42
        _ = v45
        _ = v46
    }
    var v47 int16 = (v1 << ((176) & 63))
    var v48 uint8 = 45
    v48 = 2
    var v49 float32 = float32((float64(3.14159) / 1.0))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
}
