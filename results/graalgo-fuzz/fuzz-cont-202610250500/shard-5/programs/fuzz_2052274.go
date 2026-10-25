package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }

func main() {
    v1 := []uint32{42}
    var v2 int64 = -9223372036854775808
    var v3 int8 = 0
    var v4 uint8 = 168
    var v5 float32 = 0.1
    var v6 float32 = 3.14159
    if true {
        for fi7 := 0; fi7 < 5; fi7++ {
            v1 = append(v1, 65535)
            if (v5 != (v6 + (1e10 + (v6 / 1.5)))) {
                v8 := GenS1{F0: 0, F1: (("" + "hello") + ""), F2: 65535}
                v9 := GenS1{F0: 0, F1: "\n\t", F2: 899}
                for fi10 := 0; fi10 < 14; fi10++ {
                    switch "hello" {
                    case "foo bar":
                        switch "" {
                        case "foo bar":
                            if (v4 > 241) {
                                v3 = (v3 * (v3 - v3))
                                if false {
                                    _ = len(v1)
                                    v11 := GenS1{F0: 2, F1: "a", F2: 1}
                                    _ = v11
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        case "hello":
                            var v12 uint32 = 0
                            _ = v12
                        case "\n\t":
                            v2 = (v2 | 9223372036854775807)
                        }
                    case "a":
                        v13 := GenS1{F0: 2, F1: "a", F2: 745}
                        _ = v13
                    case "":
                        fmt.Println(v6)
                        var v14 uint8 = v4
                        _ = v14
                    default:
                        var v15 float32 = v5
                        var v16 float32 = v6
                        _ = v15
                        _ = v16
                    }
                    switch -1 {
                    case -2147483648:
                        v17 := GenS1{F0: 18446744073709551615, F1: ("" + "a"), F2: 609}
                        _ = v17
                    case 0:
                        v6 = ((v5 + v6) * v5)
                    default:
                        fmt.Println(fi10, v6, v5)
                    }
                    _ = fi10
                }
                _ = v8
                _ = v9
            }
            v18 := GenS1{F0: 0, F1: "hello", F2: 437}
            _ = fi7
            _ = v18
        }
    }
    if false {
        var v19 float32 = v5
        _ = v19
    }
    var v20 uint64 = 253
    v21 := GenS2{F0: 587, F1: (32767 != 110), F2: v3}
    var v22 uint16 = 637
    v23 := GenS1{F0: (v20 - v20), F1: "", F2: v22}
    for fi24 := 0; fi24 < 5; fi24++ {
        if true {
            v4 -= v4
        }
        _ = fi24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        var v26 int8 = v3
        v27 := GenS1{F0: (v20 / (((v20 - (v20 | 1))) | 1)), F1: ("\n\t" + "foo bar"), F2: 343}
        _ = fi25
        _ = v26
        _ = v27
    }
    v1 = append(v1, 65535)
    var v28 int16 = 1
    v20 += (v20 | (^(v20)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
