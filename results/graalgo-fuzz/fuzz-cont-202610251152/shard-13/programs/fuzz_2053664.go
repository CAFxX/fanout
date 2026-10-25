package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 string }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int64; F1 bool; F2 int32 }

func main() {
    v1 := []string{"\n\t", (("foo bar" + "\n\t") + "a"), "hello"}
    var v2 int64 = 120
    var v3 int8 = 127
    var v4 int = 868
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 := GenS1{F0: 32767, F1: -1, F2: ""}
    switch 1 {
    case -998:
        var v6 float32 = 1e10
        _ = v6
    case -855:
        if false {
            _ = len(v1)
            v1 = append(v1, "a")
            for fi7 := 0; fi7 < 2; fi7++ {
                var v8 uint16 = 32767
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := GenS1{F0: (v8 / ((v8) | 1)), F1: 0, F2: ""}
                    v11 := GenS1{F0: (v8 | (^(v8))), F1: 977, F2: "\n\t"}
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    v4 -= 1
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                var v12 string = ("" + "a")
                _ = fi7
                _ = v12
                _ = v8
            }
        } else {
            fmt.Println(v4, v2)
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := GenS1{F0: 915, F1: 0, F2: ""}
                switch "\n\t" {
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v15 int8 = v3
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
    case -9223372036854775808:
        switch (^(v4)) {
        case -451:
            for fi16 := 0; fi16 < 1; fi16++ {
                v17 := GenS1{F0: 65535, F1: 1, F2: ""}
                switch "" {
                case "foo bar":
                    v18 := GenS1{F0: 764, F1: 2147483647, F2: (("\n\t" + "\n\t") + "\n\t")}
                    _ = v18
                }
                _ = fi16
                _ = v17
            }
        case 0:
            if (1.5 != 3.14159) {
                switch (v4 >> ((0) & 63)) {
                case -9223372036854775808:
                    v4 = -463
                case 128:
                    v2 = -1
                }
                if (true && (v2 < v2)) {
                    v3 -= v3
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    var v19 float32 = ((3.14159 * 1.5) + 0.0)
                    _ = v19
                }
            } else {
                fmt.Println(v3, v4, v2)
                v1 = append(v1, "foo bar")
                v3 = v3
            }
        case 32768:
            for fi20 := 0; fi20 < 3; fi20++ {
                fmt.Println(fi20, v4, fi20)
                _ = fi20
            }
        }
    }
    if (!((-2.25 / (((3.14159 + (float64(0.0) - 1e10)) - (-2.25 + ((((float64(-2.25) / float64((((float32((0.1 / 1e10)) + 1e10) + 3.14159) / 1.0))) / 1.0) - 0.0) * 1.5))) * -2.25)) <= 0.0)) {
        var v21 int64 = v2
        if true {
            v22 := GenS1{F0: 255, F1: -2147483648, F2: (("" + "foo bar") + "foo bar")}
            _ = v22
        }
        var v23 uint32 = 133
        _ = v21
        _ = v23
    }
    switch "hello" {
    case "hello":
        v3 = 119
    default:
        var v24 uint8 = 17
        _ = v24
    }
    var v25 int = v4
    v4 += 951
    var v26 uint64 = 0
    var v27 int8 = -114
    for fi28 := 0; fi28 < 3; fi28++ {
        if (9223372036854775807 >= (v2 & int64(v4))) {
            var v29 int32 = 0
            if false {
                switch 2147483647 {
                case -673:
                    v25 -= 0
                case 819:
                    v30 := GenS2{F0: v26, F1: -32768}
                    _ = v30
                default:
                    v4 -= (fi28 - (v4 & int(v29)))
                }
                v31 := GenS3{F0: "\n\t", F1: -170}
                v29 = (v29 | -1)
                _ = v31
            }
            v32 := GenS2{F0: (^(v26)), F1: 1000}
            _ = v29
            _ = v32
        } else {
            fmt.Println(v2, v3)
        }
        for fi33 := 0; fi33 < 8; fi33++ {
            var v34 uint16 = 209
            switch "" {
            case "foo bar":
                if (!false) {
                    v2 += v2
                    if (v26 == v26) {
                        v1 = append(v1, "a")
                        v27 = v3
                    } else {
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    }
                    if false {
                        var v35 int16 = -1
                        if (!(!(false && false))) {
                            var v36 uint = 0
                            _ = len(v1)
                            _ = v36
                        }
                        switch ("a" + ("" + "foo bar")) {
                        case "":
                            v3 -= v3
                        case "foo bar":
                            if false {
                                var v37 float32 = (0.1 * -2.25)
                                if false {
                                    if false {
                                        v38 := GenS4{F0: (v2 / ((v2) | 1)), F1: true, F2: 42}
                                        v27 -= 100
                                        switch 127 {
                                        case -9223372036854775808:
                                            v2 = 0
                                        case -2147483648:
                                            var v39 string = ("hello" + (("a" + "") + "a"))
                                            _ = v39
                                        }
                                        _ = v38
                                    }
                                    v40 := GenS2{F0: (v26 | (v26 % ((42) | 1))), F1: v35}
                                    var v41 uint64 = v26
                                    _ = v40
                                    _ = v41
                                } else {
                                    v42 := GenS2{F0: (v26 - 18446744073709551615), F1: v35}
                                    _ = v42
                                }
                                _ = v37
                            }
                        case "\n\t":
                            if (v26 >= v26) {
                                v1 = append(v1, (("\n\t" + "") + ("a" + "hello")))
                            }
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v35
                    } else {
                        v2 -= (v2 | (v2 << ((422) & 63)))
                    }
                } else {
                    v26 += v26
                }
            case "hello":
                var v43 float64 = 1.5
                _ = v43
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi33
            _ = v34
        }
        switch ("\n\t" + "a") {
        case "a":
            fmt.Println(v4)
        }
        _ = fi28
    }
    v26 += (v26 * (v26 & 156))
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
