package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint16; F2 int64 }
func genFn0() bool {
    var v4 int8 = -1
    switch "" {
    case "a":
        var v5 uint32 = 2147483648
        _ = v5
    case "":
        v6 := GenS1{F0: 32767}
        _ = v6
    case "hello":
        var v7 uint16 = 224
        _ = v7
    default:
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int32 = 1
            var v10 uint32 = 4294967295
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    _ = v4
    return ((float32(3.14159) * 1.5) <= (1e10 + ((-2.25 / 3.14159) + 3.14159)))
}

func main() {
    var v1 uint32 = 2
    var v2 float64 = ((1.5 + 1e10) + (1e10 * 1e10))
    var v3 int = 32768
    fmt.Println(v1, v2)
    fmt.Println(v1)
    var v11 float64 = v2
    switch v3 {
    case 1:
        var v12 int = (v3 - (v3 + (v3 / (((v3 % ((int(v1)) | 1))) | 1))))
        _ = v12
    case -2147483648:
        v13 := []int{(^(v3))}
        _ = v13
    default:
        if (297 != -2147483648) {
            var v14 float64 = v2
            v2 += ((3.14159 / v14) / ((float64((0.1 * (((3.14159 * (0.0 + float32(v14))) - 0.0) / 1.0))) * -2.25) + (float64(0.1) * 0.1)))
            if (2147483647 == 2147483647) {
                for fi15 := 0; fi15 < 3; fi15++ {
                    switch -1 {
                    case 0:
                        fmt.Println(v2)
                    case 1:
                        var v16 uint32 = v1
                        _ = v16
                    case 32767:
                        v3 = (fi15 - fi15)
                    default:
                        if false {
                            var v17 float64 = 0.0
                            _ = v17
                        } else {
                            var v18 bool = (!true)
                            var v19 int = (fi15 % ((v3) | 1))
                            var v20 int32 = -2147483648
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    }
                    v21 := GenS2{F0: 65535}
                    v22 := GenS3{F0: -2147483648, F1: 65535, F2: 0}
                    _ = fi15
                    _ = v21
                    _ = v22
                }
            }
            _ = v14
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 int32 = 42
        var v25 float32 = float32(v2)
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := GenS2{F0: 18446744073709551615}
    var v27 uint64 = 0
    switch "\n\t" {
    case "":
        switch -2147483648 {
        case -9223372036854775808:
            v28 := GenS1{F0: 65535}
            _ = v28
        case 2147483647:
            v29 := GenS3{F0: 0, F1: 546, F2: 0}
            _ = v29
        case 128:
            for fi30 := 0; fi30 < 15; fi30++ {
                v2 = v11
                var v31 bool = (4294967295 > 9223372036854775807)
                var v32 uint16 = 732
                _ = fi30
                _ = v31
                _ = v32
            }
        }
    case "hello":
        var v33 string = ((("\n\t" + "hello") + "foo bar") + ("" + ""))
        _ = v33
    case "\n\t":
        var v34 bool = ((^(v27)) > v27)
        _ = v34
    default:
        v35 := []int{v3}
        _ = v35
    }
    if (0 != 1) {
        var v36 int = v3
        if true {
            var v37 bool = true
            v1 -= (v1 | (v1 & v1))
            _ = v37
        } else {
            v38 := []int{v3, (^(v3)), v3, (v3 * -9223372036854775808)}
            _ = v38
        }
        _ = v36
    } else {
        var v39 int32 = 42
        _ = v39
    }
    v40 := []int{}
    if false {
        v1 -= v1
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
    } else {
        switch "hello" {
        case "":
            v1 -= (v1 - (v1 * (v1 | 1)))
        default:
            if true {
                v3 = (v3 >> ((255) & 63))
            } else {
                var v41 uint8 = 127
                var v42 uint64 = 42
                _ = v41
                _ = v42
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v40
}
