package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 uint16; F2 int }
type GenS8 struct { F0 int16; F1 uint; F2 string }
type GenS9 struct { F0 uint16; F1 int; F2 string }
type GenS10 struct { F0 bool; F1 int64; F2 int8 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]int)
    v3 := []int{1}
    var v4 uint = 266
    switch "foo bar" {
    case "foo bar":
        _ = len(v2)
    case "hello":
        var v5 uint8 = 42
        _ = v5
    case "a":
        v1 += 161
    default:
        var v6 int16 = 1
        _ = v6
    }
    v7 := GenS1{F0: 18446744073709551615, F1: 377, F2: 859}
    if true {
        v2[42] = 0
        v8 := GenS1{F0: 4294967295, F1: -1, F2: 42}
        if false {
            v9 := GenS2{F0: (1e10 >= 1e10), F1: 42}
            fmt.Println(v1, v4)
            _ = v9
        }
        _ = v8
    }
    v1 = v1
    v10 := GenS3{F0: 32768}
    var v11 uint8 = uint8(v4)
    fmt.Println(v11, v4, v1)
    v1 += v11
    v12 := GenS4{F0: -9223372036854775808, F1: -1}
    if (1.5 != 3.14159) {
        v13 := GenS4{F0: 42, F1: 39}
        if (true || (!(!true))) {
            v4 += (^(v4))
            var v14 int64 = 1
            v3 = append(v3, -2147483648)
            _ = v14
        }
        _ = v13
    } else {
        var v15 uint32 = 673
        fmt.Println(v4, v4, v4)
        if (!(!(true && (int(v15) > 138)))) {
            v16 := GenS1{F0: 3, F1: -100, F2: v15}
            v17 := GenS5{F0: -106, F1: 1}
            _ = v16
            _ = v17
        }
        _ = v15
    }
    fmt.Println(v1, v11)
    if (-9223372036854775808 > 93) {
        if false {
            v1 = 139
            for fi18 := 0; fi18 < 16; fi18++ {
                v2[fi18] = fi18
                for fi19 := 0; fi19 < 4; fi19++ {
                    var v20 string = ("" + "hello")
                    if false {
                        var v21 int8 = 1
                        switch 0 {
                        case 0:
                            switch (fi18 | int(v4)) {
                            case 32768:
                                var v22 int64 = -808
                                _ = v22
                            case -2147483648:
                                v1 += v11
                            }
                            v23 := GenS2{F0: (((float64(-2.25) * -2.25) * float64(0.1)) != 1.5), F1: 65535}
                            switch (^(fi18)) {
                            case -108:
                                switch (fi19 + (fi18 & fi18)) {
                                case 42:
                                    if (-2.25 != (1e10 * ((0.0 + 1e10) / 1.0))) {
                                        var v24 uint32 = 2
                                        if (123 >= 290) {
                                            v25 := GenS6{F0: 42}
                                            switch v20 {
                                            case "foo bar":
                                                var v26 uint = v4
                                                _ = v26
                                            case "\n\t":
                                                var v27 int64 = -376
                                                _ = v27
                                            default:
                                                switch ("" + "foo bar") {
                                                case "foo bar":
                                                    if false {
                                                        if false {
                                                            if false {
                                                                var v28 int16 = 32767
                                                                var v29 uint8 = (v11 & (v11 | v11))
                                                                switch fi18 {
                                                                case 127:
                                                                    if (-36 != v21) {
                                                                        if true {
                                                                            v1 = (v1 | (v11 << ((4294967295) & 63)))
                                                                            fmt.Println(v20, v24, v20)
                                                                            fmt.Println(v11)
                                                                            fmt.Println(fi18)
                                                                        }
                                                                    } else {
                                                                        var v30 float64 = (-2.25 - (3.14159 / (1e10 + (-2.25 * ((-2.25 - 0.1) - float64((-2.25 + float32(1.5))))))))
                                                                        v20 = v20
                                                                        _ = len(v2)
                                                                        _ = len(v2)
                                                                        switch fi18 {
                                                                        case -2147483648:
                                                                            var v31 uint8 = v29
                                                                            _ = v31
                                                                        }
                                                                        if ((v11 == 255) || false) {
                                                                            if ((((!false) || ((v21 & 1) == v21)) && true) || ((!(!true)) && (v24 == (^(v24))))) {
                                                                                v32 := GenS5{F0: 762, F1: v21}
                                                                                _ = v32
                                                                            }
                                                                            v4 = v4
                                                                            v33 := GenS7{F0: 971, F1: 32767, F2: -640}
                                                                            _ = v33
                                                                        }
                                                                        _ = v30
                                                                    }
                                                                case -2147483648:
                                                                    var v34 bool = (!false)
                                                                    _ = v34
                                                                default:
                                                                    fmt.Println(v24, v24, v20)
                                                                }
                                                                v3 = append(v3, (fi19 + (fi19 << (((v4 + v4)) & 63))))
                                                                var v35 uint32 = v24
                                                                _ = v28
                                                                _ = v29
                                                                _ = v35
                                                            }
                                                            if (!(!false)) {
                                                                v36 := GenS8{F0: 671, F1: v4, F2: v20}
                                                                v21 = (v21 + (v21 >> ((v4) & 63)))
                                                                _ = v36
                                                            }
                                                            v4 = (v4 * 0)
                                                        }
                                                        var v37 int16 = 42
                                                        v1 -= v1
                                                        _ = v37
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = (fi18 + 32768)
                                                        }
                                                        var v38 uint64 = 42
                                                        v1 -= 128
                                                        _ = v38
                                                    }
                                                default:
                                                    switch v20 {
                                                    case "":
                                                        var v39 float64 = float64(0.0)
                                                        _ = v39
                                                    }
                                                }
                                            }
                                            v21 -= v21
                                            _ = v25
                                        }
                                        _ = v24
                                    } else {
                                        var v40 uint8 = (^(v11))
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v41 uint16 = 2
                                        _ = v40
                                        _ = v41
                                    }
                                default:
                                    v42 := GenS1{F0: 4294967295, F1: 45, F2: 2}
                                    _ = v42
                                }
                            default:
                                v11 = (v11 << ((4294967295) & 63))
                            }
                            v11 = 127
                            v21 -= -60
                            var v43 int8 = v21
                            _ = v23
                            _ = v43
                        case 128:
                            if (356 >= 596) {
                                v3 = append(v3, fi18)
                                v3 = append(v3, 419)
                                v20 = (v20 + v20)
                                v44 := GenS5{F0: -945, F1: 42}
                                v21 = v21
                                v11 += (v11 >> (((v4 | v4)) & 63))
                                _ = v44
                            } else {
                                var v45 int32 = 2147483647
                                var v46 string = v20
                                _ = v45
                                _ = v46
                            }
                        default:
                            v47 := v2[fi19]
                            _ = v47
                        }
                        _ = v21
                    }
                    _ = fi19
                    _ = v20
                }
                fmt.Println(v11, v11)
                _ = fi18
            }
            var v48 int64 = -1
            _ = v48
        } else {
            v1 = (v11 & v11)
            v1 += v11
        }
        v49 := GenS9{F0: 303, F1: 388, F2: "a"}
        var v50 uint = (^(v4))
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        if false {
            v1 = v11
            for fi52 := 0; fi52 < 5; fi52++ {
                var v53 uint16 = 146
                if (v53 == (v53 % (((v53 << ((uint(v53)) & 63))) | 1))) {
                    var v54 string = "\n\t"
                    _ = v54
                } else {
                    if false {
                        v11 += v11
                        v53 += v53
                    }
                    v55 := GenS10{F0: (-1 <= -1), F1: 0, F2: -51}
                    _ = v55
                }
                _ = fi52
                _ = v53
            }
            if (false && true) {
                var v56 float32 = 1e10
                _ = v56
            } else {
                v11 -= uint8(v4)
                v1 += 42
            }
        } else {
            if false {
                fmt.Println(v1, v1)
            } else {
                fmt.Println(v1, v1)
                v1 = 41
            }
        }
        _ = fi51
    }
    v11 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
