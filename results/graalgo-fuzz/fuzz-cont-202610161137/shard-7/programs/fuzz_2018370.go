package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int64 = 1
    v2 := []string{"a", "", "", "\n\t"}
    var v3 int64 = v1
    var v4 int16 = -355
    v5 := GenS1{F0: 204, F1: -1, F2: 429}
    fmt.Println(v1)
    v4 -= v4
    var v6 int32 = 36
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            v1 = (-(v1))
            _ = fi8
        }
        _ = fi7
    }
    fmt.Println(v6, v3, v6)
    switch -2147483648 {
    case -1:
        var v9 float64 = 3.14159
        _ = v9
    case 127:
        for fi10 := 0; fi10 < 10; fi10++ {
            switch -1 {
            case 284:
                for fi11 := 0; fi11 < 4; fi11++ {
                    fmt.Println(fi11, fi10)
                    _ = fi11
                }
            case -530:
                v12 := GenS1{F0: 255, F1: v6, F2: 18446744073709551615}
                _ = v12
            case 9223372036854775807:
                v2 = append(v2, "foo bar")
            }
            v4 += (v4 | (v4 * -51))
            _ = fi10
        }
    case 42:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v13 int8 = int8(v6)
        _ = v13
    }
    var v14 uint32 = 4294967295
    var v15 int = 0
    v15 = v15
    v4 = 32767
    switch "foo bar" {
    case "hello":
        var v16 int32 = (v6 - 2147483647)
        _ = v16
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    case "":
        v3 = (v3 ^ v1)
    default:
        var v17 bool = false
        _ = v17
    }
    if (1 <= (v6 >> ((188) & 63))) {
        v6 -= 0
    } else {
        var v18 float32 = 3.14159
        var v19 string = "foo bar"
        fmt.Println(v3, v4, v3)
        _ = v18
        _ = v19
    }
    var v20 bool = (!(true || false))
    if (!v20) {
        if (!(!true)) {
            switch v15 {
            case -1:
                for fi21 := 0; fi21 < 6; fi21++ {
                    var v22 bool = false
                    _ = fi21
                    _ = v22
                }
            case 128:
                var v23 int32 = 42
                _ = v23
            default:
                _ = len(v2)
            }
        } else {
            for fi24 := 0; fi24 < 7; fi24++ {
                v25 := GenS1{F0: 0, F1: (v6 >> ((255) & 63)), F2: 4294967295}
                var v26 string = "foo bar"
                v14 += v14
                _ = fi24
                _ = v25
                _ = v26
            }
        }
    }
    if v20 {
        var v27 int8 = 81
        v28 := GenS1{F0: 14, F1: 814, F2: 509}
        if ((690 == (-(v15))) || false) {
            switch ("\n\t" + "") {
            case "hello":
                if ((v3 | (-(v3))) <= 1) {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        for fi30 := 0; fi30 < 3; fi30++ {
                            switch "a" {
                            case "a":
                                switch "\n\t" {
                                case "\n\t":
                                    v31 := GenS2{F0: 1, F1: int8(v4)}
                                    _ = v31
                                default:
                                    v32 := GenS3{F0: 0}
                                    _ = v32
                                }
                            }
                            v1 -= v1
                            if v20 {
                                switch -886 {
                                case 9223372036854775807:
                                    fmt.Println(fi29, v3, v4)
                                case 42:
                                    switch v15 {
                                    case 127:
                                        v33 := GenS4{F0: "a", F1: v27}
                                        _ = v33
                                    case 128:
                                        v1 = v3
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                case -2147483648:
                                    v6 += v6
                                }
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                if false {
                                    v2 = append(v2, "\n\t")
                                    _ = len(v2)
                                    v4 += int16(v6)
                                }
                            } else {
                                v34 := GenS5{F0: 18446744073709551615, F1: (v6 - (v6 / ((-638) | 1)))}
                                var v35 uint32 = (v14 + 2147483648)
                                fmt.Println(fi29)
                                _ = v34
                                _ = v35
                            }
                            _ = fi30
                        }
                        for fi36 := 0; fi36 < 13; fi36++ {
                            v37 := GenS4{F0: "a", F1: (v27 & (^(v27)))}
                            v38 := GenS3{F0: int8(fi36)}
                            if len(v2) > 0 {
                                v2[0] = ("hello" + "")
                            }
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        _ = len(v2)
                        _ = fi29
                    }
                    v2 = append(v2, "\n\t")
                    var v39 float64 = 0.0
                    _ = v39
                } else {
                    var v40 uint8 = 2
                    var v41 int32 = int32(v14)
                    _ = v40
                    _ = v41
                }
            case "\n\t":
                var v42 float32 = 1.5
                _ = v42
            }
        }
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
