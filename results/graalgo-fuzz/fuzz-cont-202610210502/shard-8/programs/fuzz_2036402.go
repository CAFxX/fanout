package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int }
func genFn0() string {
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 int16 = 1000
        _ = fi4
        _ = v5
    }
    v6 := []int{42, 4294967295}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = v6
    return "a"
}

func main() {
    var v1 int16 = 32767
    var v2 int = 9223372036854775807
    var v3 uint16 = 171
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{(fi7 | (fi7 & (v2 % ((v2) | 1)))), fi7, (v2 & (v2 - 1))}
        switch (v2 | 4294967295) {
        case 665:
            fmt.Println(fi7, v1)
        default:
            v2 = v2
        }
        _ = fi7
        _ = v8
    }
    var v9 int8 = 55
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := GenS1{F0: (v3 % (((v3 | (v3 + (v3 + 450)))) | 1)), F1: 32767}
        var v12 int = fi10
        v9 = -1
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 -= 0
    var v13 float32 = (-2.25 + 3.14159)
    v1 += v1
    switch 275 {
    case -749:
        var v14 int8 = (v9 & v9)
        _ = v14
    case 127:
        var v15 uint64 = 345
        _ = v15
    case -9223372036854775808:
        for fi16 := 0; fi16 < 17; fi16++ {
            fmt.Println(v9, v9)
            for fi17 := 0; fi17 < 10; fi17++ {
                if true {
                    if true {
                        switch (fi16 ^ (fi17 ^ (^(fi16)))) {
                        case -110:
                            v3 = v3
                            fmt.Println(fi17, v9, v9)
                            var v18 uint64 = 0
                            _ = v18
                        case 32767:
                            var v19 int = fi16
                            _ = v19
                        case -9223372036854775808:
                            v20 := GenS1{F0: v3, F1: (v3 << ((0) & 63))}
                            _ = v20
                        default:
                            fmt.Println(v2, fi17, fi16)
                        }
                        v21 := GenS1{F0: (v3 ^ 42), F1: (v3 ^ v3)}
                        _ = v21
                    } else {
                        v22 := []int{fi17, fi16}
                        var v23 float32 = float32((1.5 / 3.14159))
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS1{F0: v3, F1: v3}
                    v25 := []int{fi17, (^(v2)), fi17}
                    _ = v24
                    _ = v25
                }
                v26 := GenS2{F0: 0, F1: 0}
                switch "" {
                case "a":
                    v1 += v1
                case "":
                    var v27 int16 = (v1 << ((2) & 63))
                    _ = v27
                }
                var v28 uint32 = 2
                fmt.Println(v2, fi17, v28)
                _ = fi17
                _ = v26
                _ = v28
            }
            v1 = (-(v1))
            _ = fi16
        }
    default:
        v29 := []int{42, v2}
        _ = v29
    }
    switch v2 {
    case -448:
        fmt.Println(v13)
    case 4294967295:
        if (0 < -1) {
            for fi30 := 0; fi30 < 19; fi30++ {
                var v31 int8 = (-(v9))
                var v32 int16 = (v1 >> ((1) & 63))
                if false {
                    var v33 int16 = v1
                    _ = v33
                } else {
                    v34 := GenS2{F0: -441, F1: 255}
                    _ = v34
                }
                _ = fi30
                _ = v31
                _ = v32
            }
        } else {
            v13 = (v13 / 1.5)
        }
    }
    v9 += (v9 * 95)
    var v35 uint64 = 2
    if ((uint32(v35) <= 717) && (42 < 1)) {
        var v36 float32 = ((((v13 * float32(-2.25)) + (v13 + v13)) + (((v13 + 3.14159) / v13) / float32(3.14159))) / (0.0 - 3.14159))
        v37 := GenS3{F0: (v2 | v2)}
        _ = v36
        _ = v37
    }
    v35 += v35
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v9
}
