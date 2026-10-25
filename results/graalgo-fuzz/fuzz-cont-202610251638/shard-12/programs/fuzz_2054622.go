package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int8; F1 int; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint8 }
func genFn0(p0 int, p1 uint8) int64 {
    switch p0 {
    case -9223372036854775808:
        var v7 int16 = 935
        _ = v7
    }
    var v8 uint = 255
    var v9 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{0, -9223372036854775808, -9223372036854775808}
    var v3 int8 = 1
    v4 := []int{}
    var v5 string = "foo bar"
    var v6 uint64 = 302
    fmt.Println(v5)
    var v10 uint16 = 199
    v11 := GenS1{F0: v3, F1: 2147483647}
    var v12 uint = 42
    fmt.Println(v3)
    if (v5 > "") {
        var v13 uint = (^(v12))
        _ = v13
    } else {
        v2 = append(v2, 42)
        var v14 uint32 = 2147483647
        var v15 bool = true
        _ = v14
        _ = v15
    }
    if (1000 >= int16(v3)) {
        v16 := GenS2{F0: -1, F1: -777, F2: 4294967295}
        var v17 uint8 = 1
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v5 = ((v5 + "a") + "")
        _ = fi18
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v6, v12)
    v19 := GenS3{F0: true, F1: 2, F2: 9223372036854775807}
    fmt.Println(v5, v12)
    if true {
        if false {
            var v20 int16 = -835
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := GenS3{F0: false, F1: uint8(v6), F2: -9223372036854775808}
                v2 = append(v2, (fi21 & (fi21 + 1)))
                _ = fi21
                _ = v22
            }
            switch -9223372036854775808 {
            case 128:
                v23 := GenS3{F0: (128 < 2), F1: 128, F2: genFn0(641, 255)}
                _ = v23
            default:
                if (2 < 225) {
                    if (-981 < 926) {
                        var v24 int8 = v3
                        _ = v24
                    }
                } else {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 float64 = 1e10
                        fmt.Println(v20)
                        _ = fi25
                        _ = v26
                    }
                }
            }
            _ = v20
        } else {
            v6 += 42
            var v27 int16 = 1
            v10 = 0
            _ = v27
        }
        v28 := GenS1{F0: (-(v3)), F1: 845}
        _ = v28
    }
    switch v5 {
    case "a":
        v29 := GenS3{F0: false, F1: 0, F2: -1}
        _ = v29
    case "hello":
        if false {
            v30 := GenS4{F0: 1}
            for fi31 := 0; fi31 < 13; fi31++ {
                if ((0.1 / 1e10) > ((0.0 + (float64(((1e10 + float32(((-2.25 - 1.5) / (float64(float32((3.14159 + 1e10))) + ((1e10 * ((((-2.25 * -2.25) / -2.25) + 3.14159) / 1e10)) / -2.25))))) - 3.14159)) * (float64((-2.25 / 3.14159)) - (float64(((-2.25 + -2.25) * -2.25)) + -2.25)))) / float64((float32((1.5 - 1.5)) + float32(-2.25))))) {
                    if true {
                        if (v5 >= v5) {
                            v10 = v10
                            switch v5 {
                            case "foo bar":
                                switch (fi31 + fi31) {
                                case 4294967295:
                                    v3 += (v3 ^ (v3 + v3))
                                case -9223372036854775808:
                                    for fi32 := 0; fi32 < 20; fi32++ {
                                        switch 4294967295 {
                                        case 0:
                                            fmt.Println(fi32)
                                        }
                                        _ = fi32
                                    }
                                case 9223372036854775807:
                                    for fi33 := 0; fi33 < 14; fi33++ {
                                        v34 := GenS5{F0: (v12 | v12), F1: 170, F2: 1}
                                        var v35 uint16 = (v10 * (v10 - 658))
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                }
                            default:
                                v36 := GenS2{F0: v3, F1: (fi31 - (fi31 | -912)), F2: 2147483648}
                                _ = v36
                            }
                            var v37 float64 = -2.25
                            _ = v37
                        } else {
                            v12 = (v12 << ((v12) & 63))
                            var v38 uint32 = 1
                            v3 += v3
                            _ = v38
                        }
                        delete(v1, 2)
                    }
                    var v39 float32 = 3.14159
                    v10 -= 2
                    _ = v39
                } else {
                    fmt.Println(v5, v5)
                    v6 -= (^(v6))
                }
                _ = fi31
            }
            var v40 int16 = 42
            _ = v30
            _ = v40
        } else {
            for fi41 := 0; fi41 < 15; fi41++ {
                v3 -= v3
                _ = fi41
            }
            v42 := GenS3{F0: true, F1: 128, F2: 746}
            _ = v42
        }
    }
    var v43 uint = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
