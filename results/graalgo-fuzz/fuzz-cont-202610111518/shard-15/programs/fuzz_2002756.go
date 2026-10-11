package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
func genFn0(p0 int8) bool {
    var v5 int8 = (p0 / ((p0) | 1))
    var v6 uint64 = 18446744073709551615
    switch 9223372036854775807 {
    case 32768:
        var v7 uint32 = 42
        _ = v7
    case 42:
        if (42 != 42) {
            var v8 int = 32768
            v9 := GenS1{F0: 55, F1: 18446744073709551615}
            fmt.Println(v6)
            _ = v8
            _ = v9
        } else {
            v6 = v6
            fmt.Println(v6)
            if true {
                switch 128 {
                case -522:
                    if ((0 != -811) && true) {
                        p0 = 70
                    } else {
                        for fi10 := 0; fi10 < 16; fi10++ {
                            var v11 float32 = 0.1
                            _ = fi10
                            _ = v11
                        }
                        fmt.Println(v6, v5)
                    }
                }
            } else {
                if false {
                    var v12 int64 = 0
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 5; fi13++ {
                        for fi14 := 0; fi14 < 4; fi14++ {
                            if false {
                                v5 = (v5 % ((-128) | 1))
                                v15 := []int{(fi14 ^ int(v6)), fi13, fi13, fi14}
                                _ = v15
                            } else {
                                var v16 int8 = p0
                                _ = v16
                            }
                            if false {
                                v17 := GenS2{F0: 128}
                                _ = v17
                            } else {
                                v5 -= v5
                                var v18 float64 = 1e10
                                if true {
                                    fmt.Println(fi14, v5, p0)
                                }
                                v18 = v18
                                v19 := []int{-299}
                                var v20 uint16 = 937
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            _ = fi14
                        }
                        var v21 int = (fi13 - (fi13 >> ((2) & 63)))
                        _ = fi13
                        _ = v21
                    }
                    p0 = p0
                }
                for fi22 := 0; fi22 < 15; fi22++ {
                    v5 = -1
                    _ = fi22
                }
            }
        }
    }
    var v23 uint8 = 2
    _ = p0
    _ = v23
    _ = v5
    _ = v6
    return (65535 >= 85)
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 837
    var v3 int32 = v2
    v4 := make(map[int]bool)
    var v24 int8 = -1
    v3 = 0
    if true {
        var v25 int8 = (v24 / ((42) | 1))
        _ = v25
    } else {
        var v26 float64 = -2.25
        switch 4294967295 {
        case 56:
            v26 = v26
        case 786:
            if (376 == 955) {
                if false {
                    switch "" {
                    case "hello":
                        v2 -= v2
                    case "\n\t":
                        v26 -= float64(((1.5 - float32(float64(0.1))) * 1.5))
                    default:
                        v26 -= 1e10
                    }
                }
                v27 := []int{128, -9223372036854775808, -9223372036854775808}
                _ = v27
            }
        }
        v26 = v26
        _ = v26
    }
    var v28 uint64 = 1
    var v29 uint16 = 0
    v3 = (v3 | v3)
    if (!(!true)) {
        var v30 int64 = -9223372036854775808
        var v31 int8 = 1
        v2 = v3
        _ = v30
        _ = v31
    }
    v32 := GenS1{F0: -55, F1: 1}
    if ((242 > 3) || true) {
        v2 = (v2 & (-(v3)))
        v33 := GenS3{F0: v24, F1: v28}
        _ = v33
    } else {
        v34 := v4[42]
        _ = v34
    }
    v4[-72] = ((v28 & v28) != v28)
    if false {
        fmt.Println(v24, v28, v2)
    } else {
        for fi35 := 0; fi35 < 3; fi35++ {
            v28 = (v28 ^ (^(v28)))
            for fi36 := 0; fi36 < 10; fi36++ {
                v28 -= v28
                _ = fi36
            }
            for fi37 := 0; fi37 < 6; fi37++ {
                var v38 int8 = -128
                var v39 float64 = 3.14159
                v40 := GenS1{F0: int8(v28), F1: 943}
                if (float64((0.0 * 3.14159)) < v39) {
                    v28 = 2
                }
                v41 := []int{}
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi35
        }
    }
    v42 := v4[-1]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v42
}
