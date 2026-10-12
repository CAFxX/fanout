package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int; F1 int8 }

func main() {
    var v1 uint32 = 784
    var v2 float32 = ((1.5 - (0.0 * 1e10)) * 0.1)
    v3 := []string{("a" + ""), "", "a", ("\n\t" + "")}
    v4 := []uint32{}
    var v5 int16 = 1000
    var v6 float64 = float64((float32(0.0) / ((1.5 + v2) * v2)))
    if false {
        fmt.Println(v5, v1, v5)
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 string = ""
            _ = fi7
            _ = v8
        }
        if true {
            var v9 bool = true
            _ = v9
        }
    } else {
        fmt.Println(v5, v5, v6)
        for fi10 := 0; fi10 < 10; fi10++ {
            if true {
                v11 := GenS1{F0: 969}
                var v12 float32 = v2
                for fi13 := 0; fi13 < 4; fi13++ {
                    v5 += (v5 >> ((309) & 63))
                    var v14 float64 = float64(((v2 + v12) - 0.0))
                    _ = fi13
                    _ = v14
                }
                _ = v11
                _ = v12
            } else {
                if (-1 < -2147483648) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v15 string = "hello"
                    if false {
                        v6 += -2.25
                        var v16 string = ""
                        v16 = "foo bar"
                        _ = v16
                    } else {
                        var v17 uint = 97
                        v3 = append(v3, "hello")
                        _ = v17
                    }
                    _ = v15
                }
                if (true || true) {
                    var v18 float32 = 1e10
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v1 = v1
                        if (255 > uint16(v1)) {
                            if (42 != 1) {
                                v5 -= (^(v5))
                                _ = len(v3)
                                switch (-(fi19)) {
                                case 32767:
                                    var v20 uint16 = 790
                                    _ = v20
                                case -261:
                                    if (v5 < (-(v5))) {
                                        fmt.Println(v18)
                                    }
                                default:
                                    var v21 int32 = 2147483647
                                    _ = v21
                                }
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            v6 = (1.5 - (((1e10 - -2.25) * v6) - v6))
                            if len(v4) > 0 {
                                v4[0] = (^(v1))
                            }
                        } else {
                            var v22 int32 = -511
                            _ = v22
                        }
                        _ = fi19
                    }
                    if false {
                        var v23 int16 = (v5 - 32767)
                        _ = v23
                    } else {
                        v24 := GenS2{F0: 201, F1: (v5 % (((v5 % ((v5) | 1))) | 1)), F2: 0}
                        _ = v24
                    }
                    _ = v18
                } else {
                    var v25 string = "a"
                    _ = v25
                }
            }
            _ = fi10
        }
        var v26 int32 = -477
        _ = v26
    }
    if true {
        v3 = append(v3, "")
        v2 = (1e10 + float32(v6))
    }
    v27 := GenS3{F0: v5}
    var v28 int = 4294967295
    if (((1 != -1) && (!(14 >= 11))) && (-1 != 100)) {
        var v29 int64 = -244
        for fi30 := 0; fi30 < 17; fi30++ {
            fmt.Println(v1, fi30, v29)
            _ = fi30
        }
        var v31 uint8 = 127
        _ = v29
        _ = v31
    } else {
        _ = len(v3)
        if (-125 <= 127) {
            for fi32 := 0; fi32 < 18; fi32++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi32
            }
        } else {
            if true {
                if true {
                    v1 = (^(v1))
                    v5 = (v5 & 1)
                }
            } else {
                if (!false) {
                    _ = len(v4)
                }
                if true {
                    switch v28 {
                    case -667:
                        if len(v4) > 0 {
                            v4[0] = v1
                        }
                    case 871:
                        var v33 uint = 18446744073709551615
                        _ = v33
                    default:
                        if len(v4) > 0 {
                            v4[0] = v1
                        }
                    }
                }
                v1 -= (^(v1))
            }
            if (false || false) {
                v28 = (-(v28))
                var v34 int = (v28 - v28)
                if true {
                    var v35 int64 = 9223372036854775807
                    var v36 uint8 = 127
                    v37 := GenS2{F0: v36, F1: v5, F2: 65535}
                    _ = v35
                    _ = v36
                    _ = v37
                }
                _ = v34
            } else {
                for fi38 := 0; fi38 < 11; fi38++ {
                    _ = len(v4)
                    var v39 uint8 = uint8(v1)
                    var v40 uint = 708
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
        }
        if (!(-2147483648 > -879)) {
            switch (v28 ^ v28) {
            case -700:
                v2 += v2
            case 9223372036854775807:
                v41 := GenS4{F0: 4294967295}
                _ = v41
            }
        }
    }
    fmt.Println(v5)
    if ((1.5 < v2) && ((v2 - (v2 + ((0.0 + v2) - v2))) != v2)) {
        fmt.Println(v1, v28, v6)
        switch "a" {
        case "\n\t":
            if (v6 < ((((v6 * (((3.14159 / (float64(float32(((float64((((v2 * 1.5) / 1.0) * v2)) + (((v6 / (float64(-2.25) / (v6 - (((((v6 * v6) - float64(float32((v6 * ((((1e10 - 3.14159) / 1.5) / 3.14159) + v6))))) / -2.25) / (1.5 + v6)) - ((v6 * v6) / (3.14159 * float64(v2))))))) / 1.0) / v6)) * v6))) / 3.14159)) * (v6 * 1.5)) / (3.14159 / 1e10))) * v6) * (0.0 * (float64(v2) + float64(v2)))) - (v6 * ((v6 + v6) - (v6 + -2.25))))) {
                v1 = 65535
                if false {
                    fmt.Println(v5, v2)
                    var v42 int64 = 9223372036854775807
                    _ = v42
                } else {
                    v43 := GenS4{F0: (v28 - v28)}
                    v44 := GenS5{F0: 255, F1: 42}
                    var v45 uint64 = 18446744073709551615
                    _ = v43
                    _ = v44
                    _ = v45
                }
            }
        }
    }
    v46 := GenS6{F0: (^(v28)), F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
