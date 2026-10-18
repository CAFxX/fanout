package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 int8 }

func main() {
    var v1 bool = (-32768 >= 0)
    var v2 uint = 646
    v3 := []int{-2147483648}
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v4 string = ("a" + "hello")
    var v5 float32 = 3.14159
    v2 += (v2 / ((v2) | 1))
    fmt.Println(v4, v2)
    v1 = ((-1 < 2147483647) && v1)
    if v1 {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int32 = 42
            _ = fi6
            _ = v7
        }
        switch v4 {
        case "hello":
            switch "\n\t" {
            case "\n\t":
                v8 := GenS1{F0: 42, F1: 0, F2: 756}
                _ = v8
            case "":
                switch v4 {
                case "a":
                    var v9 float64 = (0.1 / 0.1)
                    _ = v9
                case "":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "\n\t":
                    v4 = ""
                }
            case "hello":
                var v10 string = ((v4 + (("\n\t" + v4) + v4)) + v4)
                _ = v10
            default:
                var v11 int32 = -715
                _ = v11
            }
        }
        if v1 {
            switch v4 {
            case "a":
                if v1 {
                    switch 2147483647 {
                    case 128:
                        v4 = v4
                    case -2147483648:
                        if ((42 > 9223372036854775807) || true) {
                            switch "" {
                            case "a":
                                v4 = v4
                            case "hello":
                                var v12 uint32 = 42
                                _ = v12
                            }
                            var v13 uint32 = 2
                            for fi14 := 0; fi14 < 2; fi14++ {
                                if ((!true) || (true && v1)) {
                                    if true {
                                        var v15 int8 = 100
                                        if (!((255 >= 255) && true)) {
                                            v16 := GenS2{F0: 9223372036854775807}
                                            _ = v16
                                        } else {
                                            v15 = (^(v15))
                                        }
                                        _ = v15
                                    } else {
                                        for fi17 := 0; fi17 < 7; fi17++ {
                                            if false {
                                                if (!(true && v1)) {
                                                    var v18 string = "hello"
                                                    _ = v18
                                                }
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = fi17
                                        }
                                        v19 := GenS3{F0: 42, F1: 238, F2: ((v4 != (v4 + v4)) || (!(-411 > 42)))}
                                        _ = v19
                                    }
                                    var v20 int16 = int16(v2)
                                    _ = v20
                                }
                                if (!(v4 != v4)) {
                                    for fi21 := 0; fi21 < 17; fi21++ {
                                        var v22 int8 = 1
                                        fmt.Println(v1)
                                        if v1 {
                                            var v23 uint = (^(v2))
                                            v24 := GenS4{F0: 738, F1: 65535}
                                            _ = v23
                                            _ = v24
                                        } else {
                                            fmt.Println(v1, fi21, fi21)
                                            fmt.Println(v5)
                                            if len(v3) > 0 {
                                                v3[0] = (fi14 & fi14)
                                            }
                                            v4 = (((v4 + v4) + "") + "foo bar")
                                        }
                                        var v25 float64 = (float64(1.5) / 3.14159)
                                        _ = fi21
                                        _ = v22
                                        _ = v25
                                    }
                                    v26 := GenS1{F0: 851, F1: 1, F2: 32767}
                                    _ = v26
                                }
                                _ = fi14
                            }
                            _ = v13
                        }
                    case -632:
                        v27 := GenS5{F0: v4, F1: v4}
                        _ = v27
                    default:
                        switch -1 {
                        case 2147483647:
                            fmt.Println(v1)
                        }
                    }
                    fmt.Println(v1)
                }
            default:
                switch v4 {
                case "\n\t":
                    var v28 uint16 = 2
                    _ = v28
                }
            }
        } else {
            _ = len(v3)
            if len(v3) > 0 {
                v3[0] = -1
            }
        }
    }
    var v29 float32 = v5
    var v30 string = v4
    switch 32767 {
    case 1:
        v5 = v5
    case 249:
        var v31 float64 = (1e10 - 1e10)
        _ = v31
    case -9223372036854775808:
        if v1 {
            if v1 {
                v3 = append(v3, -9223372036854775808)
                v30 = v4
                switch 32768 {
                case 42:
                    if false {
                        var v32 uint32 = 1
                        v33 := GenS4{F0: 309, F1: 42}
                        switch 128 {
                        case -255:
                            if (42 > 2) {
                                for fi34 := 0; fi34 < 5; fi34++ {
                                    var v35 int8 = 42
                                    _ = fi34
                                    _ = v35
                                }
                                v2 = v2
                            }
                        case 4294967295:
                            var v36 uint64 = 4
                            _ = v36
                        case 128:
                            v2 = v2
                        default:
                            switch -789 {
                            case 128:
                                var v37 float64 = ((-2.25 - -2.25) - 3.14159)
                                _ = v37
                            case 1:
                                for fi38 := 0; fi38 < 10; fi38++ {
                                    v1 = (!(float32(-2.25) >= v29))
                                    v32 -= v32
                                    _ = fi38
                                }
                            case 127:
                                if (false || v1) {
                                    fmt.Println(v4, v30)
                                }
                            default:
                                if (!false) {
                                    v29 = v5
                                    if true {
                                        v39 := GenS1{F0: 0, F1: 0, F2: 32767}
                                        v40 := GenS5{F0: v30, F1: v4}
                                        if false {
                                            v32 += (v32 & (v32 - 65535))
                                            fmt.Println(v5)
                                        }
                                        _ = v39
                                        _ = v40
                                    }
                                }
                            }
                        }
                        _ = v32
                        _ = v33
                    } else {
                        v29 -= (v29 - v29)
                    }
                case 490:
                    if (65535 <= 825) {
                        var v41 int32 = 2147483647
                        _ = v41
                    }
                case 48:
                    var v42 int16 = -32768
                    _ = v42
                default:
                    switch v30 {
                    case "hello":
                        var v43 int16 = -462
                        _ = v43
                    case "\n\t":
                        fmt.Println(v29, v4)
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
            } else {
                v44 := GenS6{F0: v30}
                switch 9223372036854775807 {
                case 42:
                    var v45 uint8 = 0
                    _ = v45
                default:
                    v46 := GenS2{F0: -1}
                    _ = v46
                }
                _ = v44
            }
        } else {
            var v47 uint = (v2 << (((v2 << (((v2 - (v2 - (v2 - v2)))) & 63))) & 63))
            v5 -= float32(float64(1e10))
            _ = v47
        }
    }
    if true {
        if v1 {
            var v48 uint8 = 171
            v49 := GenS7{F0: "a", F1: -1}
            _ = v48
            _ = v49
        } else {
            v4 = (v4 + (v4 + (v30 + "a")))
        }
        if ((!(false || v1)) || (v1 && v1)) {
            v30 = (v4 + v4)
            v5 -= float32(float64(-2.25))
        }
        if true {
            fmt.Println(v2, v30, v1)
        } else {
            for fi50 := 0; fi50 < 5; fi50++ {
                var v51 uint16 = 70
                v3 = append(v3, (^(fi50)))
                _ = fi50
                _ = v51
            }
        }
    } else {
        v30 = "\n\t"
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = (!(1 != 451))
    }
    var v52 uint = 215
    var v53 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v52
    _ = v53
}
