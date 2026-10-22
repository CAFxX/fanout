package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    var v1 int8 = 127
    v2 := []int{153}
    var v3 int16 = -32768
    var v4 float64 = float64(3.14159)
    fmt.Println(v3, v1, v4)
    if false {
        v3 = v3
        switch 956 {
        case -2147483648:
            if false {
                v5 := GenS1{F0: 249, F1: false}
                v4 += (((v4 / (0.1 - 1e10)) - v4) / 3.14159)
                _ = v5
            } else {
                _ = len(v2)
                if false {
                    fmt.Println(v1)
                    for fi6 := 0; fi6 < 19; fi6++ {
                        var v7 int32 = 2147483647
                        _ = fi6
                        _ = v7
                    }
                    switch 418 {
                    case 42:
                        if (37 == 70) {
                            v4 = ((float64(((1e10 * (3.14159 + 1e10)) / ((0.1 - (3.14159 - ((3.14159 / 1.0) + 0.1))) + float32(v4)))) + v4) + float64(0.0))
                            v3 = v3
                            v3 += v3
                        } else {
                            v8 := GenS1{F0: 42, F1: (((true || (25 < 956)) && true) || (825 < 698))}
                            v4 += (-2.25 + (1.5 + v4))
                            _ = v8
                        }
                    case 2147483647:
                        v9 := GenS1{F0: 2, F1: (!((42 < 2) && true))}
                        _ = v9
                    case 9223372036854775807:
                        for fi10 := 0; fi10 < 5; fi10++ {
                            if len(v2) > 0 {
                                v2[0] = 32767
                            }
                            _ = fi10
                        }
                    default:
                        if (!true) {
                            v4 = 0.0
                            v4 = float64(0.1)
                            v3 = (v3 ^ v3)
                        } else {
                            for fi11 := 0; fi11 < 16; fi11++ {
                                var v12 int = fi11
                                for fi13 := 0; fi13 < 10; fi13++ {
                                    v1 += 36
                                    _ = fi13
                                }
                                _ = fi11
                                _ = v12
                            }
                            if len(v2) > 0 {
                                v2[0] = 128
                            }
                            var v14 bool = (128 == 15)
                            _ = v14
                        }
                    }
                }
            }
        }
    } else {
        var v15 uint8 = 224
        _ = v15
    }
    var v16 uint32 = 990
    v3 -= 1
    var v17 uint32 = v16
    fmt.Println(v3)
    if (!(255 != 37)) {
        fmt.Println(v3)
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        switch (fi18 & fi18) {
        case -9223372036854775808:
            v4 = float64(0.0)
        }
        var v19 int8 = (^(v1))
        switch ("hello" + "hello") {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "":
            v17 -= v17
        }
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 uint32 = 2147483648
        var v22 bool = (("a" + "") > "hello")
        var v23 int8 = v1
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 uint8 = 128
    for fi25 := 0; fi25 < 13; fi25++ {
        for fi26 := 0; fi26 < 6; fi26++ {
            if false {
                if true {
                    if true {
                        switch fi25 {
                        case 358:
                            v3 = v3
                        case 32768:
                            var v27 int8 = (v1 / ((-1) | 1))
                            _ = v27
                        case -516:
                            fmt.Println(v16)
                        }
                        if len(v2) > 0 {
                            v2[0] = (fi26 % ((9223372036854775807) | 1))
                        }
                        var v28 int32 = 859
                        _ = v28
                    } else {
                        v24 = (^(v24))
                        v24 = (^(v24))
                        v16 = v17
                    }
                    v1 = v1
                } else {
                    var v29 uint16 = 2
                    if (!(0.1 == 0.0)) {
                        v30 := GenS1{F0: v29, F1: (!((false || (789 < 0)) || ((int8(v29) > -20) || (((0.0 == float32(v4)) && (true && (((^(v29)) != (v29 | 0)) || true))) || false))))}
                        v31 := GenS1{F0: v29, F1: (-1 > -746)}
                        if len(v2) > 0 {
                            v2[0] = int(v29)
                        }
                        _ = v30
                        _ = v31
                    }
                    _ = v29
                }
                var v32 int32 = 726
                v33 := GenS2{F0: fi25, F1: (^(v3))}
                _ = v32
                _ = v33
            }
            var v34 uint16 = 670
            v35 := GenS3{F0: 849, F1: (v34 | v34)}
            _ = fi26
            _ = v34
            _ = v35
        }
        _ = fi25
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        v1 = (v1 % (((v1 % ((-125) | 1))) | 1))
        fmt.Println(v1, v1)
        var v37 int32 = -2147483648
        _ = fi36
        _ = v37
    }
    v2 = append(v2, 32768)
    v38 := GenS2{F0: 0, F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v4
}
