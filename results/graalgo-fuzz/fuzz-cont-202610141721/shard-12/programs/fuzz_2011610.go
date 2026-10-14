package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int16 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int64; F1 string; F2 int16 }
func genFn0(p0 uint16) int32 {
    var v4 string = ("a" + "\n\t")
    v5 := []int{127, 975, -1, 0}
    if len(v5) > 0 {
        v5[0] = -991
    }
    var v6 bool = (true && false)
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    v1 := []string{"hello", ((("\n\t" + "a") + ((("foo bar" + "hello") + "") + "foo bar")) + ("\n\t" + (("" + "foo bar") + ""))), "\n\t", ("\n\t" + "")}
    var v2 int64 = -1
    var v3 uint8 = 0
    _ = len(v1)
    var v7 int16 = 795
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint64 = 0
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        _ = fi8
        _ = v9
    }
    switch "" {
    case "foo bar":
        for fi10 := 0; fi10 < 15; fi10++ {
            if false {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v3, fi10)
                case "foo bar":
                    if ((v7 ^ (v7 | (v7 & (v7 & v7)))) != v7) {
                        if (!false) {
                            var v11 int16 = (v7 | 0)
                            v12 := GenS1{F0: ("" + "hello"), F1: 249, F2: (v7 >> ((1) & 63))}
                            _ = v11
                            _ = v12
                        } else {
                            var v13 uint32 = 65535
                            if ((("\n\t" + "hello") + ("a" + "")) != (("" + ("" + ("foo bar" + (("a" + "\n\t") + ("foo bar" + ""))))) + "a")) {
                                switch "hello" {
                                case "hello":
                                    v14 := GenS1{F0: ("hello" + ("" + "a")), F1: (v3 / (((v3 ^ v3)) | 1)), F2: v7}
                                    _ = v14
                                case "foo bar":
                                    v7 -= v7
                                case "a":
                                    if (255 != 28) {
                                        if (0 > 449) {
                                            for fi15 := 0; fi15 < 9; fi15++ {
                                                var v16 int8 = int8(fi15)
                                                v13 -= v13
                                                var v17 float32 = 0.0
                                                _ = fi15
                                                _ = v16
                                                _ = v17
                                            }
                                        }
                                    } else {
                                        v2 = -1
                                        v13 = 626
                                    }
                                default:
                                    v13 -= v13
                                }
                            }
                            var v18 int8 = -128
                            _ = v13
                            _ = v18
                        }
                        var v19 uint = 1
                        _ = v19
                    } else {
                        switch "" {
                        case "\n\t":
                            v20 := GenS2{F0: -9223372036854775808, F1: 2147483648}
                            _ = v20
                        case "hello":
                            if false {
                                v3 = v3
                                fmt.Println(v3, v2, v2)
                                v21 := GenS2{F0: (v2 + (^(v2))), F1: uint32(v3)}
                                _ = v21
                            }
                        }
                        var v22 uint8 = 186
                        _ = v22
                    }
                case "a":
                    for fi23 := 0; fi23 < 5; fi23++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        switch ("a" + "\n\t") {
                        case "\n\t":
                            v24 := GenS2{F0: 0, F1: 803}
                            _ = v24
                        }
                        _ = fi23
                    }
                default:
                    _ = len(v1)
                }
                var v25 uint8 = v3
                v25 += (v25 % ((v3) | 1))
                _ = v25
            }
            fmt.Println(v3)
            _ = fi10
        }
    }
    var v26 uint8 = (v3 + v3)
    var v27 bool = true
    v28 := GenS2{F0: (-(v2)), F1: 768}
    var v29 string = ("\n\t" + ("foo bar" + "foo bar"))
    switch 570 {
    case 4294967295:
        v1 = append(v1, (v29 + "a"))
    }
    v26 -= (^(v26))
    v30 := GenS3{F0: (-(v2)), F1: "hello", F2: v7}
    v27 = (true && ((v2 + v2) == -9223372036854775808))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v7
}
