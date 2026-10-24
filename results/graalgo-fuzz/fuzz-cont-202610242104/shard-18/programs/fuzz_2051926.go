package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 int32) uint {
    v6 := []int{427, 308, 32767, 0}
    v6 = append(v6, 835)
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 42
}
func genFn1(p0 uint8, p1 uint) uint32 {
    var v7 float32 = float32(1e10)
    v8 := []int{4294967295, 364, -2147483648}
    if (!true) {
        v7 += (((v7 * float32((-2.25 * (float64(0.1) + (((3.14159 - (((((float64(v7) - (1e10 / ((((float64(((1.5 + 0.1) * (-2.25 * ((float32(0.1) - float32(float64(0.0))) / v7)))) * (1e10 - 1e10)) / (3.14159 * ((1.5 + (-2.25 * (float64(1e10) - (1.5 - (float64(v7) + 0.1))))) - 1e10))) + 3.14159) * (((1.5 * float64((-2.25 - v7))) / (float64((0.1 + (0.0 - v7))) + ((1e10 - 3.14159) + 1e10))) / 0.1)))) / 0.1) - 1e10) * (float64((3.14159 * float32(0.1))) - float64((-2.25 - (v7 - ((v7 * (v7 + v7)) * v7)))))) - 1.5)) / 1.0) * 3.14159))))) * float32(-2.25)) / v7)
        if (32767 != 32767) {
            switch 463 {
            case 42:
                for fi9 := 0; fi9 < 16; fi9++ {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    v8 = append(v8, fi9)
                    _ = fi9
                }
            case 127:
                v10 := GenS1{F0: 2147483647, F1: 128}
                _ = v10
            case -1:
                for fi11 := 0; fi11 < 3; fi11++ {
                    p0 += p0
                    _ = fi11
                }
            }
            var v12 int64 = -1
            v13 := GenS1{F0: -2147483648, F1: (p0 - 127)}
            _ = v12
            _ = v13
        }
    }
    var v14 float32 = (v7 - ((0.0 * (float32(1e10) + v7)) - 0.1))
    v15 := GenS1{F0: -2147483648, F1: (p0 << (((p1 >> ((1) & 63))) & 63))}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := make(map[int]uint32)
    var v5 uint = 1
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 string = ""
        _ = fi16
        _ = v17
    }
    if (true || true) {
        for fi18 := 0; fi18 < 5; fi18++ {
            if false {
                if true {
                    var v19 uint8 = 127
                    _ = v19
                }
                v2["foo bar"] = ("a" + "")
                if true {
                    var v20 float64 = (((-2.25 - (1.5 / 3.14159)) / float64(0.1)) - 1e10)
                    v1 += (v1 / ((genFn1(42, 1)) | 1))
                    _ = v20
                }
            } else {
                v21 := v4[fi18]
                switch (fi18 << (((v5 % (((^(v5))) | 1))) & 63)) {
                case -2147483648:
                    fmt.Println(v1, fi18)
                case -252:
                    var v22 float64 = (((float64(((-2.25 * (3.14159 + (1e10 / float32(1e10)))) * 1.5)) - ((0.0 - (0.0 - 1.5)) * float64((-2.25 / (3.14159 - 0.1))))) / 1e10) + 1e10)
                    _ = v22
                }
                _ = v21
            }
            v23 := []int{-295, (fi18 + 32768)}
            _ = fi18
            _ = v23
        }
        var v24 uint16 = 1
        v5 -= v5
        _ = v24
    }
    fmt.Println(v1)
    v2["\n\t"] = "\n\t"
    fmt.Println(v1)
    v25 := []int{}
    var v26 float64 = -2.25
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 string = (("foo bar" + "a") + ("a" + "hello"))
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
