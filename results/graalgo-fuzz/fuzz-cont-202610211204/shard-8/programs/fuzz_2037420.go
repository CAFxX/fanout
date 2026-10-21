package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }
type GenS3 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 0
    var v3 float32 = 0.0
    var v4 int64 = -1
    v5 := make(map[int]string)
    v6 := []int{2147483647, 0}
    v3 += (float32(((3.14159 - 0.0) / 0.1)) / (-2.25 / v3))
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            switch (("" + "hello") + "hello") {
            case "a":
                v3 -= (v3 / 1e10)
            }
            switch ("foo bar" + ("a" + ("hello" + "\n\t"))) {
            case "hello":
                v4 -= 0
                var v9 int64 = (v4 | (v4 ^ -9223372036854775808))
                _ = v9
            case "foo bar":
                v6 = append(v6, fi8)
            case "":
                v4 = (v4 * v4)
            default:
                var v10 float32 = v3
                _ = v10
            }
            _ = fi8
        }
        if (true || true) {
            v3 = (float32(0.1) / v3)
        } else {
            if false {
                v1[(fi7 | 32767)] = false
            } else {
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 float64 = ((((3.14159 * (0.1 - 1.5)) - (1.5 / 1.0)) / 1.0) * 0.1)
                    v3 = v3
                    var v13 int16 = 8
                    var v14 uint16 = 667
                    switch fi11 {
                    case 42:
                        v15 := GenS1{F0: 42, F1: 761, F2: v4}
                        _ = v15
                    default:
                        switch "a" {
                        case "\n\t":
                            var v16 int = fi7
                            _ = v16
                        default:
                            switch ("\n\t" + "hello") {
                            case "":
                                switch (fi11 << ((65535) & 63)) {
                                case 128:
                                    switch (fi7 >> ((uint(v14)) & 63)) {
                                    case 32767:
                                        v17 := GenS2{F0: -1, F1: -2147483648, F2: -104}
                                        _ = v17
                                    case -748:
                                        v12 -= ((0.0 - v12) - (1.5 / 1.0))
                                    default:
                                        fmt.Println(v14, v4)
                                    }
                                }
                            }
                        }
                    }
                    v18 := GenS1{F0: 431, F1: 764, F2: v4}
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v18
                }
                fmt.Println(v3, fi7, fi7)
            }
            v19 := v1[9223372036854775807]
            _ = v19
        }
        v6 = append(v6, (fi7 | fi7))
        _ = fi7
    }
    v2 = v2
    v20 := GenS1{F0: 0, F1: 355, F2: -1}
    var v21 int16 = 1
    if (((32767 < (v21 / ((-32768) | 1))) || true) || ((42 <= 0) && (1 != 1))) {
        for fi22 := 0; fi22 < 9; fi22++ {
            var v23 uint8 = v2
            v24 := GenS3{F0: (true || (4294967295 == 4294967295)), F1: 2}
            var v25 string = "hello"
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 int = -2147483648
        var v28 uint64 = 4294967295
        _ = fi26
        _ = v27
        _ = v28
    }
    v4 = (v4 >> ((2) & 63))
    fmt.Println(v21, v2, v2)
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
