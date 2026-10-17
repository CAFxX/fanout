package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 string; F1 uint32 }

func main() {
    v1 := []int{}
    var v2 uint16 = 525
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    v5 := GenS1{F0: 0, F1: 43}
    if true {
        v6 := GenS1{F0: 9223372036854775807, F1: 678}
        var v7 uint8 = 127
        _ = v6
        _ = v7
    }
    switch ("hello" + ("hello" + ("\n\t" + "foo bar"))) {
    case "foo bar":
        v8 := GenS1{F0: -469, F1: -32768}
        _ = v8
    default:
        v2 -= v2
    }
    if false {
        v2 -= (^(v2))
        if (!true) {
            var v9 int32 = -2147483648
            v1 = append(v1, 127)
            _ = v9
        }
    }
    if true {
        for fi10 := 0; fi10 < 13; fi10++ {
            v2 = (^(v2))
            _ = fi10
        }
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS2{F0: v2, F1: "", F2: (v2 * (^(v2)))}
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 int16 = 32767
                var v15 string = "foo bar"
                var v16 int16 = (^(v14))
                v17 := GenS3{F0: 681}
                v14 += v16
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi11
            _ = v12
        }
        v4["foo bar"] = 65535
    } else {
        v2 = v2
        v3[2147483648] = 42
        v18 := GenS1{F0: 42, F1: 1000}
        _ = v18
    }
    if false {
        v2 = (v2 - v2)
        v19 := GenS4{F0: 781, F1: 1}
        if (2 <= 2) {
            var v20 int16 = -1
            v21 := GenS4{F0: 767, F1: 42}
            if (!((227 != 128) && true)) {
                v2 += (v2 / ((v2) | 1))
            }
            _ = v20
            _ = v21
        } else {
            switch -9223372036854775808 {
            case 32767:
                if true {
                    if len(v1) > 0 {
                        v1[0] = -9223372036854775808
                    }
                    var v22 int8 = 100
                    _ = v22
                } else {
                    var v23 string = ("foo bar" + "\n\t")
                    switch v23 {
                    case "\n\t":
                        var v24 uint = 42
                        _ = v24
                    case "a":
                        if (2147483647 != 594) {
                            v25 := GenS5{F0: v23, F1: 925}
                            _ = v25
                        } else {
                            var v26 uint = 255
                            _ = v26
                        }
                    }
                    v23 = "a"
                    _ = v23
                }
            default:
                switch 32768 {
                case 1:
                    fmt.Println(v2)
                default:
                    var v27 uint = 4294967295
                    _ = v27
                }
            }
            for fi28 := 0; fi28 < 19; fi28++ {
                var v29 int32 = -168
                _ = fi28
                _ = v29
            }
        }
        _ = v19
    }
    v30 := GenS2{F0: (v2 + (v2 & 843)), F1: ("hello" + ("\n\t" + "hello")), F2: 700}
    v3[2] = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
