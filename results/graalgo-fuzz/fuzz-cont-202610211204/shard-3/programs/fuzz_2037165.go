package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 float64 = float64(float32(0.0))
    delete(v2, "\n\t")
    if (1 <= 0) {
        var v4 int8 = 0
        if (true || ("\n\t" != "hello")) {
            if ((((-2.25 / ((-2.25 / v3) + (v3 + -2.25))) != (((0.0 - 0.1) * (3.14159 * 0.0)) / (((0.0 - (v3 - 3.14159)) + float64(1.5)) * (v3 + float64((1.5 / (3.14159 / 1.0))))))) && false) && true) {
                var v5 uint32 = 65535
                var v6 int64 = 1
                _ = v5
                _ = v6
            }
            v7 := GenS1{F0: 886}
            var v8 int64 = -9223372036854775808
            _ = v7
            _ = v8
        }
        _ = v4
    }
    v3 = v3
    v3 += 3.14159
    var v9 int16 = 32767
    var v10 int8 = -1
    fmt.Println(v3)
    var v11 int8 = 127
    for fi12 := 0; fi12 < 14; fi12++ {
        v11 += v11
        if false {
            fmt.Println(v9, v3, v10)
            v9 += v9
            var v13 int64 = 9223372036854775807
            _ = v13
        } else {
            switch -1 {
            case 0:
                v14 := []int{}
                _ = v14
            case -14:
                for fi15 := 0; fi15 < 17; fi15++ {
                    if false {
                        fmt.Println(v11, fi12)
                    } else {
                        var v16 string = "hello"
                        var v17 uint16 = 42
                        if (false && ((v17 ^ 1) < 255)) {
                            var v18 int = (fi12 << ((0) & 63))
                            v10 = 42
                            var v19 int64 = int64(fi12)
                            _ = v18
                            _ = v19
                        } else {
                            fmt.Println(v10)
                            v9 -= v9
                            var v20 uint64 = 18446744073709551615
                            v20 -= 126
                            _ = v20
                        }
                        _ = v16
                        _ = v17
                    }
                    var v21 int32 = -992
                    v21 = 377
                    _ = fi15
                    _ = v21
                }
            default:
                switch ("a" + ("\n\t" + "hello")) {
                case "":
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v9 -= (^(v9))
                        _ = fi22
                    }
                case "\n\t":
                    fmt.Println(v9, fi12)
                }
            }
            switch (fi12 << ((1) & 63)) {
            case 2147483647:
                switch "foo bar" {
                case "a":
                    for fi23 := 0; fi23 < 4; fi23++ {
                        var v24 uint64 = 42
                        if (v24 >= 846) {
                            v11 = -1
                        } else {
                            if (97 >= (v11 % ((v11) | 1))) {
                                v9 += int16(v24)
                            } else {
                                v25 := GenS1{F0: 973}
                                if (1 < 1) {
                                    v11 = v11
                                } else {
                                    var v26 int8 = v11
                                    if true {
                                        v27 := GenS2{F0: 310}
                                        fmt.Println(fi23, v26)
                                        _ = v27
                                    } else {
                                        var v28 int16 = (-(v9))
                                        var v29 bool = (("\n\t" == "a") || false)
                                        _ = v28
                                        _ = v29
                                    }
                                    var v30 int32 = 42
                                    _ = v26
                                    _ = v30
                                }
                                v9 += 288
                                _ = v25
                            }
                        }
                        _ = fi23
                        _ = v24
                    }
                case "\n\t":
                    v31 := []int{-1, (fi12 % ((fi12) | 1))}
                    _ = v31
                case "":
                    v11 -= 8
                default:
                    v10 += v10
                }
            case -2147483648:
                if (!false) {
                    if (42 >= 106) {
                        fmt.Println(fi12)
                        if true {
                            var v32 float32 = float32((-2.25 / v3))
                            _ = v32
                        } else {
                            var v33 string = "\n\t"
                            v2[v33] = fi12
                            delete(v1, fi12)
                            _ = v33
                        }
                    } else {
                        var v34 int16 = -1
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 3; fi35++ {
                        var v36 string = "a"
                        _ = fi35
                        _ = v36
                    }
                    var v37 uint = 2
                    _ = v37
                }
            case 127:
                v38 := []int{(^(fi12)), fi12, (-(fi12))}
                _ = v38
            default:
                switch 0 {
                case 9223372036854775807:
                    v39 := []int{fi12, (fi12 % ((fi12) | 1)), fi12}
                    _ = v39
                }
            }
            v3 = v3
        }
        _ = fi12
    }
    fmt.Println(v9)
    v3 -= float64(1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v9
}
