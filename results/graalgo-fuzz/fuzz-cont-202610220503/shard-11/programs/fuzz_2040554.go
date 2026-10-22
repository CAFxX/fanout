package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int64 }
type GenS7 struct { F0 int8 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = v1
    var v3 uint = 42
    v4 := make(map[string]string)
    v5 := []int{}
    fmt.Println(v3, v1, v1)
    if true {
        v2 = v1
    } else {
        v6 := GenS1{F0: 255, F1: 814}
        if false {
            switch ((("" + ("a" + ("hello" + ""))) + ((("a" + "") + "a") + "")) + "a") {
            case "\n\t":
                v7 := GenS1{F0: 127, F1: -1}
                _ = v7
            case "hello":
                if (42 < 42) {
                    fmt.Println(v2, v3)
                    if false {
                        var v8 uint16 = 648
                        v2 = (^(v2))
                        v9 := GenS2{F0: 327, F1: 726}
                        _ = v8
                        _ = v9
                    }
                }
            case "a":
                v3 = (v3 / ((v3) | 1))
            }
        }
        var v10 float64 = -2.25
        _ = v10
        _ = v6
    }
    fmt.Println(v3)
    if ((v1 * -380) <= v2) {
        var v11 uint = v3
        v12 := GenS3{F0: 0}
        _ = v11
        _ = v12
    }
    switch "a" {
    case "hello":
        v1 -= 0
    case "":
        v1 = (-(v2))
    case "\n\t":
        fmt.Println(v1)
    default:
        switch "" {
        case "hello":
            fmt.Println(v1, v2, v1)
        case "foo bar":
            v2 -= -1
        case "\n\t":
            v2 = (v2 - 9223372036854775807)
        }
    }
    var v13 uint16 = 32767
    v13 += v13
    if true {
        if false {
            v2 += v2
        } else {
            fmt.Println(v2)
        }
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            fmt.Println(v2, v1)
            if false {
                v15 := GenS2{F0: fi14, F1: 2}
                switch fi14 {
                case 0:
                    var v16 uint32 = 0
                    _ = v16
                case 9223372036854775807:
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        for fi17 := 0; fi17 < 7; fi17++ {
                            fmt.Println(v3, v1, v1)
                            v2 += (v2 << (((v3 & 139)) & 63))
                            if false {
                                fmt.Println(v3, v2)
                                if false {
                                    v2 -= (v2 - 468)
                                    var v18 int16 = -1
                                    _ = v18
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = 0
                                    }
                                }
                            } else {
                                v3 += v3
                                v1 += v2
                            }
                            _ = fi17
                        }
                    } else {
                        v3 += (^(v3))
                        v2 += 0
                    }
                case -2147483648:
                    if (!true) {
                        switch (fi14 + -9223372036854775808) {
                        case -1:
                            v19 := GenS3{F0: 248}
                            _ = v19
                        case 2147483647:
                            v3 = 783
                        default:
                            for fi20 := 0; fi20 < 10; fi20++ {
                                var v21 int32 = -2147483648
                                var v22 uint = (v3 / (((^(v3))) | 1))
                                v2 = v1
                                if true {
                                    _ = len(v4)
                                    if false {
                                        fmt.Println(v21)
                                        v1 = -109
                                    }
                                }
                                if (0.0 >= float32((((float64((((float32(0.1) / 1e10) + -2.25) + 1e10)) - (1.5 / -2.25)) - 1e10) * 0.1))) {
                                    var v23 int8 = 32
                                    _ = v23
                                }
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        }
                        v2 += (v1 ^ (^(v1)))
                        var v24 int64 = v1
                        _ = v24
                    }
                default:
                    v25 := GenS4{F0: 0}
                    _ = v25
                }
                _ = v15
            } else {
                v5 = append(v5, (^(fi14)))
                var v26 bool = true
                v26 = (!((!(!true)) || v26))
                _ = v26
            }
            _ = fi14
        }
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 int = fi27
        _ = fi27
        _ = v28
    }
    v1 = 42
    fmt.Println(v1, v3, v13)
    fmt.Println(v2, v1, v3)
    if (((float64(3.14159) - (0.1 + (3.14159 - (((0.1 * (-2.25 / 3.14159)) + float64(0.0)) * 1e10)))) > 1.5) && true) {
        var v29 int16 = -1
        if false {
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := GenS5{F0: v29}
                if false {
                    if (fi30 < fi30) {
                        var v32 float32 = float32(float64((((0.1 - 3.14159) + ((float32(float64((1e10 / (float32(1e10) * 0.1)))) * float32(1e10)) - 1e10)) / 0.1)))
                        var v33 uint8 = 14
                        _ = v32
                        _ = v33
                    } else {
                        v5 = append(v5, (fi30 >> ((1) & 63)))
                        var v34 bool = true
                        _ = v34
                    }
                }
                _ = fi30
                _ = v31
            }
            for fi35 := 0; fi35 < 2; fi35++ {
                var v36 uint8 = uint8(v3)
                _ = fi35
                _ = v36
            }
        } else {
            v37 := GenS6{F0: 2147483647, F1: -9223372036854775808}
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            v3 -= 0
            _ = v37
        }
        _ = v29
    } else {
        v3 = 2
    }
    if (115 >= 999) {
        if (42 >= v13) {
            v38 := GenS5{F0: 1000}
            _ = v38
        }
        v39 := GenS2{F0: -9223372036854775808, F1: 4294967295}
        var v40 int16 = 1000
        _ = v39
        _ = v40
    } else {
        var v41 string = ""
        switch 0 {
        case 32767:
            v13 -= 255
        case 128:
            fmt.Println(v13, v41)
        }
        switch 9223372036854775807 {
        case 4294967295:
            switch 558 {
            case 32768:
                v42 := GenS4{F0: -1}
                _ = v42
            case 2147483647:
                for fi43 := 0; fi43 < 15; fi43++ {
                    v44 := GenS7{F0: 1}
                    _ = fi43
                    _ = v44
                }
            case 529:
                var v45 uint32 = 637
                _ = v45
            }
        }
        _ = v41
    }
    var v46 int8 = 127
    var v47 float64 = ((1e10 * 0.0) - (-2.25 - (3.14159 / -2.25)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
