package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 uint, p2 bool) int32 {
    v7 := []int{4294967295, 9223372036854775807, 122, -35}
    if len(v7) > 0 {
        v7[0] = 128
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2147483647
}

func main() {
    var v1 int16 = -32768
    var v2 int32 = 829
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    var v5 int = int(v2)
    var v6 int8 = 127
    v5 = v5
    if false {
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(v5)
                v1 = (^(v1))
                switch fi9 {
                case 4294967295:
                    var v10 uint32 = 4294967295
                    _ = v10
                }
                v1 = (v1 % ((v1) | 1))
                _ = fi9
            }
            if true {
                for fi11 := 0; fi11 < 4; fi11++ {
                    var v12 uint8 = 77
                    _ = fi11
                    _ = v12
                }
                fmt.Println(v5, v2)
                v13 := []int{(fi8 % ((fi8) | 1)), 2147483647, -368}
                _ = v13
            } else {
                if (-1 <= 42) {
                    v14 := []int{}
                    if false {
                        if true {
                            var v15 int8 = 100
                            if false {
                                for fi16 := 0; fi16 < 14; fi16++ {
                                    fmt.Println(v5)
                                    v15 += v15
                                    v2 += (-(v2))
                                    _ = fi16
                                }
                            }
                            _ = v15
                        }
                        v1 = v1
                    }
                    _ = v14
                } else {
                    v17 := []int{fi8}
                    _ = v17
                }
                fmt.Println(v6, v1)
                if (!true) {
                    v1 += (v1 / ((v1) | 1))
                    var v18 bool = true
                    _ = v18
                } else {
                    v3["a"] = 757
                }
            }
            _ = fi8
        }
        v19 := GenS1{F0: 416, F1: 382}
        fmt.Println(v6, v2)
        _ = v19
    }
    v2 -= (v2 % ((2147483647) | 1))
    v6 = -128
    var v20 uint16 = 32767
    v20 = 2
    if false {
        if false {
            var v21 int = -1
            _ = v21
        } else {
            for fi22 := 0; fi22 < 18; fi22++ {
                v23 := GenS2{F0: (^(v5))}
                var v24 int64 = 0
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    }
    v20 = v20
    switch ("\n\t" + "hello") {
    case "\n\t":
        v20 -= (v20 << ((629) & 63))
    default:
        var v25 float64 = -2.25
        _ = v25
    }
    if false {
        if true {
            v6 = v6
        }
    }
    if false {
        for fi26 := 0; fi26 < 9; fi26++ {
            v27 := GenS2{F0: 32768}
            _ = fi26
            _ = v27
        }
        v28 := GenS2{F0: v5}
        var v29 uint16 = 255
        _ = v28
        _ = v29
    } else {
        if (true || false) {
            v30 := GenS1{F0: 255, F1: 255}
            _ = v30
        } else {
            v31 := []int{v5, (-(v5))}
            if true {
                var v32 float32 = (0.0 * 1.5)
                _ = v32
            }
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
