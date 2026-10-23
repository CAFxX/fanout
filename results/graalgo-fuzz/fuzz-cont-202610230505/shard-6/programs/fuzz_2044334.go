package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 int32; F2 int16 }
type GenS6 struct { F0 bool; F1 uint8; F2 bool }
type GenS7 struct { F0 uint8; F1 int8; F2 uint64 }

func main() {
    var v1 uint8 = 0
    v2 := []string{"\n\t", ("a" + "\n\t"), "\n\t", ("a" + "foo bar")}
    var v3 int32 = -2147483648
    v4 := GenS1{F0: 0, F1: v3}
    for fi5 := 0; fi5 < 1; fi5++ {
        switch ("foo bar" + "\n\t") {
        case "":
            for fi6 := 0; fi6 < 10; fi6++ {
                switch fi6 {
                case 530:
                    var v7 uint32 = 2147483647
                    _ = v7
                default:
                    var v8 int = fi5
                    _ = v8
                }
                fmt.Println(fi5, v1)
                v9 := GenS1{F0: 381, F1: (v3 << ((255) & 63))}
                _ = fi6
                _ = v9
            }
        case "hello":
            v2 = append(v2, "a")
        case "a":
            v10 := GenS2{F0: ("a" + ("\n\t" + (("\n\t" + "\n\t") + ("hello" + "")))), F1: 2, F2: 4294967295}
            _ = v10
        default:
            v1 += v1
        }
        var v11 int16 = 626
        _ = fi5
        _ = v11
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v12 := GenS3{F0: -32768}
    v13 := GenS4{F0: -1}
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(v1, fi14, fi14)
        for fi15 := 0; fi15 < 1; fi15++ {
            if (true || true) {
                var v16 uint32 = 2
                if (497 < 0) {
                    fmt.Println(fi14, fi15)
                    var v17 bool = (42 > 32767)
                    v3 = (v3 - int32(v16))
                    if true {
                        var v18 int = 0
                        v1 = (v1 / (((v1 << ((65535) & 63))) | 1))
                        _ = v18
                    }
                    _ = v17
                }
                _ = v16
            }
            v1 = v1
            _ = fi15
        }
        _ = fi14
    }
    var v19 int = 919
    fmt.Println(v19, v19, v3)
    fmt.Println(v3, v1, v1)
    if ((1000 != 1000) && false) {
        var v20 int8 = 127
        v21 := GenS1{F0: 277, F1: (v3 ^ (v3 | (v3 & v3)))}
        if false {
            if true {
                for fi22 := 0; fi22 < 5; fi22++ {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v19 += (fi23 / ((fi23) | 1))
                        v2 = append(v2, "foo bar")
                        var v24 uint64 = 2
                        _ = fi23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 2; fi25++ {
                        if true {
                            var v26 bool = true
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v27 float64 = 1.5
                            _ = v26
                            _ = v27
                        }
                        v2 = append(v2, "")
                        _ = fi25
                    }
                    _ = fi22
                }
                var v28 uint64 = 0
                if false {
                    switch v19 {
                    case -1:
                        if (824 == uint16(v28)) {
                            if ((v20 & int8(v28)) >= v20) {
                                fmt.Println(v28, v3)
                            } else {
                                var v29 float32 = (1e10 - (-2.25 * (1e10 + 3.14159)))
                                _ = v29
                            }
                            v28 = (v28 & 322)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            v20 = 100
                            for fi30 := 0; fi30 < 5; fi30++ {
                                if (!true) {
                                    _ = len(v2)
                                } else {
                                    for fi31 := 0; fi31 < 14; fi31++ {
                                        switch -2147483648 {
                                        case 428:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if ("\n\t" != "\n\t") {
                                                v28 = (v28 & v28)
                                            } else {
                                                v1 += v1
                                            }
                                            var v32 float32 = (3.14159 / (-2.25 * 3.14159))
                                            _ = v32
                                        case -296:
                                            _ = len(v2)
                                        case 32768:
                                            if (2 <= 4294967295) {
                                                switch ("\n\t" + ((("" + "\n\t") + "a") + "foo bar")) {
                                                case "":
                                                    v19 = (^(fi30))
                                                }
                                            }
                                        default:
                                            fmt.Println(v3)
                                        }
                                        _ = fi31
                                    }
                                }
                                switch ("foo bar" + ("" + ("foo bar" + ""))) {
                                case "":
                                    v33 := GenS3{F0: -228}
                                    _ = v33
                                }
                                for fi34 := 0; fi34 < 10; fi34++ {
                                    v35 := GenS5{F0: -503, F1: (v3 ^ 1), F2: 32767}
                                    _ = fi34
                                    _ = v35
                                }
                                _ = fi30
                            }
                        }
                    case 44:
                        v36 := GenS3{F0: 0}
                        _ = v36
                    }
                    fmt.Println(v20, v20)
                }
                _ = v28
            } else {
                fmt.Println(v20, v1)
                v37 := GenS6{F0: ((!true) && false), F1: (v1 % ((243) | 1)), F2: (!("a" == "\n\t"))}
                _ = v37
            }
        }
        _ = v20
        _ = v21
    }
    v19 = (v19 & 1)
    v1 += v1
    for fi38 := 0; fi38 < 10; fi38++ {
        switch "a" {
        case "\n\t":
            var v39 string = "foo bar"
            _ = v39
        case "hello":
            if (false || (452 > (^(v19)))) {
                var v40 uint8 = (v1 - v1)
                _ = v40
            } else {
                v1 += 153
                if (!((-426 != -1) && true)) {
                    v41 := GenS5{F0: 1000, F1: v3, F2: -344}
                    _ = v41
                } else {
                    if true {
                        var v42 int8 = -1
                        _ = v42
                    }
                    if (v1 > 127) {
                        if (v3 < 548) {
                            var v43 int8 = 67
                            _ = v43
                        }
                        v1 -= (v1 % ((173) | 1))
                    }
                    var v44 int8 = -19
                    _ = v44
                }
                v1 = 42
            }
        case "":
            v45 := GenS6{F0: true, F1: v1, F2: (v19 == 9223372036854775807)}
            _ = v45
        }
        _ = fi38
    }
    if false {
        v46 := GenS7{F0: (v1 * v1), F1: 100, F2: 2}
        for fi47 := 0; fi47 < 14; fi47++ {
            _ = len(v2)
            v2 = append(v2, "foo bar")
            _ = fi47
        }
        _ = v46
    }
    switch 128 {
    case 42:
        v2 = append(v2, "foo bar")
    case 1:
        v48 := GenS6{F0: false, F1: (v1 << ((4294967295) & 63)), F2: (((0.1 + float64(1.5)) * (-2.25 + (1e10 / (3.14159 * float64(float32((0.1 - ((1.5 + 1.5) * -2.25)))))))) != (float64(-2.25) * (1.5 * float64(1.5))))}
        _ = v48
    }
    var v49 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v49
}
