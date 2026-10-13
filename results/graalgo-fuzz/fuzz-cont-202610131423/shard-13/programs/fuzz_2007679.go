package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint64; F1 bool }

func main() {
    var v1 float32 = 1.5
    v2 := []string{("hello" + "")}
    var v3 int = -2147483648
    var v4 int64 = 1
    v5 := GenS1{F0: 42, F1: 831, F2: 9223372036854775807}
    v6 := GenS1{F0: -157, F1: 978, F2: (v4 >> ((451) & 63))}
    _ = len(v2)
    fmt.Println(v1)
    v7 := GenS2{F0: 2, F1: false}
    v4 += v4
    v8 := GenS2{F0: 18446744073709551615, F1: (42 < 42)}
    v3 = (v3 * (v3 / (((v3 ^ v3)) | 1)))
    if (v3 == 32767) {
        var v9 int8 = -1
        _ = v9
    } else {
        switch (^(v3)) {
        case 128:
            var v10 int = (v3 / ((v3) | 1))
            _ = v10
        case -201:
            v11 := GenS1{F0: -543, F1: 852, F2: 1}
            _ = v11
        default:
            if false {
                var v12 float64 = (1e10 / (1e10 - 0.1))
                v4 += (v4 * v4)
                v13 := GenS2{F0: 0, F1: (!(42 <= 579))}
                _ = v12
                _ = v13
            } else {
                switch -1 {
                case -154:
                    v3 += v3
                }
                v3 = 1
                switch 0 {
                case -57:
                    for fi14 := 0; fi14 < 19; fi14++ {
                        _ = len(v2)
                        if false {
                            switch (("" + "\n\t") + ("a" + "")) {
                            case "a":
                                var v15 uint8 = 128
                                _ = v15
                            }
                        }
                        for fi16 := 0; fi16 < 6; fi16++ {
                            var v17 uint64 = 4294967295
                            _ = fi16
                            _ = v17
                        }
                        _ = fi14
                    }
                case -1:
                    fmt.Println(v3, v4, v1)
                case 1:
                    var v18 int32 = -450
                    _ = v18
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        }
        v4 = v4
        switch v3 {
        case 1:
            fmt.Println(v3, v3)
        }
    }
    if false {
        switch "\n\t" {
        case "\n\t":
            var v19 int16 = 805
            _ = v19
        case "foo bar":
            var v20 int = (v3 / ((127) | 1))
            _ = v20
        }
        if (true || (1.5 > (1.5 - float32(-2.25)))) {
            v3 -= (-(v3))
            v21 := GenS1{F0: 2147483647, F1: 479, F2: -1}
            _ = v21
        } else {
            var v22 int32 = 1
            v22 += (v22 ^ (v22 ^ v22))
            _ = v22
        }
    } else {
        var v23 float64 = (3.14159 * 1.5)
        v24 := GenS2{F0: 0, F1: false}
        if len(v2) > 0 {
            v2[0] = ("" + "\n\t")
        }
        _ = v23
        _ = v24
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v4 += v4
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
