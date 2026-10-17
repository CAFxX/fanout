package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint; F1 int8; F2 int8 }
type GenS5 struct { F0 int16; F1 string; F2 int16 }
func genFn0() int16 {
    var v7 uint16 = 1
    if (true || false) {
        v8 := []int{128, 42, 1, 1}
        _ = v8
    } else {
        var v9 string = ("" + "")
        var v10 int8 = 0
        var v11 float64 = ((3.14159 + ((3.14159 + 3.14159) - (-2.25 * (-2.25 + 3.14159)))) / (3.14159 / (-2.25 + 0.0)))
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := []int{-318, 0}
    for fi13 := 0; fi13 < 11; fi13++ {
        v7 = (v7 - 255)
        v7 = v7
        _ = fi13
    }
    _ = v12
    _ = v7
    return 446
}

func main() {
    var v1 string = "\n\t"
    var v2 int64 = 1
    v3 := []int{int(v2), 1, 0}
    var v4 uint = 1
    var v5 uint = (v4 - (v4 - (v4 >> (((v4 | v4)) & 63))))
    var v6 uint8 = 42
    var v14 uint64 = 4294967295
    if (!((!(!(-639 <= 128))) || (2 >= 65535))) {
        var v15 uint32 = 2
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := GenS1{F0: v2, F1: v1}
            v14 -= 690
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: v2, F1: v1}
        _ = v15
        _ = v18
    }
    var v19 uint = v5
    _ = len(v3)
    fmt.Println(v2)
    var v20 int8 = 49
    if false {
        for fi21 := 0; fi21 < 6; fi21++ {
            for fi22 := 0; fi22 < 16; fi22++ {
                v14 += uint64(v4)
                _ = fi22
            }
            v14 = v14
            _ = fi21
        }
        v23 := GenS2{F0: (300 >= 65535)}
        _ = v23
    }
    var v24 int32 = 42
    switch v1 {
    case "":
        v6 = 128
    case "foo bar":
        var v25 uint8 = v6
        _ = v25
    }
    if false {
        v14 += (^(v14))
        var v26 int16 = int16(v2)
        _ = v26
    } else {
        switch 127 {
        case 127:
            switch int(v14) {
            case 42:
                var v27 int32 = (v24 & v24)
                _ = v27
            case 127:
                var v28 int64 = -643
                _ = v28
            case 9223372036854775807:
                if len(v3) > 0 {
                    v3[0] = -837
                }
            }
        case -9223372036854775808:
            if ((((v2 - int64(v24)) <= (v2 - (v2 + (-(v2))))) && (-2.25 > (1e10 + 3.14159))) || (!((^(v4)) > v19))) {
                for fi29 := 0; fi29 < 5; fi29++ {
                    if (((v20 ^ v20) < -13) && (!false)) {
                        fmt.Println(v4, v2, v1)
                        switch 1 {
                        case 42:
                            v3 = append(v3, 562)
                        case 32767:
                            var v30 bool = ((v20 / (((v20 << (((^(v19))) & 63))) | 1)) < -1)
                            _ = v30
                        default:
                            v4 = (v4 + (v5 / ((v5) | 1)))
                        }
                    }
                    _ = fi29
                }
            } else {
                switch -787 {
                case 1:
                    v31 := GenS1{F0: 1, F1: v1}
                    _ = v31
                case 135:
                    _ = len(v3)
                case 42:
                    var v32 uint8 = v6
                    _ = v32
                }
            }
        case 556:
            v33 := GenS3{F0: 2, F1: 811, F2: 65535}
            _ = v33
        }
        v34 := GenS2{F0: (v6 == (v6 << ((v4) & 63)))}
        _ = v34
    }
    v35 := GenS1{F0: (v2 * v2), F1: "\n\t"}
    if len(v3) > 0 {
        v3[0] = 966
    }
    v36 := GenS4{F0: v5, F1: (v20 - (v20 + 1)), F2: 42}
    var v37 string = ("" + (v1 + "hello"))
    switch 2147483647 {
    case -2147483648:
        if false {
            switch v37 {
            case "a":
                var v38 int64 = v2
                _ = v38
            case "\n\t":
                fmt.Println(v24, v19)
            case "":
                if (32767 == 42) {
                    fmt.Println(v19)
                }
            default:
                v14 -= v14
            }
            v39 := GenS5{F0: -1, F1: "hello", F2: -32768}
            switch -9223372036854775808 {
            case 32767:
                v37 = ((v37 + "") + (v1 + v1))
            default:
                v24 = (v24 * 1)
            }
            _ = v39
        } else {
            v1 = v1
            if false {
                for fi40 := 0; fi40 < 1; fi40++ {
                    v3 = append(v3, 9223372036854775807)
                    _ = fi40
                }
            }
        }
    }
    if (32767 < -854) {
        for fi41 := 0; fi41 < 16; fi41++ {
            v14 += v14
            v24 -= v24
            _ = fi41
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 297
        }
        fmt.Println(v37)
    }
    v20 += 15
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
