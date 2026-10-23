package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int; F1 bool; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 int16; F2 uint8 }
type GenS6 struct { F0 string; F1 int64; F2 uint32 }
func genFn0(p0 int8) int8 {
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 42}
        v9 := GenS1{F0: 0}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{4294967295, 4294967295, 127, 9223372036854775807}
    if (-2147483648 == 337) {
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 uint = 18446744073709551615
            fmt.Println(fi11)
            v13 := GenS1{F0: 0}
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            switch (fi14 | fi14) {
            case 42:
                switch "a" {
                case "foo bar":
                    switch 1 {
                    case -1:
                        p0 = p0
                    case 127:
                        var v15 float32 = float32(3.14159)
                        _ = v15
                    }
                case "a":
                    if ((-2.25 * 0.0) <= -2.25) {
                        var v16 string = "hello"
                        var v17 string = v16
                        _ = v16
                        _ = v17
                    }
                default:
                    fmt.Println(p0, fi14)
                }
            case -2147483648:
                var v18 int64 = 928
                _ = v18
            default:
                p0 -= (-(p0))
            }
            _ = fi14
        }
    }
    _ = p0
    _ = v10
    return (p0 / ((p0) | 1))
}

func main() {
    v1 := []uint32{524, 2}
    var v2 int8 = 0
    var v3 int16 = 42
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    v6 := make(map[int]uint32)
    var v19 string = ""
    for fi20 := 0; fi20 < 19; fi20++ {
        _ = len(v1)
        fmt.Println(v2)
        _ = fi20
    }
    if false {
        v6[-2147483648] = 2
    }
    v19 = v19
    if true {
        v21 := GenS2{F0: (42 != 255), F1: v3}
        _ = v21
    }
    v22 := GenS3{F0: 42, F1: (false || (true && true)), F2: 253}
    if false {
        var v23 uint8 = 226
        if true {
            for fi24 := 0; fi24 < 14; fi24++ {
                v25 := GenS1{F0: 42}
                v3 = -32768
                _ = fi24
                _ = v25
            }
        } else {
            fmt.Println(v19, v3)
        }
        _ = v23
    }
    v2 = v2
    if false {
        v26 := GenS3{F0: 128, F1: (!true), F2: 2}
        for fi27 := 0; fi27 < 5; fi27++ {
            if (!true) {
                if (-928 <= int32(fi27)) {
                    v19 = ("a" + ("a" + v19))
                    v28 := GenS4{F0: 19}
                    _ = v28
                }
            } else {
                switch v19 {
                case "foo bar":
                    v29 := GenS3{F0: 32768, F1: (true && false), F2: 0}
                    _ = v29
                case "hello":
                    v3 = v3
                case "\n\t":
                    if (v2 <= v2) {
                        v2 -= v2
                        var v30 uint8 = 0
                        _ = v30
                    }
                default:
                    var v31 uint8 = 199
                    _ = v31
                }
            }
            switch v19 {
            case "a":
                _ = len(v6)
            case "\n\t":
                for fi32 := 0; fi32 < 4; fi32++ {
                    switch "hello" {
                    case "foo bar":
                        if (!false) {
                            var v33 bool = (!(42 == 0))
                            if v33 {
                                switch fi32 {
                                case -2147483648:
                                    v33 = ((0 < 1) || (0 >= 18446744073709551615))
                                default:
                                    var v34 float64 = -2.25
                                    _ = v34
                                }
                            } else {
                                switch fi32 {
                                case 128:
                                    v3 -= v3
                                case 2147483647:
                                    if (!(false && ((v19 + "foo bar") < (v19 + (v19 + "a"))))) {
                                        if ((v19 > (v19 + v19)) || (!v33)) {
                                            var v35 int16 = -32768
                                            v3 = v3
                                            if (0.1 != 1e10) {
                                                fmt.Println(v33, v3, v3)
                                                v3 += (v3 ^ v35)
                                                fmt.Println(v33, v33)
                                            }
                                            _ = v35
                                        }
                                        v33 = v33
                                        fmt.Println(fi32, v33)
                                    } else {
                                        v19 = (v19 + "\n\t")
                                        v2 = int8(fi32)
                                    }
                                case 42:
                                    v36 := GenS5{F0: v33, F1: 1000, F2: 243}
                                    _ = v36
                                }
                                v2 = (v2 & -1)
                            }
                            _ = v33
                        } else {
                            _ = len(v4)
                            var v37 bool = true
                            _ = v37
                        }
                    }
                    _ = fi32
                }
            }
            _ = fi27
        }
        _ = v26
    }
    v2 -= (v2 ^ (v2 % ((-75) | 1)))
    for fi38 := 0; fi38 < 12; fi38++ {
        var v39 int8 = (v2 << ((0) & 63))
        switch v19 {
        case "foo bar":
            if false {
                _ = len(v6)
                for fi40 := 0; fi40 < 15; fi40++ {
                    v19 = "a"
                    _ = fi40
                }
            }
        }
        var v41 int = 127
        _ = fi38
        _ = v39
        _ = v41
    }
    if (0 != 128) {
        var v42 int16 = v3
        _ = v42
    }
    v43 := GenS6{F0: "a", F1: -221, F2: 2147483647}
    v3 = v3
    v1 = append(v1, 307)
    var v44 uint = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
