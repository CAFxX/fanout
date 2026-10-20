package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 uint32 = 2147483648
    var v2 uint16 = 686
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    v5 := make(map[int]uint32)
    v6 := make(map[string]uint32)
    v1 = v1
    v7 := GenS1{F0: "foo bar", F1: -1}
    if true {
        switch "\n\t" {
        case "foo bar":
            if false {
                var v8 int32 = -1
                _ = v8
            }
        case "\n\t":
            fmt.Println(v2)
        case "":
            v9 := GenS1{F0: ("\n\t" + "\n\t"), F1: 42}
            _ = v9
        }
    } else {
        switch ("" + ("hello" + "")) {
        case "":
            _ = len(v3)
        case "a":
            switch "\n\t" {
            case "":
                var v10 uint32 = 920
                _ = v10
            case "\n\t":
                v11 := GenS1{F0: ("foo bar" + "\n\t"), F1: 983}
                _ = v11
            default:
                delete(v5, 32767)
            }
        case "hello":
            v1 = v1
        default:
            switch "" {
            case "foo bar":
                var v12 uint16 = 65535
                _ = v12
            case "a":
                v13 := GenS1{F0: "hello", F1: -9223372036854775808}
                _ = v13
            default:
                for fi14 := 0; fi14 < 15; fi14++ {
                    v2 = v2
                    _ = fi14
                }
            }
        }
        var v15 float64 = 1.5
        _ = v15
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v16 := GenS1{F0: (("\n\t" + "") + ("hello" + "")), F1: 0}
    switch -537 {
    case -816:
        v3 = append(v3, v1)
    case -26:
        v17 := GenS1{F0: (("" + "\n\t") + "a"), F1: 42}
        _ = v17
    default:
        switch -2147483648 {
        case 4294967295:
            switch -736 {
            case 42:
                v2 -= v2
            case 2147483647:
                if (((0.0 - 0.1) < (1.5 * (3.14159 + float64(1.5)))) && true) {
                    v18 := GenS2{F0: (v2 * v2), F1: (v1 ^ (^(v1)))}
                    _ = v18
                }
            case 4294967295:
                if false {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v2 += (v2 * 65535)
                            _ = fi20
                        }
                        v3 = append(v3, (v1 << ((1) & 63)))
                        var v21 int8 = 100
                        _ = fi19
                        _ = v21
                    }
                }
            }
        default:
            delete(v6, "")
        }
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    if false {
        _ = len(v4)
        if true {
            v2 += (v2 ^ 255)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v22 := v6["a"]
            fmt.Println(v22, v1)
            for fi23 := 0; fi23 < 13; fi23++ {
                v22 = (v1 & (v1 / ((v22) | 1)))
                v22 += (v1 % ((v1) | 1))
                switch ("a" + "") {
                case "hello":
                    v2 += 230
                }
                _ = fi23
            }
            _ = v22
        }
        fmt.Println(v1)
    } else {
        fmt.Println(v2, v2)
        v2 = uint16(v1)
    }
    v2 = v2
    v2 = (v2 / ((42) | 1))
    var v24 int16 = -1
    var v25 int16 = -950
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
