package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint32 {
    v7 := []int{128}
    v7 = append(v7, 9223372036854775807)
    var v8 float32 = (1.5 * 3.14159)
    var v9 int32 = 42
    _ = v7
    _ = v8
    _ = v9
    return 195
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]int)
    var v3 uint64 = 0
    var v4 uint16 = 1
    var v5 int = 9223372036854775807
    v6 := []int{v5, (v5 << ((65535) & 63)), (v5 + v5)}
    if ((^(v3)) != 18446744073709551615) {
        if true {
            var v10 int16 = -454
            switch v5 {
            case 128:
                var v11 int64 = -260
                _ = v11
            default:
                v12 := GenS1{F0: -109}
                _ = v12
            }
            _ = v10
        }
    } else {
        if false {
            v6 = append(v6, -1)
            if ((v1 + (v1 % ((uint8(v4)) | 1))) != v1) {
                v13 := GenS1{F0: -66}
                v5 += v5
                v6 = append(v6, 4294967295)
                _ = v13
            }
        }
        var v14 string = "foo bar"
        v15 := GenS1{F0: 100}
        _ = v14
        _ = v15
    }
    if false {
        fmt.Println(v1, v3, v1)
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                _ = len(v6)
                _ = fi17
            }
            switch "" {
            case "":
                if (-749 > -273) {
                    if (1 <= 255) {
                        v3 -= (v3 - (^(v3)))
                        if (1e10 >= float64(0.0)) {
                            if false {
                                _ = len(v6)
                            }
                        }
                    }
                    v6 = append(v6, -808)
                }
            case "hello":
                for fi18 := 0; fi18 < 7; fi18++ {
                    v6 = append(v6, -916)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v19 float32 = ((1e10 * (1e10 * float32((3.14159 - float64(3.14159))))) + 1.5)
                    fmt.Println(v4)
                    _ = fi18
                    _ = v19
                }
            case "\n\t":
                fmt.Println(v4)
            default:
                v20 := v2[1]
                _ = v20
            }
            _ = fi16
        }
    } else {
        switch (^(v5)) {
        case 32768:
            var v21 uint = 42
            _ = v21
        }
    }
    v22 := v2[84]
    v5 = 32768
    if (916 > 692) {
        var v23 uint64 = v3
        var v24 float32 = 0.0
        v23 -= v23
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 61}
    var v26 int = -1
    var v27 int = v5
    var v28 int32 = -2147483648
    for fi29 := 0; fi29 < 12; fi29++ {
        if (false || false) {
            var v30 float32 = (0.0 + 0.1)
            _ = v30
        } else {
            v5 = v22
            var v31 uint8 = (v1 + (v1 | 2))
            _ = v31
        }
        v28 -= -984
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
