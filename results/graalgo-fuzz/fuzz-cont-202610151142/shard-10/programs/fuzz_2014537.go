package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint16; F1 int16; F2 int8 }
type GenS5 struct { F0 uint; F1 int8; F2 uint32 }

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]int)
    v3 := []int{0, -1, 127, 4294967295}
    v4 := GenS1{F0: (("hello" + "foo bar") >= "\n\t"), F1: 42}
    if true {
        _ = len(v2)
        if false {
            _ = len(v2)
            var v5 float32 = 3.14159
            var v6 uint16 = 673
            _ = v5
            _ = v6
        }
        v7 := v2[128]
        _ = v7
    } else {
        delete(v2, 762)
    }
    v8 := GenS1{F0: true, F1: -802}
    var v9 uint = 569
    var v10 int = 274
    v1 -= uint64(v9)
    for fi11 := 0; fi11 < 11; fi11++ {
        v9 -= (^(v9))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi11
    }
    if (-1 != 1) {
        switch "foo bar" {
        case "":
            var v12 bool = (true && (231 > 2))
            _ = v12
        case "hello":
            var v13 float64 = 1e10
            _ = v13
        case "\n\t":
            var v14 uint32 = 714
            _ = v14
        }
        var v15 int64 = 430
        _ = v15
    }
    if false {
        switch v10 {
        case 128:
            v16 := GenS1{F0: (497 < (^(v9))), F1: -32768}
            _ = v16
        case -2147483648:
            switch v10 {
            case 0:
                var v17 int16 = 42
                _ = v17
            default:
                v1 = (^(v1))
            }
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            if (-32768 != -783) {
                v19 := v2[(fi18 & 42)]
                fmt.Println(v10, v19)
                switch (v10 * (^(v19))) {
                case 864:
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if false {
                            v21 := GenS1{F0: (!(!(uint32(v9) < 4294967295))), F1: -32768}
                            var v22 float32 = (3.14159 * ((-2.25 / 1.0) + (((0.0 - -2.25) * ((0.0 * 0.1) + ((1.5 + float32(float64(0.0))) / float32(1.5)))) - (float32((3.14159 / -2.25)) + ((-2.25 / (0.1 - (1e10 - 0.1))) + 1e10)))))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v21
                            _ = v22
                        }
                        _ = fi20
                    }
                case 127:
                    switch (v19 * (v19 << (((v9 ^ v9)) & 63))) {
                    case 2147483647:
                        var v23 uint8 = 1
                        _ = v23
                    default:
                        if ((0.0 * 1.5) <= 1e10) {
                            var v24 int32 = -1
                            var v25 float32 = (float32(-2.25) * (((((float32(1.5) - ((((1e10 / 1.0) / 0.1) + 3.14159) + -2.25)) + float32(1.5)) * ((0.1 + 3.14159) * -2.25)) - 0.0) + (1.5 / (0.1 * (0.1 + -2.25)))))
                            switch "foo bar" {
                            case "":
                                fmt.Println(v24)
                            case "hello":
                                v25 -= v25
                            default:
                                var v26 uint8 = 57
                                _ = v26
                            }
                            _ = v24
                            _ = v25
                        } else {
                            for fi27 := 0; fi27 < 11; fi27++ {
                                if false {
                                    v28 := GenS1{F0: false, F1: -1}
                                    v10 -= (fi27 | v19)
                                    _ = v28
                                }
                                var v29 uint64 = v1
                                var v30 uint16 = 76
                                _ = fi27
                                _ = v29
                                _ = v30
                            }
                            if true {
                                var v31 int16 = -1
                                _ = v31
                            }
                        }
                    }
                case -94:
                    for fi32 := 0; fi32 < 12; fi32++ {
                        v10 = -1
                        v10 += (v19 / ((4294967295) | 1))
                        if true {
                            v9 = v9
                        } else {
                            var v33 uint16 = 2
                            _ = v33
                        }
                        var v34 bool = true
                        v19 = fi18
                        v9 = 60
                        v2[fi18] = fi32
                        _ = fi32
                        _ = v34
                    }
                default:
                    v35 := GenS2{F0: -999, F1: v9}
                    _ = v35
                }
                _ = v19
            } else {
                fmt.Println(v9)
                for fi36 := 0; fi36 < 20; fi36++ {
                    v37 := GenS1{F0: false, F1: -32768}
                    v3 = append(v3, (-(fi18)))
                    if (((^(v1)) < v1) || ((("foo bar" + "a") + "a") > "foo bar")) {
                        fmt.Println(v1)
                    }
                    if (!true) {
                        if false {
                            if false {
                                v1 = 4294967295
                            }
                            v38 := GenS2{F0: 1, F1: 144}
                            _ = v38
                        } else {
                            v9 += (v9 * v9)
                            if ((!(3.14159 >= (0.0 - 1e10))) && true) {
                                switch 4294967295 {
                                case 4294967295:
                                    if true {
                                        if false {
                                            v9 = 2
                                        }
                                        fmt.Println(fi36, v9, v1)
                                        v39 := GenS2{F0: -1, F1: 42}
                                        _ = v39
                                    }
                                case -1:
                                    v1 = v1
                                case 85:
                                    v9 = v9
                                }
                                var v40 uint32 = uint32(v9)
                                _ = v40
                            } else {
                                switch ("foo bar" + "hello") {
                                case "a":
                                    switch 42 {
                                    case 4294967295:
                                        switch "foo bar" {
                                        case "":
                                            delete(v2, -9223372036854775808)
                                        case "foo bar":
                                            v1 = 1
                                            if len(v3) > 0 {
                                                v3[0] = (fi36 ^ v10)
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = v10
                                            }
                                        case "\n\t":
                                            if true {
                                                switch (fi36 % ((fi18) | 1)) {
                                                case 1:
                                                    if true {
                                                        if ("foo bar" >= "a") {
                                                            v3 = append(v3, fi18)
                                                            v41 := GenS2{F0: 951, F1: v9}
                                                            fmt.Println(v1, fi18)
                                                            _ = v41
                                                        }
                                                        v42 := GenS3{F0: false, F1: fi36}
                                                        _ = v42
                                                    }
                                                case -9223372036854775808:
                                                    switch ("hello" + "foo bar") {
                                                    case "a":
                                                        v2[(fi18 + (fi18 + fi36))] = 410
                                                    default:
                                                        v3 = append(v3, (fi18 ^ 78))
                                                    }
                                                default:
                                                    v1 -= 42
                                                }
                                                _ = len(v2)
                                            }
                                        default:
                                            v1 = (v1 + v1)
                                        }
                                    case 865:
                                        fmt.Println(v9, v10)
                                    default:
                                        var v43 uint = v9
                                        _ = v43
                                    }
                                case "\n\t":
                                    if ((!true) || false) {
                                        if (true && false) {
                                            if ((3.14159 <= float32(float64((3.14159 * -2.25)))) && false) {
                                                switch (fi36 - (-(v10))) {
                                                case -2147483648:
                                                    v44 := GenS4{F0: 255, F1: 32767, F2: -128}
                                                    _ = v44
                                                case 795:
                                                    v45 := GenS4{F0: 32767, F1: 1000, F2: 100}
                                                    _ = v45
                                                }
                                                _ = len(v2)
                                                v1 = (v1 - (v1 << ((v9) & 63)))
                                                var v46 uint64 = 1
                                                fmt.Println(v9, v10, v10)
                                                _ = v46
                                            }
                                            v1 = v1
                                            v9 = uint(v1)
                                        } else {
                                            if (860 != -627) {
                                                v47 := GenS4{F0: 1, F1: 1, F2: -1}
                                                _ = v47
                                            }
                                            if true {
                                                v48 := GenS1{F0: ((2 <= 42) && (568 <= (v9 ^ 255))), F1: int16(v1)}
                                                fmt.Println(v10)
                                                _ = v48
                                            }
                                        }
                                        fmt.Println(v9)
                                    }
                                }
                                v49 := GenS4{F0: 659, F1: 125, F2: 42}
                                _ = v49
                            }
                            v50 := GenS5{F0: v9, F1: 1, F2: 1}
                            _ = v50
                        }
                        v9 = (v9 << (((v9 & v9)) & 63))
                        v51 := GenS4{F0: 26, F1: -1, F2: -106}
                        _ = v51
                    } else {
                        var v52 uint = v9
                        v9 -= v9
                        _ = v52
                    }
                    _ = fi36
                    _ = v37
                }
            }
            v9 = 231
            _ = fi18
        }
    }
    var v53 int16 = -727
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v53:", v53)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v8
    _ = v9
}
