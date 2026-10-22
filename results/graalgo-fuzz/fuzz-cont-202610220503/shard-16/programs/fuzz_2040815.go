package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("\n\t" + "a"), "a", "hello"}
    v3 := []uint32{886, 149, 431}
    var v4 int32 = -2147483648
    v3 = append(v3, 2147483648)
    v5 := GenS1{F0: 2, F1: 408}
    var v6 uint32 = 784
    if (!false) {
        for fi7 := 0; fi7 < 7; fi7++ {
            v6 += (v6 * (^(v6)))
            _ = fi7
        }
    }
    if false {
        v8 := GenS1{F0: 141, F1: 0}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 int8 = 50
            fmt.Println(fi9)
            _ = fi9
            _ = v10
        }
    }
    var v11 uint8 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        if ((v4 & v4) == 955) {
            var v12 string = ("a" + "")
            var v13 uint = 1
            if (-2147483648 != 32768) {
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 int16 = -1
                    _ = fi14
                    _ = v15
                }
                switch v12 {
                case "a":
                    fmt.Println(v6)
                case "":
                    switch ("\n\t" + ((("foo bar" + v12) + ((v12 + (("" + v12) + (v12 + "foo bar"))) + v12)) + "foo bar")) {
                    case "":
                        for fi16 := 0; fi16 < 17; fi16++ {
                            if false {
                                var v17 uint = (v13 / ((v13) | 1))
                                if false {
                                    v4 += 245
                                } else {
                                    var v18 uint16 = 0
                                    switch (fi16 >> ((65535) & 63)) {
                                    case 42:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case 786:
                                        var v19 int = fi16
                                        _ = v19
                                    default:
                                        v20 := GenS2{F0: fi16, F1: 32767}
                                        _ = v20
                                    }
                                    if true {
                                        v11 -= (^(v11))
                                        v6 += v6
                                        v4 = (v4 << (((v17 ^ (v13 * (v17 / (((v13 % ((2) | 1))) | 1))))) & 63))
                                    } else {
                                        if ((fi16 << (((^(v17))) & 63)) < 32767) {
                                            v21 := v1[fi16]
                                            if (1.5 == float64(-2.25)) {
                                                if (!(false && (32767 > 0))) {
                                                    v4 -= -1
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                _ = len(v2)
                                            }
                                            v22 := GenS3{F0: 50, F1: -652}
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                    _ = v18
                                }
                                var v23 uint64 = 4294967295
                                _ = v17
                                _ = v23
                            }
                            var v24 float32 = (float32(float64(((1e10 - 1.5) * 3.14159))) - float32(-2.25))
                            _ = fi16
                            _ = v24
                        }
                    case "\n\t":
                        if true {
                            v25 := GenS1{F0: 42, F1: -785}
                            v12 = (v12 + (v12 + v12))
                            var v26 uint64 = 42
                            _ = v25
                            _ = v26
                        } else {
                            var v27 int = 42
                            v12 = v12
                            _ = v27
                        }
                    }
                default:
                    v28 := GenS1{F0: 723, F1: -1}
                    _ = v28
                }
            }
            _ = v12
            _ = v13
        } else {
            for fi29 := 0; fi29 < 7; fi29++ {
                v30 := GenS1{F0: 18446744073709551615, F1: 42}
                _ = fi29
                _ = v30
            }
        }
    } else {
        v31 := GenS3{F0: 42, F1: 42}
        v1[0] = true
        var v32 int16 = 42
        _ = v31
        _ = v32
    }
    v11 = (v11 & (v11 + (^(v11))))
    var v33 int16 = 14
    fmt.Println(v33, v11)
    v1[2147483647] = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
