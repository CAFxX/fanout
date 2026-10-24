package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 bool }

func main() {
    var v1 int64 = -760
    var v2 int = 128
    var v3 uint32 = 271
    var v4 uint64 = 535
    v3 = (v3 % (((v3 / (((v3 << ((517) & 63))) | 1))) | 1))
    fmt.Println(v4)
    var v5 int16 = -201
    if false {
        if (108 > 255) {
            var v6 int = (^(v2))
            _ = v6
        } else {
            if false {
                v7 := []int{v2, v2, (v2 << ((0) & 63)), (v2 >> ((187) & 63))}
                _ = v7
            } else {
                v8 := GenS1{F0: 0, F1: "foo bar"}
                v9 := []int{(v2 % (((v2 + (v2 >> ((429) & 63)))) | 1))}
                v10 := GenS1{F0: v1, F1: ""}
                _ = v10
                _ = v8
                _ = v9
            }
            v5 = v5
            v5 -= (v5 >> ((2) & 63))
        }
        v4 = (v4 % (((v4 * 18446744073709551615)) | 1))
    } else {
        fmt.Println(v1, v3)
    }
    if (42 != 1) {
        var v11 bool = (true && true)
        v5 = (v5 << ((4294967295) & 63))
        v12 := GenS2{F0: (float64(-2.25) <= 0.1)}
        _ = v11
        _ = v12
    } else {
        var v13 float32 = (3.14159 - (0.1 - (float32(0.0) * float32(0.1))))
        _ = v13
    }
    fmt.Println(v2, v2)
    for fi14 := 0; fi14 < 16; fi14++ {
        switch v2 {
        case 342:
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 bool = ((true || true) && true)
                var v17 uint16 = 2
                var v18 bool = v16
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        default:
            v19 := []int{(fi14 >> ((uint(v4)) & 63)), (v2 >> ((632) & 63))}
            _ = v19
        }
        v2 += (fi14 | -1)
        switch (fi14 + 42) {
        case 1:
            for fi20 := 0; fi20 < 18; fi20++ {
                if (((42 > -128) || true) || ((v1 & -9223372036854775808) == (v1 % ((v1) | 1)))) {
                    v21 := []int{fi20, (v2 % ((v2) | 1)), 950}
                    _ = len(v21)
                    _ = v21
                }
                _ = fi20
            }
        case 32767:
            v22 := []int{fi14, v2}
            _ = v22
        case -393:
            var v23 int64 = int64(v4)
            _ = v23
        default:
            var v24 string = ((("\n\t" + "") + "a") + (("" + (("" + "\n\t") + ("" + ""))) + "hello"))
            _ = v24
        }
        _ = fi14
    }
    if (((((0.1 / (float32((1.5 - (-2.25 + 1.5))) - 1.5)) / 3.14159) - ((0.0 + (3.14159 - 1.5)) / float32(3.14159))) + -2.25) > (-2.25 + ((1e10 - (float32(-2.25) / 1.0)) * 0.1))) {
        if true {
            for fi25 := 0; fi25 < 13; fi25++ {
                v2 -= 559
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 int16 = v5
                    _ = fi26
                    _ = v27
                }
                v28 := GenS2{F0: (!((!(0.1 == (-2.25 * (0.1 * (((-2.25 * 1.5) - float32((1e10 + 1.5))) / 3.14159))))) || (255 > 255)))}
                _ = fi25
                _ = v28
            }
        }
        fmt.Println(v3, v5)
        v2 -= -2147483648
    } else {
        fmt.Println(v3, v3, v3)
    }
    var v29 uint64 = (v4 ^ 42)
    var v30 uint16 = 2
    if (1 == (v29 >> ((255) & 63))) {
        if false {
            fmt.Println(v29, v29)
        } else {
            v31 := GenS1{F0: v1, F1: ""}
            v3 -= (v3 / ((v3) | 1))
            _ = v31
        }
        fmt.Println(v30, v30)
    } else {
        var v32 uint8 = 215
        v30 += 42
        var v33 int16 = (^(v5))
        _ = v32
        _ = v33
    }
    switch v2 {
    case 37:
        var v34 int16 = v5
        _ = v34
    case 526:
        v35 := []int{(v2 >> ((1) & 63)), 1, -135}
        _ = v35
    case 32767:
        var v36 int32 = 0
        _ = v36
    }
    v30 -= (v30 ^ (v30 - 65535))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
