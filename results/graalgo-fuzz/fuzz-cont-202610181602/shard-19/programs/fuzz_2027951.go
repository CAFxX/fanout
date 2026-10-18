package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint32; F2 uint8 }
func genFn0() int64 {
    if false {
        v4 := []int{-1, -2147483648, 2147483647, 128}
        if true {
            var v5 float32 = float32((-2.25 - (1e10 + 3.14159)))
            fmt.Println(v5)
            _ = v5
        } else {
            for fi6 := 0; fi6 < 8; fi6++ {
                var v7 int = (fi6 - (fi6 * (fi6 << ((2) & 63))))
                for fi8 := 0; fi8 < 14; fi8++ {
                    v9 := GenS1{F0: 314, F1: "a"}
                    _ = fi8
                    _ = v9
                }
                if len(v4) > 0 {
                    v4[0] = -708
                }
                _ = fi6
                _ = v7
            }
            var v10 int8 = -120
            _ = v10
        }
        _ = v4
    } else {
        switch "hello" {
        case "\n\t":
            var v11 int64 = 9223372036854775807
            _ = v11
        }
    }
    var v12 int16 = 0
    var v13 uint32 = 65535
    for fi14 := 0; fi14 < 18; fi14++ {
        v13 = v13
        _ = fi14
    }
    v12 -= 808
    _ = v12
    _ = v13
    return -966
}

func main() {
    var v1 bool = (true && (("foo bar" <= "a") || false))
    var v2 uint8 = 177
    v3 := []int{2147483647, -1}
    var v15 uint8 = (v2 % (((v2 & (v2 << ((511) & 63)))) | 1))
    v3 = append(v3, 32768)
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 float32 = float32(3.14159)
        if len(v3) > 0 {
            v3[0] = 863
        }
        _ = fi16
        _ = v17
    }
    switch -2147483648 {
    case -876:
        v3 = append(v3, 9223372036854775807)
    case -1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 0:
        var v18 uint64 = 222
        _ = v18
    default:
        if v1 {
            _ = len(v3)
            if ((v1 && ((v15 / (((v15 + v15)) | 1)) != (^(v2)))) && v1) {
                var v19 int8 = 127
                _ = v19
            }
            v1 = (42 > 42)
        }
    }
    if (true && (genFn0() <= -937)) {
        v20 := GenS2{F0: (339 == 130)}
        v21 := GenS3{F0: -1, F1: 2147483648, F2: v2}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 18; fi22++ {
            fmt.Println(v2)
            _ = len(v3)
            _ = fi22
        }
        if (!v1) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = len(v3)
            v1 = ((14 == 32) && (true && true))
        } else {
            _ = len(v3)
        }
    }
    v15 -= (v2 + v15)
    if (false && (42 <= -544)) {
        v15 = (v2 - 127)
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS3{F0: 2147483647, F1: 65535, F2: v2}
        fmt.Println(fi23)
        _ = fi23
        _ = v24
    }
    v1 = true
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 float32 = ((1.5 + (3.14159 * (-2.25 + 1e10))) / 1e10)
        for fi27 := 0; fi27 < 19; fi27++ {
            if (2147483648 >= 1) {
                _ = len(v3)
            } else {
                var v28 uint = 4294967295
                v29 := GenS3{F0: 386, F1: 421, F2: (^(v15))}
                _ = v28
                _ = v29
            }
            if v1 {
                v2 += 32
            }
            var v30 int32 = 1
            v1 = (42 != 571)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi27
            _ = v30
        }
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        _ = fi25
        _ = v26
    }
    v15 -= v2
    switch 128 {
    case 1:
        for fi31 := 0; fi31 < 6; fi31++ {
            fmt.Println(v2, v15)
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 float64 = (1e10 - float64((-2.25 / 0.1)))
                if true {
                    switch "" {
                    case "\n\t":
                        var v34 uint64 = 0
                        _ = v34
                    case "foo bar":
                        v1 = (false || (818 < 165))
                    }
                }
                v1 = v1
                v35 := GenS3{F0: -1, F1: 2147483647, F2: 1}
                v15 -= (^(v2))
                _ = fi32
                _ = v33
                _ = v35
            }
            switch fi31 {
            case 9223372036854775807:
                fmt.Println(v2, v2)
            default:
                v36 := GenS2{F0: false}
                _ = v36
            }
            _ = fi31
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if v1 {
        if len(v3) > 0 {
            v3[0] = 263
        }
    } else {
        var v37 int16 = -700
        _ = v37
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    v15 += 2
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
}
