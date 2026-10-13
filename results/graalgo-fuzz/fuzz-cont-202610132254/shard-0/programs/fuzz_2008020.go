package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 int32 }
func genFn0(p0 bool) uint8 {
    fmt.Println(p0)
    p0 = true
    _ = p0
    return 2
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 1
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    v5 := []string{"hello", "\n\t"}
    v6 := v3[1]
    if (3.14159 == 3.14159) {
        var v7 uint = 974
        v6 = "a"
        _ = v7
    } else {
        v1 += v1
        for fi8 := 0; fi8 < 2; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                v1 -= v1
                _ = fi9
            }
            var v10 uint32 = 4294967295
            _ = fi8
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        switch fi11 {
        case -1:
            v12 := GenS1{F0: 42, F1: uint16(v1), F2: (v2 >> ((255) & 63))}
            _ = v12
        case 0:
            if false {
                var v13 int16 = -32768
                _ = v13
            } else {
                if false {
                    v1 = v1
                }
                v14 := GenS1{F0: int16(v1), F1: 908, F2: uint8(v1)}
                _ = v14
            }
        default:
            if (605 < 443) {
                v1 = (v1 & v1)
                _ = len(v5)
                v2 += v2
            }
        }
        if false {
            var v15 uint = 2
            _ = v15
        }
        _ = fi11
    }
    var v16 string = v6
    fmt.Println(v1, v1, v16)
    _ = len(v3)
    if true {
        v5 = append(v5, (v16 + v16))
    } else {
        if true {
            v2 = (v2 | (v2 + (^(v2))))
            v16 = v6
            var v17 int64 = -116
            _ = v17
        } else {
            v18 := v3[330]
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = len(v4)
            _ = v18
        }
        v19 := GenS1{F0: 1, F1: 536, F2: 255}
        var v20 uint16 = 1
        _ = v19
        _ = v20
    }
    v3[-484] = ("a" + v16)
    v21 := GenS1{F0: 0, F1: 484, F2: (v2 & 1)}
    var v22 float64 = 3.14159
    v16 = (v6 + (v16 + ""))
    _ = len(v3)
    v23 := GenS2{F0: -1, F1: 0, F2: 1}
    var v24 int = -2147483648
    if len(v5) > 0 {
        v5[0] = v6
    }
    switch (v24 << ((980) & 63)) {
    case -2147483648:
        if false {
            _ = len(v5)
        }
    case 32767:
        fmt.Println(v6, v24)
    case 9223372036854775807:
        if (255 <= 18446744073709551615) {
            var v25 string = v6
            _ = v25
        } else {
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 uint32 = 398
                _ = fi26
                _ = v27
            }
            switch "" {
            case "":
                if false {
                    for fi28 := 0; fi28 < 18; fi28++ {
                        v2 = 0
                        _ = fi28
                    }
                    var v29 float64 = float64(((0.0 * 3.14159) / (0.1 * 1e10)))
                    _ = v29
                } else {
                    v4[("" + v6)] = ("" + "")
                    var v30 float32 = 0.1
                    v24 -= 4294967295
                    _ = v30
                }
            }
        }
    }
    var v31 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
