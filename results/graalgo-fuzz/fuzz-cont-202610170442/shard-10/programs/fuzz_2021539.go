package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{-9223372036854775808, -66, -41, 401}
    var v3 int = -2147483648
    var v4 int16 = -334
    fmt.Println(v4, v3)
    var v5 int32 = -59
    for fi6 := 0; fi6 < 11; fi6++ {
        v4 += v4
        v2 = append(v2, (fi6 * 0))
        var v7 float32 = (float32((float64((1.5 / (1e10 / (0.0 + 1e10)))) - (0.0 - 1.5))) / 0.1)
        _ = fi6
        _ = v7
    }
    v4 -= v4
    v4 = v4
    if len(v2) > 0 {
        v2[0] = (^(v3))
    }
    if len(v2) > 0 {
        v2[0] = (v3 << ((18446744073709551615) & 63))
    }
    v2 = append(v2, 411)
    switch ("a" + "") {
    case "foo bar":
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := GenS1{F0: (fi9 * 1), F1: 535, F2: "\n\t"}
                v3 += 42
                v5 += 903
                delete(v1, (fi9 | fi9))
                _ = fi9
                _ = v10
            }
            v11 := GenS2{F0: v5, F1: ((!false) || false), F2: 476}
            _ = fi8
            _ = v11
        }
    default:
        fmt.Println(v4, v4)
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        v5 -= 1
        if false {
            if true {
                var v13 int64 = 746
                for fi14 := 0; fi14 < 8; fi14++ {
                    if false {
                        v13 = v13
                        v4 = (v4 % ((-1) | 1))
                        var v15 bool = true
                        var v16 uint64 = 10
                        _ = v15
                        _ = v16
                    } else {
                        switch 42 {
                        case 0:
                            v17 := GenS2{F0: -770, F1: true, F2: 907}
                            _ = v17
                        case -2147483648:
                            var v18 float64 = (-2.25 + 0.0)
                            _ = v18
                        default:
                            if false {
                                var v19 string = "foo bar"
                                _ = v19
                            } else {
                                v4 += (v4 | v4)
                            }
                        }
                        if (!false) {
                            var v20 uint8 = 2
                            switch -9223372036854775808 {
                            case -2147483648:
                                var v21 int16 = (v4 % ((v4) | 1))
                                _ = v21
                            }
                            v22 := GenS3{F0: v13}
                            var v23 string = "a"
                            v24 := GenS3{F0: int64(v5)}
                            var v25 int8 = -128
                            _ = v20
                            _ = v22
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            v5 = 0
                            var v26 uint16 = 42
                            var v27 uint64 = 42
                            _ = v26
                            _ = v27
                        }
                        v28 := GenS2{F0: v5, F1: (float32(-2.25) < -2.25), F2: 2147483648}
                        _ = v28
                    }
                    _ = len(v2)
                    v3 = -95
                    _ = fi14
                }
                v5 -= (v5 | (v5 << ((4294967295) & 63)))
                _ = v13
            }
        } else {
            if (!false) {
                switch (("\n\t" + "") + "hello") {
                case "a":
                    v29 := GenS3{F0: -9223372036854775808}
                    _ = v29
                case "hello":
                    v30 := GenS4{F0: 9223372036854775807}
                    _ = v30
                }
                var v31 uint32 = 65535
                v3 += -800
                _ = v31
            }
            if (v4 < (v4 % (((^(v4))) | 1))) {
                for fi32 := 0; fi32 < 5; fi32++ {
                    fmt.Println(fi12, fi32, v4)
                    if len(v2) > 0 {
                        v2[0] = (fi12 << ((42) & 63))
                    }
                    switch fi12 {
                    case 189:
                        switch (-(fi32)) {
                        case 42:
                            _ = len(v2)
                        case -1:
                            var v33 int16 = v4
                            _ = v33
                        }
                    }
                    _ = fi32
                }
                var v34 int = (fi12 / (((fi12 - fi12)) | 1))
                _ = v34
            }
            _ = len(v1)
        }
        v4 = -32768
        _ = fi12
    }
    v5 = (-(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
