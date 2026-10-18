package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint16; F1 string; F2 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int8 }

func main() {
    var v1 int8 = -128
    v2 := []string{("foo bar" + ("\n\t" + ("a" + "hello")))}
    v3 := make(map[uint32]bool)
    v4 := []uint32{}
    v5 := make(map[uint32]uint32)
    var v6 float64 = 0.1
    if len(v4) > 0 {
        v4[0] = 312
    }
    if true {
        v5[971] = 2147483648
        var v7 int32 = 744
        _ = v7
    } else {
        switch -470 {
        case -1:
            v6 -= v6
        case 128:
            switch 484 {
            case 9223372036854775807:
                v8 := GenS1{F0: ((false || false) || ((float32(v6) + 3.14159) > -2.25)), F1: 18446744073709551615}
                _ = v8
            default:
                v6 = v6
            }
        case 222:
            v9 := GenS2{F0: 42, F1: "a", F2: 2147483647}
            _ = v9
        }
        v1 += -1
    }
    switch ("hello" + "") {
    case "foo bar":
        v6 = (v6 + float64(1e10))
    default:
        v6 = 0.0
    }
    switch "\n\t" {
    case "":
        v1 -= v1
    default:
        for fi10 := 0; fi10 < 9; fi10++ {
            if true {
                switch (fi10 | fi10) {
                case 127:
                    v11 := GenS1{F0: (1 <= 18446744073709551615), F1: 42}
                    _ = v11
                case 9223372036854775807:
                    v1 -= v1
                case -2147483648:
                    fmt.Println(fi10, fi10, fi10)
                }
                var v12 int8 = (-(v1))
                if true {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v13 := GenS3{F0: -32768}
                    v1 = 0
                    _ = v13
                }
                _ = v12
            }
            var v14 int64 = 832
            _ = fi10
            _ = v14
        }
    }
    var v15 uint64 = 0
    if (905 < 916) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (true && true) {
            switch "" {
            case "a":
                if (!false) {
                    var v16 string = "\n\t"
                    _ = v16
                }
            case "":
                var v17 int32 = 0
                _ = v17
            case "foo bar":
                switch 9223372036854775807 {
                case 32767:
                    fmt.Println(v6, v1)
                case -739:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                }
            }
            v15 = (v15 / ((v15) | 1))
            switch 128 {
            case 32768:
                fmt.Println(v15, v6)
            case 42:
                var v18 uint32 = 2147483648
                _ = v18
            case -9223372036854775808:
                switch "" {
                case "":
                    v1 -= (v1 ^ -128)
                case "\n\t":
                    var v19 int64 = -9223372036854775808
                    _ = v19
                case "hello":
                    if false {
                        if true {
                            if len(v4) > 0 {
                                v4[0] = 0
                            }
                            _ = len(v2)
                            v20 := GenS2{F0: 615, F1: "foo bar", F2: 466}
                            _ = v20
                        }
                        if ((127 == 1) && true) {
                            fmt.Println(v15, v15)
                        }
                    } else {
                        v21 := GenS4{F0: (v1 >> ((255) & 63)), F1: v1}
                        var v22 bool = ((-32768 >= 991) || (!(!true)))
                        _ = v21
                        _ = v22
                    }
                default:
                    v23 := GenS1{F0: true, F1: 2}
                    _ = v23
                }
            }
        } else {
            v15 += (v15 / ((0) | 1))
        }
    } else {
        v1 = 42
    }
    fmt.Println(v1, v6, v1)
    for fi24 := 0; fi24 < 9; fi24++ {
        if (false && ((((("a" + ((("foo bar" + "a") + "foo bar") + "")) + ("hello" + "")) > "\n\t") && true) || (!(922 == 961)))) {
            var v25 uint = 0
            _ = v25
        }
        if false {
            var v26 uint = 1
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v26
        } else {
            for fi27 := 0; fi27 < 4; fi27++ {
                v15 = (v15 >> ((81) & 63))
                var v28 uint8 = 255
                var v29 uint16 = 652
                var v30 float32 = 1e10
                fmt.Println(v6)
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
        var v31 int16 = -1
        _ = fi24
        _ = v31
    }
    if (!(2147483647 == 752)) {
        fmt.Println(v15)
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
