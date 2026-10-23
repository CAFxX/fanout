package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 bool; F1 int8; F2 int16 }
type GenS4 struct { F0 uint; F1 int8; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS7 struct { F0 uint64 }
func genFn0() uint32 {
    v5 := GenS1{F0: -25, F1: 18446744073709551615, F2: 2147483647}
    var v6 int16 = 1000
    if true {
        fmt.Println(v6)
        var v7 int32 = 1
        _ = v7
    }
    _ = v5
    _ = v6
    return 65535
}
func genFn1() int16 {
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint32 = 4294967295
        if (!true) {
            var v10 uint16 = 0
            switch 127 {
            case 0:
                v11 := GenS2{F0: 250, F1: 42}
                _ = v11
            case -2147483648:
                var v12 int = (fi8 % (((fi8 / ((fi8) | 1))) | 1))
                _ = v12
            }
            var v13 uint = 1
            _ = v10
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint16 = 267
        _ = fi14
        _ = v15
    }
    return -32768
}

func main() {
    var v1 int64 = -1
    var v2 int16 = 0
    var v3 float32 = 0.1
    var v4 float64 = 1e10
    v16 := []int{}
    if true {
        fmt.Println(v1, v4)
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 int = 2147483647
            _ = fi17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 19; fi19++ {
            var v20 uint64 = 0
            var v21 float64 = v4
            var v22 float32 = v3
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    }
    v1 = -549
    if true {
        v3 -= v3
        v3 = v3
    }
    switch 128 {
    case 2147483647:
        v23 := GenS1{F0: -1, F1: 42, F2: 42}
        _ = v23
    case 128:
        v24 := GenS1{F0: -1, F1: 4294967295, F2: 2147483647}
        _ = v24
    case -1:
        for fi25 := 0; fi25 < 12; fi25++ {
            v26 := GenS3{F0: true, F1: -128, F2: genFn1()}
            switch (fi25 + -2147483648) {
            case 2147483647:
                switch "" {
                case "a":
                    v1 = v1
                case "":
                    for fi27 := 0; fi27 < 13; fi27++ {
                        v28 := GenS4{F0: 0, F1: 42, F2: -128}
                        if ((v2 / ((32767) | 1)) != -93) {
                            v2 = v2
                            var v29 uint8 = 0
                            _ = v29
                        }
                        _ = fi27
                        _ = v28
                    }
                default:
                    for fi30 := 0; fi30 < 8; fi30++ {
                        if (1 >= -53) {
                            var v31 string = ("a" + ("a" + "hello"))
                            var v32 int16 = (v2 >> ((882) & 63))
                            _ = v31
                            _ = v32
                        } else {
                            switch (fi25 % ((-1) | 1)) {
                            case 0:
                                var v33 float64 = 3.14159
                                _ = v33
                            case 9223372036854775807:
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            case 1:
                                switch (^(fi25)) {
                                case 0:
                                    var v34 uint32 = 2
                                    var v35 uint16 = 1
                                    _ = v34
                                    _ = v35
                                default:
                                    _ = len(v16)
                                }
                            default:
                                if (false || true) {
                                    var v36 bool = true
                                    var v37 int16 = v2
                                    if len(v16) > 0 {
                                        v16[0] = fi30
                                    }
                                    _ = v36
                                    _ = v37
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(v2)
                                        case "":
                                            switch "" {
                                            case "a":
                                                v38 := GenS5{F0: 32767}
                                                _ = v38
                                            }
                                        default:
                                            switch -9223372036854775808 {
                                            case -224:
                                                var v39 float32 = v3
                                                _ = v39
                                            default:
                                                if true {
                                                    if false {
                                                        var v40 uint64 = 1
                                                        var v41 float32 = 3.14159
                                                        fmt.Println(v3, v4, v1)
                                                        var v42 uint64 = (v40 + 0)
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                    } else {
                                                        v43 := GenS6{F0: -2147483648, F1: 1, F2: 0}
                                                        var v44 uint32 = 1
                                                        var v45 uint64 = 366
                                                        var v46 uint16 = 0
                                                        _ = v43
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                    v16 = append(v16, (fi25 << ((312) & 63)))
                                                }
                                            }
                                            var v47 int64 = -1
                                            var v48 uint16 = 255
                                            _ = v47
                                            _ = v48
                                        }
                                    default:
                                        v49 := GenS7{F0: 654}
                                        _ = v49
                                    }
                                    fmt.Println(v2)
                                    v4 = v4
                                }
                            }
                            var v50 float64 = 3.14159
                            if true {
                                if true {
                                    if true {
                                        v50 = float64((((v3 * 0.1) / v3) / 1.5))
                                    }
                                }
                                var v51 int64 = 908
                                v1 = v51
                                _ = v51
                            }
                            v50 += float64((3.14159 - 0.0))
                            var v52 uint = 65535
                            _ = v50
                            _ = v52
                        }
                        var v53 int64 = 42
                        switch fi25 {
                        case 612:
                            v53 = -1
                        case -632:
                            v2 = 1
                        default:
                            var v54 int = (fi25 - (fi25 >> ((18446744073709551615) & 63)))
                            _ = v54
                        }
                        _ = fi30
                        _ = v53
                    }
                }
            }
            _ = fi25
            _ = v26
        }
    }
    var v55 uint16 = 32767
    _ = len(v16)
    switch ((("\n\t" + "a") + "hello") + "") {
    case "\n\t":
        fmt.Println(v1)
    case "foo bar":
        v4 = v4
    }
    for fi56 := 0; fi56 < 8; fi56++ {
        for fi57 := 0; fi57 < 19; fi57++ {
            var v58 uint16 = 171
            v59 := GenS7{F0: 18446744073709551615}
            _ = fi57
            _ = v58
            _ = v59
        }
        v2 = -961
        if false {
            if (!false) {
                var v60 uint32 = 0
                _ = v60
            }
            var v61 int = 42
            _ = v61
        }
        _ = fi56
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v55
}
