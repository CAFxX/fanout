package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 bool; F2 bool }
type GenS3 struct { F0 int64; F1 uint64; F2 int }

func main() {
    var v1 bool = (!false)
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    var v4 bool = false
    v5 := []int{128, 128, 32768, 690}
    if (v1 && (171 != 519)) {
        fmt.Println(v1, v1)
        if (v1 && ((65535 < 252) || (135 < 0))) {
            v5 = append(v5, 4294967295)
            var v6 uint64 = 42
            if v1 {
                if false {
                    var v7 int16 = 32767
                    if v1 {
                        v8 := GenS1{F0: "foo bar"}
                        var v9 int = 127
                        fmt.Println(v7, v6, v6)
                        _ = v8
                        _ = v9
                    }
                    _ = v7
                } else {
                    if (!(false || true)) {
                        var v10 float32 = (-2.25 * (((-2.25 - (((0.1 * (0.0 - (0.0 * float32(-2.25)))) * 3.14159) / 1.0)) / float32(-2.25)) + 3.14159))
                        var v11 int = -325
                        _ = v10
                        _ = v11
                    } else {
                        v12 := GenS1{F0: ("a" + "\n\t")}
                        _ = v12
                    }
                    var v13 uint = uint(v6)
                    _ = len(v3)
                    _ = v13
                }
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 int32 = 156
                    fmt.Println(v1)
                    _ = fi14
                    _ = v15
                }
                if (v4 && ((("hello" + ("foo bar" + "a")) + ("\n\t" + (("hello" + ("\n\t" + "hello")) + ("foo bar" + ("foo bar" + ("foo bar" + ("\n\t" + ""))))))) == (("" + "\n\t") + "foo bar"))) {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        var v17 uint32 = 42
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 20; fi18++ {
                        v1 = v1
                        fmt.Println(fi18)
                        _ = fi18
                    }
                }
            } else {
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch (-(fi19)) {
                    case 42:
                        for fi20 := 0; fi20 < 7; fi20++ {
                            fmt.Println(v1)
                            _ = fi20
                        }
                    default:
                        if ((2147483647 > int32(fi19)) && (759 >= 536)) {
                            var v21 uint = 349
                            _ = v21
                        } else {
                            fmt.Println(fi19, v6)
                            v22 := v3["hello"]
                            v1 = (!((fi19 | 32768) == fi19))
                            _ = v22
                        }
                    }
                    v6 += (v6 >> ((42) & 63))
                    _ = fi19
                }
            }
            _ = v6
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v23 := GenS2{F0: 128, F1: v4, F2: v4}
            _ = v23
        }
    }
    v24 := GenS1{F0: ""}
    for fi25 := 0; fi25 < 9; fi25++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi25
    }
    var v26 uint = 246
    v1 = (v4 && ((!false) || v4))
    for fi27 := 0; fi27 < 9; fi27++ {
        _ = len(v5)
        _ = fi27
    }
    v4 = (-9223372036854775808 == 365)
    delete(v2, 128)
    if (1e10 < ((3.14159 / 1.5) / (3.14159 * -2.25))) {
        v26 += v26
        fmt.Println(v4)
        v4 = (0.1 > (((1e10 - 0.0) * 0.0) - 0.1))
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        v3[(("a" + "\n\t") + "hello")] = (!((("" + ("\n\t" + ("hello" + "\n\t"))) + "") != ("foo bar" + "hello")))
        var v29 bool = false
        _ = fi28
        _ = v29
    }
    var v30 float32 = 0.1
    switch -462 {
    case 2147483647:
        v31 := GenS1{F0: "foo bar"}
        _ = v31
    case 9223372036854775807:
        var v32 float64 = 1.5
        _ = v32
    case -73:
        if v1 {
            switch ("" + ("\n\t" + "")) {
            case "":
                var v33 int8 = -36
                _ = v33
            case "a":
                _ = len(v2)
            default:
                fmt.Println(v30, v4)
            }
        } else {
            switch 127 {
            case 2147483647:
                for fi34 := 0; fi34 < 9; fi34++ {
                    for fi35 := 0; fi35 < 20; fi35++ {
                        var v36 uint16 = 355
                        _ = fi35
                        _ = v36
                    }
                    v37 := GenS3{F0: 647, F1: 399, F2: (fi34 | fi34)}
                    v30 = v30
                    _ = fi34
                    _ = v37
                }
            case -109:
                v38 := GenS1{F0: "a"}
                _ = v38
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
    default:
        var v39 int = -932
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
