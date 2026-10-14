package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
func genFn0() uint {
    v6 := []int{}
    if (((1 == 1) && ((!(!((0 < 623) && true))) || true)) || false) {
        for fi7 := 0; fi7 < 17; fi7++ {
            v6 = append(v6, (fi7 << ((0) & 63)))
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 string = ""
                switch fi7 {
                case 9223372036854775807:
                    if (fi7 < (fi7 - fi7)) {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    } else {
                        var v10 float32 = float32((0.1 - float64(-2.25)))
                        _ = v10
                    }
                }
                v9 = v9
                _ = fi8
                _ = v9
            }
            var v11 uint32 = 0
            _ = fi7
            _ = v11
        }
    } else {
        var v12 int64 = 925
        _ = v12
    }
    if (42 > 726) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        var v13 uint16 = 1
        var v14 int16 = 270
        _ = v13
        _ = v14
    }
    if true {
        v15 := GenS1{F0: "", F1: -905}
        _ = v15
    } else {
        var v16 int64 = 1
        var v17 string = (("\n\t" + ("foo bar" + ("hello" + "foo bar"))) + "\n\t")
        _ = v16
        _ = v17
    }
    _ = v6
    return 361
}
func genFn1(p0 uint64, p1 uint32) uint64 {
    var v18 uint = 4294967295
    var v19 float64 = (1e10 * ((((1.5 + ((float64(3.14159) - 0.1) * float64(float32((0.0 + (0.0 / ((0.0 - (float64(((float32((((0.1 - (float64((1e10 + 0.1)) * 3.14159)) + (((0.1 * 0.1) / 3.14159) - (1e10 - (((3.14159 + 1e10) * -2.25) - -2.25)))) + 1e10)) / 0.1) / (3.14159 + ((0.1 - -2.25) / (-2.25 + 1e10))))) / ((0.1 + -2.25) + 1.5))) / (1e10 + -2.25)))))))) - 3.14159) - (((1.5 * 1.5) - (-2.25 - float64(0.1))) / (1.5 + 1e10))) * ((0.0 * (1.5 / 3.14159)) + float64(float32(0.0)))))
    if true {
        var v20 int = 1
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return (p0 - (p0 - (p0 >> (((^(v18))) & 63))))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{2, 2147483648}
    var v3 int = 0
    var v4 int16 = 856
    var v5 string = "hello"
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int = (v3 | (v3 << ((genFn0()) & 63)))
        for fi23 := 0; fi23 < 17; fi23++ {
            _ = len(v2)
            _ = fi23
        }
        v24 := GenS1{F0: v5, F1: v4}
        _ = fi21
        _ = v22
        _ = v24
    }
    var v25 string = "a"
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS1{F0: (((v25 + (("a" + ("\n\t" + v25)) + v25)) + v25) + (v5 + "a")), F1: v4}
        _ = len(v2)
        _ = fi26
        _ = v27
    }
    v3 = v3
    if len(v2) > 0 {
        v2[0] = 65535
    }
    var v28 string = v5
    v25 = (v25 + "\n\t")
    v4 = (v4 + v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
