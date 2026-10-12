package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint; F1 uint }

func main() {
    v1 := []string{"", "\n\t", "", ("hello" + ("hello" + ("foo bar" + "a")))}
    v2 := make(map[uint32]int)
    var v3 uint16 = 32767
    v4 := []string{"foo bar"}
    v3 = (^(v3))
    if ((!(42 < 2)) || false) {
        v3 = v3
        v3 = (v3 - 42)
        v5 := GenS1{F0: "a", F1: 85, F2: true}
        _ = v5
    } else {
        fmt.Println(v3)
        if (((v3 >= (^(v3))) && (127 >= 127)) || ("" >= "hello")) {
            for fi6 := 0; fi6 < 17; fi6++ {
                v3 += (v3 + 0)
                _ = len(v4)
                v7 := GenS1{F0: ("\n\t" + "\n\t"), F1: 18446744073709551615, F2: ("a" >= ("a" + ""))}
                _ = fi6
                _ = v7
            }
            if (659 > 2) {
                switch "hello" {
                case "\n\t":
                    v3 -= v3
                default:
                    v1 = append(v1, "a")
                }
                var v8 uint8 = 0
                if (!false) {
                    if (0.1 != 1.5) {
                        for fi9 := 0; fi9 < 19; fi9++ {
                            if true {
                                var v10 float32 = 1.5
                                v11 := GenS1{F0: "hello", F1: 861, F2: (869 != 0)}
                                _ = len(v1)
                                _ = v10
                                _ = v11
                            } else {
                                switch 917 {
                                case 2147483647:
                                    v12 := v2[2]
                                    _ = v12
                                default:
                                    if true {
                                        if false {
                                            v8 = (v8 + v8)
                                        }
                                    } else {
                                        fmt.Println(v3, fi9, v8)
                                        v13 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: false}
                                        delete(v2, 42)
                                        _ = v13
                                    }
                                }
                            }
                            fmt.Println(v3, fi9, v8)
                            v3 -= v3
                            _ = fi9
                        }
                        v3 = (v3 % ((v3) | 1))
                        v14 := GenS2{F0: 969}
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 10; fi15++ {
                            if (4294967295 == 2147483647) {
                                _ = len(v2)
                            }
                            _ = fi15
                        }
                    }
                } else {
                    v16 := GenS3{F0: -2147483648, F1: 1}
                    if (true && (-1 == 1)) {
                        fmt.Println(v3)
                    }
                    if (!(false || (!true))) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        fmt.Println(v8, v8)
                        fmt.Println(v8, v3)
                    }
                    _ = v16
                }
                _ = v8
            }
            var v17 bool = (-1 == 127)
            _ = v17
        }
    }
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    v3 += (v3 - v3)
    v3 -= (v3 ^ 744)
    var v18 string = ""
    v19 := GenS3{F0: 0, F1: -389}
    if (!(-571 <= 9223372036854775807)) {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 uint32 = 65535
            _ = fi20
            _ = v21
        }
    }
    fmt.Println(v3, v18)
    switch 0 {
    case 128:
        v22 := GenS4{F0: -244, F1: 4294967295}
        _ = v22
    case -1:
        v23 := GenS2{F0: -1}
        _ = v23
    case 258:
        v3 = (^(v3))
    }
    var v24 uint64 = 42
    v3 = v3
    v18 = ((v18 + (v18 + "hello")) + (("" + ((v18 + ("foo bar" + ((v18 + v18) + "\n\t"))) + ((v18 + (v18 + v18)) + v18))) + v18))
    v25 := GenS4{F0: 32767, F1: 42}
    for fi26 := 0; fi26 < 1; fi26++ {
        v2[0] = (fi26 >> ((4294967295) & 63))
        _ = fi26
    }
    v27 := GenS5{F0: 255, F1: 1}
    v28 := GenS4{F0: 1, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
