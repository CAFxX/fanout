package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 string; F1 int64; F2 int }
func genFn0(p0 int16) uint64 {
    if true {
        v4 := GenS1{F0: 871}
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := GenS1{F0: 2}
            if true {
                switch fi5 {
                case 0:
                    var v7 bool = (!false)
                    _ = v7
                case -2147483648:
                    p0 += 0
                default:
                    v8 := GenS2{F0: (fi5 & (-(fi5)))}
                    _ = v8
                }
            } else {
                p0 -= (p0 - p0)
                var v9 uint16 = 32767
                _ = v9
            }
            switch "" {
            case "a":
                var v10 uint = 42
                _ = v10
            default:
                fmt.Println(fi5)
            }
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    v11 := []int{9223372036854775807, 32768, 4294967295}
    var v12 int16 = p0
    if len(v11) > 0 {
        v11[0] = 2147483647
    }
    _ = p0
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 float64 = 1e10
    var v2 int32 = -1
    v3 := make(map[int]int)
    var v13 int = 9223372036854775807
    var v14 uint = 90
    if (255 == 2) {
        if false {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint16 = 42
                v17 := []int{}
                var v18 int8 = int8(v2)
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            v1 += (((v1 / float64(1e10)) / v1) / ((0.0 * (float64(3.14159) + ((v1 * -2.25) / v1))) / (v1 + v1)))
            v19 := GenS1{F0: 65535}
            _ = v19
        }
        v20 := GenS3{F0: 9223372036854775807, F1: 2, F2: 42}
        _ = v20
    }
    var v21 uint8 = 2
    v22 := GenS3{F0: 91, F1: 189, F2: 1000}
    switch "foo bar" {
    case "hello":
        v21 += v21
    case "\n\t":
        fmt.Println(v1, v2)
    case "foo bar":
        switch v13 {
        case 128:
            var v23 int16 = -32768
            _ = v23
        case 383:
            if false {
                v24 := GenS4{F0: 2, F1: int(v21)}
                v13 = (v13 % (((v13 >> ((v14) & 63))) | 1))
                _ = v24
            } else {
                for fi25 := 0; fi25 < 16; fi25++ {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        if false {
                            switch fi25 {
                            case 4294967295:
                                var v27 uint32 = 65535
                                _ = v27
                            case -9223372036854775808:
                                v28 := GenS2{F0: v13}
                                _ = v28
                            case -1:
                                var v29 uint32 = 2147483647
                                _ = v29
                            default:
                                var v30 int16 = 32767
                                _ = v30
                            }
                            switch (fi25 * fi25) {
                            case 4294967295:
                                switch "" {
                                case "foo bar":
                                    fmt.Println(v1)
                                case "":
                                    v2 = v2
                                case "a":
                                    var v31 uint8 = v21
                                    _ = v31
                                }
                            case 520:
                                var v32 int8 = -51
                                _ = v32
                            }
                        } else {
                            v21 = v21
                            v13 = 2147483647
                            var v33 int = 639
                            var v34 uint32 = 805
                            v35 := GenS2{F0: (^(v33))}
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = fi26
                    }
                    _ = len(v3)
                    v13 += 32767
                    _ = fi25
                }
                if ((true || (-606 > -1)) || false) {
                    var v36 float32 = ((0.1 * (0.1 + (1.5 + -2.25))) * 0.1)
                    v37 := []int{4294967295, 988, (v13 << ((777) & 63))}
                    for fi38 := 0; fi38 < 7; fi38++ {
                        v14 -= 0
                        v39 := GenS2{F0: (^(fi38))}
                        v40 := GenS5{F0: "foo bar", F1: -9223372036854775808, F2: v13}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = v36
                    _ = v37
                } else {
                    if true {
                        v2 = 1
                        v2 += v2
                    } else {
                        fmt.Println(v21, v1)
                    }
                }
                v13 += 9223372036854775807
            }
        case -9223372036854775808:
            v41 := GenS4{F0: 2147483647, F1: v13}
            _ = v41
        }
    default:
        var v42 int8 = 67
        _ = v42
    }
    for fi43 := 0; fi43 < 4; fi43++ {
        v44 := GenS5{F0: "", F1: 1, F2: 0}
        _ = fi43
        _ = v44
    }
    v45 := v3[-708]
    for fi46 := 0; fi46 < 6; fi46++ {
        if (true || true) {
            var v47 int32 = (v2 / ((v2) | 1))
            _ = v47
        }
        _ = fi46
    }
    v2 += (-(v2))
    _ = len(v3)
    var v48 float64 = (float64((3.14159 - (1.5 - (1e10 - (1e10 - float32(v1)))))) + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v45
    _ = v48
}
