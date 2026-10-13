package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 int32 }

func main() {
    var v1 int8 = -100
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    var v4 float64 = 0.1
    v5 := make(map[string]string)
    v6 := GenS1{F0: 204, F1: -1}
    var v7 int64 = -1
    for fi8 := 0; fi8 < 1; fi8++ {
        if (!true) {
            v9 := GenS2{F0: 42, F1: 9223372036854775807}
            var v10 int = fi8
            v11 := GenS1{F0: 42, F1: 1000}
            _ = v10
            _ = v11
            _ = v9
        } else {
            v1 -= (v1 + (^(v1)))
            v12 := GenS3{F0: 4294967295}
            _ = v12
        }
        v4 -= ((v4 / (((v4 + v4) + (((0.0 / v4) - v4) + 1e10)) + 1e10)) / float64(((-2.25 * (((3.14159 / (3.14159 * 1e10)) - 1e10) / 1.5)) - (1e10 - 0.0))))
        _ = fi8
    }
    var v13 int8 = v1
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := GenS4{F0: (v13 + v13)}
        _ = fi14
        _ = v15
    }
    var v16 int16 = 1000
    v16 = 42
    var v17 uint64 = 799
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            if true {
                v13 = v1
                if false {
                    var v20 uint32 = 1
                    var v21 int64 = (v7 + -8)
                    var v22 int = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                switch ("hello" + "a") {
                case "a":
                    fmt.Println(v7, fi19)
                case "hello":
                    if (fi18 < fi19) {
                        if ((false || true) || (false || (-2147483648 == -1))) {
                            var v23 float32 = (((0.1 * (0.0 / 0.1)) / 1.0) / (-2.25 * float32(1.5)))
                            var v24 uint8 = 42
                            var v25 uint32 = 42
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    } else {
                        if true {
                            switch fi19 {
                            case 675:
                                var v26 int32 = 771
                                _ = v26
                            }
                            if false {
                                var v27 string = ""
                                _ = v27
                            } else {
                                var v28 uint8 = 110
                                switch fi19 {
                                case -2147483648:
                                    if (!true) {
                                        var v29 float32 = (1e10 * (1.5 + 1.5))
                                        var v30 uint16 = 65535
                                        _ = v29
                                        _ = v30
                                    }
                                    switch (^(fi19)) {
                                    case 33:
                                        v31 := []int{int(v28)}
                                        _ = v31
                                    default:
                                        fmt.Println(v13, v13)
                                    }
                                }
                                switch fi18 {
                                case -93:
                                    _ = len(v5)
                                }
                                v4 = (1e10 + (v4 / v4))
                                var v32 int = 127
                                var v33 uint = 4294967295
                                if false {
                                    var v34 float64 = v4
                                    _ = len(v3)
                                    v35 := GenS1{F0: (v17 / ((42) | 1)), F1: (v16 << ((uint(v17)) & 63))}
                                    if (true && false) {
                                        v33 += (v33 >> (((v33 + v33)) & 63))
                                        v36 := GenS3{F0: (v33 & 2)}
                                        _ = v36
                                    } else {
                                        fmt.Println(v33, v34)
                                        var v37 bool = true
                                        if (v37 && ((v37 || true) && ((^(v17)) >= uint64(v28)))) {
                                            v38 := GenS5{F0: 376}
                                            var v39 int8 = -128
                                            v5[""] = "\n\t"
                                            var v40 uint8 = 1
                                            var v41 bool = v37
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = v37
                                    }
                                    _ = v34
                                    _ = v35
                                }
                                _ = v28
                                _ = v32
                                _ = v33
                            }
                            switch fi19 {
                            case 4294967295:
                                var v42 int64 = v7
                                _ = v42
                            case -825:
                                v5[""] = ("a" + ("a" + ("a" + "")))
                            default:
                                fmt.Println(v17, v17, v7)
                            }
                        } else {
                            v43 := []int{(fi19 % ((2147483647) | 1)), int(v17)}
                            v17 -= v17
                            switch "a" {
                            case "a":
                                switch -327 {
                                case 4294967295:
                                    if true {
                                        var v44 uint16 = uint16(v17)
                                        _ = v44
                                    }
                                case 342:
                                    v13 -= v1
                                }
                            case "foo bar":
                                switch "a" {
                                case "\n\t":
                                    v17 = 0
                                case "foo bar":
                                    switch "hello" {
                                    case "":
                                        v4 = ((v4 * float64(float32((v4 * (v4 * (v4 - 3.14159)))))) / v4)
                                    default:
                                        v4 = float64(0.0)
                                    }
                                default:
                                    v16 = (v16 % ((1) | 1))
                                }
                            }
                            _ = v43
                        }
                        v45 := GenS3{F0: 65535}
                        _ = v45
                    }
                case "\n\t":
                    var v46 int32 = 734
                    _ = v46
                }
            }
            _ = fi19
        }
        _ = fi18
    }
    switch 32767 {
    case -492:
        fmt.Println(v13, v13, v16)
    case 32768:
        v16 = 556
    case -242:
        v17 = (v17 * v17)
    }
    _ = len(v3)
    if false {
        var v47 float32 = (0.1 * 0.1)
        v4 = v4
        _ = v47
    }
    fmt.Println(v7, v4)
    var v48 int8 = (v1 % ((v1) | 1))
    for fi49 := 0; fi49 < 15; fi49++ {
        v50 := GenS6{F0: 157, F1: 0}
        _ = fi49
        _ = v50
    }
    var v51 bool = true
    var v52 uint32 = 65535
    var v53 bool = v51
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
    _ = v7
}
