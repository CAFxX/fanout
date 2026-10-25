package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int8; F1 int16; F2 int32 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{"", ("" + ("\n\t" + (("a" + ("hello" + "foo bar")) + ""))), "foo bar", "foo bar"}
    v3 := make(map[int]bool)
    var v4 bool = false
    var v5 int8 = 1
    var v6 uint64 = 1
    v4 = (("" + "") != ("\n\t" + "foo bar"))
    if v4 {
        if (42 != 1) {
            var v7 int64 = -9223372036854775808
            switch ("\n\t" + "a") {
            case "foo bar":
                v7 = 9223372036854775807
            }
            if v4 {
                var v8 uint64 = v6
                var v9 uint = 42
                var v10 float64 = 0.0
                _ = v10
                _ = v8
                _ = v9
            } else {
                for fi11 := 0; fi11 < 3; fi11++ {
                    if false {
                        if ("" < "hello") {
                            fmt.Println(v5, v6)
                        } else {
                            for fi12 := 0; fi12 < 10; fi12++ {
                                v13 := GenS1{F0: 1, F1: 166}
                                v3[(fi11 | fi11)] = v4
                                switch "hello" {
                                case "":
                                    if (v4 || v4) {
                                        var v14 string = "foo bar"
                                        _ = v14
                                    }
                                default:
                                    v7 = v7
                                }
                                _ = fi12
                                _ = v13
                            }
                        }
                        switch "foo bar" {
                        case "foo bar":
                            if true {
                                delete(v3, (fi11 ^ 422))
                                for fi15 := 0; fi15 < 2; fi15++ {
                                    if v4 {
                                        var v16 uint = 255
                                        if v4 {
                                            switch fi15 {
                                            case -9223372036854775808:
                                                v17 := GenS1{F0: -820, F1: v16}
                                                _ = v17
                                            }
                                            v6 += 174
                                            fmt.Println(v5, v6)
                                        } else {
                                            switch (fi15 & fi11) {
                                            case -9223372036854775808:
                                                v18 := GenS1{F0: 2147483647, F1: 65535}
                                                _ = v18
                                            case 127:
                                                var v19 uint64 = 744
                                                _ = v19
                                            case 522:
                                                var v20 string = "a"
                                                switch fi11 {
                                                case 1:
                                                    var v21 uint64 = v6
                                                    _ = v21
                                                case 143:
                                                    v7 -= 646
                                                default:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                var v22 uint64 = v6
                                                _ = v20
                                                _ = v22
                                            default:
                                                fmt.Println(v5, v4)
                                            }
                                        }
                                        fmt.Println(v6, fi11, fi15)
                                        _ = v16
                                    }
                                    v23 := GenS2{F0: (v5 ^ v5), F1: 0, F2: 491}
                                    _ = fi15
                                    _ = v23
                                }
                                if ((!v4) || v4) {
                                    if v4 {
                                        var v24 int32 = 0
                                        v24 = (v24 >> ((0) & 63))
                                        var v25 string = "hello"
                                        _ = v24
                                        _ = v25
                                    } else {
                                        for fi26 := 0; fi26 < 17; fi26++ {
                                            v4 = (-154 >= 1)
                                            fmt.Println(fi11)
                                            v4 = v4
                                            delete(v3, fi26)
                                            _ = fi26
                                        }
                                    }
                                }
                            }
                        }
                        var v27 int = fi11
                        _ = v27
                    } else {
                        v28 := GenS1{F0: 42, F1: 65535}
                        _ = v28
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi11
                }
            }
            _ = v7
        }
        if ((3.14159 - 1.5) >= (((0.1 / 1.0) / 3.14159) + float32(0.1))) {
            var v29 int = -2147483648
            _ = v29
        } else {
            v5 -= v5
            v30 := v3[-1]
            _ = v30
        }
        var v31 string = "foo bar"
        _ = v31
    } else {
        switch "" {
        case "hello":
            v5 -= 127
        default:
            for fi32 := 0; fi32 < 14; fi32++ {
                v6 -= 1
                for fi33 := 0; fi33 < 2; fi33++ {
                    _ = len(v1)
                    v5 = (v5 << ((1) & 63))
                    if false {
                        v34 := GenS3{F0: 277, F1: (!false), F2: 0}
                        switch (fi32 ^ 2147483647) {
                        case 1:
                            if v4 {
                                var v35 int8 = 1
                                _ = v35
                            } else {
                                fmt.Println(v6)
                            }
                        }
                        _ = v34
                    }
                    _ = fi33
                }
                _ = fi32
            }
        }
    }
    var v36 float32 = 1.5
    _ = len(v1)
    delete(v3, 17)
    for fi37 := 0; fi37 < 9; fi37++ {
        for fi38 := 0; fi38 < 8; fi38++ {
            switch "\n\t" {
            case "\n\t":
                var v39 uint64 = (v6 * v6)
                _ = v39
            default:
                var v40 float64 = ((-2.25 * 1.5) / 0.1)
                _ = v40
            }
            v41 := GenS3{F0: 255, F1: v4, F2: 42}
            if (!((!(!(false || false))) && (!(!false)))) {
                v6 = v6
                var v42 string = "hello"
                v43 := GenS3{F0: 32767, F1: (false && true), F2: 0}
                _ = v42
                _ = v43
            }
            _ = fi38
            _ = v41
        }
        var v44 int32 = -1
        for fi45 := 0; fi45 < 16; fi45++ {
            switch "hello" {
            case "\n\t":
                v6 = (v6 % ((v6) | 1))
            case "":
                v46 := GenS3{F0: 42, F1: (false && true), F2: 2}
                _ = v46
            }
            v5 = v5
            v47 := GenS3{F0: 2, F1: (!false), F2: 42}
            _ = fi45
            _ = v47
        }
        _ = fi37
        _ = v44
    }
    fmt.Println(v4, v6, v4)
    v48 := GenS1{F0: -2147483648, F1: 572}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
