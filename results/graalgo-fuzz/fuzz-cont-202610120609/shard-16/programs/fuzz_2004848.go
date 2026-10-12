package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 string; F2 bool }
type GenS3 struct { F0 uint16; F1 uint; F2 uint64 }

func main() {
    var v1 int = 0
    var v2 uint16 = 1
    var v3 bool = (-695 >= 42)
    v4 := []int{v1, (v1 << ((42) & 63)), v1}
    var v5 uint16 = (v2 * (v2 * (v2 ^ 836)))
    var v6 uint8 = 143
    var v7 uint16 = (v5 % (((v2 ^ 65535)) | 1))
    v8 := GenS1{F0: 42, F1: 296}
    fmt.Println(v6, v2)
    var v9 float64 = 1.5
    switch v1 {
    case -2147483648:
        for fi10 := 0; fi10 < 14; fi10++ {
            if (970 >= 2) {
                fmt.Println(v3)
            }
            _ = fi10
        }
    case 32767:
        v3 = v3
    }
    var v11 float64 = (((v9 / (0.0 + 3.14159)) / (v9 - v9)) * (v9 / v9))
    var v12 int8 = -1
    if (float32(0.0) == 3.14159) {
        switch "hello" {
        case "a":
            for fi13 := 0; fi13 < 5; fi13++ {
                switch (fi13 << ((2) & 63)) {
                case -878:
                    if len(v4) > 0 {
                        v4[0] = 127
                    }
                }
                var v14 uint = 42
                var v15 uint64 = 2
                _ = fi13
                _ = v14
                _ = v15
            }
        default:
            var v16 int16 = 42
            _ = v16
        }
        var v17 float32 = (1e10 * 1.5)
        v18 := GenS1{F0: -534, F1: 2}
        _ = v17
        _ = v18
    }
    var v19 string = ""
    var v20 uint = 255
    if (!v3) {
        switch "\n\t" {
        case "":
            var v21 uint = (v20 - v20)
            _ = v21
        case "a":
            if false {
                if (-906 < (^(v1))) {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        var v23 uint64 = 2
                        v3 = true
                        _ = fi22
                        _ = v23
                    }
                    var v24 uint32 = 154
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 bool = (false && (2147483647 <= -2147483648))
                        switch v19 {
                        case "a":
                            fmt.Println(v12, v1, v3)
                        case "foo bar":
                            switch v19 {
                            case "":
                                var v27 int8 = (v12 | int8(fi25))
                                _ = v27
                            case "\n\t":
                                v2 = 255
                            }
                        case "":
                            for fi28 := 0; fi28 < 17; fi28++ {
                                v29 := GenS1{F0: -9223372036854775808, F1: 1}
                                v30 := GenS1{F0: -240, F1: 18446744073709551615}
                                v4 = append(v4, -56)
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                        }
                        v3 = v26
                        _ = fi25
                        _ = v26
                    }
                }
                var v31 uint32 = 2
                v32 := GenS1{F0: 9223372036854775807, F1: 320}
                _ = v31
                _ = v32
            }
        }
        if (v3 || (false && false)) {
            for fi33 := 0; fi33 < 1; fi33++ {
                v4 = append(v4, v1)
                var v34 int64 = 0
                _ = fi33
                _ = v34
            }
        }
    } else {
        switch v19 {
        case "foo bar":
            for fi35 := 0; fi35 < 7; fi35++ {
                for fi36 := 0; fi36 < 9; fi36++ {
                    v4 = append(v4, (fi35 >> (((v20 + 255)) & 63)))
                    _ = fi36
                }
                _ = fi35
            }
        default:
            if false {
                var v37 uint64 = 4294967295
                _ = v37
            } else {
                v38 := GenS2{F0: (^(v6)), F1: ("" + (v19 + "hello")), F2: (892 != -504)}
                var v39 string = v19
                if (v3 && false) {
                    if (!(v39 < v19)) {
                        var v40 string = "foo bar"
                        _ = v40
                    }
                }
                _ = v38
                _ = v39
            }
        }
        if (262 >= 65535) {
            v2 = (v7 / (((v2 / ((v7) | 1))) | 1))
            switch v1 {
            case -12:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 0:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 4294967295:
                var v41 uint32 = 754
                _ = v41
            }
        } else {
            v6 += 1
            fmt.Println(v9)
            v7 = (v5 + 1)
        }
    }
    if (42 != 208) {
        switch v19 {
        case "foo bar":
            _ = len(v4)
        default:
            for fi42 := 0; fi42 < 5; fi42++ {
                v43 := GenS3{F0: uint16(v20), F1: (v20 >> ((v20) & 63)), F2: 515}
                var v44 int = fi42
                v45 := GenS1{F0: 0, F1: 18446744073709551615}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
        }
        v4 = append(v4, v1)
    } else {
        if len(v4) > 0 {
            v4[0] = (v1 ^ v1)
        }
        for fi46 := 0; fi46 < 5; fi46++ {
            v47 := GenS3{F0: v7, F1: 95, F2: 0}
            _ = fi46
            _ = v47
        }
        if (!true) {
            if (0 == (v12 % ((v12) | 1))) {
                if (100 == v12) {
                    v48 := GenS3{F0: (v7 << (((v20 + (^(v20)))) & 63)), F1: (v20 - (v20 * v20)), F2: uint64(v7)}
                    v12 += v12
                    switch v1 {
                    case 32768:
                        var v49 uint16 = v5
                        _ = v49
                    case -27:
                        fmt.Println(v19)
                    default:
                        for fi50 := 0; fi50 < 14; fi50++ {
                            if ((v19 + v19) < "") {
                                for fi51 := 0; fi51 < 6; fi51++ {
                                    var v52 uint8 = v6
                                    v5 = (v5 * v5)
                                    var v53 uint16 = v2
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                                v5 = v7
                            }
                            _ = fi50
                        }
                    }
                    _ = v48
                } else {
                    var v54 int8 = -1
                    var v55 float64 = v9
                    _ = v54
                    _ = v55
                }
                fmt.Println(v19, v12)
            } else {
                v12 = v12
                _ = len(v4)
                var v56 int = v1
                _ = v56
            }
            fmt.Println(v3, v1)
        } else {
            fmt.Println(v19, v6, v11)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    }
    v57 := GenS2{F0: (v6 + (^(v6))), F1: (v19 + "foo bar"), F2: true}
    var v58 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
