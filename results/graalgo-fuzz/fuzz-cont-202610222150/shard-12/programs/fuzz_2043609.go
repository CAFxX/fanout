package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int64; F1 int64; F2 uint8 }

func main() {
    var v1 string = "hello"
    var v2 int32 = 0
    v3 := []int{-164, 128, 2147483647, 981}
    v4 := make(map[string]bool)
    if ((!(42 <= -102)) || (((0.0 / float32((-2.25 / 1.0))) / 0.1) <= 1e10)) {
        if len(v3) > 0 {
            v3[0] = 9223372036854775807
        }
    }
    fmt.Println(v1, v2)
    var v5 uint32 = 2147483648
    if false {
        v2 = (v2 + (v2 << ((18446744073709551615) & 63)))
        if (42 < -9223372036854775808) {
            if (false || false) {
                for fi6 := 0; fi6 < 17; fi6++ {
                    var v7 int = (fi6 + fi6)
                    v8 := v4[v1]
                    if (!((!v8) || false)) {
                        var v9 int8 = 100
                        v8 = (1 <= 127)
                        var v10 uint8 = 205
                        _ = v10
                        _ = v9
                    }
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                v2 = v2
            }
            v11 := v4["\n\t"]
            fmt.Println(v2)
            _ = v11
        }
    }
    v12 := GenS1{F0: 42, F1: (v2 + (v2 << ((255) & 63))), F2: (("" + "foo bar") + "")}
    for fi13 := 0; fi13 < 3; fi13++ {
        if (((1.5 - ((1e10 - 0.1) * float32(1.5))) + float32(-2.25)) < (float32((0.0 + 3.14159)) + 0.0)) {
            if (true || (!true)) {
                if (283 <= 42) {
                    if false {
                        for fi14 := 0; fi14 < 16; fi14++ {
                            v15 := GenS2{F0: int64(v2), F1: 4294967295}
                            switch fi13 {
                            case 998:
                                switch -9223372036854775808 {
                                case 0:
                                    v2 = -1
                                case -377:
                                    v1 = v1
                                }
                            case 4294967295:
                                fmt.Println(v2)
                            case 42:
                                fmt.Println(v2, fi14)
                            default:
                                switch fi14 {
                                case 4294967295:
                                    var v16 string = "a"
                                    if len(v3) > 0 {
                                        v3[0] = (fi14 * fi13)
                                    }
                                    _ = v16
                                case 9223372036854775807:
                                    fmt.Println(v2, v2, v1)
                                case -9223372036854775808:
                                    v17 := GenS1{F0: 193, F1: 39, F2: "\n\t"}
                                    _ = v17
                                }
                            }
                            _ = fi14
                            _ = v15
                        }
                    } else {
                        fmt.Println(fi13)
                        if len(v3) > 0 {
                            v3[0] = 0
                        }
                        v18 := GenS1{F0: 2, F1: v2, F2: v1}
                        _ = v18
                    }
                    var v19 uint16 = 255
                    var v20 int16 = 0
                    _ = v19
                    _ = v20
                }
                v21 := GenS2{F0: 159, F1: (v5 - 42)}
                _ = v21
            } else {
                for fi22 := 0; fi22 < 13; fi22++ {
                    v23 := GenS2{F0: 155, F1: v5}
                    _ = fi22
                    _ = v23
                }
                v5 = (v5 >> ((111) & 63))
            }
            v2 = (v2 % ((0) | 1))
            v24 := GenS2{F0: 1, F1: 656}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                v5 = v5
                if (float64(float32(((1e10 - (1e10 * (float64(((0.1 + 0.1) / 0.1)) - 1e10))) / ((((float64(1e10) * -2.25) - (3.14159 + float64(-2.25))) / 0.1) - float64(-2.25))))) > -2.25) {
                    v26 := GenS3{F0: 764, F1: 2}
                    _ = v26
                }
                _ = fi25
            }
            for fi27 := 0; fi27 < 15; fi27++ {
                if false {
                    var v28 int = (fi13 & fi27)
                    if false {
                        v28 = 32767
                        switch (fi13 - fi27) {
                        case 42:
                            var v29 uint16 = 872
                            _ = v29
                        case -914:
                            v30 := GenS4{F0: 1}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v30
                        }
                        delete(v4, v1)
                    } else {
                        v5 += 2147483647
                    }
                    var v31 int8 = 39
                    _ = v28
                    _ = v31
                }
                v5 -= (v5 / ((1) | 1))
                v1 = v1
                v1 = ("\n\t" + v1)
                var v32 int8 = 1
                _ = fi27
                _ = v32
            }
            v33 := GenS5{F0: 418, F1: (v1 + "hello")}
            _ = v33
        }
        v2 = v2
        if true {
            v34 := GenS6{F0: 9223372036854775807, F1: -673, F2: 2}
            var v35 int16 = 883
            _ = v34
            _ = v35
        }
        _ = fi13
    }
    var v36 int8 = 100
    v37 := GenS4{F0: 9223372036854775807}
    switch 32768 {
    case 9223372036854775807:
        v4["hello"] = true
    case 32767:
        v38 := v4[(v1 + "hello")]
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
