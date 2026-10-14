package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 4294967295
    v3 := make(map[int]uint32)
    v4 := []string{"\n\t", "\n\t", "a"}
    v5 := make(map[string]string)
    v6 := make(map[uint32]uint32)
    if false {
        if (false || (0 == 1)) {
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS1{F0: 0, F1: uint16(v2), F2: "\n\t"}
                _ = fi7
                _ = v8
            }
            var v9 int = 32767
            var v10 int16 = -1
            _ = v10
            _ = v9
        } else {
            _ = len(v1)
            v11 := GenS2{F0: 167}
            _ = v11
        }
        v2 -= (v2 | 65535)
        switch "foo bar" {
        case "hello":
            var v12 string = ""
            _ = v12
        default:
            v13 := GenS1{F0: 1, F1: 65535, F2: ("hello" + "a")}
            _ = v13
        }
    }
    _ = len(v4)
    switch "" {
    case "foo bar":
        fmt.Println(v2)
    case "hello":
        if (-9223372036854775808 >= -9223372036854775808) {
            if len(v4) > 0 {
                v4[0] = (("hello" + "hello") + "foo bar")
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 16; fi15++ {
                    v2 = v2
                    var v16 uint64 = 652
                    var v17 uint64 = v16
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi14
            }
            var v18 uint64 = 0
            _ = v18
        } else {
            for fi19 := 0; fi19 < 10; fi19++ {
                v2 -= v2
                _ = fi19
            }
            v2 = (v2 % (((^(v2))) | 1))
            fmt.Println(v2)
        }
    }
    var v20 uint = v2
    _ = len(v4)
    switch 9223372036854775807 {
    case -456:
        v21 := GenS2{F0: 1}
        _ = v21
    case 262:
        v2 = v20
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        switch (("foo bar" + "hello") + ("hello" + "hello")) {
        case "foo bar":
            fmt.Println(v20)
        case "a":
            var v23 float64 = ((((-2.25 + float64(3.14159)) * (-2.25 - -2.25)) * 3.14159) + -2.25)
            _ = v23
        }
        var v24 uint32 = 2
        switch (fi22 & fi22) {
        case -1:
            switch (fi22 / ((fi22) | 1)) {
            case 42:
                if false {
                    if false {
                        if true {
                            fmt.Println(v24, v24)
                            for fi25 := 0; fi25 < 11; fi25++ {
                                var v26 uint8 = 164
                                _ = fi25
                                _ = v26
                            }
                            switch "a" {
                            case "hello":
                                if true {
                                    var v27 uint = 2
                                    switch fi22 {
                                    case -1:
                                        var v28 bool = (-524 > -86)
                                        _ = v28
                                    case 4294967295:
                                        v24 += uint32(v27)
                                    case 2147483647:
                                        if true {
                                            v24 += v24
                                            var v29 int8 = 1
                                            v30 := GenS1{F0: 430, F1: 42, F2: "\n\t"}
                                            _ = v29
                                            _ = v30
                                        } else {
                                            v31 := GenS1{F0: 9223372036854775807, F1: 123, F2: ((("a" + "a") + "a") + "")}
                                            if ("foo bar" <= ((("a" + "") + "hello") + "foo bar")) {
                                                var v32 int32 = -2147483648
                                                v33 := GenS3{F0: true, F1: (-(fi22))}
                                                delete(v3, (fi22 >> ((v2) & 63)))
                                                _ = v32
                                                _ = v33
                                            }
                                            _ = v31
                                        }
                                    }
                                    _ = v27
                                }
                            default:
                                var v34 uint16 = 32767
                                _ = v34
                            }
                        }
                    }
                    for fi35 := 0; fi35 < 7; fi35++ {
                        var v36 uint8 = 127
                        v20 += v2
                        _ = len(v4)
                        var v37 uint32 = v24
                        var v38 uint16 = uint16(v37)
                        v39 := GenS4{F0: true}
                        fmt.Println(v36)
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
            case 32768:
                var v40 int64 = 42
                _ = v40
            }
        case -479:
            for fi41 := 0; fi41 < 14; fi41++ {
                if false {
                    switch "\n\t" {
                    case "":
                        if ((v20 % (((v20 ^ (v20 << ((v2) & 63)))) | 1)) != (v20 | v2)) {
                            v42 := GenS3{F0: (42 < 9223372036854775807), F1: (fi41 >> ((18446744073709551615) & 63))}
                            _ = v42
                        }
                    case "\n\t":
                        v2 = (v2 >> (((v2 << ((v20) & 63))) & 63))
                    }
                    if true {
                        v43 := v1[(fi41 ^ 127)]
                        var v44 uint64 = 2
                        switch "\n\t" {
                        case "foo bar":
                            if len(v4) > 0 {
                                v4[0] = "foo bar"
                            }
                        case "hello":
                            var v45 bool = false
                            _ = v45
                        case "":
                            v46 := GenS2{F0: (v44 + v44)}
                            _ = v46
                        default:
                            v47 := GenS3{F0: (-1 == 100), F1: 350}
                            _ = v47
                        }
                        _ = v43
                        _ = v44
                    } else {
                        v24 -= v24
                        v24 += v24
                        switch 9223372036854775807 {
                        case 4294967295:
                            if (42 > 42) {
                                v4 = append(v4, "hello")
                                v48 := GenS4{F0: false}
                                _ = v48
                            } else {
                                var v49 float32 = (3.14159 - (float32(((-2.25 / 3.14159) + 1e10)) - 1.5))
                                fmt.Println(fi22, fi22)
                                _ = v49
                            }
                        case 1:
                            v4 = append(v4, "foo bar")
                        case -416:
                            v50 := GenS5{F0: 65535}
                            _ = v50
                        }
                    }
                }
                _ = fi41
            }
        case -9223372036854775808:
            var v51 uint8 = 2
            _ = v51
        default:
            v52 := GenS4{F0: (1.5 == (3.14159 - ((3.14159 - float32(1e10)) - 0.1)))}
            _ = v52
        }
        _ = fi22
        _ = v24
    }
    for fi53 := 0; fi53 < 11; fi53++ {
        fmt.Println(fi53, v20, v2)
        _ = len(v1)
        v2 = v2
        _ = fi53
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
