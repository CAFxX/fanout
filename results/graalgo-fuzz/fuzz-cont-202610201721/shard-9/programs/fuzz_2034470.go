package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint16 }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (82 > 1)
    v3 := make(map[int]string)
    v4 := []uint32{}
    fmt.Println(v2)
    var v5 uint8 = 128
    v2 = v2
    var v6 uint16 = 255
    v1[2147483648] = 1
    var v7 int8 = -1
    fmt.Println(v2, v7, v7)
    var v8 int16 = 42
    if true {
        v9 := GenS1{F0: 221, F1: (("foo bar" + ("hello" + ("foo bar" + "a"))) + "a"), F2: v6}
        v4 = append(v4, 4294967295)
        switch "\n\t" {
        case "foo bar":
            if v2 {
                fmt.Println(v5, v2)
                v10 := GenS1{F0: 0, F1: "hello", F2: 1}
                if (!v2) {
                    _ = len(v4)
                    _ = len(v4)
                    var v11 string = ""
                    _ = v11
                }
                _ = v10
            }
        default:
            if false {
                for fi12 := 0; fi12 < 14; fi12++ {
                    fmt.Println(v2, v6, v5)
                    switch (fi12 << ((18446744073709551615) & 63)) {
                    case -1:
                        var v13 uint16 = (v6 << ((205) & 63))
                        _ = v13
                    case 0:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        var v14 uint8 = (v5 + 180)
                        _ = v14
                    }
                    v4 = append(v4, 872)
                    _ = fi12
                }
                if false {
                    if len(v4) > 0 {
                        v4[0] = 42
                    }
                    for fi15 := 0; fi15 < 14; fi15++ {
                        var v16 string = (("a" + ("foo bar" + ("foo bar" + ""))) + "\n\t")
                        _ = fi15
                        _ = v16
                    }
                } else {
                    var v17 int8 = -69
                    _ = v17
                }
                _ = len(v4)
            }
        }
        _ = v9
    } else {
        v18 := GenS1{F0: -1, F1: "foo bar", F2: 449}
        _ = v18
    }
    if v2 {
        if true {
            v19 := v1[0]
            delete(v3, v19)
            v2 = false
            _ = v19
        } else {
            for fi20 := 0; fi20 < 8; fi20++ {
                switch "" {
                case "":
                    switch 2147483647 {
                    case 9223372036854775807:
                        for fi21 := 0; fi21 < 2; fi21++ {
                            _ = len(v1)
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if ("\n\t" >= "a") {
                                    _ = len(v3)
                                }
                                var v22 int32 = -592
                                _ = v22
                            } else {
                                v8 += (v8 & (v8 ^ int16(fi20)))
                                var v23 int64 = -9223372036854775808
                                fmt.Println(v8)
                                _ = v23
                            }
                            v2 = (!v2)
                            fmt.Println(fi21)
                            _ = fi21
                        }
                    case -9223372036854775808:
                        var v24 string = ("a" + ("" + "hello"))
                        _ = v24
                    case -1:
                        var v25 uint = 4294967295
                        _ = v25
                    default:
                        var v26 int64 = 9223372036854775807
                        _ = v26
                    }
                default:
                    _ = len(v1)
                }
                switch "a" {
                case "foo bar":
                    v6 = 32767
                case "\n\t":
                    v5 -= v5
                default:
                    v27 := GenS2{F0: (v7 & v7), F1: (v7 >> ((982) & 63))}
                    _ = v27
                }
                _ = fi20
            }
        }
    }
    if v2 {
        v5 = 255
    } else {
        var v28 int8 = (v7 + (v7 ^ (v7 << ((2) & 63))))
        _ = v28
    }
    v1[42] = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
