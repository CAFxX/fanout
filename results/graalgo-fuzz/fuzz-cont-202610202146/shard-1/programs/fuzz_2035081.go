package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int64; F2 int64 }
type GenS5 struct { F0 int32; F1 int; F2 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 string; F2 uint16 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]uint32)
    var v3 float64 = float64(0.1)
    v4 := []string{"a", ("a" + ("a" + "")), "foo bar"}
    var v5 uint64 = 1
    v1 += (v1 / (((v1 / ((v1) | 1))) | 1))
    fmt.Println(v1, v5)
    if true {
        v6 := GenS1{F0: (v1 >> ((65535) & 63))}
        _ = v6
    } else {
        if true {
            var v7 string = "hello"
            switch "hello" {
            case "hello":
                var v8 uint16 = 2
                _ = v8
            }
            _ = v7
        } else {
            switch "foo bar" {
            case "hello":
                switch "hello" {
                case "":
                    var v9 string = "hello"
                    _ = v9
                default:
                    var v10 uint64 = 956
                    _ = v10
                }
            case "\n\t":
                var v11 int64 = -9223372036854775808
                _ = v11
            }
            v3 -= (v3 / (0.0 - v3))
            v12 := GenS2{F0: 546}
            _ = v12
        }
        v3 = v3
        if (false || true) {
            v1 = 1
            switch ("\n\t" + (("a" + "") + "\n\t")) {
            case "hello":
                v13 := GenS3{F0: 42, F1: 29}
                _ = v13
            case "":
                if false {
                    _ = len(v2)
                } else {
                    v14 := GenS4{F0: 65535, F1: 1, F2: int64(v5)}
                    switch -954 {
                    case 4294967295:
                        v15 := GenS5{F0: 2147483647, F1: 961, F2: 1}
                        _ = v15
                    case 42:
                        for fi16 := 0; fi16 < 1; fi16++ {
                            v5 = (v5 - (v5 / ((2) | 1)))
                            if ("foo bar" >= "a") {
                                var v17 string = ("" + "")
                                var v18 int = fi16
                                _ = v17
                                _ = v18
                            }
                            _ = fi16
                        }
                    default:
                        v1 = (v1 / (((v1 * 193)) | 1))
                    }
                    _ = v14
                }
            default:
                v4 = append(v4, ("foo bar" + ((("" + "foo bar") + ("a" + "\n\t")) + "foo bar")))
            }
            switch "a" {
            case "\n\t":
                fmt.Println(v1)
            }
        }
    }
    if (true && ((v1 | v1) < (v1 + (v1 ^ 128)))) {
        v19 := GenS2{F0: 0}
        v20 := GenS4{F0: 255, F1: -931, F2: -150}
        var v21 int8 = 1
        _ = v19
        _ = v20
        _ = v21
    }
    if (1 == v1) {
        v3 = (float64(-2.25) + v3)
        var v22 int64 = -1
        _ = v22
    }
    if (v1 == (v1 - 128)) {
        if (false || true) {
            v23 := GenS5{F0: 733, F1: 9223372036854775807, F2: 42}
            v3 -= v3
            _ = v23
        } else {
            v24 := GenS2{F0: 926}
            switch 4294967295 {
            case 282:
                for fi25 := 0; fi25 < 1; fi25++ {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v27 := GenS6{F0: -1}
                        _ = fi26
                        _ = v27
                    }
                    _ = fi25
                }
            }
            switch 32768 {
            case 32768:
                if ("foo bar" != "") {
                    v5 -= v5
                    if true {
                        if false {
                            v5 -= v5
                            v5 = 270
                            v28 := GenS7{F0: 9223372036854775807, F1: (("\n\t" + "foo bar") + ("hello" + ("hello" + ("foo bar" + "hello")))), F2: 964}
                            _ = v28
                        }
                        v29 := GenS4{F0: 0, F1: 1, F2: 42}
                        _ = v29
                    }
                    v5 = v5
                }
            case -9223372036854775808:
                for fi30 := 0; fi30 < 19; fi30++ {
                    v3 = v3
                    var v31 uint32 = 2
                    _ = fi30
                    _ = v31
                }
            case -549:
                var v32 bool = (42 > 56)
                _ = v32
            default:
                v3 = v3
            }
            _ = v24
        }
        fmt.Println(v1, v5, v3)
        for fi33 := 0; fi33 < 17; fi33++ {
            for fi34 := 0; fi34 < 17; fi34++ {
                v3 = v3
                var v35 uint64 = 371
                v1 = (^(v1))
                _ = fi34
                _ = v35
            }
            switch fi33 {
            case -2147483648:
                var v36 int64 = -9223372036854775808
                _ = v36
            default:
                var v37 uint8 = 145
                _ = v37
            }
            var v38 uint16 = 1
            _ = fi33
            _ = v38
        }
    } else {
        v2[2147483647] = 2147483648
        v39 := GenS3{F0: -1, F1: 65535}
        _ = v39
    }
    switch 4294967295 {
    case -717:
        v40 := GenS1{F0: (v1 * 106)}
        _ = v40
    case 1:
        var v41 uint32 = 0
        _ = v41
    case 9223372036854775807:
        if (false || true) {
            var v42 int32 = 2147483647
            if true {
                for fi43 := 0; fi43 < 2; fi43++ {
                    v44 := v2[2]
                    _ = fi43
                    _ = v44
                }
            }
            _ = v42
        } else {
            if false {
                var v45 float32 = 0.1
                _ = v45
            } else {
                v1 -= (v1 >> ((65535) & 63))
                v46 := GenS5{F0: int32(v5), F1: -521, F2: uint32(v5)}
                _ = v46
            }
            var v47 int16 = 32767
            _ = v47
        }
    default:
        switch "" {
        case "hello":
            if ((v5 | v5) == v5) {
                if true {
                    switch -135 {
                    case 2147483647:
                        for fi48 := 0; fi48 < 16; fi48++ {
                            v1 += v1
                            _ = fi48
                        }
                    }
                    var v49 int8 = -1
                    var v50 uint16 = 13
                    _ = v49
                    _ = v50
                }
            } else {
                v5 = v5
            }
        }
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    v1 = v1
    var v51 float32 = 3.14159
    if (0 <= 1) {
        var v52 float64 = (float64(float32(v3)) * v3)
        v51 = v51
        if false {
            var v53 float32 = 0.0
            _ = v53
        }
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
