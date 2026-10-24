package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := []string{"", "hello", "a", ("foo bar" + "hello")}
    var v2 uint32 = 2147483648
    var v3 int = 4294967295
    var v4 int64 = -683
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: v4, F1: (fi5 + v3), F2: 43}
        v7 := GenS2{F0: 1}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: v4, F1: v3, F2: -1}
    v9 := GenS3{F0: false}
    v10 := GenS1{F0: v4, F1: 4294967295, F2: 100}
    v3 = (v3 % (((v3 | -834)) | 1))
    switch "a" {
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            v4 -= v4
            _ = fi11
        }
    case "foo bar":
        if false {
            v3 = v3
            if false {
                for fi12 := 0; fi12 < 14; fi12++ {
                    v13 := GenS2{F0: 200}
                    var v14 uint = 255
                    var v15 float32 = 1.5
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                if (0.1 > (1e10 + 0.1)) {
                    switch v3 {
                    case -597:
                        var v16 uint16 = 0
                        _ = v16
                    case 1:
                        fmt.Println(v2, v2)
                    }
                    var v17 int32 = 0
                    _ = v17
                } else {
                    v2 = v2
                    v3 += (^(v3))
                }
            } else {
                var v18 uint = 1
                var v19 uint64 = 42
                _ = v18
                _ = v19
            }
        }
    }
    var v20 int16 = -603
    if true {
        v3 += 42
        switch ("foo bar" + "foo bar") {
        case "":
            fmt.Println(v2)
        case "\n\t":
            v1 = append(v1, "hello")
        case "foo bar":
            _ = len(v1)
        }
        for fi21 := 0; fi21 < 6; fi21++ {
            v1 = append(v1, "\n\t")
            _ = fi21
        }
    }
    if true {
        if true {
            for fi22 := 0; fi22 < 19; fi22++ {
                _ = len(v1)
                var v23 int64 = v4
                v4 -= (v4 * -886)
                _ = fi22
                _ = v23
            }
            var v24 uint16 = 32767
            if true {
                for fi25 := 0; fi25 < 20; fi25++ {
                    fmt.Println(v24, v24, v4)
                    v3 -= (v3 | (v3 << ((2) & 63)))
                    _ = fi25
                }
            } else {
                fmt.Println(v2, v3, v3)
                switch (v3 >> ((42) & 63)) {
                case 128:
                    v26 := GenS2{F0: v24}
                    _ = v26
                case 42:
                    if (true || false) {
                        var v27 uint64 = 4294967295
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        switch "foo bar" {
                        case "\n\t":
                            switch "" {
                            case "\n\t":
                                v28 := GenS4{F0: 409}
                                _ = v28
                            default:
                                if false {
                                    for fi29 := 0; fi29 < 3; fi29++ {
                                        v24 = (v24 ^ 65535)
                                        var v30 uint32 = (v2 ^ (v2 + (^(v2))))
                                        _ = fi29
                                        _ = v30
                                    }
                                    v31 := GenS5{F0: 0}
                                    for fi32 := 0; fi32 < 13; fi32++ {
                                        fmt.Println(v24, v24)
                                        if (!((-(v20)) > (v20 | v20))) {
                                            var v33 int32 = 42
                                            _ = v33
                                        }
                                        fmt.Println(v27, v3, v3)
                                        _ = fi32
                                    }
                                    _ = v31
                                }
                            }
                        case "hello":
                            var v34 uint64 = (v27 << ((65535) & 63))
                            _ = v34
                        default:
                            v2 += (v2 >> ((0) & 63))
                        }
                        _ = v27
                    } else {
                        var v35 float64 = 1e10
                        var v36 bool = (-2.25 > ((float32((((((((v35 - (v35 * float64(1e10))) - ((1.5 * float64((((-2.25 + 0.0) + 0.1) + -2.25))) * -2.25)) * ((((v35 / 3.14159) * 1.5) + 1e10) / v35)) - v35) * float64((-2.25 + 1e10))) - v35) - 1e10)) - float32(v35)) / 1.5))
                        v24 += (v24 % ((v24) | 1))
                        _ = v35
                        _ = v36
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                var v37 int = v3
                _ = v37
            }
            _ = v24
        } else {
            if (!true) {
                var v38 int = (v3 / ((v3) | 1))
                _ = v38
            }
            var v39 int = (v3 * (^(v3)))
            for fi40 := 0; fi40 < 5; fi40++ {
                v3 = (fi40 >> ((938) & 63))
                for fi41 := 0; fi41 < 16; fi41++ {
                    v42 := GenS6{F0: v2, F1: 42, F2: 416}
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
            _ = v39
        }
    } else {
        for fi43 := 0; fi43 < 4; fi43++ {
            if (42 == 128) {
                v44 := GenS6{F0: (v2 % ((v2) | 1)), F1: -1, F2: 37}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v45 int16 = v20
                _ = v44
                _ = v45
            } else {
                var v46 float64 = -2.25
                v3 = fi43
                _ = v46
            }
            fmt.Println(v3, v3, fi43)
            for fi47 := 0; fi47 < 7; fi47++ {
                if (false || false) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                switch ("\n\t" + "a") {
                case "foo bar":
                    v3 += v3
                }
                v48 := GenS4{F0: v4}
                switch fi47 {
                case 9223372036854775807:
                    v1 = append(v1, "a")
                case 655:
                    v1 = append(v1, "foo bar")
                default:
                    switch "foo bar" {
                    case "hello":
                        v2 = v2
                    }
                }
                _ = fi47
                _ = v48
            }
            _ = fi43
        }
        fmt.Println(v20)
        fmt.Println(v2, v4, v3)
    }
    v49 := GenS6{F0: (v2 << ((4294967295) & 63)), F1: 735, F2: 295}
    switch ("foo bar" + ("hello" + "a")) {
    case "hello":
        var v50 int32 = -2147483648
        _ = v50
    }
    if (255 <= 42) {
        v1 = append(v1, "a")
        fmt.Println(v20)
        var v51 uint16 = 2
        _ = v51
    } else {
        var v52 uint16 = 720
        v53 := GenS7{F0: 528}
        _ = v52
        _ = v53
    }
    var v54 int = -659
    v4 = (v4 - v4)
    for fi55 := 0; fi55 < 7; fi55++ {
        if (24 <= 128) {
            for fi56 := 0; fi56 < 5; fi56++ {
                v57 := GenS1{F0: (v4 ^ v4), F1: v3, F2: 0}
                var v58 uint = 255
                switch "\n\t" {
                case "":
                    var v59 float64 = float64(-2.25)
                    _ = v59
                default:
                    var v60 uint = 1
                    _ = v60
                }
                _ = fi56
                _ = v57
                _ = v58
            }
            var v61 int32 = 42
            var v62 float32 = ((0.1 / 3.14159) * -2.25)
            _ = v61
            _ = v62
        } else {
            for fi63 := 0; fi63 < 8; fi63++ {
                var v64 float32 = 0.0
                var v65 int64 = v4
                var v66 int = -9223372036854775808
                _ = fi63
                _ = v64
                _ = v65
                _ = v66
            }
            v4 += v4
        }
        _ = fi55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v49
    _ = v54
    _ = v8
    _ = v9
}
