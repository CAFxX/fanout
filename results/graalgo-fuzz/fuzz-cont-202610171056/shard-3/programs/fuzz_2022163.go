package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 int, p1 int16) uint16 {
    for fi4 := 0; fi4 < 4; fi4++ {
        fmt.Println(p0)
        _ = fi4
    }
    switch ("\n\t" + ("" + "")) {
    case "":
        v5 := []int{p0}
        _ = v5
    case "\n\t":
        p0 -= 583
    default:
        p0 = p0
    }
    var v6 uint64 = 144
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float32 = float32(1e10)
        if false {
            delete(v3, fi7)
            var v9 uint16 = 2
            _ = v9
        } else {
            var v10 float64 = (3.14159 / ((float64(v8) - 1e10) * 1e10))
            v10 = ((1e10 * float64(((v8 * float32(3.14159)) - -2.25))) / v10)
            _ = v10
        }
        switch "hello" {
        case "\n\t":
            v8 = (1.5 / v8)
        case "hello":
            var v11 float32 = ((v8 * v8) * 1.5)
            _ = v11
        default:
            v12 := []int{(fi7 & 32767), 255, (fi7 % ((1) | 1))}
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    var v13 uint = 4294967295
    v14 := GenS1{F0: 268, F1: -1}
    if true {
        for fi15 := 0; fi15 < 3; fi15++ {
            if false {
                fmt.Println(fi15, fi15)
            }
            switch "a" {
            case "a":
                v16 := v3[(fi15 * (fi15 | (^(fi15))))]
                _ = v16
            }
            _ = fi15
        }
        var v17 int64 = int64(v13)
        _ = v17
    } else {
        v13 += 1
        v13 -= (^(v13))
    }
    v13 -= (^(v13))
    if false {
        switch "hello" {
        case "\n\t":
            switch -558 {
            case 127:
                for fi18 := 0; fi18 < 5; fi18++ {
                    if true {
                        var v19 float64 = (((0.1 / 1.5) / ((0.1 / (1e10 * 0.1)) - 1.5)) * 1.5)
                        for fi20 := 0; fi20 < 7; fi20++ {
                            switch (fi18 + (-(fi20))) {
                            case 1:
                                if true {
                                    var v21 int16 = 1
                                    v13 += (v13 % (((v13 & 838)) | 1))
                                    v21 = (v21 * (v21 - -1))
                                    _ = v21
                                }
                            case 2147483647:
                                var v22 int = (fi20 - -1)
                                _ = v22
                            case 4294967295:
                                var v23 uint64 = 2
                                _ = v23
                            }
                            v24 := GenS2{F0: 494, F1: (^(v13))}
                            _ = fi20
                            _ = v24
                        }
                        _ = v19
                    } else {
                        var v25 uint = v13
                        var v26 int32 = -123
                        _ = v25
                        _ = v26
                    }
                    v27 := GenS2{F0: -2147483648, F1: (v13 / (((v13 & v13)) | 1))}
                    switch (fi18 + 128) {
                    case -2147483648:
                        v28 := []int{965, (fi18 >> ((850) & 63)), fi18, fi18}
                        _ = v28
                    case -1:
                        var v29 uint8 = 127
                        _ = v29
                    default:
                        fmt.Println(v13, fi18)
                    }
                    _ = fi18
                    _ = v27
                }
            default:
                var v30 uint32 = 790
                _ = v30
            }
        case "":
            var v31 uint64 = 42
            _ = v31
        }
    }
    var v32 int32 = -2147483648
    v33 := GenS1{F0: (v13 * (v13 & v13)), F1: -128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
