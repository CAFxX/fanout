package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 int8; F2 int8 }

func main() {
    var v1 string = "hello"
    var v2 int16 = 0
    var v3 float32 = 0.1
    v4 := []int{5}
    var v5 uint8 = 127
    var v6 uint32 = 4294967295
    var v7 uint = 312
    fmt.Println(v2, v5)
    v4 = append(v4, 128)
    if true {
        for fi8 := 0; fi8 < 14; fi8++ {
            if true {
                var v9 float32 = v3
                var v10 int8 = 100
                _ = v10
                _ = v9
            }
            if false {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v7 += v7
                    if true {
                        switch (v1 + v1) {
                        case "foo bar":
                            var v12 string = (v1 + v1)
                            _ = v12
                        case "a":
                            switch fi11 {
                            case -1:
                                if false {
                                    if (1 < 237) {
                                        if ((v1 + v1) <= ("" + ("" + v1))) {
                                            fmt.Println(v1, v7, fi8)
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v13 int32 = 42
                                            switch (-(fi8)) {
                                            case -1:
                                                var v14 float64 = (3.14159 - (-2.25 * -2.25))
                                                _ = v14
                                            default:
                                                if (!false) {
                                                    if (1 == 32767) {
                                                        var v15 uint32 = (^(v6))
                                                        if (!false) {
                                                            if (0.1 >= 1e10) {
                                                                if len(v4) > 0 {
                                                                    v4[0] = fi8
                                                                }
                                                                v13 += (v13 | (v13 & (v13 / ((v13) | 1))))
                                                            } else {
                                                                v3 -= v3
                                                                if false {
                                                                    var v16 uint = v7
                                                                    _ = v16
                                                                }
                                                            }
                                                        }
                                                        switch (("" + ("a" + v1)) + "") {
                                                        case "hello":
                                                            if true {
                                                                switch v1 {
                                                                case "":
                                                                    v17 := GenS1{F0: v15}
                                                                    _ = v17
                                                                default:
                                                                    _ = len(v4)
                                                                }
                                                            } else {
                                                                if (true || (!(false && true))) {
                                                                    v18 := GenS1{F0: (v6 << (((v7 - (v7 ^ 974))) & 63))}
                                                                    _ = v18
                                                                }
                                                                switch v1 {
                                                                case "\n\t":
                                                                    if false {
                                                                        fmt.Println(fi11)
                                                                    } else {
                                                                        v6 = (v6 << (((v7 + (^(v7)))) & 63))
                                                                        var v19 uint32 = 2147483647
                                                                        v3 += (v3 - ((v3 - 1e10) - v3))
                                                                        _ = v19
                                                                    }
                                                                case "":
                                                                    var v20 int16 = (v2 - v2)
                                                                    _ = v20
                                                                case "hello":
                                                                    v21 := GenS2{F0: -9223372036854775808, F1: 1}
                                                                    _ = v21
                                                                }
                                                            }
                                                        case "\n\t":
                                                            var v22 uint64 = 2
                                                            _ = v22
                                                        case "a":
                                                            if len(v4) > 0 {
                                                                v4[0] = (fi8 | -2147483648)
                                                            }
                                                        default:
                                                            var v23 uint = (v7 & v7)
                                                            _ = v23
                                                        }
                                                        _ = v15
                                                    }
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v24 := GenS2{F0: -1, F1: -167}
                                                    _ = v24
                                                }
                                            }
                                            _ = v13
                                        }
                                    } else {
                                        v3 = ((((((float32(1e10) + v3) / (0.1 / (0.1 + v3))) / 0.1) / v3) * v3) - (((v3 / 1.5) - ((1e10 - ((3.14159 / ((-2.25 / 1.5) + v3)) / v3)) / v3)) + (1e10 - v3)))
                                    }
                                } else {
                                    if (false || (false && (v1 >= "a"))) {
                                        v25 := GenS3{F0: -329, F1: 0}
                                        _ = v25
                                    }
                                    var v26 int16 = (^(v2))
                                    _ = v26
                                }
                            case 4294967295:
                                v7 += v7
                            default:
                                var v27 int32 = 1
                                _ = v27
                            }
                            var v28 int32 = 0
                            if true {
                                v3 -= v3
                                _ = len(v4)
                            }
                            _ = v28
                        }
                        fmt.Println(v1, v3)
                        v4 = append(v4, fi11)
                        v4 = append(v4, 32768)
                    } else {
                        var v29 int = fi8
                        v30 := GenS2{F0: 9223372036854775807, F1: 1}
                        var v31 int64 = -1
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    if len(v4) > 0 {
                        v4[0] = fi11
                    }
                    _ = fi11
                }
            }
            var v32 int = (-(fi8))
            _ = fi8
            _ = v32
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    switch v1 {
    case "":
        v33 := GenS2{F0: int64(v5), F1: 9223372036854775807}
        _ = v33
    default:
        switch 9223372036854775807 {
        case 128:
            switch 4294967295 {
            case 4294967295:
                for fi34 := 0; fi34 < 7; fi34++ {
                    v35 := GenS3{F0: 42, F1: 0}
                    _ = fi34
                    _ = v35
                }
            default:
                if false {
                    _ = len(v4)
                    v36 := GenS4{F0: 1}
                    _ = v36
                } else {
                    switch -2147483648 {
                    case 2147483647:
                        switch v1 {
                        case "":
                            var v37 int64 = 1
                            _ = v37
                        case "hello":
                            for fi38 := 0; fi38 < 8; fi38++ {
                                var v39 uint = v7
                                v7 = v39
                                _ = fi38
                                _ = v39
                            }
                        default:
                            v40 := GenS1{F0: v6}
                            _ = v40
                        }
                    case 42:
                        fmt.Println(v2)
                    case 128:
                        v7 -= (v7 & uint(v6))
                    }
                }
            }
        default:
            v41 := GenS1{F0: 42}
            _ = v41
        }
    }
    v3 += (v3 * (1e10 + (v3 * 1e10)))
    v5 -= v5
    fmt.Println(v7, v5, v6)
    v42 := GenS5{F0: -420}
    v43 := GenS6{F0: 0, F1: -128, F2: -128}
    for fi44 := 0; fi44 < 3; fi44++ {
        var v45 uint32 = 77
        v46 := GenS2{F0: -192, F1: 42}
        _ = fi44
        _ = v45
        _ = v46
    }
    v1 = (v1 + ((((v1 + v1) + (v1 + v1)) + (v1 + "foo bar")) + (v1 + "")))
    v7 += (v7 + 2)
    var v47 string = (v1 + "foo bar")
    var v48 uint = (v7 | (v7 | (v7 / ((v7) | 1))))
    v48 -= 1
    var v49 float32 = v3
    var v50 uint32 = uint32(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
