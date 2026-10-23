package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 string }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 string, p1 uint) uint {
    p1 -= p1
    fmt.Println(p0, p1)
    var v7 int = 763
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            p1 -= p1
            v10 := []int{(fi8 - (fi8 % ((fi8) | 1))), fi9, 0, fi8}
            _ = fi9
            _ = v10
        }
        var v11 int = (v7 >> (((p1 >> (((^(p1))) & 63))) & 63))
        _ = fi8
        _ = v11
    }
    v7 += (v7 << (((p1 << ((p1) & 63))) & 63))
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 float32 = 0.0
    var v2 int8 = -1
    var v3 uint8 = 127
    var v4 int = 4294967295
    v5 := make(map[int]int)
    v6 := make(map[uint32]bool)
    if (!(!(!((v2 % (((-(v2))) | 1)) < (v2 % ((v2) | 1)))))) {
        v12 := GenS1{F0: 42, F1: 65535, F2: "\n\t"}
        _ = v12
    }
    v4 += -578
    var v13 uint16 = 42
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v4, v1, v1)
        v2 -= v2
        _ = fi14
    }
    var v15 uint16 = v13
    fmt.Println(v13, v1)
    var v16 string = "\n\t"
    if (955 <= genFn0(v16, 1)) {
        v1 = v1
        if true {
            fmt.Println(v4, v16)
            if (false && true) {
                v16 = v16
            } else {
                v17 := GenS1{F0: 220, F1: 768, F2: v16}
                for fi18 := 0; fi18 < 16; fi18++ {
                    if false {
                        var v19 uint32 = 898
                        var v20 string = (v16 + v16)
                        _ = v19
                        _ = v20
                    } else {
                        if (v1 < (v1 / v1)) {
                            fmt.Println(v1)
                            var v21 int8 = (v2 + -59)
                            _ = v21
                        } else {
                            var v22 int16 = 42
                            for fi23 := 0; fi23 < 16; fi23++ {
                                v24 := GenS1{F0: 1, F1: 302, F2: (v16 + v16)}
                                _ = fi23
                                _ = v24
                            }
                            _ = v22
                        }
                    }
                    v25 := []int{v4, (v4 << ((2) & 63))}
                    _ = fi18
                    _ = v25
                }
                _ = v17
            }
            var v26 bool = (!false)
            _ = v26
        }
    }
    var v27 string = (v16 + "a")
    switch v4 {
    case 1:
        for fi28 := 0; fi28 < 19; fi28++ {
            switch v4 {
            case 128:
                var v29 bool = (!false)
                _ = v29
            }
            switch v4 {
            case -103:
                v4 += -1
            case -69:
                var v30 int16 = 0
                _ = v30
            case 2147483647:
                if true {
                    switch "a" {
                    case "a":
                        v13 -= 580
                    case "\n\t":
                        for fi31 := 0; fi31 < 5; fi31++ {
                            v1 = float32(1e10)
                            _ = fi31
                        }
                    }
                    v32 := GenS2{F0: v27, F1: 0}
                    _ = v32
                }
            default:
                var v33 uint = 495
                _ = v33
            }
            v34 := v6[2147483647]
            _ = fi28
            _ = v34
        }
    }
    fmt.Println(v4)
    v35 := GenS1{F0: 42, F1: 255, F2: v16}
    var v36 uint32 = 4294967295
    if (361 >= (^(v13))) {
        v2 = 100
        v37 := GenS1{F0: 42, F1: uint(v36), F2: "a"}
        var v38 float32 = (v1 * ((v1 - 0.0) * v1))
        _ = v37
        _ = v38
    }
    fmt.Println(v16)
    fmt.Println(v36, v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
