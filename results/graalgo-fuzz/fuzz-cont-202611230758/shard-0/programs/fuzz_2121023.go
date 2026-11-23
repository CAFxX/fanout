package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int32 }
func genFn0() int32 {
    v6 := GenS1{F0: 127}
    v7 := []int{273, 32767, -1}
    var v8 int8 = -117
    var v9 float32 = 3.14159
    switch -702 {
    case -2147483648:
        v10 := GenS1{F0: v8}
        _ = v10
    case 4294967295:
        switch -9223372036854775808 {
        case -407:
            var v11 uint8 = 135
            _ = v11
        default:
            var v12 float64 = float64((v9 - -2.25))
            _ = v12
        }
    case -9223372036854775808:
        v9 = v9
    default:
        var v13 uint = 1
        _ = v13
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 uint = 255
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 string = ("" + "\n\t")
    var v5 float64 = float64((float32((0.1 - -2.25)) + (1.5 / 3.14159)))
    var v14 string = "a"
    v15 := GenS2{F0: 0, F1: 42}
    v5 = v5
    if (v4 <= "hello") {
        v16 := []int{-1, 9223372036854775807, 907, 127}
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 3; fi18++ {
            v4 = "foo bar"
            var v19 uint = (v1 / ((v1) | 1))
            var v20 int8 = -19
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 uint = (v1 | v1)
        if (((((((("" + v4) + "a") + ((v4 + "foo bar") + v14)) + v14) + ("\n\t" + "\n\t")) + v4) + "foo bar") == ("foo bar" + "hello")) {
            switch "" {
            case "\n\t":
                var v22 int16 = -32768
                _ = v22
            case "":
                if true {
                    v1 = (v21 / ((2) | 1))
                    var v23 int64 = 733
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        v4 = v14
                        v4 = v4
                        var v25 string = (v14 + "\n\t")
                        _ = fi24
                        _ = v25
                    }
                    v1 = (^(v1))
                    v1 = v1
                }
            case "a":
                v14 = v14
            }
        } else {
            switch (fi17 * fi17) {
            case 769:
                v26 := GenS3{F0: 0, F1: 1, F2: -43}
                _ = v26
            case -2147483648:
                delete(v2, 4294967295)
            case -1:
                v27 := GenS2{F0: -111, F1: -61}
                _ = v27
            default:
                v4 = (("foo bar" + "hello") + "a")
            }
            var v28 int = (^(fi17))
            _ = v28
        }
        _ = fi17
        _ = v21
    }
    var v29 int8 = 42
    if (!("foo bar" <= v14)) {
        v14 = ("foo bar" + v4)
    } else {
        if false {
            v3[uint32(v1)] = true
            for fi30 := 0; fi30 < 3; fi30++ {
                switch (fi30 ^ (fi30 - fi30)) {
                case 32768:
                    if false {
                        if true {
                            v31 := GenS4{F0: 1}
                            v32 := GenS4{F0: 42}
                            for fi33 := 0; fi33 < 17; fi33++ {
                                v34 := GenS1{F0: 0}
                                _ = fi33
                                _ = v34
                            }
                            _ = v31
                            _ = v32
                        }
                    } else {
                        v35 := GenS1{F0: (v29 / ((v29) | 1))}
                        v29 = 1
                        var v36 uint32 = uint32(v1)
                        _ = v35
                        _ = v36
                    }
                }
                if (49 < 42) {
                    v37 := GenS5{F0: 156, F1: genFn0()}
                    var v38 float32 = 1.5
                    for fi39 := 0; fi39 < 12; fi39++ {
                        var v40 float32 = v38
                        _ = fi39
                        _ = v40
                    }
                    _ = v37
                    _ = v38
                } else {
                    v5 = (((1.5 - v5) * (v5 + 0.0)) * v5)
                    if (0 < 215) {
                        var v41 int64 = 42
                        fmt.Println(v4, v4)
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 13; fi42++ {
                            v29 = (-(v29))
                            var v43 string = v14
                            switch fi30 {
                            case 32768:
                                var v44 uint32 = 566
                                _ = v44
                            default:
                                var v45 bool = (3.14159 < (0.0 / 3.14159))
                                fmt.Println(v43)
                                switch (v14 + v43) {
                                case "":
                                    var v46 uint = (v1 << ((v1) & 63))
                                    _ = v46
                                default:
                                    var v47 int64 = -9223372036854775808
                                    v45 = true
                                    var v48 int64 = v47
                                    fmt.Println(v43, v14)
                                    _ = v47
                                    _ = v48
                                }
                                _ = v45
                            }
                            _ = fi42
                            _ = v43
                        }
                    }
                    switch 4294967295 {
                    case -2147483648:
                        switch (fi30 ^ fi30) {
                        case -888:
                            for fi49 := 0; fi49 < 20; fi49++ {
                                var v50 uint8 = 137
                                var v51 int8 = v29
                                v1 -= 255
                                v51 -= 0
                                v29 = v29
                                _ = fi49
                                _ = v50
                                _ = v51
                            }
                        case 32767:
                            for fi52 := 0; fi52 < 7; fi52++ {
                                v5 = float64(1.5)
                                var v53 uint16 = 859
                                v29 = (v29 | -17)
                                _ = fi52
                                _ = v53
                            }
                        case 42:
                            var v54 int16 = 32767
                            _ = v54
                        default:
                            if (genFn0() != -1) {
                                for fi55 := 0; fi55 < 20; fi55++ {
                                    var v56 int16 = 0
                                    if (false || true) {
                                        v56 = v56
                                        _ = len(v2)
                                    }
                                    if false {
                                        fmt.Println(fi30)
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                                var v57 float64 = v5
                                switch v4 {
                                case "":
                                    switch v4 {
                                    case "hello":
                                        var v58 int32 = 0
                                        _ = v58
                                    case "foo bar":
                                        for fi59 := 0; fi59 < 4; fi59++ {
                                            var v60 uint16 = 0
                                            fmt.Println(v57)
                                            v14 = v4
                                            _ = fi59
                                            _ = v60
                                        }
                                    case "\n\t":
                                        fmt.Println(v5, v4, v14)
                                    default:
                                        v61 := GenS1{F0: -1}
                                        _ = v61
                                    }
                                case "\n\t":
                                    v62 := []int{(fi30 + 4294967295), -9223372036854775808}
                                    _ = v62
                                }
                                _ = v57
                            }
                        }
                    case 0:
                        v1 += 42
                    }
                }
                _ = len(v3)
                _ = fi30
            }
        } else {
            for fi63 := 0; fi63 < 17; fi63++ {
                v64 := v3[28]
                var v65 int64 = -276
                var v66 int = (fi63 | (fi63 + 128))
                _ = fi63
                _ = v64
                _ = v65
                _ = v66
            }
            var v67 uint = v1
            switch (v4 + v4) {
            case "\n\t":
                var v68 uint = (v67 >> (((v67 % ((v67) | 1))) & 63))
                _ = v68
            case "a":
                v29 -= v29
            case "":
                v14 = v4
            }
            _ = v67
        }
    }
    v4 = ("" + ((v4 + "a") + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
