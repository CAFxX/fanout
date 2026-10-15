package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS6 struct { F0 uint; F1 int; F2 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 int32; F2 int32 }
type GenS9 struct { F0 int64; F1 int32; F2 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    var v4 int = -2147483648
    var v5 int32 = -332
    var v6 int32 = -2147483648
    v7 := GenS1{F0: 32767, F1: 406}
    v8 := GenS2{F0: 0, F1: 0, F2: -1}
    switch "a" {
    case "\n\t":
        v1[4294967295] = ((65535 == 4294967295) && false)
    case "":
        v9 := GenS2{F0: 4294967295, F1: 2, F2: -91}
        _ = v9
    case "foo bar":
        var v10 bool = ((((1.5 * 3.14159) * 3.14159) <= float64(0.1)) || ((^(v5)) >= (v5 & v6)))
        _ = v10
    default:
        for fi11 := 0; fi11 < 6; fi11++ {
            if true {
                switch v4 {
                case 32768:
                    v12 := v3[2147483648]
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 19; fi13++ {
                        var v14 int = -213
                        if false {
                            v15 := GenS2{F0: 18446744073709551615, F1: 2147483648, F2: -1}
                            v6 = 2147483647
                            _ = v15
                        } else {
                            v16 := []int{(v14 << ((18446744073709551615) & 63)), 2147483647, (fi13 + v14)}
                            fmt.Println(v5)
                            if (float32((1.5 + -2.25)) <= (((0.0 - (float32((float64(0.0) / 3.14159)) - 0.0)) - -2.25) / ((float32(1e10) / 3.14159) * 3.14159))) {
                                fmt.Println(v4, v5)
                                v5 = -2147483648
                                if (2 == 1) {
                                    if (!(!true)) {
                                        fmt.Println(v5, v5, fi11)
                                        v17 := GenS1{F0: 903, F1: 1}
                                        var v18 float32 = 1.5
                                        _ = v17
                                        _ = v18
                                    } else {
                                        v19 := GenS3{F0: 60}
                                        v14 -= v4
                                        _ = v19
                                    }
                                }
                            }
                            _ = v16
                        }
                        switch "" {
                        case "hello":
                            v4 -= (v14 | 808)
                        case "a":
                            var v20 uint = 65535
                            if true {
                                v2[2147483647] = (fi11 % (((fi11 | (-(fi11)))) | 1))
                                v5 = v5
                            } else {
                                v21 := GenS2{F0: 2, F1: 2147483648, F2: 1000}
                                v4 -= fi13
                                if ((1 == 1000) && true) {
                                    var v22 int32 = (v6 << ((v20) & 63))
                                    _ = v22
                                }
                                _ = v21
                            }
                            _ = v20
                        case "foo bar":
                            var v23 float32 = 0.1
                            _ = v23
                        }
                        switch "hello" {
                        case "hello":
                            if false {
                                v4 += fi11
                            }
                        case "foo bar":
                            v6 = -1
                        }
                        v24 := GenS4{F0: 2, F1: 514}
                        _ = fi13
                        _ = v14
                        _ = v24
                    }
                }
            }
            fmt.Println(v5)
            _ = fi11
        }
    }
    if (true && true) {
        v6 -= v5
        for fi25 := 0; fi25 < 14; fi25++ {
            _ = len(v1)
            _ = fi25
        }
        var v26 int32 = 875
        _ = v26
    } else {
        if ((0 >= 671) && (("" + (((("\n\t" + "a") + ("hello" + (((((("hello" + (("" + "foo bar") + "\n\t")) + "hello") + (("a" + "a") + (("\n\t" + "hello") + "a"))) + "\n\t") + "foo bar") + ("hello" + "")))) + "\n\t") + "")) > "a")) {
            fmt.Println(v4, v6)
        }
        v27 := v1[697]
        fmt.Println(v5, v5)
        _ = v27
    }
    v4 = 127
    v4 += (v4 | v4)
    for fi28 := 0; fi28 < 18; fi28++ {
        switch v4 {
        case 32767:
            v29 := v2[535]
            _ = v29
        case -194:
            for fi30 := 0; fi30 < 6; fi30++ {
                v31 := GenS2{F0: 2, F1: 65535, F2: -1}
                if false {
                    fmt.Println(fi28, fi30, v5)
                    if (4294967295 == 42) {
                        if true {
                            v32 := GenS4{F0: 177, F1: 1000}
                            _ = v32
                        } else {
                            var v33 uint16 = 364
                            v34 := GenS5{F0: 2147483647, F1: -128, F2: 123}
                            _ = v33
                            _ = v34
                        }
                        var v35 int32 = 1
                        fmt.Println(fi28, v5)
                        _ = v35
                    }
                    switch fi28 {
                    case -2147483648:
                        v36 := GenS2{F0: 0, F1: 0, F2: int16(v5)}
                        _ = v36
                    case 127:
                        v37 := GenS6{F0: 65535, F1: (fi30 | (fi28 * 32767)), F2: ("\n\t" + "hello")}
                        _ = v37
                    }
                }
                var v38 float64 = 0.0
                _ = fi30
                _ = v31
                _ = v38
            }
        default:
            var v39 uint = 4294967295
            _ = v39
        }
        if (true && (571 > 4294967295)) {
            var v40 int32 = -492
            if false {
                for fi41 := 0; fi41 < 16; fi41++ {
                    v5 = v6
                    v40 = v5
                    v4 -= 957
                    if true {
                        _ = len(v3)
                        v42 := GenS7{F0: 1}
                        v1[866] = true
                        _ = v42
                    } else {
                        v43 := v1[2147483648]
                        var v44 uint32 = 2147483648
                        _ = v43
                        _ = v44
                    }
                    var v45 int64 = -1
                    _ = fi41
                    _ = v45
                }
                switch "a" {
                case "":
                    if (fi28 < v4) {
                        fmt.Println(v4, v6, v4)
                        for fi46 := 0; fi46 < 7; fi46++ {
                            fmt.Println(fi46, v4)
                            _ = fi46
                        }
                    }
                case "hello":
                    delete(v3, 4294967295)
                case "foo bar":
                    if false {
                        v47 := []int{fi28}
                        _ = v47
                    } else {
                        if (1 < 18446744073709551615) {
                            if true {
                                v40 = -146
                                for fi48 := 0; fi48 < 1; fi48++ {
                                    switch v4 {
                                    case 4294967295:
                                        switch -2147483648 {
                                        case 32768:
                                            v49 := GenS8{F0: 6, F1: v40, F2: (v40 >> ((472) & 63))}
                                            _ = v49
                                        }
                                    case 128:
                                        v50 := []int{(fi48 % (((fi28 | -81)) | 1)), (fi28 % ((fi28) | 1)), fi28}
                                        v5 = v6
                                        _ = v50
                                    }
                                    _ = fi48
                                }
                                for fi51 := 0; fi51 < 1; fi51++ {
                                    v5 = v40
                                    fmt.Println(fi51)
                                    var v52 int32 = v6
                                    _ = fi51
                                    _ = v52
                                }
                            }
                            if (0.1 <= 1e10) {
                                v53 := []int{v4}
                                if (!false) {
                                    v6 -= (v5 - (v5 * -2147483648))
                                    var v54 uint64 = 522
                                    _ = v54
                                }
                                fmt.Println(fi28, fi28)
                                _ = v53
                            }
                            if false {
                                fmt.Println(fi28, v6)
                                v5 = v40
                                var v55 int16 = 32767
                                _ = v55
                            }
                        }
                        v6 = v6
                    }
                default:
                    v56 := []int{(^(v4)), (-(fi28)), (v4 % ((fi28) | 1)), (v4 >> ((4294967295) & 63))}
                    _ = v56
                }
                switch ("hello" + ("a" + (("foo bar" + (("\n\t" + (("hello" + "foo bar") + "hello")) + "\n\t")) + "a"))) {
                case "hello":
                    v6 = v40
                case "a":
                    v3[65535] = 0
                default:
                    fmt.Println(v4)
                }
            } else {
                v57 := GenS9{F0: -73, F1: (v6 | 645), F2: 717}
                _ = v57
            }
            _ = v40
        }
        v58 := GenS9{F0: -868, F1: v6, F2: 534}
        _ = fi28
        _ = v58
    }
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
