package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int; F1 uint64; F2 string }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 uint8; F1 string }

func main() {
    var v1 int32 = 3
    var v2 int64 = 453
    v3 := []int{128}
    switch 127 {
    case 9223372036854775807:
        v4 := GenS1{F0: 190, F1: 42, F2: -548}
        _ = v4
    case 0:
        var v5 int32 = v1
        _ = v5
    }
    v6 := GenS2{F0: "hello", F1: (^(v1))}
    v2 = (v2 * -9223372036854775808)
    fmt.Println(v1)
    v1 -= v1
    if (!false) {
        if (185 < 128) {
            var v7 float32 = -2.25
            _ = v7
        }
        v2 = v2
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: 128, F1: 255, F2: -509}
        _ = fi8
        _ = v9
    }
    var v10 int = 42
    var v11 int64 = (v2 / ((v2) | 1))
    switch v10 {
    case -36:
        switch "foo bar" {
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    default:
        v10 += 128
    }
    var v12 uint8 = 255
    var v13 int64 = (-(v11))
    switch "" {
    case "a":
        switch "a" {
        case "":
            var v14 float64 = -2.25
            _ = v14
        default:
            var v15 int16 = -1
            _ = v15
        }
    case "hello":
        if false {
            var v16 int16 = -877
            if false {
                v17 := GenS3{F0: 4294967295, F1: 0, F2: (("" + "\n\t") + "foo bar")}
                var v18 int8 = 127
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v17
                _ = v18
            }
            var v19 float32 = float32(1.5)
            _ = v16
            _ = v19
        } else {
            if ((!(2 <= 18446744073709551615)) && (false || (!true))) {
                v1 = (^(v1))
                var v20 float64 = 0.1
                _ = v20
            }
            for fi21 := 0; fi21 < 7; fi21++ {
                v2 -= v13
                fmt.Println(v13)
                _ = fi21
            }
        }
    case "foo bar":
        if true {
            var v22 int16 = 494
            v23 := GenS3{F0: (^(v10)), F1: 940, F2: "hello"}
            _ = v22
            _ = v23
        } else {
            if (v1 <= v1) {
                if true {
                    v24 := GenS4{F0: 2, F1: v11}
                    _ = v24
                } else {
                    fmt.Println(v2)
                    for fi25 := 0; fi25 < 7; fi25++ {
                        _ = len(v3)
                        _ = fi25
                    }
                    if true {
                        switch ("\n\t" + ("\n\t" + "")) {
                        case "":
                            v12 += v12
                        case "hello":
                            switch "a" {
                            case "\n\t":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "hello":
                                var v26 uint8 = v12
                                _ = v26
                            default:
                                if len(v3) > 0 {
                                    v3[0] = (v10 * v10)
                                }
                            }
                        case "\n\t":
                            v2 -= -1
                        default:
                            v27 := GenS5{F0: 1, F1: (!((!false) && (!(!true))))}
                            _ = v27
                        }
                    } else {
                        v28 := GenS1{F0: 127, F1: 0, F2: -238}
                        _ = v28
                    }
                }
                var v29 uint = 908
                fmt.Println(v2, v2, v1)
                _ = v29
            }
            if (-32768 <= 32767) {
                v2 = v13
                switch (v10 ^ -2147483648) {
                case 2147483647:
                    v2 -= (v2 >> ((4294967295) & 63))
                case -508:
                    v30 := GenS3{F0: v10, F1: 949, F2: "\n\t"}
                    _ = v30
                case -1:
                    v31 := GenS4{F0: 137, F1: (v11 / ((v11) | 1))}
                    _ = v31
                default:
                    for fi32 := 0; fi32 < 5; fi32++ {
                        var v33 int16 = 598
                        switch (fi32 & 32767) {
                        case -636:
                            v34 := GenS5{F0: 42, F1: (((!true) && (!true)) && ((^(v13)) == (v11 >> ((4294967295) & 63))))}
                            _ = v34
                        default:
                            var v35 uint8 = 63
                            _ = v35
                        }
                        v33 += v33
                        _ = fi32
                        _ = v33
                    }
                }
                var v36 uint64 = 4294967295
                _ = v36
            } else {
                v1 = v1
                fmt.Println(v13)
            }
            v1 = v1
        }
    default:
        v13 = -708
    }
    v37 := GenS6{F0: v12, F1: ""}
    v12 += (^(v12))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v6
}
