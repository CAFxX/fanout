package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 int, p1 uint64) uint8 {
    p1 = (^(p1))
    p0 += (-(p0))
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int8 = 0
    var v2 bool = (v1 != 100)
    var v3 uint = 399
    v4 := []int{-9223372036854775808, 127, -1}
    v5 := []int{-1}
    var v6 int8 = (v1 >> ((336) & 63))
    fmt.Println(v6, v2)
    switch "\n\t" {
    case "\n\t":
        if ((3.14159 + 3.14159) == 3.14159) {
            var v7 int16 = 1
            v6 -= (-(v1))
            _ = v7
        } else {
            var v8 uint8 = 0
            v1 = (v1 * 32)
            v1 += 100
            _ = v8
        }
    case "a":
        v9 := GenS1{F0: -1}
        _ = v9
    case "":
        var v10 int = 672
        _ = v10
    }
    v2 = ((1e10 >= float64(0.0)) || (!((^(v3)) < 1)))
    v11 := GenS1{F0: 1}
    switch ("a" + (("a" + "a") + ((("" + "foo bar") + "foo bar") + "hello"))) {
    case "\n\t":
        switch 4294967295 {
        case 968:
            for fi12 := 0; fi12 < 5; fi12++ {
                for fi13 := 0; fi13 < 18; fi13++ {
                    if len(v4) > 0 {
                        v4[0] = 32767
                    }
                    var v14 int32 = -2147483648
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v2)
                _ = fi12
            }
        case -517:
            v1 -= -128
        default:
            var v15 int16 = 1000
            _ = v15
        }
    }
    v16 := GenS2{F0: 65535, F1: 1000}
    if (42 <= -1) {
        v6 = v6
        v6 += (v1 & v6)
    }
    if (65535 == v3) {
        if v2 {
            v17 := GenS2{F0: 2, F1: 42}
            if (!v2) {
                fmt.Println(v2, v3)
                fmt.Println(v2, v3, v3)
            } else {
                v18 := GenS1{F0: -875}
                _ = v18
            }
            var v19 int16 = -994
            _ = v17
            _ = v19
        } else {
            fmt.Println(v3, v1, v6)
            for fi20 := 0; fi20 < 8; fi20++ {
                switch ("foo bar" + "foo bar") {
                case "":
                    switch 4294967295 {
                    case -10:
                        var v21 int = fi20
                        _ = v21
                    default:
                        v6 = int8(fi20)
                    }
                default:
                    fmt.Println(v2)
                }
                v22 := GenS1{F0: 669}
                _ = fi20
                _ = v22
            }
            for fi23 := 0; fi23 < 11; fi23++ {
                if false {
                    v24 := GenS2{F0: 65535, F1: -26}
                    _ = v24
                } else {
                    v25 := GenS1{F0: 2147483647}
                    _ = v25
                }
                var v26 float64 = -2.25
                v27 := GenS1{F0: 1}
                _ = fi23
                _ = v26
                _ = v27
            }
        }
    } else {
        v1 = -5
        if len(v4) > 0 {
            v4[0] = -716
        }
        var v28 float32 = (float32(((0.0 - 3.14159) + 0.1)) / 3.14159)
        _ = v28
    }
    _ = len(v4)
    var v29 uint16 = 1
    v30 := GenS1{F0: 714}
    if len(v5) > 0 {
        v5[0] = 32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
