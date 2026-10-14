package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int }
type GenS2 struct { F0 int32; F1 uint16; F2 int32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint8; F1 string; F2 bool }
func genFn0(p0 string) uint8 {
    p0 = p0
    v5 := GenS1{F0: 560, F1: -9223372036854775808, F2: 4294967295}
    _ = p0
    _ = v5
    return 2
}
func genFn1(p0 string) bool {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return (65535 != 419)
}

func main() {
    var v1 int64 = 755
    var v2 int8 = 42
    var v3 string = "foo bar"
    v4 := make(map[string]int)
    v6 := GenS2{F0: 2147483647, F1: 42, F2: -205}
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{fi7, (fi7 ^ fi7), (fi7 >> ((42) & 63))}
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 4294967295
    v10 := GenS3{F0: -570, F1: 127}
    switch -2147483648 {
    case 1:
        switch 32768 {
        case 32768:
            var v11 bool = true
            _ = v11
        case 32767:
            for fi12 := 0; fi12 < 14; fi12++ {
                v3 = "a"
                if ((v2 | int8(v1)) < 127) {
                    var v13 uint64 = v9
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v2 -= 127
                        var v15 float64 = 1.5
                        var v16 uint = 4294967295
                        if false {
                            v3 = v3
                            v15 += v15
                            fmt.Println(v3, v16)
                            switch (fi12 >> ((v16) & 63)) {
                            case -9223372036854775808:
                                v13 += (v13 / ((42) | 1))
                            }
                            var v17 string = v3
                            if true {
                                switch 4294967295 {
                                case 32768:
                                    v18 := GenS2{F0: -467, F1: 0, F2: -1}
                                    _ = v18
                                case 4294967295:
                                    v17 = ""
                                case -760:
                                    var v19 int8 = 127
                                    _ = v19
                                }
                                v16 = (v16 & (^(v16)))
                                if false {
                                    v9 = (v9 % (((v9 / (((v13 + (v13 & v9))) | 1))) | 1))
                                    switch fi14 {
                                    case -514:
                                        fmt.Println(v16)
                                    case -2147483648:
                                        v3 = ""
                                    case -9223372036854775808:
                                        switch (v17 + "a") {
                                        case "hello":
                                            if true {
                                                var v20 int32 = -2147483648
                                                _ = v20
                                            }
                                        case "foo bar":
                                            fmt.Println(v17, fi12)
                                        default:
                                            v15 -= v15
                                        }
                                    }
                                    var v21 int8 = -1
                                    _ = v21
                                }
                            } else {
                                var v22 bool = ((v9 / ((42) | 1)) >= v13)
                                v16 = v16
                                if (!(!(v22 || true))) {
                                    switch v3 {
                                    case "a":
                                        v22 = ((v17 + v17) >= ((v3 + "foo bar") + v17))
                                    case "hello":
                                        v17 = "foo bar"
                                    case "":
                                        if v22 {
                                            switch v17 {
                                            case "":
                                                fmt.Println(v3, v22)
                                            default:
                                                v23 := GenS3{F0: 1, F1: 42}
                                                _ = v23
                                            }
                                        } else {
                                            v24 := []int{}
                                            _ = v24
                                        }
                                    default:
                                        switch fi14 {
                                        case 127:
                                            if false {
                                                var v25 uint = (v16 << ((v16) & 63))
                                                var v26 float32 = 0.0
                                                v27 := []int{fi14, 42}
                                                _ = v25
                                                _ = v26
                                                _ = v27
                                            }
                                        case 42:
                                            v28 := GenS1{F0: 2147483647, F1: (v1 & v1), F2: int(v9)}
                                            _ = v28
                                        default:
                                            if v22 {
                                                if (!false) {
                                                    v29 := v4[v17]
                                                    _ = v29
                                                }
                                            }
                                        }
                                    }
                                    if (2 == 2) {
                                        fmt.Println(v3)
                                    }
                                    v30 := GenS1{F0: 2147483647, F1: 9223372036854775807, F2: fi14}
                                    _ = v30
                                } else {
                                    v31 := []int{fi12, (fi14 >> ((2) & 63))}
                                    _ = v31
                                }
                                _ = v22
                            }
                            _ = v17
                        } else {
                            var v32 int8 = (v2 >> ((65535) & 63))
                            v33 := GenS4{F0: v2, F1: v9, F2: 42}
                            v15 -= 1.5
                            _ = v32
                            _ = v33
                        }
                        v34 := []int{int(v1)}
                        v35 := GenS3{F0: fi14, F1: 1}
                        v16 = (^(v16))
                        v13 = (^(v13))
                        if (!(uint16(v9) == 971)) {
                            var v36 int16 = -185
                            v37 := GenS2{F0: -2147483648, F1: 611, F2: 0}
                            var v38 uint32 = 427
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            var v39 float32 = 3.14159
                            var v40 bool = (v15 == ((1.5 - 0.0) / 3.14159))
                            var v41 int32 = 10
                            v15 += ((float64((v39 / v39)) / float64((0.1 + v39))) + (v15 + ((((v15 + (v15 + (((v15 * v15) * (0.0 - v15)) + (v15 / -2.25)))) * v15) / 0.1) - v15)))
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v34
                        _ = v35
                    }
                    switch 431 {
                    case -1:
                        v1 += v1
                    case 701:
                        v42 := []int{}
                        _ = v42
                    }
                    _ = v13
                }
                if (32767 > 255) {
                    fmt.Println(v9, v9, fi12)
                    switch v3 {
                    case "\n\t":
                        var v43 uint64 = (v9 - 0)
                        _ = v43
                    case "hello":
                        var v44 int64 = (v1 % (((v1 & (v1 ^ (v1 - -9223372036854775808)))) | 1))
                        _ = v44
                    case "":
                        v3 = "\n\t"
                    }
                    v9 = (v9 / (((v9 % (((v9 / ((v9) | 1))) | 1))) | 1))
                }
                _ = fi12
            }
        default:
            if false {
                var v45 int32 = -2147483648
                _ = v45
            }
        }
    default:
        v46 := GenS5{F0: 255, F1: "", F2: (false && (!false))}
        _ = v46
    }
    _ = len(v4)
    var v47 int32 = 0
    var v48 int16 = -1
    for fi49 := 0; fi49 < 2; fi49++ {
        v2 = (v2 - 127)
        _ = fi49
    }
    for fi50 := 0; fi50 < 18; fi50++ {
        v3 = (v3 + v3)
        _ = fi50
    }
    var v51 uint = 332
    v52 := []int{-389, -54, -171, -860}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v51
    _ = v52
    _ = v6
    _ = v9
}
