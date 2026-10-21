package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 int }
type GenS3 struct { F0 int32; F1 string; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 bool; F2 int64 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    var v4 float64 = 3.14159
    var v5 bool = (false || false)
    var v6 int16 = 42
    fmt.Println(v4)
    v1 = append(v1, -515)
    var v7 float32 = (((3.14159 - (1e10 / 1.0)) + float32(1e10)) / (((0.1 - ((-2.25 / 0.1) - (0.1 + 1e10))) - 1.5) / 1e10))
    switch "a" {
    case "hello":
        v7 = v7
    case "a":
        switch -2147483648 {
        case 708:
            v8 := v2["foo bar"]
            _ = v8
        }
    }
    if v5 {
        if (v5 && (!true)) {
            _ = len(v3)
            if ((0 <= 255) || false) {
                _ = len(v2)
            }
            var v9 uint8 = 0
            _ = v9
        }
        if (false && ((v6 % (((v6 << ((65535) & 63))) | 1)) <= -1)) {
            v10 := GenS1{F0: v6}
            _ = v10
        } else {
            v7 += (v7 - v7)
            v11 := GenS2{F0: 1, F1: 100, F2: -809}
            _ = v11
        }
        var v12 string = ("hello" + (("hello" + "a") + ("\n\t" + ("" + "\n\t"))))
        _ = v12
    }
    if true {
        switch 127 {
        case 32767:
            if v5 {
                v13 := GenS2{F0: 42, F1: 249, F2: 2147483647}
                _ = len(v3)
                switch "hello" {
                case "hello":
                    var v14 int64 = 9223372036854775807
                    _ = v14
                case "\n\t":
                    var v15 uint16 = 42
                    _ = v15
                }
                _ = v13
            }
        case -543:
            v1 = append(v1, 1)
        case 1:
            v16 := v2["a"]
            _ = v16
        default:
            fmt.Println(v4, v6, v5)
        }
        _ = len(v1)
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v17 uint = 0
        v17 -= 65535
        _ = v17
    }
    if v5 {
        for fi18 := 0; fi18 < 10; fi18++ {
            var v19 uint8 = 208
            v4 += 0.1
            v5 = (!v5)
            _ = fi18
            _ = v19
        }
    }
    v20 := GenS3{F0: 991, F1: "hello", F2: 255}
    v3[536] = "a"
    fmt.Println(v5)
    v7 -= 1.5
    var v21 string = "foo bar"
    if v5 {
        if v5 {
            switch v21 {
            case "foo bar":
                fmt.Println(v5, v4, v6)
            case "hello":
                v5 = v5
            case "\n\t":
                v3[4294967295] = (v21 + (("" + ("\n\t" + (((("foo bar" + "") + ((v21 + "\n\t") + "foo bar")) + v21) + v21))) + v21))
            default:
                v22 := GenS4{F0: 2147483647}
                _ = v22
            }
            if false {
                switch 32767 {
                case 32768:
                    if v5 {
                        var v23 int32 = 2147483647
                        _ = v23
                    } else {
                        v4 = float64((v7 - float32(float64((v7 + v7)))))
                    }
                case 127:
                    var v24 bool = (v21 != v21)
                    _ = v24
                case 32767:
                    if v5 {
                        v25 := GenS5{F0: 2147483648, F1: v5, F2: 42}
                        for fi26 := 0; fi26 < 9; fi26++ {
                            var v27 int32 = 1
                            v4 = v4
                            _ = fi26
                            _ = v27
                        }
                        _ = v25
                    } else {
                        fmt.Println(v21, v7)
                    }
                }
                v28 := v3[308]
                _ = v28
            } else {
                if (!(v5 || v5)) {
                    switch (v21 + v21) {
                    case "foo bar":
                        v5 = ((v21 + (v21 + (v21 + v21))) <= "hello")
                    case "hello":
                        v29 := GenS6{F0: 255, F1: 9223372036854775807}
                        _ = v29
                    default:
                        var v30 uint32 = 65535
                        _ = v30
                    }
                    if v5 {
                        switch 931 {
                        case -1:
                            if (v6 != (v6 << ((296) & 63))) {
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    var v32 uint64 = 277
                                    for fi33 := 0; fi33 < 18; fi33++ {
                                        if (127 < 1) {
                                            v7 = (((3.14159 + (1.5 + 3.14159)) + v7) / float32((v4 + (((v4 / float64(v7)) * float64(v7)) + (v4 + v4)))))
                                            var v34 int64 = -9223372036854775808
                                            _ = v34
                                        } else {
                                            v5 = v5
                                        }
                                        var v35 float64 = float64(v7)
                                        _ = fi33
                                        _ = v35
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                            }
                        default:
                            for fi36 := 0; fi36 < 6; fi36++ {
                                v37 := GenS7{F0: 354}
                                if (!true) {
                                    var v38 uint32 = 2147483647
                                    _ = v38
                                }
                                _ = fi36
                                _ = v37
                            }
                        }
                        var v39 int = 9223372036854775807
                        _ = v39
                    }
                }
                switch 0 {
                case -9223372036854775808:
                    var v40 float64 = float64(v7)
                    _ = v40
                case 32767:
                    for fi41 := 0; fi41 < 7; fi41++ {
                        if (!(705 > (-(fi41)))) {
                            var v42 uint8 = 255
                            var v43 uint = 992
                            var v44 int64 = -1
                            _ = v42
                            _ = v43
                            _ = v44
                        } else {
                            var v45 int = 9223372036854775807
                            _ = v45
                        }
                        var v46 uint32 = 0
                        var v47 string = "a"
                        _ = fi41
                        _ = v46
                        _ = v47
                    }
                }
            }
        }
        fmt.Println(v4, v6)
        var v48 uint64 = 42
        _ = v48
    } else {
        switch v21 {
        case "hello":
            for fi49 := 0; fi49 < 7; fi49++ {
                for fi50 := 0; fi50 < 12; fi50++ {
                    fmt.Println(fi50)
                    _ = fi50
                }
                _ = fi49
            }
        default:
            for fi51 := 0; fi51 < 6; fi51++ {
                switch ("\n\t" + v21) {
                case "":
                    v4 += float64((v7 / float32(v4)))
                case "foo bar":
                    v4 += (1.5 + ((((v4 * (v4 / 3.14159)) * (1.5 - -2.25)) / v4) - -2.25))
                case "a":
                    fmt.Println(v5)
                default:
                    fmt.Println(fi51, v5, fi51)
                }
                _ = fi51
            }
        }
        fmt.Println(v6, v7, v7)
    }
    v6 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
