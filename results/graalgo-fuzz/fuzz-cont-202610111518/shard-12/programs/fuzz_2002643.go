package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []uint32{728}
    v2 := make(map[uint32]bool)
    var v3 int16 = 0
    var v4 float32 = 1.5
    v5 := make(map[int]string)
    var v6 float64 = 1.5
    v7 := GenS1{F0: 255, F1: 4294967295}
    v8 := GenS2{F0: 590}
    if (-530 > 9223372036854775807) {
        _ = len(v5)
        switch 0 {
        case 32767:
            v6 -= (v6 - (float64(v4) + (v6 - float64((-2.25 / 1.5)))))
        case 127:
            for fi9 := 0; fi9 < 1; fi9++ {
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 int64 = 0
                    _ = fi10
                    _ = v11
                }
                fmt.Println(v4, v6)
                v2[1] = (18446744073709551615 < 0)
                _ = fi9
            }
        case 636:
            if len(v1) > 0 {
                v1[0] = 2
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v12 int16 = 0
    var v13 uint = 4294967295
    switch 0 {
    case 0:
        fmt.Println(v6, v4)
    }
    var v14 int32 = 1
    fmt.Println(v12, v3, v3)
    if (((68 > 1) && ((-355 < (v3 | 1)) || (((("\n\t" + "\n\t") + "") < (("\n\t" + ("\n\t" + "hello")) + "")) && true))) && false) {
        var v15 int = 32768
        v16 := GenS2{F0: (v13 - (v13 / ((165) | 1)))}
        _ = v15
        _ = v16
    } else {
        if false {
            if true {
                v17 := GenS3{F0: 290, F1: 49, F2: v14}
                _ = v17
            } else {
                v18 := GenS4{F0: v13}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v18
            }
            for fi19 := 0; fi19 < 5; fi19++ {
                if true {
                    for fi20 := 0; fi20 < 11; fi20++ {
                        if (!false) {
                            switch "" {
                            case "foo bar":
                                var v21 uint64 = 2
                                if (v21 == v21) {
                                    v4 -= v4
                                    v1 = append(v1, 4294967295)
                                }
                                _ = v21
                            case "":
                                v6 -= v6
                                var v22 bool = ((!true) || (fi20 == fi20))
                                _ = v22
                            case "hello":
                                if false {
                                    v13 -= 2
                                    var v23 int32 = (v14 >> ((4294967295) & 63))
                                    _ = v23
                                }
                            }
                            switch (fi20 & 9223372036854775807) {
                            case 127:
                                if ((v14 + (-(v14))) >= v14) {
                                    var v24 uint = v13
                                    _ = v24
                                } else {
                                    v12 += (v3 ^ (-(v3)))
                                }
                            case 42:
                                var v25 int = fi20
                                _ = v25
                            }
                        }
                        v26 := GenS1{F0: 218, F1: v13}
                        var v27 int32 = int32(v12)
                        var v28 int = (fi19 & fi20)
                        _ = fi20
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(v6, v14)
                    for fi29 := 0; fi29 < 13; fi29++ {
                        v13 += (v13 + (v13 & 1))
                        switch ("\n\t" + "") {
                        case "":
                            v14 += v14
                        case "hello":
                            v13 += v13
                        case "\n\t":
                            var v30 int32 = v14
                            _ = v30
                        default:
                            var v31 uint8 = uint8(v13)
                            _ = v31
                        }
                        _ = fi29
                    }
                } else {
                    for fi32 := 0; fi32 < 15; fi32++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi32
                    }
                    var v33 uint32 = 2147483648
                    v3 = (v12 % (((v3 >> (((v13 + v13)) & 63))) | 1))
                    _ = v33
                }
                _ = fi19
            }
            v3 = v12
        }
    }
    fmt.Println(v12, v6, v12)
    v34 := GenS2{F0: v13}
    v35 := GenS5{F0: 42}
    switch 128 {
    case 805:
        fmt.Println(v12, v12)
    default:
        fmt.Println(v13, v4)
    }
    _ = len(v5)
    switch "hello" {
    case "\n\t":
        if false {
            _ = len(v2)
            v6 = v6
        } else {
            var v36 bool = false
            v37 := GenS4{F0: v13}
            _ = v36
            _ = v37
        }
    default:
        var v38 int8 = int8(v13)
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
