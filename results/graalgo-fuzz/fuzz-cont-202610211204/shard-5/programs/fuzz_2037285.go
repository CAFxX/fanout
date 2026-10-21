package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v1 += v1
    for fi4 := 0; fi4 < 3; fi4++ {
        if (!true) {
            _ = len(v2)
            var v5 int16 = -740
            if false {
                var v6 uint64 = 0
                _ = v6
            } else {
                switch "foo bar" {
                case "":
                    if false {
                        if true {
                            var v7 uint = 4294967295
                            switch fi4 {
                            case 485:
                                switch fi4 {
                                case 2147483647:
                                    var v8 int16 = v5
                                    _ = v8
                                }
                            case 42:
                                for fi9 := 0; fi9 < 9; fi9++ {
                                    fmt.Println(fi4, v7)
                                    _ = fi9
                                }
                            }
                            _ = v7
                        }
                        v10 := GenS1{F0: (v1 + -1), F1: 687}
                        _ = v10
                    } else {
                        v5 = -1
                        var v11 uint16 = 1
                        _ = v11
                    }
                default:
                    v1 += (v1 | v1)
                }
                v12 := GenS1{F0: (^(v1)), F1: 859}
                _ = v12
            }
            _ = v5
        }
        _ = fi4
    }
    if true {
        v13 := []int{int(v1), 542, 761}
        switch "foo bar" {
        case "a":
            v1 -= -378
        default:
            if ((true && false) && false) {
                var v14 int32 = v1
                _ = v14
            }
        }
        _ = v13
    }
    v15 := GenS1{F0: -206, F1: 42}
    _ = len(v3)
    v1 -= (v1 >> ((255) & 63))
    v16 := []int{307}
    v1 += (v1 - v1)
    v17 := GenS1{F0: 0, F1: 42}
    if true {
        v1 = (v1 % ((v1) | 1))
        if false {
            var v18 float64 = 0.0
            _ = v18
        } else {
            if false {
                var v19 int32 = (v1 + v1)
                var v20 int8 = 4
                _ = v19
                _ = v20
            }
            var v21 bool = (-1 > 1)
            v1 = -101
            _ = v21
        }
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    }
    var v22 uint32 = 65535
    if false {
        if false {
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 int = (fi23 * (fi23 + (fi23 % ((100) | 1))))
                v24 = (fi23 / (((v24 << ((1) & 63))) | 1))
                var v25 int8 = 0
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 uint = 65535
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
}
