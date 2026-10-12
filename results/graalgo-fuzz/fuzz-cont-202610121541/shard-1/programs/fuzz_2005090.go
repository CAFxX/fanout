package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 42
    v3 := []uint32{411, 1, 657}
    v4 := make(map[uint32]uint32)
    var v5 int16 = -1
    v5 = v5
    var v6 uint = 4294967295
    v3 = append(v3, 2)
    v5 += (v5 - v5)
    v7 := GenS1{F0: v5, F1: (v6 | v1), F2: 38}
    v8 := GenS1{F0: (v5 & -373), F1: (v1 >> ((42) & 63)), F2: 35}
    var v9 float32 = -2.25
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 bool = (v6 < (v1 + 150))
        if v11 {
            var v12 uint8 = 68
            _ = v12
        }
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        _ = fi10
        _ = v11
    }
    if (!(true || true)) {
        v3 = append(v3, uint32(v6))
        v13 := GenS1{F0: 42, F1: v6, F2: 0}
        _ = v13
    }
    if true {
        if (0 == 904) {
            switch "\n\t" {
            case "\n\t":
                for fi14 := 0; fi14 < 6; fi14++ {
                    var v15 float32 = (0.0 - ((v9 - ((3.14159 / float32(1e10)) + 0.0)) * v9))
                    _ = fi14
                    _ = v15
                }
            case "hello":
                v9 = (v9 * (v9 + v9))
            case "foo bar":
                if (-1 < -9223372036854775808) {
                    var v16 uint64 = 4294967295
                    _ = v16
                }
            }
            _ = len(v4)
        } else {
            var v17 uint16 = v2
            _ = v17
        }
        v18 := GenS2{F0: 726, F1: int32(v5)}
        switch ("a" + "") {
        case "":
            var v19 float64 = 1.5
            _ = v19
        case "\n\t":
            v1 += v6
        case "hello":
            var v20 uint16 = (v2 & v2)
            _ = v20
        default:
            var v21 float64 = ((3.14159 * -2.25) / 1.0)
            _ = v21
        }
        _ = v18
    } else {
        switch 1 {
        case 1:
            v22 := GenS1{F0: 1000, F1: 42, F2: 201}
            _ = v22
        default:
            if false {
                fmt.Println(v2)
                v23 := GenS1{F0: v5, F1: v1, F2: 42}
                _ = len(v4)
                _ = v23
            } else {
                var v24 string = "hello"
                _ = v24
            }
        }
        switch 4294967295 {
        case 737:
            v25 := GenS2{F0: (v5 << ((v1) & 63)), F1: -2147483648}
            _ = v25
        case 2147483647:
            switch "a" {
            case "foo bar":
                switch "" {
                case "hello":
                    switch "a" {
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "\n\t":
                        fmt.Println(v6, v6, v2)
                    case "a":
                        v5 -= 0
                    }
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        v26 := GenS1{F0: 32767, F1: v1, F2: 81}
                        _ = v26
                    }
                case "a":
                    if (992 != v2) {
                        var v27 int32 = 2147483647
                        _ = v27
                    }
                }
            case "hello":
                fmt.Println(v1, v1)
            case "":
                if len(v3) > 0 {
                    v3[0] = 2
                }
            }
        case 158:
            v28 := GenS3{F0: 4294967295, F1: -26, F2: 2}
            _ = v28
        }
        switch "hello" {
        case "a":
            v29 := GenS1{F0: 42, F1: v6, F2: 255}
            _ = v29
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            var v30 int = -25
            _ = v30
        default:
            for fi31 := 0; fi31 < 5; fi31++ {
                if false {
                    switch ("foo bar" + "foo bar") {
                    case "":
                        fmt.Println(fi31, v5, v6)
                    case "a":
                        v4[2147483648] = 65535
                    default:
                        for fi32 := 0; fi32 < 10; fi32++ {
                            v33 := GenS1{F0: 1000, F1: (^(v6)), F2: 75}
                            _ = fi32
                            _ = v33
                        }
                    }
                } else {
                    v2 += (v2 << (((v1 % (((v6 / ((42) | 1))) | 1))) & 63))
                    var v34 uint = 958
                    _ = v34
                }
                v4[490] = 2147483647
                _ = fi31
            }
        }
    }
    fmt.Println(v6, v6)
    v35 := GenS4{F0: 57}
    v4[2147483647] = 635
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
