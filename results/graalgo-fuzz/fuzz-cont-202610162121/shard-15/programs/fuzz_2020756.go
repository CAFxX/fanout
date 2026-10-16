package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0(p0 int64, p1 uint64, p2 uint16) string {
    fmt.Println(p1)
    var v5 uint16 = p2
    var v6 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ("\n\t" + "")
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 32767
    var v3 int = 32768
    var v4 bool = true
    v7 := []int{(v3 / ((v3) | 1))}
    for fi8 := 0; fi8 < 20; fi8++ {
        if false {
            if (!v4) {
                var v9 uint8 = 210
                fmt.Println(v3)
                if ((float32(1.5) + float32(1e10)) > (1.5 + -2.25)) {
                    v2 = (v2 ^ (v2 + v2))
                    v4 = (0.0 >= -2.25)
                }
                _ = v9
            } else {
                v10 := GenS1{F0: int64(v3), F1: 66}
                _ = v10
            }
        } else {
            v2 = v2
        }
        var v11 int32 = 2147483647
        _ = fi8
        _ = v11
    }
    if v4 {
        _ = len(v7)
    } else {
        if v4 {
            fmt.Println(v4, v3)
            var v12 int16 = -430
            var v13 int32 = -402
            _ = v12
            _ = v13
        }
        if v4 {
            switch v3 {
            case -51:
                var v14 uint = 866
                _ = v14
            }
        }
    }
    fmt.Println(v3)
    var v15 uint64 = 931
    var v16 int16 = -1
    var v17 int64 = -9223372036854775808
    if (v4 && ("foo bar" < "")) {
        if (false || v4) {
            switch (v3 & (v3 | -260)) {
            case 9223372036854775807:
                fmt.Println(v15, v4)
            default:
                v16 = v16
            }
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 bool = ((128 != 226) || false)
                v20 := GenS1{F0: (v17 & (v17 / ((9223372036854775807) | 1))), F1: -1}
                _ = fi18
                _ = v19
                _ = v20
            }
        }
        switch "hello" {
        case "foo bar":
            var v21 uint8 = 185
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 17; fi22++ {
            switch (fi22 | v3) {
            case 9223372036854775807:
                var v23 uint8 = 255
                _ = v23
            case 32767:
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 uint32 = 65535
                    if true {
                        v15 = (v15 - v15)
                        v4 = ((v4 && v4) && (((!("" <= ((("" + ("foo bar" + ((((("a" + "hello") + ("foo bar" + "hello")) + ("foo bar" + "hello")) + "foo bar") + ("hello" + "foo bar")))) + "a") + "\n\t"))) || v4) || (false || (!(-2147483648 >= 0)))))
                    } else {
                        var v26 uint32 = (v25 & v25)
                        _ = v26
                    }
                    fmt.Println(v16)
                    v16 -= 0
                    v3 = 4294967295
                    var v27 uint64 = (v15 - v15)
                    _ = fi24
                    _ = v25
                    _ = v27
                }
            case -636:
                fmt.Println(v15)
            default:
                v28 := GenS2{F0: (^(v15)), F1: (v16 % ((v16) | 1))}
                _ = v28
            }
            fmt.Println(v15, v3)
            _ = fi22
        }
        v17 -= v17
    }
    if ((false || (uint8(v15) > 64)) && (float32((((((float64(0.1) + (((-2.25 + 1e10) + (-2.25 * (1e10 * 0.1))) - (3.14159 * 3.14159))) + ((0.1 / 3.14159) - ((((1e10 + -2.25) / 0.1) / ((1.5 - 3.14159) + 1.5)) / 1.0))) * float64((1.5 * 0.0))) - 3.14159) - (3.14159 / (3.14159 / 3.14159)))) >= (0.0 / 1.0))) {
        switch v3 {
        case 2147483647:
            v29 := GenS1{F0: v17, F1: 1}
            _ = v29
        case 181:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case 1:
            var v30 uint = 774
            _ = v30
        default:
            v3 = (v3 + (v3 / ((v3) | 1)))
        }
    } else {
        v17 += (v17 + (v17 - (v17 ^ 9223372036854775807)))
        if true {
            fmt.Println(v2, v16)
            v17 += (^(v17))
        }
        var v31 int = v3
        _ = v31
    }
    switch (v3 / ((127) | 1)) {
    case 882:
        fmt.Println(v2, v3)
    case 0:
        v3 -= 607
    }
    v17 -= 0
    var v32 uint64 = v15
    v15 += 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v7
}
