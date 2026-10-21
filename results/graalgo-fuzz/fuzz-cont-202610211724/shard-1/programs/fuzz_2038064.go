package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 string; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int64; F1 int16; F2 int64 }

func main() {
    v1 := []uint32{0, 220, 4294967295}
    v2 := make(map[int]uint32)
    var v3 bool = true
    var v4 int32 = 757
    var v5 uint8 = 2
    var v6 uint16 = 560
    switch ("" + ("\n\t" + "hello")) {
    case "hello":
        switch 128 {
        case -9223372036854775808:
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 int8 = 42
                fmt.Println(v5, fi7, v6)
                _ = fi7
                _ = v8
            }
        case 127:
            _ = len(v1)
        case -138:
            fmt.Println(v5, v3, v4)
        default:
            v5 = 128
        }
    }
    if (v6 > 718) {
        var v9 uint16 = v6
        if v3 {
            v10 := GenS1{F0: -32768}
            _ = v10
        }
        if v3 {
            if (0.1 > 3.14159) {
                switch "hello" {
                case "hello":
                    fmt.Println(v5, v6)
                case "foo bar":
                    v3 = false
                case "\n\t":
                    v6 -= v6
                default:
                    v3 = v3
                }
                v4 = v4
            } else {
                if v3 {
                    v9 = v6
                }
                var v11 float64 = (1e10 - (0.1 * (3.14159 + 1.5)))
                if false {
                    var v12 string = ("a" + (("foo bar" + ((("\n\t" + "a") + (("" + ("hello" + ("a" + ("\n\t" + ("hello" + ("\n\t" + "foo bar")))))) + "hello")) + "foo bar")) + "a"))
                    _ = v12
                } else {
                    v13 := GenS2{F0: 42, F1: "", F2: -9223372036854775808}
                    _ = v13
                }
                _ = v11
            }
        }
        _ = v9
    } else {
        var v14 uint32 = 326
        v3 = false
        _ = v14
    }
    var v15 float64 = 3.14159
    v16 := GenS3{F0: v5, F1: 547, F2: 1}
    var v17 uint64 = 144
    v6 += v6
    fmt.Println(v17, v4, v3)
    if v3 {
        if (434 < (v17 * (v17 - 0))) {
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := GenS4{F0: v6, F1: (^(v17)), F2: 609}
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 18; fi20++ {
                v5 -= (^(v5))
                switch (fi20 + (fi20 - fi20)) {
                case -1:
                    v21 := GenS4{F0: v6, F1: v17, F2: (v17 * 18446744073709551615)}
                    _ = v21
                case 42:
                    _ = len(v1)
                case 2147483647:
                    var v22 bool = (true && (-953 <= 0))
                    _ = v22
                }
                _ = fi20
            }
            v23 := GenS5{F0: 1, F1: 32767, F2: -9223372036854775808}
            _ = v23
        } else {
            v17 -= (^(v17))
            v4 = v4
        }
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        _ = len(v1)
        var v25 uint = 462
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 uint8 = (v5 % ((255) | 1))
        var v28 uint32 = 4294967295
        if ((-(v4)) > -1) {
            for fi29 := 0; fi29 < 20; fi29++ {
                v15 = v15
                v17 = 779
                _ = fi29
            }
            v30 := GenS4{F0: v6, F1: (^(v17)), F2: (^(v17))}
            v31 := GenS4{F0: 42, F1: v17, F2: 559}
            _ = v30
            _ = v31
        } else {
            if ((v6 & (v6 % ((v6) | 1))) > (v6 << ((768) & 63))) {
                delete(v2, -187)
                v5 += (v5 % ((uint8(v28)) | 1))
            } else {
                fmt.Println(v17)
                var v32 uint = 981
                var v33 int64 = 1
                _ = v32
                _ = v33
            }
            v3 = v3
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
