package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint16; F1 uint32; F2 int }
type GenS5 struct { F0 bool; F1 uint16; F2 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{("a" + ("\n\t" + ("foo bar" + "\n\t"))), "", "hello", "\n\t"}
    v3 := []uint32{2147483648}
    var v4 int64 = 42
    v5 := []uint32{784, 391, 144, 2147483648}
    v6 := GenS1{F0: (!(false && false)), F1: (v4 >> ((0) & 63))}
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 int32 = 42
        v9 := GenS2{F0: 0}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        if false {
            var v11 uint32 = 2147483648
            var v12 int16 = 0
            var v13 int64 = (v4 ^ -507)
            _ = v11
            _ = v12
            _ = v13
        } else {
            for fi14 := 0; fi14 < 20; fi14++ {
                v1[(fi14 / ((32767) | 1))] = ("\n\t" + (("hello" + (((("foo bar" + "foo bar") + "foo bar") + "\n\t") + "")) + (("hello" + "\n\t") + ("\n\t" + "a"))))
                if false {
                    v4 = v4
                    var v15 float64 = -2.25
                    _ = v15
                }
                _ = fi14
            }
        }
        var v16 int64 = v4
        switch 217 {
        case -473:
            if ("foo bar" <= "\n\t") {
                v17 := GenS2{F0: 32767}
                v18 := GenS3{F0: -1, F1: ("foo bar" >= "")}
                _ = v17
                _ = v18
            } else {
                if false {
                    v19 := GenS3{F0: (fi10 / ((32768) | 1)), F1: ("foo bar" <= "\n\t")}
                    var v20 uint = 65535
                    fmt.Println(v16, v16)
                    _ = v19
                    _ = v20
                } else {
                    var v21 string = (("" + "foo bar") + "hello")
                    _ = len(v2)
                    _ = v21
                }
                for fi22 := 0; fi22 < 16; fi22++ {
                    v4 = (v4 + v16)
                    _ = fi22
                }
                fmt.Println(v4, v16)
            }
        case 181:
            var v23 uint64 = 42
            _ = v23
        case 342:
            v4 += (v16 * (^(v4)))
        default:
            v4 += v16
        }
        _ = fi10
        _ = v16
    }
    v4 = v4
    v24 := GenS1{F0: false, F1: -1}
    var v25 int = -2147483648
    if false {
        v26 := GenS2{F0: 32767}
        v4 += -771
        v27 := GenS2{F0: 584}
        _ = v26
        _ = v27
    } else {
        var v28 int64 = 651
        for fi29 := 0; fi29 < 6; fi29++ {
            for fi30 := 0; fi30 < 19; fi30++ {
                if (1 != 4294967295) {
                    var v31 int64 = (v28 / (((v28 - v4)) | 1))
                    v25 -= v25
                    v4 = (v31 & v4)
                    _ = v31
                }
                v32 := GenS3{F0: v25, F1: true}
                var v33 int16 = 437
                var v34 uint16 = 1
                _ = fi30
                _ = v32
                _ = v33
                _ = v34
            }
            var v35 uint32 = 65535
            fmt.Println(v4)
            _ = fi29
            _ = v35
        }
        if false {
            var v36 int8 = -123
            fmt.Println(v36, v36, v4)
            v37 := GenS3{F0: (-(v25)), F1: false}
            _ = v36
            _ = v37
        } else {
            fmt.Println(v4)
            if true {
                var v38 int64 = v28
                v25 = (v25 % (((v25 * v25)) | 1))
                _ = v38
            } else {
                v28 = v4
            }
            switch "\n\t" {
            case "":
                fmt.Println(v25)
            case "hello":
                var v39 int = -201
                _ = v39
            default:
                if (2147483647 > 471) {
                    fmt.Println(v25, v4, v25)
                    v40 := GenS1{F0: (42 != 4294967295), F1: (^(v4))}
                    _ = v40
                }
            }
        }
        _ = v28
    }
    v41 := GenS2{F0: 754}
    for fi42 := 0; fi42 < 6; fi42++ {
        if false {
            var v43 float32 = ((3.14159 * (float32(((1e10 + 1.5) + (0.1 * 3.14159))) + 1.5)) * 3.14159)
            for fi44 := 0; fi44 < 17; fi44++ {
                v4 = 152
                v45 := GenS1{F0: false, F1: (v4 % ((v4) | 1))}
                _ = fi44
                _ = v45
            }
            var v46 int = 9223372036854775807
            _ = v43
            _ = v46
        }
        switch "hello" {
        case "hello":
            v4 = 42
        case "\n\t":
            var v47 float32 = 1e10
            _ = v47
        case "":
            v48 := GenS4{F0: 708, F1: 4294967295, F2: (fi42 & (fi42 | v25))}
            _ = v48
        default:
            switch (fi42 / (((-(v25))) | 1)) {
            case 0:
                fmt.Println(v25, v4)
            case -1:
                switch ("" + "a") {
                case "\n\t":
                    var v49 float64 = 3.14159
                    _ = v49
                }
            default:
                v50 := GenS5{F0: true, F1: 0, F2: 969}
                _ = v50
            }
        }
        _ = fi42
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
