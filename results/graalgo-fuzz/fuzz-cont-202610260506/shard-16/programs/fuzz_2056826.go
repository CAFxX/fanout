package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int16; F1 int8; F2 int64 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 bool }

func main() {
    v1 := []string{("\n\t" + "foo bar"), "\n\t"}
    v2 := []string{"", "", "foo bar"}
    var v3 float64 = (float64(0.1) + 1e10)
    var v4 int = 42
    var v5 int64 = 9223372036854775807
    var v6 uint64 = 42
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    if (-323 >= -1) {
        v7 := GenS1{F0: 0, F1: 2147483648, F2: 1}
        _ = v7
    }
    switch (v4 & (v4 - (v4 << ((4294967295) & 63)))) {
    case 2147483647:
        for fi8 := 0; fi8 < 19; fi8++ {
            if true {
                v5 = 1
            } else {
                if false {
                    switch (^(v4)) {
                    case 1:
                        var v9 float64 = (v3 * -2.25)
                        _ = v9
                    }
                }
                v5 += -1
                var v10 uint64 = (v6 >> ((42) & 63))
                _ = v10
            }
            if false {
                if (175 > 0) {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    v5 = -837
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    v11 := GenS2{F0: 127}
                    v12 := GenS2{F0: 2}
                    _ = v11
                    _ = v12
                }
                if false {
                    switch v4 {
                    case 9223372036854775807:
                        fmt.Println(v4)
                    case 2147483647:
                        if (!((((v3 - (1e10 + v3)) + v3) + float64(1e10)) == (float64(float32(v3)) * v3))) {
                            if (2 != 2) {
                                v13 := GenS3{F0: 2147483647, F1: 0}
                                _ = len(v2)
                                var v14 float64 = float64(-2.25)
                                _ = v13
                                _ = v14
                            } else {
                                fmt.Println(v6)
                                v1 = append(v1, "a")
                            }
                            var v15 int64 = (v5 / ((v5) | 1))
                            _ = v15
                        } else {
                            fmt.Println(v3, v4)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    case 4294967295:
                        v6 += v6
                    default:
                        v5 -= v5
                    }
                } else {
                    v6 -= 1
                    var v16 uint8 = 2
                    if (!true) {
                        var v17 int8 = -75
                        _ = v17
                    }
                    _ = v16
                }
            }
            fmt.Println(v6, v5)
            _ = fi8
        }
    case -1:
        if true {
            switch v4 {
            case -142:
                v4 = v4
            case 32768:
                var v18 uint = 65535
                _ = v18
            }
        }
    }
    if ((true || (((!(false && false)) && (255 > 18446744073709551615)) || false)) && (128 >= 1)) {
        switch (-(v4)) {
        case 86:
            var v19 uint = 2
            _ = v19
        default:
            if true {
                if false {
                    _ = len(v2)
                    v20 := GenS1{F0: 42, F1: 2147483648, F2: v6}
                    _ = v20
                } else {
                    v21 := GenS4{F0: 214, F1: 0, F2: (v5 + 0)}
                    if true {
                        var v22 int32 = 1
                        var v23 int32 = (v22 * 2147483647)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v22
                        _ = v23
                    } else {
                        if true {
                            fmt.Println(v5, v6)
                        } else {
                            var v24 float64 = 0.1
                            fmt.Println(v3, v3)
                            var v25 float64 = (v24 / v3)
                            _ = v24
                            _ = v25
                        }
                    }
                    var v26 int32 = 571
                    _ = v21
                    _ = v26
                }
                v6 = v6
            } else {
                if (!((true && true) && (v5 <= v5))) {
                    v4 = v4
                    var v27 bool = true
                    fmt.Println(v3, v6)
                    _ = v27
                }
                if (true && true) {
                    var v28 uint64 = (v6 + v6)
                    _ = v28
                }
            }
        }
        var v29 uint32 = uint32(v6)
        if true {
            v30 := GenS5{F0: 42, F1: (v29 / (((v29 * (^(v29)))) | 1)), F2: 42}
            if ((!false) && true) {
                switch 1 {
                case 32767:
                    var v31 string = ""
                    _ = v31
                case 83:
                    for fi32 := 0; fi32 < 5; fi32++ {
                        var v33 bool = false
                        v34 := GenS6{F0: (-24 != -51), F1: (v5 >> ((861) & 63))}
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                case 42:
                    if false {
                        v4 = 0
                    }
                default:
                    v35 := GenS7{F0: 687}
                    _ = v35
                }
            } else {
                v4 -= v4
                if false {
                    switch "a" {
                    case "":
                        v4 = v4
                    }
                    var v36 uint = 2
                    _ = v36
                } else {
                    var v37 int = (v4 + 128)
                    _ = v37
                }
            }
            if (2147483647 != 1) {
                switch ("hello" + "foo bar") {
                case "foo bar":
                    v2 = append(v2, "hello")
                case "hello":
                    switch v4 {
                    case -1:
                        for fi38 := 0; fi38 < 3; fi38++ {
                            for fi39 := 0; fi39 < 12; fi39++ {
                                v3 += ((3.14159 + v3) - float64(1e10))
                                _ = fi39
                            }
                            v5 = v5
                            _ = fi38
                        }
                    case 2147483647:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 362:
                        v5 += 9223372036854775807
                    }
                }
            } else {
                v40 := GenS5{F0: 100, F1: v29, F2: 0}
                var v41 uint16 = 1
                var v42 int16 = -399
                _ = v40
                _ = v41
                _ = v42
            }
            _ = v30
        }
        _ = v29
    } else {
        var v43 uint8 = 127
        if false {
            v44 := GenS8{F0: 65535, F1: (55 != -104)}
            v43 = 255
            var v45 int = v4
            _ = v44
            _ = v45
        } else {
            var v46 uint16 = 2
            fmt.Println(v6)
            fmt.Println(v4, v3, v3)
            _ = v46
        }
        v47 := GenS8{F0: 418, F1: false}
        _ = v43
        _ = v47
    }
    v48 := GenS6{F0: true, F1: v5}
    for fi49 := 0; fi49 < 8; fi49++ {
        v4 -= v4
        var v50 int8 = -16
        _ = fi49
        _ = v50
    }
    fmt.Println(v3, v3)
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
