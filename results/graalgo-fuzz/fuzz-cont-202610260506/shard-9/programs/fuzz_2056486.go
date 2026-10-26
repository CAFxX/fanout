package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 int, p1 int8, p2 int32) int64 {
    var v7 int8 = 100
    p0 = (p0 + (^(p0)))
    var v8 int16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -9223372036854775808
}

func main() {
    v1 := []string{}
    var v2 float64 = 1e10
    var v3 int = -739
    v4 := make(map[string]uint32)
    var v5 string = ""
    var v6 uint64 = 2
    var v9 string = (((v5 + (v5 + "hello")) + (v5 + "")) + (v5 + v5))
    switch -9223372036854775808 {
    case 127:
        v10 := GenS1{F0: 42, F1: (v3 > (^(v3))), F2: 1}
        _ = v10
    case -1:
        fmt.Println(v2, v6, v3)
    default:
        v3 = 9223372036854775807
    }
    switch (^(v3)) {
    case -460:
        var v11 float64 = v2
        _ = v11
    case 42:
        switch v3 {
        case 0:
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int16 = -32768
                v14 := GenS1{F0: 1, F1: true, F2: 65535}
                switch v9 {
                case "\n\t":
                    var v15 int = fi12
                    _ = v15
                case "":
                    v16 := GenS1{F0: 1, F1: true, F2: 65535}
                    _ = v16
                case "foo bar":
                    for fi17 := 0; fi17 < 17; fi17++ {
                        v18 := GenS1{F0: 0, F1: ((float64(-2.25) / float64(((1e10 - (float32(float64(float32(3.14159))) / 0.1)) * (3.14159 / float32(3.14159))))) <= 0.0), F2: 935}
                        if (59 < -128) {
                            fmt.Println(v9, v3, v13)
                        } else {
                            var v19 int64 = 9223372036854775807
                            _ = v19
                        }
                        _ = fi17
                        _ = v18
                    }
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        case 9223372036854775807:
            if false {
                v6 += (v6 | (v6 << ((0) & 63)))
                v20 := v4[v5]
                _ = v20
            }
        case -2147483648:
            v9 = (v5 + ("\n\t" + (((v5 + (v5 + "\n\t")) + v9) + v5)))
        default:
            if (131 != 114) {
                v4["hello"] = 42
                var v21 int64 = genFn0(0, -83, 0)
                v22 := GenS2{F0: (v3 << ((65535) & 63)), F1: 1}
                _ = v21
                _ = v22
            }
        }
    case 609:
        v6 = v6
    }
    if false {
        var v23 uint64 = v6
        _ = v23
    }
    v6 = v6
    v5 = v9
    v1 = append(v1, v9)
    for fi24 := 0; fi24 < 19; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            v26 := GenS2{F0: v3, F1: 747}
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
