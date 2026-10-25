package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint16; F2 uint16 }

func main() {
    var v1 int8 = -128
    v2 := make(map[string]uint32)
    var v3 int8 = v1
    var v4 int16 = 549
    if (32767 < -2147483648) {
        v5 := v2["foo bar"]
        var v6 float64 = (3.14159 + 1.5)
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 276, F1: 53}
    v8 := []int{32768, 128, 128, 127}
    var v9 int8 = (v3 | (v3 & (^(v3))))
    v1 -= (v3 + 42)
    switch 4294967295 {
    case 32768:
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS2{F0: 4294967295}
            v3 = (v3 / ((v1) | 1))
            _ = fi10
            _ = v11
        }
    }
    switch "a" {
    case "hello":
        v9 = v3
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        if true {
            v3 = 34
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v13 uint16 = 152
            _ = v13
        } else {
            v1 += -18
            var v14 uint = 255
            v9 -= -1
            _ = v14
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v8 = append(v8, (fi15 * fi15))
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := GenS1{F0: 0, F1: 18}
            var v18 int64 = -518
            fmt.Println(fi16, fi15, v9)
            if (1 < 0) {
                if false {
                    v9 += v3
                    switch "" {
                    case "hello":
                        v3 += 42
                    default:
                        v9 = -128
                    }
                    if false {
                        v2["hello"] = 225
                        var v19 uint16 = 32767
                        if len(v8) > 0 {
                            v8[0] = (fi15 << ((255) & 63))
                        }
                        _ = v19
                    } else {
                        fmt.Println(v9, fi15)
                        v3 -= 42
                    }
                }
                if (v18 != 42) {
                    if (false || (139 > 184)) {
                        v20 := GenS2{F0: -959}
                        v9 = (^(v9))
                        if false {
                            v2["foo bar"] = 65535
                            fmt.Println(v18, v4)
                            v3 += 1
                        } else {
                            switch "a" {
                            case "foo bar":
                                if true {
                                    v4 -= int16(v1)
                                    var v21 uint16 = 42
                                    switch (fi15 * fi16) {
                                    case -9223372036854775808:
                                        if (true || true) {
                                            var v22 int8 = v1
                                            if false {
                                                v23 := GenS3{F0: (^(v4)), F1: (^(v21)), F2: 65535}
                                                _ = v23
                                            } else {
                                                v4 -= 1000
                                                v18 = 1
                                                v21 = v21
                                            }
                                            _ = v22
                                        } else {
                                            v2[("\n\t" + "a")] = 2147483648
                                        }
                                    case 32768:
                                        fmt.Println(fi16)
                                    case -2147483648:
                                        _ = len(v8)
                                    default:
                                        var v24 string = "hello"
                                        _ = v24
                                    }
                                    _ = v21
                                }
                            default:
                                var v25 uint8 = 237
                                _ = v25
                            }
                        }
                        v26 := v2["foo bar"]
                        _ = v20
                        _ = v26
                    }
                    delete(v2, "\n\t")
                    fmt.Println(v3, fi16)
                }
                _ = len(v8)
            }
            fmt.Println(v3)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    if false {
        for fi27 := 0; fi27 < 2; fi27++ {
            v9 -= v3
            _ = fi27
        }
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := GenS1{F0: 1, F1: 0}
            var v30 float32 = 1e10
            _ = fi28
            _ = v29
            _ = v30
        }
    } else {
        if false {
            for fi31 := 0; fi31 < 16; fi31++ {
                var v32 float32 = ((3.14159 / 1.0) / 0.1)
                var v33 uint64 = 4294967295
                _ = fi31
                _ = v32
                _ = v33
            }
            v34 := GenS2{F0: -460}
            for fi35 := 0; fi35 < 18; fi35++ {
                v36 := GenS3{F0: -203, F1: 1, F2: 32767}
                v37 := GenS3{F0: (v4 & v4), F1: 255, F2: 65535}
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = v34
        }
    }
    switch 32768 {
    case 4294967295:
        v3 += (v3 >> ((65535) & 63))
    case 1:
        switch -105 {
        case -9223372036854775808:
            v9 -= (v9 | -87)
        case -1:
            fmt.Println(v4, v9)
        case 2147483647:
            var v38 int16 = -32768
            _ = v38
        default:
            if (42 >= -9223372036854775808) {
                if (2 != 128) {
                    var v39 uint32 = 2147483647
                    var v40 uint8 = 255
                    _ = v39
                    _ = v40
                }
            } else {
                if true {
                    var v41 uint32 = 2
                    var v42 uint16 = 42
                    if false {
                        var v43 float64 = -2.25
                        v1 = (v1 << ((18446744073709551615) & 63))
                        _ = v43
                    }
                    _ = v41
                    _ = v42
                } else {
                    v9 = v9
                }
            }
        }
    }
    var v44 int32 = 2147483647
    fmt.Println(v1, v1)
    fmt.Println(v3, v1, v1)
    v1 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v7
    _ = v8
    _ = v9
}
