package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8 }
func genFn0() int8 {
    var v7 uint8 = 2
    var v8 int8 = -47
    v9 := []int{1, 42, int(v7)}
    if (!true) {
        fmt.Println(v8, v8)
        v10 := GenS1{F0: (v7 << ((393) & 63)), F1: 229}
        switch "hello" {
        case "hello":
            if true {
                var v11 int64 = -802
                v7 = v7
                _ = v11
            } else {
                for fi12 := 0; fi12 < 7; fi12++ {
                    v8 += 127
                    _ = fi12
                }
                if false {
                    v7 -= v7
                } else {
                    if true {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        v13 := GenS2{F0: 166, F1: (v8 - (^(v8))), F2: (v8 - v8)}
                        _ = len(v9)
                        _ = v13
                    } else {
                        if ((v7 * v7) <= (v7 & v7)) {
                            if len(v9) > 0 {
                                v9[0] = 127
                            }
                            v8 = v8
                            for fi14 := 0; fi14 < 4; fi14++ {
                                fmt.Println(v7, v8, fi14)
                                var v15 int32 = 42
                                _ = fi14
                                _ = v15
                            }
                        }
                    }
                }
                fmt.Println(v7, v7)
            }
        default:
            switch "foo bar" {
            case "":
                var v16 uint64 = 1
                _ = v16
            case "a":
                var v17 uint8 = 79
                _ = v17
            default:
                for fi18 := 0; fi18 < 3; fi18++ {
                    if false {
                        v19 := GenS3{F0: -128}
                        _ = v19
                    }
                    v9 = append(v9, (fi18 % (((fi18 * 128)) | 1)))
                    _ = fi18
                }
            }
        }
        _ = v10
    } else {
        var v20 int8 = -128
        switch "" {
        case "hello":
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        default:
            v21 := GenS1{F0: v7, F1: 0}
            _ = v21
        }
        _ = v20
    }
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int = 2147483647
    var v2 float32 = 1e10
    var v3 float64 = 3.14159
    v4 := make(map[int]int)
    var v5 bool = (((!(0 <= 9223372036854775807)) && ((true || true) && (1 < 0))) || true)
    v6 := make(map[uint32]string)
    var v22 float64 = (3.14159 + 0.0)
    for fi23 := 0; fi23 < 14; fi23++ {
        v24 := GenS4{F0: 2, F1: fi23, F2: 805}
        if (1 > 2) {
            if (!(v1 == fi23)) {
                v3 += float64((v2 - 3.14159))
            } else {
                var v25 uint64 = 42
                _ = v25
            }
            switch v1 {
            case 2147483647:
                var v26 uint16 = 42
                _ = v26
            case -2147483648:
                if (v5 || (!(v5 && false))) {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        if false {
                            if v5 {
                                if true {
                                    v28 := GenS3{F0: -70}
                                    v5 = true
                                    var v29 uint8 = 128
                                    _ = v28
                                    _ = v29
                                }
                                var v30 uint8 = 128
                                v30 += (v30 & v30)
                                _ = v30
                            } else {
                                v31 := GenS5{F0: 2147483647}
                                var v32 bool = (2 != 4294967295)
                                v33 := v6[65535]
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        } else {
                            v3 += ((0.0 - v3) * ((v22 / v22) - 1e10))
                            var v34 int = -859
                            var v35 uint8 = 77
                            _ = v34
                            _ = v35
                        }
                        _ = fi27
                    }
                } else {
                    delete(v6, 2)
                    var v36 int16 = 0
                    switch ("\n\t" + ("a" + "\n\t")) {
                    case "\n\t":
                        v1 -= (fi23 - v1)
                    case "":
                        v37 := GenS6{F0: 297, F1: 1, F2: 42}
                        _ = v37
                    case "foo bar":
                        if true {
                            v38 := []int{fi23}
                            _ = v38
                        } else {
                            var v39 uint16 = 42
                            _ = v39
                        }
                    }
                    _ = v36
                }
            }
            if (42 == -1) {
                v40 := GenS4{F0: 65535, F1: fi23, F2: 844}
                if (32767 > 1) {
                    var v41 int = (fi23 / ((188) | 1))
                    _ = v41
                }
                _ = v40
            }
        } else {
            var v42 int = (v1 >> ((18446744073709551615) & 63))
            _ = v42
        }
        _ = fi23
        _ = v24
    }
    if (65535 > 566) {
        v22 += v3
    } else {
        v5 = v5
        if false {
            v43 := GenS7{F0: v1}
            var v44 uint32 = 272
            _ = v43
            _ = v44
        } else {
            fmt.Println(v3, v22)
            v45 := GenS8{F0: 152}
            _ = v45
        }
        v2 = (v2 - 0.0)
    }
    var v46 int64 = 9223372036854775807
    v1 += v1
    var v47 int32 = -2147483648
    if v5 {
        for fi48 := 0; fi48 < 9; fi48++ {
            for fi49 := 0; fi49 < 16; fi49++ {
                v6[65535] = "hello"
                var v50 int16 = 1000
                v51 := v4[fi48]
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = fi48
        }
        for fi52 := 0; fi52 < 1; fi52++ {
            var v53 float64 = v3
            _ = fi52
            _ = v53
        }
    } else {
        _ = len(v6)
    }
    var v54 uint64 = 1
    var v55 int64 = 0
    var v56 float64 = float64(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
