package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int64 }

func main() {
    var v1 uint32 = 22
    var v2 float64 = (((1.5 * (1e10 / 1.5)) - 1e10) / float64((1e10 + -2.25)))
    v3 := make(map[int]bool)
    v4 := []uint32{v1, (v1 % ((v1) | 1))}
    var v5 uint = 255
    var v6 int32 = 1
    switch -865 {
    case 4294967295:
        for fi7 := 0; fi7 < 15; fi7++ {
            if false {
                var v8 int8 = int8(fi7)
                _ = v8
            } else {
                var v9 int16 = 1000
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 float64 = 3.14159
                    var v12 int8 = 28
                    if false {
                        var v13 uint16 = 2
                        var v14 int = fi10
                        v14 = fi7
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = v9
            }
            v15 := GenS1{F0: v5}
            v16 := GenS2{F0: 100, F1: -1}
            _ = fi7
            _ = v15
            _ = v16
        }
    default:
        fmt.Println(v2)
    }
    if false {
        switch ("\n\t" + ((("foo bar" + ("a" + "foo bar")) + "") + "a")) {
        case "\n\t":
            for fi17 := 0; fi17 < 5; fi17++ {
                if false {
                    var v18 int8 = -1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v18
                }
                switch ("a" + "a") {
                case "foo bar":
                    for fi19 := 0; fi19 < 14; fi19++ {
                        if true {
                            fmt.Println(v6, fi17, v2)
                            v20 := GenS2{F0: -128, F1: (v6 ^ (v6 * 2147483647))}
                            v6 += v6
                            _ = v20
                        }
                        switch 1 {
                        case 978:
                            v21 := GenS2{F0: 100, F1: -199}
                            _ = v21
                        case 4294967295:
                            if (3.14159 <= float32((float64(3.14159) / float64(float32(v2))))) {
                                if (92 != 0) {
                                    v2 += (((3.14159 + 1e10) * (((-2.25 + 0.1) / ((((1e10 + v2) / 1.0) + (v2 - -2.25)) - (1.5 / v2))) / 1.0)) * v2)
                                    v2 = v2
                                    var v22 uint64 = 2
                                    _ = v22
                                }
                                fmt.Println(v5, v2)
                                v1 = v1
                            } else {
                                fmt.Println(v5, v6)
                                v23 := GenS3{F0: 42}
                                _ = v23
                            }
                        case 127:
                            var v24 int32 = -1
                            _ = v24
                        }
                        var v25 int64 = -1
                        _ = fi19
                        _ = v25
                    }
                case "\n\t":
                    v26 := GenS2{F0: 1, F1: 689}
                    _ = v26
                }
                v5 = v5
                _ = fi17
            }
        default:
            v27 := v3[32768]
            _ = v27
        }
    }
    if (!false) {
        switch "" {
        case "":
            v2 -= v2
        case "hello":
            var v28 uint8 = uint8(v5)
            _ = v28
        case "a":
            if false {
                if len(v4) > 0 {
                    v4[0] = v1
                }
            } else {
                if (0.1 == (0.1 + 0.0)) {
                    fmt.Println(v5, v2, v5)
                    for fi29 := 0; fi29 < 1; fi29++ {
                        v1 -= v1
                        for fi30 := 0; fi30 < 10; fi30++ {
                            fmt.Println(v2, fi30, v5)
                            v31 := GenS4{F0: (!(1e10 == (float32(((float64((0.1 - 1e10)) + ((-2.25 / v2) * v2)) + ((v2 * v2) / (1.5 + 3.14159)))) + ((3.14159 * 1e10) * 0.0)))), F1: uint(v1)}
                            v32 := GenS4{F0: (!("hello" >= "a")), F1: (v5 % (((v5 & (v5 - 0))) | 1))}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        var v33 int = fi29
                        _ = fi29
                        _ = v33
                    }
                    var v34 bool = (!true)
                    _ = v34
                }
                v2 += 3.14159
            }
        default:
            if true {
                v35 := v3[-9223372036854775808]
                v36 := GenS5{F0: -887}
                v6 = (^(v6))
                _ = v35
                _ = v36
            } else {
                fmt.Println(v5, v6)
                for fi37 := 0; fi37 < 3; fi37++ {
                    _ = len(v4)
                    var v38 int32 = (v6 ^ (v6 >> ((v5) & 63)))
                    v2 = 0.1
                    _ = fi37
                    _ = v38
                }
                var v39 float32 = 0.0
                _ = v39
            }
        }
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        v41 := v3[(fi40 & int(v6))]
        if v41 {
            switch ("foo bar" + ("foo bar" + "")) {
            case "foo bar":
                v2 = v2
            case "\n\t":
                v42 := GenS5{F0: -28}
                _ = v42
            default:
                if true {
                    fmt.Println(v5, v1, v41)
                    fmt.Println(v2)
                } else {
                    switch ("hello" + "foo bar") {
                    case "":
                        v6 = 0
                    case "hello":
                        if (!(v6 > (v6 % (((v6 & v6)) | 1)))) {
                            var v43 uint = 768
                            _ = v43
                        } else {
                            var v44 uint16 = 0
                            fmt.Println(v2)
                            _ = v44
                        }
                    case "\n\t":
                        var v45 int8 = 0
                        _ = v45
                    default:
                        v6 = v6
                    }
                    v41 = (823 == 1)
                    if ("foo bar" > ("hello" + "foo bar")) {
                        var v46 int8 = 1
                        delete(v3, fi40)
                        var v47 uint = 65535
                        _ = v46
                        _ = v47
                    }
                }
            }
        }
        v1 += (^(v1))
        _ = fi40
        _ = v41
    }
    fmt.Println(v5, v2, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2, v2, v1)
    _ = len(v3)
    var v48 int = 9223372036854775807
    v4 = append(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
