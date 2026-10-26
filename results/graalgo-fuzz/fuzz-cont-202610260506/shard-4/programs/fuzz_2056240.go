package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 uint8 = 190
    switch "\n\t" {
    case "":
        for fi4 := 0; fi4 < 9; fi4++ {
            v5 := []int{}
            var v6 uint64 = 510
            if false {
                v7 := v1[(fi4 * (fi4 | fi4))]
                _ = v7
            }
            _ = fi4
            _ = v5
            _ = v6
        }
    case "a":
        v8 := GenS1{F0: int8(v3)}
        _ = v8
    case "\n\t":
        var v9 int64 = 497
        _ = v9
    }
    var v10 int64 = 9223372036854775807
    var v11 int16 = 1000
    fmt.Println(v10, v10, v11)
    v12 := GenS1{F0: 100}
    v2[-971] = 2147483647
    switch 737 {
    case -118:
        fmt.Println(v3, v11, v10)
    default:
        v13 := GenS1{F0: -128}
        _ = v13
    }
    var v14 uint16 = 42
    _ = len(v1)
    v14 = v14
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint16 = 597
        v16 = (^(v14))
        _ = fi15
        _ = v16
    }
    if false {
        v17 := []int{-1, -9223372036854775808}
        for fi18 := 0; fi18 < 12; fi18++ {
            v2[-1] = 647
            _ = fi18
        }
        _ = v17
    } else {
        switch 716 {
        case 865:
            v19 := GenS1{F0: 1}
            _ = v19
        case -527:
            fmt.Println(v14, v3, v10)
        default:
            v20 := GenS2{F0: ("\n\t" >= ""), F1: 2147483647}
            _ = v20
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            if true {
                if false {
                    for fi22 := 0; fi22 < 2; fi22++ {
                        fmt.Println(fi21, v14)
                        v23 := []int{fi22, (-(fi22)), (fi21 + fi22), fi22}
                        v24 := GenS3{F0: (^(v14)), F1: 0}
                        if ((fi22 <= fi21) || false) {
                            var v25 int8 = 127
                            _ = v25
                        } else {
                            fmt.Println(v11, v14)
                            if ((0 < 0) || (2147483647 > 42)) {
                                if false {
                                    fmt.Println(v10, fi22, fi22)
                                    _ = len(v23)
                                }
                                var v26 int8 = 127
                                _ = v26
                            }
                            v10 -= v10
                        }
                        v14 = v14
                        fmt.Println(fi22)
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v14, v14, fi21)
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v10 = (v10 | -275)
                        v28 := []int{(-(fi21)), -257, (fi21 << ((300) & 63)), (fi27 >> ((0) & 63))}
                        _ = fi27
                        _ = v28
                    }
                }
                if false {
                    v29 := GenS2{F0: false, F1: int32(v14)}
                    v30 := []int{(^(fi21)), -2147483648, fi21, (fi21 << ((462) & 63))}
                    _ = v29
                    _ = v30
                } else {
                    v10 = -736
                    var v31 float64 = 1.5
                    v32 := []int{(fi21 >> ((18446744073709551615) & 63)), 1, (fi21 & (fi21 << ((0) & 63)))}
                    _ = v31
                    _ = v32
                }
            }
            var v33 uint32 = 2147483648
            _ = fi21
            _ = v33
        }
        v14 -= 2
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        if true {
            v35 := GenS2{F0: false, F1: 2147483647}
            _ = v35
        } else {
            v11 = (v11 >> ((569) & 63))
        }
        if true {
            v14 = 0
            v36 := []int{fi34, (fi34 << ((255) & 63)), 798}
            _ = v36
        }
        v10 += v10
        _ = fi34
    }
    v3 -= (v3 / ((255) | 1))
    if false {
        if (100 < -1) {
            v37 := GenS2{F0: false, F1: 47}
            v38 := GenS4{F0: 42, F1: int64(v11), F2: -137}
            _ = v37
            _ = v38
        }
        fmt.Println(v11, v11, v11)
        if false {
            fmt.Println(v14)
            fmt.Println(v10, v3, v3)
        } else {
            v11 = v11
            var v39 uint64 = 592
            var v40 int64 = (^(v10))
            _ = v39
            _ = v40
        }
    } else {
        fmt.Println(v11, v14, v14)
        var v41 float32 = 3.14159
        if false {
            fmt.Println(v14)
            var v42 int16 = int16(v10)
            if true {
                var v43 float64 = (0.1 - -2.25)
                _ = v43
            }
            _ = v42
        }
        _ = v41
    }
    var v44 float64 = 1.5
    v3 = v3
    fmt.Println(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v44
}
