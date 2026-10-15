package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
type GenS3 struct { F0 int8; F1 int64; F2 int8 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 908
    var v3 float32 = 1.5
    var v4 uint16 = 42
    var v5 int32 = 0
    if (true && true) {
        for fi6 := 0; fi6 < 7; fi6++ {
            fmt.Println(v2, v3)
            _ = fi6
        }
    } else {
        if ((!(true || (v4 == v4))) && true) {
            switch "a" {
            case "a":
                v5 -= v5
            default:
                v7 := []int{-223, 2147483647, 674}
                _ = v7
            }
            for fi8 := 0; fi8 < 13; fi8++ {
                if (!(3.14159 != -2.25)) {
                    v9 := []int{(fi8 & 9223372036854775807), 340, int(v2)}
                    v5 = v5
                    switch "" {
                    case "foo bar":
                        var v10 uint32 = 705
                        _ = v10
                    }
                    _ = v9
                }
                var v11 uint = 18446744073709551615
                switch fi8 {
                case 127:
                    if (0 >= (^(fi8))) {
                        v2 += -32768
                    }
                case -282:
                    fmt.Println(v2, v3)
                case 976:
                    switch fi8 {
                    case -2147483648:
                        v12 := GenS1{F0: uint64(v11)}
                        _ = v12
                    case 9223372036854775807:
                        v13 := GenS2{F0: 42, F1: 42, F2: 255}
                        _ = v13
                    case 32768:
                        var v14 uint64 = 407
                        _ = v14
                    default:
                        fmt.Println(v4)
                    }
                default:
                    for fi15 := 0; fi15 < 14; fi15++ {
                        if (668 >= (^(v4))) {
                            v16 := v1[367]
                            v17 := GenS3{F0: 100, F1: 42, F2: 42}
                            _ = v16
                            _ = v17
                        }
                        fmt.Println(fi8, v4, v11)
                        var v18 float64 = (1e10 + -2.25)
                        _ = fi15
                        _ = v18
                    }
                }
                _ = fi8
                _ = v11
            }
            if (42 == 2) {
                var v19 int32 = 0
                if (!false) {
                    v19 = v19
                    for fi20 := 0; fi20 < 20; fi20++ {
                        if (1 > 678) {
                            for fi21 := 0; fi21 < 7; fi21++ {
                                var v22 uint64 = 42
                                switch "a" {
                                case "\n\t":
                                    switch 1 {
                                    case 0:
                                        v23 := GenS4{F0: 551, F1: 42}
                                        _ = v23
                                    }
                                default:
                                    v2 = (v2 % (((^(v2))) | 1))
                                }
                                _ = fi21
                                _ = v22
                            }
                        }
                        v5 = 2147483647
                        _ = fi20
                    }
                    v19 = v19
                } else {
                    var v24 uint16 = v4
                    for fi25 := 0; fi25 < 3; fi25++ {
                        for fi26 := 0; fi26 < 2; fi26++ {
                            fmt.Println(v19)
                            v5 -= v5
                            _ = fi26
                        }
                        fmt.Println(v19)
                        var v27 int64 = -1
                        _ = fi25
                        _ = v27
                    }
                    _ = v24
                }
                _ = v19
            } else {
                var v28 int8 = 78
                fmt.Println(v28, v5)
                switch -616 {
                case 1:
                    v3 = 0.0
                default:
                    _ = len(v1)
                }
                _ = v28
            }
        } else {
            for fi29 := 0; fi29 < 11; fi29++ {
                for fi30 := 0; fi30 < 9; fi30++ {
                    v31 := GenS4{F0: 616, F1: 1}
                    v1[0] = (fi29 & fi29)
                    v2 -= (v2 << ((683) & 63))
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
        }
        if false {
            v5 -= (-(v5))
            var v32 string = "a"
            delete(v1, 2147483647)
            _ = v32
        }
        v3 = -2.25
    }
    fmt.Println(v4, v5, v4)
    v33 := GenS5{F0: -511, F1: (v4 - (v4 % (((v4 * 65535)) | 1)))}
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(v5, v3)
        fmt.Println(fi34, v5, fi34)
        _ = fi34
    }
    fmt.Println(v4)
    v35 := GenS1{F0: uint64(v4)}
    var v36 uint32 = 2
    var v37 uint = 300
    fmt.Println(v36)
    var v38 float64 = ((1e10 * (0.1 + 0.0)) + float64(v3))
    v39 := GenS1{F0: 4294967295}
    var v40 string = ("" + "hello")
    v37 -= uint(v36)
    if (!(!((!false) && true))) {
        var v41 uint64 = 0
        var v42 uint8 = 1
        _ = v41
        _ = v42
    } else {
        v3 -= v3
        for fi43 := 0; fi43 < 8; fi43++ {
            v40 = v40
            _ = fi43
        }
    }
    var v44 int64 = -26
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v44
    _ = v5
}
