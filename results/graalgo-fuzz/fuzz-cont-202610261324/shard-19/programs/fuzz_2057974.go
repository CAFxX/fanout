package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS4 struct { F0 bool; F1 uint8; F2 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64 }
func genFn0() int {
    var v5 uint32 = 972
    var v6 uint64 = 2
    v7 := GenS1{F0: ("foo bar" + "foo bar")}
    _ = v5
    _ = v6
    _ = v7
    return -370
}
func genFn1(p0 int, p1 uint16) uint32 {
    v8 := GenS1{F0: ("\n\t" + "hello")}
    v9 := GenS2{F0: 4294967295, F1: 997}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 0
    var v3 uint16 = 806
    v4 := []uint32{315, 589, 754, 1}
    v10 := GenS1{F0: ("foo bar" + "")}
    switch "" {
    case "\n\t":
        if false {
            var v11 int64 = 134
            v12 := GenS3{F0: v3, F1: 128, F2: 0}
            var v13 uint16 = (v3 / (((v3 << ((501) & 63))) | 1))
            _ = v11
            _ = v12
            _ = v13
        } else {
            v2 -= v2
            var v14 uint = 4294967295
            _ = v14
        }
    case "":
        v15 := GenS1{F0: "foo bar"}
        _ = v15
    }
    v2 -= v2
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := GenS3{F0: v3, F1: 0, F2: -1}
            var v19 int32 = v2
            fmt.Println(v19, fi16)
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    v2 = v2
    if (!((514 == 18446744073709551615) && false)) {
        fmt.Println(v2, v2)
    } else {
        switch ("a" + "hello") {
        case "foo bar":
            if (-98 >= -114) {
                for fi20 := 0; fi20 < 11; fi20++ {
                    if true {
                        v21 := v1[(("\n\t" + "a") + ("" + ("" + ("a" + ("a" + "foo bar")))))]
                        switch "a" {
                        case "":
                            var v22 uint32 = 852
                            _ = v22
                        case "a":
                            v21 = ((v3 + (v3 >> ((658) & 63))) != (v3 % ((v3) | 1)))
                        default:
                            for fi23 := 0; fi23 < 2; fi23++ {
                                switch fi20 {
                                case 127:
                                    var v24 int32 = (^(v2))
                                    _ = v24
                                case 0:
                                    v2 = (v2 - v2)
                                }
                                v21 = (251 == 255)
                                v1["a"] = v21
                                _ = fi23
                            }
                        }
                        _ = v21
                    }
                    _ = fi20
                }
                delete(v1, ("hello" + "foo bar"))
            }
        case "hello":
            switch 2147483647 {
            case 32768:
                var v25 float32 = (3.14159 + (-2.25 - float32(0.1)))
                _ = v25
            case 32767:
                var v26 bool = true
                _ = v26
            case 284:
                v3 = 255
            default:
                switch int(v2) {
                case 32767:
                    v2 += (v2 | v2)
                case 2147483647:
                    var v27 string = ("foo bar" + "\n\t")
                    _ = v27
                case 0:
                    v28 := GenS3{F0: (v3 / ((42) | 1)), F1: 75, F2: 42}
                    _ = v28
                default:
                    v2 -= v2
                }
            }
        }
    }
    if (!((889 == 9223372036854775807) || false)) {
        _ = len(v1)
        v29 := GenS2{F0: 1, F1: 0}
        _ = v29
    }
    fmt.Println(v2, v2)
    v2 += (v2 * v2)
    v4 = append(v4, 2147483648)
    if false {
        fmt.Println(v3)
    }
    switch ("a" + "hello") {
    case "":
        for fi30 := 0; fi30 < 14; fi30++ {
            for fi31 := 0; fi31 < 5; fi31++ {
                if (float64(1.5) != 1e10) {
                    v32 := GenS1{F0: "foo bar"}
                    _ = v32
                }
                v33 := GenS2{F0: 2, F1: 0}
                _ = fi31
                _ = v33
            }
            fmt.Println(fi30)
            v34 := GenS2{F0: 2, F1: 1}
            _ = fi30
            _ = v34
        }
    default:
        v3 = (v3 | 0)
    }
    if ((v3 << ((839) & 63)) >= v3) {
        var v35 bool = true
        var v36 float64 = 3.14159
        _ = v35
        _ = v36
    }
    var v37 uint8 = 212
    var v38 int8 = 0
    switch ("" + ("foo bar" + ("" + "a"))) {
    case "hello":
        switch 4294967295 {
        case 969:
            v2 = -597
        case 0:
            fmt.Println(v2, v3)
        case -9223372036854775808:
            v39 := GenS4{F0: ((v38 - (v38 >> ((1) & 63))) < int8(v2)), F1: v37, F2: true}
            _ = v39
        }
    case "a":
        for fi40 := 0; fi40 < 5; fi40++ {
            if true {
                if (!true) {
                    if (("foo bar" + "hello") < "") {
                        v41 := GenS5{F0: (v38 >> ((18446744073709551615) & 63))}
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 5; fi42++ {
                            v3 -= (v3 << ((818) & 63))
                            fmt.Println(fi40)
                            _ = fi42
                        }
                    }
                } else {
                    v43 := GenS6{F0: -1}
                    for fi44 := 0; fi44 < 3; fi44++ {
                        if true {
                            var v45 int32 = 2147483647
                            v37 -= v37
                            _ = v45
                        } else {
                            fmt.Println(v2)
                            fmt.Println(v37, fi40, fi40)
                        }
                        _ = fi44
                    }
                    _ = v43
                }
            } else {
                var v46 uint64 = 42
                _ = v46
            }
            fmt.Println(v3, v38)
            _ = fi40
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
