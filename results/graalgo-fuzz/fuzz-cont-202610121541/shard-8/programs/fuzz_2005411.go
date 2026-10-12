package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0() int32 {
    var v5 int = 546
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int8 = -128
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := []int{9223372036854775807}
            v7 = v7
            if ((fi8 == 42) && true) {
                fmt.Println(fi6, fi6, v7)
                if (((-(v7)) != (v7 | (v7 << ((65535) & 63)))) && false) {
                    var v10 string = "\n\t"
                    _ = v10
                } else {
                    var v11 uint8 = 127
                    _ = v11
                }
                v12 := GenS1{F0: 1, F1: 42, F2: 351}
                _ = v12
            }
            if ((v7 / (((^(v7))) | 1)) < 42) {
                switch 1 {
                case -864:
                    var v13 uint16 = 255
                    _ = v13
                case -1:
                    var v14 int16 = 45
                    _ = v14
                case -562:
                    var v15 string = ("hello" + "")
                    _ = v15
                default:
                    v7 += v7
                }
            } else {
                var v16 string = ("" + ("foo bar" + "a"))
                v7 = (v7 << ((10) & 63))
                _ = v16
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v5 -= -313
    _ = v5
    return -1
}

func main() {
    var v1 int8 = -128
    var v2 uint8 = 2
    var v3 uint = 65535
    var v4 uint8 = 39
    var v17 uint8 = 2
    v18 := []int{-1}
    if (!(-32768 == 32767)) {
        var v19 float64 = ((-2.25 - 3.14159) * 1.5)
        switch ("foo bar" + "") {
        case "\n\t":
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        case "":
            _ = len(v18)
        }
        _ = v19
    } else {
        v18 = append(v18, 32767)
        switch "a" {
        case "":
            if true {
                v20 := GenS2{F0: -324}
                _ = v20
            }
        default:
            var v21 int32 = 1
            _ = v21
        }
    }
    var v22 uint = (v3 - 110)
    v1 -= v1
    var v23 int16 = 0
    v24 := GenS3{F0: 0, F1: 42}
    if (1 != v23) {
        v25 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 796}
        _ = v25
    } else {
        for fi26 := 0; fi26 < 6; fi26++ {
            if true {
                switch fi26 {
                case 2147483647:
                    switch 0 {
                    case 127:
                        for fi27 := 0; fi27 < 15; fi27++ {
                            if true {
                                if true {
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                } else {
                                    var v28 int = -9223372036854775808
                                    _ = v28
                                }
                            }
                            fmt.Println(v2)
                            if len(v18) > 0 {
                                v18[0] = (fi26 * 32768)
                            }
                            switch fi26 {
                            case -9223372036854775808:
                                v18 = append(v18, (fi27 >> ((v22) & 63)))
                            case -2147483648:
                                var v29 int32 = 2147483647
                                _ = v29
                            case 32768:
                                var v30 bool = true
                                _ = v30
                            }
                            _ = fi27
                        }
                    case 367:
                        v31 := GenS1{F0: -598, F1: v23, F2: 693}
                        _ = v31
                    case 32768:
                        if (int8(v22) != v1) {
                            v17 = v4
                            v32 := GenS3{F0: 2147483647, F1: 889}
                            for fi33 := 0; fi33 < 20; fi33++ {
                                v18 = append(v18, (fi33 + (fi26 >> ((v22) & 63))))
                                var v34 int64 = 312
                                _ = fi33
                                _ = v34
                            }
                            _ = v32
                        }
                    }
                case 42:
                    switch "foo bar" {
                    case "foo bar":
                        var v35 uint64 = 0
                        _ = v35
                    case "\n\t":
                        fmt.Println(v3)
                    case "a":
                        v36 := GenS1{F0: 1, F1: (^(v23)), F2: 42}
                        _ = v36
                    default:
                        for fi37 := 0; fi37 < 15; fi37++ {
                            var v38 int = 0
                            var v39 uint32 = 2147483648
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    }
                default:
                    for fi40 := 0; fi40 < 11; fi40++ {
                        v17 -= 255
                        v18 = append(v18, fi40)
                        v4 = (v4 * 255)
                        _ = fi40
                    }
                }
            } else {
                v1 = (v1 << (((^(v22))) & 63))
            }
            var v41 uint32 = 916
            _ = fi26
            _ = v41
        }
        fmt.Println(v22, v23)
        if len(v18) > 0 {
            v18[0] = 4294967295
        }
    }
    v42 := GenS4{F0: 4294967295, F1: (v1 >> ((v22) & 63))}
    v43 := GenS4{F0: v3, F1: (v1 - (^(v1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
