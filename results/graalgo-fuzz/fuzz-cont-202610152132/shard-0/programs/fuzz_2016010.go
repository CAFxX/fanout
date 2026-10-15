package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint64; F2 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{32767, -2147483648}
    var v3 bool = true
    v4 := []int{}
    var v5 uint32 = 4294967295
    v3 = v3
    if (!(!false)) {
        if v3 {
            v3 = ("foo bar" <= "foo bar")
        } else {
            if false {
                for fi6 := 0; fi6 < 5; fi6++ {
                    if false {
                        fmt.Println(fi6, v5)
                        var v7 int16 = 32767
                        _ = v7
                    } else {
                        for fi8 := 0; fi8 < 17; fi8++ {
                            v9 := v1[9223372036854775807]
                            _ = fi8
                            _ = v9
                        }
                        fmt.Println(v3)
                        var v10 int = fi6
                        _ = v10
                    }
                    v5 = v5
                    _ = fi6
                }
                v11 := GenS1{F0: -9223372036854775808, F1: -1}
                v12 := GenS2{F0: v3}
                _ = v11
                _ = v12
            }
            var v13 uint8 = 1
            _ = v13
        }
        fmt.Println(v3)
        if v3 {
            v2 = append(v2, -146)
            var v14 int64 = -9223372036854775808
            var v15 uint32 = v5
            _ = v14
            _ = v15
        } else {
            var v16 bool = (!(!v3))
            switch ("hello" + "a") {
            case "a":
                var v17 uint16 = 255
                _ = v17
            case "hello":
                v5 -= (v5 ^ v5)
            case "\n\t":
                for fi18 := 0; fi18 < 18; fi18++ {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v20 := GenS1{F0: (fi19 + fi18), F1: -546}
                        v3 = (!v16)
                        var v21 float64 = 0.1
                        v5 += v5
                        if false {
                            switch fi19 {
                            case 42:
                                if ((((!(!((2 >= 0) || false))) && v3) && true) || false) {
                                    var v22 float32 = (((float32(((1.5 * -2.25) / (1.5 - (1.5 / (v21 * v21))))) + ((-2.25 - 0.1) + (0.0 / 0.1))) + float32(v21)) * float32(((v21 - v21) * ((-2.25 + v21) * v21))))
                                    fmt.Println(v16, v16, v5)
                                    _ = v22
                                }
                            case 9223372036854775807:
                                var v23 float64 = v21
                                _ = v23
                            case 4294967295:
                                v5 = 911
                            }
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    var v24 float32 = 0.0
                    _ = fi18
                    _ = v24
                }
            default:
                fmt.Println(v16)
            }
            _ = v16
        }
    } else {
        if v3 {
            var v25 uint32 = 2147483648
            v26 := GenS3{F0: 985, F1: ((!v3) && v3)}
            for fi27 := 0; fi27 < 9; fi27++ {
                for fi28 := 0; fi28 < 19; fi28++ {
                    if ((2147483648 <= v25) || v3) {
                        var v29 float64 = (1e10 / 1e10)
                        v3 = ((0 >= -38) && false)
                        v30 := GenS4{F0: 127, F1: 785}
                        _ = v29
                        _ = v30
                    } else {
                        v4 = append(v4, (fi27 | fi27))
                    }
                    if (((1e10 * (3.14159 - 0.0)) - (1.5 * (((float32(0.0) + -2.25) + float32(float64(1e10))) - -2.25))) >= 3.14159) {
                        if v3 {
                            if (!(!false)) {
                                fmt.Println(fi27, fi28)
                            } else {
                                var v31 uint8 = 255
                                var v32 uint32 = v5
                                var v33 uint16 = 42
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                            var v34 int64 = 1
                            _ = v34
                        } else {
                            var v35 string = "hello"
                            var v36 uint = 2
                            var v37 int64 = -128
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                    }
                    v38 := GenS5{F0: -9223372036854775808, F1: 21, F2: 1}
                    v39 := GenS3{F0: fi27, F1: (8 >= v25)}
                    var v40 uint16 = 32767
                    if false {
                        delete(v1, (fi27 * (fi28 | (fi27 / ((-2147483648) | 1)))))
                        var v41 uint32 = 621
                        v42 := GenS6{F0: -1}
                        _ = v41
                        _ = v42
                    }
                    _ = fi28
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = len(v1)
                _ = fi27
            }
            _ = v25
            _ = v26
        } else {
            fmt.Println(v3)
            var v43 uint64 = 0
            var v44 uint64 = 4294967295
            _ = v43
            _ = v44
        }
    }
    v3 = (!(false && v3))
    switch 42 {
    case -617:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 537:
        for fi45 := 0; fi45 < 11; fi45++ {
            v3 = (!v3)
            _ = fi45
        }
    case 32768:
        fmt.Println(v3)
    default:
        v5 = v5
    }
    v5 += v5
    switch ("" + "hello") {
    case "a":
        v46 := GenS7{F0: -774}
        _ = v46
    }
    switch "" {
    case "":
        var v47 uint8 = 42
        _ = v47
    }
    if (0 <= 42) {
        _ = len(v4)
        for fi48 := 0; fi48 < 7; fi48++ {
            var v49 int64 = -1
            _ = fi48
            _ = v49
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
