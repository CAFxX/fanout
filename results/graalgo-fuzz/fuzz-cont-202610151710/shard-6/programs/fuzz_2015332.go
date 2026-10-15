package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint64; F1 uint8; F2 string }

func main() {
    var v1 int = 2147483647
    var v2 float64 = (0.1 * 1.5)
    v3 := []int{v1, (v1 | -2147483648), (v1 % (((v1 >> ((65535) & 63))) | 1)), (v1 / ((128) | 1))}
    var v4 int8 = 42
    var v5 uint16 = 539
    v6 := []uint32{uint32(v5)}
    v1 = (v1 + (v1 & v1))
    v1 = (v1 & 9223372036854775807)
    if (false && (("" < ("" + "")) && (18446744073709551615 != 42))) {
        v7 := GenS1{F0: 0}
        _ = v7
    }
    if (true && true) {
        v2 = v2
        v8 := GenS2{F0: 65535, F1: 1}
        _ = v8
    } else {
        if true {
            var v9 int16 = 0
            v5 += 1
            v2 = v2
            _ = v9
        }
        v6 = append(v6, 489)
        _ = len(v6)
    }
    if true {
        fmt.Println(v5, v1, v2)
    }
    switch 0 {
    case 32767:
        if false {
            v1 -= -9223372036854775808
            var v10 uint64 = 18446744073709551615
            _ = v10
        } else {
            switch ("" + ("hello" + ("hello" + ""))) {
            case "\n\t":
                _ = len(v6)
            case "a":
                var v11 int = v1
                _ = v11
            default:
                for fi12 := 0; fi12 < 15; fi12++ {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v14 := GenS2{F0: v5, F1: 332}
                        if false {
                            var v15 int32 = -1
                            v2 = v2
                            fmt.Println(v2)
                            _ = v15
                        } else {
                            v16 := GenS3{F0: 0, F1: (1e10 == (v2 + 0.1))}
                            if false {
                                v5 = v5
                                if true {
                                    v5 += (^(v5))
                                    if true {
                                        v5 += v5
                                        v2 -= ((1e10 / (float64(1.5) + -2.25)) * v2)
                                    }
                                    var v17 int8 = (v4 / ((v4) | 1))
                                    var v18 float64 = (v2 - float64(1.5))
                                    switch 4294967295 {
                                    case 934:
                                        v5 = (v5 << ((255) & 63))
                                    case 32767:
                                        if true {
                                            fmt.Println(fi12, fi13, v4)
                                            v19 := GenS4{F0: "hello", F1: 0}
                                            _ = v19
                                        }
                                    case 127:
                                        v17 += (v17 << ((1) & 63))
                                    default:
                                        v20 := GenS3{F0: 188, F1: false}
                                        _ = v20
                                    }
                                    if true {
                                        if len(v3) > 0 {
                                            v3[0] = fi12
                                        }
                                        fmt.Println(v2, v17, v5)
                                    }
                                    _ = v17
                                    _ = v18
                                } else {
                                    if (42 == 255) {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v21 bool = (!((float32(v2) * 0.0) == 0.0))
                                            _ = v21
                                        default:
                                            v22 := GenS3{F0: 128, F1: ("a" != "")}
                                            _ = v22
                                        }
                                        var v23 int32 = 908
                                        v6 = append(v6, 1)
                                        _ = v23
                                    }
                                    fmt.Println(v5, v5, v4)
                                }
                            }
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                    }
                    for fi24 := 0; fi24 < 2; fi24++ {
                        v6 = append(v6, 221)
                        v1 -= fi12
                        if false {
                            _ = len(v3)
                            var v25 float64 = float64((-2.25 - (1.5 / 1.0)))
                            _ = v25
                        } else {
                            if false {
                                var v26 uint32 = 2147483648
                                v6 = append(v6, v26)
                                _ = v26
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi24
                    }
                    _ = fi12
                }
            }
            var v27 float32 = float32((v2 + (1.5 * float64(1e10))))
            _ = v27
        }
    case 2147483647:
        switch ("" + "\n\t") {
        case "foo bar":
            var v28 string = "foo bar"
            _ = v28
        case "\n\t":
            fmt.Println(v5, v4, v5)
        case "a":
            var v29 uint8 = uint8(v5)
            _ = v29
        }
    default:
        fmt.Println(v4)
    }
    if len(v6) > 0 {
        v6[0] = 1
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v30 bool = (v5 < 2)
    v30 = true
    v31 := GenS5{F0: 0, F1: 144, F2: ""}
    for fi32 := 0; fi32 < 3; fi32++ {
        _ = len(v6)
        for fi33 := 0; fi33 < 16; fi33++ {
            v5 = v5
            _ = fi33
        }
        if (!(!(v30 || v30))) {
            v34 := GenS1{F0: -2147483648}
            for fi35 := 0; fi35 < 1; fi35++ {
                var v36 int64 = 42
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi35
                _ = v36
            }
            _ = v34
        } else {
            if v30 {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                switch (("\n\t" + "hello") + "foo bar") {
                case "hello":
                    v3 = append(v3, v1)
                default:
                    fmt.Println(v2)
                }
                var v37 int64 = 1
                _ = v37
            }
            switch v1 {
            case 9223372036854775807:
                v38 := GenS5{F0: 0, F1: 255, F2: "foo bar"}
                _ = v38
            case 127:
                switch (("hello" + "hello") + (("" + ("a" + ("" + (("hello" + "hello") + ((((("foo bar" + "\n\t") + "\n\t") + "hello") + "\n\t") + ("" + "a")))))) + "a")) {
                case "hello":
                    v39 := GenS2{F0: 0, F1: 999}
                    _ = v39
                case "foo bar":
                    if v30 {
                        var v40 int16 = -890
                        v41 := GenS1{F0: 42}
                        for fi42 := 0; fi42 < 2; fi42++ {
                            switch "a" {
                            case "a":
                                var v43 uint8 = 0
                                _ = v43
                            case "":
                                v4 += (v4 / ((v4) | 1))
                            case "hello":
                                if ("hello" == "foo bar") {
                                    v5 -= (v5 * (^(v5)))
                                    var v44 int64 = 9223372036854775807
                                    _ = v44
                                } else {
                                    var v45 int64 = -280
                                    v1 = (^(fi42))
                                    _ = v45
                                }
                            }
                            _ = fi42
                        }
                        _ = v40
                        _ = v41
                    }
                case "":
                    for fi46 := 0; fi46 < 11; fi46++ {
                        if (!false) {
                            switch fi32 {
                            case -1:
                                v5 += v5
                            }
                        }
                        _ = fi46
                    }
                }
            default:
                if v30 {
                    v2 = v2
                } else {
                    for fi47 := 0; fi47 < 12; fi47++ {
                        v2 = (v2 / (v2 - 1e10))
                        _ = fi47
                    }
                    var v48 bool = v30
                    _ = v48
                }
            }
            fmt.Println(v2, v5, fi32)
        }
        _ = fi32
    }
    v2 = v2
    var v49 int16 = 531
    var v50 float64 = 1e10
    for fi51 := 0; fi51 < 16; fi51++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v30 = (v49 == v49)
        _ = fi51
    }
    if (1e10 == 0.1) {
        v52 := GenS5{F0: 0, F1: 128, F2: "foo bar"}
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
