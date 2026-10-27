package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }

func main() {
    var v1 float64 = float64((float32(1.5) + 0.1))
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]uint32)
    v4 := []string{"foo bar"}
    v5 := make(map[string]int)
    var v6 uint = 894
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint8 = 128
            switch (fi8 + fi8) {
            case -2147483648:
                v6 = v6
            case 1:
                v1 += v1
                var v10 uint = (v6 >> ((2) & 63))
                _ = v10
            case 0:
                var v11 uint = (^(v6))
                if (false || (false || ((32767 != 65535) && true))) {
                    var v12 float32 = 0.1
                    if true {
                        fmt.Println(v9, fi8, v9)
                        v13 := v3[65535]
                        var v14 uint8 = v9
                        _ = v13
                        _ = v14
                    }
                    _ = v12
                }
                v9 = v9
                _ = v11
            default:
                fmt.Println(v9, fi8)
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    switch 2147483647 {
    case 487:
        switch 127 {
        case -9223372036854775808:
            v15 := GenS1{F0: -2147483648, F1: (^(v2)), F2: 0}
            _ = v15
        case 2147483647:
            v6 = v6
        }
    case 32767:
        v2 = (^(v2))
    default:
        var v16 uint8 = 73
        _ = v16
    }
    if (255 < 2) {
        v17 := GenS1{F0: 858, F1: (v2 - (v2 << (((^(v6))) & 63))), F2: -998}
        _ = v17
    }
    fmt.Println(v1, v6)
    v2 -= v2
    var v18 uint = v6
    fmt.Println(v18, v6, v6)
    switch 9223372036854775807 {
    case 4294967295:
        if false {
            var v19 uint64 = (v2 & v2)
            _ = v19
        }
    case 95:
        for fi20 := 0; fi20 < 11; fi20++ {
            v18 = 65535
            _ = fi20
        }
    case -717:
        v18 = v18
    }
    v18 -= (v18 - uint(v2))
    if false {
        v18 += (v6 & (v18 & (v18 % (((v6 - 1)) | 1))))
        v21 := GenS1{F0: 2147483647, F1: (v2 << ((42) & 63)), F2: 0}
        _ = v21
    }
    var v22 uint64 = (v2 * 4294967295)
    var v23 int = 492
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint16 = 42
        var v26 uint16 = (^(v25))
        switch "" {
        case "":
            var v27 bool = ((!(895 >= 1)) || true)
            _ = v27
        case "foo bar":
            for fi28 := 0; fi28 < 17; fi28++ {
                switch ((("a" + "hello") + "foo bar") + "foo bar") {
                case "\n\t":
                    var v29 float32 = float32(((v1 / 0.1) * v1))
                    _ = v29
                case "hello":
                    v23 = v23
                case "":
                    if true {
                        var v30 int32 = -2147483648
                        _ = v30
                    } else {
                        if true {
                            fmt.Println(v18)
                            if (255 == 1) {
                                fmt.Println(v22, v22)
                                v31 := GenS1{F0: 2147483647, F1: 1, F2: -489}
                                _ = v31
                            }
                        }
                    }
                }
                _ = fi28
            }
        case "hello":
            if (0 <= -2147483648) {
                var v32 int = 988
                for fi33 := 0; fi33 < 19; fi33++ {
                    v6 = 343
                    _ = fi33
                }
                var v34 uint32 = 1
                _ = v32
                _ = v34
            } else {
                v35 := GenS2{F0: 128}
                for fi36 := 0; fi36 < 8; fi36++ {
                    fmt.Println(v25)
                    _ = fi36
                }
                v37 := GenS3{F0: v2, F1: (v2 | v2)}
                _ = v35
                _ = v37
            }
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
