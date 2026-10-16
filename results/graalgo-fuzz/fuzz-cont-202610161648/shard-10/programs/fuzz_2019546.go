package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 bool; F2 int8 }

func main() {
    var v1 uint8 = 54
    v2 := []int{-28, 1, 1}
    v3 := []int{32768, 1, 32767, 32768}
    if len(v3) > 0 {
        v3[0] = 32768
    }
    fmt.Println(v1)
    for fi4 := 0; fi4 < 20; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            v6 := GenS1{F0: 695, F1: v1}
            if ((((!(!((1e10 <= -2.25) && ((((1 <= 255) && false) || (!(-127 == 0))) && true)))) && false) && false) || true) {
                v1 -= 128
                fmt.Println(fi4)
                v1 -= 1
                fmt.Println(fi4, v1)
                v7 := GenS2{F0: 18446744073709551615}
                _ = v7
            }
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    var v8 float64 = 1e10
    fmt.Println(v1, v8)
    if false {
        switch 2147483647 {
        case 109:
            v1 = v1
        case 1:
            var v9 uint = 429
            _ = v9
        case -2147483648:
            fmt.Println(v1, v1)
        }
    }
    var v10 int64 = -658
    if true {
        fmt.Println(v1, v1)
        v11 := GenS3{F0: -2147483648}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 uint32 = 434
            _ = fi12
            _ = v13
        }
        v14 := GenS4{F0: 25}
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int16 = 1
        switch fi15 {
        case -467:
            if ((1 >= 255) || true) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v10 = 440
            } else {
                var v17 int64 = v10
                _ = len(v3)
                switch fi15 {
                case 612:
                    if (!false) {
                        v18 := GenS3{F0: (fi15 >> ((345) & 63))}
                        fmt.Println(v17)
                        v1 = v1
                        _ = v18
                    } else {
                        if len(v3) > 0 {
                            v3[0] = (-(fi15))
                        }
                    }
                case 4294967295:
                    v19 := GenS3{F0: fi15}
                    _ = v19
                case 768:
                    if len(v3) > 0 {
                        v3[0] = fi15
                    }
                }
                _ = v17
            }
        case 127:
            v20 := GenS2{F0: 18446744073709551615}
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    v8 = 1e10
    var v21 uint16 = 2
    switch 1 {
    case 127:
        var v22 uint8 = v1
        _ = v22
    case 1:
        v23 := GenS5{F0: (v10 - 261), F1: true, F2: 1}
        _ = v23
    case 42:
        v1 = 42
    default:
        v21 += (v21 / ((v21) | 1))
    }
    if (147 <= 1) {
        if ((!(!true)) && (v10 != v10)) {
            if (2147483647 <= 4294967295) {
                var v24 int64 = (v10 & v10)
                _ = v24
            } else {
                v21 = (^(v21))
                var v25 uint = 4294967295
                _ = v25
            }
            var v26 uint = 255
            fmt.Println(v10)
            _ = v26
        } else {
            for fi27 := 0; fi27 < 4; fi27++ {
                if len(v3) > 0 {
                    v3[0] = -1
                }
                v28 := GenS1{F0: 18446744073709551615, F1: v1}
                if false {
                    if len(v2) > 0 {
                        v2[0] = int(v1)
                    }
                    if len(v3) > 0 {
                        v3[0] = fi27
                    }
                    v3 = append(v3, -2147483648)
                } else {
                    if (1e10 == -2.25) {
                        v10 += v10
                    } else {
                        v10 -= (^(v10))
                    }
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 uint = 474
                        switch fi29 {
                        case 127:
                            _ = len(v3)
                        }
                        if len(v3) > 0 {
                            v3[0] = (fi29 % ((-1) | 1))
                        }
                        _ = fi29
                        _ = v30
                    }
                }
                _ = fi27
                _ = v28
            }
        }
    } else {
        v31 := GenS4{F0: (v10 - v10)}
        v1 -= (^(v1))
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        fmt.Println(v1, fi32)
        if false {
            fmt.Println(fi32, v21)
        } else {
            if false {
                v10 += -154
                v8 = v8
                fmt.Println(v8, v21, v1)
            } else {
                for fi33 := 0; fi33 < 13; fi33++ {
                    var v34 int8 = 100
                    fmt.Println(v21, v1)
                    _ = fi33
                    _ = v34
                }
                var v35 uint8 = 105
                if true {
                    for fi36 := 0; fi36 < 19; fi36++ {
                        var v37 float32 = 1.5
                        _ = fi36
                        _ = v37
                    }
                    v10 = v10
                    var v38 uint = 65535
                    _ = v38
                }
                _ = v35
            }
            v8 += v8
            var v39 uint64 = uint64(v1)
            _ = v39
        }
        _ = fi32
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v8
}
