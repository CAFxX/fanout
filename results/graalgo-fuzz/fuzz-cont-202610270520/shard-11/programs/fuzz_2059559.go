package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int16 }
func genFn0(p0 int8, p1 uint) int {
    p0 = p0
    p0 += -1
    _ = p0
    _ = p1
    return -461
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    var v3 float32 = (0.0 / ((-2.25 / (0.1 * 3.14159)) + 3.14159))
    v4 := make(map[uint32]string)
    v5 := []int{42}
    v6 := make(map[int]string)
    if true {
        v7 := GenS1{F0: 119, F1: -83, F2: 4294967295}
        v8 := GenS2{F0: genFn0(63, 0)}
        for fi9 := 0; fi9 < 9; fi9++ {
            switch genFn0(-75, 0) {
            case 127:
                if false {
                    var v10 int8 = -1
                    _ = len(v6)
                    if (2147483647 > 42) {
                        if false {
                            switch -306 {
                            case 127:
                                v11 := GenS3{F0: 994, F1: 1000}
                                _ = v11
                            }
                            if false {
                                if (0 != -753) {
                                    for fi12 := 0; fi12 < 4; fi12++ {
                                        fmt.Println(fi12, fi9, v10)
                                        switch fi9 {
                                        case 9223372036854775807:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case 32767:
                                            switch (fi12 << ((0) & 63)) {
                                            case 635:
                                                delete(v1, ("" + "\n\t"))
                                            default:
                                                v3 += (v3 / v3)
                                            }
                                        case 1:
                                            switch fi12 {
                                            case 32768:
                                                var v13 int64 = -1
                                                _ = v13
                                            case 0:
                                                _ = len(v6)
                                            case 1:
                                                if (false || ((-2147483648 < 42) || (!((v10 > -128) && ((int64(fi9) <= -560) || false))))) {
                                                    v14 := GenS4{F0: 911, F1: 22}
                                                    var v15 bool = (!((false && false) || false))
                                                    _ = v14
                                                    _ = v15
                                                } else {
                                                    var v16 uint16 = 180
                                                    switch ("a" + "a") {
                                                    case "":
                                                        var v17 float64 = 3.14159
                                                        _ = v17
                                                    }
                                                    var v18 int64 = -1
                                                    _ = v16
                                                    _ = v18
                                                }
                                                v10 += v10
                                                v3 = v3
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                        }
                                        var v19 uint32 = 317
                                        _ = fi12
                                        _ = v19
                                    }
                                    v10 -= 42
                                    switch (fi9 ^ (-(fi9))) {
                                    case 716:
                                        var v20 bool = (-1 > 63)
                                        _ = v20
                                    default:
                                        if (164 == 2147483647) {
                                            _ = len(v5)
                                            v3 -= v3
                                        }
                                    }
                                }
                                var v21 float32 = v3
                                v21 = v3
                                _ = v21
                            } else {
                                v22 := GenS2{F0: -1}
                                v3 = v3
                                _ = v22
                            }
                        }
                        var v23 float64 = (1.5 + (-2.25 / (1e10 / 1.0)))
                        _ = v23
                    }
                    _ = v10
                } else {
                    fmt.Println(fi9)
                }
            default:
                v3 -= v3
            }
            _ = fi9
        }
        _ = v7
        _ = v8
    } else {
        fmt.Println(v3)
    }
    v5 = append(v5, 2147483647)
    v24 := GenS5{F0: 2, F1: 42, F2: 32767}
    if true {
        for fi25 := 0; fi25 < 16; fi25++ {
            v4[646] = ("\n\t" + (("a" + ("" + "\n\t")) + "foo bar"))
            var v26 int16 = -32768
            fmt.Println(fi25)
            _ = fi25
            _ = v26
        }
        var v27 uint64 = 42
        _ = v27
    }
    v3 -= v3
    for fi28 := 0; fi28 < 20; fi28++ {
        v3 = ((v3 * v3) * v3)
        v29 := GenS5{F0: 316, F1: 1, F2: -1}
        _ = fi28
        _ = v29
    }
    fmt.Println(v3)
    var v30 float64 = (1e10 + (float64(v3) / 1.0))
    v30 -= v30
    var v31 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
