package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint; F1 int64; F2 bool }
func genFn0(p0 uint32, p1 int32) int32 {
    p1 = p1
    v6 := GenS1{F0: -974, F1: 42, F2: 0}
    v7 := GenS1{F0: 917, F1: 18446744073709551615, F2: 9223372036854775807}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p1 % (((p1 ^ p1)) | 1))
}

func main() {
    var v1 int = 122
    var v2 int32 = 2147483647
    var v3 uint = 676
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 658
    var v8 uint8 = 127
    var v9 int8 = 1
    for fi10 := 0; fi10 < 20; fi10++ {
        switch (-(fi10)) {
        case 2147483647:
            delete(v4, 1)
        default:
            v2 -= v2
        }
        switch "foo bar" {
        case "hello":
            var v11 uint16 = 288
            _ = v11
        case "":
            fmt.Println(v8)
        case "a":
            v12 := GenS1{F0: 1, F1: 938, F2: 1}
            _ = v12
        }
        _ = fi10
    }
    var v13 uint32 = v5
    var v14 int8 = v9
    v15 := GenS1{F0: 1, F1: 350, F2: 83}
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int8 = -128
        v8 += (^(v8))
        _ = fi16
        _ = v17
    }
    v3 = v3
    v14 = (v14 & (v9 / ((v9) | 1)))
    if false {
        for fi18 := 0; fi18 < 5; fi18++ {
            if false {
                v19 := GenS1{F0: 42, F1: 0, F2: 42}
                var v20 float64 = float64(float32(3.14159))
                _ = v19
                _ = v20
            } else {
                v5 += v13
                for fi21 := 0; fi21 < 6; fi21++ {
                    v1 = -1
                    _ = fi21
                }
            }
            _ = fi18
        }
    } else {
        v22 := GenS1{F0: 304, F1: 18446744073709551615, F2: 1}
        _ = v22
    }
    if (!true) {
        switch (v1 >> ((v3) & 63)) {
        case 32768:
            if (((-2.25 / ((((((0.0 * (-2.25 - (3.14159 * float32(3.14159)))) * (0.0 / 1e10)) + 1e10) - 3.14159) * 1e10) * 1.5)) > (((-2.25 - ((0.1 - (1.5 / float32(-2.25))) - (1.5 + 1.5))) / ((0.1 / 0.1) - 0.1)) + 1e10)) || false) {
                fmt.Println(v2, v8)
                v23 := GenS2{F0: v8, F1: uint(v13)}
                v24 := []int{(v1 - (^(v1)))}
                _ = v23
                _ = v24
            }
        default:
            if (!false) {
                v13 = v5
                v4[uint32(v8)] = 2147483647
            }
        }
        switch -2147483648 {
        case 2147483647:
            switch v1 {
            case 35:
                var v25 uint32 = 65535
                _ = v25
            default:
                var v26 int8 = v14
                _ = v26
            }
        case -1:
            if false {
                for fi27 := 0; fi27 < 16; fi27++ {
                    if true {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            switch v1 {
                            case 770:
                                fmt.Println(fi27)
                            }
                            _ = fi28
                        }
                        v29 := []int{}
                        _ = v29
                    }
                    v8 -= (v8 >> (((v3 / (((v3 & (v3 / ((v3) | 1)))) | 1))) & 63))
                    _ = fi27
                }
                v5 -= 2147483648
                var v30 uint16 = 553
                _ = v30
            } else {
                if false {
                    if (29 <= 478) {
                        v31 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615, F2: 0}
                        v1 += v1
                        _ = v31
                    }
                    fmt.Println(v9)
                    v32 := GenS1{F0: -1, F1: 0, F2: 611}
                    _ = v32
                }
                for fi33 := 0; fi33 < 19; fi33++ {
                    if false {
                        var v34 float64 = (3.14159 * (1e10 + (((float64(-2.25) + 3.14159) + 3.14159) * ((1.5 / 1e10) + 0.1))))
                        v35 := GenS3{F0: (^(v3)), F1: 990, F2: false}
                        for fi36 := 0; fi36 < 2; fi36++ {
                            fmt.Println(v5, v2)
                            _ = fi36
                        }
                        _ = v34
                        _ = v35
                    } else {
                        v8 = 159
                        v37 := GenS2{F0: 225, F1: 189}
                        if true {
                            var v38 int16 = 0
                            var v39 int8 = v9
                            _ = v38
                            _ = v39
                        }
                        _ = v37
                    }
                    var v40 int8 = 1
                    for fi41 := 0; fi41 < 11; fi41++ {
                        v1 = v1
                        var v42 float64 = (float64(3.14159) - 3.14159)
                        _ = fi41
                        _ = v42
                    }
                    _ = fi33
                    _ = v40
                }
                var v43 uint16 = 112
                _ = v43
            }
        }
        if false {
            var v44 uint8 = (v8 - v8)
            _ = v44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
