package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0(p0 string, p1 uint64) uint32 {
    var v6 int32 = 0
    if true {
        v7 := GenS1{F0: 18446744073709551615, F1: 116}
        v8 := []int{-1}
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int32 = v6
        v6 -= v6
        if false {
            var v11 string = p0
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    var v12 uint64 = 35
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    return 270
}

func main() {
    var v1 float64 = 1e10
    var v2 uint8 = 2
    var v3 float32 = ((0.1 / float32(0.1)) / (-2.25 + 0.0))
    var v4 int8 = 42
    v5 := make(map[uint32]int)
    var v13 uint64 = 320
    v1 = float64(v3)
    var v14 float32 = v3
    v14 += float32(((-2.25 - (v1 + 0.1)) * v1))
    if ((v1 * v1) > v1) {
        v15 := []int{}
        if (971 != -2147483648) {
            v14 = v3
            var v16 int16 = 0
            _ = len(v5)
            _ = v16
        }
        switch 128 {
        case 188:
            if (v2 == (v2 % ((v2) | 1))) {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v18 := GenS1{F0: 42, F1: 0}
                    var v19 float32 = 0.0
                    switch ("" + "foo bar") {
                    case "a":
                        v5[genFn0("", 2)] = fi17
                    case "":
                        if true {
                            v2 += v2
                            var v20 int = 127
                            _ = v20
                        }
                    case "hello":
                        if (42 != 911) {
                            if (v1 == (((v1 / v1) * ((v1 - v1) + 1e10)) * (v1 / (v1 - ((v1 / v1) / ((3.14159 - v1) - ((0.1 / v1) * v1))))))) {
                                v21 := GenS1{F0: 18446744073709551615, F1: 32767}
                                v2 = 255
                                _ = v21
                            } else {
                                if (uint8(v13) == (v2 << ((518) & 63))) {
                                    v2 += 122
                                    fmt.Println(v19)
                                    if (928 > 596) {
                                        switch (fi17 & (^(fi17))) {
                                        case 4294967295:
                                            var v22 string = "a"
                                            _ = v22
                                        }
                                        v2 -= (v2 / (((v2 | (v2 * (^(v2))))) | 1))
                                    }
                                }
                                var v23 uint16 = 394
                                var v24 float64 = (v1 - -2.25)
                                _ = v23
                                _ = v24
                            }
                            fmt.Println(v13, v3, v14)
                        }
                    }
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            } else {
                v25 := GenS2{F0: 303, F1: 42}
                v4 += (v4 ^ v4)
                v26 := GenS3{F0: 475, F1: (v13 % ((v13) | 1))}
                _ = v25
                _ = v26
            }
        }
        _ = v15
    } else {
        fmt.Println(v3, v1, v14)
        for fi27 := 0; fi27 < 19; fi27++ {
            for fi28 := 0; fi28 < 2; fi28++ {
                v14 -= v3
                v2 += (v2 ^ v2)
                v29 := GenS4{F0: 522, F1: 1, F2: -1}
                if true {
                    var v30 int16 = -1
                    _ = v30
                } else {
                    if true {
                        v31 := GenS3{F0: 2, F1: 495}
                        v14 = v3
                        fmt.Println(v3)
                        _ = v31
                    } else {
                        var v32 uint16 = 532
                        v33 := GenS2{F0: 915, F1: 32767}
                        if true {
                            v1 += v1
                        } else {
                            v3 = v3
                        }
                        _ = v32
                        _ = v33
                    }
                    switch (fi27 & fi28) {
                    case -1:
                        v34 := v5[1]
                        _ = v34
                    }
                }
                v35 := GenS3{F0: 4294967295, F1: (v13 - (v13 % ((137) | 1)))}
                _ = fi28
                _ = v29
                _ = v35
            }
            _ = fi27
        }
    }
    if (443 > 255) {
        var v36 int16 = -625
        _ = v36
    }
    v2 -= 135
    v37 := GenS1{F0: 1, F1: 0}
    var v38 int = 98
    for fi39 := 0; fi39 < 1; fi39++ {
        for fi40 := 0; fi40 < 13; fi40++ {
            v3 = v3
            v41 := []int{(fi40 + 1), -1, v38, fi39}
            if (2147483647 <= 555) {
                if false {
                    if len(v41) > 0 {
                        v41[0] = (^(v38))
                    }
                    var v42 uint8 = 0
                    _ = v42
                }
                v13 += (v13 ^ 1)
            } else {
                var v43 int16 = 0
                var v44 uint8 = 110
                _ = v43
                _ = v44
            }
            _ = fi40
            _ = v41
        }
        v5[2147483648] = (fi39 * (fi39 >> ((18446744073709551615) & 63)))
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
