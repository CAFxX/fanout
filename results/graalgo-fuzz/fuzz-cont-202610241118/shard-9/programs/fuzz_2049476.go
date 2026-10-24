package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int = 608
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    v2 = v2
    var v6 uint64 = 144
    v7 := GenS1{F0: 100, F1: 1, F2: 486}
    var v8 uint64 = 1
    fmt.Println(v6, v2)
    var v9 int64 = -1
    fmt.Println(v8, v1)
    var v10 float64 = (0.1 - (1.5 / float64(1e10)))
    if false {
        var v11 int32 = (v2 >> ((1) & 63))
        v12 := []int{(^(v1)), v1}
        var v13 string = "a"
        _ = v11
        _ = v12
        _ = v13
    } else {
        _ = len(v4)
        var v14 string = ("a" + "hello")
        v15 := GenS2{F0: 0}
        _ = v14
        _ = v15
    }
    fmt.Println(v1, v9, v8)
    if (255 > 65535) {
        var v16 uint = 255
        var v17 string = "\n\t"
        if false {
            _ = len(v3)
            switch (v1 % (((v1 ^ int(v6))) | 1)) {
            case 830:
                v6 -= (^(v6))
            case 42:
                for fi18 := 0; fi18 < 17; fi18++ {
                    for fi19 := 0; fi19 < 12; fi19++ {
                        if true {
                            if (!true) {
                                v16 = v16
                                if (false && false) {
                                    v17 = v17
                                }
                            }
                        } else {
                            v1 -= (fi18 ^ -25)
                            switch (fi19 ^ (v1 + fi18)) {
                            case 42:
                                if false {
                                    v20 := []int{32767, (^(fi18))}
                                    _ = v20
                                } else {
                                    v2 = -2147483648
                                    if true {
                                        var v21 int32 = 2147483647
                                        _ = v21
                                    }
                                }
                            case -735:
                                if true {
                                    fmt.Println(v9)
                                    var v22 float64 = 1e10
                                    v1 += v1
                                    _ = v22
                                } else {
                                    v2 += (v2 * 2147483647)
                                    v16 -= 694
                                }
                            default:
                                var v23 uint64 = (v8 / ((2) | 1))
                                var v24 uint32 = 1
                                v25 := GenS2{F0: 65535}
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS2{F0: (v16 * v16)}
                            _ = v26
                        }
                        fmt.Println(v10, v6)
                        fmt.Println(v10, v8, v16)
                        _ = fi19
                    }
                    v1 -= (^(v1))
                    _ = fi18
                }
            default:
                switch v1 {
                case 32768:
                    v27 := GenS2{F0: (^(v16))}
                    _ = v27
                default:
                    fmt.Println(v2)
                }
            }
        }
        _ = v16
        _ = v17
    }
    v8 = (v8 + 4294967295)
    v28 := GenS1{F0: 32, F1: 2, F2: (v9 / ((-1) | 1))}
    v10 -= 3.14159
    v29 := v4[0]
    if (!false) {
        v30 := []int{}
        _ = v30
    } else {
        for fi31 := 0; fi31 < 11; fi31++ {
            if v29 {
                fmt.Println(fi31, v9)
                var v32 bool = (65535 != 1)
                _ = v32
            }
            v6 -= v6
            _ = fi31
        }
    }
    v33 := GenS3{F0: 231}
    v34 := GenS1{F0: -51, F1: 2, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
