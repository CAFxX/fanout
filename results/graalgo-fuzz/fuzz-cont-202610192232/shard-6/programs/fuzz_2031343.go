package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint64 }
func genFn0(p0 int64) uint {
    switch -9223372036854775808 {
    case 1:
        switch 127 {
        case 32768:
            var v5 float64 = ((3.14159 - 1e10) - -2.25)
            _ = v5
        case -170:
            p0 = p0
        case -2147483648:
            v6 := GenS1{F0: 0}
            _ = v6
        }
    }
    p0 = (p0 << ((65535) & 63))
    _ = p0
    return 301
}

func main() {
    var v1 int16 = 1
    v2 := make(map[string]uint32)
    v3 := []string{}
    v4 := []string{}
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint64 = 771
        fmt.Println(v1, v8)
        switch "" {
        case "":
            v9 := GenS2{F0: 0, F1: 255, F2: 709}
            _ = v9
        default:
            switch "a" {
            case "\n\t":
                v8 = 4294967295
            case "":
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch (("hello" + "a") + "") {
                    case "hello":
                        v8 -= (v8 << ((18446744073709551615) & 63))
                    case "":
                        var v11 int32 = 0
                        _ = v11
                    case "a":
                        var v12 int = (^(fi7))
                        _ = v12
                    default:
                        var v13 float32 = 0.0
                        _ = v13
                    }
                    _ = fi10
                }
            default:
                for fi14 := 0; fi14 < 8; fi14++ {
                    fmt.Println(fi14)
                    var v15 float64 = ((((float64(((1e10 * 0.1) * (0.0 - ((0.1 * 1e10) / 3.14159)))) * ((0.0 / 3.14159) - 3.14159)) * (0.1 + 0.1)) + 3.14159) - -2.25)
                    _ = fi14
                    _ = v15
                }
            }
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v1)
    v2["hello"] = 276
    for fi16 := 0; fi16 < 10; fi16++ {
        v1 = v1
        for fi17 := 0; fi17 < 13; fi17++ {
            fmt.Println(v1, fi17, v1)
            fmt.Println(fi16, fi16, v1)
            v1 = v1
            _ = fi17
        }
        _ = fi16
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(v1, fi18)
        _ = fi18
    }
    var v19 float64 = ((1e10 * 1.5) - 0.1)
    var v20 string = "\n\t"
    if true {
        v1 = (v1 + (v1 | v1))
        var v21 int16 = 0
        _ = v21
    }
    fmt.Println(v20, v1)
    if false {
        v1 -= 1
    }
    v20 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
