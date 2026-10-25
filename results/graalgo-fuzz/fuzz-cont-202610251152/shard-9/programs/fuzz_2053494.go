package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 9223372036854775807
    v3 := []string{"\n\t"}
    var v4 int = 9223372036854775807
    var v5 uint = 2
    v6 := []uint32{42, 4294967295}
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v7 int = v4
    for fi8 := 0; fi8 < 14; fi8++ {
        if true {
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 int = 42
                if true {
                    fmt.Println(fi9, v5)
                }
                _ = fi9
                _ = v10
            }
            fmt.Println(v7)
        }
        fmt.Println(v2)
        _ = fi8
    }
    delete(v1, "a")
    v11 := GenS1{F0: 2}
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int8 = -1
        _ = fi12
        _ = v13
    }
    v7 = (v7 & (v4 & 128))
    var v14 string = "\n\t"
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 float64 = -2.25
        v17 := GenS1{F0: 99}
        v18 := GenS1{F0: 1}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    if ((^(v5)) != 1) {
        for fi19 := 0; fi19 < 14; fi19++ {
            fmt.Println(v14)
            switch v14 {
            case "hello":
                v20 := GenS1{F0: 1}
                _ = v20
            default:
                var v21 float32 = (1.5 * 0.0)
                _ = v21
            }
            if false {
                v22 := GenS1{F0: 178}
                _ = v22
            } else {
                v6 = append(v6, 1)
                if false {
                    v2 = (v2 << ((0) & 63))
                    switch ("a" + "hello") {
                    case "":
                        var v23 int = fi19
                        _ = v23
                    case "hello":
                        for fi24 := 0; fi24 < 10; fi24++ {
                            v2 = (v2 << ((4294967295) & 63))
                            var v25 int = v4
                            switch v14 {
                            case "foo bar":
                                v26 := GenS1{F0: 127}
                                _ = v26
                            }
                            if true {
                                v14 = v14
                                fmt.Println(v14)
                            } else {
                                fmt.Println(v5, fi19, v25)
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                    var v27 int64 = v2
                    _ = v27
                }
            }
            _ = fi19
        }
        var v28 int32 = 101
        fmt.Println(v7, v14)
        _ = v28
    }
    var v29 int64 = (-(v2))
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    if false {
        var v30 uint16 = 109
        for fi31 := 0; fi31 < 18; fi31++ {
            for fi32 := 0; fi32 < 12; fi32++ {
                fmt.Println(fi32, v2, v4)
                _ = fi32
            }
            var v33 int16 = -624
            v4 = (fi31 * -9223372036854775808)
            _ = fi31
            _ = v33
        }
        _ = v30
    } else {
        var v34 float64 = ((3.14159 - -2.25) / 1.0)
        switch ("hello" + (v14 + ("" + v14))) {
        case "\n\t":
            v35 := GenS1{F0: 136}
            _ = v35
        default:
            delete(v1, "\n\t")
        }
        for fi36 := 0; fi36 < 3; fi36++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi36
        }
        _ = v34
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        v29 -= 1
        _ = fi37
    }
    if false {
        if (2147483647 > 42) {
            var v38 uint32 = 593
            switch v14 {
            case "\n\t":
                var v39 int32 = -102
                _ = v39
            case "":
                for fi40 := 0; fi40 < 15; fi40++ {
                    v14 = ("hello" + v14)
                    if true {
                        if false {
                            for fi41 := 0; fi41 < 10; fi41++ {
                                if true {
                                    if (v14 == ("foo bar" + ("" + ((v14 + v14) + v14)))) {
                                        switch fi41 {
                                        case 0:
                                            if false {
                                                v5 -= v5
                                                var v42 int64 = v2
                                                switch v14 {
                                                case "foo bar":
                                                    v43 := GenS2{F0: ((!true) || (128 == 66)), F1: -26}
                                                    _ = v43
                                                }
                                                _ = v42
                                            }
                                        }
                                        switch v14 {
                                        case "a":
                                            if false {
                                                var v44 int = v7
                                                switch int(v2) {
                                                case 128:
                                                    var v45 int8 = 42
                                                    v29 -= v29
                                                    v4 -= 897
                                                    _ = v45
                                                case -414:
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                _ = v44
                                            }
                                        default:
                                            fmt.Println(v29)
                                        }
                                    }
                                    fmt.Println(v5, v5, v2)
                                    v14 = v14
                                } else {
                                    var v46 uint64 = 1
                                    var v47 int32 = 2147483647
                                    v29 += (v29 ^ v2)
                                    _ = v46
                                    _ = v47
                                }
                                _ = fi41
                            }
                        } else {
                            switch "\n\t" {
                            case "a":
                                for fi48 := 0; fi48 < 15; fi48++ {
                                    var v49 float64 = 3.14159
                                    var v50 int16 = 32767
                                    var v51 float64 = v49
                                    _ = fi48
                                    _ = v49
                                    _ = v50
                                    _ = v51
                                }
                            case "foo bar":
                                if (101 < 1000) {
                                    var v52 uint = (^(v5))
                                    var v53 uint8 = 2
                                    _ = len(v6)
                                    _ = v52
                                    _ = v53
                                }
                            default:
                                v7 += (fi40 | 127)
                            }
                            if true {
                                if false {
                                    fmt.Println(v29)
                                    fmt.Println(v29, v5, v29)
                                    var v54 uint8 = 1
                                    _ = v54
                                } else {
                                    v55 := GenS3{F0: 0, F1: -1}
                                    _ = v55
                                }
                                fmt.Println(v2)
                            }
                        }
                    } else {
                        switch (v7 << ((168) & 63)) {
                        case 4294967295:
                            fmt.Println(v4, v7)
                        }
                    }
                    _ = fi40
                }
            }
            for fi56 := 0; fi56 < 1; fi56++ {
                var v57 int8 = 127
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi56
                _ = v57
            }
            _ = v38
        }
        if (18446744073709551615 == 18446744073709551615) {
            for fi58 := 0; fi58 < 6; fi58++ {
                if true {
                    var v59 bool = (!true)
                    _ = v59
                }
                var v60 int16 = -32768
                _ = fi58
                _ = v60
            }
        }
        v5 = 867
    } else {
        v5 += v5
        _ = len(v1)
        for fi61 := 0; fi61 < 7; fi61++ {
            var v62 int32 = -1
            v1[(v14 + (v14 + "\n\t"))] = fi61
            _ = fi61
            _ = v62
        }
    }
    var v63 int = -445
    v5 += 489
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v63:", v63)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v7
}
