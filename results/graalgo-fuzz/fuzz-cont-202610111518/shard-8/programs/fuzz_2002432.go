package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
type GenS3 struct { F0 uint; F1 uint64; F2 int16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 uint16 = 2
    v2 := []uint32{1}
    var v3 float32 = (3.14159 - (0.0 / (0.1 / ((-2.25 - 0.0) / -2.25))))
    var v4 int64 = 102
    switch "\n\t" {
    case "":
        if true {
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
            if true {
                if (!(true || (((-2.25 / -2.25) - 1e10) == (-2.25 - 1.5)))) {
                    v5 := GenS1{F0: "", F1: -2147483648, F2: 114}
                    v6 := GenS1{F0: (("" + "a") + "hello"), F1: 2147483647, F2: 0}
                    _ = v5
                    _ = v6
                }
            } else {
                v1 += (v1 / ((42) | 1))
                v7 := GenS1{F0: "hello", F1: int32(v4), F2: 4294967295}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v7
            }
            switch ("foo bar" + "hello") {
            case "":
                if false {
                    switch "hello" {
                    case "a":
                        var v8 int16 = 1
                        _ = v8
                    default:
                        _ = len(v2)
                    }
                }
            case "a":
                for fi9 := 0; fi9 < 4; fi9++ {
                    var v10 uint = 42
                    _ = fi9
                    _ = v10
                }
            case "foo bar":
                v1 -= (v1 - (^(v1)))
            default:
                switch "a" {
                case "\n\t":
                    v4 = (v4 ^ -1)
                }
            }
        } else {
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := GenS2{F0: "foo bar", F1: 0, F2: 995}
                v13 := GenS3{F0: 18446744073709551615, F1: 2, F2: -787}
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    default:
        if len(v2) > 0 {
            v2[0] = 0
        }
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        if ((18446744073709551615 >= 1) || true) {
            if (((1e10 * (0.1 - 0.1)) < ((-2.25 * 0.1) * -2.25)) || true) {
                var v15 uint8 = 208
                _ = v15
            } else {
                fmt.Println(fi14, v4, v3)
                v4 = (v4 / (((^(v4))) | 1))
            }
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 uint = 4294967295
            v17 = v17
            var v18 uint32 = 0
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 float64 = (3.14159 / ((0.0 / 1e10) + (float64(v3) * (float64(0.1) + 1.5))))
        _ = fi14
        _ = v19
    }
    fmt.Println(v4)
    v20 := GenS2{F0: "a", F1: -2147483648, F2: 1}
    switch (("hello" + "hello") + "a") {
    case "foo bar":
        v21 := GenS4{F0: -1}
        _ = v21
    case "a":
        var v22 bool = (!(-1 == -918))
        _ = v22
    case "hello":
        var v23 bool = (0.0 != v3)
        _ = v23
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    v24 := GenS5{F0: 4294967295}
    for fi25 := 0; fi25 < 4; fi25++ {
        v3 = float32(1e10)
        v26 := GenS3{F0: 999, F1: 1, F2: 917}
        _ = fi25
        _ = v26
    }
    v1 -= v1
    if false {
        switch int(v1) {
        case 169:
            fmt.Println(v1)
        case 32768:
            var v27 uint32 = 2
            _ = v27
        default:
            var v28 int64 = 717
            _ = v28
        }
        var v29 uint = 869
        _ = v29
    }
    var v30 uint = 2
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v30
    _ = v4
}
