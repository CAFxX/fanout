package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 bool; F2 uint }
func genFn0(p0 uint, p1 int32) uint32 {
    v5 := GenS1{F0: 32768}
    v6 := GenS2{F0: 355, F1: (!(p0 != (^(p0)))), F2: 0}
    p1 = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 int32) int8 {
    fmt.Println(p0)
    if true {
        if true {
            var v7 string = ("\n\t" + "\n\t")
            fmt.Println(v7, v7)
            _ = v7
        }
    }
    p0 += p0
    var v8 int8 = 0
    fmt.Println(p0)
    _ = p0
    _ = v8
    return (^(v8))
}

func main() {
    v1 := []int{-2147483648}
    v2 := []int{0, 391, 0, -707}
    var v3 int8 = 127
    var v4 int16 = -32768
    var v9 int64 = 481
    var v10 uint8 = 42
    if true {
        if true {
            if ((1 > 18446744073709551615) && ((v9 >> ((18446744073709551615) & 63)) > (v9 ^ v9))) {
                v9 -= v9
            } else {
                if len(v1) > 0 {
                    v1[0] = -268
                }
            }
        } else {
            if (!(-9223372036854775808 == 127)) {
                var v11 int = 32768
                _ = v11
            } else {
                v10 += v10
                v4 = -350
                fmt.Println(v10, v10, v10)
            }
        }
    } else {
        if len(v1) > 0 {
            v1[0] = -637
        }
    }
    var v12 uint16 = 65535
    v4 = (v4 + (v4 + v4))
    v3 += v3
    var v13 int16 = int16(v9)
    v14 := GenS2{F0: 409, F1: (true || (710 != 0)), F2: 42}
    if false {
        var v15 bool = false
        if (1e10 > (-2.25 * (((1.5 / ((0.0 + ((0.1 / 1.5) + 1e10)) * 1e10)) * 0.0) * (1e10 - 1e10)))) {
            switch (("foo bar" + ("hello" + ("\n\t" + (("hello" + (("\n\t" + "hello") + ("a" + "\n\t"))) + ("foo bar" + "hello"))))) + "\n\t") {
            case "a":
                var v16 int8 = (v3 << ((0) & 63))
                _ = v16
            default:
                fmt.Println(v10, v15)
            }
        } else {
            var v17 uint = 468
            var v18 uint8 = 31
            v19 := GenS1{F0: 9223372036854775807}
            _ = v17
            _ = v18
            _ = v19
        }
        _ = v15
    } else {
        v3 += v3
        if true {
            for fi20 := 0; fi20 < 18; fi20++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi20
            }
        } else {
            switch "a" {
            case "":
                v21 := GenS1{F0: 42}
                _ = v21
            default:
                fmt.Println(v9, v4, v3)
            }
            switch 128 {
            case -2147483648:
                v22 := GenS1{F0: 42}
                _ = v22
            case 9223372036854775807:
                for fi23 := 0; fi23 < 13; fi23++ {
                    if len(v2) > 0 {
                        v2[0] = fi23
                    }
                    var v24 bool = (!false)
                    var v25 uint64 = 2
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
