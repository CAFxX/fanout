package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int8; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 0
    var v3 uint8 = v2
    v4 := make(map[string]string)
    var v5 uint64 = 1
    v4[("a" + ("" + ("" + ("foo bar" + "hello"))))] = ("hello" + "")
    if true {
        v3 -= (v3 * v3)
        v3 = v2
    } else {
        v6 := GenS1{F0: 4294967295}
        if false {
            var v7 uint8 = v2
            if true {
                for fi8 := 0; fi8 < 20; fi8++ {
                    v2 += (v3 / ((v7) | 1))
                    fmt.Println(v5)
                    var v9 int = fi8
                    _ = fi8
                    _ = v9
                }
                var v10 string = "hello"
                _ = v10
            }
            v5 = v5
            _ = v7
        }
        v2 = 127
        _ = v6
    }
    v11 := []int{42}
    if false {
        fmt.Println(v2)
        fmt.Println(v2, v2, v5)
        if (1 >= 0) {
            v2 += (v3 * (v2 | 128))
        } else {
            switch -965 {
            case 128:
                v2 = (v2 + 2)
            }
            var v12 int = 9223372036854775807
            switch (v12 + v12) {
            case 455:
                var v13 int = v12
                _ = v13
            case 0:
                var v14 int = 641
                _ = v14
            case 2147483647:
                v12 = (v12 ^ v12)
            }
            _ = v12
        }
    } else {
        var v15 uint64 = (v5 >> ((763) & 63))
        if (((0 < 100) && (false && false)) && ((v15 ^ v5) > 18446744073709551615)) {
            if false {
                v15 += v5
                for fi16 := 0; fi16 < 20; fi16++ {
                    if (-2147483648 >= -925) {
                        v17 := GenS1{F0: 18446744073709551615}
                        delete(v1, -814)
                        _ = v17
                    } else {
                        v1[2147483647] = 65535
                    }
                    if false {
                        for fi18 := 0; fi18 < 7; fi18++ {
                            var v19 uint16 = 65535
                            v5 = (v5 / ((42) | 1))
                            v20 := GenS1{F0: 255}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        fmt.Println(v5, v5)
                    } else {
                        for fi21 := 0; fi21 < 16; fi21++ {
                            v3 -= v3
                            v5 = 42
                            if (-758 >= 899) {
                                fmt.Println(v5, fi21)
                                _ = len(v11)
                                if (!(false || ((32767 == 1) && true))) {
                                    var v22 int8 = 0
                                    _ = v22
                                } else {
                                    var v23 uint32 = 1
                                    _ = v23
                                }
                            }
                            _ = fi21
                        }
                        fmt.Println(v5, v2, v15)
                    }
                    for fi24 := 0; fi24 < 13; fi24++ {
                        switch fi24 {
                        case 32768:
                            _ = len(v4)
                        default:
                            v25 := GenS2{F0: 42, F1: (^(v2))}
                            _ = v25
                        }
                        fmt.Println(v15)
                        _ = fi24
                    }
                    _ = fi16
                }
                for fi26 := 0; fi26 < 10; fi26++ {
                    var v27 float32 = 1e10
                    v28 := GenS1{F0: 978}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            } else {
                var v29 uint32 = 581
                switch "hello" {
                case "\n\t":
                    v15 = (v5 << ((uint(v29)) & 63))
                case "hello":
                    var v30 string = (("foo bar" + ("" + "foo bar")) + "hello")
                    _ = v30
                case "a":
                    _ = len(v11)
                default:
                    v2 = (^(v3))
                }
                _ = v29
            }
        }
        _ = v15
    }
    fmt.Println(v5, v2)
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 uint = 255
        _ = fi31
        _ = v32
    }
    switch "" {
    case "foo bar":
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    case "\n\t":
        if (0 == 1) {
            v33 := GenS1{F0: 18446744073709551615}
            _ = v33
        } else {
            v34 := GenS3{F0: 0, F1: true, F2: 65535}
            _ = v34
        }
    default:
        if (255 >= 32767) {
            switch "foo bar" {
            case "":
                _ = len(v4)
            }
            var v35 uint32 = 2
            _ = v35
        }
    }
    if false {
        var v36 int = 2147483647
        for fi37 := 0; fi37 < 16; fi37++ {
            v38 := GenS2{F0: (v2 >> ((1) & 63)), F1: (v2 >> ((18446744073709551615) & 63))}
            var v39 int32 = 643
            var v40 int16 = -32768
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
        _ = v36
    } else {
        v3 -= 254
        v5 = 498
        v5 = 55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
