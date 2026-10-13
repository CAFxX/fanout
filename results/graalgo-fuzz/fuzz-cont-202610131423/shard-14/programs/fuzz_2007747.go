package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint; F1 uint64; F2 int16 }
type GenS4 struct { F0 string; F1 int32; F2 uint8 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 float64 = -2.25
    var v4 int = 127
    var v5 uint16 = 1
    fmt.Println(v4)
    var v6 uint = 882
    v6 -= v6
    var v7 int8 = 1
    v5 += 32767
    if true {
        if (true || false) {
            if false {
                if false {
                    for fi8 := 0; fi8 < 5; fi8++ {
                        switch "" {
                        case "\n\t":
                            v5 += 1
                        case "":
                            v6 += 2
                        default:
                            if (((v7 & -38) > (v7 | 23)) && ("foo bar" > "foo bar")) {
                                switch (fi8 % ((fi8) | 1)) {
                                case 94:
                                    if false {
                                        fmt.Println(v5)
                                    }
                                case -818:
                                    v1[(fi8 << (((v6 / (((v6 * (v6 << ((922) & 63)))) | 1))) & 63))] = "\n\t"
                                case 32768:
                                    v6 -= (v6 - v6)
                                default:
                                    var v9 string = ""
                                    _ = v9
                                }
                            } else {
                                v5 -= v5
                            }
                        }
                        _ = fi8
                    }
                    if true {
                        v10 := []int{v4}
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 12; fi11++ {
                        delete(v2, ("" + "\n\t"))
                        fmt.Println(v5, v6)
                        switch "foo bar" {
                        case "\n\t":
                            if false {
                                v12 := GenS1{F0: 360, F1: (v5 | (v5 & (v5 ^ v5))), F2: 893}
                                var v13 uint8 = 56
                                _ = v12
                                _ = v13
                            } else {
                                for fi14 := 0; fi14 < 20; fi14++ {
                                    if true {
                                        var v15 string = ""
                                        _ = v15
                                    } else {
                                        v16 := GenS2{F0: 18446744073709551615, F1: -32768}
                                        _ = v16
                                    }
                                    _ = fi14
                                }
                                var v17 float32 = 3.14159
                                _ = v17
                            }
                        case "hello":
                            v18 := GenS3{F0: (v6 / ((v6) | 1)), F1: 18446744073709551615, F2: 42}
                            _ = v18
                        default:
                            if true {
                                var v19 float32 = 1e10
                                for fi20 := 0; fi20 < 18; fi20++ {
                                    v4 += (-(fi11))
                                    v21 := GenS4{F0: "", F1: -1, F2: 255}
                                    _ = fi20
                                    _ = v21
                                }
                                if (true || ((v7 | -128) != (v7 << ((870) & 63)))) {
                                    for fi22 := 0; fi22 < 11; fi22++ {
                                        v5 = 42
                                        v4 += (-(v4))
                                        var v23 uint = v6
                                        _ = fi22
                                        _ = v23
                                    }
                                    v6 += v6
                                    v24 := []int{fi11}
                                    _ = v24
                                }
                                _ = v19
                            } else {
                                v7 = (v7 ^ (-(v7)))
                                fmt.Println(fi11)
                            }
                        }
                        _ = fi11
                    }
                } else {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        fmt.Println(v5, v7, v7)
                        v26 := GenS3{F0: (v6 / ((v6) | 1)), F1: 603, F2: 1000}
                        v5 = 42
                        _ = fi25
                        _ = v26
                    }
                    var v27 int64 = -9223372036854775808
                    if (32767 == v5) {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            v29 := []int{fi28}
                            for fi30 := 0; fi30 < 7; fi30++ {
                                fmt.Println(v6, v5)
                                if false {
                                    var v31 int16 = 0
                                    if (2147483647 > 2147483647) {
                                        v31 -= 1000
                                    }
                                    _ = v31
                                }
                                _ = fi30
                            }
                            if true {
                                var v32 int8 = (^(v7))
                                if len(v29) >= 2 {
                                    v29 = v29[1:len(v29)]
                                }
                                for fi33 := 0; fi33 < 14; fi33++ {
                                    fmt.Println(fi28, fi28, v32)
                                    delete(v2, ("\n\t" + ("foo bar" + "foo bar")))
                                    v1[(^(fi28))] = ("\n\t" + "foo bar")
                                    _ = fi33
                                }
                                _ = v32
                            } else {
                                for fi34 := 0; fi34 < 9; fi34++ {
                                    fmt.Println(v4, v6, fi28)
                                    if false {
                                        var v35 float64 = v3
                                        if false {
                                            v35 += -2.25
                                        }
                                        _ = v35
                                    }
                                    _ = fi34
                                }
                                if false {
                                    var v36 int8 = -52
                                    if true {
                                        var v37 uint32 = 2
                                        _ = v37
                                    } else {
                                        v27 = (-(v27))
                                        var v38 int64 = (v27 >> (((v6 - v6)) & 63))
                                        _ = v38
                                    }
                                    _ = v36
                                } else {
                                    _ = len(v1)
                                    v39 := GenS4{F0: "a", F1: int32(v27), F2: 0}
                                    for fi40 := 0; fi40 < 13; fi40++ {
                                        v5 = (v5 & (v5 ^ 2))
                                        v7 -= 113
                                        v41 := GenS4{F0: "a", F1: -106, F2: 79}
                                        _ = fi40
                                        _ = v41
                                    }
                                    _ = v39
                                }
                            }
                            _ = fi28
                            _ = v29
                        }
                        v42 := []int{4294967295, (v4 * (v4 / ((447) | 1)))}
                        _ = v42
                    } else {
                        v3 = (3.14159 * (float64(1.5) - v3))
                    }
                    _ = v27
                }
            } else {
                switch 9223372036854775807 {
                case -907:
                    v43 := GenS1{F0: 1, F1: v5, F2: v5}
                    _ = v43
                }
            }
        }
        v7 -= (v7 & (v7 + (v7 % ((v7) | 1))))
        switch "foo bar" {
        case "hello":
            switch v4 {
            case 2147483647:
                var v44 bool = true
                _ = v44
            case 127:
                fmt.Println(v6, v5, v4)
            case 32767:
                v45 := v1[(v4 | 1)]
                _ = v45
            default:
                var v46 float32 = 3.14159
                _ = v46
            }
        default:
            v47 := GenS4{F0: "", F1: -513, F2: 127}
            _ = v47
        }
    } else {
        for fi48 := 0; fi48 < 11; fi48++ {
            fmt.Println(v4, fi48, v5)
            switch "hello" {
            case "foo bar":
                if (true && false) {
                    var v49 bool = (1 <= 627)
                    v3 -= v3
                    _ = v49
                }
            case "":
                var v50 float32 = 1.5
                _ = v50
            case "hello":
                if true {
                    for fi51 := 0; fi51 < 13; fi51++ {
                        fmt.Println(v3, v5, v6)
                        v52 := GenS5{F0: false}
                        var v53 int64 = 42
                        _ = fi51
                        _ = v52
                        _ = v53
                    }
                    var v54 int64 = 1
                    _ = v54
                } else {
                    fmt.Println(v7, v7, v6)
                    v55 := []int{v4, (fi48 << ((v6) & 63))}
                    switch "a" {
                    case "foo bar":
                        var v56 uint64 = 283
                        _ = v56
                    }
                    _ = v55
                }
            default:
                var v57 uint64 = 956
                _ = v57
            }
            fmt.Println(v5, v3, v5)
            _ = fi48
        }
    }
    var v58 uint16 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v7
}
