package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint) bool {
    v7 := []int{}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = v7
    return true
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "hello"
    v3 := []int{}
    v4 := []string{("\n\t" + "hello"), (("hello" + v2) + "hello")}
    var v5 int = 32768
    v6 := make(map[uint32]string)
    var v8 uint16 = 65535
    _ = len(v6)
    v9 := GenS1{F0: v5}
    v10 := GenS1{F0: v5}
    fmt.Println(v8, v8)
    if (-1 > 42) {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 float32 = (float32(float64(1e10)) - 1e10)
            _ = fi11
            _ = v12
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 int32 = 42
    for fi14 := 0; fi14 < 20; fi14++ {
        fmt.Println(v8, v5)
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS1{F0: 306}
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v2 = v2
    fmt.Println(v8, v13, v13)
    v17 := GenS2{F0: 42, F1: (v8 % ((v8) | 1))}
    v8 += v8
    switch -117 {
    case 1:
        switch v5 {
        case 128:
            for fi18 := 0; fi18 < 7; fi18++ {
                var v19 uint32 = 2
                _ = fi18
                _ = v19
            }
        }
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        v1[(-(v5))] = false
        for fi21 := 0; fi21 < 12; fi21++ {
            switch (-(v5)) {
            case 9223372036854775807:
                if ((float32(3.14159) - (3.14159 + (0.0 + (-2.25 * float32((0.1 + (0.0 * 0.1))))))) > (-2.25 - (0.1 / 3.14159))) {
                    v5 = 0
                    if true {
                        var v22 int = (fi20 >> ((963) & 63))
                        _ = v22
                    } else {
                        var v23 float64 = -2.25
                        var v24 int = fi20
                        _ = v23
                        _ = v24
                    }
                    v5 -= 305
                } else {
                    v13 -= (v13 - -636)
                    v8 = (^(v8))
                }
            case 32768:
                if true {
                    v13 = 42
                    _ = len(v1)
                }
            case 42:
                fmt.Println(v8, v5, v8)
            default:
                var v25 int = fi20
                _ = v25
            }
            v26 := GenS2{F0: 457, F1: 42}
            _ = len(v3)
            var v27 int8 = 4
            v13 = v13
            var v28 int8 = (^(v27))
            v8 += v8
            var v29 int32 = (^(v13))
            _ = fi21
            _ = v26
            _ = v27
            _ = v28
            _ = v29
        }
        if (1 >= 1) {
            v30 := GenS1{F0: (^(fi20))}
            _ = v30
        }
        _ = fi20
    }
    v8 -= (v8 << ((353) & 63))
    if true {
        v13 = v13
        if false {
            for fi31 := 0; fi31 < 18; fi31++ {
                fmt.Println(v2, v13, v13)
                if (0 >= 2147483647) {
                    v32 := GenS1{F0: v5}
                    _ = v32
                }
                v33 := GenS3{F0: 0}
                _ = fi31
                _ = v33
            }
            for fi34 := 0; fi34 < 15; fi34++ {
                switch fi34 {
                case 161:
                    if false {
                        v35 := GenS2{F0: 481, F1: 65535}
                        switch fi34 {
                        case 4294967295:
                            var v36 int8 = -1
                            _ = v36
                        case 9223372036854775807:
                            var v37 uint16 = v8
                            _ = v37
                        default:
                            for fi38 := 0; fi38 < 8; fi38++ {
                                v39 := v1[fi38]
                                v13 += v13
                                _ = fi38
                                _ = v39
                            }
                        }
                        fmt.Println(v2, v2, v13)
                        _ = v35
                    }
                case -2147483648:
                    for fi40 := 0; fi40 < 15; fi40++ {
                        if false {
                            v2 = v2
                            var v41 int8 = 81
                            _ = v41
                        }
                        var v42 string = v2
                        _ = fi40
                        _ = v42
                    }
                }
                v43 := v1[(^(fi34))]
                _ = fi34
                _ = v43
            }
        }
        switch v2 {
        case "a":
            fmt.Println(v2, v5)
        case "\n\t":
            var v44 uint8 = 33
            _ = v44
        default:
            v45 := v6[0]
            _ = v45
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
