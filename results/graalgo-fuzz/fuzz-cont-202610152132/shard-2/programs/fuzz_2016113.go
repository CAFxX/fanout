package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 bool; F1 int64; F2 int8 }
func genFn0() uint {
    var v4 int8 = 100
    var v5 uint64 = 18446744073709551615
    fmt.Println(v4, v5)
    _ = v4
    _ = v5
    return 823
}

func main() {
    v1 := []uint32{65535, 4294967295, 4294967295, 1}
    var v2 int = 32767
    var v3 uint8 = 134
    v2 = -1
    if len(v1) > 0 {
        v1[0] = 886
    }
    switch "" {
    case "a":
        switch ("hello" + "hello") {
        case "":
            fmt.Println(v2, v2)
        case "foo bar":
            fmt.Println(v2, v2)
        case "a":
            for fi6 := 0; fi6 < 4; fi6++ {
                _ = len(v1)
                _ = fi6
            }
        default:
            v2 += (v2 % (((-(v2))) | 1))
        }
    }
    v7 := GenS1{F0: -569, F1: 2}
    if (42 != 32767) {
        v1 = append(v1, 65535)
        fmt.Println(v2, v3)
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(v3)
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 int32 = -974
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    } else {
        v2 += (v2 - (v2 & (v2 / ((42) | 1))))
        v11 := GenS2{F0: true, F1: true}
        _ = v11
    }
    var v12 float64 = 3.14159
    if true {
        if (!(v3 <= v3)) {
            v2 -= 738
            v3 -= 127
        } else {
            var v13 int16 = -345
            _ = v13
        }
    }
    if (!((81 < 104) && ((-661 > 450) || false))) {
        v14 := GenS2{F0: true, F1: true}
        if len(v1) > 0 {
            v1[0] = 42
        }
        v2 += (v2 - v2)
        _ = v14
    } else {
        var v15 float32 = float32(1.5)
        var v16 uint16 = 2
        var v17 int = (v2 + v2)
        _ = v15
        _ = v16
        _ = v17
    }
    switch "" {
    case "hello":
        switch -1 {
        case -1:
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 int32 = 1
                var v20 float32 = 1.5
                v3 -= v3
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    case "":
        if (!(1 >= 32767)) {
            var v21 float32 = 3.14159
            v12 += 0.1
            fmt.Println(v21, v3, v2)
            _ = v21
        } else {
            var v22 uint16 = 2
            switch (v2 % (((v2 & v2)) | 1)) {
            case 128:
                switch "hello" {
                case "\n\t":
                    if false {
                        v2 += 89
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case "foo bar":
                    v3 += v3
                }
            case 9223372036854775807:
                if false {
                    v23 := GenS3{F0: true, F1: 596, F2: 17}
                    _ = v23
                } else {
                    v12 = ((1.5 * 0.0) / 1e10)
                    if true {
                        for fi24 := 0; fi24 < 17; fi24++ {
                            if len(v1) > 0 {
                                v1[0] = 2147483647
                            }
                            v25 := GenS3{F0: (true && true), F1: -9223372036854775808, F2: 101}
                            v22 += 258
                            _ = fi24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 14; fi26++ {
                            for fi27 := 0; fi27 < 11; fi27++ {
                                if (2 < 0) {
                                    if ((false && false) && false) {
                                        v3 = 222
                                    }
                                }
                                v3 = 194
                                _ = fi27
                            }
                            _ = fi26
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
            default:
                switch (v2 & v2) {
                case -792:
                    var v28 float64 = v12
                    _ = v28
                default:
                    if len(v1) > 0 {
                        v1[0] = 52
                    }
                }
            }
            _ = v22
        }
    case "\n\t":
        switch "a" {
        case "":
            v2 -= (v2 + 32768)
        }
    default:
        v12 = 1.5
    }
    v2 = v2
    v29 := GenS2{F0: (!true), F1: true}
    fmt.Println(v2, v2, v12)
    var v30 int32 = 0
    for fi31 := 0; fi31 < 20; fi31++ {
        if ((-52 >= -128) && false) {
            var v32 int32 = (-(v30))
            v33 := GenS3{F0: ((!(0 >= 2)) || (3.14159 <= 0.0)), F1: 9223372036854775807, F2: 127}
            _ = v32
            _ = v33
        }
        _ = fi31
    }
    var v34 uint32 = 0
    var v35 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v7
}
