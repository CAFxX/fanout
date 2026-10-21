package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int16; F2 int8 }
func genFn0() uint64 {
    switch -1 {
    case 16:
        var v7 int8 = 100
        _ = v7
    case 127:
        var v8 int16 = 0
        _ = v8
    case 258:
        if false {
            var v9 int8 = -1
            _ = v9
        }
    default:
        var v10 uint64 = 34
        _ = v10
    }
    if (false || false) {
        var v11 float32 = float32(((float64(1.5) + 3.14159) * (((1e10 - ((1e10 - float64((3.14159 + -2.25))) / (-2.25 * 0.1))) + float64(0.1)) / 1.0)))
        _ = v11
    } else {
        if (0 > 42) {
            switch -677 {
            case 683:
                var v12 int64 = -1
                _ = v12
            case 797:
                if false {
                    switch "\n\t" {
                    case "\n\t":
                        for fi13 := 0; fi13 < 17; fi13++ {
                            v14 := []int{fi13}
                            var v15 string = ("\n\t" + "hello")
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    case "a":
                        var v16 uint64 = 42
                        _ = v16
                    }
                }
            case 32767:
                var v17 uint = 255
                _ = v17
            default:
                var v18 int8 = 110
                _ = v18
            }
            var v19 uint16 = 42
            _ = v19
        } else {
            v20 := []int{-916, 2147483647, -46, 760}
            var v21 float32 = (0.0 - -2.25)
            _ = v20
            _ = v21
        }
    }
    switch "" {
    case "":
        var v22 int8 = -1
        _ = v22
    }
    v23 := []int{42, 127}
    _ = v23
    return 289
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{2147483648}
    var v3 uint16 = 770
    v4 := []int{532, -9223372036854775808}
    v5 := []string{"hello", "foo bar", ("\n\t" + "\n\t"), (("foo bar" + "\n\t") + "\n\t")}
    v6 := []uint32{}
    switch "" {
    case "foo bar":
        if false {
            v3 = 2
        } else {
            for fi24 := 0; fi24 < 3; fi24++ {
                if (!false) {
                    for fi25 := 0; fi25 < 15; fi25++ {
                        v26 := GenS1{F0: -27}
                        fmt.Println(v3, v3)
                        if false {
                            var v27 int8 = 0
                            v27 += (v27 / ((v27) | 1))
                            _ = v27
                        }
                        _ = fi25
                        _ = v26
                    }
                } else {
                    v28 := GenS2{F0: int32(fi24)}
                    var v29 int8 = 1
                    _ = v28
                    _ = v29
                }
                if false {
                    if (738 > -1) {
                        v30 := GenS1{F0: 76}
                        v3 -= v3
                        _ = v30
                    } else {
                        fmt.Println(fi24)
                    }
                } else {
                    var v31 int16 = 0
                    _ = v31
                }
                _ = fi24
            }
        }
    default:
        v1[65535] = 1
    }
    var v32 float32 = 1e10
    v3 += v3
    v33 := GenS1{F0: 100}
    v32 += (((-2.25 / v32) + v32) / ((((float32(1e10) - v32) + (-2.25 / v32)) / 1.0) * 1e10))
    v34 := GenS2{F0: -1}
    var v35 float64 = (1.5 / 1e10)
    if (87 > 42) {
        v36 := GenS3{F0: 74, F1: -295, F2: 0}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            if (!true) {
                if (1 < genFn0()) {
                    var v37 uint64 = genFn0()
                    _ = v37
                }
            }
            var v38 string = "a"
            if (true || false) {
                var v39 int32 = 510
                _ = v39
            } else {
                v32 = 1e10
            }
            _ = v38
        }
        _ = v36
    } else {
        fmt.Println(v35, v35)
        _ = len(v6)
    }
    v40 := GenS3{F0: 100, F1: 1, F2: 89}
    var v41 int64 = 1
    if (v35 != (((float64(v32) - (v35 + 0.0)) / v35) - 0.1)) {
        fmt.Println(v35, v32, v3)
    } else {
        v41 = v41
        for fi42 := 0; fi42 < 17; fi42++ {
            v35 -= (v35 * (((float64(v32) * (0.0 - (v35 * -2.25))) - (v35 / (0.0 / (v35 - v35)))) / v35))
            _ = fi42
        }
        v43 := GenS1{F0: 112}
        _ = v43
    }
    v44 := GenS2{F0: 1}
    var v45 bool = true
    v45 = (2147483647 < 362)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
