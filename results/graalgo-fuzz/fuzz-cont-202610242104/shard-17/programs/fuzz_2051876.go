package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 int32 }
func genFn0(p0 int32) uint8 {
    if (!true) {
        p0 = 42
        v7 := GenS1{F0: (p0 >> ((18446744073709551615) & 63)), F1: 42, F2: -774}
        for fi8 := 0; fi8 < 1; fi8++ {
            if (204 == (p0 ^ (p0 << ((18446744073709551615) & 63)))) {
                var v9 int = (fi8 / ((fi8) | 1))
                _ = v9
            } else {
                p0 = 811
                for fi10 := 0; fi10 < 5; fi10++ {
                    switch fi10 {
                    case 458:
                        v11 := []int{(fi10 | fi10), 9223372036854775807, fi10, (fi8 << ((0) & 63))}
                        _ = v11
                    }
                    v12 := GenS2{F0: fi10, F1: 128}
                    p0 = p0
                    _ = fi10
                    _ = v12
                }
                p0 = (^(p0))
            }
            p0 += (-(p0))
            _ = fi8
        }
        _ = v7
    } else {
        v13 := []int{127}
        _ = v13
    }
    if (-697 <= 0) {
        var v14 uint = 0
        _ = v14
    }
    switch (("foo bar" + ("foo bar" + ("hello" + "hello"))) + "\n\t") {
    case "hello":
        var v15 bool = false
        _ = v15
    case "a":
        p0 -= (p0 % (((p0 - (p0 << ((2) & 63)))) | 1))
    case "\n\t":
        p0 -= p0
    }
    fmt.Println(p0)
    _ = p0
    return 180
}
func genFn1() int8 {
    var v16 int16 = 0
    if (18446744073709551615 < 2) {
        v16 = (v16 | -435)
        fmt.Println(v16)
    } else {
        var v17 int64 = 880
        _ = v17
    }
    v16 = (v16 * (v16 << ((4294967295) & 63)))
    _ = v16
    return 100
}

func main() {
    var v1 int32 = 1
    var v2 int16 = 32767
    var v3 uint8 = 2
    var v4 int64 = -9223372036854775808
    v5 := []string{("\n\t" + "a")}
    v6 := make(map[int]bool)
    for fi18 := 0; fi18 < 12; fi18++ {
        v4 = (v4 & int64(fi18))
        _ = fi18
    }
    v3 -= v3
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 float64 = 0.0
        _ = fi19
        _ = v20
    }
    v1 = -2147483648
    if true {
        if ((65535 <= 4294967295) && false) {
            v21 := v6[2147483647]
            _ = v21
        }
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS3{F0: v2, F1: "foo bar", F2: -277}
            fmt.Println(fi22, v3, v2)
            _ = fi22
            _ = v23
        }
    } else {
        switch ("" + "a") {
        case "hello":
            var v24 uint64 = 0
            _ = v24
        default:
            switch -221 {
            case -873:
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 int16 = v2
                    _ = fi25
                    _ = v26
                }
            case -601:
                switch "\n\t" {
                case "":
                    switch 4294967295 {
                    case 128:
                        fmt.Println(v1, v4)
                    case 2147483647:
                        v1 += int32(v2)
                    case 127:
                        v4 = v4
                    default:
                        v1 = (^(v1))
                    }
                default:
                    var v27 int64 = 0
                    _ = v27
                }
            default:
                v1 = (v1 - int32(v2))
            }
        }
    }
    fmt.Println(v4, v4, v1)
    fmt.Println(v4)
    var v28 float64 = (0.1 / 3.14159)
    v4 = v4
    var v29 uint8 = 143
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
