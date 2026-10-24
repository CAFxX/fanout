package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 int64 }
func genFn0() uint8 {
    v5 := []int{972, 42}
    var v6 int = 32767
    v6 -= (-(v6))
    switch "\n\t" {
    case "\n\t":
        v6 = -2147483648
    case "a":
        var v7 float32 = 1.5
        _ = v7
    }
    fmt.Println(v6)
    _ = v5
    _ = v6
    return 2
}

func main() {
    v1 := []string{"hello", "", "a"}
    v2 := []string{}
    var v3 float64 = (0.0 * 1.5)
    v4 := make(map[int]uint32)
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 32767}
        switch ((("a" + "foo bar") + "") + ("foo bar" + ("foo bar" + ""))) {
        case "hello":
            for fi10 := 0; fi10 < 1; fi10++ {
                var v11 int8 = 0
                _ = fi10
                _ = v11
            }
        case "a":
            v3 += (float64(3.14159) / (3.14159 + 0.1))
        default:
            for fi12 := 0; fi12 < 4; fi12++ {
                fmt.Println(fi8, v3, v3)
                _ = fi12
            }
        }
        _ = fi8
        _ = v9
    }
    v4[127] = 779
    var v13 uint16 = 32767
    var v14 int64 = 1
    v4[9223372036854775807] = 1
    var v15 uint32 = 65535
    v16 := GenS2{F0: 42, F1: v15, F2: 2}
    if ((false && (!(v15 == (v15 >> ((18446744073709551615) & 63))))) || false) {
        fmt.Println(v14, v3)
    }
    if (2 > 686) {
        v17 := GenS3{F0: -128}
        var v18 uint32 = (v15 * (v15 / ((v15) | 1)))
        v19 := GenS2{F0: int8(v14), F1: v18, F2: 402}
        _ = v17
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint8 = genFn0()
            v15 = (v15 + (v15 + v15))
            if true {
                v3 = v3
                if false {
                    if true {
                        var v22 uint32 = v15
                        v23 := GenS2{F0: 42, F1: v22, F2: 4294967295}
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 12; fi24++ {
                            if (-272 == 42) {
                                if (255 != 127) {
                                    var v25 uint8 = (^(v21))
                                    v25 += (^(v21))
                                    _ = v25
                                } else {
                                    var v26 float32 = 0.1
                                    var v27 int32 = 857
                                    _ = v26
                                    _ = v27
                                }
                                var v28 int64 = (v14 / ((v14) | 1))
                                if (true || (-1 == -721)) {
                                    if false {
                                        var v29 uint64 = 384
                                        _ = v29
                                    }
                                    switch ("hello" + (("foo bar" + "hello") + "a")) {
                                    case "a":
                                        var v30 int8 = -11
                                        _ = v30
                                    default:
                                        switch ("hello" + "\n\t") {
                                        case "":
                                            v31 := GenS4{F0: (v28 ^ (v14 << ((65535) & 63)))}
                                            _ = v31
                                        case "\n\t":
                                            switch "a" {
                                            case "hello":
                                                var v32 int = -656
                                                _ = v32
                                            default:
                                                var v33 int64 = 0
                                                _ = v33
                                            }
                                        }
                                    }
                                } else {
                                    v34 := GenS4{F0: (v14 + -132)}
                                    v3 = ((1.5 - v3) - float64(float32(3.14159)))
                                    _ = v34
                                }
                                v35 := GenS5{F0: v15}
                                v36 := GenS1{F0: -32768}
                                _ = v28
                                _ = v35
                                _ = v36
                            } else {
                                _ = len(v4)
                            }
                            v37 := GenS3{F0: -128}
                            v38 := GenS2{F0: -113, F1: (v15 / ((uint32(v13)) | 1)), F2: 2}
                            v39 := GenS1{F0: 892}
                            if true {
                                var v40 uint8 = (v21 + v21)
                                var v41 uint32 = v15
                                _ = v40
                                _ = v41
                            }
                            _ = fi24
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        switch fi20 {
                        case 1:
                            for fi42 := 0; fi42 < 19; fi42++ {
                                v15 = v15
                                _ = fi42
                            }
                        default:
                            var v43 int16 = -811
                            _ = v43
                        }
                        if false {
                            v44 := GenS5{F0: 4294967295}
                            fmt.Println(v14)
                            switch (((("\n\t" + "") + "foo bar") + "") + "") {
                            case "hello":
                                v45 := GenS6{F0: 65535, F1: (v15 ^ (v15 << ((1) & 63)))}
                                _ = v45
                            }
                            _ = v44
                        } else {
                            for fi46 := 0; fi46 < 8; fi46++ {
                                fmt.Println(v21)
                                var v47 uint32 = (v15 / ((42) | 1))
                                var v48 int = 42
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                            fmt.Println(v13)
                        }
                    }
                }
            }
            _ = fi20
            _ = v21
        }
    }
    v3 = ((v3 * float64(0.1)) / v3)
    switch 128 {
    case -9223372036854775808:
        if ("\n\t" != ((("" + "a") + "\n\t") + "foo bar")) {
            if (-2147483648 < -9223372036854775808) {
                if (v14 <= 42) {
                    for fi49 := 0; fi49 < 20; fi49++ {
                        for fi50 := 0; fi50 < 10; fi50++ {
                            var v51 uint8 = 133
                            _ = fi50
                            _ = v51
                        }
                        v13 -= 42
                        _ = fi49
                    }
                } else {
                    v52 := GenS1{F0: 1000}
                    v53 := GenS1{F0: 32767}
                    _ = v52
                    _ = v53
                }
                if false {
                    if true {
                        v13 += v13
                        if (!false) {
                            if (0 > 65535) {
                                fmt.Println(v13, v14)
                                v3 = v3
                                v3 = float64(0.0)
                            } else {
                                v54 := GenS3{F0: 127}
                                _ = v54
                            }
                            for fi55 := 0; fi55 < 18; fi55++ {
                                for fi56 := 0; fi56 < 15; fi56++ {
                                    var v57 uint32 = 2147483647
                                    fmt.Println(v15)
                                    _ = fi56
                                    _ = v57
                                }
                                switch fi55 {
                                case 32767:
                                    var v58 int = fi55
                                    _ = v58
                                case 475:
                                    for fi59 := 0; fi59 < 3; fi59++ {
                                        v60 := GenS7{F0: 1000, F1: 541}
                                        _ = fi59
                                        _ = v60
                                    }
                                default:
                                    if (float64(0.0) == float64(0.1)) {
                                        v3 = float64(1e10)
                                        _ = len(v2)
                                    }
                                }
                                v15 = v15
                                _ = fi55
                            }
                            var v61 uint32 = 211
                            _ = v61
                        }
                        delete(v4, 610)
                    } else {
                        var v62 int = -9223372036854775808
                        _ = v62
                    }
                }
            } else {
                var v63 float64 = (1.5 + v3)
                switch -2147483648 {
                case -1:
                    _ = len(v4)
                }
                _ = v63
            }
        } else {
            v64 := GenS1{F0: -833}
            for fi65 := 0; fi65 < 10; fi65++ {
                v3 -= 0.0
                _ = fi65
            }
            v14 -= v14
            _ = v64
        }
    case 0:
        v66 := GenS8{F0: -9223372036854775808}
        _ = v66
    default:
        if ((v15 << ((0) & 63)) < (v15 % (((v15 - 2)) | 1))) {
            switch "\n\t" {
            case "":
                var v67 int64 = (v14 * (v14 - (v14 * (v14 >> ((42) & 63)))))
                _ = v67
            case "foo bar":
                v15 = (v15 % ((2147483648) | 1))
            case "hello":
                v68 := GenS1{F0: -32768}
                _ = v68
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v69 int = -9223372036854775808
            _ = v69
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
