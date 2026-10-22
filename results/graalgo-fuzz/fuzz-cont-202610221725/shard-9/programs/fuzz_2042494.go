package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 uint; F2 uint }
func genFn0(p0 int, p1 int32, p2 bool) int16 {
    v7 := []int{}
    p1 -= (p1 + p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -555
}

func main() {
    var v1 int = 1
    v2 := make(map[uint32]string)
    var v3 int32 = 2147483647
    var v4 int32 = (v3 << ((255) & 63))
    var v5 uint8 = 42
    v6 := make(map[string]int)
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: 56, F1: true}
        switch "" {
        case "a":
            delete(v2, 2)
        default:
            v5 += v5
        }
        _ = fi8
        _ = v9
    }
    if true {
        var v10 string = ("hello" + ("a" + "foo bar"))
        if (!((v3 * -2147483648) == (v4 & -2147483648))) {
            var v11 float64 = (0.0 + ((0.0 + (-2.25 * -2.25)) - 3.14159))
            if false {
                var v12 bool = (!(-21 < -128))
                _ = v12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                fmt.Println(v4, v1, v4)
                var v14 uint8 = (v5 | 0)
                _ = fi13
                _ = v14
            }
            _ = v11
        } else {
            var v15 uint16 = 32767
            var v16 uint32 = 2147483648
            v17 := GenS1{F0: 42, F1: (float64((0.1 * 3.14159)) >= (0.1 / 3.14159))}
            _ = v15
            _ = v16
            _ = v17
        }
        var v18 uint64 = 18446744073709551615
        _ = v10
        _ = v18
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 float64 = 1e10
            for fi21 := 0; fi21 < 9; fi21++ {
                fmt.Println(v5, v1, v3)
                _ = len(v6)
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            if true {
                fmt.Println(fi22, v5, v3)
            }
            _ = fi22
        }
    }
    fmt.Println(v3, v5, v5)
    v23 := GenS2{F0: (-(v4))}
    var v24 uint8 = v5
    fmt.Println(v5, v24)
    var v25 int32 = (v3 & (v3 + 42))
    fmt.Println(v24, v3, v5)
    v26 := v2[uint32(v5)]
    v27 := v6[v26]
    var v28 float64 = -2.25
    v29 := []int{}
    v5 = (v24 / (((v24 - (^(v24)))) | 1))
    var v30 int16 = -1
    fmt.Println(v26, v25)
    if false {
        delete(v2, 4294967295)
        fmt.Println(v3, v5)
        if (2 >= 1) {
            fmt.Println(v26, v3, v4)
        } else {
            fmt.Println(v4, v5)
            switch v27 {
            case 4294967295:
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 int16 = v30
                    if (!true) {
                        var v33 uint32 = 1
                        if (float64((float32(v28) + 0.1)) == 1e10) {
                            switch -9223372036854775808 {
                            case -47:
                                if (183 <= 876) {
                                    for fi34 := 0; fi34 < 6; fi34++ {
                                        switch ((v26 + v26) + v26) {
                                        case "a":
                                            v35 := GenS3{F0: (^(v5)), F1: -9223372036854775808}
                                            _ = v35
                                        default:
                                            fmt.Println(v28, v25, v1)
                                        }
                                        if len(v29) > 0 {
                                            v29[0] = v27
                                        }
                                        v24 += 231
                                        v36 := GenS4{F0: 4294967295}
                                        v5 += (v5 & 127)
                                        _ = fi34
                                        _ = v36
                                    }
                                    fmt.Println(v27, v30)
                                    v37 := GenS5{F0: (v24 ^ (v24 * (v5 * v24)))}
                                    _ = v37
                                }
                            case 0:
                                var v38 bool = true
                                _ = v38
                            default:
                                var v39 int32 = v3
                                _ = v39
                            }
                        } else {
                            v26 = v26
                            if (("a" + v26) != "a") {
                                var v40 int32 = (v4 * -2147483648)
                                if true {
                                    v41 := GenS6{F0: -632, F1: 0, F2: 65535}
                                    var v42 float64 = 3.14159
                                    var v43 uint64 = 4294967295
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                                v33 = 51
                                _ = v40
                            }
                        }
                        fmt.Println(v1, v3)
                        _ = v33
                    } else {
                        var v44 bool = (!true)
                        _ = v44
                    }
                    _ = fi31
                    _ = v32
                }
            case 32767:
                fmt.Println(v27, v25, v3)
            default:
                fmt.Println(v1, v3)
            }
            v45 := v6[""]
            _ = v45
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
