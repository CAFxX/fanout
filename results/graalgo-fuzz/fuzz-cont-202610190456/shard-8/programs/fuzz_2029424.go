package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int32 }

func main() {
    var v1 float32 = float32(0.0)
    var v2 uint32 = 42
    v3 := make(map[int]uint32)
    var v4 float32 = v1
    v4 = v1
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 11; fi6++ {
        switch ("a" + "") {
        case "a":
            switch -2147483648 {
            case 4294967295:
                fmt.Println(fi6, fi6, v1)
            case 127:
                var v7 int8 = 127
                _ = v7
            default:
                switch -623 {
                case 270:
                    v2 = (^(v2))
                case 788:
                    for fi8 := 0; fi8 < 3; fi8++ {
                        v2 = 931
                        fmt.Println(v5)
                        _ = fi8
                    }
                case 32768:
                    var v9 float32 = v4
                    _ = v9
                default:
                    var v10 uint16 = 368
                    _ = v10
                }
            }
        case "foo bar":
            var v11 uint16 = 42
            _ = v11
        case "hello":
            switch fi6 {
            case -655:
                v12 := GenS1{F0: 100, F1: 472}
                _ = v12
            case 357:
                switch (-(fi6)) {
                case 128:
                    fmt.Println(v1, v4)
                default:
                    switch (fi6 - (fi6 | 172)) {
                    case -2147483648:
                        switch (fi6 | fi6) {
                        case -2147483648:
                            for fi13 := 0; fi13 < 17; fi13++ {
                                v14 := GenS1{F0: 30, F1: 2}
                                var v15 int16 = int16(fi6)
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        default:
                            if false {
                                v4 += 0.0
                                var v16 bool = (((3.14159 == (((0.1 - 0.1) * 1e10) * float64((float32(float64(v1)) * v4)))) && false) || ((3.14159 + (3.14159 + 0.0)) < 3.14159))
                                var v17 int32 = v5
                                _ = v16
                                _ = v17
                            }
                        }
                    case 32767:
                        var v18 int64 = -1
                        _ = v18
                    case -860:
                        v19 := GenS2{F0: 1, F1: v5}
                        _ = v19
                    }
                }
            case 128:
                v5 = (^(v5))
            }
        default:
            var v20 uint32 = 2147483647
            _ = v20
        }
        v1 = ((v1 / 1.5) - (v1 * float32(1e10)))
        var v21 float64 = (((((3.14159 - (1e10 * -2.25)) / (0.1 + ((3.14159 - 0.0) / 3.14159))) + (0.0 + 0.0)) - (0.1 / 1.5)) - 0.1)
        _ = fi6
        _ = v21
    }
    var v22 uint16 = 255
    fmt.Println(v2, v2, v5)
    var v23 int = -79
    v24 := v3[(v23 & (v23 << ((915) & 63)))]
    var v25 string = ("\n\t" + ((("" + ((("" + "a") + ("foo bar" + "hello")) + "")) + "hello") + "\n\t"))
    if (float64((v1 / 3.14159)) != (((1e10 - 3.14159) / ((3.14159 / (((0.0 - 1e10) * ((1.5 * ((-2.25 / 1.5) - -2.25)) / 1.0)) / 3.14159)) / 0.1)) * (float64(v4) - float64(v1)))) {
        var v26 uint32 = 864
        v27 := GenS1{F0: -128, F1: 18446744073709551615}
        _ = v26
        _ = v27
    } else {
        if false {
            v3[791] = v2
            v28 := []int{v23, int(v22)}
            _ = v28
        }
        var v29 int64 = 42
        var v30 int64 = v29
        _ = v29
        _ = v30
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
