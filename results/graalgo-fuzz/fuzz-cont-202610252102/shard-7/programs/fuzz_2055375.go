package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint8; F2 int64 }
func genFn0() uint64 {
    var v4 float32 = 0.0
    v4 = v4
    var v5 uint32 = 0
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []string{}
    var v2 int32 = 0
    var v3 int32 = 672
    var v6 int = -2147483648
    fmt.Println(v6)
    for fi7 := 0; fi7 < 13; fi7++ {
        if (!false) {
            var v8 int16 = -32768
            v2 -= (v3 - (v3 << ((4294967295) & 63)))
            var v9 bool = false
            _ = v8
            _ = v9
        } else {
            switch ("\n\t" + "\n\t") {
            case "":
                var v10 uint32 = 42
                _ = v10
            default:
                if (208 != 2) {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        v1 = append(v1, "a")
                        var v12 uint = 255
                        var v13 float64 = 1e10
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    var v14 uint8 = 255
                    for fi15 := 0; fi15 < 3; fi15++ {
                        if len(v1) > 0 {
                            v1[0] = ("" + "a")
                        }
                        var v16 uint32 = 65535
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                }
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                v3 -= v3
                _ = fi17
            }
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 float64 = float64(0.0)
            _ = fi18
            _ = v19
        }
        _ = fi7
    }
    v20 := GenS1{F0: v6}
    switch "\n\t" {
    case "\n\t":
        switch v6 {
        case -2147483648:
            v1 = append(v1, ("" + "hello"))
        case 42:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    default:
        fmt.Println(v3)
    }
    v21 := GenS1{F0: (-(v6))}
    v22 := GenS2{F0: 32767, F1: 255, F2: -9223372036854775808}
    switch (v6 ^ (v6 & 0)) {
    case 9223372036854775807:
        var v23 float64 = (float64(-2.25) * 1.5)
        _ = v23
    default:
        v6 = v6
    }
    var v24 int8 = -97
    v6 = (-(v6))
    v3 = v2
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v6
}
