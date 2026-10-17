package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int8; F1 uint8 }
func genFn0() uint32 {
    var v5 float32 = 3.14159
    switch -206 {
    case 2147483647:
        switch "hello" {
        case "\n\t":
            if false {
                switch "hello" {
                case "\n\t":
                    v5 = v5
                case "foo bar":
                    v5 -= 1e10
                case "hello":
                    v6 := []int{2147483647}
                    _ = v6
                }
            } else {
                for fi7 := 0; fi7 < 17; fi7++ {
                    fmt.Println(v5, v5)
                    _ = fi7
                }
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 int64 = 9223372036854775807
                    _ = fi8
                    _ = v9
                }
                var v10 uint64 = 2
                _ = v10
            }
        case "hello":
            v11 := GenS1{F0: 32768, F1: 1, F2: (!(-2147483648 == 1))}
            _ = v11
        case "foo bar":
            if false {
                for fi12 := 0; fi12 < 15; fi12++ {
                    v5 = ((v5 / ((1e10 * ((v5 + float32((0.1 * float64(float32(((3.14159 - -2.25) + float64(v5))))))) * 0.1)) + v5)) * (3.14159 / 1e10))
                    v5 = v5
                    v5 += 0.1
                    _ = fi12
                }
                v13 := []int{-1, -9223372036854775808, 564}
                _ = v13
            } else {
                switch "a" {
                case "foo bar":
                    var v14 float64 = ((1.5 + (3.14159 + (0.0 / ((float64((v5 / v5)) / 3.14159) - 3.14159)))) + 0.1)
                    _ = v14
                default:
                    if (974 < 65535) {
                        v15 := []int{32767, 2147483647, 670}
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = v15
                    } else {
                        v5 += 0.1
                    }
                }
                if false {
                    var v16 int = 0
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        for fi18 := 0; fi18 < 12; fi18++ {
                            v19 := GenS2{F0: 0, F1: 1}
                            v20 := GenS3{F0: 203}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        if false {
                            for fi21 := 0; fi21 < 2; fi21++ {
                                if (true && ("a" > "\n\t")) {
                                    switch (fi17 | (fi17 * -2147483648)) {
                                    case 9223372036854775807:
                                        var v22 int32 = -1
                                        _ = v22
                                    }
                                }
                                var v23 int32 = 162
                                _ = fi21
                                _ = v23
                            }
                            var v24 uint8 = 208
                            v25 := GenS1{F0: 9223372036854775807, F1: 4, F2: true}
                            _ = v24
                            _ = v25
                        } else {
                            if true {
                                if true {
                                    switch ("a" + "a") {
                                    case "foo bar":
                                        var v26 uint8 = 226
                                        _ = v26
                                    }
                                    if false {
                                        switch "a" {
                                        case "a":
                                            var v27 int8 = 0
                                            _ = v27
                                        }
                                        for fi28 := 0; fi28 < 15; fi28++ {
                                            v29 := GenS4{F0: 42}
                                            if true {
                                                fmt.Println(fi17)
                                                var v30 uint32 = 2
                                                _ = v30
                                            } else {
                                                v5 += v5
                                            }
                                            fmt.Println(fi28)
                                            _ = fi28
                                            _ = v29
                                        }
                                    }
                                } else {
                                    for fi31 := 0; fi31 < 9; fi31++ {
                                        fmt.Println(fi17)
                                        if (1 <= 0) {
                                            var v32 uint = 1
                                            _ = v32
                                        }
                                        _ = fi31
                                    }
                                    for fi33 := 0; fi33 < 12; fi33++ {
                                        fmt.Println(v5, fi17, fi33)
                                        _ = fi33
                                    }
                                }
                                v5 = 0.1
                            } else {
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    fmt.Println(v5)
                                    var v35 uint8 = 128
                                    _ = fi34
                                    _ = v35
                                }
                                if (-1 < -2147483648) {
                                    v36 := []int{(fi17 % ((fi17) | 1)), fi17, fi17, 128}
                                    if false {
                                        for fi37 := 0; fi37 < 3; fi37++ {
                                            v5 = 1e10
                                            v38 := GenS3{F0: 42}
                                            _ = fi37
                                            _ = v38
                                        }
                                    }
                                    v5 += v5
                                    _ = v36
                                }
                                for fi39 := 0; fi39 < 6; fi39++ {
                                    var v40 int8 = 127
                                    _ = fi39
                                    _ = v40
                                }
                            }
                        }
                        _ = fi17
                    }
                }
            }
        }
    default:
        for fi41 := 0; fi41 < 4; fi41++ {
            v42 := GenS5{F0: (fi41 * fi41), F1: -100}
            v5 -= float32(float64((((v5 / (float32(0.0) + v5)) - v5) + (v5 * (v5 + 1.5)))))
            var v43 uint16 = 42
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    var v44 int8 = 1
    _ = v44
    _ = v5
    return 427
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -22
    v3 := []int{1}
    v4 := make(map[int]bool)
    var v45 string = (("" + "\n\t") + ("" + "hello"))
    var v46 int16 = 240
    v47 := GenS4{F0: 42}
    var v48 int = 32767
    _ = len(v1)
    if (!(!(34 >= 0))) {
        delete(v1, 32767)
        v49 := GenS2{F0: 2, F1: (^(v2))}
        _ = v49
    }
    v50 := GenS6{F0: 1, F1: 127}
    _ = len(v1)
    fmt.Println(v48, v2)
    v51 := v4[1]
    fmt.Println(v45, v45)
    var v52 int = 128
    _ = len(v3)
    v46 = v46
    v3 = append(v3, v48)
    _ = len(v3)
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 uint64 = 42
        if (v51 && (!((!true) && v51))) {
            var v55 uint64 = 742
            var v56 float64 = 0.0
            _ = v55
            _ = v56
        } else {
            var v57 uint64 = (v54 % (((v54 / ((45) | 1))) | 1))
            _ = v57
        }
        if (((float32(1e10) * 0.1) + 3.14159) >= 3.14159) {
            if (-1 != -933) {
                v46 = 32767
            } else {
                v54 += v54
                v58 := GenS5{F0: 0, F1: (v48 | v52)}
                _ = v58
            }
            for fi59 := 0; fi59 < 7; fi59++ {
                v51 = (fi53 == (^(fi59)))
                _ = fi59
            }
            v4[422] = (128 == 174)
        }
        _ = fi53
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v50
    _ = v51
    _ = v52
}
