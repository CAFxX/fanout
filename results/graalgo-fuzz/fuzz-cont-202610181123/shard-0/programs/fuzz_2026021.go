package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 int16; F1 uint16; F2 int32 }
type GenS6 struct { F0 int; F1 bool; F2 uint16 }
func genFn0(p0 uint, p1 int8, p2 uint8) int16 {
    p1 = 16
    var v4 uint64 = 153
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32767
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 32767
    v3 := []string{(("\n\t" + "") + (("hello" + "") + "hello"))}
    v5 := GenS1{F0: 4294967295, F1: 710}
    var v6 uint64 = 605
    for fi7 := 0; fi7 < 10; fi7++ {
        if (!(0 == 705)) {
            if true {
                v2 += -2147483648
            }
        } else {
            v2 = fi7
            v8 := v1[2]
            v1[65535] = (1000 <= 1)
            _ = v8
        }
        _ = len(v3)
        _ = fi7
    }
    v3 = append(v3, "a")
    v9 := GenS2{F0: (true && (v2 < 816)), F1: (!true)}
    for fi10 := 0; fi10 < 5; fi10++ {
        switch (^(fi10)) {
        case 32767:
            v11 := GenS3{F0: v6}
            _ = v11
        case -9223372036854775808:
            v12 := v1[2147483647]
            _ = v12
        case 402:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            fmt.Println(v6)
        }
        _ = fi10
    }
    fmt.Println(v2)
    v13 := GenS2{F0: false, F1: true}
    if (false || (!(-2147483648 != 2147483647))) {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int8 = 127
            if false {
                v16 := GenS4{F0: (v15 << ((4294967295) & 63)), F1: ("" + "a")}
                _ = v16
            }
            v17 := GenS5{F0: -32768, F1: 65535, F2: 42}
            _ = fi14
            _ = v15
            _ = v17
        }
    }
    v6 -= (v6 << ((4294967295) & 63))
    v3 = append(v3, "a")
    if true {
        v18 := GenS6{F0: v2, F1: (false && ("\n\t" <= "\n\t")), F2: 0}
        fmt.Println(v2)
        _ = v18
    } else {
        if false {
            switch v2 {
            case -9223372036854775808:
                if false {
                    var v19 bool = false
                    v1[2147483647] = (4294967295 < 2147483647)
                    if (127 >= 73) {
                        var v20 uint32 = 0
                        _ = v20
                    }
                    _ = v19
                } else {
                    v21 := GenS3{F0: v6}
                    _ = v21
                }
            case 376:
                v2 -= (-(v2))
            default:
                v2 += 9223372036854775807
            }
            var v22 uint64 = 227
            _ = len(v1)
            _ = v22
        } else {
            v6 = 4294967295
            switch "a" {
            case "a":
                v2 -= v2
            default:
                switch "\n\t" {
                case "a":
                    var v23 float32 = 0.1
                    _ = v23
                }
            }
        }
        var v24 int64 = 0
        _ = v24
    }
    v25 := GenS4{F0: 0, F1: "\n\t"}
    switch v2 {
    case -381:
        var v26 uint16 = 623
        _ = v26
    default:
        delete(v1, uint32(v6))
    }
    v3 = append(v3, "\n\t")
    v6 = 42
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
