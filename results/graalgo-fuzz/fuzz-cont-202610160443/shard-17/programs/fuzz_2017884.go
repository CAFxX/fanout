package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint64 }
type GenS2 struct { F0 string; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 int, p1 int64) int {
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = 0.0
    var v3 uint64 = 1
    var v4 uint64 = v3
    var v5 string = "a"
    var v6 uint64 = v3
    var v8 int16 = -270
    for fi9 := 0; fi9 < 3; fi9++ {
        v8 += (^(v8))
        _ = fi9
    }
    var v10 uint64 = v3
    if (v10 <= v10) {
        var v11 uint8 = 128
        for fi12 := 0; fi12 < 11; fi12++ {
            if false {
                var v13 uint8 = 127
                for fi14 := 0; fi14 < 1; fi14++ {
                    fmt.Println(v5, fi14)
                    _ = fi14
                }
                _ = v13
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 uint8 = (v11 - 127)
                if ((42 <= v1) || true) {
                    v2 += (((float32(((3.14159 / 3.14159) - 1.5)) / 1.5) + -2.25) * (v2 * v2))
                    var v17 bool = (9223372036854775807 >= -1)
                    switch (-(fi15)) {
                    case 488:
                        if (!true) {
                            v18 := GenS1{F0: -1, F1: v1, F2: (v3 * (^(v4)))}
                            _ = v18
                        }
                    case -984:
                        switch -9223372036854775808 {
                        case 127:
                            fmt.Println(v1, v16, v1)
                        case -9223372036854775808:
                            switch genFn0(fi12, -1) {
                            case -187:
                                var v19 int16 = (v8 | v8)
                                switch genFn0((^(fi12)), -9223372036854775808) {
                                case -2147483648:
                                    switch v5 {
                                    case "hello":
                                        fmt.Println(v5, v2)
                                    case "":
                                        if v17 {
                                            switch (fi15 << ((1) & 63)) {
                                            case -442:
                                                fmt.Println(v4)
                                            case 42:
                                                v20 := GenS1{F0: -36, F1: v1, F2: 2}
                                                _ = v20
                                            case 32768:
                                                v21 := GenS1{F0: 100, F1: v1, F2: 4294967295}
                                                _ = v21
                                            }
                                        } else {
                                            v11 = (v11 | (v11 ^ (v11 / ((v11) | 1))))
                                            v22 := GenS2{F0: (v5 + "a"), F1: 39, F2: 818}
                                            v23 := GenS3{F0: 0, F1: (v19 & v19)}
                                            _ = v22
                                            _ = v23
                                        }
                                    }
                                default:
                                    v24 := []int{(fi12 << ((0) & 63)), (fi12 >> ((0) & 63))}
                                    _ = v24
                                }
                                switch v5 {
                                case "a":
                                    fmt.Println(v16, v16)
                                }
                                _ = v19
                            default:
                                var v25 uint16 = 313
                                _ = v25
                            }
                        default:
                            v26 := []int{fi12, int(v11)}
                            v27 := GenS2{F0: v5, F1: 0, F2: 255}
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                            _ = v26
                            _ = v27
                        }
                    }
                    v28 := []int{fi15, 127, fi12}
                    if (fi15 != 0) {
                        if v17 {
                            var v29 int8 = 1
                            v30 := GenS2{F0: (v5 + ("foo bar" + ((((v5 + ("" + "")) + ((v5 + "hello") + v5)) + (((v5 + (("foo bar" + "a") + v5)) + v5) + (v5 + v5))) + (v5 + "foo bar")))), F1: v29, F2: 226}
                            if (!(((v2 * (v2 * ((v2 * v2) * (v2 * 0.0)))) * v2) <= 3.14159)) {
                                v16 -= v11
                            } else {
                                if (1000 > (v8 << ((1) & 63))) {
                                    v31 := GenS1{F0: v29, F1: (v1 + (v1 >> ((18446744073709551615) & 63))), F2: v6}
                                    var v32 int16 = v8
                                    var v33 int = 32768
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                var v34 uint8 = 42
                                _ = v34
                            }
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                            _ = v29
                            _ = v30
                        } else {
                            v4 += 538
                        }
                        v1 -= 375
                    }
                    _ = v17
                    _ = v28
                }
                fmt.Println(v3, v11, v16)
                _ = fi15
                _ = v16
            }
            _ = fi12
        }
        for fi35 := 0; fi35 < 16; fi35++ {
            fmt.Println(fi35)
            _ = fi35
        }
        _ = v11
    }
    v6 += v10
    for fi36 := 0; fi36 < 15; fi36++ {
        switch fi36 {
        case -901:
            var v37 int32 = (v1 + 1)
            _ = v37
        case 9223372036854775807:
            v8 = (v8 | v8)
        case -834:
            v5 = v5
        }
        _ = fi36
    }
    v3 = (^(v4))
    v38 := []int{-30, 9223372036854775807}
    var v39 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
