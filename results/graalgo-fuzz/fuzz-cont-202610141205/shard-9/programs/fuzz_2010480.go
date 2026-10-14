package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 uint; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int }

func main() {
    var v1 float64 = (0.0 - 3.14159)
    v2 := make(map[string]uint32)
    var v3 int16 = 993
    v4 := make(map[string]string)
    var v5 int = -255
    for fi6 := 0; fi6 < 14; fi6++ {
        if (4294967295 < 1) {
            if true {
                _ = len(v4)
            }
            v7 := GenS1{F0: 9223372036854775807, F1: 1}
            _ = v7
        }
        switch v5 {
        case 9223372036854775807:
            var v8 float64 = 1e10
            _ = v8
        default:
            v9 := GenS2{F0: (v3 | (^(v3))), F1: 879, F2: ""}
            _ = v9
        }
        v3 += v3
        _ = fi6
    }
    v10 := GenS1{F0: -9223372036854775808, F1: -122}
    v11 := GenS3{F0: 721}
    if false {
        v12 := v4["hello"]
        switch v12 {
        case "a":
            v12 = (((v12 + v12) + v12) + v12)
        case "foo bar":
            var v13 uint64 = 200
            _ = v13
        }
        _ = v12
    } else {
        if false {
            v14 := GenS4{F0: 1, F1: (^(v5))}
            var v15 int = -1
            v16 := GenS4{F0: 42, F1: (v5 << ((4294967295) & 63))}
            _ = v14
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 14; fi17++ {
                if true {
                    v3 -= (v3 | 1000)
                }
                _ = fi17
            }
            fmt.Println(v5)
            var v18 int16 = 515
            _ = v18
        }
    }
    v19 := []int{v5, 1, (v5 / (((^(v5))) | 1)), 2147483647}
    fmt.Println(v5)
    v3 -= v3
    v3 = v3
    switch v5 {
    case -256:
        var v20 uint64 = 42
        _ = v20
    case 0:
        var v21 int16 = v3
        _ = v21
    case -61:
        if (!(-895 >= -589)) {
            switch (v5 << ((276) & 63)) {
            case -9223372036854775808:
                var v22 uint = 42
                _ = v22
            case -70:
                fmt.Println(v3)
            case 380:
                for fi23 := 0; fi23 < 20; fi23++ {
                    _ = len(v19)
                    if (318 == 2147483647) {
                        if true {
                            v3 += 1000
                            v1 = (-2.25 - 1.5)
                        }
                        fmt.Println(v1, v5, v5)
                        var v24 uint32 = 173
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 14; fi25++ {
                            _ = len(v19)
                            v5 -= int(v3)
                            fmt.Println(fi23)
                            _ = fi25
                        }
                    }
                    _ = fi23
                }
            }
            v1 = (0.1 + v1)
        } else {
            var v26 int16 = 42
            var v27 uint8 = 246
            for fi28 := 0; fi28 < 5; fi28++ {
                v29 := GenS1{F0: 9223372036854775807, F1: 42}
                _ = fi28
                _ = v29
            }
            _ = v26
            _ = v27
        }
    default:
        var v30 uint = 244
        _ = v30
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        if true {
            v3 += v3
            v19 = append(v19, (v5 | (^(v5))))
        }
        for fi32 := 0; fi32 < 19; fi32++ {
            v5 -= (fi32 * -488)
            _ = fi32
        }
        v1 = (0.1 * (0.1 + float64(float32(v1))))
        _ = fi31
    }
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
