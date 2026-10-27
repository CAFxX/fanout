package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 string; F2 int8 }
type GenS4 struct { F0 int8; F1 uint8; F2 string }
type GenS5 struct { F0 bool; F1 uint; F2 int16 }
func genFn0() int64 {
    switch 1 {
    case 2147483647:
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 string = "foo bar"
            _ = fi5
            _ = v6
        }
    case -1:
        var v7 uint32 = 42
        _ = v7
    }
    if (!true) {
        v8 := []int{1, 2147483647, 128, -159}
        var v9 int8 = -1
        v9 = -1
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 1, F1: -157, F2: 9223372036854775807}
        if (255 >= 878) {
            var v11 uint16 = 2
            var v12 int16 = 42
            _ = v11
            _ = v12
        }
        _ = v10
    }
    var v13 bool = (-2147483648 != 0)
    _ = v13
    return -744
}
func genFn1() uint {
    var v14 int = 1
    var v15 int16 = -46
    v16 := GenS1{F0: 1, F1: 724, F2: 42}
    _ = v14
    _ = v15
    _ = v16
    return 981
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 bool = true
    var v4 float64 = -2.25
    v17 := []int{961, 2147483647, 42}
    var v18 bool = true
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 uint16 = 1
        _ = fi19
        _ = v20
    }
    if len(v17) > 0 {
        v17[0] = 9223372036854775807
    }
    v21 := GenS2{F0: v18}
    for fi22 := 0; fi22 < 17; fi22++ {
        _ = len(v1)
        _ = fi22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            switch "\n\t" {
            case "a":
                v25 := GenS3{F0: -1, F1: (("hello" + "hello") + "hello"), F2: 1}
                _ = v25
            }
            _ = fi24
        }
        for fi26 := 0; fi26 < 16; fi26++ {
            fmt.Println(v3, v4)
            var v27 int16 = 32767
            fmt.Println(v18)
            _ = fi26
            _ = v27
        }
        var v28 uint32 = 65535
        _ = fi23
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 float32 = -2.25
        var v31 float32 = 1.5
        if v18 {
            for fi32 := 0; fi32 < 20; fi32++ {
                switch fi32 {
                case 4294967295:
                    if (((v31 - float32(1e10)) * v31) >= (v31 - float32(v4))) {
                        fmt.Println(fi32)
                        switch 4294967295 {
                        case -36:
                            switch (^(fi32)) {
                            case -1:
                                var v33 int64 = 567
                                _ = v33
                            case 4294967295:
                                v17 = append(v17, fi29)
                            }
                        case 194:
                            switch 0 {
                            case 42:
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            case 32768:
                                v34 := GenS4{F0: -1, F1: 128, F2: "\n\t"}
                                _ = v34
                            default:
                                fmt.Println(v4)
                            }
                        default:
                            v18 = v18
                        }
                        v35 := GenS1{F0: -1, F1: 42, F2: -748}
                        _ = v35
                    }
                case 0:
                    v36 := GenS3{F0: genFn0(), F1: ("" + "\n\t"), F2: -123}
                    _ = v36
                case 1:
                    v31 = 1e10
                default:
                    var v37 int64 = 953
                    _ = v37
                }
                if v3 {
                    v2[-415] = v18
                    var v38 int8 = 100
                    _ = v38
                }
                if false {
                    if len(v17) > 0 {
                        v17[0] = fi29
                    }
                }
                _ = fi32
            }
            delete(v1, "a")
            v39 := GenS2{F0: (!(true || v18))}
            _ = v39
        } else {
            v18 = true
            if (-1 != -32768) {
                switch fi29 {
                case 263:
                    fmt.Println(fi29, v31, v31)
                case -1:
                    v40 := GenS5{F0: v3, F1: 0, F2: 1}
                    _ = v40
                default:
                    v41 := GenS3{F0: 1, F1: "\n\t", F2: 116}
                    _ = v41
                }
                switch "\n\t" {
                case "\n\t":
                    var v42 uint8 = 1
                    _ = v42
                }
            }
            if ((127 <= int8(fi29)) && v18) {
                if len(v17) > 0 {
                    v17[0] = (fi29 / (((fi29 / ((32767) | 1))) | 1))
                }
            }
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    switch "foo bar" {
    case "":
        for fi43 := 0; fi43 < 9; fi43++ {
            var v44 int = (fi43 + (fi43 % ((fi43) | 1)))
            if v3 {
                if v18 {
                    v18 = true
                    v4 -= (((0.0 / ((v4 - v4) - 1.5)) / v4) + 0.1)
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                }
                v45 := GenS2{F0: (3.14159 == 0.1)}
                _ = v45
            } else {
                var v46 int64 = 42
                _ = v46
            }
            var v47 bool = v3
            _ = fi43
            _ = v44
            _ = v47
        }
    case "a":
        fmt.Println(v18, v3, v4)
    default:
        v4 = v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
