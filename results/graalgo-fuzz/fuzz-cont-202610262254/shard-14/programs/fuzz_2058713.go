package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint16 = 255
    v2 := []uint32{626, 929}
    v3 := make(map[uint32]bool)
    var v4 int16 = -722
    var v5 int = 4294967295
    var v6 int16 = 0
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int64 = 148
        _ = fi7
        _ = v8
    }
    v3[369] = (-1 < 42)
    fmt.Println(v5)
    fmt.Println(v4, v4)
    v9 := GenS1{F0: (v5 >> ((65535) & 63)), F1: 42, F2: v1}
    delete(v3, 757)
    v6 += (^(v4))
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            if (!true) {
                switch "foo bar" {
                case "a":
                    var v12 uint = 46
                    _ = v12
                }
                if (!false) {
                    if ("foo bar" == (("foo bar" + "\n\t") + "\n\t")) {
                        fmt.Println(v4, fi10)
                        var v13 int = fi10
                        switch (v5 / ((fi11) | 1)) {
                        case 1:
                            fmt.Println(v1, v6)
                        case 32768:
                            if (false && true) {
                                v14 := GenS1{F0: (v13 | -1), F1: 2, F2: (^(v1))}
                                _ = v14
                            } else {
                                var v15 bool = (1.5 <= 1e10)
                                fmt.Println(v5, fi10, v13)
                                v15 = false
                                delete(v3, 486)
                                v16 := GenS1{F0: (v13 % (((fi10 - v13)) | 1)), F1: 2, F2: (v1 % ((v1) | 1))}
                                _ = v15
                                _ = v16
                            }
                        case 2147483647:
                            if (v4 <= -975) {
                                var v17 uint = 4294967295
                                _ = v17
                            } else {
                                v3[42] = false
                                var v18 uint16 = v1
                                _ = v18
                            }
                        }
                        _ = v13
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = 0
                    }
                    fmt.Println(v1)
                    switch 4294967295 {
                    case 127:
                        if true {
                            v19 := GenS1{F0: (fi11 & 0), F1: 108, F2: v1}
                            v5 -= -250
                            var v20 string = ("foo bar" + "hello")
                            _ = v19
                            _ = v20
                        } else {
                            v2 = append(v2, 2)
                        }
                    case 669:
                        fmt.Println(v5, v1)
                    case 32768:
                        switch "\n\t" {
                        case "foo bar":
                            var v21 float64 = 1.5
                            _ = v21
                        case "a":
                            switch 583 {
                            case 0:
                                var v22 string = "foo bar"
                                _ = v22
                            }
                        case "\n\t":
                            v1 = v1
                        default:
                            fmt.Println(v6, fi10, v6)
                        }
                    default:
                        var v23 int = (fi10 >> ((255) & 63))
                        _ = v23
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                if (!false) {
                    if false {
                        switch (^(fi11)) {
                        case 42:
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                        case -842:
                            fmt.Println(v6)
                        default:
                            fmt.Println(v1)
                        }
                        var v24 uint8 = 179
                        _ = v24
                    } else {
                        if false {
                            fmt.Println(v4, v5, v6)
                        } else {
                            fmt.Println(v1)
                        }
                        if true {
                            v4 += -1
                            if (0 <= 42) {
                                var v25 float64 = (float64(float32(1.5)) + 3.14159)
                                v26 := GenS2{F0: 84, F1: false}
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS3{F0: (^(v5)), F1: 0}
                                if (2147483648 == 2) {
                                    fmt.Println(v1, v5)
                                    fmt.Println(v4)
                                    v5 = 42
                                }
                                _ = v27
                            }
                            var v28 uint32 = 634
                            _ = v28
                        } else {
                            var v29 int32 = 746
                            _ = v29
                        }
                    }
                    v6 = v4
                    v1 += (v1 & v1)
                }
            }
            _ = fi11
        }
        v30 := GenS4{F0: 0, F1: 2147483647}
        _ = fi10
        _ = v30
    }
    var v31 uint64 = 808
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 += (v4 >> ((2) & 63))
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    if true {
        v4 += 0
    }
    v1 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
