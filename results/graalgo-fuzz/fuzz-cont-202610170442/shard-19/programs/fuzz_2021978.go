package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
func genFn0() int {
    var v6 float64 = ((float64(float32(-2.25)) + 1e10) - 1e10)
    v7 := []int{-9223372036854775808, 107, -756}
    v6 = 1.5
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 string = "hello"
    var v2 int8 = 127
    var v3 string = v1
    v4 := make(map[uint32]int)
    var v5 int64 = 364
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 42, F1: 255}
        _ = fi8
        _ = v9
    }
    var v10 bool = true
    v11 := GenS2{F0: 0}
    var v12 int32 = -534
    v13 := []int{-76, 1, 32767, -2147483648}
    if v10 {
        var v14 int16 = -32768
        if (!(v2 <= v2)) {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 bool = v10
                _ = fi15
                _ = v16
            }
            var v17 int64 = v5
            _ = v17
        }
        _ = v14
    }
    v1 = "hello"
    v18 := GenS1{F0: 32767, F1: 970}
    if ((v5 << ((892) & 63)) <= (v5 << ((699) & 63))) {
        for fi19 := 0; fi19 < 4; fi19++ {
            if len(v13) > 0 {
                v13[0] = -9223372036854775808
            }
            var v20 int32 = 0
            v21 := GenS1{F0: 633, F1: 1}
            _ = fi19
            _ = v20
            _ = v21
        }
        fmt.Println(v10)
    } else {
        v12 -= (^(v12))
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        switch v1 {
        case "":
            for fi23 := 0; fi23 < 13; fi23++ {
                v24 := v4[1]
                v25 := GenS1{F0: 2, F1: 65535}
                v10 = true
                _ = fi23
                _ = v24
                _ = v25
            }
        case "a":
            var v26 uint32 = 2147483648
            _ = v26
        case "hello":
            var v27 string = "\n\t"
            _ = v27
        }
        v13 = append(v13, fi22)
        _ = fi22
    }
    v28 := GenS1{F0: 2, F1: 400}
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := GenS1{F0: 2, F1: 65535}
        for fi31 := 0; fi31 < 3; fi31++ {
            if false {
                var v32 uint32 = 0
                _ = len(v4)
                if (2 < 18446744073709551615) {
                    switch v3 {
                    case "":
                        v12 = v12
                    case "foo bar":
                        fmt.Println(v3, v5)
                    case "a":
                        if v10 {
                            var v33 int32 = (v12 - (v12 | v12))
                            var v34 float32 = -2.25
                            _ = v33
                            _ = v34
                        } else {
                            if false {
                                if len(v13) > 0 {
                                    v13[0] = genFn0()
                                }
                            } else {
                                if false {
                                    v1 = ((v3 + "a") + v1)
                                    v4[(v32 ^ 42)] = fi31
                                    _ = len(v4)
                                }
                                var v35 uint8 = 2
                                var v36 int8 = v2
                                _ = v35
                                _ = v36
                            }
                            fmt.Println(v12)
                            v4[v32] = (fi31 & (fi31 * (-(fi31))))
                        }
                    }
                    var v37 bool = v10
                    v38 := GenS3{F0: -939}
                    _ = v37
                    _ = v38
                }
                v39 := v4[v32]
                _ = v32
                _ = v39
            }
            var v40 int64 = 9223372036854775807
            _ = fi31
            _ = v40
        }
        _ = fi29
        _ = v30
    }
    var v41 int8 = -69
    v2 = v41
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
