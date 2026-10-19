package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int32; F1 int32; F2 uint32 }

func main() {
    v1 := []uint32{1, 42, 2147483648, 612}
    v2 := []string{"foo bar"}
    v3 := make(map[string]bool)
    v3["hello"] = false
    var v4 int16 = 1000
    if true {
        var v5 uint = 18446744073709551615
        _ = v5
    }
    v4 = v4
    var v6 bool = false
    v6 = (2147483648 == 2)
    v4 = (v4 - v4)
    if v6 {
        switch ("hello" + ("\n\t" + "\n\t")) {
        case "\n\t":
            switch 32767 {
            case 32768:
                fmt.Println(v6, v4)
            case -2147483648:
                var v7 uint64 = 18446744073709551615
                _ = v7
            case -9223372036854775808:
                var v8 uint16 = 0
                _ = v8
            default:
                var v9 float32 = 1.5
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS1{F0: 0, F1: 42}
                _ = fi10
                _ = v11
            }
        }
    }
    var v12 int16 = (v4 % ((v4) | 1))
    v13 := GenS1{F0: 1, F1: 0}
    var v14 int32 = -872
    var v15 float32 = 1e10
    v16 := GenS2{F0: "a", F1: 4294967295, F2: 128}
    if (!(!v6)) {
        v4 -= v12
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        v6 = (!(-41 != 112))
        if (127 == -128) {
            v3[""] = true
        }
        _ = fi17
    }
    switch 2147483647 {
    case 32767:
        switch 128 {
        case 42:
            var v18 float64 = 1e10
            _ = v18
        case 1:
            if len(v1) > 0 {
                v1[0] = 419
            }
        default:
            fmt.Println(v15)
        }
    case 128:
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS3{F0: -1, F1: 65535}
            for fi21 := 0; fi21 < 20; fi21++ {
                v22 := GenS4{F0: (^(v12)), F1: 255, F2: 1}
                var v23 uint16 = 969
                if (577 == v23) {
                    if (-1 > -1) {
                        var v24 int8 = -128
                        switch (("hello" + ("a" + ("hello" + (((("a" + "hello") + ((("\n\t" + "hello") + "") + ("foo bar" + ("\n\t" + (("a" + "foo bar") + "foo bar"))))) + "foo bar") + ("" + ""))))) + "\n\t") {
                        case "":
                            v15 += float32(1.5)
                        case "foo bar":
                            var v25 int32 = (v14 % (((v14 | (-(v14)))) | 1))
                            _ = v25
                        case "\n\t":
                            v26 := GenS2{F0: "a", F1: 2, F2: 125}
                            _ = v26
                        }
                        _ = v24
                    } else {
                        v27 := v3["foo bar"]
                        var v28 int64 = int64(v14)
                        fmt.Println(v6, v4, v12)
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(v14, v4)
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            v29 := GenS2{F0: "foo bar", F1: 1, F2: 42}
            _ = fi19
            _ = v20
            _ = v29
        }
    default:
        v1 = append(v1, 546)
    }
    switch "a" {
    case "hello":
        var v30 uint = 2
        _ = v30
    case "":
        var v31 int16 = (v12 * v4)
        _ = v31
    case "a":
        v4 -= (v4 * (^(v12)))
    default:
        for fi32 := 0; fi32 < 10; fi32++ {
            v33 := GenS5{F0: (v14 + 1), F1: 2147483647, F2: 1}
            v34 := GenS4{F0: v12, F1: 127, F2: 0}
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
