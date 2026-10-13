package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS3 struct { F0 bool }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]uint32)
    var v3 int32 = 902
    var v4 int16 = 0
    var v5 int64 = -1
    v1 = v1
    if true {
        v3 = (v3 % (((v3 ^ v3)) | 1))
        v6 := []int{1}
        for fi7 := 0; fi7 < 20; fi7++ {
            v3 -= (v3 * v3)
            if true {
                var v8 int32 = 1
                var v9 uint8 = 97
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        _ = v6
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{fi10, (fi10 | (-(fi10))), 1}
            var v12 uint64 = 2
            var v13 int8 = 1
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    }
    v3 = 0
    v2[-1] = 1
    v14 := GenS1{F0: v5, F1: 1, F2: 255}
    switch "hello" {
    case "foo bar":
        v15 := GenS2{F0: 0, F1: (^(v4)), F2: 719}
        _ = v15
    }
    v4 = -765
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS2{F0: 65535, F1: 995, F2: 255}
        v18 := GenS1{F0: v5, F1: -90, F2: 42}
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v5 -= (v5 * (v5 - v5))
        _ = fi19
    }
    v3 -= -1
    v3 = v3
    fmt.Println(v1, v1)
    if (!false) {
        switch "" {
        case "a":
            switch "" {
            case "":
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 uint32 = 1
                    _ = fi20
                    _ = v21
                }
            default:
                fmt.Println(v5, v1)
            }
        default:
            v4 = (v4 >> ((0) & 63))
        }
        if true {
            if true {
                for fi22 := 0; fi22 < 10; fi22++ {
                    v23 := []int{(fi22 - (fi22 - (fi22 - 868))), (fi22 % (((fi22 - 319)) | 1)), 127}
                    _ = fi22
                    _ = v23
                }
                var v24 uint16 = 157
                _ = v24
            } else {
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 int8 = 42
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v4 = v4
                        var v28 uint = 199
                        v4 += (v4 / (((v4 >> (((v28 % (((v28 & (v28 % (((v28 % ((v28) | 1))) | 1)))) | 1))) & 63))) | 1))
                        _ = fi27
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
                if (v5 == (^(v5))) {
                    if false {
                        v3 += (-(v3))
                        v5 = v5
                        v29 := []int{-1, 1, 127, -1}
                        _ = v29
                    }
                } else {
                    if true {
                        if false {
                            if (2 != 2) {
                                if true {
                                    var v30 uint8 = 127
                                    v4 -= (^(v4))
                                    if ((v5 - -1) >= 9223372036854775807) {
                                        v4 += (v4 % (((v4 >> ((1) & 63))) | 1))
                                        v1 = v1
                                    }
                                    _ = v30
                                }
                                fmt.Println(v5, v4, v4)
                                v31 := []int{-9223372036854775808}
                                _ = v31
                            } else {
                                var v32 uint8 = 42
                                v33 := GenS3{F0: true}
                                _ = v32
                                _ = v33
                            }
                            switch "hello" {
                            case "a":
                                var v34 float64 = v1
                                _ = v34
                            case "":
                                switch "foo bar" {
                                case "":
                                    v35 := []int{}
                                    _ = v35
                                case "hello":
                                    var v36 uint64 = 360
                                    _ = v36
                                }
                            }
                        } else {
                            var v37 int = 42
                            var v38 int64 = 42
                            var v39 uint8 = 128
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    }
                    v4 = -474
                }
            }
            v1 += v1
            for fi40 := 0; fi40 < 11; fi40++ {
                v5 = (v5 | 1)
                _ = fi40
            }
        } else {
            delete(v2, 166)
        }
    } else {
        var v41 string = ("\n\t" + "\n\t")
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
