package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16 }

func main() {
    var v1 float64 = 1.5
    var v2 int8 = 100
    var v3 float32 = 3.14159
    v4 := []int{-564, 9223372036854775807, 128, 127}
    var v5 uint32 = 2
    if len(v4) > 0 {
        v4[0] = -9223372036854775808
    }
    if true {
        v6 := GenS1{F0: 42, F1: 18446744073709551615}
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS2{F0: "", F1: uint16(v5)}
                _ = fi8
                _ = v9
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch "a" {
            case "":
                fmt.Println(fi7)
            default:
                v4 = append(v4, 1)
            }
            _ = fi7
        }
        if (-1 <= -115) {
            switch 0 {
            case 32767:
                v1 = v1
            }
            var v10 int8 = int8(v5)
            fmt.Println(v3, v10)
            _ = v10
        } else {
            v11 := GenS2{F0: ("a" + "a"), F1: 562}
            fmt.Println(v3)
            _ = v11
        }
        _ = v6
    } else {
        var v12 uint64 = 282
        if true {
            v13 := GenS2{F0: "", F1: uint16(v5)}
            switch 128 {
            case -9223372036854775808:
                var v14 uint64 = 707
                _ = v14
            }
            switch 9223372036854775807 {
            case -2147483648:
                switch "a" {
                case "foo bar":
                    fmt.Println(v12)
                }
            case 362:
                for fi15 := 0; fi15 < 5; fi15++ {
                    v3 = 3.14159
                    _ = fi15
                }
            default:
                for fi16 := 0; fi16 < 1; fi16++ {
                    v12 -= (v12 >> ((762) & 63))
                    _ = fi16
                }
            }
            _ = v13
        } else {
            fmt.Println(v12, v12, v1)
            _ = len(v4)
        }
        _ = v12
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v17 := GenS3{F0: "a"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v18 int = 32768
    switch 612 {
    case -2147483648:
        if true {
            v19 := GenS2{F0: ((((("a" + "foo bar") + "hello") + ((("a" + ("a" + "foo bar")) + ("" + "\n\t")) + "")) + "hello") + "hello"), F1: 65535}
            fmt.Println(v2, v5, v3)
            _ = v19
        } else {
            v20 := GenS3{F0: "hello"}
            fmt.Println(v2, v3)
            _ = v20
        }
    case 42:
        if true {
            v21 := GenS4{F0: "a", F1: 128}
            if (((v18 < v18) || (1 == v18)) && true) {
                var v22 int64 = 9223372036854775807
                _ = v22
            } else {
                for fi23 := 0; fi23 < 13; fi23++ {
                    v2 = (v2 - (v2 / (((v2 >> ((42) & 63))) | 1)))
                    _ = fi23
                }
                if false {
                    v24 := GenS3{F0: ("foo bar" + "")}
                    _ = v24
                }
                if true {
                    v1 -= v1
                    if false {
                        if (!true) {
                            for fi25 := 0; fi25 < 7; fi25++ {
                                if (-2.25 <= -2.25) {
                                    v4 = append(v4, 32768)
                                    switch ("a" + "") {
                                    case "hello":
                                        v26 := GenS5{F0: 128}
                                        _ = v26
                                    case "\n\t":
                                        var v27 float32 = (v3 * (v3 + ((v3 * v3) * v3)))
                                        _ = v27
                                    case "foo bar":
                                        if true {
                                            var v28 int = v18
                                            _ = v28
                                        }
                                    }
                                }
                                var v29 uint16 = 0
                                _ = fi25
                                _ = v29
                            }
                        } else {
                            for fi30 := 0; fi30 < 6; fi30++ {
                                v31 := GenS6{F0: (v5 % (((v5 - 697)) | 1))}
                                if (800 != 32767) {
                                    _ = len(v4)
                                    fmt.Println(v3, v18, v3)
                                }
                                _ = fi30
                                _ = v31
                            }
                        }
                    } else {
                        v32 := GenS7{F0: 255}
                        v5 = 0
                        _ = v32
                    }
                } else {
                    _ = len(v4)
                }
            }
            _ = v21
        } else {
            var v33 float32 = 0.0
            if false {
                for fi34 := 0; fi34 < 20; fi34++ {
                    for fi35 := 0; fi35 < 6; fi35++ {
                        v5 -= v5
                        v4 = append(v4, fi34)
                        _ = fi35
                    }
                    _ = fi34
                }
            } else {
                var v36 int8 = (v2 * v2)
                var v37 string = ("hello" + "a")
                _ = v36
                _ = v37
            }
            var v38 uint = 2
            _ = v33
            _ = v38
        }
    case -884:
        var v39 int = v18
        _ = v39
    }
    switch v18 {
    case -894:
        v5 -= 2147483648
    case 42:
        var v40 int64 = 0
        _ = v40
    }
    var v41 uint16 = 32767
    var v42 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
