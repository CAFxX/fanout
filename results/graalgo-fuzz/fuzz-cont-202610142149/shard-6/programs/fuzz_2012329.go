package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 string }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS7 struct { F0 uint; F1 uint16 }

func main() {
    v1 := []string{((((("hello" + ("foo bar" + "foo bar")) + ("a" + "hello")) + "hello") + "a") + ((("hello" + "a") + ("foo bar" + "foo bar")) + (("\n\t" + "hello") + ""))), ((("\n\t" + "") + "a") + (("" + "") + "hello"))}
    var v2 float32 = 1.5
    var v3 uint64 = 0
    var v4 uint32 = 707
    var v5 int8 = 0
    var v6 uint8 = 128
    var v7 uint = 255
    for fi8 := 0; fi8 < 11; fi8++ {
        v1 = append(v1, ((("foo bar" + ("\n\t" + "foo bar")) + "\n\t") + "hello"))
        _ = len(v1)
        _ = fi8
    }
    v3 -= (v3 << ((uint(v6)) & 63))
    if (-2147483648 != 256) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v6 -= 42
        switch 128 {
        case 9223372036854775807:
            v3 = (v3 << ((77) & 63))
        case 770:
            var v9 int = 462
            _ = v9
        default:
            switch "" {
            case "foo bar":
                switch "foo bar" {
                case "a":
                    fmt.Println(v3, v7)
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v6, v3, v5)
                    case "":
                        for fi10 := 0; fi10 < 1; fi10++ {
                            for fi11 := 0; fi11 < 9; fi11++ {
                                var v12 int = (fi11 ^ (fi11 * (^(fi11))))
                                v7 += v7
                                _ = fi11
                                _ = v12
                            }
                            v13 := GenS1{F0: (-1 <= -294), F1: (-(v5)), F2: "a"}
                            switch "hello" {
                            case "foo bar":
                                var v14 int32 = -213
                                _ = v14
                            case "\n\t":
                                v15 := GenS1{F0: (!(!(972 != 929))), F1: v5, F2: (("a" + "foo bar") + ("" + ("" + "\n\t")))}
                                _ = v15
                            }
                            _ = fi10
                            _ = v13
                        }
                    case "hello":
                        v1 = append(v1, "hello")
                    }
                case "hello":
                    switch -9223372036854775808 {
                    case 127:
                        if true {
                            _ = len(v1)
                            v3 = v3
                        } else {
                            v16 := GenS2{F0: v3, F1: v4, F2: v3}
                            for fi17 := 0; fi17 < 6; fi17++ {
                                var v18 uint32 = (v4 | v4)
                                _ = fi17
                                _ = v18
                            }
                            _ = v16
                        }
                    case 691:
                        var v19 int8 = v5
                        _ = v19
                    default:
                        v2 = 0.1
                    }
                }
            case "":
                var v20 uint16 = 42
                _ = v20
            }
        }
    }
    v21 := GenS3{F0: -2147483648, F1: 42, F2: v5}
    v5 += v5
    if true {
        v7 -= v7
        for fi22 := 0; fi22 < 5; fi22++ {
            v23 := GenS2{F0: (v3 & (v3 >> ((v7) & 63))), F1: 65535, F2: 0}
            v5 += 1
            _ = fi22
            _ = v23
        }
    }
    var v24 string = ((("\n\t" + "foo bar") + "a") + (("foo bar" + "") + ("foo bar" + ("a" + "foo bar"))))
    var v25 int = 7
    switch ((v24 + v24) + v24) {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = v24
        }
    case "\n\t":
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 int64 = 42
            v28 := GenS4{F0: 823, F1: 673, F2: v4}
            switch (v24 + v24) {
            case "":
                var v29 uint64 = (v3 / (((v3 ^ uint64(v4))) | 1))
                _ = v29
            }
            _ = fi26
            _ = v27
            _ = v28
        }
    case "":
        fmt.Println(v24, v6)
    default:
        v30 := GenS5{F0: uint(v6), F1: (^(v6)), F2: v6}
        _ = v30
    }
    var v31 uint16 = 65535
    v31 -= 65535
    for fi32 := 0; fi32 < 2; fi32++ {
        v7 += (v7 ^ (v7 & v7))
        _ = fi32
    }
    v33 := GenS6{F0: -2147483648, F1: v4, F2: (v6 + v6)}
    v34 := GenS7{F0: (v7 / (((v7 + uint(v6))) | 1)), F1: (v31 * (v31 | (^(v31))))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
