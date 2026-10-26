package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint32 }
func genFn0(p0 int64, p1 uint16, p2 uint) uint8 {
    var v7 uint64 = 4294967295
    v8 := GenS1{F0: 127}
    v9 := []int{668, -1}
    var v10 int64 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 int8) uint32 {
    var v11 uint16 = 35
    p0 += (^(p0))
    v12 := []int{128, 4294967295}
    var v13 uint32 = 841
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    var v1 int8 = 1
    var v2 int = 809
    var v3 uint64 = 243
    var v4 int32 = -2147483648
    var v5 int16 = -32768
    var v6 int16 = v5
    var v14 float32 = 3.14159
    v15 := []int{(v2 / (((v2 >> ((1) & 63))) | 1)), (v2 << ((0) & 63)), v2, v2}
    if (0 != (v1 >> ((4294967295) & 63))) {
        v15 = append(v15, 1)
        if false {
            if true {
                var v16 int64 = 9223372036854775807
                for fi17 := 0; fi17 < 17; fi17++ {
                    if false {
                        v15 = append(v15, fi17)
                        switch ("foo bar" + "") {
                        case "hello":
                            var v18 float64 = 1e10
                            _ = v18
                        case "a":
                            v16 -= 0
                        }
                        for fi19 := 0; fi19 < 7; fi19++ {
                            v20 := GenS2{F0: 100, F1: "", F2: 2}
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        v6 = 42
                        v5 = -463
                    }
                    v2 = fi17
                    fmt.Println(v3)
                    _ = fi17
                }
                _ = v16
            } else {
                var v21 uint16 = 2
                v15 = append(v15, int(v1))
                fmt.Println(v4)
                _ = v21
            }
            if (false && (65535 != 1)) {
                var v22 uint16 = 255
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 int16 = (v6 + v5)
                    if (true || false) {
                        var v25 float64 = -2.25
                        _ = v25
                    }
                    if len(v15) > 0 {
                        v15[0] = 42
                    }
                    _ = fi23
                    _ = v24
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    var v27 uint = 560
                    _ = len(v15)
                    for fi28 := 0; fi28 < 3; fi28++ {
                        v29 := GenS3{F0: 4294967295}
                        v30 := GenS4{F0: v1}
                        if true {
                            switch (^(v2)) {
                            case 923:
                                v31 := GenS4{F0: (v1 << ((uint(v3)) & 63))}
                                _ = v31
                            default:
                                v14 -= float32(1e10)
                            }
                            _ = len(v15)
                        }
                        var v32 int32 = (-(v4))
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v32
                    }
                    _ = fi26
                    _ = v27
                }
                _ = v22
            }
        }
        v33 := GenS4{F0: v1}
        _ = v33
    }
    switch (v2 >> ((18446744073709551615) & 63)) {
    case 407:
        var v34 uint16 = 226
        _ = v34
    case -1:
        v2 = (v2 >> ((1) & 63))
    default:
        v35 := GenS5{F0: 0}
        _ = v35
    }
    var v36 float32 = 3.14159
    fmt.Println(v5, v2, v1)
    var v37 int16 = v6
    v15 = append(v15, 32767)
    switch (v2 << ((255) & 63)) {
    case -9223372036854775808:
        var v38 uint32 = 0
        _ = v38
    case 4294967295:
        v36 += -2.25
    default:
        for fi39 := 0; fi39 < 16; fi39++ {
            v40 := GenS6{F0: 889, F1: 42, F2: 494}
            var v41 uint = 65535
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    v3 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
