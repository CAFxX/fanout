package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }
type GenS3 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint16 }
func genFn0() bool {
    var v4 uint8 = 127
    if false {
        v5 := []int{-9223372036854775808, 68, 128, 2147483647}
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 float64 = (((((1.5 / (float64(-2.25) - 0.1)) * -2.25) + (float64(3.14159) - (3.14159 - 1.5))) / (float64(float32(-2.25)) * 3.14159)) * 1e10)
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        var v8 int = -57
        _ = v8
    }
    v9 := []int{4294967295, 32767, -310}
    v10 := GenS1{F0: 100, F1: 650}
    _ = v10
    _ = v4
    _ = v9
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 uint = 42
    v2 = false
    if (2147483647 > 0) {
        v11 := GenS2{F0: true, F1: 32767, F2: ((0 >= -993) && (0 > 0))}
        v2 = (822 <= uint32(v3))
        fmt.Println(v3)
        _ = v11
    } else {
        v12 := GenS1{F0: 118, F1: 0}
        v2 = false
        _ = v12
    }
    v3 = (v3 | v3)
    if (9223372036854775807 < 2147483647) {
        var v13 int64 = 436
        _ = v13
    } else {
        switch 1 {
        case 0:
            if (32767 != 2) {
                v2 = false
                v14 := []int{488, 0}
                v15 := GenS1{F0: 127, F1: 2147483647}
                _ = v14
                _ = v15
            }
        }
        var v16 uint64 = 4294967295
        if true {
            v2 = false
        } else {
            v3 += (v3 % (((v3 | (v3 + 4294967295))) | 1))
            v2 = (!v2)
            switch ((((("" + "hello") + "hello") + "foo bar") + ("hello" + "\n\t")) + "") {
            case "a":
                if (((-2.25 < float32(float64((1.5 * (float32(float64(((float32((-2.25 - float64(1.5))) / 1e10) - 0.0))) / (((0.1 + (1.5 * 1.5)) + ((0.1 - 1e10) * 3.14159)) * float32((float64(0.1) + 1.5)))))))) && false) || (583 < 0)) {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v18 := GenS3{F0: 295, F1: v2, F2: 769}
                        if (v3 <= v3) {
                            v16 = v16
                            for fi19 := 0; fi19 < 14; fi19++ {
                                v16 = v16
                                v16 -= 0
                                _ = fi19
                            }
                            v20 := GenS3{F0: 867, F1: false, F2: 0}
                            _ = v20
                        } else {
                            v3 = (v3 - uint(v16))
                            v2 = v2
                            v21 := GenS1{F0: 1, F1: -497}
                            _ = v21
                        }
                        v22 := GenS4{F0: 953}
                        _ = fi17
                        _ = v18
                        _ = v22
                    }
                    v1[0] = 9223372036854775807
                    delete(v1, uint32(v3))
                } else {
                    v16 += (v16 - (v16 >> (((v3 ^ (^(v3)))) & 63)))
                    fmt.Println(v2, v3, v16)
                }
            case "":
                fmt.Println(v2, v2, v2)
            case "foo bar":
                v16 -= (v16 << ((2) & 63))
            default:
                v16 = v16
            }
        }
        _ = v16
    }
    var v23 int32 = 0
    v24 := []int{128, 127, int(v23), -2147483648}
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    v2 = v2
    var v25 uint = v3
    v23 = v23
    for fi26 := 0; fi26 < 2; fi26++ {
        switch fi26 {
        case -317:
            v2 = ((1 >= 65535) || v2)
        default:
            var v27 float32 = 0.0
            _ = v27
        }
        _ = fi26
    }
    if len(v24) > 0 {
        v24[0] = 128
    }
    var v28 uint64 = 399
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
}
