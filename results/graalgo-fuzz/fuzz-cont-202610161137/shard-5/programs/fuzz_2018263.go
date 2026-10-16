package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int; F1 int64 }

func main() {
    var v1 bool = false
    var v2 uint = 42
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]string)
    v5 := []string{("" + "\n\t")}
    var v6 int8 = int8(v2)
    var v7 uint = v2
    v8 := GenS1{F0: 1, F1: v6, F2: -1}
    switch (("" + ((("\n\t" + ("hello" + "")) + (("hello" + "\n\t") + (("" + ("" + "")) + (("hello" + "hello") + "")))) + ("\n\t" + ("hello" + ("hello" + ("foo bar" + ("hello" + ""))))))) + "a") {
    case "\n\t":
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint32 = 475
            _ = fi9
            _ = v10
        }
    }
    var v11 float32 = 1.5
    switch 162 {
    case -976:
        v12 := GenS2{F0: 27, F1: 127}
        _ = v12
    case -862:
        var v13 uint = (v2 ^ 255)
        _ = v13
    }
    switch "hello" {
    case "\n\t":
        if (((((("hello" + "\n\t") + "a") > "") && ((65535 == 42) || (!true))) || v1) && v1) {
            if (42 <= -345) {
                var v14 bool = (42 <= 9223372036854775807)
                switch "foo bar" {
                case "hello":
                    v2 += 65535
                case "a":
                    var v15 int = 117
                    _ = v15
                case "foo bar":
                    var v16 uint = v2
                    _ = v16
                default:
                    v11 = (1e10 * v11)
                }
                _ = v14
            } else {
                v11 = ((float32((1e10 * (3.14159 / 1e10))) * ((v11 / -2.25) + (0.1 - v11))) * 0.0)
                var v17 int8 = int8(v2)
                v11 += ((0.1 / (v11 / (v11 * ((0.0 - (v11 / 1.0)) + (1e10 + v11))))) + (v11 - (v11 - -2.25)))
                _ = v17
            }
            if true {
                fmt.Println(v7, v3, v3)
                if len(v5) > 0 {
                    v5[0] = "foo bar"
                }
            } else {
                var v18 bool = (false && (1.5 != float64(v11)))
                v5 = append(v5, "a")
                v19 := GenS2{F0: 255, F1: 2}
                _ = v18
                _ = v19
            }
        } else {
            v4[""] = "foo bar"
            switch -458 {
            case 0:
                _ = len(v5)
            case -1:
                v20 := GenS1{F0: 2, F1: v6, F2: (v6 << ((v7) & 63))}
                _ = v20
            default:
                v4[""] = (("\n\t" + (("foo bar" + "\n\t") + ("hello" + "foo bar"))) + "foo bar")
            }
        }
    case "foo bar":
        v11 = 3.14159
    case "hello":
        var v21 uint64 = v3
        _ = v21
    default:
        for fi22 := 0; fi22 < 13; fi22++ {
            if false {
                var v23 int = (fi22 + fi22)
                v7 += v2
                _ = v23
            } else {
                v2 += (v2 << ((v2) & 63))
            }
            for fi24 := 0; fi24 < 15; fi24++ {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                var v25 int64 = 334
                fmt.Println(v11)
                _ = fi24
                _ = v25
            }
            var v26 int32 = -2147483648
            _ = fi22
            _ = v26
        }
    }
    var v27 uint8 = 2
    switch "" {
    case "\n\t":
        fmt.Println(v27, v7, v11)
    default:
        if (false && v1) {
            v28 := GenS2{F0: v27, F1: v27}
            var v29 int32 = -56
            _ = v28
            _ = v29
        } else {
            v27 += (v27 / ((0) | 1))
            for fi30 := 0; fi30 < 19; fi30++ {
                v2 = 255
                fmt.Println(fi30)
                _ = fi30
            }
        }
    }
    var v31 int32 = -1
    delete(v4, "hello")
    _ = len(v4)
    v32 := GenS3{F0: 32767, F1: 557}
    var v33 uint64 = (v3 << ((592) & 63))
    v3 -= 18446744073709551615
    v3 = 4294967295
    for fi34 := 0; fi34 < 17; fi34++ {
        if (("foo bar" + "foo bar") == "foo bar") {
            switch "foo bar" {
            case "":
                _ = len(v5)
            case "\n\t":
                _ = len(v5)
            }
        } else {
            if true {
                fmt.Println(v2, fi34)
            }
        }
        _ = fi34
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
