package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16) int32 {
    p0 += 2
    var v5 float64 = (0.1 - 0.1)
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 42
    v3 := []int{803, 128}
    var v4 int8 = 0
    if (0 > (v4 + 127)) {
        if false {
            v4 += 1
            if (18446744073709551615 < 1) {
                fmt.Println(v2, v2)
                switch 42 {
                case 9223372036854775807:
                    v3 = append(v3, 970)
                case -324:
                    fmt.Println(v4)
                default:
                    var v6 float64 = 0.0
                    _ = v6
                }
                v7 := GenS1{F0: 18446744073709551615, F1: 2, F2: false}
                _ = v7
            }
            var v8 float32 = 3.14159
            _ = v8
        }
    }
    switch 32767 {
    case 464:
        var v9 uint8 = 57
        _ = v9
    case -14:
        var v10 uint32 = 885
        _ = v10
    default:
        v3 = append(v3, 32768)
    }
    v11 := v1["\n\t"]
    if false {
        var v12 int64 = 9223372036854775807
        _ = v12
    }
    var v13 int16 = -666
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            v4 = (v4 % ((v4) | 1))
            var v16 float32 = 1.5
            _ = fi15
            _ = v16
        }
        if false {
            switch "\n\t" {
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                v17 := GenS2{F0: v13}
                _ = v17
            case "foo bar":
                var v18 float64 = -2.25
                _ = v18
            default:
                var v19 string = ("hello" + ((("\n\t" + "a") + "\n\t") + "hello"))
                _ = v19
            }
        }
        v20 := GenS2{F0: (v13 + 0)}
        _ = fi14
        _ = v20
    }
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    v4 = (v4 + int8(v2))
    for fi21 := 0; fi21 < 14; fi21++ {
        v11 += 42
        _ = fi21
    }
    if false {
        switch (("" + "a") + "hello") {
        case "\n\t":
            var v22 uint32 = 4294967295
            _ = v22
        default:
            v13 = -1
        }
        if true {
            v3 = append(v3, (v11 << ((65535) & 63)))
            for fi23 := 0; fi23 < 11; fi23++ {
                v24 := GenS2{F0: (v13 << ((1) & 63))}
                _ = fi23
                _ = v24
            }
            var v25 uint = 0
            _ = v25
        } else {
            for fi26 := 0; fi26 < 6; fi26++ {
                v27 := GenS2{F0: -909}
                var v28 uint = 42
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    } else {
        var v29 bool = true
        fmt.Println(v29)
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
