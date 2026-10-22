package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int16; F1 uint16; F2 string }

func main() {
    var v1 int32 = -1
    v2 := make(map[string]bool)
    var v3 uint = 2
    v4 := make(map[int]int)
    var v5 float64 = 3.14159
    v6 := make(map[uint32]uint32)
    var v7 int = -9223372036854775808
    var v8 uint8 = 2
    _ = len(v2)
    if (547 <= 0) {
        if false {
            fmt.Println(v5, v1, v5)
            var v9 float32 = 1e10
            v9 -= v9
            _ = v9
        } else {
            if (65535 == 255) {
                var v10 int8 = 111
                _ = v10
            } else {
                var v11 int64 = -9223372036854775808
                _ = v11
            }
        }
    }
    switch ((("\n\t" + "\n\t") + ("" + "")) + ("hello" + (((((("" + "") + "a") + "") + "foo bar") + "foo bar") + ("\n\t" + ("a" + "hello"))))) {
    case "a":
        v5 -= v5
    case "hello":
        v12 := GenS1{F0: (v1 | 1), F1: "\n\t"}
        _ = v12
    default:
        var v13 int64 = int64(v7)
        _ = v13
    }
    v14 := []int{v7, int(v8), v7, -9223372036854775808}
    switch (v7 - (-(v7))) {
    case 127:
        if (true || true) {
            if len(v14) > 0 {
                v14[0] = (v7 % (((-(v7))) | 1))
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := GenS1{F0: int32(v8), F1: (("" + (("" + "\n\t") + ("hello" + "a"))) + "")}
                for fi17 := 0; fi17 < 4; fi17++ {
                    if true {
                        var v18 float64 = float64(0.0)
                        _ = v18
                    }
                    _ = fi17
                }
                if false {
                    v8 += (v8 % (((v8 - uint8(v3))) | 1))
                } else {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        switch "\n\t" {
                        case "foo bar":
                            v20 := GenS1{F0: (v1 >> (((v3 - 1)) & 63)), F1: (("a" + "") + "\n\t")}
                            _ = v20
                        case "":
                            fmt.Println(v8, v7)
                        case "hello":
                            v14 = append(v14, fi15)
                        default:
                            var v21 float64 = float64(0.1)
                            _ = v21
                        }
                        _ = fi19
                    }
                    switch (("a" + "hello") + "a") {
                    case "a":
                        var v22 int64 = 1
                        _ = v22
                    default:
                        var v23 string = ("a" + "foo bar")
                        _ = v23
                    }
                    v3 = 1
                }
                _ = fi15
                _ = v16
            }
        } else {
            if ((0.0 == -2.25) || (65535 >= 736)) {
                switch v7 {
                case 1:
                    for fi24 := 0; fi24 < 11; fi24++ {
                        _ = len(v6)
                        fmt.Println(v7, v5, v3)
                        if ("a" != "\n\t") {
                            fmt.Println(fi24, v5, v7)
                        }
                        _ = fi24
                    }
                case 706:
                    v8 = (v8 << ((v3) & 63))
                case 841:
                    v25 := GenS2{F0: "hello"}
                    _ = v25
                default:
                    v26 := GenS3{F0: 0, F1: 1}
                    _ = v26
                }
            }
            v27 := GenS4{F0: 2, F1: 970}
            v28 := v2["\n\t"]
            _ = v27
            _ = v28
        }
    case 4294967295:
        fmt.Println(v7, v1)
    case -625:
        for fi29 := 0; fi29 < 17; fi29++ {
            v30 := GenS2{F0: ("foo bar" + "hello")}
            v3 = (v3 ^ (v3 * (v3 / (((^(v3))) | 1))))
            var v31 float64 = v5
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    switch (v7 + v7) {
    case 9223372036854775807:
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    case 53:
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 uint32 = 2
            v8 += v8
            _ = fi32
            _ = v33
        }
    case 4294967295:
        fmt.Println(v8)
    default:
        v3 -= (v3 % ((v3) | 1))
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 string = ("foo bar" + ((("a" + ("" + "\n\t")) + ((("hello" + "") + "\n\t") + "hello")) + "foo bar"))
            v1 = int32(fi34)
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    var v37 uint8 = v8
    _ = len(v14)
    if true {
        switch (("a" + ("" + "a")) + ("foo bar" + ("foo bar" + "a"))) {
        case "":
            switch 4294967295 {
            case 32768:
                switch "" {
                case "\n\t":
                    v38 := GenS5{F0: -658, F1: 2, F2: "foo bar"}
                    _ = v38
                }
            case 128:
                v8 -= 236
            case -9223372036854775808:
                var v39 int8 = -1
                _ = v39
            }
        default:
            _ = len(v6)
        }
    }
    if true {
        fmt.Println(v8, v1, v37)
        v1 = (v1 >> ((v3) & 63))
    }
    var v40 string = "foo bar"
    var v41 uint64 = 0
    fmt.Println(v1, v8, v5)
    for fi42 := 0; fi42 < 16; fi42++ {
        v43 := GenS2{F0: v40}
        v37 = 200
        v41 = (v41 | (v41 & v41))
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
