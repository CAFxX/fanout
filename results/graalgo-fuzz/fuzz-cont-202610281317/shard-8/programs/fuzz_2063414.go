package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint; F2 uint64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int16; F1 uint }

func main() {
    v1 := []string{"\n\t", "hello", "hello", (("a" + "\n\t") + "foo bar")}
    var v2 uint32 = 2147483647
    v3 := []string{""}
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    var v6 int64 = 1
    fmt.Println(v2, v2)
    v6 = (v6 ^ v6)
    fmt.Println(v6)
    _ = len(v5)
    v7 := GenS1{F0: ("a" + ("" + "a"))}
    fmt.Println(v6, v2)
    v1 = append(v1, ("a" + "a"))
    v6 += (v6 / ((-9223372036854775808) | 1))
    delete(v5, 2147483647)
    if (!(!(!(v2 > (v2 & v2))))) {
        for fi8 := 0; fi8 < 6; fi8++ {
            v6 -= -1
            _ = fi8
        }
    } else {
        v6 += 42
        var v9 int16 = 32767
        _ = len(v1)
        _ = v9
    }
    v6 += (v6 / ((-557) | 1))
    v6 = (v6 | -1)
    var v10 int32 = 1
    switch -1 {
    case -462:
        for fi11 := 0; fi11 < 5; fi11++ {
            v12 := GenS1{F0: "foo bar"}
            var v13 int32 = (v10 >> ((42) & 63))
            _ = fi11
            _ = v12
            _ = v13
        }
    case 128:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 253:
        if ((false && false) && ((42 < 127) && (!false))) {
            switch "\n\t" {
            case "\n\t":
                var v14 int8 = 0
                _ = v14
            case "hello":
                var v15 int = -1
                _ = v15
            }
            v16 := GenS2{F0: v10, F1: 387, F2: 497}
            v6 = (^(v6))
            _ = v16
        } else {
            switch ("a" + ("\n\t" + ("hello" + ("\n\t" + ("\n\t" + ("\n\t" + "\n\t")))))) {
            case "hello":
                v10 = v10
            case "":
                v10 -= v10
            }
            if false {
                var v17 uint64 = 42
                _ = v17
            } else {
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                if (!false) {
                    var v18 string = "\n\t"
                    var v19 uint = 2
                    if (false && (!(v19 >= v19))) {
                        for fi20 := 0; fi20 < 4; fi20++ {
                            if (2 == 0) {
                                v19 = (v19 ^ v19)
                                var v21 uint8 = 42
                                _ = v21
                            } else {
                                fmt.Println(v19)
                            }
                            _ = fi20
                        }
                        switch 833 {
                        case 77:
                            v22 := GenS1{F0: ("foo bar" + v18)}
                            _ = v22
                        case 416:
                            switch (("\n\t" + "a") + (v18 + v18)) {
                            case "":
                                v23 := v5[42]
                                _ = v23
                            default:
                                fmt.Println(v6)
                            }
                        case -2147483648:
                            switch v18 {
                            case "\n\t":
                                v6 += v6
                            case "hello":
                                v24 := GenS2{F0: v10, F1: 18446744073709551615, F2: 861}
                                _ = v24
                            default:
                                _ = len(v5)
                            }
                        default:
                            for fi25 := 0; fi25 < 14; fi25++ {
                                v26 := GenS3{F0: v6, F1: 1}
                                _ = fi25
                                _ = v26
                            }
                        }
                    }
                    _ = v18
                    _ = v19
                } else {
                    v3 = append(v3, ("a" + ("hello" + "foo bar")))
                }
                v27 := GenS4{F0: 1000, F1: 0}
                _ = v27
            }
        }
    default:
        fmt.Println(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
