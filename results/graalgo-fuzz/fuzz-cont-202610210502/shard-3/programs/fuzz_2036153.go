package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint16 }

func main() {
    var v1 bool = ((("hello" + "foo bar") + ("a" + ("\n\t" + (("" + "hello") + "")))) >= "\n\t")
    v2 := []string{}
    v3 := []int{-663, -304, 42}
    v4 := make(map[string]int)
    switch ("a" + ("a" + ("\n\t" + "\n\t"))) {
    case "hello":
        _ = len(v2)
    case "foo bar":
        switch -1 {
        case 32767:
            var v5 string = ("foo bar" + "a")
            _ = v5
        case 141:
            v1 = false
        case -1:
            for fi6 := 0; fi6 < 7; fi6++ {
                if ("hello" != "") {
                    if (42 != 438) {
                        if v1 {
                            fmt.Println(fi6)
                        } else {
                            if (-2.25 != (((((0.0 / 1e10) + (float64(0.1) * 1e10)) / (float64(float32((3.14159 + 1e10))) / 1.0)) - float64((float32((0.1 * 3.14159)) * (1e10 - (0.1 - float32(1e10)))))) - ((((0.1 * float64((1e10 - 1.5))) * 0.1) + float64(0.0)) + ((1.5 * float64((float32((0.0 * 1e10)) * 0.1))) - 1e10)))) {
                                for fi7 := 0; fi7 < 18; fi7++ {
                                    fmt.Println(v1)
                                    v1 = v1
                                    _ = fi7
                                }
                                var v8 uint = 42
                                _ = v8
                            }
                            delete(v4, "\n\t")
                            var v9 uint = 2
                            _ = v9
                        }
                        switch (fi6 + (fi6 - -2147483648)) {
                        case -284:
                            var v10 int16 = 42
                            _ = v10
                        case 9223372036854775807:
                            var v11 int32 = -136
                            _ = v11
                        case 128:
                            if false {
                                if v1 {
                                    v12 := GenS1{F0: 195, F1: (!((!(255 <= 127)) && v1))}
                                    var v13 bool = (("foo bar" + "hello") != "a")
                                    for fi14 := 0; fi14 < 9; fi14++ {
                                        var v15 int32 = 1
                                        var v16 int32 = 1
                                        v15 -= (v16 / ((v16) | 1))
                                        if (9223372036854775807 >= -178) {
                                            if (42 <= -962) {
                                                v1 = v13
                                                fmt.Println(v13, v1)
                                            }
                                        }
                                        if ((0.1 <= 3.14159) && true) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        v13 = ((fi14 & 9223372036854775807) > 32768)
                                        v17 := GenS1{F0: 148, F1: false}
                                        var v18 int = 713
                                        v19 := GenS2{F0: "foo bar", F1: 18446744073709551615}
                                        fmt.Println(v16, fi14)
                                        var v20 int = 127
                                        _ = fi14
                                        _ = v15
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    }
                                    _ = v12
                                    _ = v13
                                } else {
                                    v1 = v1
                                }
                                if (("a" + ("\n\t" + ("" + ("\n\t" + ("hello" + ""))))) != "foo bar") {
                                    v21 := v4[("a" + ("hello" + ""))]
                                    var v22 uint = 65535
                                    v23 := GenS1{F0: 0, F1: ((0.0 - (1.5 - 1e10)) == 1.5)}
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                } else {
                                    v2 = append(v2, ("hello" + ("hello" + ("a" + ""))))
                                }
                                fmt.Println(v1, v1)
                            }
                        default:
                            var v24 uint8 = 60
                            _ = v24
                        }
                        if (1 < -1) {
                            v4[""] = fi6
                        }
                    }
                    var v25 int = 192
                    fmt.Println(v1, v1, fi6)
                    _ = v25
                }
                _ = fi6
            }
        default:
            v1 = v1
        }
    }
    v1 = (!(42 <= 65535))
    var v26 uint64 = 42
    if false {
        _ = len(v4)
    } else {
        var v27 int16 = -32768
        v28 := GenS3{F0: 2147483647, F1: v27, F2: 32767}
        if (v1 && v1) {
            var v29 uint = 110
            fmt.Println(v29, v26, v26)
            _ = v29
        }
        _ = v27
        _ = v28
    }
    v1 = true
    v1 = (9223372036854775807 != -327)
    for fi30 := 0; fi30 < 13; fi30++ {
        fmt.Println(v26, fi30, v1)
        _ = fi30
    }
    fmt.Println(v1, v26)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
