package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint8; F2 bool }
type GenS6 struct { F0 int64; F1 int32; F2 int8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    v3 := []uint32{2147483648}
    var v4 string = "a"
    v5 := make(map[int]bool)
    v6 := GenS1{F0: 597}
    var v7 uint32 = 2147483647
    if false {
        if false {
            fmt.Println(v7, v4)
        } else {
            fmt.Println(v7)
            fmt.Println(v7, v7)
            var v8 int8 = 42
            _ = v8
        }
        var v9 uint64 = 2
        _ = v9
    }
    if (224 != 255) {
        var v10 int = -798
        switch (^(v10)) {
        case 32767:
            v11 := GenS2{F0: (783 == 818)}
            _ = v11
        case 564:
            for fi12 := 0; fi12 < 16; fi12++ {
                switch v4 {
                case "\n\t":
                    var v13 int = v10
                    _ = v13
                case "hello":
                    if (!(2 <= 109)) {
                        v10 += (v10 | (v10 * (^(v10))))
                        v4 = "hello"
                    }
                case "":
                    v5[fi12] = false
                }
                fmt.Println(v4)
                _ = fi12
            }
        case 127:
            var v14 string = ""
            _ = v14
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            v4 = "\n\t"
            _ = fi15
        }
        _ = v10
    } else {
        var v16 int8 = 74
        _ = v16
    }
    v7 -= v7
    var v17 string = v4
    v17 = (("a" + v4) + (("" + "") + ("foo bar" + "hello")))
    var v18 bool = (61 == -1)
    fmt.Println(v17)
    v19 := GenS2{F0: (0.1 < 0.0)}
    switch 32767 {
    case 1:
        switch 32768 {
        case -163:
            if v18 {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v21 := GenS3{F0: 2147483647}
                    switch v17 {
                    case "a":
                        v7 = (v7 % ((v7) | 1))
                    }
                    _ = fi20
                    _ = v21
                }
                v17 = (v17 + v4)
            }
        case 42:
            fmt.Println(v17)
        case 9223372036854775807:
            var v22 int8 = -1
            _ = v22
        }
    }
    var v23 uint16 = 32767
    if v18 {
        if true {
            if (!v18) {
                v23 = (v23 ^ (v23 ^ (v23 << ((2) & 63))))
                for fi24 := 0; fi24 < 13; fi24++ {
                    if (!v18) {
                        var v25 uint = 0
                        _ = v25
                    }
                    var v26 uint16 = (v23 << ((42) & 63))
                    if false {
                        var v27 int16 = 1000
                        v28 := GenS4{F0: -128}
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS5{F0: v17, F1: 51, F2: ((v18 && true) && v18)}
                        if (v18 && true) {
                            if v18 {
                                v30 := GenS6{F0: 268, F1: 0, F2: -26}
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    var v32 int32 = -1
                                    var v33 uint8 = 30
                                    v34 := v5[int(v32)]
                                    var v35 int16 = 1000
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                                _ = v30
                            } else {
                                var v36 int16 = 0
                                v4 = ("\n\t" + (v17 + v17))
                                _ = v36
                            }
                            switch (^(fi24)) {
                            case 878:
                                v23 += 42
                            default:
                                v18 = v18
                            }
                        } else {
                            var v37 uint16 = v23
                            v26 = (v23 & 235)
                            _ = v37
                        }
                        _ = v29
                    }
                    _ = fi24
                    _ = v26
                }
                switch 32767 {
                case -458:
                    v7 = 2
                }
            }
        } else {
            fmt.Println(v4, v17, v7)
            switch 570 {
            case -914:
                _ = len(v3)
            case 0:
                var v38 uint64 = 42
                _ = v38
            }
            v39 := GenS5{F0: v4, F1: 0, F2: (128 >= -2147483648)}
            _ = v39
        }
        if (636 < 42) {
            v7 -= 65535
            if (!((v7 & (^(v7))) <= 56)) {
                var v40 int8 = -12
                _ = v40
            } else {
                switch -9223372036854775808 {
                case 0:
                    var v41 int8 = -116
                    _ = v41
                }
                switch 1 {
                case -55:
                    if v18 {
                        v4 = (((("\n\t" + v4) + "hello") + v4) + v17)
                        for fi42 := 0; fi42 < 12; fi42++ {
                            switch v4 {
                            case "a":
                                v17 = v17
                            case "":
                                v43 := GenS3{F0: 89}
                                _ = v43
                            }
                            v3 = append(v3, v7)
                            v7 = (v7 % (((v7 / ((613) | 1))) | 1))
                            _ = fi42
                        }
                        if v18 {
                            if (1.5 == (float64(1e10) * float64((((float32(-2.25) / 1.5) / 0.1) / 1.0)))) {
                                v44 := GenS7{F0: 30}
                                v7 = (v7 * (^(v7)))
                                _ = v44
                            } else {
                                fmt.Println(v23, v4)
                                for fi45 := 0; fi45 < 17; fi45++ {
                                    v7 += (^(v7))
                                    v46 := v2[(^(v7))]
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            switch 481 {
                            case -9223372036854775808:
                                var v47 float32 = ((1e10 * -2.25) - float32(3.14159))
                                _ = v47
                            }
                        } else {
                            if (!((uint8(v23) < 214) || v18)) {
                                for fi48 := 0; fi48 < 10; fi48++ {
                                    v49 := GenS3{F0: 192}
                                    _ = fi48
                                    _ = v49
                                }
                                v4 = "foo bar"
                                v5[32768] = v18
                            } else {
                                v5[2147483647] = true
                                switch 1 {
                                case 935:
                                    for fi50 := 0; fi50 < 17; fi50++ {
                                        var v51 string = "\n\t"
                                        v7 -= (v7 >> ((uint(v7)) & 63))
                                        _ = fi50
                                        _ = v51
                                    }
                                }
                            }
                            fmt.Println(v17, v18, v23)
                        }
                    }
                case 128:
                    var v52 uint64 = 776
                    _ = v52
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            v17 = (("foo bar" + v4) + "")
        }
        switch -23 {
        case 128:
            var v53 uint32 = (v7 ^ v7)
            _ = v53
        case 532:
            var v54 string = "foo bar"
            _ = v54
        }
    }
    for fi55 := 0; fi55 < 7; fi55++ {
        if (128 <= 128) {
            v56 := GenS8{F0: 87}
            _ = v56
        } else {
            var v57 int8 = 42
            v3 = append(v3, (^(v7)))
            v17 = v17
            _ = v57
        }
        _ = fi55
    }
    for fi58 := 0; fi58 < 19; fi58++ {
        var v59 float32 = 3.14159
        if ((v23 * v23) == v23) {
            var v60 float32 = v59
            fmt.Println(fi58, v18, v18)
            v17 = ""
            _ = v60
        } else {
            var v61 uint64 = 18446744073709551615
            _ = v61
        }
        _ = fi58
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
