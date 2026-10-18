package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int64; F2 int16 }
type GenS3 struct { F0 uint; F1 int8; F2 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 int; F1 uint16 }

func main() {
    var v1 int8 = -1
    var v2 int64 = 947
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 int8 = v1
    var v6 uint = 2
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := v3[("a" + ("\n\t" + "hello"))]
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        if (false || true) {
            v10 := GenS1{F0: 2147483648}
            v11 := GenS1{F0: 2147483648}
            var v12 float64 = -2.25
            _ = v10
            _ = v11
            _ = v12
        } else {
            var v13 int = fi9
            fmt.Println(fi9, v5)
            _ = v13
        }
        _ = fi9
    }
    switch (("" + "") + "foo bar") {
    case "foo bar":
        var v14 int8 = v1
        _ = v14
    case "":
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 bool = true
            switch "" {
            case "hello":
                switch -9223372036854775808 {
                case 195:
                    v17 := GenS2{F0: v6, F1: -320, F2: 67}
                    _ = v17
                case 423:
                    v1 -= v1
                case 1:
                    fmt.Println(v6, v16, v2)
                default:
                    for fi18 := 0; fi18 < 2; fi18++ {
                        switch fi18 {
                        case 128:
                            v19 := GenS3{F0: (v6 % (((v6 | 0)) | 1)), F1: (v5 % ((v5) | 1)), F2: -549}
                            _ = v19
                        case -686:
                            var v20 uint8 = 2
                            _ = v20
                        case 127:
                            if v16 {
                                v21 := []int{830, fi15, (fi15 ^ (-(fi18)))}
                                _ = v21
                            } else {
                                var v22 int64 = v2
                                v22 = v22
                                _ = v22
                            }
                        }
                        v6 = (v6 + (^(v6)))
                        _ = fi18
                    }
                }
            default:
                v23 := []int{32768, fi15, 2147483647}
                _ = v23
            }
            _ = fi15
            _ = v16
        }
    case "\n\t":
        var v24 uint = v6
        _ = v24
    default:
        if true {
            v25 := v3["\n\t"]
            _ = v25
        } else {
            var v26 int8 = (v1 - v5)
            if true {
                var v27 int8 = v1
                v2 -= 502
                if (false && true) {
                    for fi28 := 0; fi28 < 13; fi28++ {
                        switch "\n\t" {
                        case "a":
                            v1 = (v5 >> ((v6) & 63))
                        case "foo bar":
                            if true {
                                switch "foo bar" {
                                case "hello":
                                    for fi29 := 0; fi29 < 14; fi29++ {
                                        var v30 int = 627
                                        _ = fi29
                                        _ = v30
                                    }
                                default:
                                    v27 = (v1 - (v1 / ((v5) | 1)))
                                }
                                if (!false) {
                                    for fi31 := 0; fi31 < 18; fi31++ {
                                        switch "a" {
                                        case "":
                                            v32 := []int{}
                                            _ = v32
                                        default:
                                            fmt.Println(fi31)
                                        }
                                        var v33 uint = v6
                                        switch ("hello" + ("a" + "")) {
                                        case "foo bar":
                                            if (255 < 1) {
                                                switch "a" {
                                                case "foo bar":
                                                    delete(v3, "foo bar")
                                                }
                                                v34 := v3[(((("" + "foo bar") + "hello") + ((("\n\t" + (("foo bar" + "\n\t") + "\n\t")) + "a") + ((("" + "hello") + ("" + "foo bar")) + ("\n\t" + "hello")))) + (("hello" + "foo bar") + (((("foo bar" + "hello") + "") + "foo bar") + "hello")))]
                                                v35 := GenS4{F0: 65535}
                                                _ = v34
                                                _ = v35
                                            }
                                        case "\n\t":
                                            if false {
                                                fmt.Println(fi31, v33)
                                            } else {
                                                v6 = v33
                                                v1 -= v5
                                                v5 -= 67
                                            }
                                        }
                                        _ = fi31
                                        _ = v33
                                    }
                                    var v36 string = "foo bar"
                                    var v37 float64 = 1e10
                                    _ = v36
                                    _ = v37
                                }
                                if ((56 <= 127) || true) {
                                    var v38 int8 = 42
                                    _ = v38
                                } else {
                                    var v39 float64 = 1.5
                                    var v40 uint8 = 0
                                    _ = v39
                                    _ = v40
                                }
                            } else {
                                for fi41 := 0; fi41 < 19; fi41++ {
                                    fmt.Println(v6, v6)
                                    var v42 float32 = 0.0
                                    v43 := GenS5{F0: (v26 & (v26 ^ v27))}
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                }
                            }
                        case "\n\t":
                            var v44 float32 = -2.25
                            _ = v44
                        default:
                            if (1 <= 2147483647) {
                                var v45 int = (fi28 + (fi28 >> ((v6) & 63)))
                                switch 127 {
                                case -2147483648:
                                    var v46 int8 = v27
                                    _ = v46
                                case 473:
                                    v47 := GenS6{F0: v6, F1: 2147483647}
                                    _ = v47
                                }
                                var v48 uint64 = 1
                                _ = v45
                                _ = v48
                            }
                        }
                        v49 := GenS7{F0: (fi28 - fi28), F1: 65535}
                        _ = fi28
                        _ = v49
                    }
                    var v50 string = "\n\t"
                    fmt.Println(v5)
                    _ = v50
                }
                _ = v27
            }
            if false {
                v26 -= v5
                switch 0 {
                case 128:
                    for fi51 := 0; fi51 < 17; fi51++ {
                        _ = len(v4)
                        _ = fi51
                    }
                default:
                    if (true && true) {
                        for fi52 := 0; fi52 < 3; fi52++ {
                            v4[("" + "hello")] = true
                            _ = fi52
                        }
                    }
                }
            }
            _ = v26
        }
    }
    if false {
        fmt.Println(v5)
        v53 := []int{-503, 32768, -1}
        _ = v53
    }
    var v54 int = -9223372036854775808
    fmt.Println(v54, v54, v1)
    fmt.Println(v6)
    var v55 int8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
