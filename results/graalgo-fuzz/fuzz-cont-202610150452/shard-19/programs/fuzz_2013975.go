package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint64, p1 uint64, p2 int) int32 {
    var v7 uint8 = 2
    switch "a" {
    case "a":
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int16 = -1
            fmt.Println(p2, p2, v7)
            var v10 uint = 42
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        v7 += 126
    }
    v11 := GenS1{F0: 0, F1: -85, F2: -1}
    v12 := GenS2{F0: 107, F1: 0, F2: 2}
    if (!(1 <= 4294967295)) {
        var v13 uint64 = p1
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v7
    return 1
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]bool)
    var v3 float64 = v1
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    v6 := make(map[uint32]string)
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    var v16 int32 = -546
    var v17 uint64 = 1
    if (false || false) {
        fmt.Println(v17)
        for fi18 := 0; fi18 < 14; fi18++ {
            v19 := []int{}
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: -1, F1: 100, F2: 32768}
        _ = v20
    } else {
        var v21 uint16 = 1
        if (42 == 2) {
            for fi22 := 0; fi22 < 4; fi22++ {
                if true {
                    if (!(42 != 1)) {
                        v21 += 2
                        v23 := GenS2{F0: 1, F1: (v21 + 255), F2: 670}
                        _ = v23
                    }
                } else {
                    var v24 int8 = 127
                    _ = v24
                }
                _ = fi22
            }
        } else {
            var v25 int32 = genFn0(4294967295, v17, 32767)
            var v26 uint32 = 0
            if true {
                for fi27 := 0; fi27 < 3; fi27++ {
                    var v28 int32 = (v16 | v25)
                    var v29 int = (fi27 | fi27)
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                for fi30 := 0; fi30 < 18; fi30++ {
                    v31 := []int{(fi30 + fi30), fi30, (fi30 & fi30)}
                    if false {
                        for fi32 := 0; fi32 < 9; fi32++ {
                            if ((!true) || (!false)) {
                                if true {
                                    var v33 uint32 = 4294967295
                                    v1 += v1
                                    _ = v33
                                } else {
                                    delete(v4, "hello")
                                }
                            }
                            _ = fi32
                        }
                        var v34 int = (-(fi30))
                        _ = v34
                    }
                    var v35 uint16 = (v21 << ((770) & 63))
                    _ = fi30
                    _ = v31
                    _ = v35
                }
            } else {
                fmt.Println(v1, v26, v25)
                v25 = v16
                v3 = v3
            }
            _ = v25
            _ = v26
        }
        _ = v21
    }
    if (42 <= 32767) {
        v36 := GenS1{F0: -63, F1: -98, F2: 105}
        _ = v36
    }
    v37 := GenS3{F0: -2147483648, F1: v16, F2: 127}
    v38 := GenS4{F0: 4294967295}
    fmt.Println(v17)
    fmt.Println(v16, v16, v3)
    var v39 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
