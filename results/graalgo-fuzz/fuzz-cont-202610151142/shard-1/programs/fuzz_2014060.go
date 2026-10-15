package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int64; F1 int8; F2 int8 }

func main() {
    var v1 float64 = 0.1
    var v2 uint64 = 1
    var v3 string = "hello"
    var v4 float32 = 1.5
    var v5 int = 9223372036854775807
    var v6 int32 = 2147483647
    if false {
        var v7 int16 = 42
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 1000, F1: 71}
            if true {
                v2 += v2
                v10 := GenS1{F0: v7, F1: 258}
                _ = v10
            } else {
                var v11 uint = 1
                var v12 string = ""
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := []int{314}
                switch v3 {
                case "":
                    var v15 int8 = -1
                    _ = v15
                case "\n\t":
                    v2 += v2
                case "hello":
                    switch v3 {
                    case "a":
                        var v16 int32 = 2147483647
                        _ = v16
                    default:
                        var v17 int16 = (^(v7))
                        var v18 int8 = -1
                        v19 := GenS1{F0: (v7 & (v7 / ((-32768) | 1))), F1: 807}
                        switch ("foo bar" + v3) {
                        case "\n\t":
                            fmt.Println(v2, v17, v18)
                        case "foo bar":
                            var v20 uint8 = 129
                            _ = v20
                        case "hello":
                            switch fi8 {
                            case 127:
                                v3 = v3
                            default:
                                if ((985 >= 304) && true) {
                                    var v21 int64 = 0
                                    var v22 uint = 56
                                    v2 = 1
                                    _ = v21
                                    _ = v22
                                } else {
                                    v23 := GenS1{F0: -461, F1: 42}
                                    _ = v23
                                }
                            }
                            v6 = int32(v17)
                        default:
                            v24 := GenS2{F0: v7}
                            _ = v24
                        }
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
                var v25 uint64 = 1
                _ = fi13
                _ = v14
                _ = v25
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    } else {
        v1 = v1
        var v26 uint16 = 12
        v27 := GenS3{F0: (v6 >> ((1) & 63))}
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        if (!(false || false)) {
            for fi29 := 0; fi29 < 3; fi29++ {
                v30 := GenS2{F0: 1000}
                var v31 uint32 = 2147483647
                _ = fi29
                _ = v30
                _ = v31
            }
            var v32 int8 = 84
            if (!true) {
                for fi33 := 0; fi33 < 3; fi33++ {
                    switch "a" {
                    case "a":
                        var v34 int = (fi28 % (((^(v5))) | 1))
                        _ = v34
                    }
                    v35 := GenS4{F0: v2, F1: v2}
                    _ = fi33
                    _ = v35
                }
            } else {
                v1 -= v1
                switch "" {
                case "foo bar":
                    v6 = (v6 % (((v6 / ((v6) | 1))) | 1))
                case "":
                    var v36 uint16 = 0
                    _ = v36
                default:
                    v5 = (^(fi28))
                }
                switch v5 {
                case 32768:
                    v37 := GenS2{F0: -760}
                    _ = v37
                default:
                    v38 := []int{(^(fi28)), v5, fi28}
                    _ = v38
                }
            }
            _ = v32
        }
        var v39 uint8 = 2
        if ((!false) || (128 >= v39)) {
            v40 := []int{32767, v5, (v5 + v5), v5}
            var v41 int64 = 341
            _ = v40
            _ = v41
        } else {
            for fi42 := 0; fi42 < 11; fi42++ {
                var v43 uint64 = (^(v2))
                v43 = (v43 >> ((18446744073709551615) & 63))
                _ = fi42
                _ = v43
            }
        }
        _ = fi28
        _ = v39
    }
    if (473 < 65535) {
        for fi44 := 0; fi44 < 6; fi44++ {
            v3 = v3
            _ = fi44
        }
    } else {
        fmt.Println(v6)
    }
    v3 = v3
    var v45 uint8 = 51
    v2 = (^(v2))
    for fi46 := 0; fi46 < 5; fi46++ {
        fmt.Println(v2)
        for fi47 := 0; fi47 < 9; fi47++ {
            var v48 bool = true
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    v4 = (v4 - v4)
    v49 := []int{(v5 >> ((539) & 63)), 42}
    v50 := GenS5{F0: -1, F1: 100, F2: -128}
    if len(v49) > 0 {
        v49[0] = (v5 + (v5 | v5))
    }
    var v51 uint32 = 2147483647
    v4 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
