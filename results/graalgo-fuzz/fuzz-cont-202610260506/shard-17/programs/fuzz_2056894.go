package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8; F1 string; F2 int32 }
func genFn0(p0 uint8, p1 string, p2 int16) uint8 {
    var v7 string = ("a" + "a")
    v8 := []int{76, -1, 615}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 uint16 = 380
    var v2 uint8 = 1
    var v3 uint64 = 901
    v4 := []int{1}
    v5 := []uint32{65535, 2147483647}
    v6 := []int{799, 128, 922, 578}
    if len(v6) > 0 {
        v6[0] = -864
    }
    v1 = (v1 * v1)
    if (false || false) {
        v9 := GenS1{F0: 0, F1: -9223372036854775808, F2: 932}
        v3 = v3
        switch 9223372036854775807 {
        case -2147483648:
            _ = len(v6)
        case 1:
            if ((!true) || ((true || (1 > 32767)) || (3.14159 == 0.1))) {
                var v10 bool = (128 >= 2147483647)
                v5 = append(v5, 65535)
                _ = v10
            } else {
                if (725 < 0) {
                    switch 42 {
                    case 32767:
                        var v11 uint32 = 2147483647
                        _ = v11
                    case 124:
                        v5 = append(v5, 0)
                    }
                }
                v3 = (v3 & (v3 | 18446744073709551615))
            }
        case 32768:
            switch 1 {
            case 2147483647:
                fmt.Println(v2)
            case 32768:
                v6 = append(v6, 2147483647)
            case 4294967295:
                if false {
                    v12 := GenS1{F0: -119, F1: -672, F2: 65535}
                    v1 -= v1
                    _ = v12
                } else {
                    if false {
                        v3 = (v3 * 14)
                        var v13 int = 42
                        _ = v13
                    }
                    v4 = append(v4, 4294967295)
                }
            }
        }
        _ = v9
    } else {
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint = uint(v3)
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 int = fi16
                v15 -= v15
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 14; fi18++ {
                if false {
                    v15 -= v15
                    v4 = append(v4, fi18)
                    v19 := GenS1{F0: 127, F1: -1, F2: 4294967295}
                    _ = v19
                }
                fmt.Println(v15)
                var v20 float32 = float32((1e10 + float64(1e10)))
                var v21 string = "hello"
                _ = fi18
                _ = v20
                _ = v21
            }
            _ = fi14
            _ = v15
        }
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        if len(v5) > 0 {
            v5[0] = 1
        }
        fmt.Println(fi22)
        v23 := GenS1{F0: 0, F1: 1, F2: 42}
        _ = fi22
        _ = v23
    }
    fmt.Println(v2)
    if true {
        switch ("hello" + ((("a" + "") + "") + ("\n\t" + ("foo bar" + "a")))) {
        case "":
            if len(v6) > 0 {
                v6[0] = 0
            }
        case "hello":
            v24 := GenS1{F0: 62, F1: 42, F2: 2147483648}
            _ = v24
        case "\n\t":
            v2 -= v2
        default:
            if ((!true) && (float64(-2.25) > float64(-2.25))) {
                if false {
                    for fi25 := 0; fi25 < 1; fi25++ {
                        if false {
                            var v26 string = "hello"
                            fmt.Println(fi25, v1, v26)
                            _ = v26
                        } else {
                            var v27 float64 = 3.14159
                            _ = v27
                        }
                        var v28 uint32 = 126
                        _ = fi25
                        _ = v28
                    }
                    var v29 uint32 = 1
                    var v30 uint = 192
                    _ = v29
                    _ = v30
                }
            }
        }
    }
    switch ("" + "foo bar") {
    case "hello":
        for fi31 := 0; fi31 < 19; fi31++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if true {
                v32 := GenS1{F0: -78, F1: 1, F2: 2147483647}
                v1 = 255
                _ = v32
            } else {
                if len(v4) > 0 {
                    v4[0] = fi31
                }
            }
            _ = fi31
        }
    case "\n\t":
        v2 -= (v2 << ((1) & 63))
    case "a":
        for fi33 := 0; fi33 < 5; fi33++ {
            switch "" {
            case "foo bar":
                switch "foo bar" {
                case "hello":
                    v1 += (v1 ^ (^(v1)))
                case "foo bar":
                    v1 = (v1 / (((v1 / ((v1) | 1))) | 1))
                }
            case "a":
                v34 := GenS2{F0: -128, F1: "foo bar", F2: 42}
                _ = v34
            }
            fmt.Println(fi33, v1, fi33)
            _ = fi33
        }
    }
    _ = len(v4)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v3 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
