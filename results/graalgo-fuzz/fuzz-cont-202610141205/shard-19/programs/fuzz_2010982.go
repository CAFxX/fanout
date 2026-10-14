package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 bool }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint32 = 0
    var v2 uint16 = 65535
    var v3 int = -2147483648
    var v4 uint64 = 4294967295
    v5 := []string{("\n\t" + "hello"), "foo bar"}
    var v6 int = 32768
    v1 = (v1 >> ((42) & 63))
    if (-2.25 != (((1e10 * 1.5) - (((0.0 - ((0.1 / 1.5) - 0.0)) * float32(((3.14159 - -2.25) / float64(0.1)))) * -2.25)) / (3.14159 - ((1e10 / ((1.5 - -2.25) + float32((0.1 / 0.1)))) / (0.1 / float32(1.5)))))) {
        v1 = (v1 ^ (v1 & 2147483647))
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        if len(v5) > 0 {
            v5[0] = ("a" + "foo bar")
        }
        var v8 int8 = -87
        _ = fi7
        _ = v8
    }
    v2 = 1
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 bool = (!true)
        _ = fi9
        _ = v10
    }
    v1 += 2147483648
    if (!((!(1.5 == float32(float64(3.14159)))) && true)) {
        v11 := GenS1{F0: "", F1: (v4 << ((18446744073709551615) & 63))}
        _ = v11
    }
    var v12 float32 = (3.14159 - float32((-2.25 - -2.25)))
    v13 := GenS1{F0: "", F1: 4294967295}
    if true {
        fmt.Println(v2, v3)
        for fi14 := 0; fi14 < 1; fi14++ {
            switch "hello" {
            case "hello":
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 int16 = 32767
                    fmt.Println(v4)
                    if true {
                        fmt.Println(v4, v3, v16)
                    }
                    _ = fi15
                    _ = v16
                }
            case "foo bar":
                v6 = v3
            default:
                v17 := GenS1{F0: (("hello" + "") + ("a" + "")), F1: (v4 << ((683) & 63))}
                _ = v17
            }
            for fi18 := 0; fi18 < 9; fi18++ {
                switch "" {
                case "":
                    v19 := GenS1{F0: ("foo bar" + "foo bar"), F1: 42}
                    _ = v19
                case "\n\t":
                    v4 = 42
                default:
                    if false {
                        v20 := GenS1{F0: "\n\t", F1: 42}
                        fmt.Println(v1)
                        fmt.Println(v12, v2, v4)
                        _ = v20
                    }
                    v2 = (v2 >> ((42) & 63))
                    fmt.Println(fi18)
                    var v21 bool = ((("foo bar" == "") || false) || (28 == 240))
                    v5 = append(v5, (("\n\t" + ("" + "foo bar")) + ""))
                    v2 = v2
                    var v22 uint32 = v1
                    _ = v21
                    _ = v22
                }
                var v23 int = 32768
                _ = fi18
                _ = v23
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                v2 = 32767
                _ = fi24
            }
            _ = fi14
        }
    } else {
        if ((float32(1e10) != v12) && false) {
            v25 := GenS1{F0: (("\n\t" + "") + ("hello" + ("foo bar" + "a"))), F1: (v4 * 0)}
            _ = len(v5)
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS1{F0: "foo bar", F1: v4}
                _ = fi26
                _ = v27
            }
            _ = v25
        } else {
            var v28 int16 = 42
            v29 := GenS1{F0: "foo bar", F1: v4}
            v30 := GenS1{F0: (("" + "hello") + "hello"), F1: 4294967295}
            _ = v28
            _ = v29
            _ = v30
        }
    }
    if len(v5) > 0 {
        v5[0] = ""
    }
    var v31 string = ("" + "")
    if (0 != -128) {
        fmt.Println(v4, v3, v4)
        switch "" {
        case "hello":
            switch "hello" {
            case "\n\t":
                if true {
                    for fi32 := 0; fi32 < 15; fi32++ {
                        for fi33 := 0; fi33 < 10; fi33++ {
                            if false {
                                var v34 int32 = -1
                                v12 -= (3.14159 + (0.1 / v12))
                                _ = v34
                            } else {
                                switch -1 {
                                case -921:
                                    var v35 int = -732
                                    _ = v35
                                }
                                var v36 string = "a"
                                if false {
                                    var v37 bool = ((-2.25 != -2.25) || (568 < 2147483647))
                                    if true {
                                        var v38 int32 = 2147483647
                                        _ = v38
                                    } else {
                                        fmt.Println(v36, v6)
                                        switch v31 {
                                        case "foo bar":
                                            switch v6 {
                                            case 610:
                                                v39 := GenS1{F0: (v36 + ("foo bar" + v31)), F1: v4}
                                                var v40 int8 = 127
                                                _ = v39
                                                _ = v40
                                            }
                                        case "a":
                                            switch fi32 {
                                            case 1:
                                                v1 -= (^(v1))
                                            case -1:
                                                var v41 string = "a"
                                                _ = v41
                                            default:
                                                fmt.Println(v12)
                                            }
                                        case "":
                                            v42 := GenS2{F0: 1, F1: 537, F2: v37}
                                            _ = v42
                                        }
                                    }
                                    switch v36 {
                                    case "":
                                        var v43 bool = ((127 > -123) && v37)
                                        _ = len(v5)
                                        v43 = v37
                                        _ = v43
                                    case "a":
                                        var v44 float32 = v12
                                        _ = v44
                                    default:
                                        var v45 int = (v3 / (((v6 + fi33)) | 1))
                                        _ = v45
                                    }
                                    _ = v37
                                } else {
                                    var v46 bool = true
                                    _ = v46
                                }
                                _ = v36
                            }
                            var v47 float32 = v12
                            _ = fi33
                            _ = v47
                        }
                        _ = fi32
                    }
                    v5 = append(v5, v31)
                    if len(v5) > 0 {
                        v5[0] = ("" + v31)
                    }
                }
            }
        case "\n\t":
            var v48 uint = 42
            _ = v48
        case "a":
            var v49 int64 = 1
            _ = v49
        default:
            v50 := GenS3{F0: -1, F1: true}
            _ = v50
        }
        fmt.Println(v4, v1)
    } else {
        for fi51 := 0; fi51 < 3; fi51++ {
            switch v6 {
            case -2147483648:
                v52 := GenS4{F0: 0}
                _ = v52
            }
            _ = fi51
        }
        if true {
            if (1 < (v4 - 1)) {
                fmt.Println(v2, v31)
                for fi53 := 0; fi53 < 18; fi53++ {
                    if false {
                        var v54 uint32 = v1
                        _ = v54
                    }
                    v5 = append(v5, ((v31 + v31) + v31))
                    _ = fi53
                }
            }
            var v55 uint16 = (v2 >> ((42) & 63))
            v56 := GenS2{F0: -1, F1: 1000, F2: (42 == 42)}
            _ = v55
            _ = v56
        } else {
            if false {
                fmt.Println(v4)
                var v57 string = (v31 + v31)
                _ = v57
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v31, v6, v4)
        }
    }
    var v58 uint16 = v2
    v2 += v58
    var v59 float32 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
}
