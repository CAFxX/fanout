package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int; F1 uint64; F2 uint }
type GenS6 struct { F0 bool }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]int)
    var v3 bool = (-2147483648 < 128)
    var v4 bool = v3
    var v5 uint8 = 235
    v3 = (v1 >= v1)
    var v6 bool = v4
    fmt.Println(v6, v6, v3)
    fmt.Println(v6, v3)
    v6 = (!(32768 <= 0))
    for fi7 := 0; fi7 < 2; fi7++ {
        v4 = (!true)
        fmt.Println(v4)
        _ = fi7
    }
    v6 = v6
    v8 := GenS1{F0: 4294967295}
    v9 := GenS2{F0: 0}
    var v10 float64 = (0.0 - ((0.0 + v1) + ((v1 * v1) * v1)))
    if (-2.25 >= (3.14159 / (1e10 * (1e10 / 1.0)))) {
        fmt.Println(v1, v5, v1)
        for fi11 := 0; fi11 < 11; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                if (!v6) {
                    fmt.Println(fi11)
                    var v13 uint32 = 4294967295
                    _ = v13
                }
                _ = fi12
            }
            _ = fi11
        }
        v4 = (-2.25 < -2.25)
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            if v4 {
                fmt.Println(v3, v1, v4)
            } else {
                switch fi14 {
                case 1:
                    for fi15 := 0; fi15 < 9; fi15++ {
                        if v4 {
                            v16 := GenS1{F0: 18446744073709551615}
                            v1 -= ((0.1 + v10) + (((1.5 + v10) + v10) * ((v10 - 1.5) / v1)))
                            var v17 uint = 4294967295
                            _ = v16
                            _ = v17
                        } else {
                            var v18 bool = v6
                            v18 = false
                            var v19 int64 = 1
                            var v20 uint64 = 2
                            var v21 float32 = -2.25
                            v22 := []int{-9223372036854775808, int(v5)}
                            if (!v6) {
                                switch -9223372036854775808 {
                                case 825:
                                    v5 += 127
                                case -67:
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                }
                                switch ("" + "") {
                                case "a":
                                    var v23 uint8 = (v5 * 97)
                                    _ = v23
                                }
                                v21 += float32(float64(v21))
                            }
                            var v24 float32 = (v21 * 1e10)
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v24
                        }
                        v10 = (v10 + (v1 / v1))
                        _ = fi15
                    }
                default:
                    if v3 {
                        v4 = false
                    } else {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            var v26 uint64 = 656
                            if true {
                                fmt.Println(v4, v5)
                                v27 := []int{(fi25 & fi25), (fi25 + fi25), 212, (fi14 % (((fi14 | (fi14 + fi25))) | 1))}
                                if (42 >= 1) {
                                    v1 = (v10 + (3.14159 + float64(-2.25)))
                                    v10 += v1
                                    _ = len(v27)
                                }
                                _ = v27
                            } else {
                                v6 = v4
                            }
                            _ = fi25
                            _ = v26
                        }
                        if ((!(false || v3)) && (313 == 21)) {
                            for fi28 := 0; fi28 < 10; fi28++ {
                                if ((-(fi28)) >= -859) {
                                    var v29 uint32 = 667
                                    _ = v29
                                }
                                fmt.Println(v3, fi14)
                                var v30 int16 = -1
                                _ = fi28
                                _ = v30
                            }
                        }
                        for fi31 := 0; fi31 < 14; fi31++ {
                            switch fi31 {
                            case 1:
                                v10 -= float64((1e10 / 1.0))
                            }
                            var v32 int = fi31
                            _ = fi31
                            _ = v32
                        }
                    }
                }
                v33 := GenS3{F0: (^(v5)), F1: -95}
                _ = v33
            }
            _ = fi14
        }
        if (1 >= 2) {
            v4 = (-1 == -9223372036854775808)
        } else {
            var v34 int32 = 2147483647
            v3 = v6
            _ = v34
        }
        switch "hello" {
        case "":
            v35 := GenS1{F0: 836}
            _ = v35
        case "\n\t":
            for fi36 := 0; fi36 < 13; fi36++ {
                var v37 float64 = (-2.25 + -2.25)
                v38 := GenS4{F0: -9223372036854775808, F1: 421}
                v4 = (!v4)
                _ = fi36
                _ = v37
                _ = v38
            }
        default:
            v39 := []int{4294967295, -696, 2147483647}
            _ = v39
        }
    }
    var v40 int16 = 42
    v4 = (v3 && (!v6))
    v41 := GenS5{F0: int(v5), F1: 2, F2: 42}
    for fi42 := 0; fi42 < 13; fi42++ {
        v43 := []int{fi42}
        v2[542] = fi42
        _ = fi42
        _ = v43
    }
    var v44 uint = 582
    v45 := GenS6{F0: (15 > 255)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
