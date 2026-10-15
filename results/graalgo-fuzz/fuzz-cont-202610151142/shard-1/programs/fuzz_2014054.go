package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS8 struct { F0 int64; F1 int32 }
type GenS9 struct { F0 string }
func genFn0() uint32 {
    v6 := []int{42, 128, -102}
    var v7 int64 = 844
    _ = v6
    _ = v7
    return 2147483648
}

func main() {
    var v1 int16 = 0
    v2 := []string{"a", "hello", "hello", ("" + ("foo bar" + "\n\t"))}
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v5 int64 = 0
    if false {
        v5 = (v5 >> ((143) & 63))
        v8 := GenS1{F0: 2, F1: 127, F2: 42}
        var v9 uint8 = 157
        _ = v8
        _ = v9
    } else {
        v10 := GenS2{F0: 0}
        var v11 int = -1
        _ = v10
        _ = v11
    }
    if true {
        fmt.Println(v1, v1)
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            v5 += 42
            switch ("\n\t" + ("" + "")) {
            case "\n\t":
                switch "foo bar" {
                case "\n\t":
                    v13 := v3[fi12]
                    _ = v13
                case "a":
                    v14 := GenS1{F0: 161, F1: (fi12 & fi12), F2: 0}
                    _ = v14
                default:
                    var v15 int8 = -1
                    _ = v15
                }
            case "hello":
                v5 = v5
            default:
                for fi16 := 0; fi16 < 16; fi16++ {
                    v1 = 603
                    _ = fi16
                }
            }
            _ = fi12
        }
        if ("hello" < "") {
            if false {
                if ((-(v1)) > v1) {
                    v17 := GenS3{F0: 65535}
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    _ = v17
                }
                fmt.Println(v1, v5)
            } else {
                if (42 == 128) {
                    if false {
                        switch 4294967295 {
                        case 32768:
                            v1 += -1
                        case -464:
                            for fi18 := 0; fi18 < 3; fi18++ {
                                switch "a" {
                                case "\n\t":
                                    v19 := GenS4{F0: 255, F1: fi18}
                                    _ = v19
                                }
                                _ = fi18
                            }
                        case 4294967295:
                            for fi20 := 0; fi20 < 3; fi20++ {
                                var v21 int64 = 1
                                _ = fi20
                                _ = v21
                            }
                        }
                        var v22 uint64 = 2
                        v5 -= int64(v1)
                        _ = v22
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    v1 += (v1 % ((-1) | 1))
                    v23 := GenS5{F0: 0}
                    _ = v23
                }
                v5 -= v5
            }
            v1 = (v1 * (-(v1)))
            v3[9223372036854775807] = -131
        }
    }
    v1 += v1
    var v24 int32 = 905
    for fi25 := 0; fi25 < 5; fi25++ {
        if false {
            for fi26 := 0; fi26 < 2; fi26++ {
                if (-1 == -39) {
                    if (!((1 > 1) || (255 < 42))) {
                        if (true && ((false && true) || true)) {
                            if true {
                                var v27 uint8 = 51
                                if false {
                                    v4[(fi26 % ((-1) | 1))] = (!true)
                                }
                                var v28 string = "a"
                                _ = v27
                                _ = v28
                            }
                            switch 32768 {
                            case -1:
                                if (0.1 < 0.0) {
                                    v1 = -1
                                    fmt.Println(fi25, v24, fi25)
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                } else {
                                    var v29 float32 = 1.5
                                    var v30 uint = 42
                                    v31 := GenS4{F0: 255, F1: 32768}
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch 9223372036854775807 {
                                case 4294967295:
                                    if (false && true) {
                                        v5 -= (v5 | v5)
                                        var v32 int32 = -2147483648
                                        _ = v32
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case "\n\t":
                                            if len(v2) > 0 {
                                                v2[0] = "a"
                                            }
                                        case "foo bar":
                                            if (true || true) {
                                                var v33 float64 = 0.1
                                                v33 -= (1.5 - -2.25)
                                                if (fi25 == -405) {
                                                    v34 := GenS3{F0: 1}
                                                    v35 := GenS3{F0: 65535}
                                                    switch (fi25 % (((fi26 & fi25)) | 1)) {
                                                    case -2147483648:
                                                        v36 := v4[fi26]
                                                        _ = v36
                                                    case -1:
                                                        var v37 string = ("foo bar" + ("hello" + ("foo bar" + "a")))
                                                        _ = v37
                                                    }
                                                    _ = v34
                                                    _ = v35
                                                }
                                                _ = v33
                                            }
                                        }
                                    }
                                default:
                                    v5 -= (v5 + v5)
                                }
                            }
                        }
                    }
                    fmt.Println(fi26, fi25)
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi26
            }
        }
        _ = fi25
    }
    var v38 string = "hello"
    v39 := GenS2{F0: (v24 & v24)}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    if ((1.5 + ((0.1 * float32(0.0)) * ((3.14159 * 0.1) * 3.14159))) >= (3.14159 - ((((float32(0.0) / ((0.1 + -2.25) * 1e10)) - 3.14159) / (1e10 - -2.25)) + 1e10))) {
        fmt.Println(v5, v38)
        for fi40 := 0; fi40 < 9; fi40++ {
            v2 = append(v2, ((("foo bar" + ((("foo bar" + ((v38 + v38) + v38)) + v38) + (("hello" + (v38 + ((((v38 + (v38 + "foo bar")) + ("\n\t" + ("hello" + v38))) + "hello") + v38))) + "foo bar"))) + (v38 + "a")) + "hello"))
            _ = fi40
        }
    } else {
        v3[42] = -9223372036854775808
        if (!true) {
            var v41 uint16 = 0
            _ = v41
        }
        for fi42 := 0; fi42 < 4; fi42++ {
            if true {
                switch fi42 {
                case -1:
                    if false {
                        var v43 int = fi42
                        var v44 int = (-(fi42))
                        _ = v43
                        _ = v44
                    } else {
                        var v45 string = "a"
                        _ = v45
                    }
                case 526:
                    for fi46 := 0; fi46 < 3; fi46++ {
                        var v47 int8 = 100
                        _ = fi46
                        _ = v47
                    }
                case -841:
                    if false {
                        v1 = 1
                        v5 += int64(v1)
                        v5 += (v5 ^ v5)
                    } else {
                        if (v5 < v5) {
                            v48 := GenS6{F0: -10, F1: 0}
                            var v49 bool = true
                            for fi50 := 0; fi50 < 3; fi50++ {
                                v51 := GenS7{F0: 127, F1: 4294967295, F2: (v1 / ((32767) | 1))}
                                _ = fi50
                                _ = v51
                            }
                            _ = v48
                            _ = v49
                        }
                        v5 = v5
                    }
                }
                if true {
                    v2 = append(v2, (v38 + "hello"))
                    v52 := GenS1{F0: 1, F1: fi42, F2: 178}
                    _ = v52
                } else {
                    for fi53 := 0; fi53 < 10; fi53++ {
                        if true {
                            v24 -= -1
                            _ = len(v2)
                        } else {
                            fmt.Println(fi53, v38, fi53)
                            _ = len(v2)
                            v54 := GenS8{F0: (v5 ^ v5), F1: v24}
                            _ = v54
                        }
                        v24 -= 42
                        fmt.Println(fi42, v5)
                        _ = fi53
                    }
                }
                var v55 string = v38
                _ = v55
            } else {
                v3[fi42] = (fi42 >> ((42) & 63))
                v1 += (v1 + v1)
                fmt.Println(v1)
            }
            var v56 string = (v38 + ("foo bar" + v38))
            v57 := GenS3{F0: 32767}
            _ = fi42
            _ = v56
            _ = v57
        }
    }
    for fi58 := 0; fi58 < 2; fi58++ {
        if (3.14159 != 1.5) {
            var v59 float32 = (3.14159 / ((float32(3.14159) / -2.25) * (1e10 / 0.1)))
            v60 := GenS9{F0: (("hello" + ((v38 + v38) + ("hello" + ("hello" + v38)))) + v38)}
            _ = v59
            _ = v60
        }
        v5 = (-(v5))
        var v61 int64 = 9223372036854775807
        _ = fi58
        _ = v61
    }
    v62 := GenS7{F0: 123, F1: 4294967295, F2: v1}
    v63 := GenS9{F0: "\n\t"}
    if (!((1e10 * (0.0 * 1e10)) <= 0.0)) {
        v64 := v3[-256]
        _ = v64
    }
    v24 = (v24 & (v24 ^ (^(v24))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v62
    _ = v63
}
