package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 int8 = 104
    var v2 bool = true
    var v3 int16 = 0
    var v4 int16 = v3
    v5 := make(map[uint32]int)
    var v6 float64 = 1.5
    fmt.Println(v3, v2, v6)
    v7 := GenS1{F0: int64(v4), F1: 0}
    v4 = v3
    fmt.Println(v3)
    for fi8 := 0; fi8 < 17; fi8++ {
        if true {
            v9 := GenS1{F0: 9223372036854775807, F1: 0}
            v10 := GenS2{F0: 42, F1: 42, F2: 539}
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS2{F0: 42, F1: 2147483647, F2: 255}
                var v13 uint32 = 565
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 bool = (42 == fi14)
                _ = fi14
                _ = v15
            }
        }
        fmt.Println(v2, v2, v2)
        v3 -= 569
        _ = fi8
    }
    if v2 {
        v16 := GenS1{F0: 0, F1: 886}
        _ = v16
    }
    if (v3 < v4) {
        v17 := GenS3{F0: "\n\t"}
        _ = v17
    } else {
        var v18 uint = 728
        var v19 bool = false
        var v20 float64 = v6
        _ = v18
        _ = v19
        _ = v20
    }
    v3 = (^(v4))
    var v21 int8 = 0
    v1 = v21
    var v22 int64 = 245
    v23 := []int{int(v22), 575, -270, 127}
    if v2 {
        var v24 float32 = (0.0 * 0.0)
        var v25 float32 = -2.25
        v26 := GenS4{F0: 2147483648, F1: 186}
        _ = v24
        _ = v25
        _ = v26
    } else {
        v21 += (^(v21))
    }
    v23 = append(v23, -9223372036854775808)
    v22 = v22
    for fi27 := 0; fi27 < 19; fi27++ {
        v2 = false
        if false {
            v28 := GenS1{F0: -9223372036854775808, F1: 65535}
            _ = v28
        } else {
            switch fi27 {
            case -684:
                if (!v2) {
                    delete(v5, 2)
                } else {
                    var v29 bool = (0 <= 4294967295)
                    if v29 {
                        var v30 int64 = (v22 * 898)
                        v6 += 1e10
                        for fi31 := 0; fi31 < 3; fi31++ {
                            var v32 int32 = 42
                            v22 = 1
                            v33 := v5[65535]
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        _ = v30
                    } else {
                        if true {
                            switch fi27 {
                            case -9223372036854775808:
                                fmt.Println(v6, v29, v6)
                            }
                            for fi34 := 0; fi34 < 10; fi34++ {
                                fmt.Println(v29, v1)
                                _ = fi34
                            }
                        } else {
                            if v29 {
                                if v29 {
                                    for fi35 := 0; fi35 < 10; fi35++ {
                                        var v36 bool = (65535 != 781)
                                        _ = fi35
                                        _ = v36
                                    }
                                }
                                fmt.Println(v4)
                            }
                        }
                    }
                    _ = v29
                }
            }
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
