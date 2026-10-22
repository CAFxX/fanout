package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int32 }
func genFn0(p0 int64, p1 int16) int {
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int = 1
        var v7 int32 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{127, -2147483648}
    p1 = (p1 / ((p1) | 1))
    p1 = p1
    p0 = p0
    _ = p0
    _ = p1
    _ = v8
    return 0
}
func genFn1(p0 int, p1 int64, p2 int32) uint8 {
    fmt.Println(p1, p0)
    v9 := []int{p0, p0, p0, (p0 ^ genFn0(0, 0))}
    fmt.Println(p1, p2)
    v10 := GenS1{F0: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 float32 = ((0.1 * -2.25) * 0.1)
    v2 := make(map[uint32]int)
    var v3 uint64 = 2
    v4 := make(map[int]string)
    switch "" {
    case "":
        v11 := GenS2{F0: 34, F1: 566}
        _ = v11
    default:
        v1 = float32(0.0)
    }
    v3 = (v3 * (v3 | v3))
    var v12 uint64 = 4294967295
    v13 := GenS1{F0: genFn1(genFn0(42, 42), 1, 0)}
    var v14 int8 = -92
    _ = len(v4)
    if (float32((((((((float64(v1) * -2.25) * float64(v1)) / (float64(v1) / (1.5 - float64(float32(0.1))))) / 0.1) + 0.0) / (-2.25 + 0.1)) * 1.5)) < (0.0 - (float32(((1e10 / 1e10) + -2.25)) * v1))) {
        var v15 uint8 = 255
        _ = v15
    }
    switch -2147483648 {
    case -307:
        if (!(1 <= 456)) {
            v1 += v1
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 uint = 255
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := GenS2{F0: 654, F1: 1}
                    if true {
                        var v20 uint8 = 0
                        v21 := []int{fi16, genFn0(-1, -36), (fi16 * 4294967295), 0}
                        v22 := GenS1{F0: 114}
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        if ((v14 >> ((uint(v3)) & 63)) > 100) {
                            fmt.Println(v14)
                            switch "a" {
                            case "a":
                                if (!true) {
                                    var v23 float32 = (3.14159 / (float32(0.0) + v1))
                                    _ = v23
                                }
                            case "":
                                v12 += (v3 % ((v3) | 1))
                            case "hello":
                                switch 42 {
                                case 42:
                                    switch (fi18 ^ fi16) {
                                    case 32768:
                                        v17 -= 4294967295
                                    case 2147483647:
                                        var v24 uint16 = 65535
                                        _ = v24
                                    case 597:
                                        v25 := GenS1{F0: 4}
                                        _ = v25
                                    }
                                default:
                                    fmt.Println(v17, v17, v14)
                                }
                            default:
                                var v26 uint32 = 1
                                _ = v26
                            }
                        }
                        v27 := v4[65]
                        fmt.Println(v3)
                        v17 -= v17
                        _ = v27
                    }
                    _ = fi18
                    _ = v19
                }
                if true {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        v29 := []int{(fi28 ^ fi28)}
                        _ = fi28
                        _ = v29
                    }
                    v4[fi16] = "hello"
                } else {
                    for fi30 := 0; fi30 < 9; fi30++ {
                        v31 := []int{int(v17), (-(fi30))}
                        var v32 string = "foo bar"
                        var v33 int16 = 1000
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                }
                _ = fi16
                _ = v17
            }
        }
    case -530:
        if true {
            v3 = (v3 & 0)
            delete(v2, 284)
        } else {
            v34 := GenS1{F0: 205}
            _ = v34
        }
    case 32768:
        v35 := []int{42, genFn0(-1, 42), genFn0(-1, 90), 32768}
        _ = v35
    default:
        v36 := GenS1{F0: 2}
        _ = v36
    }
    v12 = (^(v3))
    var v37 string = "a"
    var v38 uint32 = 2147483648
    if (((18446744073709551615 < (v3 & 989)) || ((false && false) && (0 > genFn0(682, 1)))) || (0 >= 1)) {
        if true {
            var v39 int32 = -389
            var v40 string = "foo bar"
            _ = v39
            _ = v40
        }
        fmt.Println(v3, v12)
        fmt.Println(v1)
    } else {
        var v41 int = 0
        v42 := []int{v41, (v41 & genFn0(-9223372036854775808, 32767)), int(v14), v41}
        _ = v41
        _ = v42
    }
    var v43 uint64 = 642
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v43
}
