package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint64; F2 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16; F1 uint16 }
func genFn0() uint64 {
    v5 := GenS1{F0: 401}
    var v6 bool = true
    var v7 uint = 42
    v6 = v6
    switch (("hello" + "\n\t") + (((("a" + "a") + "") + "\n\t") + (("hello" + ("hello" + "hello")) + ""))) {
    case "foo bar":
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS2{F0: (v7 & v7)}
            _ = fi8
            _ = v9
        }
    default:
        fmt.Println(v7)
    }
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint = 643
    var v2 int16 = 1
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    v1 = v1
    v1 = (v1 - (v1 + 42))
    var v10 float32 = (((3.14159 * 0.1) / 1.5) / (3.14159 / 0.1))
    for fi11 := 0; fi11 < 3; fi11++ {
        _ = len(v4)
        fmt.Println(v10, fi11, fi11)
        if (int32(v1) == -1) {
            var v12 int8 = 90
            _ = v12
        } else {
            v4[2147483648] = uint32(v1)
            _ = len(v4)
        }
        _ = fi11
    }
    v13 := GenS2{F0: 0}
    if true {
        if true {
            v1 = v1
            switch "foo bar" {
            case "hello":
                for fi14 := 0; fi14 < 11; fi14++ {
                    fmt.Println(v2)
                    _ = fi14
                }
            case "":
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 float32 = (v10 / (v10 / 0.1))
                    if (true || false) {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            if (0 < genFn0()) {
                                v18 := GenS2{F0: (v1 / ((v1) | 1))}
                                var v19 int8 = 1
                                _ = v18
                                _ = v19
                            } else {
                                v1 = v1
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            fmt.Println(v2)
                            _ = fi17
                        }
                        var v20 uint32 = 42
                        v16 -= v16
                        _ = v20
                    } else {
                        fmt.Println(v10, fi15)
                        v21 := GenS1{F0: 255}
                        _ = v21
                    }
                    var v22 uint8 = 255
                    _ = fi15
                    _ = v16
                    _ = v22
                }
            case "foo bar":
                v1 += v1
            default:
                var v23 uint64 = 1
                _ = v23
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                v2 += v2
                _ = fi24
            }
        } else {
            if (42 == 2) {
                v1 = (v1 ^ v1)
                v1 = v1
            } else {
                v4[42] = 2
                for fi25 := 0; fi25 < 10; fi25++ {
                    if ((42 >= 442) || false) {
                        v26 := GenS3{F0: (fi25 & 4294967295), F1: 2147483647, F2: 127}
                        v2 += v2
                        v2 = v2
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 15; fi27++ {
                        if (-200 != -1) {
                            delete(v4, 2147483648)
                        } else {
                            if true {
                                var v28 int16 = (v2 / ((v2) | 1))
                                _ = v28
                            } else {
                                if (1e10 == float64(v10)) {
                                    v1 -= (v1 << ((611) & 63))
                                }
                                v29 := GenS4{F0: v2, F1: 501, F2: 245}
                                switch (fi27 >> (((v1 | 743)) & 63)) {
                                case 127:
                                    switch (("\n\t" + ("hello" + ("" + ("a" + "\n\t")))) + "foo bar") {
                                    case "foo bar":
                                        fmt.Println(v2)
                                    case "":
                                        var v30 int8 = -62
                                        _ = v30
                                    }
                                case -2147483648:
                                    v2 += -32768
                                default:
                                    v2 -= (v2 * (-(v2)))
                                }
                                _ = v29
                            }
                            v2 = 42
                        }
                        _ = fi27
                    }
                    _ = fi25
                }
            }
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            v2 -= v2
            _ = len(v4)
            _ = fi31
        }
        for fi32 := 0; fi32 < 8; fi32++ {
            for fi33 := 0; fi33 < 17; fi33++ {
                if (983 >= 481) {
                    if true {
                        v10 = float32(((((-2.25 - (-2.25 + (0.0 + (float64((0.1 - float32(0.1))) * 1.5)))) + 0.0) + ((float64(v10) + ((1.5 * 0.1) / 1.5)) * 1e10)) / 1e10))
                    } else {
                        fmt.Println(fi32, fi32)
                        v34 := GenS5{F0: ("" + ((("foo bar" + "\n\t") + "foo bar") + ""))}
                        _ = v34
                    }
                } else {
                    fmt.Println(v2, fi33, v1)
                    if (v10 <= (v10 / (float32(3.14159) + (v10 + (float32(((1e10 - 0.0) - (((0.0 / 3.14159) - (0.1 / (1.5 - float64(v10)))) * 0.0))) / 1.0))))) {
                        if (true || (128 < 2)) {
                            switch -239 {
                            case 32767:
                                var v35 string = ""
                                _ = v35
                            case 384:
                                var v36 int = 31
                                _ = v36
                            case -650:
                                v1 = 1
                            }
                        }
                        v37 := GenS5{F0: "hello"}
                        if false {
                            var v38 uint16 = 0
                            v39 := GenS4{F0: int16(v1), F1: 320, F2: 0}
                            _ = v38
                            _ = v39
                        }
                        v2 = v2
                        fmt.Println(fi33, fi33, v1)
                        if true {
                            switch "" {
                            case "":
                                v40 := GenS6{F0: 2}
                                v10 = 1e10
                                v41 := GenS6{F0: 1}
                                _ = v40
                                _ = v41
                            default:
                                v2 = v2
                            }
                            v42 := GenS2{F0: v1}
                            _ = v42
                        } else {
                            if true {
                                _ = len(v3)
                            } else {
                                var v43 uint16 = 42
                                var v44 int = (^(fi33))
                                switch "\n\t" {
                                case "\n\t":
                                    v45 := GenS4{F0: 1, F1: 1, F2: 555}
                                    _ = v45
                                case "hello":
                                    v46 := GenS7{F0: 697, F1: (v43 >> (((v1 >> ((v1) & 63))) & 63))}
                                    _ = v46
                                }
                                _ = v43
                                _ = v44
                            }
                            _ = len(v4)
                            v47 := GenS1{F0: 997}
                            var v48 uint16 = uint16(v1)
                            v49 := GenS3{F0: 128, F1: 65535, F2: 128}
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        _ = v37
                    }
                }
                if (-9223372036854775808 >= 0) {
                    var v50 float32 = (v10 / (v10 + v10))
                    if false {
                        var v51 string = "hello"
                        _ = v51
                    } else {
                        fmt.Println(v2, v10)
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    }
                    if false {
                        v52 := GenS5{F0: ""}
                        v4[4294967295] = 330
                        _ = v52
                    } else {
                        v2 = v2
                        switch "hello" {
                        case "a":
                            var v53 uint8 = 0
                            _ = v53
                        }
                    }
                    _ = v50
                }
                var v54 float64 = ((0.0 * -2.25) - 3.14159)
                _ = fi33
                _ = v54
            }
            _ = fi32
        }
    } else {
        v1 = v1
    }
    for fi55 := 0; fi55 < 15; fi55++ {
        if false {
            v56 := GenS6{F0: 42}
            _ = v56
        }
        _ = fi55
    }
    var v57 string = "hello"
    var v58 uint16 = 32767
    var v59 uint32 = 42
    v4[v59] = (v59 / (((^(v59))) | 1))
    fmt.Println(v59, v2)
    v57 = (("hello" + (("hello" + v57) + "a")) + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v59
}
