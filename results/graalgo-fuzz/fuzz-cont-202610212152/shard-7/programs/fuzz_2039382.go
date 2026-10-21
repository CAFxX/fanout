package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint16 = 834
    if true {
        switch "\n\t" {
        case "":
            var v5 int = 0
            _ = v5
        case "foo bar":
            if (!false) {
                if (!true) {
                    v4 -= 63
                }
                for fi6 := 0; fi6 < 12; fi6++ {
                    fmt.Println(fi6)
                    _ = fi6
                }
                v7 := v1[((("foo bar" + "") + "hello") + "")]
                _ = v7
            }
        case "\n\t":
            _ = len(v2)
        default:
            fmt.Println(v4)
        }
        v8 := GenS1{F0: -181, F1: 255, F2: 1000}
        v4 = v4
        _ = v8
    }
    v4 = (v4 % ((1) | 1))
    v4 = v4
    v9 := GenS1{F0: -1, F1: 255, F2: -697}
    switch ("" + "") {
    case "hello":
        if false {
            v10 := []int{128}
            v4 -= (^(v4))
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = v10
        } else {
            if (409 == 65535) {
                for fi11 := 0; fi11 < 12; fi11++ {
                    v12 := GenS2{F0: int64(fi11), F1: 42}
                    if false {
                        v13 := []int{-2147483648, fi11}
                        var v14 uint32 = 2
                        var v15 bool = ((-9223372036854775808 < 0) && true)
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 5; fi16++ {
                        var v17 bool = false
                        v4 = (v4 << ((4294967295) & 63))
                        v17 = (127 != 127)
                        _ = fi16
                        _ = v17
                    }
                    _ = fi11
                    _ = v12
                }
                delete(v1, (("\n\t" + "hello") + "\n\t"))
            } else {
                v4 = 451
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v4)
                }
                v4 -= v4
            }
            v4 -= 32767
            for fi18 := 0; fi18 < 18; fi18++ {
                v19 := GenS2{F0: -9223372036854775808, F1: 32767}
                v20 := GenS3{F0: -124}
                switch "foo bar" {
                case "hello":
                    v21 := GenS3{F0: 2147483647}
                    _ = v21
                case "":
                    if true {
                        fmt.Println(fi18, fi18)
                    } else {
                        var v22 string = ("a" + (("foo bar" + "hello") + ((("foo bar" + "a") + (("foo bar" + "hello") + "foo bar")) + "a")))
                        _ = v22
                    }
                default:
                    if (!false) {
                        v4 += v4
                        v4 = v4
                    } else {
                        if false {
                            if (42 <= 4294967295) {
                                v23 := GenS4{F0: v4, F1: 100}
                                _ = v23
                            }
                        }
                        for fi24 := 0; fi24 < 16; fi24++ {
                            v4 += (v4 >> ((255) & 63))
                            _ = fi24
                        }
                        switch (-(fi18)) {
                        case 2147483647:
                            v25 := GenS1{F0: -537, F1: 1000, F2: -1}
                            _ = v25
                        case -2147483648:
                            for fi26 := 0; fi26 < 10; fi26++ {
                                var v27 float64 = float64(1.5)
                                _ = fi26
                                _ = v27
                            }
                        default:
                            fmt.Println(fi18, v4)
                        }
                    }
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    case "foo bar":
        var v28 int32 = 432
        _ = v28
    }
    var v29 int = 9223372036854775807
    for fi30 := 0; fi30 < 18; fi30++ {
        v4 -= v4
        _ = fi30
    }
    if true {
        var v31 string = (((("hello" + "hello") + "foo bar") + "\n\t") + ("hello" + "a"))
        _ = v31
    }
    var v32 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v9
}
