package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }

func main() {
    var v1 uint = 255
    v2 := make(map[uint32]string)
    var v3 int = 32768
    var v4 string = "\n\t"
    v5 := GenS1{F0: -1, F1: 0, F2: (32767 != 32767)}
    v6 := []int{v3}
    fmt.Println(v3, v1)
    v3 = v3
    v1 = v1
    if ((!(false && true)) && (!(3.14159 >= -2.25))) {
        v4 = ""
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(v1)
            v4 = v4
            _ = fi7
        }
        fmt.Println(v3, v1)
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 104, F1: 1000, F2: ((^(v1)) <= v1)}
            if (991 >= 42) {
                v4 = (v4 + "\n\t")
                v1 -= 65535
            }
            _ = fi9
            _ = v10
        }
        if true {
            if (1.5 > (1e10 + float32(0.1))) {
                if (!((!true) && (!false))) {
                    var v11 int8 = -128
                    _ = v11
                } else {
                    for fi12 := 0; fi12 < 18; fi12++ {
                        switch 42 {
                        case 1:
                            var v13 uint64 = 354
                            _ = v13
                        case -2147483648:
                            v4 = "a"
                        }
                        _ = fi12
                    }
                    for fi14 := 0; fi14 < 18; fi14++ {
                        var v15 uint = 827
                        v3 = 0
                        var v16 bool = (-1 < -739)
                        v6 = append(v6, (fi8 | fi8))
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    v3 -= fi8
                }
                switch v3 {
                case 0:
                    v17 := GenS2{F0: 0}
                    _ = v17
                case 1:
                    if (465 != 42) {
                        for fi18 := 0; fi18 < 16; fi18++ {
                            var v19 float64 = float64((1.5 - float32((0.0 / 0.1))))
                            _ = fi18
                            _ = v19
                        }
                        switch 42 {
                        case 32768:
                            switch 128 {
                            case 1:
                                for fi20 := 0; fi20 < 1; fi20++ {
                                    if len(v6) > 0 {
                                        v6[0] = 2147483647
                                    }
                                    v3 += (v3 / ((fi20) | 1))
                                    fmt.Println(fi8, v4)
                                    v21 := v2[uint32(v1)]
                                    _ = fi20
                                    _ = v21
                                }
                            }
                        case 0:
                            v4 = "hello"
                        case 9223372036854775807:
                            switch -9223372036854775808 {
                            case -1:
                                v22 := GenS2{F0: 114}
                                _ = v22
                            case -5:
                                fmt.Println(v4, v3)
                            default:
                                v23 := v2[277]
                                _ = v23
                            }
                        }
                        fmt.Println(v1, fi8)
                    }
                case 128:
                    v24 := GenS1{F0: 146, F1: -1, F2: true}
                    _ = v24
                }
            }
            if len(v6) > 0 {
                v6[0] = (fi8 - (fi8 + (fi8 % (((fi8 | fi8)) | 1))))
            }
        } else {
            switch "" {
            case "a":
                delete(v2, 958)
            case "foo bar":
                var v25 int32 = -566
                _ = v25
            }
        }
        var v26 bool = (4294967295 >= (-(fi8)))
        _ = fi8
        _ = v26
    }
    v27 := v2[438]
    fmt.Println(v1, v3)
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if (v3 < (v3 + v3)) {
            var v28 int32 = 201
            _ = v28
        }
    } else {
        v3 += v3
        v3 -= (v3 % ((-585) | 1))
    }
    fmt.Println(v3, v27, v4)
    if len(v6) > 0 {
        v6[0] = v3
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        for fi30 := 0; fi30 < 20; fi30++ {
            if (638 < 0) {
                v31 := v2[42]
                v31 = v27
                v32 := GenS1{F0: 1, F1: 32767, F2: (true && (((!false) || (0 <= 1)) && true))}
                if (v1 != (v1 & v1)) {
                    if true {
                        var v33 uint32 = 1
                        var v34 float64 = float64((1.5 - float32(0.0)))
                        _ = v33
                        _ = v34
                    } else {
                        if true {
                            if false {
                                v35 := GenS3{F0: true}
                                _ = v35
                            } else {
                                v31 = (v31 + (v27 + v27))
                                v3 += (-(fi30))
                                var v36 float64 = (0.0 + 0.1)
                                _ = v36
                            }
                            var v37 uint16 = 255
                            var v38 int64 = 1
                            _ = v37
                            _ = v38
                        }
                        var v39 int8 = 100
                        _ = v39
                    }
                    var v40 int8 = 110
                    var v41 string = (v31 + "")
                    _ = v40
                    _ = v41
                }
                v4 = v4
                _ = v31
                _ = v32
            }
            _ = fi30
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
