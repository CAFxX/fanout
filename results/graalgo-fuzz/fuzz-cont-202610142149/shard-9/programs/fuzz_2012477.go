package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int32) uint16 {
    v6 := []int{111, -81, 127, -277}
    if true {
        p0 += 1
    }
    v7 := GenS1{F0: 1, F1: 1000}
    for fi8 := 0; fi8 < 9; fi8++ {
        switch (fi8 + (fi8 | 277)) {
        case 1:
            p0 = p0
        case -934:
            switch "" {
            case "hello":
                if false {
                    p0 += p0
                } else {
                    switch (fi8 / ((fi8) | 1)) {
                    case 42:
                        v9 := GenS1{F0: 2, F1: -32768}
                        _ = v9
                    case -1:
                        if (-67 < -1) {
                            var v10 int8 = -8
                            if (!(2 >= 145)) {
                                _ = len(v6)
                                var v11 uint16 = 65535
                                _ = v11
                            }
                            _ = v10
                        }
                    }
                }
            case "a":
                v12 := GenS1{F0: 2, F1: -1}
                _ = v12
            case "\n\t":
                p0 = p0
            default:
                fmt.Println(p0)
            }
        default:
            var v13 int64 = 9223372036854775807
            _ = v13
        }
        var v14 int = (fi8 % (((^(fi8))) | 1))
        if true {
            v14 = 128
            switch (^(v14)) {
            case 42:
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS1{F0: 1, F1: 730}
                    var v17 int = 9223372036854775807
                    switch ("foo bar" + (("a" + "\n\t") + ("\n\t" + ("foo bar" + (("hello" + (("foo bar" + "foo bar") + "a")) + ("hello" + ("hello" + "hello"))))))) {
                    case "hello":
                        var v18 uint = 0
                        _ = v18
                    case "\n\t":
                        v6 = append(v6, 127)
                    default:
                        v17 += fi8
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
        }
        _ = fi8
        _ = v14
    }
    _ = p0
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int, p1 bool, p2 uint) uint16 {
    v19 := GenS2{F0: 1}
    for fi20 := 0; fi20 < 8; fi20++ {
        if p1 {
            var v21 bool = p1
            p1 = v21
            _ = v21
        }
        var v22 float32 = -2.25
        _ = fi20
        _ = v22
    }
    v23 := []int{}
    p0 -= -9223372036854775808
    fmt.Println(p0, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v23
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 3.14159
    var v3 int16 = -300
    v4 := []string{"a", "hello", ("" + "foo bar")}
    var v5 int64 = -1
    switch "hello" {
    case "foo bar":
        var v24 int = -404
        _ = v24
    case "":
        fmt.Println(v2)
    }
    var v25 uint8 = 44
    var v26 int8 = -128
    if true {
        var v27 float32 = (1e10 / 1.0)
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        fmt.Println(v2, v2)
        _ = fi28
    }
    v5 -= -9223372036854775808
    v4 = append(v4, "")
    v3 += v3
    if true {
        switch -77 {
        case -9223372036854775808:
            v29 := GenS3{F0: 516, F1: 1000}
            _ = v29
        case -2147483648:
            for fi30 := 0; fi30 < 5; fi30++ {
                switch 436 {
                case 127:
                    var v31 int16 = v3
                    _ = v31
                case -1:
                    v32 := GenS1{F0: 65535, F1: v3}
                    _ = v32
                case -497:
                    var v33 uint64 = 789
                    _ = v33
                }
                switch 135 {
                case -146:
                    if false {
                        v5 -= v5
                        for fi34 := 0; fi34 < 11; fi34++ {
                            v35 := v1["hello"]
                            v35 += fi30
                            _ = fi34
                            _ = v35
                        }
                        switch "hello" {
                        case "":
                            fmt.Println(v26)
                        case "foo bar":
                            v3 = (^(v3))
                        }
                    } else {
                        if true {
                            v3 -= (v3 | (-(v3)))
                            v36 := GenS4{F0: v5}
                            switch (fi30 % ((276) | 1)) {
                            case 429:
                                var v37 int64 = 208
                                _ = v37
                            }
                            _ = v36
                        }
                    }
                }
                for fi38 := 0; fi38 < 17; fi38++ {
                    v3 = (v3 | int16(v25))
                    v39 := GenS5{F0: 2147483647}
                    _ = fi38
                    _ = v39
                }
                _ = fi30
            }
        case 4294967295:
            if ((v3 != v3) && true) {
                var v40 bool = false
                for fi41 := 0; fi41 < 5; fi41++ {
                    var v42 uint = 2
                    v43 := GenS3{F0: (v42 >> ((v42) & 63)), F1: -32768}
                    switch 32767 {
                    case 381:
                        v44 := GenS3{F0: v42, F1: (v3 | (v3 * v3))}
                        _ = v44
                    case 9223372036854775807:
                        if (!v40) {
                            fmt.Println(v3, v42)
                            v40 = (v5 <= v5)
                        } else {
                            for fi45 := 0; fi45 < 3; fi45++ {
                                var v46 float64 = 0.0
                                _ = fi45
                                _ = v46
                            }
                            var v47 float64 = v2
                            _ = v47
                        }
                    }
                    _ = fi41
                    _ = v42
                    _ = v43
                }
                switch "" {
                case "":
                    var v48 string = "foo bar"
                    _ = v48
                default:
                    v49 := GenS5{F0: 42}
                    _ = v49
                }
                _ = v40
            } else {
                v26 += (v26 & 100)
                fmt.Println(v5)
                switch 4294967295 {
                case 4294967295:
                    v2 = 3.14159
                case 1:
                    var v50 int32 = 956
                    _ = v50
                default:
                    for fi51 := 0; fi51 < 18; fi51++ {
                        if true {
                            var v52 int8 = int8(v3)
                            _ = v52
                        } else {
                            var v53 uint8 = 0
                            _ = v53
                        }
                        _ = fi51
                    }
                }
            }
        default:
            v3 = v3
        }
        var v54 float64 = 0.1
        for fi55 := 0; fi55 < 3; fi55++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) > 0 {
                v4[0] = ("foo bar" + "foo bar")
            }
            if len(v4) > 0 {
                v4[0] = "a"
            }
            _ = fi55
        }
        _ = v54
    } else {
        for fi56 := 0; fi56 < 2; fi56++ {
            var v57 uint8 = 2
            v25 -= (v25 * v25)
            _ = fi56
            _ = v57
        }
    }
    if false {
        v25 = (v25 & (v25 / ((v25) | 1)))
        for fi58 := 0; fi58 < 6; fi58++ {
            v2 += (v2 - v2)
            switch (fi58 - fi58) {
            case 2147483647:
                var v59 string = "a"
                _ = v59
            }
            _ = fi58
        }
    }
    var v60 int8 = v26
    switch -9223372036854775808 {
    case -51:
        var v61 float32 = 0.1
        _ = v61
    default:
        var v62 uint64 = 0
        _ = v62
    }
    var v63 uint8 = v25
    for fi64 := 0; fi64 < 5; fi64++ {
        var v65 uint = 308
        _ = fi64
        _ = v65
    }
    v66 := GenS4{F0: (v5 >> ((255) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v63
    _ = v66
}
