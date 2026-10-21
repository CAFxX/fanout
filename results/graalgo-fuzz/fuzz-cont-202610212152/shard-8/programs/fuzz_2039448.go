package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int64; F2 string }

func main() {
    var v1 uint32 = 2147483647
    var v2 int64 = 322
    v3 := []string{((("foo bar" + "\n\t") + ("a" + "a")) + ""), (("hello" + ("\n\t" + "foo bar")) + "hello"), (("a" + ("a" + ("foo bar" + "foo bar"))) + "a")}
    var v4 int8 = 1
    v5 := make(map[uint32]bool)
    if (0.0 > float32(1.5)) {
        switch "" {
        case "\n\t":
            v1 = v1
        case "":
            switch 127 {
            case -242:
                var v6 uint16 = 405
                _ = v6
            }
        case "a":
            v7 := GenS1{F0: 42, F1: (v4 << ((1) & 63))}
            _ = v7
        default:
            v4 += (v4 * 127)
        }
        v8 := GenS2{F0: false}
        v9 := v5[v1]
        _ = v8
        _ = v9
    } else {
        var v10 int32 = 864
        var v11 int = 127
        v12 := GenS1{F0: v10, F1: (v4 % ((1) | 1))}
        _ = v10
        _ = v11
        _ = v12
    }
    fmt.Println(v4, v1)
    var v13 int16 = -32768
    if (v2 < v2) {
        var v14 int8 = (v4 + v4)
        var v15 string = "hello"
        _ = v14
        _ = v15
    }
    v2 = -709
    var v16 bool = (int32(v2) <= 0)
    v1 = 671
    if v16 {
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 bool = (!(int8(fi17) < (v4 % ((42) | 1))))
            v1 = (^(v1))
            _ = fi17
            _ = v18
        }
        v5[42] = false
    } else {
        delete(v5, v1)
        v13 += 42
        _ = len(v3)
    }
    var v19 uint = 65535
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 5; fi21++ {
            v22 := GenS1{F0: 1, F1: 0}
            var v23 float32 = (1e10 / ((3.14159 - (0.0 * (3.14159 - float32(1.5)))) / 0.1))
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    var v24 int32 = 971
    v25 := GenS3{F0: 2, F1: v2, F2: "a"}
    _ = len(v5)
    if (("a" + (("" + ("foo bar" + "hello")) + "\n\t")) >= (("a" + "hello") + "")) {
        for fi26 := 0; fi26 < 13; fi26++ {
            _ = len(v3)
            v2 = -9223372036854775808
            _ = fi26
        }
    } else {
        var v27 uint64 = 4294967295
        _ = v27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := v5[(v1 | 2147483648)]
        var v30 int64 = v2
        _ = fi28
        _ = v29
        _ = v30
    }
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
