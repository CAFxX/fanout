package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = float64(3.14159)
    var v3 int8 = -60
    fmt.Println(v3, v3)
    v4 := GenS1{F0: 255}
    v5 := GenS1{F0: 1}
    if true {
        for fi6 := 0; fi6 < 2; fi6++ {
            switch fi6 {
            case 2147483647:
                v2 = (v2 + v2)
            case -534:
                v3 -= (v3 - v3)
            case 944:
                var v7 int16 = -1
                _ = v7
            }
            _ = fi6
        }
        v2 = (1e10 - float64(3.14159))
    } else {
        if (-2147483648 <= -603) {
            if (910 > -1) {
                if true {
                    if (false || (false || true)) {
                        for fi8 := 0; fi8 < 2; fi8++ {
                            switch (-(fi8)) {
                            case -1:
                                for fi9 := 0; fi9 < 8; fi9++ {
                                    if (2 > 42) {
                                        fmt.Println(v3)
                                        v10 := GenS2{F0: (true && (32767 <= -451))}
                                        var v11 int64 = -1
                                        _ = v10
                                        _ = v11
                                    } else {
                                        v12 := GenS3{F0: "a"}
                                        v3 += 127
                                        _ = v12
                                    }
                                    v2 -= (float64((1e10 / 3.14159)) + (3.14159 * ((v2 + float64(0.0)) / v2)))
                                    v2 = (((((1.5 * v2) - (v2 / 1.0)) - v2) / v2) + (v2 - v2))
                                    v13 := GenS2{F0: (2 != 255)}
                                    _ = fi9
                                    _ = v13
                                }
                            default:
                                if false {
                                    v14 := GenS4{F0: 42}
                                    _ = len(v1)
                                    _ = v14
                                } else {
                                    switch (^(fi8)) {
                                    case 32767:
                                        var v15 int = (fi8 % ((-331) | 1))
                                        _ = v15
                                    }
                                    var v16 int64 = -9223372036854775808
                                    _ = v16
                                }
                            }
                            _ = fi8
                        }
                        var v17 bool = false
                        if true {
                            if v17 {
                                v18 := GenS3{F0: ""}
                                if false {
                                    switch 4294967295 {
                                    case -1:
                                        var v19 int8 = v3
                                        _ = v19
                                    default:
                                        fmt.Println(v17)
                                    }
                                }
                                _ = v18
                            } else {
                                v20 := GenS5{F0: 127}
                                var v21 string = ""
                                _ = v20
                                _ = v21
                            }
                            v2 -= float64(1.5)
                            for fi22 := 0; fi22 < 16; fi22++ {
                                v3 -= -113
                                _ = fi22
                            }
                        }
                        _ = v17
                    }
                }
            }
            fmt.Println(v3, v3)
        }
    }
    v23 := v1["\n\t"]
    var v24 int16 = 42
    v3 += int8(v24)
    switch "" {
    case "":
        fmt.Println(v24)
    case "a":
        if true {
            switch 2147483647 {
            case 42:
                fmt.Println(v3, v3, v23)
            default:
                v25 := GenS5{F0: (v3 ^ v3)}
                _ = v25
            }
            if ((true && false) || (189 == 239)) {
                v1["foo bar"] = (v23 >> ((649) & 63))
                v24 += 32767
                switch ((("\n\t" + ("hello" + "a")) + "foo bar") + "foo bar") {
                case "foo bar":
                    var v26 string = ("\n\t" + ("" + "a"))
                    _ = v26
                case "\n\t":
                    for fi27 := 0; fi27 < 18; fi27++ {
                        for fi28 := 0; fi28 < 4; fi28++ {
                            if true {
                                var v29 int32 = 462
                                v30 := GenS5{F0: (^(v3))}
                                _ = v29
                                _ = v30
                            }
                            if (true && true) {
                                var v31 float32 = 0.1
                                _ = len(v1)
                                _ = len(v1)
                                _ = v31
                            } else {
                                v24 -= (v24 - (v24 << ((946) & 63)))
                                switch ("a" + (("\n\t" + "") + ("hello" + "foo bar"))) {
                                case "":
                                    fmt.Println(v24, fi27, v3)
                                case "hello":
                                    delete(v1, "hello")
                                default:
                                    v23 = (v23 & v23)
                                }
                                if false {
                                    v2 = v2
                                    v32 := GenS4{F0: 18446744073709551615}
                                    switch (((("hello" + "foo bar") + "hello") + ("hello" + "")) + "\n\t") {
                                    case "\n\t":
                                        var v33 bool = false
                                        _ = v33
                                    default:
                                        v34 := GenS1{F0: 915}
                                        v35 := v1["\n\t"]
                                        var v36 float32 = (0.1 - 1.5)
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v32
                                }
                            }
                            var v37 int8 = v3
                            _ = fi28
                            _ = v37
                        }
                        _ = fi27
                    }
                case "a":
                    v3 = -1
                }
            }
        }
    default:
        fmt.Println(v3, v3)
    }
    var v38 bool = (65535 != v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
