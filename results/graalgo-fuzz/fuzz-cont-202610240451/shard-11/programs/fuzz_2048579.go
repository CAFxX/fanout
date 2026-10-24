package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int8; F1 int32; F2 uint8 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 uint16; F1 bool; F2 int64 }

func main() {
    var v1 int32 = 0
    v2 := []string{"foo bar", "a"}
    var v3 uint8 = 255
    var v4 float32 = 3.14159
    var v5 int64 = 9223372036854775807
    v6 := make(map[int]bool)
    v7 := GenS1{F0: 0, F1: 1000}
    v4 += 1e10
    v8 := GenS1{F0: 127, F1: -32768}
    for fi9 := 0; fi9 < 12; fi9++ {
        v3 = (v3 & v3)
        v10 := GenS2{F0: 6, F1: v1, F2: (^(v3))}
        var v11 uint64 = 4294967295
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := v6[32768]
    for fi13 := 0; fi13 < 17; fi13++ {
        v3 = (v3 % (((v3 & 0)) | 1))
        _ = fi13
    }
    if v12 {
        if (("\n\t" <= "a") && true) {
            v3 += v3
            if (!v12) {
                if v12 {
                    fmt.Println(v5)
                    switch -994 {
                    case -344:
                        v14 := GenS2{F0: -128, F1: 862, F2: v3}
                        _ = v14
                    case 2147483647:
                        for fi15 := 0; fi15 < 16; fi15++ {
                            for fi16 := 0; fi16 < 2; fi16++ {
                                v17 := GenS3{F0: -9223372036854775808, F1: (^(fi16))}
                                if true {
                                    if v12 {
                                        fmt.Println(fi16)
                                        if (2 >= 831) {
                                            v12 = ((32767 < 255) || false)
                                            var v18 uint8 = 0
                                            var v19 uint8 = (v18 ^ v18)
                                            v5 -= int64(fi15)
                                            _ = v18
                                            _ = v19
                                        } else {
                                            _ = len(v2)
                                            v5 -= v5
                                        }
                                        switch (-(fi16)) {
                                        case -2147483648:
                                            if ((v5 << ((2) & 63)) != (v5 << ((262) & 63))) {
                                                v20 := GenS2{F0: 0, F1: -1, F2: (v3 + 2)}
                                                fmt.Println(fi16, v5, v4)
                                                v1 += v1
                                                _ = v20
                                            }
                                        case -976:
                                            v5 -= (v5 & 641)
                                        case 689:
                                            switch ("foo bar" + ("" + "hello")) {
                                            case "":
                                                var v21 uint16 = 479
                                                _ = v21
                                            }
                                        }
                                    } else {
                                        v22 := GenS2{F0: 127, F1: (^(v1)), F2: (v3 % (((v3 >> ((255) & 63))) | 1))}
                                        _ = v22
                                    }
                                    switch ("foo bar" + "hello") {
                                    case "foo bar":
                                        if v12 {
                                            v23 := GenS1{F0: 127, F1: -904}
                                            fmt.Println(v1, fi15, v12)
                                            var v24 int16 = -413
                                            _ = v23
                                            _ = v24
                                        }
                                    default:
                                        v1 -= v1
                                    }
                                }
                                _ = fi16
                                _ = v17
                            }
                            v12 = v12
                            v1 = (v1 % ((v1) | 1))
                            _ = fi15
                        }
                    case -1:
                        var v25 uint16 = 42
                        _ = v25
                    default:
                        var v26 int64 = v5
                        _ = v26
                    }
                    var v27 int8 = 127
                    _ = v27
                } else {
                    v6[-752] = v12
                    for fi28 := 0; fi28 < 11; fi28++ {
                        fmt.Println(v3, v1)
                        v4 -= v4
                        _ = fi28
                    }
                }
            } else {
                fmt.Println(v4, v3)
            }
        } else {
            fmt.Println(v4)
            if (65535 < 4294967295) {
                fmt.Println(v1, v3)
                v12 = ((65535 == 114) && v12)
                var v29 uint64 = 285
                _ = v29
            }
        }
    }
    switch 127 {
    case 9223372036854775807:
        v5 = v5
    case -658:
        v30 := GenS4{F0: 2, F1: v12, F2: v5}
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        v12 = false
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi31
    }
    switch 454 {
    case 0:
        v12 = (!(-1 >= 860))
    default:
        var v32 string = ""
        _ = v32
    }
    fmt.Println(v4, v4, v12)
    for fi33 := 0; fi33 < 10; fi33++ {
        switch (("\n\t" + "\n\t") + ("foo bar" + "foo bar")) {
        case "foo bar":
            if (42 == 2147483647) {
                _ = len(v2)
                var v34 uint8 = (v3 / ((v3) | 1))
                _ = v34
            }
        default:
            v4 = (v4 * float32(3.14159))
        }
        _ = fi33
    }
    v35 := GenS2{F0: 1, F1: v1, F2: (^(v3))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
