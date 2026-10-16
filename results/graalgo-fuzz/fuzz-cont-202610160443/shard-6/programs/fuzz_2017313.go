package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint; F2 string }
type GenS3 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int8 }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := []int{(fi5 / ((fi6) | 1)), fi5}
            var v8 float32 = 1.5
            switch fi6 {
            case 32768:
                var v9 int64 = 42
                v10 := GenS1{F0: 0, F1: 325, F2: -1}
                var v11 string = ("foo bar" + "a")
                _ = v10
                _ = v11
                _ = v9
            case 32767:
                v8 -= (0.0 - float32(1e10))
            case 309:
                v12 := GenS2{F0: 2, F1: 0, F2: (("a" + "") + "\n\t")}
                _ = v12
            default:
                fmt.Println(fi5, fi6)
            }
            var v13 int8 = 51
            v8 = v8
            _ = fi6
            _ = v13
            _ = v7
            _ = v8
        }
        fmt.Println(fi5)
        _ = fi5
    }
    var v14 string = "foo bar"
    v15 := GenS3{F0: -9223372036854775808, F1: 4294967295, F2: 876}
    switch "a" {
    case "hello":
        v16 := GenS4{F0: 127}
        _ = v16
    default:
        v17 := []int{128, 9223372036854775807, -1}
        _ = v17
    }
    v18 := []int{}
    _ = v14
    _ = v15
    _ = v18
    return 847
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 624
    var v3 uint32 = 689
    v4 := []int{-1}
    v3 += (^(v3))
    v2 -= (^(v2))
    var v19 float64 = -2.25
    fmt.Println(v19, v3)
    if ((v2 >> ((0) & 63)) <= 42) {
        v3 = 2
    } else {
        for fi20 := 0; fi20 < 8; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                v2 += v3
                v3 -= (^(v3))
                v19 = ((v19 / (float64(float32(v19)) - v19)) * ((v19 + float64(3.14159)) + v19))
                _ = fi21
            }
            if len(v4) > 0 {
                v4[0] = (fi20 + fi20)
            }
            if (-850 <= 1) {
                if false {
                    var v22 bool = ((float64(float32(((v19 - 1e10) - v19))) - 1e10) > 0.1)
                    if v22 {
                        for fi23 := 0; fi23 < 5; fi23++ {
                            fmt.Println(v22, fi20, fi20)
                            _ = len(v1)
                            _ = fi23
                        }
                        _ = len(v1)
                    }
                    var v24 int = (^(fi20))
                    _ = v22
                    _ = v24
                }
                for fi25 := 0; fi25 < 3; fi25++ {
                    fmt.Println(v2)
                    if (!(4294967295 >= 42)) {
                        if true {
                            var v26 float64 = (v19 / (0.0 * ((float64(1.5) - 1e10) - v19)))
                            v27 := v1[v3]
                            delete(v1, (v2 << ((869) & 63)))
                            _ = v26
                            _ = v27
                        } else {
                            var v28 string = ((("a" + "hello") + (("\n\t" + "hello") + "\n\t")) + "\n\t")
                            _ = v28
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    var v29 uint8 = 0
                    _ = fi25
                    _ = v29
                }
            }
            _ = fi20
        }
        v1[763] = (-1 > 0)
    }
    var v30 bool = (!(true || true))
    v31 := GenS1{F0: 255, F1: 1000, F2: int32(v2)}
    var v32 bool = v30
    v19 = v19
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
