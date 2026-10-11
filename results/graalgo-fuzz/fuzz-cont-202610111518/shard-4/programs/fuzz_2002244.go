package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 bool }

func main() {
    var v1 int32 = 1
    var v2 float64 = 0.1
    v3 := []string{"", ("\n\t" + ("foo bar" + (("hello" + "") + "hello"))), "foo bar"}
    var v4 string = ("foo bar" + "a")
    v5 := []int{-1}
    var v6 float64 = v2
    v7 := GenS1{F0: 1, F1: -9223372036854775808, F2: 65535}
    v2 = 1e10
    v2 = ((1e10 / 1.0) - v2)
    if false {
        fmt.Println(v2)
    } else {
        v8 := GenS1{F0: 1, F1: 9223372036854775807, F2: 730}
        _ = v8
    }
    v3 = append(v3, "hello")
    var v9 string = ""
    v2 -= float64(3.14159)
    if (1 == 1) {
        var v10 uint32 = 2
        switch -720 {
        case 2147483647:
            switch -771 {
            case 1:
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 int16 = 32767
                    switch ("hello" + v9) {
                    case "a":
                        v1 = -2147483648
                    case "":
                        v13 := GenS1{F0: 127, F1: 1, F2: 1}
                        _ = v13
                    }
                    _ = fi11
                    _ = v12
                }
            case 4294967295:
                var v14 uint64 = uint64(v10)
                _ = v14
            default:
                switch ("foo bar" + "hello") {
                case "\n\t":
                    v15 := GenS2{F0: 771, F1: 1, F2: (0.0 > (0.1 / v2))}
                    _ = v15
                default:
                    switch 32767 {
                    case -855:
                        var v16 int = 555
                        _ = v16
                    case -478:
                        if true {
                            v17 := GenS2{F0: 392, F1: 52, F2: (("a" + v4) > v4)}
                            if ((v1 - v1) >= 20) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v1 -= v1
                            } else {
                                if (((!false) && (!(-32768 > 42))) || false) {
                                    fmt.Println(v6, v4, v9)
                                    switch ((v4 + "a") + v9) {
                                    case "":
                                        switch ((v9 + v4) + "\n\t") {
                                        case "\n\t":
                                            var v18 float32 = -2.25
                                            _ = v18
                                        default:
                                            var v19 uint32 = (v10 | v10)
                                            _ = v19
                                        }
                                    case "a":
                                        if len(v3) > 0 {
                                            v3[0] = ""
                                        }
                                    case "\n\t":
                                        if (true || (v10 > v10)) {
                                            var v20 uint16 = 65535
                                            for fi21 := 0; fi21 < 20; fi21++ {
                                                var v22 float64 = v6
                                                var v23 uint32 = v10
                                                var v24 int = fi21
                                                _ = fi21
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            }
                                            _ = v20
                                        }
                                    }
                                }
                                fmt.Println(v1, v2)
                                for fi25 := 0; fi25 < 16; fi25++ {
                                    var v26 int = 2147483647
                                    _ = len(v5)
                                    _ = fi25
                                    _ = v26
                                }
                            }
                            _ = len(v3)
                            _ = v17
                        } else {
                            if len(v5) > 0 {
                                v5[0] = 128
                            }
                            var v27 uint16 = 106
                            var v28 int32 = 1
                            _ = v27
                            _ = v28
                        }
                    }
                }
            }
        case -9223372036854775808:
            var v29 int16 = 42
            _ = v29
        }
        _ = v10
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v3) > 0 {
        v3[0] = ("hello" + "a")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
