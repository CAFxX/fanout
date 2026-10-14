package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
func genFn0() uint32 {
    var v6 float64 = (0.0 + 0.1)
    v6 += float64((0.0 / ((1e10 + 0.0) * (1e10 / 1.0))))
    if true {
        v7 := []int{1, 32768, 127}
        _ = v7
    }
    fmt.Println(v6)
    _ = v6
    return 1
}
func genFn1(p0 int16, p1 uint32, p2 int16) int64 {
    p1 += (p1 ^ genFn0())
    if false {
        p0 -= -1
    } else {
        p1 = (p1 | (p1 - (p1 % ((genFn0()) | 1))))
        p2 = p2
        var v8 uint16 = 255
        _ = v8
    }
    var v9 int = -303
    v10 := []int{-325, (v9 | 32768), (v9 * (-(v9))), (v9 << ((554) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 430
    var v2 uint16 = 0
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v5 float32 = 3.14159
    if (-341 >= -705) {
        var v11 int64 = -621
        _ = v11
    }
    v2 = (v2 ^ 0)
    var v12 uint32 = 2147483648
    for fi13 := 0; fi13 < 2; fi13++ {
        v2 -= uint16(v12)
        switch (fi13 * 2147483647) {
        case -910:
            if (fi13 > -9223372036854775808) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 float64 = 0.0
                    fmt.Println(v15, v2)
                    var v16 float32 = v5
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                fmt.Println(v5)
            }
        default:
            switch (("\n\t" + "foo bar") + "") {
            case "a":
                var v17 uint16 = v2
                _ = v17
            case "hello":
                if (("a" > (((("hello" + "") + "\n\t") + "hello") + "\n\t")) || (!false)) {
                    for fi18 := 0; fi18 < 7; fi18++ {
                        v19 := GenS1{F0: 4294967295}
                        _ = fi18
                        _ = v19
                    }
                } else {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        fmt.Println(v1, v2)
                        _ = fi20
                    }
                    v3[(v12 << ((65535) & 63))] = true
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 int16 = 217
                        v23 := []int{21}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                }
            case "":
                v24 := GenS2{F0: (^(fi13)), F1: 255, F2: true}
                _ = v24
            default:
                var v25 int = (fi13 & 1)
                _ = v25
            }
        }
        v26 := []int{}
        _ = fi13
        _ = v26
    }
    v12 = 2
    v27 := []int{128, -9223372036854775808}
    _ = len(v3)
    if (int32(v1) == -695) {
        v28 := v3[v12]
        v29 := GenS2{F0: 2147483647, F1: 143, F2: true}
        v30 := GenS1{F0: 461}
        _ = v28
        _ = v29
        _ = v30
    }
    v1 = 42
    var v31 bool = (false && true)
    if (("" <= "foo bar") && v31) {
        if v31 {
            v31 = v31
            v2 = v2
            fmt.Println(v12)
        }
        for fi32 := 0; fi32 < 17; fi32++ {
            _ = len(v27)
            fmt.Println(v5, v2, v31)
            v1 = v1
            _ = fi32
        }
        var v33 float64 = (3.14159 + ((0.1 * 0.1) - (0.1 + 3.14159)))
        _ = v33
    }
    var v34 uint16 = 836
    fmt.Println(v2)
    v1 = v1
    if (4294967295 == 127) {
        for fi35 := 0; fi35 < 6; fi35++ {
            var v36 int = (fi35 & fi35)
            var v37 uint = uint(v1)
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    var v38 int32 = -1
    var v39 float64 = float64(float32((3.14159 - (1.5 / ((((-2.25 - float64(v5)) + 1.5) * ((1.5 - 1.5) / 0.1)) * 1e10)))))
    for fi40 := 0; fi40 < 20; fi40++ {
        v39 += (-2.25 + v39)
        for fi41 := 0; fi41 < 14; fi41++ {
            v42 := GenS1{F0: 42}
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            if v31 {
                v38 += (v38 * (v38 - (v38 / ((-574) | 1))))
                fmt.Println(v12)
                v2 = v2
            }
            var v43 int16 = 1000
            switch fi41 {
            case -167:
                delete(v4, (fi40 * fi41))
            default:
                v39 = (float64(v5) * float64(v5))
            }
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
