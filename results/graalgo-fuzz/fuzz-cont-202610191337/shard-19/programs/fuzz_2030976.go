package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int64 }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }
func genFn0() uint {
    v4 := []int{42}
    v5 := GenS1{F0: 450, F1: -928, F2: -996}
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: -2147483648, F1: 37, F2: -1}
        _ = fi6
        _ = v7
    }
    _ = v4
    _ = v5
    return 725
}

func main() {
    var v1 string = ""
    var v2 float64 = 1.5
    var v3 uint32 = 65535
    var v8 uint = 65535
    fmt.Println(v3)
    if true {
        if false {
            var v9 uint16 = 1
            fmt.Println(v2, v8)
            _ = v9
        }
        fmt.Println(v1, v3)
    } else {
        var v10 uint8 = 128
        var v11 int64 = -9223372036854775808
        switch int(v3) {
        case 128:
            var v12 uint8 = (v10 + (^(v10)))
            _ = v12
        case 32768:
            for fi13 := 0; fi13 < 7; fi13++ {
                if (v1 <= v1) {
                    var v14 uint8 = v10
                    _ = v14
                }
                _ = fi13
            }
        case 84:
            v15 := []int{-2147483648, int(v10), -1, 127}
            _ = v15
        }
        _ = v10
        _ = v11
    }
    fmt.Println(v1, v2)
    v16 := []int{}
    if false {
        v2 = (((0.0 - 3.14159) + -2.25) / v2)
        fmt.Println(v2, v3)
    }
    switch 32768 {
    case 9223372036854775807:
        fmt.Println(v1)
    case -969:
        fmt.Println(v3, v8)
    case 878:
        var v17 uint16 = 42
        _ = v17
    default:
        v8 = (v8 + (^(v8)))
    }
    switch 9223372036854775807 {
    case 4294967295:
        v2 = (v2 * 0.0)
    case 637:
        switch 907 {
        case 32767:
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        case -2147483648:
            fmt.Println(v8)
        case 42:
            v18 := GenS1{F0: int32(v3), F1: -32768, F2: 1}
            _ = v18
        }
    case -30:
        if (1.5 <= v2) {
            v1 = (((v1 + v1) + v1) + ("\n\t" + v1))
            var v19 float32 = float32(v2)
            _ = v19
        }
    default:
        switch 21 {
        case -486:
            if (4294967295 >= 18446744073709551615) {
                v20 := GenS1{F0: 0, F1: -32768, F2: -9223372036854775808}
                _ = v20
            }
        case -637:
            v1 = ""
        case 0:
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 uint8 = 42
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 float32 = 0.0
                    _ = fi23
                    _ = v24
                }
                if (0 >= uint64(v3)) {
                    switch (v1 + "a") {
                    case "":
                        var v25 float64 = v2
                        _ = v25
                    case "foo bar":
                        _ = len(v16)
                    default:
                        for fi26 := 0; fi26 < 7; fi26++ {
                            if len(v16) > 0 {
                                v16[0] = fi21
                            }
                            v8 = 677
                            _ = fi26
                        }
                    }
                    fmt.Println(v22, fi21, v2)
                    fmt.Println(v3)
                }
                _ = fi21
                _ = v22
            }
        }
    }
    if len(v16) > 0 {
        v16[0] = 128
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        if false {
            if len(v16) > 0 {
                v16[0] = (fi27 >> ((v8) & 63))
            }
            v2 = v2
        }
        _ = fi27
    }
    var v28 uint8 = 255
    if false {
        v29 := GenS1{F0: -588, F1: 270, F2: 878}
        v30 := GenS1{F0: -2147483648, F1: 0, F2: 1}
        v31 := GenS1{F0: 2147483647, F1: -1, F2: 1}
        _ = v29
        _ = v30
        _ = v31
    } else {
        v28 = (v28 / (((v28 ^ uint8(v3))) | 1))
    }
    if ((v3 * 682) <= (^(v3))) {
        v28 += v28
        var v32 int64 = int64(v28)
        v16 = append(v16, 32768)
        _ = v32
    } else {
        for fi33 := 0; fi33 < 11; fi33++ {
            v1 = "\n\t"
            var v34 int64 = 9223372036854775807
            _ = fi33
            _ = v34
        }
    }
    v16 = append(v16, 9223372036854775807)
    v35 := GenS2{F0: 1, F1: "", F2: 242}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v8
}
