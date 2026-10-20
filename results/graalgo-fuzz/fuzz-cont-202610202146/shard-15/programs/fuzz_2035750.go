package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{(("hello" + ("\n\t" + (("" + (("hello" + "foo bar") + "hello")) + "foo bar"))) + "hello"), "", "foo bar", ("hello" + "")}
    var v3 int32 = -1
    v3 += (^(v3))
    var v4 int16 = 32767
    var v5 int64 = 42
    if (false && (689 > 2)) {
        v5 = v5
        delete(v1, 428)
    }
    var v6 uint = 42
    if false {
        for fi7 := 0; fi7 < 5; fi7++ {
            if true {
                switch "hello" {
                case "foo bar":
                    if false {
                        var v8 string = "hello"
                        fmt.Println(v8, v4, v3)
                        if (fi7 >= 42) {
                            switch "hello" {
                            case "a":
                                if (v5 < 502) {
                                    var v9 int8 = 82
                                    _ = v9
                                }
                            case "":
                                if ((2 == 233) || (v6 == v6)) {
                                    var v10 uint32 = 2147483648
                                    v5 = (v5 >> ((v6) & 63))
                                    v11 := GenS1{F0: 687}
                                    _ = v10
                                    _ = v11
                                }
                            }
                            switch v8 {
                            case "a":
                                v12 := GenS1{F0: 2}
                                _ = v12
                            case "foo bar":
                                for fi13 := 0; fi13 < 8; fi13++ {
                                    v5 += v5
                                    v5 = (v5 / ((-372) | 1))
                                    _ = len(v1)
                                    _ = fi13
                                }
                            case "\n\t":
                                switch ((v8 + v8) + v8) {
                                case "a":
                                    var v14 string = ("a" + "foo bar")
                                    _ = v14
                                case "foo bar":
                                    var v15 int16 = (v4 % (((v4 << (((v6 ^ 18446744073709551615)) & 63))) | 1))
                                    _ = v15
                                case "hello":
                                    if ((((v6 & v6) > v6) || false) || ((fi7 % ((fi7) | 1)) >= -2147483648)) {
                                        for fi16 := 0; fi16 < 5; fi16++ {
                                            switch (fi7 | (^(fi16))) {
                                            case 32768:
                                                delete(v1, 0)
                                            case -89:
                                                var v17 string = ((((v8 + "a") + ((v8 + v8) + ("" + v8))) + (v8 + (v8 + v8))) + v8)
                                                _ = v17
                                            }
                                            switch v8 {
                                            case "hello":
                                                if true {
                                                    v5 += (v5 & (v5 + v5))
                                                    v18 := GenS1{F0: 878}
                                                    _ = v18
                                                }
                                            case "\n\t":
                                                if false {
                                                    switch (fi16 ^ (fi7 | 99)) {
                                                    case 9223372036854775807:
                                                        v1[2147483647] = true
                                                    default:
                                                        v6 = v6
                                                    }
                                                    v19 := GenS2{F0: v8, F1: 0}
                                                    _ = v19
                                                } else {
                                                    if false {
                                                        v8 = (("hello" + v8) + (v8 + ("hello" + v8)))
                                                        var v20 int32 = (v3 ^ (v3 ^ v3))
                                                        _ = v20
                                                    }
                                                    var v21 int = -822
                                                    _ = v21
                                                }
                                            case "":
                                                fmt.Println(v4)
                                            default:
                                                v3 = int32(v4)
                                            }
                                            _ = len(v2)
                                            switch 4294967295 {
                                            case 0:
                                                v5 = v5
                                            case 276:
                                                fmt.Println(v3, v6, v3)
                                            case 4294967295:
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            default:
                                                v1[65535] = false
                                            }
                                            _ = fi16
                                        }
                                    } else {
                                        var v22 string = v8
                                        v23 := GenS3{F0: int32(v4), F1: 103, F2: (v6 & 989)}
                                        fmt.Println(v6, v3)
                                        _ = v22
                                        _ = v23
                                    }
                                }
                            }
                            v4 = -1
                        } else {
                            fmt.Println(v5)
                        }
                        _ = v8
                    }
                case "hello":
                    if false {
                        var v24 uint64 = 624
                        var v25 int32 = 0
                        for fi26 := 0; fi26 < 9; fi26++ {
                            fmt.Println(v5)
                            v1[42] = true
                            var v27 float64 = -2.25
                            _ = fi26
                            _ = v27
                        }
                        _ = v24
                        _ = v25
                    }
                case "\n\t":
                    var v28 uint64 = 240
                    _ = v28
                default:
                    var v29 uint = 18446744073709551615
                    _ = v29
                }
                if (!true) {
                    switch (fi7 << ((42) & 63)) {
                    case -967:
                        v30 := GenS4{F0: 119}
                        _ = v30
                    case -2147483648:
                        v31 := GenS5{F0: 0, F1: 487, F2: 147}
                        _ = v31
                    case 2147483647:
                        var v32 uint8 = 255
                        _ = v32
                    default:
                        v33 := GenS6{F0: v6, F1: -2147483648}
                        _ = v33
                    }
                    v4 = 1000
                }
                fmt.Println(v4, v3)
            }
            v34 := GenS6{F0: (v6 ^ v6), F1: 42}
            if ((v5 / (((v5 | (v5 ^ v5))) | 1)) <= (^(v5))) {
                fmt.Println(fi7)
            } else {
                switch "\n\t" {
                case "a":
                    if (v4 <= v4) {
                        fmt.Println(fi7, v5, v3)
                    }
                case "\n\t":
                    v35 := GenS5{F0: 0, F1: 765, F2: 248}
                    _ = v35
                case "foo bar":
                    if (0 >= 0) {
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        var v36 float32 = -2.25
                        _ = v36
                    }
                default:
                    for fi37 := 0; fi37 < 17; fi37++ {
                        v38 := v1[0]
                        _ = fi37
                        _ = v38
                    }
                }
                var v39 uint32 = 0
                _ = v39
            }
            _ = fi7
            _ = v34
        }
        var v40 int64 = -591
        _ = v40
    } else {
        if (!((true && (859 < 42)) || (244 <= 0))) {
            v41 := GenS7{F0: v5, F1: "hello"}
            v1[42] = (0 != 42)
            if false {
                var v42 float32 = 3.14159
                fmt.Println(v42)
                _ = v42
            }
            _ = v41
        } else {
            v43 := GenS5{F0: 42, F1: 42, F2: 255}
            v44 := GenS4{F0: 32767}
            _ = v43
            _ = v44
        }
        fmt.Println(v3, v6)
    }
    var v45 float32 = 0.1
    switch "hello" {
    case "hello":
        v5 -= 1
    }
    v6 = v6
    for fi46 := 0; fi46 < 16; fi46++ {
        v45 += v45
        _ = fi46
    }
    delete(v1, 4294967295)
    fmt.Println(v5)
    for fi47 := 0; fi47 < 18; fi47++ {
        var v48 int64 = (v5 + -166)
        for fi49 := 0; fi49 < 7; fi49++ {
            var v50 uint32 = 42
            v51 := GenS8{F0: (v6 >> ((255) & 63)), F1: (v50 / ((v50) | 1))}
            _ = fi49
            _ = v50
            _ = v51
        }
        v52 := GenS3{F0: (v3 | (v3 << ((1) & 63))), F1: 255, F2: v6}
        _ = fi47
        _ = v48
        _ = v52
    }
    var v53 uint = (v6 * (^(v6)))
    v45 -= v45
    switch ("" + "a") {
    case "a":
        delete(v1, 2147483648)
    }
    if ((-128 == -118) || (2 < 51)) {
        v1[2147483648] = true
        fmt.Println(v6, v45)
        if len(v2) > 0 {
            v2[0] = ("a" + "foo bar")
        }
    } else {
        v53 += (v53 << ((v53) & 63))
        v54 := v1[841]
        var v55 int64 = (v5 * (v5 * (-(v5))))
        _ = v54
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v53
    _ = v6
}
