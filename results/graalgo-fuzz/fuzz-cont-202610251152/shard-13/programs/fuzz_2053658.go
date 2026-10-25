package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int8; F2 int16 }
type GenS5 struct { F0 int }

func main() {
    v1 := []string{}
    var v2 uint = 514
    v3 := []string{"hello", ("a" + ("\n\t" + ("a" + ("hello" + "\n\t"))))}
    v4 := make(map[string]bool)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 32767 {
    case 9223372036854775807:
        fmt.Println(v2)
    case 127:
        if false {
            v2 += v2
        } else {
            for fi5 := 0; fi5 < 8; fi5++ {
                fmt.Println(fi5, fi5)
                var v6 uint8 = 163
                _ = fi5
                _ = v6
            }
        }
    default:
        var v7 uint16 = 0
        _ = v7
    }
    if false {
        if true {
            if len(v1) > 0 {
                v1[0] = ("" + "\n\t")
            }
            var v8 int = -628
            v2 = v2
            _ = v8
        } else {
            for fi9 := 0; fi9 < 5; fi9++ {
                switch "\n\t" {
                case "foo bar":
                    var v10 int64 = -803
                    _ = v10
                }
                v11 := v4["\n\t"]
                fmt.Println(v11, v2)
                _ = fi9
                _ = v11
            }
            if false {
                if ((^(v2)) == v2) {
                    fmt.Println(v2)
                } else {
                    var v12 string = ""
                    v1 = append(v1, "\n\t")
                    var v13 int32 = 1
                    _ = v12
                    _ = v13
                }
                fmt.Println(v2)
            } else {
                var v14 uint16 = 956
                if len(v3) > 0 {
                    v3[0] = (("a" + "a") + ("foo bar" + "foo bar"))
                }
                switch 9223372036854775807 {
                case -721:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -9223372036854775808:
                    v2 -= v2
                case 467:
                    v2 += (v2 * (^(v2)))
                }
                _ = v14
            }
            v2 -= 1
        }
    } else {
        v15 := GenS1{F0: 128, F1: 76, F2: v2}
        _ = v15
    }
    v1 = append(v1, "a")
    switch -2147483648 {
    case -595:
        if true {
            if (!true) {
                switch "\n\t" {
                case "":
                    v16 := GenS2{F0: 42, F1: 42}
                    _ = v16
                case "a":
                    var v17 int64 = -9223372036854775808
                    _ = v17
                default:
                    if false {
                        for fi18 := 0; fi18 < 2; fi18++ {
                            v2 -= (v2 >> ((910) & 63))
                            _ = fi18
                        }
                    }
                }
            }
            var v19 uint64 = 593
            v19 = (v19 % (((v19 ^ (v19 * v19))) | 1))
            _ = v19
        }
    case -91:
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 int64 = -9223372036854775808
            v2 = (v2 - 0)
            fmt.Println(v2, v21, v21)
            _ = fi20
            _ = v21
        }
    case -2147483648:
        v22 := GenS2{F0: 100, F1: 1}
        _ = v22
    default:
        if (0.1 >= 1.5) {
            for fi23 := 0; fi23 < 4; fi23++ {
                _ = len(v3)
                _ = fi23
            }
            var v24 int8 = 0
            _ = v24
        } else {
            if ("\n\t" == ("a" + "\n\t")) {
                fmt.Println(v2)
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 uint64 = 567
                    if true {
                        var v27 string = ((("" + "foo bar") + "foo bar") + "")
                        _ = v27
                    } else {
                        v28 := GenS3{F0: 1000}
                        v29 := GenS4{F0: 622, F1: 100, F2: -1}
                        _ = v28
                        _ = v29
                    }
                    _ = fi25
                    _ = v26
                }
            }
            v2 += 4294967295
            if (42 < 42) {
                fmt.Println(v2)
                var v30 float32 = -2.25
                _ = v30
            } else {
                fmt.Println(v2)
                v31 := v4["foo bar"]
                _ = v31
            }
        }
    }
    v2 = 42
    var v32 uint16 = 816
    v2 = 4294967295
    var v33 int32 = -720
    delete(v4, "foo bar")
    v34 := GenS2{F0: 42, F1: v32}
    delete(v4, "a")
    if true {
        v4["hello"] = (((!false) || true) && false)
    } else {
        fmt.Println(v32)
        v35 := GenS5{F0: 128}
        var v36 int = 629
        _ = v35
        _ = v36
    }
    var v37 float64 = (0.0 * 0.1)
    v38 := GenS4{F0: 32767, F1: -73, F2: 1}
    v2 += 4294967295
    fmt.Println(v32, v32, v37)
    v39 := v4["foo bar"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
