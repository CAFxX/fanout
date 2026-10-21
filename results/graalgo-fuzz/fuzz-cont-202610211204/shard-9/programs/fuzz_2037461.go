package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint64; F2 uint64 }
func genFn0() uint32 {
    var v5 uint16 = 255
    var v6 uint16 = 32767
    var v7 int64 = 296
    if (!true) {
        if false {
            v8 := []int{}
            v9 := GenS1{F0: 42, F1: 927, F2: 0}
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 65535, F1: 2, F2: -112}
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := []int{(-(fi11)), (fi11 * fi11), 2147483647}
            if (false && true) {
                v7 = 42
            }
            _ = fi11
            _ = v12
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 uint64, p1 bool, p2 int64) int16 {
    var v13 int16 = -1
    v14 := []int{}
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 float64 = float64(0.0)
        if (!((fi15 ^ -742) >= (fi15 % (((fi15 ^ 820)) | 1)))) {
            var v17 string = ("hello" + ("a" + ""))
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    v18 := GenS1{F0: 807, F1: (^(p0)), F2: 110}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v18
    return (v13 << ((42) & 63))
}

func main() {
    v1 := []int{}
    var v2 bool = (42 > -128)
    v3 := make(map[int]int)
    var v4 int64 = -1
    if (!(797 == -376)) {
        var v19 string = "\n\t"
        switch "\n\t" {
        case "":
            switch 128 {
            case -9223372036854775808:
                switch v19 {
                case "\n\t":
                    v19 = "a"
                case "foo bar":
                    var v20 uint16 = 65535
                    _ = v20
                case "a":
                    var v21 float32 = -2.25
                    _ = v21
                default:
                    v22 := GenS2{F0: 255}
                    _ = v22
                }
            case 590:
                if true {
                    var v23 float64 = 3.14159
                    if (false && (v2 || (443 > 389))) {
                        for fi24 := 0; fi24 < 3; fi24++ {
                            var v25 int32 = 2147483647
                            v26 := v3[fi24]
                            for fi27 := 0; fi27 < 6; fi27++ {
                                v25 = (v25 >> ((65535) & 63))
                                fmt.Println(fi24, v25, v2)
                                v25 = int32(v4)
                                var v28 uint32 = 1
                                _ = fi27
                                _ = v28
                            }
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        var v29 int16 = -1
                        for fi30 := 0; fi30 < 3; fi30++ {
                            switch (fi30 % (((-(fi30))) | 1)) {
                            case -2147483648:
                                if v2 {
                                    var v31 uint32 = 381
                                    v19 = v19
                                    _ = v31
                                }
                            default:
                                v4 -= v4
                            }
                            v32 := GenS3{F0: genFn1(1, true, (v4 * 1)), F1: fi30, F2: 0}
                            _ = fi30
                            _ = v32
                        }
                        _ = v29
                    }
                    if v2 {
                        v3[-2147483648] = 653
                    }
                    _ = v23
                } else {
                    v33 := GenS1{F0: 795, F1: 1, F2: -128}
                    _ = v33
                }
            default:
                var v34 uint8 = 127
                _ = v34
            }
        default:
            _ = len(v3)
        }
        for fi35 := 0; fi35 < 7; fi35++ {
            var v36 bool = (v2 && true)
            fmt.Println(v36)
            var v37 int64 = (v4 ^ v4)
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = v19
    } else {
        delete(v3, -2147483648)
    }
    v4 -= (v4 - -9223372036854775808)
    var v38 uint8 = 255
    var v39 int8 = 48
    var v40 uint16 = 2
    if v2 {
        v41 := GenS4{F0: (v39 > v39), F1: 240, F2: 42}
        _ = v41
    } else {
        var v42 uint16 = uint16(v38)
        _ = v42
    }
    var v43 int64 = (v4 + 189)
    v2 = (0.1 == -2.25)
    var v44 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v44
}
