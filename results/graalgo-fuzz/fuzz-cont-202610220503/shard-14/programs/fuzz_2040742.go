package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 uint; F1 int; F2 int64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint16; F1 uint8; F2 string }
type GenS11 struct { F0 string; F1 uint16; F2 uint }
type GenS12 struct { F0 int }

func main() {
    v1 := []uint32{875, 65535}
    var v2 int8 = 42
    v3 := []string{("hello" + "hello"), ("a" + ("" + "hello")), ("hello" + ""), "\n\t"}
    v4 := []int{-9223372036854775808}
    v5 := GenS1{F0: 32767}
    if false {
        var v6 uint = 2
        v2 += 127
        v2 -= v2
        _ = v6
    } else {
        var v7 int = 1
        _ = v7
    }
    if len(v3) > 0 {
        v3[0] = "a"
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 uint8 = 25
                _ = fi8
                _ = v9
            }
            if true {
                for fi10 := 0; fi10 < 3; fi10++ {
                    fmt.Println(v2, v2)
                    _ = fi10
                }
            } else {
                v2 = -37
                var v11 int16 = 1
                fmt.Println(v2, v11)
                _ = v11
            }
            _ = len(v4)
        }
    case "":
        var v12 int64 = 361
        _ = v12
    case "a":
        switch "a" {
        case "foo bar":
            var v13 int32 = 0
            _ = v13
        case "hello":
            if (!((!true) || true)) {
                for fi14 := 0; fi14 < 19; fi14++ {
                    for fi15 := 0; fi15 < 13; fi15++ {
                        fmt.Println(fi15)
                        _ = fi15
                    }
                    v2 = v2
                    for fi16 := 0; fi16 < 5; fi16++ {
                        if false {
                            var v17 float64 = (1.5 / ((1e10 - 0.0) + 1.5))
                            if true {
                                var v18 string = "\n\t"
                                v1 = append(v1, 387)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v19 := GenS2{F0: fi16}
                                switch ("a" + (v18 + v18)) {
                                case "\n\t":
                                    switch "foo bar" {
                                    case "foo bar":
                                        v18 = "foo bar"
                                    default:
                                        v20 := GenS3{F0: (v18 + v18), F1: 0}
                                        _ = v20
                                    }
                                case "hello":
                                    v21 := GenS4{F0: 144}
                                    _ = v21
                                case "foo bar":
                                    v22 := GenS5{F0: 1, F1: v18}
                                    _ = v22
                                default:
                                    v23 := GenS2{F0: 4294967295}
                                    _ = v23
                                }
                                _ = v18
                                _ = v19
                            } else {
                                v17 -= v17
                                var v24 int32 = 42
                                _ = v24
                            }
                            _ = v17
                        } else {
                            var v25 uint64 = 2
                            var v26 uint = 0
                            _ = v25
                            _ = v26
                        }
                        var v27 uint64 = 0
                        v27 = (v27 - 790)
                        fmt.Println(fi14, fi14)
                        _ = fi16
                        _ = v27
                    }
                    _ = fi14
                }
                var v28 uint8 = 2
                _ = v28
            }
        case "a":
            if len(v4) > 0 {
                v4[0] = 62
            }
        default:
            v2 = (v2 + v2)
        }
    }
    v29 := GenS6{F0: 0}
    var v30 uint32 = 1
    v2 += (v2 << ((772) & 63))
    switch 32768 {
    case 5:
        var v31 uint = 18446744073709551615
        _ = v31
    case 2147483647:
        fmt.Println(v30, v30)
    }
    fmt.Println(v2)
    v2 = 0
    for fi32 := 0; fi32 < 2; fi32++ {
        var v33 bool = true
        _ = fi32
        _ = v33
    }
    v34 := GenS7{F0: 1, F1: -9223372036854775808}
    if false {
        v30 = v30
    }
    switch "" {
    case "":
        var v35 uint16 = 65535
        _ = v35
    case "hello":
        var v36 float64 = (-2.25 + 1.5)
        _ = v36
    }
    if (81 > 0) {
        v30 += 4294967295
    }
    switch "hello" {
    case "a":
        if ((v2 == (v2 ^ -15)) && true) {
            if ("\n\t" == "a") {
                v37 := GenS8{F0: 2, F1: -2147483648, F2: 0}
                _ = v37
            }
        } else {
            v30 = 42
            fmt.Println(v30, v30)
            if false {
                if (-9223372036854775808 >= -446) {
                    if true {
                        v38 := GenS9{F0: 255}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v38
                    } else {
                        fmt.Println(v30)
                        if len(v1) > 0 {
                            v1[0] = (v30 ^ v30)
                        }
                    }
                    v39 := GenS8{F0: 454, F1: -9223372036854775808, F2: -1}
                    switch -1 {
                    case 4294967295:
                        for fi40 := 0; fi40 < 18; fi40++ {
                            v30 = (v30 ^ v30)
                            _ = fi40
                        }
                    }
                    _ = v39
                }
                v41 := GenS10{F0: 32767, F1: 2, F2: ("a" + "")}
                _ = v41
            } else {
                var v42 int = 723
                fmt.Println(v42, v2)
                _ = v42
            }
        }
    case "hello":
        var v43 uint32 = 42
        _ = v43
    case "\n\t":
        if (false || false) {
            v44 := GenS10{F0: 32767, F1: 127, F2: ("hello" + "hello")}
            fmt.Println(v30, v30)
            for fi45 := 0; fi45 < 12; fi45++ {
                if (((127 == 121) && (!(true && true))) || (true || true)) {
                    for fi46 := 0; fi46 < 15; fi46++ {
                        var v47 bool = ("hello" < "")
                        var v48 float64 = (0.0 - float64((0.0 + 1e10)))
                        v48 = v48
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v30 = v30
                    v49 := GenS11{F0: "\n\t", F1: 1, F2: 397}
                    _ = v49
                }
                _ = fi45
            }
            _ = v44
        }
    default:
        var v50 uint16 = 1
        _ = v50
    }
    if true {
        for fi51 := 0; fi51 < 3; fi51++ {
            if (true && false) {
                switch (fi51 >> ((427) & 63)) {
                case 128:
                    for fi52 := 0; fi52 < 11; fi52++ {
                        v2 -= -128
                        _ = fi52
                    }
                default:
                    fmt.Println(v2)
                }
                fmt.Println(v30, fi51, v2)
                v2 = (v2 + (v2 << ((446) & 63)))
            }
            for fi53 := 0; fi53 < 1; fi53++ {
                if true {
                    v54 := GenS12{F0: fi51}
                    _ = v54
                } else {
                    v55 := GenS7{F0: 127, F1: -725}
                    v30 += (^(v30))
                    fmt.Println(v2, v30)
                    _ = v55
                }
                var v56 int32 = 373
                _ = fi53
                _ = v56
            }
            _ = fi51
        }
        var v57 int64 = 864
        _ = len(v3)
        _ = v57
    } else {
        var v58 int32 = -1
        if (((^(v30)) > 269) || true) {
            if true {
                fmt.Println(v58)
                var v59 float64 = 1e10
                _ = v59
            } else {
                v58 = v58
            }
            var v60 bool = (int8(v58) < v2)
            _ = v60
        }
        _ = v58
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
