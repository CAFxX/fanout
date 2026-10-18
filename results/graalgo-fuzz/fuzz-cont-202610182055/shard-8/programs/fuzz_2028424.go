package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int; F1 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := []string{"", ("\n\t" + "hello"), ("foo bar" + "a")}
    var v3 uint8 = 255
    var v4 int64 = -1
    var v5 uint32 = 60
    var v6 int32 = 2147483647
    delete(v1, 763)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint64 = 4294967295
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "hello":
        var v9 uint32 = v5
        _ = v9
    case "a":
        _ = len(v2)
    case "\n\t":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        switch 766 {
        case 350:
            if (!(false && false)) {
                v1[-493] = -560
                if true {
                    v6 = v6
                    v4 -= v4
                    v10 := GenS1{F0: 0, F1: true}
                    _ = v10
                }
            } else {
                var v11 float32 = 1e10
                v4 = -35
                if (false && false) {
                    for fi12 := 0; fi12 < 4; fi12++ {
                        if (fi12 > 578) {
                            fmt.Println(v5, v5, fi12)
                            var v13 float32 = v11
                            if (false && true) {
                                _ = len(v1)
                            }
                            _ = v13
                        } else {
                            v4 = v4
                            for fi14 := 0; fi14 < 9; fi14++ {
                                v15 := GenS1{F0: (v3 | v3), F1: (0 == -128)}
                                var v16 uint8 = (v3 - 121)
                                v17 := GenS1{F0: v16, F1: false}
                                _ = fi14
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                        }
                        v5 -= (v5 >> ((4294967295) & 63))
                        _ = fi12
                    }
                    v18 := GenS2{F0: 1, F1: 432}
                    v19 := GenS1{F0: 127, F1: false}
                    _ = v18
                    _ = v19
                } else {
                    var v20 uint16 = 255
                    if false {
                        fmt.Println(v11, v6, v11)
                    }
                    _ = v20
                }
                _ = v11
            }
        case -153:
            v21 := v1[111]
            _ = v21
        }
    }
    var v22 int16 = 964
    if true {
        for fi23 := 0; fi23 < 15; fi23++ {
            v24 := GenS2{F0: 1, F1: (-(fi23))}
            _ = fi23
            _ = v24
        }
    } else {
        for fi25 := 0; fi25 < 20; fi25++ {
            v3 += v3
            v22 = v22
            if len(v2) > 0 {
                v2[0] = (("a" + ("\n\t" + "a")) + "")
            }
            _ = fi25
        }
        if (2 < 4294967295) {
            var v26 uint64 = 334
            _ = v26
        } else {
            for fi27 := 0; fi27 < 14; fi27++ {
                fmt.Println(fi27, v4, v22)
                _ = fi27
            }
            if (v5 != (v5 / ((v5) | 1))) {
                var v28 string = "hello"
                v6 += (v6 / ((v6) | 1))
                _ = v28
            } else {
                v3 -= v3
            }
            switch ("\n\t" + "") {
            case "foo bar":
                var v29 bool = (0 < 21)
                _ = v29
            case "hello":
                v5 += (v5 - v5)
            case "":
                if (!(!(((v6 & v6) == 700) && true))) {
                    switch 42 {
                    case 127:
                        v4 = (v4 % ((-1) | 1))
                    case -288:
                        switch int(v5) {
                        case 127:
                            var v30 uint64 = 2
                            _ = v30
                        }
                    }
                    var v31 int8 = -128
                    _ = v31
                } else {
                    var v32 int8 = int8(v6)
                    v33 := GenS3{F0: 42, F1: v4}
                    fmt.Println(v4, v5, v4)
                    _ = v32
                    _ = v33
                }
            }
        }
    }
    switch "foo bar" {
    case "hello":
        v34 := GenS1{F0: v3, F1: (v6 <= -9)}
        _ = v34
    case "a":
        v6 = (v6 + v6)
    case "":
        fmt.Println(v6, v3, v5)
    }
    var v35 int = -1
    v6 -= (v6 << ((0) & 63))
    for fi36 := 0; fi36 < 1; fi36++ {
        switch "a" {
        case "a":
            _ = len(v2)
        }
        _ = fi36
    }
    var v37 int16 = 283
    v22 = -202
    if ((!false) && ((false && false) && (false && true))) {
        for fi38 := 0; fi38 < 2; fi38++ {
            fmt.Println(v37, v22, v3)
            for fi39 := 0; fi39 < 13; fi39++ {
                fmt.Println(v4)
                _ = fi39
            }
            _ = len(v2)
            _ = fi38
        }
        v22 -= (v37 - -400)
    } else {
        v5 = 2147483648
    }
    v40 := GenS4{F0: (v35 % (((v35 % (((v35 & (^(v35)))) | 1))) | 1)), F1: -32768}
    fmt.Println(v4)
    if (!true) {
        for fi41 := 0; fi41 < 2; fi41++ {
            if false {
                v35 += (v35 | fi41)
                for fi42 := 0; fi42 < 9; fi42++ {
                    var v43 uint64 = 448
                    if false {
                        var v44 int = (^(v35))
                        _ = v44
                    } else {
                        switch "a" {
                        case "foo bar":
                            var v45 int = fi42
                            _ = v45
                        case "":
                            switch "foo bar" {
                            case "\n\t":
                                v46 := GenS3{F0: 42, F1: 9223372036854775807}
                                _ = v46
                            default:
                                switch "foo bar" {
                                case "hello":
                                    v35 = -195
                                }
                            }
                        case "\n\t":
                            var v47 float32 = 1e10
                            _ = v47
                        }
                        v48 := v1[fi42]
                        _ = v48
                    }
                    _ = fi42
                    _ = v43
                }
            } else {
                var v49 int = int(v4)
                v6 += v6
                switch -2147483648 {
                case -240:
                    var v50 int = v35
                    _ = v50
                case 4294967295:
                    for fi51 := 0; fi51 < 19; fi51++ {
                        v6 = v6
                        var v52 int = (^(v49))
                        fmt.Println(v6)
                        _ = fi51
                        _ = v52
                    }
                }
                _ = v49
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v53 int8 = -128
            _ = fi41
            _ = v53
        }
        v54 := GenS3{F0: 0, F1: 9223372036854775807}
        _ = v54
    }
    switch "a" {
    case "foo bar":
        for fi55 := 0; fi55 < 14; fi55++ {
            if true {
                fmt.Println(v6, v3)
                v22 = v37
                v56 := GenS3{F0: 0, F1: 1}
                _ = v56
            } else {
                fmt.Println(v3, v6)
                var v57 string = (("\n\t" + ("a" + "a")) + "foo bar")
                v22 = (^(v22))
                _ = v57
            }
            for fi58 := 0; fi58 < 5; fi58++ {
                v35 += 127
                _ = fi58
            }
            v5 = 42
            _ = fi55
        }
    case "":
        v1[(v35 ^ (v35 - v35))] = v35
    case "a":
        v59 := GenS4{F0: v35, F1: (v37 * int16(v4))}
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
