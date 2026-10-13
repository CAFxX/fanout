package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int8 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint8; F1 uint64; F2 bool }
func genFn0() int64 {
    var v7 int32 = 360
    if true {
        if false {
            v8 := GenS1{F0: 57, F1: (65535 <= 0), F2: 51}
            _ = v8
        }
        fmt.Println(v7)
    }
    v9 := GenS1{F0: v7, F1: true, F2: 42}
    if true {
        v10 := GenS1{F0: (v7 * 186), F1: true, F2: -128}
        var v11 int8 = -1
        _ = v10
        _ = v11
    }
    var v12 uint8 = 42
    _ = v12
    _ = v7
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{}
    v2 := []int{176, 2147483647}
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    var v5 int32 = -1
    var v6 uint16 = 566
    switch 192 {
    case -273:
        fmt.Println(v5, v6)
    case -176:
        v5 = (v5 / ((v5) | 1))
    case 952:
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 float64 = 0.0
            for fi15 := 0; fi15 < 8; fi15++ {
                fmt.Println(fi13)
                fmt.Println(v5)
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    }
    if false {
        fmt.Println(v6, v6)
    } else {
        var v16 uint32 = 4294967295
        for fi17 := 0; fi17 < 16; fi17++ {
            if (v5 <= -1) {
                switch "a" {
                case "foo bar":
                    var v18 uint8 = 128
                    _ = v18
                case "\n\t":
                    fmt.Println(v6, v6)
                case "":
                    for fi19 := 0; fi19 < 6; fi19++ {
                        v5 += 2147483647
                        v16 = v16
                        _ = fi19
                    }
                }
            } else {
                for fi20 := 0; fi20 < 17; fi20++ {
                    v1 = append(v1, 2147483647)
                    _ = fi20
                }
                for fi21 := 0; fi21 < 2; fi21++ {
                    if false {
                        if true {
                            v5 = (v5 * v5)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(fi17)
                        }
                    }
                    _ = fi21
                }
            }
            _ = fi17
        }
        for fi22 := 0; fi22 < 2; fi22++ {
            switch "foo bar" {
            case "foo bar":
                for fi23 := 0; fi23 < 9; fi23++ {
                    if (true && false) {
                        v2 = append(v2, (fi23 ^ (fi22 - (fi22 ^ (fi23 ^ -562)))))
                        if false {
                            var v24 bool = (!false)
                            _ = v24
                        }
                    } else {
                        if true {
                            fmt.Println(v5, fi22)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch "" {
                            case "\n\t":
                                v5 = (v5 - -2147483648)
                            case "foo bar":
                                v25 := GenS2{F0: v6, F1: false}
                                _ = v25
                            }
                        }
                        v5 -= v5
                    }
                    _ = fi23
                }
            case "\n\t":
                var v26 uint16 = v6
                _ = v26
            }
            for fi27 := 0; fi27 < 10; fi27++ {
                v28 := GenS1{F0: v5, F1: (!true), F2: 84}
                v29 := GenS2{F0: (^(v6)), F1: false}
                fmt.Println(v6, fi27)
                var v30 int16 = 973
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            switch "hello" {
            case "a":
                if (true || true) {
                    var v31 int32 = v5
                    v6 -= uint16(v16)
                    _ = v31
                } else {
                    v16 = v16
                    v16 = 2147483648
                }
            case "hello":
                if len(v2) > 0 {
                    v2[0] = (fi22 / ((fi22) | 1))
                }
            }
            _ = fi22
        }
        _ = v16
    }
    _ = len(v3)
    v6 -= v6
    var v32 int16 = -713
    var v33 float64 = float64(0.1)
    v34 := GenS3{F0: 42, F1: 42, F2: (42 > 0)}
    var v35 uint = 4294967295
    v5 = (v5 / ((v5) | 1))
    var v36 int64 = -297
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
