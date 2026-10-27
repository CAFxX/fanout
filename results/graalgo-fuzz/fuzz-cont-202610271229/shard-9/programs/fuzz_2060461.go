package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint64 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 bool }

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]bool)
    var v3 float64 = (3.14159 + (3.14159 / (0.1 / 1e10)))
    var v4 int32 = 0
    v5 := make(map[string]string)
    if false {
        v6 := GenS1{F0: -12, F1: ("\n\t" + ""), F2: 42}
        if (1 < 2) {
            switch "a" {
            case "":
                switch -531 {
                case 9223372036854775807:
                    switch ("" + (("hello" + "hello") + (("hello" + "foo bar") + ("" + (((("a" + ("foo bar" + "\n\t")) + "hello") + "a") + "foo bar"))))) {
                    case "foo bar":
                        fmt.Println(v3, v1, v3)
                    case "":
                        fmt.Println(v1, v1)
                    case "\n\t":
                        v1 += v1
                    default:
                        var v7 int16 = 959
                        _ = v7
                    }
                case 127:
                    var v8 uint8 = 22
                    _ = v8
                default:
                    if true {
                        v4 = (v4 / ((1) | 1))
                        fmt.Println(v1, v4)
                    } else {
                        fmt.Println(v3)
                    }
                }
            case "hello":
                var v9 int64 = 1
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: v1, F1: "foo bar", F2: 42}
                _ = v10
            }
        } else {
            v11 := GenS1{F0: 65, F1: (((("hello" + (("foo bar" + (("\n\t" + "") + ("hello" + "\n\t"))) + (("foo bar" + "foo bar") + "hello"))) + "\n\t") + (("" + "hello") + "\n\t")) + "\n\t"), F2: 0}
            _ = v11
        }
        _ = v6
    }
    var v12 int16 = 32767
    if (163 != 42) {
        var v13 int64 = 0
        for fi14 := 0; fi14 < 6; fi14++ {
            fmt.Println(v12, fi14, v1)
            var v15 bool = true
            _ = fi14
            _ = v15
        }
        if (v12 > (v12 | (v12 & (v12 << ((65535) & 63))))) {
            var v16 uint64 = 18446744073709551615
            _ = v16
        }
        _ = v13
    }
    if true {
        delete(v5, (("\n\t" + (("\n\t" + "foo bar") + "a")) + ("" + "a")))
        var v17 float32 = float32(1.5)
        fmt.Println(v3)
        _ = v17
    } else {
        var v18 uint = 2
        fmt.Println(v18, v1, v12)
        fmt.Println(v3, v18)
        _ = v18
    }
    v19 := GenS2{F0: v12, F1: 1000}
    v20 := GenS3{F0: false}
    if true {
        v21 := GenS3{F0: (-2.25 < v3)}
        v22 := GenS2{F0: -915, F1: (v12 >> ((42) & 63))}
        v23 := v5["\n\t"]
        _ = v21
        _ = v22
        _ = v23
    } else {
        v3 = 1e10
        switch 42 {
        case -9223372036854775808:
            if (!(true && false)) {
                v12 = (^(v12))
            } else {
                fmt.Println(v12, v4)
                var v24 int8 = v1
                _ = v24
            }
        case 0:
            fmt.Println(v12, v3)
        default:
            v12 += (v12 & (v12 - (^(v12))))
        }
        var v25 uint32 = 2
        _ = v25
    }
    var v26 float32 = (1e10 + (0.1 / float32((v3 / float64(0.1)))))
    var v27 int8 = v1
    var v28 int64 = 0
    var v29 float32 = v26
    v4 = int32(v1)
    for fi30 := 0; fi30 < 12; fi30++ {
        v5["a"] = "foo bar"
        fmt.Println(v29, v1)
        v4 -= (v4 << ((59) & 63))
        _ = fi30
    }
    v31 := v5[""]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
