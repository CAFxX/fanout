package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint32 }
func genFn0(p0 int32, p1 uint, p2 uint32) int16 {
    v4 := GenS1{F0: 101}
    var v5 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 32767
}
func genFn1(p0 int) uint64 {
    if true {
        v6 := GenS2{F0: 42, F1: 4294967295}
        _ = v6
    }
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int64 = 42
        if (1e10 < 0.1) {
            v10 := []int{(p0 ^ -664), 0, p0, fi8}
            _ = v10
        } else {
            v11 := GenS2{F0: 628, F1: 1}
            var v12 uint8 = 255
            switch "" {
            case "hello":
                v9 = v9
            case "":
                v12 -= v12
            }
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    v13 := GenS3{F0: 2, F1: -1, F2: 2147483648}
    v7 += v7
    _ = p0
    _ = v13
    _ = v7
    return 0
}

func main() {
    var v1 float32 = (0.0 / 1.0)
    v2 := []uint32{903, 350, 757, 87}
    v3 := make(map[int]uint32)
    if ((0 == -1) && (-2.25 > (-2.25 + (-2.25 * float64(v1))))) {
        switch (("\n\t" + "\n\t") + "\n\t") {
        case "":
            switch 127 {
            case -262:
                fmt.Println(v1)
            default:
                if ((v1 + (0.0 * (v1 - 1.5))) >= (float32((1.5 / ((1e10 + ((0.1 * ((((float64(float32((((((0.0 + (0.1 * 1e10)) + 0.0) / 1.0) / 1.0) * 0.0))) * 0.1) - ((-2.25 * -2.25) * ((1.5 - float64(v1)) + 1e10))) - (1.5 * (float64(v1) - 1.5))) - float64((v1 - v1)))) * 1e10)) / (0.0 * (((3.14159 / (float64((v1 - -2.25)) / 1.5)) + 1.5) + 0.0))))) + v1)) {
                    v14 := GenS3{F0: 535, F1: -28, F2: 2}
                    _ = v14
                }
            }
        }
        v15 := GenS3{F0: genFn1(-9223372036854775808), F1: 462, F2: 42}
        v1 = v1
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        if true {
            var v17 bool = true
            v18 := GenS1{F0: 42}
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    if (!((0 <= 108) && (true && (2147483648 != 555)))) {
        v1 -= v1
        v3[-607] = 65535
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v19 bool = true
    fmt.Println(v19, v1)
    delete(v3, 195)
    fmt.Println(v1, v1)
    if (0 < 1) {
        for fi20 := 0; fi20 < 1; fi20++ {
            switch (fi20 | (fi20 ^ 32768)) {
            case -2147483648:
                v19 = (128 != 177)
            case -9223372036854775808:
                fmt.Println(v19, v1)
            case 32768:
                var v21 int32 = -1
                _ = v21
            }
            fmt.Println(v19, v1, fi20)
            _ = fi20
        }
        v1 += v1
        fmt.Println(v19)
    }
    delete(v3, 42)
    v19 = ((((0.0 - 0.0) / ((-2.25 * ((1.5 * (((0.1 * 0.1) / (1.5 - 1e10)) / 1.0)) * -2.25)) - 0.0)) > 1e10) && false)
    if v19 {
        for fi22 := 0; fi22 < 8; fi22++ {
            switch 128 {
            case -2147483648:
                fmt.Println(v19, fi22, v1)
            case 0:
                v19 = (!v19)
            }
            _ = fi22
        }
        v1 += -2.25
    } else {
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := v3[32768]
            v25 := GenS3{F0: 116, F1: 765, F2: 42}
            var v26 int64 = 42
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 2; fi27++ {
            fmt.Println(v1, v1, v1)
            _ = fi27
        }
        var v28 uint8 = 0
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
