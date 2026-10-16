package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int }
type GenS2 struct { F0 bool; F1 uint32; F2 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint16; F2 int8 }

func main() {
    var v1 uint8 = 255
    var v2 uint8 = 249
    v3 := make(map[int]string)
    for fi4 := 0; fi4 < 1; fi4++ {
        fmt.Println(fi4, v1, fi4)
        _ = fi4
    }
    switch 42 {
    case -1:
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 uint8 = 0
            _ = fi5
            _ = v6
        }
    default:
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int = fi7
            v8 = (v8 + 128)
            _ = fi7
            _ = v8
        }
    }
    var v9 uint16 = 0
    v2 = (v2 & (v1 + 42))
    _ = len(v3)
    fmt.Println(v1, v2)
    v1 = v2
    switch 32767 {
    case 42:
        v10 := v3[-387]
        _ = v10
    case 445:
        var v11 uint32 = 4294967295
        _ = v11
    default:
        v12 := GenS1{F0: 0, F1: 730, F2: -667}
        _ = v12
    }
    if false {
        var v13 uint8 = 217
        if (!false) {
            for fi14 := 0; fi14 < 13; fi14++ {
                for fi15 := 0; fi15 < 9; fi15++ {
                    v16 := []int{fi14, 1}
                    var v17 uint16 = (v9 + (v9 / ((v9) | 1)))
                    if false {
                        var v18 float32 = ((1.5 - 0.1) / 1e10)
                        fmt.Println(v2, v17)
                        _ = v18
                    }
                    v13 = 42
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                switch "hello" {
                case "\n\t":
                    if false {
                        delete(v3, (fi14 << ((18446744073709551615) & 63)))
                    } else {
                        for fi19 := 0; fi19 < 3; fi19++ {
                            v20 := GenS2{F0: (-445 < -1), F1: 42, F2: 42}
                            switch ("" + "") {
                            case "":
                                if true {
                                    v21 := GenS2{F0: true, F1: 2147483647, F2: 880}
                                    if true {
                                        switch (fi14 ^ fi14) {
                                        case -765:
                                            v22 := GenS3{F0: 839}
                                            _ = v22
                                        }
                                        fmt.Println(v9, v9, fi19)
                                        var v23 uint16 = v9
                                        _ = v23
                                    }
                                    v13 = v2
                                    _ = v21
                                } else {
                                    delete(v3, fi14)
                                    if (95 > 2) {
                                        if (fi14 > fi19) {
                                            v2 = 42
                                            v1 = (v1 / ((v2) | 1))
                                            fmt.Println(v9)
                                            var v24 int16 = 42
                                            fmt.Println(v2, v2)
                                            _ = v24
                                        } else {
                                            var v25 int32 = -1
                                            v13 += v13
                                            _ = v25
                                        }
                                        if (!((!((2147483648 >= 878) && false)) || (false || (298 < 515)))) {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v1 += v13
                                            case "a":
                                                var v26 int32 = -1
                                                _ = v26
                                            case "":
                                                fmt.Println(v9)
                                            }
                                        }
                                        v27 := GenS4{F0: 128, F1: (v9 % (((v9 | (v9 % ((v9) | 1)))) | 1)), F2: 100}
                                        v9 = v9
                                        v28 := []int{(fi19 & (fi19 << ((65535) & 63))), 0}
                                        _ = v27
                                        _ = v28
                                    }
                                    var v29 uint = 42
                                    _ = v29
                                }
                            case "a":
                                fmt.Println(v13, v1, v9)
                            }
                            _ = fi19
                            _ = v20
                        }
                        v2 = 112
                    }
                case "a":
                    fmt.Println(v1, v9, v2)
                case "hello":
                    fmt.Println(v13)
                }
                _ = fi14
            }
        }
        _ = v13
    }
    switch 1 {
    case 1:
        var v30 int = 962
        _ = v30
    default:
        var v31 string = "\n\t"
        _ = v31
    }
    var v32 int16 = 1000
    var v33 uint64 = 995
    v34 := GenS1{F0: 0, F1: 42, F2: -2147483648}
    v1 = (v1 / ((v1) | 1))
    v9 -= v9
    if true {
        v3[-30] = ("hello" + "a")
        var v35 float64 = (float64(3.14159) + 1e10)
        _ = v35
    } else {
        var v36 bool = (!true)
        fmt.Println(v9, v9, v2)
        _ = v36
    }
    var v37 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v9
}
