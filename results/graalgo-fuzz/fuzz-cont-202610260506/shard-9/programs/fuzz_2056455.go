package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int16 }
type GenS2 struct { F0 bool; F1 int64; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 float64 = 3.14159
    var v2 int = 128
    v3 := []uint32{326, 2147483647, 2147483647, 0}
    v2 = -515
    switch ("\n\t" + "") {
    case "hello":
        var v4 uint32 = 0
        _ = v4
    case "\n\t":
        fmt.Println(v1)
    default:
        for fi5 := 0; fi5 < 19; fi5++ {
            v1 = (v1 - v1)
            _ = fi5
        }
    }
    var v6 int = v2
    if ((!true) || false) {
        if true {
            var v7 float32 = -2.25
            _ = len(v3)
            v1 += 3.14159
            _ = v7
        } else {
            for fi8 := 0; fi8 < 2; fi8++ {
                fmt.Println(v6, fi8, v1)
                _ = fi8
            }
        }
        v9 := GenS1{F0: 2, F1: -712, F2: -1}
        v6 -= v6
        _ = v9
    } else {
        switch v2 {
        case 42:
            var v10 uint64 = 1
            _ = v10
        default:
            v11 := GenS2{F0: (825 > 42), F1: 9223372036854775807, F2: 42}
            _ = v11
        }
        var v12 int32 = -23
        if (2 != 2) {
            if false {
                var v13 uint16 = 255
                _ = v13
            } else {
                v14 := GenS3{F0: "hello"}
                _ = v14
            }
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: 4294967295, F1: int16(v2), F2: 0}
                fmt.Println(fi15, fi15, v1)
                var v17 uint16 = 0
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = v12
    }
    v6 = v6
    fmt.Println(v1, v1, v1)
    if (("" >= ("foo bar" + "\n\t")) && true) {
        v18 := GenS4{F0: -1, F1: 1}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 5; fi19++ {
            switch v2 {
            case 9223372036854775807:
                if false {
                    v6 = fi19
                }
            case 227:
                fmt.Println(v2, fi19)
            default:
                if true {
                    v20 := GenS3{F0: "\n\t"}
                    for fi21 := 0; fi21 < 8; fi21++ {
                        fmt.Println(fi21, v2)
                        if (!((831 > 903) || ((76 <= -128) && (42 != 1)))) {
                            v22 := GenS5{F0: 127, F1: 127}
                            var v23 float32 = -2.25
                            v2 = v6
                            _ = v22
                            _ = v23
                        }
                        var v24 float64 = v1
                        _ = fi21
                        _ = v24
                    }
                    var v25 int64 = -9223372036854775808
                    _ = v20
                    _ = v25
                } else {
                    v1 -= 0.0
                    var v26 int16 = 32767
                    var v27 uint32 = 4294967295
                    _ = v26
                    _ = v27
                }
            }
            v2 -= (v2 + (v6 + (v6 % ((128) | 1))))
            _ = fi19
        }
        v2 = (v2 ^ (v2 + 9223372036854775807))
    }
    if (9223372036854775807 >= (-(v2))) {
        for fi28 := 0; fi28 < 18; fi28++ {
            v1 -= (1.5 + v1)
            _ = fi28
        }
        switch 0 {
        case -9223372036854775808:
            v2 += v2
        case 4294967295:
            v29 := GenS2{F0: (1.5 <= float32(((v1 / -2.25) / 0.1))), F1: 954, F2: 1}
            _ = v29
        default:
            v2 = (v6 * (v2 + (v2 & v2)))
        }
    } else {
        var v30 int8 = 119
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
}
