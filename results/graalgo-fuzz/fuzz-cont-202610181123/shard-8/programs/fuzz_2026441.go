package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 string }
type GenS2 struct { F0 int16; F1 uint64; F2 int8 }
type GenS3 struct { F0 int; F1 bool; F2 uint }
func genFn0(p0 uint32) int64 {
    v6 := GenS1{F0: 32767, F1: 56, F2: ("\n\t" + "\n\t")}
    var v7 uint64 = 1
    v8 := GenS1{F0: 1, F1: 0, F2: "foo bar"}
    v9 := GenS1{F0: -32768, F1: 1, F2: "foo bar"}
    var v10 string = "hello"
    _ = p0
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 996
}

func main() {
    var v1 int32 = -45
    var v2 int16 = -603
    v3 := []int{807, 646}
    v4 := make(map[uint32]int)
    var v5 float64 = (1.5 + 1e10)
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 uint = 2
        var v13 uint = v12
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 float32 = 1.5
        v2 -= (^(v2))
        fmt.Println(fi14, v5, fi14)
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 1, F1: 2, F2: 127}
    fmt.Println(v1, v1, v5)
    v2 = (-(v2))
    v17 := v4[65535]
    var v18 int16 = 362
    var v19 uint64 = 836
    v20 := GenS3{F0: v17, F1: ((float64((0.1 * 0.0)) * float64(((0.0 / ((-2.25 * ((float32(((1.5 / float64(float32(v5))) - ((((float64(3.14159) - v5) + ((v5 + 1e10) / (3.14159 / v5))) + v5) + (v5 + 1e10)))) - 3.14159) + (1e10 * 1e10))) * (float32(1e10) - (3.14159 * float32(((v5 / (1e10 * (((0.1 / v5) - 3.14159) + (0.0 * v5)))) - v5)))))) - (3.14159 - float32(0.1))))) > v5), F2: 255}
    switch (v17 << ((365) & 63)) {
    case -952:
        switch v17 {
        case 11:
            for fi21 := 0; fi21 < 6; fi21++ {
                for fi22 := 0; fi22 < 19; fi22++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 += v18
                    var v23 uint32 = 0
                    v2 = (v18 ^ v18)
                    v1 -= 305
                    var v24 uint64 = 42
                    v25 := GenS2{F0: -1, F1: (v19 / ((v24) | 1)), F2: 100}
                    var v26 float32 = 0.1
                    fmt.Println(v1, v24)
                    v3 = append(v3, (fi21 << ((1) & 63)))
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                var v27 uint16 = 420
                if true {
                    var v28 uint8 = 248
                    v29 := GenS3{F0: (^(v17)), F1: true, F2: 0}
                    fmt.Println(v5, v28, v17)
                    _ = v28
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 9; fi30++ {
                        if (!false) {
                            v19 += 4294967295
                            v5 += v5
                        }
                        fmt.Println(fi21, v5, v2)
                        if ((v17 + (-(fi21))) <= (^(fi21))) {
                            v27 = (v27 & (v27 << ((255) & 63)))
                        } else {
                            fmt.Println(v17, fi30, v27)
                            v19 = (^(v19))
                            var v31 int = int(v19)
                            v31 = (v31 & (-(v31)))
                            _ = v31
                        }
                        _ = fi30
                    }
                }
                _ = fi21
                _ = v27
            }
        case 1:
            var v32 uint64 = 4294967295
            _ = v32
        }
    case 42:
        var v33 uint8 = 2
        _ = v33
    default:
        if (208 > 0) {
            v34 := GenS3{F0: v17, F1: (-2.25 != float32((v5 / ((0.0 / (1.5 + (v5 - v5))) / (-2.25 + 3.14159))))), F2: 18446744073709551615}
            var v35 int16 = v18
            _ = v34
            _ = v35
        } else {
            for fi36 := 0; fi36 < 10; fi36++ {
                if true {
                    v2 = v2
                }
                _ = fi36
            }
        }
    }
    switch "hello" {
    case "":
        v37 := GenS1{F0: (v2 >> ((4294967295) & 63)), F1: -497, F2: "a"}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
