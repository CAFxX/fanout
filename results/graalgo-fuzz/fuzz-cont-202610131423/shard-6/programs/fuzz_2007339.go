package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []string{("hello" + "hello"), "", (("\n\t" + "\n\t") + ("\n\t" + ("foo bar" + "\n\t")))}
    v2 := []uint32{435}
    var v3 int = 32768
    v4 := make(map[string]bool)
    var v5 int = (v3 * v3)
    v5 = (v5 + (v3 % (((v3 ^ v3)) | 1)))
    if (65535 > 42) {
        if ((!false) && (-610 == 42)) {
            v6 := GenS1{F0: 1}
            _ = v6
        } else {
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 uint32 = 65535
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 6; fi9++ {
                fmt.Println(fi9, v5, v3)
                _ = fi9
            }
        }
        delete(v4, "a")
        var v10 uint32 = 4294967295
        _ = v10
    }
    var v11 int32 = -2147483648
    for fi12 := 0; fi12 < 10; fi12++ {
        if ((2 >= 128) || true) {
            if (false && ((("foo bar" + ("foo bar" + ("hello" + ("foo bar" + "a")))) + (("foo bar" + "") + "\n\t")) == "")) {
                switch 2147483647 {
                case -9223372036854775808:
                    if false {
                        v13 := GenS1{F0: 673}
                        v1 = append(v1, (("" + "") + "foo bar"))
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 746}
                        v11 += v11
                        _ = v14
                    }
                case -1:
                    v5 = 127
                case 422:
                    v15 := GenS2{F0: 0}
                    _ = v15
                default:
                    v5 = (v5 | fi12)
                }
                v3 -= v3
                var v16 uint32 = 42
                _ = v16
            } else {
                v11 = (v11 ^ (v11 % ((1) | 1)))
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 int32 = (v11 % (((^(v11))) | 1))
                    _ = fi17
                    _ = v18
                }
                fmt.Println(v11, v5, v11)
            }
            var v19 int8 = -128
            var v20 int64 = 0
            _ = v19
            _ = v20
        }
        var v21 int32 = -803
        for fi22 := 0; fi22 < 13; fi22++ {
            v23 := GenS1{F0: 524}
            _ = fi22
            _ = v23
        }
        _ = fi12
        _ = v21
    }
    switch (^(v5)) {
    case 264:
        fmt.Println(v3, v11)
    case -214:
        _ = len(v2)
    }
    var v24 uint8 = 1
    switch v3 {
    case 32767:
        if (1 != 1) {
            v25 := GenS1{F0: 4294967295}
            _ = v25
        }
    default:
        switch "" {
        case "\n\t":
            v2 = append(v2, 577)
        case "":
            for fi26 := 0; fi26 < 16; fi26++ {
                v11 -= v11
                _ = fi26
            }
        case "foo bar":
            v3 += (v3 - -9223372036854775808)
        }
    }
    var v27 uint8 = (v24 & (v24 << ((18446744073709551615) & 63)))
    var v28 uint64 = 19
    for fi29 := 0; fi29 < 9; fi29++ {
        fmt.Println(v28, v28, v5)
        v27 -= 100
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
