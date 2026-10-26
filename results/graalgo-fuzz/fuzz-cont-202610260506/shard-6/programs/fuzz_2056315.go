package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint8 }
func genFn0(p0 int8, p1 string) uint8 {
    if (!(-510 > -979)) {
        switch -392 {
        case 4294967295:
            v5 := []int{-160, -2147483648, 1, 127}
            _ = v5
        case -1:
            for fi6 := 0; fi6 < 18; fi6++ {
                v7 := []int{32768, (fi6 * -9223372036854775808)}
                fmt.Println(fi6, p1)
                var v8 bool = false
                _ = fi6
                _ = v7
                _ = v8
            }
        case -2147483648:
            v9 := GenS1{F0: -32768}
            _ = v9
        default:
            var v10 float64 = (((0.1 + (1e10 + float64((-2.25 - (1.5 + -2.25))))) * 1.5) - (float64(-2.25) + 0.0))
            _ = v10
        }
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint64 = 42
            for fi13 := 0; fi13 < 13; fi13++ {
                p1 = (p1 + (p1 + p1))
                _ = fi13
            }
            for fi14 := 0; fi14 < 16; fi14++ {
                switch (fi11 >> ((65535) & 63)) {
                case 173:
                    var v15 int32 = -2147483648
                    _ = v15
                default:
                    var v16 float64 = 0.1
                    _ = v16
                }
                switch (fi14 << ((18446744073709551615) & 63)) {
                case 127:
                    p0 += p0
                case 32767:
                    v17 := GenS2{F0: (fi11 + (fi14 & (fi11 ^ fi11)))}
                    _ = v17
                default:
                    fmt.Println(p0, fi11, p0)
                }
                v18 := GenS3{F0: fi14, F1: 42}
                _ = fi14
                _ = v18
            }
            _ = fi11
            _ = v12
        }
        p1 = p1
    } else {
        var v19 int16 = 0
        p1 = p1
        _ = v19
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 208
}

func main() {
    var v1 int = 42
    var v2 int64 = -9223372036854775808
    var v3 uint32 = 65535
    var v4 int16 = -326
    v4 += int16(v2)
    for fi20 := 0; fi20 < 8; fi20++ {
        if (((0.0 <= 1.5) || true) || (true || (((true && (127 <= 2)) || ("\n\t" >= ((((("\n\t" + (("\n\t" + "foo bar") + ("\n\t" + "\n\t"))) + "foo bar") + "") + ("\n\t" + "foo bar")) + (("foo bar" + "hello") + ("" + ""))))) && (false && false)))) {
            var v21 uint64 = 4294967295
            _ = v21
        } else {
            v22 := []int{127, 2147483647, v1, v1}
            v23 := GenS3{F0: fi20, F1: uint8(v3)}
            _ = v22
            _ = v23
        }
        if true {
            for fi24 := 0; fi24 < 8; fi24++ {
                if (((-2.25 + float64(float32(0.0))) + ((0.1 - -2.25) + 1e10)) <= 1e10) {
                    v25 := []int{fi24}
                    _ = v25
                } else {
                    fmt.Println(v1, v4, v4)
                    switch "hello" {
                    case "foo bar":
                        v3 += v3
                    case "":
                        v2 = v2
                        if true {
                            v26 := GenS2{F0: int(v4)}
                            _ = v26
                        }
                    }
                    var v27 uint16 = 0
                    _ = v27
                }
                v3 = (v3 + 2147483648)
                _ = fi24
            }
        }
        _ = fi20
    }
    v28 := []int{(v1 * v1), int(v4)}
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := GenS1{F0: (v4 % ((v4) | 1))}
        _ = fi29
        _ = v30
    }
    var v31 string = ("" + "\n\t")
    v32 := GenS1{F0: v4}
    fmt.Println(v31)
    for fi33 := 0; fi33 < 19; fi33++ {
        v3 = 2147483648
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
