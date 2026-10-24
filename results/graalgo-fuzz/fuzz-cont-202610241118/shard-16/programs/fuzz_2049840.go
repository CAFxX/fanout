package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint16 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int16; F1 int }

func main() {
    var v1 uint64 = 79
    v2 := []string{("" + ("foo bar" + "\n\t")), "\n\t"}
    var v3 float32 = 3.14159
    v1 = (v1 - (v1 | (v1 | 42)))
    var v4 uint64 = v1
    v4 = (v4 << ((18446744073709551615) & 63))
    if (2147483647 >= 42) {
        fmt.Println(v1, v4, v4)
    } else {
        var v5 string = ("\n\t" + "a")
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v5
    }
    switch "" {
    case "a":
        var v6 uint8 = 0
        _ = v6
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        if false {
            var v7 uint64 = v4
            switch "a" {
            case "foo bar":
                for fi8 := 0; fi8 < 13; fi8++ {
                    var v9 int32 = 2147483647
                    _ = fi8
                    _ = v9
                }
            case "hello":
                v10 := GenS1{F0: 0, F1: 255, F2: 2}
                _ = v10
            case "":
                v3 += v3
            default:
                switch "" {
                case "":
                    if false {
                        v7 = 6
                        switch 32767 {
                        case -1:
                            for fi11 := 0; fi11 < 4; fi11++ {
                                var v12 int = 0
                                var v13 string = "a"
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                        case 128:
                            for fi14 := 0; fi14 < 4; fi14++ {
                                if false {
                                    var v15 uint32 = 65535
                                    _ = v15
                                } else {
                                    switch (fi14 >> ((200) & 63)) {
                                    case -626:
                                        for fi16 := 0; fi16 < 1; fi16++ {
                                            if true {
                                                v17 := GenS2{F0: 0, F1: 1}
                                                if len(v2) > 0 {
                                                    v2[0] = "hello"
                                                }
                                                _ = v17
                                            }
                                            var v18 bool = (false && true)
                                            if v18 {
                                                v3 -= (3.14159 - v3)
                                            }
                                            _ = fi16
                                            _ = v18
                                        }
                                    case -2147483648:
                                        for fi19 := 0; fi19 < 19; fi19++ {
                                            var v20 uint32 = 2147483647
                                            _ = fi19
                                            _ = v20
                                        }
                                    }
                                    v21 := GenS2{F0: 42, F1: 32767}
                                    fmt.Println(fi14, fi14, v4)
                                    _ = v21
                                }
                                for fi22 := 0; fi22 < 1; fi22++ {
                                    v4 -= 2
                                    fmt.Println(v7)
                                    if false {
                                        var v23 float32 = (((1.5 - ((v3 - ((float32(float64(3.14159)) / float32(1e10)) * v3)) / v3)) + (v3 - v3)) * v3)
                                        var v24 int64 = -1
                                        _ = v23
                                        _ = v24
                                    } else {
                                        fmt.Println(v1)
                                        switch (fi22 * fi22) {
                                        case 1:
                                            var v25 uint = 4294967295
                                            var v26 uint32 = 608
                                            fmt.Println(v26, fi14, v1)
                                            _ = v25
                                            _ = v26
                                        case 4294967295:
                                            v27 := GenS2{F0: 42, F1: -32768}
                                            _ = v27
                                        case 490:
                                            v7 += v4
                                        default:
                                            v3 -= 3.14159
                                            v7 += (v4 ^ (v1 & (v4 ^ (^(v7)))))
                                            var v28 int8 = -128
                                            _ = v28
                                        }
                                    }
                                    fmt.Println(fi14)
                                    _ = fi22
                                }
                                _ = fi14
                            }
                        default:
                            fmt.Println(v7, v3)
                        }
                        v1 = v7
                    }
                case "foo bar":
                    _ = len(v2)
                case "a":
                    switch "foo bar" {
                    case "\n\t":
                        var v29 uint16 = 32767
                        _ = v29
                    }
                }
            }
            _ = v7
        }
    }
    v3 = float32(1.5)
    switch 965 {
    case 472:
        var v30 int32 = -2147483648
        _ = v30
    case 902:
        for fi31 := 0; fi31 < 2; fi31++ {
            if (!(!(100 != -25))) {
                v3 = -2.25
                v32 := GenS1{F0: 42, F1: 65535, F2: 0}
                var v33 bool = ((-9223372036854775808 == (fi31 * 201)) || true)
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 6; fi34++ {
                v4 = (v4 & 42)
                v4 = v4
                var v35 int16 = 1000
                var v36 int = 1
                _ = fi34
                _ = v35
                _ = v36
            }
            for fi37 := 0; fi37 < 20; fi37++ {
                var v38 int = (fi37 * fi31)
                _ = fi37
                _ = v38
            }
            _ = fi31
        }
    default:
        for fi39 := 0; fi39 < 8; fi39++ {
            v2 = append(v2, ((("hello" + "foo bar") + "") + "\n\t"))
            var v40 uint64 = v1
            _ = fi39
            _ = v40
        }
    }
    v4 = 284
    v1 = 519
    for fi41 := 0; fi41 < 8; fi41++ {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch 9223372036854775807 {
            case 4294967295:
                var v42 uint = 144
                _ = v42
            case 128:
                var v43 int8 = 100
                _ = v43
            default:
                if len(v2) > 0 {
                    v2[0] = ("a" + "\n\t")
                }
            }
            for fi44 := 0; fi44 < 11; fi44++ {
                fmt.Println(v4, v4, v3)
                v1 = (v4 << ((2) & 63))
                _ = fi44
            }
        } else {
            var v45 int16 = -32768
            _ = v45
        }
        var v46 bool = (974 < (^(v1)))
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 uint8 = 166
            v49 := GenS1{F0: 0, F1: 0, F2: 2}
            _ = fi47
            _ = v48
            _ = v49
        }
        _ = fi41
        _ = v46
    }
    if false {
        if false {
            switch "a" {
            case "a":
                if (v4 >= (^(v1))) {
                    if true {
                        fmt.Println(v4, v4)
                        fmt.Println(v4, v3, v3)
                    }
                } else {
                    v50 := GenS1{F0: 2147483647, F1: 270, F2: 920}
                    _ = v50
                }
            }
        } else {
            if (!true) {
                var v51 int16 = 0
                v2 = append(v2, "\n\t")
                _ = v51
            } else {
                for fi52 := 0; fi52 < 4; fi52++ {
                    v2 = append(v2, (("" + ("" + "")) + ("a" + ("a" + ("a" + "\n\t")))))
                    _ = fi52
                }
                fmt.Println(v4)
                var v53 int16 = 1
                _ = v53
            }
        }
        if (757 > 4294967295) {
            fmt.Println(v4)
        } else {
            if len(v2) > 0 {
                v2[0] = "a"
            }
            v54 := GenS3{F0: 0, F1: -1}
            _ = v54
        }
        for fi55 := 0; fi55 < 9; fi55++ {
            v56 := GenS1{F0: 2147483647, F1: 413, F2: 65535}
            if (!false) {
                v2 = append(v2, (("hello" + "hello") + "hello"))
            }
            v2 = append(v2, "foo bar")
            _ = fi55
            _ = v56
        }
    } else {
        v57 := GenS2{F0: 867, F1: -32768}
        _ = v57
    }
    for fi58 := 0; fi58 < 4; fi58++ {
        var v59 int = fi58
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        _ = fi58
        _ = v59
    }
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
