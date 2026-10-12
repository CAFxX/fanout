package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0() uint8 {
    if true {
        if (3.14159 != float32(1e10)) {
            switch "a" {
            case "hello":
                v6 := []int{-9223372036854775808, 744}
                _ = v6
            case "\n\t":
                var v7 float64 = ((0.0 - 0.0) + (1.5 / 1.5))
                _ = v7
            case "foo bar":
                var v8 bool = ((2147483647 > 128) || ((-1 < -1) || ((3.14159 == (1.5 * 1.5)) || false)))
                _ = v8
            default:
                var v9 float32 = 3.14159
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 12; fi10++ {
                v11 := []int{}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 uint8 = 128
                _ = fi12
                _ = v13
            }
        }
        var v14 uint8 = 7
        _ = v14
    }
    switch -2147483648 {
    case 0:
        v15 := []int{1, 2147483647, -2147483648, 32768}
        _ = v15
    case -483:
        var v16 uint64 = 42
        _ = v16
    case -1:
        v17 := []int{}
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 float64 = 0.1
        v20 := []int{(fi18 % ((fi18) | 1))}
        _ = fi18
        _ = v19
        _ = v20
    }
    return 2
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 42
    var v3 int64 = 9223372036854775807
    var v4 string = ("foo bar" + ((("foo bar" + ("" + (("" + "hello") + (("" + "a") + "hello")))) + "") + ""))
    var v5 int64 = (-(v3))
    var v21 float64 = float64((((0.0 * 1e10) * 1.5) - float32(-2.25)))
    v4 = "foo bar"
    for fi22 := 0; fi22 < 8; fi22++ {
        v23 := []int{fi22, -93, -9223372036854775808, 2147483647}
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 float32 = 1.5
            _ = fi24
            _ = v25
        }
        fmt.Println(v5, v21)
        _ = fi22
        _ = v23
    }
    var v26 string = v4
    if (!(4294967295 == 390)) {
        if true {
            var v27 int64 = (v5 >> ((0) & 63))
            _ = v27
        }
        fmt.Println(v26, v2)
    } else {
        v26 = (((v4 + "a") + ((v4 + v26) + (("foo bar" + ((v4 + v4) + v26)) + v4))) + v4)
        v28 := GenS1{F0: 609}
        _ = v28
    }
    var v29 uint32 = 42
    v30 := []int{-963, 42, -436}
    var v31 int = -9223372036854775808
    v31 += 0
    if (1000 >= -793) {
        v2 += int64(v29)
        switch (v4 + v26) {
        case "a":
            for fi32 := 0; fi32 < 14; fi32++ {
                _ = len(v1)
                for fi33 := 0; fi33 < 17; fi33++ {
                    v34 := GenS2{F0: 1, F1: 42}
                    _ = fi33
                    _ = v34
                }
                _ = fi32
            }
        case "":
            for fi35 := 0; fi35 < 12; fi35++ {
                v4 = (v4 + v26)
                for fi36 := 0; fi36 < 20; fi36++ {
                    v37 := GenS2{F0: 17, F1: 255}
                    fmt.Println(v3, v5, v21)
                    v5 = (-(v2))
                    _ = fi36
                    _ = v37
                }
                _ = fi35
            }
        case "\n\t":
            switch v4 {
            case "hello":
                switch (v31 | v31) {
                case -1:
                    var v38 string = "foo bar"
                    _ = v38
                default:
                    var v39 int16 = 32767
                    _ = v39
                }
            case "foo bar":
                if len(v30) > 0 {
                    v30[0] = v31
                }
            case "\n\t":
                var v40 int8 = -1
                _ = v40
            }
        default:
            v2 += 640
        }
    } else {
        v2 = (v3 << ((42) & 63))
        var v41 uint32 = 65535
        switch -457 {
        case 42:
            v29 += v29
        }
        _ = v41
    }
    var v42 float64 = ((float64(3.14159) * ((float64(0.0) - v21) / v21)) - ((float64(0.1) * ((1e10 - (v21 / v21)) / ((v21 * v21) / v21))) - v21))
    v5 -= (v2 >> ((255) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v42
    _ = v5
}
