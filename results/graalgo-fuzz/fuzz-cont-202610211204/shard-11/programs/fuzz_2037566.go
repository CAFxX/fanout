package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int8 }

func main() {
    v1 := []string{"", "", "", ("a" + "hello")}
    var v2 int = 32767
    var v3 bool = true
    var v4 int16 = -650
    v5 := []uint32{2147483648, 2147483648}
    var v6 int = v2
    switch "\n\t" {
    case "\n\t":
        for fi7 := 0; fi7 < 2; fi7++ {
            v6 += (fi7 | fi7)
            _ = fi7
        }
    case "foo bar":
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "foo bar":
            for fi8 := 0; fi8 < 16; fi8++ {
                var v9 uint8 = 38
                _ = fi8
                _ = v9
            }
        }
    case "hello":
        v5 = append(v5, 2147483648)
    }
    switch "a" {
    case "foo bar":
        var v10 int8 = -1
        _ = v10
    }
    switch "\n\t" {
    case "a":
        var v11 uint8 = 185
        _ = v11
    }
    var v12 string = ""
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v2, v3, v3)
        _ = fi13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS1{F0: 0, F1: 1}
        _ = fi14
        _ = v15
    }
    v3 = v3
    var v16 bool = v3
    for fi17 := 0; fi17 < 16; fi17++ {
        v18 := GenS2{F0: 163, F1: (v4 * 1000)}
        _ = fi17
        _ = v18
    }
    if v16 {
        switch "\n\t" {
        case "hello":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        switch v6 {
        case 984:
            if v3 {
                var v19 int16 = (^(v4))
                _ = v19
            } else {
                if v3 {
                    var v20 string = "hello"
                    var v21 bool = v16
                    v20 = v12
                    _ = v20
                    _ = v21
                } else {
                    if (0 > 42) {
                        switch ("foo bar" + v12) {
                        case "":
                            v5 = append(v5, 2147483647)
                        case "foo bar":
                            v3 = false
                        default:
                            switch "foo bar" {
                            case "hello":
                                v2 = (v2 / ((128) | 1))
                            case "a":
                                fmt.Println(v4)
                            case "foo bar":
                                fmt.Println(v16, v16, v16)
                            default:
                                if true {
                                    v22 := GenS2{F0: 807, F1: 1}
                                    _ = v22
                                }
                            }
                        }
                        var v23 int16 = v4
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 4; fi24++ {
                        switch "foo bar" {
                        case "hello":
                            v3 = (255 < 186)
                        case "\n\t":
                            switch -9223372036854775808 {
                            case 32767:
                                v6 = (fi24 >> ((424) & 63))
                            case 128:
                                fmt.Println(v16, fi24)
                            case 42:
                                var v25 uint32 = 526
                                _ = v25
                            }
                        case "foo bar":
                            _ = len(v1)
                        default:
                            for fi26 := 0; fi26 < 17; fi26++ {
                                v2 = 2147483647
                                if len(v5) > 0 {
                                    v5[0] = 145
                                }
                                v27 := GenS2{F0: 42, F1: v4}
                                var v28 int8 = -128
                                v29 := GenS2{F0: 42, F1: (v4 & v4)}
                                _ = fi26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                        }
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v3 = ((v12 <= v12) && v16)
                            var v31 uint32 = 65535
                            if (15 >= 127) {
                                var v32 int16 = (v4 / ((-32768) | 1))
                                _ = v32
                            }
                            var v33 uint = 255
                            v34 := GenS1{F0: 9223372036854775807, F1: 42}
                            _ = fi30
                            _ = v31
                            _ = v33
                            _ = v34
                        }
                        var v35 bool = (38 >= 1)
                        _ = fi24
                        _ = v35
                    }
                }
                v12 = v12
            }
        default:
            v3 = true
        }
    }
    switch "" {
    case "a":
        for fi36 := 0; fi36 < 13; fi36++ {
            v6 = (fi36 + (-(v6)))
            if v3 {
                fmt.Println(fi36, v4, v12)
            }
            _ = fi36
        }
    case "foo bar":
        if (!true) {
            fmt.Println(v12, v2, v2)
        } else {
            switch v12 {
            case "":
                for fi37 := 0; fi37 < 7; fi37++ {
                    var v38 int8 = 0
                    if (!v16) {
                        v2 -= 32768
                        v3 = v16
                        if (v16 || false) {
                            var v39 uint16 = 42
                            var v40 uint8 = 128
                            _ = v39
                            _ = v40
                        }
                    }
                    var v41 string = v12
                    _ = fi37
                    _ = v38
                    _ = v41
                }
            case "\n\t":
                v4 -= (v4 & (^(v4)))
            case "hello":
                for fi42 := 0; fi42 < 18; fi42++ {
                    for fi43 := 0; fi43 < 17; fi43++ {
                        var v44 int = v6
                        _ = fi43
                        _ = v44
                    }
                    switch (v6 + v6) {
                    case -227:
                        for fi45 := 0; fi45 < 20; fi45++ {
                            var v46 bool = (-2147483648 >= -1)
                            if (!(v2 < (v2 | -9223372036854775808))) {
                                v47 := GenS3{F0: 2, F1: 18}
                                _ = v47
                            } else {
                                var v48 float32 = (0.0 / 0.1)
                                fmt.Println(fi42, v12)
                                _ = v48
                            }
                            _ = fi45
                            _ = v46
                        }
                    case -2147483648:
                        fmt.Println(v12, v4)
                    case 32768:
                        v4 += v4
                    }
                    v49 := GenS4{F0: 876, F1: 127}
                    _ = fi42
                    _ = v49
                }
            }
        }
    case "\n\t":
        v50 := GenS5{F0: 2}
        _ = v50
    default:
        switch (v6 >> ((255) & 63)) {
        case 32768:
            for fi51 := 0; fi51 < 16; fi51++ {
                var v52 uint8 = 1
                _ = fi51
                _ = v52
            }
        case 1:
            var v53 int16 = 1
            _ = v53
        default:
            v16 = false
        }
    }
    switch (("hello" + (v12 + "foo bar")) + (v12 + v12)) {
    case "foo bar":
        v54 := GenS2{F0: 4294967295, F1: (v4 ^ v4)}
        _ = v54
    case "\n\t":
        if (!(false || v16)) {
            v3 = ((!v3) || (255 != 42))
            v55 := GenS1{F0: 615, F1: 1}
            v16 = (!(275 != 255))
            _ = v55
        }
    default:
        var v56 int64 = 9223372036854775807
        _ = v56
    }
    var v57 int = 420
    switch "hello" {
    case "\n\t":
        if (2 != 1) {
            v58 := GenS6{F0: v16, F1: 42}
            fmt.Println(v3, v6, v4)
            _ = v58
        } else {
            if true {
                var v59 uint8 = 228
                _ = v59
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (!v3) {
                    if false {
                        fmt.Println(v57)
                        v3 = (!v3)
                    } else {
                        v60 := GenS5{F0: 18446744073709551615}
                        _ = v60
                    }
                    fmt.Println(v4, v6, v4)
                }
            }
            v4 = 32767
        }
    case "hello":
        v61 := GenS1{F0: 42, F1: 1}
        _ = v61
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
