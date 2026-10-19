package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint; F1 int16; F2 uint32 }
type GenS4 struct { F0 int8; F1 int; F2 int }
type GenS5 struct { F0 int16; F1 int; F2 uint }
type GenS6 struct { F0 int64 }
func genFn0(p0 uint16, p1 int16, p2 bool) int16 {
    if true {
        fmt.Println(p0, p0)
        switch -494 {
        case 0:
            p2 = p2
        case 4294967295:
            p0 = 2
        case -2147483648:
            var v4 uint8 = 0
            _ = v4
        }
    } else {
        var v5 int32 = 2147483647
        _ = v5
    }
    fmt.Println(p1)
    v6 := GenS1{F0: -52}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    v1 := []string{"a"}
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v7 bool = (63 >= 128)
    if (!v7) {
        fmt.Println(v7)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v8 bool = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        delete(v2, "foo bar")
        v10 := GenS1{F0: 100}
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 65535
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 4; fi13++ {
            var v14 uint64 = 1
            var v15 int = 2147483647
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            fmt.Println(v11, fi16)
            if v7 {
                var v17 int = fi16
                v11 = v11
                if v7 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if v7 {
                        v18 := GenS1{F0: 42}
                        switch ("hello" + "a") {
                        case "":
                            if v7 {
                                var v19 int64 = -9223372036854775808
                                _ = v19
                            }
                        case "foo bar":
                            var v20 float64 = (-2.25 - 0.1)
                            _ = v20
                        case "\n\t":
                            switch (fi16 + (^(v17))) {
                            case -2147483648:
                                switch (("hello" + "a") + "a") {
                                case "a":
                                    v21 := GenS1{F0: 1}
                                    _ = v21
                                case "foo bar":
                                    fmt.Println(v11, fi16)
                                case "hello":
                                    v22 := GenS1{F0: 0}
                                    _ = v22
                                }
                            default:
                                var v23 uint8 = 1
                                _ = v23
                            }
                        default:
                            v24 := GenS1{F0: -128}
                            _ = v24
                        }
                        _ = v18
                    }
                    switch fi16 {
                    case 1:
                        v25 := v2["foo bar"]
                        _ = v25
                    }
                } else {
                    var v26 int16 = 1000
                    v17 = 4294967295
                    _ = v26
                }
                _ = v17
            } else {
                _ = len(v3)
                v11 = (v11 >> ((1) & 63))
                fmt.Println(fi12, fi16, fi16)
            }
            switch "\n\t" {
            case "hello":
                v7 = (("a" + ("foo bar" + ("hello" + (("a" + "foo bar") + "\n\t")))) < "")
            case "":
                v7 = true
            case "\n\t":
                v7 = (v7 && v7)
            }
            _ = fi16
        }
        _ = fi12
    }
    delete(v2, "a")
    v7 = (42 > 2147483647)
    var v27 int32 = -131
    v7 = true
    if (-9223372036854775808 >= 1) {
        v28 := GenS2{F0: 255, F1: 0}
        v2[(("\n\t" + "") + ("" + "hello"))] = 32767
        if false {
            for fi29 := 0; fi29 < 7; fi29++ {
                switch "hello" {
                case "a":
                    v27 -= v27
                case "\n\t":
                    for fi30 := 0; fi30 < 7; fi30++ {
                        var v31 string = ("a" + ("\n\t" + ((("a" + "a") + "foo bar") + (("foo bar" + (("" + "a") + (("hello" + "\n\t") + ("foo bar" + "a")))) + "foo bar"))))
                        switch -2147483648 {
                        case 2147483647:
                            delete(v3, 123)
                        case 127:
                            v3[846] = 0
                        case -2147483648:
                            _ = len(v1)
                        }
                        _ = fi30
                        _ = v31
                    }
                default:
                    for fi32 := 0; fi32 < 1; fi32++ {
                        switch "a" {
                        case "\n\t":
                            var v33 int16 = -32768
                            _ = v33
                        }
                        _ = len(v2)
                        if ((fi29 + (^(fi32))) <= fi29) {
                            var v34 float32 = (3.14159 + -2.25)
                            _ = v34
                        }
                        _ = fi32
                    }
                }
                var v35 string = ""
                _ = fi29
                _ = v35
            }
            for fi36 := 0; fi36 < 7; fi36++ {
                fmt.Println(v7)
                fmt.Println(v27, v27, v7)
                _ = fi36
            }
            _ = len(v1)
        }
        _ = v28
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        v27 = -1
        v38 := GenS1{F0: -91}
        if (!v7) {
            v11 += (v11 & (^(v11)))
            v7 = v7
        } else {
            v39 := GenS1{F0: 0}
            fmt.Println(v11)
            if false {
                for fi40 := 0; fi40 < 11; fi40++ {
                    switch "\n\t" {
                    case "hello":
                        var v41 int64 = 9223372036854775807
                        _ = v41
                    case "a":
                        var v42 uint8 = 1
                        _ = v42
                    case "":
                        v3[2147483648] = (-(fi40))
                    }
                    v43 := GenS2{F0: 1, F1: 255}
                    _ = fi40
                    _ = v43
                }
                v44 := GenS1{F0: 100}
                _ = v44
            } else {
                for fi45 := 0; fi45 < 3; fi45++ {
                    v46 := GenS2{F0: 0, F1: 239}
                    if (2147483648 != 4294967295) {
                        var v47 uint64 = 2
                        var v48 float32 = 3.14159
                        if v7 {
                            v49 := GenS3{F0: 4294967295, F1: 0, F2: 2147483647}
                            _ = v49
                        }
                        _ = v47
                        _ = v48
                    }
                    var v50 float32 = float32(((-2.25 + (((3.14159 / 1.5) + 1.5) - ((1e10 / 1.0) - 1e10))) - ((0.0 / (3.14159 / (0.0 + (((0.1 / 1.0) - 3.14159) + 3.14159)))) - 0.1)))
                    _ = fi45
                    _ = v46
                    _ = v50
                }
                for fi51 := 0; fi51 < 2; fi51++ {
                    fmt.Println(v11, fi37)
                    _ = fi51
                }
                if len(v1) > 0 {
                    v1[0] = ""
                }
            }
            _ = v39
        }
        _ = fi37
        _ = v38
    }
    fmt.Println(v11, v27, v7)
    v3[0] = 9223372036854775807
    v27 = 1
    switch "hello" {
    case "":
        var v52 bool = (128 > 42)
        _ = v52
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch "" {
    case "foo bar":
        v53 := GenS2{F0: 18446744073709551615, F1: 127}
        _ = v53
    default:
        v54 := v2["hello"]
        _ = v54
    }
    for fi55 := 0; fi55 < 17; fi55++ {
        fmt.Println(v7)
        v7 = v7
        _ = fi55
    }
    for fi56 := 0; fi56 < 9; fi56++ {
        switch 2147483647 {
        case 42:
            var v57 float64 = 1e10
            _ = v57
        case 0:
            var v58 uint = 2
            _ = v58
        default:
            for fi59 := 0; fi59 < 13; fi59++ {
                if ((float64(0.1) < (((0.1 * 3.14159) - 1.5) - 1e10)) && v7) {
                    if (176 != 2) {
                        v7 = v7
                        switch fi56 {
                        case 2147483647:
                            v60 := v3[109]
                            _ = v60
                        }
                    } else {
                        v61 := GenS4{F0: 1, F1: (fi59 & fi56), F2: int(v11)}
                        switch ("hello" + "\n\t") {
                        case "hello":
                            v62 := GenS4{F0: -1, F1: fi56, F2: fi59}
                            v11 = v11
                            var v63 int64 = 0
                            _ = v62
                            _ = v63
                        case "\n\t":
                            fmt.Println(fi56, fi56, v27)
                        }
                        _ = v61
                    }
                    var v64 bool = (65535 > 4294967295)
                    v65 := GenS5{F0: int16(v27), F1: -2147483648, F2: 467}
                    _ = v64
                    _ = v65
                } else {
                    v27 -= (v27 >> ((2) & 63))
                }
                fmt.Println(fi56, v7, v27)
                _ = fi59
            }
        }
        for fi66 := 0; fi66 < 2; fi66++ {
            v7 = false
            fmt.Println(fi56)
            v67 := GenS6{F0: 107}
            switch "" {
            case "a":
                var v68 int64 = 557
                _ = v68
            }
            _ = fi66
            _ = v67
        }
        v69 := GenS6{F0: -1}
        _ = fi56
        _ = v69
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v7
}
