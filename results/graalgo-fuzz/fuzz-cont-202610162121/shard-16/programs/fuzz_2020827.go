package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int }
type GenS5 struct { F0 string; F1 uint; F2 uint8 }

func main() {
    v1 := []int{}
    var v2 float32 = float32((1e10 / 3.14159))
    var v3 uint16 = 1
    var v4 uint64 = 0
    v5 := []uint32{uint32(v3), 1, 4294967295, 2}
    var v6 int32 = -2147483648
    var v7 bool = true
    fmt.Println(v6, v6)
    v4 += (v4 + 1)
    v2 = v2
    if v7 {
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int = (fi8 / ((fi8) | 1))
            if (255 < 42) {
                var v10 int32 = 42
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 string = ("\n\t" + "hello")
                    var v13 int8 = -128
                    v10 = -2147483648
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                for fi14 := 0; fi14 < 4; fi14++ {
                    v15 := GenS1{F0: 127}
                    var v16 int64 = -1
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v10
            } else {
                for fi17 := 0; fi17 < 16; fi17++ {
                    fmt.Println(v7)
                    var v18 string = "hello"
                    _ = fi17
                    _ = v18
                }
                v6 = (v6 - -1)
            }
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 string = ("a" + "hello")
                v21 := GenS2{F0: int8(v4)}
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi8
            _ = v9
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if v7 {
            if (v4 == 419) {
                v4 = (v4 + 907)
                if false {
                    var v22 float64 = 1e10
                    _ = v22
                } else {
                    v4 -= (v4 & v4)
                }
            }
            for fi23 := 0; fi23 < 14; fi23++ {
                v24 := GenS3{F0: -1, F1: 26}
                switch 42 {
                case 9223372036854775807:
                    v25 := GenS4{F0: uint8(v3), F1: 2147483648, F2: (fi23 << ((2) & 63))}
                    _ = v25
                case -400:
                    v6 -= (v6 * (v6 & v6))
                default:
                    v26 := GenS3{F0: 257, F1: 127}
                    _ = v26
                }
                _ = fi23
                _ = v24
            }
            _ = len(v1)
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v27 float64 = 3.14159
            _ = v27
        }
        v28 := GenS4{F0: 148, F1: 2, F2: int(v4)}
        _ = v28
    }
    if (1 != -9223372036854775808) {
        fmt.Println(v2, v2, v7)
        for fi29 := 0; fi29 < 15; fi29++ {
            switch 32768 {
            case -306:
                var v30 int = 127
                _ = v30
            }
            if true {
                var v31 int = 1
                var v32 int8 = -1
                switch v31 {
                case 128:
                    v33 := GenS3{F0: -9223372036854775808, F1: 107}
                    _ = v33
                case -712:
                    fmt.Println(v31)
                }
                _ = v31
                _ = v32
            } else {
                v34 := GenS5{F0: "hello", F1: 279, F2: 2}
                var v35 uint32 = 581
                switch 920 {
                case 0:
                    if false {
                        switch fi29 {
                        case 42:
                            var v36 bool = (0 == v4)
                            _ = v36
                        case -409:
                            var v37 bool = (false && v7)
                            _ = v37
                        }
                    }
                default:
                    var v38 uint8 = 198
                    _ = v38
                }
                _ = v34
                _ = v35
            }
            _ = fi29
        }
    } else {
        v3 += 259
        v39 := GenS4{F0: 30, F1: 1, F2: -2147483648}
        _ = v39
    }
    switch ("hello" + "hello") {
    case "a":
        v1 = append(v1, 32767)
    case "foo bar":
        for fi40 := 0; fi40 < 11; fi40++ {
            v7 = (0 > 2147483647)
            _ = fi40
        }
    case "":
        var v41 int16 = 283
        _ = v41
    }
    v1 = append(v1, 4294967295)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
