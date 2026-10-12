package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint8; F1 string; F2 int }

func main() {
    v1 := []int{2147483647, -2147483648, -651, -9223372036854775808}
    var v2 float32 = float32(0.0)
    var v3 float64 = -2.25
    v4 := make(map[string]int)
    v5 := GenS1{F0: false, F1: 18446744073709551615, F2: 2147483647}
    for fi6 := 0; fi6 < 7; fi6++ {
        switch "foo bar" {
        case "\n\t":
            v4["\n\t"] = (fi6 % ((fi6) | 1))
        case "foo bar":
            v3 = v3
        case "hello":
            switch "a" {
            case "a":
                fmt.Println(v3, v2)
            case "foo bar":
                v7 := GenS1{F0: true, F1: 0, F2: 1}
                _ = v7
            }
        default:
            var v8 int8 = 16
            _ = v8
        }
        v1 = append(v1, 0)
        if false {
            switch 267 {
            case 32768:
                v4[""] = fi6
            case 4294967295:
                v3 = -2.25
            case 9223372036854775807:
                if (!(v3 <= v3)) {
                    if len(v1) > 0 {
                        v1[0] = 32768
                    }
                    delete(v4, "hello")
                }
            default:
                if true {
                    v9 := v4[("foo bar" + ("\n\t" + "hello"))]
                    _ = v9
                }
            }
            for fi10 := 0; fi10 < 17; fi10++ {
                if ((-673 != 2147483647) || false) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    switch ((((("\n\t" + "") + ("a" + "foo bar")) + ("" + (("a" + "") + "foo bar"))) + "a") + "hello") {
                    case "hello":
                        if (true || (42 != -553)) {
                            v3 = 3.14159
                        } else {
                            var v11 int = (-(fi10))
                            _ = v11
                        }
                    case "a":
                        var v12 float32 = (v2 / -2.25)
                        v4[((((("foo bar" + ("a" + "hello")) + (("foo bar" + "") + ("a" + ("" + ("\n\t" + ""))))) + "\n\t") + (("foo bar" + "\n\t") + "foo bar")) + ("hello" + ("a" + ("" + ""))))] = -819
                        _ = v12
                    }
                    if false {
                        switch fi6 {
                        case -9223372036854775808:
                            var v13 float64 = v3
                            _ = v13
                        }
                    }
                    var v14 float64 = v3
                    fmt.Println(v14, v2, v2)
                    _ = v14
                }
                v15 := GenS2{F0: 42, F1: -749}
                _ = fi10
                _ = v15
            }
        }
        _ = fi6
    }
    v3 = v3
    var v16 uint64 = 4294967295
    if (!false) {
        var v17 string = "\n\t"
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := v4[(v17 + ((("hello" + (((("foo bar" + "foo bar") + "hello") + ("foo bar" + ("a" + v17))) + (((v17 + ((v17 + v17) + v17)) + v17) + (v17 + v17)))) + v17) + (v17 + v17)))]
            switch (fi18 + 4294967295) {
            case 4294967295:
                if (((1 != 746) || true) && (1e10 == 3.14159)) {
                    var v20 uint = 1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    fmt.Println(v3, v2, v19)
                    _ = v20
                } else {
                    var v21 int = v19
                    _ = len(v1)
                    for fi22 := 0; fi22 < 17; fi22++ {
                        var v23 uint32 = 2
                        if ((82 == 100) && ((255 == 905) || (false && (0 == 0)))) {
                            var v24 uint = 303
                            v24 = (v24 >> (((v24 >> ((v24) & 63))) & 63))
                            _ = v24
                        }
                        fmt.Println(v23, fi22)
                        v17 = v17
                        v23 += (v23 % ((v23) | 1))
                        _ = fi22
                        _ = v23
                    }
                    _ = v21
                }
            default:
                v25 := GenS1{F0: true, F1: 42, F2: int32(fi18)}
                _ = v25
            }
            v17 = ((v17 + ((v17 + (("" + "a") + "a")) + v17)) + "\n\t")
            _ = fi18
            _ = v19
        }
        var v26 int16 = 32767
        _ = v17
        _ = v26
    }
    v2 = (0.0 / float32(v3))
    v2 = v2
    if true {
        v27 := GenS1{F0: true, F1: 65535, F2: 0}
        if true {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 int64 = 42
                var v30 bool = (!false)
                if false {
                    fmt.Println(v3, fi28)
                    if true {
                        v4["foo bar"] = (fi28 | fi28)
                        for fi31 := 0; fi31 < 4; fi31++ {
                            var v32 uint32 = 1
                            var v33 string = ""
                            delete(v4, v33)
                            _ = len(v4)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                    switch fi28 {
                    case 32768:
                        var v34 int64 = -9223372036854775808
                        _ = v34
                    default:
                        v16 = 18446744073709551615
                    }
                }
                _ = fi28
                _ = v29
                _ = v30
            }
            switch 32767 {
            case -563:
                _ = len(v4)
            default:
                v16 -= v16
            }
        }
        _ = v27
    } else {
        for fi35 := 0; fi35 < 2; fi35++ {
            v2 = (float32((0.0 - v3)) + v2)
            var v36 bool = false
            var v37 string = (("a" + "") + "foo bar")
            _ = fi35
            _ = v36
            _ = v37
        }
        fmt.Println(v2)
        if true {
            if true {
                var v38 float32 = (((v2 * (v2 / 3.14159)) / (v2 + (0.0 - v2))) + v2)
                v39 := GenS3{F0: ((("a" + "") + "") + ("hello" + "hello")), F1: 22}
                _ = v38
                _ = v39
            }
            v40 := GenS4{F0: 1, F1: (("" + "\n\t") + "a"), F2: 812}
            _ = v40
        } else {
            var v41 float64 = v3
            v42 := GenS3{F0: ("" + (("" + (("a" + ("a" + "hello")) + ("a" + "hello"))) + ((("a" + ("hello" + (("a" + "hello") + "hello"))) + "a") + "foo bar"))), F1: 128}
            v2 = v2
            _ = v41
            _ = v42
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    var v43 int64 = -125
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (4294967295 == 255) {
        if (-352 <= -1) {
            var v44 int16 = -1
            _ = len(v4)
            _ = v44
        } else {
            v43 += v43
            v2 -= (v2 / 1.5)
            var v45 int8 = 100
            _ = v45
        }
        v46 := GenS4{F0: 128, F1: "\n\t", F2: 0}
        for fi47 := 0; fi47 < 5; fi47++ {
            v48 := GenS1{F0: (!false), F1: 42, F2: -492}
            _ = fi47
            _ = v48
        }
        _ = v46
    } else {
        v43 = v43
    }
    if true {
        if (!((3.14159 + v2) <= v2)) {
            if true {
                if (!false) {
                    switch "foo bar" {
                    case "\n\t":
                        v49 := v4[("a" + "a")]
                        _ = v49
                    default:
                        v1 = append(v1, -2147483648)
                    }
                    if false {
                        if (false || false) {
                            fmt.Println(v3, v16)
                        } else {
                            v2 -= (v2 * float32(3.14159))
                            for fi50 := 0; fi50 < 6; fi50++ {
                                var v51 bool = (65535 == 259)
                                _ = fi50
                                _ = v51
                            }
                        }
                        var v52 uint32 = 0
                        if len(v1) > 0 {
                            v1[0] = 32768
                        }
                        _ = v52
                    }
                } else {
                    var v53 uint32 = 0
                    switch 9223372036854775807 {
                    case 24:
                        var v54 uint8 = 54
                        _ = v54
                    }
                    _ = v53
                }
                for fi55 := 0; fi55 < 7; fi55++ {
                    v16 += (^(v16))
                    _ = fi55
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
