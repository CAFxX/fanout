package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint64 }

func main() {
    var v1 int = -2147483648
    var v2 bool = ((!true) || ((!(!false)) && false))
    v3 := make(map[int]string)
    var v4 float64 = 1.5
    v5 := []int{v1, 42, 9223372036854775807, (v1 & (^(v1)))}
    if ((42 >= 2147483647) && v2) {
        delete(v3, (-(v1)))
        for fi6 := 0; fi6 < 20; fi6++ {
            for fi7 := 0; fi7 < 3; fi7++ {
                v2 = (2 != 397)
                if v2 {
                    v2 = false
                } else {
                    var v8 string = "a"
                    var v9 uint8 = 0
                    _ = v8
                    _ = v9
                }
                _ = fi7
            }
            v3[(v1 >> ((0) & 63))] = ("a" + (("\n\t" + "a") + "foo bar"))
            v1 += v1
            _ = fi6
        }
    }
    var v10 int64 = -9223372036854775808
    if true {
        v11 := GenS1{F0: 863}
        switch "\n\t" {
        case "hello":
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: 42}
                var v14 int8 = -1
                if (!(1 > (v10 * (v10 >> ((4294967295) & 63))))) {
                    if false {
                        v1 = v1
                        v10 -= (v10 + 0)
                    } else {
                        v14 = v14
                        fmt.Println(v4, v14, v10)
                    }
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        case "":
            if (25 >= 42) {
                fmt.Println(v2, v10, v10)
                if ((0.1 - -2.25) == (-2.25 / -2.25)) {
                    v2 = v2
                    if ((v1 >> ((198) & 63)) <= 2147483647) {
                        for fi15 := 0; fi15 < 14; fi15++ {
                            v16 := GenS2{F0: 65535, F1: 0}
                            _ = fi15
                            _ = v16
                        }
                        v1 += (v1 / ((v1) | 1))
                        if v2 {
                            var v17 uint = 117
                            fmt.Println(v1, v17, v2)
                            for fi18 := 0; fi18 < 1; fi18++ {
                                v19 := GenS2{F0: 0, F1: -1}
                                if (32767 > 219) {
                                    for fi20 := 0; fi20 < 9; fi20++ {
                                        fmt.Println(v4, fi20)
                                        delete(v3, fi18)
                                        v4 -= ((v4 - v4) - ((((float64(1e10) / float64(0.1)) + ((((v4 / 0.1) / (v4 / 1.0)) + ((0.1 - v4) - v4)) + (0.0 - 1.5))) / 1e10) - 0.0))
                                        _ = fi20
                                    }
                                    v21 := GenS3{F0: 128}
                                    _ = v21
                                } else {
                                    fmt.Println(v17, v17, v4)
                                    if true {
                                        for fi22 := 0; fi22 < 18; fi22++ {
                                            var v23 uint8 = 194
                                            _ = fi22
                                            _ = v23
                                        }
                                        for fi24 := 0; fi24 < 2; fi24++ {
                                            var v25 float32 = 0.0
                                            _ = fi24
                                            _ = v25
                                        }
                                    } else {
                                        delete(v3, (-(fi18)))
                                    }
                                    fmt.Println(v2)
                                }
                                _ = fi18
                                _ = v19
                            }
                            _ = v17
                        } else {
                            _ = len(v5)
                            switch (-(v1)) {
                            case 42:
                                var v26 int16 = -1
                                _ = v26
                            }
                        }
                    }
                } else {
                    var v27 bool = (42 == 2147483648)
                    _ = v27
                }
            } else {
                v1 += 316
                v10 -= 1
                v28 := GenS4{F0: 2147483647, F1: 2}
                _ = v28
            }
        }
        _ = v11
    } else {
        var v29 float32 = 0.1
        var v30 int32 = 89
        _ = v29
        _ = v30
    }
    v2 = v2
    v4 += (1.5 + v4)
    v10 = (v10 ^ v10)
    if (v2 || ((0.1 * float64(1.5)) != (v4 - (float64(((((1e10 / ((1e10 * (1.5 * 0.1)) * 1e10)) / 1e10) + (3.14159 / -2.25)) - (0.1 / 0.1))) + float64(1e10))))) {
        v5 = append(v5, v1)
    }
    fmt.Println(v2, v4)
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 uint8 = 255
        v3[9223372036854775807] = "\n\t"
        fmt.Println(v10, v2)
        _ = fi31
        _ = v32
    }
    v33 := GenS5{F0: 1, F1: (("" + ("\n\t" + "hello")) < ("\n\t" + ("\n\t" + "a"))), F2: 42}
    if v2 {
        var v34 int16 = -32768
        _ = v34
    }
    switch 32767 {
    case -2147483648:
        v2 = (!(127 != 127))
    default:
        v10 = (v10 << ((65535) & 63))
    }
    v10 = (v10 & 1)
    if v2 {
        v35 := GenS1{F0: 328}
        if (v2 || false) {
            if v2 {
                v2 = (("\n\t" + "a") >= "a")
                if true {
                    fmt.Println(v10)
                }
            } else {
                v1 -= 1000
            }
        }
        _ = v35
    } else {
        if true {
            var v36 int = (v1 ^ (v1 << ((65535) & 63)))
            _ = v36
        }
        v10 = (^(v10))
    }
    var v37 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
