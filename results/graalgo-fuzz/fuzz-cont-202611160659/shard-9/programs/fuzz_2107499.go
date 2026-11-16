package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int8; F1 int8; F2 bool }
func genFn0() uint32 {
    if (!(65535 == 1)) {
        if (!true) {
            var v6 int32 = 1
            if (65535 >= 186) {
                v7 := GenS1{F0: 320}
                for fi8 := 0; fi8 < 8; fi8++ {
                    var v9 float64 = 3.14159
                    v10 := GenS2{F0: 1, F1: 16}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = v7
            }
            _ = v6
        }
    } else {
        var v11 float64 = float64(0.0)
        fmt.Println(v11)
        _ = v11
    }
    var v12 uint = 255
    if (-991 <= 124) {
        v12 = v12
        var v13 uint32 = 2
        _ = v13
    } else {
        switch "" {
        case "a":
            for fi14 := 0; fi14 < 17; fi14++ {
                fmt.Println(fi14, v12)
                _ = fi14
            }
        case "":
            for fi15 := 0; fi15 < 5; fi15++ {
                if true {
                    v16 := []int{fi15, (fi15 >> ((v12) & 63)), fi15, fi15}
                    _ = v16
                }
                _ = fi15
            }
        case "foo bar":
            var v17 uint32 = 2
            _ = v17
        default:
            v12 = v12
        }
        var v18 uint16 = 911
        switch "\n\t" {
        case "\n\t":
            v19 := GenS2{F0: -9223372036854775808, F1: 2}
            _ = v19
        }
        _ = v18
    }
    var v20 int64 = 1
    v20 -= v20
    _ = v12
    _ = v20
    return 4294967295
}

func main() {
    var v1 int8 = 42
    v2 := []uint32{1, 1, 337}
    var v3 uint = 968
    v4 := []uint32{1, 65535, 217, 2147483647}
    v5 := make(map[uint32]int)
    if true {
        v3 += v3
        v3 = (v3 - 4294967295)
    }
    var v21 uint = (v3 & (v3 | 18446744073709551615))
    v21 += (v3 << (((v21 * (v21 * 255))) & 63))
    if (156 != 1) {
        _ = len(v5)
        v3 += v3
    } else {
        if true {
            v1 += -96
        } else {
            _ = len(v5)
        }
        v3 = v3
        var v22 string = ("hello" + "\n\t")
        _ = v22
    }
    switch 128 {
    case 128:
        v21 -= (v3 % ((42) | 1))
    case -76:
        fmt.Println(v1, v3, v1)
    case 0:
        if len(v2) > 0 {
            v2[0] = 831
        }
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 bool = false
        _ = fi23
        _ = v24
    }
    switch ("" + ("" + ((((("\n\t" + "hello") + ("" + "")) + ("\n\t" + ("foo bar" + "hello"))) + "hello") + "\n\t"))) {
    case "foo bar":
        _ = len(v2)
    case "":
        v25 := GenS2{F0: -9223372036854775808, F1: 236}
        _ = v25
    default:
        v1 = v1
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        fmt.Println(v3)
        _ = fi26
    }
    var v27 uint16 = 389
    for fi28 := 0; fi28 < 10; fi28++ {
        if true {
            var v29 int8 = (^(v1))
            v29 = (-(v29))
            v21 += (^(v21))
            _ = v29
        }
        if ((!(2 < 18446744073709551615)) && (fi28 < (fi28 & (fi28 % ((fi28) | 1))))) {
            v27 = 65535
            if true {
                fmt.Println(v21, v21, v3)
            } else {
                var v30 int16 = 0
                v1 += (v1 << (((^(v21))) & 63))
                _ = v30
            }
        }
        _ = fi28
    }
    fmt.Println(v3, v3)
    v1 = v1
    if false {
        fmt.Println(v1, v21)
        fmt.Println(v21, v3)
    }
    v27 += (v27 >> ((v21) & 63))
    v31 := GenS3{F0: v1, F1: (-(v1)), F2: (140 < 42)}
    fmt.Println(v21, v27)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
