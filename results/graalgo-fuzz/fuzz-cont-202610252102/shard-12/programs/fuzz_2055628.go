package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint16; F2 bool }

func main() {
    var v1 int8 = -99
    var v2 int8 = (v1 << ((0) & 63))
    var v3 uint8 = 42
    var v4 uint32 = 0
    switch -134 {
    case 128:
        v5 := GenS1{F0: false, F1: (v4 ^ (^(v4)))}
        _ = v5
    case -407:
        if (-873 == 0) {
            v6 := GenS1{F0: false, F1: v4}
            if (!(1000 == 0)) {
                v7 := GenS1{F0: false, F1: (v4 % ((v4) | 1))}
                _ = v7
            } else {
                v8 := GenS1{F0: ((v4 ^ 1) != (v4 ^ v4)), F1: v4}
                if (!true) {
                    for fi9 := 0; fi9 < 15; fi9++ {
                        switch "foo bar" {
                        case "":
                            switch (fi9 & fi9) {
                            case -694:
                                var v10 uint16 = 65535
                                _ = v10
                            }
                        case "hello":
                            switch "foo bar" {
                            case "foo bar":
                                var v11 int = (fi9 + (fi9 / ((fi9) | 1)))
                                _ = v11
                            case "":
                                v3 = v3
                            case "a":
                                v4 -= (v4 / ((2147483648) | 1))
                            }
                        case "\n\t":
                            var v12 uint16 = 255
                            _ = v12
                        default:
                            if false {
                                var v13 bool = (32767 != -1)
                                _ = v13
                            } else {
                                fmt.Println(v3, fi9)
                                v14 := GenS2{F0: 831}
                                _ = v14
                            }
                        }
                        _ = fi9
                    }
                    v4 = (v4 ^ (v4 | (^(v4))))
                }
                _ = v8
            }
            v15 := []int{-901}
            _ = v15
            _ = v6
        } else {
            v16 := GenS3{F0: 1}
            for fi17 := 0; fi17 < 9; fi17++ {
                v18 := []int{fi17}
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    case -9223372036854775808:
        v1 = int8(v3)
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 float32 = -2.25
        v21 := []int{(fi19 / ((fi19) | 1)), fi19}
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := []int{}
    var v23 int32 = -137
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS4{F0: 0}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v4 = (v4 >> ((669) & 63))
        v22 = append(v22, 1)
        _ = fi26
    }
    var v27 uint32 = 2147483648
    v22 = append(v22, 128)
    var v28 uint64 = 4294967295
    var v29 int16 = 32767
    v27 = v27
    for fi30 := 0; fi30 < 18; fi30++ {
        if true {
            for fi31 := 0; fi31 < 12; fi31++ {
                if true {
                    v23 -= (v23 / ((1) | 1))
                } else {
                    if true {
                        var v32 string = ("" + ("hello" + ("" + ("" + ("\n\t" + "\n\t")))))
                        v33 := GenS5{F0: 2147483647, F1: 2, F2: (1 >= (v4 | v4))}
                        if len(v22) > 0 {
                            v22[0] = -1
                        }
                        _ = v32
                        _ = v33
                    }
                }
                fmt.Println(fi31, v1, v1)
                _ = fi31
            }
            var v34 uint64 = (v28 - 0)
            _ = v34
        } else {
            for fi35 := 0; fi35 < 7; fi35++ {
                v22 = append(v22, (fi30 | (fi30 ^ -102)))
                _ = fi35
            }
            _ = len(v22)
        }
        var v36 bool = false
        _ = fi30
        _ = v36
    }
    v1 += v1
    v23 = (^(v23))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
