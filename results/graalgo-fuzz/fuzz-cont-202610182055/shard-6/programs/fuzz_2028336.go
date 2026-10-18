package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 string }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 uint = 0
    v5 := make(map[string]string)
    v6 := []uint32{1, 2147483647, 42}
    for fi7 := 0; fi7 < 1; fi7++ {
        if (!(127 != -128)) {
            if (((fi7 - fi7) < 32768) || false) {
                fmt.Println(v4, fi7)
                v5["hello"] = ("a" + "foo bar")
            }
            var v8 float32 = (0.0 / 3.14159)
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(v8, v4, v8)
                _ = fi9
            }
            _ = v8
        } else {
            v4 -= 1
            v10 := GenS1{F0: -1}
            var v11 int32 = 1
            _ = v10
            _ = v11
        }
        fmt.Println(v1, v4, v1)
        _ = fi7
    }
    var v12 uint32 = 0
    var v13 int16 = -1
    var v14 int64 = (v1 & v1)
    v15 := GenS1{F0: (v1 | v1)}
    if (int(v12) <= int(v4)) {
        if ((v13 ^ (^(v13))) != 359) {
            var v16 string = "a"
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 string = v16
                if (65535 <= 65535) {
                    v19 := GenS2{F0: false, F1: v1, F2: v16}
                    if false {
                        var v20 float32 = 1e10
                        _ = v20
                    }
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    } else {
        var v21 uint = v4
        _ = v21
    }
    if ((-738 >= 127) || false) {
        fmt.Println(v14)
        v4 = (v4 >> (((v4 | v4)) & 63))
        v22 := GenS3{F0: 218}
        _ = v22
    } else {
        v23 := GenS1{F0: v14}
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 string = ("hello" + ("" + "\n\t"))
        _ = fi24
        _ = v25
    }
    fmt.Println(v12, v12, v14)
    for fi26 := 0; fi26 < 20; fi26++ {
        fmt.Println(v4, v14)
        switch ("foo bar" + "") {
        case "foo bar":
            var v27 int64 = (v14 | 9223372036854775807)
            _ = v27
        case "":
            v1 -= (v14 * 9223372036854775807)
        }
        if (true || true) {
            var v28 float64 = (1e10 * 1.5)
            for fi29 := 0; fi29 < 3; fi29++ {
                fmt.Println(v14, v1, v12)
                var v30 uint64 = 2
                v13 = (v13 & v13)
                _ = fi29
                _ = v30
            }
            _ = v28
        } else {
            v12 -= 4294967295
        }
        _ = fi26
    }
    delete(v3, -9223372036854775808)
    v13 = v13
    switch 626 {
    case 42:
        switch 4294967295 {
        case 32768:
            switch 1 {
            case 128:
                switch "" {
                case "":
                    delete(v2, ("foo bar" + "hello"))
                case "foo bar":
                    if (109 <= 2) {
                        v14 += -1
                        v12 += 42
                    }
                default:
                    var v31 float32 = -2.25
                    _ = v31
                }
            }
        default:
            var v32 float64 = ((0.1 / 1.5) - ((1.5 / (-2.25 + (0.0 * (1e10 + (-2.25 + 3.14159))))) / 1.5))
            _ = v32
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
