package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 string }
type GenS3 struct { F0 int16; F1 uint16; F2 uint8 }

func main() {
    v1 := []int{}
    var v2 int32 = 1
    v3 := make(map[string]string)
    v3["hello"] = "a"
    v2 -= (v2 * (v2 ^ v2))
    if false {
        if false {
            var v4 float64 = (1e10 / (1e10 - 3.14159))
            if (("\n\t" <= "a") || ("\n\t" < (("a" + "") + (("" + "foo bar") + "\n\t")))) {
                for fi5 := 0; fi5 < 13; fi5++ {
                    v6 := GenS1{F0: (v2 >> ((18446744073709551615) & 63)), F1: 42}
                    _ = fi5
                    _ = v6
                }
                for fi7 := 0; fi7 < 6; fi7++ {
                    for fi8 := 0; fi8 < 4; fi8++ {
                        v4 = float64((1e10 / 3.14159))
                        v2 += (v2 & (-(v2)))
                        _ = fi8
                    }
                    v2 = 0
                    v2 = v2
                    _ = fi7
                }
                var v9 uint8 = 2
                _ = v9
            }
            _ = v4
        }
    } else {
        switch (("\n\t" + "hello") + "") {
        case "foo bar":
            if (434 != 42) {
                var v10 int16 = 822
                _ = v10
            } else {
                var v11 bool = (0 <= 2)
                switch -534 {
                case -542:
                    v11 = (v11 || (!(!v11)))
                case 127:
                    if v11 {
                        v2 = v2
                        switch "foo bar" {
                        case "":
                            for fi12 := 0; fi12 < 20; fi12++ {
                                v13 := v3["\n\t"]
                                v14 := GenS2{F0: -1, F1: 100, F2: v13}
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                        case "foo bar":
                            v11 = false
                        case "\n\t":
                            for fi15 := 0; fi15 < 16; fi15++ {
                                switch fi15 {
                                case -653:
                                    for fi16 := 0; fi16 < 17; fi16++ {
                                        v3[("\n\t" + "")] = "a"
                                        _ = fi16
                                    }
                                }
                                _ = fi15
                            }
                        default:
                            var v17 float32 = 1e10
                            _ = v17
                        }
                        var v18 float32 = (-2.25 + 1e10)
                        _ = v18
                    }
                case 42:
                    v19 := GenS1{F0: (v2 / ((v2) | 1)), F1: 42}
                    _ = v19
                default:
                    var v20 int16 = 531
                    _ = v20
                }
                v11 = true
                _ = v11
            }
        case "hello":
            fmt.Println(v2)
        case "a":
            if true {
                var v21 int8 = 1
                v21 -= (v21 / ((v21) | 1))
                v21 = (v21 % ((v21) | 1))
                _ = v21
            } else {
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
            }
        default:
            var v22 int = -2147483648
            _ = v22
        }
    }
    v2 += v2
    fmt.Println(v2)
    if (false || false) {
        v23 := GenS3{F0: -32768, F1: 32767, F2: 42}
        _ = len(v3)
        _ = v23
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 float32 = 0.1
        _ = fi24
        _ = v25
    }
    switch 42 {
    case 32768:
        v2 = (v2 % ((796) | 1))
    }
    var v26 uint8 = 1
    if true {
        v26 += v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 uint8 = v26
        _ = fi27
        _ = v28
    }
    var v29 uint32 = 2
    fmt.Println(v26, v26, v29)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v30 int64 = -9223372036854775808
    v26 = v26
    switch 128 {
    case -9223372036854775808:
        var v31 int64 = v30
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
}
