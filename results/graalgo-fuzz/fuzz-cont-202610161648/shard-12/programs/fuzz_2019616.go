package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 int }
type GenS3 struct { F0 bool; F1 uint16; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int16; F2 uint16 }
type GenS6 struct { F0 int64; F1 int }

func main() {
    v1 := []string{("" + "hello"), ("a" + "\n\t"), "hello"}
    var v2 int16 = -32768
    v3 := make(map[string]uint32)
    v2 -= (v2 - v2)
    v4 := GenS1{F0: 128, F1: 65535}
    if false {
        switch 128 {
        case 127:
            var v5 int = 42
            _ = v5
        default:
            v2 -= -32768
        }
        fmt.Println(v2)
    }
    for fi6 := 0; fi6 < 4; fi6++ {
        v2 -= int16(fi6)
        _ = fi6
    }
    v3[""] = 0
    switch -846 {
    case -2147483648:
        switch "a" {
        case "\n\t":
            var v7 uint64 = 2
            _ = v7
        }
    case 9223372036854775807:
        if (!true) {
            v8 := GenS2{F0: 74, F1: true, F2: 42}
            var v9 bool = false
            _ = v8
            _ = v9
        }
    case 1:
        v10 := GenS3{F0: true, F1: 0, F2: 42}
        _ = v10
    }
    if (true && false) {
        fmt.Println(v2)
    }
    if false {
        if true {
            switch "hello" {
            case "foo bar":
                v2 += v2
            default:
                if true {
                    v2 -= (-(v2))
                    var v11 float32 = -2.25
                    _ = v11
                }
            }
        }
        var v12 int8 = 127
        _ = v12
    } else {
        fmt.Println(v2)
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v2 = int16(fi13)
        if false {
            var v14 string = "foo bar"
            _ = v14
        } else {
            fmt.Println(fi13, v2)
        }
        if true {
            if ("hello" >= ("" + (("hello" + "\n\t") + "hello"))) {
                fmt.Println(v2, fi13)
                var v15 uint64 = 18446744073709551615
                _ = len(v3)
                _ = v15
            } else {
                _ = len(v3)
            }
            if (false || (!false)) {
                var v16 uint16 = 149
                _ = v16
            } else {
                for fi17 := 0; fi17 < 13; fi17++ {
                    v2 -= -626
                    _ = fi17
                }
                switch (("hello" + ("\n\t" + "foo bar")) + ((("a" + "a") + "\n\t") + ("hello" + "hello"))) {
                case "a":
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                }
                fmt.Println(fi13, v2)
            }
        } else {
            for fi18 := 0; fi18 < 17; fi18++ {
                switch fi13 {
                case 4294967295:
                    if ((0 != 0) && false) {
                        v2 -= (-(v2))
                    } else {
                        v19 := GenS4{F0: 368}
                        var v20 uint64 = 114
                        switch ("" + "") {
                        case "\n\t":
                            _ = len(v1)
                        }
                        _ = v19
                        _ = v20
                    }
                case -293:
                    if (-9223372036854775808 <= -9223372036854775808) {
                        v21 := GenS5{F0: -994, F1: (v2 * 32767), F2: 615}
                        _ = len(v1)
                        v2 += (v2 ^ v2)
                        _ = v21
                    } else {
                        v22 := GenS6{F0: 1, F1: (fi13 - (^(fi18)))}
                        v2 = (v2 - v2)
                        _ = v22
                    }
                }
                _ = fi18
            }
        }
        _ = fi13
    }
    if (!(42 < 32768)) {
        var v23 uint = 141
        v2 -= 0
        v3["a"] = 335
        _ = v23
    } else {
        v2 = v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
