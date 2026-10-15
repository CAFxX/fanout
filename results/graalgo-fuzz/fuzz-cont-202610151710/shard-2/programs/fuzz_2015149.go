package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"foo bar"}
    var v3 int32 = -2147483648
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v6 := make(map[int]int)
    v3 = (v3 / (((v3 % ((v3) | 1))) | 1))
    for fi7 := 0; fi7 < 14; fi7++ {
        switch "hello" {
        case "\n\t":
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 float64 = (0.1 / -2.25)
                _ = fi8
                _ = v9
            }
        default:
            v3 += -901
        }
        _ = fi7
    }
    var v10 uint = 42
    fmt.Println(v10)
    v3 += v3
    switch "\n\t" {
    case "hello":
        if true {
            var v11 string = ""
            _ = v11
        } else {
            v12 := v4[4294967295]
            switch "foo bar" {
            case "foo bar":
                var v13 uint = (v10 >> ((v10) & 63))
                _ = v13
            case "\n\t":
                v14 := GenS1{F0: 42}
                _ = v14
            }
            _ = v12
        }
    case "a":
        var v15 int32 = -2147483648
        _ = v15
    default:
        switch "a" {
        case "":
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 int8 = 115
                if false {
                    v18 := GenS1{F0: (v17 >> ((0) & 63))}
                    fmt.Println(v3)
                    v3 += (v3 << ((18446744073709551615) & 63))
                    _ = v18
                }
                delete(v1, 1)
                _ = fi16
                _ = v17
            }
        default:
            if false {
                fmt.Println(v3, v3)
            }
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        v20 := GenS2{F0: false, F1: v10, F2: (v10 & 4294967295)}
        v3 -= v3
        var v21 int32 = 0
        _ = fi19
        _ = v20
        _ = v21
    }
    if len(v2) > 0 {
        v2[0] = "a"
    }
    var v22 int16 = -1
    if (!false) {
        var v23 int16 = (-(v22))
        var v24 uint32 = 990
        _ = v23
        _ = v24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        if (!true) {
            var v26 uint64 = 62
            fmt.Println(v3)
            var v27 bool = (!((-2.25 + (0.0 - ((float64((0.1 * 1e10)) * (1.5 + 0.0)) * 1e10))) < ((float64(1e10) * (((1.5 - (0.0 - float64((float32(-2.25) - 0.0)))) + 1.5) * 3.14159)) - ((1e10 + 0.1) - (0.1 - float64(((float32(-2.25) * 0.1) / 1.0)))))))
            _ = v26
            _ = v27
        } else {
            var v28 int8 = 74
            var v29 bool = (!(true || false))
            v29 = v29
            _ = v28
            _ = v29
        }
        _ = fi25
    }
    v22 = (-(v22))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
