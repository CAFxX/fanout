package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS6 struct { F0 int8 }

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 128
    var v3 uint8 = v2
    var v4 int16 = -996
    v5 := []string{("hello" + "a"), (("foo bar" + "\n\t") + ("foo bar" + ("hello" + ((("\n\t" + "\n\t") + "foo bar") + "a")))), ("" + (("hello" + "a") + "foo bar"))}
    switch "" {
    case "hello":
        v3 = (^(v2))
    case "foo bar":
        v2 -= v3
    case "a":
        v2 = (v2 | (^(v2)))
    default:
        for fi6 := 0; fi6 < 4; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                v3 -= (v3 / (((^(v3))) | 1))
                _ = fi7
            }
            fmt.Println(v2)
            _ = fi6
        }
    }
    if (v3 <= v2) {
        if (!false) {
            v8 := GenS1{F0: 1000}
            _ = v8
        }
        v1 = float32(0.1)
    } else {
        v9 := GenS1{F0: 0}
        _ = v9
    }
    if (false && true) {
        fmt.Println(v2)
        v1 -= v1
        for fi10 := 0; fi10 < 13; fi10++ {
            v1 += v1
            _ = fi10
        }
    } else {
        v3 = (v3 - (v3 | v3))
    }
    var v11 uint = 255
    _ = len(v5)
    v12 := GenS2{F0: (v4 + v4)}
    v13 := GenS3{F0: 0, F1: -307}
    if false {
        v1 += 0.1
        v4 -= v4
    } else {
        v11 += v11
        fmt.Println(v3)
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v11 -= v11
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 bool = (!false)
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    switch -9223372036854775808 {
    case -118:
        var v17 float32 = 0.0
        _ = v17
    case -1:
        v11 -= (v11 << (((^(v11))) & 63))
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        if false {
            if false {
                fmt.Println(v11, v4, v11)
                if (v2 > v2) {
                    fmt.Println(v2)
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v19 int32 = 0
                    _ = v19
                }
            } else {
                v11 = v11
                if false {
                    for fi20 := 0; fi20 < 2; fi20++ {
                        switch 422 {
                        case 127:
                            v5 = append(v5, "hello")
                        case 227:
                            if true {
                                switch ("foo bar" + "\n\t") {
                                case "":
                                    _ = len(v5)
                                }
                                fmt.Println(v4)
                                var v21 uint16 = 32767
                                _ = v21
                            } else {
                                if false {
                                    v4 -= v4
                                    switch "foo bar" {
                                    case "foo bar":
                                        v22 := GenS2{F0: (v4 >> (((v11 ^ (v11 | (v11 / (((v11 & v11)) | 1))))) & 63))}
                                        _ = v22
                                    }
                                } else {
                                    _ = len(v5)
                                }
                                fmt.Println(fi20, v3)
                            }
                        default:
                            if (42 <= 1) {
                                if (true || (false && true)) {
                                    if ((true && (!false)) && (-1 == -697)) {
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v23 float64 = 0.0
                                            _ = v23
                                        case "":
                                            if ((v4 / (((v4 >> ((0) & 63))) | 1)) > 32767) {
                                                v1 -= v1
                                                v3 = (v3 - 57)
                                                switch "a" {
                                                case "hello":
                                                    var v24 int64 = 1
                                                    _ = v24
                                                case "a":
                                                    v4 -= v4
                                                case "foo bar":
                                                    v25 := GenS4{F0: 645, F1: "foo bar"}
                                                    _ = v25
                                                default:
                                                    var v26 int = (fi18 & fi18)
                                                    _ = v26
                                                }
                                            } else {
                                                v3 = 1
                                                if len(v5) > 0 {
                                                    v5[0] = "a"
                                                }
                                            }
                                        default:
                                            v1 += (float32(((3.14159 * ((float64((0.1 + v1)) - 1.5) / (3.14159 * (0.1 + (0.1 * -2.25))))) * 3.14159)) - (v1 + float32((((float64(1.5) * 0.0) - 0.0) / 1.0))))
                                        }
                                    } else {
                                        var v27 uint32 = 2147483648
                                        _ = v27
                                    }
                                } else {
                                    if false {
                                        _ = len(v5)
                                    } else {
                                        var v28 string = "foo bar"
                                        v11 -= 4294967295
                                        var v29 string = v28
                                        var v30 uint = v11
                                        if (42 != -128) {
                                            if len(v5) > 0 {
                                                v5[0] = v28
                                            }
                                        }
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                if (fi18 != 2147483647) {
                                    v1 = (v1 * -2.25)
                                } else {
                                    v31 := GenS5{F0: 4294967295, F1: 2, F2: 1}
                                    switch "foo bar" {
                                    case "hello":
                                        var v32 uint8 = v3
                                        _ = v32
                                    case "":
                                        fmt.Println(fi18, v11)
                                    default:
                                        _ = len(v5)
                                    }
                                    _ = v31
                                }
                            }
                            if false {
                                v11 = 255
                            }
                        }
                        v5 = append(v5, "foo bar")
                        _ = fi20
                    }
                } else {
                    v5 = append(v5, "foo bar")
                    v11 = 65535
                }
            }
        } else {
            switch "\n\t" {
            case "a":
                if (255 <= 0) {
                    fmt.Println(v11)
                    fmt.Println(v11, v3, fi18)
                } else {
                    var v33 int = 791
                    var v34 string = ""
                    _ = v33
                    _ = v34
                }
            case "\n\t":
                _ = len(v5)
            }
            switch -542 {
            case 42:
                v35 := GenS2{F0: (^(v4))}
                _ = v35
            }
        }
        _ = fi18
    }
    v36 := GenS3{F0: 0, F1: 224}
    v1 = 0.0
    v37 := GenS6{F0: 109}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
