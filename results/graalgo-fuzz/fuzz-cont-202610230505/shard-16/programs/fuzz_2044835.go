package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 string; F1 int64; F2 string }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    var v4 uint8 = 127
    var v5 int32 = -2147483648
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 0, F1: 701}
        _ = fi6
        _ = v7
    }
    if ((!(!(!true))) && (!true)) {
        v8 := GenS1{F0: 385, F1: uint(v4)}
        v9 := GenS1{F0: 2, F1: 0}
        _ = v8
        _ = v9
    } else {
        v5 = (v5 * (v5 + v5))
        if (!false) {
            v1 = v1
            v10 := GenS2{F0: 4294967295}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS1{F0: 2, F1: 255}
                delete(v2, 1)
                _ = fi11
                _ = v12
            }
        }
        fmt.Println(v1)
    }
    v13 := GenS3{F0: v1}
    fmt.Println(v5)
    for fi14 := 0; fi14 < 20; fi14++ {
        v5 -= v5
        _ = fi14
    }
    if true {
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := v2[39]
            _ = fi15
            _ = v16
        }
    }
    v17 := v2[65535]
    if (((("a" + (("a" + v1) + v1)) + (v1 + (v1 + v1))) + ((v1 + ((v1 + ((v1 + (v1 + ("\n\t" + (v1 + v1)))) + "foo bar")) + (v1 + v1))) + v1)) != v1) {
        switch -541 {
        case 42:
            for fi18 := 0; fi18 < 15; fi18++ {
                v4 = 0
                _ = fi18
            }
        default:
            v17 = (v17 % ((611) | 1))
        }
        fmt.Println(v17, v5)
        var v19 int8 = 100
        _ = v19
    } else {
        if true {
            var v20 int = -2147483648
            _ = v20
        } else {
            var v21 uint = 84
            _ = v21
        }
        for fi22 := 0; fi22 < 13; fi22++ {
            v5 -= v5
            for fi23 := 0; fi23 < 8; fi23++ {
                var v24 uint64 = 687
                switch fi23 {
                case 4294967295:
                    v1 = (v1 + v1)
                default:
                    v24 += v24
                    if (0 >= 42) {
                        if (false && (false || false)) {
                            v4 = v4
                        } else {
                            var v25 float64 = 3.14159
                            v26 := GenS4{F0: v5, F1: int32(v17)}
                            _ = len(v3)
                            _ = v25
                            _ = v26
                        }
                        if (!true) {
                            if (!true) {
                                var v27 int32 = v5
                                _ = v27
                            }
                            var v28 uint16 = 154
                            _ = v28
                        }
                        fmt.Println(fi23)
                    } else {
                        v17 -= v17
                    }
                    v29 := GenS1{F0: 65535, F1: 42}
                    _ = v29
                }
                _ = fi23
                _ = v24
            }
            for fi30 := 0; fi30 < 18; fi30++ {
                switch fi30 {
                case 1:
                    if (!(((v4 ^ (^(v4))) != v4) || false)) {
                        var v31 uint8 = v4
                        v17 -= (v17 - v17)
                        v31 += v31
                        _ = v31
                    }
                case -747:
                    switch "a" {
                    case "":
                        v32 := []int{fi30, 2147483647, fi30}
                        _ = v32
                    }
                case 9223372036854775807:
                    var v33 uint64 = 2
                    _ = v33
                default:
                    var v34 uint32 = v17
                    _ = v34
                }
                _ = fi30
            }
            _ = fi22
        }
    }
    var v35 int64 = 1
    v36 := v3[-1]
    if true {
        v37 := GenS5{F0: (("foo bar" + (v1 + (v1 + (v1 + v1)))) + "foo bar"), F1: v35, F2: "hello"}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
