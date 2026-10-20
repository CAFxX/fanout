package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := []int{119, -93, 4294967295, -2147483648}
    var v4 int8 = 101
    var v5 int32 = 2147483647
    if (!true) {
        v5 = -2147483648
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        v3 = append(v3, 1)
    }
    if true {
        if len(v3) > 0 {
            v3[0] = 262
        }
        var v7 float32 = (-2.25 / 1e10)
        _ = v7
    }
    var v8 bool = false
    if len(v3) > 0 {
        v3[0] = 846
    }
    v9 := GenS1{F0: 1, F1: ("foo bar" + "")}
    if (-32768 != -333) {
        var v10 float64 = (((-2.25 - 0.1) + -2.25) * (float64(1e10) * (0.1 - 1e10)))
        if true {
            if (v8 && (v10 < 1e10)) {
                v4 += v4
                switch ("foo bar" + "") {
                case "hello":
                    v11 := v1[65535]
                    _ = v11
                default:
                    var v12 int64 = 42
                    _ = v12
                }
            } else {
                switch -1 {
                case 1:
                    _ = len(v3)
                }
                switch (("hello" + ("hello" + "")) + "hello") {
                case "a":
                    for fi13 := 0; fi13 < 12; fi13++ {
                        for fi14 := 0; fi14 < 15; fi14++ {
                            if ((1.5 < 0.1) && v8) {
                                var v15 int = fi14
                                v3 = append(v3, fi14)
                                _ = v15
                            }
                            if v8 {
                                var v16 uint16 = 32767
                                if v8 {
                                    var v17 int = (fi14 & fi14)
                                    var v18 bool = false
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v19 uint16 = (v16 + 0)
                                    v20 := GenS1{F0: 2147483647, F1: "hello"}
                                    _ = v19
                                    _ = v20
                                }
                                _ = v16
                            } else {
                                v21 := v2[fi14]
                                if false {
                                    fmt.Println(v21)
                                    fmt.Println(v8)
                                    v3 = append(v3, fi13)
                                }
                                _ = v21
                            }
                            _ = fi14
                        }
                        v22 := GenS1{F0: 2, F1: ("hello" + ("foo bar" + ""))}
                        v23 := GenS2{F0: 0, F1: 42}
                        _ = fi13
                        _ = v22
                        _ = v23
                    }
                case "":
                    switch ("" + (("a" + ("a" + "foo bar")) + "\n\t")) {
                    case "":
                        switch "foo bar" {
                        case "a":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            for fi24 := 0; fi24 < 20; fi24++ {
                                for fi25 := 0; fi25 < 1; fi25++ {
                                    v8 = ((float64(0.0) + 3.14159) != 1e10)
                                    var v26 int16 = -669
                                    if (("" + "a") >= "") {
                                        v8 = v8
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                                var v27 uint = 359
                                _ = fi24
                                _ = v27
                            }
                        }
                    }
                }
                for fi28 := 0; fi28 < 7; fi28++ {
                    _ = len(v3)
                    _ = fi28
                }
            }
            v10 -= 1.5
        } else {
            var v29 int64 = 9223372036854775807
            fmt.Println(v10, v29)
            switch 42 {
            case 32768:
                fmt.Println(v29, v5)
            default:
                _ = len(v3)
            }
            _ = v29
        }
        if (!(v8 && (0 == 1))) {
            v30 := GenS1{F0: 2, F1: ((("foo bar" + "foo bar") + "a") + "")}
            for fi31 := 0; fi31 < 11; fi31++ {
                v8 = v8
                var v32 float64 = v10
                v8 = ("foo bar" >= ("hello" + ""))
                _ = fi31
                _ = v32
            }
            _ = v30
        }
        _ = v10
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        if (1 > 9223372036854775807) {
            fmt.Println(v5, v8, v4)
            var v34 float32 = (3.14159 / (((0.1 - 1e10) / -2.25) - 3.14159))
            var v35 float64 = 1.5
            _ = v34
            _ = v35
        }
        for fi36 := 0; fi36 < 18; fi36++ {
            v4 += v4
            var v37 int32 = v5
            _ = fi36
            _ = v37
        }
        var v38 float64 = 1e10
        _ = fi33
        _ = v38
    }
    v3 = append(v3, int(v4))
    v39 := v1[65535]
    v40 := GenS3{F0: 2}
    fmt.Println(v8)
    v41 := GenS4{F0: -128}
    switch v39 {
    case "a":
        for fi42 := 0; fi42 < 14; fi42++ {
            if len(v3) > 0 {
                v3[0] = int(v5)
            }
            var v43 uint = 65535
            switch fi42 {
            case 32768:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case 4294967295:
                v44 := GenS2{F0: -328, F1: (^(v5))}
                _ = v44
            case 9223372036854775807:
                for fi45 := 0; fi45 < 13; fi45++ {
                    v46 := GenS1{F0: 741, F1: v39}
                    v5 = v5
                    _ = fi45
                    _ = v46
                }
            }
            _ = fi42
            _ = v43
        }
    case "foo bar":
        var v47 float64 = 3.14159
        _ = v47
    default:
        fmt.Println(v39)
    }
    v4 = (v4 % ((114) | 1))
    var v48 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v48
    _ = v5
    _ = v8
    _ = v9
}
