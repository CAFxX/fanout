package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 uint64 = 191
    v4 := make(map[uint32]string)
    var v5 uint16 = 892
    v6 := []string{"\n\t", (("hello" + (("\n\t" + "hello") + "foo bar")) + "\n\t"), (("\n\t" + (("a" + "\n\t") + "a")) + "")}
    v3 = 2
    for fi7 := 0; fi7 < 1; fi7++ {
        v3 += v3
        if (false && true) {
            fmt.Println(v5, fi7, fi7)
            var v8 int64 = 42
            _ = v8
        } else {
            switch (fi7 - fi7) {
            case -483:
                v9 := GenS1{F0: (^(v3)), F1: "hello"}
                _ = v9
            case -919:
                fmt.Println(v3, v3)
            case 360:
                var v10 string = "a"
                _ = v10
            default:
                var v11 uint32 = 2147483647
                _ = v11
            }
        }
        _ = fi7
    }
    var v12 int64 = -9223372036854775808
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v5, v5, v5)
    default:
        v3 = v3
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: (v3 - 365), F1: "hello"}
        fmt.Println(fi13)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS2{F0: 1, F1: (v5 % (((v5 << ((272) & 63))) | 1))}
            v18 := GenS2{F0: 61, F1: (v5 | (v5 >> ((4294967295) & 63)))}
            if false {
                v4[4294967295] = ("foo bar" + "")
                v3 -= 755
            }
            var v19 int = fi15
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        var v20 float32 = 3.14159
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 uint64 = (v3 & (^(v3)))
            switch "hello" {
            case "a":
                if len(v6) > 0 {
                    v6[0] = "a"
                }
            case "hello":
                v23 := GenS2{F0: 165, F1: v5}
                _ = v23
            case "foo bar":
                if false {
                    v24 := GenS2{F0: 1, F1: (v5 * (v5 - 65535))}
                    _ = v24
                }
            }
            var v25 uint32 = 702
            _ = fi21
            _ = v22
            _ = v25
        }
        _ = fi15
        _ = v20
    }
    var v26 float32 = (1.5 + ((1e10 / float32(-2.25)) - (-2.25 * (3.14159 - 0.1))))
    v27 := GenS1{F0: v3, F1: ("" + "foo bar")}
    switch -691 {
    case 128:
        fmt.Println(v12, v26)
    case 922:
        switch ("foo bar" + "") {
        case "a":
            fmt.Println(v3, v5)
        case "":
            _ = len(v2)
        case "foo bar":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    case -696:
        if (98 < -57) {
            for fi28 := 0; fi28 < 20; fi28++ {
                fmt.Println(v26, v5, v5)
                var v29 int64 = (v12 * (v12 >> ((1) & 63)))
                _ = fi28
                _ = v29
            }
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 float32 = float32(0.0)
                _ = fi30
                _ = v31
            }
            v12 = v12
        }
    }
    var v32 int32 = 2147483647
    _ = len(v6)
    fmt.Println(v3, v12)
    _ = len(v6)
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 int32 = v32
        v32 += (^(v32))
        v34 = v32
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
