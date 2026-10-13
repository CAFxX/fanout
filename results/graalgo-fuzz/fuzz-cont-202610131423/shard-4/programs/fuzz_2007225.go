package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint16) int {
    p0 -= 662
    v6 := []int{}
    switch 127 {
    case -885:
        v7 := GenS1{F0: 2}
        _ = v7
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = v6
    return 32768
}

func main() {
    v1 := []int{127, 0}
    v2 := make(map[int]int)
    var v3 string = ((("a" + "hello") + "hello") + "a")
    var v4 float32 = -2.25
    v5 := []uint32{}
    for fi8 := 0; fi8 < 17; fi8++ {
        if (!false) {
            _ = len(v2)
            var v9 uint32 = 268
            if true {
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 float32 = float32(3.14159)
                    var v12 int32 = -2147483648
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            } else {
                switch fi8 {
                case -2147483648:
                    var v13 int = fi8
                    _ = v13
                case -9223372036854775808:
                    for fi14 := 0; fi14 < 9; fi14++ {
                        if len(v5) > 0 {
                            v5[0] = 4294967295
                        }
                        v15 := GenS2{F0: (fi8 & genFn0(32767))}
                        v16 := GenS2{F0: fi14}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                case 511:
                    v17 := GenS1{F0: 107}
                    _ = v17
                default:
                    if false {
                        v2[(fi8 * fi8)] = (fi8 - -560)
                    } else {
                        var v18 uint32 = v9
                        var v19 uint8 = 128
                        fmt.Println(v19, v3)
                        _ = v18
                        _ = v19
                    }
                }
                if ((fi8 * (^(fi8))) > fi8) {
                    var v20 float32 = ((v4 * v4) * ((((0.0 + 1e10) + v4) + (1.5 + 3.14159)) * 3.14159))
                    _ = v20
                } else {
                    v21 := GenS2{F0: fi8}
                    _ = v21
                }
            }
            _ = v9
        } else {
            var v22 uint64 = 4294967295
            if true {
                var v23 uint16 = 32767
                var v24 float32 = v4
                _ = v23
                _ = v24
            } else {
                v4 = v4
                var v25 uint = 2
                var v26 uint = (v25 - (v25 & (^(v25))))
                _ = v25
                _ = v26
            }
            _ = v22
        }
        var v27 uint8 = 127
        _ = fi8
        _ = v27
    }
    fmt.Println(v4, v4)
    var v28 uint32 = 164
    for fi29 := 0; fi29 < 20; fi29++ {
        v30 := GenS3{F0: 42, F1: 42}
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        switch fi31 {
        case 4294967295:
            v28 = v28
        case -9223372036854775808:
            v28 = 119
        case 179:
            _ = len(v1)
        }
        var v32 uint32 = (v28 | (v28 * v28))
        _ = fi31
        _ = v32
    }
    if ((-587 > 2147483647) || (!true)) {
        _ = len(v2)
        for fi33 := 0; fi33 < 18; fi33++ {
            fmt.Println(fi33)
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 uint16 = 255
                v28 = 1
                fmt.Println(fi33)
                _ = fi34
                _ = v35
            }
            var v36 float64 = (3.14159 + ((3.14159 / 3.14159) / 1.0))
            _ = fi33
            _ = v36
        }
    }
    if len(v5) > 0 {
        v5[0] = v28
    }
    if false {
        var v37 string = (v3 + v3)
        _ = v37
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
