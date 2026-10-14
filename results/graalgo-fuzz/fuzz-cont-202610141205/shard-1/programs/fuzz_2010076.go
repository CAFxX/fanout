package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
type GenS3 struct { F0 int8; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 uint8 = 155
    fmt.Println(v4)
    v4 += v4
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: 1, F1: ((1.5 - 1.5) == (-2.25 * 1.5)), F2: v4}
            v4 = (^(v4))
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 982, F1: (((0.0 + -2.25) + 1.5) <= 0.1), F2: v4}
        _ = fi5
        _ = v8
    }
    v9 := v2[""]
    for fi10 := 0; fi10 < 19; fi10++ {
        v4 += (v4 >> ((125) & 63))
        for fi11 := 0; fi11 < 18; fi11++ {
            var v12 uint = 255
            _ = fi11
            _ = v12
        }
        fmt.Println(v9, fi10)
        _ = fi10
    }
    switch 42 {
    case 284:
        v3[("a" + "a")] = 2147483647
    default:
        switch 4294967295 {
        case -679:
            for fi13 := 0; fi13 < 13; fi13++ {
                if (274 == 446) {
                    for fi14 := 0; fi14 < 17; fi14++ {
                        fmt.Println(v9)
                        switch "\n\t" {
                        case "a":
                            v9 = ("\n\t" > "foo bar")
                        case "":
                            v15 := GenS1{F0: 42, F1: true, F2: v4}
                            _ = v15
                        case "\n\t":
                            var v16 bool = v9
                            _ = v16
                        }
                        if (v9 || v9) {
                            if (!(18446744073709551615 != 18446744073709551615)) {
                                var v17 int8 = 100
                                _ = v17
                            }
                        } else {
                            fmt.Println(v9, v4)
                            v18 := v3[("foo bar" + "a")]
                            _ = v18
                        }
                        _ = fi14
                    }
                } else {
                    v19 := []int{(fi13 + fi13), (fi13 >> ((65535) & 63))}
                    v9 = true
                    _ = v19
                }
                switch "a" {
                case "a":
                    for fi20 := 0; fi20 < 8; fi20++ {
                        var v21 string = "foo bar"
                        var v22 uint16 = 42
                        v23 := GenS1{F0: 2, F1: false, F2: v4}
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                }
                switch "" {
                case "foo bar":
                    if true {
                        v4 += (v4 ^ v4)
                        for fi24 := 0; fi24 < 9; fi24++ {
                            v3["foo bar"] = 0
                            v25 := GenS1{F0: 42, F1: ((v9 && v9) && ((0.1 <= -2.25) || (137 <= 1))), F2: v4}
                            switch -9223372036854775808 {
                            case 1:
                                v26 := v2[("" + "\n\t")]
                                _ = v26
                            case -2147483648:
                                var v27 int8 = -38
                                _ = v27
                            default:
                                var v28 bool = (!((1 > 988) && v9))
                                _ = v28
                            }
                            var v29 int16 = 616
                            var v30 int = 2147483647
                            _ = fi24
                            _ = v25
                            _ = v29
                            _ = v30
                        }
                        v31 := GenS2{F0: v9, F1: -350, F2: 868}
                        _ = v31
                    }
                case "hello":
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 int = 4294967295
                        fmt.Println(fi13)
                        _ = fi32
                        _ = v33
                    }
                default:
                    if v9 {
                        fmt.Println(v9, v9)
                    } else {
                        v34 := GenS3{F0: 1, F1: 558, F2: 42}
                        var v35 int8 = 15
                        v36 := GenS2{F0: v9, F1: 0, F2: 712}
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                }
                _ = fi13
            }
        case -662:
            _ = len(v1)
        case 9223372036854775807:
            fmt.Println(v4, v4)
        }
    }
    v9 = ((v4 & v4) <= v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
