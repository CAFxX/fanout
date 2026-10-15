package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int64; F1 int16; F2 int16 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 uint8, p1 string) uint32 {
    var v7 string = "a"
    p1 = v7
    _ = p0
    _ = p1
    _ = v7
    return 201
}

func main() {
    v1 := []int{}
    var v2 uint32 = 2147483648
    var v3 uint64 = 0
    var v4 uint = 387
    var v5 int32 = -60
    var v6 uint16 = 2
    v8 := GenS1{F0: 1, F1: -1, F2: (^(v2))}
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v3, v6)
        _ = fi9
    }
    var v10 bool = (true && ("hello" <= (("hello" + "") + "a")))
    v5 = v5
    if v10 {
        if v10 {
            v3 = v3
            v2 = v2
            var v11 float32 = 1e10
            _ = v11
        } else {
            var v12 int = 4294967295
            for fi13 := 0; fi13 < 19; fi13++ {
                switch -802 {
                case -838:
                    switch 32768 {
                    case -449:
                        v14 := GenS2{F0: 1, F1: 1, F2: 42}
                        _ = v14
                    case -2147483648:
                        if ((v10 || (130 == 0)) || false) {
                            for fi15 := 0; fi15 < 20; fi15++ {
                                v16 := GenS2{F0: -1, F1: 1000, F2: -32768}
                                var v17 int16 = int16(v3)
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 7; fi18++ {
                                fmt.Println(v12)
                                var v19 int = fi13
                                v2 = (v2 - (v2 + 65535))
                                _ = fi18
                                _ = v19
                            }
                        }
                    case 697:
                        if len(v1) > 0 {
                            v1[0] = -608
                        }
                    }
                default:
                    v20 := GenS1{F0: int64(v2), F1: 829, F2: uint32(v3)}
                    _ = v20
                }
                for fi21 := 0; fi21 < 5; fi21++ {
                    v22 := GenS3{F0: 0, F1: 1000}
                    v3 = v3
                    if (1e10 >= 0.0) {
                        var v23 int8 = -12
                        v2 += genFn0(255, ("" + "a"))
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi13
            }
            v6 = 32767
            _ = v12
        }
    }
    v2 += genFn0(1, ("foo bar" + "\n\t"))
    var v24 uint16 = (v6 - v6)
    for fi25 := 0; fi25 < 8; fi25++ {
        v2 = (v2 % ((genFn0(42, "foo bar")) | 1))
        _ = fi25
    }
    switch 32767 {
    case -351:
        switch "foo bar" {
        case "foo bar":
            v5 = (-(v5))
        case "hello":
            v26 := GenS2{F0: 1, F1: -32768, F2: 1}
            _ = v26
        }
    case 0:
        v27 := GenS2{F0: 77, F1: -244, F2: 42}
        _ = v27
    default:
        var v28 int = 2147483647
        _ = v28
    }
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    if false {
        var v29 int = 0
        var v30 int32 = -2147483648
        _ = v29
        _ = v30
    }
    var v31 uint32 = v2
    if v10 {
        if len(v1) > 0 {
            v1[0] = 545
        }
        v5 -= (^(v5))
    } else {
        switch int(v3) {
        case 42:
            for fi32 := 0; fi32 < 3; fi32++ {
                var v33 int64 = -1
                v34 := GenS3{F0: (v33 * (v33 * (v33 * (-(v33))))), F1: 32767}
                _ = fi32
                _ = v33
                _ = v34
            }
        default:
            var v35 int = -257
            _ = v35
        }
        switch "hello" {
        case "hello":
            for fi36 := 0; fi36 < 17; fi36++ {
                v4 -= (v4 & (^(v4)))
                v2 += (v2 * 303)
                _ = fi36
            }
        case "\n\t":
            switch "" {
            case "a":
                for fi37 := 0; fi37 < 4; fi37++ {
                    for fi38 := 0; fi38 < 15; fi38++ {
                        v10 = (!false)
                        _ = fi38
                    }
                    _ = fi37
                }
            case "hello":
                var v39 int8 = -128
                _ = v39
            default:
                var v40 uint16 = v6
                _ = v40
            }
        default:
            switch "a" {
            case "\n\t":
                var v41 uint = v4
                _ = v41
            default:
                fmt.Println(v3)
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
