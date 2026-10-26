package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int; F1 int; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 float32 = 1.5
    fmt.Println(v3)
    fmt.Println(v3)
    var v4 uint = 511
    v5 := GenS1{F0: 255, F1: 1, F2: 4294967295}
    switch "foo bar" {
    case "a":
        switch "" {
        case "foo bar":
            var v6 int8 = -1
            _ = v6
        }
    case "":
        if (-9223372036854775808 == 15) {
            v7 := []int{9223372036854775807}
            v8 := v2["foo bar"]
            if false {
                if ((!(32767 == 42)) || (!((!v8) && (!v8)))) {
                    var v9 int8 = -128
                    _ = v9
                } else {
                    v10 := GenS1{F0: 32767, F1: -9223372036854775808, F2: 42}
                    var v11 int32 = 2147483647
                    _ = v10
                    _ = v11
                }
            } else {
                delete(v2, "hello")
            }
            _ = v7
            _ = v8
        }
    case "foo bar":
        fmt.Println(v3, v3)
    default:
        if (("foo bar" + "hello") != "\n\t") {
            fmt.Println(v4, v3)
            var v12 string = ("" + ("hello" + "foo bar"))
            _ = v12
        } else {
            for fi13 := 0; fi13 < 10; fi13++ {
                var v14 float32 = v3
                v15 := GenS1{F0: 1, F1: 9223372036854775807, F2: uint64(v4)}
                var v16 int16 = -527
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            v17 := GenS2{F0: 737, F1: -2147483648, F2: 779}
            _ = v17
        }
    }
    v2[("" + "hello")] = false
    if true {
        if false {
            fmt.Println(v4, v3)
            v3 = v3
            switch "\n\t" {
            case "\n\t":
                switch 0 {
                case 42:
                    v18 := GenS3{F0: 127}
                    _ = v18
                default:
                    v4 = (v4 | (v4 + 42))
                }
            case "a":
                fmt.Println(v3, v4)
            case "hello":
                v19 := []int{-598, 1, 882}
                _ = v19
            default:
                v20 := GenS3{F0: uint8(v4)}
                _ = v20
            }
        }
        if (!false) {
            var v21 float32 = v3
            if true {
                if (true || true) {
                    var v22 int16 = 32767
                    v22 = v22
                    fmt.Println(v4, v4, v22)
                    _ = v22
                }
                var v23 int16 = -247
                for fi24 := 0; fi24 < 17; fi24++ {
                    for fi25 := 0; fi25 < 9; fi25++ {
                        v3 += v21
                        v26 := []int{(^(fi25)), fi25, (fi24 + -37)}
                        _ = fi25
                        _ = v26
                    }
                    _ = fi24
                }
                _ = v23
            }
            if (200 != 165) {
                v27 := []int{-9223372036854775808, -2147483648}
                _ = v27
            }
            _ = v21
        } else {
            if false {
                v3 += 1.5
            } else {
                fmt.Println(v3, v4)
                fmt.Println(v4, v4)
                v28 := GenS2{F0: -297, F1: 2147483647, F2: 2}
                _ = v28
            }
            var v29 string = ("hello" + ("hello" + "foo bar"))
            var v30 float32 = (1.5 / 1.0)
            _ = v29
            _ = v30
        }
    } else {
        v31 := GenS4{F0: 42}
        var v32 uint = (v4 % ((v4) | 1))
        _ = v31
        _ = v32
    }
    v4 += v4
    v4 = (v4 >> ((468) & 63))
    v33 := []int{}
    var v34 string = (((("foo bar" + "") + "hello") + "hello") + (("foo bar" + "\n\t") + "foo bar"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
