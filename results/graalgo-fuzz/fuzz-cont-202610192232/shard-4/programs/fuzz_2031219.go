package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint16; F2 int64 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int8; F1 int; F2 int16 }

func main() {
    var v1 int32 = 646
    var v2 int16 = -1
    v3 := make(map[string]string)
    v4 := []uint32{}
    fmt.Println(v2)
    v5 := GenS1{F0: 42, F1: 2, F2: 128}
    var v6 uint8 = 42
    if (0.0 <= 0.0) {
        fmt.Println(v1)
    }
    if (((^(v6)) < (v6 | 128)) && (((0.0 / 1.0) + 0.1) > (1e10 - -2.25))) {
        if true {
            switch "a" {
            case "hello":
                if false {
                    v1 = v1
                    v7 := GenS1{F0: (v6 - (v6 >> ((0) & 63))), F1: (v6 + v6), F2: 127}
                    v4 = append(v4, 226)
                    _ = v7
                }
            case "a":
                v1 = -2147483648
            case "foo bar":
                var v8 float64 = float64(1.5)
                _ = v8
            default:
                _ = len(v4)
            }
            v3[("a" + "a")] = "foo bar"
            v6 = (v6 ^ v6)
        }
        if (!true) {
            var v9 int32 = 104
            if (!false) {
                switch 9223372036854775807 {
                case 4294967295:
                    for fi10 := 0; fi10 < 20; fi10++ {
                        var v11 uint16 = 276
                        var v12 uint64 = 4294967295
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case 42:
                    switch 32767 {
                    case -9223372036854775808:
                        var v13 int8 = 0
                        _ = v13
                    case -1:
                        if ((!(-102 < -2147483648)) && true) {
                            for fi14 := 0; fi14 < 2; fi14++ {
                                var v15 string = ""
                                _ = fi14
                                _ = v15
                            }
                            for fi16 := 0; fi16 < 5; fi16++ {
                                v4 = append(v4, 2147483648)
                                v17 := GenS1{F0: 12, F1: v6, F2: (fi16 / ((-857) | 1))}
                                _ = fi16
                                _ = v17
                            }
                            v18 := GenS2{F0: (2 < 4294967295)}
                            _ = v18
                        }
                    case -420:
                        fmt.Println(v1)
                    default:
                        if false {
                            for fi19 := 0; fi19 < 16; fi19++ {
                                var v20 uint16 = 65535
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v9 -= (^(v1))
                                _ = fi19
                                _ = v20
                            }
                            switch 583 {
                            case 127:
                                var v21 string = "foo bar"
                                _ = v21
                            default:
                                switch 32767 {
                                case 127:
                                    fmt.Println(v9, v1)
                                case -1:
                                    v9 = v9
                                case -117:
                                    if (((v2 - -32768) != v2) || (!false)) {
                                        if false {
                                            v2 = (v2 * 0)
                                        }
                                        v9 -= 0
                                    } else {
                                        if true {
                                            v3[("" + "a")] = "hello"
                                        }
                                        v6 += (^(v6))
                                    }
                                }
                            }
                            delete(v3, (("foo bar" + "\n\t") + ""))
                        } else {
                            if (130 < 462) {
                                v2 -= -32768
                            } else {
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    for fi23 := 0; fi23 < 19; fi23++ {
                                        v2 = (-(v2))
                                        _ = fi23
                                    }
                                    _ = fi22
                                }
                                var v24 float32 = 1.5
                                v25 := GenS3{F0: (32767 == (v2 ^ (v2 / ((42) | 1)))), F1: 734, F2: -753}
                                _ = v24
                                _ = v25
                            }
                            v3["a"] = ""
                            for fi26 := 0; fi26 < 12; fi26++ {
                                _ = len(v4)
                                _ = fi26
                            }
                        }
                    }
                case 0:
                    v4 = append(v4, 0)
                }
            }
            _ = v9
        }
    } else {
        v1 = v1
        switch ("a" + ("hello" + "foo bar")) {
        case "a":
            fmt.Println(v2)
        case "hello":
            if true {
                var v27 uint32 = 1
                var v28 string = "foo bar"
                var v29 float32 = -2.25
                _ = v27
                _ = v28
                _ = v29
            } else {
                for fi30 := 0; fi30 < 16; fi30++ {
                    v31 := GenS4{F0: true, F1: -389}
                    v1 = v1
                    _ = fi30
                    _ = v31
                }
                for fi32 := 0; fi32 < 5; fi32++ {
                    var v33 int32 = 0
                    _ = fi32
                    _ = v33
                }
                var v34 string = ("\n\t" + "hello")
                _ = v34
            }
        default:
            fmt.Println(v6, v1)
        }
        if (!true) {
            v35 := GenS4{F0: false, F1: -75}
            v1 -= (v1 - (v1 ^ (v1 & (v1 | (v1 % (((v1 ^ v1)) | 1))))))
            var v36 string = "foo bar"
            _ = v35
            _ = v36
        }
    }
    v2 = 42
    var v37 int32 = v1
    fmt.Println(v1, v2, v37)
    switch 127 {
    case 562:
        fmt.Println(v1, v2)
    case -113:
        var v38 int = 1
        _ = v38
    case 42:
        var v39 int64 = 0
        _ = v39
    }
    fmt.Println(v2, v6, v6)
    switch "\n\t" {
    case "a":
        v40 := GenS5{F0: 8, F1: 4294967295, F2: 0}
        _ = v40
    case "":
        fmt.Println(v2, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
