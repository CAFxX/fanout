package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 string; F1 uint8 }
func genFn0(p0 uint64, p1 uint64, p2 bool) string {
    v4 := []int{992, 1, -2147483648, 127}
    if (!p2) {
        if false {
            fmt.Println(p1)
            for fi5 := 0; fi5 < 8; fi5++ {
                var v6 bool = (!(true || p2))
                _ = fi5
                _ = v6
            }
        } else {
            v7 := GenS1{F0: 603, F1: 32767}
            p1 -= p1
            _ = v7
        }
        if (!p2) {
            var v8 uint64 = 18446744073709551615
            v8 += 129
            _ = len(v4)
            _ = v8
        }
        if p2 {
            p1 += p1
        } else {
            v9 := GenS1{F0: 0, F1: 42}
            p1 -= (p1 & (p0 >> ((42) & 63)))
            p0 = (^(p1))
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(p2, p2)
            if false {
                v11 := GenS2{F0: 32767, F1: 65535, F2: 490}
                _ = v11
            } else {
                v12 := GenS3{F0: 0, F1: -683}
                var v13 int = 0
                _ = v12
                _ = v13
            }
            if p2 {
                if ((((fi10 >> ((18446744073709551615) & 63)) <= 505) || (2 > 2147483647)) || true) {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        switch "a" {
                        case "hello":
                            var v15 uint16 = 255
                            _ = v15
                        default:
                            var v16 bool = (false && (18446744073709551615 < 656))
                            _ = v16
                        }
                        p0 += (p0 % (((p0 * p1)) | 1))
                        _ = fi14
                    }
                } else {
                    fmt.Println(p0)
                }
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            } else {
                p2 = true
                var v17 int16 = 32767
                _ = v17
            }
            _ = fi10
        }
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 float64 = (3.14159 * (1.5 * (((-2.25 * float64(0.1)) - -2.25) + 0.1)))
            _ = fi18
            _ = v19
        }
    }
    p0 += (p1 - p0)
    switch 32768 {
    case 42:
        if ((-2.25 <= float64((((-2.25 * 0.1) / 0.1) - float32((1.5 / 1.5))))) && p2) {
            p0 = (p1 & (p0 & (p0 >> ((383) & 63))))
        } else {
            p2 = true
            fmt.Println(p2)
            var v20 uint64 = 257
            _ = v20
        }
    default:
        p0 = (p0 % ((414) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ("hello" + ("foo bar" + "\n\t"))
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]uint32)
    v3 := []string{"foo bar", "foo bar", (("" + "foo bar") + "\n\t")}
    v1 += v1
    v1 = (-(v1))
    v1 -= 2147483647
    v21 := GenS4{F0: -759, F1: 1}
    v1 = -961
    v1 -= v1
    v22 := GenS1{F0: 32767, F1: -954}
    v1 = v1
    if false {
        v1 = -2147483648
        v1 = (v1 | (^(v1)))
        switch 9223372036854775807 {
        case 128:
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 int = fi23
                if (!true) {
                    var v25 string = "\n\t"
                    var v26 int32 = v1
                    _ = v25
                    _ = v26
                }
                for fi27 := 0; fi27 < 18; fi27++ {
                    fmt.Println(v1, fi23)
                    var v28 int64 = -533
                    v29 := GenS5{F0: "foo bar", F1: 128}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = fi23
                _ = v24
            }
        default:
            switch -355 {
            case 4294967295:
                var v30 uint16 = 0
                _ = v30
            default:
                if len(v3) > 0 {
                    v3[0] = "a"
                }
            }
        }
    } else {
        if ("a" <= "hello") {
            v1 = v1
            v1 += (^(v1))
            var v31 uint8 = 128
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
