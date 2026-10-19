package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint; F2 int8 }
type GenS5 struct { F0 int16; F1 int; F2 int32 }
func genFn0(p0 bool, p1 uint) string {
    var v4 int32 = -2147483648
    v5 := GenS1{F0: 2, F1: 32768, F2: 1}
    var v6 uint16 = 790
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return ("\n\t" + ((("\n\t" + "\n\t") + "foo bar") + "hello"))
}
func genFn1(p0 int) int {
    v7 := []int{-876, p0, (p0 % ((-144) | 1)), -564}
    p0 = (p0 - p0)
    _ = p0
    _ = v7
    return p0
}

func main() {
    var v1 float64 = (1.5 / 1e10)
    var v2 int8 = 42
    var v3 int64 = 9223372036854775807
    for fi8 := 0; fi8 < 14; fi8++ {
        v3 = (v3 & (v3 << ((978) & 63)))
        v1 += (0.1 + ((v1 * v1) - v1))
        _ = fi8
    }
    if true {
        v9 := GenS2{F0: v3}
        v1 = -2.25
        v10 := GenS1{F0: 969, F1: -2147483648, F2: (v2 + (v2 / (((v2 | v2)) | 1)))}
        _ = v10
        _ = v9
    } else {
        v1 -= 0.1
    }
    v11 := []int{1, 9223372036854775807, -319, 2147483647}
    if (1 != 42) {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 uint64 = 4294967295
            switch genFn1(genFn1((fi12 - (fi12 - (fi12 | 42))))) {
            case -64:
                for fi14 := 0; fi14 < 4; fi14++ {
                    var v15 float32 = 1e10
                    v15 += v15
                    if (v13 < (v13 * 0)) {
                        v15 += ((v15 * 3.14159) + (v15 * (v15 / ((1.5 + v15) / 1.0))))
                    } else {
                        v3 -= 1
                        switch "" {
                        case "a":
                            var v16 int64 = v3
                            _ = v16
                        }
                    }
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    v11 = append(v11, genFn1(2147483647))
                    _ = fi14
                    _ = v15
                }
            case -110:
                v1 -= float64(1e10)
            }
            v17 := GenS3{F0: -158}
            _ = fi12
            _ = v13
            _ = v17
        }
        if len(v11) > 0 {
            v11[0] = 32768
        }
        v18 := GenS4{F0: true, F1: 42, F2: 42}
        _ = v18
    } else {
        v19 := GenS2{F0: -9223372036854775808}
        v3 = v3
        if (128 != 255) {
            v3 = v3
            var v20 uint8 = 133
            _ = v20
        }
        _ = v19
    }
    if (true || true) {
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 6; fi22++ {
                v2 = 42
                v3 = v3
                _ = fi22
            }
            v2 = v2
            if ((v3 << ((2) & 63)) < v3) {
                v23 := GenS1{F0: 42, F1: -9223372036854775808, F2: v2}
                var v24 float64 = -2.25
                _ = v23
                _ = v24
            }
            _ = fi21
        }
        if (!false) {
            switch "" {
            case "\n\t":
                fmt.Println(v3, v1)
            case "a":
                v25 := GenS5{F0: 1000, F1: genFn1(-16), F2: -2147483648}
                _ = v25
            }
            var v26 int32 = 0
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = v26
        } else {
            fmt.Println(v1)
            var v27 int = 752
            _ = v27
        }
        var v28 float64 = v1
        _ = v28
    }
    var v29 string = "foo bar"
    fmt.Println(v3, v3)
    v3 = (v3 + (v3 | v3))
    v1 = float64((-2.25 / (0.1 + (float32(v1) / float32(v1)))))
    var v30 bool = (v1 < float64((-2.25 / (1.5 + 1e10))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
