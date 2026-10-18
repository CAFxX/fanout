package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS4 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]uint32)
    var v3 int64 = 734
    var v4 uint8 = 2
    var v5 uint32 = 1
    if ((!(255 != (v4 * uint8(v1)))) && true) {
        var v6 float64 = 1e10
        _ = v6
    }
    if (794 != 32767) {
        v2[("" + "foo bar")] = (v5 << ((42) & 63))
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint32 = 269
        _ = fi7
        _ = v8
    }
    var v9 uint8 = v4
    switch -9223372036854775808 {
    case -1:
        _ = len(v2)
    default:
        fmt.Println(v3, v3)
    }
    if (!false) {
        v9 = 0
        if (0 != 42) {
            if false {
                delete(v2, "hello")
                var v10 float64 = -2.25
                if (((v3 & (v3 % (((v3 & v3)) | 1))) < -9223372036854775808) || ("a" == "a")) {
                    switch "" {
                    case "a":
                        v10 = v10
                    case "hello":
                        fmt.Println(v9, v9, v3)
                    case "foo bar":
                        fmt.Println(v5, v5)
                    default:
                        v4 -= 1
                    }
                    v11 := []int{32768, 339, 32768, -418}
                    _ = v11
                } else {
                    var v12 int = 128
                    switch "" {
                    case "\n\t":
                        v13 := GenS1{F0: "hello", F1: 255, F2: 2}
                        _ = v13
                    case "foo bar":
                        var v14 int8 = -120
                        _ = v14
                    default:
                        v15 := GenS1{F0: (("foo bar" + (("" + ((("" + "\n\t") + "a") + "a")) + ("a" + ("a" + "")))) + ((("" + ("\n\t" + "foo bar")) + ("\n\t" + "")) + "a")), F1: 18446744073709551615, F2: 882}
                        _ = v15
                    }
                    _ = v12
                }
                _ = v10
            }
            for fi16 := 0; fi16 < 15; fi16++ {
                v1 -= (v1 | (v1 + (v1 << ((1) & 63))))
                switch (fi16 - fi16) {
                case 9223372036854775807:
                    v1 += (v1 / (((^(v1))) | 1))
                case 32767:
                    var v17 int8 = 100
                    _ = v17
                case 0:
                    var v18 uint32 = v5
                    _ = v18
                default:
                    v19 := GenS2{F0: v3, F1: (v1 << ((18446744073709551615) & 63))}
                    _ = v19
                }
                _ = fi16
            }
            v5 = v5
        } else {
            switch "hello" {
            case "a":
                fmt.Println(v1)
            default:
                fmt.Println(v5, v5)
            }
            if false {
                v20 := []int{32768, -781, 384, int(v1)}
                v21 := GenS3{F0: v1, F1: 42, F2: (^(v9))}
                _ = v20
                _ = v21
            }
            v9 += (v9 << ((4294967295) & 63))
        }
    } else {
        var v22 float64 = 1.5
        v22 = float64(-2.25)
        v1 += v1
        _ = v22
    }
    switch "a" {
    case "foo bar":
        v2["foo bar"] = (^(v5))
    case "":
        fmt.Println(v9, v3, v1)
    case "a":
        v5 = 2147483647
    }
    fmt.Println(v4)
    switch int(v1) {
    case 9223372036854775807:
        v5 += (^(v5))
    case 0:
        switch -9223372036854775808 {
        case 801:
            var v23 uint8 = v4
            _ = v23
        case 255:
            var v24 int8 = 42
            _ = v24
        case -234:
            var v25 float32 = float32(0.0)
            _ = v25
        default:
            var v26 int64 = v3
            _ = v26
        }
    case 32767:
        v5 = (v5 << ((18446744073709551615) & 63))
    default:
        var v27 bool = (!(false || (65535 < 4294967295)))
        _ = v27
    }
    fmt.Println(v4, v1, v5)
    v28 := GenS4{F0: v3, F1: -1}
    var v29 int32 = -389
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
