package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 int8; F1 string; F2 uint32 }
type GenS8 struct { F0 uint; F1 int64; F2 uint16 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{}
    v3 := []int{42, 2147483647}
    var v4 uint8 = 128
    if true {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        var v5 int16 = 643
        v5 += (v5 / ((42) | 1))
        _ = v5
    } else {
        if (1 > 255) {
            switch "hello" {
            case "":
                v6 := GenS1{F0: -56, F1: ("a" + "foo bar")}
                _ = v6
            case "a":
                switch 2147483647 {
                case -361:
                    fmt.Println(v1, v4)
                case -755:
                    var v7 int32 = 42
                    _ = v7
                case 127:
                    v4 += (v4 % (((v4 - (v4 << ((65535) & 63)))) | 1))
                default:
                    if (!(0 != 2147483647)) {
                        v4 += v4
                        switch (("" + ("foo bar" + ("\n\t" + ((("hello" + ("hello" + (((("a" + "hello") + "") + "a") + "foo bar"))) + "foo bar") + ("hello" + "a"))))) + "\n\t") {
                        case "a":
                            if (v4 == 1) {
                                for fi8 := 0; fi8 < 3; fi8++ {
                                    v9 := GenS2{F0: 255, F1: -1}
                                    v3 = append(v3, -639)
                                    _ = fi8
                                    _ = v9
                                }
                                fmt.Println(v1, v1)
                                var v10 int16 = 32767
                                _ = v10
                            } else {
                                for fi11 := 0; fi11 < 2; fi11++ {
                                    v12 := GenS3{F0: 18446744073709551615}
                                    _ = fi11
                                    _ = v12
                                }
                                _ = len(v3)
                                v1 -= v1
                            }
                        }
                        v4 -= v4
                    }
                }
            }
            switch "\n\t" {
            case "hello":
                for fi13 := 0; fi13 < 13; fi13++ {
                    v4 = 255
                    _ = fi13
                }
            case "":
                v14 := GenS4{F0: (("hello" + "foo bar") + "")}
                _ = v14
            default:
                fmt.Println(v4, v4)
            }
        }
        var v15 int8 = 33
        var v16 bool = (!false)
        _ = v15
        _ = v16
    }
    v4 += 2
    if false {
        _ = len(v2)
    }
    var v17 uint64 = 419
    _ = len(v2)
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint16 = 282
        if false {
            v20 := GenS5{F0: 99, F1: (v19 % ((42) | 1))}
            var v21 int8 = 0
            v22 := GenS6{F0: 127, F1: -32768}
            _ = v20
            _ = v21
            _ = v22
        } else {
            var v23 uint64 = 1
            if (v1 > (v1 * (-(v1)))) {
                var v24 uint16 = v19
                for fi25 := 0; fi25 < 14; fi25++ {
                    if false {
                        fmt.Println(fi18, fi25)
                        switch "foo bar" {
                        case "foo bar":
                            if false {
                                var v26 int16 = 0
                                _ = v26
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(v4, v17)
                                v27 := GenS1{F0: -1, F1: "hello"}
                                _ = v27
                            }
                        case "hello":
                            var v28 int8 = 127
                            _ = v28
                        }
                    } else {
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        v1 = v1
                    }
                    v29 := GenS7{F0: -128, F1: "", F2: 859}
                    v30 := GenS8{F0: 18446744073709551615, F1: (v1 & (v1 / ((-483) | 1))), F2: v19}
                    v24 += v24
                    v1 += v1
                    v31 := GenS4{F0: (("\n\t" + ("" + ("\n\t" + "foo bar"))) + "foo bar")}
                    _ = fi25
                    _ = v29
                    _ = v30
                    _ = v31
                }
                switch 1 {
                case 9223372036854775807:
                    if false {
                        if false {
                            v1 -= v1
                        }
                    } else {
                        fmt.Println(v4, v4, v17)
                        fmt.Println(v1)
                    }
                case 32768:
                    v2 = append(v2, "")
                case -172:
                    v32 := GenS3{F0: (v23 | (^(v23)))}
                    _ = v32
                }
                _ = v24
            } else {
                v4 -= (v4 << ((uint(v17)) & 63))
            }
            _ = v23
        }
        var v33 int8 = -1
        _ = fi18
        _ = v19
        _ = v33
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = append(v3, -1)
    var v34 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
