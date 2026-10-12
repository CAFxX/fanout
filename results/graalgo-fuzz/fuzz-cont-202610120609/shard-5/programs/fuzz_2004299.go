package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() bool {
    if false {
        v6 := []int{-9223372036854775808, 32768, 9223372036854775807}
        for fi7 := 0; fi7 < 10; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
        _ = v6
    } else {
        v8 := []int{0, 0, -1}
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint16 = 32767
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    v11 := GenS1{F0: 2}
    if ((-131 == 42) && (!false)) {
        v12 := []int{1, -2147483648}
        _ = v12
    }
    _ = v11
    return false
}
func genFn1(p0 int8) uint16 {
    fmt.Println(p0)
    switch 42 {
    case 127:
        p0 += p0
    case 47:
        if true {
            if (793 >= 255) {
                var v13 uint8 = 255
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 uint32 = 4294967295
                for fi16 := 0; fi16 < 15; fi16++ {
                    if false {
                        if false {
                            var v17 int32 = 42
                            _ = v17
                        }
                    }
                    v18 := GenS1{F0: 65535}
                    v15 = (v15 & v15)
                    _ = fi16
                    _ = v18
                }
                if false {
                    if false {
                        v15 -= (v15 + (v15 << ((1) & 63)))
                        var v19 uint64 = 18446744073709551615
                        _ = v19
                    }
                } else {
                    var v20 float64 = 1.5
                    v20 = (((v20 - 3.14159) / 1.0) * v20)
                    _ = v20
                }
                _ = fi14
                _ = v15
            }
            for fi21 := 0; fi21 < 11; fi21++ {
                for fi22 := 0; fi22 < 16; fi22++ {
                    p0 -= p0
                    v23 := []int{(fi22 % ((fi22) | 1)), fi22}
                    _ = fi22
                    _ = v23
                }
                var v24 bool = false
                _ = fi21
                _ = v24
            }
        }
    case 240:
        p0 = p0
    default:
        v25 := []int{-1}
        _ = v25
    }
    var v26 uint = 65535
    p0 = p0
    _ = p0
    _ = v26
    return 2
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 int8 = -128
    var v4 int16 = 32767
    var v5 uint8 = 2
    v5 += (v5 - 0)
    v27 := GenS1{F0: 255}
    var v28 uint32 = uint32(v5)
    _ = len(v2)
    v1["hello"] = true
    if false {
        if false {
            v29 := GenS1{F0: 1}
            _ = v29
        } else {
            v30 := GenS1{F0: 1}
            _ = v30
        }
        var v31 uint = 14
        fmt.Println(v5, v28, v5)
        _ = v31
    } else {
        if (!false) {
            delete(v2, "")
            if true {
                _ = len(v2)
                var v32 float64 = (3.14159 * (-2.25 - (1.5 / 0.1)))
                _ = v32
            }
            var v33 uint16 = 255
            _ = v33
        } else {
            var v34 int32 = 42
            if (true && true) {
                fmt.Println(v3)
                for fi35 := 0; fi35 < 8; fi35++ {
                    v28 = (v28 >> ((0) & 63))
                    v36 := GenS1{F0: 829}
                    v37 := GenS1{F0: 574}
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            switch 1 {
            case -9223372036854775808:
                var v38 int = -9223372036854775808
                _ = v38
            case -889:
                var v39 uint64 = 18446744073709551615
                _ = v39
            }
            _ = v34
        }
        var v40 uint64 = 458
        _ = v40
    }
    _ = len(v2)
    v41 := []int{4294967295, 2147483647, -893}
    var v42 uint32 = (^(v28))
    v41 = append(v41, -804)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
