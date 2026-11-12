package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint16; F1 string }

func main() {
    var v1 float64 = (-2.25 * 3.14159)
    var v2 uint8 = 127
    v3 := make(map[int]string)
    var v4 uint = 0
    var v5 bool = true
    var v6 string = "hello"
    _ = len(v3)
    if (!((194 == v2) || ((false || ((1 <= -1) && true)) || (true && (((v1 + ((v1 - v1) - (((((v1 - v1) + ((3.14159 - v1) - (v1 * float64(3.14159)))) / ((v1 * ((float64((-2.25 * float32((-2.25 - v1)))) + float64(0.0)) / v1)) * v1)) + v1) / (v1 / 1.0)))) + ((3.14159 * v1) + v1)) == 0.1))))) {
        if v5 {
            for fi7 := 0; fi7 < 20; fi7++ {
                var v8 float64 = 0.0
                var v9 bool = false
                _ = fi7
                _ = v8
                _ = v9
            }
            v1 -= (-2.25 + 0.0)
            v1 = 1e10
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            switch (fi10 - fi10) {
            case 128:
                if (((v1 * (0.0 + ((0.0 + -2.25) + 1.5))) * v1) > ((v1 + v1) + v1)) {
                    v11 := GenS1{F0: v4, F1: v4}
                    v12 := []int{-259, -226}
                    if v5 {
                        fmt.Println(v1, v1, v4)
                    }
                    _ = v11
                    _ = v12
                } else {
                    for fi13 := 0; fi13 < 11; fi13++ {
                        if (!(!v5)) {
                            if (v5 && v5) {
                                v14 := GenS2{F0: (v4 >> (((^(v4))) & 63))}
                                v6 = v6
                                switch fi13 {
                                case 733:
                                    v2 = v2
                                case 32767:
                                    var v15 float32 = (-2.25 * float32(v1))
                                    _ = v15
                                case 387:
                                    v1 -= -2.25
                                }
                                _ = v14
                            } else {
                                fmt.Println(v6)
                            }
                            if v5 {
                                v16 := GenS1{F0: 194, F1: (v4 << ((63) & 63))}
                                _ = v16
                            }
                            switch fi10 {
                            case 1:
                                switch fi13 {
                                case 0:
                                    v17 := GenS3{F0: 128}
                                    _ = v17
                                default:
                                    var v18 int8 = -1
                                    _ = v18
                                }
                            default:
                                var v19 int64 = 541
                                _ = v19
                            }
                        }
                        switch "\n\t" {
                        case "\n\t":
                            if (!v5) {
                                if v5 {
                                    fmt.Println(v4)
                                    v20 := []int{127, fi13, fi13, -2147483648}
                                    if v5 {
                                        if len(v20) > 0 {
                                            v20[0] = fi10
                                        }
                                    } else {
                                        v1 = ((1e10 + 1e10) / 1e10)
                                    }
                                    v21 := GenS4{F0: 1000, F1: 237, F2: 65535}
                                    var v22 uint8 = v2
                                    _ = v20
                                    _ = v21
                                    _ = v22
                                } else {
                                    v2 = (v2 % ((255) | 1))
                                }
                                if ((true && (1e10 <= (0.0 / 1.5))) || (!v5)) {
                                    var v23 int = -275
                                    _ = v23
                                }
                                fmt.Println(v5)
                            }
                        case "":
                            v24 := v3[1]
                            _ = v24
                        case "a":
                            var v25 uint64 = 310
                            _ = v25
                        }
                        v4 += (v4 >> ((489) & 63))
                        _ = fi13
                    }
                    var v26 uint16 = 189
                    v27 := v3[32768]
                    _ = v26
                    _ = v27
                }
            case 32767:
                var v28 int64 = -598
                _ = v28
            case -9223372036854775808:
                v6 = v6
            default:
                v4 = (v4 & v4)
            }
            v29 := []int{}
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 bool = v5
                switch fi30 {
                case 128:
                    _ = len(v29)
                case 42:
                    if ((float64(float32(v1)) / float64(3.14159)) != 1e10) {
                        var v32 uint8 = (v2 & v2)
                        fmt.Println(fi10, v31, fi30)
                        v6 = "\n\t"
                        v32 -= (v32 & v32)
                        if ((!v5) && ("" <= ("hello" + ""))) {
                            fmt.Println(v31, v4)
                            v33 := GenS5{F0: false}
                            _ = v33
                        }
                        _ = v32
                    }
                default:
                    v3[fi30] = ""
                }
                v34 := GenS5{F0: v5}
                _ = fi30
                _ = v31
                _ = v34
            }
            _ = fi10
            _ = v29
        }
        v35 := GenS6{F0: 1, F1: -368, F2: 42}
        _ = v35
    } else {
        if v5 {
            v36 := []int{127, 1}
            delete(v3, 4294967295)
            for fi37 := 0; fi37 < 11; fi37++ {
                v38 := GenS7{F0: 0, F1: ((("" + (v6 + v6)) + v6) + v6)}
                var v39 uint16 = 42
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = v36
        }
        switch "hello" {
        case "a":
            v5 = ((!(32768 >= 9223372036854775807)) || v5)
        case "":
            v1 = 3.14159
        case "\n\t":
            v3[747] = ""
        }
    }
    if (v6 < v6) {
        if (!((1 > -2147483648) && (!v5))) {
            v3[-2147483648] = v6
        }
        v40 := []int{9223372036854775807}
        _ = v40
    } else {
        v1 += 1.5
        for fi41 := 0; fi41 < 5; fi41++ {
            v42 := GenS6{F0: v4, F1: 42, F2: 690}
            _ = fi41
            _ = v42
        }
    }
    fmt.Println(v6, v5, v2)
    fmt.Println(v2)
    delete(v3, -1)
    for fi43 := 0; fi43 < 18; fi43++ {
        v6 = v6
        var v44 bool = (1 != 2147483647)
        v45 := v3[fi43]
        _ = fi43
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 14; fi46++ {
        delete(v3, (-(fi46)))
        switch (v6 + "a") {
        case "\n\t":
            _ = len(v3)
        }
        if (!(((-2.25 * 1.5) - v1) >= v1)) {
            switch 32767 {
            case 1:
                var v47 uint64 = 0
                _ = v47
            default:
                v3[fi46] = (("" + (v6 + v6)) + ("\n\t" + v6))
            }
        }
        _ = fi46
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
