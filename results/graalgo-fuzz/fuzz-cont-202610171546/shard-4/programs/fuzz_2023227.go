package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int8; F1 uint16; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 uint = 1
    var v2 bool = false
    var v3 int8 = -128
    var v4 float32 = -2.25
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    if v2 {
        switch "hello" {
        case "":
            if ((v3 << (((v1 ^ (v1 << ((v1) & 63)))) & 63)) < (v3 >> (((v1 | (v1 | 775))) & 63))) {
                v7 := GenS1{F0: 365, F1: (v1 & (^(v1)))}
                fmt.Println(v4, v1)
                v3 -= int8(v1)
                _ = v7
            }
        default:
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 float64 = 1.5
                if (870 > 1) {
                    v2 = (!(0 != 127))
                    switch "foo bar" {
                    case "a":
                        for fi10 := 0; fi10 < 5; fi10++ {
                            var v11 int8 = -87
                            delete(v6, fi10)
                            switch -283 {
                            case 0:
                                var v12 uint8 = 17
                                _ = v12
                            }
                            var v13 float64 = 1.5
                            _ = fi10
                            _ = v11
                            _ = v13
                        }
                    }
                }
                var v14 uint8 = 255
                _ = fi8
                _ = v14
                _ = v9
            }
        }
        if (!(42 >= v1)) {
            v15 := GenS1{F0: 916, F1: v1}
            _ = v15
        }
    }
    v1 = (^(v1))
    v16 := GenS2{F0: (v3 >> (((^(v1))) & 63)), F1: 1, F2: (v1 << (((v1 * (v1 - (v1 - (v1 >> ((v1) & 63)))))) & 63))}
    var v17 int32 = 42
    v1 = (v1 % ((v1) | 1))
    if v2 {
        var v18 uint16 = 986
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                v17 -= 1
                _ = fi20
            }
            v21 := GenS3{F0: (v1 + 2)}
            _ = fi19
            _ = v21
        }
        if (!v2) {
            var v22 uint32 = 0
            _ = v22
        } else {
            var v23 float32 = (1.5 - float32(float64((v4 / (float32(((-2.25 - 1.5) * (-2.25 / 1.5))) / (v4 * v4))))))
            switch "foo bar" {
            case "a":
                v24 := GenS1{F0: -32768, F1: (v1 & 4294967295)}
                _ = v24
            default:
                v1 = (v1 + 1)
            }
            if false {
                if v2 {
                    var v25 string = ""
                    switch v25 {
                    case "foo bar":
                        var v26 int16 = -1
                        _ = v26
                    case "hello":
                        if false {
                            delete(v5, ("a" + ("\n\t" + (("a" + v25) + (v25 + ("" + "foo bar"))))))
                            fmt.Println(v18)
                            if ((v1 | (v1 * 678)) != v1) {
                                if (42 == 0) {
                                    var v27 float64 = float64(float32(0.1))
                                    v25 = ("foo bar" + v25)
                                    _ = v27
                                }
                            }
                        } else {
                            if (127 >= 297) {
                                fmt.Println(v17, v25, v18)
                                var v28 int16 = -32768
                                _ = v28
                            } else {
                                fmt.Println(v1)
                                v29 := GenS4{F0: 9223372036854775807, F1: 2}
                                _ = v29
                            }
                            var v30 uint32 = 964
                            if v2 {
                                v1 -= 4294967295
                                var v31 int16 = 955
                                var v32 int64 = 42
                                _ = v31
                                _ = v32
                            } else {
                                var v33 uint = v1
                                fmt.Println(v23, v3, v2)
                                _ = v33
                            }
                            _ = v30
                        }
                    case "\n\t":
                        v34 := []int{-523, 127}
                        _ = v34
                    default:
                        var v35 uint8 = 127
                        _ = v35
                    }
                    v1 += uint(v18)
                    _ = v25
                } else {
                    var v36 uint32 = 2
                    _ = v36
                }
                switch 636 {
                case -614:
                    var v37 int = 934
                    _ = v37
                case 32768:
                    for fi38 := 0; fi38 < 4; fi38++ {
                        fmt.Println(v23, v1, v2)
                        if true {
                            for fi39 := 0; fi39 < 6; fi39++ {
                                v1 = v1
                                v40 := v6[(-(fi38))]
                                _ = fi39
                                _ = v40
                            }
                        }
                        v18 -= (v18 / (((v18 >> ((uint(v18)) & 63))) | 1))
                        _ = fi38
                    }
                }
                v41 := []int{}
                _ = v41
            } else {
                fmt.Println(v1)
                v23 += v4
            }
            _ = v23
        }
        _ = v18
    }
    switch -2147483648 {
    case 42:
        v42 := GenS4{F0: -968, F1: 255}
        _ = v42
    case -9223372036854775808:
        v43 := GenS5{F0: 520}
        _ = v43
    case 128:
        var v44 int16 = 314
        _ = v44
    }
    v45 := []int{-1, -240, 664}
    for fi46 := 0; fi46 < 1; fi46++ {
        var v47 int = fi46
        _ = fi46
        _ = v47
    }
    var v48 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v5
    _ = v6
}
