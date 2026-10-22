package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int16; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint8 }

func main() {
    var v1 bool = (12 > 0)
    var v2 float32 = 1e10
    v3 := make(map[uint32]string)
    v4 := []string{"hello", ("a" + ("foo bar" + "hello")), "foo bar", "foo bar"}
    if (!v1) {
        fmt.Println(v2, v1)
    } else {
        var v5 bool = ((1.5 * (v2 - float32(1.5))) >= (v2 - (1e10 / ((v2 + -2.25) - float32(((1.5 - (float64((v2 / v2)) / 0.1)) - -2.25))))))
        _ = v5
    }
    v1 = (false && v1)
    v6 := GenS1{F0: 127, F1: ("a" + ""), F2: 4294967295}
    if (-638 >= -752) {
        var v7 uint32 = 920
        v8 := GenS1{F0: 32767, F1: ("foo bar" + "a"), F2: 263}
        _ = v7
        _ = v8
    }
    v2 = 3.14159
    v2 = v2
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := v3[2147483648]
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 2}
    for fi12 := 0; fi12 < 1; fi12++ {
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        switch fi12 {
        case 128:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int32 = 825
                if true {
                    var v15 int16 = 32767
                    _ = v15
                } else {
                    switch ("a" + "") {
                    case "foo bar":
                        switch (fi12 ^ fi13) {
                        case -748:
                            v4 = append(v4, "")
                        case -9223372036854775808:
                            switch (^(fi13)) {
                            case 32767:
                                if ((!(65535 <= 18446744073709551615)) && false) {
                                    v16 := GenS1{F0: fi12, F1: "hello", F2: 0}
                                    _ = v16
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                }
                                fmt.Println(v14)
                            case 543:
                                fmt.Println(v1, v1)
                            case 127:
                                switch fi12 {
                                case 308:
                                    var v17 uint16 = 32767
                                    _ = v17
                                case 0:
                                    var v18 uint16 = 42
                                    _ = v18
                                case 127:
                                    fmt.Println(v14, v1)
                                default:
                                    if (v1 && (((((("a" + ("a" + "\n\t")) + (("\n\t" + ("a" + "hello")) + "foo bar")) + "\n\t") + ("" + "foo bar")) + "\n\t") == "")) {
                                        v19 := GenS1{F0: (-(fi12)), F1: "hello", F2: 458}
                                        v14 -= v14
                                        v2 -= v2
                                        _ = v19
                                    }
                                }
                            }
                            var v20 uint16 = 796
                            v14 -= (v14 ^ v14)
                            _ = v20
                        default:
                            v14 = v14
                        }
                    case "hello":
                        fmt.Println(v14, v1)
                    default:
                        fmt.Println(v2, v14)
                        if v1 {
                            var v21 uint16 = 255
                            v2 = ((float32(1e10) - v2) / ((v2 / v2) - v2))
                            _ = v21
                        } else {
                            v22 := GenS2{F0: 42}
                            fmt.Println(v14)
                            _ = v22
                        }
                    }
                    v23 := GenS2{F0: 0}
                    v2 += 1.5
                    var v24 int8 = -128
                    _ = v23
                    _ = v24
                }
                _ = fi13
                _ = v14
            }
        case 574:
            v3[1] = (("\n\t" + "foo bar") + ("\n\t" + ("" + "a")))
        case 127:
            for fi25 := 0; fi25 < 9; fi25++ {
                v1 = v1
                _ = fi25
            }
        default:
            var v26 int64 = -9223372036854775808
            _ = v26
        }
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 float64 = float64(((1e10 * v2) / 3.14159))
            v29 := v3[2147483647]
            v1 = (-1 > -1)
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi12
    }
    if false {
        if v1 {
            v1 = (!v1)
        } else {
            for fi30 := 0; fi30 < 11; fi30++ {
                _ = len(v4)
                v2 = (3.14159 - v2)
                _ = fi30
            }
        }
        v31 := GenS3{F0: 0, F1: 448, F2: 0}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v31
    } else {
        var v32 uint32 = 0
        for fi33 := 0; fi33 < 9; fi33++ {
            v34 := GenS4{F0: 0, F1: 1}
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    fmt.Println(v2, v1)
    var v35 uint = 65535
    if (100 == 29) {
        for fi36 := 0; fi36 < 15; fi36++ {
            v35 += 0
            v1 = (true && (!v1))
            _ = fi36
        }
        for fi37 := 0; fi37 < 16; fi37++ {
            var v38 float64 = (float64((v2 * (v2 / ((v2 - float32(((1.5 + 1e10) / 0.1))) / (v2 + ((0.0 + float32(3.14159)) - 1e10)))))) - -2.25)
            _ = fi37
            _ = v38
        }
    }
    v39 := GenS1{F0: 127, F1: "", F2: 18446744073709551615}
    var v40 int8 = 14
    v1 = (v1 && v1)
    var v41 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v6
}
